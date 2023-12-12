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

// EtlDagVisitorsEnumsLanguageType Map the dag input to LanguageType.
type EtlDagVisitorsEnumsLanguageType string

// List of etl__dag__visitors__enums__LanguageType
const (
	PYSPARK EtlDagVisitorsEnumsLanguageType = "pyspark"
	PANDAS EtlDagVisitorsEnumsLanguageType = "pandas"
	JAVA EtlDagVisitorsEnumsLanguageType = "java"
	DASK EtlDagVisitorsEnumsLanguageType = "dask"
	BEAM EtlDagVisitorsEnumsLanguageType = "beam"
	FLINK EtlDagVisitorsEnumsLanguageType = "flink"
)

// All allowed values of EtlDagVisitorsEnumsLanguageType enum
var AllowedEtlDagVisitorsEnumsLanguageTypeEnumValues = []EtlDagVisitorsEnumsLanguageType{
	"pyspark",
	"pandas",
	"java",
	"dask",
	"beam",
	"flink",
}

func (v *EtlDagVisitorsEnumsLanguageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EtlDagVisitorsEnumsLanguageType(value)
	for _, existing := range AllowedEtlDagVisitorsEnumsLanguageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EtlDagVisitorsEnumsLanguageType", value)
}

// NewEtlDagVisitorsEnumsLanguageTypeFromValue returns a pointer to a valid EtlDagVisitorsEnumsLanguageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEtlDagVisitorsEnumsLanguageTypeFromValue(v string) (*EtlDagVisitorsEnumsLanguageType, error) {
	ev := EtlDagVisitorsEnumsLanguageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EtlDagVisitorsEnumsLanguageType: valid values are %v", v, AllowedEtlDagVisitorsEnumsLanguageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EtlDagVisitorsEnumsLanguageType) IsValid() bool {
	for _, existing := range AllowedEtlDagVisitorsEnumsLanguageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to etl__dag__visitors__enums__LanguageType value
func (v EtlDagVisitorsEnumsLanguageType) Ptr() *EtlDagVisitorsEnumsLanguageType {
	return &v
}

type NullableEtlDagVisitorsEnumsLanguageType struct {
	value *EtlDagVisitorsEnumsLanguageType
	isSet bool
}

func (v NullableEtlDagVisitorsEnumsLanguageType) Get() *EtlDagVisitorsEnumsLanguageType {
	return v.value
}

func (v *NullableEtlDagVisitorsEnumsLanguageType) Set(val *EtlDagVisitorsEnumsLanguageType) {
	v.value = val
	v.isSet = true
}

func (v NullableEtlDagVisitorsEnumsLanguageType) IsSet() bool {
	return v.isSet
}

func (v *NullableEtlDagVisitorsEnumsLanguageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtlDagVisitorsEnumsLanguageType(val *EtlDagVisitorsEnumsLanguageType) *NullableEtlDagVisitorsEnumsLanguageType {
	return &NullableEtlDagVisitorsEnumsLanguageType{value: val, isSet: true}
}

func (v NullableEtlDagVisitorsEnumsLanguageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtlDagVisitorsEnumsLanguageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
