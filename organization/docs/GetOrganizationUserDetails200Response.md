# GetOrganizationUserDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Displaycount** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Totalcount** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Organizations** | Pointer to [**[]GetOrganizationUserDetails200ResponseOrganizationsInner**](GetOrganizationUserDetails200ResponseOrganizationsInner.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOrganizationUserDetails200Response

`func NewGetOrganizationUserDetails200Response() *GetOrganizationUserDetails200Response`

NewGetOrganizationUserDetails200Response instantiates a new GetOrganizationUserDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUserDetails200ResponseWithDefaults

`func NewGetOrganizationUserDetails200ResponseWithDefaults() *GetOrganizationUserDetails200Response`

NewGetOrganizationUserDetails200ResponseWithDefaults instantiates a new GetOrganizationUserDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplaycount

`func (o *GetOrganizationUserDetails200Response) GetDisplaycount() int32`

GetDisplaycount returns the Displaycount field if non-nil, zero value otherwise.

### GetDisplaycountOk

`func (o *GetOrganizationUserDetails200Response) GetDisplaycountOk() (*int32, bool)`

GetDisplaycountOk returns a tuple with the Displaycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaycount

`func (o *GetOrganizationUserDetails200Response) SetDisplaycount(v int32)`

SetDisplaycount sets Displaycount field to given value.

### HasDisplaycount

`func (o *GetOrganizationUserDetails200Response) HasDisplaycount() bool`

HasDisplaycount returns a boolean if a field has been set.

### GetMsg

`func (o *GetOrganizationUserDetails200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetOrganizationUserDetails200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetOrganizationUserDetails200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetOrganizationUserDetails200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTotalcount

`func (o *GetOrganizationUserDetails200Response) GetTotalcount() int32`

GetTotalcount returns the Totalcount field if non-nil, zero value otherwise.

### GetTotalcountOk

`func (o *GetOrganizationUserDetails200Response) GetTotalcountOk() (*int32, bool)`

GetTotalcountOk returns a tuple with the Totalcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalcount

`func (o *GetOrganizationUserDetails200Response) SetTotalcount(v int32)`

SetTotalcount sets Totalcount field to given value.

### HasTotalcount

`func (o *GetOrganizationUserDetails200Response) HasTotalcount() bool`

HasTotalcount returns a boolean if a field has been set.

### GetOffset

`func (o *GetOrganizationUserDetails200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetOrganizationUserDetails200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetOrganizationUserDetails200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetOrganizationUserDetails200Response) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMax

`func (o *GetOrganizationUserDetails200Response) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetOrganizationUserDetails200Response) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetOrganizationUserDetails200Response) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetOrganizationUserDetails200Response) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetOrganizationUserDetails200Response) GetOrganizations() []GetOrganizationUserDetails200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetOrganizationUserDetails200Response) GetOrganizationsOk() (*[]GetOrganizationUserDetails200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetOrganizationUserDetails200Response) SetOrganizations(v []GetOrganizationUserDetails200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetOrganizationUserDetails200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetOrganizationUserDetails200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetOrganizationUserDetails200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetOrganizationUserDetails200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetOrganizationUserDetails200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationUserDetails200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationUserDetails200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationUserDetails200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationUserDetails200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


