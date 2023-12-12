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

// ResponseGetUserFollowingUsersUserIdFollowingGetInner - struct for ResponseGetUserFollowingUsersUserIdFollowingGetInner
type ResponseGetUserFollowingUsersUserIdFollowingGetInner struct {
	Project1 *Project1
	User1 *User1
}

// Project1AsResponseGetUserFollowingUsersUserIdFollowingGetInner is a convenience function that returns Project1 wrapped in ResponseGetUserFollowingUsersUserIdFollowingGetInner
func Project1AsResponseGetUserFollowingUsersUserIdFollowingGetInner(v *Project1) ResponseGetUserFollowingUsersUserIdFollowingGetInner {
	return ResponseGetUserFollowingUsersUserIdFollowingGetInner{
		Project1: v,
	}
}

// User1AsResponseGetUserFollowingUsersUserIdFollowingGetInner is a convenience function that returns User1 wrapped in ResponseGetUserFollowingUsersUserIdFollowingGetInner
func User1AsResponseGetUserFollowingUsersUserIdFollowingGetInner(v *User1) ResponseGetUserFollowingUsersUserIdFollowingGetInner {
	return ResponseGetUserFollowingUsersUserIdFollowingGetInner{
		User1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResponseGetUserFollowingUsersUserIdFollowingGetInner) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ResponseGetUserFollowingUsersUserIdFollowingGetInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResponseGetUserFollowingUsersUserIdFollowingGetInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResponseGetUserFollowingUsersUserIdFollowingGetInner) MarshalJSON() ([]byte, error) {
	if src.Project1 != nil {
		return json.Marshal(&src.Project1)
	}

	if src.User1 != nil {
		return json.Marshal(&src.User1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResponseGetUserFollowingUsersUserIdFollowingGetInner) GetActualInstance() (interface{}) {
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

type NullableResponseGetUserFollowingUsersUserIdFollowingGetInner struct {
	value *ResponseGetUserFollowingUsersUserIdFollowingGetInner
	isSet bool
}

func (v NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) Get() *ResponseGetUserFollowingUsersUserIdFollowingGetInner {
	return v.value
}

func (v *NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) Set(val *ResponseGetUserFollowingUsersUserIdFollowingGetInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseGetUserFollowingUsersUserIdFollowingGetInner(val *ResponseGetUserFollowingUsersUserIdFollowingGetInner) *NullableResponseGetUserFollowingUsersUserIdFollowingGetInner {
	return &NullableResponseGetUserFollowingUsersUserIdFollowingGetInner{value: val, isSet: true}
}

func (v NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseGetUserFollowingUsersUserIdFollowingGetInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

