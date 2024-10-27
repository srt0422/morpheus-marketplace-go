// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

// BlockchainSessionUserService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainSessionUserService] method instead.
type BlockchainSessionUserService struct {
	Options []option.RequestOption
}

// NewBlockchainSessionUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainSessionUserService(opts ...option.RequestOption) (r *BlockchainSessionUserService) {
	r = &BlockchainSessionUserService{}
	r.Options = opts
	return
}

// Retrieves a list of all active and closed sessions associated with a specific
// user.
func (r *BlockchainSessionUserService) List(ctx context.Context, query BlockchainSessionUserListParams, opts ...option.RequestOption) (res *[]BlockchainSessionUserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainSessionUserListResponse struct {
	// Closure timestamp.
	ClosedAt time.Time `json:"closedAt" format:"date-time"`
	// Creation timestamp.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Model associated with the session.
	ModelID string `json:"modelID"`
	// Provider associated with the session.
	ProviderID string `json:"providerID"`
	// Unique identifier for the session.
	SessionID string `json:"sessionID"`
	// Status of the session (active, closed, etc.).
	Status string                                `json:"status"`
	JSON   blockchainSessionUserListResponseJSON `json:"-"`
}

// blockchainSessionUserListResponseJSON contains the JSON metadata for the struct
// [BlockchainSessionUserListResponse]
type blockchainSessionUserListResponseJSON struct {
	ClosedAt    apijson.Field
	CreatedAt   apijson.Field
	ModelID     apijson.Field
	ProviderID  apijson.Field
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSessionUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSessionUserListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSessionUserListParams struct {
	// Ethereum address of the user.
	User param.Field[string] `query:"user,required"`
	// Limit for pagination.
	Limit param.Field[int64] `query:"limit"`
	// Offset for pagination.
	Offset param.Field[string] `query:"offset" format:"biginteger"`
}

// URLQuery serializes [BlockchainSessionUserListParams]'s query parameters as
// `url.Values`.
func (r BlockchainSessionUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
