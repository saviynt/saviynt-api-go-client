# UpdateSecuritySystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Systemname** | **string** | The security system name to update. | 
**DisplayName** | Pointer to **string** | The display name for the security system. | [optional] 
**Hostname** | Pointer to **string** | Hostname. | [optional] 
**Port** | Pointer to **string** | Port number (provided but not used). | [optional] 
**PolicyRule** | Pointer to **string** | Password policy for the security system (Release v24.5+). | [optional] 
**PolicyRuleServiceAccount** | Pointer to **string** | Password policy for service accounts (Release v24.5+). | [optional] 
**AccessAddWorkflow** | Pointer to **string** | Workflow for adding access. | [optional] 
**AccessRemoveWorkflow** | Pointer to **string** | Workflow for removing access. | [optional] 
**AddServiceAccountWorkflow** | Pointer to **string** | Workflow for adding a service account. | [optional] 
**RemoveServiceAccountWorkflow** | Pointer to **string** | Workflow for removing a service account. | [optional] 
**Connectionparameters** | Pointer to **string** | Connection parameters in JSON format. | [optional] 
**AutomatedProvisioning** | Pointer to **string** | Flag indicating automated provisioning. | [optional] 
**Connectionname** | Pointer to **string** | Connection name for identity reconciliation. | [optional] 
**ProvisioningConnection** | Pointer to **string** | Provisioning connection. | [optional] 
**ServiceDeskConnection** | Pointer to **string** | Service desk connection. | [optional] 
**Provisioningcomments** | Pointer to **string** | Comments regarding provisioning. | [optional] 
**Action** | Pointer to **string** | Action to perform, either &#39;enable&#39; or &#39;disable&#39;. | [optional] 
**DefaultSystem** | Pointer to **string** | Default system flag. | [optional] 
**ReconApplication** | Pointer to **string** | Flag indicating reconciliation application usage. | [optional] 
**ProvisioningTries** | Pointer to **string** | Number of tries for provisioning/de-provisioning (1-20). | [optional] 
**InherentSODReportFields** | Pointer to **string** | Comma-separated list of SOD report fields to filter out (Release v24.5+). | [optional] 

## Methods

### NewUpdateSecuritySystemRequest

`func NewUpdateSecuritySystemRequest(systemname string, ) *UpdateSecuritySystemRequest`

NewUpdateSecuritySystemRequest instantiates a new UpdateSecuritySystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecuritySystemRequestWithDefaults

`func NewUpdateSecuritySystemRequestWithDefaults() *UpdateSecuritySystemRequest`

NewUpdateSecuritySystemRequestWithDefaults instantiates a new UpdateSecuritySystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemname

`func (o *UpdateSecuritySystemRequest) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *UpdateSecuritySystemRequest) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *UpdateSecuritySystemRequest) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.


### GetDisplayName

`func (o *UpdateSecuritySystemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateSecuritySystemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateSecuritySystemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateSecuritySystemRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateSecuritySystemRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateSecuritySystemRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateSecuritySystemRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateSecuritySystemRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *UpdateSecuritySystemRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateSecuritySystemRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateSecuritySystemRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateSecuritySystemRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPolicyRule

`func (o *UpdateSecuritySystemRequest) GetPolicyRule() string`

GetPolicyRule returns the PolicyRule field if non-nil, zero value otherwise.

### GetPolicyRuleOk

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleOk() (*string, bool)`

GetPolicyRuleOk returns a tuple with the PolicyRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRule

`func (o *UpdateSecuritySystemRequest) SetPolicyRule(v string)`

SetPolicyRule sets PolicyRule field to given value.

### HasPolicyRule

`func (o *UpdateSecuritySystemRequest) HasPolicyRule() bool`

HasPolicyRule returns a boolean if a field has been set.

### GetPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleServiceAccount() string`

GetPolicyRuleServiceAccount returns the PolicyRuleServiceAccount field if non-nil, zero value otherwise.

### GetPolicyRuleServiceAccountOk

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleServiceAccountOk() (*string, bool)`

GetPolicyRuleServiceAccountOk returns a tuple with the PolicyRuleServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) SetPolicyRuleServiceAccount(v string)`

SetPolicyRuleServiceAccount sets PolicyRuleServiceAccount field to given value.

### HasPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) HasPolicyRuleServiceAccount() bool`

HasPolicyRuleServiceAccount returns a boolean if a field has been set.

### GetAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) GetAccessAddWorkflow() string`

GetAccessAddWorkflow returns the AccessAddWorkflow field if non-nil, zero value otherwise.

### GetAccessAddWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAccessAddWorkflowOk() (*string, bool)`

GetAccessAddWorkflowOk returns a tuple with the AccessAddWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) SetAccessAddWorkflow(v string)`

SetAccessAddWorkflow sets AccessAddWorkflow field to given value.

### HasAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) HasAccessAddWorkflow() bool`

