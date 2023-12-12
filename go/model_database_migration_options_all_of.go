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

// DatabaseMigrationOptionsAllOf struct for DatabaseMigrationOptionsAllOf
type DatabaseMigrationOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	DatawarehouseId *string `json:"datawarehouse_id,omitempty"`
	Repository *GitLibrary `json:"repository,omitempty"`
}

// NewDatabaseMigrationOptionsAllOf instantiates a new DatabaseMigrationOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseMigrationOptionsAllOf() *DatabaseMigrationOptionsAllOf {
	this := DatabaseMigrationOptionsAllOf{}
	var type_ string = "database-migration"
	this.Type = &type_
	return &this
}

// NewDatabaseMigrationOptionsAllOfWithDefaults instantiates a new DatabaseMigrationOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseMigrationOptionsAllOfWithDefaults() *DatabaseMigrationOptionsAllOf {
	this := DatabaseMigrationOptionsAllOf{}
	var type_ string = "database-migration"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DatabaseMigrationOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseMigrationOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatabaseMigrationOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DatabaseMigrationOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetDatawarehouseId returns the DatawarehouseId field value if set, zero value otherwise.
func (o *DatabaseMigrationOptionsAllOf) GetDatawarehouseId() string {
	if o == nil || o.DatawarehouseId == nil {
		var ret string
		return ret
	}
	return *o.DatawarehouseId
}

// GetDatawarehouseIdOk returns a tuple with the DatawarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseMigrationOptionsAllOf) GetDatawarehouseIdOk() (*string, bool) {
	if o == nil || o.DatawarehouseId == nil {
		return nil, false
	}
	return o.DatawarehouseId, true
}

// HasDatawarehouseId returns a boolean if a field has been set.
func (o *DatabaseMigrationOptionsAllOf) HasDatawarehouseId() bool {
	if o != nil && o.DatawarehouseId != nil {
		return true
	}

	return false
}

// SetDatawarehouseId gets a reference to the given string and assigns it to the DatawarehouseId field.
func (o *DatabaseMigrationOptionsAllOf) SetDatawarehouseId(v string) {
	o.DatawarehouseId = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *DatabaseMigrationOptionsAllOf) GetRepository() GitLibrary {
	if o == nil || o.Repository == nil {
		var ret GitLibrary
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseMigrationOptionsAllOf) GetRepositoryOk() (*GitLibrary, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *DatabaseMigrationOptionsAllOf) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GitLibrary and assigns it to the Repository field.
func (o *DatabaseMigrationOptionsAllOf) SetRepository(v GitLibrary) {
	o.Repository = &v
}

func (o DatabaseMigrationOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DatawarehouseId != nil {
		toSerialize["datawarehouse_id"] = o.DatawarehouseId
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseMigrationOptionsAllOf struct {
	value *DatabaseMigrationOptionsAllOf
	isSet bool
}

func (v NullableDatabaseMigrationOptionsAllOf) Get() *DatabaseMigrationOptionsAllOf {
	return v.value
}

func (v *NullableDatabaseMigrationOptionsAllOf) Set(val *DatabaseMigrationOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseMigrationOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseMigrationOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseMigrationOptionsAllOf(val *DatabaseMigrationOptionsAllOf) *NullableDatabaseMigrationOptionsAllOf {
	return &NullableDatabaseMigrationOptionsAllOf{value: val, isSet: true}
}

func (v NullableDatabaseMigrationOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseMigrationOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


