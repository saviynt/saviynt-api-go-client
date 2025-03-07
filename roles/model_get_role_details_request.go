package roles

import (
	// "context"
	"encoding/json"
	"fmt"
	"reflect"
)

var _ MappedNullable = &GetRoleRequest{}

type GetRoleRequest struct {
	UserName        *string          `json:"username,omitempty"`
	RequestedObject *string          `json:"requestedObject,omitempty"`
	RoleType        *string          `json:"roletype,omitempty"`
	Requestable     *string          `json:"requestable,omitempty"`
	Status          *string          `json:"status,omitempty"`
	RoleName        *string          `json:"role_name,omitempty"`
	Description     *string          `json:"description,omitempty"`
	DisplayName     *string          `json:"displayname,omitempty"`
	Glossary        *string          `json:"glossary,omitempty"`
	MiningInstance  *string          `json:"mininginstance,omitempty"`
	Risk            *string          `json:"risk,omitempty"`
	UpdateUser      *string          `json:"updateuser,omitempty"`
	SystemID        *string          `json:"systemid,omitempty"`
	SoxCritical     *string          `json:"soxcritical,omitempty"`
	SysCritical     *string          `json:"syscritical,omitempty"`
	Level           *string          `json:"level,omitempty"`
	Privileged      *string          `json:"privileged,omitempty"`
	Confidentiality *string          `json:"confidentiality,omitempty"`

	// Custom properties 1 to 60 (individually defined)
	CustomProperty1  *string `json:"customproperty1,omitempty"`
	CustomProperty2  *string `json:"customproperty2,omitempty"`
	CustomProperty3  *string `json:"customproperty3,omitempty"`
	CustomProperty4  *string `json:"customproperty4,omitempty"`
	CustomProperty5  *string `json:"customproperty5,omitempty"`
	CustomProperty6  *string `json:"customproperty6,omitempty"`
	CustomProperty7  *string `json:"customproperty7,omitempty"`
	CustomProperty8  *string `json:"customproperty8,omitempty"`
	CustomProperty9  *string `json:"customproperty9,omitempty"`
	CustomProperty10 *string `json:"customproperty10,omitempty"`
	CustomProperty11 *string `json:"customproperty11,omitempty"`
	CustomProperty12 *string `json:"customproperty12,omitempty"`
	CustomProperty13 *string `json:"customproperty13,omitempty"`
	CustomProperty14 *string `json:"customproperty14,omitempty"`
	CustomProperty15 *string `json:"customproperty15,omitempty"`
	CustomProperty16 *string `json:"customproperty16,omitempty"`
	CustomProperty17 *string `json:"customproperty17,omitempty"`
	CustomProperty18 *string `json:"customproperty18,omitempty"`
	CustomProperty19 *string `json:"customproperty19,omitempty"`
	CustomProperty20 *string `json:"customproperty20,omitempty"`
	CustomProperty21 *string `json:"customproperty21,omitempty"`
	CustomProperty22 *string `json:"customproperty22,omitempty"`
	CustomProperty23 *string `json:"customproperty23,omitempty"`
	CustomProperty24 *string `json:"customproperty24,omitempty"`
	CustomProperty25 *string `json:"customproperty25,omitempty"`
	CustomProperty26 *string `json:"customproperty26,omitempty"`
	CustomProperty27 *string `json:"customproperty27,omitempty"`
	CustomProperty28 *string `json:"customproperty28,omitempty"`
	CustomProperty29 *string `json:"customproperty29,omitempty"`
	CustomProperty30 *string `json:"customproperty30,omitempty"`
	CustomProperty31 *string `json:"customproperty31,omitempty"`
	CustomProperty32 *string `json:"customproperty32,omitempty"`
	CustomProperty33 *string `json:"customproperty33,omitempty"`
	CustomProperty34 *string `json:"customproperty34,omitempty"`
	CustomProperty35 *string `json:"customproperty35,omitempty"`
	CustomProperty36 *string `json:"customproperty36,omitempty"`
	CustomProperty37 *string `json:"customproperty37,omitempty"`
	CustomProperty38 *string `json:"customproperty38,omitempty"`
	CustomProperty39 *string `json:"customproperty39,omitempty"`
	CustomProperty40 *string `json:"customproperty40,omitempty"`
	CustomProperty41 *string `json:"customproperty41,omitempty"`
	CustomProperty42 *string `json:"customproperty42,omitempty"`
	CustomProperty43 *string `json:"customproperty43,omitempty"`
	CustomProperty44 *string `json:"customproperty44,omitempty"`
	CustomProperty45 *string `json:"customproperty45,omitempty"`
	CustomProperty46 *string `json:"customproperty46,omitempty"`
	CustomProperty47 *string `json:"customproperty47,omitempty"`
	CustomProperty48 *string `json:"customproperty48,omitempty"`
	CustomProperty49 *string `json:"customproperty49,omitempty"`
	CustomProperty50 *string `json:"customproperty50,omitempty"`
	CustomProperty51 *string `json:"customproperty51,omitempty"`
	CustomProperty52 *string `json:"customproperty52,omitempty"`
	CustomProperty53 *string `json:"customproperty53,omitempty"`
	CustomProperty54 *string `json:"customproperty54,omitempty"`
	CustomProperty55 *string `json:"customproperty55,omitempty"`
	CustomProperty56 *string `json:"customproperty56,omitempty"`
	CustomProperty57 *string `json:"customproperty57,omitempty"`
	CustomProperty58 *string `json:"customproperty58,omitempty"`
	CustomProperty59 *string `json:"customproperty59,omitempty"`
	CustomProperty60 *string `json:"customproperty60,omitempty"`
}

