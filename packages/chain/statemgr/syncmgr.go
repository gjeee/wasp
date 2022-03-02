package statemgr

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/iotaledger/wasp/packages/chain/messages"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/trie"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/util"
)

func (sm *stateManager) outputPulled(output *iscp.AliasOutputWithID) bool {
	sm.log.Debugf("outputPulled: output index %v id %v", output.GetStateIndex(), iscp.OID(output.ID()))
	if !sm.syncingBlocks.isSyncing(output.GetStateIndex()) {
		// not interested
		sm.log.Debugf("outputPulled: not interested in output for state index %v", output.GetStateIndex())
		return false
	}
	return sm.syncingBlocks.approveBlockCandidates(output)
}

func (sm *stateManager) stateOutputReceived(output *iscp.AliasOutputWithID, timestamp time.Time) bool {
	sm.log.Debugf("stateOutputReceived: received output index: %v, id: %v, timestamp: %v",
		output.GetStateIndex(), iscp.OID(output.ID()), timestamp)
	if sm.solidState.BlockIndex() > output.GetStateIndex() {
		sm.log.Warnf("stateOutputReceived: out of order state output: state manager is already at state %v", sm.solidState.BlockIndex())
		return false
	}
	if sm.stateOutput != nil {
		switch {
		case sm.stateOutput.GetStateIndex() == output.GetStateIndex():
			if sm.stateOutput.ID() == output.ID() {
				// it is just a duplicate
				sm.log.Debugf("stateOutputReceived ignoring: repeated state output")
				return false
			}
			/*if !output.GetIsGovernanceUpdated() {
				sm.log.Panicf("L1 inconsistency: governance transition expected in %s", iscp.OID(output.ID()))
			}*/
			// it is a state controller address rotation

		case sm.stateOutput.GetStateIndex() > output.GetStateIndex():
			sm.log.Warnf("stateOutputReceived: out of order state output: stateOutput index is already larger: %v", sm.stateOutput.GetStateIndex())
			return false
		}
	}
	sm.stateOutput = output
	sm.stateOutputTimestamp = timestamp
	sm.log.Debugf("stateOutputReceived: stateOutput set to index %v id %v timestamp %v", output.GetStateIndex(), iscp.OID(output.ID()), timestamp)
	sm.syncingBlocks.approveBlockCandidates(output)
	return true
}

