/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the AssetProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetProfile{}

// AssetProfile struct for AssetProfile
type AssetProfile struct {
	Issuer *string `json:"issuer,omitempty"`
	AssetCode *string `json:"assetCode,omitempty"`
	AssetCodeType *string `json:"assetCodeType,omitempty"`
	IssuanceDate *string `json:"issuanceDate,omitempty"`
	ExpirationDate string `json:"expirationDate"`
	VerificationEndPoint *string `json:"verificationEndPoint,omitempty"`
	DigitalSignature *string `json:"digitalSignature,omitempty"`
	ProspectusLink *string `json:"prospectusLink,omitempty"`
	KeyInformationLink []interface{} `json:"keyInformationLink,omitempty"`
	KeyWord []interface{} `json:"keyWord,omitempty"`
	TransferRestriction []interface{} `json:"transferRestriction,omitempty"`
	LedgerRequirements []interface{} `json:"ledgerRequirements,omitempty"`
}

// NewAssetProfile instantiates a new AssetProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetProfile(expirationDate string) *AssetProfile {
	this := AssetProfile{}
	this.ExpirationDate = expirationDate
	return &this
}

// NewAssetProfileWithDefaults instantiates a new AssetProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetProfileWithDefaults() *AssetProfile {
	this := AssetProfile{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *AssetProfile) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *AssetProfile) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *AssetProfile) SetIssuer(v string) {
	o.Issuer = &v
}

// GetAssetCode returns the AssetCode field value if set, zero value otherwise.
func (o *AssetProfile) GetAssetCode() string {
	if o == nil || IsNil(o.AssetCode) {
		var ret string
		return ret
	}
	return *o.AssetCode
}

// GetAssetCodeOk returns a tuple with the AssetCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetAssetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetCode) {
		return nil, false
	}
	return o.AssetCode, true
}

// HasAssetCode returns a boolean if a field has been set.
func (o *AssetProfile) HasAssetCode() bool {
	if o != nil && !IsNil(o.AssetCode) {
		return true
	}

	return false
}

// SetAssetCode gets a reference to the given string and assigns it to the AssetCode field.
func (o *AssetProfile) SetAssetCode(v string) {
	o.AssetCode = &v
}

// GetAssetCodeType returns the AssetCodeType field value if set, zero value otherwise.
func (o *AssetProfile) GetAssetCodeType() string {
	if o == nil || IsNil(o.AssetCodeType) {
		var ret string
		return ret
	}
	return *o.AssetCodeType
}

// GetAssetCodeTypeOk returns a tuple with the AssetCodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetAssetCodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetCodeType) {
		return nil, false
	}
	return o.AssetCodeType, true
}

// HasAssetCodeType returns a boolean if a field has been set.
func (o *AssetProfile) HasAssetCodeType() bool {
	if o != nil && !IsNil(o.AssetCodeType) {
		return true
	}

	return false
}

// SetAssetCodeType gets a reference to the given string and assigns it to the AssetCodeType field.
func (o *AssetProfile) SetAssetCodeType(v string) {
	o.AssetCodeType = &v
}

// GetIssuanceDate returns the IssuanceDate field value if set, zero value otherwise.
func (o *AssetProfile) GetIssuanceDate() string {
	if o == nil || IsNil(o.IssuanceDate) {
		var ret string
		return ret
	}
	return *o.IssuanceDate
}

// GetIssuanceDateOk returns a tuple with the IssuanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetIssuanceDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssuanceDate) {
		return nil, false
	}
	return o.IssuanceDate, true
}

// HasIssuanceDate returns a boolean if a field has been set.
func (o *AssetProfile) HasIssuanceDate() bool {
	if o != nil && !IsNil(o.IssuanceDate) {
		return true
	}

	return false
}

// SetIssuanceDate gets a reference to the given string and assigns it to the IssuanceDate field.
func (o *AssetProfile) SetIssuanceDate(v string) {
	o.IssuanceDate = &v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *AssetProfile) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *AssetProfile) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetVerificationEndPoint returns the VerificationEndPoint field value if set, zero value otherwise.
