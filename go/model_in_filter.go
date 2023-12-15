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

// checks if the InFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InFilter{}

// InFilter Filter values that are in a list or a string.  Attributes ---------- type : RelationalOperatorType     Type of the filter.
type InFilter struct {
	// Left operator for filter.
	Left string `json:"left"`
	Right Right `json:"right"`
	// Operator for IN list/range filter.
	Type *string `json:"type,omitempty"`
}

type _InFilter InFilter

// NewInFilter instantiates a new InFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInFilter(left string, right Right) *InFilter {
	this := InFilter{}
	this.Left = left
	this.Right = right
	var type_ string = "in"
	this.Type = &type_
	return &this
}

// NewInFilterWithDefaults instantiates a new InFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInFilterWithDefaults() *InFilter {
	this := InFilter{}
	var type_ string = "in"
	this.Type = &type_
	return &this
}

// GetLeft returns the Left field value
func (o *InFilter) GetLeft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *InFilter) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *InFilter) SetLeft(v string) {
	o.Left = v
}

// GetRight returns the Right field value
func (o *InFilter) GetRight() Right {
	if o == nil {
		var ret Right
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *InFilter) GetRightOk() (*Right, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *InFilter) SetRight(v Right) {
	o.Right = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InFilter) SetType(v string) {
	o.Type = &v
}

func (o InFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["left"] = o.Left
	toSerialize["right"] = o.Right
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *InFilter) UnmarshalJSON(bytes []byte) (err error) {
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

	varInFilter := _InFilter{}

	err = json.Unmarshal(bytes, &varInFilter)

	if err != nil {
		return err
	}

	*o = InFilter(varInFilter)

	return err
}

type NullableInFilter struct {
	value *InFilter
	isSet bool
}

func (v NullableInFilter) Get() *InFilter {
	return v.value
}

func (v *NullableInFilter) Set(val *InFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableInFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableInFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInFilter(val *InFilter) *NullableInFilter {
	return &NullableInFilter{value: val, isSet: true}
}

func (v NullableInFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


