# UpdateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizationname** | **string** |  | 
**Username** | **string** |  | 
**Updatedorgname** | Pointer to **string** |  | [optional] 
**Primarycontact** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Customproperty5** | Pointer to **string** |  | [optional] 
**Customproperty10** | Pointer to **string** |  | [optional] 
**Organizationtype** | Pointer to **string** |  | [optional] 
**Parentorganization** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to [**[]UpdateOrganizationRequestEndpointsInner**](UpdateOrganizationRequestEndpointsInner.md) |  | [optional] 
**Roles** | Pointer to [**[]UpdateOrganizationRequestRolesInner**](UpdateOrganizationRequestRolesInner.md) |  | [optional] 
**Rules** | Pointer to [**[]UpdateOrganizationRequestRulesInner**](UpdateOrganizationRequestRulesInner.md) |  | [optional] 
**Entitlements** | Pointer to [**[]UpdateOrganizationRequestEntitlementsInner**](UpdateOrganizationRequestEntitlementsInner.md) |  | [optional] 
**Users** | Pointer to [**[]UpdateOrganizationRequestUsersInner**](UpdateOrganizationRequestUsersInner.md) |  | [optional] 
**Owners** | Pointer to [**[]UpdateOrganizationRequestOwnersInner**](UpdateOrganizationRequestOwnersInner.md) |  | [optional] 
**Attributes** | Pointer to [**[]UpdateOrganizationRequestAttributesInner**](UpdateOrganizationRequestAttributesInner.md) |  | [optional] 

## Methods

### NewUpdateOrganizationRequest

`func NewUpdateOrganizationRequest(organizationname string, username string, ) *UpdateOrganizationRequest`

NewUpdateOrganizationRequest instantiates a new UpdateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationRequestWithDefaults

`func NewUpdateOrganizationRequestWithDefaults() *UpdateOrganizationRequest`

NewUpdateOrganizationRequestWithDefaults instantiates a new UpdateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationname

`func (o *UpdateOrganizationRequest) GetOrganizationname() string`

GetOrganizationname returns the Organizationname field if non-nil, zero value otherwise.

### GetOrganizationnameOk

`func (o *UpdateOrganizationRequest) GetOrganizationnameOk() (*string, bool)`

GetOrganizationnameOk returns a tuple with the Organizationname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationname

`func (o *UpdateOrganizationRequest) SetOrganizationname(v string)`

SetOrganizationname sets Organizationname field to given value.


### GetUsername