func (o *AssetProfile) GetVerificationEndPoint() string {
	if o == nil || IsNil(o.VerificationEndPoint) {
		var ret string
		return ret
	}
	return *o.VerificationEndPoint
}

// GetVerificationEndPointOk returns a tuple with the VerificationEndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetVerificationEndPointOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationEndPoint) {
		return nil, false
	}
	return o.VerificationEndPoint, true
}

// HasVerificationEndPoint returns a boolean if a field has been set.
func (o *AssetProfile) HasVerificationEndPoint() bool {
	if o != nil && !IsNil(o.VerificationEndPoint) {
		return true
	}

	return false
}

// SetVerificationEndPoint gets a reference to the given string and assigns it to the VerificationEndPoint field.
func (o *AssetProfile) SetVerificationEndPoint(v string) {
	o.VerificationEndPoint = &v
}

// GetDigitalSignature returns the DigitalSignature field value if set, zero value otherwise.
func (o *AssetProfile) GetDigitalSignature() string {
	if o == nil || IsNil(o.DigitalSignature) {
		var ret string
		return ret
	}
	return *o.DigitalSignature
}

// GetDigitalSignatureOk returns a tuple with the DigitalSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetDigitalSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.DigitalSignature) {
		return nil, false
	}
	return o.DigitalSignature, true
}

// HasDigitalSignature returns a boolean if a field has been set.
func (o *AssetProfile) HasDigitalSignature() bool {
	if o != nil && !IsNil(o.DigitalSignature) {
		return true
	}

	return false
}

// SetDigitalSignature gets a reference to the given string and assigns it to the DigitalSignature field.
func (o *AssetProfile) SetDigitalSignature(v string) {
	o.DigitalSignature = &v
}

// GetProspectusLink returns the ProspectusLink field value if set, zero value otherwise.
func (o *AssetProfile) GetProspectusLink() string {
	if o == nil || IsNil(o.ProspectusLink) {
		var ret string
		return ret
	}
	return *o.ProspectusLink
}

// GetProspectusLinkOk returns a tuple with the ProspectusLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetProspectusLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ProspectusLink) {
		return nil, false
	}
	return o.ProspectusLink, true
}

// HasProspectusLink returns a boolean if a field has been set.
func (o *AssetProfile) HasProspectusLink() bool {
	if o != nil && !IsNil(o.ProspectusLink) {
		return true
	}

	return false
}

// SetProspectusLink gets a reference to the given string and assigns it to the ProspectusLink field.
func (o *AssetProfile) SetProspectusLink(v string) {
	o.ProspectusLink = &v
}

// GetKeyInformationLink returns the KeyInformationLink field value if set, zero value otherwise.
func (o *AssetProfile) GetKeyInformationLink() []interface{} {
	if o == nil || IsNil(o.KeyInformationLink) {
		var ret []interface{}
		return ret
	}
	return o.KeyInformationLink
}

// GetKeyInformationLinkOk returns a tuple with the KeyInformationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetKeyInformationLinkOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.KeyInformationLink) {
		return nil, false
	}
	return o.KeyInformationLink, true
}

// HasKeyInformationLink returns a boolean if a field has been set.
func (o *AssetProfile) HasKeyInformationLink() bool {
	if o != nil && !IsNil(o.KeyInformationLink) {
		return true
	}

	return false
}

// SetKeyInformationLink gets a reference to the given []interface{} and assigns it to the KeyInformationLink field.
func (o *AssetProfile) SetKeyInformationLink(v []interface{}) {
	o.KeyInformationLink = v
}

// GetKeyWord returns the KeyWord field value if set, zero value otherwise.
func (o *AssetProfile) GetKeyWord() []interface{} {
	if o == nil || IsNil(o.KeyWord) {
		var ret []interface{}
		return ret
	}
	return o.KeyWord
}

// GetKeyWordOk returns a tuple with the KeyWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetKeyWordOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.KeyWord) {
		return nil, false
	}
	return o.KeyWord, true
}

