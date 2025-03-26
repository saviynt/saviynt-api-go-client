# TestConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | Name of the connection | 
**Connectiontype** | **string** | Connection type (e.g., &#39;AD&#39; for Active Directory). | 
**Description** | Pointer to **string** | Description for the connection. | [optional] 
**Defaultsavroles** | Pointer to **string** | Default SAV roles for managing the connection. | [optional] 
**EmailTemplate** | Pointer to **string** | Email template for notifications. | [optional] 
**SslCertificate** | Pointer to **string** | SSL certificates to secure the connection. | [optional] 
**VaultConnection** | Pointer to **string** | Specifies the type of vault connection being used (e.g., Hashicorp, AWS Secrets Manager). | [optional] 
**VaultConfiguration** | Pointer to **string** | JSON string specifying vault configuration | [optional] 
**Saveinvault** | Pointer to **string** | Flag indicating whether the encrypted attribute should be saved in the configured vault. | [optional] 
**URL** | **string** | Domain URL list (comma Separated) | 
**Domain** | Pointer to **string** | Domain for ADSI authentication | [optional] 
**USERNAME** | **string** | The username to use for SOAP authentication. | 
**PASSWORD** | **string** | The password of the username used for authentication. | 
**CONNECTION_URL** | **string** | Server URL where ADSI agent is deployed on IIS. You could have http/https URL with actual port information. Used only to save and test connection and retrieve forest information | 
**PROVISIONING_URL** | Pointer to **string** | Server URL with specific provisioning URL where ADSI agent is deployed on IIS. You could have http/https URL with actual port information. | [optional] 
**FORESTLIST** | **string** | Forest List (Comma Separated) which we need to connect using ADSI connector. | 
**DEFAULT_USER_ROLE** | Pointer to **string** | Default SAV Role to be assigned to all the new users that gets imported - only valid for User Import | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Update existing User | [optional] 
**FOREST_DETAILS** | Pointer to **string** |  | [optional] 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** |  | [optional] 
**CreateUpdateMappings** | Pointer to **string** |  | [optional] 
**IMPORTDATACOOKIES** | Pointer to **string** |  | [optional] 
**PASSWDPOLICYJSON** | Pointer to **string** |  | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**SEARCHFILTER** | Pointer to **string** | Search Filter can be used to specify the BaseDN of the directory from where the data needs to be imported. You can have multiple BaseDNs here separated by ###. | [optional] 
**OBJECTFILTER** | Pointer to **string** | Object Filter is used to filter the objects that will be returned.This filter will be same for all domains. | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** | Controls the AD Attribute to Saviynt Account Mapping (AD attributes must be in lower case) | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Specify this parameter if you want to read and import the status of an account and an entitlement. | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | This field provides details as to which endpoint in Saviynt should the AD accounts be associated to after the accounts have been imported | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** | Controls the AD Attribute to Saviynt Account Mapping (AD attributes must be in lower case) | [optional] 
**GroupSearchBaseDN** | Pointer to **string** | The DN from which the search for all the groups begins. You can have multiple BaseDNs here separated by ###. | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** |  | [optional] 
**STATUSKEYJSON** | Pointer to **string** | JSON to specify Users status | [optional] 
**GroupImportMapping** | Pointer to **string** | Mapping used during accessimport to specify which attribute of a group maps to which attribute on Saviynt | [optional] 
**ImportNestedMembership** | Pointer to **string** | Brings all indirect/nested membership of an account/group during Account/Access import. By Default, it is FALSE. Once true, it will return value in an array with key name \&quot;nestedEntitlementList\&quot;. Recommended to map this value in a column with datatype LONGTEXT. | [optional] 
**PAGE_SIZE** | Pointer to **string** | Specify the number of objects to return in a page for each import request from Workday | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** | Rule to generate account name. | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | JSON Similar to Create Account to mention the changes that needs to be made at the user level after updating the account | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON Similar to Create Account to mention the changes that needs to be made at the user level after enabling the account | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for disabling an active account. | [optional] 
**REMOVEACCOUNTJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for deleting/suspending an account. | [optional] 
**ADDACCESSJSON** | Pointer to **string** | JSON to ADD Access (cross domain/forest group membership) to an account. | [optional] 
**REMOVEACCESSJSON** | Pointer to **string** | JSON to REMOVE Access (cross domain/forest group membership) to an account. | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** | JSON to Reset and Change Password. | [optional] 
**MOVEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**CREATEGROUPJSON** | Pointer to **string** | JSON to Create Group in a multi-domain/forest Setup. | [optional] 
**UPDATEGROUPJSON** | Pointer to **string** | JSON to Update Group in a multi-domain/forest Setup. | [optional] 
**REMOVEGROUPJSON** | Pointer to **string** | JSON to Update Group in a multi-domain/forest Setup. | [optional] 
**ADDACCESSENTITLEMENTJSON** | Pointer to **string** | JSON to Add Group to a Group. | [optional] 
**CUSTOMCONFIGJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCESSENTITLEMENTJSON** | Pointer to **string** | JSON to Remove Group from a Group. | [optional] 
**CREATESERVICEACCOUNTJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Create the New Service Account. | [optional] 
**UPDATESERVICEACCOUNTJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Update existing Service Account. | [optional] 
**REMOVESERVICEACCOUNTJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed to delete an existing service account. | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**MESSAGESERVER** | Pointer to **string** | set it to TRUE if the Message Server is going to be used for connecting to SAP | [optional] 
**JCO_ASHOST** | Pointer to **string** | HostName for connection for Import | [optional] 
**JCO_SYSNR** | Pointer to **string** | System Number for the SAP Instance for Import | [optional] 
**JCO_CLIENT** | Pointer to **string** | Client Number for the SAP Instance for Import | [optional] 
**JCO_USER** | Pointer to **string** | Username to connect to SAP using JCO for Import | [optional] 
**JCO_LANG** | Pointer to **string** | LANGUAGE for the SAP Instance for Import | [optional] 
**JCOR3NAME** | Pointer to **string** | Message Server SystemID for connection for Import | [optional] 
**JCO_MSHOST** | Pointer to **string** | Message Server Hostname for connection for Import | [optional] 
**JCO_MSSERV** | Pointer to **string** | Message Server Port Number for connection for Import | [optional] 
**JCO_GROUP** | Pointer to **string** | Message Server logon Group connection for Import | [optional] 
**SNC** | Pointer to **string** | TRUE/FALSE, set it to TRUE if Secure Network Connection (SNC) is setup between ABAP and ASHOST | [optional] 
**JCO_SNC_MODE** | Pointer to **string** | Enable/Disable Secure Network Connection (SNC) Mode. 0 - Disabled, 1 - Enabled | [optional] 
**JCO_SNC_PARTNERNAME** | Pointer to **string** | String used to generate secured certificate in SAP server to be used by Saviynt | [optional] 
**JCO_SNC_MYNAME** | Pointer to **string** | String used to generate the secured certificate on the server where Saviynt is deployed | [optional] 
**JCO_SNC_LIBRARY** | Pointer to **string** | Location of SNC library on the Saviynt server | [optional] 
**JCO_SNC_QOP** | Pointer to **string** | The Quality of Protection Level | [optional] 
**TABLES** | Pointer to **string** | Tables to be Imported | [optional] 
**SYSTEMNAME** | Pointer to **string** |  | [optional] 
**TERMINATEDUSERGROUP** | Pointer to **string** | User group for all terminated users | [optional] 
**TERMINATED_USER_ROLE_ACTION** | Pointer to **string** | Action to take for the user roles when the user gets terminated or roles are removed | [optional] 
**PROV_JCO_ASHOST** | Pointer to **string** | HostName for connection for Provisioning | [optional] 
**PROV_JCO_SYSNR** | Pointer to **string** | System Number for the SAP Instance for Provisioning | [optional] 
**PROV_JCO_CLIENT** | Pointer to **string** | Client Number for the SAP Instance for Provisioning | [optional] 
**PROV_JCO_USER** | Pointer to **string** | Username to connect to to SAP using JCO for Provisioning | [optional] 
**PROV_PASSWORD** | Pointer to **string** | Password for connection for Provisioning | [optional] 
**PROV_JCO_LANG** | Pointer to **string** | LANGUAGE for the SAP Instance for Provisioning | [optional] 
**PROVJCOR3NAME** | Pointer to **string** | Message Server SystemID for connection for Provisioning | [optional] 
**PROV_JCO_MSHOST** | Pointer to **string** | Message Server Hostname for connection for Provisioning | [optional] 
**PROV_JCO_MSSERV** | Pointer to **string** | Message Server Port Number for connection for Provisioning | [optional] 
**PROV_JCO_GROUP** | Pointer to **string** | Message Server logon Group connection for Provisioning | [optional] 
**PROV_CUA_ENABLED** | Pointer to **string** | Is CUA Enabled | [optional] 
**PROV_CUA_SNC** | Pointer to **string** | Property for PROV_CUA_SNC | [optional] 
**RESET_PWD_FOR_NEWACCOUNT** | Pointer to **string** | For CUA enabled if the user already exists should the Password be reset for accounts for other systems | [optional] 
**ENFORCEPASSWORDCHANGE** | Pointer to **string** | Set it to FALSE if while changing the password the productive password should be set | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Specify the Min length for the random password | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Specify the Max length for the random password | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Specify the Number of Upper case alphabets required for the random password | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Specify the Number of special chars required for the random password | [optional] 
**HANAREFTABLEJSON** | Pointer to **string** | JSON to specify the mapping for Reference table | [optional] 
**USERIMPORTJSON** | Pointer to **string** | Property for USERIMPORTJSON | [optional] 
**SETCUASYSTEM** | Pointer to **string** | set it to FALSE if using an older CUA System to not support setting sub-systems | [optional] 
**FIREFIGHTERID_GRANT_ACCESS_JSON** | Pointer to **string** | JSON similar to Update Account to mention SNC and other related changes that needs to be done  | [optional] 
**FIREFIGHTERID_REVOKE_ACCESS_JSON** | Pointer to **string** | JSON similar to Update Account to reset SNC and other related changes that were done during GRANT | [optional] 
**EXTERNAL_SOD_EVAL_JSON** | Pointer to **string** | JSON to populate values for External SOD Evaluation | [optional] 
**EXTERNAL_SOD_EVAL_JSON_DETAIL** | Pointer to **string** | Property for EXTERNAL_SOD_EVAL_JSON_DETAIL | [optional] 
**LOGS_TABLE_FILTER** | Pointer to **string** | Property for LOGS_TABLE_FILTER | [optional] 
**SAPTABLE_FILTER_LANG** | Pointer to **string** | Property for SAPTABLE_FILTER_LANG | [optional] 
**ALTERNATE_OUTPUT_PARAMETER_ET_DATA** | Pointer to **string** | Based on the flag &#39;&#39;USE_ET_DATA_4_RETURN&#39;&#39; in SAP, set this value to make use of export parameter ET_DATA | [optional] 
**AUDIT_LOG_JSON** | Pointer to **string** | Property for AUDIT_LOG_JSON | [optional] 
**ECCORS4HANA** | Pointer to **string** | Property for ECC_OR_S4HANA | [optional] 
**DATA_IMPORT_FILTER** | Pointer to **string** | Property for DATA_IMPORT_FILTER | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**CLIENT_ID** | **string** |  | 
**CLIENT_SECRET** | **string** |  | 
**REFRESH_TOKEN** | Pointer to **string** |  | [optional] 
**REDIRECT_URI** | Pointer to **string** |  | [optional] 
**INSTANCE_URL** | Pointer to **string** |  | [optional] 
**OBJECT_TO_BE_IMPORTED** | Pointer to **string** |  | [optional] 
**FEATURE_LICENSE_JSON** | Pointer to **string** |  | [optional] 
**CUSTOM_CREATEACCOUNT_URL** | Pointer to **string** |  | [optional] 
**ACCOUNT_FILTER_QUERY** | Pointer to **string** |  | [optional] 
**ACCOUNT_FIELD_QUERY** | Pointer to **string** |  | [optional] 
**FIELD_MAPPING_JSON** | Pointer to **string** |  | [optional] 
**MODIFYACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ConnectionJSON** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportUserJSON** | Pointer to **string** | Property for ImportUserJSON | [optional] 
**ImportAccountEntJSON** | Pointer to **string** |  | [optional] 
**CreateAccountJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Create the New Account | [optional] 
**UpdateAccountJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Update existing Account | [optional] 
**EnableAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for enabling a disabled account | [optional] 
**DisableAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for disabling a enabled account | [optional] 
**AddAccessJSON** | Pointer to **string** | JSON to ADD Access to an account | [optional] 
**RemoveAccessJSON** | Pointer to **string** | JSON to REMOVE Access from an account | [optional] 
**UpdateUserJSON** | Pointer to **string** | Property for UpdateUserJSON | [optional] 
**ChangePassJSON** | Pointer to **string** |  | [optional] 
**RemoveAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for deleting/suspending an account | [optional] 
**TicketStatusJSON** | Pointer to **string** |  | [optional] 
**CreateTicketJSON** | Pointer to **string** |  | [optional] 
**PasswdPolicyJSON** | Pointer to **string** |  | [optional] 
**AddFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**RemoveFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**SendOtpJSON** | Pointer to **string** |  | [optional] 
**ValidateOtpJSON** | Pointer to **string** |  | [optional] 
**BASEURL** | **string** |  | 
**TENANT_ID** | **string** |  | 
**LOGIN_URL** | **string** |  | 
**USER_FILTER** | Pointer to **string** | Property for USER_FILTER | [optional] 
**USER_IMPORT_MAPPING** | Pointer to **string** | Specify the User import Mapping | [optional] 
**ACCOUNT_IMPORT_MAPPING** | Pointer to **string** | Specify the Account import Mapping | [optional] 
**ORGANIZATION_FILTER** | Pointer to **string** |  | [optional] 
**SCOPE** | Pointer to **string** | If present - 2.0 API will be used. Accepts space separated values. If blank 1.0 API is used | [optional] 
**USERS_LAST_IMPORT_TIME** | Pointer to **string** | Property for USERS_LAST_IMPORT_TIME | [optional] 
**ACCOUNTS_LAST_IMPORT_TIME** | Pointer to **string** | Property for ACCOUNTS_LAST_IMPORT_TIME | [optional] 
**ACCESS_LAST_IMPORT_TIME** | Pointer to **string** | Property for ACCESS_LAST_IMPORT_TIME | [optional] 
**BASE_URL** | Pointer to **string** | Base URL of the Workday tenant instance | [optional] 
**API_VERSION** | Pointer to **string** | Version of the SOAP API used for the connection. | [optional] 
**TENANT_NAME** | Pointer to **string** | The name of your tenant. | [optional] 
**REPORT_OWNER** | Pointer to **string** | Account name of the report owner required for constructing the default URLs of RaaS reports | [optional] 
**USE_OAUTH** | **string** | By default, the reports are accessed through OAuth authentication | 
**INCLUDE_REFERENCE_DESCRIPTORS** | Pointer to **string** | if it is TRUE, descriptor attribute will be avialable in response. | [optional] 
**USE_ENHANCED_ORGROLE** | Pointer to **string** | Set TRUE to utilize enhanced Organizational Role setup | [optional] 
**USEX509AUTHFORSOAP** | Pointer to **string** |  it is TRUE certificate based authentication will happen. | [optional] 
**X509KEY** | Pointer to **string** | Specify the privateKey to be used | [optional] 
**X509CERT** | Pointer to **string** | Specify the certificate to be used | [optional] 
**USER_IMPORT_PAYLOAD** | Pointer to **string** | Define the request payload for importing users | [optional] 
**ACCOUNT_IMPORT_PAYLOAD** | Pointer to **string** | Define the request payload for importing accounts | [optional] 
**ACCESS_IMPORT_LIST** | Pointer to **string** | Specify the access types to import in a comma separated list | [optional] 
**RAAS_MAPPING_JSON** | Pointer to **string** | Use this parameter if you want to override the default report mapping | [optional] 
**ACCESS_IMPORT_MAPPING** | Pointer to **string** | Define the access attribute mapping in addition to the default non-editable mapping for the Workday access objects | [optional] 
**ORGROLE_IMPORT_PAYLOAD** | Pointer to **string** | Specify the custom SOAP body for the organization role import | [optional] 
**CREATE_ACCOUNT_PAYLOAD** | Pointer to **string** | Define the request payload for creating an account in Workday | [optional] 
**UPDATE_ACCOUNT_PAYLOAD** | Pointer to **string** | Define the request payload for updating an account in Workday | [optional] 
**UPDATE_USER_PAYLOAD** | Pointer to **string** | Define the request payload for updating an user in Workday | [optional] 
**ASSIGN_ORGROLE_PAYLOAD** | Pointer to **string** | Define the request payload for assigning an organizational role to an account in Workday | [optional] 
**REMOVE_ORGROLE_PAYLOAD** | Pointer to **string** | Define the request payload for assigning an organizational role to an account in Workday | [optional] 
**STATUS_KEY_JSON** | Pointer to **string** | Specify the mapping of user status | [optional] 
**USERATTRIBUTEJSON** | Pointer to **string** | JSON that specifies which job related attributes are to be stored as User Attributes | [optional] 
**CUSTOM_CONFIG** | Pointer to **string** |  | [optional] 

## Methods

### NewTestConnectionRequest

`func NewTestConnectionRequest(connectionName string, connectiontype string, uRL string, uSERNAME string, pASSWORD string, cONNECTIONURL string, fORESTLIST string, cLIENTID string, cLIENTSECRET string, bASEURL string, tENANTID string, lOGINURL string, uSEOAUTH string, ) *TestConnectionRequest`

NewTestConnectionRequest instantiates a new TestConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestConnectionRequestWithDefaults

`func NewTestConnectionRequestWithDefaults() *TestConnectionRequest`

NewTestConnectionRequestWithDefaults instantiates a new TestConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *TestConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *TestConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *TestConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetConnectiontype

`func (o *TestConnectionRequest) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *TestConnectionRequest) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *TestConnectionRequest) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.


### GetDescription

`func (o *TestConnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestConnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestConnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestConnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *TestConnectionRequest) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *TestConnectionRequest) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *TestConnectionRequest) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *TestConnectionRequest) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetEmailTemplate

