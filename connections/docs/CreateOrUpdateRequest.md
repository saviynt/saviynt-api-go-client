# CreateOrUpdateRequest

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
**URL** | **string** | Host Name for connection | 
**USERNAME** | **string** | Username for connection | 
**PASSWORD** | **string** | Password for connection | 
**CONNECTION_URL** | **string** | ADSI remote agent Connection URL | 
**PROVISIONING_URL** | Pointer to **string** | ADSI remote agent Provisioning URL | [optional] 
**FORESTLIST** | **string** | Forest List (Comma Separated) which we need to manage | 
**DEFAULT_USER_ROLE** | Pointer to **string** | Default SAV Role to be assigned to all the new users that gets imported via User Import | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Update an existing Account,Objects Exposed-(randomPassword,task,user,accountName,role,updatetaskuser,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**SEARCHFILTER** | Pointer to **string** | Account Search Filter to specify the starting point of the directory from where the accounts needs to be imported. You can have multiple BaseDNs here separated by ###. | [optional] 
**OBJECTFILTER** | Pointer to **string** | Object Filter is used to filter the objects that will be returned.This filter will be same for all domains. | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** | Map EIC and AD attributes for account import (AD attributes must be in lower case) | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Applicable for Accounts full import only.If this config isdefined with status/threshold values, it will take precedence over account_not_in_file_action defined in ACCOUNTIMPORT xml.If this config is defined with only correlateInactiveAccounts, then account_not_in_file_action will used as normal.The attributes of statusAndThresholdConfig json are:statusColumn: Property in saviynt which stores the status of target system.activeStatus:All possible values that denotes the Active status of the target system.accountThresholdValue: No. of accounts to be deleted in Saviynt &gt;&#x3D; accountThresholdValue, it performs NO ACTION, else it disables the accounts.inactivateAccountsNotInFile: If true,accounts not in file are marked as Inactive. If false, accounts not in file are marked as SUSPENDED FROM IMPORT SERVICE.CorrelateInactiveAccounts: If true, correlates disabled accounts as well with the users. | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | Account attribute that contains group membership | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** | Map EIC and AD attributes for user import (AD attributes must be in lower case) | [optional] 
**GroupSearchBaseDN** | Pointer to **string** | Group Search Filter to specify the starting point of the directory from where the groups needs to be imported. You can have multiple BaseDNs here separated by ###. | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** | Evaluate the uniqueness of an attribute | [optional] 
**STATUSKEYJSON** | Pointer to **string** | JSON configuration to specify Users status | [optional] 
**GroupImportMapping** | Pointer to **string** | Map AD group attribute to EIC entitlement attribute for import | [optional] 
**ImportNestedMembership** | Pointer to **string** | Specify if you want the connector to import all indirect or nested membership of an account or a group during access import | [optional] 
**PAGE_SIZE** | Pointer to **string** | Specify the number of objects to return in a page for each import request from Workday | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** | Rule to generate account name. | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Create the New Account,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ) | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Update an existing Account,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Enable,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Disable Account,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**REMOVEACCOUNTJSON** | Pointer to **string** | Specify the actions to be performed for deleting an account. | [optional] 
**ADDACCESSJSON** | Pointer to **string** | Configuration to ADD Access (cross domain/forest group membership) to an account. | [optional] 
**REMOVEACCESSJSON** | Pointer to **string** | Configuration to REMOVE Access (cross domain/forest group membership) to an account. | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** | Configuration to Reset and Change Password. | [optional] 
**CREATEGROUPJSON** | Pointer to **string** | Configuration to Create a Group | [optional] 
**UPDATEGROUPJSON** | Pointer to **string** | Configuration to Update a Group | [optional] 
**REMOVEGROUPJSON** | Pointer to **string** | Configuration to Delete a Group | [optional] 
**ADDACCESSENTITLEMENTJSON** | Pointer to **string** | Configuration to Add nested group hierarchy | [optional] 
**CUSTOMCONFIGJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCESSENTITLEMENTJSON** | Pointer to **string** | Configuration to Remove nested group hierarchy | [optional] 
**CREATESERVICEACCOUNTJSON** | Pointer to **string** | Specify the Field Value which will be used to Create the New Service Account. | [optional] 
**UPDATESERVICEACCOUNTJSON** | Pointer to **string** | Specify the Field Value which will be used to update the existing Service Account. | [optional] 
**REMOVESERVICEACCOUNTJSON** | Pointer to **string** | Specify the actions to be performed while deleting a service account. | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Property for MODIFYUSERDATAJSON | [optional] 
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
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password | [optional] 
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
**DRIVERNAME** | **string** | Driver name for the connection | 
**CONNECTIONPROPERTIES** | Pointer to **string** | Properties that needs to be added when connecting to the database | [optional] 
**GRANTACCESSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to provide acccess,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ).  | [optional] 
**REVOKEACCESSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to revoke access,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**CHANGEPASSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to change password,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ).  | [optional] 
**DELETEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to delete an account,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ACCOUNTEXISTSJSON** | Pointer to **string** | JSON to specify the Query which will be used to check whether an account exists,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ACCOUNTSIMPORT** | Pointer to **string** | Accounts Import XML file content | [optional] 
**ENTITLEMENTVALUEIMPORT** | Pointer to **string** | Entitlement Value Import XML file content | [optional] 
**ROLEOWNERIMPORT** | Pointer to **string** | Role Owner Import XML file contentT | [optional] 
**ROLESIMPORT** | Pointer to **string** | Roles Import XML file content | [optional] 
**SYSTEMIMPORT** | Pointer to **string** | System Import XML file content | [optional] 
**USERIMPORT** | Pointer to **string** | User Import XML file content | [optional] 
**MAX_PAGINATION_SIZE** | Pointer to **string** | Defines the max number of records from the target to be processed in each page | [optional] 
**CLI_COMMAND_JSON** | Pointer to **string** | JSON to specify the commands which can be executed in target server. | [optional] 

## Methods

### NewCreateOrUpdateRequest

`func NewCreateOrUpdateRequest(connectionName string, connectiontype string, uRL string, uSERNAME string, pASSWORD string, cONNECTIONURL string, fORESTLIST string, cLIENTID string, cLIENTSECRET string, bASEURL string, tENANTID string, lOGINURL string, uSEOAUTH string, dRIVERNAME string, ) *CreateOrUpdateRequest`

NewCreateOrUpdateRequest instantiates a new CreateOrUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateRequestWithDefaults

`func NewCreateOrUpdateRequestWithDefaults() *CreateOrUpdateRequest`

NewCreateOrUpdateRequestWithDefaults instantiates a new CreateOrUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *CreateOrUpdateRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CreateOrUpdateRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CreateOrUpdateRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetConnectiontype

`func (o *CreateOrUpdateRequest) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *CreateOrUpdateRequest) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *CreateOrUpdateRequest) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.


### GetDescription

`func (o *CreateOrUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *CreateOrUpdateRequest) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *CreateOrUpdateRequest) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *CreateOrUpdateRequest) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *CreateOrUpdateRequest) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetEmailTemplate

`func (o *CreateOrUpdateRequest) GetEmailTemplate() string`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *CreateOrUpdateRequest) GetEmailTemplateOk() (*string, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *CreateOrUpdateRequest) SetEmailTemplate(v string)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *CreateOrUpdateRequest) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetSslCertificate

`func (o *CreateOrUpdateRequest) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *CreateOrUpdateRequest) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *CreateOrUpdateRequest) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *CreateOrUpdateRequest) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetVaultConnection

`func (o *CreateOrUpdateRequest) GetVaultConnection() string`

GetVaultConnection returns the VaultConnection field if non-nil, zero value otherwise.

### GetVaultConnectionOk

`func (o *CreateOrUpdateRequest) GetVaultConnectionOk() (*string, bool)`

GetVaultConnectionOk returns a tuple with the VaultConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConnection

`func (o *CreateOrUpdateRequest) SetVaultConnection(v string)`

SetVaultConnection sets VaultConnection field to given value.

### HasVaultConnection

`func (o *CreateOrUpdateRequest) HasVaultConnection() bool`

HasVaultConnection returns a boolean if a field has been set.

### GetVaultConfiguration

`func (o *CreateOrUpdateRequest) GetVaultConfiguration() string`

GetVaultConfiguration returns the VaultConfiguration field if non-nil, zero value otherwise.

### GetVaultConfigurationOk

`func (o *CreateOrUpdateRequest) GetVaultConfigurationOk() (*string, bool)`

GetVaultConfigurationOk returns a tuple with the VaultConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConfiguration

`func (o *CreateOrUpdateRequest) SetVaultConfiguration(v string)`

SetVaultConfiguration sets VaultConfiguration field to given value.

### HasVaultConfiguration

`func (o *CreateOrUpdateRequest) HasVaultConfiguration() bool`

HasVaultConfiguration returns a boolean if a field has been set.

### GetSaveinvault

`func (o *CreateOrUpdateRequest) GetSaveinvault() string`

GetSaveinvault returns the Saveinvault field if non-nil, zero value otherwise.

### GetSaveinvaultOk

`func (o *CreateOrUpdateRequest) GetSaveinvaultOk() (*string, bool)`

GetSaveinvaultOk returns a tuple with the Saveinvault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveinvault

`func (o *CreateOrUpdateRequest) SetSaveinvault(v string)`

SetSaveinvault sets Saveinvault field to given value.

### HasSaveinvault

`func (o *CreateOrUpdateRequest) HasSaveinvault() bool`

HasSaveinvault returns a boolean if a field has been set.

### GetURL

`func (o *CreateOrUpdateRequest) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *CreateOrUpdateRequest) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *CreateOrUpdateRequest) SetURL(v string)`

SetURL sets URL field to given value.


### GetUSERNAME

`func (o *CreateOrUpdateRequest) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *CreateOrUpdateRequest) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *CreateOrUpdateRequest) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetPASSWORD

`func (o *CreateOrUpdateRequest) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *CreateOrUpdateRequest) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *CreateOrUpdateRequest) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.


### GetCONNECTION_URL

`func (o *CreateOrUpdateRequest) GetCONNECTION_URL() string`

GetCONNECTION_URL returns the CONNECTION_URL field if non-nil, zero value otherwise.

### GetCONNECTION_URLOk

`func (o *CreateOrUpdateRequest) GetCONNECTION_URLOk() (*string, bool)`

GetCONNECTION_URLOk returns a tuple with the CONNECTION_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTION_URL

`func (o *CreateOrUpdateRequest) SetCONNECTION_URL(v string)`

SetCONNECTION_URL sets CONNECTION_URL field to given value.


### GetPROVISIONING_URL

`func (o *CreateOrUpdateRequest) GetPROVISIONING_URL() string`

GetPROVISIONING_URL returns the PROVISIONING_URL field if non-nil, zero value otherwise.

### GetPROVISIONING_URLOk

`func (o *CreateOrUpdateRequest) GetPROVISIONING_URLOk() (*string, bool)`

GetPROVISIONING_URLOk returns a tuple with the PROVISIONING_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISIONING_URL

`func (o *CreateOrUpdateRequest) SetPROVISIONING_URL(v string)`

SetPROVISIONING_URL sets PROVISIONING_URL field to given value.

### HasPROVISIONING_URL

`func (o *CreateOrUpdateRequest) HasPROVISIONING_URL() bool`

HasPROVISIONING_URL returns a boolean if a field has been set.

### GetFORESTLIST

`func (o *CreateOrUpdateRequest) GetFORESTLIST() string`

GetFORESTLIST returns the FORESTLIST field if non-nil, zero value otherwise.

### GetFORESTLISTOk

`func (o *CreateOrUpdateRequest) GetFORESTLISTOk() (*string, bool)`

GetFORESTLISTOk returns a tuple with the FORESTLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFORESTLIST

`func (o *CreateOrUpdateRequest) SetFORESTLIST(v string)`

SetFORESTLIST sets FORESTLIST field to given value.


### GetDEFAULT_USER_ROLE

`func (o *CreateOrUpdateRequest) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *CreateOrUpdateRequest) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *CreateOrUpdateRequest) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *CreateOrUpdateRequest) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *CreateOrUpdateRequest) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *CreateOrUpdateRequest) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *CreateOrUpdateRequest) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *CreateOrUpdateRequest) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *CreateOrUpdateRequest) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *CreateOrUpdateRequest) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *CreateOrUpdateRequest) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *CreateOrUpdateRequest) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *CreateOrUpdateRequest) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *CreateOrUpdateRequest) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *CreateOrUpdateRequest) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *CreateOrUpdateRequest) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *CreateOrUpdateRequest) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *CreateOrUpdateRequest) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *CreateOrUpdateRequest) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *CreateOrUpdateRequest) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *CreateOrUpdateRequest) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateRequest) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *CreateOrUpdateRequest) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateRequest) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *CreateOrUpdateRequest) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *CreateOrUpdateRequest) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *CreateOrUpdateRequest) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *CreateOrUpdateRequest) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *CreateOrUpdateRequest) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *CreateOrUpdateRequest) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *CreateOrUpdateRequest) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *CreateOrUpdateRequest) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *CreateOrUpdateRequest) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *CreateOrUpdateRequest) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *CreateOrUpdateRequest) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *CreateOrUpdateRequest) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *CreateOrUpdateRequest) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *CreateOrUpdateRequest) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *CreateOrUpdateRequest) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *CreateOrUpdateRequest) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *CreateOrUpdateRequest) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *CreateOrUpdateRequest) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *CreateOrUpdateRequest) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *CreateOrUpdateRequest) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *CreateOrUpdateRequest) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *CreateOrUpdateRequest) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *CreateOrUpdateRequest) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetImportNestedMembership

`func (o *CreateOrUpdateRequest) GetImportNestedMembership() string`

GetImportNestedMembership returns the ImportNestedMembership field if non-nil, zero value otherwise.

### GetImportNestedMembershipOk

`func (o *CreateOrUpdateRequest) GetImportNestedMembershipOk() (*string, bool)`

GetImportNestedMembershipOk returns a tuple with the ImportNestedMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNestedMembership

`func (o *CreateOrUpdateRequest) SetImportNestedMembership(v string)`

SetImportNestedMembership sets ImportNestedMembership field to given value.

### HasImportNestedMembership

`func (o *CreateOrUpdateRequest) HasImportNestedMembership() bool`

HasImportNestedMembership returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *CreateOrUpdateRequest) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *CreateOrUpdateRequest) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *CreateOrUpdateRequest) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *CreateOrUpdateRequest) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *CreateOrUpdateRequest) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *CreateOrUpdateRequest) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *CreateOrUpdateRequest) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *CreateOrUpdateRequest) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetREMOVEACCOUNTJSON() string`

GetREMOVEACCOUNTJSON returns the REMOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetREMOVEACCOUNTJSONOk() (*string, bool)`

