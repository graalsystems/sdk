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

// StartTaskAllOf struct for StartTaskAllOf
type StartTaskAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewStartTaskAllOf instantiates a new StartTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartTaskAllOf() *StartTaskAllOf {
	this := StartTaskAllOf{}
	var type_ string = "start"
	this.Type = &type_
	return &this
}

// NewStartTaskAllOfWithDefaults instantiates a new StartTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartTaskAllOfWithDefaults() *StartTaskAllOf {
	this := StartTaskAllOf{}
	var type_ string = "start"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StartTaskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartTaskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StartTaskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StartTaskAllOf) SetType(v string) {
	o.Type = &v
}

func (o StartTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableStartTaskAllOf struct {
	value *StartTaskAllOf
	isSet bool
}

func (v NullableStartTaskAllOf) Get() *StartTaskAllOf {
	return v.value
}

func (v *NullableStartTaskAllOf) Set(val *StartTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStartTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStartTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartTaskAllOf(val *StartTaskAllOf) *NullableStartTaskAllOf {
	return &NullableStartTaskAllOf{value: val, isSet: true}
}

func (v NullableStartTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