`func (o *TestConnectionRequest) GetEmailTemplate() string`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *TestConnectionRequest) GetEmailTemplateOk() (*string, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *TestConnectionRequest) SetEmailTemplate(v string)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *TestConnectionRequest) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetSslCertificate

`func (o *TestConnectionRequest) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *TestConnectionRequest) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *TestConnectionRequest) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *TestConnectionRequest) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetVaultConnection

`func (o *TestConnectionRequest) GetVaultConnection() string`

GetVaultConnection returns the VaultConnection field if non-nil, zero value otherwise.

### GetVaultConnectionOk

`func (o *TestConnectionRequest) GetVaultConnectionOk() (*string, bool)`

GetVaultConnectionOk returns a tuple with the VaultConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConnection

`func (o *TestConnectionRequest) SetVaultConnection(v string)`

SetVaultConnection sets VaultConnection field to given value.

### HasVaultConnection

`func (o *TestConnectionRequest) HasVaultConnection() bool`

HasVaultConnection returns a boolean if a field has been set.

### GetVaultConfiguration

`func (o *TestConnectionRequest) GetVaultConfiguration() string`

GetVaultConfiguration returns the VaultConfiguration field if non-nil, zero value otherwise.

### GetVaultConfigurationOk

`func (o *TestConnectionRequest) GetVaultConfigurationOk() (*string, bool)`

GetVaultConfigurationOk returns a tuple with the VaultConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConfiguration

`func (o *TestConnectionRequest) SetVaultConfiguration(v string)`

SetVaultConfiguration sets VaultConfiguration field to given value.

### HasVaultConfiguration

`func (o *TestConnectionRequest) HasVaultConfiguration() bool`

HasVaultConfiguration returns a boolean if a field has been set.

### GetSaveinvault

`func (o *TestConnectionRequest) GetSaveinvault() string`

GetSaveinvault returns the Saveinvault field if non-nil, zero value otherwise.

### GetSaveinvaultOk

`func (o *TestConnectionRequest) GetSaveinvaultOk() (*string, bool)`

GetSaveinvaultOk returns a tuple with the Saveinvault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveinvault

`func (o *TestConnectionRequest) SetSaveinvault(v string)`

SetSaveinvault sets Saveinvault field to given value.

### HasSaveinvault

`func (o *TestConnectionRequest) HasSaveinvault() bool`

HasSaveinvault returns a boolean if a field has been set.

### GetURL

`func (o *TestConnectionRequest) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *TestConnectionRequest) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *TestConnectionRequest) SetURL(v string)`

SetURL sets URL field to given value.


### GetDomain

`func (o *TestConnectionRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TestConnectionRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TestConnectionRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TestConnectionRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUSERNAME

`func (o *TestConnectionRequest) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *TestConnectionRequest) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *TestConnectionRequest) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetPASSWORD

`func (o *TestConnectionRequest) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *TestConnectionRequest) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *TestConnectionRequest) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.


### GetCONNECTION_URL

`func (o *TestConnectionRequest) GetCONNECTION_URL() string`

GetCONNECTION_URL returns the CONNECTION_URL field if non-nil, zero value otherwise.

### GetCONNECTION_URLOk

`func (o *TestConnectionRequest) GetCONNECTION_URLOk() (*string, bool)`

GetCONNECTION_URLOk returns a tuple with the CONNECTION_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTION_URL

`func (o *TestConnectionRequest) SetCONNECTION_URL(v string)`

SetCONNECTION_URL sets CONNECTION_URL field to given value.


### GetPROVISIONING_URL

`func (o *TestConnectionRequest) GetPROVISIONING_URL() string`

GetPROVISIONING_URL returns the PROVISIONING_URL field if non-nil, zero value otherwise.

### GetPROVISIONING_URLOk

`func (o *TestConnectionRequest) GetPROVISIONING_URLOk() (*string, bool)`

GetPROVISIONING_URLOk returns a tuple with the PROVISIONING_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISIONING_URL

`func (o *TestConnectionRequest) SetPROVISIONING_URL(v string)`

SetPROVISIONING_URL sets PROVISIONING_URL field to given value.

### HasPROVISIONING_URL

`func (o *TestConnectionRequest) HasPROVISIONING_URL() bool`

HasPROVISIONING_URL returns a boolean if a field has been set.

### GetFORESTLIST

`func (o *TestConnectionRequest) GetFORESTLIST() string`

GetFORESTLIST returns the FORESTLIST field if non-nil, zero value otherwise.

### GetFORESTLISTOk

`func (o *TestConnectionRequest) GetFORESTLISTOk() (*string, bool)`

GetFORESTLISTOk returns a tuple with the FORESTLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFORESTLIST

`func (o *TestConnectionRequest) SetFORESTLIST(v string)`

SetFORESTLIST sets FORESTLIST field to given value.


### GetDEFAULT_USER_ROLE

`func (o *TestConnectionRequest) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *TestConnectionRequest) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *TestConnectionRequest) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *TestConnectionRequest) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *TestConnectionRequest) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *TestConnectionRequest) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *TestConnectionRequest) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *TestConnectionRequest) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetFOREST_DETAILS

`func (o *TestConnectionRequest) GetFOREST_DETAILS() string`

GetFOREST_DETAILS returns the FOREST_DETAILS field if non-nil, zero value otherwise.

### GetFOREST_DETAILSOk

`func (o *TestConnectionRequest) GetFOREST_DETAILSOk() (*string, bool)`

GetFOREST_DETAILSOk returns a tuple with the FOREST_DETAILS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFOREST_DETAILS

`func (o *TestConnectionRequest) SetFOREST_DETAILS(v string)`

SetFOREST_DETAILS sets FOREST_DETAILS field to given value.

### HasFOREST_DETAILS

`func (o *TestConnectionRequest) HasFOREST_DETAILS() bool`

HasFOREST_DETAILS returns a boolean if a field has been set.

### GetENABLEGROUPMANAGEMENT

`func (o *TestConnectionRequest) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *TestConnectionRequest) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *TestConnectionRequest) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *TestConnectionRequest) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *TestConnectionRequest) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *TestConnectionRequest) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *TestConnectionRequest) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *TestConnectionRequest) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetIMPORTDATACOOKIES

`func (o *TestConnectionRequest) GetIMPORTDATACOOKIES() string`

GetIMPORTDATACOOKIES returns the IMPORTDATACOOKIES field if non-nil, zero value otherwise.

### GetIMPORTDATACOOKIESOk

`func (o *TestConnectionRequest) GetIMPORTDATACOOKIESOk() (*string, bool)`

GetIMPORTDATACOOKIESOk returns a tuple with the IMPORTDATACOOKIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTDATACOOKIES

`func (o *TestConnectionRequest) SetIMPORTDATACOOKIES(v string)`

SetIMPORTDATACOOKIES sets IMPORTDATACOOKIES field to given value.

### HasIMPORTDATACOOKIES

`func (o *TestConnectionRequest) HasIMPORTDATACOOKIES() bool`

HasIMPORTDATACOOKIES returns a boolean if a field has been set.

### GetPASSWDPOLICYJSON

`func (o *TestConnectionRequest) GetPASSWDPOLICYJSON() string`

GetPASSWDPOLICYJSON returns the PASSWDPOLICYJSON field if non-nil, zero value otherwise.

### GetPASSWDPOLICYJSONOk

`func (o *TestConnectionRequest) GetPASSWDPOLICYJSONOk() (*string, bool)`

GetPASSWDPOLICYJSONOk returns a tuple with the PASSWDPOLICYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWDPOLICYJSON

`func (o *TestConnectionRequest) SetPASSWDPOLICYJSON(v string)`

SetPASSWDPOLICYJSON sets PASSWDPOLICYJSON field to given value.

### HasPASSWDPOLICYJSON

`func (o *TestConnectionRequest) HasPASSWDPOLICYJSON() bool`

HasPASSWDPOLICYJSON returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *TestConnectionRequest) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *TestConnectionRequest) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *TestConnectionRequest) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *TestConnectionRequest) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *TestConnectionRequest) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *TestConnectionRequest) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *TestConnectionRequest) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *TestConnectionRequest) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *TestConnectionRequest) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *TestConnectionRequest) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *TestConnectionRequest) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *TestConnectionRequest) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *TestConnectionRequest) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *TestConnectionRequest) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *TestConnectionRequest) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *TestConnectionRequest) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *TestConnectionRequest) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *TestConnectionRequest) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *TestConnectionRequest) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *TestConnectionRequest) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *TestConnectionRequest) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *TestConnectionRequest) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *TestConnectionRequest) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *TestConnectionRequest) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *TestConnectionRequest) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *TestConnectionRequest) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *TestConnectionRequest) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *TestConnectionRequest) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *TestConnectionRequest) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *TestConnectionRequest) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *TestConnectionRequest) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *TestConnectionRequest) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *TestConnectionRequest) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *TestConnectionRequest) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *TestConnectionRequest) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *TestConnectionRequest) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *TestConnectionRequest) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *TestConnectionRequest) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *TestConnectionRequest) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *TestConnectionRequest) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *TestConnectionRequest) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *TestConnectionRequest) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *TestConnectionRequest) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *TestConnectionRequest) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetImportNestedMembership

`func (o *TestConnectionRequest) GetImportNestedMembership() string`

GetImportNestedMembership returns the ImportNestedMembership field if non-nil, zero value otherwise.

### GetImportNestedMembershipOk

`func (o *TestConnectionRequest) GetImportNestedMembershipOk() (*string, bool)`

GetImportNestedMembershipOk returns a tuple with the ImportNestedMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNestedMembership

`func (o *TestConnectionRequest) SetImportNestedMembership(v string)`

SetImportNestedMembership sets ImportNestedMembership field to given value.

### HasImportNestedMembership

`func (o *TestConnectionRequest) HasImportNestedMembership() bool`

HasImportNestedMembership returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *TestConnectionRequest) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *TestConnectionRequest) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *TestConnectionRequest) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *TestConnectionRequest) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *TestConnectionRequest) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *TestConnectionRequest) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *TestConnectionRequest) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *TestConnectionRequest) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *TestConnectionRequest) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *TestConnectionRequest) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *TestConnectionRequest) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *TestConnectionRequest) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *TestConnectionRequest) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *TestConnectionRequest) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *TestConnectionRequest) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *TestConnectionRequest) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *TestConnectionRequest) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *TestConnectionRequest) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *TestConnectionRequest) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *TestConnectionRequest) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVEACCOUNTJSON

`func (o *TestConnectionRequest) GetREMOVEACCOUNTJSON() string`

