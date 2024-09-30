/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the PayloadProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayloadProfile{}

// PayloadProfile struct for PayloadProfile
type PayloadProfile struct {
	AssetProfile AssetProfile `json:"assetProfile"`
	Capabilities *string `json:"capabilities,omitempty"`
}

// NewPayloadProfile instantiates a new PayloadProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadProfile(assetProfile AssetProfile) *PayloadProfile {
	this := PayloadProfile{}
	this.AssetProfile = assetProfile
	return &this
}

// NewPayloadProfileWithDefaults instantiates a new PayloadProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadProfileWithDefaults() *PayloadProfile {
	this := PayloadProfile{}
	return &this
}

// GetAssetProfile returns the AssetProfile field value
func (o *PayloadProfile) GetAssetProfile() AssetProfile {
	if o == nil {
		var ret AssetProfile
		return ret
	}

	return o.AssetProfile
}

// GetAssetProfileOk returns a tuple with the AssetProfile field value
// and a boolean to check if the value has been set.
func (o *PayloadProfile) GetAssetProfileOk() (*AssetProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetProfile, true
}

// SetAssetProfile sets field value
func (o *PayloadProfile) SetAssetProfile(v AssetProfile) {
	o.AssetProfile = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *PayloadProfile) GetCapabilities() string {
	if o == nil || IsNil(o.Capabilities) {
		var ret string
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadProfile) GetCapabilitiesOk() (*string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *PayloadProfile) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given string and assigns it to the Capabilities field.
func (o *PayloadProfile) SetCapabilities(v string) {
	o.Capabilities = &v
}

func (o PayloadProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayloadProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetProfile"] = o.AssetProfile
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullablePayloadProfile struct {
	value *PayloadProfile
	isSet bool
}

func (v NullablePayloadProfile) Get() *PayloadProfile {
	return v.value
}

func (v *NullablePayloadProfile) Set(val *PayloadProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadProfile(val *PayloadProfile) *NullablePayloadProfile {
	return &NullablePayloadProfile{value: val, isSet: true}
}

func (v NullablePayloadProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


