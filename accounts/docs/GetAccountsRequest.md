# GetAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Filter accounts by username. | [optional] 
**Endpoint** | Pointer to **string** | Filter accounts by endpoint. | [optional] 
**Max** | Pointer to **string** | Maximum number of records to return. | [optional] 
**Offset** | Pointer to **string** | Pagination offset. | [optional] 
**AccountQuery** | Pointer to **string** | A query string to filter accounts. | [optional] 
**Advsearchcriteria** | Pointer to [**GetAccountsRequestAdvsearchcriteria**](GetAccountsRequestAdvsearchcriteria.md) |  | [optional] 

## Methods

### NewGetAccountsRequest

`func NewGetAccountsRequest() *GetAccountsRequest`

NewGetAccountsRequest instantiates a new GetAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsRequestWithDefaults

`func NewGetAccountsRequestWithDefaults() *GetAccountsRequest`

NewGetAccountsRequestWithDefaults instantiates a new GetAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GetAccountsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAccountsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAccountsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAccountsRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpoint

`func (o *GetAccountsRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetAccountsRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetAccountsRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GetAccountsRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetMax

`func (o *GetAccountsRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetAccountsRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetAccountsRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetAccountsRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetAccountsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetAccountsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetAccountsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetAccountsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetAccountQuery

`func (o *GetAccountsRequest) GetAccountQuery() string`

GetAccountQuery returns the AccountQuery field if non-nil, zero value otherwise.

### GetAccountQueryOk

`func (o *GetAccountsRequest) GetAccountQueryOk() (*string, bool)`

GetAccountQueryOk returns a tuple with the AccountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQuery

`func (o *GetAccountsRequest) SetAccountQuery(v string)`

SetAccountQuery sets AccountQuery field to given value.

### HasAccountQuery

`func (o *GetAccountsRequest) HasAccountQuery() bool`

HasAccountQuery returns a boolean if a field has been set.

### GetAdvsearchcriteria

`func (o *GetAccountsRequest) GetAdvsearchcriteria() GetAccountsRequestAdvsearchcriteria`

GetAdvsearchcriteria returns the Advsearchcriteria field if non-nil, zero value otherwise.

### GetAdvsearchcriteriaOk

`func (o *GetAccountsRequest) GetAdvsearchcriteriaOk() (*GetAccountsRequestAdvsearchcriteria, bool)`

GetAdvsearchcriteriaOk returns a tuple with the Advsearchcriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvsearchcriteria

`func (o *GetAccountsRequest) SetAdvsearchcriteria(v GetAccountsRequestAdvsearchcriteria)`

SetAdvsearchcriteria sets Advsearchcriteria field to given value.

### HasAdvsearchcriteria

`func (o *GetAccountsRequest) HasAdvsearchcriteria() bool`

HasAdvsearchcriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


