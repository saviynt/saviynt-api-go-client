# SuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message | [optional] 
**ErrorCode** | Pointer to **string** | Error code (0 indicates success) | [optional] 

## Methods

### NewSuccessResponse

`func NewSuccessResponse() *SuccessResponse`

NewSuccessResponse instantiates a new SuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessResponseWithDefaults

`func NewSuccessResponseWithDefaults() *SuccessResponse`

NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *SuccessResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SuccessResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SuccessResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SuccessResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *SuccessResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SuccessResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SuccessResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SuccessResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


