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

// BlockchainModelService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainModelService] method instead.
type BlockchainModelService struct {
	Options  []option.RequestOption
	Bids     *BlockchainModelBidService
	Minstake *BlockchainModelMinstakeService
	Stats    *BlockchainModelStatService
}

// NewBlockchainModelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainModelService(opts ...option.RequestOption) (r *BlockchainModelService) {
	r = &BlockchainModelService{}
	r.Options = opts
	r.Bids = NewBlockchainModelBidService(opts...)
	r.Minstake = NewBlockchainModelMinstakeService(opts...)
	r.Stats = NewBlockchainModelStatService(opts...)
	return
}

// Registers a new model on the blockchain.
func (r *BlockchainModelService) New(ctx context.Context, body BlockchainModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a list of all models available on the blockchain.
func (r *BlockchainModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]BlockchainModelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes a model from the blockchain by its ID.
func (r *BlockchainModelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainModelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Checks if a model exists on the blockchain.
func (r *BlockchainModelService) Exists(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainModelExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/exists", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resets the statistics of a model.
func (r *BlockchainModelService) Resetstats(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainModelResetstatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/resetstats", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Opens a session with a specific model, associating it with a provider on the
// blockchain.
func (r *BlockchainModelService) Session(ctx context.Context, id string, body BlockchainModelSessionParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/models/%s/session", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Model struct {
	Details ModelDetails `json:"details"`
	// Unique identifier for the model.
	ModelID string    `json:"modelID"`
	JSON    modelJSON `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	Details     apijson.Field
	ModelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

type ModelDetails struct {
	Fee     string           `json:"fee" format:"biginteger"`
	IpfsID  string           `json:"ipfsID"`
	ModelID string           `json:"modelID"`
	Name    string           `json:"name"`
	Stake   string           `json:"stake" format:"biginteger"`
	Tags    []string         `json:"tags"`
	JSON    modelDetailsJSON `json:"-"`
}

// modelDetailsJSON contains the JSON metadata for the struct [ModelDetails]
type modelDetailsJSON struct {
	Fee         apijson.Field
	IpfsID      apijson.Field
	ModelID     apijson.Field
	Name        apijson.Field
	Stake       apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelDetailsJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelListResponse struct {
	Fee     string                          `json:"fee" format:"biginteger"`
	IpfsID  string                          `json:"ipfsID"`
	ModelID string                          `json:"modelID"`
	Name    string                          `json:"name"`
	Stake   string                          `json:"stake" format:"biginteger"`
	Tags    []string                        `json:"tags"`
	JSON    blockchainModelListResponseJSON `json:"-"`
}

// blockchainModelListResponseJSON contains the JSON metadata for the struct
// [BlockchainModelListResponse]
type blockchainModelListResponseJSON struct {
	Fee         apijson.Field
	IpfsID      apijson.Field
	ModelID     apijson.Field
	Name        apijson.Field
	Stake       apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelListResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelDeleteResponse struct {
	// Transaction hash.
	Tx   string                            `json:"tx"`
	JSON blockchainModelDeleteResponseJSON `json:"-"`
}

// blockchainModelDeleteResponseJSON contains the JSON metadata for the struct
// [BlockchainModelDeleteResponse]
type blockchainModelDeleteResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainModelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelExistsResponse struct {
	// Indicates whether the model exists.
	Exists bool                              `json:"exists"`
	JSON   blockchainModelExistsResponseJSON `json:"-"`
}

// blockchainModelExistsResponseJSON contains the JSON metadata for the struct
// [BlockchainModelExistsResponse]
type blockchainModelExistsResponseJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainModelExistsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelExistsResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelResetstatsResponse struct {
	// Success message.
	Message string                                `json:"message"`
	JSON    blockchainModelResetstatsResponseJSON `json:"-"`
}

// blockchainModelResetstatsResponseJSON contains the JSON metadata for the struct
// [BlockchainModelResetstatsResponse]
type blockchainModelResetstatsResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainModelResetstatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainModelResetstatsResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainModelNewParams struct {
	// Fee amount required for model usage.
	Fee param.Field[string] `json:"fee,required" format:"biginteger"`
	// IPFS hash storing the model’s data.
	IpfsID param.Field[string] `json:"ipfsID,required"`
	// Unique identifier for the model.
	ModelID param.Field[string] `json:"modelID,required"`
	// Name of the model.
	Name param.Field[string] `json:"name,required"`
	// Stake amount for the model.
	Stake param.Field[string] `json:"stake,required" format:"biginteger"`
	// Descriptive tags for categorizing the model.
	Tags param.Field[[]string] `json:"tags"`
}

func (r BlockchainModelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainModelSessionParams struct {
	// The duration of the session in seconds.
	SessionDuration param.Field[string] `json:"sessionDuration,required" format:"biginteger"`
}

func (r BlockchainModelSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
