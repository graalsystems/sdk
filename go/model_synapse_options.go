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

// checks if the SynapseOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SynapseOptions{}

// SynapseOptions struct for SynapseOptions
type SynapseOptions struct {
	DataWarehouseOptions
	Host *string `json:"host,omitempty"`
	Port *float32 `json:"port,omitempty"`
	Database *string `json:"database,omitempty"`
	Schema *string `json:"schema,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSynapseOptions instantiates a new SynapseOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynapseOptions() *SynapseOptions {
	this := SynapseOptions{}
	var type_ string = "synapse"
	this.Type = &type_
	return &this
}

// NewSynapseOptionsWithDefaults instantiates a new SynapseOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynapseOptionsWithDefaults() *SynapseOptions {
	this := SynapseOptions{}
	var type_ string = "synapse"
	this.Type = &type_
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SynapseOptions) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynapseOptions) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SynapseOptions) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SynapseOptions) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SynapseOptions) GetPort() float32 {
	if o == nil || IsNil(o.Port) {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynapseOptions) GetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SynapseOptions) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *SynapseOptions) SetPort(v float32) {
	o.Port = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *SynapseOptions) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynapseOptions) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *SynapseOptions) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *SynapseOptions) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SynapseOptions) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynapseOptions) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SynapseOptions) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *SynapseOptions) SetSchema(v string) {
	o.Schema = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SynapseOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynapseOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SynapseOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SynapseOptions) SetType(v string) {
	o.Type = &v
}

func (o SynapseOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SynapseOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataWarehouseOptions, errDataWarehouseOptions := json.Marshal(o.DataWarehouseOptions)
	if errDataWarehouseOptions != nil {
		return map[string]interface{}{}, errDataWarehouseOptions
	}
	errDataWarehouseOptions = json.Unmarshal([]byte(serializedDataWarehouseOptions), &toSerialize)
	if errDataWarehouseOptions != nil {
		return map[string]interface{}{}, errDataWarehouseOptions
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSynapseOptions struct {
	value *SynapseOptions
	isSet bool
}

func (v NullableSynapseOptions) Get() *SynapseOptions {
	return v.value
}

func (v *NullableSynapseOptions) Set(val *SynapseOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSynapseOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSynapseOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynapseOptions(val *SynapseOptions) *NullableSynapseOptions {
	return &NullableSynapseOptions{value: val, isSet: true}
}

func (v NullableSynapseOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynapseOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


