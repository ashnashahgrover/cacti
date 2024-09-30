/*
Hyperledger Cacti Plugin - Connector Aries

Can communicate with other Aries agents and Cacti Aries connectors

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-aries

import (
	"encoding/json"
)

// checks if the WatchProofStateProgressV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchProofStateProgressV1{}

// WatchProofStateProgressV1 Values pushed on each proof state change.
type WatchProofStateProgressV1 struct {
	ProofRecord AriesProofExchangeRecordV1 `json:"proofRecord"`
	PreviousState NullableString `json:"previousState"`
}

// NewWatchProofStateProgressV1 instantiates a new WatchProofStateProgressV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchProofStateProgressV1(proofRecord AriesProofExchangeRecordV1, previousState NullableString) *WatchProofStateProgressV1 {
	this := WatchProofStateProgressV1{}
	this.ProofRecord = proofRecord
	this.PreviousState = previousState
	return &this
}

// NewWatchProofStateProgressV1WithDefaults instantiates a new WatchProofStateProgressV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchProofStateProgressV1WithDefaults() *WatchProofStateProgressV1 {
	this := WatchProofStateProgressV1{}
	return &this
}

// GetProofRecord returns the ProofRecord field value
func (o *WatchProofStateProgressV1) GetProofRecord() AriesProofExchangeRecordV1 {
	if o == nil {
		var ret AriesProofExchangeRecordV1
		return ret
	}

	return o.ProofRecord
}

// GetProofRecordOk returns a tuple with the ProofRecord field value
// and a boolean to check if the value has been set.
func (o *WatchProofStateProgressV1) GetProofRecordOk() (*AriesProofExchangeRecordV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProofRecord, true
}

// SetProofRecord sets field value
func (o *WatchProofStateProgressV1) SetProofRecord(v AriesProofExchangeRecordV1) {
	o.ProofRecord = v
}

// GetPreviousState returns the PreviousState field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchProofStateProgressV1) GetPreviousState() string {
	if o == nil || o.PreviousState.Get() == nil {
		var ret string
		return ret
	}

	return *o.PreviousState.Get()
}

// GetPreviousStateOk returns a tuple with the PreviousState field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchProofStateProgressV1) GetPreviousStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousState.Get(), o.PreviousState.IsSet()
}

// SetPreviousState sets field value
func (o *WatchProofStateProgressV1) SetPreviousState(v string) {
	o.PreviousState.Set(&v)
}

func (o WatchProofStateProgressV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchProofStateProgressV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["proofRecord"] = o.ProofRecord
	toSerialize["previousState"] = o.PreviousState.Get()
	return toSerialize, nil
}

type NullableWatchProofStateProgressV1 struct {
	value *WatchProofStateProgressV1
	isSet bool
}

func (v NullableWatchProofStateProgressV1) Get() *WatchProofStateProgressV1 {
	return v.value
}

func (v *NullableWatchProofStateProgressV1) Set(val *WatchProofStateProgressV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchProofStateProgressV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchProofStateProgressV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchProofStateProgressV1(val *WatchProofStateProgressV1) *NullableWatchProofStateProgressV1 {
	return &NullableWatchProofStateProgressV1{value: val, isSet: true}
}

func (v NullableWatchProofStateProgressV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchProofStateProgressV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


