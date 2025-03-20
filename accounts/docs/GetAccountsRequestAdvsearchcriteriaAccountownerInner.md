# GetAccountsRequestAdvsearchcriteriaAccountownerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of owner, e.g., user or usergroup | 
**Value** | **string** | Username or usergroup name | 
**Rank** | Pointer to **string** | Optional rank (e.g., &#39;1&#39; to &#39;5&#39;) | [optional] 

## Methods

### NewGetAccountsRequestAdvsearchcriteriaAccountownerInner

`func NewGetAccountsRequestAdvsearchcriteriaAccountownerInner(type_ string, value string, ) *GetAccountsRequestAdvsearchcriteriaAccountownerInner`

NewGetAccountsRequestAdvsearchcriteriaAccountownerInner instantiates a new GetAccountsRequestAdvsearchcriteriaAccountownerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsRequestAdvsearchcriteriaAccountownerInnerWithDefaults

`func NewGetAccountsRequestAdvsearchcriteriaAccountownerInnerWithDefaults() *GetAccountsRequestAdvsearchcriteriaAccountownerInner`

NewGetAccountsRequestAdvsearchcriteriaAccountownerInnerWithDefaults instantiates a new GetAccountsRequestAdvsearchcriteriaAccountownerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetRank

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *GetAccountsRequestAdvsearchcriteriaAccountownerInner) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