GetREMOVEACCOUNTJSONOk returns a tuple with the REMOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetREMOVEACCOUNTJSON(v string)`

SetREMOVEACCOUNTJSON sets REMOVEACCOUNTJSON field to given value.

### HasREMOVEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasREMOVEACCOUNTJSON() bool`

HasREMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetADDACCESSJSON

`func (o *CreateOrUpdateRequest) GetADDACCESSJSON() string`

GetADDACCESSJSON returns the ADDACCESSJSON field if non-nil, zero value otherwise.

### GetADDACCESSJSONOk

`func (o *CreateOrUpdateRequest) GetADDACCESSJSONOk() (*string, bool)`

GetADDACCESSJSONOk returns a tuple with the ADDACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSJSON

`func (o *CreateOrUpdateRequest) SetADDACCESSJSON(v string)`

SetADDACCESSJSON sets ADDACCESSJSON field to given value.

### HasADDACCESSJSON

`func (o *CreateOrUpdateRequest) HasADDACCESSJSON() bool`

HasADDACCESSJSON returns a boolean if a field has been set.

### GetREMOVEACCESSJSON

`func (o *CreateOrUpdateRequest) GetREMOVEACCESSJSON() string`

GetREMOVEACCESSJSON returns the REMOVEACCESSJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSJSONOk

`func (o *CreateOrUpdateRequest) GetREMOVEACCESSJSONOk() (*string, bool)`

GetREMOVEACCESSJSONOk returns a tuple with the REMOVEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSJSON

`func (o *CreateOrUpdateRequest) SetREMOVEACCESSJSON(v string)`

SetREMOVEACCESSJSON sets REMOVEACCESSJSON field to given value.

### HasREMOVEACCESSJSON

`func (o *CreateOrUpdateRequest) HasREMOVEACCESSJSON() bool`

HasREMOVEACCESSJSON returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateRequest) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *CreateOrUpdateRequest) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateRequest) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *CreateOrUpdateRequest) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetCREATEGROUPJSON

`func (o *CreateOrUpdateRequest) GetCREATEGROUPJSON() string`

GetCREATEGROUPJSON returns the CREATEGROUPJSON field if non-nil, zero value otherwise.

### GetCREATEGROUPJSONOk

`func (o *CreateOrUpdateRequest) GetCREATEGROUPJSONOk() (*string, bool)`

GetCREATEGROUPJSONOk returns a tuple with the CREATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEGROUPJSON

`func (o *CreateOrUpdateRequest) SetCREATEGROUPJSON(v string)`

SetCREATEGROUPJSON sets CREATEGROUPJSON field to given value.

### HasCREATEGROUPJSON

`func (o *CreateOrUpdateRequest) HasCREATEGROUPJSON() bool`

HasCREATEGROUPJSON returns a boolean if a field has been set.

### GetUPDATEGROUPJSON

`func (o *CreateOrUpdateRequest) GetUPDATEGROUPJSON() string`

GetUPDATEGROUPJSON returns the UPDATEGROUPJSON field if non-nil, zero value otherwise.

### GetUPDATEGROUPJSONOk

`func (o *CreateOrUpdateRequest) GetUPDATEGROUPJSONOk() (*string, bool)`

GetUPDATEGROUPJSONOk returns a tuple with the UPDATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEGROUPJSON

`func (o *CreateOrUpdateRequest) SetUPDATEGROUPJSON(v string)`

SetUPDATEGROUPJSON sets UPDATEGROUPJSON field to given value.

### HasUPDATEGROUPJSON

`func (o *CreateOrUpdateRequest) HasUPDATEGROUPJSON() bool`

HasUPDATEGROUPJSON returns a boolean if a field has been set.

### GetREMOVEGROUPJSON

`func (o *CreateOrUpdateRequest) GetREMOVEGROUPJSON() string`

GetREMOVEGROUPJSON returns the REMOVEGROUPJSON field if non-nil, zero value otherwise.

### GetREMOVEGROUPJSONOk

`func (o *CreateOrUpdateRequest) GetREMOVEGROUPJSONOk() (*string, bool)`

GetREMOVEGROUPJSONOk returns a tuple with the REMOVEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEGROUPJSON

`func (o *CreateOrUpdateRequest) SetREMOVEGROUPJSON(v string)`

SetREMOVEGROUPJSON sets REMOVEGROUPJSON field to given value.

### HasREMOVEGROUPJSON

`func (o *CreateOrUpdateRequest) HasREMOVEGROUPJSON() bool`

HasREMOVEGROUPJSON returns a boolean if a field has been set.

### GetADDACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) GetADDACCESSENTITLEMENTJSON() string`

GetADDACCESSENTITLEMENTJSON returns the ADDACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetADDACCESSENTITLEMENTJSONOk

`func (o *CreateOrUpdateRequest) GetADDACCESSENTITLEMENTJSONOk() (*string, bool)`

GetADDACCESSENTITLEMENTJSONOk returns a tuple with the ADDACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) SetADDACCESSENTITLEMENTJSON(v string)`

SetADDACCESSENTITLEMENTJSON sets ADDACCESSENTITLEMENTJSON field to given value.

### HasADDACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) HasADDACCESSENTITLEMENTJSON() bool`

HasADDACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCUSTOMCONFIGJSON

`func (o *CreateOrUpdateRequest) GetCUSTOMCONFIGJSON() string`

GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field if non-nil, zero value otherwise.

### GetCUSTOMCONFIGJSONOk

`func (o *CreateOrUpdateRequest) GetCUSTOMCONFIGJSONOk() (*string, bool)`

GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMCONFIGJSON

`func (o *CreateOrUpdateRequest) SetCUSTOMCONFIGJSON(v string)`

SetCUSTOMCONFIGJSON sets CUSTOMCONFIGJSON field to given value.

### HasCUSTOMCONFIGJSON

`func (o *CreateOrUpdateRequest) HasCUSTOMCONFIGJSON() bool`

HasCUSTOMCONFIGJSON returns a boolean if a field has been set.

### GetREMOVEACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) GetREMOVEACCESSENTITLEMENTJSON() string`

GetREMOVEACCESSENTITLEMENTJSON returns the REMOVEACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSENTITLEMENTJSONOk

`func (o *CreateOrUpdateRequest) GetREMOVEACCESSENTITLEMENTJSONOk() (*string, bool)`

GetREMOVEACCESSENTITLEMENTJSONOk returns a tuple with the REMOVEACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) SetREMOVEACCESSENTITLEMENTJSON(v string)`

SetREMOVEACCESSENTITLEMENTJSON sets REMOVEACCESSENTITLEMENTJSON field to given value.

### HasREMOVEACCESSENTITLEMENTJSON

`func (o *CreateOrUpdateRequest) HasREMOVEACCESSENTITLEMENTJSON() bool`

HasREMOVEACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCREATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetCREATESERVICEACCOUNTJSON() string`

GetCREATESERVICEACCOUNTJSON returns the CREATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATESERVICEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetCREATESERVICEACCOUNTJSONOk() (*string, bool)`

GetCREATESERVICEACCOUNTJSONOk returns a tuple with the CREATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetCREATESERVICEACCOUNTJSON(v string)`

SetCREATESERVICEACCOUNTJSON sets CREATESERVICEACCOUNTJSON field to given value.

### HasCREATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasCREATESERVICEACCOUNTJSON() bool`

HasCREATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetUPDATESERVICEACCOUNTJSON() string`

GetUPDATESERVICEACCOUNTJSON returns the UPDATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATESERVICEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetUPDATESERVICEACCOUNTJSONOk() (*string, bool)`

GetUPDATESERVICEACCOUNTJSONOk returns a tuple with the UPDATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetUPDATESERVICEACCOUNTJSON(v string)`

SetUPDATESERVICEACCOUNTJSON sets UPDATESERVICEACCOUNTJSON field to given value.

### HasUPDATESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasUPDATESERVICEACCOUNTJSON() bool`

HasUPDATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetREMOVESERVICEACCOUNTJSON() string`

GetREMOVESERVICEACCOUNTJSON returns the REMOVESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVESERVICEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetREMOVESERVICEACCOUNTJSONOk() (*string, bool)`

GetREMOVESERVICEACCOUNTJSONOk returns a tuple with the REMOVESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetREMOVESERVICEACCOUNTJSON(v string)`

SetREMOVESERVICEACCOUNTJSON sets REMOVESERVICEACCOUNTJSON field to given value.

### HasREMOVESERVICEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasREMOVESERVICEACCOUNTJSON() bool`

HasREMOVESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *CreateOrUpdateRequest) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *CreateOrUpdateRequest) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *CreateOrUpdateRequest) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *CreateOrUpdateRequest) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *CreateOrUpdateRequest) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *CreateOrUpdateRequest) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *CreateOrUpdateRequest) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *CreateOrUpdateRequest) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetMESSAGESERVER

`func (o *CreateOrUpdateRequest) GetMESSAGESERVER() string`

GetMESSAGESERVER returns the MESSAGESERVER field if non-nil, zero value otherwise.

### GetMESSAGESERVEROk

`func (o *CreateOrUpdateRequest) GetMESSAGESERVEROk() (*string, bool)`

GetMESSAGESERVEROk returns a tuple with the MESSAGESERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMESSAGESERVER

`func (o *CreateOrUpdateRequest) SetMESSAGESERVER(v string)`

SetMESSAGESERVER sets MESSAGESERVER field to given value.

### HasMESSAGESERVER

`func (o *CreateOrUpdateRequest) HasMESSAGESERVER() bool`

HasMESSAGESERVER returns a boolean if a field has been set.

### GetJCO_ASHOST

`func (o *CreateOrUpdateRequest) GetJCO_ASHOST() string`

GetJCO_ASHOST returns the JCO_ASHOST field if non-nil, zero value otherwise.

### GetJCO_ASHOSTOk

`func (o *CreateOrUpdateRequest) GetJCO_ASHOSTOk() (*string, bool)`

GetJCO_ASHOSTOk returns a tuple with the JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_ASHOST

`func (o *CreateOrUpdateRequest) SetJCO_ASHOST(v string)`

SetJCO_ASHOST sets JCO_ASHOST field to given value.

### HasJCO_ASHOST

`func (o *CreateOrUpdateRequest) HasJCO_ASHOST() bool`

HasJCO_ASHOST returns a boolean if a field has been set.

### GetJCO_SYSNR

`func (o *CreateOrUpdateRequest) GetJCO_SYSNR() string`

GetJCO_SYSNR returns the JCO_SYSNR field if non-nil, zero value otherwise.

### GetJCO_SYSNROk

`func (o *CreateOrUpdateRequest) GetJCO_SYSNROk() (*string, bool)`

GetJCO_SYSNROk returns a tuple with the JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SYSNR

`func (o *CreateOrUpdateRequest) SetJCO_SYSNR(v string)`

SetJCO_SYSNR sets JCO_SYSNR field to given value.

### HasJCO_SYSNR

`func (o *CreateOrUpdateRequest) HasJCO_SYSNR() bool`

HasJCO_SYSNR returns a boolean if a field has been set.

### GetJCO_CLIENT

`func (o *CreateOrUpdateRequest) GetJCO_CLIENT() string`

GetJCO_CLIENT returns the JCO_CLIENT field if non-nil, zero value otherwise.

### GetJCO_CLIENTOk

`func (o *CreateOrUpdateRequest) GetJCO_CLIENTOk() (*string, bool)`

GetJCO_CLIENTOk returns a tuple with the JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_CLIENT

`func (o *CreateOrUpdateRequest) SetJCO_CLIENT(v string)`

SetJCO_CLIENT sets JCO_CLIENT field to given value.

### HasJCO_CLIENT

`func (o *CreateOrUpdateRequest) HasJCO_CLIENT() bool`

HasJCO_CLIENT returns a boolean if a field has been set.

### GetJCO_USER

`func (o *CreateOrUpdateRequest) GetJCO_USER() string`

GetJCO_USER returns the JCO_USER field if non-nil, zero value otherwise.

### GetJCO_USEROk

`func (o *CreateOrUpdateRequest) GetJCO_USEROk() (*string, bool)`

GetJCO_USEROk returns a tuple with the JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_USER

`func (o *CreateOrUpdateRequest) SetJCO_USER(v string)`

SetJCO_USER sets JCO_USER field to given value.

### HasJCO_USER

`func (o *CreateOrUpdateRequest) HasJCO_USER() bool`

HasJCO_USER returns a boolean if a field has been set.

### GetJCO_LANG

`func (o *CreateOrUpdateRequest) GetJCO_LANG() string`

GetJCO_LANG returns the JCO_LANG field if non-nil, zero value otherwise.

### GetJCO_LANGOk

`func (o *CreateOrUpdateRequest) GetJCO_LANGOk() (*string, bool)`

GetJCO_LANGOk returns a tuple with the JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_LANG

`func (o *CreateOrUpdateRequest) SetJCO_LANG(v string)`

SetJCO_LANG sets JCO_LANG field to given value.

### HasJCO_LANG

`func (o *CreateOrUpdateRequest) HasJCO_LANG() bool`

HasJCO_LANG returns a boolean if a field has been set.

### GetJCOR3NAME

`func (o *CreateOrUpdateRequest) GetJCOR3NAME() string`

GetJCOR3NAME returns the JCOR3NAME field if non-nil, zero value otherwise.

### GetJCOR3NAMEOk

`func (o *CreateOrUpdateRequest) GetJCOR3NAMEOk() (*string, bool)`

GetJCOR3NAMEOk returns a tuple with the JCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCOR3NAME

`func (o *CreateOrUpdateRequest) SetJCOR3NAME(v string)`

SetJCOR3NAME sets JCOR3NAME field to given value.

### HasJCOR3NAME

`func (o *CreateOrUpdateRequest) HasJCOR3NAME() bool`

HasJCOR3NAME returns a boolean if a field has been set.

### GetJCO_MSHOST

`func (o *CreateOrUpdateRequest) GetJCO_MSHOST() string`

GetJCO_MSHOST returns the JCO_MSHOST field if non-nil, zero value otherwise.

### GetJCO_MSHOSTOk

`func (o *CreateOrUpdateRequest) GetJCO_MSHOSTOk() (*string, bool)`

GetJCO_MSHOSTOk returns a tuple with the JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSHOST

`func (o *CreateOrUpdateRequest) SetJCO_MSHOST(v string)`

SetJCO_MSHOST sets JCO_MSHOST field to given value.

### HasJCO_MSHOST

`func (o *CreateOrUpdateRequest) HasJCO_MSHOST() bool`

HasJCO_MSHOST returns a boolean if a field has been set.

### GetJCO_MSSERV

`func (o *CreateOrUpdateRequest) GetJCO_MSSERV() string`

GetJCO_MSSERV returns the JCO_MSSERV field if non-nil, zero value otherwise.

### GetJCO_MSSERVOk

`func (o *CreateOrUpdateRequest) GetJCO_MSSERVOk() (*string, bool)`

