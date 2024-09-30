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

// checks if the StatusResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseV1{}

// StatusResponseV1 Response with plugin status report.
type StatusResponseV1 struct {
	// Plugin instance id.
	InstanceId string `json:"instanceId"`
	// True if successfully connected to the database, false otherwise.
	Connected bool `json:"connected"`
	// True if web services were correctly exported.
	WebServicesRegistered bool `json:"webServicesRegistered"`
	// Total number of tokens being monitored by the plugin.
	MonitoredTokensCount float32 `json:"monitoredTokensCount"`
	OperationsRunning []TrackedOperationV1 `json:"operationsRunning"`
	// True if block monitoring is running, false otherwise.
	MonitorRunning bool `json:"monitorRunning"`
	// Number of the last block seen by the block monitor.
	LastSeenBlock float32 `json:"lastSeenBlock"`
}

// NewStatusResponseV1 instantiates a new StatusResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseV1(instanceId string, connected bool, webServicesRegistered bool, monitoredTokensCount float32, operationsRunning []TrackedOperationV1, monitorRunning bool, lastSeenBlock float32) *StatusResponseV1 {
	this := StatusResponseV1{}
	this.InstanceId = instanceId
	this.Connected = connected
	this.WebServicesRegistered = webServicesRegistered
	this.MonitoredTokensCount = monitoredTokensCount
	this.OperationsRunning = operationsRunning
	this.MonitorRunning = monitorRunning
	this.LastSeenBlock = lastSeenBlock
	return &this
}

// NewStatusResponseV1WithDefaults instantiates a new StatusResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseV1WithDefaults() *StatusResponseV1 {
	this := StatusResponseV1{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *StatusResponseV1) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *StatusResponseV1) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetConnected returns the Connected field value
func (o *StatusResponseV1) GetConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connected, true
}

// SetConnected sets field value
func (o *StatusResponseV1) SetConnected(v bool) {
	o.Connected = v
}

// GetWebServicesRegistered returns the WebServicesRegistered field value
func (o *StatusResponseV1) GetWebServicesRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WebServicesRegistered
}

// GetWebServicesRegisteredOk returns a tuple with the WebServicesRegistered field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetWebServicesRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebServicesRegistered, true
}

// SetWebServicesRegistered sets field value
func (o *StatusResponseV1) SetWebServicesRegistered(v bool) {
	o.WebServicesRegistered = v
}

// GetMonitoredTokensCount returns the MonitoredTokensCount field value
func (o *StatusResponseV1) GetMonitoredTokensCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonitoredTokensCount
}

// GetMonitoredTokensCountOk returns a tuple with the MonitoredTokensCount field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetMonitoredTokensCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoredTokensCount, true
}

// SetMonitoredTokensCount sets field value
func (o *StatusResponseV1) SetMonitoredTokensCount(v float32) {
	o.MonitoredTokensCount = v
}

// GetOperationsRunning returns the OperationsRunning field value
func (o *StatusResponseV1) GetOperationsRunning() []TrackedOperationV1 {
	if o == nil {
		var ret []TrackedOperationV1
		return ret
	}

	return o.OperationsRunning
}

// GetOperationsRunningOk returns a tuple with the OperationsRunning field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetOperationsRunningOk() ([]TrackedOperationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationsRunning, true
}

// SetOperationsRunning sets field value
func (o *StatusResponseV1) SetOperationsRunning(v []TrackedOperationV1) {
	o.OperationsRunning = v
}

// GetMonitorRunning returns the MonitorRunning field value
func (o *StatusResponseV1) GetMonitorRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MonitorRunning
}

// GetMonitorRunningOk returns a tuple with the MonitorRunning field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetMonitorRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorRunning, true
}

// SetMonitorRunning sets field value
func (o *StatusResponseV1) SetMonitorRunning(v bool) {
	o.MonitorRunning = v
}

// GetLastSeenBlock returns the LastSeenBlock field value
func (o *StatusResponseV1) GetLastSeenBlock() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastSeenBlock
}

// GetLastSeenBlockOk returns a tuple with the LastSeenBlock field value
// and a boolean to check if the value has been set.
func (o *StatusResponseV1) GetLastSeenBlockOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeenBlock, true
}

// SetLastSeenBlock sets field value
func (o *StatusResponseV1) SetLastSeenBlock(v float32) {
	o.LastSeenBlock = v
}

func (o StatusResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["connected"] = o.Connected
	toSerialize["webServicesRegistered"] = o.WebServicesRegistered
	toSerialize["monitoredTokensCount"] = o.MonitoredTokensCount
	toSerialize["operationsRunning"] = o.OperationsRunning
	toSerialize["monitorRunning"] = o.MonitorRunning
	toSerialize["lastSeenBlock"] = o.LastSeenBlock
	return toSerialize, nil
}

type NullableStatusResponseV1 struct {
	value *StatusResponseV1
	isSet bool
}

func (v NullableStatusResponseV1) Get() *StatusResponseV1 {
	return v.value
}

func (v *NullableStatusResponseV1) Set(val *StatusResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseV1(val *StatusResponseV1) *NullableStatusResponseV1 {
	return &NullableStatusResponseV1{value: val, isSet: true}
}

func (v NullableStatusResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


