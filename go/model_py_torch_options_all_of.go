/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// PyTorchOptionsAllOf struct for PyTorchOptionsAllOf
type PyTorchOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	Module *string `json:"module,omitempty"`
	NumberReplicas *float32 `json:"number_replicas,omitempty"`
	MasterInstanceType *string `json:"master_instance_type,omitempty"`
	WorkerInstanceType *string `json:"worker_instance_type,omitempty"`
}

// NewPyTorchOptionsAllOf instantiates a new PyTorchOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPyTorchOptionsAllOf() *PyTorchOptionsAllOf {
	this := PyTorchOptionsAllOf{}
	var type_ string = "python"
	this.Type = &type_
	return &this
}

// NewPyTorchOptionsAllOfWithDefaults instantiates a new PyTorchOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPyTorchOptionsAllOfWithDefaults() *PyTorchOptionsAllOf {
	this := PyTorchOptionsAllOf{}
	var type_ string = "python"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PyTorchOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyTorchOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PyTorchOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PyTorchOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *PyTorchOptionsAllOf) GetModule() string {
	if o == nil || o.Module == nil {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyTorchOptionsAllOf) GetModuleOk() (*string, bool) {
	if o == nil || o.Module == nil {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *PyTorchOptionsAllOf) HasModule() bool {
	if o != nil && o.Module != nil {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *PyTorchOptionsAllOf) SetModule(v string) {
	o.Module = &v
}

// GetNumberReplicas returns the NumberReplicas field value if set, zero value otherwise.
func (o *PyTorchOptionsAllOf) GetNumberReplicas() float32 {
	if o == nil || o.NumberReplicas == nil {
		var ret float32
		return ret
	}
	return *o.NumberReplicas
}

// GetNumberReplicasOk returns a tuple with the NumberReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyTorchOptionsAllOf) GetNumberReplicasOk() (*float32, bool) {
	if o == nil || o.NumberReplicas == nil {
		return nil, false
	}
	return o.NumberReplicas, true
}

// HasNumberReplicas returns a boolean if a field has been set.
func (o *PyTorchOptionsAllOf) HasNumberReplicas() bool {
	if o != nil && o.NumberReplicas != nil {
		return true
	}

	return false
}

// SetNumberReplicas gets a reference to the given float32 and assigns it to the NumberReplicas field.
func (o *PyTorchOptionsAllOf) SetNumberReplicas(v float32) {
	o.NumberReplicas = &v
}

// GetMasterInstanceType returns the MasterInstanceType field value if set, zero value otherwise.
func (o *PyTorchOptionsAllOf) GetMasterInstanceType() string {
	if o == nil || o.MasterInstanceType == nil {
		var ret string
		return ret
	}
	return *o.MasterInstanceType
}

// GetMasterInstanceTypeOk returns a tuple with the MasterInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyTorchOptionsAllOf) GetMasterInstanceTypeOk() (*string, bool) {
	if o == nil || o.MasterInstanceType == nil {
		return nil, false
	}
	return o.MasterInstanceType, true
}

// HasMasterInstanceType returns a boolean if a field has been set.
func (o *PyTorchOptionsAllOf) HasMasterInstanceType() bool {
	if o != nil && o.MasterInstanceType != nil {
		return true
	}

	return false
}

// SetMasterInstanceType gets a reference to the given string and assigns it to the MasterInstanceType field.
func (o *PyTorchOptionsAllOf) SetMasterInstanceType(v string) {
	o.MasterInstanceType = &v
}

// GetWorkerInstanceType returns the WorkerInstanceType field value if set, zero value otherwise.
func (o *PyTorchOptionsAllOf) GetWorkerInstanceType() string {
	if o == nil || o.WorkerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.WorkerInstanceType
}

// GetWorkerInstanceTypeOk returns a tuple with the WorkerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyTorchOptionsAllOf) GetWorkerInstanceTypeOk() (*string, bool) {
	if o == nil || o.WorkerInstanceType == nil {
		return nil, false
	}
	return o.WorkerInstanceType, true
}

// HasWorkerInstanceType returns a boolean if a field has been set.
func (o *PyTorchOptionsAllOf) HasWorkerInstanceType() bool {
	if o != nil && o.WorkerInstanceType != nil {
		return true
	}

	return false
}

// SetWorkerInstanceType gets a reference to the given string and assigns it to the WorkerInstanceType field.
func (o *PyTorchOptionsAllOf) SetWorkerInstanceType(v string) {
	o.WorkerInstanceType = &v
}

func (o PyTorchOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Module != nil {
		toSerialize["module"] = o.Module
	}
	if o.NumberReplicas != nil {
		toSerialize["number_replicas"] = o.NumberReplicas
	}
	if o.MasterInstanceType != nil {
		toSerialize["master_instance_type"] = o.MasterInstanceType
	}
	if o.WorkerInstanceType != nil {
		toSerialize["worker_instance_type"] = o.WorkerInstanceType
	}
	return json.Marshal(toSerialize)
}

type NullablePyTorchOptionsAllOf struct {
	value *PyTorchOptionsAllOf
	isSet bool
}

func (v NullablePyTorchOptionsAllOf) Get() *PyTorchOptionsAllOf {
	return v.value
}

func (v *NullablePyTorchOptionsAllOf) Set(val *PyTorchOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePyTorchOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePyTorchOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePyTorchOptionsAllOf(val *PyTorchOptionsAllOf) *NullablePyTorchOptionsAllOf {
	return &NullablePyTorchOptionsAllOf{value: val, isSet: true}
}

func (v NullablePyTorchOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePyTorchOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


