// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/apijson"
)

type Bid []BidItem

type BidItem struct {
	BidID          string      `json:"bidID"`
	ModelID        string      `json:"modelID"`
	PricePerSecond string      `json:"pricePerSecond" format:"biginteger"`
	ProviderID     string      `json:"providerID"`
	Status         string      `json:"status"`
	JSON           bidItemJSON `json:"-"`
}

// bidItemJSON contains the JSON metadata for the struct [BidItem]
type bidItemJSON struct {
	BidID          apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	ProviderID     apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BidItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bidItemJSON) RawJSON() string {
	return r.raw
}

type Session struct {
	// Additional session details.
	Details interface{} `json:"details"`
	// Unique identifier for the session.
	SessionID string      `json:"sessionID"`
	JSON      sessionJSON `json:"-"`
}

// sessionJSON contains the JSON metadata for the struct [Session]
type sessionJSON struct {
	Details     apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Session) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionJSON) RawJSON() string {
	return r.raw
}
