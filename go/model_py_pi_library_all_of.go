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

// PyPiLibraryAllOf struct for PyPiLibraryAllOf
type PyPiLibraryAllOf struct {
	Type *string `json:"type,omitempty"`
	Dep *string `json:"dep,omitempty"`
}

// NewPyPiLibraryAllOf instantiates a new PyPiLibraryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPyPiLibraryAllOf() *PyPiLibraryAllOf {
	this := PyPiLibraryAllOf{}
	var type_ string = "pypi"
	this.Type = &type_
	return &this
}

// NewPyPiLibraryAllOfWithDefaults instantiates a new PyPiLibraryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPyPiLibraryAllOfWithDefaults() *PyPiLibraryAllOf {
	this := PyPiLibraryAllOf{}
	var type_ string = "pypi"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PyPiLibraryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyPiLibraryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PyPiLibraryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PyPiLibraryAllOf) SetType(v string) {
	o.Type = &v
}

// GetDep returns the Dep field value if set, zero value otherwise.
func (o *PyPiLibraryAllOf) GetDep() string {
	if o == nil || o.Dep == nil {
		var ret string
		return ret
	}
	return *o.Dep
}

// GetDepOk returns a tuple with the Dep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyPiLibraryAllOf) GetDepOk() (*string, bool) {
	if o == nil || o.Dep == nil {
		return nil, false
	}
	return o.Dep, true
}

// HasDep returns a boolean if a field has been set.
func (o *PyPiLibraryAllOf) HasDep() bool {
	if o != nil && o.Dep != nil {
		return true
	}

	return false
}

// SetDep gets a reference to the given string and assigns it to the Dep field.
func (o *PyPiLibraryAllOf) SetDep(v string) {
	o.Dep = &v
}

func (o PyPiLibraryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Dep != nil {
		toSerialize["dep"] = o.Dep
	}
	return json.Marshal(toSerialize)
}

type NullablePyPiLibraryAllOf struct {
	value *PyPiLibraryAllOf
	isSet bool
}

func (v NullablePyPiLibraryAllOf) Get() *PyPiLibraryAllOf {
	return v.value
}

func (v *NullablePyPiLibraryAllOf) Set(val *PyPiLibraryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePyPiLibraryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePyPiLibraryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePyPiLibraryAllOf(val *PyPiLibraryAllOf) *NullablePyPiLibraryAllOf {
	return &NullablePyPiLibraryAllOf{value: val, isSet: true}
}

func (v NullablePyPiLibraryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePyPiLibraryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