GetJCO_MSSERVOk returns a tuple with the JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSSERV

`func (o *CreateOrUpdateRequest) SetJCO_MSSERV(v string)`

SetJCO_MSSERV sets JCO_MSSERV field to given value.

### HasJCO_MSSERV

`func (o *CreateOrUpdateRequest) HasJCO_MSSERV() bool`

HasJCO_MSSERV returns a boolean if a field has been set.

### GetJCO_GROUP

`func (o *CreateOrUpdateRequest) GetJCO_GROUP() string`

GetJCO_GROUP returns the JCO_GROUP field if non-nil, zero value otherwise.

### GetJCO_GROUPOk

`func (o *CreateOrUpdateRequest) GetJCO_GROUPOk() (*string, bool)`

GetJCO_GROUPOk returns a tuple with the JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_GROUP

`func (o *CreateOrUpdateRequest) SetJCO_GROUP(v string)`

SetJCO_GROUP sets JCO_GROUP field to given value.

### HasJCO_GROUP

`func (o *CreateOrUpdateRequest) HasJCO_GROUP() bool`

HasJCO_GROUP returns a boolean if a field has been set.

### GetSNC

`func (o *CreateOrUpdateRequest) GetSNC() string`

GetSNC returns the SNC field if non-nil, zero value otherwise.

### GetSNCOk

`func (o *CreateOrUpdateRequest) GetSNCOk() (*string, bool)`

GetSNCOk returns a tuple with the SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNC

`func (o *CreateOrUpdateRequest) SetSNC(v string)`

SetSNC sets SNC field to given value.

### HasSNC

`func (o *CreateOrUpdateRequest) HasSNC() bool`

HasSNC returns a boolean if a field has been set.

### GetJCO_SNC_MODE

`func (o *CreateOrUpdateRequest) GetJCO_SNC_MODE() string`

GetJCO_SNC_MODE returns the JCO_SNC_MODE field if non-nil, zero value otherwise.

### GetJCO_SNC_MODEOk

`func (o *CreateOrUpdateRequest) GetJCO_SNC_MODEOk() (*string, bool)`

GetJCO_SNC_MODEOk returns a tuple with the JCO_SNC_MODE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MODE

`func (o *CreateOrUpdateRequest) SetJCO_SNC_MODE(v string)`

SetJCO_SNC_MODE sets JCO_SNC_MODE field to given value.

### HasJCO_SNC_MODE

`func (o *CreateOrUpdateRequest) HasJCO_SNC_MODE() bool`

HasJCO_SNC_MODE returns a boolean if a field has been set.

### GetJCO_SNC_PARTNERNAME

`func (o *CreateOrUpdateRequest) GetJCO_SNC_PARTNERNAME() string`

GetJCO_SNC_PARTNERNAME returns the JCO_SNC_PARTNERNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_PARTNERNAMEOk

`func (o *CreateOrUpdateRequest) GetJCO_SNC_PARTNERNAMEOk() (*string, bool)`

GetJCO_SNC_PARTNERNAMEOk returns a tuple with the JCO_SNC_PARTNERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_PARTNERNAME

`func (o *CreateOrUpdateRequest) SetJCO_SNC_PARTNERNAME(v string)`

SetJCO_SNC_PARTNERNAME sets JCO_SNC_PARTNERNAME field to given value.

### HasJCO_SNC_PARTNERNAME

`func (o *CreateOrUpdateRequest) HasJCO_SNC_PARTNERNAME() bool`

HasJCO_SNC_PARTNERNAME returns a boolean if a field has been set.

### GetJCO_SNC_MYNAME

`func (o *CreateOrUpdateRequest) GetJCO_SNC_MYNAME() string`

GetJCO_SNC_MYNAME returns the JCO_SNC_MYNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_MYNAMEOk

`func (o *CreateOrUpdateRequest) GetJCO_SNC_MYNAMEOk() (*string, bool)`

GetJCO_SNC_MYNAMEOk returns a tuple with the JCO_SNC_MYNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MYNAME

`func (o *CreateOrUpdateRequest) SetJCO_SNC_MYNAME(v string)`

SetJCO_SNC_MYNAME sets JCO_SNC_MYNAME field to given value.

### HasJCO_SNC_MYNAME

`func (o *CreateOrUpdateRequest) HasJCO_SNC_MYNAME() bool`

HasJCO_SNC_MYNAME returns a boolean if a field has been set.

### GetJCO_SNC_LIBRARY

`func (o *CreateOrUpdateRequest) GetJCO_SNC_LIBRARY() string`

GetJCO_SNC_LIBRARY returns the JCO_SNC_LIBRARY field if non-nil, zero value otherwise.

### GetJCO_SNC_LIBRARYOk

`func (o *CreateOrUpdateRequest) GetJCO_SNC_LIBRARYOk() (*string, bool)`

GetJCO_SNC_LIBRARYOk returns a tuple with the JCO_SNC_LIBRARY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_LIBRARY

`func (o *CreateOrUpdateRequest) SetJCO_SNC_LIBRARY(v string)`

SetJCO_SNC_LIBRARY sets JCO_SNC_LIBRARY field to given value.

### HasJCO_SNC_LIBRARY

`func (o *CreateOrUpdateRequest) HasJCO_SNC_LIBRARY() bool`

HasJCO_SNC_LIBRARY returns a boolean if a field has been set.

### GetJCO_SNC_QOP

`func (o *CreateOrUpdateRequest) GetJCO_SNC_QOP() string`

GetJCO_SNC_QOP returns the JCO_SNC_QOP field if non-nil, zero value otherwise.

### GetJCO_SNC_QOPOk

`func (o *CreateOrUpdateRequest) GetJCO_SNC_QOPOk() (*string, bool)`

GetJCO_SNC_QOPOk returns a tuple with the JCO_SNC_QOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_QOP

`func (o *CreateOrUpdateRequest) SetJCO_SNC_QOP(v string)`

SetJCO_SNC_QOP sets JCO_SNC_QOP field to given value.

### HasJCO_SNC_QOP

`func (o *CreateOrUpdateRequest) HasJCO_SNC_QOP() bool`

HasJCO_SNC_QOP returns a boolean if a field has been set.

### GetTABLES

`func (o *CreateOrUpdateRequest) GetTABLES() string`

GetTABLES returns the TABLES field if non-nil, zero value otherwise.

### GetTABLESOk

`func (o *CreateOrUpdateRequest) GetTABLESOk() (*string, bool)`

GetTABLESOk returns a tuple with the TABLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTABLES

`func (o *CreateOrUpdateRequest) SetTABLES(v string)`

SetTABLES sets TABLES field to given value.

### HasTABLES

`func (o *CreateOrUpdateRequest) HasTABLES() bool`

HasTABLES returns a boolean if a field has been set.

### GetSYSTEMNAME

`func (o *CreateOrUpdateRequest) GetSYSTEMNAME() string`

GetSYSTEMNAME returns the SYSTEMNAME field if non-nil, zero value otherwise.

### GetSYSTEMNAMEOk

`func (o *CreateOrUpdateRequest) GetSYSTEMNAMEOk() (*string, bool)`

GetSYSTEMNAMEOk returns a tuple with the SYSTEMNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMNAME

`func (o *CreateOrUpdateRequest) SetSYSTEMNAME(v string)`

SetSYSTEMNAME sets SYSTEMNAME field to given value.

### HasSYSTEMNAME

`func (o *CreateOrUpdateRequest) HasSYSTEMNAME() bool`

HasSYSTEMNAME returns a boolean if a field has been set.

### GetTERMINATEDUSERGROUP

`func (o *CreateOrUpdateRequest) GetTERMINATEDUSERGROUP() string`

GetTERMINATEDUSERGROUP returns the TERMINATEDUSERGROUP field if non-nil, zero value otherwise.

### GetTERMINATEDUSERGROUPOk

`func (o *CreateOrUpdateRequest) GetTERMINATEDUSERGROUPOk() (*string, bool)`

GetTERMINATEDUSERGROUPOk returns a tuple with the TERMINATEDUSERGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATEDUSERGROUP

`func (o *CreateOrUpdateRequest) SetTERMINATEDUSERGROUP(v string)`

SetTERMINATEDUSERGROUP sets TERMINATEDUSERGROUP field to given value.

### HasTERMINATEDUSERGROUP

`func (o *CreateOrUpdateRequest) HasTERMINATEDUSERGROUP() bool`

HasTERMINATEDUSERGROUP returns a boolean if a field has been set.

### GetTERMINATED_USER_ROLE_ACTION

`func (o *CreateOrUpdateRequest) GetTERMINATED_USER_ROLE_ACTION() string`

GetTERMINATED_USER_ROLE_ACTION returns the TERMINATED_USER_ROLE_ACTION field if non-nil, zero value otherwise.

### GetTERMINATED_USER_ROLE_ACTIONOk

`func (o *CreateOrUpdateRequest) GetTERMINATED_USER_ROLE_ACTIONOk() (*string, bool)`

GetTERMINATED_USER_ROLE_ACTIONOk returns a tuple with the TERMINATED_USER_ROLE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATED_USER_ROLE_ACTION

`func (o *CreateOrUpdateRequest) SetTERMINATED_USER_ROLE_ACTION(v string)`

SetTERMINATED_USER_ROLE_ACTION sets TERMINATED_USER_ROLE_ACTION field to given value.

### HasTERMINATED_USER_ROLE_ACTION

`func (o *CreateOrUpdateRequest) HasTERMINATED_USER_ROLE_ACTION() bool`

HasTERMINATED_USER_ROLE_ACTION returns a boolean if a field has been set.

### GetPROV_JCO_ASHOST

`func (o *CreateOrUpdateRequest) GetPROV_JCO_ASHOST() string`

GetPROV_JCO_ASHOST returns the PROV_JCO_ASHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_ASHOSTOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_ASHOSTOk() (*string, bool)`

GetPROV_JCO_ASHOSTOk returns a tuple with the PROV_JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_ASHOST

`func (o *CreateOrUpdateRequest) SetPROV_JCO_ASHOST(v string)`

SetPROV_JCO_ASHOST sets PROV_JCO_ASHOST field to given value.

### HasPROV_JCO_ASHOST

`func (o *CreateOrUpdateRequest) HasPROV_JCO_ASHOST() bool`

HasPROV_JCO_ASHOST returns a boolean if a field has been set.

### GetPROV_JCO_SYSNR

`func (o *CreateOrUpdateRequest) GetPROV_JCO_SYSNR() string`

GetPROV_JCO_SYSNR returns the PROV_JCO_SYSNR field if non-nil, zero value otherwise.

### GetPROV_JCO_SYSNROk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_SYSNROk() (*string, bool)`

GetPROV_JCO_SYSNROk returns a tuple with the PROV_JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_SYSNR

`func (o *CreateOrUpdateRequest) SetPROV_JCO_SYSNR(v string)`

SetPROV_JCO_SYSNR sets PROV_JCO_SYSNR field to given value.

### HasPROV_JCO_SYSNR

`func (o *CreateOrUpdateRequest) HasPROV_JCO_SYSNR() bool`

HasPROV_JCO_SYSNR returns a boolean if a field has been set.

### GetPROV_JCO_CLIENT

`func (o *CreateOrUpdateRequest) GetPROV_JCO_CLIENT() string`

GetPROV_JCO_CLIENT returns the PROV_JCO_CLIENT field if non-nil, zero value otherwise.

### GetPROV_JCO_CLIENTOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_CLIENTOk() (*string, bool)`

GetPROV_JCO_CLIENTOk returns a tuple with the PROV_JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_CLIENT

`func (o *CreateOrUpdateRequest) SetPROV_JCO_CLIENT(v string)`

SetPROV_JCO_CLIENT sets PROV_JCO_CLIENT field to given value.

### HasPROV_JCO_CLIENT

`func (o *CreateOrUpdateRequest) HasPROV_JCO_CLIENT() bool`

HasPROV_JCO_CLIENT returns a boolean if a field has been set.

### GetPROV_JCO_USER

`func (o *CreateOrUpdateRequest) GetPROV_JCO_USER() string`

GetPROV_JCO_USER returns the PROV_JCO_USER field if non-nil, zero value otherwise.

### GetPROV_JCO_USEROk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_USEROk() (*string, bool)`

GetPROV_JCO_USEROk returns a tuple with the PROV_JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_USER

`func (o *CreateOrUpdateRequest) SetPROV_JCO_USER(v string)`

SetPROV_JCO_USER sets PROV_JCO_USER field to given value.

### HasPROV_JCO_USER

`func (o *CreateOrUpdateRequest) HasPROV_JCO_USER() bool`

HasPROV_JCO_USER returns a boolean if a field has been set.

### GetPROV_PASSWORD

`func (o *CreateOrUpdateRequest) GetPROV_PASSWORD() string`

GetPROV_PASSWORD returns the PROV_PASSWORD field if non-nil, zero value otherwise.

### GetPROV_PASSWORDOk

`func (o *CreateOrUpdateRequest) GetPROV_PASSWORDOk() (*string, bool)`

GetPROV_PASSWORDOk returns a tuple with the PROV_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_PASSWORD

`func (o *CreateOrUpdateRequest) SetPROV_PASSWORD(v string)`

SetPROV_PASSWORD sets PROV_PASSWORD field to given value.

### HasPROV_PASSWORD

`func (o *CreateOrUpdateRequest) HasPROV_PASSWORD() bool`

HasPROV_PASSWORD returns a boolean if a field has been set.

### GetPROV_JCO_LANG

`func (o *CreateOrUpdateRequest) GetPROV_JCO_LANG() string`

GetPROV_JCO_LANG returns the PROV_JCO_LANG field if non-nil, zero value otherwise.

### GetPROV_JCO_LANGOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_LANGOk() (*string, bool)`

GetPROV_JCO_LANGOk returns a tuple with the PROV_JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_LANG

`func (o *CreateOrUpdateRequest) SetPROV_JCO_LANG(v string)`

SetPROV_JCO_LANG sets PROV_JCO_LANG field to given value.

### HasPROV_JCO_LANG

`func (o *CreateOrUpdateRequest) HasPROV_JCO_LANG() bool`

HasPROV_JCO_LANG returns a boolean if a field has been set.

### GetPROVJCOR3NAME

`func (o *CreateOrUpdateRequest) GetPROVJCOR3NAME() string`

GetPROVJCOR3NAME returns the PROVJCOR3NAME field if non-nil, zero value otherwise.

### GetPROVJCOR3NAMEOk

`func (o *CreateOrUpdateRequest) GetPROVJCOR3NAMEOk() (*string, bool)`

GetPROVJCOR3NAMEOk returns a tuple with the PROVJCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVJCOR3NAME

`func (o *CreateOrUpdateRequest) SetPROVJCOR3NAME(v string)`

SetPROVJCOR3NAME sets PROVJCOR3NAME field to given value.

### HasPROVJCOR3NAME

