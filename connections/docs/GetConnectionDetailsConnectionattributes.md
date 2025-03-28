# GetConnectionDetailsConnectionattributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LAST_IMPORT_TIME** | Pointer to **string** | Timestamp of the last import operation. | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | JSON template for creating an account. | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for disabling an active account. | [optional] 
**GroupSearchBaseDN** | Pointer to **string** | Base DN for group search in LDAP. | [optional] 
**ASSWORD_NOOFSPLCHARS** | Pointer to **string** | Number of special characters in password. | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password. | [optional] 
**ConnectionType** | Pointer to **string** | Type of connection, e.g., MS_BASED_CONNECTOR. | [optional] 
**STATUSKEYJSON** | Pointer to **string** | JSON configuration for account statuses. | [optional] 
**SEARCHFILTER** | Pointer to **string** | LDAP search filter. | [optional] 
**ConfigJSON** | Pointer to **string** | We can use this attribute to define the provisioningLimit,showLogs and connectionTimeoutConfig. | [optional] 
**REMOVEACCOUNTACTION** | Pointer to **string** | Action performed on account removal. | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** | Map LDAP user attribute to EIC account attribute for import | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** | Rule to generate account name. | [optional] 
**ADVSEARCH** | Pointer to **string** | Advanced search criteria. | [optional] 
**USERNAME** | Pointer to **string** | LDAP or system admin username. | [optional] 
**PASSWORD** | Pointer to **string** | Set the Password | [optional] 
**LDAP_OR_AD** | Pointer to **string** | Type of Endpoint - LDAP or AD, default value is AD | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | Attribute used for entitlements. | [optional] 
**SETRANDOMPASSWORD** | Pointer to **string** | Set random password option. | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Minimum password length. | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Maximum password length. | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Specify the number of special characters required for the random password. | [optional] 
**SETDEFAULTPAGESIZE** | Pointer to **string** | Default page size for data retrieval. | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**REUSEINACTIVEACCOUNT** | Pointer to **string** | Flag to determine if inactive accounts can be reused. | [optional] 
**IMPORTJSON** | Pointer to **string** | JSON import configurations. | [optional] 
**CreateUpdateMappings** | Pointer to **string** | Mapping used during group creation and updation. | [optional] 
**ADVANCE_FILTER_JSON** | Pointer to **string** | Advanced filter criteria in JSON format. | [optional] 
**ORGIMPORTJSON** | Pointer to **string** | JSON configuration for importing organizations. | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** | LDAP pagination page size. | [optional] 
**BASE** | Pointer to **string** | LDAP base DN. | [optional] 
**DC_LOCATOR** | Pointer to **string** | Domain Controller locator. | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | JSON configuration for account status thresholds. | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** | JSON actions for password reset/change. | [optional] 
**SUPPORTEMPTYSTRING** | Pointer to **string** | Set to TRUE to send an empty value or null string during provisioning. | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON configuration to enable account actions. | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** | Map LDAP user attribute to EIC user attribute for import. | [optional] 
**DEFAULT_USER_ROLE** | Pointer to **string** | Default SAV Role to be assigned to all the new users imported via User Import. | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | To update an existing account, specify the value of the UPDATEACCOUNTJSON parameter. | [optional] 
**REUSEACCOUNTJSON** | Pointer to **string** | To retain the suspended user accounts, specify the value of the REUSEACCOUNTJSON parameter. | [optional] 
**ENFORCE_TREE_DELETION** | Pointer to **string** | Enforce tree deletion in LDAP. | [optional] 
**FILTER** | Pointer to **string** | Filters the objects that will be returned. | [optional] 
**OBJECTFILTER** | Pointer to **string** | LDAP object filter. | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | JSON template for updating user attributes. | [optional] 
**Saveconnection** | Pointer to **string** | Flag to save connection permanently. | [optional] 
**Systemname** | Pointer to **string** | Associated system name. | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Number of special characters in password. | [optional] 
**GroupImportMapping** | Pointer to **string** | Mapping for importing groups. | [optional] 
**UNLOCKACCOUNTJSON** | Pointer to **string** | JSON template for unlocking accounts. | [optional] 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** | Specify TRUE to enable Group Management for AD. | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**ORG_BASE** | Pointer to **string** | Base DN for organization searches. | [optional] 
**ORGANIZATION_ATTRIBUTE** | Pointer to **string** | Attribute used for organization management. | [optional] 
**CREATEORGJSON** | Pointer to **string** | JSON template for creating an organization. | [optional] 
**UPDATEORGJSON** | Pointer to **string** | JSON template for updating an organization. | [optional] 
**MAX_CHANGENUMBER** | Pointer to **string** | Maximum change number for incremental syncs. | [optional] 
**INCREMENTAL_CONFIG** | Pointer to **string** | Configuration for incremental data sync. | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** | Rules for checking unique users in JSON format. | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**RESTConnectionAttributesConnectionTimeoutConfig**](RESTConnectionAttributesConnectionTimeoutConfig.md) |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**UpdateUserJSON** | Pointer to **string** | Property for UpdateUserJSON | [optional] 
**ChangePassJSON** | Pointer to **string** |  | [optional] 
**RemoveAccountJSON** | Pointer to **string** |  | [optional] 
**TicketStatusJSON** | Pointer to **string** |  | [optional] 
**CreateTicketJSON** | Pointer to **string** |  | [optional] 
**PasswdPolicyJSON** | Pointer to **string** |  | [optional] 
**AddFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**RemoveFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**SendOtpJSON** | Pointer to **string** |  | [optional] 
**ValidateOtpJSON** | Pointer to **string** |  | [optional] 
**CreateAccountJSON** | Pointer to **string** |  | [optional] 
**UpdateAccountJSON** | Pointer to **string** |  | [optional] 
**EnableAccountJSON** | Pointer to **string** |  | [optional] 
**DisableAccountJSON** | Pointer to **string** |  | [optional] 
**AddAccessJSON** | Pointer to **string** |  | [optional] 
**RemoveAccessJSON** | Pointer to **string** |  | [optional] 
**ImportUserJSON** | Pointer to **string** |  | [optional] 
**ImportAccountEntJSON** | Pointer to **string** |  | [optional] 
**ConnectionJSON** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetConnectionDetailsConnectionattributes

