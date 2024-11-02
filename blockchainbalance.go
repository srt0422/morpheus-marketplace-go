// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
)

// BlockchainBalanceService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainBalanceService] method instead.
type BlockchainBalanceService struct {
	Options []option.RequestOption
}

// NewBlockchainBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainBalanceService(opts ...option.RequestOption) (r *BlockchainBalanceService) {
	r = &BlockchainBalanceService{}
	r.Options = opts
	return
}

// Retrieve balance
func (r *BlockchainBalanceService) Get(ctx context.Context, opts ...option.RequestOption) (res *shared.Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
