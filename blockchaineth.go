// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/param"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
	"github.com/stainless-sdks/morpheus-marketplace-go/shared"
)

// BlockchainEthService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainEthService] method instead.
type BlockchainEthService struct {
	Options []option.RequestOption
}

// NewBlockchainEthService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainEthService(opts ...option.RequestOption) (r *BlockchainEthService) {
	r = &BlockchainEthService{}
	r.Options = opts
	return
}

// Send ETH to a specified address
func (r *BlockchainEthService) Send(ctx context.Context, body BlockchainEthSendParams, opts ...option.RequestOption) (res *shared.Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/eth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
