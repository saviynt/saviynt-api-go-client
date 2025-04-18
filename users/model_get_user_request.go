/*
Saviynt Users API

Saviynt Users API

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package users

import (
	"encoding/json"
)

// checks if the GetUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserRequest{}

// GetUserRequest Note - By default, if no responsefields param is passed, required attrs that will always return are username, email, statuskey, firstname, lastname, employeeid along with other attributes with nonblank values only.
type GetUserRequest struct {
	// Specify the username for which you want to get the user attribute details
	Username *string `json:"username,omitempty"`
	// User attributes which you want to see in the response(for encrypted values, mention ecp<1-5>, and for hashed values, mention hcp<1-5>)
	Responsefields []interface{} `json:"responsefields,omitempty"`
	Max            *int32        `json:"max,omitempty"`
	Offset         *int32        `json:"offset,omitempty"`
	Order          *string       `json:"order,omitempty"`
	// username
	Manager *string `json:"manager,omitempty"`
	// userkey OR `secondaryManager` - username
	Secondarymanager *string `json:"secondarymanager,omitempty"`
	// \"0\"/\"1\" to display encrypted security answers for the user
	Showsecurityanswers *string                `json:"showsecurityanswers,omitempty"`
	Filtercriteria      map[string]interface{} `json:"filtercriteria,omitempty"`
	SearchCriteria      *string                `json:"searchCriteria,omitempty"`
	Advsearchcriteria   *string                `json:"advsearchcriteria,omitempty"`
	UserQuery           *string                `json:"userQuery,omitempty"`
}

// NewGetUserRequest instantiates a new GetUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserRequest() *GetUserRequest {
	this := GetUserRequest{}
	return &this
}

// NewGetUserRequestWithDefaults instantiates a new GetUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserRequestWithDefaults() *GetUserRequest {
	this := GetUserRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetUserRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetUserRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetUserRequest) SetUsername(v string) {
	o.Username = &v
}

// GetResponsefields returns the Responsefields field value if set, zero value otherwise.
func (o *GetUserRequest) GetResponsefields() []interface{} {
	if o == nil || IsNil(o.Responsefields) {
		var ret []interface{}
		return ret
	}
	return o.Responsefields
}

// GetResponsefieldsOk returns a tuple with the Responsefields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetResponsefieldsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Responsefields) {
		return nil, false
	}
	return o.Responsefields, true
}

// HasResponsefields returns a boolean if a field has been set.
func (o *GetUserRequest) HasResponsefields() bool {
	if o != nil && !IsNil(o.Responsefields) {
		return true
	}

	return false
}

// SetResponsefields gets a reference to the given []interface{} and assigns it to the Responsefields field.
func (o *GetUserRequest) SetResponsefields(v []interface{}) {
	o.Responsefields = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *GetUserRequest) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *GetUserRequest) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *GetUserRequest) SetMax(v int32) {
	o.Max = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetUserRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetUserRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetUserRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GetUserRequest) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GetUserRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *GetUserRequest) SetOrder(v string) {
	o.Order = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *GetUserRequest) GetManager() string {
	if o == nil || IsNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetManagerOk() (*string, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *GetUserRequest) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *GetUserRequest) SetManager(v string) {
	o.Manager = &v
}

// GetSecondarymanager returns the Secondarymanager field value if set, zero value otherwise.
func (o *GetUserRequest) GetSecondarymanager() string {
	if o == nil || IsNil(o.Secondarymanager) {
		var ret string
		return ret
	}
	return *o.Secondarymanager
}

// GetSecondarymanagerOk returns a tuple with the Secondarymanager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetSecondarymanagerOk() (*string, bool) {
	if o == nil || IsNil(o.Secondarymanager) {
		return nil, false
	}
	return o.Secondarymanager, true
}

// HasSecondarymanager returns a boolean if a field has been set.
func (o *GetUserRequest) HasSecondarymanager() bool {
	if o != nil && !IsNil(o.Secondarymanager) {
		return true
	}

	return false
}

// SetSecondarymanager gets a reference to the given string and assigns it to the Secondarymanager field.
func (o *GetUserRequest) SetSecondarymanager(v string) {
	o.Secondarymanager = &v
}

// GetShowsecurityanswers returns the Showsecurityanswers field value if set, zero value otherwise.
func (o *GetUserRequest) GetShowsecurityanswers() string {
	if o == nil || IsNil(o.Showsecurityanswers) {
		var ret string
		return ret
	}
	return *o.Showsecurityanswers
}

// GetShowsecurityanswersOk returns a tuple with the Showsecurityanswers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetShowsecurityanswersOk() (*string, bool) {
	if o == nil || IsNil(o.Showsecurityanswers) {
		return nil, false
	}
	return o.Showsecurityanswers, true
}

// HasShowsecurityanswers returns a boolean if a field has been set.
func (o *GetUserRequest) HasShowsecurityanswers() bool {
	if o != nil && !IsNil(o.Showsecurityanswers) {
		return true
	}

	return false
}

// SetShowsecurityanswers gets a reference to the given string and assigns it to the Showsecurityanswers field.
func (o *GetUserRequest) SetShowsecurityanswers(v string) {
	o.Showsecurityanswers = &v
}

// GetFiltercriteria returns the Filtercriteria field value if set, zero value otherwise.
func (o *GetUserRequest) GetFiltercriteria() map[string]interface{} {
	if o == nil || IsNil(o.Filtercriteria) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filtercriteria
}

// GetFiltercriteriaOk returns a tuple with the Filtercriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetFiltercriteriaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filtercriteria) {
		return map[string]interface{}{}, false
	}
	return o.Filtercriteria, true
}

// HasFiltercriteria returns a boolean if a field has been set.
func (o *GetUserRequest) HasFiltercriteria() bool {
	if o != nil && !IsNil(o.Filtercriteria) {
		return true
	}

	return false
}

// SetFiltercriteria gets a reference to the given map[string]interface{} and assigns it to the Filtercriteria field.
func (o *GetUserRequest) SetFiltercriteria(v map[string]interface{}) {
	o.Filtercriteria = v
}

// GetSearchCriteria returns the SearchCriteria field value if set, zero value otherwise.
func (o *GetUserRequest) GetSearchCriteria() string {
	if o == nil || IsNil(o.SearchCriteria) {
		var ret string
		return ret
	}
	return *o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetSearchCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SearchCriteria) {
		return nil, false
	}
	return o.SearchCriteria, true
}

// HasSearchCriteria returns a boolean if a field has been set.
func (o *GetUserRequest) HasSearchCriteria() bool {
	if o != nil && !IsNil(o.SearchCriteria) {
		return true
	}

	return false
}

// SetSearchCriteria gets a reference to the given string and assigns it to the SearchCriteria field.
func (o *GetUserRequest) SetSearchCriteria(v string) {
	o.SearchCriteria = &v
}

// GetAdvsearchcriteria returns the Advsearchcriteria field value if set, zero value otherwise.
func (o *GetUserRequest) GetAdvsearchcriteria() string {
	if o == nil || IsNil(o.Advsearchcriteria) {
		var ret string
		return ret
	}
	return *o.Advsearchcriteria
}

// GetAdvsearchcriteriaOk returns a tuple with the Advsearchcriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetAdvsearchcriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.Advsearchcriteria) {
		return nil, false
	}
	return o.Advsearchcriteria, true
}

// HasAdvsearchcriteria returns a boolean if a field has been set.
func (o *GetUserRequest) HasAdvsearchcriteria() bool {
	if o != nil && !IsNil(o.Advsearchcriteria) {
		return true
	}

	return false
}

// SetAdvsearchcriteria gets a reference to the given string and assigns it to the Advsearchcriteria field.
func (o *GetUserRequest) SetAdvsearchcriteria(v string) {
	o.Advsearchcriteria = &v
}

// GetUserQuery returns the UserQuery field value if set, zero value otherwise.
func (o *GetUserRequest) GetUserQuery() string {
	if o == nil || IsNil(o.UserQuery) {
		var ret string
		return ret
	}
	return *o.UserQuery
}

// GetUserQueryOk returns a tuple with the UserQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetUserQueryOk() (*string, bool) {
	if o == nil || IsNil(o.UserQuery) {
		return nil, false
	}
	return o.UserQuery, true
}

// HasUserQuery returns a boolean if a field has been set.
func (o *GetUserRequest) HasUserQuery() bool {
	if o != nil && !IsNil(o.UserQuery) {
		return true
	}

	return false
}

// SetUserQuery gets a reference to the given string and assigns it to the UserQuery field.
func (o *GetUserRequest) SetUserQuery(v string) {
	o.UserQuery = &v
}

func (o GetUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Responsefields) {
		toSerialize["responsefields"] = o.Responsefields
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.Secondarymanager) {
		toSerialize["secondarymanager"] = o.Secondarymanager
	}
	if !IsNil(o.Showsecurityanswers) {
		toSerialize["showsecurityanswers"] = o.Showsecurityanswers
	}
	if !IsNil(o.Filtercriteria) {
		toSerialize["filtercriteria"] = o.Filtercriteria
	}
	if !IsNil(o.SearchCriteria) {
		toSerialize["searchCriteria"] = o.SearchCriteria
	}
	if !IsNil(o.Advsearchcriteria) {
		toSerialize["advsearchcriteria"] = o.Advsearchcriteria
	}
	if !IsNil(o.UserQuery) {
		toSerialize["userQuery"] = o.UserQuery
	}
	return toSerialize, nil
}

type NullableGetUserRequest struct {
	value *GetUserRequest
	isSet bool
}

func (v NullableGetUserRequest) Get() *GetUserRequest {
	return v.value
}

func (v *NullableGetUserRequest) Set(val *GetUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRequest(val *GetUserRequest) *NullableGetUserRequest {
	return &NullableGetUserRequest{value: val, isSet: true}
}

func (v NullableGetUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
