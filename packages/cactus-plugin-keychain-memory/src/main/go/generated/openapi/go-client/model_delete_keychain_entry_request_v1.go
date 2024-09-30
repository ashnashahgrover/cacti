/*
Hyperledger Cactus Plugin - Keychain Memory 

Contains/describes the Hyperledger Cacti Keychain Memory plugin.

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-memory

import (
	"encoding/json"
)

// checks if the DeleteKeychainEntryRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteKeychainEntryRequestV1{}

// DeleteKeychainEntryRequestV1 struct for DeleteKeychainEntryRequestV1
type DeleteKeychainEntryRequestV1 struct {
	// The key for the entry to check the presence of on the keychain.
	Key string `json:"key"`
}

// NewDeleteKeychainEntryRequestV1 instantiates a new DeleteKeychainEntryRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteKeychainEntryRequestV1(key string) *DeleteKeychainEntryRequestV1 {
	this := DeleteKeychainEntryRequestV1{}
	this.Key = key
	return &this
}

// NewDeleteKeychainEntryRequestV1WithDefaults instantiates a new DeleteKeychainEntryRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteKeychainEntryRequestV1WithDefaults() *DeleteKeychainEntryRequestV1 {
	this := DeleteKeychainEntryRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *DeleteKeychainEntryRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeleteKeychainEntryRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeleteKeychainEntryRequestV1) SetKey(v string) {
	o.Key = v
}

func (o DeleteKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteKeychainEntryRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableDeleteKeychainEntryRequestV1 struct {
	value *DeleteKeychainEntryRequestV1
	isSet bool
}

func (v NullableDeleteKeychainEntryRequestV1) Get() *DeleteKeychainEntryRequestV1 {
	return v.value
}

func (v *NullableDeleteKeychainEntryRequestV1) Set(val *DeleteKeychainEntryRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteKeychainEntryRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteKeychainEntryRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteKeychainEntryRequestV1(val *DeleteKeychainEntryRequestV1) *NullableDeleteKeychainEntryRequestV1 {
	return &NullableDeleteKeychainEntryRequestV1{value: val, isSet: true}
}

func (v NullableDeleteKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteKeychainEntryRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


