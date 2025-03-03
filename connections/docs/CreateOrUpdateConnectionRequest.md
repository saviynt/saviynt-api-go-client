# CreateOrUpdateConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | Pointer to **string** |  | [optional] 
**Connectiontype** | **string** |  | 
**Connectionkey** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GroupSearchBaseDN** | Pointer to **string** |  | [optional] 
**Saveconnection** | Pointer to **string** |  | [optional] 
**Systemname** | Pointer to **string** |  | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** |  | [optional] 
**BASE** | Pointer to **string** |  | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** |  | [optional] 
**CREATEORGJSON** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENFORCE_TREE_DELETION** | Pointer to **string** |  | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**IMPORTJSON** | Pointer to **string** |  | [optional] 
**JCO_ASHOST** | Pointer to **string** |  | [optional] 
**JCO_CLIENT** | Pointer to **string** |  | [optional] 
**JCO_SYSNR** | Pointer to **string** |  | [optional] 
**JCO_USER** | Pointer to **string** |  | [optional] 
**LDAP_OR_AD** | Pointer to **string** |  | [optional] 
**OBJECTFILTER** | Pointer to **string** |  | [optional] 
**ORGANIZATION_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ORGIMPORTJSON** | Pointer to **string** |  | [optional] 
**ORG_BASE** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** |  | [optional] 
**READ_OPERATIONAL_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**REMOVEACCOUNTACTION** | Pointer to **string** |  | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** |  | [optional] 
**REUSEINACTIVEACCOUNT** | Pointer to **string** |  | [optional] 
**SEARCHFILTER** | Pointer to **string** |  | [optional] 
**SETDEFAULTPAGESIZE** | Pointer to **string** |  | [optional] 
**SETRANDOMPASSWORD** | Pointer to **string** |  | [optional] 
**STATUSKEYJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**SUPPORTEMPTYSTRING** | Pointer to **string** |  | [optional] 
**UNLOCKACCOUNTJSON** | Pointer to **string** |  | [optional] 
**UPDATEORGJSON** | Pointer to **string** |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrUpdateConnectionRequest

`func NewCreateOrUpdateConnectionRequest(connectiontype string, ) *CreateOrUpdateConnectionRequest`

NewCreateOrUpdateConnectionRequest instantiates a new CreateOrUpdateConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConnectionRequestWithDefaults

`func NewCreateOrUpdateConnectionRequestWithDefaults() *CreateOrUpdateConnectionRequest`

NewCreateOrUpdateConnectionRequestWithDefaults instantiates a new CreateOrUpdateConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *CreateOrUpdateConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CreateOrUpdateConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CreateOrUpdateConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *CreateOrUpdateConnectionRequest) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectiontype

`func (o *CreateOrUpdateConnectionRequest) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *CreateOrUpdateConnectionRequest) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *CreateOrUpdateConnectionRequest) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.


### GetConnectionkey

`func (o *CreateOrUpdateConnectionRequest) GetConnectionkey() string`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *CreateOrUpdateConnectionRequest) GetConnectionkeyOk() (*string, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *CreateOrUpdateConnectionRequest) SetConnectionkey(v string)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *CreateOrUpdateConnectionRequest) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateConnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateConnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateConnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateConnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *CreateOrUpdateConnectionRequest) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *CreateOrUpdateConnectionRequest) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *CreateOrUpdateConnectionRequest) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *CreateOrUpdateConnectionRequest) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetSaveconnection

`func (o *CreateOrUpdateConnectionRequest) GetSaveconnection() string`

GetSaveconnection returns the Saveconnection field if non-nil, zero value otherwise.

### GetSaveconnectionOk

`func (o *CreateOrUpdateConnectionRequest) GetSaveconnectionOk() (*string, bool)`

GetSaveconnectionOk returns a tuple with the Saveconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveconnection

`func (o *CreateOrUpdateConnectionRequest) SetSaveconnection(v string)`

SetSaveconnection sets Saveconnection field to given value.

### HasSaveconnection

`func (o *CreateOrUpdateConnectionRequest) HasSaveconnection() bool`

HasSaveconnection returns a boolean if a field has been set.

### GetSystemname

`func (o *CreateOrUpdateConnectionRequest) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *CreateOrUpdateConnectionRequest) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *CreateOrUpdateConnectionRequest) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.

