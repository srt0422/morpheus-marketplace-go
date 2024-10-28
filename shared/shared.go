// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
)

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

type SessionList struct {
	// List of sessions
	Sessions []Session       `json:"sessions,required"`
	JSON     sessionListJSON `json:"-"`
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
