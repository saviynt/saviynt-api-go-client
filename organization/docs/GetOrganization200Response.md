# GetOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Displaycount** | Pointer to **float32** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**[]GetOrganization200ResponseOrganizationsInner**](GetOrganization200ResponseOrganizationsInner.md) |  | [optional] 

## Methods

### NewGetOrganization200Response

`func NewGetOrganization200Response() *GetOrganization200Response`

NewGetOrganization200Response instantiates a new GetOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganization200ResponseWithDefaults

`func NewGetOrganization200ResponseWithDefaults() *GetOrganization200Response`

NewGetOrganization200ResponseWithDefaults instantiates a new GetOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplaycount

`func (o *GetOrganization200Response) GetDisplaycount() float32`

GetDisplaycount returns the Displaycount field if non-nil, zero value otherwise.

### GetDisplaycountOk

`func (o *GetOrganization200Response) GetDisplaycountOk() (*float32, bool)`

GetDisplaycountOk returns a tuple with the Displaycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaycount

`func (o *GetOrganization200Response) SetDisplaycount(v float32)`

SetDisplaycount sets Displaycount field to given value.

### HasDisplaycount

`func (o *GetOrganization200Response) HasDisplaycount() bool`

HasDisplaycount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetOrganization200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetOrganization200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetOrganization200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetOrganization200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *GetOrganization200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetOrganization200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetOrganization200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetOrganization200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetOrganization200Response) GetOrganizations() []GetOrganization200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetOrganization200Response) GetOrganizationsOk() (*[]GetOrganization200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetOrganization200Response) SetOrganizations(v []GetOrganization200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetOrganization200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