### HasSystemname

`func (o *CreateOrUpdateConnectionRequest) HasSystemname() bool`

HasSystemname returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *CreateOrUpdateConnectionRequest) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *CreateOrUpdateConnectionRequest) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *CreateOrUpdateConnectionRequest) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *CreateOrUpdateConnectionRequest) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *CreateOrUpdateConnectionRequest) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetBASE

`func (o *CreateOrUpdateConnectionRequest) GetBASE() string`

GetBASE returns the BASE field if non-nil, zero value otherwise.

### GetBASEOk

`func (o *CreateOrUpdateConnectionRequest) GetBASEOk() (*string, bool)`

GetBASEOk returns a tuple with the BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE

`func (o *CreateOrUpdateConnectionRequest) SetBASE(v string)`

SetBASE sets BASE field to given value.

### HasBASE

`func (o *CreateOrUpdateConnectionRequest) HasBASE() bool`

HasBASE returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *CreateOrUpdateConnectionRequest) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *CreateOrUpdateConnectionRequest) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *CreateOrUpdateConnectionRequest) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *CreateOrUpdateConnectionRequest) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetCREATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) GetCREATEORGJSON() string`

GetCREATEORGJSON returns the CREATEORGJSON field if non-nil, zero value otherwise.

### GetCREATEORGJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetCREATEORGJSONOk() (*string, bool)`

GetCREATEORGJSONOk returns a tuple with the CREATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) SetCREATEORGJSON(v string)`

SetCREATEORGJSON sets CREATEORGJSON field to given value.

### HasCREATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) HasCREATEORGJSON() bool`

HasCREATEORGJSON returns a boolean if a field has been set.

### GetConfigJSON

`func (o *CreateOrUpdateConnectionRequest) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *CreateOrUpdateConnectionRequest) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *CreateOrUpdateConnectionRequest) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENFORCE_TREE_DELETION

`func (o *CreateOrUpdateConnectionRequest) GetENFORCE_TREE_DELETION() string`

GetENFORCE_TREE_DELETION returns the ENFORCE_TREE_DELETION field if non-nil, zero value otherwise.

### GetENFORCE_TREE_DELETIONOk

`func (o *CreateOrUpdateConnectionRequest) GetENFORCE_TREE_DELETIONOk() (*string, bool)`

GetENFORCE_TREE_DELETIONOk returns a tuple with the ENFORCE_TREE_DELETION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCE_TREE_DELETION

`func (o *CreateOrUpdateConnectionRequest) SetENFORCE_TREE_DELETION(v string)`

SetENFORCE_TREE_DELETION sets ENFORCE_TREE_DELETION field to given value.

### HasENFORCE_TREE_DELETION

`func (o *CreateOrUpdateConnectionRequest) HasENFORCE_TREE_DELETION() bool`

HasENFORCE_TREE_DELETION returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *CreateOrUpdateConnectionRequest) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) GetIMPORTJSON() string`

GetIMPORTJSON returns the IMPORTJSON field if non-nil, zero value otherwise.

### GetIMPORTJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetIMPORTJSONOk() (*string, bool)`

GetIMPORTJSONOk returns a tuple with the IMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) SetIMPORTJSON(v string)`

SetIMPORTJSON sets IMPORTJSON field to given value.

### HasIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) HasIMPORTJSON() bool`

HasIMPORTJSON returns a boolean if a field has been set.

### GetJCO_ASHOST

`func (o *CreateOrUpdateConnectionRequest) GetJCO_ASHOST() string`

