package v1

import (
	"github.com/artela-network/artela-rollkit/x/aspect/store"
	v0 "github.com/artela-network/artela-rollkit/x/aspect/store/v0"
	"github.com/artela-network/artela-rollkit/x/aspect/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

var _ store.AspectStateStore = (*stateStore)(nil)

// stateStore is the version 1 state Store
type stateStore struct {
	BaseStore

	ctx *types.AspectStoreContext
}

// NewStateStore creates a new instance of account state.
func NewStateStore(ctx *types.AspectStoreContext) store.AspectStateStore {
	// for state Store, we have already charged gas in host api,
	// so no need to charge it again in the Store
	store := runtime.KVStoreAdapter(ctx.StoreService().OpenKVStore(ctx.CosmosContext()))

	return &stateStore{
		BaseStore: NewBaseStore(ctx.CosmosContext().Logger(), v0.NewNoOpGasMeter(ctx), store),
		ctx:       ctx,
	}
}

// SetState sets the state of the aspect with the given ID and key.
func (s *stateStore) SetState(key []byte, value []byte) {
	aspectID := s.ctx.AspectID
	stateKey := store.NewKeyBuilder(V1AspectStateKeyPrefix).AppendBytes(aspectID.Bytes()).AppendBytes(key).Build()
	// no need to check error here, since we are using noop gas meter in state store
	_ = s.Store(stateKey, value)
	s.ctx.Logger().Debug("set aspect state", "key", string(key), "value", abbreviateHex(value))
	return
}

// GetState returns the state of the aspect with the given ID and key.
func (s *stateStore) GetState(key []byte) []byte {
	aspectID := s.ctx.AspectID
	stateKey := store.NewKeyBuilder(V1AspectStateKeyPrefix).AppendBytes(aspectID.Bytes()).AppendBytes(key).Build()
	// no need to check error here, since we are using noop gas meter in state store
	data, _ := s.Load(stateKey)
	s.ctx.Logger().Debug("get aspect state", "key", string(key), "value", abbreviateHex(data))
	return data
}
