/*
Connection Management API

Use this API to create a connection in Saviynt Identity Cloud.  The Authorization header must have \"Bearer {token}\".  **Mandatory Parameters:** - **connectionname**: Specify the name to identify the connection. - **connectiontype**: Specify a connection type. For example, if your target application is Active Directory, specify the connection type as \"AD\".  **Optional Parameters:** - **description**: Provide a description for the connection. - **defaultsavroles**: Specify the SAV role(s) required for managing this connection along with its associated security systems, endpoints, accounts, and entitlements. - **emailTemplate**: Specify the email template applicable for notifications. - **sslCertificate**: Specify the SSL certificate(s) to secure the connection between EIC and the target application. - **vaultConfiguration**: Specify the path of the vault to obtain secret data (suffix the connector name to make it unique). - **saveinvault**: Set to true to save the encrypted attribute in the configured vault.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connections

import (
	"encoding/json"
)

// checks if the GetConnectionDetailsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConnectionDetailsRequest{}

// GetConnectionDetailsRequest struct for GetConnectionDetailsRequest
type GetConnectionDetailsRequest struct {
	// Name of the connection
	Connectionname *string `json:"connectionname,omitempty"`
	// Connection key
	Connectionkey *string `json:"connectionkey,omitempty"`
}

// NewGetConnectionDetailsRequest instantiates a new GetConnectionDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionDetailsRequest() *GetConnectionDetailsRequest {
	this := GetConnectionDetailsRequest{}
	return &this
}

// NewGetConnectionDetailsRequestWithDefaults instantiates a new GetConnectionDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionDetailsRequestWithDefaults() *GetConnectionDetailsRequest {
	this := GetConnectionDetailsRequest{}
	return &this
}

// GetConnectionname returns the Connectionname field value if set, zero value otherwise.
func (o *GetConnectionDetailsRequest) GetConnectionname() string {
	if o == nil || IsNil(o.Connectionname) {
		var ret string
		return ret
	}
	return *o.Connectionname
}

// GetConnectionnameOk returns a tuple with the Connectionname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionDetailsRequest) GetConnectionnameOk() (*string, bool) {
	if o == nil || IsNil(o.Connectionname) {
		return nil, false
	}
	return o.Connectionname, true
}

// HasConnectionname returns a boolean if a field has been set.
func (o *GetConnectionDetailsRequest) HasConnectionname() bool {
	if o != nil && !IsNil(o.Connectionname) {
		return true
	}

	return false
}

// SetConnectionname gets a reference to the given string and assigns it to the Connectionname field.
func (o *GetConnectionDetailsRequest) SetConnectionname(v string) {
	o.Connectionname = &v
}

// GetConnectionkey returns the Connectionkey field value if set, zero value otherwise.
func (o *GetConnectionDetailsRequest) GetConnectionkey() string {
	if o == nil || IsNil(o.Connectionkey) {
		var ret string
		return ret
	}
	return *o.Connectionkey
}

// GetConnectionkeyOk returns a tuple with the Connectionkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionDetailsRequest) GetConnectionkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Connectionkey) {
		return nil, false
	}
	return o.Connectionkey, true
}

// HasConnectionkey returns a boolean if a field has been set.
func (o *GetConnectionDetailsRequest) HasConnectionkey() bool {
	if o != nil && !IsNil(o.Connectionkey) {
		return true
	}

	return false
}

// SetConnectionkey gets a reference to the given string and assigns it to the Connectionkey field.
func (o *GetConnectionDetailsRequest) SetConnectionkey(v string) {
	o.Connectionkey = &v
}

func (o GetConnectionDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConnectionDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connectionname) {
		toSerialize["connectionname"] = o.Connectionname
	}
	if !IsNil(o.Connectionkey) {
		toSerialize["connectionkey"] = o.Connectionkey
	}
	return toSerialize, nil
}

type NullableGetConnectionDetailsRequest struct {
	value *GetConnectionDetailsRequest
	isSet bool
}

func (v NullableGetConnectionDetailsRequest) Get() *GetConnectionDetailsRequest {
	return v.value
}

func (v *NullableGetConnectionDetailsRequest) Set(val *GetConnectionDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionDetailsRequest(val *GetConnectionDetailsRequest) *NullableGetConnectionDetailsRequest {
	return &NullableGetConnectionDetailsRequest{value: val, isSet: true}
}

func (v NullableGetConnectionDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
