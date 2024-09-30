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

// checks if the ClientV1RequestClientGatewayConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientV1RequestClientGatewayConfiguration{}

// ClientV1RequestClientGatewayConfiguration struct for ClientV1RequestClientGatewayConfiguration
type ClientV1RequestClientGatewayConfiguration struct {
	ApiHost string `json:"apiHost"`
}

// NewClientV1RequestClientGatewayConfiguration instantiates a new ClientV1RequestClientGatewayConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientV1RequestClientGatewayConfiguration(apiHost string) *ClientV1RequestClientGatewayConfiguration {
	this := ClientV1RequestClientGatewayConfiguration{}
	this.ApiHost = apiHost
	return &this
}

// NewClientV1RequestClientGatewayConfigurationWithDefaults instantiates a new ClientV1RequestClientGatewayConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientV1RequestClientGatewayConfigurationWithDefaults() *ClientV1RequestClientGatewayConfiguration {
	this := ClientV1RequestClientGatewayConfiguration{}
	return &this
}

// GetApiHost returns the ApiHost field value
func (o *ClientV1RequestClientGatewayConfiguration) GetApiHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiHost
}

// GetApiHostOk returns a tuple with the ApiHost field value
// and a boolean to check if the value has been set.
func (o *ClientV1RequestClientGatewayConfiguration) GetApiHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiHost, true
}

// SetApiHost sets field value
func (o *ClientV1RequestClientGatewayConfiguration) SetApiHost(v string) {
	o.ApiHost = v
}

func (o ClientV1RequestClientGatewayConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientV1RequestClientGatewayConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiHost"] = o.ApiHost
	return toSerialize, nil
}

type NullableClientV1RequestClientGatewayConfiguration struct {
	value *ClientV1RequestClientGatewayConfiguration
	isSet bool
}

func (v NullableClientV1RequestClientGatewayConfiguration) Get() *ClientV1RequestClientGatewayConfiguration {
	return v.value
}

func (v *NullableClientV1RequestClientGatewayConfiguration) Set(val *ClientV1RequestClientGatewayConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableClientV1RequestClientGatewayConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableClientV1RequestClientGatewayConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientV1RequestClientGatewayConfiguration(val *ClientV1RequestClientGatewayConfiguration) *NullableClientV1RequestClientGatewayConfiguration {
	return &NullableClientV1RequestClientGatewayConfiguration{value: val, isSet: true}
}

func (v NullableClientV1RequestClientGatewayConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientV1RequestClientGatewayConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