HasAccessAddWorkflow returns a boolean if a field has been set.

### GetAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) GetAccessRemoveWorkflow() string`

GetAccessRemoveWorkflow returns the AccessRemoveWorkflow field if non-nil, zero value otherwise.

### GetAccessRemoveWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAccessRemoveWorkflowOk() (*string, bool)`

GetAccessRemoveWorkflowOk returns a tuple with the AccessRemoveWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) SetAccessRemoveWorkflow(v string)`

SetAccessRemoveWorkflow sets AccessRemoveWorkflow field to given value.

### HasAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) HasAccessRemoveWorkflow() bool`

HasAccessRemoveWorkflow returns a boolean if a field has been set.

### GetAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) GetAddServiceAccountWorkflow() string`

GetAddServiceAccountWorkflow returns the AddServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetAddServiceAccountWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAddServiceAccountWorkflowOk() (*string, bool)`

GetAddServiceAccountWorkflowOk returns a tuple with the AddServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) SetAddServiceAccountWorkflow(v string)`

SetAddServiceAccountWorkflow sets AddServiceAccountWorkflow field to given value.

### HasAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) HasAddServiceAccountWorkflow() bool`

HasAddServiceAccountWorkflow returns a boolean if a field has been set.

### GetRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) GetRemoveServiceAccountWorkflow() string`

GetRemoveServiceAccountWorkflow returns the RemoveServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetRemoveServiceAccountWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetRemoveServiceAccountWorkflowOk() (*string, bool)`

GetRemoveServiceAccountWorkflowOk returns a tuple with the RemoveServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) SetRemoveServiceAccountWorkflow(v string)`

SetRemoveServiceAccountWorkflow sets RemoveServiceAccountWorkflow field to given value.

### HasRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) HasRemoveServiceAccountWorkflow() bool`

HasRemoveServiceAccountWorkflow returns a boolean if a field has been set.

### GetConnectionparameters

`func (o *UpdateSecuritySystemRequest) GetConnectionparameters() string`

GetConnectionparameters returns the Connectionparameters field if non-nil, zero value otherwise.

### GetConnectionparametersOk

`func (o *UpdateSecuritySystemRequest) GetConnectionparametersOk() (*string, bool)`

GetConnectionparametersOk returns a tuple with the Connectionparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionparameters

`func (o *UpdateSecuritySystemRequest) SetConnectionparameters(v string)`

SetConnectionparameters sets Connectionparameters field to given value.

### HasConnectionparameters

`func (o *UpdateSecuritySystemRequest) HasConnectionparameters() bool`

HasConnectionparameters returns a boolean if a field has been set.

### GetAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) GetAutomatedProvisioning() string`

GetAutomatedProvisioning returns the AutomatedProvisioning field if non-nil, zero value otherwise.

### GetAutomatedProvisioningOk

`func (o *UpdateSecuritySystemRequest) GetAutomatedProvisioningOk() (*string, bool)`

GetAutomatedProvisioningOk returns a tuple with the AutomatedProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) SetAutomatedProvisioning(v string)`

SetAutomatedProvisioning sets AutomatedProvisioning field to given value.

### HasAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) HasAutomatedProvisioning() bool`

HasAutomatedProvisioning returns a boolean if a field has been set.

### GetConnectionname

`func (o *UpdateSecuritySystemRequest) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *UpdateSecuritySystemRequest) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *UpdateSecuritySystemRequest) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *UpdateSecuritySystemRequest) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetProvisioningConnection

`func (o *UpdateSecuritySystemRequest) GetProvisioningConnection() string`

GetProvisioningConnection returns the ProvisioningConnection field if non-nil, zero value otherwise.

### GetProvisioningConnectionOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningConnectionOk() (*string, bool)`

GetProvisioningConnectionOk returns a tuple with the ProvisioningConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConnection

`func (o *UpdateSecuritySystemRequest) SetProvisioningConnection(v string)`

