package kv

import (
	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/wasp/packages/isc/coreutil"
	"golang.org/x/xerrors"
)

var ErrNotImplementedInHiveOptimisticWrapper = xerrors.New("HiveKVStoreReaderOptimistic does not implement write methods")

// HiveKVStoreReaderOptimistic wraps reading methods of the hive.go V store into optimistic reader checks
type HiveKVStoreReaderOptimistic struct {
	db       kvstore.KVStore
	baseline coreutil.StateBaseline
}

func WrapHiveKVStoreOptimistic(db kvstore.KVStore, baseline coreutil.StateBaseline) kvstore.KVStore {
	return &HiveKVStoreReaderOptimistic{
		db:       db,
		baseline: baseline,
	}
}

func (h *HiveKVStoreReaderOptimistic) WithRealm(realm kvstore.Realm) (kvstore.KVStore, error) {
	return h.db.WithRealm(realm)
}

func (h *HiveKVStoreReaderOptimistic) Realm() kvstore.Realm {
	return h.db.Realm()
}

func (h *HiveKVStoreReaderOptimistic) Iterate(prefix kvstore.KeyPrefix, kvConsumerFunc kvstore.IteratorKeyValueConsumerFunc, direction ...kvstore.IterDirection) error {
	if !h.baseline.IsValid() {
		return coreutil.ErrorStateInvalidated
	}
	err := h.db.Iterate(prefix, kvConsumerFunc, direction...)
	if !h.baseline.IsValid() {
		return coreutil.ErrorStateInvalidated
	}
	return err
}

func (h *HiveKVStoreReaderOptimistic) IterateKeys(prefix kvstore.KeyPrefix, consumerFunc kvstore.IteratorKeyConsumerFunc, direction ...kvstore.IterDirection) error {
	if !h.baseline.IsValid() {
		return coreutil.ErrorStateInvalidated
	}
	err := h.db.IterateKeys(prefix, consumerFunc, direction...)
	if !h.baseline.IsValid() {
		return coreutil.ErrorStateInvalidated
	}
	return err
}

func (h *HiveKVStoreReaderOptimistic) Clear() error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Get(key kvstore.Key) (value kvstore.Value, err error) {
	if !h.baseline.IsValid() {
		return nil, coreutil.ErrorStateInvalidated
	}
	ret, err := h.db.Get(key)
	if !h.baseline.IsValid() {
		return nil, coreutil.ErrorStateInvalidated
	}
	return ret, err
}

func (h *HiveKVStoreReaderOptimistic) Set(key kvstore.Key, value kvstore.Value) error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Has(key kvstore.Key) (bool, error) {
	if !h.baseline.IsValid() {
		return false, coreutil.ErrorStateInvalidated
	}
	ret, err := h.db.Has(key)
	if !h.baseline.IsValid() {
		return false, coreutil.ErrorStateInvalidated
	}
	return ret, err
}

func (h *HiveKVStoreReaderOptimistic) Delete(key kvstore.Key) error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) DeletePrefix(prefix kvstore.KeyPrefix) error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Flush() error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Close() error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Batched() (kvstore.BatchedMutations, error) {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}
