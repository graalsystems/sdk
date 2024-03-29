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

// FlinkOptionsAllOf struct for FlinkOptionsAllOf
type FlinkOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	MainLibrary *Library `json:"main_library,omitempty"`
	MainClassName *string `json:"main_class_name,omitempty"`
	Loggers *[]string `json:"loggers,omitempty"`
	JobManagerInstanceType *string `json:"job_manager_instance_type,omitempty"`
	TaskManagerInstanceType *string `json:"task_manager_instance_type,omitempty"`
	NumberTaskManagers *float32 `json:"number_task_managers,omitempty"`
}

// NewFlinkOptionsAllOf instantiates a new FlinkOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkOptionsAllOf() *FlinkOptionsAllOf {
	this := FlinkOptionsAllOf{}
	var type_ string = "flink"
	this.Type = &type_
	return &this
}

// NewFlinkOptionsAllOfWithDefaults instantiates a new FlinkOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkOptionsAllOfWithDefaults() *FlinkOptionsAllOf {
	this := FlinkOptionsAllOf{}
	var type_ string = "flink"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlinkOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetMainLibrary returns the MainLibrary field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetMainLibrary() Library {
	if o == nil || o.MainLibrary == nil {
		var ret Library
		return ret
	}
	return *o.MainLibrary
}

// GetMainLibraryOk returns a tuple with the MainLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetMainLibraryOk() (*Library, bool) {
	if o == nil || o.MainLibrary == nil {
		return nil, false
	}
	return o.MainLibrary, true
}

// HasMainLibrary returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasMainLibrary() bool {
	if o != nil && o.MainLibrary != nil {
		return true
	}

	return false
}

// SetMainLibrary gets a reference to the given Library and assigns it to the MainLibrary field.
func (o *FlinkOptionsAllOf) SetMainLibrary(v Library) {
	o.MainLibrary = &v
}

// GetMainClassName returns the MainClassName field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetMainClassName() string {
	if o == nil || o.MainClassName == nil {
		var ret string
		return ret
	}
	return *o.MainClassName
}

// GetMainClassNameOk returns a tuple with the MainClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetMainClassNameOk() (*string, bool) {
	if o == nil || o.MainClassName == nil {
		return nil, false
	}
	return o.MainClassName, true
}

// HasMainClassName returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasMainClassName() bool {
	if o != nil && o.MainClassName != nil {
		return true
	}

	return false
}

// SetMainClassName gets a reference to the given string and assigns it to the MainClassName field.
func (o *FlinkOptionsAllOf) SetMainClassName(v string) {
	o.MainClassName = &v
}

// GetLoggers returns the Loggers field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetLoggers() []string {
	if o == nil || o.Loggers == nil {
		var ret []string
		return ret
	}
	return *o.Loggers
}

// GetLoggersOk returns a tuple with the Loggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetLoggersOk() (*[]string, bool) {
	if o == nil || o.Loggers == nil {
		return nil, false
	}
	return o.Loggers, true
}

// HasLoggers returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasLoggers() bool {
	if o != nil && o.Loggers != nil {
		return true
	}

	return false
}

// SetLoggers gets a reference to the given []string and assigns it to the Loggers field.
func (o *FlinkOptionsAllOf) SetLoggers(v []string) {
	o.Loggers = &v
}

// GetJobManagerInstanceType returns the JobManagerInstanceType field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetJobManagerInstanceType() string {
	if o == nil || o.JobManagerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.JobManagerInstanceType
}

// GetJobManagerInstanceTypeOk returns a tuple with the JobManagerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetJobManagerInstanceTypeOk() (*string, bool) {
	if o == nil || o.JobManagerInstanceType == nil {
		return nil, false
	}
	return o.JobManagerInstanceType, true
}

// HasJobManagerInstanceType returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasJobManagerInstanceType() bool {
	if o != nil && o.JobManagerInstanceType != nil {
		return true
	}

	return false
}

// SetJobManagerInstanceType gets a reference to the given string and assigns it to the JobManagerInstanceType field.
func (o *FlinkOptionsAllOf) SetJobManagerInstanceType(v string) {
	o.JobManagerInstanceType = &v
}

// GetTaskManagerInstanceType returns the TaskManagerInstanceType field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetTaskManagerInstanceType() string {
	if o == nil || o.TaskManagerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.TaskManagerInstanceType
}

// GetTaskManagerInstanceTypeOk returns a tuple with the TaskManagerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetTaskManagerInstanceTypeOk() (*string, bool) {
	if o == nil || o.TaskManagerInstanceType == nil {
		return nil, false
	}
	return o.TaskManagerInstanceType, true
}

// HasTaskManagerInstanceType returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasTaskManagerInstanceType() bool {
	if o != nil && o.TaskManagerInstanceType != nil {
		return true
	}

	return false
}

// SetTaskManagerInstanceType gets a reference to the given string and assigns it to the TaskManagerInstanceType field.
func (o *FlinkOptionsAllOf) SetTaskManagerInstanceType(v string) {
	o.TaskManagerInstanceType = &v
}

// GetNumberTaskManagers returns the NumberTaskManagers field value if set, zero value otherwise.
func (o *FlinkOptionsAllOf) GetNumberTaskManagers() float32 {
	if o == nil || o.NumberTaskManagers == nil {
		var ret float32
		return ret
	}
	return *o.NumberTaskManagers
}

// GetNumberTaskManagersOk returns a tuple with the NumberTaskManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptionsAllOf) GetNumberTaskManagersOk() (*float32, bool) {
	if o == nil || o.NumberTaskManagers == nil {
		return nil, false
	}
	return o.NumberTaskManagers, true
}

// HasNumberTaskManagers returns a boolean if a field has been set.
func (o *FlinkOptionsAllOf) HasNumberTaskManagers() bool {
	if o != nil && o.NumberTaskManagers != nil {
		return true
	}

	return false
}

// SetNumberTaskManagers gets a reference to the given float32 and assigns it to the NumberTaskManagers field.
func (o *FlinkOptionsAllOf) SetNumberTaskManagers(v float32) {
	o.NumberTaskManagers = &v
}

func (o FlinkOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MainLibrary != nil {
		toSerialize["main_library"] = o.MainLibrary
	}
	if o.MainClassName != nil {
		toSerialize["main_class_name"] = o.MainClassName
	}
	if o.Loggers != nil {
		toSerialize["loggers"] = o.Loggers
	}
	if o.JobManagerInstanceType != nil {
		toSerialize["job_manager_instance_type"] = o.JobManagerInstanceType
	}
	if o.TaskManagerInstanceType != nil {
		toSerialize["task_manager_instance_type"] = o.TaskManagerInstanceType
	}
	if o.NumberTaskManagers != nil {
		toSerialize["number_task_managers"] = o.NumberTaskManagers
	}
	return json.Marshal(toSerialize)
}

type NullableFlinkOptionsAllOf struct {
	value *FlinkOptionsAllOf
	isSet bool
}

func (v NullableFlinkOptionsAllOf) Get() *FlinkOptionsAllOf {
	return v.value
}

func (v *NullableFlinkOptionsAllOf) Set(val *FlinkOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkOptionsAllOf(val *FlinkOptionsAllOf) *NullableFlinkOptionsAllOf {
	return &NullableFlinkOptionsAllOf{value: val, isSet: true}
}

func (v NullableFlinkOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


