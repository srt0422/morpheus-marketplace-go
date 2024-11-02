// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"net/http"
	"net/url"

	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
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

// List user sessions
func (r *BlockchainSessionUserService) List(ctx context.Context, query BlockchainSessionUserListParams, opts ...option.RequestOption) (res *shared.SessionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BlockchainSessionUserListParams struct {
	// User identifier
	User param.Field[string] `query:"user,required"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Number of results to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BlockchainSessionUserListParams]'s query parameters as
// `url.Values`.
func (r BlockchainSessionUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
