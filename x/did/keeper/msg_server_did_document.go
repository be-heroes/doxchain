package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocumentRequest) (result *types.MsgCreateDidDocumentResponse, err error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, types.ErrImpersonation
	}

	if msg.DidDocument.Id.IsModuleIdentifier() {
		return nil, types.ErrImpersonation
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	found := k.Keeper.HasDidDocument(ctx, msg.DidDocument.Id.GetW3CIdentifier())

	if found {
		return nil, types.ErrDidDocumentExists
	}

	k.Keeper.SetDidDocument(ctx, msg.DidDocument, false)

	result.DidDocumentW3CIdentifier = msg.DidDocument.Id.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocumentRequest) (result *types.MsgUpdateDidDocumentResponse, err error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, types.ErrImpersonation
	}

	if msg.DidDocument.Id.IsModuleIdentifier() {
		return nil, types.ErrImpersonation
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDidDocument(ctx, msg.DidDocument.Id.GetW3CIdentifier())

	if found && msg.Creator != match.Id.Creator {
		return nil, types.ErrImpersonation
	}

	k.Keeper.SetDidDocument(ctx, msg.DidDocument, true)

	result.DidDocumentW3CIdentifier = msg.DidDocument.Id.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (result *types.MsgDeleteDidDocumentResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	if !found {
		return nil, types.ErrDidDocumentNotFound
	}

	if msg.Creator != match.Id.Creator {
		return nil, types.ErrImpersonation
	}

	k.Keeper.RemoveDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	return result, nil
}
