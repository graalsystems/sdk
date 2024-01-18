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

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

type ITask interface{}

// Task struct for Task
type Task struct {
	ITask
	Depends []string `json:"depends,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Type    *string  `json:"type,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetDepends returns the Depends field value if set, zero value otherwise.
func (o *Task) GetDepends() []string {
	if o == nil || IsNil(o.Depends) {
		var ret []string
		return ret
	}
	return o.Depends
}

// GetDependsOk returns a tuple with the Depends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDependsOk() ([]string, bool) {
	if o == nil || IsNil(o.Depends) {
		return nil, false
	}
	return o.Depends, true
}

// HasDepends returns a boolean if a field has been set.
func (o *Task) HasDepends() bool {
	if o != nil && !IsNil(o.Depends) {
		return true
	}

	return false
}

// SetDepends gets a reference to the given []string and assigns it to the Depends field.
func (o *Task) SetDepends(v []string) {
	o.Depends = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Task) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Task) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Task) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Task) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Task) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Task) SetType(v string) {
	o.Type = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Depends) {
		toSerialize["depends"] = o.Depends
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
