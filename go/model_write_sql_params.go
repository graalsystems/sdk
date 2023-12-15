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

// checks if the WriteSQLParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteSQLParams{}

// WriteSQLParams struct for WriteSQLParams
type WriteSQLParams struct {
	// Type SQL server.
	Type *string `json:"type,omitempty"`
	Object SQLDataWriter `json:"object"`
	// List of dictionaries representing the columns (name and type). For each dictionnary the         'name' represents the column name and the 'type' their corresponding type.
	Structure []map[string]string `json:"structure,omitempty"`
	Connector DatabaseConnector `json:"connector"`
}

type _WriteSQLParams WriteSQLParams

// NewWriteSQLParams instantiates a new WriteSQLParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteSQLParams(object SQLDataWriter, connector DatabaseConnector) *WriteSQLParams {
	this := WriteSQLParams{}
	var type_ string = "sql"
	this.Type = &type_
	this.Object = object
	this.Connector = connector
	return &this
}

// NewWriteSQLParamsWithDefaults instantiates a new WriteSQLParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteSQLParamsWithDefaults() *WriteSQLParams {
	this := WriteSQLParams{}
	var type_ string = "sql"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WriteSQLParams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteSQLParams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WriteSQLParams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WriteSQLParams) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value
func (o *WriteSQLParams) GetObject() SQLDataWriter {
	if o == nil {
		var ret SQLDataWriter
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *WriteSQLParams) GetObjectOk() (*SQLDataWriter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *WriteSQLParams) SetObject(v SQLDataWriter) {
	o.Object = v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *WriteSQLParams) GetStructure() []map[string]string {
	if o == nil || IsNil(o.Structure) {
		var ret []map[string]string
		return ret
	}
	return o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteSQLParams) GetStructureOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Structure) {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *WriteSQLParams) HasStructure() bool {
	if o != nil && !IsNil(o.Structure) {
		return true
	}

	return false
}

// SetStructure gets a reference to the given []map[string]string and assigns it to the Structure field.
func (o *WriteSQLParams) SetStructure(v []map[string]string) {
	o.Structure = v
}

// GetConnector returns the Connector field value
func (o *WriteSQLParams) GetConnector() DatabaseConnector {
	if o == nil {
		var ret DatabaseConnector
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *WriteSQLParams) GetConnectorOk() (*DatabaseConnector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *WriteSQLParams) SetConnector(v DatabaseConnector) {
	o.Connector = v
}

func (o WriteSQLParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteSQLParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.Structure) {
		toSerialize["structure"] = o.Structure
	}
	toSerialize["connector"] = o.Connector
	return toSerialize, nil
}

func (o *WriteSQLParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"connector",
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

	varWriteSQLParams := _WriteSQLParams{}

	err = json.Unmarshal(bytes, &varWriteSQLParams)

	if err != nil {
		return err
	}

	*o = WriteSQLParams(varWriteSQLParams)

	return err
}

type NullableWriteSQLParams struct {
	value *WriteSQLParams
	isSet bool
}

func (v NullableWriteSQLParams) Get() *WriteSQLParams {
	return v.value
}

func (v *NullableWriteSQLParams) Set(val *WriteSQLParams) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteSQLParams) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteSQLParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteSQLParams(val *WriteSQLParams) *NullableWriteSQLParams {
	return &NullableWriteSQLParams{value: val, isSet: true}
}

func (v NullableWriteSQLParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteSQLParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