`func (o *CreateOrUpdateRequest) HasPROVJCOR3NAME() bool`

HasPROVJCOR3NAME returns a boolean if a field has been set.

### GetPROV_JCO_MSHOST

`func (o *CreateOrUpdateRequest) GetPROV_JCO_MSHOST() string`

GetPROV_JCO_MSHOST returns the PROV_JCO_MSHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_MSHOSTOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_MSHOSTOk() (*string, bool)`

GetPROV_JCO_MSHOSTOk returns a tuple with the PROV_JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSHOST

`func (o *CreateOrUpdateRequest) SetPROV_JCO_MSHOST(v string)`

SetPROV_JCO_MSHOST sets PROV_JCO_MSHOST field to given value.

### HasPROV_JCO_MSHOST

`func (o *CreateOrUpdateRequest) HasPROV_JCO_MSHOST() bool`

HasPROV_JCO_MSHOST returns a boolean if a field has been set.

### GetPROV_JCO_MSSERV

`func (o *CreateOrUpdateRequest) GetPROV_JCO_MSSERV() string`

GetPROV_JCO_MSSERV returns the PROV_JCO_MSSERV field if non-nil, zero value otherwise.

### GetPROV_JCO_MSSERVOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_MSSERVOk() (*string, bool)`

GetPROV_JCO_MSSERVOk returns a tuple with the PROV_JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSSERV

`func (o *CreateOrUpdateRequest) SetPROV_JCO_MSSERV(v string)`

SetPROV_JCO_MSSERV sets PROV_JCO_MSSERV field to given value.

### HasPROV_JCO_MSSERV

`func (o *CreateOrUpdateRequest) HasPROV_JCO_MSSERV() bool`

HasPROV_JCO_MSSERV returns a boolean if a field has been set.

### GetPROV_JCO_GROUP

`func (o *CreateOrUpdateRequest) GetPROV_JCO_GROUP() string`

GetPROV_JCO_GROUP returns the PROV_JCO_GROUP field if non-nil, zero value otherwise.

### GetPROV_JCO_GROUPOk

`func (o *CreateOrUpdateRequest) GetPROV_JCO_GROUPOk() (*string, bool)`

GetPROV_JCO_GROUPOk returns a tuple with the PROV_JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_GROUP

`func (o *CreateOrUpdateRequest) SetPROV_JCO_GROUP(v string)`

SetPROV_JCO_GROUP sets PROV_JCO_GROUP field to given value.

### HasPROV_JCO_GROUP

`func (o *CreateOrUpdateRequest) HasPROV_JCO_GROUP() bool`

HasPROV_JCO_GROUP returns a boolean if a field has been set.

### GetPROV_CUA_ENABLED

`func (o *CreateOrUpdateRequest) GetPROV_CUA_ENABLED() string`

GetPROV_CUA_ENABLED returns the PROV_CUA_ENABLED field if non-nil, zero value otherwise.

### GetPROV_CUA_ENABLEDOk

`func (o *CreateOrUpdateRequest) GetPROV_CUA_ENABLEDOk() (*string, bool)`

GetPROV_CUA_ENABLEDOk returns a tuple with the PROV_CUA_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_ENABLED

`func (o *CreateOrUpdateRequest) SetPROV_CUA_ENABLED(v string)`

SetPROV_CUA_ENABLED sets PROV_CUA_ENABLED field to given value.

### HasPROV_CUA_ENABLED

`func (o *CreateOrUpdateRequest) HasPROV_CUA_ENABLED() bool`

HasPROV_CUA_ENABLED returns a boolean if a field has been set.

### GetPROV_CUA_SNC

`func (o *CreateOrUpdateRequest) GetPROV_CUA_SNC() string`

GetPROV_CUA_SNC returns the PROV_CUA_SNC field if non-nil, zero value otherwise.

### GetPROV_CUA_SNCOk

`func (o *CreateOrUpdateRequest) GetPROV_CUA_SNCOk() (*string, bool)`

GetPROV_CUA_SNCOk returns a tuple with the PROV_CUA_SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_SNC

`func (o *CreateOrUpdateRequest) SetPROV_CUA_SNC(v string)`

SetPROV_CUA_SNC sets PROV_CUA_SNC field to given value.

### HasPROV_CUA_SNC

`func (o *CreateOrUpdateRequest) HasPROV_CUA_SNC() bool`

HasPROV_CUA_SNC returns a boolean if a field has been set.

### GetRESET_PWD_FOR_NEWACCOUNT

`func (o *CreateOrUpdateRequest) GetRESET_PWD_FOR_NEWACCOUNT() string`

GetRESET_PWD_FOR_NEWACCOUNT returns the RESET_PWD_FOR_NEWACCOUNT field if non-nil, zero value otherwise.

### GetRESET_PWD_FOR_NEWACCOUNTOk

`func (o *CreateOrUpdateRequest) GetRESET_PWD_FOR_NEWACCOUNTOk() (*string, bool)`

GetRESET_PWD_FOR_NEWACCOUNTOk returns a tuple with the RESET_PWD_FOR_NEWACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESET_PWD_FOR_NEWACCOUNT

`func (o *CreateOrUpdateRequest) SetRESET_PWD_FOR_NEWACCOUNT(v string)`

SetRESET_PWD_FOR_NEWACCOUNT sets RESET_PWD_FOR_NEWACCOUNT field to given value.

### HasRESET_PWD_FOR_NEWACCOUNT

`func (o *CreateOrUpdateRequest) HasRESET_PWD_FOR_NEWACCOUNT() bool`

HasRESET_PWD_FOR_NEWACCOUNT returns a boolean if a field has been set.

### GetENFORCEPASSWORDCHANGE

`func (o *CreateOrUpdateRequest) GetENFORCEPASSWORDCHANGE() string`

GetENFORCEPASSWORDCHANGE returns the ENFORCEPASSWORDCHANGE field if non-nil, zero value otherwise.

### GetENFORCEPASSWORDCHANGEOk

`func (o *CreateOrUpdateRequest) GetENFORCEPASSWORDCHANGEOk() (*string, bool)`

GetENFORCEPASSWORDCHANGEOk returns a tuple with the ENFORCEPASSWORDCHANGE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCEPASSWORDCHANGE

`func (o *CreateOrUpdateRequest) SetENFORCEPASSWORDCHANGE(v string)`

SetENFORCEPASSWORDCHANGE sets ENFORCEPASSWORDCHANGE field to given value.

### HasENFORCEPASSWORDCHANGE

`func (o *CreateOrUpdateRequest) HasENFORCEPASSWORDCHANGE() bool`

HasENFORCEPASSWORDCHANGE returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateRequest) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *CreateOrUpdateRequest) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateRequest) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *CreateOrUpdateRequest) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateRequest) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *CreateOrUpdateRequest) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateRequest) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *CreateOrUpdateRequest) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateRequest) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *CreateOrUpdateRequest) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateRequest) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *CreateOrUpdateRequest) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *CreateOrUpdateRequest) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateRequest) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *CreateOrUpdateRequest) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetHANAREFTABLEJSON

`func (o *CreateOrUpdateRequest) GetHANAREFTABLEJSON() string`

GetHANAREFTABLEJSON returns the HANAREFTABLEJSON field if non-nil, zero value otherwise.

### GetHANAREFTABLEJSONOk

`func (o *CreateOrUpdateRequest) GetHANAREFTABLEJSONOk() (*string, bool)`

GetHANAREFTABLEJSONOk returns a tuple with the HANAREFTABLEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHANAREFTABLEJSON

`func (o *CreateOrUpdateRequest) SetHANAREFTABLEJSON(v string)`

SetHANAREFTABLEJSON sets HANAREFTABLEJSON field to given value.

### HasHANAREFTABLEJSON

`func (o *CreateOrUpdateRequest) HasHANAREFTABLEJSON() bool`

HasHANAREFTABLEJSON returns a boolean if a field has been set.

### GetUSERIMPORTJSON

`func (o *CreateOrUpdateRequest) GetUSERIMPORTJSON() string`

GetUSERIMPORTJSON returns the USERIMPORTJSON field if non-nil, zero value otherwise.

### GetUSERIMPORTJSONOk

`func (o *CreateOrUpdateRequest) GetUSERIMPORTJSONOk() (*string, bool)`

GetUSERIMPORTJSONOk returns a tuple with the USERIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORTJSON

`func (o *CreateOrUpdateRequest) SetUSERIMPORTJSON(v string)`

SetUSERIMPORTJSON sets USERIMPORTJSON field to given value.

### HasUSERIMPORTJSON

`func (o *CreateOrUpdateRequest) HasUSERIMPORTJSON() bool`

HasUSERIMPORTJSON returns a boolean if a field has been set.

### GetSETCUASYSTEM

`func (o *CreateOrUpdateRequest) GetSETCUASYSTEM() string`

GetSETCUASYSTEM returns the SETCUASYSTEM field if non-nil, zero value otherwise.

### GetSETCUASYSTEMOk

`func (o *CreateOrUpdateRequest) GetSETCUASYSTEMOk() (*string, bool)`

GetSETCUASYSTEMOk returns a tuple with the SETCUASYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETCUASYSTEM

`func (o *CreateOrUpdateRequest) SetSETCUASYSTEM(v string)`

SetSETCUASYSTEM sets SETCUASYSTEM field to given value.

### HasSETCUASYSTEM

`func (o *CreateOrUpdateRequest) HasSETCUASYSTEM() bool`

HasSETCUASYSTEM returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *CreateOrUpdateRequest) GetFIREFIGHTERID_GRANT_ACCESS_JSON() string`

GetFIREFIGHTERID_GRANT_ACCESS_JSON returns the FIREFIGHTERID_GRANT_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_JSONOk

`func (o *CreateOrUpdateRequest) GetFIREFIGHTERID_GRANT_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *CreateOrUpdateRequest) SetFIREFIGHTERID_GRANT_ACCESS_JSON(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_JSON sets FIREFIGHTERID_GRANT_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *CreateOrUpdateRequest) HasFIREFIGHTERID_GRANT_ACCESS_JSON() bool`

HasFIREFIGHTERID_GRANT_ACCESS_JSON returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *CreateOrUpdateRequest) GetFIREFIGHTERID_REVOKE_ACCESS_JSON() string`

GetFIREFIGHTERID_REVOKE_ACCESS_JSON returns the FIREFIGHTERID_REVOKE_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk

`func (o *CreateOrUpdateRequest) GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *CreateOrUpdateRequest) SetFIREFIGHTERID_REVOKE_ACCESS_JSON(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_JSON sets FIREFIGHTERID_REVOKE_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *CreateOrUpdateRequest) HasFIREFIGHTERID_REVOKE_ACCESS_JSON() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_JSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON

`func (o *CreateOrUpdateRequest) GetEXTERNAL_SOD_EVAL_JSON() string`

GetEXTERNAL_SOD_EVAL_JSON returns the EXTERNAL_SOD_EVAL_JSON field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSONOk

`func (o *CreateOrUpdateRequest) GetEXTERNAL_SOD_EVAL_JSONOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSONOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON

`func (o *CreateOrUpdateRequest) SetEXTERNAL_SOD_EVAL_JSON(v string)`

SetEXTERNAL_SOD_EVAL_JSON sets EXTERNAL_SOD_EVAL_JSON field to given value.

### HasEXTERNAL_SOD_EVAL_JSON

`func (o *CreateOrUpdateRequest) HasEXTERNAL_SOD_EVAL_JSON() bool`

HasEXTERNAL_SOD_EVAL_JSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *CreateOrUpdateRequest) GetEXTERNAL_SOD_EVAL_JSON_DETAIL() string`

GetEXTERNAL_SOD_EVAL_JSON_DETAIL returns the EXTERNAL_SOD_EVAL_JSON_DETAIL field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSON_DETAILOk

`func (o *CreateOrUpdateRequest) GetEXTERNAL_SOD_EVAL_JSON_DETAILOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSON_DETAILOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON_DETAIL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *CreateOrUpdateRequest) SetEXTERNAL_SOD_EVAL_JSON_DETAIL(v string)`

SetEXTERNAL_SOD_EVAL_JSON_DETAIL sets EXTERNAL_SOD_EVAL_JSON_DETAIL field to given value.

### HasEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *CreateOrUpdateRequest) HasEXTERNAL_SOD_EVAL_JSON_DETAIL() bool`

HasEXTERNAL_SOD_EVAL_JSON_DETAIL returns a boolean if a field has been set.

### GetLOGS_TABLE_FILTER

`func (o *CreateOrUpdateRequest) GetLOGS_TABLE_FILTER() string`

GetLOGS_TABLE_FILTER returns the LOGS_TABLE_FILTER field if non-nil, zero value otherwise.

### GetLOGS_TABLE_FILTEROk

`func (o *CreateOrUpdateRequest) GetLOGS_TABLE_FILTEROk() (*string, bool)`

GetLOGS_TABLE_FILTEROk returns a tuple with the LOGS_TABLE_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGS_TABLE_FILTER

`func (o *CreateOrUpdateRequest) SetLOGS_TABLE_FILTER(v string)`

SetLOGS_TABLE_FILTER sets LOGS_TABLE_FILTER field to given value.

### HasLOGS_TABLE_FILTER

`func (o *CreateOrUpdateRequest) HasLOGS_TABLE_FILTER() bool`

HasLOGS_TABLE_FILTER returns a boolean if a field has been set.

### GetSAPTABLE_FILTER_LANG

`func (o *CreateOrUpdateRequest) GetSAPTABLE_FILTER_LANG() string`

GetSAPTABLE_FILTER_LANG returns the SAPTABLE_FILTER_LANG field if non-nil, zero value otherwise.

### GetSAPTABLE_FILTER_LANGOk

`func (o *CreateOrUpdateRequest) GetSAPTABLE_FILTER_LANGOk() (*string, bool)`

GetSAPTABLE_FILTER_LANGOk returns a tuple with the SAPTABLE_FILTER_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAPTABLE_FILTER_LANG

`func (o *CreateOrUpdateRequest) SetSAPTABLE_FILTER_LANG(v string)`

SetSAPTABLE_FILTER_LANG sets SAPTABLE_FILTER_LANG field to given value.

### HasSAPTABLE_FILTER_LANG

`func (o *CreateOrUpdateRequest) HasSAPTABLE_FILTER_LANG() bool`

HasSAPTABLE_FILTER_LANG returns a boolean if a field has been set.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *CreateOrUpdateRequest) GetALTERNATE_OUTPUT_PARAMETER_ET_DATA() string`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATA returns the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if non-nil, zero value otherwise.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk

`func (o *CreateOrUpdateRequest) GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk() (*string, bool)`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk returns a tuple with the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *CreateOrUpdateRequest) SetALTERNATE_OUTPUT_PARAMETER_ET_DATA(v string)`

