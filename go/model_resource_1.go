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

// Resource1 - struct for Resource1
type Resource1 struct {
	Project1 *Project1
	User1 *User1
}

// Project1AsResource1 is a convenience function that returns Project1 wrapped in Resource1
func Project1AsResource1(v *Project1) Resource1 {
	return Resource1{
		Project1: v,
	}
}

// User1AsResource1 is a convenience function that returns User1 wrapped in Resource1
func User1AsResource1(v *User1) Resource1 {
	return Resource1{
		User1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Resource1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Project1
	err = newStrictDecoder(data).Decode(&dst.Project1)
	if err == nil {
		jsonProject1, _ := json.Marshal(dst.Project1)
		if string(jsonProject1) == "{}" { // empty struct
			dst.Project1 = nil
		} else {
			match++
		}
	} else {
		dst.Project1 = nil
	}

	// try to unmarshal data into User1
	err = newStrictDecoder(data).Decode(&dst.User1)
	if err == nil {
		jsonUser1, _ := json.Marshal(dst.User1)
		if string(jsonUser1) == "{}" { // empty struct
			dst.User1 = nil
		} else {
			match++
		}
	} else {
		dst.User1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Project1 = nil
		dst.User1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Resource1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Resource1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Resource1) MarshalJSON() ([]byte, error) {
	if src.Project1 != nil {
		return json.Marshal(&src.Project1)
	}

	if src.User1 != nil {
		return json.Marshal(&src.User1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Resource1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Project1 != nil {
		return obj.Project1
	}

	if obj.User1 != nil {
		return obj.User1
	}

	// all schemas are nil
	return nil
}

type NullableResource1 struct {
	value *Resource1
	isSet bool
}

func (v NullableResource1) Get() *Resource1 {
	return v.value
}

func (v *NullableResource1) Set(val *Resource1) {
	v.value = val
	v.isSet = true
}

func (v NullableResource1) IsSet() bool {
	return v.isSet
}

func (v *NullableResource1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource1(val *Resource1) *NullableResource1 {
	return &NullableResource1{value: val, isSet: true}
}

func (v NullableResource1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


