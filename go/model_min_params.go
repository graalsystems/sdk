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
	"fmt"
)

// checks if the MinParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinParams{}

// MinParams struct for MinParams
type MinParams struct {
	// Minimum aggregation.
	Type *string `json:"type,omitempty"`
	// Column(s) to compute the minimum.
	Columns []string `json:"columns"`
}

type _MinParams MinParams

// NewMinParams instantiates a new MinParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinParams(columns []string) *MinParams {
	this := MinParams{}
	var type_ string = "min"
	this.Type = &type_
	this.Columns = columns
	return &this
}

// NewMinParamsWithDefaults instantiates a new MinParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinParamsWithDefaults() *MinParams {
	this := MinParams{}
	var type_ string = "min"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MinParams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinParams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MinParams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MinParams) SetType(v string) {
	o.Type = &v
}

// GetColumns returns the Columns field value
func (o *MinParams) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *MinParams) GetColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *MinParams) SetColumns(v []string) {
	o.Columns = v
}

func (o MinParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["columns"] = o.Columns
	return toSerialize, nil
}

func (o *MinParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"columns",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMinParams := _MinParams{}

	err = json.Unmarshal(bytes, &varMinParams)

	if err != nil {
		return err
	}

	*o = MinParams(varMinParams)

	return err
}

type NullableMinParams struct {
	value *MinParams
	isSet bool
}

func (v NullableMinParams) Get() *MinParams {
	return v.value
}

func (v *NullableMinParams) Set(val *MinParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMinParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMinParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinParams(val *MinParams) *NullableMinParams {
	return &NullableMinParams{value: val, isSet: true}
}

func (v NullableMinParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


