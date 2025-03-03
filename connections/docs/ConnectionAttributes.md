# ConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCOUNT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ACCOUNT_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**ACCOUNT_ATTRIBUTE_LABEL** | Pointer to **string** |  | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** |  | [optional] 
**ADVANCE_FILTER_JSON** | Pointer to **string** |  | [optional] 
**ADVSEARCH** | Pointer to **string** |  | [optional] 
**APPLICATION_NAME** | Pointer to **string** |  | [optional] 
**BASE** | Pointer to **string** |  | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**CONNECTOR_FILE_PATH** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**CREATEORGJSON** | Pointer to **string** |  | [optional] 
**CUSTOM_ACCOUNT_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**CUSTOMER_ID** | Pointer to **string** |  | [optional] 
**DC_LOCATOR** | Pointer to **string** |  | [optional] 
**DEFAULT_USER_ROLE** | Pointer to **string** |  | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** |  | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**ENFORCE_TREE_DELETION** | Pointer to **string** |  | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**FILEFOLDER_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**FILTER** | Pointer to **string** |  | [optional] 
**GROUP_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**IMPORT_USERS** | Pointer to **string** |  | [optional] 
**IMPORTJSON** | Pointer to **string** |  | [optional] 
**INCREMENTAL_CONFIG** | Pointer to **string** |  | [optional] 
**LAST_IMPORT_TIME** | Pointer to **string** |  | [optional] 
**LDAP_OR_AD** | Pointer to **string** |  | [optional] 
**MAX_CHANGENUMBER** | Pointer to **string** |  | [optional] 
**MKTPLACEAPP_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**OBJECTFILTER** | Pointer to **string** |  | [optional] 
**ORG_BASE** | Pointer to **string** |  | [optional] 
**ORGANIZATION_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ORGIMPORTJSON** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** |  | [optional] 
**QUARANTINE_FOLDER_ID** | Pointer to **string** |  | [optional] 
**READ_OPERATIONAL_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**REMOVEACCOUNTACTION** | Pointer to **string** |  | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** |  | [optional] 
**REUSEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**REUSEINACTIVEACCOUNT** | Pointer to **string** |  | [optional] 
**SCOPES** | Pointer to **string** |  | [optional] 
**SEARCHFILTER** | Pointer to **string** |  | [optional] 
**SERVICE_ACCOUNT_ID** | Pointer to **string** |  | [optional] 
**SERVICE_ACCOUNT_KEY_JSON** | Pointer to **string** |  | [optional] 
**SERVICE_ACCOUNT_USER** | Pointer to **string** |  | [optional] 
**SETDEFAULTPAGESIZE** | Pointer to **string** |  | [optional] 
**SETRANDOMPASSWORD** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**STATUSKEYJSON** | Pointer to **string** |  | [optional] 
**SUPPORTEMPTYSTRING** | Pointer to **string** |  | [optional] 
**SUSPEND_USER_ORGANIZATIONAL_UNIT** | Pointer to **string** |  | [optional] 
**UNLOCKACCOUNTJSON** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**UPDATEORGJSON** | Pointer to **string** |  | [optional] 
**UPDATEUSERJSON** | Pointer to **string** |  | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**USER_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**CreateUpdateMappings** | Pointer to **string** |  | [optional] 
**GroupImportMapping** | Pointer to **string** |  | [optional] 
**GroupSearchBaseDN** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnectionAttributes

`func NewConnectionAttributes() *ConnectionAttributes`

NewConnectionAttributes instantiates a new ConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionAttributesWithDefaults

`func NewConnectionAttributesWithDefaults() *ConnectionAttributes`

NewConnectionAttributesWithDefaults instantiates a new ConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCOUNT_ATTRIBUTE

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *ConnectionAttributes) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *ConnectionAttributes) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTES() string`

GetACCOUNT_ATTRIBUTES returns the ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetACCOUNT_ATTRIBUTESOk returns a tuple with the ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) SetACCOUNT_ATTRIBUTES(v string)`

SetACCOUNT_ATTRIBUTES sets ACCOUNT_ATTRIBUTES field to given value.

### HasACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) HasACCOUNT_ATTRIBUTES() bool`

HasACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE_LABEL

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTE_LABEL() string`

GetACCOUNT_ATTRIBUTE_LABEL returns the ACCOUNT_ATTRIBUTE_LABEL field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTE_LABELOk

`func (o *ConnectionAttributes) GetACCOUNT_ATTRIBUTE_LABELOk() (*string, bool)`

GetACCOUNT_ATTRIBUTE_LABELOk returns a tuple with the ACCOUNT_ATTRIBUTE_LABEL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE_LABEL

`func (o *ConnectionAttributes) SetACCOUNT_ATTRIBUTE_LABEL(v string)`

SetACCOUNT_ATTRIBUTE_LABEL sets ACCOUNT_ATTRIBUTE_LABEL field to given value.

### HasACCOUNT_ATTRIBUTE_LABEL

`func (o *ConnectionAttributes) HasACCOUNT_ATTRIBUTE_LABEL() bool`

HasACCOUNT_ATTRIBUTE_LABEL returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *ConnectionAttributes) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *ConnectionAttributes) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *ConnectionAttributes) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *ConnectionAttributes) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetADVANCE_FILTER_JSON

`func (o *ConnectionAttributes) GetADVANCE_FILTER_JSON() string`

GetADVANCE_FILTER_JSON returns the ADVANCE_FILTER_JSON field if non-nil, zero value otherwise.

### GetADVANCE_FILTER_JSONOk

`func (o *ConnectionAttributes) GetADVANCE_FILTER_JSONOk() (*string, bool)`

GetADVANCE_FILTER_JSONOk returns a tuple with the ADVANCE_FILTER_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVANCE_FILTER_JSON

`func (o *ConnectionAttributes) SetADVANCE_FILTER_JSON(v string)`

SetADVANCE_FILTER_JSON sets ADVANCE_FILTER_JSON field to given value.

### HasADVANCE_FILTER_JSON

`func (o *ConnectionAttributes) HasADVANCE_FILTER_JSON() bool`

HasADVANCE_FILTER_JSON returns a boolean if a field has been set.

### GetADVSEARCH

`func (o *ConnectionAttributes) GetADVSEARCH() string`

GetADVSEARCH returns the ADVSEARCH field if non-nil, zero value otherwise.

### GetADVSEARCHOk

`func (o *ConnectionAttributes) GetADVSEARCHOk() (*string, bool)`

GetADVSEARCHOk returns a tuple with the ADVSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVSEARCH

`func (o *ConnectionAttributes) SetADVSEARCH(v string)`

SetADVSEARCH sets ADVSEARCH field to given value.

### HasADVSEARCH

`func (o *ConnectionAttributes) HasADVSEARCH() bool`

HasADVSEARCH returns a boolean if a field has been set.

### GetAPPLICATION_NAME

`func (o *ConnectionAttributes) GetAPPLICATION_NAME() string`

GetAPPLICATION_NAME returns the APPLICATION_NAME field if non-nil, zero value otherwise.

### GetAPPLICATION_NAMEOk

`func (o *ConnectionAttributes) GetAPPLICATION_NAMEOk() (*string, bool)`

GetAPPLICATION_NAMEOk returns a tuple with the APPLICATION_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPLICATION_NAME

`func (o *ConnectionAttributes) SetAPPLICATION_NAME(v string)`

SetAPPLICATION_NAME sets APPLICATION_NAME field to given value.

### HasAPPLICATION_NAME

`func (o *ConnectionAttributes) HasAPPLICATION_NAME() bool`

HasAPPLICATION_NAME returns a boolean if a field has been set.

### GetBASE

`func (o *ConnectionAttributes) GetBASE() string`

GetBASE returns the BASE field if non-nil, zero value otherwise.

### GetBASEOk

`func (o *ConnectionAttributes) GetBASEOk() (*string, bool)`

GetBASEOk returns a tuple with the BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE

`func (o *ConnectionAttributes) SetBASE(v string)`

SetBASE sets BASE field to given value.

### HasBASE

`func (o *ConnectionAttributes) HasBASE() bool`

HasBASE returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *ConnectionAttributes) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *ConnectionAttributes) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *ConnectionAttributes) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *ConnectionAttributes) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetConfigJSON

`func (o *ConnectionAttributes) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *ConnectionAttributes) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *ConnectionAttributes) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *ConnectionAttributes) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetCONNECTOR_FILE_PATH

`func (o *ConnectionAttributes) GetCONNECTOR_FILE_PATH() string`

GetCONNECTOR_FILE_PATH returns the CONNECTOR_FILE_PATH field if non-nil, zero value otherwise.

### GetCONNECTOR_FILE_PATHOk

`func (o *ConnectionAttributes) GetCONNECTOR_FILE_PATHOk() (*string, bool)`

GetCONNECTOR_FILE_PATHOk returns a tuple with the CONNECTOR_FILE_PATH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTOR_FILE_PATH

`func (o *ConnectionAttributes) SetCONNECTOR_FILE_PATH(v string)`

SetCONNECTOR_FILE_PATH sets CONNECTOR_FILE_PATH field to given value.

### HasCONNECTOR_FILE_PATH

`func (o *ConnectionAttributes) HasCONNECTOR_FILE_PATH() bool`

HasCONNECTOR_FILE_PATH returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *ConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *ConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *ConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *ConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetCREATEORGJSON

`func (o *ConnectionAttributes) GetCREATEORGJSON() string`

GetCREATEORGJSON returns the CREATEORGJSON field if non-nil, zero value otherwise.

### GetCREATEORGJSONOk

`func (o *ConnectionAttributes) GetCREATEORGJSONOk() (*string, bool)`

GetCREATEORGJSONOk returns a tuple with the CREATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEORGJSON

`func (o *ConnectionAttributes) SetCREATEORGJSON(v string)`

SetCREATEORGJSON sets CREATEORGJSON field to given value.

### HasCREATEORGJSON

`func (o *ConnectionAttributes) HasCREATEORGJSON() bool`

HasCREATEORGJSON returns a boolean if a field has been set.

### GetCUSTOM_ACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) GetCUSTOM_ACCOUNT_ATTRIBUTES() string`

