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

// LowCodeOptionsAllOf struct for LowCodeOptionsAllOf
type LowCodeOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	Mode *string `json:"mode,omitempty"`
	ExecutorInstanceType *string `json:"executor_instance_type,omitempty"`
	NumberExecutors *float32 `json:"number_executors,omitempty"`
	Definition *map[string]interface{} `json:"definition,omitempty"`
}

// NewLowCodeOptionsAllOf instantiates a new LowCodeOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowCodeOptionsAllOf() *LowCodeOptionsAllOf {
	this := LowCodeOptionsAllOf{}
	var type_ string = "lowcode"
	this.Type = &type_
	return &this
}

// NewLowCodeOptionsAllOfWithDefaults instantiates a new LowCodeOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowCodeOptionsAllOfWithDefaults() *LowCodeOptionsAllOf {
	this := LowCodeOptionsAllOf{}
	var type_ string = "lowcode"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LowCodeOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LowCodeOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LowCodeOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *LowCodeOptionsAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptionsAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *LowCodeOptionsAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *LowCodeOptionsAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetExecutorInstanceType returns the ExecutorInstanceType field value if set, zero value otherwise.
func (o *LowCodeOptionsAllOf) GetExecutorInstanceType() string {
	if o == nil || o.ExecutorInstanceType == nil {
		var ret string
		return ret
	}
	return *o.ExecutorInstanceType
}

// GetExecutorInstanceTypeOk returns a tuple with the ExecutorInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptionsAllOf) GetExecutorInstanceTypeOk() (*string, bool) {
	if o == nil || o.ExecutorInstanceType == nil {
		return nil, false
	}
	return o.ExecutorInstanceType, true
}

// HasExecutorInstanceType returns a boolean if a field has been set.
func (o *LowCodeOptionsAllOf) HasExecutorInstanceType() bool {
	if o != nil && o.ExecutorInstanceType != nil {
		return true
	}

	return false
}

// SetExecutorInstanceType gets a reference to the given string and assigns it to the ExecutorInstanceType field.
func (o *LowCodeOptionsAllOf) SetExecutorInstanceType(v string) {
	o.ExecutorInstanceType = &v
}

// GetNumberExecutors returns the NumberExecutors field value if set, zero value otherwise.
func (o *LowCodeOptionsAllOf) GetNumberExecutors() float32 {
	if o == nil || o.NumberExecutors == nil {
		var ret float32
		return ret
	}
	return *o.NumberExecutors
}

// GetNumberExecutorsOk returns a tuple with the NumberExecutors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptionsAllOf) GetNumberExecutorsOk() (*float32, bool) {
	if o == nil || o.NumberExecutors == nil {
		return nil, false
	}
	return o.NumberExecutors, true
}

// HasNumberExecutors returns a boolean if a field has been set.
func (o *LowCodeOptionsAllOf) HasNumberExecutors() bool {
	if o != nil && o.NumberExecutors != nil {
		return true
	}

	return false
}

// SetNumberExecutors gets a reference to the given float32 and assigns it to the NumberExecutors field.
func (o *LowCodeOptionsAllOf) SetNumberExecutors(v float32) {
	o.NumberExecutors = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *LowCodeOptionsAllOf) GetDefinition() map[string]interface{} {
	if o == nil || o.Definition == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LowCodeOptionsAllOf) GetDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *LowCodeOptionsAllOf) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given map[string]interface{} and assigns it to the Definition field.
func (o *LowCodeOptionsAllOf) SetDefinition(v map[string]interface{}) {
	o.Definition = &v
}

func (o LowCodeOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.ExecutorInstanceType != nil {
		toSerialize["executor_instance_type"] = o.ExecutorInstanceType
	}
	if o.NumberExecutors != nil {
		toSerialize["number_executors"] = o.NumberExecutors
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableLowCodeOptionsAllOf struct {
	value *LowCodeOptionsAllOf
	isSet bool
}

func (v NullableLowCodeOptionsAllOf) Get() *LowCodeOptionsAllOf {
	return v.value
}

func (v *NullableLowCodeOptionsAllOf) Set(val *LowCodeOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLowCodeOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLowCodeOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowCodeOptionsAllOf(val *LowCodeOptionsAllOf) *NullableLowCodeOptionsAllOf {
	return &NullableLowCodeOptionsAllOf{value: val, isSet: true}
}

func (v NullableLowCodeOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowCodeOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

