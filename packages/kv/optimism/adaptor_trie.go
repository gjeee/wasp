package optimism

import (
	"github.com/iotaledger/trie.go/trie"
	"github.com/iotaledger/wasp/packages/isc/coreutil"
)

type TrieKVReaderOptimisticWrapper struct {
	reader   trie.KVReader
	baseline coreutil.StateBaseline
}

func WrapTrieKVReader(reader trie.KVReader, baseline coreutil.StateBaseline) trie.KVReader {
	return &TrieKVReaderOptimisticWrapper{
		reader:   reader,
		baseline: baseline,
	}
}

func (r *TrieKVReaderOptimisticWrapper) mustValidState() {
	if !r.baseline.IsValid() {
		panic(coreutil.ErrorStateInvalidated)
	}
}

func (r *TrieKVReaderOptimisticWrapper) Get(key []byte) []byte {
	r.mustValidState()
	defer r.mustValidState()

	return r.reader.Get(key)
}

func (r *TrieKVReaderOptimisticWrapper) Has(key []byte) bool {
	r.mustValidState()
	defer r.mustValidState()

	return r.reader.Has(key)
}
