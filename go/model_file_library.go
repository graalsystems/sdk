/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

import (
	"encoding/json"
)

// FileLibrary struct for FileLibrary
type FileLibrary struct {
	Library
	Type *string `json:"type,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewFileLibrary instantiates a new FileLibrary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLibrary() *FileLibrary {
	this := FileLibrary{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// NewFileLibraryWithDefaults instantiates a new FileLibrary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLibraryWithDefaults() *FileLibrary {
	this := FileLibrary{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileLibrary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLibrary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileLibrary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileLibrary) SetType(v string) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FileLibrary) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLibrary) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FileLibrary) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FileLibrary) SetKey(v string) {
	o.Key = &v
}

func (o FileLibrary) MarshalJSON() ([]byte, error) {
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableFileLibrary struct {
	value *FileLibrary
	isSet bool
}

func (v NullableFileLibrary) Get() *FileLibrary {
	return v.value
}

func (v *NullableFileLibrary) Set(val *FileLibrary) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLibrary) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLibrary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLibrary(val *FileLibrary) *NullableFileLibrary {
	return &NullableFileLibrary{value: val, isSet: true}
}

func (v NullableFileLibrary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLibrary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


