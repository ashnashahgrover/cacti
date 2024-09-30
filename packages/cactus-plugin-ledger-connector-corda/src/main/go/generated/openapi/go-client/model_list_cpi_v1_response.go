/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the ListCpiV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCpiV1Response{}

// ListCpiV1Response struct for ListCpiV1Response
type ListCpiV1Response struct {
	Cpis []ListCpiV1ResponseCpisInner `json:"cpis,omitempty"`
}

// NewListCpiV1Response instantiates a new ListCpiV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCpiV1Response() *ListCpiV1Response {
	this := ListCpiV1Response{}
	return &this
}

// NewListCpiV1ResponseWithDefaults instantiates a new ListCpiV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCpiV1ResponseWithDefaults() *ListCpiV1Response {
	this := ListCpiV1Response{}
	return &this
}

// GetCpis returns the Cpis field value if set, zero value otherwise.
func (o *ListCpiV1Response) GetCpis() []ListCpiV1ResponseCpisInner {
	if o == nil || IsNil(o.Cpis) {
		var ret []ListCpiV1ResponseCpisInner
		return ret
	}
	return o.Cpis
}

// GetCpisOk returns a tuple with the Cpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1Response) GetCpisOk() ([]ListCpiV1ResponseCpisInner, bool) {
	if o == nil || IsNil(o.Cpis) {
		return nil, false
	}
	return o.Cpis, true
}

// HasCpis returns a boolean if a field has been set.
func (o *ListCpiV1Response) HasCpis() bool {
	if o != nil && !IsNil(o.Cpis) {
		return true
	}

	return false
}

// SetCpis gets a reference to the given []ListCpiV1ResponseCpisInner and assigns it to the Cpis field.
func (o *ListCpiV1Response) SetCpis(v []ListCpiV1ResponseCpisInner) {
	o.Cpis = v
}

func (o ListCpiV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCpiV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpis) {
		toSerialize["cpis"] = o.Cpis
	}
	return toSerialize, nil
}

type NullableListCpiV1Response struct {
	value *ListCpiV1Response
	isSet bool
}

func (v NullableListCpiV1Response) Get() *ListCpiV1Response {
	return v.value
}

func (v *NullableListCpiV1Response) Set(val *ListCpiV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListCpiV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListCpiV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCpiV1Response(val *ListCpiV1Response) *NullableListCpiV1Response {
	return &NullableListCpiV1Response{value: val, isSet: true}
}

func (v NullableListCpiV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCpiV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


