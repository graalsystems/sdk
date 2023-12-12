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

// checks if the PickleWriter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PickleWriter{}

// PickleWriter struct for PickleWriter
type PickleWriter struct {
	// Name of the path.
	Path string `json:"path"`
	// Define file type for Pickle.
	Type *string `json:"type,omitempty"`
}

type _PickleWriter PickleWriter

// NewPickleWriter instantiates a new PickleWriter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPickleWriter(path string) *PickleWriter {
	this := PickleWriter{}
	this.Path = path
	var type_ string = "pickle"
	this.Type = &type_
	return &this
}

// NewPickleWriterWithDefaults instantiates a new PickleWriter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPickleWriterWithDefaults() *PickleWriter {
	this := PickleWriter{}
	var type_ string = "pickle"
	this.Type = &type_
	return &this
}

// GetPath returns the Path field value
func (o *PickleWriter) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PickleWriter) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PickleWriter) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PickleWriter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickleWriter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PickleWriter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PickleWriter) SetType(v string) {
	o.Type = &v
}

func (o PickleWriter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PickleWriter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *PickleWriter) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varPickleWriter := _PickleWriter{}

	err = json.Unmarshal(bytes, &varPickleWriter)

	if err != nil {
		return err
	}

	*o = PickleWriter(varPickleWriter)

	return err
}

type NullablePickleWriter struct {
	value *PickleWriter
	isSet bool
}

func (v NullablePickleWriter) Get() *PickleWriter {
	return v.value
}

func (v *NullablePickleWriter) Set(val *PickleWriter) {
	v.value = val
	v.isSet = true
}

func (v NullablePickleWriter) IsSet() bool {
	return v.isSet
}

func (v *NullablePickleWriter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePickleWriter(val *PickleWriter) *NullablePickleWriter {
	return &NullablePickleWriter{value: val, isSet: true}
}

func (v NullablePickleWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePickleWriter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

