# CreateOrUpdateConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**ErrorCode** | **string** |  | 
**ConnectionKey** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateOrUpdateConnectionResponse

`func NewCreateOrUpdateConnectionResponse(msg string, errorCode string, ) *CreateOrUpdateConnectionResponse`

NewCreateOrUpdateConnectionResponse instantiates a new CreateOrUpdateConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConnectionResponseWithDefaults

`func NewCreateOrUpdateConnectionResponseWithDefaults() *CreateOrUpdateConnectionResponse`

NewCreateOrUpdateConnectionResponseWithDefaults instantiates a new CreateOrUpdateConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateOrUpdateConnectionResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateOrUpdateConnectionResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateOrUpdateConnectionResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *CreateOrUpdateConnectionResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateOrUpdateConnectionResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateOrUpdateConnectionResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetConnectionKey

`func (o *CreateOrUpdateConnectionResponse) GetConnectionKey() int32`

GetConnectionKey returns the ConnectionKey field if non-nil, zero value otherwise.

### GetConnectionKeyOk

`func (o *CreateOrUpdateConnectionResponse) GetConnectionKeyOk() (*int32, bool)`

GetConnectionKeyOk returns a tuple with the ConnectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionKey

`func (o *CreateOrUpdateConnectionResponse) SetConnectionKey(v int32)`

SetConnectionKey sets ConnectionKey field to given value.

### HasConnectionKey

`func (o *CreateOrUpdateConnectionResponse) HasConnectionKey() bool`

HasConnectionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


