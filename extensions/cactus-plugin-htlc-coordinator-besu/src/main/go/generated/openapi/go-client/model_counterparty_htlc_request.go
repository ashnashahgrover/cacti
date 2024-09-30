/*
Hyperledger Cactus Plugin - HTLC Coordinator

Can exchange assets between networks

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-htlc-coordinator-besu

import (
	"encoding/json"
)

// checks if the CounterpartyHTLCRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CounterpartyHTLCRequest{}

// CounterpartyHTLCRequest struct for CounterpartyHTLCRequest
type CounterpartyHTLCRequest struct {
	HtlcPackage HtlcPackage `json:"htlcPackage"`
	// connector Instance Id for the connector plugin
	ConnectorInstanceId string `json:"connectorInstanceId"`
	// keychainId for the keychain plugin
	KeychainId string `json:"keychainId"`
	// Id for the HTLC
	HtlcId string `json:"htlcId"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	Gas *float32 `json:"gas,omitempty"`
}

// NewCounterpartyHTLCRequest instantiates a new CounterpartyHTLCRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCounterpartyHTLCRequest(htlcPackage HtlcPackage, connectorInstanceId string, keychainId string, htlcId string, web3SigningCredential Web3SigningCredential) *CounterpartyHTLCRequest {
	this := CounterpartyHTLCRequest{}
	this.HtlcPackage = htlcPackage
	this.ConnectorInstanceId = connectorInstanceId
	this.KeychainId = keychainId
	this.HtlcId = htlcId
	this.Web3SigningCredential = web3SigningCredential
	return &this
}

// NewCounterpartyHTLCRequestWithDefaults instantiates a new CounterpartyHTLCRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCounterpartyHTLCRequestWithDefaults() *CounterpartyHTLCRequest {
	this := CounterpartyHTLCRequest{}
	return &this
}

// GetHtlcPackage returns the HtlcPackage field value
func (o *CounterpartyHTLCRequest) GetHtlcPackage() HtlcPackage {
	if o == nil {
		var ret HtlcPackage
		return ret
	}

	return o.HtlcPackage
}

// GetHtlcPackageOk returns a tuple with the HtlcPackage field value
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetHtlcPackageOk() (*HtlcPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtlcPackage, true
}

// SetHtlcPackage sets field value
func (o *CounterpartyHTLCRequest) SetHtlcPackage(v HtlcPackage) {
	o.HtlcPackage = v
}

// GetConnectorInstanceId returns the ConnectorInstanceId field value
func (o *CounterpartyHTLCRequest) GetConnectorInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorInstanceId
}

// GetConnectorInstanceIdOk returns a tuple with the ConnectorInstanceId field value
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetConnectorInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorInstanceId, true
}

// SetConnectorInstanceId sets field value
func (o *CounterpartyHTLCRequest) SetConnectorInstanceId(v string) {
	o.ConnectorInstanceId = v
}

// GetKeychainId returns the KeychainId field value
func (o *CounterpartyHTLCRequest) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *CounterpartyHTLCRequest) SetKeychainId(v string) {
	o.KeychainId = v
}

// GetHtlcId returns the HtlcId field value
func (o *CounterpartyHTLCRequest) GetHtlcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtlcId
}

// GetHtlcIdOk returns a tuple with the HtlcId field value
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetHtlcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtlcId, true
}

// SetHtlcId sets field value
func (o *CounterpartyHTLCRequest) SetHtlcId(v string) {
	o.HtlcId = v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *CounterpartyHTLCRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *CounterpartyHTLCRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *CounterpartyHTLCRequest) GetGas() float32 {
	if o == nil || IsNil(o.Gas) {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CounterpartyHTLCRequest) GetGasOk() (*float32, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *CounterpartyHTLCRequest) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *CounterpartyHTLCRequest) SetGas(v float32) {
	o.Gas = &v
}

func (o CounterpartyHTLCRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CounterpartyHTLCRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["htlcPackage"] = o.HtlcPackage
	toSerialize["connectorInstanceId"] = o.ConnectorInstanceId
	toSerialize["keychainId"] = o.KeychainId
	toSerialize["htlcId"] = o.HtlcId
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	return toSerialize, nil
}

type NullableCounterpartyHTLCRequest struct {
	value *CounterpartyHTLCRequest
	isSet bool
}

func (v NullableCounterpartyHTLCRequest) Get() *CounterpartyHTLCRequest {
	return v.value
}

func (v *NullableCounterpartyHTLCRequest) Set(val *CounterpartyHTLCRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterpartyHTLCRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterpartyHTLCRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterpartyHTLCRequest(val *CounterpartyHTLCRequest) *NullableCounterpartyHTLCRequest {
	return &NullableCounterpartyHTLCRequest{value: val, isSet: true}
}

func (v NullableCounterpartyHTLCRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterpartyHTLCRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


