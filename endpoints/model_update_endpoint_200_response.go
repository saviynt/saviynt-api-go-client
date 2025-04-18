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

// checks if the UpdateEndpoint200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEndpoint200Response{}

// UpdateEndpoint200Response struct for UpdateEndpoint200Response
type UpdateEndpoint200Response struct {
	// A message indicating the outcome of the operation.
	Msg *string `json:"msg,omitempty"`
	// An error code where '0' signifies success and '1' signifies an unsuccessful operation.
	ErrorCode *string `json:"errorCode,omitempty"`
}

// NewUpdateEndpoint200Response instantiates a new UpdateEndpoint200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEndpoint200Response() *UpdateEndpoint200Response {
	this := UpdateEndpoint200Response{}
	return &this
}

// NewUpdateEndpoint200ResponseWithDefaults instantiates a new UpdateEndpoint200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEndpoint200ResponseWithDefaults() *UpdateEndpoint200Response {
	this := UpdateEndpoint200Response{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *UpdateEndpoint200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpoint200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *UpdateEndpoint200Response) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *UpdateEndpoint200Response) SetMsg(v string) {
	o.Msg = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *UpdateEndpoint200Response) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEndpoint200Response) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *UpdateEndpoint200Response) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *UpdateEndpoint200Response) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o UpdateEndpoint200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEndpoint200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableUpdateEndpoint200Response struct {
	value *UpdateEndpoint200Response
	isSet bool
}

func (v NullableUpdateEndpoint200Response) Get() *UpdateEndpoint200Response {
	return v.value
}

func (v *NullableUpdateEndpoint200Response) Set(val *UpdateEndpoint200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEndpoint200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEndpoint200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEndpoint200Response(val *UpdateEndpoint200Response) *NullableUpdateEndpoint200Response {
	return &NullableUpdateEndpoint200Response{value: val, isSet: true}
}

func (v NullableUpdateEndpoint200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEndpoint200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
