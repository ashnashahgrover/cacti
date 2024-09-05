/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the SignTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignTransactionResponse{}

// SignTransactionResponse struct for SignTransactionResponse
type SignTransactionResponse struct {
	// The signatures of ledger from the corresponding transaction hash.
	Signature string `json:"signature"`
}

// NewSignTransactionResponse instantiates a new SignTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignTransactionResponse(signature string) *SignTransactionResponse {
	this := SignTransactionResponse{}
	this.Signature = signature
	return &this
}

// NewSignTransactionResponseWithDefaults instantiates a new SignTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignTransactionResponseWithDefaults() *SignTransactionResponse {
	this := SignTransactionResponse{}
	return &this
}

// GetSignature returns the Signature field value
func (o *SignTransactionResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *SignTransactionResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *SignTransactionResponse) SetSignature(v string) {
	o.Signature = v
}

func (o SignTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableSignTransactionResponse struct {
	value *SignTransactionResponse
	isSet bool
}

func (v NullableSignTransactionResponse) Get() *SignTransactionResponse {
	return v.value
}

func (v *NullableSignTransactionResponse) Set(val *SignTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignTransactionResponse(val *SignTransactionResponse) *NullableSignTransactionResponse {
	return &NullableSignTransactionResponse{value: val, isSet: true}
}

func (v NullableSignTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

