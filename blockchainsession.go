// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainSessionService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainSessionService] method instead.
type BlockchainSessionService struct {
	Options []option.RequestOption
}

// NewBlockchainSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainSessionService(opts ...option.RequestOption) (r *BlockchainSessionService) {
	r = &BlockchainSessionService{}
	r.Options = opts
	return
}