GetJCO_ASHOST returns the JCO_ASHOST field if non-nil, zero value otherwise.

### GetJCO_ASHOSTOk

`func (o *CreateOrUpdateConnectionRequest) GetJCO_ASHOSTOk() (*string, bool)`

GetJCO_ASHOSTOk returns a tuple with the JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_ASHOST

`func (o *CreateOrUpdateConnectionRequest) SetJCO_ASHOST(v string)`

SetJCO_ASHOST sets JCO_ASHOST field to given value.

### HasJCO_ASHOST

`func (o *CreateOrUpdateConnectionRequest) HasJCO_ASHOST() bool`

HasJCO_ASHOST returns a boolean if a field has been set.

### GetJCO_CLIENT

`func (o *CreateOrUpdateConnectionRequest) GetJCO_CLIENT() string`

GetJCO_CLIENT returns the JCO_CLIENT field if non-nil, zero value otherwise.

### GetJCO_CLIENTOk

`func (o *CreateOrUpdateConnectionRequest) GetJCO_CLIENTOk() (*string, bool)`

GetJCO_CLIENTOk returns a tuple with the JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_CLIENT

`func (o *CreateOrUpdateConnectionRequest) SetJCO_CLIENT(v string)`

SetJCO_CLIENT sets JCO_CLIENT field to given value.

### HasJCO_CLIENT

`func (o *CreateOrUpdateConnectionRequest) HasJCO_CLIENT() bool`

HasJCO_CLIENT returns a boolean if a field has been set.

### GetJCO_SYSNR

`func (o *CreateOrUpdateConnectionRequest) GetJCO_SYSNR() string`

GetJCO_SYSNR returns the JCO_SYSNR field if non-nil, zero value otherwise.

### GetJCO_SYSNROk

`func (o *CreateOrUpdateConnectionRequest) GetJCO_SYSNROk() (*string, bool)`

GetJCO_SYSNROk returns a tuple with the JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SYSNR

`func (o *CreateOrUpdateConnectionRequest) SetJCO_SYSNR(v string)`

SetJCO_SYSNR sets JCO_SYSNR field to given value.

### HasJCO_SYSNR

`func (o *CreateOrUpdateConnectionRequest) HasJCO_SYSNR() bool`

HasJCO_SYSNR returns a boolean if a field has been set.

### GetJCO_USER

`func (o *CreateOrUpdateConnectionRequest) GetJCO_USER() string`

GetJCO_USER returns the JCO_USER field if non-nil, zero value otherwise.

### GetJCO_USEROk

`func (o *CreateOrUpdateConnectionRequest) GetJCO_USEROk() (*string, bool)`

GetJCO_USEROk returns a tuple with the JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_USER

`func (o *CreateOrUpdateConnectionRequest) SetJCO_USER(v string)`

SetJCO_USER sets JCO_USER field to given value.

### HasJCO_USER

`func (o *CreateOrUpdateConnectionRequest) HasJCO_USER() bool`

HasJCO_USER returns a boolean if a field has been set.

### GetLDAP_OR_AD

`func (o *CreateOrUpdateConnectionRequest) GetLDAP_OR_AD() string`

GetLDAP_OR_AD returns the LDAP_OR_AD field if non-nil, zero value otherwise.

### GetLDAP_OR_ADOk

`func (o *CreateOrUpdateConnectionRequest) GetLDAP_OR_ADOk() (*string, bool)`

GetLDAP_OR_ADOk returns a tuple with the LDAP_OR_AD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_OR_AD

`func (o *CreateOrUpdateConnectionRequest) SetLDAP_OR_AD(v string)`

SetLDAP_OR_AD sets LDAP_OR_AD field to given value.

### HasLDAP_OR_AD

`func (o *CreateOrUpdateConnectionRequest) HasLDAP_OR_AD() bool`

