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

// checks if the OrFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrFilter{}

// OrFilter Filter values that verify at least one condition.  Attributes ---------- type : RelationalOperatorType     Type of the filter. conditions : union_filter     Union of all relational filters.
type OrFilter struct {
	// Operator for OR filter.
	Type *string `json:"type,omitempty"`
	// Condition for OR filter.
	Conditions []ANDConditionInner `json:"conditions"`
}

type _OrFilter OrFilter

// NewOrFilter instantiates a new OrFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrFilter(conditions []ANDConditionInner) *OrFilter {
	this := OrFilter{}
	var type_ string = "or"
	this.Type = &type_
	this.Conditions = conditions
	return &this
}

// NewOrFilterWithDefaults instantiates a new OrFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrFilterWithDefaults() *OrFilter {
	this := OrFilter{}
	var type_ string = "or"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrFilter) SetType(v string) {
	o.Type = &v
}

// GetConditions returns the Conditions field value
func (o *OrFilter) GetConditions() []ANDConditionInner {
	if o == nil {
		var ret []ANDConditionInner
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *OrFilter) GetConditionsOk() ([]ANDConditionInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *OrFilter) SetConditions(v []ANDConditionInner) {
	o.Conditions = v
}

func (o OrFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["conditions"] = o.Conditions
	return toSerialize, nil
}

func (o *OrFilter) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conditions",
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

	varOrFilter := _OrFilter{}

	err = json.Unmarshal(bytes, &varOrFilter)

	if err != nil {
		return err
	}

	*o = OrFilter(varOrFilter)

	return err
}

type NullableOrFilter struct {
	value *OrFilter
	isSet bool
}

func (v NullableOrFilter) Get() *OrFilter {
	return v.value
}

func (v *NullableOrFilter) Set(val *OrFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrFilter(val *OrFilter) *NullableOrFilter {
	return &NullableOrFilter{value: val, isSet: true}
}

func (v NullableOrFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

