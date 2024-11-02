// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
	Budget   *BlockchainSessionBudgetService
	User     *BlockchainSessionUserService
	Provider *BlockchainSessionProviderService
}

// NewBlockchainSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainSessionService(opts ...option.RequestOption) (r *BlockchainSessionService) {
	r = &BlockchainSessionService{}
	r.Options = opts
	r.Budget = NewBlockchainSessionBudgetService(opts...)
	r.User = NewBlockchainSessionUserService(opts...)
	r.Provider = NewBlockchainSessionProviderService(opts...)
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
