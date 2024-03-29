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

// Right struct for Right
type Right struct {
	Array *[]string
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Right) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into []string
	err = json.Unmarshal(data, &dst.Array);
	if err == nil {
		json_array, _ := json.Marshal(dst.Array)
		if string(json_array) == "{}" { // empty struct
			dst.Array = nil
		} else {
			return nil // data stored in dst.Array, return on the first match
		}
	} else {
		dst.Array = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.String)
		if string(jsonstring) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Right)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Right) MarshalJSON() ([]byte, error) {
	if src.Array != nil {
		return json.Marshal(&src.Array)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRight struct {
	value *Right
	isSet bool
}

func (v NullableRight) Get() *Right {
	return v.value
}

func (v *NullableRight) Set(val *Right) {
	v.value = val
	v.isSet = true
}

func (v NullableRight) IsSet() bool {
	return v.isSet
}

func (v *NullableRight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRight(val *Right) *NullableRight {
	return &NullableRight{value: val, isSet: true}
}

func (v NullableRight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