// HasKeyWord returns a boolean if a field has been set.
func (o *AssetProfile) HasKeyWord() bool {
	if o != nil && !IsNil(o.KeyWord) {
		return true
	}

	return false
}

// SetKeyWord gets a reference to the given []interface{} and assigns it to the KeyWord field.
func (o *AssetProfile) SetKeyWord(v []interface{}) {
	o.KeyWord = v
}

// GetTransferRestriction returns the TransferRestriction field value if set, zero value otherwise.
func (o *AssetProfile) GetTransferRestriction() []interface{} {
	if o == nil || IsNil(o.TransferRestriction) {
		var ret []interface{}
		return ret
	}
	return o.TransferRestriction
}

// GetTransferRestrictionOk returns a tuple with the TransferRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetTransferRestrictionOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.TransferRestriction) {
		return nil, false
	}
	return o.TransferRestriction, true
}

// HasTransferRestriction returns a boolean if a field has been set.
func (o *AssetProfile) HasTransferRestriction() bool {
	if o != nil && !IsNil(o.TransferRestriction) {
		return true
	}

	return false
}

// SetTransferRestriction gets a reference to the given []interface{} and assigns it to the TransferRestriction field.
func (o *AssetProfile) SetTransferRestriction(v []interface{}) {
	o.TransferRestriction = v
}

// GetLedgerRequirements returns the LedgerRequirements field value if set, zero value otherwise.
func (o *AssetProfile) GetLedgerRequirements() []interface{} {
	if o == nil || IsNil(o.LedgerRequirements) {
		var ret []interface{}
		return ret
	}
	return o.LedgerRequirements
}

// GetLedgerRequirementsOk returns a tuple with the LedgerRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProfile) GetLedgerRequirementsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.LedgerRequirements) {
		return nil, false
	}
	return o.LedgerRequirements, true
}

// HasLedgerRequirements returns a boolean if a field has been set.
func (o *AssetProfile) HasLedgerRequirements() bool {
	if o != nil && !IsNil(o.LedgerRequirements) {
		return true
	}

	return false
}

// SetLedgerRequirements gets a reference to the given []interface{} and assigns it to the LedgerRequirements field.
func (o *AssetProfile) SetLedgerRequirements(v []interface{}) {
	o.LedgerRequirements = v
}

func (o AssetProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.AssetCode) {
		toSerialize["assetCode"] = o.AssetCode
	}
	if !IsNil(o.AssetCodeType) {
		toSerialize["assetCodeType"] = o.AssetCodeType
	}
	if !IsNil(o.IssuanceDate) {
		toSerialize["issuanceDate"] = o.IssuanceDate
	}
	toSerialize["expirationDate"] = o.ExpirationDate
	if !IsNil(o.VerificationEndPoint) {
		toSerialize["verificationEndPoint"] = o.VerificationEndPoint
	}
	if !IsNil(o.DigitalSignature) {
		toSerialize["digitalSignature"] = o.DigitalSignature
	}
	if !IsNil(o.ProspectusLink) {
		toSerialize["prospectusLink"] = o.ProspectusLink
	}
	if !IsNil(o.KeyInformationLink) {
		toSerialize["keyInformationLink"] = o.KeyInformationLink
	}
	if !IsNil(o.KeyWord) {
		toSerialize["keyWord"] = o.KeyWord
	}
	if !IsNil(o.TransferRestriction) {
		toSerialize["transferRestriction"] = o.TransferRestriction
	}
	if !IsNil(o.LedgerRequirements) {
		toSerialize["ledgerRequirements"] = o.LedgerRequirements
	}
	return toSerialize, nil
}

type NullableAssetProfile struct {
	value *AssetProfile
	isSet bool
}

func (v NullableAssetProfile) Get() *AssetProfile {
	return v.value
}

func (v *NullableAssetProfile) Set(val *AssetProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetProfile(val *AssetProfile) *NullableAssetProfile {
	return &NullableAssetProfile{value: val, isSet: true}
}

func (v NullableAssetProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