GetCUSTOM_ACCOUNT_ATTRIBUTES returns the CUSTOM_ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetCUSTOM_ACCOUNT_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetCUSTOM_ACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetCUSTOM_ACCOUNT_ATTRIBUTESOk returns a tuple with the CUSTOM_ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_ACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) SetCUSTOM_ACCOUNT_ATTRIBUTES(v string)`

SetCUSTOM_ACCOUNT_ATTRIBUTES sets CUSTOM_ACCOUNT_ATTRIBUTES field to given value.

### HasCUSTOM_ACCOUNT_ATTRIBUTES

`func (o *ConnectionAttributes) HasCUSTOM_ACCOUNT_ATTRIBUTES() bool`

HasCUSTOM_ACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetCUSTOMER_ID

`func (o *ConnectionAttributes) GetCUSTOMER_ID() string`

GetCUSTOMER_ID returns the CUSTOMER_ID field if non-nil, zero value otherwise.

### GetCUSTOMER_IDOk

`func (o *ConnectionAttributes) GetCUSTOMER_IDOk() (*string, bool)`

GetCUSTOMER_IDOk returns a tuple with the CUSTOMER_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMER_ID

`func (o *ConnectionAttributes) SetCUSTOMER_ID(v string)`

SetCUSTOMER_ID sets CUSTOMER_ID field to given value.

### HasCUSTOMER_ID

`func (o *ConnectionAttributes) HasCUSTOMER_ID() bool`

HasCUSTOMER_ID returns a boolean if a field has been set.

### GetDC_LOCATOR

`func (o *ConnectionAttributes) GetDC_LOCATOR() string`

GetDC_LOCATOR returns the DC_LOCATOR field if non-nil, zero value otherwise.

### GetDC_LOCATOROk

`func (o *ConnectionAttributes) GetDC_LOCATOROk() (*string, bool)`

GetDC_LOCATOROk returns a tuple with the DC_LOCATOR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDC_LOCATOR

`func (o *ConnectionAttributes) SetDC_LOCATOR(v string)`

SetDC_LOCATOR sets DC_LOCATOR field to given value.

### HasDC_LOCATOR

`func (o *ConnectionAttributes) HasDC_LOCATOR() bool`

HasDC_LOCATOR returns a boolean if a field has been set.

### GetDEFAULT_USER_ROLE

`func (o *ConnectionAttributes) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *ConnectionAttributes) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *ConnectionAttributes) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *ConnectionAttributes) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *ConnectionAttributes) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *ConnectionAttributes) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *ConnectionAttributes) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *ConnectionAttributes) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *ConnectionAttributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *ConnectionAttributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *ConnectionAttributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *ConnectionAttributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEGROUPMANAGEMENT

`func (o *ConnectionAttributes) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *ConnectionAttributes) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *ConnectionAttributes) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *ConnectionAttributes) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *ConnectionAttributes) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *ConnectionAttributes) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *ConnectionAttributes) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *ConnectionAttributes) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetENFORCE_TREE_DELETION

`func (o *ConnectionAttributes) GetENFORCE_TREE_DELETION() string`

GetENFORCE_TREE_DELETION returns the ENFORCE_TREE_DELETION field if non-nil, zero value otherwise.

### GetENFORCE_TREE_DELETIONOk

`func (o *ConnectionAttributes) GetENFORCE_TREE_DELETIONOk() (*string, bool)`

GetENFORCE_TREE_DELETIONOk returns a tuple with the ENFORCE_TREE_DELETION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCE_TREE_DELETION

`func (o *ConnectionAttributes) SetENFORCE_TREE_DELETION(v string)`

SetENFORCE_TREE_DELETION sets ENFORCE_TREE_DELETION field to given value.

### HasENFORCE_TREE_DELETION

`func (o *ConnectionAttributes) HasENFORCE_TREE_DELETION() bool`

HasENFORCE_TREE_DELETION returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *ConnectionAttributes) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *ConnectionAttributes) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *ConnectionAttributes) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *ConnectionAttributes) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetFILEFOLDER_ATTRIBUTES

`func (o *ConnectionAttributes) GetFILEFOLDER_ATTRIBUTES() string`

GetFILEFOLDER_ATTRIBUTES returns the FILEFOLDER_ATTRIBUTES field if non-nil, zero value otherwise.

### GetFILEFOLDER_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetFILEFOLDER_ATTRIBUTESOk() (*string, bool)`

GetFILEFOLDER_ATTRIBUTESOk returns a tuple with the FILEFOLDER_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILEFOLDER_ATTRIBUTES

`func (o *ConnectionAttributes) SetFILEFOLDER_ATTRIBUTES(v string)`

SetFILEFOLDER_ATTRIBUTES sets FILEFOLDER_ATTRIBUTES field to given value.

### HasFILEFOLDER_ATTRIBUTES

`func (o *ConnectionAttributes) HasFILEFOLDER_ATTRIBUTES() bool`

HasFILEFOLDER_ATTRIBUTES returns a boolean if a field has been set.

### GetFILTER

`func (o *ConnectionAttributes) GetFILTER() string`

GetFILTER returns the FILTER field if non-nil, zero value otherwise.

### GetFILTEROk

`func (o *ConnectionAttributes) GetFILTEROk() (*string, bool)`

GetFILTEROk returns a tuple with the FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILTER

`func (o *ConnectionAttributes) SetFILTER(v string)`

SetFILTER sets FILTER field to given value.

### HasFILTER

`func (o *ConnectionAttributes) HasFILTER() bool`

HasFILTER returns a boolean if a field has been set.

### GetGROUP_ATTRIBUTES

`func (o *ConnectionAttributes) GetGROUP_ATTRIBUTES() string`

