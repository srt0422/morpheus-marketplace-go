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
)

// BlockchainProviderService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainProviderService] method instead.
type BlockchainProviderService struct {
	Options []option.RequestOption
	Bids    *BlockchainProviderBidService
}

// NewBlockchainProviderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainProviderService(opts ...option.RequestOption) (r *BlockchainProviderService) {
	r = &BlockchainProviderService{}
	r.Options = opts
	r.Bids = NewBlockchainProviderBidService(opts...)
	return
}

// Create a new provider
func (r *BlockchainProviderService) New(ctx context.Context, body BlockchainProviderNewParams, opts ...option.RequestOption) (res *Provider, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List providers
func (r *BlockchainProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Provider, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a provider
func (r *BlockchainProviderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Provider struct {
	// Unique identifier of the provider
	ID string `json:"id,required"`
	// Endpoint URL of the provider
	Endpoint string `json:"endpoint,required"`
	// Amount staked by the provider
	Stake string       `json:"stake,required"`
	JSON  providerJSON `json:"-"`
}

// providerJSON contains the JSON metadata for the struct [Provider]
type providerJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Stake       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Provider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerJSON) RawJSON() string {
	return r.raw
}

type BlockchainProviderNewParams struct {
	// Endpoint URL of the provider
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Amount to stake for the provider
	Stake param.Field[string] `json:"stake,required"`
}

func (r BlockchainProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