func (sm *stateManager) doSyncActionIfNeeded() {
	if sm.stateOutput == nil {
		sm.log.Debugf("doSyncAction not needed: stateOutput is nil")
		return
	}
	switch {
	case sm.solidState.BlockIndex() == sm.stateOutput.GetStateIndex():
		sm.log.Debugf("doSyncAction not needed: state is already synced at index #%d", sm.stateOutput.GetStateIndex())
		if sm.domain.HaveMainPeers() {
			sm.domain.SetFallbackMode(false)
		}
		return
	case sm.solidState.BlockIndex() > sm.stateOutput.GetStateIndex():
		sm.log.Debugf("BlockIndex=%v, StateIndex=%v", sm.solidState.BlockIndex(), sm.stateOutput.GetStateIndex())
		sm.log.Panicf("doSyncAction inconsistency: solid state index is larger than state output index")
	}
	// not synced
	startSyncFromIndex := sm.solidState.BlockIndex() + 1
	sm.log.Debugf("doSyncAction: trying to sync state from index %v to %v", startSyncFromIndex, sm.stateOutput.GetStateIndex())
	if !sm.domain.HaveMainPeers() || sm.syncingBlocks.blockPollFallbackNeeded() {
		sm.domain.SetFallbackMode(true)
	}
	for i := startSyncFromIndex; i <= sm.stateOutput.GetStateIndex(); i++ {
		requestBlockRetryTime := sm.syncingBlocks.getRequestBlockRetryTime(i)
		blockCandidatesCount := sm.syncingBlocks.getBlockCandidatesCount(i)
		if blockCandidatesCount == 0 {
			if sm.candidateBlockInWAL(i) {
				blockCandidatesCount++
				sm.syncingBlocks.setReceivedFromWAL(i)
			}
		}
		approvedBlockCandidatesCount := sm.syncingBlocks.getApprovedBlockCandidatesCount(i)
		sm.log.Debugf("doSyncAction: trying to sync state for index %v; requestBlockRetryTime %v, blockCandidates count %v, approved blockCandidates count %v",
			i, requestBlockRetryTime, blockCandidatesCount, approvedBlockCandidatesCount)
		// TODO: temporary if. We need to find a solution to synchronize over large gaps. Making state snapshots may help.
		if i > startSyncFromIndex+maxBlocksToCommitConst {
			sm.chain.EnqueueDismissChain(fmt.Sprintf("StateManager.doSyncActionIfNeeded: too many blocks to catch up: %v", sm.stateOutput.GetStateIndex()-startSyncFromIndex+1))
			return
		}
		currentTime := time.Now()
		if !sm.syncingBlocks.isObtainedFromWAL(i) && currentTime.After(requestBlockRetryTime) {
			// have to pull
			sm.log.Debugf("doSyncAction: requesting block index %v, fallback=%v from %v random peers.", i, sm.domain.GetFallbackMode(), numberOfNodesToRequestBlockFromConst)
			getBlockMsg := &messages.GetBlockMsg{BlockIndex: i}
			for _, p := range sm.domain.GetRandomOtherPeers(numberOfNodesToRequestBlockFromConst) {
				sm.domain.SendMsgByPubKey(p, peering.PeerMessageReceiverStateManager, peerMsgTypeGetBlock, util.MustBytes(getBlockMsg))
				sm.syncingBlocks.blocksPulled()
				sm.log.Debugf("doSyncAction: requesting block index %v,from %v", i, p.AsString())
			}
			sm.syncingBlocks.startSyncingIfNeeded(i)
			sm.syncingBlocks.setRequestBlockRetryTime(i, currentTime.Add(sm.timers.GetBlockRetry))
			if blockCandidatesCount == 0 {
				return
			}
		}
		if approvedBlockCandidatesCount > 0 {
			sm.log.Debugf("doSyncAction: trying to find candidates to commit from index %v to %v", startSyncFromIndex, i)
			candidates, tentativeState, ok := sm.getCandidatesToCommit(make([]*candidateBlock, 0, i-startSyncFromIndex+1), sm.solidState.Copy(), startSyncFromIndex, i)
			if ok {
				sm.log.Debugf("doSyncAction: candidates to commit found, committing")
				sm.commitCandidates(candidates, tentativeState)
				sm.log.Debugf("doSyncAction: blocks from index %v to %v committed", startSyncFromIndex, i)
				return
			}
		}
	}
}

func (sm *stateManager) candidateBlockInWAL(i uint32) bool {
	if !sm.wal.Contains(i) {
		sm.log.Debugf("candidateBlockInWAL: block with index %d not found in wal.", i)
		return false
	}
	blockBytes, err := sm.wal.Read(i)
	if err != nil {
		sm.log.Debugf("candidateBlockInWAL: error reading block bytes for %d. %v", i, err)
		return false
	}
	block, err := state.BlockFromBytes(blockBytes)
	if err != nil {
		sm.log.Debugf("candidateBlockInWAL: error reading block bytes for %d. %v", i, err)
		return false
	}
	nextState := sm.solidState.Copy()
	err = nextState.ApplyBlock(block)
	if err != nil {
		sm.log.Debugf("candidateBlockInWAL: error applying block %d. %v", i, err)
		return false
	}
	_, candidate := sm.syncingBlocks.addBlockCandidate(block, nextState)
	if candidate == nil {
		return false
	}
	candidate.approveIfRightOutput(sm.stateOutput)
	return true
}

