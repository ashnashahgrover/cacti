/*
Hyperledger Cactus Plugin - Connector Polkadot

Can perform basic tasks on a Polkadot parachain

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-polkadot

import (
	"encoding/json"
)

// checks if the InvokeContractRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvokeContractRequest{}

// InvokeContractRequest struct for InvokeContractRequest
type InvokeContractRequest struct {
	InvocationType PolkadotContractInvocationType `json:"invocationType"`
	AccountAddress string `json:"accountAddress"`
	Web3SigningCredential Web3SigningCredential `json:"web3SigningCredential"`
	Metadata PolkadotTransactionConfigTransferSubmittable `json:"metadata"`
	ContractAddress string `json:"contractAddress"`
	// The name of the contract method to invoke.
	MethodName string `json:"methodName"`
	GasLimit DeployContractInkRequestGasLimit `json:"gasLimit"`
	StorageDepositLimit NullableDeployContractInkRequestStorageDepositLimit `json:"storageDepositLimit,omitempty"`
	Balance *DeployContractInkRequestBalance `json:"balance,omitempty"`
	// The list of arguments to pass in to the contract method being invoked
	Params []interface{} `json:"params,omitempty"`
}

// NewInvokeContractRequest instantiates a new InvokeContractRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeContractRequest(invocationType PolkadotContractInvocationType, accountAddress string, web3SigningCredential Web3SigningCredential, metadata PolkadotTransactionConfigTransferSubmittable, contractAddress string, methodName string, gasLimit DeployContractInkRequestGasLimit) *InvokeContractRequest {
	this := InvokeContractRequest{}
	this.InvocationType = invocationType
	this.AccountAddress = accountAddress
	this.Web3SigningCredential = web3SigningCredential
	this.Metadata = metadata
	this.ContractAddress = contractAddress
	this.MethodName = methodName
	this.GasLimit = gasLimit
	return &this
}

// NewInvokeContractRequestWithDefaults instantiates a new InvokeContractRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeContractRequestWithDefaults() *InvokeContractRequest {
	this := InvokeContractRequest{}
	return &this
}

// GetInvocationType returns the InvocationType field value
func (o *InvokeContractRequest) GetInvocationType() PolkadotContractInvocationType {
	if o == nil {
		var ret PolkadotContractInvocationType
		return ret
	}

	return o.InvocationType
}

// GetInvocationTypeOk returns a tuple with the InvocationType field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetInvocationTypeOk() (*PolkadotContractInvocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvocationType, true
}

// SetInvocationType sets field value
func (o *InvokeContractRequest) SetInvocationType(v PolkadotContractInvocationType) {
	o.InvocationType = v
}

// GetAccountAddress returns the AccountAddress field value
func (o *InvokeContractRequest) GetAccountAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountAddress
}

// GetAccountAddressOk returns a tuple with the AccountAddress field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetAccountAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountAddress, true
}

// SetAccountAddress sets field value
func (o *InvokeContractRequest) SetAccountAddress(v string) {
	o.AccountAddress = v
}

// GetWeb3SigningCredential returns the Web3SigningCredential field value
func (o *InvokeContractRequest) GetWeb3SigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.Web3SigningCredential
}

// GetWeb3SigningCredentialOk returns a tuple with the Web3SigningCredential field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetWeb3SigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Web3SigningCredential, true
}

// SetWeb3SigningCredential sets field value
func (o *InvokeContractRequest) SetWeb3SigningCredential(v Web3SigningCredential) {
	o.Web3SigningCredential = v
}

// GetMetadata returns the Metadata field value
func (o *InvokeContractRequest) GetMetadata() PolkadotTransactionConfigTransferSubmittable {
	if o == nil {
		var ret PolkadotTransactionConfigTransferSubmittable
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetMetadataOk() (*PolkadotTransactionConfigTransferSubmittable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *InvokeContractRequest) SetMetadata(v PolkadotTransactionConfigTransferSubmittable) {
	o.Metadata = v
}

// GetContractAddress returns the ContractAddress field value
func (o *InvokeContractRequest) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *InvokeContractRequest) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetMethodName returns the MethodName field value
func (o *InvokeContractRequest) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *InvokeContractRequest) SetMethodName(v string) {
	o.MethodName = v
}

// GetGasLimit returns the GasLimit field value
func (o *InvokeContractRequest) GetGasLimit() DeployContractInkRequestGasLimit {
	if o == nil {
		var ret DeployContractInkRequestGasLimit
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetGasLimitOk() (*DeployContractInkRequestGasLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *InvokeContractRequest) SetGasLimit(v DeployContractInkRequestGasLimit) {
	o.GasLimit = v
}

// GetStorageDepositLimit returns the StorageDepositLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvokeContractRequest) GetStorageDepositLimit() DeployContractInkRequestStorageDepositLimit {
	if o == nil || IsNil(o.StorageDepositLimit.Get()) {
		var ret DeployContractInkRequestStorageDepositLimit
		return ret
	}
	return *o.StorageDepositLimit.Get()
}

// GetStorageDepositLimitOk returns a tuple with the StorageDepositLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvokeContractRequest) GetStorageDepositLimitOk() (*DeployContractInkRequestStorageDepositLimit, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDepositLimit.Get(), o.StorageDepositLimit.IsSet()
}

// HasStorageDepositLimit returns a boolean if a field has been set.
func (o *InvokeContractRequest) HasStorageDepositLimit() bool {
	if o != nil && o.StorageDepositLimit.IsSet() {
		return true
	}

	return false
}

// SetStorageDepositLimit gets a reference to the given NullableDeployContractInkRequestStorageDepositLimit and assigns it to the StorageDepositLimit field.
func (o *InvokeContractRequest) SetStorageDepositLimit(v DeployContractInkRequestStorageDepositLimit) {
	o.StorageDepositLimit.Set(&v)
}
// SetStorageDepositLimitNil sets the value for StorageDepositLimit to be an explicit nil
func (o *InvokeContractRequest) SetStorageDepositLimitNil() {
	o.StorageDepositLimit.Set(nil)
}

// UnsetStorageDepositLimit ensures that no value is present for StorageDepositLimit, not even an explicit nil
func (o *InvokeContractRequest) UnsetStorageDepositLimit() {
	o.StorageDepositLimit.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *InvokeContractRequest) GetBalance() DeployContractInkRequestBalance {
	if o == nil || IsNil(o.Balance) {
		var ret DeployContractInkRequestBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetBalanceOk() (*DeployContractInkRequestBalance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *InvokeContractRequest) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given DeployContractInkRequestBalance and assigns it to the Balance field.
func (o *InvokeContractRequest) SetBalance(v DeployContractInkRequestBalance) {
	o.Balance = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *InvokeContractRequest) GetParams() []interface{} {
	if o == nil || IsNil(o.Params) {
		var ret []interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractRequest) GetParamsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *InvokeContractRequest) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []interface{} and assigns it to the Params field.
func (o *InvokeContractRequest) SetParams(v []interface{}) {
	o.Params = v
}

func (o InvokeContractRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvokeContractRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invocationType"] = o.InvocationType
	toSerialize["accountAddress"] = o.AccountAddress
	toSerialize["web3SigningCredential"] = o.Web3SigningCredential
	toSerialize["metadata"] = o.Metadata
	toSerialize["contractAddress"] = o.ContractAddress
	toSerialize["methodName"] = o.MethodName
	toSerialize["gasLimit"] = o.GasLimit
	if o.StorageDepositLimit.IsSet() {
		toSerialize["storageDepositLimit"] = o.StorageDepositLimit.Get()
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableInvokeContractRequest struct {
	value *InvokeContractRequest
	isSet bool
}

func (v NullableInvokeContractRequest) Get() *InvokeContractRequest {
	return v.value
}

func (v *NullableInvokeContractRequest) Set(val *InvokeContractRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeContractRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeContractRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeContractRequest(val *InvokeContractRequest) *NullableInvokeContractRequest {
	return &NullableInvokeContractRequest{value: val, isSet: true}
}

func (v NullableInvokeContractRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeContractRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


