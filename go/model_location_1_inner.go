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

// Location1Inner struct for Location1Inner
type Location1Inner struct {
	int32 *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Location1Inner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Location1Inner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Location1Inner) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocation1Inner struct {
	value *Location1Inner
	isSet bool
}

func (v NullableLocation1Inner) Get() *Location1Inner {
	return v.value
}

func (v *NullableLocation1Inner) Set(val *Location1Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation1Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation1Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation1Inner(val *Location1Inner) *NullableLocation1Inner {
	return &NullableLocation1Inner{value: val, isSet: true}
}

func (v NullableLocation1Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation1Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

