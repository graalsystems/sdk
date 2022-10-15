/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// AdminTargetAllOf struct for AdminTargetAllOf
type AdminTargetAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewAdminTargetAllOf instantiates a new AdminTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminTargetAllOf() *AdminTargetAllOf {
	this := AdminTargetAllOf{}
	var type_ string = "admin"
	this.Type = &type_
	return &this
}

// NewAdminTargetAllOfWithDefaults instantiates a new AdminTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminTargetAllOfWithDefaults() *AdminTargetAllOf {
	this := AdminTargetAllOf{}
	var type_ string = "admin"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdminTargetAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTargetAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdminTargetAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdminTargetAllOf) SetType(v string) {
	o.Type = &v
}

func (o AdminTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAdminTargetAllOf struct {
	value *AdminTargetAllOf
	isSet bool
}

func (v NullableAdminTargetAllOf) Get() *AdminTargetAllOf {
	return v.value
}

func (v *NullableAdminTargetAllOf) Set(val *AdminTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminTargetAllOf(val *AdminTargetAllOf) *NullableAdminTargetAllOf {
	return &NullableAdminTargetAllOf{value: val, isSet: true}
}

func (v NullableAdminTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

