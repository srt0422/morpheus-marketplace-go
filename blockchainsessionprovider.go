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

// BlockchainSessionProviderService contains methods and other services that help
// with interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainSessionProviderService] method instead.
type BlockchainSessionProviderService struct {
	Options []option.RequestOption
}

// NewBlockchainSessionProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBlockchainSessionProviderService(opts ...option.RequestOption) (r *BlockchainSessionProviderService) {
	r = &BlockchainSessionProviderService{}
	r.Options = opts
	return
}

// Retrieves a list of sessions for a specified provider.
func (r *BlockchainSessionProviderService) List(ctx context.Context, query BlockchainSessionProviderListParams, opts ...option.RequestOption) (res *[]BlockchainSessionProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainSessionProviderListResponse struct {
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
	Status string                                    `json:"status"`
	JSON   blockchainSessionProviderListResponseJSON `json:"-"`
}

// blockchainSessionProviderListResponseJSON contains the JSON metadata for the
// struct [BlockchainSessionProviderListResponse]
type blockchainSessionProviderListResponseJSON struct {
	ClosedAt    apijson.Field
	CreatedAt   apijson.Field
	ModelID     apijson.Field
	ProviderID  apijson.Field
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSessionProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSessionProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSessionProviderListParams struct {
	// Ethereum address of the provider.
	Provider param.Field[string] `query:"provider,required"`
	// Limit for pagination.
	Limit param.Field[int64] `query:"limit"`
	// Offset for pagination.
	Offset param.Field[string] `query:"offset" format:"biginteger"`
}

// URLQuery serializes [BlockchainSessionProviderListParams]'s query parameters as
// `url.Values`.
func (r BlockchainSessionProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
