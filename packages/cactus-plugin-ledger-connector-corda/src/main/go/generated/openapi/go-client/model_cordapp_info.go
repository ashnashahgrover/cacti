/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the CordappInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CordappInfo{}

// CordappInfo A CordappInfo describes a single CorDapp currently installed on the node
type CordappInfo struct {
	JarHash SHA256 `json:"jarHash"`
	// The name of the licence this CorDapp is released under
	Licence string `json:"licence"`
	// The minimum platform version the node must be at for the CorDapp to run
	MinimumPlatformVersion int32 `json:"minimumPlatformVersion"`
	// The name of the JAR file that defines the CorDapp
	Name string `json:"name"`
	// The name of the CorDapp
	ShortName string `json:"shortName"`
	// The target platform version this CorDapp has been tested against
	TargetPlatformVersion int32 `json:"targetPlatformVersion"`
	// A description of what sort of CorDapp this is - either a contract, workflow, or a combination.
	Type string `json:"type"`
	// The vendor of this CorDapp
	Vendor string `json:"vendor"`
	// The version of this CorDapp
	Version string `json:"version"`
}

// NewCordappInfo instantiates a new CordappInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCordappInfo(jarHash SHA256, licence string, minimumPlatformVersion int32, name string, shortName string, targetPlatformVersion int32, type_ string, vendor string, version string) *CordappInfo {
	this := CordappInfo{}
	this.JarHash = jarHash
	this.Licence = licence
	this.MinimumPlatformVersion = minimumPlatformVersion
	this.Name = name
	this.ShortName = shortName
	this.TargetPlatformVersion = targetPlatformVersion
	this.Type = type_
	this.Vendor = vendor
	this.Version = version
	return &this
}

// NewCordappInfoWithDefaults instantiates a new CordappInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCordappInfoWithDefaults() *CordappInfo {
	this := CordappInfo{}
	return &this
}

// GetJarHash returns the JarHash field value
func (o *CordappInfo) GetJarHash() SHA256 {
	if o == nil {
		var ret SHA256
		return ret
	}

	return o.JarHash
}

// GetJarHashOk returns a tuple with the JarHash field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetJarHashOk() (SHA256, bool) {
	if o == nil {
		return SHA256{}, false
	}
	return o.JarHash, true
}

// SetJarHash sets field value
func (o *CordappInfo) SetJarHash(v SHA256) {
	o.JarHash = v
}

// GetLicence returns the Licence field value
func (o *CordappInfo) GetLicence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Licence
}

// GetLicenceOk returns a tuple with the Licence field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetLicenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Licence, true
}

// SetLicence sets field value
func (o *CordappInfo) SetLicence(v string) {
	o.Licence = v
}

// GetMinimumPlatformVersion returns the MinimumPlatformVersion field value
func (o *CordappInfo) GetMinimumPlatformVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinimumPlatformVersion
}

// GetMinimumPlatformVersionOk returns a tuple with the MinimumPlatformVersion field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetMinimumPlatformVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumPlatformVersion, true
}

// SetMinimumPlatformVersion sets field value
func (o *CordappInfo) SetMinimumPlatformVersion(v int32) {
	o.MinimumPlatformVersion = v
}

// GetName returns the Name field value
func (o *CordappInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CordappInfo) SetName(v string) {
	o.Name = v
}

// GetShortName returns the ShortName field value
func (o *CordappInfo) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *CordappInfo) SetShortName(v string) {
	o.ShortName = v
}

// GetTargetPlatformVersion returns the TargetPlatformVersion field value
func (o *CordappInfo) GetTargetPlatformVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetPlatformVersion
}

// GetTargetPlatformVersionOk returns a tuple with the TargetPlatformVersion field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetTargetPlatformVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPlatformVersion, true
}

// SetTargetPlatformVersion sets field value
func (o *CordappInfo) SetTargetPlatformVersion(v int32) {
	o.TargetPlatformVersion = v
}

// GetType returns the Type field value
func (o *CordappInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CordappInfo) SetType(v string) {
	o.Type = v
}

// GetVendor returns the Vendor field value
func (o *CordappInfo) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *CordappInfo) SetVendor(v string) {
	o.Vendor = v
}

// GetVersion returns the Version field value
func (o *CordappInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CordappInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CordappInfo) SetVersion(v string) {
	o.Version = v
}

func (o CordappInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CordappInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jarHash"] = o.JarHash
	toSerialize["licence"] = o.Licence
	toSerialize["minimumPlatformVersion"] = o.MinimumPlatformVersion
	toSerialize["name"] = o.Name
	toSerialize["shortName"] = o.ShortName
	toSerialize["targetPlatformVersion"] = o.TargetPlatformVersion
	toSerialize["type"] = o.Type
	toSerialize["vendor"] = o.Vendor
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableCordappInfo struct {
	value *CordappInfo
	isSet bool
}

func (v NullableCordappInfo) Get() *CordappInfo {
	return v.value
}

func (v *NullableCordappInfo) Set(val *CordappInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCordappInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCordappInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCordappInfo(val *CordappInfo) *NullableCordappInfo {
	return &NullableCordappInfo{value: val, isSet: true}
}

func (v NullableCordappInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCordappInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


