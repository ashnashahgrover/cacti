/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
	"fmt"
)

// ReceiptType Enumerates the possible types of receipts that can be waited for by someone or something that has requested the execution of a transaction on a ledger.
type ReceiptType string

// List of ReceiptType
const (
	NODE_TX_POOL_ACK ReceiptType = "NODE_TX_POOL_ACK"
	LEDGER_BLOCK_ACK ReceiptType = "LEDGER_BLOCK_ACK"
)

// All allowed values of ReceiptType enum
var AllowedReceiptTypeEnumValues = []ReceiptType{
	"NODE_TX_POOL_ACK",
	"LEDGER_BLOCK_ACK",
}

func (v *ReceiptType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReceiptType(value)
	for _, existing := range AllowedReceiptTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReceiptType", value)
}

// NewReceiptTypeFromValue returns a pointer to a valid ReceiptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReceiptTypeFromValue(v string) (*ReceiptType, error) {
	ev := ReceiptType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReceiptType: valid values are %v", v, AllowedReceiptTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReceiptType) IsValid() bool {
	for _, existing := range AllowedReceiptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReceiptType value
func (v ReceiptType) Ptr() *ReceiptType {
	return &v
}

type NullableReceiptType struct {
	value *ReceiptType
	isSet bool
}

func (v NullableReceiptType) Get() *ReceiptType {
	return v.value
}

func (v *NullableReceiptType) Set(val *ReceiptType) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptType) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptType(val *ReceiptType) *NullableReceiptType {
	return &NullableReceiptType{value: val, isSet: true}
}

func (v NullableReceiptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

