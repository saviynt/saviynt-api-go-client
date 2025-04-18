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

// checks if the DBConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DBConnectionResponse{}

// DBConnectionResponse struct for DBConnectionResponse
type DBConnectionResponse struct {
	// API response message
	Msg *string `json:"msg,omitempty"`
	// Email template for the connection
	Emailtemplate *string `json:"emailtemplate,omitempty"`
	// Updator account for the connection
	Updatedby *string `json:"updatedby,omitempty"`
	// Name of the connection
	Connectionname *string `json:"connectionname,omitempty"`
	// Connection key
	Connectionkey *int32 `json:"connectionkey,omitempty"`
	// Description for the connection
	Description *string `json:"description,omitempty"`
	// Connection type
	Connectiontype *string `json:"connectiontype,omitempty"`
	// Connection creation time
	Createdon *string `json:"createdon,omitempty"`
	// Creator account for the connection
	Createdby *string `json:"createdby,omitempty"`
	// Error code
	Errorcode            *int32                  `json:"errorcode,omitempty"`
	Status               *int32                  `json:"status,omitempty"`
	Defaultsavroles      *string                 `json:"defaultsavroles,omitempty"`
	Connectionattributes *DBConnectionAttributes `json:"connectionattributes,omitempty"`
}

// NewDBConnectionResponse instantiates a new DBConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBConnectionResponse() *DBConnectionResponse {
	this := DBConnectionResponse{}
	return &this
}

// NewDBConnectionResponseWithDefaults instantiates a new DBConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBConnectionResponseWithDefaults() *DBConnectionResponse {
	this := DBConnectionResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *DBConnectionResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetEmailtemplate returns the Emailtemplate field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetEmailtemplate() string {
	if o == nil || IsNil(o.Emailtemplate) {
		var ret string
		return ret
	}
	return *o.Emailtemplate
}

// GetEmailtemplateOk returns a tuple with the Emailtemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetEmailtemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Emailtemplate) {
		return nil, false
	}
	return o.Emailtemplate, true
}

// HasEmailtemplate returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasEmailtemplate() bool {
	if o != nil && !IsNil(o.Emailtemplate) {
		return true
	}

	return false
}

// SetEmailtemplate gets a reference to the given string and assigns it to the Emailtemplate field.
func (o *DBConnectionResponse) SetEmailtemplate(v string) {
	o.Emailtemplate = &v
}

// GetUpdatedby returns the Updatedby field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetUpdatedby() string {
	if o == nil || IsNil(o.Updatedby) {
		var ret string
		return ret
	}
	return *o.Updatedby
}

// GetUpdatedbyOk returns a tuple with the Updatedby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetUpdatedbyOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedby) {
		return nil, false
	}
	return o.Updatedby, true
}

// HasUpdatedby returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasUpdatedby() bool {
	if o != nil && !IsNil(o.Updatedby) {
		return true
	}

	return false
}

// SetUpdatedby gets a reference to the given string and assigns it to the Updatedby field.
func (o *DBConnectionResponse) SetUpdatedby(v string) {
	o.Updatedby = &v
}

// GetConnectionname returns the Connectionname field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetConnectionname() string {
	if o == nil || IsNil(o.Connectionname) {
		var ret string
		return ret
	}
	return *o.Connectionname
}

// GetConnectionnameOk returns a tuple with the Connectionname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetConnectionnameOk() (*string, bool) {
	if o == nil || IsNil(o.Connectionname) {
		return nil, false
	}
	return o.Connectionname, true
}

// HasConnectionname returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasConnectionname() bool {
	if o != nil && !IsNil(o.Connectionname) {
		return true
	}

	return false
}

// SetConnectionname gets a reference to the given string and assigns it to the Connectionname field.
func (o *DBConnectionResponse) SetConnectionname(v string) {
	o.Connectionname = &v
}

// GetConnectionkey returns the Connectionkey field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetConnectionkey() int32 {
	if o == nil || IsNil(o.Connectionkey) {
		var ret int32
		return ret
	}
	return *o.Connectionkey
}

// GetConnectionkeyOk returns a tuple with the Connectionkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetConnectionkeyOk() (*int32, bool) {
	if o == nil || IsNil(o.Connectionkey) {
		return nil, false
	}
	return o.Connectionkey, true
}

// HasConnectionkey returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasConnectionkey() bool {
	if o != nil && !IsNil(o.Connectionkey) {
		return true
	}

	return false
}

// SetConnectionkey gets a reference to the given int32 and assigns it to the Connectionkey field.
func (o *DBConnectionResponse) SetConnectionkey(v int32) {
	o.Connectionkey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DBConnectionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetConnectiontype returns the Connectiontype field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetConnectiontype() string {
	if o == nil || IsNil(o.Connectiontype) {
		var ret string
		return ret
	}
	return *o.Connectiontype
}

// GetConnectiontypeOk returns a tuple with the Connectiontype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetConnectiontypeOk() (*string, bool) {
	if o == nil || IsNil(o.Connectiontype) {
		return nil, false
	}
	return o.Connectiontype, true
}

// HasConnectiontype returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasConnectiontype() bool {
	if o != nil && !IsNil(o.Connectiontype) {
		return true
	}

	return false
}

// SetConnectiontype gets a reference to the given string and assigns it to the Connectiontype field.
func (o *DBConnectionResponse) SetConnectiontype(v string) {
	o.Connectiontype = &v
}

