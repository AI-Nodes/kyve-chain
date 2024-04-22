package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueueState returns a queue state object based on the identifier as
// there are multiple queues present in the stakers module
func (k Keeper) GetQueueState(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER) (state types.QueueState) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	b := store.Get(identifier)

	if b == nil {
		return state
	}

	k.cdc.MustUnmarshal(b, &state)
	return
}

// SetQueueState sets a endBlocker queue state based on the identifier.
// The identifier is used to distinguish between different queues.
func (k Keeper) SetQueueState(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER, state types.QueueState) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	b := k.cdc.MustMarshal(&state)
	store.Set(identifier, b)
}
