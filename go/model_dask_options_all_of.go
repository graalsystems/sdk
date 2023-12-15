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

// DaskOptionsAllOf struct for DaskOptionsAllOf
type DaskOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	PackageName *string `json:"package_name,omitempty"`
	ModuleName *string `json:"module_name,omitempty"`
	FunctionName *string `json:"function_name,omitempty"`
	NumberWorkers *float32 `json:"number_workers,omitempty"`
	NumberWorkersMax *float32 `json:"number_workers_max,omitempty"`
	DriverInstanceType *string `json:"driver_instance_type,omitempty"`
	WorkerInstanceType *string `json:"worker_instance_type,omitempty"`
}

// NewDaskOptionsAllOf instantiates a new DaskOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaskOptionsAllOf() *DaskOptionsAllOf {
	this := DaskOptionsAllOf{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// NewDaskOptionsAllOfWithDefaults instantiates a new DaskOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaskOptionsAllOfWithDefaults() *DaskOptionsAllOf {
	this := DaskOptionsAllOf{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DaskOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *DaskOptionsAllOf) SetPackageName(v string) {
	o.PackageName = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetModuleName() string {
	if o == nil || o.ModuleName == nil {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetModuleNameOk() (*string, bool) {
	if o == nil || o.ModuleName == nil {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasModuleName() bool {
	if o != nil && o.ModuleName != nil {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *DaskOptionsAllOf) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetFunctionName() string {
	if o == nil || o.FunctionName == nil {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetFunctionNameOk() (*string, bool) {
	if o == nil || o.FunctionName == nil {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasFunctionName() bool {
	if o != nil && o.FunctionName != nil {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *DaskOptionsAllOf) SetFunctionName(v string) {
	o.FunctionName = &v
}

// GetNumberWorkers returns the NumberWorkers field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetNumberWorkers() float32 {
	if o == nil || o.NumberWorkers == nil {
		var ret float32
		return ret
	}
	return *o.NumberWorkers
}

// GetNumberWorkersOk returns a tuple with the NumberWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetNumberWorkersOk() (*float32, bool) {
	if o == nil || o.NumberWorkers == nil {
		return nil, false
	}
	return o.NumberWorkers, true
}

// HasNumberWorkers returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasNumberWorkers() bool {
	if o != nil && o.NumberWorkers != nil {
		return true
	}

	return false
}

// SetNumberWorkers gets a reference to the given float32 and assigns it to the NumberWorkers field.
func (o *DaskOptionsAllOf) SetNumberWorkers(v float32) {
	o.NumberWorkers = &v
}

// GetNumberWorkersMax returns the NumberWorkersMax field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetNumberWorkersMax() float32 {
	if o == nil || o.NumberWorkersMax == nil {
		var ret float32
		return ret
	}
	return *o.NumberWorkersMax
}

// GetNumberWorkersMaxOk returns a tuple with the NumberWorkersMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetNumberWorkersMaxOk() (*float32, bool) {
	if o == nil || o.NumberWorkersMax == nil {
		return nil, false
	}
	return o.NumberWorkersMax, true
}

// HasNumberWorkersMax returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasNumberWorkersMax() bool {
	if o != nil && o.NumberWorkersMax != nil {
		return true
	}

	return false
}

// SetNumberWorkersMax gets a reference to the given float32 and assigns it to the NumberWorkersMax field.
func (o *DaskOptionsAllOf) SetNumberWorkersMax(v float32) {
	o.NumberWorkersMax = &v
}

// GetDriverInstanceType returns the DriverInstanceType field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetDriverInstanceType() string {
	if o == nil || o.DriverInstanceType == nil {
		var ret string
		return ret
	}
	return *o.DriverInstanceType
}

// GetDriverInstanceTypeOk returns a tuple with the DriverInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetDriverInstanceTypeOk() (*string, bool) {
	if o == nil || o.DriverInstanceType == nil {
		return nil, false
	}
	return o.DriverInstanceType, true
}

// HasDriverInstanceType returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasDriverInstanceType() bool {
	if o != nil && o.DriverInstanceType != nil {
		return true
	}

	return false
}

// SetDriverInstanceType gets a reference to the given string and assigns it to the DriverInstanceType field.
func (o *DaskOptionsAllOf) SetDriverInstanceType(v string) {
	o.DriverInstanceType = &v
}

// GetWorkerInstanceType returns the WorkerInstanceType field value if set, zero value otherwise.
func (o *DaskOptionsAllOf) GetWorkerInstanceType() string {
	if o == nil || o.WorkerInstanceType == nil {
		var ret string
		return ret
	}
	return *o.WorkerInstanceType
}

// GetWorkerInstanceTypeOk returns a tuple with the WorkerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptionsAllOf) GetWorkerInstanceTypeOk() (*string, bool) {
	if o == nil || o.WorkerInstanceType == nil {
		return nil, false
	}
	return o.WorkerInstanceType, true
}

// HasWorkerInstanceType returns a boolean if a field has been set.
func (o *DaskOptionsAllOf) HasWorkerInstanceType() bool {
	if o != nil && o.WorkerInstanceType != nil {
		return true
	}

	return false
}

// SetWorkerInstanceType gets a reference to the given string and assigns it to the WorkerInstanceType field.
func (o *DaskOptionsAllOf) SetWorkerInstanceType(v string) {
	o.WorkerInstanceType = &v
}

func (o DaskOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.PackageName != nil {
		toSerialize["package_name"] = o.PackageName
	}
	if o.ModuleName != nil {
		toSerialize["module_name"] = o.ModuleName
	}
	if o.FunctionName != nil {
		toSerialize["function_name"] = o.FunctionName
	}
	if o.NumberWorkers != nil {
		toSerialize["number_workers"] = o.NumberWorkers
	}
	if o.NumberWorkersMax != nil {
		toSerialize["number_workers_max"] = o.NumberWorkersMax
	}
	if o.DriverInstanceType != nil {
		toSerialize["driver_instance_type"] = o.DriverInstanceType
	}
	if o.WorkerInstanceType != nil {
		toSerialize["worker_instance_type"] = o.WorkerInstanceType
	}
	return json.Marshal(toSerialize)
}

type NullableDaskOptionsAllOf struct {
	value *DaskOptionsAllOf
	isSet bool
}

func (v NullableDaskOptionsAllOf) Get() *DaskOptionsAllOf {
	return v.value
}

func (v *NullableDaskOptionsAllOf) Set(val *DaskOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDaskOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDaskOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaskOptionsAllOf(val *DaskOptionsAllOf) *NullableDaskOptionsAllOf {
	return &NullableDaskOptionsAllOf{value: val, isSet: true}
}

func (v NullableDaskOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaskOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


