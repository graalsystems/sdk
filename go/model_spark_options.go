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

// checks if the SparkOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SparkOptions{}

// SparkOptions struct for SparkOptions
type SparkOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	Libraries []string `json:"libraries,omitempty"`
	Conf map[string]map[string]interface{} `json:"conf,omitempty"`
	MainLibrary *Library `json:"main_library,omitempty"`
	PyFiles []Library `json:"py_files,omitempty"`
	MainClassName *string `json:"main_class_name,omitempty"`
	Loggers []string `json:"loggers,omitempty"`
	ExecutorInstanceType *string `json:"executor_instance_type,omitempty"`
	DriverInstanceType *string `json:"driver_instance_type,omitempty"`
	NumberExecutors *float32 `json:"number_executors,omitempty"`
}

// NewSparkOptions instantiates a new SparkOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparkOptions() *SparkOptions {
	this := SparkOptions{}
	var type_ string = "spark"
	this.Type = &type_
	return &this
}

// NewSparkOptionsWithDefaults instantiates a new SparkOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparkOptionsWithDefaults() *SparkOptions {
	this := SparkOptions{}
	var type_ string = "spark"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SparkOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SparkOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SparkOptions) SetType(v string) {
	o.Type = &v
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *SparkOptions) GetLibraries() []string {
	if o == nil || IsNil(o.Libraries) {
		var ret []string
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetLibrariesOk() ([]string, bool) {
	if o == nil || IsNil(o.Libraries) {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *SparkOptions) HasLibraries() bool {
	if o != nil && !IsNil(o.Libraries) {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []string and assigns it to the Libraries field.
func (o *SparkOptions) SetLibraries(v []string) {
	o.Libraries = v
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *SparkOptions) GetConf() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Conf) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetConfOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conf) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *SparkOptions) HasConf() bool {
	if o != nil && !IsNil(o.Conf) {
		return true
	}

	return false
}

// SetConf gets a reference to the given map[string]map[string]interface{} and assigns it to the Conf field.
func (o *SparkOptions) SetConf(v map[string]map[string]interface{}) {
	o.Conf = v
}

// GetMainLibrary returns the MainLibrary field value if set, zero value otherwise.
func (o *SparkOptions) GetMainLibrary() Library {
	if o == nil || IsNil(o.MainLibrary) {
		var ret Library
		return ret
	}
	return *o.MainLibrary
}

// GetMainLibraryOk returns a tuple with the MainLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetMainLibraryOk() (*Library, bool) {
	if o == nil || IsNil(o.MainLibrary) {
		return nil, false
	}
	return o.MainLibrary, true
}

// HasMainLibrary returns a boolean if a field has been set.
func (o *SparkOptions) HasMainLibrary() bool {
	if o != nil && !IsNil(o.MainLibrary) {
		return true
	}

	return false
}

// SetMainLibrary gets a reference to the given Library and assigns it to the MainLibrary field.
func (o *SparkOptions) SetMainLibrary(v Library) {
	o.MainLibrary = &v
}

// GetPyFiles returns the PyFiles field value if set, zero value otherwise.
func (o *SparkOptions) GetPyFiles() []Library {
	if o == nil || IsNil(o.PyFiles) {
		var ret []Library
		return ret
	}
	return o.PyFiles
}

// GetPyFilesOk returns a tuple with the PyFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetPyFilesOk() ([]Library, bool) {
	if o == nil || IsNil(o.PyFiles) {
		return nil, false
	}
	return o.PyFiles, true
}

// HasPyFiles returns a boolean if a field has been set.
func (o *SparkOptions) HasPyFiles() bool {
	if o != nil && !IsNil(o.PyFiles) {
		return true
	}

	return false
}

// SetPyFiles gets a reference to the given []Library and assigns it to the PyFiles field.
func (o *SparkOptions) SetPyFiles(v []Library) {
	o.PyFiles = v
}

// GetMainClassName returns the MainClassName field value if set, zero value otherwise.
func (o *SparkOptions) GetMainClassName() string {
	if o == nil || IsNil(o.MainClassName) {
		var ret string
		return ret
	}
	return *o.MainClassName
}

// GetMainClassNameOk returns a tuple with the MainClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetMainClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.MainClassName) {
		return nil, false
	}
	return o.MainClassName, true
}

// HasMainClassName returns a boolean if a field has been set.
func (o *SparkOptions) HasMainClassName() bool {
	if o != nil && !IsNil(o.MainClassName) {
		return true
	}

	return false
}

// SetMainClassName gets a reference to the given string and assigns it to the MainClassName field.
func (o *SparkOptions) SetMainClassName(v string) {
	o.MainClassName = &v
}

// GetLoggers returns the Loggers field value if set, zero value otherwise.
func (o *SparkOptions) GetLoggers() []string {
	if o == nil || IsNil(o.Loggers) {
		var ret []string
		return ret
	}
	return o.Loggers
}

