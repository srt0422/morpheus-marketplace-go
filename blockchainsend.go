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

// BlockchainSendService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainSendService] method instead.
type BlockchainSendService struct {
	Options []option.RequestOption
}

// NewBlockchainSendService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainSendService(opts ...option.RequestOption) (r *BlockchainSendService) {
	r = &BlockchainSendService{}
	r.Options = opts
	return
}

// Transfers ETH to a specified address.
func (r *BlockchainSendService) Eth(ctx context.Context, body BlockchainSendEthParams, opts ...option.RequestOption) (res *BlockchainSendEthResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/eth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transfers MOR tokens to a specified address.
func (r *BlockchainSendService) Mor(ctx context.Context, body BlockchainSendMorParams, opts ...option.RequestOption) (res *BlockchainSendMorResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/send/mor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockchainSendEthResponse struct {
	// Transaction hash.
	Tx   string                        `json:"tx"`
	JSON blockchainSendEthResponseJSON `json:"-"`
}

// blockchainSendEthResponseJSON contains the JSON metadata for the struct
// [BlockchainSendEthResponse]
type blockchainSendEthResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSendEthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSendEthResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSendMorResponse struct {
	// Transaction hash.
	Tx   string                        `json:"tx"`
	JSON blockchainSendMorResponseJSON `json:"-"`
}

// blockchainSendMorResponseJSON contains the JSON metadata for the struct
// [BlockchainSendMorResponse]
type blockchainSendMorResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSendMorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSendMorResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSendEthParams struct {
	// Amount of ETH to be transferred.
	Amount param.Field[string] `json:"amount,required" format:"biginteger"`
	// Recipient ETH address.
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainSendEthParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainSendMorParams struct {
	// Amount of MOR tokens to be transferred.
	Amount param.Field[string] `json:"amount,required" format:"biginteger"`
	// Recipient ETH address.
	To param.Field[string] `json:"to,required"`
}

func (r BlockchainSendMorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
