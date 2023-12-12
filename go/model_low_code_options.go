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

// checks if the LowCodeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowCodeOptions{}

// LowCodeOptions struct for LowCodeOptions
type LowCodeOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	Mode *string `json:"mode,omitempty"`
	ExecutorInstanceType *string `json:"executor_instance_type,omitempty"`
	NumberExecutors *float32 `json:"number_executors,omitempty"`
	Definition map[string]interface{} `json:"definition,omitempty"`
}

// NewLowCodeOptions instantiates a new LowCodeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowCodeOptions() *LowCodeOptions {
	this := LowCodeOptions{}
	var type_ string = "lowcode"
	this.Type = &type_
	return &this
}

// NewLowCodeOptionsWithDefaults instantiates a new LowCodeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowCodeOptionsWithDefaults() *LowCodeOptions {
	this := LowCodeOptions{}
	var type_ string = "lowcode"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LowCodeOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LowCodeOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LowCodeOptions) SetType(v string) {
	o.Type = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *LowCodeOptions) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptions) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *LowCodeOptions) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *LowCodeOptions) SetMode(v string) {
	o.Mode = &v
}

// GetExecutorInstanceType returns the ExecutorInstanceType field value if set, zero value otherwise.
func (o *LowCodeOptions) GetExecutorInstanceType() string {
	if o == nil || IsNil(o.ExecutorInstanceType) {
		var ret string
		return ret
	}
	return *o.ExecutorInstanceType
}

// GetExecutorInstanceTypeOk returns a tuple with the ExecutorInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptions) GetExecutorInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutorInstanceType) {
		return nil, false
	}
	return o.ExecutorInstanceType, true
}

// HasExecutorInstanceType returns a boolean if a field has been set.
func (o *LowCodeOptions) HasExecutorInstanceType() bool {
	if o != nil && !IsNil(o.ExecutorInstanceType) {
		return true
	}

	return false
}

// SetExecutorInstanceType gets a reference to the given string and assigns it to the ExecutorInstanceType field.
func (o *LowCodeOptions) SetExecutorInstanceType(v string) {
	o.ExecutorInstanceType = &v
}

// GetNumberExecutors returns the NumberExecutors field value if set, zero value otherwise.
func (o *LowCodeOptions) GetNumberExecutors() float32 {
	if o == nil || IsNil(o.NumberExecutors) {
		var ret float32
		return ret
	}
	return *o.NumberExecutors
}

// GetNumberExecutorsOk returns a tuple with the NumberExecutors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptions) GetNumberExecutorsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberExecutors) {
		return nil, false
	}
	return o.NumberExecutors, true
}

// HasNumberExecutors returns a boolean if a field has been set.
func (o *LowCodeOptions) HasNumberExecutors() bool {
	if o != nil && !IsNil(o.NumberExecutors) {
		return true
	}

	return false
}

// SetNumberExecutors gets a reference to the given float32 and assigns it to the NumberExecutors field.
func (o *LowCodeOptions) SetNumberExecutors(v float32) {
	o.NumberExecutors = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *LowCodeOptions) GetDefinition() map[string]interface{} {
	if o == nil || IsNil(o.Definition) {
		var ret map[string]interface{}
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptions) GetDefinitionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Definition) {
		return map[string]interface{}{}, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *LowCodeOptions) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given map[string]interface{} and assigns it to the Definition field.
func (o *LowCodeOptions) SetDefinition(v map[string]interface{}) {
	o.Definition = v
}

func (o LowCodeOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LowCodeOptions) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ExecutorInstanceType) {
		toSerialize["executor_instance_type"] = o.ExecutorInstanceType
	}
	if !IsNil(o.NumberExecutors) {
		toSerialize["number_executors"] = o.NumberExecutors
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return toSerialize, nil
}

type NullableLowCodeOptions struct {
	value *LowCodeOptions
	isSet bool
}

func (v NullableLowCodeOptions) Get() *LowCodeOptions {
	return v.value
}

func (v *NullableLowCodeOptions) Set(val *LowCodeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLowCodeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLowCodeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowCodeOptions(val *LowCodeOptions) *NullableLowCodeOptions {
	return &NullableLowCodeOptions{value: val, isSet: true}
}

func (v NullableLowCodeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowCodeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


