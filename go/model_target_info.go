/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

import (
	"encoding/json"
)

// TargetInfo struct for TargetInfo
type TargetInfo struct {
	TotalNodes *Capacity `json:"total_nodes,omitempty"`
	TotalCpus *Capacity `json:"total_cpus,omitempty"`
	TotalMemory *ByteCapacity `json:"total_memory,omitempty"`
	AvailabilityZones *[]string `json:"availability_zones,omitempty"`
}

// NewTargetInfo instantiates a new TargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetInfo() *TargetInfo {
	this := TargetInfo{}
	return &this
}

// NewTargetInfoWithDefaults instantiates a new TargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetInfoWithDefaults() *TargetInfo {
	this := TargetInfo{}
	return &this
}

// GetTotalNodes returns the TotalNodes field value if set, zero value otherwise.
func (o *TargetInfo) GetTotalNodes() Capacity {
	if o == nil || o.TotalNodes == nil {
		var ret Capacity
		return ret
	}
	return *o.TotalNodes
}

// GetTotalNodesOk returns a tuple with the TotalNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetTotalNodesOk() (*Capacity, bool) {
	if o == nil || o.TotalNodes == nil {
		return nil, false
	}
	return o.TotalNodes, true
}

// HasTotalNodes returns a boolean if a field has been set.
func (o *TargetInfo) HasTotalNodes() bool {
	if o != nil && o.TotalNodes != nil {
		return true
	}

	return false
}

// SetTotalNodes gets a reference to the given Capacity and assigns it to the TotalNodes field.
func (o *TargetInfo) SetTotalNodes(v Capacity) {
	o.TotalNodes = &v
}

// GetTotalCpus returns the TotalCpus field value if set, zero value otherwise.
func (o *TargetInfo) GetTotalCpus() Capacity {
	if o == nil || o.TotalCpus == nil {
		var ret Capacity
		return ret
	}
	return *o.TotalCpus
}

// GetTotalCpusOk returns a tuple with the TotalCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetTotalCpusOk() (*Capacity, bool) {
	if o == nil || o.TotalCpus == nil {
		return nil, false
	}
	return o.TotalCpus, true
}

// HasTotalCpus returns a boolean if a field has been set.
func (o *TargetInfo) HasTotalCpus() bool {
	if o != nil && o.TotalCpus != nil {
		return true
	}

	return false
}

// SetTotalCpus gets a reference to the given Capacity and assigns it to the TotalCpus field.
func (o *TargetInfo) SetTotalCpus(v Capacity) {
	o.TotalCpus = &v
}

// GetTotalMemory returns the TotalMemory field value if set, zero value otherwise.
func (o *TargetInfo) GetTotalMemory() ByteCapacity {
	if o == nil || o.TotalMemory == nil {
		var ret ByteCapacity
		return ret
	}
	return *o.TotalMemory
}

// GetTotalMemoryOk returns a tuple with the TotalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetTotalMemoryOk() (*ByteCapacity, bool) {
	if o == nil || o.TotalMemory == nil {
		return nil, false
	}
	return o.TotalMemory, true
}

// HasTotalMemory returns a boolean if a field has been set.
func (o *TargetInfo) HasTotalMemory() bool {
	if o != nil && o.TotalMemory != nil {
		return true
	}

	return false
}

// SetTotalMemory gets a reference to the given ByteCapacity and assigns it to the TotalMemory field.
func (o *TargetInfo) SetTotalMemory(v ByteCapacity) {
	o.TotalMemory = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *TargetInfo) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return *o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *TargetInfo) HasAvailabilityZones() bool {
	if o != nil && o.AvailabilityZones != nil {
		return true
	}

	return false
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *TargetInfo) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = &v
}

func (o TargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalNodes != nil {
		toSerialize["total_nodes"] = o.TotalNodes
	}
	if o.TotalCpus != nil {
		toSerialize["total_cpus"] = o.TotalCpus
	}
	if o.TotalMemory != nil {
		toSerialize["total_memory"] = o.TotalMemory
	}
	if o.AvailabilityZones != nil {
		toSerialize["availability_zones"] = o.AvailabilityZones
	}
	return json.Marshal(toSerialize)
}

type NullableTargetInfo struct {
	value *TargetInfo
	isSet bool
}

func (v NullableTargetInfo) Get() *TargetInfo {
	return v.value
}

func (v *NullableTargetInfo) Set(val *TargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetInfo(val *TargetInfo) *NullableTargetInfo {
	return &NullableTargetInfo{value: val, isSet: true}
}

func (v NullableTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


