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

// checks if the GetConnectionsResponseConnectionListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConnectionsResponseConnectionListInner{}

// GetConnectionsResponseConnectionListInner struct for GetConnectionsResponseConnectionListInner
type GetConnectionsResponseConnectionListInner struct {
	// Name of the Connection
	CONNECTIONNAME *string `json:"CONNECTIONNAME,omitempty"`
	// Type of the Connection
	CONNECTIONTYPE *string `json:"CONNECTIONTYPE,omitempty"`
	// Connection description
	CONNECTIONDESCRIPTION *string `json:"CONNECTIONDESCRIPTION,omitempty"`
	// Connection status
	STATUS *int32 `json:"STATUS,omitempty"`
	// User who created the connection
	CREATEDBY *string `json:"CREATEDBY,omitempty"`
	// Connection creation time
	CREATEDON *string `json:"CREATEDON,omitempty"`
	// User who updated the connection
	UPDATEDBY *string `json:"UPDATEDBY,omitempty"`
	// Time of connection updation
	UPDATEDON *string `json:"UPDATEDON,omitempty"`
}

// NewGetConnectionsResponseConnectionListInner instantiates a new GetConnectionsResponseConnectionListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConnectionsResponseConnectionListInner() *GetConnectionsResponseConnectionListInner {
	this := GetConnectionsResponseConnectionListInner{}
	return &this
}

// NewGetConnectionsResponseConnectionListInnerWithDefaults instantiates a new GetConnectionsResponseConnectionListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConnectionsResponseConnectionListInnerWithDefaults() *GetConnectionsResponseConnectionListInner {
	this := GetConnectionsResponseConnectionListInner{}
	return &this
}

// GetCONNECTIONNAME returns the CONNECTIONNAME field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONNAME() string {
	if o == nil || IsNil(o.CONNECTIONNAME) {
		var ret string
		return ret
	}
	return *o.CONNECTIONNAME
}

// GetCONNECTIONNAMEOk returns a tuple with the CONNECTIONNAME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONNAMEOk() (*string, bool) {
	if o == nil || IsNil(o.CONNECTIONNAME) {
		return nil, false
	}
	return o.CONNECTIONNAME, true
}

// HasCONNECTIONNAME returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONNAME() bool {
	if o != nil && !IsNil(o.CONNECTIONNAME) {
		return true
	}

	return false
}

// SetCONNECTIONNAME gets a reference to the given string and assigns it to the CONNECTIONNAME field.
func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONNAME(v string) {
	o.CONNECTIONNAME = &v
}

// GetCONNECTIONTYPE returns the CONNECTIONTYPE field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONTYPE() string {
	if o == nil || IsNil(o.CONNECTIONTYPE) {
		var ret string
		return ret
	}
	return *o.CONNECTIONTYPE
}

// GetCONNECTIONTYPEOk returns a tuple with the CONNECTIONTYPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONTYPEOk() (*string, bool) {
	if o == nil || IsNil(o.CONNECTIONTYPE) {
		return nil, false
	}
	return o.CONNECTIONTYPE, true
}

// HasCONNECTIONTYPE returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONTYPE() bool {
	if o != nil && !IsNil(o.CONNECTIONTYPE) {
		return true
	}

	return false
}

// SetCONNECTIONTYPE gets a reference to the given string and assigns it to the CONNECTIONTYPE field.
func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONTYPE(v string) {
	o.CONNECTIONTYPE = &v
}

// GetCONNECTIONDESCRIPTION returns the CONNECTIONDESCRIPTION field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONDESCRIPTION() string {
	if o == nil || IsNil(o.CONNECTIONDESCRIPTION) {
		var ret string
		return ret
	}
	return *o.CONNECTIONDESCRIPTION
}

// GetCONNECTIONDESCRIPTIONOk returns a tuple with the CONNECTIONDESCRIPTION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetCONNECTIONDESCRIPTIONOk() (*string, bool) {
	if o == nil || IsNil(o.CONNECTIONDESCRIPTION) {
		return nil, false
	}
	return o.CONNECTIONDESCRIPTION, true
}

// HasCONNECTIONDESCRIPTION returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasCONNECTIONDESCRIPTION() bool {
	if o != nil && !IsNil(o.CONNECTIONDESCRIPTION) {
		return true
	}

	return false
}

// SetCONNECTIONDESCRIPTION gets a reference to the given string and assigns it to the CONNECTIONDESCRIPTION field.
func (o *GetConnectionsResponseConnectionListInner) SetCONNECTIONDESCRIPTION(v string) {
	o.CONNECTIONDESCRIPTION = &v
}

// GetSTATUS returns the STATUS field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetSTATUS() int32 {
	if o == nil || IsNil(o.STATUS) {
		var ret int32
		return ret
	}
	return *o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetSTATUSOk() (*int32, bool) {
	if o == nil || IsNil(o.STATUS) {
		return nil, false
	}
	return o.STATUS, true
}

// HasSTATUS returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasSTATUS() bool {
	if o != nil && !IsNil(o.STATUS) {
		return true
	}

	return false
}

// SetSTATUS gets a reference to the given int32 and assigns it to the STATUS field.
func (o *GetConnectionsResponseConnectionListInner) SetSTATUS(v int32) {
	o.STATUS = &v
}