`func NewGetConnectionDetailsConnectionattributes() *GetConnectionDetailsConnectionattributes`

NewGetConnectionDetailsConnectionattributes instantiates a new GetConnectionDetailsConnectionattributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionDetailsConnectionattributesWithDefaults

`func NewGetConnectionDetailsConnectionattributesWithDefaults() *GetConnectionDetailsConnectionattributes`

NewGetConnectionDetailsConnectionattributesWithDefaults instantiates a new GetConnectionDetailsConnectionattributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLAST_IMPORT_TIME

`func (o *GetConnectionDetailsConnectionattributes) GetLAST_IMPORT_TIME() string`

GetLAST_IMPORT_TIME returns the LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetLAST_IMPORT_TIMEOk

`func (o *GetConnectionDetailsConnectionattributes) GetLAST_IMPORT_TIMEOk() (*string, bool)`

GetLAST_IMPORT_TIMEOk returns a tuple with the LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLAST_IMPORT_TIME

`func (o *GetConnectionDetailsConnectionattributes) SetLAST_IMPORT_TIME(v string)`

SetLAST_IMPORT_TIME sets LAST_IMPORT_TIME field to given value.

### HasLAST_IMPORT_TIME

`func (o *GetConnectionDetailsConnectionattributes) HasLAST_IMPORT_TIME() bool`

HasLAST_IMPORT_TIME returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *GetConnectionDetailsConnectionattributes) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *GetConnectionDetailsConnectionattributes) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *GetConnectionDetailsConnectionattributes) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *GetConnectionDetailsConnectionattributes) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) GetASSWORD_NOOFSPLCHARS() string`

GetASSWORD_NOOFSPLCHARS returns the ASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetASSWORD_NOOFSPLCHARSOk

`func (o *GetConnectionDetailsConnectionattributes) GetASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetASSWORD_NOOFSPLCHARSOk returns a tuple with the ASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) SetASSWORD_NOOFSPLCHARS(v string)`

SetASSWORD_NOOFSPLCHARS sets ASSWORD_NOOFSPLCHARS field to given value.

### HasASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) HasASSWORD_NOOFSPLCHARS() bool`

HasASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetConnectionType

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetConnectionDetailsConnectionattributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetConnectionDetailsConnectionattributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *GetConnectionDetailsConnectionattributes) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *GetConnectionDetailsConnectionattributes) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *GetConnectionDetailsConnectionattributes) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *GetConnectionDetailsConnectionattributes) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *GetConnectionDetailsConnectionattributes) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *GetConnectionDetailsConnectionattributes) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *GetConnectionDetailsConnectionattributes) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetConfigJSON

