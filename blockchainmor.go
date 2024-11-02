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

// BlockchainMorService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainMorService] method instead.
type BlockchainMorService struct {
	Options []option.RequestOption
}

// NewBlockchainMorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainMorService(opts ...option.RequestOption) (r *BlockchainMorService) {
	r = &BlockchainMorService{}
	r.Options = opts
	return
}

// Send MOR to a specified address
func (r *BlockchainMorService) Send(ctx context.Context, body BlockchainMorSendParams, opts ...option.RequestOption) (res *shared.Balance, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/mor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
