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

// checks if the WatchBlocksOptionsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksOptionsV1{}

// WatchBlocksOptionsV1 Options passed when subscribing to block monitoring.
type WatchBlocksOptionsV1 struct {
	Type *BlockTypeV1 `json:"type,omitempty"`
	// Number of block to start monitoring from.
	StartBlock *string `json:"startBlock,omitempty"`
	BaseConfig *Iroha2BaseConfig `json:"baseConfig,omitempty"`
}

// NewWatchBlocksOptionsV1 instantiates a new WatchBlocksOptionsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksOptionsV1() *WatchBlocksOptionsV1 {
	this := WatchBlocksOptionsV1{}
	return &this
}

// NewWatchBlocksOptionsV1WithDefaults instantiates a new WatchBlocksOptionsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksOptionsV1WithDefaults() *WatchBlocksOptionsV1 {
	this := WatchBlocksOptionsV1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WatchBlocksOptionsV1) GetType() BlockTypeV1 {
	if o == nil || IsNil(o.Type) {
		var ret BlockTypeV1
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetTypeOk() (*BlockTypeV1, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WatchBlocksOptionsV1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BlockTypeV1 and assigns it to the Type field.
func (o *WatchBlocksOptionsV1) SetType(v BlockTypeV1) {
	o.Type = &v
}

// GetStartBlock returns the StartBlock field value if set, zero value otherwise.
func (o *WatchBlocksOptionsV1) GetStartBlock() string {
	if o == nil || IsNil(o.StartBlock) {
		var ret string
		return ret
	}
	return *o.StartBlock
}

// GetStartBlockOk returns a tuple with the StartBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetStartBlockOk() (*string, bool) {
	if o == nil || IsNil(o.StartBlock) {
		return nil, false
	}
	return o.StartBlock, true
}

// HasStartBlock returns a boolean if a field has been set.
func (o *WatchBlocksOptionsV1) HasStartBlock() bool {
	if o != nil && !IsNil(o.StartBlock) {
		return true
	}

	return false
}

// SetStartBlock gets a reference to the given string and assigns it to the StartBlock field.
func (o *WatchBlocksOptionsV1) SetStartBlock(v string) {
	o.StartBlock = &v
}

// GetBaseConfig returns the BaseConfig field value if set, zero value otherwise.
func (o *WatchBlocksOptionsV1) GetBaseConfig() Iroha2BaseConfig {
	if o == nil || IsNil(o.BaseConfig) {
		var ret Iroha2BaseConfig
		return ret
	}
	return *o.BaseConfig
}

// GetBaseConfigOk returns a tuple with the BaseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksOptionsV1) GetBaseConfigOk() (*Iroha2BaseConfig, bool) {
	if o == nil || IsNil(o.BaseConfig) {
		return nil, false
	}
	return o.BaseConfig, true
}

// HasBaseConfig returns a boolean if a field has been set.
func (o *WatchBlocksOptionsV1) HasBaseConfig() bool {
	if o != nil && !IsNil(o.BaseConfig) {
		return true
	}

	return false
}

// SetBaseConfig gets a reference to the given Iroha2BaseConfig and assigns it to the BaseConfig field.
func (o *WatchBlocksOptionsV1) SetBaseConfig(v Iroha2BaseConfig) {
	o.BaseConfig = &v
}

func (o WatchBlocksOptionsV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksOptionsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.StartBlock) {
		toSerialize["startBlock"] = o.StartBlock
	}
	if !IsNil(o.BaseConfig) {
		toSerialize["baseConfig"] = o.BaseConfig
	}
	return toSerialize, nil
}

type NullableWatchBlocksOptionsV1 struct {
	value *WatchBlocksOptionsV1
	isSet bool
}

func (v NullableWatchBlocksOptionsV1) Get() *WatchBlocksOptionsV1 {
	return v.value
}

func (v *NullableWatchBlocksOptionsV1) Set(val *WatchBlocksOptionsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksOptionsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksOptionsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksOptionsV1(val *WatchBlocksOptionsV1) *NullableWatchBlocksOptionsV1 {
	return &NullableWatchBlocksOptionsV1{value: val, isSet: true}
}

func (v NullableWatchBlocksOptionsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksOptionsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


