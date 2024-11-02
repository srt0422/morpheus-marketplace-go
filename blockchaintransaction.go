// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
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

// List transactions
func (r *BlockchainTransactionService) List(ctx context.Context, opts ...option.RequestOption) (res *TransactionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionList struct {
	// List of transactions
	Transactions []TransactionListTransaction `json:"transactions,required"`
	JSON         transactionListJSON          `json:"-"`
}

// transactionListJSON contains the JSON metadata for the struct [TransactionList]
type transactionListJSON struct {
	Transactions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TransactionList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListJSON) RawJSON() string {
	return r.raw
}

type TransactionListTransaction struct {
	// Transaction ID
	ID string `json:"id,required"`
	// Amount involved in the transaction
	Amount string `json:"amount,required"`
	// Type of transaction
	Type string                         `json:"type,required"`
	JSON transactionListTransactionJSON `json:"-"`
}

// transactionListTransactionJSON contains the JSON metadata for the struct
// [TransactionListTransaction]
type transactionListTransactionJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListTransactionJSON) RawJSON() string {
	return r.raw
}
