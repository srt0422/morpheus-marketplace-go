// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
)

// ProxySessionService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProxySessionService] method instead.
type ProxySessionService struct {
	Options []option.RequestOption
}

// NewProxySessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProxySessionService(opts ...option.RequestOption) (r *ProxySessionService) {
	r = &ProxySessionService{}
	r.Options = opts
	return
}

// Initiate a proxy session
func (r *ProxySessionService) Initiate(ctx context.Context, body ProxySessionInitiateParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "proxy/sessions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Claim provider balance
func (r *ProxySessionService) ProviderClaim(ctx context.Context, id string, body ProxySessionProviderClaimParams, opts ...option.RequestOption) (res *ClaimableBalance, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("proxy/sessions/%s/providerClaim", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get provider claimable balance
func (r *ProxySessionService) ProviderClaimableBalance(ctx context.Context, id string, opts ...option.RequestOption) (res *ClaimableBalance, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("proxy/sessions/%s/providerClaimableBalance", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClaimableBalance struct {
	// Amount claimable by the provider
	Balance string               `json:"balance,required"`
	JSON    claimableBalanceJSON `json:"-"`
}

// claimableBalanceJSON contains the JSON metadata for the struct
// [ClaimableBalance]
type claimableBalanceJSON struct {
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClaimableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r claimableBalanceJSON) RawJSON() string {
	return r.raw
}

type ProxySessionInitiateParams struct {
	// Model ID to initiate session with
	ModelID param.Field[string] `json:"modelId,required"`
	// Duration of the session
	SessionDuration param.Field[string] `json:"sessionDuration,required"`
}

func (r ProxySessionInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProxySessionProviderClaimParams struct {
	// Claim identifier
	Claim param.Field[string] `json:"claim,required"`
}

func (r ProxySessionProviderClaimParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