`func (o *GetConnectionDetailsConnectionattributes) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *GetConnectionDetailsConnectionattributes) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *GetConnectionDetailsConnectionattributes) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetREMOVEACCOUNTACTION

`func (o *GetConnectionDetailsConnectionattributes) GetREMOVEACCOUNTACTION() string`

GetREMOVEACCOUNTACTION returns the REMOVEACCOUNTACTION field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTACTIONOk

`func (o *GetConnectionDetailsConnectionattributes) GetREMOVEACCOUNTACTIONOk() (*string, bool)`

GetREMOVEACCOUNTACTIONOk returns a tuple with the REMOVEACCOUNTACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTACTION

`func (o *GetConnectionDetailsConnectionattributes) SetREMOVEACCOUNTACTION(v string)`

SetREMOVEACCOUNTACTION sets REMOVEACCOUNTACTION field to given value.

### HasREMOVEACCOUNTACTION

`func (o *GetConnectionDetailsConnectionattributes) HasREMOVEACCOUNTACTION() bool`

HasREMOVEACCOUNTACTION returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *GetConnectionDetailsConnectionattributes) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *GetConnectionDetailsConnectionattributes) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *GetConnectionDetailsConnectionattributes) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *GetConnectionDetailsConnectionattributes) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *GetConnectionDetailsConnectionattributes) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetADVSEARCH

`func (o *GetConnectionDetailsConnectionattributes) GetADVSEARCH() string`

GetADVSEARCH returns the ADVSEARCH field if non-nil, zero value otherwise.

### GetADVSEARCHOk

`func (o *GetConnectionDetailsConnectionattributes) GetADVSEARCHOk() (*string, bool)`

GetADVSEARCHOk returns a tuple with the ADVSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVSEARCH

`func (o *GetConnectionDetailsConnectionattributes) SetADVSEARCH(v string)`

SetADVSEARCH sets ADVSEARCH field to given value.

### HasADVSEARCH

`func (o *GetConnectionDetailsConnectionattributes) HasADVSEARCH() bool`

HasADVSEARCH returns a boolean if a field has been set.

### GetUSERNAME

`func (o *GetConnectionDetailsConnectionattributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *GetConnectionDetailsConnectionattributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *GetConnectionDetailsConnectionattributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *GetConnectionDetailsConnectionattributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetLDAP_OR_AD

`func (o *GetConnectionDetailsConnectionattributes) GetLDAP_OR_AD() string`

GetLDAP_OR_AD returns the LDAP_OR_AD field if non-nil, zero value otherwise.

### GetLDAP_OR_ADOk

`func (o *GetConnectionDetailsConnectionattributes) GetLDAP_OR_ADOk() (*string, bool)`

GetLDAP_OR_ADOk returns a tuple with the LDAP_OR_AD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_OR_AD

`func (o *GetConnectionDetailsConnectionattributes) SetLDAP_OR_AD(v string)`

SetLDAP_OR_AD sets LDAP_OR_AD field to given value.

### HasLDAP_OR_AD

`func (o *GetConnectionDetailsConnectionattributes) HasLDAP_OR_AD() bool`

HasLDAP_OR_AD returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *GetConnectionDetailsConnectionattributes) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetSETRANDOMPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) GetSETRANDOMPASSWORD() string`

GetSETRANDOMPASSWORD returns the SETRANDOMPASSWORD field if non-nil, zero value otherwise.

### GetSETRANDOMPASSWORDOk

`func (o *GetConnectionDetailsConnectionattributes) GetSETRANDOMPASSWORDOk() (*string, bool)`

GetSETRANDOMPASSWORDOk returns a tuple with the SETRANDOMPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETRANDOMPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) SetSETRANDOMPASSWORD(v string)`

SetSETRANDOMPASSWORD sets SETRANDOMPASSWORD field to given value.

### HasSETRANDOMPASSWORD

`func (o *GetConnectionDetailsConnectionattributes) HasSETRANDOMPASSWORD() bool`

HasSETRANDOMPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetSETDEFAULTPAGESIZE

`func (o *GetConnectionDetailsConnectionattributes) GetSETDEFAULTPAGESIZE() string`

GetSETDEFAULTPAGESIZE returns the SETDEFAULTPAGESIZE field if non-nil, zero value otherwise.

### GetSETDEFAULTPAGESIZEOk

`func (o *GetConnectionDetailsConnectionattributes) GetSETDEFAULTPAGESIZEOk() (*string, bool)`

GetSETDEFAULTPAGESIZEOk returns a tuple with the SETDEFAULTPAGESIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETDEFAULTPAGESIZE

`func (o *GetConnectionDetailsConnectionattributes) SetSETDEFAULTPAGESIZE(v string)`

SetSETDEFAULTPAGESIZE sets SETDEFAULTPAGESIZE field to given value.

### HasSETDEFAULTPAGESIZE

`func (o *GetConnectionDetailsConnectionattributes) HasSETDEFAULTPAGESIZE() bool`

