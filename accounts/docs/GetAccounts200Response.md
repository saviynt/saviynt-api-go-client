# GetAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Displaycount** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Accountdetails** | Pointer to [**[]GetAccounts200ResponseAccountdetailsInner**](GetAccounts200ResponseAccountdetailsInner.md) |  | [optional] 

## Methods

### NewGetAccounts200Response

`func NewGetAccounts200Response() *GetAccounts200Response`

NewGetAccounts200Response instantiates a new GetAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccounts200ResponseWithDefaults

`func NewGetAccounts200ResponseWithDefaults() *GetAccounts200Response`

NewGetAccounts200ResponseWithDefaults instantiates a new GetAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetAccounts200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAccounts200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAccounts200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetAccounts200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetAccounts200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetAccounts200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetAccounts200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetAccounts200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDisplaycount

`func (o *GetAccounts200Response) GetDisplaycount() int32`

GetDisplaycount returns the Displaycount field if non-nil, zero value otherwise.

### GetDisplaycountOk

`func (o *GetAccounts200Response) GetDisplaycountOk() (*int32, bool)`

GetDisplaycountOk returns a tuple with the Displaycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaycount

`func (o *GetAccounts200Response) SetDisplaycount(v int32)`

SetDisplaycount sets Displaycount field to given value.

### HasDisplaycount

`func (o *GetAccounts200Response) HasDisplaycount() bool`

HasDisplaycount returns a boolean if a field has been set.

### GetTotal

`func (o *GetAccounts200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAccounts200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAccounts200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAccounts200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetAccountdetails

`func (o *GetAccounts200Response) GetAccountdetails() []GetAccounts200ResponseAccountdetailsInner`

GetAccountdetails returns the Accountdetails field if non-nil, zero value otherwise.

### GetAccountdetailsOk

`func (o *GetAccounts200Response) GetAccountdetailsOk() (*[]GetAccounts200ResponseAccountdetailsInner, bool)`

GetAccountdetailsOk returns a tuple with the Accountdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountdetails

`func (o *GetAccounts200Response) SetAccountdetails(v []GetAccounts200ResponseAccountdetailsInner)`

SetAccountdetails sets Accountdetails field to given value.

### HasAccountdetails

`func (o *GetAccounts200Response) HasAccountdetails() bool`

HasAccountdetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


