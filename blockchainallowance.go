// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apiquery"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/param"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
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

// Fetches the MOR token allowance allocated for a specific spender.
func (r *BlockchainAllowanceService) Get(ctx context.Context, query BlockchainAllowanceGetParams, opts ...option.RequestOption) (res *Allowance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/allowance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Approves a specified MOR token allowance for a designated spender.
func (r *BlockchainAllowanceService) Approve(ctx context.Context, body BlockchainAllowanceApproveParams, opts ...option.RequestOption) (res *BlockchainAllowanceApproveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/allowance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Allowance struct {
	// MOR token allowance.
	Allowance string        `json:"allowance" format:"biginteger"`
	JSON      allowanceJSON `json:"-"`
}

// allowanceJSON contains the JSON metadata for the struct [Allowance]
type allowanceJSON struct {
	Allowance   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Allowance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r allowanceJSON) RawJSON() string {
	return r.raw
}

type BlockchainAllowanceApproveResponse struct {
	// Transaction hash.
	Tx   string                                 `json:"tx"`
	JSON blockchainAllowanceApproveResponseJSON `json:"-"`
}

// blockchainAllowanceApproveResponseJSON contains the JSON metadata for the struct
// [BlockchainAllowanceApproveResponse]
type blockchainAllowanceApproveResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainAllowanceApproveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainAllowanceApproveResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainAllowanceGetParams struct {
	// Ethereum address of the spender.
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

type BlockchainAllowanceApproveParams struct {
	// Amount to be approved or transferred.
	Amount param.Field[string] `query:"amount,required" format:"biginteger"`
	// Ethereum address of the spender.
	Spender param.Field[string] `query:"spender,required"`
}

// URLQuery serializes [BlockchainAllowanceApproveParams]'s query parameters as
// `url.Values`.
func (r BlockchainAllowanceApproveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
