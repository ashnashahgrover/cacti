/*
Hyperledger Cactus Plugin - Persistence Ethereum

Synchronizes state of an ethereum ledger into a DB that can later be viewed in GUI

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-persistence-ethereum

import (
	"encoding/json"
)

// checks if the TrackedOperationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackedOperationV1{}

// TrackedOperationV1 Persistence plugin operation that is tracked and returned in status report.
type TrackedOperationV1 struct {
	// Start time of the operation.
	StartAt string `json:"startAt"`
	// Operation name.
	Operation string `json:"operation"`
}

// NewTrackedOperationV1 instantiates a new TrackedOperationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackedOperationV1(startAt string, operation string) *TrackedOperationV1 {
	this := TrackedOperationV1{}
	this.StartAt = startAt
	this.Operation = operation
	return &this
}

// NewTrackedOperationV1WithDefaults instantiates a new TrackedOperationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackedOperationV1WithDefaults() *TrackedOperationV1 {
	this := TrackedOperationV1{}
	return &this
}

// GetStartAt returns the StartAt field value
func (o *TrackedOperationV1) GetStartAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *TrackedOperationV1) GetStartAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *TrackedOperationV1) SetStartAt(v string) {
	o.StartAt = v
}

// GetOperation returns the Operation field value
func (o *TrackedOperationV1) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *TrackedOperationV1) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *TrackedOperationV1) SetOperation(v string) {
	o.Operation = v
}

func (o TrackedOperationV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackedOperationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startAt"] = o.StartAt
	toSerialize["operation"] = o.Operation
	return toSerialize, nil
}

type NullableTrackedOperationV1 struct {
	value *TrackedOperationV1
	isSet bool
}

func (v NullableTrackedOperationV1) Get() *TrackedOperationV1 {
	return v.value
}

func (v *NullableTrackedOperationV1) Set(val *TrackedOperationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedOperationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedOperationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedOperationV1(val *TrackedOperationV1) *NullableTrackedOperationV1 {
	return &NullableTrackedOperationV1{value: val, isSet: true}
}

func (v NullableTrackedOperationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedOperationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