HasSETDEFAULTPAGESIZE returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *GetConnectionDetailsConnectionattributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *GetConnectionDetailsConnectionattributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *GetConnectionDetailsConnectionattributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *GetConnectionDetailsConnectionattributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetREUSEINACTIVEACCOUNT

`func (o *GetConnectionDetailsConnectionattributes) GetREUSEINACTIVEACCOUNT() string`

GetREUSEINACTIVEACCOUNT returns the REUSEINACTIVEACCOUNT field if non-nil, zero value otherwise.

### GetREUSEINACTIVEACCOUNTOk

`func (o *GetConnectionDetailsConnectionattributes) GetREUSEINACTIVEACCOUNTOk() (*string, bool)`

GetREUSEINACTIVEACCOUNTOk returns a tuple with the REUSEINACTIVEACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEINACTIVEACCOUNT

`func (o *GetConnectionDetailsConnectionattributes) SetREUSEINACTIVEACCOUNT(v string)`

SetREUSEINACTIVEACCOUNT sets REUSEINACTIVEACCOUNT field to given value.

### HasREUSEINACTIVEACCOUNT

`func (o *GetConnectionDetailsConnectionattributes) HasREUSEINACTIVEACCOUNT() bool`

HasREUSEINACTIVEACCOUNT returns a boolean if a field has been set.

### GetIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetIMPORTJSON() string`

GetIMPORTJSON returns the IMPORTJSON field if non-nil, zero value otherwise.

### GetIMPORTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetIMPORTJSONOk() (*string, bool)`

GetIMPORTJSONOk returns a tuple with the IMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetIMPORTJSON(v string)`

SetIMPORTJSON sets IMPORTJSON field to given value.

### HasIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasIMPORTJSON() bool`

HasIMPORTJSON returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *GetConnectionDetailsConnectionattributes) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *GetConnectionDetailsConnectionattributes) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *GetConnectionDetailsConnectionattributes) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *GetConnectionDetailsConnectionattributes) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetADVANCE_FILTER_JSON

`func (o *GetConnectionDetailsConnectionattributes) GetADVANCE_FILTER_JSON() string`

GetADVANCE_FILTER_JSON returns the ADVANCE_FILTER_JSON field if non-nil, zero value otherwise.

### GetADVANCE_FILTER_JSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetADVANCE_FILTER_JSONOk() (*string, bool)`

GetADVANCE_FILTER_JSONOk returns a tuple with the ADVANCE_FILTER_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVANCE_FILTER_JSON

`func (o *GetConnectionDetailsConnectionattributes) SetADVANCE_FILTER_JSON(v string)`

SetADVANCE_FILTER_JSON sets ADVANCE_FILTER_JSON field to given value.

### HasADVANCE_FILTER_JSON

`func (o *GetConnectionDetailsConnectionattributes) HasADVANCE_FILTER_JSON() bool`

HasADVANCE_FILTER_JSON returns a boolean if a field has been set.

### GetORGIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetORGIMPORTJSON() string`

GetORGIMPORTJSON returns the ORGIMPORTJSON field if non-nil, zero value otherwise.

### GetORGIMPORTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetORGIMPORTJSONOk() (*string, bool)`

GetORGIMPORTJSONOk returns a tuple with the ORGIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetORGIMPORTJSON(v string)`

SetORGIMPORTJSON sets ORGIMPORTJSON field to given value.

### HasORGIMPORTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasORGIMPORTJSON() bool`

HasORGIMPORTJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *GetConnectionDetailsConnectionattributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *GetConnectionDetailsConnectionattributes) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *GetConnectionDetailsConnectionattributes) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *GetConnectionDetailsConnectionattributes) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *GetConnectionDetailsConnectionattributes) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetBASE

`func (o *GetConnectionDetailsConnectionattributes) GetBASE() string`

GetBASE returns the BASE field if non-nil, zero value otherwise.

### GetBASEOk

`func (o *GetConnectionDetailsConnectionattributes) GetBASEOk() (*string, bool)`

GetBASEOk returns a tuple with the BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE

`func (o *GetConnectionDetailsConnectionattributes) SetBASE(v string)`

SetBASE sets BASE field to given value.

### HasBASE

`func (o *GetConnectionDetailsConnectionattributes) HasBASE() bool`

HasBASE returns a boolean if a field has been set.

### GetDC_LOCATOR

`func (o *GetConnectionDetailsConnectionattributes) GetDC_LOCATOR() string`

GetDC_LOCATOR returns the DC_LOCATOR field if non-nil, zero value otherwise.

### GetDC_LOCATOROk

`func (o *GetConnectionDetailsConnectionattributes) GetDC_LOCATOROk() (*string, bool)`

