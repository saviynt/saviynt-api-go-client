# CreateEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Response message | 
**ErrorCode** | **string** | Error code (0 indicates success) | 

## Methods

### NewCreateEndpointResponse

`func NewCreateEndpointResponse(msg string, errorCode string, ) *CreateEndpointResponse`

NewCreateEndpointResponse instantiates a new CreateEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointResponseWithDefaults

`func NewCreateEndpointResponseWithDefaults() *CreateEndpointResponse`

NewCreateEndpointResponseWithDefaults instantiates a new CreateEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateEndpointResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateEndpointResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateEndpointResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *CreateEndpointResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateEndpointResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateEndpointResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