HasLDAP_OR_AD returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *CreateOrUpdateConnectionRequest) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *CreateOrUpdateConnectionRequest) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *CreateOrUpdateConnectionRequest) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *CreateOrUpdateConnectionRequest) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetORGANIZATION_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) GetORGANIZATION_ATTRIBUTE() string`

GetORGANIZATION_ATTRIBUTE returns the ORGANIZATION_ATTRIBUTE field if non-nil, zero value otherwise.

### GetORGANIZATION_ATTRIBUTEOk

`func (o *CreateOrUpdateConnectionRequest) GetORGANIZATION_ATTRIBUTEOk() (*string, bool)`

GetORGANIZATION_ATTRIBUTEOk returns a tuple with the ORGANIZATION_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) SetORGANIZATION_ATTRIBUTE(v string)`

SetORGANIZATION_ATTRIBUTE sets ORGANIZATION_ATTRIBUTE field to given value.

### HasORGANIZATION_ATTRIBUTE

`func (o *CreateOrUpdateConnectionRequest) HasORGANIZATION_ATTRIBUTE() bool`

HasORGANIZATION_ATTRIBUTE returns a boolean if a field has been set.

### GetORGIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) GetORGIMPORTJSON() string`

GetORGIMPORTJSON returns the ORGIMPORTJSON field if non-nil, zero value otherwise.

### GetORGIMPORTJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetORGIMPORTJSONOk() (*string, bool)`