GetREMOVEACCOUNTJSON returns the REMOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetREMOVEACCOUNTJSONOk() (*string, bool)`

GetREMOVEACCOUNTJSONOk returns a tuple with the REMOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTJSON

`func (o *TestConnectionRequest) SetREMOVEACCOUNTJSON(v string)`

SetREMOVEACCOUNTJSON sets REMOVEACCOUNTJSON field to given value.

### HasREMOVEACCOUNTJSON

`func (o *TestConnectionRequest) HasREMOVEACCOUNTJSON() bool`

HasREMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetADDACCESSJSON

`func (o *TestConnectionRequest) GetADDACCESSJSON() string`

GetADDACCESSJSON returns the ADDACCESSJSON field if non-nil, zero value otherwise.

### GetADDACCESSJSONOk

`func (o *TestConnectionRequest) GetADDACCESSJSONOk() (*string, bool)`

GetADDACCESSJSONOk returns a tuple with the ADDACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSJSON

`func (o *TestConnectionRequest) SetADDACCESSJSON(v string)`

SetADDACCESSJSON sets ADDACCESSJSON field to given value.

### HasADDACCESSJSON

`func (o *TestConnectionRequest) HasADDACCESSJSON() bool`

HasADDACCESSJSON returns a boolean if a field has been set.

### GetREMOVEACCESSJSON

`func (o *TestConnectionRequest) GetREMOVEACCESSJSON() string`

GetREMOVEACCESSJSON returns the REMOVEACCESSJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSJSONOk

`func (o *TestConnectionRequest) GetREMOVEACCESSJSONOk() (*string, bool)`

GetREMOVEACCESSJSONOk returns a tuple with the REMOVEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSJSON

`func (o *TestConnectionRequest) SetREMOVEACCESSJSON(v string)`

SetREMOVEACCESSJSON sets REMOVEACCESSJSON field to given value.

### HasREMOVEACCESSJSON

`func (o *TestConnectionRequest) HasREMOVEACCESSJSON() bool`

HasREMOVEACCESSJSON returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *TestConnectionRequest) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *TestConnectionRequest) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *TestConnectionRequest) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *TestConnectionRequest) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetMOVEACCOUNTJSON

`func (o *TestConnectionRequest) GetMOVEACCOUNTJSON() string`

GetMOVEACCOUNTJSON returns the MOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetMOVEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetMOVEACCOUNTJSONOk() (*string, bool)`

GetMOVEACCOUNTJSONOk returns a tuple with the MOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMOVEACCOUNTJSON

`func (o *TestConnectionRequest) SetMOVEACCOUNTJSON(v string)`

SetMOVEACCOUNTJSON sets MOVEACCOUNTJSON field to given value.

### HasMOVEACCOUNTJSON

`func (o *TestConnectionRequest) HasMOVEACCOUNTJSON() bool`

HasMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetCREATEGROUPJSON

`func (o *TestConnectionRequest) GetCREATEGROUPJSON() string`

GetCREATEGROUPJSON returns the CREATEGROUPJSON field if non-nil, zero value otherwise.

### GetCREATEGROUPJSONOk

`func (o *TestConnectionRequest) GetCREATEGROUPJSONOk() (*string, bool)`

GetCREATEGROUPJSONOk returns a tuple with the CREATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEGROUPJSON

`func (o *TestConnectionRequest) SetCREATEGROUPJSON(v string)`

SetCREATEGROUPJSON sets CREATEGROUPJSON field to given value.

### HasCREATEGROUPJSON

`func (o *TestConnectionRequest) HasCREATEGROUPJSON() bool`

HasCREATEGROUPJSON returns a boolean if a field has been set.

### GetUPDATEGROUPJSON

`func (o *TestConnectionRequest) GetUPDATEGROUPJSON() string`

GetUPDATEGROUPJSON returns the UPDATEGROUPJSON field if non-nil, zero value otherwise.

### GetUPDATEGROUPJSONOk

`func (o *TestConnectionRequest) GetUPDATEGROUPJSONOk() (*string, bool)`

GetUPDATEGROUPJSONOk returns a tuple with the UPDATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEGROUPJSON

`func (o *TestConnectionRequest) SetUPDATEGROUPJSON(v string)`

SetUPDATEGROUPJSON sets UPDATEGROUPJSON field to given value.

### HasUPDATEGROUPJSON

`func (o *TestConnectionRequest) HasUPDATEGROUPJSON() bool`

HasUPDATEGROUPJSON returns a boolean if a field has been set.

### GetREMOVEGROUPJSON

`func (o *TestConnectionRequest) GetREMOVEGROUPJSON() string`

GetREMOVEGROUPJSON returns the REMOVEGROUPJSON field if non-nil, zero value otherwise.

### GetREMOVEGROUPJSONOk

`func (o *TestConnectionRequest) GetREMOVEGROUPJSONOk() (*string, bool)`

GetREMOVEGROUPJSONOk returns a tuple with the REMOVEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEGROUPJSON

`func (o *TestConnectionRequest) SetREMOVEGROUPJSON(v string)`

SetREMOVEGROUPJSON sets REMOVEGROUPJSON field to given value.

### HasREMOVEGROUPJSON

`func (o *TestConnectionRequest) HasREMOVEGROUPJSON() bool`

HasREMOVEGROUPJSON returns a boolean if a field has been set.

### GetADDACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) GetADDACCESSENTITLEMENTJSON() string`

GetADDACCESSENTITLEMENTJSON returns the ADDACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetADDACCESSENTITLEMENTJSONOk

`func (o *TestConnectionRequest) GetADDACCESSENTITLEMENTJSONOk() (*string, bool)`

GetADDACCESSENTITLEMENTJSONOk returns a tuple with the ADDACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) SetADDACCESSENTITLEMENTJSON(v string)`

SetADDACCESSENTITLEMENTJSON sets ADDACCESSENTITLEMENTJSON field to given value.

### HasADDACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) HasADDACCESSENTITLEMENTJSON() bool`

HasADDACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCUSTOMCONFIGJSON

`func (o *TestConnectionRequest) GetCUSTOMCONFIGJSON() string`

GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field if non-nil, zero value otherwise.

### GetCUSTOMCONFIGJSONOk

`func (o *TestConnectionRequest) GetCUSTOMCONFIGJSONOk() (*string, bool)`

GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMCONFIGJSON

`func (o *TestConnectionRequest) SetCUSTOMCONFIGJSON(v string)`

SetCUSTOMCONFIGJSON sets CUSTOMCONFIGJSON field to given value.

### HasCUSTOMCONFIGJSON

`func (o *TestConnectionRequest) HasCUSTOMCONFIGJSON() bool`

HasCUSTOMCONFIGJSON returns a boolean if a field has been set.

### GetREMOVEACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) GetREMOVEACCESSENTITLEMENTJSON() string`

GetREMOVEACCESSENTITLEMENTJSON returns the REMOVEACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSENTITLEMENTJSONOk

`func (o *TestConnectionRequest) GetREMOVEACCESSENTITLEMENTJSONOk() (*string, bool)`

GetREMOVEACCESSENTITLEMENTJSONOk returns a tuple with the REMOVEACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) SetREMOVEACCESSENTITLEMENTJSON(v string)`

SetREMOVEACCESSENTITLEMENTJSON sets REMOVEACCESSENTITLEMENTJSON field to given value.

### HasREMOVEACCESSENTITLEMENTJSON

`func (o *TestConnectionRequest) HasREMOVEACCESSENTITLEMENTJSON() bool`

HasREMOVEACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCREATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) GetCREATESERVICEACCOUNTJSON() string`

GetCREATESERVICEACCOUNTJSON returns the CREATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATESERVICEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetCREATESERVICEACCOUNTJSONOk() (*string, bool)`

GetCREATESERVICEACCOUNTJSONOk returns a tuple with the CREATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) SetCREATESERVICEACCOUNTJSON(v string)`

SetCREATESERVICEACCOUNTJSON sets CREATESERVICEACCOUNTJSON field to given value.

### HasCREATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) HasCREATESERVICEACCOUNTJSON() bool`

HasCREATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) GetUPDATESERVICEACCOUNTJSON() string`

GetUPDATESERVICEACCOUNTJSON returns the UPDATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATESERVICEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetUPDATESERVICEACCOUNTJSONOk() (*string, bool)`

GetUPDATESERVICEACCOUNTJSONOk returns a tuple with the UPDATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) SetUPDATESERVICEACCOUNTJSON(v string)`

SetUPDATESERVICEACCOUNTJSON sets UPDATESERVICEACCOUNTJSON field to given value.

### HasUPDATESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) HasUPDATESERVICEACCOUNTJSON() bool`

HasUPDATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) GetREMOVESERVICEACCOUNTJSON() string`

GetREMOVESERVICEACCOUNTJSON returns the REMOVESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVESERVICEACCOUNTJSONOk

`func (o *TestConnectionRequest) GetREMOVESERVICEACCOUNTJSONOk() (*string, bool)`

GetREMOVESERVICEACCOUNTJSONOk returns a tuple with the REMOVESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) SetREMOVESERVICEACCOUNTJSON(v string)`

SetREMOVESERVICEACCOUNTJSON sets REMOVESERVICEACCOUNTJSON field to given value.

### HasREMOVESERVICEACCOUNTJSON

`func (o *TestConnectionRequest) HasREMOVESERVICEACCOUNTJSON() bool`

HasREMOVESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *TestConnectionRequest) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *TestConnectionRequest) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *TestConnectionRequest) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *TestConnectionRequest) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *TestConnectionRequest) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *TestConnectionRequest) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *TestConnectionRequest) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *TestConnectionRequest) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetMESSAGESERVER

`func (o *TestConnectionRequest) GetMESSAGESERVER() string`

GetMESSAGESERVER returns the MESSAGESERVER field if non-nil, zero value otherwise.

### GetMESSAGESERVEROk

`func (o *TestConnectionRequest) GetMESSAGESERVEROk() (*string, bool)`

GetMESSAGESERVEROk returns a tuple with the MESSAGESERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMESSAGESERVER

`func (o *TestConnectionRequest) SetMESSAGESERVER(v string)`

SetMESSAGESERVER sets MESSAGESERVER field to given value.

### HasMESSAGESERVER

`func (o *TestConnectionRequest) HasMESSAGESERVER() bool`

HasMESSAGESERVER returns a boolean if a field has been set.

### GetJCO_ASHOST

`func (o *TestConnectionRequest) GetJCO_ASHOST() string`

GetJCO_ASHOST returns the JCO_ASHOST field if non-nil, zero value otherwise.

### GetJCO_ASHOSTOk

`func (o *TestConnectionRequest) GetJCO_ASHOSTOk() (*string, bool)`

GetJCO_ASHOSTOk returns a tuple with the JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_ASHOST

`func (o *TestConnectionRequest) SetJCO_ASHOST(v string)`

SetJCO_ASHOST sets JCO_ASHOST field to given value.

### HasJCO_ASHOST

`func (o *TestConnectionRequest) HasJCO_ASHOST() bool`

HasJCO_ASHOST returns a boolean if a field has been set.

### GetJCO_SYSNR

`func (o *TestConnectionRequest) GetJCO_SYSNR() string`

GetJCO_SYSNR returns the JCO_SYSNR field if non-nil, zero value otherwise.

### GetJCO_SYSNROk

`func (o *TestConnectionRequest) GetJCO_SYSNROk() (*string, bool)`

GetJCO_SYSNROk returns a tuple with the JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SYSNR

`func (o *TestConnectionRequest) SetJCO_SYSNR(v string)`

SetJCO_SYSNR sets JCO_SYSNR field to given value.

### HasJCO_SYSNR

`func (o *TestConnectionRequest) HasJCO_SYSNR() bool`

HasJCO_SYSNR returns a boolean if a field has been set.

### GetJCO_CLIENT

`func (o *TestConnectionRequest) GetJCO_CLIENT() string`

GetJCO_CLIENT returns the JCO_CLIENT field if non-nil, zero value otherwise.

### GetJCO_CLIENTOk

`func (o *TestConnectionRequest) GetJCO_CLIENTOk() (*string, bool)`

GetJCO_CLIENTOk returns a tuple with the JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_CLIENT

`func (o *TestConnectionRequest) SetJCO_CLIENT(v string)`

SetJCO_CLIENT sets JCO_CLIENT field to given value.

### HasJCO_CLIENT

`func (o *TestConnectionRequest) HasJCO_CLIENT() bool`

HasJCO_CLIENT returns a boolean if a field has been set.

### GetJCO_USER

`func (o *TestConnectionRequest) GetJCO_USER() string`

GetJCO_USER returns the JCO_USER field if non-nil, zero value otherwise.

### GetJCO_USEROk

`func (o *TestConnectionRequest) GetJCO_USEROk() (*string, bool)`

GetJCO_USEROk returns a tuple with the JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_USER

`func (o *TestConnectionRequest) SetJCO_USER(v string)`

SetJCO_USER sets JCO_USER field to given value.

### HasJCO_USER

`func (o *TestConnectionRequest) HasJCO_USER() bool`

HasJCO_USER returns a boolean if a field has been set.

### GetJCO_LANG

`func (o *TestConnectionRequest) GetJCO_LANG() string`

GetJCO_LANG returns the JCO_LANG field if non-nil, zero value otherwise.

### GetJCO_LANGOk

`func (o *TestConnectionRequest) GetJCO_LANGOk() (*string, bool)`

GetJCO_LANGOk returns a tuple with the JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_LANG

`func (o *TestConnectionRequest) SetJCO_LANG(v string)`

SetJCO_LANG sets JCO_LANG field to given value.

### HasJCO_LANG

`func (o *TestConnectionRequest) HasJCO_LANG() bool`

HasJCO_LANG returns a boolean if a field has been set.

### GetJCOR3NAME

`func (o *TestConnectionRequest) GetJCOR3NAME() string`

GetJCOR3NAME returns the JCOR3NAME field if non-nil, zero value otherwise.

### GetJCOR3NAMEOk

`func (o *TestConnectionRequest) GetJCOR3NAMEOk() (*string, bool)`

GetJCOR3NAMEOk returns a tuple with the JCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCOR3NAME

`func (o *TestConnectionRequest) SetJCOR3NAME(v string)`

SetJCOR3NAME sets JCOR3NAME field to given value.

### HasJCOR3NAME

`func (o *TestConnectionRequest) HasJCOR3NAME() bool`

HasJCOR3NAME returns a boolean if a field has been set.

### GetJCO_MSHOST

`func (o *TestConnectionRequest) GetJCO_MSHOST() string`

GetJCO_MSHOST returns the JCO_MSHOST field if non-nil, zero value otherwise.

### GetJCO_MSHOSTOk

`func (o *TestConnectionRequest) GetJCO_MSHOSTOk() (*string, bool)`

GetJCO_MSHOSTOk returns a tuple with the JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSHOST

`func (o *TestConnectionRequest) SetJCO_MSHOST(v string)`

SetJCO_MSHOST sets JCO_MSHOST field to given value.

### HasJCO_MSHOST

`func (o *TestConnectionRequest) HasJCO_MSHOST() bool`

HasJCO_MSHOST returns a boolean if a field has been set.

### GetJCO_MSSERV

`func (o *TestConnectionRequest) GetJCO_MSSERV() string`

GetJCO_MSSERV returns the JCO_MSSERV field if non-nil, zero value otherwise.

### GetJCO_MSSERVOk

`func (o *TestConnectionRequest) GetJCO_MSSERVOk() (*string, bool)`

GetJCO_MSSERVOk returns a tuple with the JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSSERV

`func (o *TestConnectionRequest) SetJCO_MSSERV(v string)`

SetJCO_MSSERV sets JCO_MSSERV field to given value.

### HasJCO_MSSERV

`func (o *TestConnectionRequest) HasJCO_MSSERV() bool`

HasJCO_MSSERV returns a boolean if a field has been set.

### GetJCO_GROUP

`func (o *TestConnectionRequest) GetJCO_GROUP() string`

GetJCO_GROUP returns the JCO_GROUP field if non-nil, zero value otherwise.

