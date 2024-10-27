// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
	"github.com/stainless-sdks/morpheus-marketplace-go/shared"
)

// BlockchainProviderBidActiveService contains methods and other services that help
// with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainProviderBidActiveService] method instead.
type BlockchainProviderBidActiveService struct {
	Options []option.RequestOption
}

// NewBlockchainProviderBidActiveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBlockchainProviderBidActiveService(opts ...option.RequestOption) (r *BlockchainProviderBidActiveService) {
	r = &BlockchainProviderBidActiveService{}
	r.Options = opts
	return
}

// Retrieves active bids associated with a specific provider.
func (r *BlockchainProviderBidActiveService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.Bid, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/providers/%s/bids/active", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BlockchainProviderBidActiveListResponse struct {
	BidID          string                                      `json:"bidID"`
	ModelID        string                                      `json:"modelID"`
	PricePerSecond string                                      `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                                      `json:"providerID"`
	Status         string                                      `json:"status"`
	JSON           blockchainProviderBidActiveListResponseJSON `json:"-"`
}

// blockchainProviderBidActiveListResponseJSON contains the JSON metadata for the
// struct [BlockchainProviderBidActiveListResponse]
type blockchainProviderBidActiveListResponseJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainProviderBidActiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainProviderBidActiveListResponseJSON) RawJSON() string {
	return r.raw
}