func NewGetRoleRequest() *GetRoleRequest {
	this := GetRoleRequest{}
	return &this
}

func NewGetRoleRequestWithDefaults() *GetRoleRequest {
	this := GetRoleRequest{}
	return &this
}

// UserName functions
func (o *GetRoleRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		return ""
	}
	return *o.UserName
}

func (o *GetRoleRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

func (o *GetRoleRequest) HasUserName() bool {
	return o != nil && !IsNil(o.UserName)
}

func (o *GetRoleRequest) SetUserName(v string) {
	o.UserName = &v
}

// RequestedObject functions
func (o *GetRoleRequest) GetRequestedObject() string {
	if o == nil || IsNil(o.RequestedObject) {
		return ""
	}
	return *o.RequestedObject
}

func (o *GetRoleRequest) GetRequestedObjectOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedObject) {
		return nil, false
	}
	return o.RequestedObject, true
}

func (o *GetRoleRequest) HasRequestedObject() bool {
	return o != nil && !IsNil(o.RequestedObject)
}

func (o *GetRoleRequest) SetRequestedObject(v string) {
	o.RequestedObject = &v
}

// RoleType functions
func (o *GetRoleRequest) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		return ""
	}
	return *o.RoleType
}

func (o *GetRoleRequest) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

func (o *GetRoleRequest) HasRoleType() bool {
	return o != nil && !IsNil(o.RoleType)
}

func (o *GetRoleRequest) SetRoleType(v string) {
	o.RoleType = &v
}

// Requestable functions
func (o *GetRoleRequest) GetRequestable() string {
	if o == nil || IsNil(o.Requestable) {
		return ""
	}
	return *o.Requestable
}

func (o *GetRoleRequest) GetRequestableOk() (*string, bool) {
	if o == nil || IsNil(o.Requestable) {
		return nil, false
	}
	return o.Requestable, true
}

func (o *GetRoleRequest) HasRequestable() bool {
	return o != nil && !IsNil(o.Requestable)
}

func (o *GetRoleRequest) SetRequestable(v string) {
	o.Requestable = &v
}

// Status functions
func (o *GetRoleRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		return ""
	}
	return *o.Status
}