SetALTERNATE_OUTPUT_PARAMETER_ET_DATA sets ALTERNATE_OUTPUT_PARAMETER_ET_DATA field to given value.

### HasALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *CreateOrUpdateRequest) HasALTERNATE_OUTPUT_PARAMETER_ET_DATA() bool`

HasALTERNATE_OUTPUT_PARAMETER_ET_DATA returns a boolean if a field has been set.

### GetAUDIT_LOG_JSON

`func (o *CreateOrUpdateRequest) GetAUDIT_LOG_JSON() string`

GetAUDIT_LOG_JSON returns the AUDIT_LOG_JSON field if non-nil, zero value otherwise.

### GetAUDIT_LOG_JSONOk

`func (o *CreateOrUpdateRequest) GetAUDIT_LOG_JSONOk() (*string, bool)`

GetAUDIT_LOG_JSONOk returns a tuple with the AUDIT_LOG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_LOG_JSON

`func (o *CreateOrUpdateRequest) SetAUDIT_LOG_JSON(v string)`

SetAUDIT_LOG_JSON sets AUDIT_LOG_JSON field to given value.

### HasAUDIT_LOG_JSON

`func (o *CreateOrUpdateRequest) HasAUDIT_LOG_JSON() bool`

HasAUDIT_LOG_JSON returns a boolean if a field has been set.

### GetECCORS4HANA

`func (o *CreateOrUpdateRequest) GetECCORS4HANA() string`

GetECCORS4HANA returns the ECCORS4HANA field if non-nil, zero value otherwise.

### GetECCORS4HANAOk

`func (o *CreateOrUpdateRequest) GetECCORS4HANAOk() (*string, bool)`

GetECCORS4HANAOk returns a tuple with the ECCORS4HANA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCORS4HANA

`func (o *CreateOrUpdateRequest) SetECCORS4HANA(v string)`

SetECCORS4HANA sets ECCORS4HANA field to given value.

### HasECCORS4HANA

`func (o *CreateOrUpdateRequest) HasECCORS4HANA() bool`

HasECCORS4HANA returns a boolean if a field has been set.

### GetDATA_IMPORT_FILTER

`func (o *CreateOrUpdateRequest) GetDATA_IMPORT_FILTER() string`

GetDATA_IMPORT_FILTER returns the DATA_IMPORT_FILTER field if non-nil, zero value otherwise.

### GetDATA_IMPORT_FILTEROk

`func (o *CreateOrUpdateRequest) GetDATA_IMPORT_FILTEROk() (*string, bool)`

GetDATA_IMPORT_FILTEROk returns a tuple with the DATA_IMPORT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_IMPORT_FILTER

`func (o *CreateOrUpdateRequest) SetDATA_IMPORT_FILTER(v string)`

SetDATA_IMPORT_FILTER sets DATA_IMPORT_FILTER field to given value.

### HasDATA_IMPORT_FILTER

`func (o *CreateOrUpdateRequest) HasDATA_IMPORT_FILTER() bool`

HasDATA_IMPORT_FILTER returns a boolean if a field has been set.

### GetConfigJSON

`func (o *CreateOrUpdateRequest) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *CreateOrUpdateRequest) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *CreateOrUpdateRequest) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *CreateOrUpdateRequest) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *CreateOrUpdateRequest) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *CreateOrUpdateRequest) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *CreateOrUpdateRequest) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *CreateOrUpdateRequest) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *CreateOrUpdateRequest) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *CreateOrUpdateRequest) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetREFRESH_TOKEN

`func (o *CreateOrUpdateRequest) GetREFRESH_TOKEN() string`

GetREFRESH_TOKEN returns the REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetREFRESH_TOKENOk

`func (o *CreateOrUpdateRequest) GetREFRESH_TOKENOk() (*string, bool)`

GetREFRESH_TOKENOk returns a tuple with the REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREFRESH_TOKEN

`func (o *CreateOrUpdateRequest) SetREFRESH_TOKEN(v string)`

SetREFRESH_TOKEN sets REFRESH_TOKEN field to given value.

### HasREFRESH_TOKEN

`func (o *CreateOrUpdateRequest) HasREFRESH_TOKEN() bool`

HasREFRESH_TOKEN returns a boolean if a field has been set.

### GetREDIRECT_URI

`func (o *CreateOrUpdateRequest) GetREDIRECT_URI() string`

GetREDIRECT_URI returns the REDIRECT_URI field if non-nil, zero value otherwise.

### GetREDIRECT_URIOk

`func (o *CreateOrUpdateRequest) GetREDIRECT_URIOk() (*string, bool)`

GetREDIRECT_URIOk returns a tuple with the REDIRECT_URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREDIRECT_URI

`func (o *CreateOrUpdateRequest) SetREDIRECT_URI(v string)`

SetREDIRECT_URI sets REDIRECT_URI field to given value.

### HasREDIRECT_URI

`func (o *CreateOrUpdateRequest) HasREDIRECT_URI() bool`

HasREDIRECT_URI returns a boolean if a field has been set.

### GetINSTANCE_URL

`func (o *CreateOrUpdateRequest) GetINSTANCE_URL() string`

GetINSTANCE_URL returns the INSTANCE_URL field if non-nil, zero value otherwise.

### GetINSTANCE_URLOk

`func (o *CreateOrUpdateRequest) GetINSTANCE_URLOk() (*string, bool)`

GetINSTANCE_URLOk returns a tuple with the INSTANCE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINSTANCE_URL

`func (o *CreateOrUpdateRequest) SetINSTANCE_URL(v string)`

SetINSTANCE_URL sets INSTANCE_URL field to given value.

### HasINSTANCE_URL

`func (o *CreateOrUpdateRequest) HasINSTANCE_URL() bool`

HasINSTANCE_URL returns a boolean if a field has been set.

### GetOBJECT_TO_BE_IMPORTED

`func (o *CreateOrUpdateRequest) GetOBJECT_TO_BE_IMPORTED() string`

GetOBJECT_TO_BE_IMPORTED returns the OBJECT_TO_BE_IMPORTED field if non-nil, zero value otherwise.

### GetOBJECT_TO_BE_IMPORTEDOk

`func (o *CreateOrUpdateRequest) GetOBJECT_TO_BE_IMPORTEDOk() (*string, bool)`

GetOBJECT_TO_BE_IMPORTEDOk returns a tuple with the OBJECT_TO_BE_IMPORTED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECT_TO_BE_IMPORTED

`func (o *CreateOrUpdateRequest) SetOBJECT_TO_BE_IMPORTED(v string)`

SetOBJECT_TO_BE_IMPORTED sets OBJECT_TO_BE_IMPORTED field to given value.

### HasOBJECT_TO_BE_IMPORTED

`func (o *CreateOrUpdateRequest) HasOBJECT_TO_BE_IMPORTED() bool`

HasOBJECT_TO_BE_IMPORTED returns a boolean if a field has been set.

### GetFEATURE_LICENSE_JSON

`func (o *CreateOrUpdateRequest) GetFEATURE_LICENSE_JSON() string`

GetFEATURE_LICENSE_JSON returns the FEATURE_LICENSE_JSON field if non-nil, zero value otherwise.

### GetFEATURE_LICENSE_JSONOk

`func (o *CreateOrUpdateRequest) GetFEATURE_LICENSE_JSONOk() (*string, bool)`

GetFEATURE_LICENSE_JSONOk returns a tuple with the FEATURE_LICENSE_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFEATURE_LICENSE_JSON

`func (o *CreateOrUpdateRequest) SetFEATURE_LICENSE_JSON(v string)`

SetFEATURE_LICENSE_JSON sets FEATURE_LICENSE_JSON field to given value.

### HasFEATURE_LICENSE_JSON

`func (o *CreateOrUpdateRequest) HasFEATURE_LICENSE_JSON() bool`

HasFEATURE_LICENSE_JSON returns a boolean if a field has been set.

### GetCUSTOM_CREATEACCOUNT_URL

`func (o *CreateOrUpdateRequest) GetCUSTOM_CREATEACCOUNT_URL() string`

GetCUSTOM_CREATEACCOUNT_URL returns the CUSTOM_CREATEACCOUNT_URL field if non-nil, zero value otherwise.

### GetCUSTOM_CREATEACCOUNT_URLOk

`func (o *CreateOrUpdateRequest) GetCUSTOM_CREATEACCOUNT_URLOk() (*string, bool)`

GetCUSTOM_CREATEACCOUNT_URLOk returns a tuple with the CUSTOM_CREATEACCOUNT_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CREATEACCOUNT_URL

`func (o *CreateOrUpdateRequest) SetCUSTOM_CREATEACCOUNT_URL(v string)`

SetCUSTOM_CREATEACCOUNT_URL sets CUSTOM_CREATEACCOUNT_URL field to given value.

### HasCUSTOM_CREATEACCOUNT_URL

`func (o *CreateOrUpdateRequest) HasCUSTOM_CREATEACCOUNT_URL() bool`

HasCUSTOM_CREATEACCOUNT_URL returns a boolean if a field has been set.

### GetACCOUNT_FILTER_QUERY

`func (o *CreateOrUpdateRequest) GetACCOUNT_FILTER_QUERY() string`

GetACCOUNT_FILTER_QUERY returns the ACCOUNT_FILTER_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FILTER_QUERYOk

`func (o *CreateOrUpdateRequest) GetACCOUNT_FILTER_QUERYOk() (*string, bool)`

GetACCOUNT_FILTER_QUERYOk returns a tuple with the ACCOUNT_FILTER_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FILTER_QUERY

`func (o *CreateOrUpdateRequest) SetACCOUNT_FILTER_QUERY(v string)`

SetACCOUNT_FILTER_QUERY sets ACCOUNT_FILTER_QUERY field to given value.

### HasACCOUNT_FILTER_QUERY

`func (o *CreateOrUpdateRequest) HasACCOUNT_FILTER_QUERY() bool`

HasACCOUNT_FILTER_QUERY returns a boolean if a field has been set.

### GetACCOUNT_FIELD_QUERY

`func (o *CreateOrUpdateRequest) GetACCOUNT_FIELD_QUERY() string`

GetACCOUNT_FIELD_QUERY returns the ACCOUNT_FIELD_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FIELD_QUERYOk

`func (o *CreateOrUpdateRequest) GetACCOUNT_FIELD_QUERYOk() (*string, bool)`

GetACCOUNT_FIELD_QUERYOk returns a tuple with the ACCOUNT_FIELD_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FIELD_QUERY

`func (o *CreateOrUpdateRequest) SetACCOUNT_FIELD_QUERY(v string)`

SetACCOUNT_FIELD_QUERY sets ACCOUNT_FIELD_QUERY field to given value.

### HasACCOUNT_FIELD_QUERY

`func (o *CreateOrUpdateRequest) HasACCOUNT_FIELD_QUERY() bool`

HasACCOUNT_FIELD_QUERY returns a boolean if a field has been set.

### GetFIELD_MAPPING_JSON

`func (o *CreateOrUpdateRequest) GetFIELD_MAPPING_JSON() string`

GetFIELD_MAPPING_JSON returns the FIELD_MAPPING_JSON field if non-nil, zero value otherwise.

### GetFIELD_MAPPING_JSONOk

`func (o *CreateOrUpdateRequest) GetFIELD_MAPPING_JSONOk() (*string, bool)`

GetFIELD_MAPPING_JSONOk returns a tuple with the FIELD_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIELD_MAPPING_JSON

`func (o *CreateOrUpdateRequest) SetFIELD_MAPPING_JSON(v string)`

SetFIELD_MAPPING_JSON sets FIELD_MAPPING_JSON field to given value.

### HasFIELD_MAPPING_JSON

`func (o *CreateOrUpdateRequest) HasFIELD_MAPPING_JSON() bool`

HasFIELD_MAPPING_JSON returns a boolean if a field has been set.

### GetMODIFYACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetMODIFYACCOUNTJSON() string`

GetMODIFYACCOUNTJSON returns the MODIFYACCOUNTJSON field if non-nil, zero value otherwise.

### GetMODIFYACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetMODIFYACCOUNTJSONOk() (*string, bool)`

GetMODIFYACCOUNTJSONOk returns a tuple with the MODIFYACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetMODIFYACCOUNTJSON(v string)`

SetMODIFYACCOUNTJSON sets MODIFYACCOUNTJSON field to given value.

### HasMODIFYACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasMODIFYACCOUNTJSON() bool`

HasMODIFYACCOUNTJSON returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *CreateOrUpdateRequest) GetConnectionJSON() map[string]interface{}`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *CreateOrUpdateRequest) GetConnectionJSONOk() (*map[string]interface{}, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *CreateOrUpdateRequest) SetConnectionJSON(v map[string]interface{})`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *CreateOrUpdateRequest) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *CreateOrUpdateRequest) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *CreateOrUpdateRequest) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *CreateOrUpdateRequest) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *CreateOrUpdateRequest) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetImportAccountEntJSON

`func (o *CreateOrUpdateRequest) GetImportAccountEntJSON() string`

GetImportAccountEntJSON returns the ImportAccountEntJSON field if non-nil, zero value otherwise.

### GetImportAccountEntJSONOk

`func (o *CreateOrUpdateRequest) GetImportAccountEntJSONOk() (*string, bool)`

GetImportAccountEntJSONOk returns a tuple with the ImportAccountEntJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAccountEntJSON

`func (o *CreateOrUpdateRequest) SetImportAccountEntJSON(v string)`

SetImportAccountEntJSON sets ImportAccountEntJSON field to given value.

### HasImportAccountEntJSON

`func (o *CreateOrUpdateRequest) HasImportAccountEntJSON() bool`

HasImportAccountEntJSON returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *CreateOrUpdateRequest) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *CreateOrUpdateRequest) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *CreateOrUpdateRequest) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *CreateOrUpdateRequest) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *CreateOrUpdateRequest) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *CreateOrUpdateRequest) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *CreateOrUpdateRequest) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *CreateOrUpdateRequest) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *CreateOrUpdateRequest) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *CreateOrUpdateRequest) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *CreateOrUpdateRequest) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *CreateOrUpdateRequest) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *CreateOrUpdateRequest) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *CreateOrUpdateRequest) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *CreateOrUpdateRequest) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *CreateOrUpdateRequest) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *CreateOrUpdateRequest) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *CreateOrUpdateRequest) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *CreateOrUpdateRequest) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *CreateOrUpdateRequest) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *CreateOrUpdateRequest) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *CreateOrUpdateRequest) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *CreateOrUpdateRequest) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *CreateOrUpdateRequest) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetUpdateUserJSON

`func (o *CreateOrUpdateRequest) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *CreateOrUpdateRequest) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *CreateOrUpdateRequest) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *CreateOrUpdateRequest) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *CreateOrUpdateRequest) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *CreateOrUpdateRequest) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *CreateOrUpdateRequest) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *CreateOrUpdateRequest) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *CreateOrUpdateRequest) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *CreateOrUpdateRequest) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *CreateOrUpdateRequest) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *CreateOrUpdateRequest) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetTicketStatusJSON

`func (o *CreateOrUpdateRequest) GetTicketStatusJSON() string`

GetTicketStatusJSON returns the TicketStatusJSON field if non-nil, zero value otherwise.

### GetTicketStatusJSONOk

`func (o *CreateOrUpdateRequest) GetTicketStatusJSONOk() (*string, bool)`

GetTicketStatusJSONOk returns a tuple with the TicketStatusJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatusJSON

`func (o *CreateOrUpdateRequest) SetTicketStatusJSON(v string)`

SetTicketStatusJSON sets TicketStatusJSON field to given value.

### HasTicketStatusJSON

`func (o *CreateOrUpdateRequest) HasTicketStatusJSON() bool`

HasTicketStatusJSON returns a boolean if a field has been set.

### GetCreateTicketJSON

`func (o *CreateOrUpdateRequest) GetCreateTicketJSON() string`

GetCreateTicketJSON returns the CreateTicketJSON field if non-nil, zero value otherwise.

### GetCreateTicketJSONOk

`func (o *CreateOrUpdateRequest) GetCreateTicketJSONOk() (*string, bool)`

GetCreateTicketJSONOk returns a tuple with the CreateTicketJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTicketJSON

`func (o *CreateOrUpdateRequest) SetCreateTicketJSON(v string)`

SetCreateTicketJSON sets CreateTicketJSON field to given value.

### HasCreateTicketJSON

`func (o *CreateOrUpdateRequest) HasCreateTicketJSON() bool`

HasCreateTicketJSON returns a boolean if a field has been set.

### GetPasswdPolicyJSON

`func (o *CreateOrUpdateRequest) GetPasswdPolicyJSON() string`

GetPasswdPolicyJSON returns the PasswdPolicyJSON field if non-nil, zero value otherwise.

### GetPasswdPolicyJSONOk

`func (o *CreateOrUpdateRequest) GetPasswdPolicyJSONOk() (*string, bool)`

GetPasswdPolicyJSONOk returns a tuple with the PasswdPolicyJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswdPolicyJSON

`func (o *CreateOrUpdateRequest) SetPasswdPolicyJSON(v string)`

SetPasswdPolicyJSON sets PasswdPolicyJSON field to given value.

### HasPasswdPolicyJSON

`func (o *CreateOrUpdateRequest) HasPasswdPolicyJSON() bool`

HasPasswdPolicyJSON returns a boolean if a field has been set.

### GetAddFFIDAccessJSON

`func (o *CreateOrUpdateRequest) GetAddFFIDAccessJSON() string`

GetAddFFIDAccessJSON returns the AddFFIDAccessJSON field if non-nil, zero value otherwise.

### GetAddFFIDAccessJSONOk

`func (o *CreateOrUpdateRequest) GetAddFFIDAccessJSONOk() (*string, bool)`

GetAddFFIDAccessJSONOk returns a tuple with the AddFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFFIDAccessJSON

`func (o *CreateOrUpdateRequest) SetAddFFIDAccessJSON(v string)`

SetAddFFIDAccessJSON sets AddFFIDAccessJSON field to given value.

### HasAddFFIDAccessJSON

`func (o *CreateOrUpdateRequest) HasAddFFIDAccessJSON() bool`

HasAddFFIDAccessJSON returns a boolean if a field has been set.

### GetRemoveFFIDAccessJSON

`func (o *CreateOrUpdateRequest) GetRemoveFFIDAccessJSON() string`

GetRemoveFFIDAccessJSON returns the RemoveFFIDAccessJSON field if non-nil, zero value otherwise.

### GetRemoveFFIDAccessJSONOk

`func (o *CreateOrUpdateRequest) GetRemoveFFIDAccessJSONOk() (*string, bool)`

GetRemoveFFIDAccessJSONOk returns a tuple with the RemoveFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFFIDAccessJSON

`func (o *CreateOrUpdateRequest) SetRemoveFFIDAccessJSON(v string)`

SetRemoveFFIDAccessJSON sets RemoveFFIDAccessJSON field to given value.

### HasRemoveFFIDAccessJSON

`func (o *CreateOrUpdateRequest) HasRemoveFFIDAccessJSON() bool`

HasRemoveFFIDAccessJSON returns a boolean if a field has been set.

### GetSendOtpJSON

`func (o *CreateOrUpdateRequest) GetSendOtpJSON() string`

GetSendOtpJSON returns the SendOtpJSON field if non-nil, zero value otherwise.

### GetSendOtpJSONOk

`func (o *CreateOrUpdateRequest) GetSendOtpJSONOk() (*string, bool)`

GetSendOtpJSONOk returns a tuple with the SendOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOtpJSON

`func (o *CreateOrUpdateRequest) SetSendOtpJSON(v string)`

SetSendOtpJSON sets SendOtpJSON field to given value.

### HasSendOtpJSON

`func (o *CreateOrUpdateRequest) HasSendOtpJSON() bool`

HasSendOtpJSON returns a boolean if a field has been set.

### GetValidateOtpJSON

`func (o *CreateOrUpdateRequest) GetValidateOtpJSON() string`

GetValidateOtpJSON returns the ValidateOtpJSON field if non-nil, zero value otherwise.

### GetValidateOtpJSONOk

`func (o *CreateOrUpdateRequest) GetValidateOtpJSONOk() (*string, bool)`

GetValidateOtpJSONOk returns a tuple with the ValidateOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOtpJSON

`func (o *CreateOrUpdateRequest) SetValidateOtpJSON(v string)`

SetValidateOtpJSON sets ValidateOtpJSON field to given value.

### HasValidateOtpJSON

`func (o *CreateOrUpdateRequest) HasValidateOtpJSON() bool`

HasValidateOtpJSON returns a boolean if a field has been set.

### GetBASEURL

`func (o *CreateOrUpdateRequest) GetBASEURL() string`

GetBASEURL returns the BASEURL field if non-nil, zero value otherwise.

### GetBASEURLOk

`func (o *CreateOrUpdateRequest) GetBASEURLOk() (*string, bool)`

GetBASEURLOk returns a tuple with the BASEURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASEURL

`func (o *CreateOrUpdateRequest) SetBASEURL(v string)`

SetBASEURL sets BASEURL field to given value.


### GetTENANT_ID

`func (o *CreateOrUpdateRequest) GetTENANT_ID() string`

GetTENANT_ID returns the TENANT_ID field if non-nil, zero value otherwise.

### GetTENANT_IDOk

`func (o *CreateOrUpdateRequest) GetTENANT_IDOk() (*string, bool)`

GetTENANT_IDOk returns a tuple with the TENANT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_ID

`func (o *CreateOrUpdateRequest) SetTENANT_ID(v string)`

SetTENANT_ID sets TENANT_ID field to given value.


### GetLOGIN_URL

`func (o *CreateOrUpdateRequest) GetLOGIN_URL() string`

GetLOGIN_URL returns the LOGIN_URL field if non-nil, zero value otherwise.

### GetLOGIN_URLOk

`func (o *CreateOrUpdateRequest) GetLOGIN_URLOk() (*string, bool)`

GetLOGIN_URLOk returns a tuple with the LOGIN_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGIN_URL

`func (o *CreateOrUpdateRequest) SetLOGIN_URL(v string)`

SetLOGIN_URL sets LOGIN_URL field to given value.


### GetUSER_FILTER

`func (o *CreateOrUpdateRequest) GetUSER_FILTER() string`

GetUSER_FILTER returns the USER_FILTER field if non-nil, zero value otherwise.

### GetUSER_FILTEROk

`func (o *CreateOrUpdateRequest) GetUSER_FILTEROk() (*string, bool)`

GetUSER_FILTEROk returns a tuple with the USER_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_FILTER

`func (o *CreateOrUpdateRequest) SetUSER_FILTER(v string)`

SetUSER_FILTER sets USER_FILTER field to given value.

### HasUSER_FILTER

`func (o *CreateOrUpdateRequest) HasUSER_FILTER() bool`

HasUSER_FILTER returns a boolean if a field has been set.

### GetUSER_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) GetUSER_IMPORT_MAPPING() string`

GetUSER_IMPORT_MAPPING returns the USER_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetUSER_IMPORT_MAPPINGOk

`func (o *CreateOrUpdateRequest) GetUSER_IMPORT_MAPPINGOk() (*string, bool)`

GetUSER_IMPORT_MAPPINGOk returns a tuple with the USER_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) SetUSER_IMPORT_MAPPING(v string)`

SetUSER_IMPORT_MAPPING sets USER_IMPORT_MAPPING field to given value.

### HasUSER_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) HasUSER_IMPORT_MAPPING() bool`

HasUSER_IMPORT_MAPPING returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) GetACCOUNT_IMPORT_MAPPING() string`

GetACCOUNT_IMPORT_MAPPING returns the ACCOUNT_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_MAPPINGOk

`func (o *CreateOrUpdateRequest) GetACCOUNT_IMPORT_MAPPINGOk() (*string, bool)`

GetACCOUNT_IMPORT_MAPPINGOk returns a tuple with the ACCOUNT_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) SetACCOUNT_IMPORT_MAPPING(v string)`

SetACCOUNT_IMPORT_MAPPING sets ACCOUNT_IMPORT_MAPPING field to given value.

### HasACCOUNT_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) HasACCOUNT_IMPORT_MAPPING() bool`

HasACCOUNT_IMPORT_MAPPING returns a boolean if a field has been set.

### GetORGANIZATION_FILTER

`func (o *CreateOrUpdateRequest) GetORGANIZATION_FILTER() string`

GetORGANIZATION_FILTER returns the ORGANIZATION_FILTER field if non-nil, zero value otherwise.

### GetORGANIZATION_FILTEROk

`func (o *CreateOrUpdateRequest) GetORGANIZATION_FILTEROk() (*string, bool)`

GetORGANIZATION_FILTEROk returns a tuple with the ORGANIZATION_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_FILTER

`func (o *CreateOrUpdateRequest) SetORGANIZATION_FILTER(v string)`

SetORGANIZATION_FILTER sets ORGANIZATION_FILTER field to given value.

### HasORGANIZATION_FILTER

`func (o *CreateOrUpdateRequest) HasORGANIZATION_FILTER() bool`

HasORGANIZATION_FILTER returns a boolean if a field has been set.

### GetSCOPE

`func (o *CreateOrUpdateRequest) GetSCOPE() string`

GetSCOPE returns the SCOPE field if non-nil, zero value otherwise.

### GetSCOPEOk

`func (o *CreateOrUpdateRequest) GetSCOPEOk() (*string, bool)`

GetSCOPEOk returns a tuple with the SCOPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCOPE

`func (o *CreateOrUpdateRequest) SetSCOPE(v string)`

SetSCOPE sets SCOPE field to given value.

### HasSCOPE

`func (o *CreateOrUpdateRequest) HasSCOPE() bool`

HasSCOPE returns a boolean if a field has been set.

### GetUSERS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) GetUSERS_LAST_IMPORT_TIME() string`

GetUSERS_LAST_IMPORT_TIME returns the USERS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetUSERS_LAST_IMPORT_TIMEOk

`func (o *CreateOrUpdateRequest) GetUSERS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetUSERS_LAST_IMPORT_TIMEOk returns a tuple with the USERS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) SetUSERS_LAST_IMPORT_TIME(v string)`

SetUSERS_LAST_IMPORT_TIME sets USERS_LAST_IMPORT_TIME field to given value.

### HasUSERS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) HasUSERS_LAST_IMPORT_TIME() bool`

HasUSERS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetACCOUNTS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) GetACCOUNTS_LAST_IMPORT_TIME() string`

GetACCOUNTS_LAST_IMPORT_TIME returns the ACCOUNTS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCOUNTS_LAST_IMPORT_TIMEOk

`func (o *CreateOrUpdateRequest) GetACCOUNTS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCOUNTS_LAST_IMPORT_TIMEOk returns a tuple with the ACCOUNTS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) SetACCOUNTS_LAST_IMPORT_TIME(v string)`

SetACCOUNTS_LAST_IMPORT_TIME sets ACCOUNTS_LAST_IMPORT_TIME field to given value.

### HasACCOUNTS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) HasACCOUNTS_LAST_IMPORT_TIME() bool`

HasACCOUNTS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetACCESS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) GetACCESS_LAST_IMPORT_TIME() string`

GetACCESS_LAST_IMPORT_TIME returns the ACCESS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCESS_LAST_IMPORT_TIMEOk

`func (o *CreateOrUpdateRequest) GetACCESS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCESS_LAST_IMPORT_TIMEOk returns a tuple with the ACCESS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) SetACCESS_LAST_IMPORT_TIME(v string)`

SetACCESS_LAST_IMPORT_TIME sets ACCESS_LAST_IMPORT_TIME field to given value.

### HasACCESS_LAST_IMPORT_TIME

`func (o *CreateOrUpdateRequest) HasACCESS_LAST_IMPORT_TIME() bool`

HasACCESS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetBASE_URL

`func (o *CreateOrUpdateRequest) GetBASE_URL() string`

GetBASE_URL returns the BASE_URL field if non-nil, zero value otherwise.

### GetBASE_URLOk

`func (o *CreateOrUpdateRequest) GetBASE_URLOk() (*string, bool)`

GetBASE_URLOk returns a tuple with the BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE_URL

`func (o *CreateOrUpdateRequest) SetBASE_URL(v string)`

SetBASE_URL sets BASE_URL field to given value.

### HasBASE_URL

`func (o *CreateOrUpdateRequest) HasBASE_URL() bool`

HasBASE_URL returns a boolean if a field has been set.

### GetAPI_VERSION

`func (o *CreateOrUpdateRequest) GetAPI_VERSION() string`

GetAPI_VERSION returns the API_VERSION field if non-nil, zero value otherwise.

### GetAPI_VERSIONOk

`func (o *CreateOrUpdateRequest) GetAPI_VERSIONOk() (*string, bool)`

GetAPI_VERSIONOk returns a tuple with the API_VERSION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_VERSION

`func (o *CreateOrUpdateRequest) SetAPI_VERSION(v string)`

SetAPI_VERSION sets API_VERSION field to given value.

### HasAPI_VERSION

`func (o *CreateOrUpdateRequest) HasAPI_VERSION() bool`

HasAPI_VERSION returns a boolean if a field has been set.

### GetTENANT_NAME

`func (o *CreateOrUpdateRequest) GetTENANT_NAME() string`

GetTENANT_NAME returns the TENANT_NAME field if non-nil, zero value otherwise.

### GetTENANT_NAMEOk

`func (o *CreateOrUpdateRequest) GetTENANT_NAMEOk() (*string, bool)`

GetTENANT_NAMEOk returns a tuple with the TENANT_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_NAME

`func (o *CreateOrUpdateRequest) SetTENANT_NAME(v string)`

SetTENANT_NAME sets TENANT_NAME field to given value.

### HasTENANT_NAME

`func (o *CreateOrUpdateRequest) HasTENANT_NAME() bool`

HasTENANT_NAME returns a boolean if a field has been set.

### GetREPORT_OWNER

`func (o *CreateOrUpdateRequest) GetREPORT_OWNER() string`

GetREPORT_OWNER returns the REPORT_OWNER field if non-nil, zero value otherwise.

### GetREPORT_OWNEROk

`func (o *CreateOrUpdateRequest) GetREPORT_OWNEROk() (*string, bool)`

GetREPORT_OWNEROk returns a tuple with the REPORT_OWNER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREPORT_OWNER

`func (o *CreateOrUpdateRequest) SetREPORT_OWNER(v string)`

SetREPORT_OWNER sets REPORT_OWNER field to given value.

### HasREPORT_OWNER

`func (o *CreateOrUpdateRequest) HasREPORT_OWNER() bool`

HasREPORT_OWNER returns a boolean if a field has been set.

### GetUSE_OAUTH

`func (o *CreateOrUpdateRequest) GetUSE_OAUTH() string`

GetUSE_OAUTH returns the USE_OAUTH field if non-nil, zero value otherwise.

### GetUSE_OAUTHOk

`func (o *CreateOrUpdateRequest) GetUSE_OAUTHOk() (*string, bool)`

GetUSE_OAUTHOk returns a tuple with the USE_OAUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_OAUTH

`func (o *CreateOrUpdateRequest) SetUSE_OAUTH(v string)`

SetUSE_OAUTH sets USE_OAUTH field to given value.


### GetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *CreateOrUpdateRequest) GetINCLUDE_REFERENCE_DESCRIPTORS() string`