### GetJCO_GROUPOk

`func (o *TestConnectionRequest) GetJCO_GROUPOk() (*string, bool)`

GetJCO_GROUPOk returns a tuple with the JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_GROUP

`func (o *TestConnectionRequest) SetJCO_GROUP(v string)`

SetJCO_GROUP sets JCO_GROUP field to given value.

### HasJCO_GROUP

`func (o *TestConnectionRequest) HasJCO_GROUP() bool`

HasJCO_GROUP returns a boolean if a field has been set.

### GetSNC

`func (o *TestConnectionRequest) GetSNC() string`

GetSNC returns the SNC field if non-nil, zero value otherwise.

### GetSNCOk

`func (o *TestConnectionRequest) GetSNCOk() (*string, bool)`

GetSNCOk returns a tuple with the SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNC

`func (o *TestConnectionRequest) SetSNC(v string)`

SetSNC sets SNC field to given value.

### HasSNC

`func (o *TestConnectionRequest) HasSNC() bool`

HasSNC returns a boolean if a field has been set.

### GetJCO_SNC_MODE

`func (o *TestConnectionRequest) GetJCO_SNC_MODE() string`

GetJCO_SNC_MODE returns the JCO_SNC_MODE field if non-nil, zero value otherwise.

### GetJCO_SNC_MODEOk

`func (o *TestConnectionRequest) GetJCO_SNC_MODEOk() (*string, bool)`

GetJCO_SNC_MODEOk returns a tuple with the JCO_SNC_MODE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MODE

`func (o *TestConnectionRequest) SetJCO_SNC_MODE(v string)`

SetJCO_SNC_MODE sets JCO_SNC_MODE field to given value.

### HasJCO_SNC_MODE

`func (o *TestConnectionRequest) HasJCO_SNC_MODE() bool`

HasJCO_SNC_MODE returns a boolean if a field has been set.

### GetJCO_SNC_PARTNERNAME

`func (o *TestConnectionRequest) GetJCO_SNC_PARTNERNAME() string`

GetJCO_SNC_PARTNERNAME returns the JCO_SNC_PARTNERNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_PARTNERNAMEOk

`func (o *TestConnectionRequest) GetJCO_SNC_PARTNERNAMEOk() (*string, bool)`

GetJCO_SNC_PARTNERNAMEOk returns a tuple with the JCO_SNC_PARTNERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_PARTNERNAME

`func (o *TestConnectionRequest) SetJCO_SNC_PARTNERNAME(v string)`

SetJCO_SNC_PARTNERNAME sets JCO_SNC_PARTNERNAME field to given value.

### HasJCO_SNC_PARTNERNAME

`func (o *TestConnectionRequest) HasJCO_SNC_PARTNERNAME() bool`

HasJCO_SNC_PARTNERNAME returns a boolean if a field has been set.

### GetJCO_SNC_MYNAME

`func (o *TestConnectionRequest) GetJCO_SNC_MYNAME() string`

GetJCO_SNC_MYNAME returns the JCO_SNC_MYNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_MYNAMEOk

`func (o *TestConnectionRequest) GetJCO_SNC_MYNAMEOk() (*string, bool)`

GetJCO_SNC_MYNAMEOk returns a tuple with the JCO_SNC_MYNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MYNAME

`func (o *TestConnectionRequest) SetJCO_SNC_MYNAME(v string)`

SetJCO_SNC_MYNAME sets JCO_SNC_MYNAME field to given value.

### HasJCO_SNC_MYNAME

`func (o *TestConnectionRequest) HasJCO_SNC_MYNAME() bool`

HasJCO_SNC_MYNAME returns a boolean if a field has been set.

### GetJCO_SNC_LIBRARY

`func (o *TestConnectionRequest) GetJCO_SNC_LIBRARY() string`

GetJCO_SNC_LIBRARY returns the JCO_SNC_LIBRARY field if non-nil, zero value otherwise.

### GetJCO_SNC_LIBRARYOk

`func (o *TestConnectionRequest) GetJCO_SNC_LIBRARYOk() (*string, bool)`

GetJCO_SNC_LIBRARYOk returns a tuple with the JCO_SNC_LIBRARY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_LIBRARY

`func (o *TestConnectionRequest) SetJCO_SNC_LIBRARY(v string)`

SetJCO_SNC_LIBRARY sets JCO_SNC_LIBRARY field to given value.

### HasJCO_SNC_LIBRARY

`func (o *TestConnectionRequest) HasJCO_SNC_LIBRARY() bool`

HasJCO_SNC_LIBRARY returns a boolean if a field has been set.

### GetJCO_SNC_QOP

`func (o *TestConnectionRequest) GetJCO_SNC_QOP() string`

GetJCO_SNC_QOP returns the JCO_SNC_QOP field if non-nil, zero value otherwise.

### GetJCO_SNC_QOPOk

`func (o *TestConnectionRequest) GetJCO_SNC_QOPOk() (*string, bool)`

GetJCO_SNC_QOPOk returns a tuple with the JCO_SNC_QOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_QOP

`func (o *TestConnectionRequest) SetJCO_SNC_QOP(v string)`

SetJCO_SNC_QOP sets JCO_SNC_QOP field to given value.

### HasJCO_SNC_QOP

`func (o *TestConnectionRequest) HasJCO_SNC_QOP() bool`

HasJCO_SNC_QOP returns a boolean if a field has been set.

### GetTABLES

`func (o *TestConnectionRequest) GetTABLES() string`

GetTABLES returns the TABLES field if non-nil, zero value otherwise.

### GetTABLESOk

`func (o *TestConnectionRequest) GetTABLESOk() (*string, bool)`

GetTABLESOk returns a tuple with the TABLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTABLES

`func (o *TestConnectionRequest) SetTABLES(v string)`

SetTABLES sets TABLES field to given value.

### HasTABLES

`func (o *TestConnectionRequest) HasTABLES() bool`

HasTABLES returns a boolean if a field has been set.

### GetSYSTEMNAME

`func (o *TestConnectionRequest) GetSYSTEMNAME() string`

GetSYSTEMNAME returns the SYSTEMNAME field if non-nil, zero value otherwise.

### GetSYSTEMNAMEOk

`func (o *TestConnectionRequest) GetSYSTEMNAMEOk() (*string, bool)`

GetSYSTEMNAMEOk returns a tuple with the SYSTEMNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMNAME

`func (o *TestConnectionRequest) SetSYSTEMNAME(v string)`

SetSYSTEMNAME sets SYSTEMNAME field to given value.

### HasSYSTEMNAME

`func (o *TestConnectionRequest) HasSYSTEMNAME() bool`

HasSYSTEMNAME returns a boolean if a field has been set.

### GetTERMINATEDUSERGROUP

`func (o *TestConnectionRequest) GetTERMINATEDUSERGROUP() string`

GetTERMINATEDUSERGROUP returns the TERMINATEDUSERGROUP field if non-nil, zero value otherwise.

### GetTERMINATEDUSERGROUPOk

`func (o *TestConnectionRequest) GetTERMINATEDUSERGROUPOk() (*string, bool)`

GetTERMINATEDUSERGROUPOk returns a tuple with the TERMINATEDUSERGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATEDUSERGROUP

`func (o *TestConnectionRequest) SetTERMINATEDUSERGROUP(v string)`

SetTERMINATEDUSERGROUP sets TERMINATEDUSERGROUP field to given value.

### HasTERMINATEDUSERGROUP

`func (o *TestConnectionRequest) HasTERMINATEDUSERGROUP() bool`

HasTERMINATEDUSERGROUP returns a boolean if a field has been set.

### GetTERMINATED_USER_ROLE_ACTION

`func (o *TestConnectionRequest) GetTERMINATED_USER_ROLE_ACTION() string`

GetTERMINATED_USER_ROLE_ACTION returns the TERMINATED_USER_ROLE_ACTION field if non-nil, zero value otherwise.

### GetTERMINATED_USER_ROLE_ACTIONOk

`func (o *TestConnectionRequest) GetTERMINATED_USER_ROLE_ACTIONOk() (*string, bool)`

GetTERMINATED_USER_ROLE_ACTIONOk returns a tuple with the TERMINATED_USER_ROLE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATED_USER_ROLE_ACTION

`func (o *TestConnectionRequest) SetTERMINATED_USER_ROLE_ACTION(v string)`

SetTERMINATED_USER_ROLE_ACTION sets TERMINATED_USER_ROLE_ACTION field to given value.

### HasTERMINATED_USER_ROLE_ACTION

`func (o *TestConnectionRequest) HasTERMINATED_USER_ROLE_ACTION() bool`

HasTERMINATED_USER_ROLE_ACTION returns a boolean if a field has been set.

### GetPROV_JCO_ASHOST

`func (o *TestConnectionRequest) GetPROV_JCO_ASHOST() string`

GetPROV_JCO_ASHOST returns the PROV_JCO_ASHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_ASHOSTOk

`func (o *TestConnectionRequest) GetPROV_JCO_ASHOSTOk() (*string, bool)`

GetPROV_JCO_ASHOSTOk returns a tuple with the PROV_JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_ASHOST

`func (o *TestConnectionRequest) SetPROV_JCO_ASHOST(v string)`

SetPROV_JCO_ASHOST sets PROV_JCO_ASHOST field to given value.

### HasPROV_JCO_ASHOST

`func (o *TestConnectionRequest) HasPROV_JCO_ASHOST() bool`

HasPROV_JCO_ASHOST returns a boolean if a field has been set.

### GetPROV_JCO_SYSNR

`func (o *TestConnectionRequest) GetPROV_JCO_SYSNR() string`

GetPROV_JCO_SYSNR returns the PROV_JCO_SYSNR field if non-nil, zero value otherwise.

### GetPROV_JCO_SYSNROk

`func (o *TestConnectionRequest) GetPROV_JCO_SYSNROk() (*string, bool)`

GetPROV_JCO_SYSNROk returns a tuple with the PROV_JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_SYSNR

`func (o *TestConnectionRequest) SetPROV_JCO_SYSNR(v string)`

SetPROV_JCO_SYSNR sets PROV_JCO_SYSNR field to given value.

### HasPROV_JCO_SYSNR

`func (o *TestConnectionRequest) HasPROV_JCO_SYSNR() bool`

HasPROV_JCO_SYSNR returns a boolean if a field has been set.

### GetPROV_JCO_CLIENT

`func (o *TestConnectionRequest) GetPROV_JCO_CLIENT() string`

GetPROV_JCO_CLIENT returns the PROV_JCO_CLIENT field if non-nil, zero value otherwise.

### GetPROV_JCO_CLIENTOk

`func (o *TestConnectionRequest) GetPROV_JCO_CLIENTOk() (*string, bool)`

GetPROV_JCO_CLIENTOk returns a tuple with the PROV_JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_CLIENT

`func (o *TestConnectionRequest) SetPROV_JCO_CLIENT(v string)`

SetPROV_JCO_CLIENT sets PROV_JCO_CLIENT field to given value.

### HasPROV_JCO_CLIENT

`func (o *TestConnectionRequest) HasPROV_JCO_CLIENT() bool`

HasPROV_JCO_CLIENT returns a boolean if a field has been set.

### GetPROV_JCO_USER

`func (o *TestConnectionRequest) GetPROV_JCO_USER() string`

GetPROV_JCO_USER returns the PROV_JCO_USER field if non-nil, zero value otherwise.

### GetPROV_JCO_USEROk

`func (o *TestConnectionRequest) GetPROV_JCO_USEROk() (*string, bool)`

GetPROV_JCO_USEROk returns a tuple with the PROV_JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_USER

`func (o *TestConnectionRequest) SetPROV_JCO_USER(v string)`

SetPROV_JCO_USER sets PROV_JCO_USER field to given value.

### HasPROV_JCO_USER

`func (o *TestConnectionRequest) HasPROV_JCO_USER() bool`

HasPROV_JCO_USER returns a boolean if a field has been set.

### GetPROV_PASSWORD

`func (o *TestConnectionRequest) GetPROV_PASSWORD() string`

GetPROV_PASSWORD returns the PROV_PASSWORD field if non-nil, zero value otherwise.

### GetPROV_PASSWORDOk

`func (o *TestConnectionRequest) GetPROV_PASSWORDOk() (*string, bool)`

GetPROV_PASSWORDOk returns a tuple with the PROV_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_PASSWORD

`func (o *TestConnectionRequest) SetPROV_PASSWORD(v string)`

SetPROV_PASSWORD sets PROV_PASSWORD field to given value.

### HasPROV_PASSWORD

`func (o *TestConnectionRequest) HasPROV_PASSWORD() bool`

HasPROV_PASSWORD returns a boolean if a field has been set.

### GetPROV_JCO_LANG

`func (o *TestConnectionRequest) GetPROV_JCO_LANG() string`

GetPROV_JCO_LANG returns the PROV_JCO_LANG field if non-nil, zero value otherwise.

### GetPROV_JCO_LANGOk

`func (o *TestConnectionRequest) GetPROV_JCO_LANGOk() (*string, bool)`

GetPROV_JCO_LANGOk returns a tuple with the PROV_JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_LANG

`func (o *TestConnectionRequest) SetPROV_JCO_LANG(v string)`

SetPROV_JCO_LANG sets PROV_JCO_LANG field to given value.

### HasPROV_JCO_LANG

`func (o *TestConnectionRequest) HasPROV_JCO_LANG() bool`

HasPROV_JCO_LANG returns a boolean if a field has been set.

### GetPROVJCOR3NAME

`func (o *TestConnectionRequest) GetPROVJCOR3NAME() string`

GetPROVJCOR3NAME returns the PROVJCOR3NAME field if non-nil, zero value otherwise.

### GetPROVJCOR3NAMEOk

`func (o *TestConnectionRequest) GetPROVJCOR3NAMEOk() (*string, bool)`

GetPROVJCOR3NAMEOk returns a tuple with the PROVJCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVJCOR3NAME

`func (o *TestConnectionRequest) SetPROVJCOR3NAME(v string)`

SetPROVJCOR3NAME sets PROVJCOR3NAME field to given value.

### HasPROVJCOR3NAME

`func (o *TestConnectionRequest) HasPROVJCOR3NAME() bool`

HasPROVJCOR3NAME returns a boolean if a field has been set.

### GetPROV_JCO_MSHOST

`func (o *TestConnectionRequest) GetPROV_JCO_MSHOST() string`

GetPROV_JCO_MSHOST returns the PROV_JCO_MSHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_MSHOSTOk

`func (o *TestConnectionRequest) GetPROV_JCO_MSHOSTOk() (*string, bool)`

GetPROV_JCO_MSHOSTOk returns a tuple with the PROV_JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSHOST

`func (o *TestConnectionRequest) SetPROV_JCO_MSHOST(v string)`

SetPROV_JCO_MSHOST sets PROV_JCO_MSHOST field to given value.

### HasPROV_JCO_MSHOST

`func (o *TestConnectionRequest) HasPROV_JCO_MSHOST() bool`

