/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
	"fmt"
)

// FabricContractInvocationType the model 'FabricContractInvocationType'
type FabricContractInvocationType string

// List of FabricContractInvocationType
const (
	SEND FabricContractInvocationType = "FabricContractInvocationType.SEND"
	CALL FabricContractInvocationType = "FabricContractInvocationType.CALL"
	SENDPRIVATE FabricContractInvocationType = "FabricContractInvocationType.SENDPRIVATE"
)

// All allowed values of FabricContractInvocationType enum
var AllowedFabricContractInvocationTypeEnumValues = []FabricContractInvocationType{
	"FabricContractInvocationType.SEND",
	"FabricContractInvocationType.CALL",
	"FabricContractInvocationType.SENDPRIVATE",
}

func (v *FabricContractInvocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FabricContractInvocationType(value)
	for _, existing := range AllowedFabricContractInvocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FabricContractInvocationType", value)
}

// NewFabricContractInvocationTypeFromValue returns a pointer to a valid FabricContractInvocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFabricContractInvocationTypeFromValue(v string) (*FabricContractInvocationType, error) {
	ev := FabricContractInvocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FabricContractInvocationType: valid values are %v", v, AllowedFabricContractInvocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FabricContractInvocationType) IsValid() bool {
	for _, existing := range AllowedFabricContractInvocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FabricContractInvocationType value
func (v FabricContractInvocationType) Ptr() *FabricContractInvocationType {
	return &v
}

type NullableFabricContractInvocationType struct {
	value *FabricContractInvocationType
	isSet bool
}

func (v NullableFabricContractInvocationType) Get() *FabricContractInvocationType {
	return v.value
}

func (v *NullableFabricContractInvocationType) Set(val *FabricContractInvocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricContractInvocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricContractInvocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricContractInvocationType(val *FabricContractInvocationType) *NullableFabricContractInvocationType {
	return &NullableFabricContractInvocationType{value: val, isSet: true}
}

func (v NullableFabricContractInvocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricContractInvocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