func (sm *stateManager) getCandidatesToCommit(candidateAcc []*candidateBlock, calculatedPrevState state.VirtualStateAccess, fromStateIndex, toStateIndex uint32) ([]*candidateBlock, state.VirtualStateAccess, bool) {
	sm.log.Debugf("getCandidatesToCommit from %v to %v", fromStateIndex, toStateIndex)
	if fromStateIndex > toStateIndex {
		// state hashes must be equal
		finalStateCommitment := trie.RootCommitment(calculatedPrevState.TrieAccess())
		finalCandidateCommitment := candidateAcc[len(candidateAcc)-1].getNextStateCommitment()
		if trie.EqualCommitments(finalStateCommitment, finalCandidateCommitment) {
			sm.log.Debugf("getCandidatesToCommit from %v to %v: tentative state obtained, however its hash does not match last candidate expected hash: %v != %v",
				fromStateIndex, toStateIndex, finalStateCommitment.String(), finalCandidateCommitment.String())
			return nil, nil, false
		}
		sm.log.Debugf("getCandidatesToCommit from %v to %v: tentative state obtained, its hash matches last candidate expected hash: %v",
			fromStateIndex, toStateIndex, finalStateCommitment.String())
		return candidateAcc, calculatedPrevState, true
	}

	var stateCandidateBlocks []*candidateBlock
	if fromStateIndex == toStateIndex {
		stateCandidateBlocks = sm.syncingBlocks.getApprovedBlockCandidates(fromStateIndex)
	} else {
		stateCandidateBlocks = sm.syncingBlocks.getBlockCandidates(fromStateIndex)
	}
	sort.Slice(stateCandidateBlocks, func(i, j int) bool {
		return stateCandidateBlocks[i].getVotes() > stateCandidateBlocks[j].getVotes()
	})

	for i, stateCandidateBlock := range stateCandidateBlocks {
		sm.log.Debugf("getCandidatesToCommit from %v to %v: checking block %v of %v", fromStateIndex, toStateIndex, i+1, len(stateCandidateBlocks))
		candidatePrevStateCommitment := stateCandidateBlock.getBlock().PreviousStateCommitment(state.CommitmentModel)
		calculatedPrevStateCommitment := trie.RootCommitment(calculatedPrevState.TrieAccess())
		if trie.EqualCommitments(candidatePrevStateCommitment, calculatedPrevStateCommitment) {
			sm.log.Errorf("getCandidatesToCommit from %v to %v: candidate previous state hash does not match calculated state hash: %v <> %v",
				fromStateIndex, toStateIndex, candidatePrevStateCommitment.String(), calculatedPrevStateCommitment.String())
			return nil, nil, false
		}
		calculatedState, err := stateCandidateBlock.getNextState(calculatedPrevState)
		if err != nil {
			sm.log.Errorf("getCandidatesToCommit from %v to %v: failed to apply synced block index #%d: %v",
				fromStateIndex, toStateIndex, stateCandidateBlock.getBlock().BlockIndex(), err)
			return nil, nil, false
		}
		resultBlocks, tentativeState, ok := sm.getCandidatesToCommit(append(candidateAcc, stateCandidateBlock), calculatedState, fromStateIndex+1, toStateIndex)
		if ok {
			return resultBlocks, tentativeState, true
		}
	}
	return nil, nil, false
}

func (sm *stateManager) commitCandidates(candidates []*candidateBlock, tentativeState state.VirtualStateAccess) {
	blocks := make([]state.Block, len(candidates))
	for i, candidate := range candidates {
		block := candidate.getBlock()
		blocks[i] = block
		sm.syncingBlocks.deleteSyncingBlock(block.BlockIndex())
	}
	from := blocks[0].BlockIndex()
	to := blocks[len(blocks)-1].BlockIndex()
	sm.log.Debugf("commitCandidates: syncing of state indexes from %v to %v is stopped", from, to)

	// TODO: maybe commit in 10 (or some const) block batches?
	//      This would save from large commits and huge memory usage to store blocks

	// invalidate solid state.
	// - If any VM task is running with the assumption of the previous state, it is obsolete and will self-cancel
	// - any view call will return 'state invalidated message'
	sm.chain.GlobalStateSync().InvalidateSolidIndex()
	err := tentativeState.Save(blocks...)
	for _, block := range blocks {
		sm.stateManagerMetrics.RecordBlockSize(block.BlockIndex(), float64(len(block.Bytes())))
	}
	sm.chain.GlobalStateSync().SetSolidIndex(tentativeState.BlockIndex())

	if err != nil {
		sm.log.Errorf("commitCandidates: failed to commit synced changes into DB. Restart syncing. %w", err)
		if strings.Contains(err.Error(), "space left on device") {
			sm.log.Panicf("Terminating WASP, no space left on disc.")
		}
		sm.syncingBlocks.restartSyncing()
		return
	}
	sm.solidState = tentativeState

	sm.log.Debugf("commitCandidates: committing of block indices from %v to %v was successful", from, to)
}
