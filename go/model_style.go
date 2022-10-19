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
	"time"
)

// Style struct for Style
type Style struct {
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
}

// NewStyle instantiates a new Style object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStyle() *Style {
	this := Style{}
	return &this
}

// NewStyleWithDefaults instantiates a new Style object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStyleWithDefaults() *Style {
	this := Style{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Style) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Style) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Style) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Style) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Style) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Style) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Style) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Style) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Style) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Style) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Style) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Style) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

func (o Style) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableStyle struct {
	value *Style
	isSet bool
}

func (v NullableStyle) Get() *Style {
	return v.value
}

func (v *NullableStyle) Set(val *Style) {
	v.value = val
	v.isSet = true
}

func (v NullableStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStyle(val *Style) *NullableStyle {
	return &NullableStyle{value: val, isSet: true}
}

func (v NullableStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