HasPROV_JCO_MSHOST returns a boolean if a field has been set.

### GetPROV_JCO_MSSERV

`func (o *TestConnectionRequest) GetPROV_JCO_MSSERV() string`

GetPROV_JCO_MSSERV returns the PROV_JCO_MSSERV field if non-nil, zero value otherwise.

### GetPROV_JCO_MSSERVOk

`func (o *TestConnectionRequest) GetPROV_JCO_MSSERVOk() (*string, bool)`

GetPROV_JCO_MSSERVOk returns a tuple with the PROV_JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSSERV

`func (o *TestConnectionRequest) SetPROV_JCO_MSSERV(v string)`

SetPROV_JCO_MSSERV sets PROV_JCO_MSSERV field to given value.

### HasPROV_JCO_MSSERV

`func (o *TestConnectionRequest) HasPROV_JCO_MSSERV() bool`

HasPROV_JCO_MSSERV returns a boolean if a field has been set.

### GetPROV_JCO_GROUP

`func (o *TestConnectionRequest) GetPROV_JCO_GROUP() string`

GetPROV_JCO_GROUP returns the PROV_JCO_GROUP field if non-nil, zero value otherwise.

### GetPROV_JCO_GROUPOk

`func (o *TestConnectionRequest) GetPROV_JCO_GROUPOk() (*string, bool)`

GetPROV_JCO_GROUPOk returns a tuple with the PROV_JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_GROUP

`func (o *TestConnectionRequest) SetPROV_JCO_GROUP(v string)`

SetPROV_JCO_GROUP sets PROV_JCO_GROUP field to given value.

### HasPROV_JCO_GROUP

`func (o *TestConnectionRequest) HasPROV_JCO_GROUP() bool`

HasPROV_JCO_GROUP returns a boolean if a field has been set.

### GetPROV_CUA_ENABLED

`func (o *TestConnectionRequest) GetPROV_CUA_ENABLED() string`

GetPROV_CUA_ENABLED returns the PROV_CUA_ENABLED field if non-nil, zero value otherwise.

### GetPROV_CUA_ENABLEDOk

`func (o *TestConnectionRequest) GetPROV_CUA_ENABLEDOk() (*string, bool)`

GetPROV_CUA_ENABLEDOk returns a tuple with the PROV_CUA_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_ENABLED

`func (o *TestConnectionRequest) SetPROV_CUA_ENABLED(v string)`

SetPROV_CUA_ENABLED sets PROV_CUA_ENABLED field to given value.

### HasPROV_CUA_ENABLED

`func (o *TestConnectionRequest) HasPROV_CUA_ENABLED() bool`

HasPROV_CUA_ENABLED returns a boolean if a field has been set.

### GetPROV_CUA_SNC

`func (o *TestConnectionRequest) GetPROV_CUA_SNC() string`

GetPROV_CUA_SNC returns the PROV_CUA_SNC field if non-nil, zero value otherwise.

### GetPROV_CUA_SNCOk

`func (o *TestConnectionRequest) GetPROV_CUA_SNCOk() (*string, bool)`

GetPROV_CUA_SNCOk returns a tuple with the PROV_CUA_SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_SNC

`func (o *TestConnectionRequest) SetPROV_CUA_SNC(v string)`

SetPROV_CUA_SNC sets PROV_CUA_SNC field to given value.

### HasPROV_CUA_SNC

`func (o *TestConnectionRequest) HasPROV_CUA_SNC() bool`

HasPROV_CUA_SNC returns a boolean if a field has been set.

### GetRESET_PWD_FOR_NEWACCOUNT

`func (o *TestConnectionRequest) GetRESET_PWD_FOR_NEWACCOUNT() string`

GetRESET_PWD_FOR_NEWACCOUNT returns the RESET_PWD_FOR_NEWACCOUNT field if non-nil, zero value otherwise.

### GetRESET_PWD_FOR_NEWACCOUNTOk

`func (o *TestConnectionRequest) GetRESET_PWD_FOR_NEWACCOUNTOk() (*string, bool)`

GetRESET_PWD_FOR_NEWACCOUNTOk returns a tuple with the RESET_PWD_FOR_NEWACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESET_PWD_FOR_NEWACCOUNT

`func (o *TestConnectionRequest) SetRESET_PWD_FOR_NEWACCOUNT(v string)`

SetRESET_PWD_FOR_NEWACCOUNT sets RESET_PWD_FOR_NEWACCOUNT field to given value.

### HasRESET_PWD_FOR_NEWACCOUNT

`func (o *TestConnectionRequest) HasRESET_PWD_FOR_NEWACCOUNT() bool`

HasRESET_PWD_FOR_NEWACCOUNT returns a boolean if a field has been set.

### GetENFORCEPASSWORDCHANGE

`func (o *TestConnectionRequest) GetENFORCEPASSWORDCHANGE() string`

GetENFORCEPASSWORDCHANGE returns the ENFORCEPASSWORDCHANGE field if non-nil, zero value otherwise.

### GetENFORCEPASSWORDCHANGEOk

`func (o *TestConnectionRequest) GetENFORCEPASSWORDCHANGEOk() (*string, bool)`

GetENFORCEPASSWORDCHANGEOk returns a tuple with the ENFORCEPASSWORDCHANGE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCEPASSWORDCHANGE

`func (o *TestConnectionRequest) SetENFORCEPASSWORDCHANGE(v string)`

SetENFORCEPASSWORDCHANGE sets ENFORCEPASSWORDCHANGE field to given value.

### HasENFORCEPASSWORDCHANGE

`func (o *TestConnectionRequest) HasENFORCEPASSWORDCHANGE() bool`

HasENFORCEPASSWORDCHANGE returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *TestConnectionRequest) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *TestConnectionRequest) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *TestConnectionRequest) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *TestConnectionRequest) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *TestConnectionRequest) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *TestConnectionRequest) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *TestConnectionRequest) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *TestConnectionRequest) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *TestConnectionRequest) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *TestConnectionRequest) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *TestConnectionRequest) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *TestConnectionRequest) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *TestConnectionRequest) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *TestConnectionRequest) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *TestConnectionRequest) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *TestConnectionRequest) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *TestConnectionRequest) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *TestConnectionRequest) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *TestConnectionRequest) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *TestConnectionRequest) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetHANAREFTABLEJSON

`func (o *TestConnectionRequest) GetHANAREFTABLEJSON() string`

GetHANAREFTABLEJSON returns the HANAREFTABLEJSON field if non-nil, zero value otherwise.

### GetHANAREFTABLEJSONOk

`func (o *TestConnectionRequest) GetHANAREFTABLEJSONOk() (*string, bool)`

GetHANAREFTABLEJSONOk returns a tuple with the HANAREFTABLEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHANAREFTABLEJSON

`func (o *TestConnectionRequest) SetHANAREFTABLEJSON(v string)`

SetHANAREFTABLEJSON sets HANAREFTABLEJSON field to given value.

### HasHANAREFTABLEJSON

`func (o *TestConnectionRequest) HasHANAREFTABLEJSON() bool`

HasHANAREFTABLEJSON returns a boolean if a field has been set.

### GetUSERIMPORTJSON

`func (o *TestConnectionRequest) GetUSERIMPORTJSON() string`

GetUSERIMPORTJSON returns the USERIMPORTJSON field if non-nil, zero value otherwise.

### GetUSERIMPORTJSONOk

`func (o *TestConnectionRequest) GetUSERIMPORTJSONOk() (*string, bool)`

GetUSERIMPORTJSONOk returns a tuple with the USERIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORTJSON

`func (o *TestConnectionRequest) SetUSERIMPORTJSON(v string)`

SetUSERIMPORTJSON sets USERIMPORTJSON field to given value.

### HasUSERIMPORTJSON

`func (o *TestConnectionRequest) HasUSERIMPORTJSON() bool`

HasUSERIMPORTJSON returns a boolean if a field has been set.

### GetSETCUASYSTEM

`func (o *TestConnectionRequest) GetSETCUASYSTEM() string`

GetSETCUASYSTEM returns the SETCUASYSTEM field if non-nil, zero value otherwise.

### GetSETCUASYSTEMOk

`func (o *TestConnectionRequest) GetSETCUASYSTEMOk() (*string, bool)`

GetSETCUASYSTEMOk returns a tuple with the SETCUASYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETCUASYSTEM

`func (o *TestConnectionRequest) SetSETCUASYSTEM(v string)`

SetSETCUASYSTEM sets SETCUASYSTEM field to given value.

### HasSETCUASYSTEM

`func (o *TestConnectionRequest) HasSETCUASYSTEM() bool`

HasSETCUASYSTEM returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *TestConnectionRequest) GetFIREFIGHTERID_GRANT_ACCESS_JSON() string`

GetFIREFIGHTERID_GRANT_ACCESS_JSON returns the FIREFIGHTERID_GRANT_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_JSONOk

`func (o *TestConnectionRequest) GetFIREFIGHTERID_GRANT_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *TestConnectionRequest) SetFIREFIGHTERID_GRANT_ACCESS_JSON(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_JSON sets FIREFIGHTERID_GRANT_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *TestConnectionRequest) HasFIREFIGHTERID_GRANT_ACCESS_JSON() bool`

HasFIREFIGHTERID_GRANT_ACCESS_JSON returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *TestConnectionRequest) GetFIREFIGHTERID_REVOKE_ACCESS_JSON() string`

GetFIREFIGHTERID_REVOKE_ACCESS_JSON returns the FIREFIGHTERID_REVOKE_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk

`func (o *TestConnectionRequest) GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *TestConnectionRequest) SetFIREFIGHTERID_REVOKE_ACCESS_JSON(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_JSON sets FIREFIGHTERID_REVOKE_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *TestConnectionRequest) HasFIREFIGHTERID_REVOKE_ACCESS_JSON() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_JSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON

`func (o *TestConnectionRequest) GetEXTERNAL_SOD_EVAL_JSON() string`

GetEXTERNAL_SOD_EVAL_JSON returns the EXTERNAL_SOD_EVAL_JSON field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSONOk

`func (o *TestConnectionRequest) GetEXTERNAL_SOD_EVAL_JSONOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSONOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON

`func (o *TestConnectionRequest) SetEXTERNAL_SOD_EVAL_JSON(v string)`

SetEXTERNAL_SOD_EVAL_JSON sets EXTERNAL_SOD_EVAL_JSON field to given value.

### HasEXTERNAL_SOD_EVAL_JSON

`func (o *TestConnectionRequest) HasEXTERNAL_SOD_EVAL_JSON() bool`

HasEXTERNAL_SOD_EVAL_JSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *TestConnectionRequest) GetEXTERNAL_SOD_EVAL_JSON_DETAIL() string`

GetEXTERNAL_SOD_EVAL_JSON_DETAIL returns the EXTERNAL_SOD_EVAL_JSON_DETAIL field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSON_DETAILOk

`func (o *TestConnectionRequest) GetEXTERNAL_SOD_EVAL_JSON_DETAILOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSON_DETAILOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON_DETAIL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *TestConnectionRequest) SetEXTERNAL_SOD_EVAL_JSON_DETAIL(v string)`

SetEXTERNAL_SOD_EVAL_JSON_DETAIL sets EXTERNAL_SOD_EVAL_JSON_DETAIL field to given value.

### HasEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *TestConnectionRequest) HasEXTERNAL_SOD_EVAL_JSON_DETAIL() bool`

HasEXTERNAL_SOD_EVAL_JSON_DETAIL returns a boolean if a field has been set.

### GetLOGS_TABLE_FILTER

`func (o *TestConnectionRequest) GetLOGS_TABLE_FILTER() string`

GetLOGS_TABLE_FILTER returns the LOGS_TABLE_FILTER field if non-nil, zero value otherwise.

### GetLOGS_TABLE_FILTEROk

`func (o *TestConnectionRequest) GetLOGS_TABLE_FILTEROk() (*string, bool)`

GetLOGS_TABLE_FILTEROk returns a tuple with the LOGS_TABLE_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGS_TABLE_FILTER

`func (o *TestConnectionRequest) SetLOGS_TABLE_FILTER(v string)`

SetLOGS_TABLE_FILTER sets LOGS_TABLE_FILTER field to given value.

### HasLOGS_TABLE_FILTER

`func (o *TestConnectionRequest) HasLOGS_TABLE_FILTER() bool`

HasLOGS_TABLE_FILTER returns a boolean if a field has been set.

### GetSAPTABLE_FILTER_LANG

`func (o *TestConnectionRequest) GetSAPTABLE_FILTER_LANG() string`

GetSAPTABLE_FILTER_LANG returns the SAPTABLE_FILTER_LANG field if non-nil, zero value otherwise.

### GetSAPTABLE_FILTER_LANGOk

`func (o *TestConnectionRequest) GetSAPTABLE_FILTER_LANGOk() (*string, bool)`

GetSAPTABLE_FILTER_LANGOk returns a tuple with the SAPTABLE_FILTER_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAPTABLE_FILTER_LANG

`func (o *TestConnectionRequest) SetSAPTABLE_FILTER_LANG(v string)`

SetSAPTABLE_FILTER_LANG sets SAPTABLE_FILTER_LANG field to given value.

### HasSAPTABLE_FILTER_LANG

`func (o *TestConnectionRequest) HasSAPTABLE_FILTER_LANG() bool`

HasSAPTABLE_FILTER_LANG returns a boolean if a field has been set.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *TestConnectionRequest) GetALTERNATE_OUTPUT_PARAMETER_ET_DATA() string`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATA returns the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if non-nil, zero value otherwise.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk

`func (o *TestConnectionRequest) GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk() (*string, bool)`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk returns a tuple with the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *TestConnectionRequest) SetALTERNATE_OUTPUT_PARAMETER_ET_DATA(v string)`

SetALTERNATE_OUTPUT_PARAMETER_ET_DATA sets ALTERNATE_OUTPUT_PARAMETER_ET_DATA field to given value.

### HasALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *TestConnectionRequest) HasALTERNATE_OUTPUT_PARAMETER_ET_DATA() bool`

HasALTERNATE_OUTPUT_PARAMETER_ET_DATA returns a boolean if a field has been set.

### GetAUDIT_LOG_JSON

`func (o *TestConnectionRequest) GetAUDIT_LOG_JSON() string`

GetAUDIT_LOG_JSON returns the AUDIT_LOG_JSON field if non-nil, zero value otherwise.

### GetAUDIT_LOG_JSONOk

`func (o *TestConnectionRequest) GetAUDIT_LOG_JSONOk() (*string, bool)`

GetAUDIT_LOG_JSONOk returns a tuple with the AUDIT_LOG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_LOG_JSON

`func (o *TestConnectionRequest) SetAUDIT_LOG_JSON(v string)`

SetAUDIT_LOG_JSON sets AUDIT_LOG_JSON field to given value.

### HasAUDIT_LOG_JSON

`func (o *TestConnectionRequest) HasAUDIT_LOG_JSON() bool`

