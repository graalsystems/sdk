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
	"fmt"
)

// checks if the GreaterThanFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GreaterThanFilter{}

// GreaterThanFilter Filter values that are greater than a given value.  Attributes ---------- type : RelationalOperatorType     Type of the filter.
type GreaterThanFilter struct {
	// Left operator for filter.
	Left string `json:"left"`
	// Right operator for filter.
	Right string `json:"right"`
	// Operator for GREATER THAN filter.
	Type *string `json:"type,omitempty"`
}

type _GreaterThanFilter GreaterThanFilter

// NewGreaterThanFilter instantiates a new GreaterThanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterThanFilter(left string, right string) *GreaterThanFilter {
	this := GreaterThanFilter{}
	this.Left = left
	this.Right = right
	var type_ string = "greater"
	this.Type = &type_
	return &this
}

// NewGreaterThanFilterWithDefaults instantiates a new GreaterThanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterThanFilterWithDefaults() *GreaterThanFilter {
	this := GreaterThanFilter{}
	var type_ string = "greater"
	this.Type = &type_
	return &this
}

// GetLeft returns the Left field value
func (o *GreaterThanFilter) GetLeft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *GreaterThanFilter) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *GreaterThanFilter) SetLeft(v string) {
	o.Left = v
}

// GetRight returns the Right field value
func (o *GreaterThanFilter) GetRight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *GreaterThanFilter) GetRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *GreaterThanFilter) SetRight(v string) {
	o.Right = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GreaterThanFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GreaterThanFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GreaterThanFilter) SetType(v string) {
	o.Type = &v
}

func (o GreaterThanFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreaterThanFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["left"] = o.Left
	toSerialize["right"] = o.Right
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *GreaterThanFilter) UnmarshalJSON(bytes []byte) (err error) {
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

	varGreaterThanFilter := _GreaterThanFilter{}

	err = json.Unmarshal(bytes, &varGreaterThanFilter)

	if err != nil {
		return err
	}

	*o = GreaterThanFilter(varGreaterThanFilter)

	return err
}

type NullableGreaterThanFilter struct {
	value *GreaterThanFilter
	isSet bool
}

func (v NullableGreaterThanFilter) Get() *GreaterThanFilter {
	return v.value
}

func (v *NullableGreaterThanFilter) Set(val *GreaterThanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterThanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterThanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterThanFilter(val *GreaterThanFilter) *NullableGreaterThanFilter {
	return &NullableGreaterThanFilter{value: val, isSet: true}
}

func (v NullableGreaterThanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterThanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


