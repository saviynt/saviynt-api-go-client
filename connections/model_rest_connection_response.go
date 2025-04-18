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

// checks if the RESTConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RESTConnectionResponse{}

// RESTConnectionResponse struct for RESTConnectionResponse
type RESTConnectionResponse struct {
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
	Errorcode            *int32                    `json:"errorcode,omitempty"`
	Status               *int32                    `json:"status,omitempty"`
	Defaultsavroles      *string                   `json:"defaultsavroles,omitempty"`
	Connectionattributes *RESTConnectionAttributes `json:"connectionattributes,omitempty"`
}

// NewRESTConnectionResponse instantiates a new RESTConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRESTConnectionResponse() *RESTConnectionResponse {
	this := RESTConnectionResponse{}
	return &this
}

// NewRESTConnectionResponseWithDefaults instantiates a new RESTConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRESTConnectionResponseWithDefaults() *RESTConnectionResponse {
	this := RESTConnectionResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *RESTConnectionResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetEmailtemplate returns the Emailtemplate field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetEmailtemplate() string {
	if o == nil || IsNil(o.Emailtemplate) {
		var ret string
		return ret
	}
	return *o.Emailtemplate
}

// GetEmailtemplateOk returns a tuple with the Emailtemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetEmailtemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Emailtemplate) {
		return nil, false
	}
	return o.Emailtemplate, true
}

// HasEmailtemplate returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasEmailtemplate() bool {
	if o != nil && !IsNil(o.Emailtemplate) {
		return true
	}

	return false
}

// SetEmailtemplate gets a reference to the given string and assigns it to the Emailtemplate field.
func (o *RESTConnectionResponse) SetEmailtemplate(v string) {
	o.Emailtemplate = &v
}

// GetUpdatedby returns the Updatedby field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetUpdatedby() string {
	if o == nil || IsNil(o.Updatedby) {
		var ret string
		return ret
	}
	return *o.Updatedby
}

// GetUpdatedbyOk returns a tuple with the Updatedby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetUpdatedbyOk() (*string, bool) {
	if o == nil || IsNil(o.Updatedby) {
		return nil, false
	}
	return o.Updatedby, true
}

// HasUpdatedby returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasUpdatedby() bool {
	if o != nil && !IsNil(o.Updatedby) {
		return true
	}

	return false
}

// SetUpdatedby gets a reference to the given string and assigns it to the Updatedby field.
func (o *RESTConnectionResponse) SetUpdatedby(v string) {
	o.Updatedby = &v
}

// GetConnectionname returns the Connectionname field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetConnectionname() string {
	if o == nil || IsNil(o.Connectionname) {
		var ret string
		return ret
	}
	return *o.Connectionname
}

// GetConnectionnameOk returns a tuple with the Connectionname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetConnectionnameOk() (*string, bool) {
	if o == nil || IsNil(o.Connectionname) {
		return nil, false
	}
	return o.Connectionname, true
}

// HasConnectionname returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasConnectionname() bool {
	if o != nil && !IsNil(o.Connectionname) {
		return true
	}

	return false
}

// SetConnectionname gets a reference to the given string and assigns it to the Connectionname field.
func (o *RESTConnectionResponse) SetConnectionname(v string) {
	o.Connectionname = &v
}

// GetConnectionkey returns the Connectionkey field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetConnectionkey() int32 {
	if o == nil || IsNil(o.Connectionkey) {
		var ret int32
		return ret
	}
	return *o.Connectionkey
}

// GetConnectionkeyOk returns a tuple with the Connectionkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetConnectionkeyOk() (*int32, bool) {
	if o == nil || IsNil(o.Connectionkey) {
		return nil, false
	}
	return o.Connectionkey, true
}

// HasConnectionkey returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasConnectionkey() bool {
	if o != nil && !IsNil(o.Connectionkey) {
		return true
	}

	return false
}

// SetConnectionkey gets a reference to the given int32 and assigns it to the Connectionkey field.
func (o *RESTConnectionResponse) SetConnectionkey(v int32) {
	o.Connectionkey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RESTConnectionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetConnectiontype returns the Connectiontype field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetConnectiontype() string {
	if o == nil || IsNil(o.Connectiontype) {
		var ret string
		return ret
	}
	return *o.Connectiontype
}

// GetConnectiontypeOk returns a tuple with the Connectiontype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetConnectiontypeOk() (*string, bool) {
	if o == nil || IsNil(o.Connectiontype) {
		return nil, false
	}
	return o.Connectiontype, true
}

// HasConnectiontype returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasConnectiontype() bool {
	if o != nil && !IsNil(o.Connectiontype) {
		return true
	}

	return false
}

