/*
Saviynt mTLS Authentication API

mTLS Authentication

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mtlsauthentication

import (
	"encoding/json"
)

// checks if the CertificateThumbprint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateThumbprint{}

// CertificateThumbprint struct for CertificateThumbprint
type CertificateThumbprint struct {
	Expiry     *string `json:"expiry,omitempty"`
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// NewCertificateThumbprint instantiates a new CertificateThumbprint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateThumbprint() *CertificateThumbprint {
	this := CertificateThumbprint{}
	return &this
}

// NewCertificateThumbprintWithDefaults instantiates a new CertificateThumbprint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateThumbprintWithDefaults() *CertificateThumbprint {
	this := CertificateThumbprint{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *CertificateThumbprint) GetExpiry() string {
	if o == nil || IsNil(o.Expiry) {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateThumbprint) GetExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *CertificateThumbprint) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *CertificateThumbprint) SetExpiry(v string) {
	o.Expiry = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *CertificateThumbprint) GetThumbprint() string {
	if o == nil || IsNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateThumbprint) GetThumbprintOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *CertificateThumbprint) HasThumbprint() bool {
	if o != nil && !IsNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *CertificateThumbprint) SetThumbprint(v string) {
	o.Thumbprint = &v
}

func (o CertificateThumbprint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateThumbprint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Thumbprint) {
		toSerialize["thumbprint"] = o.Thumbprint
	}
	return toSerialize, nil
}

type NullableCertificateThumbprint struct {
	value *CertificateThumbprint
	isSet bool
}

func (v NullableCertificateThumbprint) Get() *CertificateThumbprint {
	return v.value
}

func (v *NullableCertificateThumbprint) Set(val *CertificateThumbprint) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateThumbprint) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateThumbprint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateThumbprint(val *CertificateThumbprint) *NullableCertificateThumbprint {
	return &NullableCertificateThumbprint{value: val, isSet: true}
}

func (v NullableCertificateThumbprint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateThumbprint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
