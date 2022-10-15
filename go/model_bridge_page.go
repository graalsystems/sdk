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

// BridgePage typed Page
type BridgePage struct {
	Number *int32 `json:"number,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Content *[]Bridge `json:"content,omitempty"`
}

// NewBridgePage instantiates a new BridgePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgePage() *BridgePage {
	this := BridgePage{}
	return &this
}

// NewBridgePageWithDefaults instantiates a new BridgePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgePageWithDefaults() *BridgePage {
	this := BridgePage{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *BridgePage) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgePage) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *BridgePage) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *BridgePage) SetNumber(v int32) {
	o.Number = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BridgePage) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgePage) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BridgePage) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *BridgePage) SetSize(v int32) {
	o.Size = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BridgePage) GetContent() []Bridge {
	if o == nil || o.Content == nil {
		var ret []Bridge
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgePage) GetContentOk() (*[]Bridge, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BridgePage) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []Bridge and assigns it to the Content field.
func (o *BridgePage) SetContent(v []Bridge) {
	o.Content = &v
}

func (o BridgePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableBridgePage struct {
	value *BridgePage
	isSet bool
}

func (v NullableBridgePage) Get() *BridgePage {
	return v.value
}

func (v *NullableBridgePage) Set(val *BridgePage) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgePage) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgePage(val *BridgePage) *NullableBridgePage {
	return &NullableBridgePage{value: val, isSet: true}
}

func (v NullableBridgePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


