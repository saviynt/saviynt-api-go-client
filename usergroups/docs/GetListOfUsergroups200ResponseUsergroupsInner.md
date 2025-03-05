# GetListOfUsergroups200ResponseUsergroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]GetListOfUsergroups200ResponseUsergroupsInnerEntitlementsInner**](GetListOfUsergroups200ResponseUsergroupsInnerEntitlementsInner.md) |  | [optional] 
**Groupid** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to **[]interface{}** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**UserGroupdescription** | Pointer to **string** |  | [optional] 
**UserGroupname** | Pointer to **string** |  | [optional] 
**Usergroupkey** | Pointer to **float32** |  | [optional] 
**Users** | Pointer to [**[]GetListOfUsergroups200ResponseUsergroupsInnerUsersInner**](GetListOfUsergroups200ResponseUsergroupsInnerUsersInner.md) |  | [optional] 

## Methods

### NewGetListOfUsergroups200ResponseUsergroupsInner

`func NewGetListOfUsergroups200ResponseUsergroupsInner() *GetListOfUsergroups200ResponseUsergroupsInner`

NewGetListOfUsergroups200ResponseUsergroupsInner instantiates a new GetListOfUsergroups200ResponseUsergroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListOfUsergroups200ResponseUsergroupsInnerWithDefaults

`func NewGetListOfUsergroups200ResponseUsergroupsInnerWithDefaults() *GetListOfUsergroups200ResponseUsergroupsInner`

NewGetListOfUsergroups200ResponseUsergroupsInnerWithDefaults instantiates a new GetListOfUsergroups200ResponseUsergroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetEntitlements() []GetListOfUsergroups200ResponseUsergroupsInnerEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetEntitlementsOk() (*[]GetListOfUsergroups200ResponseUsergroupsInnerEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetEntitlements(v []GetListOfUsergroups200ResponseUsergroupsInnerEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetGroupid

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetGroupid() string`

GetGroupid returns the Groupid field if non-nil, zero value otherwise.

### GetGroupidOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetGroupidOk() (*string, bool)`

GetGroupidOk returns a tuple with the Groupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupid

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetGroupid(v string)`

SetGroupid sets Groupid field to given value.

### HasGroupid

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasGroupid() bool`

HasGroupid returns a boolean if a field has been set.

### GetOwners

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetOwners() []interface{}`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetOwnersOk() (*[]interface{}, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetOwners(v []interface{})`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetRisk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetUserGroupdescription

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUserGroupdescription() string`

GetUserGroupdescription returns the UserGroupdescription field if non-nil, zero value otherwise.

### GetUserGroupdescriptionOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUserGroupdescriptionOk() (*string, bool)`

GetUserGroupdescriptionOk returns a tuple with the UserGroupdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupdescription

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetUserGroupdescription(v string)`

SetUserGroupdescription sets UserGroupdescription field to given value.

### HasUserGroupdescription

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasUserGroupdescription() bool`

HasUserGroupdescription returns a boolean if a field has been set.

### GetUserGroupname

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUserGroupname() string`

GetUserGroupname returns the UserGroupname field if non-nil, zero value otherwise.

### GetUserGroupnameOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUserGroupnameOk() (*string, bool)`

GetUserGroupnameOk returns a tuple with the UserGroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupname

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetUserGroupname(v string)`

SetUserGroupname sets UserGroupname field to given value.

### HasUserGroupname

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasUserGroupname() bool`

HasUserGroupname returns a boolean if a field has been set.

### GetUsergroupkey

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUsergroupkey() float32`

GetUsergroupkey returns the Usergroupkey field if non-nil, zero value otherwise.

### GetUsergroupkeyOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUsergroupkeyOk() (*float32, bool)`

GetUsergroupkeyOk returns a tuple with the Usergroupkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroupkey

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetUsergroupkey(v float32)`

SetUsergroupkey sets Usergroupkey field to given value.

### HasUsergroupkey

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasUsergroupkey() bool`

HasUsergroupkey returns a boolean if a field has been set.

### GetUsers

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUsers() []GetListOfUsergroups200ResponseUsergroupsInnerUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) GetUsersOk() (*[]GetListOfUsergroups200ResponseUsergroupsInnerUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) SetUsers(v []GetListOfUsergroups200ResponseUsergroupsInnerUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetListOfUsergroups200ResponseUsergroupsInner) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


