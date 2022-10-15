/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// PyPiLibrary struct for PyPiLibrary
type PyPiLibrary struct {
	Library
	Type *string `json:"type,omitempty"`
	Dep *string `json:"dep,omitempty"`
}

// NewPyPiLibrary instantiates a new PyPiLibrary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPyPiLibrary() *PyPiLibrary {
	this := PyPiLibrary{}
	var type_ string = "pypi"
	this.Type = &type_
	return &this
}

// NewPyPiLibraryWithDefaults instantiates a new PyPiLibrary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPyPiLibraryWithDefaults() *PyPiLibrary {
	this := PyPiLibrary{}
	var type_ string = "pypi"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PyPiLibrary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyPiLibrary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PyPiLibrary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PyPiLibrary) SetType(v string) {
	o.Type = &v
}

// GetDep returns the Dep field value if set, zero value otherwise.
func (o *PyPiLibrary) GetDep() string {
	if o == nil || o.Dep == nil {
		var ret string
		return ret
	}
	return *o.Dep
}

// GetDepOk returns a tuple with the Dep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PyPiLibrary) GetDepOk() (*string, bool) {
	if o == nil || o.Dep == nil {
		return nil, false
	}
	return o.Dep, true
}

// HasDep returns a boolean if a field has been set.
func (o *PyPiLibrary) HasDep() bool {
	if o != nil && o.Dep != nil {
		return true
	}

	return false
}

// SetDep gets a reference to the given string and assigns it to the Dep field.
func (o *PyPiLibrary) SetDep(v string) {
	o.Dep = &v
}

func (o PyPiLibrary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedLibrary, errLibrary := json.Marshal(o.Library)
	if errLibrary != nil {
		return []byte{}, errLibrary
	}
	errLibrary = json.Unmarshal([]byte(serializedLibrary), &toSerialize)
	if errLibrary != nil {
		return []byte{}, errLibrary
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Dep != nil {
		toSerialize["dep"] = o.Dep
	}
	return json.Marshal(toSerialize)
}

type NullablePyPiLibrary struct {
	value *PyPiLibrary
	isSet bool
}

func (v NullablePyPiLibrary) Get() *PyPiLibrary {
	return v.value
}

func (v *NullablePyPiLibrary) Set(val *PyPiLibrary) {
	v.value = val
	v.isSet = true
}

func (v NullablePyPiLibrary) IsSet() bool {
	return v.isSet
}

func (v *NullablePyPiLibrary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePyPiLibrary(val *PyPiLibrary) *NullablePyPiLibrary {
	return &NullablePyPiLibrary{value: val, isSet: true}
}

func (v NullablePyPiLibrary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePyPiLibrary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


