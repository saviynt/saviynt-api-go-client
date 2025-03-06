# GetEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) |  | [optional] 
**DisplayCount** | Pointer to **int32** | Number of endpoints returned | [optional] 
**ErrorCode** | **string** | Error code (0 indicates success) | 
**TotalCount** | Pointer to **int32** | Total number of endpoints matching the criteria | [optional] 
**Message** | **string** | Response message | 

## Methods

### NewGetEndpointsResponse

`func NewGetEndpointsResponse(errorCode string, message string, ) *GetEndpointsResponse`

NewGetEndpointsResponse instantiates a new GetEndpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsResponseWithDefaults

`func NewGetEndpointsResponseWithDefaults() *GetEndpointsResponse`

NewGetEndpointsResponseWithDefaults instantiates a new GetEndpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *GetEndpointsResponse) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *GetEndpointsResponse) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *GetEndpointsResponse) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *GetEndpointsResponse) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetEndpointsResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetEndpointsResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetEndpointsResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetEndpointsResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEndpointsResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEndpointsResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEndpointsResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetTotalCount

`func (o *GetEndpointsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetEndpointsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetEndpointsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetEndpointsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetMessage

`func (o *GetEndpointsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEndpointsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEndpointsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


