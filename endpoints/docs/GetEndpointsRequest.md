# GetEndpointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | Filter by connection type | [optional] 
**Endpointkey** | Pointer to **[]int32** | Filter by endpoint key(s) | [optional] 
**Endpointname** | Pointer to **string** | Filter by endpoint name | [optional] 
**Max** | Pointer to **int32** | Maximum number of results to return. Maximium is 500. | [optional] 
**Offset** | Pointer to **int32** | Pagination offset | [optional] 
**FilterCriteria** | Pointer to **map[string]interface{}** | Custom filter criteria | [optional] 

## Methods

### NewGetEndpointsRequest

`func NewGetEndpointsRequest() *GetEndpointsRequest`

NewGetEndpointsRequest instantiates a new GetEndpointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsRequestWithDefaults

`func NewGetEndpointsRequestWithDefaults() *GetEndpointsRequest`

NewGetEndpointsRequestWithDefaults instantiates a new GetEndpointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *GetEndpointsRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetEndpointsRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetEndpointsRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetEndpointsRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndpointkey

`func (o *GetEndpointsRequest) GetEndpointkey() []int32`

GetEndpointkey returns the Endpointkey field if non-nil, zero value otherwise.

### GetEndpointkeyOk

`func (o *GetEndpointsRequest) GetEndpointkeyOk() (*[]int32, bool)`

GetEndpointkeyOk returns a tuple with the Endpointkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointkey

`func (o *GetEndpointsRequest) SetEndpointkey(v []int32)`

SetEndpointkey sets Endpointkey field to given value.

### HasEndpointkey

`func (o *GetEndpointsRequest) HasEndpointkey() bool`

HasEndpointkey returns a boolean if a field has been set.

### GetEndpointname

`func (o *GetEndpointsRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *GetEndpointsRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *GetEndpointsRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.

### HasEndpointname

`func (o *GetEndpointsRequest) HasEndpointname() bool`

HasEndpointname returns a boolean if a field has been set.

### GetMax

`func (o *GetEndpointsRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetEndpointsRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetEndpointsRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetEndpointsRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetEndpointsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetEndpointsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetEndpointsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetEndpointsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *GetEndpointsRequest) GetFilterCriteria() map[string]interface{}`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *GetEndpointsRequest) GetFilterCriteriaOk() (*map[string]interface{}, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *GetEndpointsRequest) SetFilterCriteria(v map[string]interface{})`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *GetEndpointsRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


