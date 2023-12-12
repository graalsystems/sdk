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

// Connector Type of connector
type Connector struct {
	LocalConnector *LocalConnector
	S3Connector *S3Connector
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Connector) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocalConnector
	err = json.Unmarshal(data, &dst.LocalConnector);
	if err == nil {
		jsonLocalConnector, _ := json.Marshal(dst.LocalConnector)
		if string(jsonLocalConnector) == "{}" { // empty struct
			dst.LocalConnector = nil
		} else {
			return nil // data stored in dst.LocalConnector, return on the first match
		}
	} else {
		dst.LocalConnector = nil
	}

	// try to unmarshal JSON data into S3Connector
	err = json.Unmarshal(data, &dst.S3Connector);
	if err == nil {
		jsonS3Connector, _ := json.Marshal(dst.S3Connector)
		if string(jsonS3Connector) == "{}" { // empty struct
			dst.S3Connector = nil
		} else {
			return nil // data stored in dst.S3Connector, return on the first match
		}
	} else {
		dst.S3Connector = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Connector)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Connector) MarshalJSON() ([]byte, error) {
	if src.LocalConnector != nil {
		return json.Marshal(&src.LocalConnector)
	}

	if src.S3Connector != nil {
		return json.Marshal(&src.S3Connector)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConnector struct {
	value *Connector
	isSet bool
}

func (v NullableConnector) Get() *Connector {
	return v.value
}

func (v *NullableConnector) Set(val *Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnector(val *Connector) *NullableConnector {
	return &NullableConnector{value: val, isSet: true}
}

func (v NullableConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

