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

// MXNetOptionsAllOf struct for MXNetOptionsAllOf
type MXNetOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	Module *string `json:"module,omitempty"`
	NumberReplicas *float32 `json:"number_replicas,omitempty"`
	ServerInstanceType *string `json:"server_instance_type,omitempty"`
	WorkerInstanceType *string `json:"worker_instance_type,omitempty"`
	SchedulerInstanceType *string `json:"scheduler_instance_type,omitempty"`
}

// NewMXNetOptionsAllOf instantiates a new MXNetOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMXNetOptionsAllOf() *MXNetOptionsAllOf {
	this := MXNetOptionsAllOf{}
	var type_ string = "mxnet"
	this.Type = &type_
	return &this
}

// NewMXNetOptionsAllOfWithDefaults instantiates a new MXNetOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMXNetOptionsAllOfWithDefaults() *MXNetOptionsAllOf {
	this := MXNetOptionsAllOf{}
	var type_ string = "mxnet"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MXNetOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetModule() string {
	if o == nil || o.Module == nil {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetModuleOk() (*string, bool) {
	if o == nil || o.Module == nil {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasModule() bool {
	if o != nil && o.Module != nil {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *MXNetOptionsAllOf) SetModule(v string) {
	o.Module = &v
}

// GetNumberReplicas returns the NumberReplicas field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetNumberReplicas() float32 {
	if o == nil || o.NumberReplicas == nil {
		var ret float32
		return ret
	}
	return *o.NumberReplicas
}

// GetNumberReplicasOk returns a tuple with the NumberReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetNumberReplicasOk() (*float32, bool) {
	if o == nil || o.NumberReplicas == nil {
		return nil, false
	}
	return o.NumberReplicas, true
}

// HasNumberReplicas returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasNumberReplicas() bool {
	if o != nil && o.NumberReplicas != nil {
		return true
	}

	return false
}

// SetNumberReplicas gets a reference to the given float32 and assigns it to the NumberReplicas field.
func (o *MXNetOptionsAllOf) SetNumberReplicas(v float32) {
	o.NumberReplicas = &v
}

// GetServerInstanceType returns the ServerInstanceType field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetServerInstanceType() string {
	if o == nil || o.ServerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.ServerInstanceType
}

// GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetServerInstanceTypeOk() (*string, bool) {
	if o == nil || o.ServerInstanceType == nil {
		return nil, false
	}
	return o.ServerInstanceType, true
}

// HasServerInstanceType returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasServerInstanceType() bool {
	if o != nil && o.ServerInstanceType != nil {
		return true
	}

	return false
}

// SetServerInstanceType gets a reference to the given string and assigns it to the ServerInstanceType field.
func (o *MXNetOptionsAllOf) SetServerInstanceType(v string) {
	o.ServerInstanceType = &v
}

// GetWorkerInstanceType returns the WorkerInstanceType field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetWorkerInstanceType() string {
	if o == nil || o.WorkerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.WorkerInstanceType
}

// GetWorkerInstanceTypeOk returns a tuple with the WorkerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetWorkerInstanceTypeOk() (*string, bool) {
	if o == nil || o.WorkerInstanceType == nil {
		return nil, false
	}
	return o.WorkerInstanceType, true
}

// HasWorkerInstanceType returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasWorkerInstanceType() bool {
	if o != nil && o.WorkerInstanceType != nil {
		return true
	}

	return false
}

// SetWorkerInstanceType gets a reference to the given string and assigns it to the WorkerInstanceType field.
func (o *MXNetOptionsAllOf) SetWorkerInstanceType(v string) {
	o.WorkerInstanceType = &v
}

// GetSchedulerInstanceType returns the SchedulerInstanceType field value if set, zero value otherwise.
func (o *MXNetOptionsAllOf) GetSchedulerInstanceType() string {
	if o == nil || o.SchedulerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.SchedulerInstanceType
}

// GetSchedulerInstanceTypeOk returns a tuple with the SchedulerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MXNetOptionsAllOf) GetSchedulerInstanceTypeOk() (*string, bool) {
	if o == nil || o.SchedulerInstanceType == nil {
		return nil, false
	}
	return o.SchedulerInstanceType, true
}

// HasSchedulerInstanceType returns a boolean if a field has been set.
func (o *MXNetOptionsAllOf) HasSchedulerInstanceType() bool {
	if o != nil && o.SchedulerInstanceType != nil {
		return true
	}

	return false
}

// SetSchedulerInstanceType gets a reference to the given string and assigns it to the SchedulerInstanceType field.
func (o *MXNetOptionsAllOf) SetSchedulerInstanceType(v string) {
	o.SchedulerInstanceType = &v
}

func (o MXNetOptionsAllOf) MarshalJSON() ([]byte, error) {
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
	if o.ServerInstanceType != nil {
		toSerialize["server_instance_type"] = o.ServerInstanceType
	}
	if o.WorkerInstanceType != nil {
		toSerialize["worker_instance_type"] = o.WorkerInstanceType
	}
	if o.SchedulerInstanceType != nil {
		toSerialize["scheduler_instance_type"] = o.SchedulerInstanceType
	}
	return json.Marshal(toSerialize)
}

type NullableMXNetOptionsAllOf struct {
	value *MXNetOptionsAllOf
	isSet bool
}

func (v NullableMXNetOptionsAllOf) Get() *MXNetOptionsAllOf {
	return v.value
}

func (v *NullableMXNetOptionsAllOf) Set(val *MXNetOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMXNetOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMXNetOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMXNetOptionsAllOf(val *MXNetOptionsAllOf) *NullableMXNetOptionsAllOf {
	return &NullableMXNetOptionsAllOf{value: val, isSet: true}
}

func (v NullableMXNetOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMXNetOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


