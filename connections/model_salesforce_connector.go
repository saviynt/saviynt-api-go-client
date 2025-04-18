/*
Connection Management API

Use this API to create a connection in Saviynt Identity Cloud.  The Authorization header must have \"Bearer {token}\".  **Mandatory Parameters:** - **connectionname**: Specify the name to identify the connection. - **connectiontype**: Specify a connection type. For example, if your target application is Active Directory, specify the connection type as \"AD\".  **Optional Parameters:** - **description**: Provide a description for the connection. - **defaultsavroles**: Specify the SAV role(s) required for managing this connection along with its associated security systems, endpoints, accounts, and entitlements. - **emailTemplate**: Specify the email template applicable for notifications. - **sslCertificate**: Specify the SSL certificate(s) to secure the connection between EIC and the target application. - **vaultConfiguration**: Specify the path of the vault to obtain secret data (suffix the connector name to make it unique). - **saveinvault**: Set to true to save the encrypted attribute in the configured vault.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connections

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SalesforceConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesforceConnector{}

// SalesforceConnector struct for SalesforceConnector
type SalesforceConnector struct {
	BaseConnector
	CLIENT_ID                *string `json:"CLIENT_ID,omitempty"`
	CLIENT_SECRET            *string `json:"CLIENT_SECRET,omitempty"`
	REFRESH_TOKEN            *string `json:"REFRESH_TOKEN,omitempty"`
	REDIRECT_URI             *string `json:"REDIRECT_URI,omitempty"`
	INSTANCE_URL             *string `json:"INSTANCE_URL,omitempty"`
	OBJECT_TO_BE_IMPORTED    *string `json:"OBJECT_TO_BE_IMPORTED,omitempty"`
	FEATURE_LICENSE_JSON     *string `json:"FEATURE_LICENSE_JSON,omitempty"`
	CUSTOM_CREATEACCOUNT_URL *string `json:"CUSTOM_CREATEACCOUNT_URL,omitempty"`
	CREATEACCOUNTJSON        *string `json:"CREATEACCOUNTJSON,omitempty"`
	ACCOUNT_FILTER_QUERY     *string `json:"ACCOUNT_FILTER_QUERY,omitempty"`
	ACCOUNT_FIELD_QUERY      *string `json:"ACCOUNT_FIELD_QUERY,omitempty"`
	FIELD_MAPPING_JSON       *string `json:"FIELD_MAPPING_JSON,omitempty"`
	MODIFYACCOUNTJSON        *string `json:"MODIFYACCOUNTJSON,omitempty"`
	// The attributes of statusAndThresholdConfig json
	STATUS_THRESHOLD_CONFIG *string `json:"STATUS_THRESHOLD_CONFIG,omitempty"`
	CUSTOMCONFIGJSON        *string `json:"CUSTOMCONFIGJSON,omitempty"`
	PAM_CONFIG              *string `json:"PAM_CONFIG,omitempty"`
}

type _SalesforceConnector SalesforceConnector

// NewSalesforceConnector instantiates a new SalesforceConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesforceConnector(connectionName string, connectiontype string) *SalesforceConnector {
	this := SalesforceConnector{}
	this.ConnectionName = connectionName
	this.Connectiontype = connectiontype
	return &this
}

// NewSalesforceConnectorWithDefaults instantiates a new SalesforceConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesforceConnectorWithDefaults() *SalesforceConnector {
	this := SalesforceConnector{}
	return &this
}

// GetCLIENT_ID returns the CLIENT_ID field value if set, zero value otherwise.
func (o *SalesforceConnector) GetCLIENT_ID() string {
	if o == nil || IsNil(o.CLIENT_ID) {
		var ret string
		return ret
	}
	return *o.CLIENT_ID
}

// GetCLIENT_IDOk returns a tuple with the CLIENT_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetCLIENT_IDOk() (*string, bool) {
	if o == nil || IsNil(o.CLIENT_ID) {
		return nil, false
	}
	return o.CLIENT_ID, true
}

// HasCLIENT_ID returns a boolean if a field has been set.
func (o *SalesforceConnector) HasCLIENT_ID() bool {
	if o != nil && !IsNil(o.CLIENT_ID) {
		return true
	}

	return false
}

// SetCLIENT_ID gets a reference to the given string and assigns it to the CLIENT_ID field.
func (o *SalesforceConnector) SetCLIENT_ID(v string) {
	o.CLIENT_ID = &v
}

// GetCLIENT_SECRET returns the CLIENT_SECRET field value if set, zero value otherwise.
func (o *SalesforceConnector) GetCLIENT_SECRET() string {
	if o == nil || IsNil(o.CLIENT_SECRET) {
		var ret string
		return ret
	}
	return *o.CLIENT_SECRET
}

// GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetCLIENT_SECRETOk() (*string, bool) {
	if o == nil || IsNil(o.CLIENT_SECRET) {
		return nil, false
	}
	return o.CLIENT_SECRET, true
}

// HasCLIENT_SECRET returns a boolean if a field has been set.
func (o *SalesforceConnector) HasCLIENT_SECRET() bool {
	if o != nil && !IsNil(o.CLIENT_SECRET) {
		return true
	}

	return false
}

// SetCLIENT_SECRET gets a reference to the given string and assigns it to the CLIENT_SECRET field.
func (o *SalesforceConnector) SetCLIENT_SECRET(v string) {
	o.CLIENT_SECRET = &v
}

// GetREFRESH_TOKEN returns the REFRESH_TOKEN field value if set, zero value otherwise.
func (o *SalesforceConnector) GetREFRESH_TOKEN() string {
	if o == nil || IsNil(o.REFRESH_TOKEN) {
		var ret string
		return ret
	}
	return *o.REFRESH_TOKEN
}

// GetREFRESH_TOKENOk returns a tuple with the REFRESH_TOKEN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetREFRESH_TOKENOk() (*string, bool) {
	if o == nil || IsNil(o.REFRESH_TOKEN) {
		return nil, false
	}
	return o.REFRESH_TOKEN, true
}

// HasREFRESH_TOKEN returns a boolean if a field has been set.
func (o *SalesforceConnector) HasREFRESH_TOKEN() bool {
	if o != nil && !IsNil(o.REFRESH_TOKEN) {
		return true
	}

	return false
}

// SetREFRESH_TOKEN gets a reference to the given string and assigns it to the REFRESH_TOKEN field.
func (o *SalesforceConnector) SetREFRESH_TOKEN(v string) {
	o.REFRESH_TOKEN = &v
}

// GetREDIRECT_URI returns the REDIRECT_URI field value if set, zero value otherwise.
func (o *SalesforceConnector) GetREDIRECT_URI() string {
	if o == nil || IsNil(o.REDIRECT_URI) {
		var ret string
		return ret
	}
	return *o.REDIRECT_URI
}

// GetREDIRECT_URIOk returns a tuple with the REDIRECT_URI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetREDIRECT_URIOk() (*string, bool) {
	if o == nil || IsNil(o.REDIRECT_URI) {
		return nil, false
	}
	return o.REDIRECT_URI, true
}

// HasREDIRECT_URI returns a boolean if a field has been set.
func (o *SalesforceConnector) HasREDIRECT_URI() bool {
	if o != nil && !IsNil(o.REDIRECT_URI) {
		return true
	}

	return false
}

// SetREDIRECT_URI gets a reference to the given string and assigns it to the REDIRECT_URI field.
func (o *SalesforceConnector) SetREDIRECT_URI(v string) {
	o.REDIRECT_URI = &v
}

// GetINSTANCE_URL returns the INSTANCE_URL field value if set, zero value otherwise.
func (o *SalesforceConnector) GetINSTANCE_URL() string {
	if o == nil || IsNil(o.INSTANCE_URL) {
		var ret string
		return ret
	}
	return *o.INSTANCE_URL
}

// GetINSTANCE_URLOk returns a tuple with the INSTANCE_URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetINSTANCE_URLOk() (*string, bool) {
	if o == nil || IsNil(o.INSTANCE_URL) {
		return nil, false
	}
	return o.INSTANCE_URL, true
}

// HasINSTANCE_URL returns a boolean if a field has been set.
func (o *SalesforceConnector) HasINSTANCE_URL() bool {
	if o != nil && !IsNil(o.INSTANCE_URL) {
		return true
	}

	return false
}

// SetINSTANCE_URL gets a reference to the given string and assigns it to the INSTANCE_URL field.
func (o *SalesforceConnector) SetINSTANCE_URL(v string) {
	o.INSTANCE_URL = &v
}

// GetOBJECT_TO_BE_IMPORTED returns the OBJECT_TO_BE_IMPORTED field value if set, zero value otherwise.
func (o *SalesforceConnector) GetOBJECT_TO_BE_IMPORTED() string {
	if o == nil || IsNil(o.OBJECT_TO_BE_IMPORTED) {
		var ret string
		return ret
	}
	return *o.OBJECT_TO_BE_IMPORTED
}

// GetOBJECT_TO_BE_IMPORTEDOk returns a tuple with the OBJECT_TO_BE_IMPORTED field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetOBJECT_TO_BE_IMPORTEDOk() (*string, bool) {
	if o == nil || IsNil(o.OBJECT_TO_BE_IMPORTED) {
		return nil, false
	}
	return o.OBJECT_TO_BE_IMPORTED, true
}

// HasOBJECT_TO_BE_IMPORTED returns a boolean if a field has been set.
func (o *SalesforceConnector) HasOBJECT_TO_BE_IMPORTED() bool {
	if o != nil && !IsNil(o.OBJECT_TO_BE_IMPORTED) {
		return true
	}

	return false
}

// SetOBJECT_TO_BE_IMPORTED gets a reference to the given string and assigns it to the OBJECT_TO_BE_IMPORTED field.
func (o *SalesforceConnector) SetOBJECT_TO_BE_IMPORTED(v string) {
	o.OBJECT_TO_BE_IMPORTED = &v
}

// GetFEATURE_LICENSE_JSON returns the FEATURE_LICENSE_JSON field value if set, zero value otherwise.
func (o *SalesforceConnector) GetFEATURE_LICENSE_JSON() string {
	if o == nil || IsNil(o.FEATURE_LICENSE_JSON) {
		var ret string
		return ret
	}
	return *o.FEATURE_LICENSE_JSON
}

// GetFEATURE_LICENSE_JSONOk returns a tuple with the FEATURE_LICENSE_JSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetFEATURE_LICENSE_JSONOk() (*string, bool) {
	if o == nil || IsNil(o.FEATURE_LICENSE_JSON) {
		return nil, false
	}
	return o.FEATURE_LICENSE_JSON, true
}

// HasFEATURE_LICENSE_JSON returns a boolean if a field has been set.
func (o *SalesforceConnector) HasFEATURE_LICENSE_JSON() bool {
	if o != nil && !IsNil(o.FEATURE_LICENSE_JSON) {
		return true
	}

	return false
}

// SetFEATURE_LICENSE_JSON gets a reference to the given string and assigns it to the FEATURE_LICENSE_JSON field.
func (o *SalesforceConnector) SetFEATURE_LICENSE_JSON(v string) {
	o.FEATURE_LICENSE_JSON = &v
}

// GetCUSTOM_CREATEACCOUNT_URL returns the CUSTOM_CREATEACCOUNT_URL field value if set, zero value otherwise.
func (o *SalesforceConnector) GetCUSTOM_CREATEACCOUNT_URL() string {
	if o == nil || IsNil(o.CUSTOM_CREATEACCOUNT_URL) {
		var ret string
		return ret
	}
	return *o.CUSTOM_CREATEACCOUNT_URL
}

// GetCUSTOM_CREATEACCOUNT_URLOk returns a tuple with the CUSTOM_CREATEACCOUNT_URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetCUSTOM_CREATEACCOUNT_URLOk() (*string, bool) {
	if o == nil || IsNil(o.CUSTOM_CREATEACCOUNT_URL) {
		return nil, false
	}
	return o.CUSTOM_CREATEACCOUNT_URL, true
}

// HasCUSTOM_CREATEACCOUNT_URL returns a boolean if a field has been set.
func (o *SalesforceConnector) HasCUSTOM_CREATEACCOUNT_URL() bool {
	if o != nil && !IsNil(o.CUSTOM_CREATEACCOUNT_URL) {
		return true
	}

	return false
}

// SetCUSTOM_CREATEACCOUNT_URL gets a reference to the given string and assigns it to the CUSTOM_CREATEACCOUNT_URL field.
func (o *SalesforceConnector) SetCUSTOM_CREATEACCOUNT_URL(v string) {
	o.CUSTOM_CREATEACCOUNT_URL = &v
}

// GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field value if set, zero value otherwise.
func (o *SalesforceConnector) GetCREATEACCOUNTJSON() string {
	if o == nil || IsNil(o.CREATEACCOUNTJSON) {
		var ret string
		return ret
	}
	return *o.CREATEACCOUNTJSON
}

// GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetCREATEACCOUNTJSONOk() (*string, bool) {
	if o == nil || IsNil(o.CREATEACCOUNTJSON) {
		return nil, false
	}
	return o.CREATEACCOUNTJSON, true
}

// HasCREATEACCOUNTJSON returns a boolean if a field has been set.
func (o *SalesforceConnector) HasCREATEACCOUNTJSON() bool {
	if o != nil && !IsNil(o.CREATEACCOUNTJSON) {
		return true
	}

	return false
}

// SetCREATEACCOUNTJSON gets a reference to the given string and assigns it to the CREATEACCOUNTJSON field.
func (o *SalesforceConnector) SetCREATEACCOUNTJSON(v string) {
	o.CREATEACCOUNTJSON = &v
}

// GetACCOUNT_FILTER_QUERY returns the ACCOUNT_FILTER_QUERY field value if set, zero value otherwise.
func (o *SalesforceConnector) GetACCOUNT_FILTER_QUERY() string {
	if o == nil || IsNil(o.ACCOUNT_FILTER_QUERY) {
		var ret string
		return ret
	}
	return *o.ACCOUNT_FILTER_QUERY
}

// GetACCOUNT_FILTER_QUERYOk returns a tuple with the ACCOUNT_FILTER_QUERY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetACCOUNT_FILTER_QUERYOk() (*string, bool) {
	if o == nil || IsNil(o.ACCOUNT_FILTER_QUERY) {
		return nil, false
	}
	return o.ACCOUNT_FILTER_QUERY, true
}

// HasACCOUNT_FILTER_QUERY returns a boolean if a field has been set.
func (o *SalesforceConnector) HasACCOUNT_FILTER_QUERY() bool {
	if o != nil && !IsNil(o.ACCOUNT_FILTER_QUERY) {
		return true
	}

	return false
}

// SetACCOUNT_FILTER_QUERY gets a reference to the given string and assigns it to the ACCOUNT_FILTER_QUERY field.
func (o *SalesforceConnector) SetACCOUNT_FILTER_QUERY(v string) {
	o.ACCOUNT_FILTER_QUERY = &v
}

// GetACCOUNT_FIELD_QUERY returns the ACCOUNT_FIELD_QUERY field value if set, zero value otherwise.
func (o *SalesforceConnector) GetACCOUNT_FIELD_QUERY() string {
	if o == nil || IsNil(o.ACCOUNT_FIELD_QUERY) {
		var ret string
		return ret
	}
	return *o.ACCOUNT_FIELD_QUERY
}

// GetACCOUNT_FIELD_QUERYOk returns a tuple with the ACCOUNT_FIELD_QUERY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetACCOUNT_FIELD_QUERYOk() (*string, bool) {
	if o == nil || IsNil(o.ACCOUNT_FIELD_QUERY) {
		return nil, false
	}
	return o.ACCOUNT_FIELD_QUERY, true
}

// HasACCOUNT_FIELD_QUERY returns a boolean if a field has been set.
func (o *SalesforceConnector) HasACCOUNT_FIELD_QUERY() bool {
	if o != nil && !IsNil(o.ACCOUNT_FIELD_QUERY) {
		return true
	}

	return false
}

// SetACCOUNT_FIELD_QUERY gets a reference to the given string and assigns it to the ACCOUNT_FIELD_QUERY field.
func (o *SalesforceConnector) SetACCOUNT_FIELD_QUERY(v string) {
	o.ACCOUNT_FIELD_QUERY = &v
}

// GetFIELD_MAPPING_JSON returns the FIELD_MAPPING_JSON field value if set, zero value otherwise.
func (o *SalesforceConnector) GetFIELD_MAPPING_JSON() string {
	if o == nil || IsNil(o.FIELD_MAPPING_JSON) {
		var ret string
		return ret
	}
	return *o.FIELD_MAPPING_JSON
}

// GetFIELD_MAPPING_JSONOk returns a tuple with the FIELD_MAPPING_JSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetFIELD_MAPPING_JSONOk() (*string, bool) {
	if o == nil || IsNil(o.FIELD_MAPPING_JSON) {
		return nil, false
	}
	return o.FIELD_MAPPING_JSON, true
}

// HasFIELD_MAPPING_JSON returns a boolean if a field has been set.
func (o *SalesforceConnector) HasFIELD_MAPPING_JSON() bool {
	if o != nil && !IsNil(o.FIELD_MAPPING_JSON) {
		return true
	}

	return false
}

// SetFIELD_MAPPING_JSON gets a reference to the given string and assigns it to the FIELD_MAPPING_JSON field.
func (o *SalesforceConnector) SetFIELD_MAPPING_JSON(v string) {
	o.FIELD_MAPPING_JSON = &v
}

// GetMODIFYACCOUNTJSON returns the MODIFYACCOUNTJSON field value if set, zero value otherwise.
func (o *SalesforceConnector) GetMODIFYACCOUNTJSON() string {
	if o == nil || IsNil(o.MODIFYACCOUNTJSON) {
		var ret string
		return ret
	}
	return *o.MODIFYACCOUNTJSON
}

// GetMODIFYACCOUNTJSONOk returns a tuple with the MODIFYACCOUNTJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetMODIFYACCOUNTJSONOk() (*string, bool) {
	if o == nil || IsNil(o.MODIFYACCOUNTJSON) {
		return nil, false
	}
	return o.MODIFYACCOUNTJSON, true
}

// HasMODIFYACCOUNTJSON returns a boolean if a field has been set.
func (o *SalesforceConnector) HasMODIFYACCOUNTJSON() bool {
	if o != nil && !IsNil(o.MODIFYACCOUNTJSON) {
		return true
	}

	return false
}

// SetMODIFYACCOUNTJSON gets a reference to the given string and assigns it to the MODIFYACCOUNTJSON field.
func (o *SalesforceConnector) SetMODIFYACCOUNTJSON(v string) {
	o.MODIFYACCOUNTJSON = &v
}

// GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field value if set, zero value otherwise.
func (o *SalesforceConnector) GetSTATUS_THRESHOLD_CONFIG() string {
	if o == nil || IsNil(o.STATUS_THRESHOLD_CONFIG) {
		var ret string
		return ret
	}
	return *o.STATUS_THRESHOLD_CONFIG
}

// GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool) {
	if o == nil || IsNil(o.STATUS_THRESHOLD_CONFIG) {
		return nil, false
	}
	return o.STATUS_THRESHOLD_CONFIG, true
}

// HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.
func (o *SalesforceConnector) HasSTATUS_THRESHOLD_CONFIG() bool {
	if o != nil && !IsNil(o.STATUS_THRESHOLD_CONFIG) {
		return true
	}

	return false
}

// SetSTATUS_THRESHOLD_CONFIG gets a reference to the given string and assigns it to the STATUS_THRESHOLD_CONFIG field.
func (o *SalesforceConnector) SetSTATUS_THRESHOLD_CONFIG(v string) {
	o.STATUS_THRESHOLD_CONFIG = &v
}

// GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field value if set, zero value otherwise.
func (o *SalesforceConnector) GetCUSTOMCONFIGJSON() string {
	if o == nil || IsNil(o.CUSTOMCONFIGJSON) {
		var ret string
		return ret
	}
	return *o.CUSTOMCONFIGJSON
}

// GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetCUSTOMCONFIGJSONOk() (*string, bool) {
	if o == nil || IsNil(o.CUSTOMCONFIGJSON) {
		return nil, false
	}
	return o.CUSTOMCONFIGJSON, true
}

// HasCUSTOMCONFIGJSON returns a boolean if a field has been set.
func (o *SalesforceConnector) HasCUSTOMCONFIGJSON() bool {
	if o != nil && !IsNil(o.CUSTOMCONFIGJSON) {
		return true
	}

	return false
}

// SetCUSTOMCONFIGJSON gets a reference to the given string and assigns it to the CUSTOMCONFIGJSON field.
func (o *SalesforceConnector) SetCUSTOMCONFIGJSON(v string) {
	o.CUSTOMCONFIGJSON = &v
}

// GetPAM_CONFIG returns the PAM_CONFIG field value if set, zero value otherwise.
func (o *SalesforceConnector) GetPAM_CONFIG() string {
	if o == nil || IsNil(o.PAM_CONFIG) {
		var ret string
		return ret
	}
	return *o.PAM_CONFIG
}

// GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceConnector) GetPAM_CONFIGOk() (*string, bool) {
	if o == nil || IsNil(o.PAM_CONFIG) {
		return nil, false
	}
	return o.PAM_CONFIG, true
}

// HasPAM_CONFIG returns a boolean if a field has been set.
func (o *SalesforceConnector) HasPAM_CONFIG() bool {
	if o != nil && !IsNil(o.PAM_CONFIG) {
		return true
	}

	return false
}

// SetPAM_CONFIG gets a reference to the given string and assigns it to the PAM_CONFIG field.
func (o *SalesforceConnector) SetPAM_CONFIG(v string) {
	o.PAM_CONFIG = &v
}

func (o SalesforceConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesforceConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseConnector, errBaseConnector := json.Marshal(o.BaseConnector)
	if errBaseConnector != nil {
		return map[string]interface{}{}, errBaseConnector
	}
	errBaseConnector = json.Unmarshal([]byte(serializedBaseConnector), &toSerialize)
	if errBaseConnector != nil {
		return map[string]interface{}{}, errBaseConnector
	}
	if !IsNil(o.CLIENT_ID) {
		toSerialize["CLIENT_ID"] = o.CLIENT_ID
	}
	if !IsNil(o.CLIENT_SECRET) {
		toSerialize["CLIENT_SECRET"] = o.CLIENT_SECRET
	}
	if !IsNil(o.REFRESH_TOKEN) {
		toSerialize["REFRESH_TOKEN"] = o.REFRESH_TOKEN
	}
	if !IsNil(o.REDIRECT_URI) {
		toSerialize["REDIRECT_URI"] = o.REDIRECT_URI
	}
	if !IsNil(o.INSTANCE_URL) {
		toSerialize["INSTANCE_URL"] = o.INSTANCE_URL
	}
	if !IsNil(o.OBJECT_TO_BE_IMPORTED) {
		toSerialize["OBJECT_TO_BE_IMPORTED"] = o.OBJECT_TO_BE_IMPORTED
	}
	if !IsNil(o.FEATURE_LICENSE_JSON) {
		toSerialize["FEATURE_LICENSE_JSON"] = o.FEATURE_LICENSE_JSON
	}
	if !IsNil(o.CUSTOM_CREATEACCOUNT_URL) {
		toSerialize["CUSTOM_CREATEACCOUNT_URL"] = o.CUSTOM_CREATEACCOUNT_URL
	}
	if !IsNil(o.CREATEACCOUNTJSON) {
		toSerialize["CREATEACCOUNTJSON"] = o.CREATEACCOUNTJSON
	}
	if !IsNil(o.ACCOUNT_FILTER_QUERY) {
		toSerialize["ACCOUNT_FILTER_QUERY"] = o.ACCOUNT_FILTER_QUERY
	}
	if !IsNil(o.ACCOUNT_FIELD_QUERY) {
		toSerialize["ACCOUNT_FIELD_QUERY"] = o.ACCOUNT_FIELD_QUERY
	}
	if !IsNil(o.FIELD_MAPPING_JSON) {
		toSerialize["FIELD_MAPPING_JSON"] = o.FIELD_MAPPING_JSON
	}
	if !IsNil(o.MODIFYACCOUNTJSON) {
		toSerialize["MODIFYACCOUNTJSON"] = o.MODIFYACCOUNTJSON
	}
	if !IsNil(o.STATUS_THRESHOLD_CONFIG) {
		toSerialize["STATUS_THRESHOLD_CONFIG"] = o.STATUS_THRESHOLD_CONFIG
	}
	if !IsNil(o.CUSTOMCONFIGJSON) {
		toSerialize["CUSTOMCONFIGJSON"] = o.CUSTOMCONFIGJSON
	}
	if !IsNil(o.PAM_CONFIG) {
		toSerialize["PAM_CONFIG"] = o.PAM_CONFIG
	}
	return toSerialize, nil
}

func (o *SalesforceConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionName",
		"connectiontype",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSalesforceConnector := _SalesforceConnector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSalesforceConnector)

	if err != nil {
		return err
	}

	*o = SalesforceConnector(varSalesforceConnector)

	return err
}

type NullableSalesforceConnector struct {
	value *SalesforceConnector
	isSet bool
}

func (v NullableSalesforceConnector) Get() *SalesforceConnector {
	return v.value
}

func (v *NullableSalesforceConnector) Set(val *SalesforceConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesforceConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesforceConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesforceConnector(val *SalesforceConnector) *NullableSalesforceConnector {
	return &NullableSalesforceConnector{value: val, isSet: true}
}

func (v NullableSalesforceConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesforceConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