GetORGIMPORTJSONOk returns a tuple with the ORGIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) SetORGIMPORTJSON(v string)`

SetORGIMPORTJSON sets ORGIMPORTJSON field to given value.

### HasORGIMPORTJSON

`func (o *CreateOrUpdateConnectionRequest) HasORGIMPORTJSON() bool`

HasORGIMPORTJSON returns a boolean if a field has been set.

### GetORG_BASE

`func (o *CreateOrUpdateConnectionRequest) GetORG_BASE() string`

GetORG_BASE returns the ORG_BASE field if non-nil, zero value otherwise.

### GetORG_BASEOk

`func (o *CreateOrUpdateConnectionRequest) GetORG_BASEOk() (*string, bool)`

GetORG_BASEOk returns a tuple with the ORG_BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORG_BASE

`func (o *CreateOrUpdateConnectionRequest) SetORG_BASE(v string)`

SetORG_BASE sets ORG_BASE field to given value.

### HasORG_BASE

`func (o *CreateOrUpdateConnectionRequest) HasORG_BASE() bool`

HasORG_BASE returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *CreateOrUpdateConnectionRequest) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *CreateOrUpdateConnectionRequest) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *CreateOrUpdateConnectionRequest) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *CreateOrUpdateConnectionRequest) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *CreateOrUpdateConnectionRequest) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *CreateOrUpdateConnectionRequest) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *CreateOrUpdateConnectionRequest) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *CreateOrUpdateConnectionRequest) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPASSWORD

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *CreateOrUpdateConnectionRequest) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateConnectionRequest) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateConnectionRequest) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetREAD_OPERATIONAL_ATTRIBUTES

`func (o *CreateOrUpdateConnectionRequest) GetREAD_OPERATIONAL_ATTRIBUTES() string`

GetREAD_OPERATIONAL_ATTRIBUTES returns the READ_OPERATIONAL_ATTRIBUTES field if non-nil, zero value otherwise.

### GetREAD_OPERATIONAL_ATTRIBUTESOk

`func (o *CreateOrUpdateConnectionRequest) GetREAD_OPERATIONAL_ATTRIBUTESOk() (*string, bool)`

GetREAD_OPERATIONAL_ATTRIBUTESOk returns a tuple with the READ_OPERATIONAL_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_OPERATIONAL_ATTRIBUTES

`func (o *CreateOrUpdateConnectionRequest) SetREAD_OPERATIONAL_ATTRIBUTES(v string)`

SetREAD_OPERATIONAL_ATTRIBUTES sets READ_OPERATIONAL_ATTRIBUTES field to given value.

### HasREAD_OPERATIONAL_ATTRIBUTES

`func (o *CreateOrUpdateConnectionRequest) HasREAD_OPERATIONAL_ATTRIBUTES() bool`

HasREAD_OPERATIONAL_ATTRIBUTES returns a boolean if a field has been set.

### GetREMOVEACCOUNTACTION

`func (o *CreateOrUpdateConnectionRequest) GetREMOVEACCOUNTACTION() string`

GetREMOVEACCOUNTACTION returns the REMOVEACCOUNTACTION field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTACTIONOk

`func (o *CreateOrUpdateConnectionRequest) GetREMOVEACCOUNTACTIONOk() (*string, bool)`

GetREMOVEACCOUNTACTIONOk returns a tuple with the REMOVEACCOUNTACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTACTION

`func (o *CreateOrUpdateConnectionRequest) SetREMOVEACCOUNTACTION(v string)`

SetREMOVEACCOUNTACTION sets REMOVEACCOUNTACTION field to given value.

### HasREMOVEACCOUNTACTION

`func (o *CreateOrUpdateConnectionRequest) HasREMOVEACCOUNTACTION() bool`

HasREMOVEACCOUNTACTION returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateConnectionRequest) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateConnectionRequest) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateConnectionRequest) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetREUSEINACTIVEACCOUNT

`func (o *CreateOrUpdateConnectionRequest) GetREUSEINACTIVEACCOUNT() string`

GetREUSEINACTIVEACCOUNT returns the REUSEINACTIVEACCOUNT field if non-nil, zero value otherwise.

### GetREUSEINACTIVEACCOUNTOk

`func (o *CreateOrUpdateConnectionRequest) GetREUSEINACTIVEACCOUNTOk() (*string, bool)`

GetREUSEINACTIVEACCOUNTOk returns a tuple with the REUSEINACTIVEACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEINACTIVEACCOUNT

`func (o *CreateOrUpdateConnectionRequest) SetREUSEINACTIVEACCOUNT(v string)`

SetREUSEINACTIVEACCOUNT sets REUSEINACTIVEACCOUNT field to given value.

### HasREUSEINACTIVEACCOUNT

`func (o *CreateOrUpdateConnectionRequest) HasREUSEINACTIVEACCOUNT() bool`

HasREUSEINACTIVEACCOUNT returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *CreateOrUpdateConnectionRequest) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *CreateOrUpdateConnectionRequest) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *CreateOrUpdateConnectionRequest) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *CreateOrUpdateConnectionRequest) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetSETDEFAULTPAGESIZE

`func (o *CreateOrUpdateConnectionRequest) GetSETDEFAULTPAGESIZE() string`

GetSETDEFAULTPAGESIZE returns the SETDEFAULTPAGESIZE field if non-nil, zero value otherwise.

### GetSETDEFAULTPAGESIZEOk

`func (o *CreateOrUpdateConnectionRequest) GetSETDEFAULTPAGESIZEOk() (*string, bool)`

GetSETDEFAULTPAGESIZEOk returns a tuple with the SETDEFAULTPAGESIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETDEFAULTPAGESIZE

`func (o *CreateOrUpdateConnectionRequest) SetSETDEFAULTPAGESIZE(v string)`

SetSETDEFAULTPAGESIZE sets SETDEFAULTPAGESIZE field to given value.

### HasSETDEFAULTPAGESIZE

`func (o *CreateOrUpdateConnectionRequest) HasSETDEFAULTPAGESIZE() bool`

HasSETDEFAULTPAGESIZE returns a boolean if a field has been set.

### GetSETRANDOMPASSWORD

`func (o *CreateOrUpdateConnectionRequest) GetSETRANDOMPASSWORD() string`

GetSETRANDOMPASSWORD returns the SETRANDOMPASSWORD field if non-nil, zero value otherwise.

### GetSETRANDOMPASSWORDOk

`func (o *CreateOrUpdateConnectionRequest) GetSETRANDOMPASSWORDOk() (*string, bool)`

GetSETRANDOMPASSWORDOk returns a tuple with the SETRANDOMPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETRANDOMPASSWORD

`func (o *CreateOrUpdateConnectionRequest) SetSETRANDOMPASSWORD(v string)`

SetSETRANDOMPASSWORD sets SETRANDOMPASSWORD field to given value.

### HasSETRANDOMPASSWORD

`func (o *CreateOrUpdateConnectionRequest) HasSETRANDOMPASSWORD() bool`

HasSETRANDOMPASSWORD returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *CreateOrUpdateConnectionRequest) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *CreateOrUpdateConnectionRequest) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *CreateOrUpdateConnectionRequest) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateConnectionRequest) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *CreateOrUpdateConnectionRequest) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateConnectionRequest) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateConnectionRequest) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetSUPPORTEMPTYSTRING

`func (o *CreateOrUpdateConnectionRequest) GetSUPPORTEMPTYSTRING() string`

GetSUPPORTEMPTYSTRING returns the SUPPORTEMPTYSTRING field if non-nil, zero value otherwise.

### GetSUPPORTEMPTYSTRINGOk

`func (o *CreateOrUpdateConnectionRequest) GetSUPPORTEMPTYSTRINGOk() (*string, bool)`

GetSUPPORTEMPTYSTRINGOk returns a tuple with the SUPPORTEMPTYSTRING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUPPORTEMPTYSTRING

`func (o *CreateOrUpdateConnectionRequest) SetSUPPORTEMPTYSTRING(v string)`

SetSUPPORTEMPTYSTRING sets SUPPORTEMPTYSTRING field to given value.

### HasSUPPORTEMPTYSTRING

`func (o *CreateOrUpdateConnectionRequest) HasSUPPORTEMPTYSTRING() bool`

HasSUPPORTEMPTYSTRING returns a boolean if a field has been set.

### GetUNLOCKACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) GetUNLOCKACCOUNTJSON() string`

