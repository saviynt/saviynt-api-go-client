# GetEntitlementValuesForEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** |  | 
**EntitlementType** | Pointer to **string** |  | [optional] 
**Entownerwithrank** | Pointer to **string** |  | [optional] 
**Returnentitlementmap** | Pointer to **string** | If &#39;true&#39;, returns entitlementmap details. | [optional] 
**Max** | Pointer to **string** | Maximum number of records to return. | [optional] 
**Offset** | Pointer to **string** | Pagination offset. | [optional] 

## Methods

### NewGetEntitlementValuesForEndpointRequest

`func NewGetEntitlementValuesForEndpointRequest(endpoint string, ) *GetEntitlementValuesForEndpointRequest`

NewGetEntitlementValuesForEndpointRequest instantiates a new GetEntitlementValuesForEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlementValuesForEndpointRequestWithDefaults

`func NewGetEntitlementValuesForEndpointRequestWithDefaults() *GetEntitlementValuesForEndpointRequest`

NewGetEntitlementValuesForEndpointRequestWithDefaults instantiates a new GetEntitlementValuesForEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *GetEntitlementValuesForEndpointRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetEntitlementValuesForEndpointRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetEntitlementValuesForEndpointRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementType

`func (o *GetEntitlementValuesForEndpointRequest) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *GetEntitlementValuesForEndpointRequest) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *GetEntitlementValuesForEndpointRequest) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.

### HasEntitlementType

`func (o *GetEntitlementValuesForEndpointRequest) HasEntitlementType() bool`

HasEntitlementType returns a boolean if a field has been set.

### GetEntownerwithrank

`func (o *GetEntitlementValuesForEndpointRequest) GetEntownerwithrank() string`

GetEntownerwithrank returns the Entownerwithrank field if non-nil, zero value otherwise.

### GetEntownerwithrankOk

`func (o *GetEntitlementValuesForEndpointRequest) GetEntownerwithrankOk() (*string, bool)`

GetEntownerwithrankOk returns a tuple with the Entownerwithrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntownerwithrank

`func (o *GetEntitlementValuesForEndpointRequest) SetEntownerwithrank(v string)`

SetEntownerwithrank sets Entownerwithrank field to given value.

### HasEntownerwithrank

`func (o *GetEntitlementValuesForEndpointRequest) HasEntownerwithrank() bool`

HasEntownerwithrank returns a boolean if a field has been set.

### GetReturnentitlementmap

`func (o *GetEntitlementValuesForEndpointRequest) GetReturnentitlementmap() string`

GetReturnentitlementmap returns the Returnentitlementmap field if non-nil, zero value otherwise.

### GetReturnentitlementmapOk

`func (o *GetEntitlementValuesForEndpointRequest) GetReturnentitlementmapOk() (*string, bool)`

GetReturnentitlementmapOk returns a tuple with the Returnentitlementmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnentitlementmap

`func (o *GetEntitlementValuesForEndpointRequest) SetReturnentitlementmap(v string)`

SetReturnentitlementmap sets Returnentitlementmap field to given value.

### HasReturnentitlementmap

`func (o *GetEntitlementValuesForEndpointRequest) HasReturnentitlementmap() bool`

HasReturnentitlementmap returns a boolean if a field has been set.

### GetMax

`func (o *GetEntitlementValuesForEndpointRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetEntitlementValuesForEndpointRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetEntitlementValuesForEndpointRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetEntitlementValuesForEndpointRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetEntitlementValuesForEndpointRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetEntitlementValuesForEndpointRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetEntitlementValuesForEndpointRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetEntitlementValuesForEndpointRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