// GetLoggersOk returns a tuple with the Loggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetLoggersOk() ([]string, bool) {
	if o == nil || IsNil(o.Loggers) {
		return nil, false
	}
	return o.Loggers, true
}

// HasLoggers returns a boolean if a field has been set.
func (o *SparkOptions) HasLoggers() bool {
	if o != nil && !IsNil(o.Loggers) {
		return true
	}

	return false
}

// SetLoggers gets a reference to the given []string and assigns it to the Loggers field.
func (o *SparkOptions) SetLoggers(v []string) {
	o.Loggers = v
}

// GetExecutorInstanceType returns the ExecutorInstanceType field value if set, zero value otherwise.
func (o *SparkOptions) GetExecutorInstanceType() string {
	if o == nil || IsNil(o.ExecutorInstanceType) {
		var ret string
		return ret
	}
	return *o.ExecutorInstanceType
}

// GetExecutorInstanceTypeOk returns a tuple with the ExecutorInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetExecutorInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutorInstanceType) {
		return nil, false
	}
	return o.ExecutorInstanceType, true
}

// HasExecutorInstanceType returns a boolean if a field has been set.
func (o *SparkOptions) HasExecutorInstanceType() bool {
	if o != nil && !IsNil(o.ExecutorInstanceType) {
		return true
	}

	return false
}

// SetExecutorInstanceType gets a reference to the given string and assigns it to the ExecutorInstanceType field.
func (o *SparkOptions) SetExecutorInstanceType(v string) {
	o.ExecutorInstanceType = &v
}

// GetDriverInstanceType returns the DriverInstanceType field value if set, zero value otherwise.
func (o *SparkOptions) GetDriverInstanceType() string {
	if o == nil || IsNil(o.DriverInstanceType) {
		var ret string
		return ret
	}
	return *o.DriverInstanceType
}

// GetDriverInstanceTypeOk returns a tuple with the DriverInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetDriverInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DriverInstanceType) {
		return nil, false
	}
	return o.DriverInstanceType, true
}

// HasDriverInstanceType returns a boolean if a field has been set.
func (o *SparkOptions) HasDriverInstanceType() bool {
	if o != nil && !IsNil(o.DriverInstanceType) {
		return true
	}

	return false
}

// SetDriverInstanceType gets a reference to the given string and assigns it to the DriverInstanceType field.
func (o *SparkOptions) SetDriverInstanceType(v string) {
	o.DriverInstanceType = &v
}

// GetNumberExecutors returns the NumberExecutors field value if set, zero value otherwise.
func (o *SparkOptions) GetNumberExecutors() float32 {
	if o == nil || IsNil(o.NumberExecutors) {
		var ret float32
		return ret
	}
	return *o.NumberExecutors
}

// GetNumberExecutorsOk returns a tuple with the NumberExecutors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparkOptions) GetNumberExecutorsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberExecutors) {
		return nil, false
	}
	return o.NumberExecutors, true
}

// HasNumberExecutors returns a boolean if a field has been set.
func (o *SparkOptions) HasNumberExecutors() bool {
	if o != nil && !IsNil(o.NumberExecutors) {
		return true
	}

	return false
}

// SetNumberExecutors gets a reference to the given float32 and assigns it to the NumberExecutors field.
func (o *SparkOptions) SetNumberExecutors(v float32) {
	o.NumberExecutors = &v
}

func (o SparkOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SparkOptions) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Libraries) {
		toSerialize["libraries"] = o.Libraries
	}
	if !IsNil(o.Conf) {
		toSerialize["conf"] = o.Conf
	}
	if !IsNil(o.MainLibrary) {
		toSerialize["main_library"] = o.MainLibrary
	}
	if !IsNil(o.PyFiles) {
		toSerialize["py_files"] = o.PyFiles
	}
	if !IsNil(o.MainClassName) {
		toSerialize["main_class_name"] = o.MainClassName
	}
	if !IsNil(o.Loggers) {
		toSerialize["loggers"] = o.Loggers
	}
	if !IsNil(o.ExecutorInstanceType) {
		toSerialize["executor_instance_type"] = o.ExecutorInstanceType
	}
	if !IsNil(o.DriverInstanceType) {
		toSerialize["driver_instance_type"] = o.DriverInstanceType
	}
	if !IsNil(o.NumberExecutors) {
		toSerialize["number_executors"] = o.NumberExecutors
	}
	return toSerialize, nil
}

type NullableSparkOptions struct {
	value *SparkOptions
	isSet bool
}

func (v NullableSparkOptions) Get() *SparkOptions {
	return v.value
}

func (v *NullableSparkOptions) Set(val *SparkOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSparkOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSparkOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparkOptions(val *SparkOptions) *NullableSparkOptions {
	return &NullableSparkOptions{value: val, isSet: true}
}

func (v NullableSparkOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparkOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


