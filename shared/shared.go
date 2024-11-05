// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/srt0422/morpheus-marketplace-go"
	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
)

type Balance struct {
	// Current balance after the transaction
	Balance string      `json:"balance,required"`
	JSON    balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	Balance     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
}

type BidList struct {
	// List of bids
	Bids []morpheusmarketplace.Bid `json:"bids,required"`
	JSON bidListJSON               `json:"-"`
}

// bidListJSON contains the JSON metadata for the struct [BidList]
type bidListJSON struct {
	Bids        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BidList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bidListJSON) RawJSON() string {
	return r.raw
}

type Session struct {
	// Unique identifier of the session
	ID string `json:"id,required"`
	// Duration of the session in seconds
	SessionDuration string `json:"sessionDuration,required"`
	// Status of the session
	Status string      `json:"status,required"`
	JSON   sessionJSON `json:"-"`
}

// sessionJSON contains the JSON metadata for the struct [Session]
type sessionJSON struct {
	ID              apijson.Field
	SessionDuration apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Session) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionJSON) RawJSON() string {
	return r.raw
}
