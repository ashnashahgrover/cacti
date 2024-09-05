/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
)

// checks if the RunTransactionV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionV1Request{}

// RunTransactionV1Request struct for RunTransactionV1Request
type RunTransactionV1Request struct {
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	TransactionConfig XdaiTransactionConfig `json:"transactionConfig"`
	ConsistencyStrategy ConsistencyStrategy `json:"consistencyStrategy"`
}

// NewRunTransactionV1Request instantiates a new RunTransactionV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionV1Request(web3SigningCredential Web3SigningCredential, transactionConfig XdaiTransactionConfig, consistencyStrategy ConsistencyStrategy) *RunTransactionV1Request {
	this := RunTransactionV1Request{}
	this.Web3SigningCredential = web3SigningCredential
	this.TransactionConfig = transactionConfig
	this.ConsistencyStrategy = consistencyStrategy
	return &this
}

// NewRunTransactionV1RequestWithDefaults instantiates a new RunTransactionV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionV1RequestWithDefaults() *RunTransactionV1Request {
	this := RunTransactionV1Request{}
	return &this
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *RunTransactionV1Request) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *RunTransactionV1Request) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *RunTransactionV1Request) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetTransactionConfig returns the TransactionConfig field value
func (o *RunTransactionV1Request) GetTransactionConfig() XdaiTransactionConfig {
	if o == nil {
		var ret XdaiTransactionConfig
		return ret
	}

	return o.TransactionConfig
}

// GetTransactionConfigOk returns a tuple with the TransactionConfig field value
// and a boolean to check if the value has been set.
func (o *RunTransactionV1Request) GetTransactionConfigOk() (*XdaiTransactionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionConfig, true
}

// SetTransactionConfig sets field value
func (o *RunTransactionV1Request) SetTransactionConfig(v XdaiTransactionConfig) {
	o.TransactionConfig = v
}

// GetConsistencyStrategy returns the ConsistencyStrategy field value
func (o *RunTransactionV1Request) GetConsistencyStrategy() ConsistencyStrategy {
	if o == nil {
		var ret ConsistencyStrategy
		return ret
	}

	return o.ConsistencyStrategy
}

// GetConsistencyStrategyOk returns a tuple with the ConsistencyStrategy field value
// and a boolean to check if the value has been set.
func (o *RunTransactionV1Request) GetConsistencyStrategyOk() (*ConsistencyStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsistencyStrategy, true
}

// SetConsistencyStrategy sets field value
func (o *RunTransactionV1Request) SetConsistencyStrategy(v ConsistencyStrategy) {
	o.ConsistencyStrategy = v
}

func (o RunTransactionV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["transactionConfig"] = o.TransactionConfig
	toSerialize["consistencyStrategy"] = o.ConsistencyStrategy
	return toSerialize, nil
}

type NullableRunTransactionV1Request struct {
	value *RunTransactionV1Request
	isSet bool
}

func (v NullableRunTransactionV1Request) Get() *RunTransactionV1Request {
	return v.value
}

func (v *NullableRunTransactionV1Request) Set(val *RunTransactionV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionV1Request(val *RunTransactionV1Request) *NullableRunTransactionV1Request {
	return &NullableRunTransactionV1Request{value: val, isSet: true}
}

func (v NullableRunTransactionV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

