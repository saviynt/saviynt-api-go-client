# CreateUpdateEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** |  | 
**EntitlementValue** | **string** |  | 
**Entitlementtype** | **string** |  | 
**Entitlementattributes** | Pointer to **map[string]interface{}** | Additional attributes for the entitlement. | [optional] 
**Entitlementowner1** | Pointer to **string** |  | [optional] 
**Entitlementowner3** | Pointer to **string** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**NewEntitlementValue** | Pointer to **string** |  | [optional] 
**Entitlementcasecheck** | Pointer to **bool** |  | [optional] [default to false]
**EntitlementValuekey** | Pointer to **string** |  | [optional] 
**UpdatedentitlementValue** | Pointer to **string** |  | [optional] 
**Entitlementmap** | Pointer to [**CreateUpdateEntitlementRequestEntitlementmap**](CreateUpdateEntitlementRequestEntitlementmap.md) |  | [optional] 

## Methods

### NewCreateUpdateEntitlementRequest

`func NewCreateUpdateEntitlementRequest(endpoint string, entitlementValue string, entitlementtype string, ) *CreateUpdateEntitlementRequest`

NewCreateUpdateEntitlementRequest instantiates a new CreateUpdateEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateEntitlementRequestWithDefaults

`func NewCreateUpdateEntitlementRequestWithDefaults() *CreateUpdateEntitlementRequest`

NewCreateUpdateEntitlementRequestWithDefaults instantiates a new CreateUpdateEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CreateUpdateEntitlementRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateUpdateEntitlementRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateUpdateEntitlementRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementValue

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValue() string`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValueOk() (*string, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *CreateUpdateEntitlementRequest) SetEntitlementValue(v string)`

SetEntitlementValue sets EntitlementValue field to given value.


### GetEntitlementtype

`func (o *CreateUpdateEntitlementRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *CreateUpdateEntitlementRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.


### GetEntitlementattributes

`func (o *CreateUpdateEntitlementRequest) GetEntitlementattributes() map[string]interface{}`

GetEntitlementattributes returns the Entitlementattributes field if non-nil, zero value otherwise.

### GetEntitlementattributesOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementattributesOk() (*map[string]interface{}, bool)`

GetEntitlementattributesOk returns a tuple with the Entitlementattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementattributes

`func (o *CreateUpdateEntitlementRequest) SetEntitlementattributes(v map[string]interface{})`

SetEntitlementattributes sets Entitlementattributes field to given value.

### HasEntitlementattributes

`func (o *CreateUpdateEntitlementRequest) HasEntitlementattributes() bool`

HasEntitlementattributes returns a boolean if a field has been set.

### GetEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner1() string`

GetEntitlementowner1 returns the Entitlementowner1 field if non-nil, zero value otherwise.

### GetEntitlementowner1Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner1Ok() (*string, bool)`

GetEntitlementowner1Ok returns a tuple with the Entitlementowner1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner1(v string)`

SetEntitlementowner1 sets Entitlementowner1 field to given value.

### HasEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner1() bool`

HasEntitlementowner1 returns a boolean if a field has been set.

### GetEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner3() string`

GetEntitlementowner3 returns the Entitlementowner3 field if non-nil, zero value otherwise.

### GetEntitlementowner3Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner3Ok() (*string, bool)`

GetEntitlementowner3Ok returns a tuple with the Entitlementowner3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner3(v string)`

SetEntitlementowner3 sets Entitlementowner3 field to given value.

### HasEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner3() bool`

HasEntitlementowner3 returns a boolean if a field has been set.

### GetEntitlementID

`func (o *CreateUpdateEntitlementRequest) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *CreateUpdateEntitlementRequest) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *CreateUpdateEntitlementRequest) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetNewEntitlementValue

`func (o *CreateUpdateEntitlementRequest) GetNewEntitlementValue() string`

GetNewEntitlementValue returns the NewEntitlementValue field if non-nil, zero value otherwise.

### GetNewEntitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetNewEntitlementValueOk() (*string, bool)`

GetNewEntitlementValueOk returns a tuple with the NewEntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEntitlementValue

`func (o *CreateUpdateEntitlementRequest) SetNewEntitlementValue(v string)`

SetNewEntitlementValue sets NewEntitlementValue field to given value.

### HasNewEntitlementValue

`func (o *CreateUpdateEntitlementRequest) HasNewEntitlementValue() bool`

HasNewEntitlementValue returns a boolean if a field has been set.

### GetEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) GetEntitlementcasecheck() bool`

GetEntitlementcasecheck returns the Entitlementcasecheck field if non-nil, zero value otherwise.

### GetEntitlementcasecheckOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementcasecheckOk() (*bool, bool)`

GetEntitlementcasecheckOk returns a tuple with the Entitlementcasecheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) SetEntitlementcasecheck(v bool)`

SetEntitlementcasecheck sets Entitlementcasecheck field to given value.

### HasEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) HasEntitlementcasecheck() bool`

HasEntitlementcasecheck returns a boolean if a field has been set.

### GetEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValuekey() string`

GetEntitlementValuekey returns the EntitlementValuekey field if non-nil, zero value otherwise.

### GetEntitlementValuekeyOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValuekeyOk() (*string, bool)`

GetEntitlementValuekeyOk returns a tuple with the EntitlementValuekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) SetEntitlementValuekey(v string)`

SetEntitlementValuekey sets EntitlementValuekey field to given value.

### HasEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) HasEntitlementValuekey() bool`

HasEntitlementValuekey returns a boolean if a field has been set.

### GetUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) GetUpdatedentitlementValue() string`

GetUpdatedentitlementValue returns the UpdatedentitlementValue field if non-nil, zero value otherwise.

### GetUpdatedentitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetUpdatedentitlementValueOk() (*string, bool)`

GetUpdatedentitlementValueOk returns a tuple with the UpdatedentitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) SetUpdatedentitlementValue(v string)`

SetUpdatedentitlementValue sets UpdatedentitlementValue field to given value.

### HasUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) HasUpdatedentitlementValue() bool`

HasUpdatedentitlementValue returns a boolean if a field has been set.

### GetEntitlementmap

`func (o *CreateUpdateEntitlementRequest) GetEntitlementmap() CreateUpdateEntitlementRequestEntitlementmap`

GetEntitlementmap returns the Entitlementmap field if non-nil, zero value otherwise.

### GetEntitlementmapOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementmapOk() (*CreateUpdateEntitlementRequestEntitlementmap, bool)`

GetEntitlementmapOk returns a tuple with the Entitlementmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementmap

`func (o *CreateUpdateEntitlementRequest) SetEntitlementmap(v CreateUpdateEntitlementRequestEntitlementmap)`

SetEntitlementmap sets Entitlementmap field to given value.

### HasEntitlementmap

`func (o *CreateUpdateEntitlementRequest) HasEntitlementmap() bool`

HasEntitlementmap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


