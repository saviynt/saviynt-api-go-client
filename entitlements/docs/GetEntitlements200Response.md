# GetEntitlements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlementdetails** | Pointer to [**[]GetEntitlements200ResponseEntitlementdetailsInner**](GetEntitlements200ResponseEntitlementdetailsInner.md) |  | [optional] 
**EntitlementsCount** | Pointer to **float32** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**TotalEntitlementCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetEntitlements200Response

`func NewGetEntitlements200Response() *GetEntitlements200Response`

NewGetEntitlements200Response instantiates a new GetEntitlements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlements200ResponseWithDefaults

`func NewGetEntitlements200ResponseWithDefaults() *GetEntitlements200Response`

NewGetEntitlements200ResponseWithDefaults instantiates a new GetEntitlements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementdetails

`func (o *GetEntitlements200Response) GetEntitlementdetails() []GetEntitlements200ResponseEntitlementdetailsInner`

GetEntitlementdetails returns the Entitlementdetails field if non-nil, zero value otherwise.

### GetEntitlementdetailsOk

`func (o *GetEntitlements200Response) GetEntitlementdetailsOk() (*[]GetEntitlements200ResponseEntitlementdetailsInner, bool)`

GetEntitlementdetailsOk returns a tuple with the Entitlementdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementdetails

`func (o *GetEntitlements200Response) SetEntitlementdetails(v []GetEntitlements200ResponseEntitlementdetailsInner)`

SetEntitlementdetails sets Entitlementdetails field to given value.

### HasEntitlementdetails

`func (o *GetEntitlements200Response) HasEntitlementdetails() bool`

HasEntitlementdetails returns a boolean if a field has been set.

### GetEntitlementsCount

`func (o *GetEntitlements200Response) GetEntitlementsCount() float32`

GetEntitlementsCount returns the EntitlementsCount field if non-nil, zero value otherwise.

### GetEntitlementsCountOk

`func (o *GetEntitlements200Response) GetEntitlementsCountOk() (*float32, bool)`

GetEntitlementsCountOk returns a tuple with the EntitlementsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsCount

`func (o *GetEntitlements200Response) SetEntitlementsCount(v float32)`

SetEntitlementsCount sets EntitlementsCount field to given value.

### HasEntitlementsCount

`func (o *GetEntitlements200Response) HasEntitlementsCount() bool`

HasEntitlementsCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEntitlements200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEntitlements200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEntitlements200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetEntitlements200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *GetEntitlements200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetEntitlements200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetEntitlements200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetEntitlements200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTotalEntitlementCount

`func (o *GetEntitlements200Response) GetTotalEntitlementCount() float32`

GetTotalEntitlementCount returns the TotalEntitlementCount field if non-nil, zero value otherwise.

### GetTotalEntitlementCountOk

`func (o *GetEntitlements200Response) GetTotalEntitlementCountOk() (*float32, bool)`

GetTotalEntitlementCountOk returns a tuple with the TotalEntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntitlementCount

`func (o *GetEntitlements200Response) SetTotalEntitlementCount(v float32)`

SetTotalEntitlementCount sets TotalEntitlementCount field to given value.

### HasTotalEntitlementCount

`func (o *GetEntitlements200Response) HasTotalEntitlementCount() bool`

HasTotalEntitlementCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


