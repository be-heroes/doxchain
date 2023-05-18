package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetClientRegistry set a specific ClientRegistry in the store from its index
func (k Keeper) SetClientRegistry(ctx sdk.Context, ClientRegistry types.ClientRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&ClientRegistry)
	store.Set(types.ClientRegistryKey(
		ClientRegistry.Owner.Creator,
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
	clientRegistry, found := k.GetClientRegistry(ctx, clientRegistration.Id.Creator)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetW3CIdentifier() == clientRegistration.Id.GetW3CIdentifier() {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				break
			}
		}

		clientRegistry.Registrations = append(clientRegistry.Registrations, clientRegistration)

		k.SetClientRegistry(ctx, clientRegistry)
	}
}

func (k Keeper) RemoveClientRegistration(ctx sdk.Context, creator string, name string) {
	clientRegistry, found := k.GetClientRegistry(ctx, creator)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Name == name {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				k.SetClientRegistry(ctx, clientRegistry)

				break
			}
		}
	}
}

func (k Keeper) GetClientRegistration(ctx sdk.Context, creator string, fullyQualifiedDidIdentifier string) types.ClientRegistration {
	clientRegistry, found := k.GetClientRegistry(ctx, creator)

	if found {
		for _, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetW3CIdentifier() == fullyQualifiedDidIdentifier {
				return existingClientRegistration
			}
		}
	}
	
	return types.ClientRegistration{}
}

func (k Keeper) SetClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRelationshipRegistryEntry types.ClientRegistrationRelationshipRegistryEntry) error {
	ownerRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationshipRegistryEntry.Owner.Creator, clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier())

	if &ownerRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid owner")
	}

	destinationRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationshipRegistryEntry.Destination.Creator, clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier())

	if &destinationRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid destination")
	}

	for _, aclEntry := range clientRegistrationRelationshipRegistryEntry.AccessClientList.Entries {
		matchOwner := false;
		matchDestination := false;

		for _, aclEntryOwner := range ownerRegistration.AccessClientList.Entries {
			if aclEntryOwner.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchOwner = true;
			}
		}

		for _, aclEntryDestination := range destinationRegistration.AccessClientList.Entries {
			if aclEntryDestination.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchDestination = true;
			}
		}

		if !matchOwner || !matchDestination {
			return sdkerrors.Wrap(types.AccessClientListError, "Illigal relationship. Acl must match both owner and destination acls")
		}
	}

	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))
	
	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRelationshipRegistryEntry.Owner.Creator,
	))
		
	if b == nil {
		clientRegistrationRelationshipRegistry = types.ClientRegistrationRelationshipRegistry{}
	} else {
		k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)	
	}
	
	for index, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier() == existingEntry.Owner.GetW3CIdentifier() && clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier() == existingEntry.Destination.GetW3CIdentifier() {
			clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships[:index], clientRegistrationRelationshipRegistry.Relationships[index+1:]...)
		}
	}

	clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships, clientRegistrationRelationshipRegistryEntry)

	b = k.cdc.MustMarshal(&clientRegistrationRelationshipRegistry)

	store.Set(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRelationshipRegistryEntry.Owner.Creator,
	), b)

	return nil
}

func (k Keeper) RemoveClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRelationshipRegistryEntry types.ClientRegistrationRelationshipRegistryEntry) error {
	ownerRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationshipRegistryEntry.Owner.Creator, clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier())

	if &ownerRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid owner")
	}

	destinationRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationshipRegistryEntry.Destination.Creator, clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier())

	if &destinationRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid destination")
	}

	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))
	
	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRelationshipRegistryEntry.Owner.Creator,
	))
		
	if b == nil {
		clientRegistrationRelationshipRegistry = types.ClientRegistrationRelationshipRegistry{}
	} else {
		k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)	
	}
	
	for index, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier() == existingEntry.Owner.GetW3CIdentifier() && clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier() == existingEntry.Destination.GetW3CIdentifier() {
			clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships[:index], clientRegistrationRelationshipRegistry.Relationships[index+1:]...)
		}
	}

	b = k.cdc.MustMarshal(&clientRegistrationRelationshipRegistry)

	store.Set(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRelationshipRegistryEntry.Owner.Creator,
	), b)

	return nil
}