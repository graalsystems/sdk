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
	"fmt"
)

// checks if the PredictModelTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictModelTask{}

// PredictModelTask Defines a task in a DAG.  Attributes ---------- id : str     Identifier of a task. params : Params     Parameters of a task.  Methods ------- accept(visitor)     Visit a task using a specified visitor. to_node()     Returns the information about the task (id and parameters). to_edge()     Gets all the dependencies of the task.
type PredictModelTask struct {
	// Identifier of the task.
	Id string `json:"id"`
	Params map[string]interface{} `json:"params"`
	// List of all dependencies of the task.
	Dependency []string `json:"dependency"`
	// Type of the predict model task.
	Type *string `json:"type,omitempty"`
}

type _PredictModelTask PredictModelTask

// NewPredictModelTask instantiates a new PredictModelTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictModelTask(id string, params map[string]interface{}, dependency []string) *PredictModelTask {
	this := PredictModelTask{}
	this.Id = id
	this.Params = params
	this.Dependency = dependency
	var type_ string = "predict_model"
	this.Type = &type_
	return &this
}

// NewPredictModelTaskWithDefaults instantiates a new PredictModelTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictModelTaskWithDefaults() *PredictModelTask {
	this := PredictModelTask{}
	var type_ string = "predict_model"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *PredictModelTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PredictModelTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PredictModelTask) SetId(v string) {
	o.Id = v
}

// GetParams returns the Params field value
func (o *PredictModelTask) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *PredictModelTask) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *PredictModelTask) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetDependency returns the Dependency field value
func (o *PredictModelTask) GetDependency() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dependency
}

// GetDependencyOk returns a tuple with the Dependency field value
// and a boolean to check if the value has been set.
func (o *PredictModelTask) GetDependencyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dependency, true
}

// SetDependency sets field value
func (o *PredictModelTask) SetDependency(v []string) {
	o.Dependency = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PredictModelTask) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictModelTask) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PredictModelTask) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PredictModelTask) SetType(v string) {
	o.Type = &v
}

func (o PredictModelTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictModelTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["params"] = o.Params
	toSerialize["dependency"] = o.Dependency
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *PredictModelTask) UnmarshalJSON(bytes []byte) (err error) {
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

	varPredictModelTask := _PredictModelTask{}

	err = json.Unmarshal(bytes, &varPredictModelTask)

	if err != nil {
		return err
	}

	*o = PredictModelTask(varPredictModelTask)

	return err
}

type NullablePredictModelTask struct {
	value *PredictModelTask
	isSet bool
}

func (v NullablePredictModelTask) Get() *PredictModelTask {
	return v.value
}

func (v *NullablePredictModelTask) Set(val *PredictModelTask) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictModelTask) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictModelTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictModelTask(val *PredictModelTask) *NullablePredictModelTask {
	return &NullablePredictModelTask{value: val, isSet: true}
}

func (v NullablePredictModelTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictModelTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


