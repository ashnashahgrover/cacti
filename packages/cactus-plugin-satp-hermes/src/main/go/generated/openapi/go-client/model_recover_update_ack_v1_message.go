/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the RecoverUpdateAckV1Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverUpdateAckV1Message{}

// RecoverUpdateAckV1Message struct for RecoverUpdateAckV1Message
type RecoverUpdateAckV1Message struct {
	SessionID string `json:"sessionID"`
	Success bool `json:"success"`
	ChangedEntriesHash []string `json:"changedEntriesHash"`
	Signature string `json:"signature"`
}

// NewRecoverUpdateAckV1Message instantiates a new RecoverUpdateAckV1Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverUpdateAckV1Message(sessionID string, success bool, changedEntriesHash []string, signature string) *RecoverUpdateAckV1Message {
	this := RecoverUpdateAckV1Message{}
	this.SessionID = sessionID
	this.Success = success
	this.ChangedEntriesHash = changedEntriesHash
	this.Signature = signature
	return &this
}

// NewRecoverUpdateAckV1MessageWithDefaults instantiates a new RecoverUpdateAckV1Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverUpdateAckV1MessageWithDefaults() *RecoverUpdateAckV1Message {
	this := RecoverUpdateAckV1Message{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *RecoverUpdateAckV1Message) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *RecoverUpdateAckV1Message) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *RecoverUpdateAckV1Message) SetSessionID(v string) {
	o.SessionID = v
}

// GetSuccess returns the Success field value
func (o *RecoverUpdateAckV1Message) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RecoverUpdateAckV1Message) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RecoverUpdateAckV1Message) SetSuccess(v bool) {
	o.Success = v
}

// GetChangedEntriesHash returns the ChangedEntriesHash field value
func (o *RecoverUpdateAckV1Message) GetChangedEntriesHash() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ChangedEntriesHash
}

// GetChangedEntriesHashOk returns a tuple with the ChangedEntriesHash field value
// and a boolean to check if the value has been set.
func (o *RecoverUpdateAckV1Message) GetChangedEntriesHashOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangedEntriesHash, true
}

// SetChangedEntriesHash sets field value
func (o *RecoverUpdateAckV1Message) SetChangedEntriesHash(v []string) {
	o.ChangedEntriesHash = v
}

// GetSignature returns the Signature field value
func (o *RecoverUpdateAckV1Message) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RecoverUpdateAckV1Message) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RecoverUpdateAckV1Message) SetSignature(v string) {
	o.Signature = v
}

func (o RecoverUpdateAckV1Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverUpdateAckV1Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["success"] = o.Success
	toSerialize["changedEntriesHash"] = o.ChangedEntriesHash
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableRecoverUpdateAckV1Message struct {
	value *RecoverUpdateAckV1Message
	isSet bool
}

func (v NullableRecoverUpdateAckV1Message) Get() *RecoverUpdateAckV1Message {
	return v.value
}

func (v *NullableRecoverUpdateAckV1Message) Set(val *RecoverUpdateAckV1Message) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverUpdateAckV1Message) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverUpdateAckV1Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverUpdateAckV1Message(val *RecoverUpdateAckV1Message) *NullableRecoverUpdateAckV1Message {
	return &NullableRecoverUpdateAckV1Message{value: val, isSet: true}
}

func (v NullableRecoverUpdateAckV1Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverUpdateAckV1Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


