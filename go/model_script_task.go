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

// ScriptTask struct for ScriptTask
type ScriptTask struct {
	Task
	Type *string `json:"type,omitempty"`
	Script *string `json:"script,omitempty"`
}

// NewScriptTask instantiates a new ScriptTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptTask() *ScriptTask {
	this := ScriptTask{}
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// NewScriptTaskWithDefaults instantiates a new ScriptTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptTaskWithDefaults() *ScriptTask {
	this := ScriptTask{}
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScriptTask) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptTask) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScriptTask) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScriptTask) SetType(v string) {
	o.Type = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *ScriptTask) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptTask) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *ScriptTask) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *ScriptTask) SetScript(v string) {
	o.Script = &v
}

func (o ScriptTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTask, errTask := json.Marshal(o.Task)
	if errTask != nil {
		return []byte{}, errTask
	}
	errTask = json.Unmarshal([]byte(serializedTask), &toSerialize)
	if errTask != nil {
		return []byte{}, errTask
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullableScriptTask struct {
	value *ScriptTask
	isSet bool
}

func (v NullableScriptTask) Get() *ScriptTask {
	return v.value
}

func (v *NullableScriptTask) Set(val *ScriptTask) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptTask) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptTask(val *ScriptTask) *NullableScriptTask {
	return &NullableScriptTask{value: val, isSet: true}
}

func (v NullableScriptTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

