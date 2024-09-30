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

// checks if the AcceptInvitationV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptInvitationV1Response{}

// AcceptInvitationV1Response Response for AcceptInvitation endpoint.
type AcceptInvitationV1Response struct {
	// ID that can be used to track status of the connection
	OutOfBandId string `json:"outOfBandId"`
}

// NewAcceptInvitationV1Response instantiates a new AcceptInvitationV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptInvitationV1Response(outOfBandId string) *AcceptInvitationV1Response {
	this := AcceptInvitationV1Response{}
	this.OutOfBandId = outOfBandId
	return &this
}

// NewAcceptInvitationV1ResponseWithDefaults instantiates a new AcceptInvitationV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptInvitationV1ResponseWithDefaults() *AcceptInvitationV1Response {
	this := AcceptInvitationV1Response{}
	return &this
}

// GetOutOfBandId returns the OutOfBandId field value
func (o *AcceptInvitationV1Response) GetOutOfBandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutOfBandId
}

// GetOutOfBandIdOk returns a tuple with the OutOfBandId field value
// and a boolean to check if the value has been set.
func (o *AcceptInvitationV1Response) GetOutOfBandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutOfBandId, true
}

// SetOutOfBandId sets field value
func (o *AcceptInvitationV1Response) SetOutOfBandId(v string) {
	o.OutOfBandId = v
}

func (o AcceptInvitationV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptInvitationV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["outOfBandId"] = o.OutOfBandId
	return toSerialize, nil
}

type NullableAcceptInvitationV1Response struct {
	value *AcceptInvitationV1Response
	isSet bool
}

func (v NullableAcceptInvitationV1Response) Get() *AcceptInvitationV1Response {
	return v.value
}

func (v *NullableAcceptInvitationV1Response) Set(val *AcceptInvitationV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptInvitationV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptInvitationV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptInvitationV1Response(val *AcceptInvitationV1Response) *NullableAcceptInvitationV1Response {
	return &NullableAcceptInvitationV1Response{value: val, isSet: true}
}

func (v NullableAcceptInvitationV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptInvitationV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


