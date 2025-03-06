# UpdateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessquery** | Pointer to **string** | Access query | [optional] 
**AccountNameRule** | Pointer to **string** |  | [optional] 
**BlockInflightRequest** | Pointer to **string** | Flag to block inflight requests | [optional] 
**Connectionconfig** | Pointer to **string** | Connection configuration in JSON format (escaped). Supports valid case-senstive string values such as “None”, “DropDownSingle”, “Table”, and “TableOnlyAdd” | [optional] 
**CreateEntTaskforRemoveAcc** | Pointer to **string** | Flag to create enterprise task for remove account | [optional] 
**Description** | Pointer to **string** | Description of the endpoint | [optional] 
**DisableModifyAccount** | Pointer to **string** | Flag to disable modify account functionality | [optional] 
**DisableNewAccountRequestIfAccountExists** | Pointer to **string** | Flag to disable new account request if account exists | [optional] 
**DisableRemoveAccount** | Pointer to **string** | Flag to disable remove account functionality | [optional] 
**DisplayName** | Pointer to **string** | Name displayed in the user interface | [optional] 
**EnableCopyAccess** | Pointer to **string** | Flag to enable copy access | [optional] 
**Endpointname** | **string** | Name of the endpoint to update | 
**Owner** | Pointer to **string** | Owner identifier | [optional] 
**OwnerType** | Pointer to **string** | Type of owner | [optional] 
**RequestOption** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **string** | Supports valid boolean values are true and falserequestedQuery/selectedQuery Supports string values in double quotes | [optional] 
**ResourceOwner** | Pointer to **string** | Resource owner identifier | [optional] 
**ResourceOwnerType** | Pointer to **string** | Type of resource owner | [optional] 
**RequestableRoleType** | Pointer to [**[]UpdateEndpointRequestRequestableRoleTypeInner**](UpdateEndpointRequestRequestableRoleTypeInner.md) | Configuration for requestable role types (available from v23.9) | [optional] 
**RoleType** | Pointer to **string** | Supports valid case-senstive string values such as “Enabler”, “Transactional”, “EmergencyAccess”, “Enterprise”, and “Application” | [optional] 
**Securitysystem** | Pointer to **string** | Security system that encapsulates the endpoint | [optional] 
**ShowOn** | Pointer to **string** | Supports valid case-sensitive string values such as “All”, “ShowOnApplicationRequest”, and “ShowOnServiceAccountRequest”. specify the value to display the roles on Application Requests, Service Account Requests, or both. This parameter is available from the Release v24.6 onwards. | [optional] 
**UserAccountCorrelationRule** | Pointer to **string** | Rule for correlating users to accounts | [optional] 

## Methods

### NewUpdateEndpointRequest

`func NewUpdateEndpointRequest(endpointname string, ) *UpdateEndpointRequest`

NewUpdateEndpointRequest instantiates a new UpdateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEndpointRequestWithDefaults

`func NewUpdateEndpointRequestWithDefaults() *UpdateEndpointRequest`

NewUpdateEndpointRequestWithDefaults instantiates a new UpdateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessquery

`func (o *UpdateEndpointRequest) GetAccessquery() string`

GetAccessquery returns the Accessquery field if non-nil, zero value otherwise.

### GetAccessqueryOk

`func (o *UpdateEndpointRequest) GetAccessqueryOk() (*string, bool)`

GetAccessqueryOk returns a tuple with the Accessquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessquery

`func (o *UpdateEndpointRequest) SetAccessquery(v string)`

SetAccessquery sets Accessquery field to given value.

### HasAccessquery

`func (o *UpdateEndpointRequest) HasAccessquery() bool`

HasAccessquery returns a boolean if a field has been set.

### GetAccountNameRule

`func (o *UpdateEndpointRequest) GetAccountNameRule() string`

GetAccountNameRule returns the AccountNameRule field if non-nil, zero value otherwise.

### GetAccountNameRuleOk

`func (o *UpdateEndpointRequest) GetAccountNameRuleOk() (*string, bool)`

GetAccountNameRuleOk returns a tuple with the AccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameRule

`func (o *UpdateEndpointRequest) SetAccountNameRule(v string)`

SetAccountNameRule sets AccountNameRule field to given value.

### HasAccountNameRule

`func (o *UpdateEndpointRequest) HasAccountNameRule() bool`

HasAccountNameRule returns a boolean if a field has been set.

### GetBlockInflightRequest

`func (o *UpdateEndpointRequest) GetBlockInflightRequest() string`

GetBlockInflightRequest returns the BlockInflightRequest field if non-nil, zero value otherwise.

### GetBlockInflightRequestOk

`func (o *UpdateEndpointRequest) GetBlockInflightRequestOk() (*string, bool)`

GetBlockInflightRequestOk returns a tuple with the BlockInflightRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInflightRequest

`func (o *UpdateEndpointRequest) SetBlockInflightRequest(v string)`

