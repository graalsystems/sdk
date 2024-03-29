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

// checks if the ExternalOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalOptions{}

// ExternalOptions struct for ExternalOptions
type ExternalOptions struct {
	Host *string `json:"host,omitempty"`
	Port *float32 `json:"port,omitempty"`
	Database *string `json:"database,omitempty"`
	Schema *string `json:"schema,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewExternalOptions instantiates a new ExternalOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalOptions() *ExternalOptions {
	this := ExternalOptions{}
	return &this
}

// NewExternalOptionsWithDefaults instantiates a new ExternalOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalOptionsWithDefaults() *ExternalOptions {
	this := ExternalOptions{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ExternalOptions) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOptions) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ExternalOptions) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ExternalOptions) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ExternalOptions) GetPort() float32 {
	if o == nil || IsNil(o.Port) {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOptions) GetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ExternalOptions) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *ExternalOptions) SetPort(v float32) {
	o.Port = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ExternalOptions) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOptions) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ExternalOptions) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ExternalOptions) SetDatabase(v string) {
	o.Database = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ExternalOptions) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOptions) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ExternalOptions) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ExternalOptions) SetSchema(v string) {
	o.Schema = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExternalOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExternalOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExternalOptions) SetType(v string) {
	o.Type = &v
}

func (o ExternalOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableExternalOptions struct {
	value *ExternalOptions
	isSet bool
}

func (v NullableExternalOptions) Get() *ExternalOptions {
	return v.value
}

func (v *NullableExternalOptions) Set(val *ExternalOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalOptions(val *ExternalOptions) *NullableExternalOptions {
	return &NullableExternalOptions{value: val, isSet: true}
}

func (v NullableExternalOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


