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

// checks if the WorkflowTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTask{}

// WorkflowTask struct for WorkflowTask
type WorkflowTask struct {
	Task
	Type *string `json:"type,omitempty"`
	Ref *string `json:"ref,omitempty"`
}

// NewWorkflowTask instantiates a new WorkflowTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTask() *WorkflowTask {
	this := WorkflowTask{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// NewWorkflowTaskWithDefaults instantiates a new WorkflowTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskWithDefaults() *WorkflowTask {
	this := WorkflowTask{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTask) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTask) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTask) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTask) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *WorkflowTask) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTask) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *WorkflowTask) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *WorkflowTask) SetRef(v string) {
	o.Ref = &v
}

func (o WorkflowTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTask, errTask := json.Marshal(o.Task)
	if errTask != nil {
		return map[string]interface{}{}, errTask
	}
	errTask = json.Unmarshal([]byte(serializedTask), &toSerialize)
	if errTask != nil {
		return map[string]interface{}{}, errTask
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	return toSerialize, nil
}

type NullableWorkflowTask struct {
	value *WorkflowTask
	isSet bool
}

func (v NullableWorkflowTask) Get() *WorkflowTask {
	return v.value
}

func (v *NullableWorkflowTask) Set(val *WorkflowTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTask(val *WorkflowTask) *NullableWorkflowTask {
	return &NullableWorkflowTask{value: val, isSet: true}
}

func (v NullableWorkflowTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


