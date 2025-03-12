# GetOrganizationUserDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The name of the user to fetch organization details for. | 
**Max** | Pointer to **string** | The maximum number of records to return. | [optional] 
**Offset** | Pointer to **string** | The starting offset of records. | [optional] 
**OrgQuery** | Pointer to **string** | A query to filter organizations (e.g. c.organizationname like &#39;%TestOrg5%&#39;). | [optional] 
**Verbose** | Pointer to **bool** | Set to true to get additional user details. | [optional] 
**Fetchonlycount** | Pointer to **bool** | Set to true to fetch only the count. | [optional] 
**Orgmemebership** | Pointer to **string** | Specify if you want only details for a particular organization type (e.g. primary or secondary). | [optional] 

## Methods

### NewGetOrganizationUserDetailsRequest

`func NewGetOrganizationUserDetailsRequest(username string, ) *GetOrganizationUserDetailsRequest`

NewGetOrganizationUserDetailsRequest instantiates a new GetOrganizationUserDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUserDetailsRequestWithDefaults

`func NewGetOrganizationUserDetailsRequestWithDefaults() *GetOrganizationUserDetailsRequest`

NewGetOrganizationUserDetailsRequestWithDefaults instantiates a new GetOrganizationUserDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GetOrganizationUserDetailsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationUserDetailsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationUserDetailsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetMax

`func (o *GetOrganizationUserDetailsRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetOrganizationUserDetailsRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetOrganizationUserDetailsRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetOrganizationUserDetailsRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetOrganizationUserDetailsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetOrganizationUserDetailsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetOrganizationUserDetailsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetOrganizationUserDetailsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrgQuery

`func (o *GetOrganizationUserDetailsRequest) GetOrgQuery() string`

GetOrgQuery returns the OrgQuery field if non-nil, zero value otherwise.

### GetOrgQueryOk

`func (o *GetOrganizationUserDetailsRequest) GetOrgQueryOk() (*string, bool)`

GetOrgQueryOk returns a tuple with the OrgQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgQuery

`func (o *GetOrganizationUserDetailsRequest) SetOrgQuery(v string)`

SetOrgQuery sets OrgQuery field to given value.

### HasOrgQuery

`func (o *GetOrganizationUserDetailsRequest) HasOrgQuery() bool`

HasOrgQuery returns a boolean if a field has been set.

### GetVerbose

`func (o *GetOrganizationUserDetailsRequest) GetVerbose() bool`

GetVerbose returns the Verbose field if non-nil, zero value otherwise.

### GetVerboseOk

`func (o *GetOrganizationUserDetailsRequest) GetVerboseOk() (*bool, bool)`

GetVerboseOk returns a tuple with the Verbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbose

`func (o *GetOrganizationUserDetailsRequest) SetVerbose(v bool)`

SetVerbose sets Verbose field to given value.

### HasVerbose

`func (o *GetOrganizationUserDetailsRequest) HasVerbose() bool`

HasVerbose returns a boolean if a field has been set.

### GetFetchonlycount

`func (o *GetOrganizationUserDetailsRequest) GetFetchonlycount() bool`

GetFetchonlycount returns the Fetchonlycount field if non-nil, zero value otherwise.

### GetFetchonlycountOk

`func (o *GetOrganizationUserDetailsRequest) GetFetchonlycountOk() (*bool, bool)`

GetFetchonlycountOk returns a tuple with the Fetchonlycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchonlycount

`func (o *GetOrganizationUserDetailsRequest) SetFetchonlycount(v bool)`

SetFetchonlycount sets Fetchonlycount field to given value.

### HasFetchonlycount

`func (o *GetOrganizationUserDetailsRequest) HasFetchonlycount() bool`

HasFetchonlycount returns a boolean if a field has been set.

### GetOrgmemebership

`func (o *GetOrganizationUserDetailsRequest) GetOrgmemebership() string`

GetOrgmemebership returns the Orgmemebership field if non-nil, zero value otherwise.

### GetOrgmemebershipOk

`func (o *GetOrganizationUserDetailsRequest) GetOrgmemebershipOk() (*string, bool)`

GetOrgmemebershipOk returns a tuple with the Orgmemebership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgmemebership

`func (o *GetOrganizationUserDetailsRequest) SetOrgmemebership(v string)`

SetOrgmemebership sets Orgmemebership field to given value.

### HasOrgmemebership

`func (o *GetOrganizationUserDetailsRequest) HasOrgmemebership() bool`

HasOrgmemebership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


