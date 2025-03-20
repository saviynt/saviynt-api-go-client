# GetListOfPrivilegesForEntitlementType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivilegeDetails** | Pointer to [**[]GetListOfPrivilegesForEntitlementType200ResponsePrivilegeDetailsInner**](GetListOfPrivilegesForEntitlementType200ResponsePrivilegeDetailsInner.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewGetListOfPrivilegesForEntitlementType200Response

`func NewGetListOfPrivilegesForEntitlementType200Response() *GetListOfPrivilegesForEntitlementType200Response`

NewGetListOfPrivilegesForEntitlementType200Response instantiates a new GetListOfPrivilegesForEntitlementType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListOfPrivilegesForEntitlementType200ResponseWithDefaults

`func NewGetListOfPrivilegesForEntitlementType200ResponseWithDefaults() *GetListOfPrivilegesForEntitlementType200Response`

NewGetListOfPrivilegesForEntitlementType200ResponseWithDefaults instantiates a new GetListOfPrivilegesForEntitlementType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilegeDetails

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetPrivilegeDetails() []GetListOfPrivilegesForEntitlementType200ResponsePrivilegeDetailsInner`

GetPrivilegeDetails returns the PrivilegeDetails field if non-nil, zero value otherwise.

### GetPrivilegeDetailsOk

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetPrivilegeDetailsOk() (*[]GetListOfPrivilegesForEntitlementType200ResponsePrivilegeDetailsInner, bool)`

GetPrivilegeDetailsOk returns a tuple with the PrivilegeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeDetails

`func (o *GetListOfPrivilegesForEntitlementType200Response) SetPrivilegeDetails(v []GetListOfPrivilegesForEntitlementType200ResponsePrivilegeDetailsInner)`

SetPrivilegeDetails sets PrivilegeDetails field to given value.

### HasPrivilegeDetails

`func (o *GetListOfPrivilegesForEntitlementType200Response) HasPrivilegeDetails() bool`

HasPrivilegeDetails returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetListOfPrivilegesForEntitlementType200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetListOfPrivilegesForEntitlementType200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetListOfPrivilegesForEntitlementType200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetListOfPrivilegesForEntitlementType200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetListOfPrivilegesForEntitlementType200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


