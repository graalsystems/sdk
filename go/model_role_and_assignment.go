/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems

import (
	"encoding/json"
)

// checks if the RoleAndAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAndAssignment{}

// RoleAndAssignment struct for RoleAndAssignment
type RoleAndAssignment struct {
	Role *Role `json:"role,omitempty"`
	Assignment *RoleAssignment `json:"assignment,omitempty"`
}

// NewRoleAndAssignment instantiates a new RoleAndAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAndAssignment() *RoleAndAssignment {
	this := RoleAndAssignment{}
	return &this
}

// NewRoleAndAssignmentWithDefaults instantiates a new RoleAndAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAndAssignmentWithDefaults() *RoleAndAssignment {
	this := RoleAndAssignment{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleAndAssignment) GetRole() Role {
	if o == nil || IsNil(o.Role) {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAndAssignment) GetRoleOk() (*Role, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleAndAssignment) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *RoleAndAssignment) SetRole(v Role) {
	o.Role = &v
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *RoleAndAssignment) GetAssignment() RoleAssignment {
	if o == nil || IsNil(o.Assignment) {
		var ret RoleAssignment
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAndAssignment) GetAssignmentOk() (*RoleAssignment, bool) {
	if o == nil || IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *RoleAndAssignment) HasAssignment() bool {
	if o != nil && !IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given RoleAssignment and assigns it to the Assignment field.
func (o *RoleAndAssignment) SetAssignment(v RoleAssignment) {
	o.Assignment = &v
}

func (o RoleAndAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAndAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Assignment) {
		toSerialize["assignment"] = o.Assignment
	}
	return toSerialize, nil
}

type NullableRoleAndAssignment struct {
	value *RoleAndAssignment
	isSet bool
}

func (v NullableRoleAndAssignment) Get() *RoleAndAssignment {
	return v.value
}

func (v *NullableRoleAndAssignment) Set(val *RoleAndAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAndAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAndAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAndAssignment(val *RoleAndAssignment) *NullableRoleAndAssignment {
	return &NullableRoleAndAssignment{value: val, isSet: true}
}

func (v NullableRoleAndAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAndAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


