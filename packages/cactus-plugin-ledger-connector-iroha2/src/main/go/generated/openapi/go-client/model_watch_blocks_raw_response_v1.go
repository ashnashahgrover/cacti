/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the WatchBlocksRawResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksRawResponseV1{}

// WatchBlocksRawResponseV1 Default JSON-encoded string full block data.
type WatchBlocksRawResponseV1 struct {
	BlockData string `json:"blockData"`
}

// NewWatchBlocksRawResponseV1 instantiates a new WatchBlocksRawResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksRawResponseV1(blockData string) *WatchBlocksRawResponseV1 {
	this := WatchBlocksRawResponseV1{}
	this.BlockData = blockData
	return &this
}

// NewWatchBlocksRawResponseV1WithDefaults instantiates a new WatchBlocksRawResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksRawResponseV1WithDefaults() *WatchBlocksRawResponseV1 {
	this := WatchBlocksRawResponseV1{}
	return &this
}

// GetBlockData returns the BlockData field value
func (o *WatchBlocksRawResponseV1) GetBlockData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockData
}

// GetBlockDataOk returns a tuple with the BlockData field value
// and a boolean to check if the value has been set.
func (o *WatchBlocksRawResponseV1) GetBlockDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockData, true
}

// SetBlockData sets field value
func (o *WatchBlocksRawResponseV1) SetBlockData(v string) {
	o.BlockData = v
}

func (o WatchBlocksRawResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksRawResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blockData"] = o.BlockData
	return toSerialize, nil
}

type NullableWatchBlocksRawResponseV1 struct {
	value *WatchBlocksRawResponseV1
	isSet bool
}

func (v NullableWatchBlocksRawResponseV1) Get() *WatchBlocksRawResponseV1 {
	return v.value
}

func (v *NullableWatchBlocksRawResponseV1) Set(val *WatchBlocksRawResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksRawResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksRawResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksRawResponseV1(val *WatchBlocksRawResponseV1) *NullableWatchBlocksRawResponseV1 {
	return &NullableWatchBlocksRawResponseV1{value: val, isSet: true}
}

func (v NullableWatchBlocksRawResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksRawResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


