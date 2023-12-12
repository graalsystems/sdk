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

// checks if the TrinoOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrinoOptions{}

// TrinoOptions struct for TrinoOptions
type TrinoOptions struct {
	DataWarehouseOptions
	Env *map[string]string `json:"env,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	Type *string `json:"type,omitempty"`
	NumberWorkers *float32 `json:"number_workers,omitempty"`
	Conf map[string]map[string]interface{} `json:"conf,omitempty"`
}

// NewTrinoOptions instantiates a new TrinoOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrinoOptions() *TrinoOptions {
	this := TrinoOptions{}
	var type_ string = "trino"
	this.Type = &type_
	return &this
}

// NewTrinoOptionsWithDefaults instantiates a new TrinoOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrinoOptionsWithDefaults() *TrinoOptions {
	this := TrinoOptions{}
	var type_ string = "trino"
	this.Type = &type_
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TrinoOptions) GetEnv() map[string]string {
	if o == nil || IsNil(o.Env) {
		var ret map[string]string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptions) GetEnvOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TrinoOptions) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given map[string]string and assigns it to the Env field.
func (o *TrinoOptions) SetEnv(v map[string]string) {
	o.Env = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *TrinoOptions) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptions) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *TrinoOptions) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *TrinoOptions) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrinoOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrinoOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TrinoOptions) SetType(v string) {
	o.Type = &v
}

// GetNumberWorkers returns the NumberWorkers field value if set, zero value otherwise.
func (o *TrinoOptions) GetNumberWorkers() float32 {
	if o == nil || IsNil(o.NumberWorkers) {
		var ret float32
		return ret
	}
	return *o.NumberWorkers
}

// GetNumberWorkersOk returns a tuple with the NumberWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptions) GetNumberWorkersOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberWorkers) {
		return nil, false
	}
	return o.NumberWorkers, true
}

// HasNumberWorkers returns a boolean if a field has been set.
func (o *TrinoOptions) HasNumberWorkers() bool {
	if o != nil && !IsNil(o.NumberWorkers) {
		return true
	}

	return false
}

// SetNumberWorkers gets a reference to the given float32 and assigns it to the NumberWorkers field.
func (o *TrinoOptions) SetNumberWorkers(v float32) {
	o.NumberWorkers = &v
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *TrinoOptions) GetConf() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Conf) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptions) GetConfOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conf) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *TrinoOptions) HasConf() bool {
	if o != nil && !IsNil(o.Conf) {
		return true
	}

	return false
}

// SetConf gets a reference to the given map[string]map[string]interface{} and assigns it to the Conf field.
func (o *TrinoOptions) SetConf(v map[string]map[string]interface{}) {
	o.Conf = v
}

func (o TrinoOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrinoOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataWarehouseOptions, errDataWarehouseOptions := json.Marshal(o.DataWarehouseOptions)
	if errDataWarehouseOptions != nil {
		return map[string]interface{}{}, errDataWarehouseOptions
	}
	errDataWarehouseOptions = json.Unmarshal([]byte(serializedDataWarehouseOptions), &toSerialize)
	if errDataWarehouseOptions != nil {
		return map[string]interface{}{}, errDataWarehouseOptions
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NumberWorkers) {
		toSerialize["number_workers"] = o.NumberWorkers
	}
	if !IsNil(o.Conf) {
		toSerialize["conf"] = o.Conf
	}
	return toSerialize, nil
}

type NullableTrinoOptions struct {
	value *TrinoOptions
	isSet bool
}

func (v NullableTrinoOptions) Get() *TrinoOptions {
	return v.value
}

func (v *NullableTrinoOptions) Set(val *TrinoOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTrinoOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTrinoOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrinoOptions(val *TrinoOptions) *NullableTrinoOptions {
	return &NullableTrinoOptions{value: val, isSet: true}
}

func (v NullableTrinoOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrinoOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


