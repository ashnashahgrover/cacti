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

// FabricSigningCredentialType different type of identity provider for singing fabric messages supported by this package
type FabricSigningCredentialType string

// List of FabricSigningCredentialType
const (
	X_509 FabricSigningCredentialType = "X.509"
	VAULT_X_509 FabricSigningCredentialType = "Vault-X.509"
	WS_X_509 FabricSigningCredentialType = "WS-X.509"
)

// All allowed values of FabricSigningCredentialType enum
var AllowedFabricSigningCredentialTypeEnumValues = []FabricSigningCredentialType{
	"X.509",
	"Vault-X.509",
	"WS-X.509",
}

func (v *FabricSigningCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FabricSigningCredentialType(value)
	for _, existing := range AllowedFabricSigningCredentialTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FabricSigningCredentialType", value)
}

// NewFabricSigningCredentialTypeFromValue returns a pointer to a valid FabricSigningCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFabricSigningCredentialTypeFromValue(v string) (*FabricSigningCredentialType, error) {
	ev := FabricSigningCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FabricSigningCredentialType: valid values are %v", v, AllowedFabricSigningCredentialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FabricSigningCredentialType) IsValid() bool {
	for _, existing := range AllowedFabricSigningCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FabricSigningCredentialType value
func (v FabricSigningCredentialType) Ptr() *FabricSigningCredentialType {
	return &v
}

type NullableFabricSigningCredentialType struct {
	value *FabricSigningCredentialType
	isSet bool
}

func (v NullableFabricSigningCredentialType) Get() *FabricSigningCredentialType {
	return v.value
}

func (v *NullableFabricSigningCredentialType) Set(val *FabricSigningCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSigningCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSigningCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSigningCredentialType(val *FabricSigningCredentialType) *NullableFabricSigningCredentialType {
	return &NullableFabricSigningCredentialType{value: val, isSet: true}
}

func (v NullableFabricSigningCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSigningCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