func (o *GetRoleRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

func (o *GetRoleRequest) HasStatus() bool {
	return o != nil && !IsNil(o.Status)
}

func (o *GetRoleRequest) SetStatus(v string) {
	o.Status = &v
}

// RoleName functions
func (o *GetRoleRequest) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		return ""
	}
	return *o.RoleName
}

func (o *GetRoleRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

func (o *GetRoleRequest) HasRoleName() bool {
	return o != nil && !IsNil(o.RoleName)
}

func (o *GetRoleRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// Description functions
func (o *GetRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		return ""
	}
	return *o.Description
}

func (o *GetRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

func (o *GetRoleRequest) HasDescription() bool {
	return o != nil && !IsNil(o.Description)
}

func (o *GetRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// DisplayName functions
func (o *GetRoleRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		return ""
	}
	return *o.DisplayName
}

func (o *GetRoleRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

func (o *GetRoleRequest) HasDisplayName() bool {
	return o != nil && !IsNil(o.DisplayName)
}

func (o *GetRoleRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// Glossary functions
func (o *GetRoleRequest) GetGlossary() string {
	if o == nil || IsNil(o.Glossary) {
		return ""
	}
	return *o.Glossary
}

func (o *GetRoleRequest) GetGlossaryOk() (*string, bool) {
	if o == nil || IsNil(o.Glossary) {
		return nil, false
	}
	return o.Glossary, true
}

func (o *GetRoleRequest) HasGlossary() bool {
	return o != nil && !IsNil(o.Glossary)
}

func (o *GetRoleRequest) SetGlossary(v string) {
	o.Glossary = &v
}

// MiningInstance functions
func (o *GetRoleRequest) GetMiningInstance() string {
	if o == nil || IsNil(o.MiningInstance) {
		return ""
	}
	return *o.MiningInstance
}

func (o *GetRoleRequest) GetMiningInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.MiningInstance) {
		return nil, false
	}
	return o.MiningInstance, true
}

func (o *GetRoleRequest) HasMiningInstance() bool {
	return o != nil && !IsNil(o.MiningInstance)
}

func (o *GetRoleRequest) SetMiningInstance(v string) {
	o.MiningInstance = &v
}

func (o *GetRoleRequest) GetRisk() string {
	if o == nil || IsNil(o.Risk) {
		return ""
	}
	return *o.Risk
}

func (o *GetRoleRequest) GetRiskOk() (*string, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

func (o *GetRoleRequest) HasRisk() bool {
	return o != nil && !IsNil(o.Risk)
}

func (o *GetRoleRequest) SetRisk(v string) {
	o.Risk = &v
}

func (o *GetRoleRequest) GetUpdateUser() string {
	if o == nil || IsNil(o.UpdateUser) {
		return ""
	}
	return *o.UpdateUser
}

func (o *GetRoleRequest) GetUpdateUserOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateUser) {
		return nil, false
	}
	return o.UpdateUser, true
}

func (o *GetRoleRequest) HasUpdateUser() bool {
	return o != nil && !IsNil(o.UpdateUser)
}

func (o *GetRoleRequest) SetUpdateUser(v string) {
	o.UpdateUser = &v
}

func (o *GetRoleRequest) GetSystemID() string {
	if o == nil || IsNil(o.SystemID) {
		return ""
	}
	return *o.SystemID
}

func (o *GetRoleRequest) GetSystemIDOk() (*string, bool) {
	if o == nil || IsNil(o.SystemID) {
		return nil, false
	}
	return o.SystemID, true
}

func (o *GetRoleRequest) HasSystemID() bool {
	return o != nil && !IsNil(o.SystemID)
}

func (o *GetRoleRequest) SetSystemID(v string) {
	o.SystemID = &v
}

func (o *GetRoleRequest) GetSoxCritical() string {
	if o == nil || IsNil(o.SoxCritical) {
		return ""
	}
	return *o.SoxCritical
}

func (o *GetRoleRequest) GetSoxCriticalOk() (*string, bool) {
	if o == nil || IsNil(o.SoxCritical) {
		return nil, false
	}
	return o.SoxCritical, true
}

func (o *GetRoleRequest) HasSoxCritical() bool {
	return o != nil && !IsNil(o.SoxCritical)
}

func (o *GetRoleRequest) SetSoxCritical(v string) {
	o.SoxCritical = &v
}

func (o *GetRoleRequest) GetSysCritical() string {
	if o == nil || IsNil(o.SysCritical) {
		return ""
	}
	return *o.SysCritical
}

func (o *GetRoleRequest) GetSysCriticalOk() (*string, bool) {
	if o == nil || IsNil(o.SysCritical) {
		return nil, false
	}
	return o.SysCritical, true
}

func (o *GetRoleRequest) HasSysCritical() bool {
	return o != nil && !IsNil(o.SysCritical)
}

func (o *GetRoleRequest) SetSysCritical(v string) {
	o.SysCritical = &v
}


func (o *GetRoleRequest) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		return ""
	}
	return *o.Level
}

