# UpdateAccountRequestAccountownerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Owner type (user or usergroup). | 
**Value** | **string** | Username or usergroup name. | 
**Rank** | Pointer to **string** | Optional rank. Allowed values: &#39;1&#39;-&#39;5&#39;, &#39;26&#39; (for primary certifier), &#39;27&#39; (for secondary certifier). | [optional] 
**Action** | **string** | Action to perform, either &#39;add&#39; or &#39;remove&#39;. | 

## Methods

### NewUpdateAccountRequestAccountownerInner

`func NewUpdateAccountRequestAccountownerInner(type_ string, value string, action string, ) *UpdateAccountRequestAccountownerInner`

NewUpdateAccountRequestAccountownerInner instantiates a new UpdateAccountRequestAccountownerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestAccountownerInnerWithDefaults

`func NewUpdateAccountRequestAccountownerInnerWithDefaults() *UpdateAccountRequestAccountownerInner`

NewUpdateAccountRequestAccountownerInnerWithDefaults instantiates a new UpdateAccountRequestAccountownerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateAccountRequestAccountownerInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAccountRequestAccountownerInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAccountRequestAccountownerInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *UpdateAccountRequestAccountownerInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAccountRequestAccountownerInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateAccountRequestAccountownerInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetRank

`func (o *UpdateAccountRequestAccountownerInner) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *UpdateAccountRequestAccountownerInner) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *UpdateAccountRequestAccountownerInner) SetRank(v string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *UpdateAccountRequestAccountownerInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetAction

`func (o *UpdateAccountRequestAccountownerInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateAccountRequestAccountownerInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateAccountRequestAccountownerInner) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


