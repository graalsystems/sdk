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

// PostgreSQLOptionsAllOf struct for PostgreSQLOptionsAllOf
type PostgreSQLOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewPostgreSQLOptionsAllOf instantiates a new PostgreSQLOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgreSQLOptionsAllOf() *PostgreSQLOptionsAllOf {
	this := PostgreSQLOptionsAllOf{}
	var type_ string = "postgresql"
	this.Type = &type_
	return &this
}

// NewPostgreSQLOptionsAllOfWithDefaults instantiates a new PostgreSQLOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgreSQLOptionsAllOfWithDefaults() *PostgreSQLOptionsAllOf {
	this := PostgreSQLOptionsAllOf{}
	var type_ string = "postgresql"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostgreSQLOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgreSQLOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostgreSQLOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostgreSQLOptionsAllOf) SetType(v string) {
	o.Type = &v
}

func (o PostgreSQLOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePostgreSQLOptionsAllOf struct {
	value *PostgreSQLOptionsAllOf
	isSet bool
}

func (v NullablePostgreSQLOptionsAllOf) Get() *PostgreSQLOptionsAllOf {
	return v.value
}

func (v *NullablePostgreSQLOptionsAllOf) Set(val *PostgreSQLOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgreSQLOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgreSQLOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgreSQLOptionsAllOf(val *PostgreSQLOptionsAllOf) *NullablePostgreSQLOptionsAllOf {
	return &NullablePostgreSQLOptionsAllOf{value: val, isSet: true}
}

func (v NullablePostgreSQLOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgreSQLOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

