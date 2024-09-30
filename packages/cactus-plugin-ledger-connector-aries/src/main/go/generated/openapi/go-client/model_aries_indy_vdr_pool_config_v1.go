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

// checks if the AriesIndyVdrPoolConfigV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AriesIndyVdrPoolConfigV1{}

// AriesIndyVdrPoolConfigV1 Indy VDR network configuration
type AriesIndyVdrPoolConfigV1 struct {
	// Indy genesis transactions.
	GenesisTransactions string `json:"genesisTransactions"`
	// Flag to specify whether this is production or development ledger.
	IsProduction bool `json:"isProduction"`
	// Indy namespace
	IndyNamespace string `json:"indyNamespace"`
	// Connect to the ledger on startup flag
	ConnectOnStartup *bool `json:"connectOnStartup,omitempty"`
}

// NewAriesIndyVdrPoolConfigV1 instantiates a new AriesIndyVdrPoolConfigV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAriesIndyVdrPoolConfigV1(genesisTransactions string, isProduction bool, indyNamespace string) *AriesIndyVdrPoolConfigV1 {
	this := AriesIndyVdrPoolConfigV1{}
	this.GenesisTransactions = genesisTransactions
	this.IsProduction = isProduction
	this.IndyNamespace = indyNamespace
	return &this
}

// NewAriesIndyVdrPoolConfigV1WithDefaults instantiates a new AriesIndyVdrPoolConfigV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAriesIndyVdrPoolConfigV1WithDefaults() *AriesIndyVdrPoolConfigV1 {
	this := AriesIndyVdrPoolConfigV1{}
	return &this
}

// GetGenesisTransactions returns the GenesisTransactions field value
func (o *AriesIndyVdrPoolConfigV1) GetGenesisTransactions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GenesisTransactions
}

// GetGenesisTransactionsOk returns a tuple with the GenesisTransactions field value
// and a boolean to check if the value has been set.
func (o *AriesIndyVdrPoolConfigV1) GetGenesisTransactionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenesisTransactions, true
}

// SetGenesisTransactions sets field value
func (o *AriesIndyVdrPoolConfigV1) SetGenesisTransactions(v string) {
	o.GenesisTransactions = v
}

// GetIsProduction returns the IsProduction field value
func (o *AriesIndyVdrPoolConfigV1) GetIsProduction() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value
// and a boolean to check if the value has been set.
func (o *AriesIndyVdrPoolConfigV1) GetIsProductionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsProduction, true
}

// SetIsProduction sets field value
func (o *AriesIndyVdrPoolConfigV1) SetIsProduction(v bool) {
	o.IsProduction = v
}

// GetIndyNamespace returns the IndyNamespace field value
func (o *AriesIndyVdrPoolConfigV1) GetIndyNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndyNamespace
}

// GetIndyNamespaceOk returns a tuple with the IndyNamespace field value
// and a boolean to check if the value has been set.
func (o *AriesIndyVdrPoolConfigV1) GetIndyNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndyNamespace, true
}

// SetIndyNamespace sets field value
func (o *AriesIndyVdrPoolConfigV1) SetIndyNamespace(v string) {
	o.IndyNamespace = v
}

// GetConnectOnStartup returns the ConnectOnStartup field value if set, zero value otherwise.
func (o *AriesIndyVdrPoolConfigV1) GetConnectOnStartup() bool {
	if o == nil || IsNil(o.ConnectOnStartup) {
		var ret bool
		return ret
	}
	return *o.ConnectOnStartup
}

// GetConnectOnStartupOk returns a tuple with the ConnectOnStartup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AriesIndyVdrPoolConfigV1) GetConnectOnStartupOk() (*bool, bool) {
	if o == nil || IsNil(o.ConnectOnStartup) {
		return nil, false
	}
	return o.ConnectOnStartup, true
}

// HasConnectOnStartup returns a boolean if a field has been set.
func (o *AriesIndyVdrPoolConfigV1) HasConnectOnStartup() bool {
	if o != nil && !IsNil(o.ConnectOnStartup) {
		return true
	}

	return false
}

// SetConnectOnStartup gets a reference to the given bool and assigns it to the ConnectOnStartup field.
func (o *AriesIndyVdrPoolConfigV1) SetConnectOnStartup(v bool) {
	o.ConnectOnStartup = &v
}

func (o AriesIndyVdrPoolConfigV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AriesIndyVdrPoolConfigV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["genesisTransactions"] = o.GenesisTransactions
	toSerialize["isProduction"] = o.IsProduction
	toSerialize["indyNamespace"] = o.IndyNamespace
	if !IsNil(o.ConnectOnStartup) {
		toSerialize["connectOnStartup"] = o.ConnectOnStartup
	}
	return toSerialize, nil
}

type NullableAriesIndyVdrPoolConfigV1 struct {
	value *AriesIndyVdrPoolConfigV1
	isSet bool
}

func (v NullableAriesIndyVdrPoolConfigV1) Get() *AriesIndyVdrPoolConfigV1 {
	return v.value
}

func (v *NullableAriesIndyVdrPoolConfigV1) Set(val *AriesIndyVdrPoolConfigV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAriesIndyVdrPoolConfigV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAriesIndyVdrPoolConfigV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAriesIndyVdrPoolConfigV1(val *AriesIndyVdrPoolConfigV1) *NullableAriesIndyVdrPoolConfigV1 {
	return &NullableAriesIndyVdrPoolConfigV1{value: val, isSet: true}
}

func (v NullableAriesIndyVdrPoolConfigV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAriesIndyVdrPoolConfigV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


