// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
)

// BlockchainService contains methods and other services that help with interacting
// with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainService] method instead.
type BlockchainService struct {
	Options []option.RequestOption
	Models  *BlockchainModelService
}

// NewBlockchainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainService(opts ...option.RequestOption) (r *BlockchainService) {
	r = &BlockchainService{}
	r.Options = opts
	r.Models = NewBlockchainModelService(opts...)
	return
}

// Send ETH to a specified address
func (r *BlockchainService) SendEth(ctx context.Context, body BlockchainSendEthParams, opts ...option.RequestOption) (res *shared.Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/eth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send MOR to a specified address
func (r *BlockchainService) SendMor(ctx context.Context, body BlockchainSendMorParams, opts ...option.RequestOption) (res *shared.Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/mor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockchainSendEthParams struct {
	// Amount of ETH to send
	Amount param.Field[string] `json:"amount,required"`
	// Ethereum address to send ETH to
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainSendEthParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainSendMorParams struct {
	// Amount of MOR to send
	Amount param.Field[string] `json:"amount,required"`
	// Ethereum address to send MOR to
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainSendMorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
