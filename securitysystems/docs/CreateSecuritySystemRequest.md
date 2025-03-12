# CreateSecuritySystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Systemname** | **string** | Name of the security system. | 
**DisplayName** | **string** | Display name for the security system. | 
**Hostname** | Pointer to **string** | Hostname for the security system. | [optional] 
**Port** | Pointer to **string** | Port number. | [optional] 
**AccessAddWorkflow** | Pointer to **string** | Workflow for adding access. | [optional] 
**AccessRemoveWorkflow** | Pointer to **string** | Workflow for removing access. | [optional] 
**AddServiceAccountWorkflow** | Pointer to **string** | Workflow for adding a service account. | [optional] 
**RemoveServiceAccountWorkflow** | Pointer to **string** | Workflow for removing a service account. | [optional] 
**Connectionparameters** | Pointer to **string** | JSON string with connection parameters. | [optional] 
**AutomatedProvisioning** | Pointer to **string** | Flag to indicate automated provisioning. | [optional] 
**ProvisioningTries** | Pointer to **string** | Number of provisioning tries. | [optional] 

## Methods

### NewCreateSecuritySystemRequest

`func NewCreateSecuritySystemRequest(systemname string, displayName string, ) *CreateSecuritySystemRequest`

NewCreateSecuritySystemRequest instantiates a new CreateSecuritySystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecuritySystemRequestWithDefaults

`func NewCreateSecuritySystemRequestWithDefaults() *CreateSecuritySystemRequest`

NewCreateSecuritySystemRequestWithDefaults instantiates a new CreateSecuritySystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemname

`func (o *CreateSecuritySystemRequest) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *CreateSecuritySystemRequest) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *CreateSecuritySystemRequest) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.


### GetDisplayName

`func (o *CreateSecuritySystemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateSecuritySystemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateSecuritySystemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHostname

`func (o *CreateSecuritySystemRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateSecuritySystemRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateSecuritySystemRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateSecuritySystemRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *CreateSecuritySystemRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateSecuritySystemRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateSecuritySystemRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateSecuritySystemRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) GetAccessAddWorkflow() string`

GetAccessAddWorkflow returns the AccessAddWorkflow field if non-nil, zero value otherwise.

### GetAccessAddWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAccessAddWorkflowOk() (*string, bool)`

GetAccessAddWorkflowOk returns a tuple with the AccessAddWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) SetAccessAddWorkflow(v string)`

SetAccessAddWorkflow sets AccessAddWorkflow field to given value.

### HasAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) HasAccessAddWorkflow() bool`

HasAccessAddWorkflow returns a boolean if a field has been set.

### GetAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) GetAccessRemoveWorkflow() string`

GetAccessRemoveWorkflow returns the AccessRemoveWorkflow field if non-nil, zero value otherwise.

### GetAccessRemoveWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAccessRemoveWorkflowOk() (*string, bool)`

GetAccessRemoveWorkflowOk returns a tuple with the AccessRemoveWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) SetAccessRemoveWorkflow(v string)`

SetAccessRemoveWorkflow sets AccessRemoveWorkflow field to given value.

### HasAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) HasAccessRemoveWorkflow() bool`

HasAccessRemoveWorkflow returns a boolean if a field has been set.

### GetAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) GetAddServiceAccountWorkflow() string`

GetAddServiceAccountWorkflow returns the AddServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetAddServiceAccountWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAddServiceAccountWorkflowOk() (*string, bool)`

GetAddServiceAccountWorkflowOk returns a tuple with the AddServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) SetAddServiceAccountWorkflow(v string)`

SetAddServiceAccountWorkflow sets AddServiceAccountWorkflow field to given value.

### HasAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) HasAddServiceAccountWorkflow() bool`

HasAddServiceAccountWorkflow returns a boolean if a field has been set.

### GetRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) GetRemoveServiceAccountWorkflow() string`

GetRemoveServiceAccountWorkflow returns the RemoveServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetRemoveServiceAccountWorkflowOk

`func (o *CreateSecuritySystemRequest) GetRemoveServiceAccountWorkflowOk() (*string, bool)`

GetRemoveServiceAccountWorkflowOk returns a tuple with the RemoveServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) SetRemoveServiceAccountWorkflow(v string)`

SetRemoveServiceAccountWorkflow sets RemoveServiceAccountWorkflow field to given value.

### HasRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) HasRemoveServiceAccountWorkflow() bool`

HasRemoveServiceAccountWorkflow returns a boolean if a field has been set.

### GetConnectionparameters

`func (o *CreateSecuritySystemRequest) GetConnectionparameters() string`

GetConnectionparameters returns the Connectionparameters field if non-nil, zero value otherwise.

### GetConnectionparametersOk

`func (o *CreateSecuritySystemRequest) GetConnectionparametersOk() (*string, bool)`

GetConnectionparametersOk returns a tuple with the Connectionparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionparameters

`func (o *CreateSecuritySystemRequest) SetConnectionparameters(v string)`

SetConnectionparameters sets Connectionparameters field to given value.

### HasConnectionparameters

`func (o *CreateSecuritySystemRequest) HasConnectionparameters() bool`

HasConnectionparameters returns a boolean if a field has been set.

### GetAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) GetAutomatedProvisioning() string`

GetAutomatedProvisioning returns the AutomatedProvisioning field if non-nil, zero value otherwise.

### GetAutomatedProvisioningOk

`func (o *CreateSecuritySystemRequest) GetAutomatedProvisioningOk() (*string, bool)`

GetAutomatedProvisioningOk returns a tuple with the AutomatedProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) SetAutomatedProvisioning(v string)`

SetAutomatedProvisioning sets AutomatedProvisioning field to given value.

### HasAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) HasAutomatedProvisioning() bool`

HasAutomatedProvisioning returns a boolean if a field has been set.

### GetProvisioningTries

`func (o *CreateSecuritySystemRequest) GetProvisioningTries() string`

GetProvisioningTries returns the ProvisioningTries field if non-nil, zero value otherwise.

### GetProvisioningTriesOk

`func (o *CreateSecuritySystemRequest) GetProvisioningTriesOk() (*string, bool)`

GetProvisioningTriesOk returns a tuple with the ProvisioningTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTries

`func (o *CreateSecuritySystemRequest) SetProvisioningTries(v string)`

SetProvisioningTries sets ProvisioningTries field to given value.

### HasProvisioningTries

`func (o *CreateSecuritySystemRequest) HasProvisioningTries() bool`

HasProvisioningTries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


