# VaultConfigurationKeyMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AUTHTOKEN** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to [**VaultConfigurationKeyMappingPassword**](VaultConfigurationKeyMappingPassword.md) |  | [optional] 

## Methods

### NewVaultConfigurationKeyMapping

`func NewVaultConfigurationKeyMapping() *VaultConfigurationKeyMapping`

NewVaultConfigurationKeyMapping instantiates a new VaultConfigurationKeyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultConfigurationKeyMappingWithDefaults

`func NewVaultConfigurationKeyMappingWithDefaults() *VaultConfigurationKeyMapping`

NewVaultConfigurationKeyMappingWithDefaults instantiates a new VaultConfigurationKeyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAUTHTOKEN

`func (o *VaultConfigurationKeyMapping) GetAUTHTOKEN() string`

GetAUTHTOKEN returns the AUTHTOKEN field if non-nil, zero value otherwise.

### GetAUTHTOKENOk

`func (o *VaultConfigurationKeyMapping) GetAUTHTOKENOk() (*string, bool)`

GetAUTHTOKENOk returns a tuple with the AUTHTOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHTOKEN

`func (o *VaultConfigurationKeyMapping) SetAUTHTOKEN(v string)`

SetAUTHTOKEN sets AUTHTOKEN field to given value.

### HasAUTHTOKEN

`func (o *VaultConfigurationKeyMapping) HasAUTHTOKEN() bool`

HasAUTHTOKEN returns a boolean if a field has been set.

### GetPASSWORD

`func (o *VaultConfigurationKeyMapping) GetPASSWORD() VaultConfigurationKeyMappingPassword`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *VaultConfigurationKeyMapping) GetPASSWORDOk() (*VaultConfigurationKeyMappingPassword, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *VaultConfigurationKeyMapping) SetPASSWORD(v VaultConfigurationKeyMappingPassword)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *VaultConfigurationKeyMapping) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


