# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | The security system (not the display name). | 
**Endpoint** | **string** | The endpoint (not the display name). | 
**Name** | **string** | The account name. | 
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
**Description** | Pointer to **string** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Accountid** | Pointer to **string** |  | [optional] 
**Passwordchangestatus** | Pointer to **string** |  | [optional] 
**Privileged** | Pointer to **string** |  | [optional] 
**Usergroup** | Pointer to **string** |  | [optional] 
**Updateuser** | Pointer to **string** | Username of the user updating the account. | [optional] 
**Status** | Pointer to **string** | Account status, e.g., Manually Suspended, Manually Provisioned, SUSPENDED FROM IMPORT SERVICE, or numeric values like 1, 2, 3, 4 (default: 1). | [optional] 
**Accounttype** | Pointer to **string** | The account type imported from a third-party system, e.g., Service Account, Shared Account, FIREFIGHTERID. | [optional] 
**Incorrectlogons** | Pointer to **string** |  | [optional] 
**Orphan** | Pointer to **string** |  | [optional] 
**Validfrom** | Pointer to **string** | Date in MM-dd-yyyy format. | [optional] 
**Validthrough** | Pointer to **string** | Date in MM-dd-yyyy format. | [optional] 
**Lastlogondate** | Pointer to **string** | Date in MM-dd-yyyy format. | [optional] 
**Passwordlockdate** | Pointer to **string** | Date in MM-dd-yyyy format. | [optional] 
**Lastpasswordchange** | Pointer to **string** | Date in MM-dd-yyyy format. | [optional] 
**Accountowner** | Pointer to [**[]UpdateAccountRequestAccountownerInner**](UpdateAccountRequestAccountownerInner.md) | List of account owner objects. Each object must include type, value, and action; rank is optional (default: &#39;1&#39;). | [optional] 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest(securitysystem string, endpoint string, name string, ) *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *UpdateAccountRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *UpdateAccountRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *UpdateAccountRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *UpdateAccountRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateAccountRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateAccountRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetName

`func (o *UpdateAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCustomproperty1

`func (o *UpdateAccountRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *UpdateAccountRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *UpdateAccountRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *UpdateAccountRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *UpdateAccountRequest) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *UpdateAccountRequest) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *UpdateAccountRequest) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *UpdateAccountRequest) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *UpdateAccountRequest) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *UpdateAccountRequest) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *UpdateAccountRequest) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *UpdateAccountRequest) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *UpdateAccountRequest) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *UpdateAccountRequest) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *UpdateAccountRequest) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *UpdateAccountRequest) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *UpdateAccountRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *UpdateAccountRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *UpdateAccountRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *UpdateAccountRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *UpdateAccountRequest) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *UpdateAccountRequest) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *UpdateAccountRequest) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *UpdateAccountRequest) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *UpdateAccountRequest) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *UpdateAccountRequest) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *UpdateAccountRequest) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *UpdateAccountRequest) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *UpdateAccountRequest) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *UpdateAccountRequest) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *UpdateAccountRequest) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *UpdateAccountRequest) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *UpdateAccountRequest) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *UpdateAccountRequest) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *UpdateAccountRequest) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *UpdateAccountRequest) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *UpdateAccountRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *UpdateAccountRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *UpdateAccountRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *UpdateAccountRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *UpdateAccountRequest) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *UpdateAccountRequest) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *UpdateAccountRequest) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *UpdateAccountRequest) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *UpdateAccountRequest) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *UpdateAccountRequest) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *UpdateAccountRequest) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *UpdateAccountRequest) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *UpdateAccountRequest) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *UpdateAccountRequest) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *UpdateAccountRequest) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *UpdateAccountRequest) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *UpdateAccountRequest) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *UpdateAccountRequest) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *UpdateAccountRequest) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *UpdateAccountRequest) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *UpdateAccountRequest) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *UpdateAccountRequest) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *UpdateAccountRequest) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *UpdateAccountRequest) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *UpdateAccountRequest) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *UpdateAccountRequest) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *UpdateAccountRequest) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *UpdateAccountRequest) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *UpdateAccountRequest) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *UpdateAccountRequest) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *UpdateAccountRequest) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *UpdateAccountRequest) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *UpdateAccountRequest) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *UpdateAccountRequest) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *UpdateAccountRequest) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *UpdateAccountRequest) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *UpdateAccountRequest) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *UpdateAccountRequest) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *UpdateAccountRequest) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *UpdateAccountRequest) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *UpdateAccountRequest) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *UpdateAccountRequest) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *UpdateAccountRequest) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *UpdateAccountRequest) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *UpdateAccountRequest) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *UpdateAccountRequest) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *UpdateAccountRequest) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *UpdateAccountRequest) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *UpdateAccountRequest) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *UpdateAccountRequest) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *UpdateAccountRequest) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *UpdateAccountRequest) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *UpdateAccountRequest) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *UpdateAccountRequest) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *UpdateAccountRequest) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *UpdateAccountRequest) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *UpdateAccountRequest) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *UpdateAccountRequest) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *UpdateAccountRequest) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *UpdateAccountRequest) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *UpdateAccountRequest) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *UpdateAccountRequest) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *UpdateAccountRequest) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *UpdateAccountRequest) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *UpdateAccountRequest) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *UpdateAccountRequest) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *UpdateAccountRequest) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *UpdateAccountRequest) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *UpdateAccountRequest) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *UpdateAccountRequest) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *UpdateAccountRequest) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *UpdateAccountRequest) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *UpdateAccountRequest) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *UpdateAccountRequest) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *UpdateAccountRequest) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *UpdateAccountRequest) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *UpdateAccountRequest) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *UpdateAccountRequest) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *UpdateAccountRequest) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *UpdateAccountRequest) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *UpdateAccountRequest) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *UpdateAccountRequest) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *UpdateAccountRequest) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *UpdateAccountRequest) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *UpdateAccountRequest) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *UpdateAccountRequest) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *UpdateAccountRequest) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *UpdateAccountRequest) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *UpdateAccountRequest) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *UpdateAccountRequest) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *UpdateAccountRequest) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *UpdateAccountRequest) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *UpdateAccountRequest) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *UpdateAccountRequest) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *UpdateAccountRequest) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *UpdateAccountRequest) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *UpdateAccountRequest) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *UpdateAccountRequest) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *UpdateAccountRequest) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *UpdateAccountRequest) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *UpdateAccountRequest) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *UpdateAccountRequest) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *UpdateAccountRequest) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *UpdateAccountRequest) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *UpdateAccountRequest) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *UpdateAccountRequest) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *UpdateAccountRequest) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *UpdateAccountRequest) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *UpdateAccountRequest) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *UpdateAccountRequest) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *UpdateAccountRequest) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *UpdateAccountRequest) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *UpdateAccountRequest) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *UpdateAccountRequest) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *UpdateAccountRequest) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *UpdateAccountRequest) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *UpdateAccountRequest) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *UpdateAccountRequest) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *UpdateAccountRequest) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *UpdateAccountRequest) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *UpdateAccountRequest) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *UpdateAccountRequest) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *UpdateAccountRequest) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *UpdateAccountRequest) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetCustomproperty41