GetUNLOCKACCOUNTJSON returns the UNLOCKACCOUNTJSON field if non-nil, zero value otherwise.

### GetUNLOCKACCOUNTJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetUNLOCKACCOUNTJSONOk() (*string, bool)`

GetUNLOCKACCOUNTJSONOk returns a tuple with the UNLOCKACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCKACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) SetUNLOCKACCOUNTJSON(v string)`

SetUNLOCKACCOUNTJSON sets UNLOCKACCOUNTJSON field to given value.

### HasUNLOCKACCOUNTJSON

`func (o *CreateOrUpdateConnectionRequest) HasUNLOCKACCOUNTJSON() bool`

HasUNLOCKACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) GetUPDATEORGJSON() string`

GetUPDATEORGJSON returns the UPDATEORGJSON field if non-nil, zero value otherwise.

### GetUPDATEORGJSONOk

`func (o *CreateOrUpdateConnectionRequest) GetUPDATEORGJSONOk() (*string, bool)`

GetUPDATEORGJSONOk returns a tuple with the UPDATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) SetUPDATEORGJSON(v string)`

SetUPDATEORGJSON sets UPDATEORGJSON field to given value.

### HasUPDATEORGJSON

`func (o *CreateOrUpdateConnectionRequest) HasUPDATEORGJSON() bool`

HasUPDATEORGJSON returns a boolean if a field has been set.

### GetURL

`func (o *CreateOrUpdateConnectionRequest) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *CreateOrUpdateConnectionRequest) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *CreateOrUpdateConnectionRequest) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *CreateOrUpdateConnectionRequest) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetUSERNAME

`func (o *CreateOrUpdateConnectionRequest) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *CreateOrUpdateConnectionRequest) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *CreateOrUpdateConnectionRequest) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *CreateOrUpdateConnectionRequest) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


