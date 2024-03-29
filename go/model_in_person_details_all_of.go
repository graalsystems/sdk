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

// InPersonDetailsAllOf struct for InPersonDetailsAllOf
type InPersonDetailsAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewInPersonDetailsAllOf instantiates a new InPersonDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInPersonDetailsAllOf() *InPersonDetailsAllOf {
	this := InPersonDetailsAllOf{}
	var type_ string = "in_person"
	this.Type = &type_
	return &this
}

// NewInPersonDetailsAllOfWithDefaults instantiates a new InPersonDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInPersonDetailsAllOfWithDefaults() *InPersonDetailsAllOf {
	this := InPersonDetailsAllOf{}
	var type_ string = "in_person"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InPersonDetailsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InPersonDetailsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InPersonDetailsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InPersonDetailsAllOf) SetType(v string) {
	o.Type = &v
}

func (o InPersonDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInPersonDetailsAllOf struct {
	value *InPersonDetailsAllOf
	isSet bool
}

func (v NullableInPersonDetailsAllOf) Get() *InPersonDetailsAllOf {
	return v.value
}

func (v *NullableInPersonDetailsAllOf) Set(val *InPersonDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInPersonDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInPersonDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInPersonDetailsAllOf(val *InPersonDetailsAllOf) *NullableInPersonDetailsAllOf {
	return &NullableInPersonDetailsAllOf{value: val, isSet: true}
}

func (v NullableInPersonDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInPersonDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


