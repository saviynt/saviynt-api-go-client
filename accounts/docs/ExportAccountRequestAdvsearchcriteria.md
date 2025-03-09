# ExportAccountRequestAdvsearchcriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Accounttype** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Customproperty1** | Pointer to **string** |  | [optional] 
**Customproperty2** | Pointer to **string** |  | [optional] 
**Customproperty3** | Pointer to **string** |  | [optional] 
**Customproperty4** | Pointer to **string** |  | [optional] 
**Customproperty5** | Pointer to **string** |  | [optional] 
**Customproperty6** | Pointer to **string** |  | [optional] 
**Customproperty7** | Pointer to **string** |  | [optional] 
**Customproperty8** | Pointer to **string** |  | [optional] 
**Customproperty9** | Pointer to **string** |  | [optional] 
**Customproperty10** | Pointer to **string** |  | [optional] 
**Customproperty11** | Pointer to **string** |  | [optional] 
**Customproperty12** | Pointer to **string** |  | [optional] 
**Customproperty13** | Pointer to **string** |  | [optional] 
**Customproperty14** | Pointer to **string** |  | [optional] 
**Customproperty15** | Pointer to **string** |  | [optional] 
**Customproperty16** | Pointer to **string** |  | [optional] 
**Customproperty17** | Pointer to **string** |  | [optional] 
**Customproperty18** | Pointer to **string** |  | [optional] 
**Customproperty19** | Pointer to **string** |  | [optional] 
**Customproperty20** | Pointer to **string** |  | [optional] 
**Customproperty21** | Pointer to **string** |  | [optional] 
**Customproperty22** | Pointer to **string** |  | [optional] 
**Customproperty23** | Pointer to **string** |  | [optional] 
**Customproperty24** | Pointer to **string** |  | [optional] 
**Customproperty25** | Pointer to **string** |  | [optional] 
**Customproperty26** | Pointer to **string** |  | [optional] 
**Customproperty27** | Pointer to **string** |  | [optional] 
**Customproperty28** | Pointer to **string** |  | [optional] 
**Customproperty29** | Pointer to **string** |  | [optional] 
**Customproperty30** | Pointer to **string** |  | [optional] 
**Customproperty31** | Pointer to **string** |  | [optional] 
**Customproperty32** | Pointer to **string** |  | [optional] 
**Customproperty33** | Pointer to **string** |  | [optional] 
**Customproperty34** | Pointer to **string** |  | [optional] 
**Customproperty35** | Pointer to **string** |  | [optional] 
**Customproperty36** | Pointer to **string** |  | [optional] 
**Customproperty37** | Pointer to **string** |  | [optional] 
**Customproperty38** | Pointer to **string** |  | [optional] 
**Customproperty39** | Pointer to **string** |  | [optional] 
**Customproperty40** | Pointer to **string** |  | [optional] 
**Customproperty41** | Pointer to **string** |  | [optional] 
**Customproperty42** | Pointer to **string** |  | [optional] 
**Customproperty43** | Pointer to **string** |  | [optional] 
**Customproperty44** | Pointer to **string** |  | [optional] 
**Customproperty45** | Pointer to **string** |  | [optional] 
**Customproperty46** | Pointer to **string** |  | [optional] 
**Customproperty47** | Pointer to **string** |  | [optional] 
**Customproperty48** | Pointer to **string** |  | [optional] 
**Customproperty49** | Pointer to **string** |  | [optional] 
**Customproperty50** | Pointer to **string** |  | [optional] 
**Customproperty51** | Pointer to **string** |  | [optional] 
**Customproperty52** | Pointer to **string** |  | [optional] 
**Customproperty53** | Pointer to **string** |  | [optional] 
**Customproperty54** | Pointer to **string** |  | [optional] 
**Customproperty55** | Pointer to **string** |  | [optional] 
**Customproperty56** | Pointer to **string** |  | [optional] 
**AccountID** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Updateuser** | Pointer to **string** |  | [optional] 
**ValidfromDate** | Pointer to **string** |  | [optional] 
**Validthrough** | Pointer to **string** |  | [optional] 
**Createdon** | Pointer to **string** |  | [optional] 
**Lastlogondate** | Pointer to **string** |  | [optional] 
**Lastpasswordchange** | Pointer to **string** |  | [optional] 
**Updatedate** | Pointer to **string** |  | [optional] 
**Orphan** | Pointer to **bool** |  | [optional] 
**Accountowner** | Pointer to [**[]ExportAccountRequestAdvsearchcriteriaAccountownerInner**](ExportAccountRequestAdvsearchcriteriaAccountownerInner.md) | List of account owner objects. | [optional] 

