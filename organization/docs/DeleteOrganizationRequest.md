# DeleteOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizationname** | Pointer to **string** | The name of the organization to delete. | [optional] 
**Username** | Pointer to **string** | The username of the user deleting the organization. | [optional] 

## Methods

### NewDeleteOrganizationRequest

`func NewDeleteOrganizationRequest() *DeleteOrganizationRequest`

NewDeleteOrganizationRequest instantiates a new DeleteOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteOrganizationRequestWithDefaults

`func NewDeleteOrganizationRequestWithDefaults() *DeleteOrganizationRequest`

NewDeleteOrganizationRequestWithDefaults instantiates a new DeleteOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationname

`func (o *DeleteOrganizationRequest) GetOrganizationname() string`

GetOrganizationname returns the Organizationname field if non-nil, zero value otherwise.

### GetOrganizationnameOk

`func (o *DeleteOrganizationRequest) GetOrganizationnameOk() (*string, bool)`

GetOrganizationnameOk returns a tuple with the Organizationname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationname

`func (o *DeleteOrganizationRequest) SetOrganizationname(v string)`

SetOrganizationname sets Organizationname field to given value.

### HasOrganizationname

`func (o *DeleteOrganizationRequest) HasOrganizationname() bool`

HasOrganizationname returns a boolean if a field has been set.

### GetUsername

`func (o *DeleteOrganizationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeleteOrganizationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeleteOrganizationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeleteOrganizationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


