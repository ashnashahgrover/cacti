/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
)

// checks if the CactiTransactionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CactiTransactionV1{}

// CactiTransactionV1 Sawtooth transaction with additional fields filled by Cacti connector.
type CactiTransactionV1 struct {
	Header SawtoothTransactionHeaderV1 `json:"header"`
	HeaderSignature string `json:"header_signature"`
	Payload string `json:"payload"`
	// Decoded payload of sawtooth transaction.
	PayloadDecoded interface{} `json:"payload_decoded"`
}

// NewCactiTransactionV1 instantiates a new CactiTransactionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCactiTransactionV1(header SawtoothTransactionHeaderV1, headerSignature string, payload string, payloadDecoded interface{}) *CactiTransactionV1 {
	this := CactiTransactionV1{}
	this.Header = header
	this.HeaderSignature = headerSignature
	this.Payload = payload
	this.PayloadDecoded = payloadDecoded
	return &this
}

// NewCactiTransactionV1WithDefaults instantiates a new CactiTransactionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCactiTransactionV1WithDefaults() *CactiTransactionV1 {
	this := CactiTransactionV1{}
	return &this
}

// GetHeader returns the Header field value
func (o *CactiTransactionV1) GetHeader() SawtoothTransactionHeaderV1 {
	if o == nil {
		var ret SawtoothTransactionHeaderV1
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *CactiTransactionV1) GetHeaderOk() (*SawtoothTransactionHeaderV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *CactiTransactionV1) SetHeader(v SawtoothTransactionHeaderV1) {
	o.Header = v
}

// GetHeaderSignature returns the HeaderSignature field value
func (o *CactiTransactionV1) GetHeaderSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderSignature
}

// GetHeaderSignatureOk returns a tuple with the HeaderSignature field value
// and a boolean to check if the value has been set.
func (o *CactiTransactionV1) GetHeaderSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderSignature, true
}

// SetHeaderSignature sets field value
func (o *CactiTransactionV1) SetHeaderSignature(v string) {
	o.HeaderSignature = v
}

// GetPayload returns the Payload field value
func (o *CactiTransactionV1) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *CactiTransactionV1) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *CactiTransactionV1) SetPayload(v string) {
	o.Payload = v
}

// GetPayloadDecoded returns the PayloadDecoded field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CactiTransactionV1) GetPayloadDecoded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PayloadDecoded
}

// GetPayloadDecodedOk returns a tuple with the PayloadDecoded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CactiTransactionV1) GetPayloadDecodedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PayloadDecoded) {
		return nil, false
	}
	return &o.PayloadDecoded, true
}

// SetPayloadDecoded sets field value
func (o *CactiTransactionV1) SetPayloadDecoded(v interface{}) {
	o.PayloadDecoded = v
}

func (o CactiTransactionV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CactiTransactionV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["header"] = o.Header
	toSerialize["header_signature"] = o.HeaderSignature
	toSerialize["payload"] = o.Payload
	if o.PayloadDecoded != nil {
		toSerialize["payload_decoded"] = o.PayloadDecoded
	}
	return toSerialize, nil
}

type NullableCactiTransactionV1 struct {
	value *CactiTransactionV1
	isSet bool
}

func (v NullableCactiTransactionV1) Get() *CactiTransactionV1 {
	return v.value
}

func (v *NullableCactiTransactionV1) Set(val *CactiTransactionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCactiTransactionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCactiTransactionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCactiTransactionV1(val *CactiTransactionV1) *NullableCactiTransactionV1 {
	return &NullableCactiTransactionV1{value: val, isSet: true}
}

func (v NullableCactiTransactionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCactiTransactionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


