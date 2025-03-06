# RemoveEntitlementFromRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requestor** | Pointer to **string** | The person requesting the removal (optional). | [optional] 
**Rolename** | Pointer to **string** | The name of the role. If both rolename and rolekey are provided, rolekey takes precedence. | [optional] 
**Rolekey** | Pointer to **string** | The key of the role. Takes precedence over rolename if provided. | [optional] 
**Comments** | Pointer to **string** | Optional comments about the removal. | [optional] 
**Entitlements** | Pointer to [**[]RemoveEntitlementFromRoleRequestEntitlementsInner**](RemoveEntitlementFromRoleRequestEntitlementsInner.md) |  | [optional] 

## Methods

### NewRemoveEntitlementFromRoleRequest

`func NewRemoveEntitlementFromRoleRequest() *RemoveEntitlementFromRoleRequest`

NewRemoveEntitlementFromRoleRequest instantiates a new RemoveEntitlementFromRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveEntitlementFromRoleRequestWithDefaults

`func NewRemoveEntitlementFromRoleRequestWithDefaults() *RemoveEntitlementFromRoleRequest`

NewRemoveEntitlementFromRoleRequestWithDefaults instantiates a new RemoveEntitlementFromRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestor

`func (o *RemoveEntitlementFromRoleRequest) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *RemoveEntitlementFromRoleRequest) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *RemoveEntitlementFromRoleRequest) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *RemoveEntitlementFromRoleRequest) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetRolename

`func (o *RemoveEntitlementFromRoleRequest) GetRolename() string`

GetRolename returns the Rolename field if non-nil, zero value otherwise.

### GetRolenameOk

`func (o *RemoveEntitlementFromRoleRequest) GetRolenameOk() (*string, bool)`

GetRolenameOk returns a tuple with the Rolename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolename

`func (o *RemoveEntitlementFromRoleRequest) SetRolename(v string)`

SetRolename sets Rolename field to given value.

### HasRolename

`func (o *RemoveEntitlementFromRoleRequest) HasRolename() bool`

HasRolename returns a boolean if a field has been set.

### GetRolekey

`func (o *RemoveEntitlementFromRoleRequest) GetRolekey() string`

GetRolekey returns the Rolekey field if non-nil, zero value otherwise.

### GetRolekeyOk

`func (o *RemoveEntitlementFromRoleRequest) GetRolekeyOk() (*string, bool)`

GetRolekeyOk returns a tuple with the Rolekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolekey

`func (o *RemoveEntitlementFromRoleRequest) SetRolekey(v string)`

SetRolekey sets Rolekey field to given value.

### HasRolekey

`func (o *RemoveEntitlementFromRoleRequest) HasRolekey() bool`

HasRolekey returns a boolean if a field has been set.

### GetComments

`func (o *RemoveEntitlementFromRoleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RemoveEntitlementFromRoleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RemoveEntitlementFromRoleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RemoveEntitlementFromRoleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetEntitlements

`func (o *RemoveEntitlementFromRoleRequest) GetEntitlements() []RemoveEntitlementFromRoleRequestEntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *RemoveEntitlementFromRoleRequest) GetEntitlementsOk() (*[]RemoveEntitlementFromRoleRequestEntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *RemoveEntitlementFromRoleRequest) SetEntitlements(v []RemoveEntitlementFromRoleRequestEntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *RemoveEntitlementFromRoleRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


