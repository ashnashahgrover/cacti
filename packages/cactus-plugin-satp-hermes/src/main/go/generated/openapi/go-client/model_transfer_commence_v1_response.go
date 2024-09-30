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

// checks if the TransferCommenceV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferCommenceV1Response{}

// TransferCommenceV1Response struct for TransferCommenceV1Response
type TransferCommenceV1Response struct {
	SessionID string `json:"sessionID"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	HashCommenceRequest string `json:"hashCommenceRequest"`
	ServerTransferNumber NullableInt32 `json:"serverTransferNumber,omitempty"`
	Signature string `json:"signature"`
	MessageType string `json:"messageType"`
	MessageHash *string `json:"messageHash,omitempty"`
	SequenceNumber float32 `json:"sequenceNumber"`
}

// NewTransferCommenceV1Response instantiates a new TransferCommenceV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCommenceV1Response(sessionID string, clientIdentityPubkey string, serverIdentityPubkey string, hashCommenceRequest string, signature string, messageType string, sequenceNumber float32) *TransferCommenceV1Response {
	this := TransferCommenceV1Response{}
	this.SessionID = sessionID
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.HashCommenceRequest = hashCommenceRequest
	this.Signature = signature
	this.MessageType = messageType
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewTransferCommenceV1ResponseWithDefaults instantiates a new TransferCommenceV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCommenceV1ResponseWithDefaults() *TransferCommenceV1Response {
	this := TransferCommenceV1Response{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *TransferCommenceV1Response) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *TransferCommenceV1Response) SetSessionID(v string) {
	o.SessionID = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *TransferCommenceV1Response) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *TransferCommenceV1Response) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *TransferCommenceV1Response) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *TransferCommenceV1Response) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetHashCommenceRequest returns the HashCommenceRequest field value
func (o *TransferCommenceV1Response) GetHashCommenceRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashCommenceRequest
}

// GetHashCommenceRequestOk returns a tuple with the HashCommenceRequest field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetHashCommenceRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashCommenceRequest, true
}

// SetHashCommenceRequest sets field value
func (o *TransferCommenceV1Response) SetHashCommenceRequest(v string) {
	o.HashCommenceRequest = v
}

// GetServerTransferNumber returns the ServerTransferNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferCommenceV1Response) GetServerTransferNumber() int32 {
	if o == nil || IsNil(o.ServerTransferNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ServerTransferNumber.Get()
}

// GetServerTransferNumberOk returns a tuple with the ServerTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferCommenceV1Response) GetServerTransferNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerTransferNumber.Get(), o.ServerTransferNumber.IsSet()
}

// HasServerTransferNumber returns a boolean if a field has been set.
func (o *TransferCommenceV1Response) HasServerTransferNumber() bool {
	if o != nil && o.ServerTransferNumber.IsSet() {
		return true
	}

	return false
}

// SetServerTransferNumber gets a reference to the given NullableInt32 and assigns it to the ServerTransferNumber field.
func (o *TransferCommenceV1Response) SetServerTransferNumber(v int32) {
	o.ServerTransferNumber.Set(&v)
}
// SetServerTransferNumberNil sets the value for ServerTransferNumber to be an explicit nil
func (o *TransferCommenceV1Response) SetServerTransferNumberNil() {
	o.ServerTransferNumber.Set(nil)
}

// UnsetServerTransferNumber ensures that no value is present for ServerTransferNumber, not even an explicit nil
func (o *TransferCommenceV1Response) UnsetServerTransferNumber() {
	o.ServerTransferNumber.Unset()
}

// GetSignature returns the Signature field value
func (o *TransferCommenceV1Response) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *TransferCommenceV1Response) SetSignature(v string) {
	o.Signature = v
}

// GetMessageType returns the MessageType field value
func (o *TransferCommenceV1Response) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *TransferCommenceV1Response) SetMessageType(v string) {
	o.MessageType = v
}

// GetMessageHash returns the MessageHash field value if set, zero value otherwise.
func (o *TransferCommenceV1Response) GetMessageHash() string {
	if o == nil || IsNil(o.MessageHash) {
		var ret string
		return ret
	}
	return *o.MessageHash
}

// GetMessageHashOk returns a tuple with the MessageHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetMessageHashOk() (*string, bool) {
	if o == nil || IsNil(o.MessageHash) {
		return nil, false
	}
	return o.MessageHash, true
}

// HasMessageHash returns a boolean if a field has been set.
func (o *TransferCommenceV1Response) HasMessageHash() bool {
	if o != nil && !IsNil(o.MessageHash) {
		return true
	}

	return false
}

// SetMessageHash gets a reference to the given string and assigns it to the MessageHash field.
func (o *TransferCommenceV1Response) SetMessageHash(v string) {
	o.MessageHash = &v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *TransferCommenceV1Response) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Response) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *TransferCommenceV1Response) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

func (o TransferCommenceV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferCommenceV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["hashCommenceRequest"] = o.HashCommenceRequest
	if o.ServerTransferNumber.IsSet() {
		toSerialize["serverTransferNumber"] = o.ServerTransferNumber.Get()
	}
	toSerialize["signature"] = o.Signature
	toSerialize["messageType"] = o.MessageType
	if !IsNil(o.MessageHash) {
		toSerialize["messageHash"] = o.MessageHash
	}
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableTransferCommenceV1Response struct {
	value *TransferCommenceV1Response
	isSet bool
}

func (v NullableTransferCommenceV1Response) Get() *TransferCommenceV1Response {
	return v.value
}

func (v *NullableTransferCommenceV1Response) Set(val *TransferCommenceV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCommenceV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCommenceV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCommenceV1Response(val *TransferCommenceV1Response) *NullableTransferCommenceV1Response {
	return &NullableTransferCommenceV1Response{value: val, isSet: true}
}

func (v NullableTransferCommenceV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCommenceV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