HasAUDIT_LOG_JSON returns a boolean if a field has been set.

### GetECCORS4HANA

`func (o *TestConnectionRequest) GetECCORS4HANA() string`

GetECCORS4HANA returns the ECCORS4HANA field if non-nil, zero value otherwise.

### GetECCORS4HANAOk

`func (o *TestConnectionRequest) GetECCORS4HANAOk() (*string, bool)`

GetECCORS4HANAOk returns a tuple with the ECCORS4HANA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCORS4HANA

`func (o *TestConnectionRequest) SetECCORS4HANA(v string)`

SetECCORS4HANA sets ECCORS4HANA field to given value.

### HasECCORS4HANA

`func (o *TestConnectionRequest) HasECCORS4HANA() bool`

HasECCORS4HANA returns a boolean if a field has been set.

### GetDATA_IMPORT_FILTER

`func (o *TestConnectionRequest) GetDATA_IMPORT_FILTER() string`

GetDATA_IMPORT_FILTER returns the DATA_IMPORT_FILTER field if non-nil, zero value otherwise.

### GetDATA_IMPORT_FILTEROk

`func (o *TestConnectionRequest) GetDATA_IMPORT_FILTEROk() (*string, bool)`

GetDATA_IMPORT_FILTEROk returns a tuple with the DATA_IMPORT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_IMPORT_FILTER

`func (o *TestConnectionRequest) SetDATA_IMPORT_FILTER(v string)`

SetDATA_IMPORT_FILTER sets DATA_IMPORT_FILTER field to given value.

### HasDATA_IMPORT_FILTER

`func (o *TestConnectionRequest) HasDATA_IMPORT_FILTER() bool`

HasDATA_IMPORT_FILTER returns a boolean if a field has been set.

### GetConfigJSON

`func (o *TestConnectionRequest) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *TestConnectionRequest) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *TestConnectionRequest) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *TestConnectionRequest) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *TestConnectionRequest) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *TestConnectionRequest) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *TestConnectionRequest) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *TestConnectionRequest) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *TestConnectionRequest) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *TestConnectionRequest) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetREFRESH_TOKEN

`func (o *TestConnectionRequest) GetREFRESH_TOKEN() string`

GetREFRESH_TOKEN returns the REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetREFRESH_TOKENOk

`func (o *TestConnectionRequest) GetREFRESH_TOKENOk() (*string, bool)`

GetREFRESH_TOKENOk returns a tuple with the REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREFRESH_TOKEN

`func (o *TestConnectionRequest) SetREFRESH_TOKEN(v string)`

SetREFRESH_TOKEN sets REFRESH_TOKEN field to given value.

### HasREFRESH_TOKEN

`func (o *TestConnectionRequest) HasREFRESH_TOKEN() bool`

HasREFRESH_TOKEN returns a boolean if a field has been set.

### GetREDIRECT_URI

`func (o *TestConnectionRequest) GetREDIRECT_URI() string`

GetREDIRECT_URI returns the REDIRECT_URI field if non-nil, zero value otherwise.

### GetREDIRECT_URIOk

`func (o *TestConnectionRequest) GetREDIRECT_URIOk() (*string, bool)`

GetREDIRECT_URIOk returns a tuple with the REDIRECT_URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREDIRECT_URI

`func (o *TestConnectionRequest) SetREDIRECT_URI(v string)`

SetREDIRECT_URI sets REDIRECT_URI field to given value.

### HasREDIRECT_URI

`func (o *TestConnectionRequest) HasREDIRECT_URI() bool`

HasREDIRECT_URI returns a boolean if a field has been set.

### GetINSTANCE_URL

`func (o *TestConnectionRequest) GetINSTANCE_URL() string`

GetINSTANCE_URL returns the INSTANCE_URL field if non-nil, zero value otherwise.

### GetINSTANCE_URLOk

`func (o *TestConnectionRequest) GetINSTANCE_URLOk() (*string, bool)`

GetINSTANCE_URLOk returns a tuple with the INSTANCE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINSTANCE_URL

`func (o *TestConnectionRequest) SetINSTANCE_URL(v string)`

SetINSTANCE_URL sets INSTANCE_URL field to given value.

### HasINSTANCE_URL

`func (o *TestConnectionRequest) HasINSTANCE_URL() bool`

HasINSTANCE_URL returns a boolean if a field has been set.

### GetOBJECT_TO_BE_IMPORTED

`func (o *TestConnectionRequest) GetOBJECT_TO_BE_IMPORTED() string`

GetOBJECT_TO_BE_IMPORTED returns the OBJECT_TO_BE_IMPORTED field if non-nil, zero value otherwise.

### GetOBJECT_TO_BE_IMPORTEDOk

`func (o *TestConnectionRequest) GetOBJECT_TO_BE_IMPORTEDOk() (*string, bool)`

GetOBJECT_TO_BE_IMPORTEDOk returns a tuple with the OBJECT_TO_BE_IMPORTED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECT_TO_BE_IMPORTED

`func (o *TestConnectionRequest) SetOBJECT_TO_BE_IMPORTED(v string)`

SetOBJECT_TO_BE_IMPORTED sets OBJECT_TO_BE_IMPORTED field to given value.

### HasOBJECT_TO_BE_IMPORTED

`func (o *TestConnectionRequest) HasOBJECT_TO_BE_IMPORTED() bool`

HasOBJECT_TO_BE_IMPORTED returns a boolean if a field has been set.

### GetFEATURE_LICENSE_JSON

`func (o *TestConnectionRequest) GetFEATURE_LICENSE_JSON() string`

GetFEATURE_LICENSE_JSON returns the FEATURE_LICENSE_JSON field if non-nil, zero value otherwise.

### GetFEATURE_LICENSE_JSONOk

`func (o *TestConnectionRequest) GetFEATURE_LICENSE_JSONOk() (*string, bool)`

GetFEATURE_LICENSE_JSONOk returns a tuple with the FEATURE_LICENSE_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFEATURE_LICENSE_JSON

`func (o *TestConnectionRequest) SetFEATURE_LICENSE_JSON(v string)`

SetFEATURE_LICENSE_JSON sets FEATURE_LICENSE_JSON field to given value.

### HasFEATURE_LICENSE_JSON

`func (o *TestConnectionRequest) HasFEATURE_LICENSE_JSON() bool`

HasFEATURE_LICENSE_JSON returns a boolean if a field has been set.

### GetCUSTOM_CREATEACCOUNT_URL

`func (o *TestConnectionRequest) GetCUSTOM_CREATEACCOUNT_URL() string`

GetCUSTOM_CREATEACCOUNT_URL returns the CUSTOM_CREATEACCOUNT_URL field if non-nil, zero value otherwise.

### GetCUSTOM_CREATEACCOUNT_URLOk

`func (o *TestConnectionRequest) GetCUSTOM_CREATEACCOUNT_URLOk() (*string, bool)`

GetCUSTOM_CREATEACCOUNT_URLOk returns a tuple with the CUSTOM_CREATEACCOUNT_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CREATEACCOUNT_URL

`func (o *TestConnectionRequest) SetCUSTOM_CREATEACCOUNT_URL(v string)`

SetCUSTOM_CREATEACCOUNT_URL sets CUSTOM_CREATEACCOUNT_URL field to given value.

### HasCUSTOM_CREATEACCOUNT_URL

`func (o *TestConnectionRequest) HasCUSTOM_CREATEACCOUNT_URL() bool`

HasCUSTOM_CREATEACCOUNT_URL returns a boolean if a field has been set.

### GetACCOUNT_FILTER_QUERY

`func (o *TestConnectionRequest) GetACCOUNT_FILTER_QUERY() string`

GetACCOUNT_FILTER_QUERY returns the ACCOUNT_FILTER_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FILTER_QUERYOk

`func (o *TestConnectionRequest) GetACCOUNT_FILTER_QUERYOk() (*string, bool)`

GetACCOUNT_FILTER_QUERYOk returns a tuple with the ACCOUNT_FILTER_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FILTER_QUERY

`func (o *TestConnectionRequest) SetACCOUNT_FILTER_QUERY(v string)`

SetACCOUNT_FILTER_QUERY sets ACCOUNT_FILTER_QUERY field to given value.

### HasACCOUNT_FILTER_QUERY

`func (o *TestConnectionRequest) HasACCOUNT_FILTER_QUERY() bool`

HasACCOUNT_FILTER_QUERY returns a boolean if a field has been set.

### GetACCOUNT_FIELD_QUERY

`func (o *TestConnectionRequest) GetACCOUNT_FIELD_QUERY() string`

GetACCOUNT_FIELD_QUERY returns the ACCOUNT_FIELD_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FIELD_QUERYOk

`func (o *TestConnectionRequest) GetACCOUNT_FIELD_QUERYOk() (*string, bool)`

GetACCOUNT_FIELD_QUERYOk returns a tuple with the ACCOUNT_FIELD_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FIELD_QUERY

`func (o *TestConnectionRequest) SetACCOUNT_FIELD_QUERY(v string)`

SetACCOUNT_FIELD_QUERY sets ACCOUNT_FIELD_QUERY field to given value.

### HasACCOUNT_FIELD_QUERY

`func (o *TestConnectionRequest) HasACCOUNT_FIELD_QUERY() bool`

HasACCOUNT_FIELD_QUERY returns a boolean if a field has been set.

### GetFIELD_MAPPING_JSON

`func (o *TestConnectionRequest) GetFIELD_MAPPING_JSON() string`

GetFIELD_MAPPING_JSON returns the FIELD_MAPPING_JSON field if non-nil, zero value otherwise.

### GetFIELD_MAPPING_JSONOk

`func (o *TestConnectionRequest) GetFIELD_MAPPING_JSONOk() (*string, bool)`

GetFIELD_MAPPING_JSONOk returns a tuple with the FIELD_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIELD_MAPPING_JSON

`func (o *TestConnectionRequest) SetFIELD_MAPPING_JSON(v string)`

SetFIELD_MAPPING_JSON sets FIELD_MAPPING_JSON field to given value.

### HasFIELD_MAPPING_JSON

`func (o *TestConnectionRequest) HasFIELD_MAPPING_JSON() bool`

HasFIELD_MAPPING_JSON returns a boolean if a field has been set.

### GetMODIFYACCOUNTJSON

`func (o *TestConnectionRequest) GetMODIFYACCOUNTJSON() string`

GetMODIFYACCOUNTJSON returns the MODIFYACCOUNTJSON field if non-nil, zero value otherwise.

### GetMODIFYACCOUNTJSONOk

`func (o *TestConnectionRequest) GetMODIFYACCOUNTJSONOk() (*string, bool)`

GetMODIFYACCOUNTJSONOk returns a tuple with the MODIFYACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYACCOUNTJSON

`func (o *TestConnectionRequest) SetMODIFYACCOUNTJSON(v string)`

SetMODIFYACCOUNTJSON sets MODIFYACCOUNTJSON field to given value.

### HasMODIFYACCOUNTJSON

`func (o *TestConnectionRequest) HasMODIFYACCOUNTJSON() bool`

HasMODIFYACCOUNTJSON returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *TestConnectionRequest) GetConnectionJSON() map[string]interface{}`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *TestConnectionRequest) GetConnectionJSONOk() (*map[string]interface{}, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *TestConnectionRequest) SetConnectionJSON(v map[string]interface{})`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *TestConnectionRequest) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *TestConnectionRequest) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *TestConnectionRequest) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *TestConnectionRequest) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *TestConnectionRequest) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetImportAccountEntJSON

`func (o *TestConnectionRequest) GetImportAccountEntJSON() string`

GetImportAccountEntJSON returns the ImportAccountEntJSON field if non-nil, zero value otherwise.

### GetImportAccountEntJSONOk

`func (o *TestConnectionRequest) GetImportAccountEntJSONOk() (*string, bool)`

GetImportAccountEntJSONOk returns a tuple with the ImportAccountEntJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAccountEntJSON

`func (o *TestConnectionRequest) SetImportAccountEntJSON(v string)`

SetImportAccountEntJSON sets ImportAccountEntJSON field to given value.

### HasImportAccountEntJSON

`func (o *TestConnectionRequest) HasImportAccountEntJSON() bool`

HasImportAccountEntJSON returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *TestConnectionRequest) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *TestConnectionRequest) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *TestConnectionRequest) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *TestConnectionRequest) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *TestConnectionRequest) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *TestConnectionRequest) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *TestConnectionRequest) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *TestConnectionRequest) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *TestConnectionRequest) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *TestConnectionRequest) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *TestConnectionRequest) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *TestConnectionRequest) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *TestConnectionRequest) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *TestConnectionRequest) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *TestConnectionRequest) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *TestConnectionRequest) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *TestConnectionRequest) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *TestConnectionRequest) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *TestConnectionRequest) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *TestConnectionRequest) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *TestConnectionRequest) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *TestConnectionRequest) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *TestConnectionRequest) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *TestConnectionRequest) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetUpdateUserJSON

`func (o *TestConnectionRequest) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *TestConnectionRequest) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *TestConnectionRequest) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *TestConnectionRequest) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *TestConnectionRequest) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *TestConnectionRequest) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *TestConnectionRequest) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *TestConnectionRequest) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *TestConnectionRequest) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *TestConnectionRequest) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *TestConnectionRequest) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *TestConnectionRequest) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetTicketStatusJSON

`func (o *TestConnectionRequest) GetTicketStatusJSON() string`

GetTicketStatusJSON returns the TicketStatusJSON field if non-nil, zero value otherwise.

### GetTicketStatusJSONOk

`func (o *TestConnectionRequest) GetTicketStatusJSONOk() (*string, bool)`

GetTicketStatusJSONOk returns a tuple with the TicketStatusJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatusJSON

`func (o *TestConnectionRequest) SetTicketStatusJSON(v string)`

SetTicketStatusJSON sets TicketStatusJSON field to given value.

### HasTicketStatusJSON

`func (o *TestConnectionRequest) HasTicketStatusJSON() bool`

HasTicketStatusJSON returns a boolean if a field has been set.

### GetCreateTicketJSON

`func (o *TestConnectionRequest) GetCreateTicketJSON() string`

GetCreateTicketJSON returns the CreateTicketJSON field if non-nil, zero value otherwise.

### GetCreateTicketJSONOk

`func (o *TestConnectionRequest) GetCreateTicketJSONOk() (*string, bool)`

GetCreateTicketJSONOk returns a tuple with the CreateTicketJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTicketJSON

`func (o *TestConnectionRequest) SetCreateTicketJSON(v string)`

SetCreateTicketJSON sets CreateTicketJSON field to given value.

### HasCreateTicketJSON

`func (o *TestConnectionRequest) HasCreateTicketJSON() bool`

HasCreateTicketJSON returns a boolean if a field has been set.

### GetPasswdPolicyJSON

`func (o *TestConnectionRequest) GetPasswdPolicyJSON() string`

GetPasswdPolicyJSON returns the PasswdPolicyJSON field if non-nil, zero value otherwise.

