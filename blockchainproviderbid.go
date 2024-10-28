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

// BlockchainProviderBidService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainProviderBidService] method instead.
type BlockchainProviderBidService struct {
	Options []option.RequestOption
	Active  *BlockchainProviderBidActiveService
}

// NewBlockchainProviderBidService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainProviderBidService(opts ...option.RequestOption) (r *BlockchainProviderBidService) {
	r = &BlockchainProviderBidService{}
	r.Options = opts
	r.Active = NewBlockchainProviderBidActiveService(opts...)
	return
}

// List bids for a provider
func (r *BlockchainProviderBidService) List(ctx context.Context, id string, query BlockchainProviderBidListParams, opts ...option.RequestOption) (res *BlockchainProviderBidListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/providers/%s/bids", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainProviderBidListResponse struct {
	// List of bids
	Bids []shared.Bid                          `json:"bids,required"`
	JSON blockchainProviderBidListResponseJSON `json:"-"`
}

// blockchainProviderBidListResponseJSON contains the JSON metadata for the struct
// [BlockchainProviderBidListResponse]
type blockchainProviderBidListResponseJSON struct {
	Bids        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainProviderBidListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainProviderBidListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainProviderBidListParams struct {
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Number of results to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BlockchainProviderBidListParams]'s query parameters as
// `url.Values`.
func (r BlockchainProviderBidListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
