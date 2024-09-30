/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
)

// checks if the WatchBlocksV1CactiTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1CactiTransactionsResponse{}

// WatchBlocksV1CactiTransactionsResponse Custom response containing block transactions summary.
type WatchBlocksV1CactiTransactionsResponse struct {
	// List of sawtooth transactions matching specifid (optional) filter
	CactiTransactionsEvents []CactiTransactionV1 `json:"cactiTransactionsEvents"`
}

// NewWatchBlocksV1CactiTransactionsResponse instantiates a new WatchBlocksV1CactiTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1CactiTransactionsResponse(cactiTransactionsEvents []CactiTransactionV1) *WatchBlocksV1CactiTransactionsResponse {
	this := WatchBlocksV1CactiTransactionsResponse{}
	this.CactiTransactionsEvents = cactiTransactionsEvents
	return &this
}

// NewWatchBlocksV1CactiTransactionsResponseWithDefaults instantiates a new WatchBlocksV1CactiTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1CactiTransactionsResponseWithDefaults() *WatchBlocksV1CactiTransactionsResponse {
	this := WatchBlocksV1CactiTransactionsResponse{}
	return &this
}

// GetCactiTransactionsEvents returns the CactiTransactionsEvents field value
func (o *WatchBlocksV1CactiTransactionsResponse) GetCactiTransactionsEvents() []CactiTransactionV1 {
	if o == nil {
		var ret []CactiTransactionV1
		return ret
	}

	return o.CactiTransactionsEvents
}

// GetCactiTransactionsEventsOk returns a tuple with the CactiTransactionsEvents field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1CactiTransactionsResponse) GetCactiTransactionsEventsOk() ([]CactiTransactionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.CactiTransactionsEvents, true
}

// SetCactiTransactionsEvents sets field value
func (o *WatchBlocksV1CactiTransactionsResponse) SetCactiTransactionsEvents(v []CactiTransactionV1) {
	o.CactiTransactionsEvents = v
}

func (o WatchBlocksV1CactiTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1CactiTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cactiTransactionsEvents"] = o.CactiTransactionsEvents
	return toSerialize, nil
}

type NullableWatchBlocksV1CactiTransactionsResponse struct {
	value *WatchBlocksV1CactiTransactionsResponse
	isSet bool
}

func (v NullableWatchBlocksV1CactiTransactionsResponse) Get() *WatchBlocksV1CactiTransactionsResponse {
	return v.value
}

func (v *NullableWatchBlocksV1CactiTransactionsResponse) Set(val *WatchBlocksV1CactiTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1CactiTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1CactiTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1CactiTransactionsResponse(val *WatchBlocksV1CactiTransactionsResponse) *NullableWatchBlocksV1CactiTransactionsResponse {
	return &NullableWatchBlocksV1CactiTransactionsResponse{value: val, isSet: true}
}

func (v NullableWatchBlocksV1CactiTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1CactiTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


