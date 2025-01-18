# Delegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegatekey** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parentusername** | Pointer to **string** |  | [optional] 
**ParentFirstName** | Pointer to **string** |  | [optional] 
**ParentLastName** | Pointer to **string** |  | [optional] 
**DelegateFirstName** | Pointer to **string** |  | [optional] 
**DelegateLastName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Delegateusername** | Pointer to **string** |  | [optional] 
**Startdate** | Pointer to **string** | date format 07/14/2020 | [optional] 
**Enddate** | Pointer to **string** | date with format 07/14/2020 | [optional] 

## Methods

### NewDelegate

`func NewDelegate() *Delegate`

NewDelegate instantiates a new Delegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateWithDefaults

`func NewDelegateWithDefaults() *Delegate`

NewDelegateWithDefaults instantiates a new Delegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatekey

`func (o *Delegate) GetDelegatekey() string`

GetDelegatekey returns the Delegatekey field if non-nil, zero value otherwise.

### GetDelegatekeyOk

`func (o *Delegate) GetDelegatekeyOk() (*string, bool)`

GetDelegatekeyOk returns a tuple with the Delegatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatekey

`func (o *Delegate) SetDelegatekey(v string)`

SetDelegatekey sets Delegatekey field to given value.

### HasDelegatekey

`func (o *Delegate) HasDelegatekey() bool`

HasDelegatekey returns a boolean if a field has been set.

### GetName

`func (o *Delegate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Delegate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Delegate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Delegate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Delegate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Delegate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Delegate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Delegate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentusername

`func (o *Delegate) GetParentusername() string`

GetParentusername returns the Parentusername field if non-nil, zero value otherwise.

### GetParentusernameOk

`func (o *Delegate) GetParentusernameOk() (*string, bool)`

GetParentusernameOk returns a tuple with the Parentusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentusername

`func (o *Delegate) SetParentusername(v string)`

SetParentusername sets Parentusername field to given value.

### HasParentusername

`func (o *Delegate) HasParentusername() bool`

HasParentusername returns a boolean if a field has been set.

### GetParentFirstName

`func (o *Delegate) GetParentFirstName() string`

GetParentFirstName returns the ParentFirstName field if non-nil, zero value otherwise.

### GetParentFirstNameOk

`func (o *Delegate) GetParentFirstNameOk() (*string, bool)`

GetParentFirstNameOk returns a tuple with the ParentFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFirstName

`func (o *Delegate) SetParentFirstName(v string)`

SetParentFirstName sets ParentFirstName field to given value.

### HasParentFirstName

`func (o *Delegate) HasParentFirstName() bool`

HasParentFirstName returns a boolean if a field has been set.

### GetParentLastName

`func (o *Delegate) GetParentLastName() string`

GetParentLastName returns the ParentLastName field if non-nil, zero value otherwise.

### GetParentLastNameOk

`func (o *Delegate) GetParentLastNameOk() (*string, bool)`

GetParentLastNameOk returns a tuple with the ParentLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLastName

`func (o *Delegate) SetParentLastName(v string)`

SetParentLastName sets ParentLastName field to given value.

### HasParentLastName

`func (o *Delegate) HasParentLastName() bool`

HasParentLastName returns a boolean if a field has been set.

### GetDelegateFirstName

`func (o *Delegate) GetDelegateFirstName() string`

GetDelegateFirstName returns the DelegateFirstName field if non-nil, zero value otherwise.

### GetDelegateFirstNameOk

`func (o *Delegate) GetDelegateFirstNameOk() (*string, bool)`

GetDelegateFirstNameOk returns a tuple with the DelegateFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateFirstName

`func (o *Delegate) SetDelegateFirstName(v string)`

SetDelegateFirstName sets DelegateFirstName field to given value.

### HasDelegateFirstName

`func (o *Delegate) HasDelegateFirstName() bool`

HasDelegateFirstName returns a boolean if a field has been set.

### GetDelegateLastName

`func (o *Delegate) GetDelegateLastName() string`

GetDelegateLastName returns the DelegateLastName field if non-nil, zero value otherwise.

### GetDelegateLastNameOk

`func (o *Delegate) GetDelegateLastNameOk() (*string, bool)`

GetDelegateLastNameOk returns a tuple with the DelegateLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateLastName

`func (o *Delegate) SetDelegateLastName(v string)`

SetDelegateLastName sets DelegateLastName field to given value.

### HasDelegateLastName

`func (o *Delegate) HasDelegateLastName() bool`

HasDelegateLastName returns a boolean if a field has been set.

### GetStatus

`func (o *Delegate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Delegate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Delegate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Delegate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDelegateusername

`func (o *Delegate) GetDelegateusername() string`

GetDelegateusername returns the Delegateusername field if non-nil, zero value otherwise.

### GetDelegateusernameOk

`func (o *Delegate) GetDelegateusernameOk() (*string, bool)`

GetDelegateusernameOk returns a tuple with the Delegateusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateusername

`func (o *Delegate) SetDelegateusername(v string)`

SetDelegateusername sets Delegateusername field to given value.

### HasDelegateusername

`func (o *Delegate) HasDelegateusername() bool`

HasDelegateusername returns a boolean if a field has been set.

### GetStartdate

`func (o *Delegate) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *Delegate) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *Delegate) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *Delegate) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *Delegate) GetEnddate() string`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *Delegate) GetEnddateOk() (*string, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *Delegate) SetEnddate(v string)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *Delegate) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