`func (o *UpdateAccountRequest) GetCustomproperty41() string`

GetCustomproperty41 returns the Customproperty41 field if non-nil, zero value otherwise.

### GetCustomproperty41Ok

`func (o *UpdateAccountRequest) GetCustomproperty41Ok() (*string, bool)`

GetCustomproperty41Ok returns a tuple with the Customproperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41

`func (o *UpdateAccountRequest) SetCustomproperty41(v string)`

SetCustomproperty41 sets Customproperty41 field to given value.

### HasCustomproperty41

`func (o *UpdateAccountRequest) HasCustomproperty41() bool`

HasCustomproperty41 returns a boolean if a field has been set.

### GetCustomproperty42

`func (o *UpdateAccountRequest) GetCustomproperty42() string`

GetCustomproperty42 returns the Customproperty42 field if non-nil, zero value otherwise.

### GetCustomproperty42Ok

`func (o *UpdateAccountRequest) GetCustomproperty42Ok() (*string, bool)`

GetCustomproperty42Ok returns a tuple with the Customproperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42

`func (o *UpdateAccountRequest) SetCustomproperty42(v string)`

SetCustomproperty42 sets Customproperty42 field to given value.

### HasCustomproperty42

`func (o *UpdateAccountRequest) HasCustomproperty42() bool`

HasCustomproperty42 returns a boolean if a field has been set.

### GetCustomproperty43

`func (o *UpdateAccountRequest) GetCustomproperty43() string`

GetCustomproperty43 returns the Customproperty43 field if non-nil, zero value otherwise.

### GetCustomproperty43Ok

`func (o *UpdateAccountRequest) GetCustomproperty43Ok() (*string, bool)`

GetCustomproperty43Ok returns a tuple with the Customproperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43

`func (o *UpdateAccountRequest) SetCustomproperty43(v string)`

