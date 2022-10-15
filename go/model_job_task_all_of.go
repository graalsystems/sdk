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

// JobTaskAllOf struct for JobTaskAllOf
type JobTaskAllOf struct {
	Type *string `json:"type,omitempty"`
	Ref *string `json:"ref,omitempty"`
}

// NewJobTaskAllOf instantiates a new JobTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobTaskAllOf() *JobTaskAllOf {
	this := JobTaskAllOf{}
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// NewJobTaskAllOfWithDefaults instantiates a new JobTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobTaskAllOfWithDefaults() *JobTaskAllOf {
	this := JobTaskAllOf{}
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JobTaskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobTaskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JobTaskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JobTaskAllOf) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *JobTaskAllOf) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobTaskAllOf) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *JobTaskAllOf) HasRef() bool {
	if o != nil && o.Ref != nil {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *JobTaskAllOf) SetRef(v string) {
	o.Ref = &v
}

func (o JobTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	return json.Marshal(toSerialize)
}

type NullableJobTaskAllOf struct {
	value *JobTaskAllOf
	isSet bool
}

func (v NullableJobTaskAllOf) Get() *JobTaskAllOf {
	return v.value
}

func (v *NullableJobTaskAllOf) Set(val *JobTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJobTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJobTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobTaskAllOf(val *JobTaskAllOf) *NullableJobTaskAllOf {
	return &NullableJobTaskAllOf{value: val, isSet: true}
}

func (v NullableJobTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


