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

// checks if the ServicesRuntimes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicesRuntimes{}

// ServicesRuntimes struct for ServicesRuntimes
type ServicesRuntimes struct {
	Spark *bool `json:"spark,omitempty"`
	Mxnet *bool `json:"mxnet,omitempty"`
	Tensorflow *bool `json:"tensorflow,omitempty"`
	Xgboost *bool `json:"xgboost,omitempty"`
	Mpi map[string]interface{} `json:"mpi,omitempty"`
	Pytorch map[string]interface{} `json:"pytorch,omitempty"`
}

// NewServicesRuntimes instantiates a new ServicesRuntimes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesRuntimes() *ServicesRuntimes {
	this := ServicesRuntimes{}
	return &this
}

// NewServicesRuntimesWithDefaults instantiates a new ServicesRuntimes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesRuntimesWithDefaults() *ServicesRuntimes {
	this := ServicesRuntimes{}
	return &this
}

// GetSpark returns the Spark field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetSpark() bool {
	if o == nil || IsNil(o.Spark) {
		var ret bool
		return ret
	}
	return *o.Spark
}

// GetSparkOk returns a tuple with the Spark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetSparkOk() (*bool, bool) {
	if o == nil || IsNil(o.Spark) {
		return nil, false
	}
	return o.Spark, true
}

// HasSpark returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasSpark() bool {
	if o != nil && !IsNil(o.Spark) {
		return true
	}

	return false
}

// SetSpark gets a reference to the given bool and assigns it to the Spark field.
func (o *ServicesRuntimes) SetSpark(v bool) {
	o.Spark = &v
}

// GetMxnet returns the Mxnet field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetMxnet() bool {
	if o == nil || IsNil(o.Mxnet) {
		var ret bool
		return ret
	}
	return *o.Mxnet
}

// GetMxnetOk returns a tuple with the Mxnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetMxnetOk() (*bool, bool) {
	if o == nil || IsNil(o.Mxnet) {
		return nil, false
	}
	return o.Mxnet, true
}

// HasMxnet returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasMxnet() bool {
	if o != nil && !IsNil(o.Mxnet) {
		return true
	}

	return false
}

// SetMxnet gets a reference to the given bool and assigns it to the Mxnet field.
func (o *ServicesRuntimes) SetMxnet(v bool) {
	o.Mxnet = &v
}

// GetTensorflow returns the Tensorflow field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetTensorflow() bool {
	if o == nil || IsNil(o.Tensorflow) {
		var ret bool
		return ret
	}
	return *o.Tensorflow
}

// GetTensorflowOk returns a tuple with the Tensorflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetTensorflowOk() (*bool, bool) {
	if o == nil || IsNil(o.Tensorflow) {
		return nil, false
	}
	return o.Tensorflow, true
}

// HasTensorflow returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasTensorflow() bool {
	if o != nil && !IsNil(o.Tensorflow) {
		return true
	}

	return false
}

// SetTensorflow gets a reference to the given bool and assigns it to the Tensorflow field.
func (o *ServicesRuntimes) SetTensorflow(v bool) {
	o.Tensorflow = &v
}

// GetXgboost returns the Xgboost field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetXgboost() bool {
	if o == nil || IsNil(o.Xgboost) {
		var ret bool
		return ret
	}
	return *o.Xgboost
}

// GetXgboostOk returns a tuple with the Xgboost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetXgboostOk() (*bool, bool) {
	if o == nil || IsNil(o.Xgboost) {
		return nil, false
	}
	return o.Xgboost, true
}

// HasXgboost returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasXgboost() bool {
	if o != nil && !IsNil(o.Xgboost) {
		return true
	}

	return false
}

// SetXgboost gets a reference to the given bool and assigns it to the Xgboost field.
func (o *ServicesRuntimes) SetXgboost(v bool) {
	o.Xgboost = &v
}

// GetMpi returns the Mpi field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetMpi() map[string]interface{} {
	if o == nil || IsNil(o.Mpi) {
		var ret map[string]interface{}
		return ret
	}
	return o.Mpi
}

// GetMpiOk returns a tuple with the Mpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetMpiOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Mpi) {
		return map[string]interface{}{}, false
	}
	return o.Mpi, true
}

// HasMpi returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasMpi() bool {
	if o != nil && !IsNil(o.Mpi) {
		return true
	}

	return false
}

// SetMpi gets a reference to the given map[string]interface{} and assigns it to the Mpi field.
func (o *ServicesRuntimes) SetMpi(v map[string]interface{}) {
	o.Mpi = v
}

// GetPytorch returns the Pytorch field value if set, zero value otherwise.
func (o *ServicesRuntimes) GetPytorch() map[string]interface{} {
	if o == nil || IsNil(o.Pytorch) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pytorch
}

// GetPytorchOk returns a tuple with the Pytorch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRuntimes) GetPytorchOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Pytorch) {
		return map[string]interface{}{}, false
	}
	return o.Pytorch, true
}

// HasPytorch returns a boolean if a field has been set.
func (o *ServicesRuntimes) HasPytorch() bool {
	if o != nil && !IsNil(o.Pytorch) {
		return true
	}

	return false
}

// SetPytorch gets a reference to the given map[string]interface{} and assigns it to the Pytorch field.
func (o *ServicesRuntimes) SetPytorch(v map[string]interface{}) {
	o.Pytorch = v
}

func (o ServicesRuntimes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicesRuntimes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Spark) {
		toSerialize["spark"] = o.Spark
	}
	if !IsNil(o.Mxnet) {
		toSerialize["mxnet"] = o.Mxnet
	}
	if !IsNil(o.Tensorflow) {
		toSerialize["tensorflow"] = o.Tensorflow
	}
	if !IsNil(o.Xgboost) {
		toSerialize["xgboost"] = o.Xgboost
	}
	if !IsNil(o.Mpi) {
		toSerialize["mpi"] = o.Mpi
	}
	if !IsNil(o.Pytorch) {
		toSerialize["pytorch"] = o.Pytorch
	}
	return toSerialize, nil
}

type NullableServicesRuntimes struct {
	value *ServicesRuntimes
	isSet bool
}

func (v NullableServicesRuntimes) Get() *ServicesRuntimes {
	return v.value
}

func (v *NullableServicesRuntimes) Set(val *ServicesRuntimes) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesRuntimes) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesRuntimes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesRuntimes(val *ServicesRuntimes) *NullableServicesRuntimes {
	return &NullableServicesRuntimes{value: val, isSet: true}
}

func (v NullableServicesRuntimes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesRuntimes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


