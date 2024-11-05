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

// BlockchainModelSessionService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelSessionService] method instead.
type BlockchainModelSessionService struct {
	Options []option.RequestOption
}

// NewBlockchainModelSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockchainModelSessionService(opts ...option.RequestOption) (r *BlockchainModelSessionService) {
	r = &BlockchainModelSessionService{}
	r.Options = opts
	return
}

// Start a session for a model
func (r *BlockchainModelSessionService) New(ctx context.Context, id string, body BlockchainModelSessionNewParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/session", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockchainModelSessionNewParams struct {
	// Duration of the session
	SessionDuration param.Field[string] `json:"sessionDuration,required"`
}

func (r BlockchainModelSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
