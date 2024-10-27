// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

// BlockchainTokenSupplyService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainTokenSupplyService] method instead.
type BlockchainTokenSupplyService struct {
	Options []option.RequestOption
}

// NewBlockchainTokenSupplyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainTokenSupplyService(opts ...option.RequestOption) (r *BlockchainTokenSupplyService) {
	r = &BlockchainTokenSupplyService{}
	r.Options = opts
	return
}

// Fetches the total supply of MOR tokens.
func (r *BlockchainTokenSupplyService) Get(ctx context.Context, opts ...option.RequestOption) (res *TokenSupply, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/token/supply"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TokenSupply struct {
	// Total supply of MOR tokens.
	Supply string          `json:"supply" format:"biginteger"`
	JSON   tokenSupplyJSON `json:"-"`
}

// tokenSupplyJSON contains the JSON metadata for the struct [TokenSupply]
type tokenSupplyJSON struct {
	Supply      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSupply) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSupplyJSON) RawJSON() string {
	return r.raw
}