`func (o *UpdateOrganizationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateOrganizationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateOrganizationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUpdatedorgname

`func (o *UpdateOrganizationRequest) GetUpdatedorgname() string`

GetUpdatedorgname returns the Updatedorgname field if non-nil, zero value otherwise.

### GetUpdatedorgnameOk

`func (o *UpdateOrganizationRequest) GetUpdatedorgnameOk() (*string, bool)`

GetUpdatedorgnameOk returns a tuple with the Updatedorgname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedorgname

`func (o *UpdateOrganizationRequest) SetUpdatedorgname(v string)`

SetUpdatedorgname sets Updatedorgname field to given value.

### HasUpdatedorgname

`func (o *UpdateOrganizationRequest) HasUpdatedorgname() bool`

HasUpdatedorgname returns a boolean if a field has been set.

### GetPrimarycontact

`func (o *UpdateOrganizationRequest) GetPrimarycontact() string`

GetPrimarycontact returns the Primarycontact field if non-nil, zero value otherwise.

### GetPrimarycontactOk

`func (o *UpdateOrganizationRequest) GetPrimarycontactOk() (*string, bool)`

GetPrimarycontactOk returns a tuple with the Primarycontact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarycontact

`func (o *UpdateOrganizationRequest) SetPrimarycontact(v string)`

SetPrimarycontact sets Primarycontact field to given value.

### HasPrimarycontact

`func (o *UpdateOrganizationRequest) HasPrimarycontact() bool`

HasPrimarycontact returns a boolean if a field has been set.

### GetComments

`func (o *UpdateOrganizationRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateOrganizationRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateOrganizationRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateOrganizationRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateOrganizationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateOrganizationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateOrganizationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateOrganizationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *UpdateOrganizationRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *UpdateOrganizationRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *UpdateOrganizationRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *UpdateOrganizationRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *UpdateOrganizationRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *UpdateOrganizationRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *UpdateOrganizationRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *UpdateOrganizationRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetOrganizationtype

`func (o *UpdateOrganizationRequest) GetOrganizationtype() string`

GetOrganizationtype returns the Organizationtype field if non-nil, zero value otherwise.

### GetOrganizationtypeOk

`func (o *UpdateOrganizationRequest) GetOrganizationtypeOk() (*string, bool)`

GetOrganizationtypeOk returns a tuple with the Organizationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationtype

`func (o *UpdateOrganizationRequest) SetOrganizationtype(v string)`

SetOrganizationtype sets Organizationtype field to given value.

### HasOrganizationtype

`func (o *UpdateOrganizationRequest) HasOrganizationtype() bool`

HasOrganizationtype returns a boolean if a field has been set.

### GetParentorganization

`func (o *UpdateOrganizationRequest) GetParentorganization() string`

GetParentorganization returns the Parentorganization field if non-nil, zero value otherwise.

### GetParentorganizationOk

`func (o *UpdateOrganizationRequest) GetParentorganizationOk() (*string, bool)`

GetParentorganizationOk returns a tuple with the Parentorganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentorganization

`func (o *UpdateOrganizationRequest) SetParentorganization(v string)`

SetParentorganization sets Parentorganization field to given value.

### HasParentorganization

`func (o *UpdateOrganizationRequest) HasParentorganization() bool`

HasParentorganization returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateOrganizationRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateOrganizationRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateOrganizationRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateOrganizationRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEndpoints

`func (o *UpdateOrganizationRequest) GetEndpoints() []UpdateOrganizationRequestEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *UpdateOrganizationRequest) GetEndpointsOk() (*[]UpdateOrganizationRequestEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *UpdateOrganizationRequest) SetEndpoints(v []UpdateOrganizationRequestEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *UpdateOrganizationRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateOrganizationRequest) GetRoles() []UpdateOrganizationRequestRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateOrganizationRequest) GetRolesOk() (*[]UpdateOrganizationRequestRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateOrganizationRequest) SetRoles(v []UpdateOrganizationRequestRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateOrganizationRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRules

`func (o *UpdateOrganizationRequest) GetRules() []UpdateOrganizationRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateOrganizationRequest) GetRulesOk() (*[]UpdateOrganizationRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateOrganizationRequest) SetRules(v []UpdateOrganizationRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateOrganizationRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetEntitlements

`func (o *UpdateOrganizationRequest) GetEntitlements() []UpdateOrganizationRequestEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *UpdateOrganizationRequest) GetEntitlementsOk() (*[]UpdateOrganizationRequestEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *UpdateOrganizationRequest) SetEntitlements(v []UpdateOrganizationRequestEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *UpdateOrganizationRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateOrganizationRequest) GetUsers() []UpdateOrganizationRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateOrganizationRequest) GetUsersOk() (*[]UpdateOrganizationRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateOrganizationRequest) SetUsers(v []UpdateOrganizationRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateOrganizationRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetOwners

`func (o *UpdateOrganizationRequest) GetOwners() []UpdateOrganizationRequestOwnersInner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *UpdateOrganizationRequest) GetOwnersOk() (*[]UpdateOrganizationRequestOwnersInner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *UpdateOrganizationRequest) SetOwners(v []UpdateOrganizationRequestOwnersInner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *UpdateOrganizationRequest) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateOrganizationRequest) GetAttributes() []UpdateOrganizationRequestAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateOrganizationRequest) GetAttributesOk() (*[]UpdateOrganizationRequestAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateOrganizationRequest) SetAttributes(v []UpdateOrganizationRequestAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateOrganizationRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


