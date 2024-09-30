/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
	"fmt"
)

// BlockTypeV1 Iroha V2 block response type.
type BlockTypeV1 string

// List of BlockTypeV1
const (
	Raw BlockTypeV1 = "raw"
	Binary BlockTypeV1 = "binary"
)

// All allowed values of BlockTypeV1 enum
var AllowedBlockTypeV1EnumValues = []BlockTypeV1{
	"raw",
	"binary",
}

func (v *BlockTypeV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BlockTypeV1(value)
	for _, existing := range AllowedBlockTypeV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BlockTypeV1", value)
}

// NewBlockTypeV1FromValue returns a pointer to a valid BlockTypeV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBlockTypeV1FromValue(v string) (*BlockTypeV1, error) {
	ev := BlockTypeV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BlockTypeV1: valid values are %v", v, AllowedBlockTypeV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BlockTypeV1) IsValid() bool {
	for _, existing := range AllowedBlockTypeV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlockTypeV1 value
func (v BlockTypeV1) Ptr() *BlockTypeV1 {
	return &v
}

type NullableBlockTypeV1 struct {
	value *BlockTypeV1
	isSet bool
}

func (v NullableBlockTypeV1) Get() *BlockTypeV1 {
	return v.value
}

func (v *NullableBlockTypeV1) Set(val *BlockTypeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTypeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTypeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTypeV1(val *BlockTypeV1) *NullableBlockTypeV1 {
	return &NullableBlockTypeV1{value: val, isSet: true}
}

func (v NullableBlockTypeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTypeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

