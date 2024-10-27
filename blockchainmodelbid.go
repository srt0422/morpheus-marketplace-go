// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
)

// BlockchainModelBidService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelBidService] method instead.
type BlockchainModelBidService struct {
	Options []option.RequestOption
}

// NewBlockchainModelBidService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainModelBidService(opts ...option.RequestOption) (r *BlockchainModelBidService) {
	r = &BlockchainModelBidService{}
	r.Options = opts
	return
}

// Retrieves a list of bids associated with a specific model.
func (r *BlockchainModelBidService) List(ctx context.Context, id string, query BlockchainModelBidListParams, opts ...option.RequestOption) (res *shared.Bid, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/bids", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetches active bids associated with a specific model.
func (r *BlockchainModelBidService) Active(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.Bid, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/bids/active", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves rated bids for a specified model.
func (r *BlockchainModelBidService) Rated(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.Bid, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/bids/rated", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BlockchainModelBidListResponse struct {
	BidID          string                             `json:"bidID"`
	ModelID        string                             `json:"modelID"`
	PricePerSecond string                             `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                             `json:"providerID"`
	Status         string                             `json:"status"`
	JSON           blockchainModelBidListResponseJSON `json:"-"`
}

// blockchainModelBidListResponseJSON contains the JSON metadata for the struct
// [BlockchainModelBidListResponse]
type blockchainModelBidListResponseJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainModelBidListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelBidListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelBidActiveResponse struct {
	BidID          string                               `json:"bidID"`
	ModelID        string                               `json:"modelID"`
	PricePerSecond string                               `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                               `json:"providerID"`
	Status         string                               `json:"status"`
	JSON           blockchainModelBidActiveResponseJSON `json:"-"`
}

// blockchainModelBidActiveResponseJSON contains the JSON metadata for the struct
// [BlockchainModelBidActiveResponse]
type blockchainModelBidActiveResponseJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainModelBidActiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelBidActiveResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelBidRatedResponse struct {
	BidID          string                              `json:"bidID"`
	ModelID        string                              `json:"modelID"`
	PricePerSecond string                              `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                              `json:"providerID"`
	Status         string                              `json:"status"`
	JSON           blockchainModelBidRatedResponseJSON `json:"-"`
}

// blockchainModelBidRatedResponseJSON contains the JSON metadata for the struct
// [BlockchainModelBidRatedResponse]
type blockchainModelBidRatedResponseJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainModelBidRatedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelBidRatedResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelBidListParams struct {
	// Limit for pagination.
	Limit param.Field[int64] `query:"limit"`
	// Offset for pagination.
	Offset param.Field[string] `query:"offset" format:"biginteger"`
}

// URLQuery serializes [BlockchainModelBidListParams]'s query parameters as
// `url.Values`.
func (r BlockchainModelBidListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
