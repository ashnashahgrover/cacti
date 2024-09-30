/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
	"fmt"
)

// WatchBlocksV1 the model 'WatchBlocksV1'
type WatchBlocksV1 string

// List of WatchBlocksV1
const (
	Subscribe WatchBlocksV1 = "org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Subscribe"
	Next WatchBlocksV1 = "org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Next"
	Unsubscribe WatchBlocksV1 = "org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Unsubscribe"
	Error WatchBlocksV1 = "org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Error"
	Complete WatchBlocksV1 = "org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Complete"
)

// All allowed values of WatchBlocksV1 enum
var AllowedWatchBlocksV1EnumValues = []WatchBlocksV1{
	"org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Subscribe",
	"org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Next",
	"org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Unsubscribe",
	"org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Error",
	"org.hyperledger.cacti.api.async.ethereum.WatchBlocksV1.Complete",
}

func (v *WatchBlocksV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WatchBlocksV1(value)
	for _, existing := range AllowedWatchBlocksV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WatchBlocksV1", value)
}

// NewWatchBlocksV1FromValue returns a pointer to a valid WatchBlocksV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatchBlocksV1FromValue(v string) (*WatchBlocksV1, error) {
	ev := WatchBlocksV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WatchBlocksV1: valid values are %v", v, AllowedWatchBlocksV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatchBlocksV1) IsValid() bool {
	for _, existing := range AllowedWatchBlocksV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatchBlocksV1 value
func (v WatchBlocksV1) Ptr() *WatchBlocksV1 {
	return &v
}

type NullableWatchBlocksV1 struct {
	value *WatchBlocksV1
	isSet bool
}

func (v NullableWatchBlocksV1) Get() *WatchBlocksV1 {
	return v.value
}

func (v *NullableWatchBlocksV1) Set(val *WatchBlocksV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1(val *WatchBlocksV1) *NullableWatchBlocksV1 {
	return &NullableWatchBlocksV1{value: val, isSet: true}
}

func (v NullableWatchBlocksV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