### GetPasswdPolicyJSONOk

`func (o *TestConnectionRequest) GetPasswdPolicyJSONOk() (*string, bool)`

GetPasswdPolicyJSONOk returns a tuple with the PasswdPolicyJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswdPolicyJSON

`func (o *TestConnectionRequest) SetPasswdPolicyJSON(v string)`

SetPasswdPolicyJSON sets PasswdPolicyJSON field to given value.

### HasPasswdPolicyJSON

`func (o *TestConnectionRequest) HasPasswdPolicyJSON() bool`

HasPasswdPolicyJSON returns a boolean if a field has been set.

### GetAddFFIDAccessJSON

`func (o *TestConnectionRequest) GetAddFFIDAccessJSON() string`

GetAddFFIDAccessJSON returns the AddFFIDAccessJSON field if non-nil, zero value otherwise.

### GetAddFFIDAccessJSONOk

`func (o *TestConnectionRequest) GetAddFFIDAccessJSONOk() (*string, bool)`

GetAddFFIDAccessJSONOk returns a tuple with the AddFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFFIDAccessJSON

`func (o *TestConnectionRequest) SetAddFFIDAccessJSON(v string)`

SetAddFFIDAccessJSON sets AddFFIDAccessJSON field to given value.

### HasAddFFIDAccessJSON

`func (o *TestConnectionRequest) HasAddFFIDAccessJSON() bool`

HasAddFFIDAccessJSON returns a boolean if a field has been set.

### GetRemoveFFIDAccessJSON

`func (o *TestConnectionRequest) GetRemoveFFIDAccessJSON() string`

GetRemoveFFIDAccessJSON returns the RemoveFFIDAccessJSON field if non-nil, zero value otherwise.

### GetRemoveFFIDAccessJSONOk

`func (o *TestConnectionRequest) GetRemoveFFIDAccessJSONOk() (*string, bool)`

GetRemoveFFIDAccessJSONOk returns a tuple with the RemoveFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFFIDAccessJSON

`func (o *TestConnectionRequest) SetRemoveFFIDAccessJSON(v string)`

SetRemoveFFIDAccessJSON sets RemoveFFIDAccessJSON field to given value.

### HasRemoveFFIDAccessJSON

`func (o *TestConnectionRequest) HasRemoveFFIDAccessJSON() bool`

HasRemoveFFIDAccessJSON returns a boolean if a field has been set.

### GetSendOtpJSON

`func (o *TestConnectionRequest) GetSendOtpJSON() string`

GetSendOtpJSON returns the SendOtpJSON field if non-nil, zero value otherwise.

### GetSendOtpJSONOk

`func (o *TestConnectionRequest) GetSendOtpJSONOk() (*string, bool)`

GetSendOtpJSONOk returns a tuple with the SendOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOtpJSON

`func (o *TestConnectionRequest) SetSendOtpJSON(v string)`

SetSendOtpJSON sets SendOtpJSON field to given value.

### HasSendOtpJSON

`func (o *TestConnectionRequest) HasSendOtpJSON() bool`

HasSendOtpJSON returns a boolean if a field has been set.

### GetValidateOtpJSON

`func (o *TestConnectionRequest) GetValidateOtpJSON() string`

GetValidateOtpJSON returns the ValidateOtpJSON field if non-nil, zero value otherwise.

### GetValidateOtpJSONOk

`func (o *TestConnectionRequest) GetValidateOtpJSONOk() (*string, bool)`

GetValidateOtpJSONOk returns a tuple with the ValidateOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOtpJSON

`func (o *TestConnectionRequest) SetValidateOtpJSON(v string)`

SetValidateOtpJSON sets ValidateOtpJSON field to given value.

### HasValidateOtpJSON

`func (o *TestConnectionRequest) HasValidateOtpJSON() bool`

HasValidateOtpJSON returns a boolean if a field has been set.

### GetBASEURL

`func (o *TestConnectionRequest) GetBASEURL() string`

GetBASEURL returns the BASEURL field if non-nil, zero value otherwise.

### GetBASEURLOk

`func (o *TestConnectionRequest) GetBASEURLOk() (*string, bool)`

GetBASEURLOk returns a tuple with the BASEURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASEURL

`func (o *TestConnectionRequest) SetBASEURL(v string)`

SetBASEURL sets BASEURL field to given value.


### GetTENANT_ID

`func (o *TestConnectionRequest) GetTENANT_ID() string`

GetTENANT_ID returns the TENANT_ID field if non-nil, zero value otherwise.

### GetTENANT_IDOk

`func (o *TestConnectionRequest) GetTENANT_IDOk() (*string, bool)`

GetTENANT_IDOk returns a tuple with the TENANT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_ID

`func (o *TestConnectionRequest) SetTENANT_ID(v string)`

SetTENANT_ID sets TENANT_ID field to given value.


### GetLOGIN_URL

`func (o *TestConnectionRequest) GetLOGIN_URL() string`

GetLOGIN_URL returns the LOGIN_URL field if non-nil, zero value otherwise.

### GetLOGIN_URLOk

`func (o *TestConnectionRequest) GetLOGIN_URLOk() (*string, bool)`

GetLOGIN_URLOk returns a tuple with the LOGIN_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGIN_URL

`func (o *TestConnectionRequest) SetLOGIN_URL(v string)`

SetLOGIN_URL sets LOGIN_URL field to given value.


### GetUSER_FILTER

`func (o *TestConnectionRequest) GetUSER_FILTER() string`

GetUSER_FILTER returns the USER_FILTER field if non-nil, zero value otherwise.

### GetUSER_FILTEROk

`func (o *TestConnectionRequest) GetUSER_FILTEROk() (*string, bool)`

GetUSER_FILTEROk returns a tuple with the USER_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_FILTER

`func (o *TestConnectionRequest) SetUSER_FILTER(v string)`

SetUSER_FILTER sets USER_FILTER field to given value.

### HasUSER_FILTER

`func (o *TestConnectionRequest) HasUSER_FILTER() bool`

HasUSER_FILTER returns a boolean if a field has been set.

### GetUSER_IMPORT_MAPPING

`func (o *TestConnectionRequest) GetUSER_IMPORT_MAPPING() string`

GetUSER_IMPORT_MAPPING returns the USER_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetUSER_IMPORT_MAPPINGOk

`func (o *TestConnectionRequest) GetUSER_IMPORT_MAPPINGOk() (*string, bool)`

GetUSER_IMPORT_MAPPINGOk returns a tuple with the USER_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_MAPPING

`func (o *TestConnectionRequest) SetUSER_IMPORT_MAPPING(v string)`

SetUSER_IMPORT_MAPPING sets USER_IMPORT_MAPPING field to given value.

### HasUSER_IMPORT_MAPPING

`func (o *TestConnectionRequest) HasUSER_IMPORT_MAPPING() bool`

HasUSER_IMPORT_MAPPING returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_MAPPING

`func (o *TestConnectionRequest) GetACCOUNT_IMPORT_MAPPING() string`

GetACCOUNT_IMPORT_MAPPING returns the ACCOUNT_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_MAPPINGOk

`func (o *TestConnectionRequest) GetACCOUNT_IMPORT_MAPPINGOk() (*string, bool)`

GetACCOUNT_IMPORT_MAPPINGOk returns a tuple with the ACCOUNT_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_MAPPING

`func (o *TestConnectionRequest) SetACCOUNT_IMPORT_MAPPING(v string)`

SetACCOUNT_IMPORT_MAPPING sets ACCOUNT_IMPORT_MAPPING field to given value.

### HasACCOUNT_IMPORT_MAPPING

`func (o *TestConnectionRequest) HasACCOUNT_IMPORT_MAPPING() bool`

HasACCOUNT_IMPORT_MAPPING returns a boolean if a field has been set.

### GetORGANIZATION_FILTER

`func (o *TestConnectionRequest) GetORGANIZATION_FILTER() string`

GetORGANIZATION_FILTER returns the ORGANIZATION_FILTER field if non-nil, zero value otherwise.

### GetORGANIZATION_FILTEROk

`func (o *TestConnectionRequest) GetORGANIZATION_FILTEROk() (*string, bool)`

GetORGANIZATION_FILTEROk returns a tuple with the ORGANIZATION_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_FILTER

`func (o *TestConnectionRequest) SetORGANIZATION_FILTER(v string)`

SetORGANIZATION_FILTER sets ORGANIZATION_FILTER field to given value.

### HasORGANIZATION_FILTER

`func (o *TestConnectionRequest) HasORGANIZATION_FILTER() bool`

HasORGANIZATION_FILTER returns a boolean if a field has been set.

### GetSCOPE

`func (o *TestConnectionRequest) GetSCOPE() string`

GetSCOPE returns the SCOPE field if non-nil, zero value otherwise.

### GetSCOPEOk

`func (o *TestConnectionRequest) GetSCOPEOk() (*string, bool)`

GetSCOPEOk returns a tuple with the SCOPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCOPE

`func (o *TestConnectionRequest) SetSCOPE(v string)`

SetSCOPE sets SCOPE field to given value.

### HasSCOPE

`func (o *TestConnectionRequest) HasSCOPE() bool`

HasSCOPE returns a boolean if a field has been set.

### GetUSERS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) GetUSERS_LAST_IMPORT_TIME() string`

GetUSERS_LAST_IMPORT_TIME returns the USERS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetUSERS_LAST_IMPORT_TIMEOk

`func (o *TestConnectionRequest) GetUSERS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetUSERS_LAST_IMPORT_TIMEOk returns a tuple with the USERS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) SetUSERS_LAST_IMPORT_TIME(v string)`

SetUSERS_LAST_IMPORT_TIME sets USERS_LAST_IMPORT_TIME field to given value.

### HasUSERS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) HasUSERS_LAST_IMPORT_TIME() bool`

HasUSERS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetACCOUNTS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) GetACCOUNTS_LAST_IMPORT_TIME() string`

GetACCOUNTS_LAST_IMPORT_TIME returns the ACCOUNTS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCOUNTS_LAST_IMPORT_TIMEOk

`func (o *TestConnectionRequest) GetACCOUNTS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCOUNTS_LAST_IMPORT_TIMEOk returns a tuple with the ACCOUNTS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) SetACCOUNTS_LAST_IMPORT_TIME(v string)`

SetACCOUNTS_LAST_IMPORT_TIME sets ACCOUNTS_LAST_IMPORT_TIME field to given value.

### HasACCOUNTS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) HasACCOUNTS_LAST_IMPORT_TIME() bool`

HasACCOUNTS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetACCESS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) GetACCESS_LAST_IMPORT_TIME() string`

GetACCESS_LAST_IMPORT_TIME returns the ACCESS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCESS_LAST_IMPORT_TIMEOk

`func (o *TestConnectionRequest) GetACCESS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCESS_LAST_IMPORT_TIMEOk returns a tuple with the ACCESS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) SetACCESS_LAST_IMPORT_TIME(v string)`

SetACCESS_LAST_IMPORT_TIME sets ACCESS_LAST_IMPORT_TIME field to given value.

### HasACCESS_LAST_IMPORT_TIME

`func (o *TestConnectionRequest) HasACCESS_LAST_IMPORT_TIME() bool`

HasACCESS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetBASE_URL

`func (o *TestConnectionRequest) GetBASE_URL() string`

GetBASE_URL returns the BASE_URL field if non-nil, zero value otherwise.

### GetBASE_URLOk

`func (o *TestConnectionRequest) GetBASE_URLOk() (*string, bool)`

GetBASE_URLOk returns a tuple with the BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE_URL

`func (o *TestConnectionRequest) SetBASE_URL(v string)`

SetBASE_URL sets BASE_URL field to given value.

### HasBASE_URL

`func (o *TestConnectionRequest) HasBASE_URL() bool`

HasBASE_URL returns a boolean if a field has been set.

### GetAPI_VERSION

`func (o *TestConnectionRequest) GetAPI_VERSION() string`

GetAPI_VERSION returns the API_VERSION field if non-nil, zero value otherwise.

### GetAPI_VERSIONOk

`func (o *TestConnectionRequest) GetAPI_VERSIONOk() (*string, bool)`

GetAPI_VERSIONOk returns a tuple with the API_VERSION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_VERSION

`func (o *TestConnectionRequest) SetAPI_VERSION(v string)`

SetAPI_VERSION sets API_VERSION field to given value.

### HasAPI_VERSION

`func (o *TestConnectionRequest) HasAPI_VERSION() bool`

HasAPI_VERSION returns a boolean if a field has been set.

### GetTENANT_NAME

`func (o *TestConnectionRequest) GetTENANT_NAME() string`

GetTENANT_NAME returns the TENANT_NAME field if non-nil, zero value otherwise.

### GetTENANT_NAMEOk

`func (o *TestConnectionRequest) GetTENANT_NAMEOk() (*string, bool)`

GetTENANT_NAMEOk returns a tuple with the TENANT_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_NAME

`func (o *TestConnectionRequest) SetTENANT_NAME(v string)`

SetTENANT_NAME sets TENANT_NAME field to given value.

### HasTENANT_NAME

`func (o *TestConnectionRequest) HasTENANT_NAME() bool`

HasTENANT_NAME returns a boolean if a field has been set.

### GetREPORT_OWNER

`func (o *TestConnectionRequest) GetREPORT_OWNER() string`

GetREPORT_OWNER returns the REPORT_OWNER field if non-nil, zero value otherwise.

### GetREPORT_OWNEROk

`func (o *TestConnectionRequest) GetREPORT_OWNEROk() (*string, bool)`

GetREPORT_OWNEROk returns a tuple with the REPORT_OWNER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREPORT_OWNER

`func (o *TestConnectionRequest) SetREPORT_OWNER(v string)`

SetREPORT_OWNER sets REPORT_OWNER field to given value.

### HasREPORT_OWNER

`func (o *TestConnectionRequest) HasREPORT_OWNER() bool`

HasREPORT_OWNER returns a boolean if a field has been set.

### GetUSE_OAUTH

`func (o *TestConnectionRequest) GetUSE_OAUTH() string`

GetUSE_OAUTH returns the USE_OAUTH field if non-nil, zero value otherwise.

### GetUSE_OAUTHOk

`func (o *TestConnectionRequest) GetUSE_OAUTHOk() (*string, bool)`

GetUSE_OAUTHOk returns a tuple with the USE_OAUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_OAUTH

`func (o *TestConnectionRequest) SetUSE_OAUTH(v string)`

SetUSE_OAUTH sets USE_OAUTH field to given value.


### GetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *TestConnectionRequest) GetINCLUDE_REFERENCE_DESCRIPTORS() string`

GetINCLUDE_REFERENCE_DESCRIPTORS returns the INCLUDE_REFERENCE_DESCRIPTORS field if non-nil, zero value otherwise.

### GetINCLUDE_REFERENCE_DESCRIPTORSOk

`func (o *TestConnectionRequest) GetINCLUDE_REFERENCE_DESCRIPTORSOk() (*string, bool)`

