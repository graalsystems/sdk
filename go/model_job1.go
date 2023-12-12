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

// checks if the Job1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Job1{}

// Job1 Object Job.
type Job1 struct {
	Id string `json:"id"`
	Type *string `json:"type,omitempty"`
	Parent *Project1 `json:"parent,omitempty"`
}

type _Job1 Job1

// NewJob1 instantiates a new Job1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob1(id string) *Job1 {
	this := Job1{}
	this.Id = id
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// NewJob1WithDefaults instantiates a new Job1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJob1WithDefaults() *Job1 {
	this := Job1{}
	var type_ string = "job"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *Job1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Job1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Job1) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Job1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Job1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Job1) SetType(v string) {
	o.Type = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Job1) GetParent() Project1 {
	if o == nil || IsNil(o.Parent) {
		var ret Project1
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job1) GetParentOk() (*Project1, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Job1) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given Project1 and assigns it to the Parent field.
func (o *Job1) SetParent(v Project1) {
	o.Parent = &v
}

func (o Job1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Job1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

func (o *Job1) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varJob1 := _Job1{}

	err = json.Unmarshal(bytes, &varJob1)

	if err != nil {
		return err
	}

	*o = Job1(varJob1)

	return err
}

type NullableJob1 struct {
	value *Job1
	isSet bool
}

func (v NullableJob1) Get() *Job1 {
	return v.value
}

func (v *NullableJob1) Set(val *Job1) {
	v.value = val
	v.isSet = true
}

func (v NullableJob1) IsSet() bool {
	return v.isSet
}

func (v *NullableJob1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob1(val *Job1) *NullableJob1 {
	return &NullableJob1{value: val, isSet: true}
}

func (v NullableJob1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

