// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
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
	Providers    *BlockchainProviderService
	Balance      *BlockchainBalanceService
	Allowance    *BlockchainAllowanceService
	LatestBlock  *BlockchainLatestBlockService
	Token        *BlockchainTokenService
	Transactions *BlockchainTransactionService
	Bids         *BlockchainBidService
	Sessions     *BlockchainSessionService
}

// NewBlockchainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainService(opts ...option.RequestOption) (r *BlockchainService) {
	r = &BlockchainService{}
	r.Options = opts
	r.Models = NewBlockchainModelService(opts...)
	r.Providers = NewBlockchainProviderService(opts...)
	r.Balance = NewBlockchainBalanceService(opts...)
	r.Allowance = NewBlockchainAllowanceService(opts...)
	r.LatestBlock = NewBlockchainLatestBlockService(opts...)
	r.Token = NewBlockchainTokenService(opts...)
	r.Transactions = NewBlockchainTransactionService(opts...)
	r.Bids = NewBlockchainBidService(opts...)
	r.Sessions = NewBlockchainSessionService(opts...)
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

// Send ETH to a specified address
func (r *BlockchainService) EthSend(ctx context.Context, body BlockchainEthSendParams, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/eth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send MOR to a specified address
func (r *BlockchainService) MorSend(ctx context.Context, body BlockchainMorSendParams, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/mor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Balance struct {
	// Current balance after the transaction
	Balance string      `json:"balance,required"`
	JSON    balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
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

type BlockchainEthSendParams struct {
	// Amount of ETH to send
	Amount param.Field[string] `json:"amount,required"`
	// Ethereum address to send ETH to
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainEthSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainMorSendParams struct {
	// Amount of MOR to send
	Amount param.Field[string] `json:"amount,required"`
	// Ethereum address to send MOR to
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainMorSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
