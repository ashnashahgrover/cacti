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

// checks if the CreateNewConnectionInvitationV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNewConnectionInvitationV1Request{}

// CreateNewConnectionInvitationV1Request Request for CreateNewConnectionInvitation endpoint.
type CreateNewConnectionInvitationV1Request struct {
	// Aries label of an agent to use to generate an invitation
	AgentName string `json:"agentName"`
	// Invitation URL domain to use. If not specified, then connector default domain will be used
	InvitationDomain *string `json:"invitationDomain,omitempty"`
}

// NewCreateNewConnectionInvitationV1Request instantiates a new CreateNewConnectionInvitationV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNewConnectionInvitationV1Request(agentName string) *CreateNewConnectionInvitationV1Request {
	this := CreateNewConnectionInvitationV1Request{}
	this.AgentName = agentName
	return &this
}

// NewCreateNewConnectionInvitationV1RequestWithDefaults instantiates a new CreateNewConnectionInvitationV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNewConnectionInvitationV1RequestWithDefaults() *CreateNewConnectionInvitationV1Request {
	this := CreateNewConnectionInvitationV1Request{}
	return &this
}

// GetAgentName returns the AgentName field value
func (o *CreateNewConnectionInvitationV1Request) GetAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value
// and a boolean to check if the value has been set.
func (o *CreateNewConnectionInvitationV1Request) GetAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentName, true
}

// SetAgentName sets field value
func (o *CreateNewConnectionInvitationV1Request) SetAgentName(v string) {
	o.AgentName = v
}

// GetInvitationDomain returns the InvitationDomain field value if set, zero value otherwise.
func (o *CreateNewConnectionInvitationV1Request) GetInvitationDomain() string {
	if o == nil || IsNil(o.InvitationDomain) {
		var ret string
		return ret
	}
	return *o.InvitationDomain
}

// GetInvitationDomainOk returns a tuple with the InvitationDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNewConnectionInvitationV1Request) GetInvitationDomainOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationDomain) {
		return nil, false
	}
	return o.InvitationDomain, true
}

// HasInvitationDomain returns a boolean if a field has been set.
func (o *CreateNewConnectionInvitationV1Request) HasInvitationDomain() bool {
	if o != nil && !IsNil(o.InvitationDomain) {
		return true
	}

	return false
}

// SetInvitationDomain gets a reference to the given string and assigns it to the InvitationDomain field.
func (o *CreateNewConnectionInvitationV1Request) SetInvitationDomain(v string) {
	o.InvitationDomain = &v
}

func (o CreateNewConnectionInvitationV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNewConnectionInvitationV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentName"] = o.AgentName
	if !IsNil(o.InvitationDomain) {
		toSerialize["invitationDomain"] = o.InvitationDomain
	}
	return toSerialize, nil
}

type NullableCreateNewConnectionInvitationV1Request struct {
	value *CreateNewConnectionInvitationV1Request
	isSet bool
}

func (v NullableCreateNewConnectionInvitationV1Request) Get() *CreateNewConnectionInvitationV1Request {
	return v.value
}

func (v *NullableCreateNewConnectionInvitationV1Request) Set(val *CreateNewConnectionInvitationV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNewConnectionInvitationV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNewConnectionInvitationV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNewConnectionInvitationV1Request(val *CreateNewConnectionInvitationV1Request) *NullableCreateNewConnectionInvitationV1Request {
	return &NullableCreateNewConnectionInvitationV1Request{value: val, isSet: true}
}

func (v NullableCreateNewConnectionInvitationV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNewConnectionInvitationV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