GetDC_LOCATOROk returns a tuple with the DC_LOCATOR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDC_LOCATOR

`func (o *GetConnectionDetailsConnectionattributes) SetDC_LOCATOR(v string)`

SetDC_LOCATOR sets DC_LOCATOR field to given value.

### HasDC_LOCATOR

`func (o *GetConnectionDetailsConnectionattributes) HasDC_LOCATOR() bool`

HasDC_LOCATOR returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *GetConnectionDetailsConnectionattributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *GetConnectionDetailsConnectionattributes) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *GetConnectionDetailsConnectionattributes) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *GetConnectionDetailsConnectionattributes) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetSUPPORTEMPTYSTRING

`func (o *GetConnectionDetailsConnectionattributes) GetSUPPORTEMPTYSTRING() string`

GetSUPPORTEMPTYSTRING returns the SUPPORTEMPTYSTRING field if non-nil, zero value otherwise.

### GetSUPPORTEMPTYSTRINGOk

`func (o *GetConnectionDetailsConnectionattributes) GetSUPPORTEMPTYSTRINGOk() (*string, bool)`

GetSUPPORTEMPTYSTRINGOk returns a tuple with the SUPPORTEMPTYSTRING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUPPORTEMPTYSTRING

`func (o *GetConnectionDetailsConnectionattributes) SetSUPPORTEMPTYSTRING(v string)`

SetSUPPORTEMPTYSTRING sets SUPPORTEMPTYSTRING field to given value.

### HasSUPPORTEMPTYSTRING

`func (o *GetConnectionDetailsConnectionattributes) HasSUPPORTEMPTYSTRING() bool`

HasSUPPORTEMPTYSTRING returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *GetConnectionDetailsConnectionattributes) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetDEFAULT_USER_ROLE

`func (o *GetConnectionDetailsConnectionattributes) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *GetConnectionDetailsConnectionattributes) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *GetConnectionDetailsConnectionattributes) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *GetConnectionDetailsConnectionattributes) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *GetConnectionDetailsConnectionattributes) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *GetConnectionDetailsConnectionattributes) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *GetConnectionDetailsConnectionattributes) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *GetConnectionDetailsConnectionattributes) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetREUSEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetREUSEACCOUNTJSON() string`

GetREUSEACCOUNTJSON returns the REUSEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREUSEACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetREUSEACCOUNTJSONOk() (*string, bool)`

GetREUSEACCOUNTJSONOk returns a tuple with the REUSEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetREUSEACCOUNTJSON(v string)`

SetREUSEACCOUNTJSON sets REUSEACCOUNTJSON field to given value.

### HasREUSEACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasREUSEACCOUNTJSON() bool`

HasREUSEACCOUNTJSON returns a boolean if a field has been set.

### GetENFORCE_TREE_DELETION

`func (o *GetConnectionDetailsConnectionattributes) GetENFORCE_TREE_DELETION() string`

GetENFORCE_TREE_DELETION returns the ENFORCE_TREE_DELETION field if non-nil, zero value otherwise.

### GetENFORCE_TREE_DELETIONOk

`func (o *GetConnectionDetailsConnectionattributes) GetENFORCE_TREE_DELETIONOk() (*string, bool)`

GetENFORCE_TREE_DELETIONOk returns a tuple with the ENFORCE_TREE_DELETION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCE_TREE_DELETION

`func (o *GetConnectionDetailsConnectionattributes) SetENFORCE_TREE_DELETION(v string)`

SetENFORCE_TREE_DELETION sets ENFORCE_TREE_DELETION field to given value.

### HasENFORCE_TREE_DELETION

`func (o *GetConnectionDetailsConnectionattributes) HasENFORCE_TREE_DELETION() bool`

HasENFORCE_TREE_DELETION returns a boolean if a field has been set.

### GetFILTER

`func (o *GetConnectionDetailsConnectionattributes) GetFILTER() string`

GetFILTER returns the FILTER field if non-nil, zero value otherwise.

### GetFILTEROk

`func (o *GetConnectionDetailsConnectionattributes) GetFILTEROk() (*string, bool)`

GetFILTEROk returns a tuple with the FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILTER

`func (o *GetConnectionDetailsConnectionattributes) SetFILTER(v string)`

SetFILTER sets FILTER field to given value.

### HasFILTER

`func (o *GetConnectionDetailsConnectionattributes) HasFILTER() bool`

HasFILTER returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *GetConnectionDetailsConnectionattributes) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *GetConnectionDetailsConnectionattributes) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *GetConnectionDetailsConnectionattributes) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *GetConnectionDetailsConnectionattributes) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetSaveconnection

`func (o *GetConnectionDetailsConnectionattributes) GetSaveconnection() string`

