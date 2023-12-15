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
	"fmt"
)

// checks if the DropTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DropTask{}

// DropTask Defines a task in a DAG.  Attributes ---------- id : str     Identifier of a task. params : Params     Parameters of a task.  Methods ------- accept(visitor)     Visit a task using a specified visitor. to_node()     Returns the information about the task (id and parameters). to_edge()     Gets all the dependencies of the task.
type DropTask struct {
	// Identifier of the task.
	Id string `json:"id"`
	Params DropParams `json:"params"`
	// List of all dependencies of the task.
	Dependency []string `json:"dependency"`
	// Type of the drop task.
	Type *string `json:"type,omitempty"`
}

type _DropTask DropTask

// NewDropTask instantiates a new DropTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropTask(id string, params DropParams, dependency []string) *DropTask {
	this := DropTask{}
	this.Id = id
	this.Params = params
	this.Dependency = dependency
	var type_ string = "drop"
	this.Type = &type_
	return &this
}

// NewDropTaskWithDefaults instantiates a new DropTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropTaskWithDefaults() *DropTask {
	this := DropTask{}
	var type_ string = "drop"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *DropTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DropTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DropTask) SetId(v string) {
	o.Id = v
}

// GetParams returns the Params field value
func (o *DropTask) GetParams() DropParams {
	if o == nil {
		var ret DropParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *DropTask) GetParamsOk() (*DropParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *DropTask) SetParams(v DropParams) {
	o.Params = v
}

// GetDependency returns the Dependency field value
func (o *DropTask) GetDependency() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dependency
}

// GetDependencyOk returns a tuple with the Dependency field value
// and a boolean to check if the value has been set.
func (o *DropTask) GetDependencyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependency, true
}

// SetDependency sets field value
func (o *DropTask) SetDependency(v []string) {
	o.Dependency = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DropTask) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropTask) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DropTask) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DropTask) SetType(v string) {
	o.Type = &v
}

func (o DropTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DropTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["params"] = o.Params
	toSerialize["dependency"] = o.Dependency
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *DropTask) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"params",
		"dependency",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDropTask := _DropTask{}

	err = json.Unmarshal(bytes, &varDropTask)

	if err != nil {
		return err
	}

	*o = DropTask(varDropTask)

	return err
}

type NullableDropTask struct {
	value *DropTask
	isSet bool
}

func (v NullableDropTask) Get() *DropTask {
	return v.value
}

func (v *NullableDropTask) Set(val *DropTask) {
	v.value = val
	v.isSet = true
}

func (v NullableDropTask) IsSet() bool {
	return v.isSet
}

func (v *NullableDropTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropTask(val *DropTask) *NullableDropTask {
	return &NullableDropTask{value: val, isSet: true}
}

func (v NullableDropTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


