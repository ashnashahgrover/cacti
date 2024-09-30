/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
	"time"
)

// checks if the ListCpiV1ResponseCpisInnerCpksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCpiV1ResponseCpisInnerCpksInner{}

// ListCpiV1ResponseCpisInnerCpksInner struct for ListCpiV1ResponseCpisInnerCpksInner
type ListCpiV1ResponseCpisInnerCpksInner struct {
	Hash *string `json:"hash,omitempty"`
	Id *CPIIDV1 `json:"id,omitempty"`
	Libraries []string `json:"libraries,omitempty"`
	MainBundle *string `json:"mainBundle,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewListCpiV1ResponseCpisInnerCpksInner instantiates a new ListCpiV1ResponseCpisInnerCpksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCpiV1ResponseCpisInnerCpksInner() *ListCpiV1ResponseCpisInnerCpksInner {
	this := ListCpiV1ResponseCpisInnerCpksInner{}
	return &this
}

// NewListCpiV1ResponseCpisInnerCpksInnerWithDefaults instantiates a new ListCpiV1ResponseCpisInnerCpksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCpiV1ResponseCpisInnerCpksInnerWithDefaults() *ListCpiV1ResponseCpisInnerCpksInner {
	this := ListCpiV1ResponseCpisInnerCpksInner{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetId() CPIIDV1 {
	if o == nil || IsNil(o.Id) {
		var ret CPIIDV1
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetIdOk() (*CPIIDV1, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given CPIIDV1 and assigns it to the Id field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetId(v CPIIDV1) {
	o.Id = &v
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetLibraries() []string {
	if o == nil || IsNil(o.Libraries) {
		var ret []string
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetLibrariesOk() ([]string, bool) {
	if o == nil || IsNil(o.Libraries) {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasLibraries() bool {
	if o != nil && !IsNil(o.Libraries) {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []string and assigns it to the Libraries field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetLibraries(v []string) {
	o.Libraries = v
}

// GetMainBundle returns the MainBundle field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetMainBundle() string {
	if o == nil || IsNil(o.MainBundle) {
		var ret string
		return ret
	}
	return *o.MainBundle
}

// GetMainBundleOk returns a tuple with the MainBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetMainBundleOk() (*string, bool) {
	if o == nil || IsNil(o.MainBundle) {
		return nil, false
	}
	return o.MainBundle, true
}

// HasMainBundle returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasMainBundle() bool {
	if o != nil && !IsNil(o.MainBundle) {
		return true
	}

	return false
}

// SetMainBundle gets a reference to the given string and assigns it to the MainBundle field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetMainBundle(v string) {
	o.MainBundle = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListCpiV1ResponseCpisInnerCpksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListCpiV1ResponseCpisInnerCpksInner) SetType(v string) {
	o.Type = &v
}

func (o ListCpiV1ResponseCpisInnerCpksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCpiV1ResponseCpisInnerCpksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Libraries) {
		toSerialize["libraries"] = o.Libraries
	}
	if !IsNil(o.MainBundle) {
		toSerialize["mainBundle"] = o.MainBundle
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableListCpiV1ResponseCpisInnerCpksInner struct {
	value *ListCpiV1ResponseCpisInnerCpksInner
	isSet bool
}

func (v NullableListCpiV1ResponseCpisInnerCpksInner) Get() *ListCpiV1ResponseCpisInnerCpksInner {
	return v.value
}

func (v *NullableListCpiV1ResponseCpisInnerCpksInner) Set(val *ListCpiV1ResponseCpisInnerCpksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListCpiV1ResponseCpisInnerCpksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListCpiV1ResponseCpisInnerCpksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCpiV1ResponseCpisInnerCpksInner(val *ListCpiV1ResponseCpisInnerCpksInner) *NullableListCpiV1ResponseCpisInnerCpksInner {
	return &NullableListCpiV1ResponseCpisInnerCpksInner{value: val, isSet: true}
}

func (v NullableListCpiV1ResponseCpisInnerCpksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCpiV1ResponseCpisInnerCpksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


