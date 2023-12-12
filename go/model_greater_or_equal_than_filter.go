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

// checks if the GreaterOrEqualThanFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GreaterOrEqualThanFilter{}

// GreaterOrEqualThanFilter Filter values that are greater or equal to a given value.  Attributes ---------- type : RelationalOperatorType     Type of the filter.
type GreaterOrEqualThanFilter struct {
	// Left operator for filter.
	Left string `json:"left"`
	// Right operator for filter.
	Right string `json:"right"`
	// Operator for GREATER OR EQUAL filter.
	Type *string `json:"type,omitempty"`
}

type _GreaterOrEqualThanFilter GreaterOrEqualThanFilter

// NewGreaterOrEqualThanFilter instantiates a new GreaterOrEqualThanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterOrEqualThanFilter(left string, right string) *GreaterOrEqualThanFilter {
	this := GreaterOrEqualThanFilter{}
	this.Left = left
	this.Right = right
	var type_ string = "greater_or_equal"
	this.Type = &type_
	return &this
}

// NewGreaterOrEqualThanFilterWithDefaults instantiates a new GreaterOrEqualThanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterOrEqualThanFilterWithDefaults() *GreaterOrEqualThanFilter {
	this := GreaterOrEqualThanFilter{}
	var type_ string = "greater_or_equal"
	this.Type = &type_
	return &this
}

// GetLeft returns the Left field value
func (o *GreaterOrEqualThanFilter) GetLeft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *GreaterOrEqualThanFilter) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *GreaterOrEqualThanFilter) SetLeft(v string) {
	o.Left = v
}

// GetRight returns the Right field value
func (o *GreaterOrEqualThanFilter) GetRight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *GreaterOrEqualThanFilter) GetRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *GreaterOrEqualThanFilter) SetRight(v string) {
	o.Right = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GreaterOrEqualThanFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterOrEqualThanFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GreaterOrEqualThanFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GreaterOrEqualThanFilter) SetType(v string) {
	o.Type = &v
}

func (o GreaterOrEqualThanFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreaterOrEqualThanFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["left"] = o.Left
	toSerialize["right"] = o.Right
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *GreaterOrEqualThanFilter) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"left",
		"right",
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

	varGreaterOrEqualThanFilter := _GreaterOrEqualThanFilter{}

	err = json.Unmarshal(bytes, &varGreaterOrEqualThanFilter)

	if err != nil {
		return err
	}

	*o = GreaterOrEqualThanFilter(varGreaterOrEqualThanFilter)

	return err
}

type NullableGreaterOrEqualThanFilter struct {
	value *GreaterOrEqualThanFilter
	isSet bool
}

func (v NullableGreaterOrEqualThanFilter) Get() *GreaterOrEqualThanFilter {
	return v.value
}

func (v *NullableGreaterOrEqualThanFilter) Set(val *GreaterOrEqualThanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterOrEqualThanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterOrEqualThanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterOrEqualThanFilter(val *GreaterOrEqualThanFilter) *NullableGreaterOrEqualThanFilter {
	return &NullableGreaterOrEqualThanFilter{value: val, isSet: true}
}

func (v NullableGreaterOrEqualThanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterOrEqualThanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