GetGROUP_ATTRIBUTES returns the GROUP_ATTRIBUTES field if non-nil, zero value otherwise.

### GetGROUP_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetGROUP_ATTRIBUTESOk() (*string, bool)`

GetGROUP_ATTRIBUTESOk returns a tuple with the GROUP_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUP_ATTRIBUTES

`func (o *ConnectionAttributes) SetGROUP_ATTRIBUTES(v string)`

SetGROUP_ATTRIBUTES sets GROUP_ATTRIBUTES field to given value.

### HasGROUP_ATTRIBUTES

`func (o *ConnectionAttributes) HasGROUP_ATTRIBUTES() bool`

HasGROUP_ATTRIBUTES returns a boolean if a field has been set.

### GetIMPORT_USERS

`func (o *ConnectionAttributes) GetIMPORT_USERS() string`

GetIMPORT_USERS returns the IMPORT_USERS field if non-nil, zero value otherwise.

### GetIMPORT_USERSOk

`func (o *ConnectionAttributes) GetIMPORT_USERSOk() (*string, bool)`

GetIMPORT_USERSOk returns a tuple with the IMPORT_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORT_USERS

`func (o *ConnectionAttributes) SetIMPORT_USERS(v string)`

SetIMPORT_USERS sets IMPORT_USERS field to given value.

### HasIMPORT_USERS

`func (o *ConnectionAttributes) HasIMPORT_USERS() bool`

HasIMPORT_USERS returns a boolean if a field has been set.

### GetIMPORTJSON

`func (o *ConnectionAttributes) GetIMPORTJSON() string`

GetIMPORTJSON returns the IMPORTJSON field if non-nil, zero value otherwise.

### GetIMPORTJSONOk

`func (o *ConnectionAttributes) GetIMPORTJSONOk() (*string, bool)`

GetIMPORTJSONOk returns a tuple with the IMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTJSON

`func (o *ConnectionAttributes) SetIMPORTJSON(v string)`

SetIMPORTJSON sets IMPORTJSON field to given value.

### HasIMPORTJSON

`func (o *ConnectionAttributes) HasIMPORTJSON() bool`

HasIMPORTJSON returns a boolean if a field has been set.

### GetINCREMENTAL_CONFIG

`func (o *ConnectionAttributes) GetINCREMENTAL_CONFIG() string`

GetINCREMENTAL_CONFIG returns the INCREMENTAL_CONFIG field if non-nil, zero value otherwise.

### GetINCREMENTAL_CONFIGOk

`func (o *ConnectionAttributes) GetINCREMENTAL_CONFIGOk() (*string, bool)`

GetINCREMENTAL_CONFIGOk returns a tuple with the INCREMENTAL_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCREMENTAL_CONFIG

`func (o *ConnectionAttributes) SetINCREMENTAL_CONFIG(v string)`

SetINCREMENTAL_CONFIG sets INCREMENTAL_CONFIG field to given value.

### HasINCREMENTAL_CONFIG

`func (o *ConnectionAttributes) HasINCREMENTAL_CONFIG() bool`

HasINCREMENTAL_CONFIG returns a boolean if a field has been set.

### GetLAST_IMPORT_TIME

`func (o *ConnectionAttributes) GetLAST_IMPORT_TIME() string`

GetLAST_IMPORT_TIME returns the LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetLAST_IMPORT_TIMEOk

`func (o *ConnectionAttributes) GetLAST_IMPORT_TIMEOk() (*string, bool)`

GetLAST_IMPORT_TIMEOk returns a tuple with the LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLAST_IMPORT_TIME

`func (o *ConnectionAttributes) SetLAST_IMPORT_TIME(v string)`

SetLAST_IMPORT_TIME sets LAST_IMPORT_TIME field to given value.

### HasLAST_IMPORT_TIME

`func (o *ConnectionAttributes) HasLAST_IMPORT_TIME() bool`

HasLAST_IMPORT_TIME returns a boolean if a field has been set.

### GetLDAP_OR_AD

`func (o *ConnectionAttributes) GetLDAP_OR_AD() string`

GetLDAP_OR_AD returns the LDAP_OR_AD field if non-nil, zero value otherwise.

### GetLDAP_OR_ADOk

`func (o *ConnectionAttributes) GetLDAP_OR_ADOk() (*string, bool)`

GetLDAP_OR_ADOk returns a tuple with the LDAP_OR_AD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_OR_AD

`func (o *ConnectionAttributes) SetLDAP_OR_AD(v string)`

SetLDAP_OR_AD sets LDAP_OR_AD field to given value.

### HasLDAP_OR_AD

`func (o *ConnectionAttributes) HasLDAP_OR_AD() bool`

HasLDAP_OR_AD returns a boolean if a field has been set.

### GetMAX_CHANGENUMBER

`func (o *ConnectionAttributes) GetMAX_CHANGENUMBER() string`

GetMAX_CHANGENUMBER returns the MAX_CHANGENUMBER field if non-nil, zero value otherwise.

### GetMAX_CHANGENUMBEROk

`func (o *ConnectionAttributes) GetMAX_CHANGENUMBEROk() (*string, bool)`

GetMAX_CHANGENUMBEROk returns a tuple with the MAX_CHANGENUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_CHANGENUMBER