GetSaveconnection returns the Saveconnection field if non-nil, zero value otherwise.

### GetSaveconnectionOk

`func (o *GetConnectionDetailsConnectionattributes) GetSaveconnectionOk() (*string, bool)`

GetSaveconnectionOk returns a tuple with the Saveconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveconnection

`func (o *GetConnectionDetailsConnectionattributes) SetSaveconnection(v string)`

SetSaveconnection sets Saveconnection field to given value.

### HasSaveconnection

`func (o *GetConnectionDetailsConnectionattributes) HasSaveconnection() bool`

HasSaveconnection returns a boolean if a field has been set.

### GetSystemname

`func (o *GetConnectionDetailsConnectionattributes) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *GetConnectionDetailsConnectionattributes) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *GetConnectionDetailsConnectionattributes) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.

### HasSystemname

`func (o *GetConnectionDetailsConnectionattributes) HasSystemname() bool`

HasSystemname returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *GetConnectionDetailsConnectionattributes) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *GetConnectionDetailsConnectionattributes) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *GetConnectionDetailsConnectionattributes) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *GetConnectionDetailsConnectionattributes) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *GetConnectionDetailsConnectionattributes) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *GetConnectionDetailsConnectionattributes) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetUNLOCKACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUNLOCKACCOUNTJSON() string`

GetUNLOCKACCOUNTJSON returns the UNLOCKACCOUNTJSON field if non-nil, zero value otherwise.

### GetUNLOCKACCOUNTJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUNLOCKACCOUNTJSONOk() (*string, bool)`

GetUNLOCKACCOUNTJSONOk returns a tuple with the UNLOCKACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCKACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUNLOCKACCOUNTJSON(v string)`

SetUNLOCKACCOUNTJSON sets UNLOCKACCOUNTJSON field to given value.

### HasUNLOCKACCOUNTJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUNLOCKACCOUNTJSON() bool`

HasUNLOCKACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEGROUPMANAGEMENT

`func (o *GetConnectionDetailsConnectionattributes) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *GetConnectionDetailsConnectionattributes) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *GetConnectionDetailsConnectionattributes) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *GetConnectionDetailsConnectionattributes) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *GetConnectionDetailsConnectionattributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *GetConnectionDetailsConnectionattributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *GetConnectionDetailsConnectionattributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetORG_BASE

`func (o *GetConnectionDetailsConnectionattributes) GetORG_BASE() string`

GetORG_BASE returns the ORG_BASE field if non-nil, zero value otherwise.

### GetORG_BASEOk

`func (o *GetConnectionDetailsConnectionattributes) GetORG_BASEOk() (*string, bool)`

GetORG_BASEOk returns a tuple with the ORG_BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORG_BASE

`func (o *GetConnectionDetailsConnectionattributes) SetORG_BASE(v string)`

SetORG_BASE sets ORG_BASE field to given value.

### HasORG_BASE

`func (o *GetConnectionDetailsConnectionattributes) HasORG_BASE() bool`

HasORG_BASE returns a boolean if a field has been set.

### GetORGANIZATION_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) GetORGANIZATION_ATTRIBUTE() string`

GetORGANIZATION_ATTRIBUTE returns the ORGANIZATION_ATTRIBUTE field if non-nil, zero value otherwise.

### GetORGANIZATION_ATTRIBUTEOk

`func (o *GetConnectionDetailsConnectionattributes) GetORGANIZATION_ATTRIBUTEOk() (*string, bool)`

GetORGANIZATION_ATTRIBUTEOk returns a tuple with the ORGANIZATION_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) SetORGANIZATION_ATTRIBUTE(v string)`

SetORGANIZATION_ATTRIBUTE sets ORGANIZATION_ATTRIBUTE field to given value.

### HasORGANIZATION_ATTRIBUTE

`func (o *GetConnectionDetailsConnectionattributes) HasORGANIZATION_ATTRIBUTE() bool`

HasORGANIZATION_ATTRIBUTE returns a boolean if a field has been set.

### GetCREATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) GetCREATEORGJSON() string`

GetCREATEORGJSON returns the CREATEORGJSON field if non-nil, zero value otherwise.

### GetCREATEORGJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetCREATEORGJSONOk() (*string, bool)`

GetCREATEORGJSONOk returns a tuple with the CREATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) SetCREATEORGJSON(v string)`

SetCREATEORGJSON sets CREATEORGJSON field to given value.

### HasCREATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) HasCREATEORGJSON() bool`

HasCREATEORGJSON returns a boolean if a field has been set.

### GetUPDATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEORGJSON() string`

