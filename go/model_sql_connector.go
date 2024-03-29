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

// checks if the SQLConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SQLConnector{}

// SQLConnector Connect to a SQL server.  Attributes ---------- type : ConnectorType     SQL connector type. options : SQLConnectorOptions     Options to connect to a SQL server.
type SQLConnector struct {
	// Connection to a SQL server.
	Type *string `json:"type,omitempty"`
	// Options to connect to a SQL server.
	Options map[string]interface{} `json:"options,omitempty"`
}

// NewSQLConnector instantiates a new SQLConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSQLConnector() *SQLConnector {
	this := SQLConnector{}
	var type_ string = "sql"
	this.Type = &type_
	return &this
}

// NewSQLConnectorWithDefaults instantiates a new SQLConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSQLConnectorWithDefaults() *SQLConnector {
	this := SQLConnector{}
	var type_ string = "sql"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SQLConnector) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SQLConnector) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SQLConnector) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SQLConnector) SetType(v string) {
	o.Type = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SQLConnector) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SQLConnector) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SQLConnector) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *SQLConnector) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o SQLConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SQLConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableSQLConnector struct {
	value *SQLConnector
	isSet bool
}

func (v NullableSQLConnector) Get() *SQLConnector {
	return v.value
}

func (v *NullableSQLConnector) Set(val *SQLConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableSQLConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableSQLConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSQLConnector(val *SQLConnector) *NullableSQLConnector {
	return &NullableSQLConnector{value: val, isSet: true}
}

func (v NullableSQLConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSQLConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