## Methods

### NewExportAccountRequestAdvsearchcriteria

`func NewExportAccountRequestAdvsearchcriteria() *ExportAccountRequestAdvsearchcriteria`

NewExportAccountRequestAdvsearchcriteria instantiates a new ExportAccountRequestAdvsearchcriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAccountRequestAdvsearchcriteriaWithDefaults

`func NewExportAccountRequestAdvsearchcriteriaWithDefaults() *ExportAccountRequestAdvsearchcriteria`

NewExportAccountRequestAdvsearchcriteriaWithDefaults instantiates a new ExportAccountRequestAdvsearchcriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *ExportAccountRequestAdvsearchcriteria) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *ExportAccountRequestAdvsearchcriteria) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetDescription

`func (o *ExportAccountRequestAdvsearchcriteria) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportAccountRequestAdvsearchcriteria) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportAccountRequestAdvsearchcriteria) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *ExportAccountRequestAdvsearchcriteria) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ExportAccountRequestAdvsearchcriteria) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ExportAccountRequestAdvsearchcriteria) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAccounttype

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccounttype() string`

GetAccounttype returns the Accounttype field if non-nil, zero value otherwise.

### GetAccounttypeOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccounttypeOk() (*string, bool)`

GetAccounttypeOk returns a tuple with the Accounttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounttype

`func (o *ExportAccountRequestAdvsearchcriteria) SetAccounttype(v string)`

SetAccounttype sets Accounttype field to given value.

### HasAccounttype

`func (o *ExportAccountRequestAdvsearchcriteria) HasAccounttype() bool`

HasAccounttype returns a boolean if a field has been set.

### GetStatus

`func (o *ExportAccountRequestAdvsearchcriteria) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportAccountRequestAdvsearchcriteria) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExportAccountRequestAdvsearchcriteria) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomproperty1

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetCustomproperty41

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty41() string`

GetCustomproperty41 returns the Customproperty41 field if non-nil, zero value otherwise.

### GetCustomproperty41Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty41Ok() (*string, bool)`

GetCustomproperty41Ok returns a tuple with the Customproperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty41(v string)`

SetCustomproperty41 sets Customproperty41 field to given value.

### HasCustomproperty41

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty41() bool`

HasCustomproperty41 returns a boolean if a field has been set.

### GetCustomproperty42

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty42() string`

GetCustomproperty42 returns the Customproperty42 field if non-nil, zero value otherwise.

### GetCustomproperty42Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty42Ok() (*string, bool)`

GetCustomproperty42Ok returns a tuple with the Customproperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty42(v string)`

SetCustomproperty42 sets Customproperty42 field to given value.

### HasCustomproperty42

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty42() bool`

HasCustomproperty42 returns a boolean if a field has been set.

### GetCustomproperty43

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty43() string`

GetCustomproperty43 returns the Customproperty43 field if non-nil, zero value otherwise.

### GetCustomproperty43Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty43Ok() (*string, bool)`

GetCustomproperty43Ok returns a tuple with the Customproperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty43(v string)`

SetCustomproperty43 sets Customproperty43 field to given value.

### HasCustomproperty43

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty43() bool`

HasCustomproperty43 returns a boolean if a field has been set.

### GetCustomproperty44

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty44() string`

GetCustomproperty44 returns the Customproperty44 field if non-nil, zero value otherwise.

### GetCustomproperty44Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty44Ok() (*string, bool)`