`func (o *ConnectionAttributes) SetMAX_CHANGENUMBER(v string)`

SetMAX_CHANGENUMBER sets MAX_CHANGENUMBER field to given value.

### HasMAX_CHANGENUMBER

`func (o *ConnectionAttributes) HasMAX_CHANGENUMBER() bool`

HasMAX_CHANGENUMBER returns a boolean if a field has been set.

### GetMKTPLACEAPP_ATTRIBUTES

`func (o *ConnectionAttributes) GetMKTPLACEAPP_ATTRIBUTES() string`

GetMKTPLACEAPP_ATTRIBUTES returns the MKTPLACEAPP_ATTRIBUTES field if non-nil, zero value otherwise.

### GetMKTPLACEAPP_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetMKTPLACEAPP_ATTRIBUTESOk() (*string, bool)`

GetMKTPLACEAPP_ATTRIBUTESOk returns a tuple with the MKTPLACEAPP_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMKTPLACEAPP_ATTRIBUTES

`func (o *ConnectionAttributes) SetMKTPLACEAPP_ATTRIBUTES(v string)`

SetMKTPLACEAPP_ATTRIBUTES sets MKTPLACEAPP_ATTRIBUTES field to given value.

### HasMKTPLACEAPP_ATTRIBUTES

`func (o *ConnectionAttributes) HasMKTPLACEAPP_ATTRIBUTES() bool`

HasMKTPLACEAPP_ATTRIBUTES returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *ConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *ConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *ConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *ConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *ConnectionAttributes) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *ConnectionAttributes) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *ConnectionAttributes) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *ConnectionAttributes) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetORG_BASE

`func (o *ConnectionAttributes) GetORG_BASE() string`

GetORG_BASE returns the ORG_BASE field if non-nil, zero value otherwise.

### GetORG_BASEOk

`func (o *ConnectionAttributes) GetORG_BASEOk() (*string, bool)`

GetORG_BASEOk returns a tuple with the ORG_BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORG_BASE

`func (o *ConnectionAttributes) SetORG_BASE(v string)`

SetORG_BASE sets ORG_BASE field to given value.

### HasORG_BASE

`func (o *ConnectionAttributes) HasORG_BASE() bool`

HasORG_BASE returns a boolean if a field has been set.

### GetORGANIZATION_ATTRIBUTE

`func (o *ConnectionAttributes) GetORGANIZATION_ATTRIBUTE() string`

GetORGANIZATION_ATTRIBUTE returns the ORGANIZATION_ATTRIBUTE field if non-nil, zero value otherwise.

### GetORGANIZATION_ATTRIBUTEOk

`func (o *ConnectionAttributes) GetORGANIZATION_ATTRIBUTEOk() (*string, bool)`

GetORGANIZATION_ATTRIBUTEOk returns a tuple with the ORGANIZATION_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_ATTRIBUTE

`func (o *ConnectionAttributes) SetORGANIZATION_ATTRIBUTE(v string)`

SetORGANIZATION_ATTRIBUTE sets ORGANIZATION_ATTRIBUTE field to given value.

### HasORGANIZATION_ATTRIBUTE

`func (o *ConnectionAttributes) HasORGANIZATION_ATTRIBUTE() bool`

HasORGANIZATION_ATTRIBUTE returns a boolean if a field has been set.

### GetORGIMPORTJSON

`func (o *ConnectionAttributes) GetORGIMPORTJSON() string`

GetORGIMPORTJSON returns the ORGIMPORTJSON field if non-nil, zero value otherwise.

### GetORGIMPORTJSONOk

`func (o *ConnectionAttributes) GetORGIMPORTJSONOk() (*string, bool)`

GetORGIMPORTJSONOk returns a tuple with the ORGIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGIMPORTJSON

`func (o *ConnectionAttributes) SetORGIMPORTJSON(v string)`

SetORGIMPORTJSON sets ORGIMPORTJSON field to given value.

### HasORGIMPORTJSON

`func (o *ConnectionAttributes) HasORGIMPORTJSON() bool`

HasORGIMPORTJSON returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *ConnectionAttributes) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *ConnectionAttributes) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *ConnectionAttributes) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *ConnectionAttributes) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *ConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *ConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *ConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *ConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPASSWORD

`func (o *ConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *ConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *ConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *ConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *ConnectionAttributes) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *ConnectionAttributes) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *ConnectionAttributes) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *ConnectionAttributes) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *ConnectionAttributes) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *ConnectionAttributes) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *ConnectionAttributes) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *ConnectionAttributes) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *ConnectionAttributes) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *ConnectionAttributes) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *ConnectionAttributes) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *ConnectionAttributes) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *ConnectionAttributes) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *ConnectionAttributes) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *ConnectionAttributes) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *ConnectionAttributes) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *ConnectionAttributes) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *ConnectionAttributes) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *ConnectionAttributes) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *ConnectionAttributes) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetQUARANTINE_FOLDER_ID

`func (o *ConnectionAttributes) GetQUARANTINE_FOLDER_ID() string`

GetQUARANTINE_FOLDER_ID returns the QUARANTINE_FOLDER_ID field if non-nil, zero value otherwise.

### GetQUARANTINE_FOLDER_IDOk

`func (o *ConnectionAttributes) GetQUARANTINE_FOLDER_IDOk() (*string, bool)`

GetQUARANTINE_FOLDER_IDOk returns a tuple with the QUARANTINE_FOLDER_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQUARANTINE_FOLDER_ID

`func (o *ConnectionAttributes) SetQUARANTINE_FOLDER_ID(v string)`

SetQUARANTINE_FOLDER_ID sets QUARANTINE_FOLDER_ID field to given value.

### HasQUARANTINE_FOLDER_ID

`func (o *ConnectionAttributes) HasQUARANTINE_FOLDER_ID() bool`

HasQUARANTINE_FOLDER_ID returns a boolean if a field has been set.

### GetREAD_OPERATIONAL_ATTRIBUTES

`func (o *ConnectionAttributes) GetREAD_OPERATIONAL_ATTRIBUTES() string`

GetREAD_OPERATIONAL_ATTRIBUTES returns the READ_OPERATIONAL_ATTRIBUTES field if non-nil, zero value otherwise.

### GetREAD_OPERATIONAL_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetREAD_OPERATIONAL_ATTRIBUTESOk() (*string, bool)`

