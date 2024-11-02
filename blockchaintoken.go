// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

// BlockchainTokenService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainTokenService] method instead.
type BlockchainTokenService struct {
	Options []option.RequestOption
	Supply  *BlockchainTokenSupplyService
}

// NewBlockchainTokenService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainTokenService(opts ...option.RequestOption) (r *BlockchainTokenService) {
	r = &BlockchainTokenService{}
	r.Options = opts
	r.Supply = NewBlockchainTokenSupplyService(opts...)
	return
}
