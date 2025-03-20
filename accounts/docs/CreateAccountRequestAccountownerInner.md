# CreateAccountRequestAccountownerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of owner, e.g., \&quot;user\&quot; or \&quot;usergroup\&quot;. | 
**Value** | **string** | Username or usergroup name. | 
**Rank** | Pointer to **string** | Optional rank from \&quot;1\&quot; to \&quot;5\&quot;. Default is \&quot;1\&quot;. | [optional] 

## Methods

### NewCreateAccountRequestAccountownerInner

`func NewCreateAccountRequestAccountownerInner(type_ string, value string, ) *CreateAccountRequestAccountownerInner`

NewCreateAccountRequestAccountownerInner instantiates a new CreateAccountRequestAccountownerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountRequestAccountownerInnerWithDefaults

`func NewCreateAccountRequestAccountownerInnerWithDefaults() *CreateAccountRequestAccountownerInner`

NewCreateAccountRequestAccountownerInnerWithDefaults instantiates a new CreateAccountRequestAccountownerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateAccountRequestAccountownerInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountRequestAccountownerInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountRequestAccountownerInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CreateAccountRequestAccountownerInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateAccountRequestAccountownerInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateAccountRequestAccountownerInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetRank

`func (o *CreateAccountRequestAccountownerInner) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *CreateAccountRequestAccountownerInner) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *CreateAccountRequestAccountownerInner) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *CreateAccountRequestAccountownerInner) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