GetREAD_OPERATIONAL_ATTRIBUTESOk returns a tuple with the READ_OPERATIONAL_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_OPERATIONAL_ATTRIBUTES

`func (o *ConnectionAttributes) SetREAD_OPERATIONAL_ATTRIBUTES(v string)`

SetREAD_OPERATIONAL_ATTRIBUTES sets READ_OPERATIONAL_ATTRIBUTES field to given value.

### HasREAD_OPERATIONAL_ATTRIBUTES

`func (o *ConnectionAttributes) HasREAD_OPERATIONAL_ATTRIBUTES() bool`

HasREAD_OPERATIONAL_ATTRIBUTES returns a boolean if a field has been set.

### GetREMOVEACCOUNTACTION

`func (o *ConnectionAttributes) GetREMOVEACCOUNTACTION() string`

GetREMOVEACCOUNTACTION returns the REMOVEACCOUNTACTION field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTACTIONOk

`func (o *ConnectionAttributes) GetREMOVEACCOUNTACTIONOk() (*string, bool)`

GetREMOVEACCOUNTACTIONOk returns a tuple with the REMOVEACCOUNTACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTACTION

`func (o *ConnectionAttributes) SetREMOVEACCOUNTACTION(v string)`

SetREMOVEACCOUNTACTION sets REMOVEACCOUNTACTION field to given value.

### HasREMOVEACCOUNTACTION

`func (o *ConnectionAttributes) HasREMOVEACCOUNTACTION() bool`

HasREMOVEACCOUNTACTION returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *ConnectionAttributes) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *ConnectionAttributes) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *ConnectionAttributes) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *ConnectionAttributes) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetREUSEACCOUNTJSON

`func (o *ConnectionAttributes) GetREUSEACCOUNTJSON() string`

GetREUSEACCOUNTJSON returns the REUSEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREUSEACCOUNTJSONOk

`func (o *ConnectionAttributes) GetREUSEACCOUNTJSONOk() (*string, bool)`

GetREUSEACCOUNTJSONOk returns a tuple with the REUSEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEACCOUNTJSON

`func (o *ConnectionAttributes) SetREUSEACCOUNTJSON(v string)`

SetREUSEACCOUNTJSON sets REUSEACCOUNTJSON field to given value.

### HasREUSEACCOUNTJSON

`func (o *ConnectionAttributes) HasREUSEACCOUNTJSON() bool`

HasREUSEACCOUNTJSON returns a boolean if a field has been set.

### GetREUSEINACTIVEACCOUNT

`func (o *ConnectionAttributes) GetREUSEINACTIVEACCOUNT() string`

GetREUSEINACTIVEACCOUNT returns the REUSEINACTIVEACCOUNT field if non-nil, zero value otherwise.

### GetREUSEINACTIVEACCOUNTOk

`func (o *ConnectionAttributes) GetREUSEINACTIVEACCOUNTOk() (*string, bool)`

GetREUSEINACTIVEACCOUNTOk returns a tuple with the REUSEINACTIVEACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEINACTIVEACCOUNT

`func (o *ConnectionAttributes) SetREUSEINACTIVEACCOUNT(v string)`

SetREUSEINACTIVEACCOUNT sets REUSEINACTIVEACCOUNT field to given value.

### HasREUSEINACTIVEACCOUNT

`func (o *ConnectionAttributes) HasREUSEINACTIVEACCOUNT() bool`

HasREUSEINACTIVEACCOUNT returns a boolean if a field has been set.

### GetSCOPES

`func (o *ConnectionAttributes) GetSCOPES() string`

GetSCOPES returns the SCOPES field if non-nil, zero value otherwise.

### GetSCOPESOk

`func (o *ConnectionAttributes) GetSCOPESOk() (*string, bool)`

GetSCOPESOk returns a tuple with the SCOPES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCOPES

`func (o *ConnectionAttributes) SetSCOPES(v string)`

SetSCOPES sets SCOPES field to given value.

### HasSCOPES

`func (o *ConnectionAttributes) HasSCOPES() bool`

HasSCOPES returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *ConnectionAttributes) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *ConnectionAttributes) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *ConnectionAttributes) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *ConnectionAttributes) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetSERVICE_ACCOUNT_ID

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_ID() string`

GetSERVICE_ACCOUNT_ID returns the SERVICE_ACCOUNT_ID field if non-nil, zero value otherwise.

### GetSERVICE_ACCOUNT_IDOk

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_IDOk() (*string, bool)`

