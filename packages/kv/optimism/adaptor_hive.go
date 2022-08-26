package optimism

import (
	"github.com/iotaledger/hive.go/kvstore"
	"github.com/iotaledger/wasp/packages/isc/coreutil"
	"golang.org/x/xerrors"
)

var ErrNotImplementedInHiveOptimisticWrapper = xerrors.New("HiveKVStoreReaderOptimistic does not implement write methods")

// HiveKVStoreReaderOptimistic wraps reading methods of the hive.go KV store into optimistic reader checks
// Panic on any write methods and on state invalidated
type HiveKVStoreReaderOptimistic struct {
	db       kvstore.KVStore
	baseline coreutil.StateBaseline
}

func WrapHiveKVStoreReader(db kvstore.KVStore, baseline coreutil.StateBaseline) kvstore.KVStore {
	return &HiveKVStoreReaderOptimistic{
		db:       db,
		baseline: baseline,
	}
}

func (h *HiveKVStoreReaderOptimistic) mustValidState() {
	if !h.baseline.IsValid() {
		panic(coreutil.ErrorStateInvalidated)
	}
}

func (h *HiveKVStoreReaderOptimistic) WithRealm(realm kvstore.Realm) (kvstore.KVStore, error) {
	return h.db.WithRealm(realm)
}

func (h *HiveKVStoreReaderOptimistic) Realm() kvstore.Realm {
	return h.db.Realm()
}

func (h *HiveKVStoreReaderOptimistic) Iterate(prefix kvstore.KeyPrefix, kvConsumerFunc kvstore.IteratorKeyValueConsumerFunc, direction ...kvstore.IterDirection) error {
	h.mustValidState()
	defer h.mustValidState()

	return h.db.Iterate(prefix, kvConsumerFunc, direction...)
}

func (h *HiveKVStoreReaderOptimistic) IterateKeys(prefix kvstore.KeyPrefix, consumerFunc kvstore.IteratorKeyConsumerFunc, direction ...kvstore.IterDirection) error {
	h.mustValidState()
	defer h.mustValidState()

	return h.db.IterateKeys(prefix, consumerFunc, direction...)
}

func (h *HiveKVStoreReaderOptimistic) Clear() error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Get(key kvstore.Key) (value kvstore.Value, err error) {
	h.mustValidState()
	defer h.mustValidState()

	return h.db.Get(key)
}

func (h *HiveKVStoreReaderOptimistic) Set(key kvstore.Key, value kvstore.Value) error {
	panic(ErrNotImplementedInHiveOptimisticWrapper)
}

func (h *HiveKVStoreReaderOptimistic) Has(key kvstore.Key) (bool, error) {
	h.mustValidState()
	defer h.mustValidState()

	return h.db.Has(key)
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
