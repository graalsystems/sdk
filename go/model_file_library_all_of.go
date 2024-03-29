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

// FileLibraryAllOf struct for FileLibraryAllOf
type FileLibraryAllOf struct {
	Type *string `json:"type,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewFileLibraryAllOf instantiates a new FileLibraryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLibraryAllOf() *FileLibraryAllOf {
	this := FileLibraryAllOf{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// NewFileLibraryAllOfWithDefaults instantiates a new FileLibraryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLibraryAllOfWithDefaults() *FileLibraryAllOf {
	this := FileLibraryAllOf{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileLibraryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLibraryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileLibraryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileLibraryAllOf) SetType(v string) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FileLibraryAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLibraryAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FileLibraryAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FileLibraryAllOf) SetKey(v string) {
	o.Key = &v
}

func (o FileLibraryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableFileLibraryAllOf struct {
	value *FileLibraryAllOf
	isSet bool
}

func (v NullableFileLibraryAllOf) Get() *FileLibraryAllOf {
	return v.value
}

func (v *NullableFileLibraryAllOf) Set(val *FileLibraryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLibraryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLibraryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLibraryAllOf(val *FileLibraryAllOf) *NullableFileLibraryAllOf {
	return &NullableFileLibraryAllOf{value: val, isSet: true}
}

func (v NullableFileLibraryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLibraryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


