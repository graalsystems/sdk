/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

import (
	"encoding/json"
)

// RunOnceScheduleAllOf struct for RunOnceScheduleAllOf
type RunOnceScheduleAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewRunOnceScheduleAllOf instantiates a new RunOnceScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunOnceScheduleAllOf() *RunOnceScheduleAllOf {
	this := RunOnceScheduleAllOf{}
	var type_ string = "once"
	this.Type = &type_
	return &this
}

// NewRunOnceScheduleAllOfWithDefaults instantiates a new RunOnceScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunOnceScheduleAllOfWithDefaults() *RunOnceScheduleAllOf {
	this := RunOnceScheduleAllOf{}
	var type_ string = "once"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RunOnceScheduleAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunOnceScheduleAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RunOnceScheduleAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RunOnceScheduleAllOf) SetType(v string) {
	o.Type = &v
}

func (o RunOnceScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRunOnceScheduleAllOf struct {
	value *RunOnceScheduleAllOf
	isSet bool
}

func (v NullableRunOnceScheduleAllOf) Get() *RunOnceScheduleAllOf {
	return v.value
}

func (v *NullableRunOnceScheduleAllOf) Set(val *RunOnceScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRunOnceScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRunOnceScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunOnceScheduleAllOf(val *RunOnceScheduleAllOf) *NullableRunOnceScheduleAllOf {
	return &NullableRunOnceScheduleAllOf{value: val, isSet: true}
}

func (v NullableRunOnceScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunOnceScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


