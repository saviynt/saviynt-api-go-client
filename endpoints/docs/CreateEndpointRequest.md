# CreateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessquery** | Pointer to **string** | Access query | [optional] 
**AccountNameRule** | Pointer to **string** | Account name rule | [optional] 
**BlockInflightRequest** | Pointer to **string** | Flag to block inflight requests | [optional] 
**Connectionconfig** | Pointer to **string** | Connection configuration in JSON format (escaped) | [optional] 
**Description** | Pointer to **string** | Description of the endpoint | [optional] 
**DisableNewAccountRequestIfAccountExists** | Pointer to **string** | Flag to disable new account request if account exists | [optional] 
**DisableModifyAccount** | Pointer to **string** | Flag to disable modify account functionality | [optional] 
**DisableRemoveAccount** | Pointer to **string** | Flag to disable remove account functionality | [optional] 
**DisplayName** | **string** | Name displayed in the user interface (can be different from endpoint name) | 
**EnableCopyAccess** | Pointer to **string** | Flag to enable copy access | [optional] 
**Endpointname** | **string** | Logical name for the endpoint that helps identify it easily | 
**Owner** | Pointer to **string** | Owner identifier | [optional] 
**OwnerType** | Pointer to **string** | Type of owner (e.g., User, Usergroup) | [optional] 
**ResourceOwnerType** | Pointer to **string** | Type of resource owner | [optional] 
**ResourceOwner** | Pointer to **string** | Resource owner identifier | [optional] 
**Securitysystem** | **string** | The security system for which you want to create an endpoint. The security system encapsulates the endpoint along with other endpoints sharing the same connections, workflows, or more. | 
**UserAccountCorrelationRule** | Pointer to **string** | Rule for correlating users to accounts | [optional] 

## Methods

### NewCreateEndpointRequest

`func NewCreateEndpointRequest(displayName string, endpointname string, securitysystem string, ) *CreateEndpointRequest`

NewCreateEndpointRequest instantiates a new CreateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointRequestWithDefaults

`func NewCreateEndpointRequestWithDefaults() *CreateEndpointRequest`

NewCreateEndpointRequestWithDefaults instantiates a new CreateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessquery

`func (o *CreateEndpointRequest) GetAccessquery() string`

GetAccessquery returns the Accessquery field if non-nil, zero value otherwise.

### GetAccessqueryOk

`func (o *CreateEndpointRequest) GetAccessqueryOk() (*string, bool)`

GetAccessqueryOk returns a tuple with the Accessquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessquery

`func (o *CreateEndpointRequest) SetAccessquery(v string)`

SetAccessquery sets Accessquery field to given value.

### HasAccessquery

`func (o *CreateEndpointRequest) HasAccessquery() bool`

HasAccessquery returns a boolean if a field has been set.

### GetAccountNameRule

`func (o *CreateEndpointRequest) GetAccountNameRule() string`

GetAccountNameRule returns the AccountNameRule field if non-nil, zero value otherwise.

### GetAccountNameRuleOk

`func (o *CreateEndpointRequest) GetAccountNameRuleOk() (*string, bool)`

GetAccountNameRuleOk returns a tuple with the AccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameRule

`func (o *CreateEndpointRequest) SetAccountNameRule(v string)`

SetAccountNameRule sets AccountNameRule field to given value.

### HasAccountNameRule

`func (o *CreateEndpointRequest) HasAccountNameRule() bool`

HasAccountNameRule returns a boolean if a field has been set.

### GetBlockInflightRequest

`func (o *CreateEndpointRequest) GetBlockInflightRequest() string`

GetBlockInflightRequest returns the BlockInflightRequest field if non-nil, zero value otherwise.

### GetBlockInflightRequestOk

`func (o *CreateEndpointRequest) GetBlockInflightRequestOk() (*string, bool)`

GetBlockInflightRequestOk returns a tuple with the BlockInflightRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInflightRequest

`func (o *CreateEndpointRequest) SetBlockInflightRequest(v string)`

SetBlockInflightRequest sets BlockInflightRequest field to given value.

### HasBlockInflightRequest

`func (o *CreateEndpointRequest) HasBlockInflightRequest() bool`

HasBlockInflightRequest returns a boolean if a field has been set.

### GetConnectionconfig

`func (o *CreateEndpointRequest) GetConnectionconfig() string`

GetConnectionconfig returns the Connectionconfig field if non-nil, zero value otherwise.

### GetConnectionconfigOk

`func (o *CreateEndpointRequest) GetConnectionconfigOk() (*string, bool)`

GetConnectionconfigOk returns a tuple with the Connectionconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionconfig

`func (o *CreateEndpointRequest) SetConnectionconfig(v string)`

SetConnectionconfig sets Connectionconfig field to given value.

### HasConnectionconfig

`func (o *CreateEndpointRequest) HasConnectionconfig() bool`

HasConnectionconfig returns a boolean if a field has been set.

### GetDescription