SetBlockInflightRequest sets BlockInflightRequest field to given value.

### HasBlockInflightRequest

`func (o *UpdateEndpointRequest) HasBlockInflightRequest() bool`

HasBlockInflightRequest returns a boolean if a field has been set.

### GetConnectionconfig

`func (o *UpdateEndpointRequest) GetConnectionconfig() string`

GetConnectionconfig returns the Connectionconfig field if non-nil, zero value otherwise.

### GetConnectionconfigOk

`func (o *UpdateEndpointRequest) GetConnectionconfigOk() (*string, bool)`

GetConnectionconfigOk returns a tuple with the Connectionconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionconfig

`func (o *UpdateEndpointRequest) SetConnectionconfig(v string)`

SetConnectionconfig sets Connectionconfig field to given value.

### HasConnectionconfig

`func (o *UpdateEndpointRequest) HasConnectionconfig() bool`

HasConnectionconfig returns a boolean if a field has been set.

### GetCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAcc() string`

GetCreateEntTaskforRemoveAcc returns the CreateEntTaskforRemoveAcc field if non-nil, zero value otherwise.

### GetCreateEntTaskforRemoveAccOk

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAccOk() (*string, bool)`

GetCreateEntTaskforRemoveAccOk returns a tuple with the CreateEntTaskforRemoveAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) SetCreateEntTaskforRemoveAcc(v string)`

SetCreateEntTaskforRemoveAcc sets CreateEntTaskforRemoveAcc field to given value.

### HasCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) HasCreateEntTaskforRemoveAcc() bool`

HasCreateEntTaskforRemoveAcc returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableModifyAccount

`func (o *UpdateEndpointRequest) GetDisableModifyAccount() string`

GetDisableModifyAccount returns the DisableModifyAccount field if non-nil, zero value otherwise.

### GetDisableModifyAccountOk

`func (o *UpdateEndpointRequest) GetDisableModifyAccountOk() (*string, bool)`

GetDisableModifyAccountOk returns a tuple with the DisableModifyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModifyAccount

`func (o *UpdateEndpointRequest) SetDisableModifyAccount(v string)`

SetDisableModifyAccount sets DisableModifyAccount field to given value.

### HasDisableModifyAccount

`func (o *UpdateEndpointRequest) HasDisableModifyAccount() bool`

HasDisableModifyAccount returns a boolean if a field has been set.

### GetDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) GetDisableNewAccountRequestIfAccountExists() string`

GetDisableNewAccountRequestIfAccountExists returns the DisableNewAccountRequestIfAccountExists field if non-nil, zero value otherwise.

### GetDisableNewAccountRequestIfAccountExistsOk

`func (o *UpdateEndpointRequest) GetDisableNewAccountRequestIfAccountExistsOk() (*string, bool)`

GetDisableNewAccountRequestIfAccountExistsOk returns a tuple with the DisableNewAccountRequestIfAccountExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) SetDisableNewAccountRequestIfAccountExists(v string)`

SetDisableNewAccountRequestIfAccountExists sets DisableNewAccountRequestIfAccountExists field to given value.

### HasDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) HasDisableNewAccountRequestIfAccountExists() bool`

HasDisableNewAccountRequestIfAccountExists returns a boolean if a field has been set.

### GetDisableRemoveAccount

`func (o *UpdateEndpointRequest) GetDisableRemoveAccount() string`

GetDisableRemoveAccount returns the DisableRemoveAccount field if non-nil, zero value otherwise.

### GetDisableRemoveAccountOk

`func (o *UpdateEndpointRequest) GetDisableRemoveAccountOk() (*string, bool)`

GetDisableRemoveAccountOk returns a tuple with the DisableRemoveAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRemoveAccount

`func (o *UpdateEndpointRequest) SetDisableRemoveAccount(v string)`

SetDisableRemoveAccount sets DisableRemoveAccount field to given value.

### HasDisableRemoveAccount

`func (o *UpdateEndpointRequest) HasDisableRemoveAccount() bool`

HasDisableRemoveAccount returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateEndpointRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateEndpointRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateEndpointRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateEndpointRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableCopyAccess

`func (o *UpdateEndpointRequest) GetEnableCopyAccess() string`

GetEnableCopyAccess returns the EnableCopyAccess field if non-nil, zero value otherwise.

### GetEnableCopyAccessOk

`func (o *UpdateEndpointRequest) GetEnableCopyAccessOk() (*string, bool)`

GetEnableCopyAccessOk returns a tuple with the EnableCopyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyAccess

`func (o *UpdateEndpointRequest) SetEnableCopyAccess(v string)`

SetEnableCopyAccess sets EnableCopyAccess field to given value.

### HasEnableCopyAccess

`func (o *UpdateEndpointRequest) HasEnableCopyAccess() bool`

HasEnableCopyAccess returns a boolean if a field has been set.

### GetEndpointname

