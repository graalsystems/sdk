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

// EndpointPageAllOf struct for EndpointPageAllOf
type EndpointPageAllOf struct {
	Content *[]Endpoint `json:"content,omitempty"`
}

// NewEndpointPageAllOf instantiates a new EndpointPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointPageAllOf() *EndpointPageAllOf {
	this := EndpointPageAllOf{}
	return &this
}

// NewEndpointPageAllOfWithDefaults instantiates a new EndpointPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointPageAllOfWithDefaults() *EndpointPageAllOf {
	this := EndpointPageAllOf{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *EndpointPageAllOf) GetContent() []Endpoint {
	if o == nil || o.Content == nil {
		var ret []Endpoint
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPageAllOf) GetContentOk() (*[]Endpoint, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *EndpointPageAllOf) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []Endpoint and assigns it to the Content field.
func (o *EndpointPageAllOf) SetContent(v []Endpoint) {
	o.Content = &v
}

func (o EndpointPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointPageAllOf struct {
	value *EndpointPageAllOf
	isSet bool
}

func (v NullableEndpointPageAllOf) Get() *EndpointPageAllOf {
	return v.value
}

func (v *NullableEndpointPageAllOf) Set(val *EndpointPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointPageAllOf(val *EndpointPageAllOf) *NullableEndpointPageAllOf {
	return &NullableEndpointPageAllOf{value: val, isSet: true}
}

func (v NullableEndpointPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


