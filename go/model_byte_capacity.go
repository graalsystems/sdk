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

// checks if the ByteCapacity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ByteCapacity{}

// ByteCapacity struct for ByteCapacity
type ByteCapacity struct {
	Min *string `json:"min,omitempty"`
	Max *string `json:"max,omitempty"`
	Current *string `json:"current,omitempty"`
}

// NewByteCapacity instantiates a new ByteCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewByteCapacity() *ByteCapacity {
	this := ByteCapacity{}
	return &this
}

// NewByteCapacityWithDefaults instantiates a new ByteCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewByteCapacityWithDefaults() *ByteCapacity {
	this := ByteCapacity{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ByteCapacity) GetMin() string {
	if o == nil || IsNil(o.Min) {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ByteCapacity) GetMinOk() (*string, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ByteCapacity) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *ByteCapacity) SetMin(v string) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ByteCapacity) GetMax() string {
	if o == nil || IsNil(o.Max) {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ByteCapacity) GetMaxOk() (*string, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ByteCapacity) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *ByteCapacity) SetMax(v string) {
	o.Max = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *ByteCapacity) GetCurrent() string {
	if o == nil || IsNil(o.Current) {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ByteCapacity) GetCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *ByteCapacity) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *ByteCapacity) SetCurrent(v string) {
	o.Current = &v
}

func (o ByteCapacity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ByteCapacity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	return toSerialize, nil
}

type NullableByteCapacity struct {
	value *ByteCapacity
	isSet bool
}

func (v NullableByteCapacity) Get() *ByteCapacity {
	return v.value
}

func (v *NullableByteCapacity) Set(val *ByteCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableByteCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableByteCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableByteCapacity(val *ByteCapacity) *NullableByteCapacity {
	return &NullableByteCapacity{value: val, isSet: true}
}

func (v NullableByteCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableByteCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