SetCustomproperty43 sets Customproperty43 field to given value.

### HasCustomproperty43

`func (o *UpdateAccountRequest) HasCustomproperty43() bool`

HasCustomproperty43 returns a boolean if a field has been set.

### GetCustomproperty44

`func (o *UpdateAccountRequest) GetCustomproperty44() string`

GetCustomproperty44 returns the Customproperty44 field if non-nil, zero value otherwise.

### GetCustomproperty44Ok

`func (o *UpdateAccountRequest) GetCustomproperty44Ok() (*string, bool)`

GetCustomproperty44Ok returns a tuple with the Customproperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44

`func (o *UpdateAccountRequest) SetCustomproperty44(v string)`

SetCustomproperty44 sets Customproperty44 field to given value.

### HasCustomproperty44

`func (o *UpdateAccountRequest) HasCustomproperty44() bool`

HasCustomproperty44 returns a boolean if a field has been set.

### GetCustomproperty45

`func (o *UpdateAccountRequest) GetCustomproperty45() string`

GetCustomproperty45 returns the Customproperty45 field if non-nil, zero value otherwise.

### GetCustomproperty45Ok

`func (o *UpdateAccountRequest) GetCustomproperty45Ok() (*string, bool)`

GetCustomproperty45Ok returns a tuple with the Customproperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45

`func (o *UpdateAccountRequest) SetCustomproperty45(v string)`

SetCustomproperty45 sets Customproperty45 field to given value.

### HasCustomproperty45

`func (o *UpdateAccountRequest) HasCustomproperty45() bool`

HasCustomproperty45 returns a boolean if a field has been set.

### GetCustomproperty46

`func (o *UpdateAccountRequest) GetCustomproperty46() string`

GetCustomproperty46 returns the Customproperty46 field if non-nil, zero value otherwise.

### GetCustomproperty46Ok

`func (o *UpdateAccountRequest) GetCustomproperty46Ok() (*string, bool)`

GetCustomproperty46Ok returns a tuple with the Customproperty46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty46

`func (o *UpdateAccountRequest) SetCustomproperty46(v string)`

SetCustomproperty46 sets Customproperty46 field to given value.

### HasCustomproperty46

`func (o *UpdateAccountRequest) HasCustomproperty46() bool`

HasCustomproperty46 returns a boolean if a field has been set.

### GetCustomproperty47

`func (o *UpdateAccountRequest) GetCustomproperty47() string`

GetCustomproperty47 returns the Customproperty47 field if non-nil, zero value otherwise.

### GetCustomproperty47Ok

`func (o *UpdateAccountRequest) GetCustomproperty47Ok() (*string, bool)`

GetCustomproperty47Ok returns a tuple with the Customproperty47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty47

`func (o *UpdateAccountRequest) SetCustomproperty47(v string)`

SetCustomproperty47 sets Customproperty47 field to given value.

### HasCustomproperty47

`func (o *UpdateAccountRequest) HasCustomproperty47() bool`

HasCustomproperty47 returns a boolean if a field has been set.

### GetCustomproperty48

`func (o *UpdateAccountRequest) GetCustomproperty48() string`

GetCustomproperty48 returns the Customproperty48 field if non-nil, zero value otherwise.

### GetCustomproperty48Ok

`func (o *UpdateAccountRequest) GetCustomproperty48Ok() (*string, bool)`

GetCustomproperty48Ok returns a tuple with the Customproperty48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty48

`func (o *UpdateAccountRequest) SetCustomproperty48(v string)`

SetCustomproperty48 sets Customproperty48 field to given value.

### HasCustomproperty48

`func (o *UpdateAccountRequest) HasCustomproperty48() bool`

HasCustomproperty48 returns a boolean if a field has been set.

### GetCustomproperty49

`func (o *UpdateAccountRequest) GetCustomproperty49() string`

GetCustomproperty49 returns the Customproperty49 field if non-nil, zero value otherwise.

### GetCustomproperty49Ok

`func (o *UpdateAccountRequest) GetCustomproperty49Ok() (*string, bool)`

GetCustomproperty49Ok returns a tuple with the Customproperty49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty49

`func (o *UpdateAccountRequest) SetCustomproperty49(v string)`

SetCustomproperty49 sets Customproperty49 field to given value.

### HasCustomproperty49

`func (o *UpdateAccountRequest) HasCustomproperty49() bool`

HasCustomproperty49 returns a boolean if a field has been set.

### GetCustomproperty50

