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

// checks if the WatchBlocksV1FullResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1FullResponse{}

// WatchBlocksV1FullResponse Response that returns entire sawtooth block.
type WatchBlocksV1FullResponse struct {
	FullBlock SawtoothBlockV1 `json:"fullBlock"`
}

// NewWatchBlocksV1FullResponse instantiates a new WatchBlocksV1FullResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1FullResponse(fullBlock SawtoothBlockV1) *WatchBlocksV1FullResponse {
	this := WatchBlocksV1FullResponse{}
	this.FullBlock = fullBlock
	return &this
}

// NewWatchBlocksV1FullResponseWithDefaults instantiates a new WatchBlocksV1FullResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1FullResponseWithDefaults() *WatchBlocksV1FullResponse {
	this := WatchBlocksV1FullResponse{}
	return &this
}

// GetFullBlock returns the FullBlock field value
func (o *WatchBlocksV1FullResponse) GetFullBlock() SawtoothBlockV1 {
	if o == nil {
		var ret SawtoothBlockV1
		return ret
	}

	return o.FullBlock
}

// GetFullBlockOk returns a tuple with the FullBlock field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1FullResponse) GetFullBlockOk() (*SawtoothBlockV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullBlock, true
}

// SetFullBlock sets field value
func (o *WatchBlocksV1FullResponse) SetFullBlock(v SawtoothBlockV1) {
	o.FullBlock = v
}

func (o WatchBlocksV1FullResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1FullResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullBlock"] = o.FullBlock
	return toSerialize, nil
}

type NullableWatchBlocksV1FullResponse struct {
	value *WatchBlocksV1FullResponse
	isSet bool
}

func (v NullableWatchBlocksV1FullResponse) Get() *WatchBlocksV1FullResponse {
	return v.value
}

func (v *NullableWatchBlocksV1FullResponse) Set(val *WatchBlocksV1FullResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1FullResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1FullResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1FullResponse(val *WatchBlocksV1FullResponse) *NullableWatchBlocksV1FullResponse {
	return &NullableWatchBlocksV1FullResponse{value: val, isSet: true}
}

func (v NullableWatchBlocksV1FullResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1FullResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


