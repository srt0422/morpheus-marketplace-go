// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
)

// BlockchainModelBidActiveService contains methods and other services that help
// with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelBidActiveService] method instead.
type BlockchainModelBidActiveService struct {
	Options []option.RequestOption
}

// NewBlockchainModelBidActiveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBlockchainModelBidActiveService(opts ...option.RequestOption) (r *BlockchainModelBidActiveService) {
	r = &BlockchainModelBidActiveService{}
	r.Options = opts
	return
}

// List active bids for a model
func (r *BlockchainModelBidActiveService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.BidList, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/bids/active", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
