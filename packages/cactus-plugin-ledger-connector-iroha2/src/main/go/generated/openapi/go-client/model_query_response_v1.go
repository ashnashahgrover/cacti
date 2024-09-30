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

// checks if the QueryResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryResponseV1{}

// QueryResponseV1 Response with the query results.
type QueryResponseV1 struct {
	// Query response data that varies between different queries.
	Response interface{} `json:"response"`
}

// NewQueryResponseV1 instantiates a new QueryResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryResponseV1(response interface{}) *QueryResponseV1 {
	this := QueryResponseV1{}
	this.Response = response
	return &this
}

// NewQueryResponseV1WithDefaults instantiates a new QueryResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryResponseV1WithDefaults() *QueryResponseV1 {
	this := QueryResponseV1{}
	return &this
}

// GetResponse returns the Response field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *QueryResponseV1) GetResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryResponseV1) GetResponseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *QueryResponseV1) SetResponse(v interface{}) {
	o.Response = v
}

func (o QueryResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableQueryResponseV1 struct {
	value *QueryResponseV1
	isSet bool
}

func (v NullableQueryResponseV1) Get() *QueryResponseV1 {
	return v.value
}

func (v *NullableQueryResponseV1) Set(val *QueryResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryResponseV1(val *QueryResponseV1) *NullableQueryResponseV1 {
	return &NullableQueryResponseV1{value: val, isSet: true}
}

func (v NullableQueryResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


