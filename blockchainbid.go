// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainBidService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainBidService] method instead.
type BlockchainBidService struct {
	Options []option.RequestOption
}

// NewBlockchainBidService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainBidService(opts ...option.RequestOption) (r *BlockchainBidService) {
	r = &BlockchainBidService{}
	r.Options = opts
	return
}
