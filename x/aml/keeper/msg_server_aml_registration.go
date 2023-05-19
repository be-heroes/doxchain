package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAMLRegistration(goCtx context.Context, msg *types.MsgCreateAMLRegistrationRequest) (*types.MsgCreateAMLRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, isFound := k.GetAMLRegistration(ctx, msg.Creator)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}
	
	if msg.Creator != msg.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	k.SetAMLRegistration(
		ctx,
		types.AMLRegistration{
			Owner:      msg.Owner,
			Approved: false,
		},
	)

	return &types.MsgCreateAMLRegistrationResponse{}, nil
}

func (k msgServer) DeleteAMLRegistration(goCtx context.Context, msg *types.MsgDeleteAMLRegistrationRequest) (*types.MsgDeleteAMLRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetAMLRegistration(ctx, msg.Creator)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation is not supported")
	}

	k.RemoveAMLRegistration(ctx, msg.Creator)

	return &types.MsgDeleteAMLRegistrationResponse{}, nil
}