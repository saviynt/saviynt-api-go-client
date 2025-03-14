/*
Saviynt Connections API

Connections

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connections

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateOrUpdateConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateConnectionResponse{}

// CreateOrUpdateConnectionResponse struct for CreateOrUpdateConnectionResponse
type CreateOrUpdateConnectionResponse struct {
	Msg string `json:"msg"`
	ErrorCode string `json:"errorCode"`
	ConnectionKey *int32 `json:"connectionKey,omitempty"`
}

type _CreateOrUpdateConnectionResponse CreateOrUpdateConnectionResponse

// NewCreateOrUpdateConnectionResponse instantiates a new CreateOrUpdateConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateConnectionResponse(msg string, errorCode string) *CreateOrUpdateConnectionResponse {
	this := CreateOrUpdateConnectionResponse{}
	this.Msg = msg
	this.ErrorCode = errorCode
	return &this
}

// NewCreateOrUpdateConnectionResponseWithDefaults instantiates a new CreateOrUpdateConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateConnectionResponseWithDefaults() *CreateOrUpdateConnectionResponse {
	this := CreateOrUpdateConnectionResponse{}
	return &this
}

// GetMsg returns the Msg field value
func (o *CreateOrUpdateConnectionResponse) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectionResponse) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *CreateOrUpdateConnectionResponse) SetMsg(v string) {
	o.Msg = v
}

// GetErrorCode returns the ErrorCode field value
func (o *CreateOrUpdateConnectionResponse) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectionResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CreateOrUpdateConnectionResponse) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetConnectionKey returns the ConnectionKey field value if set, zero value otherwise.
func (o *CreateOrUpdateConnectionResponse) GetConnectionKey() int32 {
	if o == nil || IsNil(o.ConnectionKey) {
		var ret int32
		return ret
	}
	return *o.ConnectionKey
}

// GetConnectionKeyOk returns a tuple with the ConnectionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectionResponse) GetConnectionKeyOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectionKey) {
		return nil, false
	}
	return o.ConnectionKey, true
}

// HasConnectionKey returns a boolean if a field has been set.
func (o *CreateOrUpdateConnectionResponse) HasConnectionKey() bool {
	if o != nil && !IsNil(o.ConnectionKey) {
		return true
	}

	return false
}

// SetConnectionKey gets a reference to the given int32 and assigns it to the ConnectionKey field.
func (o *CreateOrUpdateConnectionResponse) SetConnectionKey(v int32) {
	o.ConnectionKey = &v
}

func (o CreateOrUpdateConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msg"] = o.Msg
	toSerialize["errorCode"] = o.ErrorCode
	if !IsNil(o.ConnectionKey) {
		toSerialize["connectionKey"] = o.ConnectionKey
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"msg",
		"errorCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateOrUpdateConnectionResponse := _CreateOrUpdateConnectionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateConnectionResponse)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateConnectionResponse(varCreateOrUpdateConnectionResponse)

	return err
}

type NullableCreateOrUpdateConnectionResponse struct {
	value *CreateOrUpdateConnectionResponse
	isSet bool
}

func (v NullableCreateOrUpdateConnectionResponse) Get() *CreateOrUpdateConnectionResponse {
	return v.value
}

func (v *NullableCreateOrUpdateConnectionResponse) Set(val *CreateOrUpdateConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateConnectionResponse(val *CreateOrUpdateConnectionResponse) *NullableCreateOrUpdateConnectionResponse {
	return &NullableCreateOrUpdateConnectionResponse{value: val, isSet: true}
}

func (v NullableCreateOrUpdateConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