func (o *GetRoleRequest) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

func (o *GetRoleRequest) HasLevel() bool {
	return o != nil && !IsNil(o.Level)
}

func (o *GetRoleRequest) SetLevel(v string) {
	o.Level = &v
}

func (o *GetRoleRequest) GetPrivileged() string {
	if o == nil || IsNil(o.Privileged) {
		return ""
	}
	return *o.Privileged
}

func (o *GetRoleRequest) GetPrivilegedOk() (*string, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

func (o *GetRoleRequest) HasPrivileged() bool {
	return o != nil && !IsNil(o.Privileged)
}

func (o *GetRoleRequest) SetPrivileged(v string) {
	o.Privileged = &v
}

func (o *GetRoleRequest) GetConfidentiality() string {
	if o == nil || IsNil(o.Confidentiality) {
		return ""
	}
	return *o.Confidentiality
}

func (o *GetRoleRequest) GetConfidentialityOk() (*string, bool) {
	if o == nil || IsNil(o.Confidentiality) {
		return nil, false
	}
	return o.Confidentiality, true
}

func (o *GetRoleRequest) HasConfidentiality() bool {
	return o != nil && !IsNil(o.Confidentiality)
}

func (o *GetRoleRequest) SetConfidentiality(v string) {
	o.Confidentiality = &v
}

func (o *GetRoleRequest) GetCustomProperty(index int) string {
	if o == nil {
		return ""
	}

	fieldName := fmt.Sprintf("CustomProperty%d", index)
	fieldValue := reflect.ValueOf(o).Elem().FieldByName(fieldName)

	if !fieldValue.IsValid() || fieldValue.IsNil() {
		return ""
	}

	return fieldValue.Elem().String()
}

func (o *GetRoleRequest) GetCustomPropertyOk(index int) (*string, bool) {
	if o == nil {
		return nil, false
	}

	fieldName := fmt.Sprintf("CustomProperty%d", index)
	fieldValue := reflect.ValueOf(o).Elem().FieldByName(fieldName)

	if !fieldValue.IsValid() || fieldValue.IsNil() {
		return nil, false
	}

	return fieldValue.Interface().(*string), true
}

// HasCustomProperty checks if a custom property is set.
func (o *GetRoleRequest) HasCustomProperty(index int) bool {
	if o == nil {
		return false
	}

	fieldName := fmt.Sprintf("CustomProperty%d", index)
	fieldValue := reflect.ValueOf(o).Elem().FieldByName(fieldName)

	return fieldValue.IsValid() && !fieldValue.IsNil()
}

// SetCustomProperty sets the value of a custom property by index.
func (o *GetRoleRequest) SetCustomProperty(index int, value string) {
	if o == nil {
		return
	}

	fieldName := fmt.Sprintf("CustomProperty%d", index)
	field := reflect.ValueOf(o).Elem().FieldByName(fieldName)

	if field.IsValid() && field.CanSet() {
		v := value
		field.Set(reflect.ValueOf(&v))
	}
}


func (o GetRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

// ToMap converts GetRoleRequest into a map for JSON serialization
func (o GetRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["username"] = o.UserName
	}
	if !IsNil(o.RequestedObject) {
		toSerialize["requestedObject"] = o.RequestedObject
	}
	if !IsNil(o.RoleType) {
		toSerialize["roletype"] = o.RoleType
	}
	if !IsNil(o.Requestable) {
		toSerialize["requestable"] = o.Requestable
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.RoleName) {
		toSerialize["role_name"] = o.RoleName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayname"] = o.DisplayName
	}
	if !IsNil(o.Glossary) {
		toSerialize["glossary"] = o.Glossary
	}
	if !IsNil(o.MiningInstance) {
		toSerialize["mininginstance"] = o.MiningInstance
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	if !IsNil(o.UpdateUser) {
		toSerialize["updateuser"] = o.UpdateUser
	}
	if !IsNil(o.SystemID) {
		toSerialize["systemid"] = o.SystemID
	}
	if !IsNil(o.SoxCritical) {
		toSerialize["soxcritical"] = o.SoxCritical
	}
	if !IsNil(o.SysCritical) {
		toSerialize["syscritical"] = o.SysCritical
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.Confidentiality) {
		toSerialize["confidentiality"] = o.Confidentiality
	}
	customProperties := []*string{
		o.CustomProperty1, o.CustomProperty2, o.CustomProperty3, o.CustomProperty4, o.CustomProperty5,
		o.CustomProperty6, o.CustomProperty7, o.CustomProperty8, o.CustomProperty9, o.CustomProperty10,
		o.CustomProperty11, o.CustomProperty12, o.CustomProperty13, o.CustomProperty14, o.CustomProperty15,
		o.CustomProperty16, o.CustomProperty17, o.CustomProperty18, o.CustomProperty19, o.CustomProperty20,
		o.CustomProperty21, o.CustomProperty22, o.CustomProperty23, o.CustomProperty24, o.CustomProperty25,
		o.CustomProperty26, o.CustomProperty27, o.CustomProperty28, o.CustomProperty29, o.CustomProperty30,
		o.CustomProperty31, o.CustomProperty32, o.CustomProperty33, o.CustomProperty34, o.CustomProperty35,
		o.CustomProperty36, o.CustomProperty37, o.CustomProperty38, o.CustomProperty39, o.CustomProperty40,
		o.CustomProperty41, o.CustomProperty42, o.CustomProperty43, o.CustomProperty44, o.CustomProperty45,
		o.CustomProperty46, o.CustomProperty47, o.CustomProperty48, o.CustomProperty49, o.CustomProperty50,
		o.CustomProperty51, o.CustomProperty52, o.CustomProperty53, o.CustomProperty54, o.CustomProperty55,
		o.CustomProperty56, o.CustomProperty57, o.CustomProperty58, o.CustomProperty59, o.CustomProperty60,
	}

	for i, value := range customProperties {
		if !IsNil(value) {
			fieldName := fmt.Sprintf("customproperty%d", i+1)
			toSerialize[fieldName] = value
		}
	}

	return toSerialize, nil
}

// NullableGetRoleRequest handles optional values
type NullableGetRoleRequest struct {
	value *GetRoleRequest
	isSet bool
}

// Get retrieves the stored value
func (v NullableGetRoleRequest) Get() *GetRoleRequest {
	return v.value
}

// Set assigns a new value
func (v *NullableGetRoleRequest) Set(val *GetRoleRequest) {
	v.value = val
	v.isSet = true
}

// IsSet checks if a value is set
func (v NullableGetRoleRequest) IsSet() bool {
	return v.isSet
}

// Unset clears the stored value
func (v *NullableGetRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGetRoleRequest initializes a new nullable object
func NewNullableGetRoleRequest(val *GetRoleRequest) *NullableGetRoleRequest {
	return &NullableGetRoleRequest{value: val, isSet: true}
}

// MarshalJSON serializes NullableGetRoleRequest
func (v NullableGetRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes NullableGetRoleRequest
func (v *NullableGetRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
