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

// TrinoOptionsAllOf struct for TrinoOptionsAllOf
type TrinoOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	NumberWorkers *float32 `json:"number_workers,omitempty"`
	Conf *map[string]map[string]interface{} `json:"conf,omitempty"`
}

// NewTrinoOptionsAllOf instantiates a new TrinoOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrinoOptionsAllOf() *TrinoOptionsAllOf {
	this := TrinoOptionsAllOf{}
	var type_ string = "trino"
	this.Type = &type_
	return &this
}

// NewTrinoOptionsAllOfWithDefaults instantiates a new TrinoOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrinoOptionsAllOfWithDefaults() *TrinoOptionsAllOf {
	this := TrinoOptionsAllOf{}
	var type_ string = "trino"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrinoOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrinoOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TrinoOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetNumberWorkers returns the NumberWorkers field value if set, zero value otherwise.
func (o *TrinoOptionsAllOf) GetNumberWorkers() float32 {
	if o == nil || o.NumberWorkers == nil {
		var ret float32
		return ret
	}
	return *o.NumberWorkers
}

// GetNumberWorkersOk returns a tuple with the NumberWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptionsAllOf) GetNumberWorkersOk() (*float32, bool) {
	if o == nil || o.NumberWorkers == nil {
		return nil, false
	}
	return o.NumberWorkers, true
}

// HasNumberWorkers returns a boolean if a field has been set.
func (o *TrinoOptionsAllOf) HasNumberWorkers() bool {
	if o != nil && o.NumberWorkers != nil {
		return true
	}

	return false
}

// SetNumberWorkers gets a reference to the given float32 and assigns it to the NumberWorkers field.
func (o *TrinoOptionsAllOf) SetNumberWorkers(v float32) {
	o.NumberWorkers = &v
}

// GetConf returns the Conf field value if set, zero value otherwise.
func (o *TrinoOptionsAllOf) GetConf() map[string]map[string]interface{} {
	if o == nil || o.Conf == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Conf
}

// GetConfOk returns a tuple with the Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrinoOptionsAllOf) GetConfOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Conf == nil {
		return nil, false
	}
	return o.Conf, true
}

// HasConf returns a boolean if a field has been set.
func (o *TrinoOptionsAllOf) HasConf() bool {
	if o != nil && o.Conf != nil {
		return true
	}

	return false
}

// SetConf gets a reference to the given map[string]map[string]interface{} and assigns it to the Conf field.
func (o *TrinoOptionsAllOf) SetConf(v map[string]map[string]interface{}) {
	o.Conf = &v
}

func (o TrinoOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.NumberWorkers != nil {
		toSerialize["number_workers"] = o.NumberWorkers
	}
	if o.Conf != nil {
		toSerialize["conf"] = o.Conf
	}
	return json.Marshal(toSerialize)
}

type NullableTrinoOptionsAllOf struct {
	value *TrinoOptionsAllOf
	isSet bool
}

func (v NullableTrinoOptionsAllOf) Get() *TrinoOptionsAllOf {
	return v.value
}

func (v *NullableTrinoOptionsAllOf) Set(val *TrinoOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTrinoOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTrinoOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrinoOptionsAllOf(val *TrinoOptionsAllOf) *NullableTrinoOptionsAllOf {
	return &NullableTrinoOptionsAllOf{value: val, isSet: true}
}

func (v NullableTrinoOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrinoOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


