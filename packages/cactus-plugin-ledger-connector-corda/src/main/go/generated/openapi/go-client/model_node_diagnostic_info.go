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

// checks if the NodeDiagnosticInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeDiagnosticInfo{}

// NodeDiagnosticInfo A NodeDiagnosticInfo holds information about the current node version.
type NodeDiagnosticInfo struct {
	// A list of CorDapps currently installed on this node
	Cordapps []CordappInfo `json:"cordapps"`
	// The platform version of this node. This number represents a released API version, and should be used to make functionality decisions (e.g. enabling an app feature only if an underlying platform feature exists)
	PlatformVersion int32 `json:"platformVersion"`
	// The git commit hash this node was built from
	Revision string `json:"revision"`
	// The vendor of this node
	Vendor string `json:"vendor"`
	// The current node version string, e.g. 4.3, 4.4-SNAPSHOT. Note that this string is effectively freeform, and so should only be used for providing diagnostic information. It should not be used to make functionality decisions (the platformVersion is a better fit for this).
	Version string `json:"version"`
}

// NewNodeDiagnosticInfo instantiates a new NodeDiagnosticInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeDiagnosticInfo(cordapps []CordappInfo, platformVersion int32, revision string, vendor string, version string) *NodeDiagnosticInfo {
	this := NodeDiagnosticInfo{}
	this.Cordapps = cordapps
	this.PlatformVersion = platformVersion
	this.Revision = revision
	this.Vendor = vendor
	this.Version = version
	return &this
}

// NewNodeDiagnosticInfoWithDefaults instantiates a new NodeDiagnosticInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDiagnosticInfoWithDefaults() *NodeDiagnosticInfo {
	this := NodeDiagnosticInfo{}
	return &this
}

// GetCordapps returns the Cordapps field value
func (o *NodeDiagnosticInfo) GetCordapps() []CordappInfo {
	if o == nil {
		var ret []CordappInfo
		return ret
	}

	return o.Cordapps
}

// GetCordappsOk returns a tuple with the Cordapps field value
// and a boolean to check if the value has been set.
func (o *NodeDiagnosticInfo) GetCordappsOk() ([]CordappInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cordapps, true
}

// SetCordapps sets field value
func (o *NodeDiagnosticInfo) SetCordapps(v []CordappInfo) {
	o.Cordapps = v
}

// GetPlatformVersion returns the PlatformVersion field value
func (o *NodeDiagnosticInfo) GetPlatformVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlatformVersion
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field value
// and a boolean to check if the value has been set.
func (o *NodeDiagnosticInfo) GetPlatformVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformVersion, true
}

// SetPlatformVersion sets field value
func (o *NodeDiagnosticInfo) SetPlatformVersion(v int32) {
	o.PlatformVersion = v
}

// GetRevision returns the Revision field value
func (o *NodeDiagnosticInfo) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *NodeDiagnosticInfo) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *NodeDiagnosticInfo) SetRevision(v string) {
	o.Revision = v
}

// GetVendor returns the Vendor field value
func (o *NodeDiagnosticInfo) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *NodeDiagnosticInfo) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *NodeDiagnosticInfo) SetVendor(v string) {
	o.Vendor = v
}

// GetVersion returns the Version field value
func (o *NodeDiagnosticInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NodeDiagnosticInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NodeDiagnosticInfo) SetVersion(v string) {
	o.Version = v
}

func (o NodeDiagnosticInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeDiagnosticInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cordapps"] = o.Cordapps
	toSerialize["platformVersion"] = o.PlatformVersion
	toSerialize["revision"] = o.Revision
	toSerialize["vendor"] = o.Vendor
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableNodeDiagnosticInfo struct {
	value *NodeDiagnosticInfo
	isSet bool
}

func (v NullableNodeDiagnosticInfo) Get() *NodeDiagnosticInfo {
	return v.value
}

func (v *NullableNodeDiagnosticInfo) Set(val *NodeDiagnosticInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeDiagnosticInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeDiagnosticInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeDiagnosticInfo(val *NodeDiagnosticInfo) *NullableNodeDiagnosticInfo {
	return &NullableNodeDiagnosticInfo{value: val, isSet: true}
}

func (v NullableNodeDiagnosticInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeDiagnosticInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