`func (o *UpdateAccountRequest) GetCustomproperty50() string`

GetCustomproperty50 returns the Customproperty50 field if non-nil, zero value otherwise.

### GetCustomproperty50Ok

`func (o *UpdateAccountRequest) GetCustomproperty50Ok() (*string, bool)`

GetCustomproperty50Ok returns a tuple with the Customproperty50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty50

`func (o *UpdateAccountRequest) SetCustomproperty50(v string)`

SetCustomproperty50 sets Customproperty50 field to given value.

### HasCustomproperty50

`func (o *UpdateAccountRequest) HasCustomproperty50() bool`

HasCustomproperty50 returns a boolean if a field has been set.

### GetCustomproperty51

`func (o *UpdateAccountRequest) GetCustomproperty51() string`

GetCustomproperty51 returns the Customproperty51 field if non-nil, zero value otherwise.

### GetCustomproperty51Ok

`func (o *UpdateAccountRequest) GetCustomproperty51Ok() (*string, bool)`

GetCustomproperty51Ok returns a tuple with the Customproperty51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty51

`func (o *UpdateAccountRequest) SetCustomproperty51(v string)`

SetCustomproperty51 sets Customproperty51 field to given value.

### HasCustomproperty51

`func (o *UpdateAccountRequest) HasCustomproperty51() bool`

HasCustomproperty51 returns a boolean if a field has been set.

### GetCustomproperty52

`func (o *UpdateAccountRequest) GetCustomproperty52() string`

GetCustomproperty52 returns the Customproperty52 field if non-nil, zero value otherwise.

### GetCustomproperty52Ok

`func (o *UpdateAccountRequest) GetCustomproperty52Ok() (*string, bool)`

GetCustomproperty52Ok returns a tuple with the Customproperty52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty52

`func (o *UpdateAccountRequest) SetCustomproperty52(v string)`

SetCustomproperty52 sets Customproperty52 field to given value.

### HasCustomproperty52

`func (o *UpdateAccountRequest) HasCustomproperty52() bool`

HasCustomproperty52 returns a boolean if a field has been set.

### GetCustomproperty53

`func (o *UpdateAccountRequest) GetCustomproperty53() string`

GetCustomproperty53 returns the Customproperty53 field if non-nil, zero value otherwise.

### GetCustomproperty53Ok

`func (o *UpdateAccountRequest) GetCustomproperty53Ok() (*string, bool)`

GetCustomproperty53Ok returns a tuple with the Customproperty53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty53

`func (o *UpdateAccountRequest) SetCustomproperty53(v string)`

SetCustomproperty53 sets Customproperty53 field to given value.

### HasCustomproperty53

`func (o *UpdateAccountRequest) HasCustomproperty53() bool`

HasCustomproperty53 returns a boolean if a field has been set.

### GetCustomproperty54

`func (o *UpdateAccountRequest) GetCustomproperty54() string`

GetCustomproperty54 returns the Customproperty54 field if non-nil, zero value otherwise.

### GetCustomproperty54Ok

`func (o *UpdateAccountRequest) GetCustomproperty54Ok() (*string, bool)`

GetCustomproperty54Ok returns a tuple with the Customproperty54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty54

`func (o *UpdateAccountRequest) SetCustomproperty54(v string)`

SetCustomproperty54 sets Customproperty54 field to given value.

### HasCustomproperty54

`func (o *UpdateAccountRequest) HasCustomproperty54() bool`

HasCustomproperty54 returns a boolean if a field has been set.

### GetCustomproperty55

`func (o *UpdateAccountRequest) GetCustomproperty55() string`

GetCustomproperty55 returns the Customproperty55 field if non-nil, zero value otherwise.

### GetCustomproperty55Ok

`func (o *UpdateAccountRequest) GetCustomproperty55Ok() (*string, bool)`

GetCustomproperty55Ok returns a tuple with the Customproperty55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty55

`func (o *UpdateAccountRequest) SetCustomproperty55(v string)`

SetCustomproperty55 sets Customproperty55 field to given value.

### HasCustomproperty55

`func (o *UpdateAccountRequest) HasCustomproperty55() bool`

HasCustomproperty55 returns a boolean if a field has been set.

### GetCustomproperty56

`func (o *UpdateAccountRequest) GetCustomproperty56() string`

GetCustomproperty56 returns the Customproperty56 field if non-nil, zero value otherwise.

### GetCustomproperty56Ok

`func (o *UpdateAccountRequest) GetCustomproperty56Ok() (*string, bool)`

