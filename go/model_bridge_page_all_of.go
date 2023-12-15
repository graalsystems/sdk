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

// BridgePageAllOf struct for BridgePageAllOf
type BridgePageAllOf struct {
	Content *[]Bridge `json:"content,omitempty"`
}

// NewBridgePageAllOf instantiates a new BridgePageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgePageAllOf() *BridgePageAllOf {
	this := BridgePageAllOf{}
	return &this
}

// NewBridgePageAllOfWithDefaults instantiates a new BridgePageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgePageAllOfWithDefaults() *BridgePageAllOf {
	this := BridgePageAllOf{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BridgePageAllOf) GetContent() []Bridge {
	if o == nil || o.Content == nil {
		var ret []Bridge
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BridgePageAllOf) GetContentOk() (*[]Bridge, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BridgePageAllOf) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []Bridge and assigns it to the Content field.
func (o *BridgePageAllOf) SetContent(v []Bridge) {
	o.Content = &v
}

func (o BridgePageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableBridgePageAllOf struct {
	value *BridgePageAllOf
	isSet bool
}

func (v NullableBridgePageAllOf) Get() *BridgePageAllOf {
	return v.value
}

func (v *NullableBridgePageAllOf) Set(val *BridgePageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgePageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgePageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgePageAllOf(val *BridgePageAllOf) *NullableBridgePageAllOf {
	return &NullableBridgePageAllOf{value: val, isSet: true}
}

func (v NullableBridgePageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgePageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


