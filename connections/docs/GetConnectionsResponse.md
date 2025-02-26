# GetConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionList** | [**[]Connection**](Connection.md) |  | 
**ErrorCode** | **string** |  | 
**Msg** | **string** |  | 
**DisplayCount** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetConnectionsResponse

`func NewGetConnectionsResponse(connectionList []Connection, errorCode string, msg string, ) *GetConnectionsResponse`

NewGetConnectionsResponse instantiates a new GetConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionsResponseWithDefaults

`func NewGetConnectionsResponseWithDefaults() *GetConnectionsResponse`

NewGetConnectionsResponseWithDefaults instantiates a new GetConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionList

`func (o *GetConnectionsResponse) GetConnectionList() []Connection`

GetConnectionList returns the ConnectionList field if non-nil, zero value otherwise.

### GetConnectionListOk

`func (o *GetConnectionsResponse) GetConnectionListOk() (*[]Connection, bool)`

GetConnectionListOk returns a tuple with the ConnectionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionList

`func (o *GetConnectionsResponse) SetConnectionList(v []Connection)`

SetConnectionList sets ConnectionList field to given value.


### GetErrorCode

`func (o *GetConnectionsResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetConnectionsResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetConnectionsResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMsg

`func (o *GetConnectionsResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetConnectionsResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetConnectionsResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetDisplayCount

`func (o *GetConnectionsResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetConnectionsResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetConnectionsResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetConnectionsResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetConnectionsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetConnectionsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetConnectionsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetConnectionsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


