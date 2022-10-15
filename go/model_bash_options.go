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

// BashOptions struct for BashOptions
type BashOptions struct {
	Options
	Type *string `json:"type,omitempty"`
	Lines *[]string `json:"lines,omitempty"`
}

// NewBashOptions instantiates a new BashOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBashOptions() *BashOptions {
	this := BashOptions{}
	var type_ string = "bash"
	this.Type = &type_
	return &this
}

// NewBashOptionsWithDefaults instantiates a new BashOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBashOptionsWithDefaults() *BashOptions {
	this := BashOptions{}
	var type_ string = "bash"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BashOptions) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BashOptions) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BashOptions) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BashOptions) SetType(v string) {
	o.Type = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *BashOptions) GetLines() []string {
	if o == nil || o.Lines == nil {
		var ret []string
		return ret
	}
	return *o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BashOptions) GetLinesOk() (*[]string, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *BashOptions) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given []string and assigns it to the Lines field.
func (o *BashOptions) SetLines(v []string) {
	o.Lines = &v
}

func (o BashOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOptions, errOptions := json.Marshal(o.Options)
	if errOptions != nil {
		return []byte{}, errOptions
	}
	errOptions = json.Unmarshal([]byte(serializedOptions), &toSerialize)
	if errOptions != nil {
		return []byte{}, errOptions
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	return json.Marshal(toSerialize)
}

type NullableBashOptions struct {
	value *BashOptions
	isSet bool
}

func (v NullableBashOptions) Get() *BashOptions {
	return v.value
}

func (v *NullableBashOptions) Set(val *BashOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBashOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBashOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBashOptions(val *BashOptions) *NullableBashOptions {
	return &NullableBashOptions{value: val, isSet: true}
}

func (v NullableBashOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBashOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


