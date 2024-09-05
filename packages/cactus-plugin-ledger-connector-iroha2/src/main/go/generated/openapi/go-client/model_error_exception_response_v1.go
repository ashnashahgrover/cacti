/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the ErrorExceptionResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorExceptionResponseV1{}

// ErrorExceptionResponseV1 Error response from the connector.
type ErrorExceptionResponseV1 struct {
	// Short error description message.
	Message string `json:"message"`
	// Detailed error information.
	Error string `json:"error"`
}

// NewErrorExceptionResponseV1 instantiates a new ErrorExceptionResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorExceptionResponseV1(message string, error_ string) *ErrorExceptionResponseV1 {
	this := ErrorExceptionResponseV1{}
	this.Message = message
	this.Error = error_
	return &this
}

// NewErrorExceptionResponseV1WithDefaults instantiates a new ErrorExceptionResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorExceptionResponseV1WithDefaults() *ErrorExceptionResponseV1 {
	this := ErrorExceptionResponseV1{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorExceptionResponseV1) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorExceptionResponseV1) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorExceptionResponseV1) SetMessage(v string) {
	o.Message = v
}

// GetError returns the Error field value
func (o *ErrorExceptionResponseV1) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorExceptionResponseV1) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorExceptionResponseV1) SetError(v string) {
	o.Error = v
}

func (o ErrorExceptionResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorExceptionResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableErrorExceptionResponseV1 struct {
	value *ErrorExceptionResponseV1
	isSet bool
}

func (v NullableErrorExceptionResponseV1) Get() *ErrorExceptionResponseV1 {
	return v.value
}

func (v *NullableErrorExceptionResponseV1) Set(val *ErrorExceptionResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorExceptionResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorExceptionResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorExceptionResponseV1(val *ErrorExceptionResponseV1) *NullableErrorExceptionResponseV1 {
	return &NullableErrorExceptionResponseV1{value: val, isSet: true}
}

func (v NullableErrorExceptionResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorExceptionResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

