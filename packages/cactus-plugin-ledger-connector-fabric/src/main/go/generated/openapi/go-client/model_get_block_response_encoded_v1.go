/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the GetBlockResponseEncodedV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBlockResponseEncodedV1{}

// GetBlockResponseEncodedV1 When skipDecode is true then encoded block Buffer is returned.
type GetBlockResponseEncodedV1 struct {
	EncodedBlock string `json:"encodedBlock"`
}

// NewGetBlockResponseEncodedV1 instantiates a new GetBlockResponseEncodedV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockResponseEncodedV1(encodedBlock string) *GetBlockResponseEncodedV1 {
	this := GetBlockResponseEncodedV1{}
	this.EncodedBlock = encodedBlock
	return &this
}

// NewGetBlockResponseEncodedV1WithDefaults instantiates a new GetBlockResponseEncodedV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockResponseEncodedV1WithDefaults() *GetBlockResponseEncodedV1 {
	this := GetBlockResponseEncodedV1{}
	return &this
}

// GetEncodedBlock returns the EncodedBlock field value
func (o *GetBlockResponseEncodedV1) GetEncodedBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncodedBlock
}

// GetEncodedBlockOk returns a tuple with the EncodedBlock field value
// and a boolean to check if the value has been set.
func (o *GetBlockResponseEncodedV1) GetEncodedBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncodedBlock, true
}

// SetEncodedBlock sets field value
func (o *GetBlockResponseEncodedV1) SetEncodedBlock(v string) {
	o.EncodedBlock = v
}

func (o GetBlockResponseEncodedV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBlockResponseEncodedV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encodedBlock"] = o.EncodedBlock
	return toSerialize, nil
}

type NullableGetBlockResponseEncodedV1 struct {
	value *GetBlockResponseEncodedV1
	isSet bool
}

func (v NullableGetBlockResponseEncodedV1) Get() *GetBlockResponseEncodedV1 {
	return v.value
}

func (v *NullableGetBlockResponseEncodedV1) Set(val *GetBlockResponseEncodedV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockResponseEncodedV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockResponseEncodedV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockResponseEncodedV1(val *GetBlockResponseEncodedV1) *NullableGetBlockResponseEncodedV1 {
	return &NullableGetBlockResponseEncodedV1{value: val, isSet: true}
}

func (v NullableGetBlockResponseEncodedV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockResponseEncodedV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