GetINCLUDE_REFERENCE_DESCRIPTORS returns the INCLUDE_REFERENCE_DESCRIPTORS field if non-nil, zero value otherwise.

### GetINCLUDE_REFERENCE_DESCRIPTORSOk

`func (o *CreateOrUpdateRequest) GetINCLUDE_REFERENCE_DESCRIPTORSOk() (*string, bool)`

GetINCLUDE_REFERENCE_DESCRIPTORSOk returns a tuple with the INCLUDE_REFERENCE_DESCRIPTORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *CreateOrUpdateRequest) SetINCLUDE_REFERENCE_DESCRIPTORS(v string)`

SetINCLUDE_REFERENCE_DESCRIPTORS sets INCLUDE_REFERENCE_DESCRIPTORS field to given value.

### HasINCLUDE_REFERENCE_DESCRIPTORS

`func (o *CreateOrUpdateRequest) HasINCLUDE_REFERENCE_DESCRIPTORS() bool`

HasINCLUDE_REFERENCE_DESCRIPTORS returns a boolean if a field has been set.

### GetUSE_ENHANCED_ORGROLE

`func (o *CreateOrUpdateRequest) GetUSE_ENHANCED_ORGROLE() string`

GetUSE_ENHANCED_ORGROLE returns the USE_ENHANCED_ORGROLE field if non-nil, zero value otherwise.

### GetUSE_ENHANCED_ORGROLEOk

`func (o *CreateOrUpdateRequest) GetUSE_ENHANCED_ORGROLEOk() (*string, bool)`

GetUSE_ENHANCED_ORGROLEOk returns a tuple with the USE_ENHANCED_ORGROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_ENHANCED_ORGROLE

`func (o *CreateOrUpdateRequest) SetUSE_ENHANCED_ORGROLE(v string)`

SetUSE_ENHANCED_ORGROLE sets USE_ENHANCED_ORGROLE field to given value.

### HasUSE_ENHANCED_ORGROLE

`func (o *CreateOrUpdateRequest) HasUSE_ENHANCED_ORGROLE() bool`

HasUSE_ENHANCED_ORGROLE returns a boolean if a field has been set.

### GetUSEX509AUTHFORSOAP

`func (o *CreateOrUpdateRequest) GetUSEX509AUTHFORSOAP() string`

GetUSEX509AUTHFORSOAP returns the USEX509AUTHFORSOAP field if non-nil, zero value otherwise.

### GetUSEX509AUTHFORSOAPOk

`func (o *CreateOrUpdateRequest) GetUSEX509AUTHFORSOAPOk() (*string, bool)`

GetUSEX509AUTHFORSOAPOk returns a tuple with the USEX509AUTHFORSOAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSEX509AUTHFORSOAP

`func (o *CreateOrUpdateRequest) SetUSEX509AUTHFORSOAP(v string)`

SetUSEX509AUTHFORSOAP sets USEX509AUTHFORSOAP field to given value.

### HasUSEX509AUTHFORSOAP

`func (o *CreateOrUpdateRequest) HasUSEX509AUTHFORSOAP() bool`

HasUSEX509AUTHFORSOAP returns a boolean if a field has been set.

### GetX509KEY

`func (o *CreateOrUpdateRequest) GetX509KEY() string`

GetX509KEY returns the X509KEY field if non-nil, zero value otherwise.

### GetX509KEYOk

`func (o *CreateOrUpdateRequest) GetX509KEYOk() (*string, bool)`

GetX509KEYOk returns a tuple with the X509KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509KEY

`func (o *CreateOrUpdateRequest) SetX509KEY(v string)`

SetX509KEY sets X509KEY field to given value.

### HasX509KEY

`func (o *CreateOrUpdateRequest) HasX509KEY() bool`

HasX509KEY returns a boolean if a field has been set.

### GetX509CERT

`func (o *CreateOrUpdateRequest) GetX509CERT() string`

GetX509CERT returns the X509CERT field if non-nil, zero value otherwise.

### GetX509CERTOk

`func (o *CreateOrUpdateRequest) GetX509CERTOk() (*string, bool)`

GetX509CERTOk returns a tuple with the X509CERT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509CERT

`func (o *CreateOrUpdateRequest) SetX509CERT(v string)`

SetX509CERT sets X509CERT field to given value.

### HasX509CERT

`func (o *CreateOrUpdateRequest) HasX509CERT() bool`

HasX509CERT returns a boolean if a field has been set.

### GetUSER_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) GetUSER_IMPORT_PAYLOAD() string`

GetUSER_IMPORT_PAYLOAD returns the USER_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetUSER_IMPORT_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetUSER_IMPORT_PAYLOADOk() (*string, bool)`

GetUSER_IMPORT_PAYLOADOk returns a tuple with the USER_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) SetUSER_IMPORT_PAYLOAD(v string)`

SetUSER_IMPORT_PAYLOAD sets USER_IMPORT_PAYLOAD field to given value.

### HasUSER_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) HasUSER_IMPORT_PAYLOAD() bool`

HasUSER_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) GetACCOUNT_IMPORT_PAYLOAD() string`

GetACCOUNT_IMPORT_PAYLOAD returns the ACCOUNT_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetACCOUNT_IMPORT_PAYLOADOk() (*string, bool)`

GetACCOUNT_IMPORT_PAYLOADOk returns a tuple with the ACCOUNT_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) SetACCOUNT_IMPORT_PAYLOAD(v string)`

SetACCOUNT_IMPORT_PAYLOAD sets ACCOUNT_IMPORT_PAYLOAD field to given value.

### HasACCOUNT_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) HasACCOUNT_IMPORT_PAYLOAD() bool`

HasACCOUNT_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetACCESS_IMPORT_LIST

`func (o *CreateOrUpdateRequest) GetACCESS_IMPORT_LIST() string`

GetACCESS_IMPORT_LIST returns the ACCESS_IMPORT_LIST field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_LISTOk

`func (o *CreateOrUpdateRequest) GetACCESS_IMPORT_LISTOk() (*string, bool)`

GetACCESS_IMPORT_LISTOk returns a tuple with the ACCESS_IMPORT_LIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_LIST

`func (o *CreateOrUpdateRequest) SetACCESS_IMPORT_LIST(v string)`

SetACCESS_IMPORT_LIST sets ACCESS_IMPORT_LIST field to given value.

### HasACCESS_IMPORT_LIST

`func (o *CreateOrUpdateRequest) HasACCESS_IMPORT_LIST() bool`

HasACCESS_IMPORT_LIST returns a boolean if a field has been set.

### GetRAAS_MAPPING_JSON

`func (o *CreateOrUpdateRequest) GetRAAS_MAPPING_JSON() string`

GetRAAS_MAPPING_JSON returns the RAAS_MAPPING_JSON field if non-nil, zero value otherwise.

### GetRAAS_MAPPING_JSONOk

`func (o *CreateOrUpdateRequest) GetRAAS_MAPPING_JSONOk() (*string, bool)`

GetRAAS_MAPPING_JSONOk returns a tuple with the RAAS_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAAS_MAPPING_JSON

`func (o *CreateOrUpdateRequest) SetRAAS_MAPPING_JSON(v string)`

SetRAAS_MAPPING_JSON sets RAAS_MAPPING_JSON field to given value.

### HasRAAS_MAPPING_JSON

`func (o *CreateOrUpdateRequest) HasRAAS_MAPPING_JSON() bool`

HasRAAS_MAPPING_JSON returns a boolean if a field has been set.

### GetACCESS_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) GetACCESS_IMPORT_MAPPING() string`

GetACCESS_IMPORT_MAPPING returns the ACCESS_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_MAPPINGOk

`func (o *CreateOrUpdateRequest) GetACCESS_IMPORT_MAPPINGOk() (*string, bool)`

GetACCESS_IMPORT_MAPPINGOk returns a tuple with the ACCESS_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) SetACCESS_IMPORT_MAPPING(v string)`

SetACCESS_IMPORT_MAPPING sets ACCESS_IMPORT_MAPPING field to given value.

### HasACCESS_IMPORT_MAPPING

`func (o *CreateOrUpdateRequest) HasACCESS_IMPORT_MAPPING() bool`

HasACCESS_IMPORT_MAPPING returns a boolean if a field has been set.

### GetORGROLE_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) GetORGROLE_IMPORT_PAYLOAD() string`

GetORGROLE_IMPORT_PAYLOAD returns the ORGROLE_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetORGROLE_IMPORT_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetORGROLE_IMPORT_PAYLOADOk() (*string, bool)`

GetORGROLE_IMPORT_PAYLOADOk returns a tuple with the ORGROLE_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGROLE_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) SetORGROLE_IMPORT_PAYLOAD(v string)`

SetORGROLE_IMPORT_PAYLOAD sets ORGROLE_IMPORT_PAYLOAD field to given value.

### HasORGROLE_IMPORT_PAYLOAD

`func (o *CreateOrUpdateRequest) HasORGROLE_IMPORT_PAYLOAD() bool`

HasORGROLE_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetCREATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) GetCREATE_ACCOUNT_PAYLOAD() string`

GetCREATE_ACCOUNT_PAYLOAD returns the CREATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetCREATE_ACCOUNT_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetCREATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetCREATE_ACCOUNT_PAYLOADOk returns a tuple with the CREATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) SetCREATE_ACCOUNT_PAYLOAD(v string)`

SetCREATE_ACCOUNT_PAYLOAD sets CREATE_ACCOUNT_PAYLOAD field to given value.

### HasCREATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) HasCREATE_ACCOUNT_PAYLOAD() bool`

HasCREATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetUPDATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) GetUPDATE_ACCOUNT_PAYLOAD() string`

GetUPDATE_ACCOUNT_PAYLOAD returns the UPDATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_ACCOUNT_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetUPDATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetUPDATE_ACCOUNT_PAYLOADOk returns a tuple with the UPDATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) SetUPDATE_ACCOUNT_PAYLOAD(v string)`

SetUPDATE_ACCOUNT_PAYLOAD sets UPDATE_ACCOUNT_PAYLOAD field to given value.

### HasUPDATE_ACCOUNT_PAYLOAD

`func (o *CreateOrUpdateRequest) HasUPDATE_ACCOUNT_PAYLOAD() bool`

HasUPDATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetUPDATE_USER_PAYLOAD

`func (o *CreateOrUpdateRequest) GetUPDATE_USER_PAYLOAD() string`

GetUPDATE_USER_PAYLOAD returns the UPDATE_USER_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_USER_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetUPDATE_USER_PAYLOADOk() (*string, bool)`

GetUPDATE_USER_PAYLOADOk returns a tuple with the UPDATE_USER_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USER_PAYLOAD

`func (o *CreateOrUpdateRequest) SetUPDATE_USER_PAYLOAD(v string)`

SetUPDATE_USER_PAYLOAD sets UPDATE_USER_PAYLOAD field to given value.

### HasUPDATE_USER_PAYLOAD

`func (o *CreateOrUpdateRequest) HasUPDATE_USER_PAYLOAD() bool`

HasUPDATE_USER_PAYLOAD returns a boolean if a field has been set.

### GetASSIGN_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) GetASSIGN_ORGROLE_PAYLOAD() string`

GetASSIGN_ORGROLE_PAYLOAD returns the ASSIGN_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetASSIGN_ORGROLE_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetASSIGN_ORGROLE_PAYLOADOk() (*string, bool)`

GetASSIGN_ORGROLE_PAYLOADOk returns a tuple with the ASSIGN_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASSIGN_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) SetASSIGN_ORGROLE_PAYLOAD(v string)`

SetASSIGN_ORGROLE_PAYLOAD sets ASSIGN_ORGROLE_PAYLOAD field to given value.

### HasASSIGN_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) HasASSIGN_ORGROLE_PAYLOAD() bool`

HasASSIGN_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetREMOVE_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) GetREMOVE_ORGROLE_PAYLOAD() string`

GetREMOVE_ORGROLE_PAYLOAD returns the REMOVE_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetREMOVE_ORGROLE_PAYLOADOk

`func (o *CreateOrUpdateRequest) GetREMOVE_ORGROLE_PAYLOADOk() (*string, bool)`

GetREMOVE_ORGROLE_PAYLOADOk returns a tuple with the REMOVE_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) SetREMOVE_ORGROLE_PAYLOAD(v string)`

SetREMOVE_ORGROLE_PAYLOAD sets REMOVE_ORGROLE_PAYLOAD field to given value.

### HasREMOVE_ORGROLE_PAYLOAD

`func (o *CreateOrUpdateRequest) HasREMOVE_ORGROLE_PAYLOAD() bool`

HasREMOVE_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetSTATUS_KEY_JSON

`func (o *CreateOrUpdateRequest) GetSTATUS_KEY_JSON() string`

GetSTATUS_KEY_JSON returns the STATUS_KEY_JSON field if non-nil, zero value otherwise.

### GetSTATUS_KEY_JSONOk

`func (o *CreateOrUpdateRequest) GetSTATUS_KEY_JSONOk() (*string, bool)`

GetSTATUS_KEY_JSONOk returns a tuple with the STATUS_KEY_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_KEY_JSON

