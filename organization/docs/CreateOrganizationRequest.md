# CreateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizationname** | **string** | Unique name of the organization. | 
**Username** | **string** | User creating the organization. | 
**Primarycontact** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Customproperty1** | Pointer to **string** |  | [optional] 
**Customproperty5** | Pointer to **string** |  | [optional] 
**Customproperty10** | Pointer to **string** |  | [optional] 
**Organizationtype** | Pointer to **string** |  | [optional] 
**Parentorganization** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to [**[]CreateOrganizationRequestEndpointsInner**](CreateOrganizationRequestEndpointsInner.md) |  | [optional] 
**Roles** | Pointer to [**[]CreateOrganizationRequestRolesInner**](CreateOrganizationRequestRolesInner.md) |  | [optional] 
**Rules** | Pointer to [**[]CreateOrganizationRequestRulesInner**](CreateOrganizationRequestRulesInner.md) |  | [optional] 
**Entitlements** | Pointer to [**[]CreateOrganizationRequestEntitlementsInner**](CreateOrganizationRequestEntitlementsInner.md) |  | [optional] 
**Users** | Pointer to [**[]CreateOrganizationRequestUsersInner**](CreateOrganizationRequestUsersInner.md) |  | [optional] 
**Owners** | Pointer to [**[]CreateOrganizationRequestOwnersInner**](CreateOrganizationRequestOwnersInner.md) |  | [optional] 
**Attributes** | Pointer to [**[]CreateOrganizationRequestAttributesInner**](CreateOrganizationRequestAttributesInner.md) |  | [optional] 

## Methods

### NewCreateOrganizationRequest

`func NewCreateOrganizationRequest(organizationname string, username string, ) *CreateOrganizationRequest`

NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRequestWithDefaults

`func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest`

NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationname

`func (o *CreateOrganizationRequest) GetOrganizationname() string`

GetOrganizationname returns the Organizationname field if non-nil, zero value otherwise.

### GetOrganizationnameOk

`func (o *CreateOrganizationRequest) GetOrganizationnameOk() (*string, bool)`

GetOrganizationnameOk returns a tuple with the Organizationname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationname

`func (o *CreateOrganizationRequest) SetOrganizationname(v string)`

SetOrganizationname sets Organizationname field to given value.


### GetUsername

`func (o *CreateOrganizationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateOrganizationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateOrganizationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPrimarycontact

`func (o *CreateOrganizationRequest) GetPrimarycontact() string`

GetPrimarycontact returns the Primarycontact field if non-nil, zero value otherwise.

### GetPrimarycontactOk

`func (o *CreateOrganizationRequest) GetPrimarycontactOk() (*string, bool)`

GetPrimarycontactOk returns a tuple with the Primarycontact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarycontact

`func (o *CreateOrganizationRequest) SetPrimarycontact(v string)`

SetPrimarycontact sets Primarycontact field to given value.

### HasPrimarycontact

`func (o *CreateOrganizationRequest) HasPrimarycontact() bool`

HasPrimarycontact returns a boolean if a field has been set.

### GetComments

`func (o *CreateOrganizationRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CreateOrganizationRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CreateOrganizationRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CreateOrganizationRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOrganizationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrganizationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrganizationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateOrganizationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomproperty1

`func (o *CreateOrganizationRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *CreateOrganizationRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *CreateOrganizationRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *CreateOrganizationRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *CreateOrganizationRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *CreateOrganizationRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *CreateOrganizationRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *CreateOrganizationRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *CreateOrganizationRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *CreateOrganizationRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *CreateOrganizationRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *CreateOrganizationRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetOrganizationtype

`func (o *CreateOrganizationRequest) GetOrganizationtype() string`

GetOrganizationtype returns the Organizationtype field if non-nil, zero value otherwise.

### GetOrganizationtypeOk

`func (o *CreateOrganizationRequest) GetOrganizationtypeOk() (*string, bool)`

GetOrganizationtypeOk returns a tuple with the Organizationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationtype

`func (o *CreateOrganizationRequest) SetOrganizationtype(v string)`

SetOrganizationtype sets Organizationtype field to given value.

### HasOrganizationtype

`func (o *CreateOrganizationRequest) HasOrganizationtype() bool`

HasOrganizationtype returns a boolean if a field has been set.

### GetParentorganization

`func (o *CreateOrganizationRequest) GetParentorganization() string`

GetParentorganization returns the Parentorganization field if non-nil, zero value otherwise.

### GetParentorganizationOk

`func (o *CreateOrganizationRequest) GetParentorganizationOk() (*string, bool)`

GetParentorganizationOk returns a tuple with the Parentorganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentorganization

`func (o *CreateOrganizationRequest) SetParentorganization(v string)`

SetParentorganization sets Parentorganization field to given value.

### HasParentorganization

`func (o *CreateOrganizationRequest) HasParentorganization() bool`

HasParentorganization returns a boolean if a field has been set.

### GetLocation

`func (o *CreateOrganizationRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateOrganizationRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateOrganizationRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateOrganizationRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEndpoints

`func (o *CreateOrganizationRequest) GetEndpoints() []CreateOrganizationRequestEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateOrganizationRequest) GetEndpointsOk() (*[]CreateOrganizationRequestEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateOrganizationRequest) SetEndpoints(v []CreateOrganizationRequestEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateOrganizationRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetRoles

`func (o *CreateOrganizationRequest) GetRoles() []CreateOrganizationRequestRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateOrganizationRequest) GetRolesOk() (*[]CreateOrganizationRequestRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateOrganizationRequest) SetRoles(v []CreateOrganizationRequestRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateOrganizationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRules

`func (o *CreateOrganizationRequest) GetRules() []CreateOrganizationRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateOrganizationRequest) GetRulesOk() (*[]CreateOrganizationRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateOrganizationRequest) SetRules(v []CreateOrganizationRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateOrganizationRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetEntitlements

`func (o *CreateOrganizationRequest) GetEntitlements() []CreateOrganizationRequestEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CreateOrganizationRequest) GetEntitlementsOk() (*[]CreateOrganizationRequestEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CreateOrganizationRequest) SetEntitlements(v []CreateOrganizationRequestEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CreateOrganizationRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetUsers

`func (o *CreateOrganizationRequest) GetUsers() []CreateOrganizationRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateOrganizationRequest) GetUsersOk() (*[]CreateOrganizationRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateOrganizationRequest) SetUsers(v []CreateOrganizationRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateOrganizationRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetOwners

`func (o *CreateOrganizationRequest) GetOwners() []CreateOrganizationRequestOwnersInner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *CreateOrganizationRequest) GetOwnersOk() (*[]CreateOrganizationRequestOwnersInner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *CreateOrganizationRequest) SetOwners(v []CreateOrganizationRequestOwnersInner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *CreateOrganizationRequest) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetAttributes

`func (o *CreateOrganizationRequest) GetAttributes() []CreateOrganizationRequestAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateOrganizationRequest) GetAttributesOk() (*[]CreateOrganizationRequestAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateOrganizationRequest) SetAttributes(v []CreateOrganizationRequestAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateOrganizationRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


