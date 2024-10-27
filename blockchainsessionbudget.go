// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainSessionBudgetService contains methods and other services that help
// with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainSessionBudgetService] method instead.
type BlockchainSessionBudgetService struct {
	Options []option.RequestOption
}

// NewBlockchainSessionBudgetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainSessionBudgetService(opts ...option.RequestOption) (r *BlockchainSessionBudgetService) {
	r = &BlockchainSessionBudgetService{}
	r.Options = opts
	return
}

// Fetches the total budget for sessions on the current day.
func (r *BlockchainSessionBudgetService) List(ctx context.Context, opts ...option.RequestOption) (res *BlockchainSessionBudgetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/budget"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BlockchainSessionBudgetListResponse struct {
	// Total budget for sessions on the current day.
	Budget string                                  `json:"budget" format:"biginteger"`
	JSON   blockchainSessionBudgetListResponseJSON `json:"-"`
}

// blockchainSessionBudgetListResponseJSON contains the JSON metadata for the
// struct [BlockchainSessionBudgetListResponse]
type blockchainSessionBudgetListResponseJSON struct {
	Budget      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSessionBudgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSessionBudgetListResponseJSON) RawJSON() string {
	return r.raw
}
