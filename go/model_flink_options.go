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

// checks if the FlinkOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlinkOptions{}

// FlinkOptions struct for FlinkOptions
type FlinkOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	MainLibrary *Library `json:"main_library,omitempty"`
	MainClassName *string `json:"main_class_name,omitempty"`
	Loggers []string `json:"loggers,omitempty"`
	JobManagerInstanceType *string `json:"job_manager_instance_type,omitempty"`
	TaskManagerInstanceType *string `json:"task_manager_instance_type,omitempty"`
	NumberTaskManagers *float32 `json:"number_task_managers,omitempty"`
}

// NewFlinkOptions instantiates a new FlinkOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkOptions() *FlinkOptions {
	this := FlinkOptions{}
	var type_ string = "flink"
	this.Type = &type_
	return &this
}

// NewFlinkOptionsWithDefaults instantiates a new FlinkOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkOptionsWithDefaults() *FlinkOptions {
	this := FlinkOptions{}
	var type_ string = "flink"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlinkOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlinkOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlinkOptions) SetType(v string) {
	o.Type = &v
}

// GetMainLibrary returns the MainLibrary field value if set, zero value otherwise.
func (o *FlinkOptions) GetMainLibrary() Library {
	if o == nil || IsNil(o.MainLibrary) {
		var ret Library
		return ret
	}
	return *o.MainLibrary
}

// GetMainLibraryOk returns a tuple with the MainLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetMainLibraryOk() (*Library, bool) {
	if o == nil || IsNil(o.MainLibrary) {
		return nil, false
	}
	return o.MainLibrary, true
}

// HasMainLibrary returns a boolean if a field has been set.
func (o *FlinkOptions) HasMainLibrary() bool {
	if o != nil && !IsNil(o.MainLibrary) {
		return true
	}

	return false
}

// SetMainLibrary gets a reference to the given Library and assigns it to the MainLibrary field.
func (o *FlinkOptions) SetMainLibrary(v Library) {
	o.MainLibrary = &v
}

// GetMainClassName returns the MainClassName field value if set, zero value otherwise.
func (o *FlinkOptions) GetMainClassName() string {
	if o == nil || IsNil(o.MainClassName) {
		var ret string
		return ret
	}
	return *o.MainClassName
}

// GetMainClassNameOk returns a tuple with the MainClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetMainClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.MainClassName) {
		return nil, false
	}
	return o.MainClassName, true
}

// HasMainClassName returns a boolean if a field has been set.
func (o *FlinkOptions) HasMainClassName() bool {
	if o != nil && !IsNil(o.MainClassName) {
		return true
	}

	return false
}

// SetMainClassName gets a reference to the given string and assigns it to the MainClassName field.
func (o *FlinkOptions) SetMainClassName(v string) {
	o.MainClassName = &v
}

// GetLoggers returns the Loggers field value if set, zero value otherwise.
func (o *FlinkOptions) GetLoggers() []string {
	if o == nil || IsNil(o.Loggers) {
		var ret []string
		return ret
	}
	return o.Loggers
}

// GetLoggersOk returns a tuple with the Loggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetLoggersOk() ([]string, bool) {
	if o == nil || IsNil(o.Loggers) {
		return nil, false
	}
	return o.Loggers, true
}

// HasLoggers returns a boolean if a field has been set.
func (o *FlinkOptions) HasLoggers() bool {
	if o != nil && !IsNil(o.Loggers) {
		return true
	}

	return false
}

// SetLoggers gets a reference to the given []string and assigns it to the Loggers field.
func (o *FlinkOptions) SetLoggers(v []string) {
	o.Loggers = v
}

// GetJobManagerInstanceType returns the JobManagerInstanceType field value if set, zero value otherwise.
func (o *FlinkOptions) GetJobManagerInstanceType() string {
	if o == nil || IsNil(o.JobManagerInstanceType) {
		var ret string
		return ret
	}
	return *o.JobManagerInstanceType
}

// GetJobManagerInstanceTypeOk returns a tuple with the JobManagerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetJobManagerInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JobManagerInstanceType) {
		return nil, false
	}
	return o.JobManagerInstanceType, true
}

