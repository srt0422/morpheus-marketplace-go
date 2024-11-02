// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// ProxySessionProviderClaimableBalanceService contains methods and other services
// that help with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProxySessionProviderClaimableBalanceService] method instead.
type ProxySessionProviderClaimableBalanceService struct {
	Options []option.RequestOption
}

// NewProxySessionProviderClaimableBalanceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewProxySessionProviderClaimableBalanceService(opts ...option.RequestOption) (r *ProxySessionProviderClaimableBalanceService) {
	r = &ProxySessionProviderClaimableBalanceService{}
	r.Options = opts
	return
}

// Get provider claimable balance
func (r *ProxySessionProviderClaimableBalanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ClaimableBalance, err error) {
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
