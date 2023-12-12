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

// checks if the AggParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggParams{}

// AggParams struct for AggParams
type AggParams struct {
	// Column(s) to use to group the dataset.
	GroupColumns []string `json:"group_columns"`
	// List of aggregations to compute.
	Aggregations []AggregationsInner `json:"aggregations"`
}

type _AggParams AggParams

// NewAggParams instantiates a new AggParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggParams(groupColumns []string, aggregations []AggregationsInner) *AggParams {
	this := AggParams{}
	this.GroupColumns = groupColumns
	this.Aggregations = aggregations
	return &this
}

// NewAggParamsWithDefaults instantiates a new AggParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggParamsWithDefaults() *AggParams {
	this := AggParams{}
	return &this
}

// GetGroupColumns returns the GroupColumns field value
func (o *AggParams) GetGroupColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupColumns
}

// GetGroupColumnsOk returns a tuple with the GroupColumns field value
// and a boolean to check if the value has been set.
func (o *AggParams) GetGroupColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupColumns, true
}

// SetGroupColumns sets field value
func (o *AggParams) SetGroupColumns(v []string) {
	o.GroupColumns = v
}

// GetAggregations returns the Aggregations field value
func (o *AggParams) GetAggregations() []AggregationsInner {
	if o == nil {
		var ret []AggregationsInner
		return ret
	}

	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value
// and a boolean to check if the value has been set.
func (o *AggParams) GetAggregationsOk() ([]AggregationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// SetAggregations sets field value
func (o *AggParams) SetAggregations(v []AggregationsInner) {
	o.Aggregations = v
}

func (o AggParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_columns"] = o.GroupColumns
	toSerialize["aggregations"] = o.Aggregations
	return toSerialize, nil
}

func (o *AggParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_columns",
		"aggregations",
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

	varAggParams := _AggParams{}

	err = json.Unmarshal(bytes, &varAggParams)

	if err != nil {
		return err
	}

	*o = AggParams(varAggParams)

	return err
}

type NullableAggParams struct {
	value *AggParams
	isSet bool
}

func (v NullableAggParams) Get() *AggParams {
	return v.value
}

func (v *NullableAggParams) Set(val *AggParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAggParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAggParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggParams(val *AggParams) *NullableAggParams {
	return &NullableAggParams{value: val, isSet: true}
}

func (v NullableAggParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

