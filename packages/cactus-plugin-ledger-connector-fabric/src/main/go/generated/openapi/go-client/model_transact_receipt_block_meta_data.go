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

// checks if the TransactReceiptBlockMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactReceiptBlockMetaData{}

// TransactReceiptBlockMetaData struct for TransactReceiptBlockMetaData
type TransactReceiptBlockMetaData struct {
	Mspid *string `json:"mspid,omitempty"`
	BlockCreatorID *string `json:"blockCreatorID,omitempty"`
	Signature *string `json:"signature,omitempty"`
}

// NewTransactReceiptBlockMetaData instantiates a new TransactReceiptBlockMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactReceiptBlockMetaData() *TransactReceiptBlockMetaData {
	this := TransactReceiptBlockMetaData{}
	return &this
}

// NewTransactReceiptBlockMetaDataWithDefaults instantiates a new TransactReceiptBlockMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactReceiptBlockMetaDataWithDefaults() *TransactReceiptBlockMetaData {
	this := TransactReceiptBlockMetaData{}
	return &this
}

// GetMspid returns the Mspid field value if set, zero value otherwise.
func (o *TransactReceiptBlockMetaData) GetMspid() string {
	if o == nil || IsNil(o.Mspid) {
		var ret string
		return ret
	}
	return *o.Mspid
}

// GetMspidOk returns a tuple with the Mspid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptBlockMetaData) GetMspidOk() (*string, bool) {
	if o == nil || IsNil(o.Mspid) {
		return nil, false
	}
	return o.Mspid, true
}

// HasMspid returns a boolean if a field has been set.
func (o *TransactReceiptBlockMetaData) HasMspid() bool {
	if o != nil && !IsNil(o.Mspid) {
		return true
	}

	return false
}

// SetMspid gets a reference to the given string and assigns it to the Mspid field.
func (o *TransactReceiptBlockMetaData) SetMspid(v string) {
	o.Mspid = &v
}

// GetBlockCreatorID returns the BlockCreatorID field value if set, zero value otherwise.
func (o *TransactReceiptBlockMetaData) GetBlockCreatorID() string {
	if o == nil || IsNil(o.BlockCreatorID) {
		var ret string
		return ret
	}
	return *o.BlockCreatorID
}

// GetBlockCreatorIDOk returns a tuple with the BlockCreatorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptBlockMetaData) GetBlockCreatorIDOk() (*string, bool) {
	if o == nil || IsNil(o.BlockCreatorID) {
		return nil, false
	}
	return o.BlockCreatorID, true
}

// HasBlockCreatorID returns a boolean if a field has been set.
func (o *TransactReceiptBlockMetaData) HasBlockCreatorID() bool {
	if o != nil && !IsNil(o.BlockCreatorID) {
		return true
	}

	return false
}

// SetBlockCreatorID gets a reference to the given string and assigns it to the BlockCreatorID field.
func (o *TransactReceiptBlockMetaData) SetBlockCreatorID(v string) {
	o.BlockCreatorID = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *TransactReceiptBlockMetaData) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptBlockMetaData) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *TransactReceiptBlockMetaData) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *TransactReceiptBlockMetaData) SetSignature(v string) {
	o.Signature = &v
}

func (o TransactReceiptBlockMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactReceiptBlockMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mspid) {
		toSerialize["mspid"] = o.Mspid
	}
	if !IsNil(o.BlockCreatorID) {
		toSerialize["blockCreatorID"] = o.BlockCreatorID
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	return toSerialize, nil
}

type NullableTransactReceiptBlockMetaData struct {
	value *TransactReceiptBlockMetaData
	isSet bool
}

func (v NullableTransactReceiptBlockMetaData) Get() *TransactReceiptBlockMetaData {
	return v.value
}

func (v *NullableTransactReceiptBlockMetaData) Set(val *TransactReceiptBlockMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactReceiptBlockMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactReceiptBlockMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactReceiptBlockMetaData(val *TransactReceiptBlockMetaData) *NullableTransactReceiptBlockMetaData {
	return &NullableTransactReceiptBlockMetaData{value: val, isSet: true}
}

func (v NullableTransactReceiptBlockMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactReceiptBlockMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


