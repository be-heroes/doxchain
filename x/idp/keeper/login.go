package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	"github.com/be-heroes/doxchain/x/idp/types"
)

func (k Keeper) Login(ctx sdk.Context, msg types.MsgAuthenticationRequest) (types.MsgAuthenticationResponse, error) {
	response := types.MsgAuthenticationResponse{}
	isAuthorized, err := k.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.Tenant, msg.Creator, msg.Creator, time.Minute*60)
	tokenString, err := jwtToken.SignedString([]byte(msg.Creator))

	if err != nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not issue refresh token")
	}

	response.Token = tokenString

	return response, nil
}

func (k Keeper) AuthorizeCreator(ctx sdk.Context, fullyQualifiedW3CIdentifier string, creator string) (bool, error) {
	acl, err := k.GetAccessClientList(ctx, fullyQualifiedW3CIdentifier)

	if err != nil {
		return false, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.User.Creator == creator {
			return true, nil
		}
	}

	return false, nil
}