GetSERVICE_ACCOUNT_IDOk returns a tuple with the SERVICE_ACCOUNT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVICE_ACCOUNT_ID

`func (o *ConnectionAttributes) SetSERVICE_ACCOUNT_ID(v string)`

SetSERVICE_ACCOUNT_ID sets SERVICE_ACCOUNT_ID field to given value.

### HasSERVICE_ACCOUNT_ID

`func (o *ConnectionAttributes) HasSERVICE_ACCOUNT_ID() bool`

HasSERVICE_ACCOUNT_ID returns a boolean if a field has been set.

### GetSERVICE_ACCOUNT_KEY_JSON

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_KEY_JSON() string`

GetSERVICE_ACCOUNT_KEY_JSON returns the SERVICE_ACCOUNT_KEY_JSON field if non-nil, zero value otherwise.

### GetSERVICE_ACCOUNT_KEY_JSONOk

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_KEY_JSONOk() (*string, bool)`

GetSERVICE_ACCOUNT_KEY_JSONOk returns a tuple with the SERVICE_ACCOUNT_KEY_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVICE_ACCOUNT_KEY_JSON

`func (o *ConnectionAttributes) SetSERVICE_ACCOUNT_KEY_JSON(v string)`

SetSERVICE_ACCOUNT_KEY_JSON sets SERVICE_ACCOUNT_KEY_JSON field to given value.

### HasSERVICE_ACCOUNT_KEY_JSON

`func (o *ConnectionAttributes) HasSERVICE_ACCOUNT_KEY_JSON() bool`

HasSERVICE_ACCOUNT_KEY_JSON returns a boolean if a field has been set.