GetCustomproperty44Ok returns a tuple with the Customproperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty44(v string)`

SetCustomproperty44 sets Customproperty44 field to given value.

### HasCustomproperty44

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty44() bool`

HasCustomproperty44 returns a boolean if a field has been set.

### GetCustomproperty45

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty45() string`

GetCustomproperty45 returns the Customproperty45 field if non-nil, zero value otherwise.

### GetCustomproperty45Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty45Ok() (*string, bool)`

GetCustomproperty45Ok returns a tuple with the Customproperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty45(v string)`

SetCustomproperty45 sets Customproperty45 field to given value.

### HasCustomproperty45

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty45() bool`

HasCustomproperty45 returns a boolean if a field has been set.

### GetCustomproperty46

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty46() string`

GetCustomproperty46 returns the Customproperty46 field if non-nil, zero value otherwise.

### GetCustomproperty46Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty46Ok() (*string, bool)`

GetCustomproperty46Ok returns a tuple with the Customproperty46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty46

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty46(v string)`

SetCustomproperty46 sets Customproperty46 field to given value.

### HasCustomproperty46

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty46() bool`

HasCustomproperty46 returns a boolean if a field has been set.

### GetCustomproperty47

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty47() string`

GetCustomproperty47 returns the Customproperty47 field if non-nil, zero value otherwise.

### GetCustomproperty47Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty47Ok() (*string, bool)`

GetCustomproperty47Ok returns a tuple with the Customproperty47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty47

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty47(v string)`

SetCustomproperty47 sets Customproperty47 field to given value.

### HasCustomproperty47

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty47() bool`

HasCustomproperty47 returns a boolean if a field has been set.

### GetCustomproperty48

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty48() string`

GetCustomproperty48 returns the Customproperty48 field if non-nil, zero value otherwise.

### GetCustomproperty48Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty48Ok() (*string, bool)`

GetCustomproperty48Ok returns a tuple with the Customproperty48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty48

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty48(v string)`

SetCustomproperty48 sets Customproperty48 field to given value.

### HasCustomproperty48

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty48() bool`

HasCustomproperty48 returns a boolean if a field has been set.

### GetCustomproperty49

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty49() string`

GetCustomproperty49 returns the Customproperty49 field if non-nil, zero value otherwise.

### GetCustomproperty49Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty49Ok() (*string, bool)`

GetCustomproperty49Ok returns a tuple with the Customproperty49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty49

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty49(v string)`

SetCustomproperty49 sets Customproperty49 field to given value.

### HasCustomproperty49

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty49() bool`

HasCustomproperty49 returns a boolean if a field has been set.

### GetCustomproperty50

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty50() string`

GetCustomproperty50 returns the Customproperty50 field if non-nil, zero value otherwise.

### GetCustomproperty50Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty50Ok() (*string, bool)`

GetCustomproperty50Ok returns a tuple with the Customproperty50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty50

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty50(v string)`

SetCustomproperty50 sets Customproperty50 field to given value.

### HasCustomproperty50

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty50() bool`

HasCustomproperty50 returns a boolean if a field has been set.

### GetCustomproperty51

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty51() string`

GetCustomproperty51 returns the Customproperty51 field if non-nil, zero value otherwise.

### GetCustomproperty51Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty51Ok() (*string, bool)`

GetCustomproperty51Ok returns a tuple with the Customproperty51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty51

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty51(v string)`

SetCustomproperty51 sets Customproperty51 field to given value.

### HasCustomproperty51

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty51() bool`

HasCustomproperty51 returns a boolean if a field has been set.

### GetCustomproperty52

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty52() string`

GetCustomproperty52 returns the Customproperty52 field if non-nil, zero value otherwise.

### GetCustomproperty52Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty52Ok() (*string, bool)`

GetCustomproperty52Ok returns a tuple with the Customproperty52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty52

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty52(v string)`

SetCustomproperty52 sets Customproperty52 field to given value.

### HasCustomproperty52

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty52() bool`

