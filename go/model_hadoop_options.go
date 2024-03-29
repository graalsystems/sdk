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

// checks if the HadoopOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HadoopOptions{}

// HadoopOptions struct for HadoopOptions
type HadoopOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	ApplicationType *string `json:"application_type,omitempty"`
	ApplicationId *string `json:"application_id,omitempty"`
	JobId *string `json:"job_id,omitempty"`
	Command *string `json:"command,omitempty"`
}

// NewHadoopOptions instantiates a new HadoopOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHadoopOptions() *HadoopOptions {
	this := HadoopOptions{}
	var type_ string = "hadoop"
	this.Type = &type_
	return &this
}

// NewHadoopOptionsWithDefaults instantiates a new HadoopOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHadoopOptionsWithDefaults() *HadoopOptions {
	this := HadoopOptions{}
	var type_ string = "hadoop"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HadoopOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HadoopOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HadoopOptions) SetType(v string) {
	o.Type = &v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *HadoopOptions) GetApplicationType() string {
	if o == nil || IsNil(o.ApplicationType) {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptions) GetApplicationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationType) {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *HadoopOptions) HasApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *HadoopOptions) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *HadoopOptions) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptions) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *HadoopOptions) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *HadoopOptions) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *HadoopOptions) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptions) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *HadoopOptions) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *HadoopOptions) SetJobId(v string) {
	o.JobId = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *HadoopOptions) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptions) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *HadoopOptions) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *HadoopOptions) SetCommand(v string) {
	o.Command = &v
}

func (o HadoopOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HadoopOptions) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ApplicationType) {
		toSerialize["application_type"] = o.ApplicationType
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	return toSerialize, nil
}

type NullableHadoopOptions struct {
	value *HadoopOptions
	isSet bool
}

func (v NullableHadoopOptions) Get() *HadoopOptions {
	return v.value
}

func (v *NullableHadoopOptions) Set(val *HadoopOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableHadoopOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableHadoopOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHadoopOptions(val *HadoopOptions) *NullableHadoopOptions {
	return &NullableHadoopOptions{value: val, isSet: true}
}

func (v NullableHadoopOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHadoopOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


