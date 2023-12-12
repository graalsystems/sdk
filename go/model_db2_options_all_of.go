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

// DB2OptionsAllOf struct for DB2OptionsAllOf
type DB2OptionsAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewDB2OptionsAllOf instantiates a new DB2OptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDB2OptionsAllOf() *DB2OptionsAllOf {
	this := DB2OptionsAllOf{}
	var type_ string = "db2"
	this.Type = &type_
	return &this
}

// NewDB2OptionsAllOfWithDefaults instantiates a new DB2OptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDB2OptionsAllOfWithDefaults() *DB2OptionsAllOf {
	this := DB2OptionsAllOf{}
	var type_ string = "db2"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DB2OptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DB2OptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DB2OptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DB2OptionsAllOf) SetType(v string) {
	o.Type = &v
}

func (o DB2OptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDB2OptionsAllOf struct {
	value *DB2OptionsAllOf
	isSet bool
}

func (v NullableDB2OptionsAllOf) Get() *DB2OptionsAllOf {
	return v.value
}

func (v *NullableDB2OptionsAllOf) Set(val *DB2OptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDB2OptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDB2OptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDB2OptionsAllOf(val *DB2OptionsAllOf) *NullableDB2OptionsAllOf {
	return &NullableDB2OptionsAllOf{value: val, isSet: true}
}

func (v NullableDB2OptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDB2OptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


