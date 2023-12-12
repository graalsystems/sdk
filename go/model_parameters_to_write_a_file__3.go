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

// ParametersToWriteAFile3 - Options for writing a file (header, file path, separator,         delimiter, ...).
type ParametersToWriteAFile3 struct {
	CsvWriter *CsvWriter
	ParquetWriter *ParquetWriter
}

// CsvWriterAsParametersToWriteAFile3 is a convenience function that returns CsvWriter wrapped in ParametersToWriteAFile3
func CsvWriterAsParametersToWriteAFile3(v *CsvWriter) ParametersToWriteAFile3 {
	return ParametersToWriteAFile3{
		CsvWriter: v,
	}
}

// ParquetWriterAsParametersToWriteAFile3 is a convenience function that returns ParquetWriter wrapped in ParametersToWriteAFile3
func ParquetWriterAsParametersToWriteAFile3(v *ParquetWriter) ParametersToWriteAFile3 {
	return ParametersToWriteAFile3{
		ParquetWriter: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ParametersToWriteAFile3) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CsvWriter
	err = newStrictDecoder(data).Decode(&dst.CsvWriter)
	if err == nil {
		jsonCsvWriter, _ := json.Marshal(dst.CsvWriter)
		if string(jsonCsvWriter) == "{}" { // empty struct
			dst.CsvWriter = nil
		} else {
			match++
		}
	} else {
		dst.CsvWriter = nil
	}

	// try to unmarshal data into ParquetWriter
	err = newStrictDecoder(data).Decode(&dst.ParquetWriter)
	if err == nil {
		jsonParquetWriter, _ := json.Marshal(dst.ParquetWriter)
		if string(jsonParquetWriter) == "{}" { // empty struct
			dst.ParquetWriter = nil
		} else {
			match++
		}
	} else {
		dst.ParquetWriter = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CsvWriter = nil
		dst.ParquetWriter = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ParametersToWriteAFile3)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ParametersToWriteAFile3)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ParametersToWriteAFile3) MarshalJSON() ([]byte, error) {
	if src.CsvWriter != nil {
		return json.Marshal(&src.CsvWriter)
	}

	if src.ParquetWriter != nil {
		return json.Marshal(&src.ParquetWriter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ParametersToWriteAFile3) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CsvWriter != nil {
		return obj.CsvWriter
	}

	if obj.ParquetWriter != nil {
		return obj.ParquetWriter
	}

	// all schemas are nil
	return nil
}

type NullableParametersToWriteAFile3 struct {
	value *ParametersToWriteAFile3
	isSet bool
}

func (v NullableParametersToWriteAFile3) Get() *ParametersToWriteAFile3 {
	return v.value
}

func (v *NullableParametersToWriteAFile3) Set(val *ParametersToWriteAFile3) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersToWriteAFile3) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersToWriteAFile3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersToWriteAFile3(val *ParametersToWriteAFile3) *NullableParametersToWriteAFile3 {
	return &NullableParametersToWriteAFile3{value: val, isSet: true}
}

func (v NullableParametersToWriteAFile3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersToWriteAFile3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


