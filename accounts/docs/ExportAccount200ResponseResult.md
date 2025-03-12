# ExportAccount200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]ExportAccount200ResponseResultAccountsInner**](ExportAccount200ResponseResultAccountsInner.md) |  | [optional] 

## Methods

### NewExportAccount200ResponseResult

`func NewExportAccount200ResponseResult() *ExportAccount200ResponseResult`

NewExportAccount200ResponseResult instantiates a new ExportAccount200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAccount200ResponseResultWithDefaults

`func NewExportAccount200ResponseResultWithDefaults() *ExportAccount200ResponseResult`

NewExportAccount200ResponseResultWithDefaults instantiates a new ExportAccount200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ExportAccount200ResponseResult) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExportAccount200ResponseResult) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExportAccount200ResponseResult) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExportAccount200ResponseResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetAccounts

`func (o *ExportAccount200ResponseResult) GetAccounts() []ExportAccount200ResponseResultAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ExportAccount200ResponseResult) GetAccountsOk() (*[]ExportAccount200ResponseResultAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ExportAccount200ResponseResult) SetAccounts(v []ExportAccount200ResponseResultAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ExportAccount200ResponseResult) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


