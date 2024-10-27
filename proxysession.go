// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/param"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
	"github.com/stainless-sdks/morpheus-marketplace-go/shared"
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

// Initializes a new session with the provider for usage.
func (r *ProxySessionService) Initiate(ctx context.Context, body ProxySessionInitiateParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "proxy/sessions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows a provider to claim their balance for a specific session.
func (r *ProxySessionService) ProviderClaim(ctx context.Context, id string, body ProxySessionProviderClaimParams, opts ...option.RequestOption) (res *ProxySessionProviderClaimResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("proxy/sessions/%s/providerClaim", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the claimable balance for a provider from a session.
func (r *ProxySessionService) ProviderClaimableBalance(ctx context.Context, id string, opts ...option.RequestOption) (res *ProxySessionProviderClaimableBalanceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("proxy/sessions/%s/providerClaimableBalance", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProxySessionProviderClaimResponse struct {
	// Transaction hash.
	Tx   string                                `json:"tx"`
	JSON proxySessionProviderClaimResponseJSON `json:"-"`
}

// proxySessionProviderClaimResponseJSON contains the JSON metadata for the struct
// [ProxySessionProviderClaimResponse]
type proxySessionProviderClaimResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxySessionProviderClaimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxySessionProviderClaimResponseJSON) RawJSON() string {
	return r.raw
}

type ProxySessionProviderClaimableBalanceResponse struct {
	// Claimable balance for the provider.
	ClaimableBalance string                                           `json:"claimableBalance" format:"biginteger"`
	JSON             proxySessionProviderClaimableBalanceResponseJSON `json:"-"`
}

// proxySessionProviderClaimableBalanceResponseJSON contains the JSON metadata for
// the struct [ProxySessionProviderClaimableBalanceResponse]
type proxySessionProviderClaimableBalanceResponseJSON struct {
	ClaimableBalance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProxySessionProviderClaimableBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxySessionProviderClaimableBalanceResponseJSON) RawJSON() string {
	return r.raw
}

type ProxySessionInitiateParams struct {
	// Model ID for the session.
	ModelID param.Field[string] `json:"modelId,required"`
	// Duration for which the session will remain active.
	SessionDuration param.Field[string] `json:"sessionDuration,required" format:"biginteger"`
}

func (r ProxySessionInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProxySessionProviderClaimParams struct {
	// Amount to claim.
	Claim param.Field[string] `json:"claim,required" format:"biginteger"`
}

func (r ProxySessionProviderClaimParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
