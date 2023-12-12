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

// checks if the DistinctParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistinctParams{}

// DistinctParams struct for DistinctParams
type DistinctParams struct {
	// Return dataframe with duplicate rows removed. Certain                 columns can be used for identifying duplicates, by default use                 all of the columns.
	Columns []string `json:"columns,omitempty"`
}

// NewDistinctParams instantiates a new DistinctParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistinctParams() *DistinctParams {
	this := DistinctParams{}
	return &this
}

// NewDistinctParamsWithDefaults instantiates a new DistinctParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistinctParamsWithDefaults() *DistinctParams {
	this := DistinctParams{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DistinctParams) GetColumns() []string {
	if o == nil || IsNil(o.Columns) {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistinctParams) GetColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DistinctParams) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *DistinctParams) SetColumns(v []string) {
	o.Columns = v
}

func (o DistinctParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistinctParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableDistinctParams struct {
	value *DistinctParams
	isSet bool
}

func (v NullableDistinctParams) Get() *DistinctParams {
	return v.value
}

func (v *NullableDistinctParams) Set(val *DistinctParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDistinctParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDistinctParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistinctParams(val *DistinctParams) *NullableDistinctParams {
	return &NullableDistinctParams{value: val, isSet: true}
}

func (v NullableDistinctParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistinctParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


