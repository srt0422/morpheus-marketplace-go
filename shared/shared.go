// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/srt0422/morpheus-marketplace-go/internal/apijson"
)

type Allowance struct {
	// Current allowance amount
	Allowance string        `json:"allowance,required"`
	JSON      allowanceJSON `json:"-"`
}

// allowanceJSON contains the JSON metadata for the struct [Allowance]
type allowanceJSON struct {
	Allowance   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Allowance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r allowanceJSON) RawJSON() string {
	return r.raw
}

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

type Bid struct {
	// Unique identifier of the bid
	ID string `json:"id,required"`
	// ID of the model the bid is for
	ModelID string `json:"modelID,required"`
	// Bid price per second
	PricePerSecond string  `json:"pricePerSecond,required"`
	JSON           bidJSON `json:"-"`
}

func (r *Bid) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// bidJSON contains the JSON metadata for the struct [Bid]
type bidJSON struct {
	ID             apijson.Field
	ModelID        apijson.Field
	PricePerSecond apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r bidJSON) RawJSON() string {
	return r.raw
}

type BidList struct {
	// List of bids
	Bids []Bid       `json:"bids,required"`
	JSON bidListJSON `json:"-"`
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

type LatestBlock struct {
	// Latest block number on the blockchain
	BlockNumber string          `json:"blockNumber,required"`
	JSON        latestBlockJSON `json:"-"`
}

// latestBlockJSON contains the JSON metadata for the struct [LatestBlock]
type latestBlockJSON struct {
	BlockNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LatestBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r latestBlockJSON) RawJSON() string {
	return r.raw
}

type Provider struct {
	// Unique identifier of the provider
	ID string `json:"id,required"`
	// Endpoint URL of the provider
	Endpoint string `json:"endpoint,required"`
	// Amount staked by the provider
	Stake string       `json:"stake,required"`
	JSON  providerJSON `json:"-"`
}

// providerJSON contains the JSON metadata for the struct [Provider]
type providerJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Stake       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Provider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerJSON) RawJSON() string {
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

type TokenSupply struct {
	// Total supply of the token
	Supply string          `json:"supply,required"`
	JSON   tokenSupplyJSON `json:"-"`
}

// tokenSupplyJSON contains the JSON metadata for the struct [TokenSupply]
type tokenSupplyJSON struct {
	Supply      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenSupply) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenSupplyJSON) RawJSON() string {
	return r.raw
}
