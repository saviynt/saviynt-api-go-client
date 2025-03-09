# ExportAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ExportAccount200ResponseResult**](ExportAccount200ResponseResult.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewExportAccount200Response

`func NewExportAccount200Response() *ExportAccount200Response`

NewExportAccount200Response instantiates a new ExportAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAccount200ResponseWithDefaults

`func NewExportAccount200ResponseWithDefaults() *ExportAccount200Response`

NewExportAccount200ResponseWithDefaults instantiates a new ExportAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ExportAccount200Response) GetResult() ExportAccount200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExportAccount200Response) GetResultOk() (*ExportAccount200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExportAccount200Response) SetResult(v ExportAccount200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ExportAccount200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetErrorCode

`func (o *ExportAccount200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ExportAccount200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ExportAccount200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ExportAccount200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *ExportAccount200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExportAccount200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExportAccount200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ExportAccount200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