GetUPDATEORGJSON returns the UPDATEORGJSON field if non-nil, zero value otherwise.

### GetUPDATEORGJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUPDATEORGJSONOk() (*string, bool)`

GetUPDATEORGJSONOk returns a tuple with the UPDATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUPDATEORGJSON(v string)`

SetUPDATEORGJSON sets UPDATEORGJSON field to given value.

### HasUPDATEORGJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUPDATEORGJSON() bool`

HasUPDATEORGJSON returns a boolean if a field has been set.

### GetMAX_CHANGENUMBER

`func (o *GetConnectionDetailsConnectionattributes) GetMAX_CHANGENUMBER() string`

GetMAX_CHANGENUMBER returns the MAX_CHANGENUMBER field if non-nil, zero value otherwise.

### GetMAX_CHANGENUMBEROk

`func (o *GetConnectionDetailsConnectionattributes) GetMAX_CHANGENUMBEROk() (*string, bool)`

GetMAX_CHANGENUMBEROk returns a tuple with the MAX_CHANGENUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_CHANGENUMBER

`func (o *GetConnectionDetailsConnectionattributes) SetMAX_CHANGENUMBER(v string)`

SetMAX_CHANGENUMBER sets MAX_CHANGENUMBER field to given value.

### HasMAX_CHANGENUMBER

`func (o *GetConnectionDetailsConnectionattributes) HasMAX_CHANGENUMBER() bool`

HasMAX_CHANGENUMBER returns a boolean if a field has been set.

### GetINCREMENTAL_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) GetINCREMENTAL_CONFIG() string`

GetINCREMENTAL_CONFIG returns the INCREMENTAL_CONFIG field if non-nil, zero value otherwise.

### GetINCREMENTAL_CONFIGOk

`func (o *GetConnectionDetailsConnectionattributes) GetINCREMENTAL_CONFIGOk() (*string, bool)`

GetINCREMENTAL_CONFIGOk returns a tuple with the INCREMENTAL_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCREMENTAL_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) SetINCREMENTAL_CONFIG(v string)`

SetINCREMENTAL_CONFIG sets INCREMENTAL_CONFIG field to given value.

### HasINCREMENTAL_CONFIG

`func (o *GetConnectionDetailsConnectionattributes) HasINCREMENTAL_CONFIG() bool`

HasINCREMENTAL_CONFIG returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *GetConnectionDetailsConnectionattributes) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *GetConnectionDetailsConnectionattributes) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *GetConnectionDetailsConnectionattributes) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *GetConnectionDetailsConnectionattributes) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionTimeoutConfig() RESTConnectionAttributesConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionTimeoutConfigOk() (*RESTConnectionAttributesConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *GetConnectionDetailsConnectionattributes) SetConnectionTimeoutConfig(v RESTConnectionAttributesConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *GetConnectionDetailsConnectionattributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *GetConnectionDetailsConnectionattributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *GetConnectionDetailsConnectionattributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *GetConnectionDetailsConnectionattributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *GetConnectionDetailsConnectionattributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetUpdateUserJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *GetConnectionDetailsConnectionattributes) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *GetConnectionDetailsConnectionattributes) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *GetConnectionDetailsConnectionattributes) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetTicketStatusJSON

`func (o *GetConnectionDetailsConnectionattributes) GetTicketStatusJSON() string`

GetTicketStatusJSON returns the TicketStatusJSON field if non-nil, zero value otherwise.

### GetTicketStatusJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetTicketStatusJSONOk() (*string, bool)`

GetTicketStatusJSONOk returns a tuple with the TicketStatusJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatusJSON

`func (o *GetConnectionDetailsConnectionattributes) SetTicketStatusJSON(v string)`

SetTicketStatusJSON sets TicketStatusJSON field to given value.

### HasTicketStatusJSON

`func (o *GetConnectionDetailsConnectionattributes) HasTicketStatusJSON() bool`

HasTicketStatusJSON returns a boolean if a field has been set.

### GetCreateTicketJSON

`func (o *GetConnectionDetailsConnectionattributes) GetCreateTicketJSON() string`

GetCreateTicketJSON returns the CreateTicketJSON field if non-nil, zero value otherwise.

### GetCreateTicketJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetCreateTicketJSONOk() (*string, bool)`

GetCreateTicketJSONOk returns a tuple with the CreateTicketJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTicketJSON

`func (o *GetConnectionDetailsConnectionattributes) SetCreateTicketJSON(v string)`

SetCreateTicketJSON sets CreateTicketJSON field to given value.

### HasCreateTicketJSON

`func (o *GetConnectionDetailsConnectionattributes) HasCreateTicketJSON() bool`

