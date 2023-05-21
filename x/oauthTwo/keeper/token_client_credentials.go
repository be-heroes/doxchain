package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgTokenRequest) (response types.MsgTokenResponse, err error) {
	switch msg.ClientAssertionType {
	case "urn:ietf:params:oauth:client-assertion-type:jwt-bearer":
		//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
		return response, sdkerrors.Wrap(types.TokenServiceError, "Assertion is not supported: urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	default:
		jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.TenantW3CIdentifier, msg.Creator, msg.ClientRegistrationAppIdW3CIdentifier, time.Minute*3)
		claims := jwtToken.Claims.(jwt.MapClaims)
		signedToken, err := jwtToken.SignedString([]byte(msg.ClientSecret))

		if err != nil {
			return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
		}

		response.AccessToken = signedToken
		response.TokenType = types.Bearer.String()
		response.ExpiresIn = claims["exp"].(int64)

		tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, msg.TenantW3CIdentifier)

		if !found {
			return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
		}

		tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenRegistryEntry{
			Owner: *didUtils.NewDidTokenFactory().Create(msg.Creator, didUtils.GetWellFormedDidUrl(types.ModuleName, fmt.Sprintf("%T", msg), msg.Creator)),
			Jti: claims["jti"].(string),
			ExpiresAt: response.ExpiresIn,
		})

		k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)
	}

	return response, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
