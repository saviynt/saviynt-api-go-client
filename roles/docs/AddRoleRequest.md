# AddRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rolename** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewAddRoleRequest

`func NewAddRoleRequest() *AddRoleRequest`

NewAddRoleRequest instantiates a new AddRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoleRequestWithDefaults

`func NewAddRoleRequestWithDefaults() *AddRoleRequest`

NewAddRoleRequestWithDefaults instantiates a new AddRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolename

`func (o *AddRoleRequest) GetRolename() string`

GetRolename returns the Rolename field if non-nil, zero value otherwise.

### GetRolenameOk

`func (o *AddRoleRequest) GetRolenameOk() (*string, bool)`

GetRolenameOk returns a tuple with the Rolename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolename

`func (o *AddRoleRequest) SetRolename(v string)`

SetRolename sets Rolename field to given value.

### HasRolename

`func (o *AddRoleRequest) HasRolename() bool`

HasRolename returns a boolean if a field has been set.

### GetUsername

`func (o *AddRoleRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddRoleRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddRoleRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddRoleRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