// GetCREATEDBY returns the CREATEDBY field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetCREATEDBY() string {
	if o == nil || IsNil(o.CREATEDBY) {
		var ret string
		return ret
	}
	return *o.CREATEDBY
}

// GetCREATEDBYOk returns a tuple with the CREATEDBY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetCREATEDBYOk() (*string, bool) {
	if o == nil || IsNil(o.CREATEDBY) {
		return nil, false
	}
	return o.CREATEDBY, true
}

// HasCREATEDBY returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasCREATEDBY() bool {
	if o != nil && !IsNil(o.CREATEDBY) {
		return true
	}

	return false
}

// SetCREATEDBY gets a reference to the given string and assigns it to the CREATEDBY field.
func (o *GetConnectionsResponseConnectionListInner) SetCREATEDBY(v string) {
	o.CREATEDBY = &v
}

// GetCREATEDON returns the CREATEDON field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetCREATEDON() string {
	if o == nil || IsNil(o.CREATEDON) {
		var ret string
		return ret
	}
	return *o.CREATEDON
}

// GetCREATEDONOk returns a tuple with the CREATEDON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetCREATEDONOk() (*string, bool) {
	if o == nil || IsNil(o.CREATEDON) {
		return nil, false
	}
	return o.CREATEDON, true
}

// HasCREATEDON returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasCREATEDON() bool {
	if o != nil && !IsNil(o.CREATEDON) {
		return true
	}

	return false
}

// SetCREATEDON gets a reference to the given string and assigns it to the CREATEDON field.
func (o *GetConnectionsResponseConnectionListInner) SetCREATEDON(v string) {
	o.CREATEDON = &v
}

// GetUPDATEDBY returns the UPDATEDBY field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDBY() string {
	if o == nil || IsNil(o.UPDATEDBY) {
		var ret string
		return ret
	}
	return *o.UPDATEDBY
}

// GetUPDATEDBYOk returns a tuple with the UPDATEDBY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDBYOk() (*string, bool) {
	if o == nil || IsNil(o.UPDATEDBY) {
		return nil, false
	}
	return o.UPDATEDBY, true
}

// HasUPDATEDBY returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasUPDATEDBY() bool {
	if o != nil && !IsNil(o.UPDATEDBY) {
		return true
	}

	return false
}

// SetUPDATEDBY gets a reference to the given string and assigns it to the UPDATEDBY field.
func (o *GetConnectionsResponseConnectionListInner) SetUPDATEDBY(v string) {
	o.UPDATEDBY = &v
}

// GetUPDATEDON returns the UPDATEDON field value if set, zero value otherwise.
func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDON() string {
	if o == nil || IsNil(o.UPDATEDON) {
		var ret string
		return ret
	}
	return *o.UPDATEDON
}

// GetUPDATEDONOk returns a tuple with the UPDATEDON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConnectionsResponseConnectionListInner) GetUPDATEDONOk() (*string, bool) {
	if o == nil || IsNil(o.UPDATEDON) {
		return nil, false
	}
	return o.UPDATEDON, true
}

// HasUPDATEDON returns a boolean if a field has been set.
func (o *GetConnectionsResponseConnectionListInner) HasUPDATEDON() bool {
	if o != nil && !IsNil(o.UPDATEDON) {
		return true
	}

	return false
}

// SetUPDATEDON gets a reference to the given string and assigns it to the UPDATEDON field.
func (o *GetConnectionsResponseConnectionListInner) SetUPDATEDON(v string) {
	o.UPDATEDON = &v
}

func (o GetConnectionsResponseConnectionListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConnectionsResponseConnectionListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CONNECTIONNAME) {
		toSerialize["CONNECTIONNAME"] = o.CONNECTIONNAME
	}
	if !IsNil(o.CONNECTIONTYPE) {
		toSerialize["CONNECTIONTYPE"] = o.CONNECTIONTYPE
	}
	if !IsNil(o.CONNECTIONDESCRIPTION) {
		toSerialize["CONNECTIONDESCRIPTION"] = o.CONNECTIONDESCRIPTION
	}
	if !IsNil(o.STATUS) {
		toSerialize["STATUS"] = o.STATUS
	}
	if !IsNil(o.CREATEDBY) {
		toSerialize["CREATEDBY"] = o.CREATEDBY
	}
	if !IsNil(o.CREATEDON) {
		toSerialize["CREATEDON"] = o.CREATEDON
	}
	if !IsNil(o.UPDATEDBY) {
		toSerialize["UPDATEDBY"] = o.UPDATEDBY
	}
	if !IsNil(o.UPDATEDON) {
		toSerialize["UPDATEDON"] = o.UPDATEDON
	}
	return toSerialize, nil
}

type NullableGetConnectionsResponseConnectionListInner struct {
	value *GetConnectionsResponseConnectionListInner
	isSet bool
}

func (v NullableGetConnectionsResponseConnectionListInner) Get() *GetConnectionsResponseConnectionListInner {
	return v.value
}

func (v *NullableGetConnectionsResponseConnectionListInner) Set(val *GetConnectionsResponseConnectionListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConnectionsResponseConnectionListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConnectionsResponseConnectionListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConnectionsResponseConnectionListInner(val *GetConnectionsResponseConnectionListInner) *NullableGetConnectionsResponseConnectionListInner {
	return &NullableGetConnectionsResponseConnectionListInner{value: val, isSet: true}
}

func (v NullableGetConnectionsResponseConnectionListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConnectionsResponseConnectionListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
