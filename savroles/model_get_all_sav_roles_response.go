/*
Saviynt SAV Roles API

Saviynt SAV Roles API

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package savroles

import (
	"encoding/json"
)

// checks if the GetAllSAVRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllSAVRolesResponse{}

// GetAllSAVRolesResponse
type GetAllSAVRolesResponse struct {
	Savroles []SAVRole `json:"savroles,omitempty"`
}

// NewGetAllSAVRolesResponse instantiates a new GetAllSAVRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllSAVRolesResponse() *GetAllSAVRolesResponse {
	this := GetAllSAVRolesResponse{}
	return &this
}

// NewGetAllSAVRolesResponseWithDefaults instantiates a new GetAllSAVRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllSAVRolesResponseWithDefaults() *GetAllSAVRolesResponse {
	this := GetAllSAVRolesResponse{}
	return &this
}

// GetSavroles returns the Savroles field value if set, zero value otherwise.
func (o *GetAllSAVRolesResponse) GetSavroles() []SAVRole {
	if o == nil || IsNil(o.Savroles) {
		var ret []SAVRole
		return ret
	}
	return o.Savroles
}

// GetSavrolesOk returns a tuple with the Savroles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSAVRolesResponse) GetSavrolesOk() ([]SAVRole, bool) {
	if o == nil || IsNil(o.Savroles) {
		return nil, false
	}
	return o.Savroles, true
}

// HasSavroles returns a boolean if a field has been set.
func (o *GetAllSAVRolesResponse) HasSavroles() bool {
	if o != nil && !IsNil(o.Savroles) {
		return true
	}

	return false
}

// SetSavroles gets a reference to the given []SAVRole and assigns it to the Savroles field.
func (o *GetAllSAVRolesResponse) SetSavroles(v []SAVRole) {
	o.Savroles = v
}

func (o GetAllSAVRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllSAVRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Savroles) {
		toSerialize["savroles"] = o.Savroles
	}
	return toSerialize, nil
}

type NullableGetAllSAVRolesResponse struct {
	value *GetAllSAVRolesResponse
	isSet bool
}

func (v NullableGetAllSAVRolesResponse) Get() *GetAllSAVRolesResponse {
	return v.value
}

func (v *NullableGetAllSAVRolesResponse) Set(val *GetAllSAVRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllSAVRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllSAVRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllSAVRolesResponse(val *GetAllSAVRolesResponse) *NullableGetAllSAVRolesResponse {
	return &NullableGetAllSAVRolesResponse{value: val, isSet: true}
}

func (v NullableGetAllSAVRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllSAVRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
