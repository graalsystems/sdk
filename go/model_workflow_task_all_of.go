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

// WorkflowTaskAllOf struct for WorkflowTaskAllOf
type WorkflowTaskAllOf struct {
	Type *string `json:"type,omitempty"`
	Ref *string `json:"ref,omitempty"`
}

// NewWorkflowTaskAllOf instantiates a new WorkflowTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskAllOf() *WorkflowTaskAllOf {
	this := WorkflowTaskAllOf{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// NewWorkflowTaskAllOfWithDefaults instantiates a new WorkflowTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskAllOfWithDefaults() *WorkflowTaskAllOf {
	this := WorkflowTaskAllOf{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTaskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTaskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTaskAllOf) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *WorkflowTaskAllOf) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskAllOf) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *WorkflowTaskAllOf) HasRef() bool {
	if o != nil && o.Ref != nil {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *WorkflowTaskAllOf) SetRef(v string) {
	o.Ref = &v
}

func (o WorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskAllOf struct {
	value *WorkflowTaskAllOf
	isSet bool
}

func (v NullableWorkflowTaskAllOf) Get() *WorkflowTaskAllOf {
	return v.value
}

func (v *NullableWorkflowTaskAllOf) Set(val *WorkflowTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskAllOf(val *WorkflowTaskAllOf) *NullableWorkflowTaskAllOf {
	return &NullableWorkflowTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

