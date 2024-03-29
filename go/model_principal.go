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

// Principal - struct for Principal
type Principal struct {
	Identity *Identity
	User *User
}

// IdentityAsPrincipal is a convenience function that returns Identity wrapped in Principal
func IdentityAsPrincipal(v *Identity) Principal {
	return Principal{
		Identity: v,
	}
}

// UserAsPrincipal is a convenience function that returns User wrapped in Principal
func UserAsPrincipal(v *User) Principal {
	return Principal{
		User: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Principal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Identity
	err = newStrictDecoder(data).Decode(&dst.Identity)
	if err == nil {
		jsonIdentity, _ := json.Marshal(dst.Identity)
		if string(jsonIdentity) == "{}" { // empty struct
			dst.Identity = nil
		} else {
			match++
		}
	} else {
		dst.Identity = nil
	}

	// try to unmarshal data into User
	err = newStrictDecoder(data).Decode(&dst.User)
	if err == nil {
		jsonUser, _ := json.Marshal(dst.User)
		if string(jsonUser) == "{}" { // empty struct
			dst.User = nil
		} else {
			match++
		}
	} else {
		dst.User = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Identity = nil
		dst.User = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Principal)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Principal)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Principal) MarshalJSON() ([]byte, error) {
	if src.Identity != nil {
		return json.Marshal(&src.Identity)
	}

	if src.User != nil {
		return json.Marshal(&src.User)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Principal) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Identity != nil {
		return obj.Identity
	}

	if obj.User != nil {
		return obj.User
	}

	// all schemas are nil
	return nil
}

type NullablePrincipal struct {
	value *Principal
	isSet bool
}

func (v NullablePrincipal) Get() *Principal {
	return v.value
}

func (v *NullablePrincipal) Set(val *Principal) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipal(val *Principal) *NullablePrincipal {
	return &NullablePrincipal{value: val, isSet: true}
}

func (v NullablePrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


