/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the LockEvidenceV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockEvidenceV1Response{}

// LockEvidenceV1Response struct for LockEvidenceV1Response
type LockEvidenceV1Response struct {
	SessionID string `json:"sessionID"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	HashLockEvidenceRequest string `json:"hashLockEvidenceRequest"`
	ServerTransferNumber NullableInt32 `json:"serverTransferNumber,omitempty"`
	Signature string `json:"signature"`
	MessageType string `json:"messageType"`
	SequenceNumber float32 `json:"sequenceNumber"`
}

// NewLockEvidenceV1Response instantiates a new LockEvidenceV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockEvidenceV1Response(sessionID string, clientIdentityPubkey string, serverIdentityPubkey string, hashLockEvidenceRequest string, signature string, messageType string, sequenceNumber float32) *LockEvidenceV1Response {
	this := LockEvidenceV1Response{}
	this.SessionID = sessionID
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.HashLockEvidenceRequest = hashLockEvidenceRequest
	this.Signature = signature
	this.MessageType = messageType
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewLockEvidenceV1ResponseWithDefaults instantiates a new LockEvidenceV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockEvidenceV1ResponseWithDefaults() *LockEvidenceV1Response {
	this := LockEvidenceV1Response{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *LockEvidenceV1Response) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *LockEvidenceV1Response) SetSessionID(v string) {
	o.SessionID = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *LockEvidenceV1Response) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *LockEvidenceV1Response) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *LockEvidenceV1Response) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *LockEvidenceV1Response) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetHashLockEvidenceRequest returns the HashLockEvidenceRequest field value
func (o *LockEvidenceV1Response) GetHashLockEvidenceRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashLockEvidenceRequest
}

// GetHashLockEvidenceRequestOk returns a tuple with the HashLockEvidenceRequest field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetHashLockEvidenceRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashLockEvidenceRequest, true
}

// SetHashLockEvidenceRequest sets field value
func (o *LockEvidenceV1Response) SetHashLockEvidenceRequest(v string) {
	o.HashLockEvidenceRequest = v
}

// GetServerTransferNumber returns the ServerTransferNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LockEvidenceV1Response) GetServerTransferNumber() int32 {
	if o == nil || IsNil(o.ServerTransferNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ServerTransferNumber.Get()
}

// GetServerTransferNumberOk returns a tuple with the ServerTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LockEvidenceV1Response) GetServerTransferNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerTransferNumber.Get(), o.ServerTransferNumber.IsSet()
}

// HasServerTransferNumber returns a boolean if a field has been set.
func (o *LockEvidenceV1Response) HasServerTransferNumber() bool {
	if o != nil && o.ServerTransferNumber.IsSet() {
		return true
	}

	return false
}

// SetServerTransferNumber gets a reference to the given NullableInt32 and assigns it to the ServerTransferNumber field.
func (o *LockEvidenceV1Response) SetServerTransferNumber(v int32) {
	o.ServerTransferNumber.Set(&v)
}
// SetServerTransferNumberNil sets the value for ServerTransferNumber to be an explicit nil
func (o *LockEvidenceV1Response) SetServerTransferNumberNil() {
	o.ServerTransferNumber.Set(nil)
}

// UnsetServerTransferNumber ensures that no value is present for ServerTransferNumber, not even an explicit nil
func (o *LockEvidenceV1Response) UnsetServerTransferNumber() {
	o.ServerTransferNumber.Unset()
}

// GetSignature returns the Signature field value
func (o *LockEvidenceV1Response) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *LockEvidenceV1Response) SetSignature(v string) {
	o.Signature = v
}

// GetMessageType returns the MessageType field value
func (o *LockEvidenceV1Response) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *LockEvidenceV1Response) SetMessageType(v string) {
	o.MessageType = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *LockEvidenceV1Response) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Response) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *LockEvidenceV1Response) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

func (o LockEvidenceV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockEvidenceV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["hashLockEvidenceRequest"] = o.HashLockEvidenceRequest
	if o.ServerTransferNumber.IsSet() {
		toSerialize["serverTransferNumber"] = o.ServerTransferNumber.Get()
	}
	toSerialize["signature"] = o.Signature
	toSerialize["messageType"] = o.MessageType
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableLockEvidenceV1Response struct {
	value *LockEvidenceV1Response
	isSet bool
}

func (v NullableLockEvidenceV1Response) Get() *LockEvidenceV1Response {
	return v.value
}

func (v *NullableLockEvidenceV1Response) Set(val *LockEvidenceV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLockEvidenceV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLockEvidenceV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockEvidenceV1Response(val *LockEvidenceV1Response) *NullableLockEvidenceV1Response {
	return &NullableLockEvidenceV1Response{value: val, isSet: true}
}

func (v NullableLockEvidenceV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockEvidenceV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

