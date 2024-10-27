// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
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
	Options  []option.RequestOption
	User     *BlockchainSessionUserService
	Provider *BlockchainSessionProviderService
	Budget   *BlockchainSessionBudgetService
}

// NewBlockchainSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainSessionService(opts ...option.RequestOption) (r *BlockchainSessionService) {
	r = &BlockchainSessionService{}
	r.Options = opts
	r.User = NewBlockchainSessionUserService(opts...)
	r.Provider = NewBlockchainSessionProviderService(opts...)
	r.Budget = NewBlockchainSessionBudgetService(opts...)
	return
}

// Opens a session with a provider, including staking requirements.
func (r *BlockchainSessionService) New(ctx context.Context, body BlockchainSessionNewParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details of a specific session by its ID.
func (r *BlockchainSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Closes an active session using its unique session ID.
func (r *BlockchainSessionService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainSessionCloseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/sessions/%s/close", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type BlockchainSessionGetResponse struct {
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
	Status string                           `json:"status"`
	JSON   blockchainSessionGetResponseJSON `json:"-"`
}

// blockchainSessionGetResponseJSON contains the JSON metadata for the struct
// [BlockchainSessionGetResponse]
type blockchainSessionGetResponseJSON struct {
	ClosedAt    apijson.Field
	CreatedAt   apijson.Field
	ModelID     apijson.Field
	ProviderID  apijson.Field
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSessionCloseResponse struct {
	// Transaction hash confirming the closure.
	Tx   string                             `json:"tx"`
	JSON blockchainSessionCloseResponseJSON `json:"-"`
}

// blockchainSessionCloseResponseJSON contains the JSON metadata for the struct
// [BlockchainSessionCloseResponse]
type blockchainSessionCloseResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainSessionCloseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainSessionCloseResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainSessionNewParams struct {
	// Approval data for session initiation.
	Approval param.Field[string] `json:"approval,required"`
	// Digital signature associated with the approval.
	ApprovalSig param.Field[string] `json:"approvalSig,required"`
	// Amount of tokens staked for the session.
	Stake param.Field[string] `json:"stake,required" format:"biginteger"`
}

func (r BlockchainSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
