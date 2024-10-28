// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"
	"net/url"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainAllowanceService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainAllowanceService] method instead.
type BlockchainAllowanceService struct {
	Options []option.RequestOption
}

// NewBlockchainAllowanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainAllowanceService(opts ...option.RequestOption) (r *BlockchainAllowanceService) {
	r = &BlockchainAllowanceService{}
	r.Options = opts
	return
}

// Retrieve allowance
func (r *BlockchainAllowanceService) Get(ctx context.Context, query BlockchainAllowanceGetParams, opts ...option.RequestOption) (res *BlockchainAllowanceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/allowance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainAllowanceGetResponse struct {
	// Current allowance amount
	Allowance string                             `json:"allowance,required"`
	JSON      blockchainAllowanceGetResponseJSON `json:"-"`
}

// blockchainAllowanceGetResponseJSON contains the JSON metadata for the struct
// [BlockchainAllowanceGetResponse]
type blockchainAllowanceGetResponseJSON struct {
	Allowance   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainAllowanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainAllowanceGetResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainAllowanceGetParams struct {
	// Spender Ethereum address
	Spender param.Field[string] `query:"spender,required"`
}

// URLQuery serializes [BlockchainAllowanceGetParams]'s query parameters as
// `url.Values`.
func (r BlockchainAllowanceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
