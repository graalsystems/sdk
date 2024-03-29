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

// checks if the DagUnverified type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DagUnverified{}

// DagUnverified Creates a DAG (Directed Acyclic Graph) from a JSON file.  Attributes ---------- id : str     Identifier of the DAG. tasks : ListIntVisited     Tasks that make up the DAG.  Methods ------- check_dependencies(tasks)     Checks if the tasks dependencies are valid. check_circle(tasks)     Checks that there are no circular dependencies in the DAG. get_zipped_code(language: LanguageType)     Returns the requirements and code in a zipped folder. get_nodes(tasks)     Returns all the information of the DAG as a list of tasks. get_edges(tasks)     Get all two-by-two dependencies between tasks. to_graph(tasks)     Create a graph from the DAG with the edges and nodes.
type DagUnverified struct {
	// Identifier of the Directed Acyclic Graph (DAG).
	Id string `json:"id"`
	// List of tasks that compose the DAG.
	Tasks []TaskListInner `json:"tasks"`
}

type _DagUnverified DagUnverified

// NewDagUnverified instantiates a new DagUnverified object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDagUnverified(id string, tasks []TaskListInner) *DagUnverified {
	this := DagUnverified{}
	this.Id = id
	this.Tasks = tasks
	return &this
}

// NewDagUnverifiedWithDefaults instantiates a new DagUnverified object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDagUnverifiedWithDefaults() *DagUnverified {
	this := DagUnverified{}
	return &this
}

// GetId returns the Id field value
func (o *DagUnverified) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DagUnverified) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DagUnverified) SetId(v string) {
	o.Id = v
}

// GetTasks returns the Tasks field value
func (o *DagUnverified) GetTasks() []TaskListInner {
	if o == nil {
		var ret []TaskListInner
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *DagUnverified) GetTasksOk() ([]TaskListInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *DagUnverified) SetTasks(v []TaskListInner) {
	o.Tasks = v
}

func (o DagUnverified) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DagUnverified) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tasks"] = o.Tasks
	return toSerialize, nil
}

func (o *DagUnverified) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tasks",
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

	varDagUnverified := _DagUnverified{}

	err = json.Unmarshal(bytes, &varDagUnverified)

	if err != nil {
		return err
	}

	*o = DagUnverified(varDagUnverified)

	return err
}

type NullableDagUnverified struct {
	value *DagUnverified
	isSet bool
}

func (v NullableDagUnverified) Get() *DagUnverified {
	return v.value
}

func (v *NullableDagUnverified) Set(val *DagUnverified) {
	v.value = val
	v.isSet = true
}

func (v NullableDagUnverified) IsSet() bool {
	return v.isSet
}

func (v *NullableDagUnverified) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDagUnverified(val *DagUnverified) *NullableDagUnverified {
	return &NullableDagUnverified{value: val, isSet: true}
}

func (v NullableDagUnverified) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDagUnverified) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


