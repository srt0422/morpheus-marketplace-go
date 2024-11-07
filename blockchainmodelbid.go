// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
	Active  *BlockchainModelBidActiveService
	Rated   *BlockchainModelBidRatedService
}

// NewBlockchainModelBidService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainModelBidService(opts ...option.RequestOption) (r *BlockchainModelBidService) {
	r = &BlockchainModelBidService{}
	r.Options = opts
	r.Active = NewBlockchainModelBidActiveService(opts...)
	r.Rated = NewBlockchainModelBidRatedService(opts...)
	return
}

// List bids for a model
func (r *BlockchainModelBidService) List(ctx context.Context, id string, query BlockchainModelBidListParams, opts ...option.RequestOption) (res *shared.BidList, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/bids", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainModelBidListParams struct {
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Number of results to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BlockchainModelBidListParams]'s query parameters as
// `url.Values`.
func (r BlockchainModelBidListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
