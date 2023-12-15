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

// checks if the CsvReader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsvReader{}

// CsvReader struct for CsvReader
type CsvReader struct {
	// Path to CSV file.
	Path string `json:"path"`
	// Delimiter to use between columns.
	Separator *string `json:"separator,omitempty"`
	// Is the first row the header?
	Header *bool `json:"header,omitempty"`
	// Alias for separator.
	Delimiter *string `json:"delimiter,omitempty"`
	// Define file type for CSV.
	Type *string `json:"type,omitempty"`
}

type _CsvReader CsvReader

// NewCsvReader instantiates a new CsvReader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsvReader(path string) *CsvReader {
	this := CsvReader{}
	this.Path = path
	var type_ string = "csv"
	this.Type = &type_
	return &this
}

// NewCsvReaderWithDefaults instantiates a new CsvReader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsvReaderWithDefaults() *CsvReader {
	this := CsvReader{}
	var type_ string = "csv"
	this.Type = &type_
	return &this
}

// GetPath returns the Path field value
func (o *CsvReader) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CsvReader) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CsvReader) SetPath(v string) {
	o.Path = v
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *CsvReader) GetSeparator() string {
	if o == nil || IsNil(o.Separator) {
		var ret string
		return ret
	}
	return *o.Separator
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvReader) GetSeparatorOk() (*string, bool) {
	if o == nil || IsNil(o.Separator) {
		return nil, false
	}
	return o.Separator, true
}

// HasSeparator returns a boolean if a field has been set.
func (o *CsvReader) HasSeparator() bool {
	if o != nil && !IsNil(o.Separator) {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *CsvReader) SetSeparator(v string) {
	o.Separator = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *CsvReader) GetHeader() bool {
	if o == nil || IsNil(o.Header) {
		var ret bool
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvReader) GetHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CsvReader) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given bool and assigns it to the Header field.
func (o *CsvReader) SetHeader(v bool) {
	o.Header = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *CsvReader) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter) {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvReader) GetDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.Delimiter) {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *CsvReader) HasDelimiter() bool {
	if o != nil && !IsNil(o.Delimiter) {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *CsvReader) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CsvReader) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvReader) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CsvReader) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CsvReader) SetType(v string) {
	o.Type = &v
}

func (o CsvReader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsvReader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if !IsNil(o.Separator) {
		toSerialize["separator"] = o.Separator
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Delimiter) {
		toSerialize["delimiter"] = o.Delimiter
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *CsvReader) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varCsvReader := _CsvReader{}

	err = json.Unmarshal(bytes, &varCsvReader)

	if err != nil {
		return err
	}

	*o = CsvReader(varCsvReader)

	return err
}

type NullableCsvReader struct {
	value *CsvReader
	isSet bool
}

func (v NullableCsvReader) Get() *CsvReader {
	return v.value
}

func (v *NullableCsvReader) Set(val *CsvReader) {
	v.value = val
	v.isSet = true
}

func (v NullableCsvReader) IsSet() bool {
	return v.isSet
}

func (v *NullableCsvReader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsvReader(val *CsvReader) *NullableCsvReader {
	return &NullableCsvReader{value: val, isSet: true}
}

func (v NullableCsvReader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsvReader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


