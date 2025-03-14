/*
Saviynt Endpoints API

Saviynt Endpoints API

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpoints

import (
	"encoding/json"
)

// checks if the UpdateEndpointResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEndpointResponse{}

// UpdateEndpointResponse struct for UpdateEndpointResponse
type UpdateEndpointResponse struct {
	Msg *string `json:"msg,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
}

// NewUpdateEndpointResponse instantiates a new UpdateEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEndpointResponse() *UpdateEndpointResponse {
	this := UpdateEndpointResponse{}
	return &this
}

// NewUpdateEndpointResponseWithDefaults instantiates a new UpdateEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEndpointResponseWithDefaults() *UpdateEndpointResponse {
	this := UpdateEndpointResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *UpdateEndpointResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpointResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *UpdateEndpointResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *UpdateEndpointResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *UpdateEndpointResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpointResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *UpdateEndpointResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *UpdateEndpointResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o UpdateEndpointResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEndpointResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableUpdateEndpointResponse struct {
	value *UpdateEndpointResponse
	isSet bool
}

func (v NullableUpdateEndpointResponse) Get() *UpdateEndpointResponse {
	return v.value
}

func (v *NullableUpdateEndpointResponse) Set(val *UpdateEndpointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEndpointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEndpointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEndpointResponse(val *UpdateEndpointResponse) *NullableUpdateEndpointResponse {
	return &NullableUpdateEndpointResponse{value: val, isSet: true}
}

func (v NullableUpdateEndpointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEndpointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


