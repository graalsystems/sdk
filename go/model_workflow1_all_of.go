/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// Workflow1AllOf struct for Workflow1AllOf
type Workflow1AllOf struct {
	Type *string `json:"type,omitempty"`
	Parent *Project1 `json:"parent,omitempty"`
}

// NewWorkflow1AllOf instantiates a new Workflow1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow1AllOf() *Workflow1AllOf {
	this := Workflow1AllOf{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// NewWorkflow1AllOfWithDefaults instantiates a new Workflow1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflow1AllOfWithDefaults() *Workflow1AllOf {
	this := Workflow1AllOf{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Workflow1AllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow1AllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Workflow1AllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Workflow1AllOf) SetType(v string) {
	o.Type = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Workflow1AllOf) GetParent() Project1 {
	if o == nil || o.Parent == nil {
		var ret Project1
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow1AllOf) GetParentOk() (*Project1, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Workflow1AllOf) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given Project1 and assigns it to the Parent field.
func (o *Workflow1AllOf) SetParent(v Project1) {
	o.Parent = &v
}

func (o Workflow1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflow1AllOf struct {
	value *Workflow1AllOf
	isSet bool
}

func (v NullableWorkflow1AllOf) Get() *Workflow1AllOf {
	return v.value
}

func (v *NullableWorkflow1AllOf) Set(val *Workflow1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow1AllOf(val *Workflow1AllOf) *NullableWorkflow1AllOf {
	return &NullableWorkflow1AllOf{value: val, isSet: true}
}

func (v NullableWorkflow1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


