# DeleteConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectionkey** | Pointer to **[]string** |  | [optional] 
**Updateuser** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteConnectionRequest

`func NewDeleteConnectionRequest() *DeleteConnectionRequest`

NewDeleteConnectionRequest instantiates a new DeleteConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteConnectionRequestWithDefaults

`func NewDeleteConnectionRequestWithDefaults() *DeleteConnectionRequest`

NewDeleteConnectionRequestWithDefaults instantiates a new DeleteConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionkey

`func (o *DeleteConnectionRequest) GetConnectionkey() []string`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *DeleteConnectionRequest) GetConnectionkeyOk() (*[]string, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *DeleteConnectionRequest) SetConnectionkey(v []string)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *DeleteConnectionRequest) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetUpdateuser

`func (o *DeleteConnectionRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *DeleteConnectionRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *DeleteConnectionRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *DeleteConnectionRequest) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


