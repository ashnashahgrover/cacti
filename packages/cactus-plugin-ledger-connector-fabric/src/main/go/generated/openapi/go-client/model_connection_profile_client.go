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

// checks if the ConnectionProfileClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionProfileClient{}

// ConnectionProfileClient struct for ConnectionProfileClient
type ConnectionProfileClient struct {
	Organization *string `json:"organization,omitempty"`
}

// NewConnectionProfileClient instantiates a new ConnectionProfileClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionProfileClient() *ConnectionProfileClient {
	this := ConnectionProfileClient{}
	return &this
}

// NewConnectionProfileClientWithDefaults instantiates a new ConnectionProfileClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionProfileClientWithDefaults() *ConnectionProfileClient {
	this := ConnectionProfileClient{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConnectionProfileClient) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionProfileClient) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConnectionProfileClient) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *ConnectionProfileClient) SetOrganization(v string) {
	o.Organization = &v
}

func (o ConnectionProfileClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionProfileClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableConnectionProfileClient struct {
	value *ConnectionProfileClient
	isSet bool
}

func (v NullableConnectionProfileClient) Get() *ConnectionProfileClient {
	return v.value
}

func (v *NullableConnectionProfileClient) Set(val *ConnectionProfileClient) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionProfileClient) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionProfileClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionProfileClient(val *ConnectionProfileClient) *NullableConnectionProfileClient {
	return &NullableConnectionProfileClient{value: val, isSet: true}
}

func (v NullableConnectionProfileClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionProfileClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