// SetConnectiontype gets a reference to the given string and assigns it to the Connectiontype field.
func (o *RESTConnectionResponse) SetConnectiontype(v string) {
	o.Connectiontype = &v
}

// GetCreatedon returns the Createdon field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetCreatedon() string {
	if o == nil || IsNil(o.Createdon) {
		var ret string
		return ret
	}
	return *o.Createdon
}

// GetCreatedonOk returns a tuple with the Createdon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetCreatedonOk() (*string, bool) {
	if o == nil || IsNil(o.Createdon) {
		return nil, false
	}
	return o.Createdon, true
}

// HasCreatedon returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasCreatedon() bool {
	if o != nil && !IsNil(o.Createdon) {
		return true
	}

	return false
}

// SetCreatedon gets a reference to the given string and assigns it to the Createdon field.
func (o *RESTConnectionResponse) SetCreatedon(v string) {
	o.Createdon = &v
}

// GetCreatedby returns the Createdby field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetCreatedby() string {
	if o == nil || IsNil(o.Createdby) {
		var ret string
		return ret
	}
	return *o.Createdby
}

// GetCreatedbyOk returns a tuple with the Createdby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetCreatedbyOk() (*string, bool) {
	if o == nil || IsNil(o.Createdby) {
		return nil, false
	}
	return o.Createdby, true
}

// HasCreatedby returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasCreatedby() bool {
	if o != nil && !IsNil(o.Createdby) {
		return true
	}

	return false
}

// SetCreatedby gets a reference to the given string and assigns it to the Createdby field.
func (o *RESTConnectionResponse) SetCreatedby(v string) {
	o.Createdby = &v
}

// GetErrorcode returns the Errorcode field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetErrorcode() int32 {
	if o == nil || IsNil(o.Errorcode) {
		var ret int32
		return ret
	}
	return *o.Errorcode
}

// GetErrorcodeOk returns a tuple with the Errorcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetErrorcodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Errorcode) {
		return nil, false
	}
	return o.Errorcode, true
}

// HasErrorcode returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasErrorcode() bool {
	if o != nil && !IsNil(o.Errorcode) {
		return true
	}

	return false
}

// SetErrorcode gets a reference to the given int32 and assigns it to the Errorcode field.
func (o *RESTConnectionResponse) SetErrorcode(v int32) {
	o.Errorcode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *RESTConnectionResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetDefaultsavroles returns the Defaultsavroles field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetDefaultsavroles() string {
	if o == nil || IsNil(o.Defaultsavroles) {
		var ret string
		return ret
	}
	return *o.Defaultsavroles
}

// GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetDefaultsavrolesOk() (*string, bool) {
	if o == nil || IsNil(o.Defaultsavroles) {
		return nil, false
	}
	return o.Defaultsavroles, true
}

// HasDefaultsavroles returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasDefaultsavroles() bool {
	if o != nil && !IsNil(o.Defaultsavroles) {
		return true
	}

	return false
}

// SetDefaultsavroles gets a reference to the given string and assigns it to the Defaultsavroles field.
func (o *RESTConnectionResponse) SetDefaultsavroles(v string) {
	o.Defaultsavroles = &v
}

// GetConnectionattributes returns the Connectionattributes field value if set, zero value otherwise.
func (o *RESTConnectionResponse) GetConnectionattributes() RESTConnectionAttributes {
	if o == nil || IsNil(o.Connectionattributes) {
		var ret RESTConnectionAttributes
		return ret
	}
	return *o.Connectionattributes
}

// GetConnectionattributesOk returns a tuple with the Connectionattributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTConnectionResponse) GetConnectionattributesOk() (*RESTConnectionAttributes, bool) {
	if o == nil || IsNil(o.Connectionattributes) {
		return nil, false
	}
	return o.Connectionattributes, true
}

// HasConnectionattributes returns a boolean if a field has been set.
func (o *RESTConnectionResponse) HasConnectionattributes() bool {
	if o != nil && !IsNil(o.Connectionattributes) {
		return true
	}

	return false
}

// SetConnectionattributes gets a reference to the given RESTConnectionAttributes and assigns it to the Connectionattributes field.
func (o *RESTConnectionResponse) SetConnectionattributes(v RESTConnectionAttributes) {
	o.Connectionattributes = &v
}

func (o RESTConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RESTConnectionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableRESTConnectionResponse struct {
	value *RESTConnectionResponse
	isSet bool
}

func (v NullableRESTConnectionResponse) Get() *RESTConnectionResponse {
	return v.value
}

func (v *NullableRESTConnectionResponse) Set(val *RESTConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRESTConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRESTConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRESTConnectionResponse(val *RESTConnectionResponse) *NullableRESTConnectionResponse {
	return &NullableRESTConnectionResponse{value: val, isSet: true}
}

func (v NullableRESTConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRESTConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