// GetCreatedon returns the Createdon field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetCreatedon() string {
	if o == nil || IsNil(o.Createdon) {
		var ret string
		return ret
	}
	return *o.Createdon
}

// GetCreatedonOk returns a tuple with the Createdon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetCreatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Createdon) {
		return nil, false
	}
	return o.Createdon, true
}

// HasCreatedon returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasCreatedon() bool {
	if o != nil && !IsNil(o.Createdon) {
		return true
	}

	return false
}

// SetCreatedon gets a reference to the given string and assigns it to the Createdon field.
func (o *DBConnectionResponse) SetCreatedon(v string) {
	o.Createdon = &v
}

// GetCreatedby returns the Createdby field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetCreatedby() string {
	if o == nil || IsNil(o.Createdby) {
		var ret string
		return ret
	}
	return *o.Createdby
}

// GetCreatedbyOk returns a tuple with the Createdby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetCreatedbyOk() (*string, bool) {
	if o == nil || IsNil(o.Createdby) {
		return nil, false
	}
	return o.Createdby, true
}

// HasCreatedby returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasCreatedby() bool {
	if o != nil && !IsNil(o.Createdby) {
		return true
	}

	return false
}

// SetCreatedby gets a reference to the given string and assigns it to the Createdby field.
func (o *DBConnectionResponse) SetCreatedby(v string) {
	o.Createdby = &v
}

// GetErrorcode returns the Errorcode field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetErrorcode() int32 {
	if o == nil || IsNil(o.Errorcode) {
		var ret int32
		return ret
	}
	return *o.Errorcode
}

// GetErrorcodeOk returns a tuple with the Errorcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetErrorcodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Errorcode) {
		return nil, false
	}
	return o.Errorcode, true
}

// HasErrorcode returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasErrorcode() bool {
	if o != nil && !IsNil(o.Errorcode) {
		return true
	}

	return false
}

// SetErrorcode gets a reference to the given int32 and assigns it to the Errorcode field.
func (o *DBConnectionResponse) SetErrorcode(v int32) {
	o.Errorcode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *DBConnectionResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetDefaultsavroles returns the Defaultsavroles field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetDefaultsavroles() string {
	if o == nil || IsNil(o.Defaultsavroles) {
		var ret string
		return ret
	}
	return *o.Defaultsavroles
}

// GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetDefaultsavrolesOk() (*string, bool) {
	if o == nil || IsNil(o.Defaultsavroles) {
		return nil, false
	}
	return o.Defaultsavroles, true
}

// HasDefaultsavroles returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasDefaultsavroles() bool {
	if o != nil && !IsNil(o.Defaultsavroles) {
		return true
	}

	return false
}

// SetDefaultsavroles gets a reference to the given string and assigns it to the Defaultsavroles field.
func (o *DBConnectionResponse) SetDefaultsavroles(v string) {
	o.Defaultsavroles = &v
}

// GetConnectionattributes returns the Connectionattributes field value if set, zero value otherwise.
func (o *DBConnectionResponse) GetConnectionattributes() DBConnectionAttributes {
	if o == nil || IsNil(o.Connectionattributes) {
		var ret DBConnectionAttributes
		return ret
	}
	return *o.Connectionattributes
}

// GetConnectionattributesOk returns a tuple with the Connectionattributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBConnectionResponse) GetConnectionattributesOk() (*DBConnectionAttributes, bool) {
	if o == nil || IsNil(o.Connectionattributes) {
		return nil, false
	}
	return o.Connectionattributes, true
}

// HasConnectionattributes returns a boolean if a field has been set.
func (o *DBConnectionResponse) HasConnectionattributes() bool {
	if o != nil && !IsNil(o.Connectionattributes) {
		return true
	}

	return false
}

// SetConnectionattributes gets a reference to the given DBConnectionAttributes and assigns it to the Connectionattributes field.
func (o *DBConnectionResponse) SetConnectionattributes(v DBConnectionAttributes) {
	o.Connectionattributes = &v
}

func (o DBConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DBConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Emailtemplate) {
		toSerialize["emailtemplate"] = o.Emailtemplate
	}
	if !IsNil(o.Updatedby) {
		toSerialize["updatedby"] = o.Updatedby
	}
	if !IsNil(o.Connectionname) {
		toSerialize["connectionname"] = o.Connectionname
	}
	if !IsNil(o.Connectionkey) {
		toSerialize["connectionkey"] = o.Connectionkey
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Connectiontype) {
		toSerialize["connectiontype"] = o.Connectiontype
	}
	if !IsNil(o.Createdon) {
		toSerialize["createdon"] = o.Createdon
	}
	if !IsNil(o.Createdby) {
		toSerialize["createdby"] = o.Createdby
	}
	if !IsNil(o.Errorcode) {
		toSerialize["errorcode"] = o.Errorcode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Defaultsavroles) {
		toSerialize["defaultsavroles"] = o.Defaultsavroles
	}
	if !IsNil(o.Connectionattributes) {
		toSerialize["connectionattributes"] = o.Connectionattributes
	}
	return toSerialize, nil
}

type NullableDBConnectionResponse struct {
	value *DBConnectionResponse
	isSet bool
}

func (v NullableDBConnectionResponse) Get() *DBConnectionResponse {
	return v.value
}

func (v *NullableDBConnectionResponse) Set(val *DBConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDBConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDBConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBConnectionResponse(val *DBConnectionResponse) *NullableDBConnectionResponse {
	return &NullableDBConnectionResponse{value: val, isSet: true}
}

func (v NullableDBConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
