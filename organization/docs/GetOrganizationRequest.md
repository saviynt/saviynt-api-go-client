# GetOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizationname** | Pointer to **string** | Filter organizations by name. | [optional] 
**Max** | Pointer to **string** | Maximum number of records to return. | [optional] 
**Offset** | Pointer to **string** | Offset for pagination. | [optional] 
**Filtercriteria** | Pointer to **string** | A comma‐separated list of criteria to filter organizations. | [optional] 
**OrgQuery** | Pointer to **string** | A SQL-like query (e.g., c.organizationname like &#39;%TestOrg5%&#39;). | [optional] 

## Methods

### NewGetOrganizationRequest

`func NewGetOrganizationRequest() *GetOrganizationRequest`

NewGetOrganizationRequest instantiates a new GetOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationRequestWithDefaults

`func NewGetOrganizationRequestWithDefaults() *GetOrganizationRequest`

NewGetOrganizationRequestWithDefaults instantiates a new GetOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationname

`func (o *GetOrganizationRequest) GetOrganizationname() string`

GetOrganizationname returns the Organizationname field if non-nil, zero value otherwise.

### GetOrganizationnameOk

`func (o *GetOrganizationRequest) GetOrganizationnameOk() (*string, bool)`

GetOrganizationnameOk returns a tuple with the Organizationname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationname

`func (o *GetOrganizationRequest) SetOrganizationname(v string)`

SetOrganizationname sets Organizationname field to given value.

### HasOrganizationname

`func (o *GetOrganizationRequest) HasOrganizationname() bool`

HasOrganizationname returns a boolean if a field has been set.

### GetMax

`func (o *GetOrganizationRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetOrganizationRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetOrganizationRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetOrganizationRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetOrganizationRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetOrganizationRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetOrganizationRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetOrganizationRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFiltercriteria

`func (o *GetOrganizationRequest) GetFiltercriteria() string`

GetFiltercriteria returns the Filtercriteria field if non-nil, zero value otherwise.

### GetFiltercriteriaOk

`func (o *GetOrganizationRequest) GetFiltercriteriaOk() (*string, bool)`

GetFiltercriteriaOk returns a tuple with the Filtercriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltercriteria

`func (o *GetOrganizationRequest) SetFiltercriteria(v string)`

SetFiltercriteria sets Filtercriteria field to given value.

### HasFiltercriteria

`func (o *GetOrganizationRequest) HasFiltercriteria() bool`

HasFiltercriteria returns a boolean if a field has been set.

### GetOrgQuery

`func (o *GetOrganizationRequest) GetOrgQuery() string`

GetOrgQuery returns the OrgQuery field if non-nil, zero value otherwise.

### GetOrgQueryOk

`func (o *GetOrganizationRequest) GetOrgQueryOk() (*string, bool)`

GetOrgQueryOk returns a tuple with the OrgQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgQuery

`func (o *GetOrganizationRequest) SetOrgQuery(v string)`

SetOrgQuery sets OrgQuery field to given value.

### HasOrgQuery

`func (o *GetOrganizationRequest) HasOrgQuery() bool`

HasOrgQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


