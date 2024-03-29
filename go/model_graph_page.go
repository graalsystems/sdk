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
)

// checks if the GraphPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphPage{}

// GraphPage typed Page
type GraphPage struct {
	Number *int32 `json:"number,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Content []Graph `json:"content,omitempty"`
}

// NewGraphPage instantiates a new GraphPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphPage() *GraphPage {
	this := GraphPage{}
	return &this
}

// NewGraphPageWithDefaults instantiates a new GraphPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphPageWithDefaults() *GraphPage {
	this := GraphPage{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GraphPage) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphPage) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GraphPage) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GraphPage) SetNumber(v int32) {
	o.Number = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GraphPage) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphPage) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GraphPage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GraphPage) SetSize(v int32) {
	o.Size = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GraphPage) GetContent() []Graph {
	if o == nil || IsNil(o.Content) {
		var ret []Graph
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphPage) GetContentOk() ([]Graph, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GraphPage) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []Graph and assigns it to the Content field.
func (o *GraphPage) SetContent(v []Graph) {
	o.Content = v
}

func (o GraphPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphPage) ToMap() (map[string]interface{}, error) {
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

type NullableGraphPage struct {
	value *GraphPage
	isSet bool
}

func (v NullableGraphPage) Get() *GraphPage {
	return v.value
}

func (v *NullableGraphPage) Set(val *GraphPage) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphPage) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphPage(val *GraphPage) *NullableGraphPage {
	return &NullableGraphPage{value: val, isSet: true}
}

func (v NullableGraphPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


