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

// checks if the CreateEndpointResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEndpointResponse{}

// CreateEndpointResponse struct for CreateEndpointResponse
type CreateEndpointResponse struct {
	// Response message
	Msg *string `json:"msg,omitempty"`
	// Error code (0 indicates success)
	ErrorCode *string `json:"errorCode,omitempty"`
}

// NewCreateEndpointResponse instantiates a new CreateEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEndpointResponse() *CreateEndpointResponse {
	this := CreateEndpointResponse{}
	return &this
}

// NewCreateEndpointResponseWithDefaults instantiates a new CreateEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEndpointResponseWithDefaults() *CreateEndpointResponse {
	this := CreateEndpointResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateEndpointResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateEndpointResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateEndpointResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreateEndpointResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreateEndpointResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreateEndpointResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o CreateEndpointResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEndpointResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableCreateEndpointResponse struct {
	value *CreateEndpointResponse
	isSet bool
}

func (v NullableCreateEndpointResponse) Get() *CreateEndpointResponse {
	return v.value
}

func (v *NullableCreateEndpointResponse) Set(val *CreateEndpointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndpointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndpointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndpointResponse(val *CreateEndpointResponse) *NullableCreateEndpointResponse {
	return &NullableCreateEndpointResponse{value: val, isSet: true}
}

func (v NullableCreateEndpointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndpointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