// HasJobManagerInstanceType returns a boolean if a field has been set.
func (o *FlinkOptions) HasJobManagerInstanceType() bool {
	if o != nil && !IsNil(o.JobManagerInstanceType) {
		return true
	}

	return false
}

// SetJobManagerInstanceType gets a reference to the given string and assigns it to the JobManagerInstanceType field.
func (o *FlinkOptions) SetJobManagerInstanceType(v string) {
	o.JobManagerInstanceType = &v
}

// GetTaskManagerInstanceType returns the TaskManagerInstanceType field value if set, zero value otherwise.
func (o *FlinkOptions) GetTaskManagerInstanceType() string {
	if o == nil || IsNil(o.TaskManagerInstanceType) {
		var ret string
		return ret
	}
	return *o.TaskManagerInstanceType
}

// GetTaskManagerInstanceTypeOk returns a tuple with the TaskManagerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetTaskManagerInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaskManagerInstanceType) {
		return nil, false
	}
	return o.TaskManagerInstanceType, true
}

// HasTaskManagerInstanceType returns a boolean if a field has been set.
func (o *FlinkOptions) HasTaskManagerInstanceType() bool {
	if o != nil && !IsNil(o.TaskManagerInstanceType) {
		return true
	}

	return false
}

// SetTaskManagerInstanceType gets a reference to the given string and assigns it to the TaskManagerInstanceType field.
func (o *FlinkOptions) SetTaskManagerInstanceType(v string) {
	o.TaskManagerInstanceType = &v
}

// GetNumberTaskManagers returns the NumberTaskManagers field value if set, zero value otherwise.
func (o *FlinkOptions) GetNumberTaskManagers() float32 {
	if o == nil || IsNil(o.NumberTaskManagers) {
		var ret float32
		return ret
	}
	return *o.NumberTaskManagers
}

// GetNumberTaskManagersOk returns a tuple with the NumberTaskManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkOptions) GetNumberTaskManagersOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberTaskManagers) {
		return nil, false
	}
	return o.NumberTaskManagers, true
}

// HasNumberTaskManagers returns a boolean if a field has been set.
func (o *FlinkOptions) HasNumberTaskManagers() bool {
	if o != nil && !IsNil(o.NumberTaskManagers) {
		return true
	}

	return false
}

// SetNumberTaskManagers gets a reference to the given float32 and assigns it to the NumberTaskManagers field.
func (o *FlinkOptions) SetNumberTaskManagers(v float32) {
	o.NumberTaskManagers = &v
}

func (o FlinkOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlinkOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOptions, errOptions := json.Marshal(o.Options)
	if errOptions != nil {
		return map[string]interface{}{}, errOptions
	}
	errOptions = json.Unmarshal([]byte(serializedOptions), &toSerialize)
	if errOptions != nil {
		return map[string]interface{}{}, errOptions
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MainLibrary) {
		toSerialize["main_library"] = o.MainLibrary
	}
	if !IsNil(o.MainClassName) {
		toSerialize["main_class_name"] = o.MainClassName
	}
	if !IsNil(o.Loggers) {
		toSerialize["loggers"] = o.Loggers
	}
	if !IsNil(o.JobManagerInstanceType) {
		toSerialize["job_manager_instance_type"] = o.JobManagerInstanceType
	}
	if !IsNil(o.TaskManagerInstanceType) {
		toSerialize["task_manager_instance_type"] = o.TaskManagerInstanceType
	}
	if !IsNil(o.NumberTaskManagers) {
		toSerialize["number_task_managers"] = o.NumberTaskManagers
	}
	return toSerialize, nil
}

type NullableFlinkOptions struct {
	value *FlinkOptions
	isSet bool
}

func (v NullableFlinkOptions) Get() *FlinkOptions {
	return v.value
}

func (v *NullableFlinkOptions) Set(val *FlinkOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkOptions(val *FlinkOptions) *NullableFlinkOptions {
	return &NullableFlinkOptions{value: val, isSet: true}
}

func (v NullableFlinkOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


