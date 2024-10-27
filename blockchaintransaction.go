// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainTransactionService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainTransactionService] method instead.
type BlockchainTransactionService struct {
	Options []option.RequestOption
}

// NewBlockchainTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainTransactionService(opts ...option.RequestOption) (r *BlockchainTransactionService) {
	r = &BlockchainTransactionService{}
	r.Options = opts
	return
}

// Retrieves ETH and MOR token transactions for the user.
func (r *BlockchainTransactionService) List(ctx context.Context, query BlockchainTransactionListParams, opts ...option.RequestOption) (res *TransactionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TransactionList []TransactionListItem

type TransactionListItem struct {
	Amount    string                  `json:"amount" format:"biginteger"`
	From      string                  `json:"from"`
	Timestamp time.Time               `json:"timestamp" format:"date-time"`
	To        string                  `json:"to"`
	TxHash    string                  `json:"txHash"`
	JSON      transactionListItemJSON `json:"-"`
}

// transactionListItemJSON contains the JSON metadata for the struct
// [TransactionListItem]
type transactionListItemJSON struct {
	Amount      apijson.Field
	From        apijson.Field
	Timestamp   apijson.Field
	To          apijson.Field
	TxHash      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListItemJSON) RawJSON() string {
	return r.raw
}

type BlockchainTransactionListResponse struct {
	Amount    string                                `json:"amount" format:"biginteger"`
	From      string                                `json:"from"`
	Timestamp time.Time                             `json:"timestamp" format:"date-time"`
	To        string                                `json:"to"`
	TxHash    string                                `json:"txHash"`
	JSON      blockchainTransactionListResponseJSON `json:"-"`
}

// blockchainTransactionListResponseJSON contains the JSON metadata for the struct
// [BlockchainTransactionListResponse]
type blockchainTransactionListResponseJSON struct {
	Amount      apijson.Field
	From        apijson.Field
	Timestamp   apijson.Field
	To          apijson.Field
	TxHash      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainTransactionListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainTransactionListParams struct {
	// Limit for pagination.
	Limit param.Field[int64] `query:"limit"`
	// Page number for pagination.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [BlockchainTransactionListParams]'s query parameters as
// `url.Values`.
func (r BlockchainTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