SetProvisioningConnection sets ProvisioningConnection field to given value.

### HasProvisioningConnection

`func (o *UpdateSecuritySystemRequest) HasProvisioningConnection() bool`

HasProvisioningConnection returns a boolean if a field has been set.

### GetServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) GetServiceDeskConnection() string`

GetServiceDeskConnection returns the ServiceDeskConnection field if non-nil, zero value otherwise.

### GetServiceDeskConnectionOk

`func (o *UpdateSecuritySystemRequest) GetServiceDeskConnectionOk() (*string, bool)`

GetServiceDeskConnectionOk returns a tuple with the ServiceDeskConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) SetServiceDeskConnection(v string)`

SetServiceDeskConnection sets ServiceDeskConnection field to given value.

### HasServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) HasServiceDeskConnection() bool`

HasServiceDeskConnection returns a boolean if a field has been set.

### GetProvisioningcomments

`func (o *UpdateSecuritySystemRequest) GetProvisioningcomments() string`

GetProvisioningcomments returns the Provisioningcomments field if non-nil, zero value otherwise.

### GetProvisioningcommentsOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningcommentsOk() (*string, bool)`

GetProvisioningcommentsOk returns a tuple with the Provisioningcomments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningcomments

`func (o *UpdateSecuritySystemRequest) SetProvisioningcomments(v string)`

SetProvisioningcomments sets Provisioningcomments field to given value.

### HasProvisioningcomments

`func (o *UpdateSecuritySystemRequest) HasProvisioningcomments() bool`

HasProvisioningcomments returns a boolean if a field has been set.

### GetAction

`func (o *UpdateSecuritySystemRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateSecuritySystemRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateSecuritySystemRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateSecuritySystemRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDefaultSystem

`func (o *UpdateSecuritySystemRequest) GetDefaultSystem() string`

GetDefaultSystem returns the DefaultSystem field if non-nil, zero value otherwise.

### GetDefaultSystemOk

`func (o *UpdateSecuritySystemRequest) GetDefaultSystemOk() (*string, bool)`

GetDefaultSystemOk returns a tuple with the DefaultSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSystem

`func (o *UpdateSecuritySystemRequest) SetDefaultSystem(v string)`

SetDefaultSystem sets DefaultSystem field to given value.

### HasDefaultSystem

`func (o *UpdateSecuritySystemRequest) HasDefaultSystem() bool`

HasDefaultSystem returns a boolean if a field has been set.

### GetReconApplication

`func (o *UpdateSecuritySystemRequest) GetReconApplication() string`

GetReconApplication returns the ReconApplication field if non-nil, zero value otherwise.

### GetReconApplicationOk

`func (o *UpdateSecuritySystemRequest) GetReconApplicationOk() (*string, bool)`

GetReconApplicationOk returns a tuple with the ReconApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconApplication

`func (o *UpdateSecuritySystemRequest) SetReconApplication(v string)`

SetReconApplication sets ReconApplication field to given value.

### HasReconApplication

`func (o *UpdateSecuritySystemRequest) HasReconApplication() bool`

HasReconApplication returns a boolean if a field has been set.

### GetProvisioningTries

`func (o *UpdateSecuritySystemRequest) GetProvisioningTries() string`

GetProvisioningTries returns the ProvisioningTries field if non-nil, zero value otherwise.

### GetProvisioningTriesOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningTriesOk() (*string, bool)`

GetProvisioningTriesOk returns a tuple with the ProvisioningTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTries

`func (o *UpdateSecuritySystemRequest) SetProvisioningTries(v string)`

SetProvisioningTries sets ProvisioningTries field to given value.

### HasProvisioningTries

`func (o *UpdateSecuritySystemRequest) HasProvisioningTries() bool`

HasProvisioningTries returns a boolean if a field has been set.

### GetInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) GetInherentSODReportFields() string`

GetInherentSODReportFields returns the InherentSODReportFields field if non-nil, zero value otherwise.

### GetInherentSODReportFieldsOk

`func (o *UpdateSecuritySystemRequest) GetInherentSODReportFieldsOk() (*string, bool)`

GetInherentSODReportFieldsOk returns a tuple with the InherentSODReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) SetInherentSODReportFields(v string)`

SetInherentSODReportFields sets InherentSODReportFields field to given value.

### HasInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) HasInherentSODReportFields() bool`

HasInherentSODReportFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


