/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the StopMonitorV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopMonitorV1Request{}

// StopMonitorV1Request struct for StopMonitorV1Request
type StopMonitorV1Request struct {
	// ID of a client application that wants to monitor the state changes
	ClientAppId string `json:"clientAppId"`
	// The fully qualified name of the Corda state to monitor
	StateFullClassName string `json:"stateFullClassName"`
}

// NewStopMonitorV1Request instantiates a new StopMonitorV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopMonitorV1Request(clientAppId string, stateFullClassName string) *StopMonitorV1Request {
	this := StopMonitorV1Request{}
	this.ClientAppId = clientAppId
	this.StateFullClassName = stateFullClassName
	return &this
}

// NewStopMonitorV1RequestWithDefaults instantiates a new StopMonitorV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopMonitorV1RequestWithDefaults() *StopMonitorV1Request {
	this := StopMonitorV1Request{}
	return &this
}

// GetClientAppId returns the ClientAppId field value
func (o *StopMonitorV1Request) GetClientAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientAppId
}

// GetClientAppIdOk returns a tuple with the ClientAppId field value
// and a boolean to check if the value has been set.
func (o *StopMonitorV1Request) GetClientAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAppId, true
}

// SetClientAppId sets field value
func (o *StopMonitorV1Request) SetClientAppId(v string) {
	o.ClientAppId = v
}

// GetStateFullClassName returns the StateFullClassName field value
func (o *StopMonitorV1Request) GetStateFullClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateFullClassName
}

// GetStateFullClassNameOk returns a tuple with the StateFullClassName field value
// and a boolean to check if the value has been set.
func (o *StopMonitorV1Request) GetStateFullClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateFullClassName, true
}

// SetStateFullClassName sets field value
func (o *StopMonitorV1Request) SetStateFullClassName(v string) {
	o.StateFullClassName = v
}

func (o StopMonitorV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopMonitorV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientAppId"] = o.ClientAppId
	toSerialize["stateFullClassName"] = o.StateFullClassName
	return toSerialize, nil
}

type NullableStopMonitorV1Request struct {
	value *StopMonitorV1Request
	isSet bool
}

func (v NullableStopMonitorV1Request) Get() *StopMonitorV1Request {
	return v.value
}

func (v *NullableStopMonitorV1Request) Set(val *StopMonitorV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableStopMonitorV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableStopMonitorV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopMonitorV1Request(val *StopMonitorV1Request) *NullableStopMonitorV1Request {
	return &NullableStopMonitorV1Request{value: val, isSet: true}
}

func (v NullableStopMonitorV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopMonitorV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


