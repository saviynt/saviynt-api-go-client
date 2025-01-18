# FetchDelegatesListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | this is the parentusername | 
**Max** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | values can be ACTIVE / INACTIVE | [optional] 

## Methods

### NewFetchDelegatesListRequest

`func NewFetchDelegatesListRequest(userName string, ) *FetchDelegatesListRequest`

NewFetchDelegatesListRequest instantiates a new FetchDelegatesListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchDelegatesListRequestWithDefaults

`func NewFetchDelegatesListRequestWithDefaults() *FetchDelegatesListRequest`

NewFetchDelegatesListRequestWithDefaults instantiates a new FetchDelegatesListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *FetchDelegatesListRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FetchDelegatesListRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FetchDelegatesListRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetMax

`func (o *FetchDelegatesListRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FetchDelegatesListRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FetchDelegatesListRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *FetchDelegatesListRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *FetchDelegatesListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FetchDelegatesListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FetchDelegatesListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *FetchDelegatesListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *FetchDelegatesListRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchDelegatesListRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchDelegatesListRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FetchDelegatesListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


