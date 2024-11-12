// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
	"github.com/srt0422/morpheus-marketplace-go/internal/apiquery"
	"github.com/srt0422/morpheus-marketplace-go/internal/param"
	"github.com/srt0422/morpheus-marketplace-go/internal/requestconfig"
	"github.com/srt0422/morpheus-marketplace-go/option"
	"github.com/srt0422/morpheus-marketplace-go/shared"
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

// Create a new session
func (r *BlockchainSessionService) New(ctx context.Context, body BlockchainSessionNewParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a session
func (r *BlockchainSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get session budget
func (r *BlockchainSessionService) Budget(ctx context.Context, opts ...option.RequestOption) (res *Budget, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/budget"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Close a session
func (r *BlockchainSessionService) Close(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/sessions/%s/close", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// List provider sessions
func (r *BlockchainSessionService) Provider(ctx context.Context, query BlockchainSessionProviderParams, opts ...option.RequestOption) (res *SessionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List user sessions
func (r *BlockchainSessionService) User(ctx context.Context, query BlockchainSessionUserParams, opts ...option.RequestOption) (res *SessionList, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Budget struct {
	// Current session budget
	Budget string     `json:"budget,required"`
	JSON   budgetJSON `json:"-"`
}

// budgetJSON contains the JSON metadata for the struct [Budget]
type budgetJSON struct {
	Budget      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Budget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetJSON) RawJSON() string {
	return r.raw
}

type SessionList struct {
	// List of sessions
	Sessions []shared.Session `json:"sessions,required"`
	JSON     sessionListJSON  `json:"-"`
}

// sessionListJSON contains the JSON metadata for the struct [SessionList]
type sessionListJSON struct {
	Sessions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListJSON) RawJSON() string {
	return r.raw
}

type BlockchainSessionNewParams struct {
	// Approval identifier
	Approval param.Field[string] `json:"approval,required"`
	// Signature for the approval
	ApprovalSig param.Field[string] `json:"approvalSig,required"`
	// Stake amount for the session
	Stake param.Field[string] `json:"stake,required"`
}

func (r BlockchainSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainSessionProviderParams struct {
	// Provider identifier
	Provider param.Field[string] `query:"provider,required"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Number of results to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BlockchainSessionProviderParams]'s query parameters as
// `url.Values`.
func (r BlockchainSessionProviderParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BlockchainSessionUserParams struct {
	// User identifier
	User param.Field[string] `query:"user,required"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Number of results to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BlockchainSessionUserParams]'s query parameters as
// `url.Values`.
func (r BlockchainSessionUserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
