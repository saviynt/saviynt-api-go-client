/*
Account Management API

API for managing accounts in Saviynt/SSM. - **Create Endpoint**: Creates a new endpoint. - **Update Endpoint**: Updates an existing endpoint based on its name and roletype. - **Get Endpoint List**: Returns a list of endpoints based on search criteria.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpoints

import (
	"encoding/json"
)

// checks if the UpdateEndpointRequestEmailTemplateInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEndpointRequestEmailTemplateInner{}

// UpdateEndpointRequestEmailTemplateInner struct for UpdateEndpointRequestEmailTemplateInner
type UpdateEndpointRequestEmailTemplateInner struct {
	//
	EmailTemplateType *string `json:"emailTemplateType,omitempty"`
	//
	TaskType *string `json:"taskType,omitempty"`
	//
	EmailTemplate *string `json:"emailTemplate,omitempty"`
}

// NewUpdateEndpointRequestEmailTemplateInner instantiates a new UpdateEndpointRequestEmailTemplateInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEndpointRequestEmailTemplateInner() *UpdateEndpointRequestEmailTemplateInner {
	this := UpdateEndpointRequestEmailTemplateInner{}
	return &this
}

// NewUpdateEndpointRequestEmailTemplateInnerWithDefaults instantiates a new UpdateEndpointRequestEmailTemplateInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEndpointRequestEmailTemplateInnerWithDefaults() *UpdateEndpointRequestEmailTemplateInner {
	this := UpdateEndpointRequestEmailTemplateInner{}
	return &this
}

// GetEmailTemplateType returns the EmailTemplateType field value if set, zero value otherwise.
func (o *UpdateEndpointRequestEmailTemplateInner) GetEmailTemplateType() string {
	if o == nil || IsNil(o.EmailTemplateType) {
		var ret string
		return ret
	}
	return *o.EmailTemplateType
}

// GetEmailTemplateTypeOk returns a tuple with the EmailTemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) GetEmailTemplateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateType) {
		return nil, false
	}
	return o.EmailTemplateType, true
}

// HasEmailTemplateType returns a boolean if a field has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) HasEmailTemplateType() bool {
	if o != nil && !IsNil(o.EmailTemplateType) {
		return true
	}

	return false
}

// SetEmailTemplateType gets a reference to the given string and assigns it to the EmailTemplateType field.
func (o *UpdateEndpointRequestEmailTemplateInner) SetEmailTemplateType(v string) {
	o.EmailTemplateType = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *UpdateEndpointRequestEmailTemplateInner) GetTaskType() string {
	if o == nil || IsNil(o.TaskType) {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) GetTaskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) HasTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *UpdateEndpointRequestEmailTemplateInner) SetTaskType(v string) {
	o.TaskType = &v
}

// GetEmailTemplate returns the EmailTemplate field value if set, zero value otherwise.
func (o *UpdateEndpointRequestEmailTemplateInner) GetEmailTemplate() string {
	if o == nil || IsNil(o.EmailTemplate) {
		var ret string
		return ret
	}
	return *o.EmailTemplate
}

// GetEmailTemplateOk returns a tuple with the EmailTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) GetEmailTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplate) {
		return nil, false
	}
	return o.EmailTemplate, true
}

// HasEmailTemplate returns a boolean if a field has been set.
func (o *UpdateEndpointRequestEmailTemplateInner) HasEmailTemplate() bool {
	if o != nil && !IsNil(o.EmailTemplate) {
		return true
	}

	return false
}

// SetEmailTemplate gets a reference to the given string and assigns it to the EmailTemplate field.
func (o *UpdateEndpointRequestEmailTemplateInner) SetEmailTemplate(v string) {
	o.EmailTemplate = &v
}

func (o UpdateEndpointRequestEmailTemplateInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEndpointRequestEmailTemplateInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailTemplateType) {
		toSerialize["emailTemplateType"] = o.EmailTemplateType
	}
	if !IsNil(o.TaskType) {
		toSerialize["taskType"] = o.TaskType
	}
	if !IsNil(o.EmailTemplate) {
		toSerialize["emailTemplate"] = o.EmailTemplate
	}
	return toSerialize, nil
}

type NullableUpdateEndpointRequestEmailTemplateInner struct {
	value *UpdateEndpointRequestEmailTemplateInner
	isSet bool
}

func (v NullableUpdateEndpointRequestEmailTemplateInner) Get() *UpdateEndpointRequestEmailTemplateInner {
	return v.value
}

func (v *NullableUpdateEndpointRequestEmailTemplateInner) Set(val *UpdateEndpointRequestEmailTemplateInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEndpointRequestEmailTemplateInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEndpointRequestEmailTemplateInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEndpointRequestEmailTemplateInner(val *UpdateEndpointRequestEmailTemplateInner) *NullableUpdateEndpointRequestEmailTemplateInner {
	return &NullableUpdateEndpointRequestEmailTemplateInner{value: val, isSet: true}
}

func (v NullableUpdateEndpointRequestEmailTemplateInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEndpointRequestEmailTemplateInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
