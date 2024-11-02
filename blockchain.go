// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/param"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

// BlockchainService contains methods and other services that help with interacting
// with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainService] method instead.
type BlockchainService struct {
	Options      []option.RequestOption
	Eth          *BlockchainEthService
	Mor          *BlockchainMorService
	Models       *BlockchainModelService
	Bids         *BlockchainBidService
	Sessions     *BlockchainSessionService
	Providers    *BlockchainProviderService
	Balance      *BlockchainBalanceService
	Allowance    *BlockchainAllowanceService
	LatestBlock  *BlockchainLatestBlockService
	Token        *BlockchainTokenService
	Transactions *BlockchainTransactionService
}

// NewBlockchainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainService(opts ...option.RequestOption) (r *BlockchainService) {
	r = &BlockchainService{}
	r.Options = opts
	r.Eth = NewBlockchainEthService(opts...)
	r.Mor = NewBlockchainMorService(opts...)
	r.Models = NewBlockchainModelService(opts...)
	r.Bids = NewBlockchainBidService(opts...)
	r.Sessions = NewBlockchainSessionService(opts...)
	r.Providers = NewBlockchainProviderService(opts...)
	r.Balance = NewBlockchainBalanceService(opts...)
	r.Allowance = NewBlockchainAllowanceService(opts...)
	r.LatestBlock = NewBlockchainLatestBlockService(opts...)
	r.Token = NewBlockchainTokenService(opts...)
	r.Transactions = NewBlockchainTransactionService(opts...)
	return
}

// Approve allowance
func (r *BlockchainService) Approve(ctx context.Context, body BlockchainApproveParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "blockchain/approve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type BlockchainApproveParams struct {
	// Amount to approve
	Amount param.Field[string] `json:"amount,required"`
	// Spender Ethereum address
	Spender param.Field[string] `json:"spender,required"`
}

func (r BlockchainApproveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
