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

// ParametersToWriteAFile1 - Options for reading a file (header, file path, separator,         delimiter, ...).
type ParametersToWriteAFile1 struct {
	CsvReader *CsvReader
	ParquetReader *ParquetReader
}

// CsvReaderAsParametersToWriteAFile1 is a convenience function that returns CsvReader wrapped in ParametersToWriteAFile1
func CsvReaderAsParametersToWriteAFile1(v *CsvReader) ParametersToWriteAFile1 {
	return ParametersToWriteAFile1{
		CsvReader: v,
	}
}

// ParquetReaderAsParametersToWriteAFile1 is a convenience function that returns ParquetReader wrapped in ParametersToWriteAFile1
func ParquetReaderAsParametersToWriteAFile1(v *ParquetReader) ParametersToWriteAFile1 {
	return ParametersToWriteAFile1{
		ParquetReader: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ParametersToWriteAFile1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CsvReader
	err = newStrictDecoder(data).Decode(&dst.CsvReader)
	if err == nil {
		jsonCsvReader, _ := json.Marshal(dst.CsvReader)
		if string(jsonCsvReader) == "{}" { // empty struct
			dst.CsvReader = nil
		} else {
			match++
		}
	} else {
		dst.CsvReader = nil
	}

	// try to unmarshal data into ParquetReader
	err = newStrictDecoder(data).Decode(&dst.ParquetReader)
	if err == nil {
		jsonParquetReader, _ := json.Marshal(dst.ParquetReader)
		if string(jsonParquetReader) == "{}" { // empty struct
			dst.ParquetReader = nil
		} else {
			match++
		}
	} else {
		dst.ParquetReader = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CsvReader = nil
		dst.ParquetReader = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ParametersToWriteAFile1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ParametersToWriteAFile1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ParametersToWriteAFile1) MarshalJSON() ([]byte, error) {
	if src.CsvReader != nil {
		return json.Marshal(&src.CsvReader)
	}

	if src.ParquetReader != nil {
		return json.Marshal(&src.ParquetReader)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ParametersToWriteAFile1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CsvReader != nil {
		return obj.CsvReader
	}

	if obj.ParquetReader != nil {
		return obj.ParquetReader
	}

	// all schemas are nil
	return nil
}

type NullableParametersToWriteAFile1 struct {
	value *ParametersToWriteAFile1
	isSet bool
}

func (v NullableParametersToWriteAFile1) Get() *ParametersToWriteAFile1 {
	return v.value
}

func (v *NullableParametersToWriteAFile1) Set(val *ParametersToWriteAFile1) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersToWriteAFile1) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersToWriteAFile1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersToWriteAFile1(val *ParametersToWriteAFile1) *NullableParametersToWriteAFile1 {
	return &NullableParametersToWriteAFile1{value: val, isSet: true}
}

func (v NullableParametersToWriteAFile1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersToWriteAFile1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


