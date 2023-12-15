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
	"fmt"
)

// checks if the WriteFileParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteFileParams{}

// WriteFileParams struct for WriteFileParams
type WriteFileParams struct {
	// Type of object file.
	Type *string `json:"type,omitempty"`
	Connector Connector1 `json:"connector"`
	// List of dictionaries representing the columns         (name and type). For each dictionnary the 'name' represents the column         name and the 'type' their corresponding type.
	Structure []map[string]string `json:"structure,omitempty"`
	Object ParametersToWriteAFile3 `json:"object"`
}

type _WriteFileParams WriteFileParams

// NewWriteFileParams instantiates a new WriteFileParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteFileParams(connector Connector1, object ParametersToWriteAFile3) *WriteFileParams {
	this := WriteFileParams{}
	var type_ string = "file"
	this.Type = &type_
	this.Connector = connector
	this.Object = object
	return &this
}

// NewWriteFileParamsWithDefaults instantiates a new WriteFileParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteFileParamsWithDefaults() *WriteFileParams {
	this := WriteFileParams{}
	var type_ string = "file"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WriteFileParams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteFileParams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WriteFileParams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WriteFileParams) SetType(v string) {
	o.Type = &v
}

// GetConnector returns the Connector field value
func (o *WriteFileParams) GetConnector() Connector1 {
	if o == nil {
		var ret Connector1
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *WriteFileParams) GetConnectorOk() (*Connector1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *WriteFileParams) SetConnector(v Connector1) {
	o.Connector = v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *WriteFileParams) GetStructure() []map[string]string {
	if o == nil || IsNil(o.Structure) {
		var ret []map[string]string
		return ret
	}
	return o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteFileParams) GetStructureOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Structure) {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *WriteFileParams) HasStructure() bool {
	if o != nil && !IsNil(o.Structure) {
		return true
	}

	return false
}

// SetStructure gets a reference to the given []map[string]string and assigns it to the Structure field.
func (o *WriteFileParams) SetStructure(v []map[string]string) {
	o.Structure = v
}

// GetObject returns the Object field value
func (o *WriteFileParams) GetObject() ParametersToWriteAFile3 {
	if o == nil {
		var ret ParametersToWriteAFile3
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *WriteFileParams) GetObjectOk() (*ParametersToWriteAFile3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *WriteFileParams) SetObject(v ParametersToWriteAFile3) {
	o.Object = v
}

func (o WriteFileParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteFileParams) ToMap() (map[string]interface{}, error) {
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

func (o *WriteFileParams) UnmarshalJSON(bytes []byte) (err error) {
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

	varWriteFileParams := _WriteFileParams{}

	err = json.Unmarshal(bytes, &varWriteFileParams)

	if err != nil {
		return err
	}

	*o = WriteFileParams(varWriteFileParams)

	return err
}

type NullableWriteFileParams struct {
	value *WriteFileParams
	isSet bool
}

func (v NullableWriteFileParams) Get() *WriteFileParams {
	return v.value
}

func (v *NullableWriteFileParams) Set(val *WriteFileParams) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteFileParams) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteFileParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteFileParams(val *WriteFileParams) *NullableWriteFileParams {
	return &NullableWriteFileParams{value: val, isSet: true}
}

func (v NullableWriteFileParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteFileParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


