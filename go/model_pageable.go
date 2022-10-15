/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// Pageable minimal Pageable query parameters
type Pageable struct {
	Page *int32 `json:"page,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewPageable instantiates a new Pageable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageable() *Pageable {
	this := Pageable{}
	return &this
}

// NewPageableWithDefaults instantiates a new Pageable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableWithDefaults() *Pageable {
	this := Pageable{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *Pageable) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *Pageable) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *Pageable) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Pageable) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pageable) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Pageable) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Pageable) SetSize(v int32) {
	o.Size = &v
}

func (o Pageable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullablePageable struct {
	value *Pageable
	isSet bool
}

func (v NullablePageable) Get() *Pageable {
	return v.value
}

func (v *NullablePageable) Set(val *Pageable) {
	v.value = val
	v.isSet = true
}

func (v NullablePageable) IsSet() bool {
	return v.isSet
}

func (v *NullablePageable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageable(val *Pageable) *NullablePageable {
	return &NullablePageable{value: val, isSet: true}
}

func (v NullablePageable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