HasCreateTicketJSON returns a boolean if a field has been set.

### GetPasswdPolicyJSON

`func (o *GetConnectionDetailsConnectionattributes) GetPasswdPolicyJSON() string`

GetPasswdPolicyJSON returns the PasswdPolicyJSON field if non-nil, zero value otherwise.

### GetPasswdPolicyJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetPasswdPolicyJSONOk() (*string, bool)`

GetPasswdPolicyJSONOk returns a tuple with the PasswdPolicyJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswdPolicyJSON

`func (o *GetConnectionDetailsConnectionattributes) SetPasswdPolicyJSON(v string)`

SetPasswdPolicyJSON sets PasswdPolicyJSON field to given value.

### HasPasswdPolicyJSON

`func (o *GetConnectionDetailsConnectionattributes) HasPasswdPolicyJSON() bool`

HasPasswdPolicyJSON returns a boolean if a field has been set.

### GetAddFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) GetAddFFIDAccessJSON() string`

GetAddFFIDAccessJSON returns the AddFFIDAccessJSON field if non-nil, zero value otherwise.

### GetAddFFIDAccessJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetAddFFIDAccessJSONOk() (*string, bool)`

GetAddFFIDAccessJSONOk returns a tuple with the AddFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) SetAddFFIDAccessJSON(v string)`

SetAddFFIDAccessJSON sets AddFFIDAccessJSON field to given value.

### HasAddFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) HasAddFFIDAccessJSON() bool`

HasAddFFIDAccessJSON returns a boolean if a field has been set.

### GetRemoveFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveFFIDAccessJSON() string`

GetRemoveFFIDAccessJSON returns the RemoveFFIDAccessJSON field if non-nil, zero value otherwise.

### GetRemoveFFIDAccessJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveFFIDAccessJSONOk() (*string, bool)`

GetRemoveFFIDAccessJSONOk returns a tuple with the RemoveFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) SetRemoveFFIDAccessJSON(v string)`

SetRemoveFFIDAccessJSON sets RemoveFFIDAccessJSON field to given value.

### HasRemoveFFIDAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) HasRemoveFFIDAccessJSON() bool`

HasRemoveFFIDAccessJSON returns a boolean if a field has been set.

### GetSendOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) GetSendOtpJSON() string`

GetSendOtpJSON returns the SendOtpJSON field if non-nil, zero value otherwise.

### GetSendOtpJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetSendOtpJSONOk() (*string, bool)`

GetSendOtpJSONOk returns a tuple with the SendOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) SetSendOtpJSON(v string)`

SetSendOtpJSON sets SendOtpJSON field to given value.

### HasSendOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) HasSendOtpJSON() bool`

HasSendOtpJSON returns a boolean if a field has been set.

### GetValidateOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) GetValidateOtpJSON() string`

GetValidateOtpJSON returns the ValidateOtpJSON field if non-nil, zero value otherwise.

### GetValidateOtpJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetValidateOtpJSONOk() (*string, bool)`

GetValidateOtpJSONOk returns a tuple with the ValidateOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) SetValidateOtpJSON(v string)`

SetValidateOtpJSON sets ValidateOtpJSON field to given value.

### HasValidateOtpJSON

`func (o *GetConnectionDetailsConnectionattributes) HasValidateOtpJSON() bool`

HasValidateOtpJSON returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *GetConnectionDetailsConnectionattributes) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *GetConnectionDetailsConnectionattributes) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *GetConnectionDetailsConnectionattributes) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *GetConnectionDetailsConnectionattributes) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *GetConnectionDetailsConnectionattributes) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetImportAccountEntJSON

`func (o *GetConnectionDetailsConnectionattributes) GetImportAccountEntJSON() string`

GetImportAccountEntJSON returns the ImportAccountEntJSON field if non-nil, zero value otherwise.

### GetImportAccountEntJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetImportAccountEntJSONOk() (*string, bool)`

GetImportAccountEntJSONOk returns a tuple with the ImportAccountEntJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAccountEntJSON

`func (o *GetConnectionDetailsConnectionattributes) SetImportAccountEntJSON(v string)`

SetImportAccountEntJSON sets ImportAccountEntJSON field to given value.

### HasImportAccountEntJSON

`func (o *GetConnectionDetailsConnectionattributes) HasImportAccountEntJSON() bool`

HasImportAccountEntJSON returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionJSON() map[string]interface{}`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *GetConnectionDetailsConnectionattributes) GetConnectionJSONOk() (*map[string]interface{}, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *GetConnectionDetailsConnectionattributes) SetConnectionJSON(v map[string]interface{})`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *GetConnectionDetailsConnectionattributes) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


