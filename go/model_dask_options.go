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

// checks if the DaskOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaskOptions{}

// DaskOptions struct for DaskOptions
type DaskOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	PackageName *string `json:"package_name,omitempty"`
	ModuleName *string `json:"module_name,omitempty"`
	FunctionName *string `json:"function_name,omitempty"`
	NumberWorkers *float32 `json:"number_workers,omitempty"`
	NumberWorkersMax *float32 `json:"number_workers_max,omitempty"`
	DriverInstanceType *string `json:"driver_instance_type,omitempty"`
	WorkerInstanceType *string `json:"worker_instance_type,omitempty"`
}

// NewDaskOptions instantiates a new DaskOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaskOptions() *DaskOptions {
	this := DaskOptions{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// NewDaskOptionsWithDefaults instantiates a new DaskOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaskOptionsWithDefaults() *DaskOptions {
	this := DaskOptions{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DaskOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DaskOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DaskOptions) SetType(v string) {
	o.Type = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *DaskOptions) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *DaskOptions) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *DaskOptions) SetPackageName(v string) {
	o.PackageName = &v
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise.
func (o *DaskOptions) GetModuleName() string {
	if o == nil || IsNil(o.ModuleName) {
		var ret string
		return ret
	}
	return *o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetModuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleName) {
		return nil, false
	}
	return o.ModuleName, true
}

// HasModuleName returns a boolean if a field has been set.
func (o *DaskOptions) HasModuleName() bool {
	if o != nil && !IsNil(o.ModuleName) {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given string and assigns it to the ModuleName field.
func (o *DaskOptions) SetModuleName(v string) {
	o.ModuleName = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *DaskOptions) GetFunctionName() string {
	if o == nil || IsNil(o.FunctionName) {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetFunctionNameOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionName) {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *DaskOptions) HasFunctionName() bool {
	if o != nil && !IsNil(o.FunctionName) {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *DaskOptions) SetFunctionName(v string) {
	o.FunctionName = &v
}

// GetNumberWorkers returns the NumberWorkers field value if set, zero value otherwise.
func (o *DaskOptions) GetNumberWorkers() float32 {
	if o == nil || IsNil(o.NumberWorkers) {
		var ret float32
		return ret
	}
	return *o.NumberWorkers
}

// GetNumberWorkersOk returns a tuple with the NumberWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetNumberWorkersOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberWorkers) {
		return nil, false
	}
	return o.NumberWorkers, true
}

// HasNumberWorkers returns a boolean if a field has been set.
func (o *DaskOptions) HasNumberWorkers() bool {
	if o != nil && !IsNil(o.NumberWorkers) {
		return true
	}

	return false
}

// SetNumberWorkers gets a reference to the given float32 and assigns it to the NumberWorkers field.
func (o *DaskOptions) SetNumberWorkers(v float32) {
	o.NumberWorkers = &v
}

// GetNumberWorkersMax returns the NumberWorkersMax field value if set, zero value otherwise.
func (o *DaskOptions) GetNumberWorkersMax() float32 {
	if o == nil || IsNil(o.NumberWorkersMax) {
		var ret float32
		return ret
	}
	return *o.NumberWorkersMax
}

// GetNumberWorkersMaxOk returns a tuple with the NumberWorkersMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetNumberWorkersMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberWorkersMax) {
		return nil, false
	}
	return o.NumberWorkersMax, true
}

// HasNumberWorkersMax returns a boolean if a field has been set.
func (o *DaskOptions) HasNumberWorkersMax() bool {
	if o != nil && !IsNil(o.NumberWorkersMax) {
		return true
	}

	return false
}

// SetNumberWorkersMax gets a reference to the given float32 and assigns it to the NumberWorkersMax field.
func (o *DaskOptions) SetNumberWorkersMax(v float32) {
	o.NumberWorkersMax = &v
}

// GetDriverInstanceType returns the DriverInstanceType field value if set, zero value otherwise.
func (o *DaskOptions) GetDriverInstanceType() string {
	if o == nil || IsNil(o.DriverInstanceType) {
		var ret string
		return ret
	}
	return *o.DriverInstanceType
}

// GetDriverInstanceTypeOk returns a tuple with the DriverInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetDriverInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DriverInstanceType) {
		return nil, false
	}
	return o.DriverInstanceType, true
}

// HasDriverInstanceType returns a boolean if a field has been set.
func (o *DaskOptions) HasDriverInstanceType() bool {
	if o != nil && !IsNil(o.DriverInstanceType) {
		return true
	}

	return false
}

// SetDriverInstanceType gets a reference to the given string and assigns it to the DriverInstanceType field.
func (o *DaskOptions) SetDriverInstanceType(v string) {
	o.DriverInstanceType = &v
}

// GetWorkerInstanceType returns the WorkerInstanceType field value if set, zero value otherwise.
func (o *DaskOptions) GetWorkerInstanceType() string {
	if o == nil || IsNil(o.WorkerInstanceType) {
		var ret string
		return ret
	}
	return *o.WorkerInstanceType
}

// GetWorkerInstanceTypeOk returns a tuple with the WorkerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaskOptions) GetWorkerInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkerInstanceType) {
		return nil, false
	}
	return o.WorkerInstanceType, true
}

// HasWorkerInstanceType returns a boolean if a field has been set.
func (o *DaskOptions) HasWorkerInstanceType() bool {
	if o != nil && !IsNil(o.WorkerInstanceType) {
		return true
	}

	return false
}

// SetWorkerInstanceType gets a reference to the given string and assigns it to the WorkerInstanceType field.
func (o *DaskOptions) SetWorkerInstanceType(v string) {
	o.WorkerInstanceType = &v
}

func (o DaskOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaskOptions) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PackageName) {
		toSerialize["package_name"] = o.PackageName
	}
	if !IsNil(o.ModuleName) {
		toSerialize["module_name"] = o.ModuleName
	}
	if !IsNil(o.FunctionName) {
		toSerialize["function_name"] = o.FunctionName
	}
	if !IsNil(o.NumberWorkers) {
		toSerialize["number_workers"] = o.NumberWorkers
	}
	if !IsNil(o.NumberWorkersMax) {
		toSerialize["number_workers_max"] = o.NumberWorkersMax
	}
	if !IsNil(o.DriverInstanceType) {
		toSerialize["driver_instance_type"] = o.DriverInstanceType
	}
	if !IsNil(o.WorkerInstanceType) {
		toSerialize["worker_instance_type"] = o.WorkerInstanceType
	}
	return toSerialize, nil
}

type NullableDaskOptions struct {
	value *DaskOptions
	isSet bool
}

func (v NullableDaskOptions) Get() *DaskOptions {
	return v.value
}

func (v *NullableDaskOptions) Set(val *DaskOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDaskOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDaskOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaskOptions(val *DaskOptions) *NullableDaskOptions {
	return &NullableDaskOptions{value: val, isSet: true}
}

func (v NullableDaskOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaskOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


