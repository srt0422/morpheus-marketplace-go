// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainService contains methods and other services that help with interacting
// with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainService] method instead.
type BlockchainService struct {
	Options      []option.RequestOption
	Models       *BlockchainModelService
	Bids         *BlockchainBidService
	Sessions     *BlockchainSessionService
	Providers    *BlockchainProviderService
	Balance      *BlockchainBalanceService
	Allowance    *BlockchainAllowanceService
	Send         *BlockchainSendService
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
	r.Models = NewBlockchainModelService(opts...)
	r.Bids = NewBlockchainBidService(opts...)
	r.Sessions = NewBlockchainSessionService(opts...)
	r.Providers = NewBlockchainProviderService(opts...)
	r.Balance = NewBlockchainBalanceService(opts...)
	r.Allowance = NewBlockchainAllowanceService(opts...)
	r.Send = NewBlockchainSendService(opts...)
	r.LatestBlock = NewBlockchainLatestBlockService(opts...)
	r.Token = NewBlockchainTokenService(opts...)
	r.Transactions = NewBlockchainTransactionService(opts...)
	return
}
