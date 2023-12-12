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

// checks if the ReadFileParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadFileParams{}

// ReadFileParams struct for ReadFileParams
type ReadFileParams struct {
	// Type of object file.
	Type *string `json:"type,omitempty"`
	Connector Connector1 `json:"connector"`
	// List of dictionaries representing the columns         (name and type). For each dictionnary the 'name' represents the column         name and the 'type' their corresponding type.
	Structure []map[string]string `json:"structure,omitempty"`
	Object ParametersToWriteAFile1 `json:"object"`
}

type _ReadFileParams ReadFileParams

// NewReadFileParams instantiates a new ReadFileParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFileParams(connector Connector1, object ParametersToWriteAFile1) *ReadFileParams {
	this := ReadFileParams{}
	var type_ string = "file"
	this.Type = &type_
	this.Connector = connector
	this.Object = object
	return &this
}

// NewReadFileParamsWithDefaults instantiates a new ReadFileParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFileParamsWithDefaults() *ReadFileParams {
	this := ReadFileParams{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReadFileParams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFileParams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReadFileParams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReadFileParams) SetType(v string) {
	o.Type = &v
}

// GetConnector returns the Connector field value
func (o *ReadFileParams) GetConnector() Connector1 {
	if o == nil {
		var ret Connector1
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ReadFileParams) GetConnectorOk() (*Connector1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ReadFileParams) SetConnector(v Connector1) {
	o.Connector = v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *ReadFileParams) GetStructure() []map[string]string {
	if o == nil || IsNil(o.Structure) {
		var ret []map[string]string
		return ret
	}
	return o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFileParams) GetStructureOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Structure) {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *ReadFileParams) HasStructure() bool {
	if o != nil && !IsNil(o.Structure) {
		return true
	}

	return false
}

// SetStructure gets a reference to the given []map[string]string and assigns it to the Structure field.
func (o *ReadFileParams) SetStructure(v []map[string]string) {
	o.Structure = v
}

// GetObject returns the Object field value
func (o *ReadFileParams) GetObject() ParametersToWriteAFile1 {
	if o == nil {
		var ret ParametersToWriteAFile1
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ReadFileParams) GetObjectOk() (*ParametersToWriteAFile1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ReadFileParams) SetObject(v ParametersToWriteAFile1) {
	o.Object = v
}

func (o ReadFileParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadFileParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["connector"] = o.Connector
	if !IsNil(o.Structure) {
		toSerialize["structure"] = o.Structure
	}
	toSerialize["object"] = o.Object
	return toSerialize, nil
}

func (o *ReadFileParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connector",
		"object",
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

	varReadFileParams := _ReadFileParams{}

	err = json.Unmarshal(bytes, &varReadFileParams)

	if err != nil {
		return err
	}

	*o = ReadFileParams(varReadFileParams)

	return err
}

type NullableReadFileParams struct {
	value *ReadFileParams
	isSet bool
}

func (v NullableReadFileParams) Get() *ReadFileParams {
	return v.value
}

func (v *NullableReadFileParams) Set(val *ReadFileParams) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFileParams) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFileParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFileParams(val *ReadFileParams) *NullableReadFileParams {
	return &NullableReadFileParams{value: val, isSet: true}
}

func (v NullableReadFileParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFileParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