GetCustomproperty56Ok returns a tuple with the Customproperty56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty56

`func (o *UpdateAccountRequest) SetCustomproperty56(v string)`

SetCustomproperty56 sets Customproperty56 field to given value.

### HasCustomproperty56

`func (o *UpdateAccountRequest) HasCustomproperty56() bool`

HasCustomproperty56 returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAccountRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayname

`func (o *UpdateAccountRequest) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *UpdateAccountRequest) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *UpdateAccountRequest) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *UpdateAccountRequest) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetComments

`func (o *UpdateAccountRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateAccountRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateAccountRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateAccountRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAccountid

`func (o *UpdateAccountRequest) GetAccountid() string`

GetAccountid returns the Accountid field if non-nil, zero value otherwise.

### GetAccountidOk

`func (o *UpdateAccountRequest) GetAccountidOk() (*string, bool)`

GetAccountidOk returns a tuple with the Accountid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountid

`func (o *UpdateAccountRequest) SetAccountid(v string)`

SetAccountid sets Accountid field to given value.

### HasAccountid

`func (o *UpdateAccountRequest) HasAccountid() bool`

HasAccountid returns a boolean if a field has been set.

### GetPasswordchangestatus

`func (o *UpdateAccountRequest) GetPasswordchangestatus() string`

GetPasswordchangestatus returns the Passwordchangestatus field if non-nil, zero value otherwise.

### GetPasswordchangestatusOk

`func (o *UpdateAccountRequest) GetPasswordchangestatusOk() (*string, bool)`

GetPasswordchangestatusOk returns a tuple with the Passwordchangestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordchangestatus

`func (o *UpdateAccountRequest) SetPasswordchangestatus(v string)`

SetPasswordchangestatus sets Passwordchangestatus field to given value.

### HasPasswordchangestatus

`func (o *UpdateAccountRequest) HasPasswordchangestatus() bool`

HasPasswordchangestatus returns a boolean if a field has been set.

### GetPrivileged

`func (o *UpdateAccountRequest) GetPrivileged() string`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *UpdateAccountRequest) GetPrivilegedOk() (*string, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *UpdateAccountRequest) SetPrivileged(v string)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *UpdateAccountRequest) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetUsergroup

`func (o *UpdateAccountRequest) GetUsergroup() string`

GetUsergroup returns the Usergroup field if non-nil, zero value otherwise.

### GetUsergroupOk

`func (o *UpdateAccountRequest) GetUsergroupOk() (*string, bool)`

GetUsergroupOk returns a tuple with the Usergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroup

`func (o *UpdateAccountRequest) SetUsergroup(v string)`

SetUsergroup sets Usergroup field to given value.

### HasUsergroup

`func (o *UpdateAccountRequest) HasUsergroup() bool`

HasUsergroup returns a boolean if a field has been set.

### GetUpdateuser

