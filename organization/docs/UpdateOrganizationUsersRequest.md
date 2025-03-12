# UpdateOrganizationUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizationname** | Pointer to **string** | The secondary organization name. | [optional] 
**Users** | Pointer to [**[]UpdateOrganizationUsersRequestUsersInner**](UpdateOrganizationUsersRequestUsersInner.md) | A list of users with their update actions. | [optional] 

## Methods

### NewUpdateOrganizationUsersRequest

`func NewUpdateOrganizationUsersRequest() *UpdateOrganizationUsersRequest`

NewUpdateOrganizationUsersRequest instantiates a new UpdateOrganizationUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationUsersRequestWithDefaults

`func NewUpdateOrganizationUsersRequestWithDefaults() *UpdateOrganizationUsersRequest`

NewUpdateOrganizationUsersRequestWithDefaults instantiates a new UpdateOrganizationUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationname

`func (o *UpdateOrganizationUsersRequest) GetOrganizationname() string`

GetOrganizationname returns the Organizationname field if non-nil, zero value otherwise.

### GetOrganizationnameOk

`func (o *UpdateOrganizationUsersRequest) GetOrganizationnameOk() (*string, bool)`

GetOrganizationnameOk returns a tuple with the Organizationname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationname

`func (o *UpdateOrganizationUsersRequest) SetOrganizationname(v string)`

SetOrganizationname sets Organizationname field to given value.

### HasOrganizationname

`func (o *UpdateOrganizationUsersRequest) HasOrganizationname() bool`

HasOrganizationname returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateOrganizationUsersRequest) GetUsers() []UpdateOrganizationUsersRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateOrganizationUsersRequest) GetUsersOk() (*[]UpdateOrganizationUsersRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateOrganizationUsersRequest) SetUsers(v []UpdateOrganizationUsersRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateOrganizationUsersRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