`func (o *UpdateEndpointRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *UpdateEndpointRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *UpdateEndpointRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.


### GetOwner

`func (o *UpdateEndpointRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateEndpointRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateEndpointRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateEndpointRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerType

`func (o *UpdateEndpointRequest) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *UpdateEndpointRequest) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *UpdateEndpointRequest) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *UpdateEndpointRequest) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetRequestOption

`func (o *UpdateEndpointRequest) GetRequestOption() string`

GetRequestOption returns the RequestOption field if non-nil, zero value otherwise.

### GetRequestOptionOk

`func (o *UpdateEndpointRequest) GetRequestOptionOk() (*string, bool)`

GetRequestOptionOk returns a tuple with the RequestOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOption

`func (o *UpdateEndpointRequest) SetRequestOption(v string)`

SetRequestOption sets RequestOption field to given value.

### HasRequestOption

`func (o *UpdateEndpointRequest) HasRequestOption() bool`

HasRequestOption returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateEndpointRequest) GetRequired() string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateEndpointRequest) GetRequiredOk() (*string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateEndpointRequest) SetRequired(v string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateEndpointRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetResourceOwner

`func (o *UpdateEndpointRequest) GetResourceOwner() string`

GetResourceOwner returns the ResourceOwner field if non-nil, zero value otherwise.

### GetResourceOwnerOk

`func (o *UpdateEndpointRequest) GetResourceOwnerOk() (*string, bool)`

GetResourceOwnerOk returns a tuple with the ResourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwner

`func (o *UpdateEndpointRequest) SetResourceOwner(v string)`

SetResourceOwner sets ResourceOwner field to given value.

### HasResourceOwner

`func (o *UpdateEndpointRequest) HasResourceOwner() bool`

HasResourceOwner returns a boolean if a field has been set.

### GetResourceOwnerType

`func (o *UpdateEndpointRequest) GetResourceOwnerType() string`

GetResourceOwnerType returns the ResourceOwnerType field if non-nil, zero value otherwise.

### GetResourceOwnerTypeOk

`func (o *UpdateEndpointRequest) GetResourceOwnerTypeOk() (*string, bool)`

GetResourceOwnerTypeOk returns a tuple with the ResourceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerType

`func (o *UpdateEndpointRequest) SetResourceOwnerType(v string)`

SetResourceOwnerType sets ResourceOwnerType field to given value.

### HasResourceOwnerType

`func (o *UpdateEndpointRequest) HasResourceOwnerType() bool`

HasResourceOwnerType returns a boolean if a field has been set.

### GetRequestableRoleType

`func (o *UpdateEndpointRequest) GetRequestableRoleType() []UpdateEndpointRequestRequestableRoleTypeInner`

GetRequestableRoleType returns the RequestableRoleType field if non-nil, zero value otherwise.

### GetRequestableRoleTypeOk

`func (o *UpdateEndpointRequest) GetRequestableRoleTypeOk() (*[]UpdateEndpointRequestRequestableRoleTypeInner, bool)`

GetRequestableRoleTypeOk returns a tuple with the RequestableRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableRoleType

`func (o *UpdateEndpointRequest) SetRequestableRoleType(v []UpdateEndpointRequestRequestableRoleTypeInner)`

SetRequestableRoleType sets RequestableRoleType field to given value.

### HasRequestableRoleType

`func (o *UpdateEndpointRequest) HasRequestableRoleType() bool`

HasRequestableRoleType returns a boolean if a field has been set.

### GetRoleType

`func (o *UpdateEndpointRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *UpdateEndpointRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *UpdateEndpointRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *UpdateEndpointRequest) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *UpdateEndpointRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *UpdateEndpointRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *UpdateEndpointRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *UpdateEndpointRequest) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetShowOn

`func (o *UpdateEndpointRequest) GetShowOn() string`

GetShowOn returns the ShowOn field if non-nil, zero value otherwise.

### GetShowOnOk

`func (o *UpdateEndpointRequest) GetShowOnOk() (*string, bool)`

GetShowOnOk returns a tuple with the ShowOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOn

`func (o *UpdateEndpointRequest) SetShowOn(v string)`

SetShowOn sets ShowOn field to given value.

### HasShowOn

`func (o *UpdateEndpointRequest) HasShowOn() bool`

HasShowOn returns a boolean if a field has been set.

### GetUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) GetUserAccountCorrelationRule() string`

GetUserAccountCorrelationRule returns the UserAccountCorrelationRule field if non-nil, zero value otherwise.

### GetUserAccountCorrelationRuleOk

`func (o *UpdateEndpointRequest) GetUserAccountCorrelationRuleOk() (*string, bool)`

GetUserAccountCorrelationRuleOk returns a tuple with the UserAccountCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) SetUserAccountCorrelationRule(v string)`

SetUserAccountCorrelationRule sets UserAccountCorrelationRule field to given value.

### HasUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) HasUserAccountCorrelationRule() bool`

HasUserAccountCorrelationRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


