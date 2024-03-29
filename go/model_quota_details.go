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

// checks if the QuotaDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaDetails{}

// QuotaDetails struct for QuotaDetails
type QuotaDetails struct {
	Details
	Type *string `json:"type,omitempty"`
	Scope *string `json:"scope,omitempty"`
	NewValue *float32 `json:"new_value,omitempty"`
}

// NewQuotaDetails instantiates a new QuotaDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaDetails() *QuotaDetails {
	this := QuotaDetails{}
	var type_ string = "quota"
	this.Type = &type_
	return &this
}

// NewQuotaDetailsWithDefaults instantiates a new QuotaDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaDetailsWithDefaults() *QuotaDetails {
	this := QuotaDetails{}
	var type_ string = "quota"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QuotaDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QuotaDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QuotaDetails) SetType(v string) {
	o.Type = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *QuotaDetails) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDetails) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *QuotaDetails) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *QuotaDetails) SetScope(v string) {
	o.Scope = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *QuotaDetails) GetNewValue() float32 {
	if o == nil || IsNil(o.NewValue) {
		var ret float32
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDetails) GetNewValueOk() (*float32, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *QuotaDetails) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given float32 and assigns it to the NewValue field.
func (o *QuotaDetails) SetNewValue(v float32) {
	o.NewValue = &v
}

func (o QuotaDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDetails, errDetails := json.Marshal(o.Details)
	if errDetails != nil {
		return map[string]interface{}{}, errDetails
	}
	errDetails = json.Unmarshal([]byte(serializedDetails), &toSerialize)
	if errDetails != nil {
		return map[string]interface{}{}, errDetails
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.NewValue) {
		toSerialize["new_value"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableQuotaDetails struct {
	value *QuotaDetails
	isSet bool
}

func (v NullableQuotaDetails) Get() *QuotaDetails {
	return v.value
}

func (v *NullableQuotaDetails) Set(val *QuotaDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaDetails(val *QuotaDetails) *NullableQuotaDetails {
	return &NullableQuotaDetails{value: val, isSet: true}
}

func (v NullableQuotaDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


