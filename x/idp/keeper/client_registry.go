package keeper

import (
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetClientRegistry set a specific ClientRegistry in the store from its index
func (k Keeper) SetClientRegistry(ctx sdk.Context, ClientRegistry types.ClientRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&ClientRegistry)
	store.Set(types.ClientRegistryKey(
		ClientRegistry.Creator,
	), b)
}

// GetClientRegistry returns a ClientRegistry from its index
func (k Keeper) GetClientRegistry(
	ctx sdk.Context,
	creator string,

) (val types.ClientRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))

	b := store.Get(types.ClientRegistryKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClientRegistry removes a ClientRegistry from the store
func (k Keeper) RemoveClientRegistry(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	store.Delete(types.ClientRegistryKey(
		creator,
	))
}

// GetAllClientRegistry returns all ClientRegistry
func (k Keeper) GetAllClientRegistry(ctx sdk.Context) (list []types.ClientRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetClientRegistration(ctx sdk.Context, clientRegistration types.ClientRegistration) {
	clientRegistry, _ := k.GetClientRegistry(ctx, clientRegistration.Id.Creator)

	for index, existingClientRegistration := range clientRegistry.Registrations {
		if existingClientRegistration.Id.GetFullyQualifiedDidIdentifier() == clientRegistration.Id.GetFullyQualifiedDidIdentifier() {
			clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

			break
		}
	}

	clientRegistry.Registrations = append(clientRegistry.Registrations, &clientRegistration)

	k.SetClientRegistry(ctx, clientRegistry)
}

func (k Keeper) RemoveClientRegistration(ctx sdk.Context, creator string, name string) {
	clientRegistry, _ := k.GetClientRegistry(ctx, creator)

	for index, existingClientRegistration := range clientRegistry.Registrations {
		if existingClientRegistration.Name == name {
			clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

			k.SetClientRegistry(ctx, clientRegistry)
			break
		}
	}
}
