// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
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

// Get token supply
func (r *BlockchainTokenSupplyService) Get(ctx context.Context, opts ...option.RequestOption) (res *shared.TokenSupply, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/token/supply"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
