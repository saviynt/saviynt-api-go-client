# CreateUpdateUsergroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]CreateUpdateUsergroupRequestEntitlementsInner**](CreateUpdateUsergroupRequestEntitlementsInner.md) |  | [optional] 
**Owners** | Pointer to [**[]CreateUpdateUsergroupRequestOwnersInner**](CreateUpdateUsergroupRequestOwnersInner.md) |  | [optional] 
**UserGroupdescription** | Pointer to **string** |  | [optional] 
**Usergroup** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]CreateUpdateUsergroupRequestUsersInner**](CreateUpdateUsergroupRequestUsersInner.md) |  | [optional] 

## Methods

### NewCreateUpdateUsergroupRequest

`func NewCreateUpdateUsergroupRequest() *CreateUpdateUsergroupRequest`

NewCreateUpdateUsergroupRequest instantiates a new CreateUpdateUsergroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateUsergroupRequestWithDefaults

`func NewCreateUpdateUsergroupRequestWithDefaults() *CreateUpdateUsergroupRequest`

NewCreateUpdateUsergroupRequestWithDefaults instantiates a new CreateUpdateUsergroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *CreateUpdateUsergroupRequest) GetEntitlements() []CreateUpdateUsergroupRequestEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CreateUpdateUsergroupRequest) GetEntitlementsOk() (*[]CreateUpdateUsergroupRequestEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CreateUpdateUsergroupRequest) SetEntitlements(v []CreateUpdateUsergroupRequestEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CreateUpdateUsergroupRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetOwners

`func (o *CreateUpdateUsergroupRequest) GetOwners() []CreateUpdateUsergroupRequestOwnersInner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *CreateUpdateUsergroupRequest) GetOwnersOk() (*[]CreateUpdateUsergroupRequestOwnersInner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *CreateUpdateUsergroupRequest) SetOwners(v []CreateUpdateUsergroupRequestOwnersInner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *CreateUpdateUsergroupRequest) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetUserGroupdescription

`func (o *CreateUpdateUsergroupRequest) GetUserGroupdescription() string`

GetUserGroupdescription returns the UserGroupdescription field if non-nil, zero value otherwise.

### GetUserGroupdescriptionOk

`func (o *CreateUpdateUsergroupRequest) GetUserGroupdescriptionOk() (*string, bool)`

GetUserGroupdescriptionOk returns a tuple with the UserGroupdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupdescription

`func (o *CreateUpdateUsergroupRequest) SetUserGroupdescription(v string)`

SetUserGroupdescription sets UserGroupdescription field to given value.

### HasUserGroupdescription

`func (o *CreateUpdateUsergroupRequest) HasUserGroupdescription() bool`

HasUserGroupdescription returns a boolean if a field has been set.

### GetUsergroup

`func (o *CreateUpdateUsergroupRequest) GetUsergroup() string`

GetUsergroup returns the Usergroup field if non-nil, zero value otherwise.

### GetUsergroupOk

`func (o *CreateUpdateUsergroupRequest) GetUsergroupOk() (*string, bool)`

GetUsergroupOk returns a tuple with the Usergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroup

`func (o *CreateUpdateUsergroupRequest) SetUsergroup(v string)`

SetUsergroup sets Usergroup field to given value.

### HasUsergroup

`func (o *CreateUpdateUsergroupRequest) HasUsergroup() bool`

HasUsergroup returns a boolean if a field has been set.

### GetUsername

`func (o *CreateUpdateUsergroupRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUpdateUsergroupRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUpdateUsergroupRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateUpdateUsergroupRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUsers

`func (o *CreateUpdateUsergroupRequest) GetUsers() []CreateUpdateUsergroupRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateUpdateUsergroupRequest) GetUsersOk() (*[]CreateUpdateUsergroupRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateUpdateUsergroupRequest) SetUsers(v []CreateUpdateUsergroupRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateUpdateUsergroupRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