`func (o *CreateOrUpdateRequest) SetSTATUS_KEY_JSON(v string)`

SetSTATUS_KEY_JSON sets STATUS_KEY_JSON field to given value.

### HasSTATUS_KEY_JSON

`func (o *CreateOrUpdateRequest) HasSTATUS_KEY_JSON() bool`

HasSTATUS_KEY_JSON returns a boolean if a field has been set.

### GetUSERATTRIBUTEJSON

`func (o *CreateOrUpdateRequest) GetUSERATTRIBUTEJSON() string`

GetUSERATTRIBUTEJSON returns the USERATTRIBUTEJSON field if non-nil, zero value otherwise.

### GetUSERATTRIBUTEJSONOk

`func (o *CreateOrUpdateRequest) GetUSERATTRIBUTEJSONOk() (*string, bool)`

GetUSERATTRIBUTEJSONOk returns a tuple with the USERATTRIBUTEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERATTRIBUTEJSON

`func (o *CreateOrUpdateRequest) SetUSERATTRIBUTEJSON(v string)`

SetUSERATTRIBUTEJSON sets USERATTRIBUTEJSON field to given value.

### HasUSERATTRIBUTEJSON

`func (o *CreateOrUpdateRequest) HasUSERATTRIBUTEJSON() bool`

HasUSERATTRIBUTEJSON returns a boolean if a field has been set.

### GetCUSTOM_CONFIG

`func (o *CreateOrUpdateRequest) GetCUSTOM_CONFIG() string`

GetCUSTOM_CONFIG returns the CUSTOM_CONFIG field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIGOk

`func (o *CreateOrUpdateRequest) GetCUSTOM_CONFIGOk() (*string, bool)`

GetCUSTOM_CONFIGOk returns a tuple with the CUSTOM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG

`func (o *CreateOrUpdateRequest) SetCUSTOM_CONFIG(v string)`

SetCUSTOM_CONFIG sets CUSTOM_CONFIG field to given value.

### HasCUSTOM_CONFIG

`func (o *CreateOrUpdateRequest) HasCUSTOM_CONFIG() bool`

HasCUSTOM_CONFIG returns a boolean if a field has been set.

### GetDRIVERNAME

`func (o *CreateOrUpdateRequest) GetDRIVERNAME() string`

GetDRIVERNAME returns the DRIVERNAME field if non-nil, zero value otherwise.

### GetDRIVERNAMEOk

`func (o *CreateOrUpdateRequest) GetDRIVERNAMEOk() (*string, bool)`

GetDRIVERNAMEOk returns a tuple with the DRIVERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRIVERNAME

`func (o *CreateOrUpdateRequest) SetDRIVERNAME(v string)`

SetDRIVERNAME sets DRIVERNAME field to given value.


### GetCONNECTIONPROPERTIES

`func (o *CreateOrUpdateRequest) GetCONNECTIONPROPERTIES() string`

GetCONNECTIONPROPERTIES returns the CONNECTIONPROPERTIES field if non-nil, zero value otherwise.

### GetCONNECTIONPROPERTIESOk

`func (o *CreateOrUpdateRequest) GetCONNECTIONPROPERTIESOk() (*string, bool)`

GetCONNECTIONPROPERTIESOk returns a tuple with the CONNECTIONPROPERTIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONPROPERTIES

`func (o *CreateOrUpdateRequest) SetCONNECTIONPROPERTIES(v string)`

SetCONNECTIONPROPERTIES sets CONNECTIONPROPERTIES field to given value.

### HasCONNECTIONPROPERTIES

`func (o *CreateOrUpdateRequest) HasCONNECTIONPROPERTIES() bool`

HasCONNECTIONPROPERTIES returns a boolean if a field has been set.

### GetGRANTACCESSJSON

`func (o *CreateOrUpdateRequest) GetGRANTACCESSJSON() string`

GetGRANTACCESSJSON returns the GRANTACCESSJSON field if non-nil, zero value otherwise.

### GetGRANTACCESSJSONOk

`func (o *CreateOrUpdateRequest) GetGRANTACCESSJSONOk() (*string, bool)`

GetGRANTACCESSJSONOk returns a tuple with the GRANTACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRANTACCESSJSON

`func (o *CreateOrUpdateRequest) SetGRANTACCESSJSON(v string)`

SetGRANTACCESSJSON sets GRANTACCESSJSON field to given value.

### HasGRANTACCESSJSON

`func (o *CreateOrUpdateRequest) HasGRANTACCESSJSON() bool`

HasGRANTACCESSJSON returns a boolean if a field has been set.

### GetREVOKEACCESSJSON

`func (o *CreateOrUpdateRequest) GetREVOKEACCESSJSON() string`

GetREVOKEACCESSJSON returns the REVOKEACCESSJSON field if non-nil, zero value otherwise.

### GetREVOKEACCESSJSONOk

`func (o *CreateOrUpdateRequest) GetREVOKEACCESSJSONOk() (*string, bool)`

GetREVOKEACCESSJSONOk returns a tuple with the REVOKEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREVOKEACCESSJSON

`func (o *CreateOrUpdateRequest) SetREVOKEACCESSJSON(v string)`

SetREVOKEACCESSJSON sets REVOKEACCESSJSON field to given value.

### HasREVOKEACCESSJSON

`func (o *CreateOrUpdateRequest) HasREVOKEACCESSJSON() bool`

HasREVOKEACCESSJSON returns a boolean if a field has been set.

### GetCHANGEPASSJSON

`func (o *CreateOrUpdateRequest) GetCHANGEPASSJSON() string`

GetCHANGEPASSJSON returns the CHANGEPASSJSON field if non-nil, zero value otherwise.

### GetCHANGEPASSJSONOk

`func (o *CreateOrUpdateRequest) GetCHANGEPASSJSONOk() (*string, bool)`

GetCHANGEPASSJSONOk returns a tuple with the CHANGEPASSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGEPASSJSON

`func (o *CreateOrUpdateRequest) SetCHANGEPASSJSON(v string)`

SetCHANGEPASSJSON sets CHANGEPASSJSON field to given value.

### HasCHANGEPASSJSON

`func (o *CreateOrUpdateRequest) HasCHANGEPASSJSON() bool`

HasCHANGEPASSJSON returns a boolean if a field has been set.

### GetDELETEACCOUNTJSON

`func (o *CreateOrUpdateRequest) GetDELETEACCOUNTJSON() string`

GetDELETEACCOUNTJSON returns the DELETEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDELETEACCOUNTJSONOk

`func (o *CreateOrUpdateRequest) GetDELETEACCOUNTJSONOk() (*string, bool)`

GetDELETEACCOUNTJSONOk returns a tuple with the DELETEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEACCOUNTJSON

`func (o *CreateOrUpdateRequest) SetDELETEACCOUNTJSON(v string)`

SetDELETEACCOUNTJSON sets DELETEACCOUNTJSON field to given value.

### HasDELETEACCOUNTJSON

`func (o *CreateOrUpdateRequest) HasDELETEACCOUNTJSON() bool`

HasDELETEACCOUNTJSON returns a boolean if a field has been set.

### GetACCOUNTEXISTSJSON

`func (o *CreateOrUpdateRequest) GetACCOUNTEXISTSJSON() string`

GetACCOUNTEXISTSJSON returns the ACCOUNTEXISTSJSON field if non-nil, zero value otherwise.

### GetACCOUNTEXISTSJSONOk

`func (o *CreateOrUpdateRequest) GetACCOUNTEXISTSJSONOk() (*string, bool)`

GetACCOUNTEXISTSJSONOk returns a tuple with the ACCOUNTEXISTSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTEXISTSJSON

`func (o *CreateOrUpdateRequest) SetACCOUNTEXISTSJSON(v string)`

SetACCOUNTEXISTSJSON sets ACCOUNTEXISTSJSON field to given value.

### HasACCOUNTEXISTSJSON

`func (o *CreateOrUpdateRequest) HasACCOUNTEXISTSJSON() bool`

HasACCOUNTEXISTSJSON returns a boolean if a field has been set.

### GetACCOUNTSIMPORT

`func (o *CreateOrUpdateRequest) GetACCOUNTSIMPORT() string`

GetACCOUNTSIMPORT returns the ACCOUNTSIMPORT field if non-nil, zero value otherwise.

### GetACCOUNTSIMPORTOk

`func (o *CreateOrUpdateRequest) GetACCOUNTSIMPORTOk() (*string, bool)`

GetACCOUNTSIMPORTOk returns a tuple with the ACCOUNTSIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTSIMPORT

`func (o *CreateOrUpdateRequest) SetACCOUNTSIMPORT(v string)`

SetACCOUNTSIMPORT sets ACCOUNTSIMPORT field to given value.

### HasACCOUNTSIMPORT

`func (o *CreateOrUpdateRequest) HasACCOUNTSIMPORT() bool`

HasACCOUNTSIMPORT returns a boolean if a field has been set.

### GetENTITLEMENTVALUEIMPORT

`func (o *CreateOrUpdateRequest) GetENTITLEMENTVALUEIMPORT() string`

GetENTITLEMENTVALUEIMPORT returns the ENTITLEMENTVALUEIMPORT field if non-nil, zero value otherwise.

### GetENTITLEMENTVALUEIMPORTOk

`func (o *CreateOrUpdateRequest) GetENTITLEMENTVALUEIMPORTOk() (*string, bool)`

GetENTITLEMENTVALUEIMPORTOk returns a tuple with the ENTITLEMENTVALUEIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTVALUEIMPORT

`func (o *CreateOrUpdateRequest) SetENTITLEMENTVALUEIMPORT(v string)`

SetENTITLEMENTVALUEIMPORT sets ENTITLEMENTVALUEIMPORT field to given value.

### HasENTITLEMENTVALUEIMPORT

`func (o *CreateOrUpdateRequest) HasENTITLEMENTVALUEIMPORT() bool`

HasENTITLEMENTVALUEIMPORT returns a boolean if a field has been set.

### GetROLEOWNERIMPORT

`func (o *CreateOrUpdateRequest) GetROLEOWNERIMPORT() string`

GetROLEOWNERIMPORT returns the ROLEOWNERIMPORT field if non-nil, zero value otherwise.

### GetROLEOWNERIMPORTOk

`func (o *CreateOrUpdateRequest) GetROLEOWNERIMPORTOk() (*string, bool)`

GetROLEOWNERIMPORTOk returns a tuple with the ROLEOWNERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLEOWNERIMPORT

`func (o *CreateOrUpdateRequest) SetROLEOWNERIMPORT(v string)`

SetROLEOWNERIMPORT sets ROLEOWNERIMPORT field to given value.

### HasROLEOWNERIMPORT

`func (o *CreateOrUpdateRequest) HasROLEOWNERIMPORT() bool`

HasROLEOWNERIMPORT returns a boolean if a field has been set.

### GetROLESIMPORT

`func (o *CreateOrUpdateRequest) GetROLESIMPORT() string`

GetROLESIMPORT returns the ROLESIMPORT field if non-nil, zero value otherwise.

### GetROLESIMPORTOk

`func (o *CreateOrUpdateRequest) GetROLESIMPORTOk() (*string, bool)`

GetROLESIMPORTOk returns a tuple with the ROLESIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLESIMPORT

`func (o *CreateOrUpdateRequest) SetROLESIMPORT(v string)`

SetROLESIMPORT sets ROLESIMPORT field to given value.

### HasROLESIMPORT

`func (o *CreateOrUpdateRequest) HasROLESIMPORT() bool`

HasROLESIMPORT returns a boolean if a field has been set.

### GetSYSTEMIMPORT

`func (o *CreateOrUpdateRequest) GetSYSTEMIMPORT() string`

GetSYSTEMIMPORT returns the SYSTEMIMPORT field if non-nil, zero value otherwise.

### GetSYSTEMIMPORTOk

`func (o *CreateOrUpdateRequest) GetSYSTEMIMPORTOk() (*string, bool)`

GetSYSTEMIMPORTOk returns a tuple with the SYSTEMIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMIMPORT

`func (o *CreateOrUpdateRequest) SetSYSTEMIMPORT(v string)`

SetSYSTEMIMPORT sets SYSTEMIMPORT field to given value.

### HasSYSTEMIMPORT

`func (o *CreateOrUpdateRequest) HasSYSTEMIMPORT() bool`

HasSYSTEMIMPORT returns a boolean if a field has been set.

### GetUSERIMPORT

`func (o *CreateOrUpdateRequest) GetUSERIMPORT() string`

GetUSERIMPORT returns the USERIMPORT field if non-nil, zero value otherwise.

### GetUSERIMPORTOk

`func (o *CreateOrUpdateRequest) GetUSERIMPORTOk() (*string, bool)`

GetUSERIMPORTOk returns a tuple with the USERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORT

`func (o *CreateOrUpdateRequest) SetUSERIMPORT(v string)`

SetUSERIMPORT sets USERIMPORT field to given value.

### HasUSERIMPORT

`func (o *CreateOrUpdateRequest) HasUSERIMPORT() bool`

HasUSERIMPORT returns a boolean if a field has been set.

### GetMAX_PAGINATION_SIZE

`func (o *CreateOrUpdateRequest) GetMAX_PAGINATION_SIZE() string`

GetMAX_PAGINATION_SIZE returns the MAX_PAGINATION_SIZE field if non-nil, zero value otherwise.

### GetMAX_PAGINATION_SIZEOk

`func (o *CreateOrUpdateRequest) GetMAX_PAGINATION_SIZEOk() (*string, bool)`

GetMAX_PAGINATION_SIZEOk returns a tuple with the MAX_PAGINATION_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_PAGINATION_SIZE

`func (o *CreateOrUpdateRequest) SetMAX_PAGINATION_SIZE(v string)`

SetMAX_PAGINATION_SIZE sets MAX_PAGINATION_SIZE field to given value.

### HasMAX_PAGINATION_SIZE

`func (o *CreateOrUpdateRequest) HasMAX_PAGINATION_SIZE() bool`

HasMAX_PAGINATION_SIZE returns a boolean if a field has been set.

### GetCLI_COMMAND_JSON

`func (o *CreateOrUpdateRequest) GetCLI_COMMAND_JSON() string`

GetCLI_COMMAND_JSON returns the CLI_COMMAND_JSON field if non-nil, zero value otherwise.

### GetCLI_COMMAND_JSONOk

`func (o *CreateOrUpdateRequest) GetCLI_COMMAND_JSONOk() (*string, bool)`

GetCLI_COMMAND_JSONOk returns a tuple with the CLI_COMMAND_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLI_COMMAND_JSON

`func (o *CreateOrUpdateRequest) SetCLI_COMMAND_JSON(v string)`

SetCLI_COMMAND_JSON sets CLI_COMMAND_JSON field to given value.

### HasCLI_COMMAND_JSON

`func (o *CreateOrUpdateRequest) HasCLI_COMMAND_JSON() bool`

HasCLI_COMMAND_JSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