HasCustomproperty52 returns a boolean if a field has been set.

### GetCustomproperty53

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty53() string`

GetCustomproperty53 returns the Customproperty53 field if non-nil, zero value otherwise.

### GetCustomproperty53Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty53Ok() (*string, bool)`

GetCustomproperty53Ok returns a tuple with the Customproperty53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty53

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty53(v string)`

SetCustomproperty53 sets Customproperty53 field to given value.

### HasCustomproperty53

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty53() bool`

HasCustomproperty53 returns a boolean if a field has been set.

### GetCustomproperty54

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty54() string`

GetCustomproperty54 returns the Customproperty54 field if non-nil, zero value otherwise.

### GetCustomproperty54Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty54Ok() (*string, bool)`

GetCustomproperty54Ok returns a tuple with the Customproperty54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty54

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty54(v string)`

SetCustomproperty54 sets Customproperty54 field to given value.

### HasCustomproperty54

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty54() bool`

HasCustomproperty54 returns a boolean if a field has been set.

### GetCustomproperty55

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty55() string`

GetCustomproperty55 returns the Customproperty55 field if non-nil, zero value otherwise.

### GetCustomproperty55Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty55Ok() (*string, bool)`

GetCustomproperty55Ok returns a tuple with the Customproperty55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty55

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty55(v string)`

SetCustomproperty55 sets Customproperty55 field to given value.

### HasCustomproperty55

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty55() bool`

HasCustomproperty55 returns a boolean if a field has been set.

### GetCustomproperty56

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty56() string`

GetCustomproperty56 returns the Customproperty56 field if non-nil, zero value otherwise.

### GetCustomproperty56Ok

`func (o *ExportAccountRequestAdvsearchcriteria) GetCustomproperty56Ok() (*string, bool)`

GetCustomproperty56Ok returns a tuple with the Customproperty56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty56

`func (o *ExportAccountRequestAdvsearchcriteria) SetCustomproperty56(v string)`

SetCustomproperty56 sets Customproperty56 field to given value.

### HasCustomproperty56

`func (o *ExportAccountRequestAdvsearchcriteria) HasCustomproperty56() bool`

HasCustomproperty56 returns a boolean if a field has been set.

### GetAccountID

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *ExportAccountRequestAdvsearchcriteria) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *ExportAccountRequestAdvsearchcriteria) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExportAccountRequestAdvsearchcriteria) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExportAccountRequestAdvsearchcriteria) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExportAccountRequestAdvsearchcriteria) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ExportAccountRequestAdvsearchcriteria) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportAccountRequestAdvsearchcriteria) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExportAccountRequestAdvsearchcriteria) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreator

`func (o *ExportAccountRequestAdvsearchcriteria) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ExportAccountRequestAdvsearchcriteria) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ExportAccountRequestAdvsearchcriteria) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetUpdateuser

`func (o *ExportAccountRequestAdvsearchcriteria) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *ExportAccountRequestAdvsearchcriteria) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *ExportAccountRequestAdvsearchcriteria) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetValidfromDate

`func (o *ExportAccountRequestAdvsearchcriteria) GetValidfromDate() string`

GetValidfromDate returns the ValidfromDate field if non-nil, zero value otherwise.

### GetValidfromDateOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetValidfromDateOk() (*string, bool)`

GetValidfromDateOk returns a tuple with the ValidfromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidfromDate

`func (o *ExportAccountRequestAdvsearchcriteria) SetValidfromDate(v string)`

SetValidfromDate sets ValidfromDate field to given value.

### HasValidfromDate

`func (o *ExportAccountRequestAdvsearchcriteria) HasValidfromDate() bool`

HasValidfromDate returns a boolean if a field has been set.

### GetValidthrough

`func (o *ExportAccountRequestAdvsearchcriteria) GetValidthrough() string`

GetValidthrough returns the Validthrough field if non-nil, zero value otherwise.

### GetValidthroughOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetValidthroughOk() (*string, bool)`

GetValidthroughOk returns a tuple with the Validthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidthrough

