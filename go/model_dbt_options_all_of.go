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

// DbtOptionsAllOf struct for DbtOptionsAllOf
type DbtOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	Profile *string `json:"profile,omitempty"`
	Adapter *string `json:"adapter,omitempty"`
}

// NewDbtOptionsAllOf instantiates a new DbtOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbtOptionsAllOf() *DbtOptionsAllOf {
	this := DbtOptionsAllOf{}
	var type_ string = "dbt"
	this.Type = &type_
	return &this
}

// NewDbtOptionsAllOfWithDefaults instantiates a new DbtOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbtOptionsAllOfWithDefaults() *DbtOptionsAllOf {
	this := DbtOptionsAllOf{}
	var type_ string = "dbt"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DbtOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbtOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DbtOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DbtOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *DbtOptionsAllOf) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbtOptionsAllOf) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *DbtOptionsAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *DbtOptionsAllOf) SetProfile(v string) {
	o.Profile = &v
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *DbtOptionsAllOf) GetAdapter() string {
	if o == nil || o.Adapter == nil {
		var ret string
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbtOptionsAllOf) GetAdapterOk() (*string, bool) {
	if o == nil || o.Adapter == nil {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *DbtOptionsAllOf) HasAdapter() bool {
	if o != nil && o.Adapter != nil {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given string and assigns it to the Adapter field.
func (o *DbtOptionsAllOf) SetAdapter(v string) {
	o.Adapter = &v
}

func (o DbtOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Adapter != nil {
		toSerialize["adapter"] = o.Adapter
	}
	return json.Marshal(toSerialize)
}

type NullableDbtOptionsAllOf struct {
	value *DbtOptionsAllOf
	isSet bool
}

func (v NullableDbtOptionsAllOf) Get() *DbtOptionsAllOf {
	return v.value
}

func (v *NullableDbtOptionsAllOf) Set(val *DbtOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDbtOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDbtOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbtOptionsAllOf(val *DbtOptionsAllOf) *NullableDbtOptionsAllOf {
	return &NullableDbtOptionsAllOf{value: val, isSet: true}
}

func (v NullableDbtOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbtOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

