// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

// BlockchainLatestBlockService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainLatestBlockService] method instead.
type BlockchainLatestBlockService struct {
	Options []option.RequestOption
}

// NewBlockchainLatestBlockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainLatestBlockService(opts ...option.RequestOption) (r *BlockchainLatestBlockService) {
	r = &BlockchainLatestBlockService{}
	r.Options = opts
	return
}

// Retrieves the latest block number from the blockchain.
func (r *BlockchainLatestBlockService) Get(ctx context.Context, opts ...option.RequestOption) (res *LatestBlock, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/latestBlock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LatestBlock struct {
	// Latest block number.
	Block int64           `json:"block"`
	JSON  latestBlockJSON `json:"-"`
}

// latestBlockJSON contains the JSON metadata for the struct [LatestBlock]
type latestBlockJSON struct {
	Block       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LatestBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r latestBlockJSON) RawJSON() string {
	return r.raw
}