### GetSERVICE_ACCOUNT_USER

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_USER() string`

GetSERVICE_ACCOUNT_USER returns the SERVICE_ACCOUNT_USER field if non-nil, zero value otherwise.

### GetSERVICE_ACCOUNT_USEROk

`func (o *ConnectionAttributes) GetSERVICE_ACCOUNT_USEROk() (*string, bool)`

GetSERVICE_ACCOUNT_USEROk returns a tuple with the SERVICE_ACCOUNT_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVICE_ACCOUNT_USER

`func (o *ConnectionAttributes) SetSERVICE_ACCOUNT_USER(v string)`

SetSERVICE_ACCOUNT_USER sets SERVICE_ACCOUNT_USER field to given value.

### HasSERVICE_ACCOUNT_USER

`func (o *ConnectionAttributes) HasSERVICE_ACCOUNT_USER() bool`

HasSERVICE_ACCOUNT_USER returns a boolean if a field has been set.

### GetSETDEFAULTPAGESIZE

`func (o *ConnectionAttributes) GetSETDEFAULTPAGESIZE() string`

GetSETDEFAULTPAGESIZE returns the SETDEFAULTPAGESIZE field if non-nil, zero value otherwise.

### GetSETDEFAULTPAGESIZEOk

`func (o *ConnectionAttributes) GetSETDEFAULTPAGESIZEOk() (*string, bool)`

GetSETDEFAULTPAGESIZEOk returns a tuple with the SETDEFAULTPAGESIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETDEFAULTPAGESIZE

`func (o *ConnectionAttributes) SetSETDEFAULTPAGESIZE(v string)`

SetSETDEFAULTPAGESIZE sets SETDEFAULTPAGESIZE field to given value.

### HasSETDEFAULTPAGESIZE

`func (o *ConnectionAttributes) HasSETDEFAULTPAGESIZE() bool`

HasSETDEFAULTPAGESIZE returns a boolean if a field has been set.

### GetSETRANDOMPASSWORD

`func (o *ConnectionAttributes) GetSETRANDOMPASSWORD() string`

GetSETRANDOMPASSWORD returns the SETRANDOMPASSWORD field if non-nil, zero value otherwise.

### GetSETRANDOMPASSWORDOk

`func (o *ConnectionAttributes) GetSETRANDOMPASSWORDOk() (*string, bool)`

GetSETRANDOMPASSWORDOk returns a tuple with the SETRANDOMPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETRANDOMPASSWORD

`func (o *ConnectionAttributes) SetSETRANDOMPASSWORD(v string)`

SetSETRANDOMPASSWORD sets SETRANDOMPASSWORD field to given value.

### HasSETRANDOMPASSWORD

`func (o *ConnectionAttributes) HasSETRANDOMPASSWORD() bool`

HasSETRANDOMPASSWORD returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *ConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *ConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *ConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *ConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *ConnectionAttributes) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *ConnectionAttributes) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *ConnectionAttributes) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *ConnectionAttributes) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetSUPPORTEMPTYSTRING

`func (o *ConnectionAttributes) GetSUPPORTEMPTYSTRING() string`

GetSUPPORTEMPTYSTRING returns the SUPPORTEMPTYSTRING field if non-nil, zero value otherwise.

### GetSUPPORTEMPTYSTRINGOk

`func (o *ConnectionAttributes) GetSUPPORTEMPTYSTRINGOk() (*string, bool)`

GetSUPPORTEMPTYSTRINGOk returns a tuple with the SUPPORTEMPTYSTRING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUPPORTEMPTYSTRING

`func (o *ConnectionAttributes) SetSUPPORTEMPTYSTRING(v string)`

SetSUPPORTEMPTYSTRING sets SUPPORTEMPTYSTRING field to given value.

### HasSUPPORTEMPTYSTRING

`func (o *ConnectionAttributes) HasSUPPORTEMPTYSTRING() bool`

HasSUPPORTEMPTYSTRING returns a boolean if a field has been set.

### GetSUSPEND_USER_ORGANIZATIONAL_UNIT

`func (o *ConnectionAttributes) GetSUSPEND_USER_ORGANIZATIONAL_UNIT() string`

GetSUSPEND_USER_ORGANIZATIONAL_UNIT returns the SUSPEND_USER_ORGANIZATIONAL_UNIT field if non-nil, zero value otherwise.

### GetSUSPEND_USER_ORGANIZATIONAL_UNITOk

`func (o *ConnectionAttributes) GetSUSPEND_USER_ORGANIZATIONAL_UNITOk() (*string, bool)`

GetSUSPEND_USER_ORGANIZATIONAL_UNITOk returns a tuple with the SUSPEND_USER_ORGANIZATIONAL_UNIT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUSPEND_USER_ORGANIZATIONAL_UNIT

`func (o *ConnectionAttributes) SetSUSPEND_USER_ORGANIZATIONAL_UNIT(v string)`

SetSUSPEND_USER_ORGANIZATIONAL_UNIT sets SUSPEND_USER_ORGANIZATIONAL_UNIT field to given value.

### HasSUSPEND_USER_ORGANIZATIONAL_UNIT

`func (o *ConnectionAttributes) HasSUSPEND_USER_ORGANIZATIONAL_UNIT() bool`

HasSUSPEND_USER_ORGANIZATIONAL_UNIT returns a boolean if a field has been set.

### GetUNLOCKACCOUNTJSON

`func (o *ConnectionAttributes) GetUNLOCKACCOUNTJSON() string`

GetUNLOCKACCOUNTJSON returns the UNLOCKACCOUNTJSON field if non-nil, zero value otherwise.

### GetUNLOCKACCOUNTJSONOk

`func (o *ConnectionAttributes) GetUNLOCKACCOUNTJSONOk() (*string, bool)`

GetUNLOCKACCOUNTJSONOk returns a tuple with the UNLOCKACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCKACCOUNTJSON

`func (o *ConnectionAttributes) SetUNLOCKACCOUNTJSON(v string)`

SetUNLOCKACCOUNTJSON sets UNLOCKACCOUNTJSON field to given value.

### HasUNLOCKACCOUNTJSON

`func (o *ConnectionAttributes) HasUNLOCKACCOUNTJSON() bool`

HasUNLOCKACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *ConnectionAttributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *ConnectionAttributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *ConnectionAttributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *ConnectionAttributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEORGJSON

`func (o *ConnectionAttributes) GetUPDATEORGJSON() string`

GetUPDATEORGJSON returns the UPDATEORGJSON field if non-nil, zero value otherwise.

### GetUPDATEORGJSONOk

`func (o *ConnectionAttributes) GetUPDATEORGJSONOk() (*string, bool)`

GetUPDATEORGJSONOk returns a tuple with the UPDATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEORGJSON

`func (o *ConnectionAttributes) SetUPDATEORGJSON(v string)`

SetUPDATEORGJSON sets UPDATEORGJSON field to given value.

### HasUPDATEORGJSON

`func (o *ConnectionAttributes) HasUPDATEORGJSON() bool`

HasUPDATEORGJSON returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *ConnectionAttributes) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *ConnectionAttributes) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *ConnectionAttributes) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *ConnectionAttributes) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *ConnectionAttributes) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *ConnectionAttributes) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *ConnectionAttributes) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *ConnectionAttributes) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetUSER_ATTRIBUTES

`func (o *ConnectionAttributes) GetUSER_ATTRIBUTES() string`

GetUSER_ATTRIBUTES returns the USER_ATTRIBUTES field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTESOk

`func (o *ConnectionAttributes) GetUSER_ATTRIBUTESOk() (*string, bool)`

GetUSER_ATTRIBUTESOk returns a tuple with the USER_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTES

`func (o *ConnectionAttributes) SetUSER_ATTRIBUTES(v string)`

SetUSER_ATTRIBUTES sets USER_ATTRIBUTES field to given value.

### HasUSER_ATTRIBUTES

`func (o *ConnectionAttributes) HasUSER_ATTRIBUTES() bool`

HasUSER_ATTRIBUTES returns a boolean if a field has been set.

### GetURL

`func (o *ConnectionAttributes) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *ConnectionAttributes) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *ConnectionAttributes) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *ConnectionAttributes) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetUSERNAME

`func (o *ConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *ConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *ConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *ConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *ConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *ConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *ConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *ConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetConnectionType

`func (o *ConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *ConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *ConnectionAttributes) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *ConnectionAttributes) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *ConnectionAttributes) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *ConnectionAttributes) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *ConnectionAttributes) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *ConnectionAttributes) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *ConnectionAttributes) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *ConnectionAttributes) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *ConnectionAttributes) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *ConnectionAttributes) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *ConnectionAttributes) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *ConnectionAttributes) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *ConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *ConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *ConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *ConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *ConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *ConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *ConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *ConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