`func (o *ExportAccountRequestAdvsearchcriteria) SetValidthrough(v string)`

SetValidthrough sets Validthrough field to given value.

### HasValidthrough

`func (o *ExportAccountRequestAdvsearchcriteria) HasValidthrough() bool`

HasValidthrough returns a boolean if a field has been set.

### GetCreatedon

`func (o *ExportAccountRequestAdvsearchcriteria) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *ExportAccountRequestAdvsearchcriteria) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *ExportAccountRequestAdvsearchcriteria) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetLastlogondate

`func (o *ExportAccountRequestAdvsearchcriteria) GetLastlogondate() string`

GetLastlogondate returns the Lastlogondate field if non-nil, zero value otherwise.

### GetLastlogondateOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetLastlogondateOk() (*string, bool)`

GetLastlogondateOk returns a tuple with the Lastlogondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastlogondate

`func (o *ExportAccountRequestAdvsearchcriteria) SetLastlogondate(v string)`

SetLastlogondate sets Lastlogondate field to given value.

### HasLastlogondate

`func (o *ExportAccountRequestAdvsearchcriteria) HasLastlogondate() bool`

HasLastlogondate returns a boolean if a field has been set.

### GetLastpasswordchange

`func (o *ExportAccountRequestAdvsearchcriteria) GetLastpasswordchange() string`

GetLastpasswordchange returns the Lastpasswordchange field if non-nil, zero value otherwise.

### GetLastpasswordchangeOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetLastpasswordchangeOk() (*string, bool)`

GetLastpasswordchangeOk returns a tuple with the Lastpasswordchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastpasswordchange

`func (o *ExportAccountRequestAdvsearchcriteria) SetLastpasswordchange(v string)`

SetLastpasswordchange sets Lastpasswordchange field to given value.

### HasLastpasswordchange

`func (o *ExportAccountRequestAdvsearchcriteria) HasLastpasswordchange() bool`

HasLastpasswordchange returns a boolean if a field has been set.

### GetUpdatedate

`func (o *ExportAccountRequestAdvsearchcriteria) GetUpdatedate() string`

GetUpdatedate returns the Updatedate field if non-nil, zero value otherwise.

### GetUpdatedateOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetUpdatedateOk() (*string, bool)`

GetUpdatedateOk returns a tuple with the Updatedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedate

`func (o *ExportAccountRequestAdvsearchcriteria) SetUpdatedate(v string)`

SetUpdatedate sets Updatedate field to given value.

### HasUpdatedate

`func (o *ExportAccountRequestAdvsearchcriteria) HasUpdatedate() bool`

HasUpdatedate returns a boolean if a field has been set.

### GetOrphan

`func (o *ExportAccountRequestAdvsearchcriteria) GetOrphan() bool`

GetOrphan returns the Orphan field if non-nil, zero value otherwise.

### GetOrphanOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetOrphanOk() (*bool, bool)`

GetOrphanOk returns a tuple with the Orphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphan

`func (o *ExportAccountRequestAdvsearchcriteria) SetOrphan(v bool)`

SetOrphan sets Orphan field to given value.

### HasOrphan

`func (o *ExportAccountRequestAdvsearchcriteria) HasOrphan() bool`

HasOrphan returns a boolean if a field has been set.

### GetAccountowner

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountowner() []ExportAccountRequestAdvsearchcriteriaAccountownerInner`

GetAccountowner returns the Accountowner field if non-nil, zero value otherwise.

### GetAccountownerOk

`func (o *ExportAccountRequestAdvsearchcriteria) GetAccountownerOk() (*[]ExportAccountRequestAdvsearchcriteriaAccountownerInner, bool)`

GetAccountownerOk returns a tuple with the Accountowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountowner

`func (o *ExportAccountRequestAdvsearchcriteria) SetAccountowner(v []ExportAccountRequestAdvsearchcriteriaAccountownerInner)`

SetAccountowner sets Accountowner field to given value.

### HasAccountowner

`func (o *ExportAccountRequestAdvsearchcriteria) HasAccountowner() bool`

HasAccountowner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


