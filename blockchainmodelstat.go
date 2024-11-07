// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainModelStatService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelStatService] method instead.
type BlockchainModelStatService struct {
	Options []option.RequestOption
}

// NewBlockchainModelStatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainModelStatService(opts ...option.RequestOption) (r *BlockchainModelStatService) {
	r = &BlockchainModelStatService{}
	r.Options = opts
	return
}

// Retrieve statistics for a model
func (r *BlockchainModelStatService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ModelStats, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/stats", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelStats struct {
	// ID of the model
	ModelID string `json:"modelID,required"`
	// Statistics related to the model
	Stats map[string]interface{} `json:"stats,required"`
	JSON  modelStatsJSON         `json:"-"`
}

// modelStatsJSON contains the JSON metadata for the struct [ModelStats]
type modelStatsJSON struct {
	ModelID     apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelStatsJSON) RawJSON() string {
	return r.raw
}
