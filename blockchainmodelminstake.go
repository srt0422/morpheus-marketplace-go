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

// BlockchainModelMinstakeService contains methods and other services that help
// with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelMinstakeService] method instead.
type BlockchainModelMinstakeService struct {
	Options []option.RequestOption
}

// NewBlockchainModelMinstakeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainModelMinstakeService(opts ...option.RequestOption) (r *BlockchainModelMinstakeService) {
	r = &BlockchainModelMinstakeService{}
	r.Options = opts
	return
}

// Retrieves the current minimum stake required for models.
func (r *BlockchainModelMinstakeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MinStake, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models/minstake"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sets the minimum stake required for models.
func (r *BlockchainModelMinstakeService) Set(ctx context.Context, body BlockchainModelMinstakeSetParams, opts ...option.RequestOption) (res *BlockchainModelMinstakeSetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models/minstake"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MinStake struct {
	// Current minimum stake required for models.
	MinStake string       `json:"minStake" format:"biginteger"`
	JSON     minStakeJSON `json:"-"`
}

// minStakeJSON contains the JSON metadata for the struct [MinStake]
type minStakeJSON struct {
	MinStake    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MinStake) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minStakeJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelMinstakeSetResponse struct {
	// Success message.
	Message string                                 `json:"message"`
	JSON    blockchainModelMinstakeSetResponseJSON `json:"-"`
}

// blockchainModelMinstakeSetResponseJSON contains the JSON metadata for the struct
// [BlockchainModelMinstakeSetResponse]
type blockchainModelMinstakeSetResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainModelMinstakeSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelMinstakeSetResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelMinstakeSetParams struct {
	// Minimum stake amount.
	MinStake param.Field[string] `json:"minStake,required" format:"biginteger"`
}

func (r BlockchainModelMinstakeSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
