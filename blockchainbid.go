// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/param"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/requestconfig"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
	"github.com/stainless-sdks/morpheus-marketplace-go/shared"
)

// BlockchainBidService contains methods and other services that help with
// interacting with the morpheus-marketplace API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockchainBidService] method instead.
type BlockchainBidService struct {
	Options []option.RequestOption
}

// NewBlockchainBidService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockchainBidService(opts ...option.RequestOption) (r *BlockchainBidService) {
	r = &BlockchainBidService{}
	r.Options = opts
	return
}

// Places a bid for a model on the blockchain.
func (r *BlockchainBidService) New(ctx context.Context, body BlockchainBidNewParams, opts ...option.RequestOption) (res *BlockchainBidNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/bids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves bid details by its ID.
func (r *BlockchainBidService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainBidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/bids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes a bid from the blockchain by its ID.
func (r *BlockchainBidService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BlockchainBidDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/bids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Opens a session based on a specific bid ID, linking the provider and model via a
// bid.
func (r *BlockchainBidService) Session(ctx context.Context, id string, body BlockchainBidSessionParams, opts ...option.RequestOption) (res *shared.Session, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/bids/%s/session", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockchainBidNewResponse struct {
	// Unique identifier for the bid.
	BidID   string                          `json:"bidID"`
	Details BlockchainBidNewResponseDetails `json:"details"`
	JSON    blockchainBidNewResponseJSON    `json:"-"`
}

// blockchainBidNewResponseJSON contains the JSON metadata for the struct
// [BlockchainBidNewResponse]
type blockchainBidNewResponseJSON struct {
	BidID       apijson.Field
	Details     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainBidNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainBidNewResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidNewResponseDetails struct {
	BidID          string                              `json:"bidID"`
	ModelID        string                              `json:"modelID"`
	PricePerSecond string                              `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                              `json:"providerID"`
	Status         string                              `json:"status"`
	JSON           blockchainBidNewResponseDetailsJSON `json:"-"`
}

// blockchainBidNewResponseDetailsJSON contains the JSON metadata for the struct
// [BlockchainBidNewResponseDetails]
type blockchainBidNewResponseDetailsJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainBidNewResponseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainBidNewResponseDetailsJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidGetResponse struct {
	// Unique identifier for the bid.
	BidID   string                          `json:"bidID"`
	Details BlockchainBidGetResponseDetails `json:"details"`
	JSON    blockchainBidGetResponseJSON    `json:"-"`
}

// blockchainBidGetResponseJSON contains the JSON metadata for the struct
// [BlockchainBidGetResponse]
type blockchainBidGetResponseJSON struct {
	BidID       apijson.Field
	Details     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainBidGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainBidGetResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidGetResponseDetails struct {
	BidID          string                              `json:"bidID"`
	ModelID        string                              `json:"modelID"`
	PricePerSecond string                              `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string                              `json:"providerID"`
	Status         string                              `json:"status"`
	JSON           blockchainBidGetResponseDetailsJSON `json:"-"`
}

// blockchainBidGetResponseDetailsJSON contains the JSON metadata for the struct
// [BlockchainBidGetResponseDetails]
type blockchainBidGetResponseDetailsJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockchainBidGetResponseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainBidGetResponseDetailsJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidDeleteResponse struct {
	// Transaction hash.
	Tx   string                          `json:"tx"`
	JSON blockchainBidDeleteResponseJSON `json:"-"`
}

// blockchainBidDeleteResponseJSON contains the JSON metadata for the struct
// [BlockchainBidDeleteResponse]
type blockchainBidDeleteResponseJSON struct {
	Tx          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockchainBidDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockchainBidDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidNewParams struct {
	// ID of the model the bid is associated with.
	ModelID param.Field[string] `json:"modelID,required"`
	// Price per second of model usage.
	PricePerSecond param.Field[string] `json:"pricePerSecond,required" format:"biginteger"`
}

func (r BlockchainBidNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainBidSessionParams struct {
	// The duration of the session in seconds.
	SessionDuration param.Field[string] `json:"sessionDuration,required" format:"biginteger"`
}

func (r BlockchainBidSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
