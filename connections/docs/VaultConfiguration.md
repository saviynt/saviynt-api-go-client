# VaultConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Vault secret path. | [optional] 
**KeyMapping** | Pointer to [**VaultConfigurationKeyMapping**](VaultConfigurationKeyMapping.md) |  | [optional] 

## Methods

### NewVaultConfiguration

`func NewVaultConfiguration() *VaultConfiguration`

NewVaultConfiguration instantiates a new VaultConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultConfigurationWithDefaults

`func NewVaultConfigurationWithDefaults() *VaultConfiguration`

NewVaultConfigurationWithDefaults instantiates a new VaultConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *VaultConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VaultConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VaultConfiguration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VaultConfiguration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetKeyMapping

`func (o *VaultConfiguration) GetKeyMapping() VaultConfigurationKeyMapping`

GetKeyMapping returns the KeyMapping field if non-nil, zero value otherwise.

### GetKeyMappingOk

`func (o *VaultConfiguration) GetKeyMappingOk() (*VaultConfigurationKeyMapping, bool)`

GetKeyMappingOk returns a tuple with the KeyMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMapping

`func (o *VaultConfiguration) SetKeyMapping(v VaultConfigurationKeyMapping)`

SetKeyMapping sets KeyMapping field to given value.

### HasKeyMapping

`func (o *VaultConfiguration) HasKeyMapping() bool`

HasKeyMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


