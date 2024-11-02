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

// Create a new bid
func (r *BlockchainBidService) New(ctx context.Context, body BlockchainBidNewParams, opts ...option.RequestOption) (res *Bid, err error) {
	opts = append(r.Options[:], opts...)
	path := "blockchain/bids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a bid
func (r *BlockchainBidService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Bid, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/bids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a bid
func (r *BlockchainBidService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("blockchain/bids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Start a session for a bid
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

type Bid struct {
	// Unique identifier of the bid
	ID string `json:"id,required"`
	// ID of the model the bid is for
	ModelID string `json:"modelID,required"`
	// Bid price per second
	PricePerSecond string  `json:"pricePerSecond,required"`
	JSON           bidJSON `json:"-"`
}

// bidJSON contains the JSON metadata for the struct [Bid]
type bidJSON struct {
	ID             apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Bid) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bidJSON) RawJSON() string {
	return r.raw
}

type BlockchainBidNewParams struct {
	// ID of the model to bid on
	ModelID param.Field[string] `json:"modelID,required"`
	// Bid price per second
	PricePerSecond param.Field[string] `json:"pricePerSecond,required"`
}

func (r BlockchainBidNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockchainBidSessionParams struct {
	// Duration of the session
	SessionDuration param.Field[string] `json:"sessionDuration,required"`
}

func (r BlockchainBidSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