`func (o *UpdateAccountRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *UpdateAccountRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *UpdateAccountRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *UpdateAccountRequest) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAccountRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAccountRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAccountRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAccountRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccounttype

`func (o *UpdateAccountRequest) GetAccounttype() string`

GetAccounttype returns the Accounttype field if non-nil, zero value otherwise.

### GetAccounttypeOk

`func (o *UpdateAccountRequest) GetAccounttypeOk() (*string, bool)`

GetAccounttypeOk returns a tuple with the Accounttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounttype

`func (o *UpdateAccountRequest) SetAccounttype(v string)`

SetAccounttype sets Accounttype field to given value.

### HasAccounttype

`func (o *UpdateAccountRequest) HasAccounttype() bool`

HasAccounttype returns a boolean if a field has been set.

### GetIncorrectlogons

`func (o *UpdateAccountRequest) GetIncorrectlogons() string`

GetIncorrectlogons returns the Incorrectlogons field if non-nil, zero value otherwise.

### GetIncorrectlogonsOk

`func (o *UpdateAccountRequest) GetIncorrectlogonsOk() (*string, bool)`

GetIncorrectlogonsOk returns a tuple with the Incorrectlogons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorrectlogons

`func (o *UpdateAccountRequest) SetIncorrectlogons(v string)`

SetIncorrectlogons sets Incorrectlogons field to given value.

### HasIncorrectlogons

`func (o *UpdateAccountRequest) HasIncorrectlogons() bool`

HasIncorrectlogons returns a boolean if a field has been set.

### GetOrphan

`func (o *UpdateAccountRequest) GetOrphan() string`

GetOrphan returns the Orphan field if non-nil, zero value otherwise.

### GetOrphanOk

`func (o *UpdateAccountRequest) GetOrphanOk() (*string, bool)`

GetOrphanOk returns a tuple with the Orphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphan

`func (o *UpdateAccountRequest) SetOrphan(v string)`

SetOrphan sets Orphan field to given value.

### HasOrphan

`func (o *UpdateAccountRequest) HasOrphan() bool`

HasOrphan returns a boolean if a field has been set.

### GetValidfrom

`func (o *UpdateAccountRequest) GetValidfrom() string`

GetValidfrom returns the Validfrom field if non-nil, zero value otherwise.

### GetValidfromOk

`func (o *UpdateAccountRequest) GetValidfromOk() (*string, bool)`

GetValidfromOk returns a tuple with the Validfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidfrom

`func (o *UpdateAccountRequest) SetValidfrom(v string)`

SetValidfrom sets Validfrom field to given value.

### HasValidfrom

`func (o *UpdateAccountRequest) HasValidfrom() bool`

HasValidfrom returns a boolean if a field has been set.

### GetValidthrough

`func (o *UpdateAccountRequest) GetValidthrough() string`

GetValidthrough returns the Validthrough field if non-nil, zero value otherwise.

### GetValidthroughOk

`func (o *UpdateAccountRequest) GetValidthroughOk() (*string, bool)`

GetValidthroughOk returns a tuple with the Validthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidthrough

`func (o *UpdateAccountRequest) SetValidthrough(v string)`

SetValidthrough sets Validthrough field to given value.

### HasValidthrough

`func (o *UpdateAccountRequest) HasValidthrough() bool`

HasValidthrough returns a boolean if a field has been set.

### GetLastlogondate

`func (o *UpdateAccountRequest) GetLastlogondate() string`

GetLastlogondate returns the Lastlogondate field if non-nil, zero value otherwise.

### GetLastlogondateOk

`func (o *UpdateAccountRequest) GetLastlogondateOk() (*string, bool)`

GetLastlogondateOk returns a tuple with the Lastlogondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastlogondate

`func (o *UpdateAccountRequest) SetLastlogondate(v string)`

SetLastlogondate sets Lastlogondate field to given value.

### HasLastlogondate

`func (o *UpdateAccountRequest) HasLastlogondate() bool`

HasLastlogondate returns a boolean if a field has been set.

### GetPasswordlockdate

`func (o *UpdateAccountRequest) GetPasswordlockdate() string`

GetPasswordlockdate returns the Passwordlockdate field if non-nil, zero value otherwise.

### GetPasswordlockdateOk

`func (o *UpdateAccountRequest) GetPasswordlockdateOk() (*string, bool)`

GetPasswordlockdateOk returns a tuple with the Passwordlockdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlockdate

`func (o *UpdateAccountRequest) SetPasswordlockdate(v string)`

SetPasswordlockdate sets Passwordlockdate field to given value.

### HasPasswordlockdate

`func (o *UpdateAccountRequest) HasPasswordlockdate() bool`

HasPasswordlockdate returns a boolean if a field has been set.

### GetLastpasswordchange

`func (o *UpdateAccountRequest) GetLastpasswordchange() string`

GetLastpasswordchange returns the Lastpasswordchange field if non-nil, zero value otherwise.

### GetLastpasswordchangeOk

`func (o *UpdateAccountRequest) GetLastpasswordchangeOk() (*string, bool)`

GetLastpasswordchangeOk returns a tuple with the Lastpasswordchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastpasswordchange

`func (o *UpdateAccountRequest) SetLastpasswordchange(v string)`

SetLastpasswordchange sets Lastpasswordchange field to given value.

### HasLastpasswordchange

`func (o *UpdateAccountRequest) HasLastpasswordchange() bool`

HasLastpasswordchange returns a boolean if a field has been set.

### GetAccountowner

`func (o *UpdateAccountRequest) GetAccountowner() []UpdateAccountRequestAccountownerInner`

GetAccountowner returns the Accountowner field if non-nil, zero value otherwise.

### GetAccountownerOk

`func (o *UpdateAccountRequest) GetAccountownerOk() (*[]UpdateAccountRequestAccountownerInner, bool)`

GetAccountownerOk returns a tuple with the Accountowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountowner

`func (o *UpdateAccountRequest) SetAccountowner(v []UpdateAccountRequestAccountownerInner)`

SetAccountowner sets Accountowner field to given value.

### HasAccountowner

`func (o *UpdateAccountRequest) HasAccountowner() bool`

HasAccountowner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


