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
)

// BlockchainModelService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelService] method instead.
type BlockchainModelService struct {
	Options  []option.RequestOption
	Sessions *BlockchainModelSessionService
	Bids     *BlockchainModelBidService
	Stats    *BlockchainModelStatService
}

// NewBlockchainModelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainModelService(opts ...option.RequestOption) (r *BlockchainModelService) {
	r = &BlockchainModelService{}
	r.Options = opts
	r.Sessions = NewBlockchainModelSessionService(opts...)
	r.Bids = NewBlockchainModelBidService(opts...)
	r.Stats = NewBlockchainModelStatService(opts...)
	return
}

// Create a new model
func (r *BlockchainModelService) New(ctx context.Context, body BlockchainModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all available models
func (r *BlockchainModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a model
func (r *BlockchainModelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Model struct {
	// Unique identifier of the model
	ID string `json:"id,required"`
	// Fee for using the model
	Fee string `json:"fee,required"`
	// IPFS ID where the model is stored
	IpfsID string `json:"ipfsID,required"`
	// Model ID provided by the user
	ModelID string `json:"modelID,required"`
	// Name of the model
	Name string `json:"name,required"`
	// Amount staked for the model
	Stake string `json:"stake,required"`
	// Tags associated with the model
	Tags []string  `json:"tags"`
	JSON modelJSON `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	ID          apijson.Field
	Fee         apijson.Field
	IpfsID      apijson.Field
	ModelID     apijson.Field
	Name        apijson.Field
	Stake       apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelNewParams struct {
	// Fee for using the model
	Fee param.Field[string] `json:"fee,required"`
	// IPFS ID where the model is stored
	IpfsID param.Field[string] `json:"ipfsID,required"`
	// Model ID provided by the user
	ModelID param.Field[string] `json:"modelID,required"`
	// Name of the model
	Name param.Field[string] `json:"name,required"`
	// Amount to stake for the model
	Stake param.Field[string] `json:"stake,required"`
	// Tags associated with the model
	Tags param.Field[[]string] `json:"tags"`
}

func (r BlockchainModelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