GetINCLUDE_REFERENCE_DESCRIPTORSOk returns a tuple with the INCLUDE_REFERENCE_DESCRIPTORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *TestConnectionRequest) SetINCLUDE_REFERENCE_DESCRIPTORS(v string)`

SetINCLUDE_REFERENCE_DESCRIPTORS sets INCLUDE_REFERENCE_DESCRIPTORS field to given value.

### HasINCLUDE_REFERENCE_DESCRIPTORS

`func (o *TestConnectionRequest) HasINCLUDE_REFERENCE_DESCRIPTORS() bool`

HasINCLUDE_REFERENCE_DESCRIPTORS returns a boolean if a field has been set.

### GetUSE_ENHANCED_ORGROLE

`func (o *TestConnectionRequest) GetUSE_ENHANCED_ORGROLE() string`

GetUSE_ENHANCED_ORGROLE returns the USE_ENHANCED_ORGROLE field if non-nil, zero value otherwise.

### GetUSE_ENHANCED_ORGROLEOk

`func (o *TestConnectionRequest) GetUSE_ENHANCED_ORGROLEOk() (*string, bool)`

GetUSE_ENHANCED_ORGROLEOk returns a tuple with the USE_ENHANCED_ORGROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_ENHANCED_ORGROLE

`func (o *TestConnectionRequest) SetUSE_ENHANCED_ORGROLE(v string)`

SetUSE_ENHANCED_ORGROLE sets USE_ENHANCED_ORGROLE field to given value.

### HasUSE_ENHANCED_ORGROLE

`func (o *TestConnectionRequest) HasUSE_ENHANCED_ORGROLE() bool`

HasUSE_ENHANCED_ORGROLE returns a boolean if a field has been set.

### GetUSEX509AUTHFORSOAP

`func (o *TestConnectionRequest) GetUSEX509AUTHFORSOAP() string`

GetUSEX509AUTHFORSOAP returns the USEX509AUTHFORSOAP field if non-nil, zero value otherwise.

### GetUSEX509AUTHFORSOAPOk

`func (o *TestConnectionRequest) GetUSEX509AUTHFORSOAPOk() (*string, bool)`

GetUSEX509AUTHFORSOAPOk returns a tuple with the USEX509AUTHFORSOAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSEX509AUTHFORSOAP

`func (o *TestConnectionRequest) SetUSEX509AUTHFORSOAP(v string)`

SetUSEX509AUTHFORSOAP sets USEX509AUTHFORSOAP field to given value.

### HasUSEX509AUTHFORSOAP

`func (o *TestConnectionRequest) HasUSEX509AUTHFORSOAP() bool`

HasUSEX509AUTHFORSOAP returns a boolean if a field has been set.

### GetX509KEY

`func (o *TestConnectionRequest) GetX509KEY() string`

GetX509KEY returns the X509KEY field if non-nil, zero value otherwise.

### GetX509KEYOk

`func (o *TestConnectionRequest) GetX509KEYOk() (*string, bool)`

GetX509KEYOk returns a tuple with the X509KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509KEY

`func (o *TestConnectionRequest) SetX509KEY(v string)`

SetX509KEY sets X509KEY field to given value.

### HasX509KEY

`func (o *TestConnectionRequest) HasX509KEY() bool`

HasX509KEY returns a boolean if a field has been set.

### GetX509CERT

`func (o *TestConnectionRequest) GetX509CERT() string`

GetX509CERT returns the X509CERT field if non-nil, zero value otherwise.

### GetX509CERTOk

`func (o *TestConnectionRequest) GetX509CERTOk() (*string, bool)`

GetX509CERTOk returns a tuple with the X509CERT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509CERT

`func (o *TestConnectionRequest) SetX509CERT(v string)`

SetX509CERT sets X509CERT field to given value.

### HasX509CERT

`func (o *TestConnectionRequest) HasX509CERT() bool`

HasX509CERT returns a boolean if a field has been set.

### GetUSER_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) GetUSER_IMPORT_PAYLOAD() string`

GetUSER_IMPORT_PAYLOAD returns the USER_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetUSER_IMPORT_PAYLOADOk

`func (o *TestConnectionRequest) GetUSER_IMPORT_PAYLOADOk() (*string, bool)`

GetUSER_IMPORT_PAYLOADOk returns a tuple with the USER_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) SetUSER_IMPORT_PAYLOAD(v string)`

SetUSER_IMPORT_PAYLOAD sets USER_IMPORT_PAYLOAD field to given value.

### HasUSER_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) HasUSER_IMPORT_PAYLOAD() bool`

HasUSER_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) GetACCOUNT_IMPORT_PAYLOAD() string`

GetACCOUNT_IMPORT_PAYLOAD returns the ACCOUNT_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_PAYLOADOk

`func (o *TestConnectionRequest) GetACCOUNT_IMPORT_PAYLOADOk() (*string, bool)`

GetACCOUNT_IMPORT_PAYLOADOk returns a tuple with the ACCOUNT_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) SetACCOUNT_IMPORT_PAYLOAD(v string)`

SetACCOUNT_IMPORT_PAYLOAD sets ACCOUNT_IMPORT_PAYLOAD field to given value.

### HasACCOUNT_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) HasACCOUNT_IMPORT_PAYLOAD() bool`

HasACCOUNT_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetACCESS_IMPORT_LIST

`func (o *TestConnectionRequest) GetACCESS_IMPORT_LIST() string`

GetACCESS_IMPORT_LIST returns the ACCESS_IMPORT_LIST field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_LISTOk

`func (o *TestConnectionRequest) GetACCESS_IMPORT_LISTOk() (*string, bool)`

GetACCESS_IMPORT_LISTOk returns a tuple with the ACCESS_IMPORT_LIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_LIST

`func (o *TestConnectionRequest) SetACCESS_IMPORT_LIST(v string)`

SetACCESS_IMPORT_LIST sets ACCESS_IMPORT_LIST field to given value.

### HasACCESS_IMPORT_LIST

`func (o *TestConnectionRequest) HasACCESS_IMPORT_LIST() bool`

HasACCESS_IMPORT_LIST returns a boolean if a field has been set.

### GetRAAS_MAPPING_JSON

`func (o *TestConnectionRequest) GetRAAS_MAPPING_JSON() string`

GetRAAS_MAPPING_JSON returns the RAAS_MAPPING_JSON field if non-nil, zero value otherwise.

### GetRAAS_MAPPING_JSONOk

`func (o *TestConnectionRequest) GetRAAS_MAPPING_JSONOk() (*string, bool)`

GetRAAS_MAPPING_JSONOk returns a tuple with the RAAS_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAAS_MAPPING_JSON

`func (o *TestConnectionRequest) SetRAAS_MAPPING_JSON(v string)`

SetRAAS_MAPPING_JSON sets RAAS_MAPPING_JSON field to given value.

### HasRAAS_MAPPING_JSON

`func (o *TestConnectionRequest) HasRAAS_MAPPING_JSON() bool`

HasRAAS_MAPPING_JSON returns a boolean if a field has been set.

### GetACCESS_IMPORT_MAPPING

`func (o *TestConnectionRequest) GetACCESS_IMPORT_MAPPING() string`

GetACCESS_IMPORT_MAPPING returns the ACCESS_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_MAPPINGOk

`func (o *TestConnectionRequest) GetACCESS_IMPORT_MAPPINGOk() (*string, bool)`

GetACCESS_IMPORT_MAPPINGOk returns a tuple with the ACCESS_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_MAPPING

`func (o *TestConnectionRequest) SetACCESS_IMPORT_MAPPING(v string)`

SetACCESS_IMPORT_MAPPING sets ACCESS_IMPORT_MAPPING field to given value.

### HasACCESS_IMPORT_MAPPING

`func (o *TestConnectionRequest) HasACCESS_IMPORT_MAPPING() bool`

HasACCESS_IMPORT_MAPPING returns a boolean if a field has been set.

### GetORGROLE_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) GetORGROLE_IMPORT_PAYLOAD() string`

GetORGROLE_IMPORT_PAYLOAD returns the ORGROLE_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetORGROLE_IMPORT_PAYLOADOk

`func (o *TestConnectionRequest) GetORGROLE_IMPORT_PAYLOADOk() (*string, bool)`

GetORGROLE_IMPORT_PAYLOADOk returns a tuple with the ORGROLE_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGROLE_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) SetORGROLE_IMPORT_PAYLOAD(v string)`

SetORGROLE_IMPORT_PAYLOAD sets ORGROLE_IMPORT_PAYLOAD field to given value.

### HasORGROLE_IMPORT_PAYLOAD

`func (o *TestConnectionRequest) HasORGROLE_IMPORT_PAYLOAD() bool`

HasORGROLE_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetCREATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) GetCREATE_ACCOUNT_PAYLOAD() string`

GetCREATE_ACCOUNT_PAYLOAD returns the CREATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetCREATE_ACCOUNT_PAYLOADOk

`func (o *TestConnectionRequest) GetCREATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetCREATE_ACCOUNT_PAYLOADOk returns a tuple with the CREATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) SetCREATE_ACCOUNT_PAYLOAD(v string)`

SetCREATE_ACCOUNT_PAYLOAD sets CREATE_ACCOUNT_PAYLOAD field to given value.

### HasCREATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) HasCREATE_ACCOUNT_PAYLOAD() bool`

HasCREATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetUPDATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) GetUPDATE_ACCOUNT_PAYLOAD() string`

GetUPDATE_ACCOUNT_PAYLOAD returns the UPDATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_ACCOUNT_PAYLOADOk

`func (o *TestConnectionRequest) GetUPDATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetUPDATE_ACCOUNT_PAYLOADOk returns a tuple with the UPDATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) SetUPDATE_ACCOUNT_PAYLOAD(v string)`

SetUPDATE_ACCOUNT_PAYLOAD sets UPDATE_ACCOUNT_PAYLOAD field to given value.

### HasUPDATE_ACCOUNT_PAYLOAD

`func (o *TestConnectionRequest) HasUPDATE_ACCOUNT_PAYLOAD() bool`

HasUPDATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetUPDATE_USER_PAYLOAD

`func (o *TestConnectionRequest) GetUPDATE_USER_PAYLOAD() string`

GetUPDATE_USER_PAYLOAD returns the UPDATE_USER_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_USER_PAYLOADOk

`func (o *TestConnectionRequest) GetUPDATE_USER_PAYLOADOk() (*string, bool)`

GetUPDATE_USER_PAYLOADOk returns a tuple with the UPDATE_USER_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USER_PAYLOAD

`func (o *TestConnectionRequest) SetUPDATE_USER_PAYLOAD(v string)`

SetUPDATE_USER_PAYLOAD sets UPDATE_USER_PAYLOAD field to given value.

### HasUPDATE_USER_PAYLOAD

`func (o *TestConnectionRequest) HasUPDATE_USER_PAYLOAD() bool`

HasUPDATE_USER_PAYLOAD returns a boolean if a field has been set.

### GetASSIGN_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) GetASSIGN_ORGROLE_PAYLOAD() string`

GetASSIGN_ORGROLE_PAYLOAD returns the ASSIGN_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetASSIGN_ORGROLE_PAYLOADOk

`func (o *TestConnectionRequest) GetASSIGN_ORGROLE_PAYLOADOk() (*string, bool)`

GetASSIGN_ORGROLE_PAYLOADOk returns a tuple with the ASSIGN_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASSIGN_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) SetASSIGN_ORGROLE_PAYLOAD(v string)`

SetASSIGN_ORGROLE_PAYLOAD sets ASSIGN_ORGROLE_PAYLOAD field to given value.

### HasASSIGN_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) HasASSIGN_ORGROLE_PAYLOAD() bool`

HasASSIGN_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetREMOVE_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) GetREMOVE_ORGROLE_PAYLOAD() string`

GetREMOVE_ORGROLE_PAYLOAD returns the REMOVE_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetREMOVE_ORGROLE_PAYLOADOk

`func (o *TestConnectionRequest) GetREMOVE_ORGROLE_PAYLOADOk() (*string, bool)`

GetREMOVE_ORGROLE_PAYLOADOk returns a tuple with the REMOVE_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) SetREMOVE_ORGROLE_PAYLOAD(v string)`

SetREMOVE_ORGROLE_PAYLOAD sets REMOVE_ORGROLE_PAYLOAD field to given value.

### HasREMOVE_ORGROLE_PAYLOAD

`func (o *TestConnectionRequest) HasREMOVE_ORGROLE_PAYLOAD() bool`

HasREMOVE_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetSTATUS_KEY_JSON

`func (o *TestConnectionRequest) GetSTATUS_KEY_JSON() string`

GetSTATUS_KEY_JSON returns the STATUS_KEY_JSON field if non-nil, zero value otherwise.

### GetSTATUS_KEY_JSONOk

`func (o *TestConnectionRequest) GetSTATUS_KEY_JSONOk() (*string, bool)`

GetSTATUS_KEY_JSONOk returns a tuple with the STATUS_KEY_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_KEY_JSON

`func (o *TestConnectionRequest) SetSTATUS_KEY_JSON(v string)`

SetSTATUS_KEY_JSON sets STATUS_KEY_JSON field to given value.

### HasSTATUS_KEY_JSON

`func (o *TestConnectionRequest) HasSTATUS_KEY_JSON() bool`

HasSTATUS_KEY_JSON returns a boolean if a field has been set.

### GetUSERATTRIBUTEJSON

`func (o *TestConnectionRequest) GetUSERATTRIBUTEJSON() string`

GetUSERATTRIBUTEJSON returns the USERATTRIBUTEJSON field if non-nil, zero value otherwise.

### GetUSERATTRIBUTEJSONOk

`func (o *TestConnectionRequest) GetUSERATTRIBUTEJSONOk() (*string, bool)`

GetUSERATTRIBUTEJSONOk returns a tuple with the USERATTRIBUTEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERATTRIBUTEJSON

`func (o *TestConnectionRequest) SetUSERATTRIBUTEJSON(v string)`

SetUSERATTRIBUTEJSON sets USERATTRIBUTEJSON field to given value.

### HasUSERATTRIBUTEJSON

`func (o *TestConnectionRequest) HasUSERATTRIBUTEJSON() bool`

HasUSERATTRIBUTEJSON returns a boolean if a field has been set.

### GetCUSTOM_CONFIG

`func (o *TestConnectionRequest) GetCUSTOM_CONFIG() string`

GetCUSTOM_CONFIG returns the CUSTOM_CONFIG field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIGOk

`func (o *TestConnectionRequest) GetCUSTOM_CONFIGOk() (*string, bool)`

GetCUSTOM_CONFIGOk returns a tuple with the CUSTOM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG

`func (o *TestConnectionRequest) SetCUSTOM_CONFIG(v string)`

SetCUSTOM_CONFIG sets CUSTOM_CONFIG field to given value.

### HasCUSTOM_CONFIG

`func (o *TestConnectionRequest) HasCUSTOM_CONFIG() bool`

HasCUSTOM_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


