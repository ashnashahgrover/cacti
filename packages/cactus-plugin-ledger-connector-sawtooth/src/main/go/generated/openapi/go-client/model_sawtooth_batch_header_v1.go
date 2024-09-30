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

// checks if the SawtoothBatchHeaderV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SawtoothBatchHeaderV1{}

// SawtoothBatchHeaderV1 struct for SawtoothBatchHeaderV1
type SawtoothBatchHeaderV1 struct {
	SignerPublicKey string `json:"signer_public_key"`
	TransactionIds []string `json:"transaction_ids"`
}

// NewSawtoothBatchHeaderV1 instantiates a new SawtoothBatchHeaderV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSawtoothBatchHeaderV1(signerPublicKey string, transactionIds []string) *SawtoothBatchHeaderV1 {
	this := SawtoothBatchHeaderV1{}
	this.SignerPublicKey = signerPublicKey
	this.TransactionIds = transactionIds
	return &this
}

// NewSawtoothBatchHeaderV1WithDefaults instantiates a new SawtoothBatchHeaderV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSawtoothBatchHeaderV1WithDefaults() *SawtoothBatchHeaderV1 {
	this := SawtoothBatchHeaderV1{}
	return &this
}

// GetSignerPublicKey returns the SignerPublicKey field value
func (o *SawtoothBatchHeaderV1) GetSignerPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerPublicKey
}

// GetSignerPublicKeyOk returns a tuple with the SignerPublicKey field value
// and a boolean to check if the value has been set.
func (o *SawtoothBatchHeaderV1) GetSignerPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerPublicKey, true
}

// SetSignerPublicKey sets field value
func (o *SawtoothBatchHeaderV1) SetSignerPublicKey(v string) {
	o.SignerPublicKey = v
}

// GetTransactionIds returns the TransactionIds field value
func (o *SawtoothBatchHeaderV1) GetTransactionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TransactionIds
}

// GetTransactionIdsOk returns a tuple with the TransactionIds field value
// and a boolean to check if the value has been set.
func (o *SawtoothBatchHeaderV1) GetTransactionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionIds, true
}

// SetTransactionIds sets field value
func (o *SawtoothBatchHeaderV1) SetTransactionIds(v []string) {
	o.TransactionIds = v
}

func (o SawtoothBatchHeaderV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SawtoothBatchHeaderV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_public_key"] = o.SignerPublicKey
	toSerialize["transaction_ids"] = o.TransactionIds
	return toSerialize, nil
}

type NullableSawtoothBatchHeaderV1 struct {
	value *SawtoothBatchHeaderV1
	isSet bool
}

func (v NullableSawtoothBatchHeaderV1) Get() *SawtoothBatchHeaderV1 {
	return v.value
}

func (v *NullableSawtoothBatchHeaderV1) Set(val *SawtoothBatchHeaderV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSawtoothBatchHeaderV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSawtoothBatchHeaderV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSawtoothBatchHeaderV1(val *SawtoothBatchHeaderV1) *NullableSawtoothBatchHeaderV1 {
	return &NullableSawtoothBatchHeaderV1{value: val, isSet: true}
}

func (v NullableSawtoothBatchHeaderV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSawtoothBatchHeaderV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


