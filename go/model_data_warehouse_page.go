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
)

// checks if the DataWarehousePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataWarehousePage{}

// DataWarehousePage typed Page
type DataWarehousePage struct {
	Number *int32 `json:"number,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Content []DataWarehouse `json:"content,omitempty"`
}

// NewDataWarehousePage instantiates a new DataWarehousePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataWarehousePage() *DataWarehousePage {
	this := DataWarehousePage{}
	return &this
}

// NewDataWarehousePageWithDefaults instantiates a new DataWarehousePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWarehousePageWithDefaults() *DataWarehousePage {
	this := DataWarehousePage{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *DataWarehousePage) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehousePage) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *DataWarehousePage) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *DataWarehousePage) SetNumber(v int32) {
	o.Number = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DataWarehousePage) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehousePage) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DataWarehousePage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *DataWarehousePage) SetSize(v int32) {
	o.Size = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DataWarehousePage) GetContent() []DataWarehouse {
	if o == nil || IsNil(o.Content) {
		var ret []DataWarehouse
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehousePage) GetContentOk() ([]DataWarehouse, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DataWarehousePage) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []DataWarehouse and assigns it to the Content field.
func (o *DataWarehousePage) SetContent(v []DataWarehouse) {
	o.Content = v
}

func (o DataWarehousePage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataWarehousePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableDataWarehousePage struct {
	value *DataWarehousePage
	isSet bool
}

func (v NullableDataWarehousePage) Get() *DataWarehousePage {
	return v.value
}

func (v *NullableDataWarehousePage) Set(val *DataWarehousePage) {
	v.value = val
	v.isSet = true
}

func (v NullableDataWarehousePage) IsSet() bool {
	return v.isSet
}

func (v *NullableDataWarehousePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataWarehousePage(val *DataWarehousePage) *NullableDataWarehousePage {
	return &NullableDataWarehousePage{value: val, isSet: true}
}

func (v NullableDataWarehousePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataWarehousePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