`func (o *CreateEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableNewAccountRequestIfAccountExists

`func (o *CreateEndpointRequest) GetDisableNewAccountRequestIfAccountExists() string`

GetDisableNewAccountRequestIfAccountExists returns the DisableNewAccountRequestIfAccountExists field if non-nil, zero value otherwise.

### GetDisableNewAccountRequestIfAccountExistsOk

`func (o *CreateEndpointRequest) GetDisableNewAccountRequestIfAccountExistsOk() (*string, bool)`

GetDisableNewAccountRequestIfAccountExistsOk returns a tuple with the DisableNewAccountRequestIfAccountExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewAccountRequestIfAccountExists

`func (o *CreateEndpointRequest) SetDisableNewAccountRequestIfAccountExists(v string)`

SetDisableNewAccountRequestIfAccountExists sets DisableNewAccountRequestIfAccountExists field to given value.

### HasDisableNewAccountRequestIfAccountExists

`func (o *CreateEndpointRequest) HasDisableNewAccountRequestIfAccountExists() bool`

HasDisableNewAccountRequestIfAccountExists returns a boolean if a field has been set.

### GetDisableModifyAccount

`func (o *CreateEndpointRequest) GetDisableModifyAccount() string`

GetDisableModifyAccount returns the DisableModifyAccount field if non-nil, zero value otherwise.

### GetDisableModifyAccountOk

`func (o *CreateEndpointRequest) GetDisableModifyAccountOk() (*string, bool)`

GetDisableModifyAccountOk returns a tuple with the DisableModifyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModifyAccount

`func (o *CreateEndpointRequest) SetDisableModifyAccount(v string)`

SetDisableModifyAccount sets DisableModifyAccount field to given value.

### HasDisableModifyAccount

`func (o *CreateEndpointRequest) HasDisableModifyAccount() bool`

HasDisableModifyAccount returns a boolean if a field has been set.

### GetDisableRemoveAccount

`func (o *CreateEndpointRequest) GetDisableRemoveAccount() string`

GetDisableRemoveAccount returns the DisableRemoveAccount field if non-nil, zero value otherwise.

### GetDisableRemoveAccountOk

`func (o *CreateEndpointRequest) GetDisableRemoveAccountOk() (*string, bool)`

GetDisableRemoveAccountOk returns a tuple with the DisableRemoveAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRemoveAccount

`func (o *CreateEndpointRequest) SetDisableRemoveAccount(v string)`

SetDisableRemoveAccount sets DisableRemoveAccount field to given value.

### HasDisableRemoveAccount

`func (o *CreateEndpointRequest) HasDisableRemoveAccount() bool`

HasDisableRemoveAccount returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateEndpointRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateEndpointRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateEndpointRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnableCopyAccess

`func (o *CreateEndpointRequest) GetEnableCopyAccess() string`

GetEnableCopyAccess returns the EnableCopyAccess field if non-nil, zero value otherwise.

### GetEnableCopyAccessOk

`func (o *CreateEndpointRequest) GetEnableCopyAccessOk() (*string, bool)`

GetEnableCopyAccessOk returns a tuple with the EnableCopyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyAccess

`func (o *CreateEndpointRequest) SetEnableCopyAccess(v string)`

SetEnableCopyAccess sets EnableCopyAccess field to given value.

### HasEnableCopyAccess

`func (o *CreateEndpointRequest) HasEnableCopyAccess() bool`

HasEnableCopyAccess returns a boolean if a field has been set.

### GetEndpointname

`func (o *CreateEndpointRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *CreateEndpointRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *CreateEndpointRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.


### GetOwner

`func (o *CreateEndpointRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateEndpointRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateEndpointRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateEndpointRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerType

`func (o *CreateEndpointRequest) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *CreateEndpointRequest) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *CreateEndpointRequest) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *CreateEndpointRequest) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetResourceOwnerType

`func (o *CreateEndpointRequest) GetResourceOwnerType() string`

GetResourceOwnerType returns the ResourceOwnerType field if non-nil, zero value otherwise.

### GetResourceOwnerTypeOk

`func (o *CreateEndpointRequest) GetResourceOwnerTypeOk() (*string, bool)`

GetResourceOwnerTypeOk returns a tuple with the ResourceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerType

`func (o *CreateEndpointRequest) SetResourceOwnerType(v string)`

SetResourceOwnerType sets ResourceOwnerType field to given value.

### HasResourceOwnerType

`func (o *CreateEndpointRequest) HasResourceOwnerType() bool`

HasResourceOwnerType returns a boolean if a field has been set.

### GetResourceOwner

`func (o *CreateEndpointRequest) GetResourceOwner() string`

GetResourceOwner returns the ResourceOwner field if non-nil, zero value otherwise.

### GetResourceOwnerOk

`func (o *CreateEndpointRequest) GetResourceOwnerOk() (*string, bool)`

GetResourceOwnerOk returns a tuple with the ResourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwner

`func (o *CreateEndpointRequest) SetResourceOwner(v string)`

SetResourceOwner sets ResourceOwner field to given value.

### HasResourceOwner

`func (o *CreateEndpointRequest) HasResourceOwner() bool`

HasResourceOwner returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *CreateEndpointRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *CreateEndpointRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *CreateEndpointRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetUserAccountCorrelationRule

`func (o *CreateEndpointRequest) GetUserAccountCorrelationRule() string`

GetUserAccountCorrelationRule returns the UserAccountCorrelationRule field if non-nil, zero value otherwise.

### GetUserAccountCorrelationRuleOk

`func (o *CreateEndpointRequest) GetUserAccountCorrelationRuleOk() (*string, bool)`

GetUserAccountCorrelationRuleOk returns a tuple with the UserAccountCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountCorrelationRule

`func (o *CreateEndpointRequest) SetUserAccountCorrelationRule(v string)`

SetUserAccountCorrelationRule sets UserAccountCorrelationRule field to given value.

### HasUserAccountCorrelationRule

`func (o *CreateEndpointRequest) HasUserAccountCorrelationRule() bool`

HasUserAccountCorrelationRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


