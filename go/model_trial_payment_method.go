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

// TrialPaymentMethod struct for TrialPaymentMethod
type TrialPaymentMethod struct {
	PaymentMethod
	Type *string `json:"type,omitempty"`
	Expire *time.Time `json:"expire,omitempty"`
}

// NewTrialPaymentMethod instantiates a new TrialPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialPaymentMethod() *TrialPaymentMethod {
	this := TrialPaymentMethod{}
	var type_ string = "trial"
	this.Type = &type_
	return &this
}

// NewTrialPaymentMethodWithDefaults instantiates a new TrialPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialPaymentMethodWithDefaults() *TrialPaymentMethod {
	this := TrialPaymentMethod{}
	var type_ string = "trial"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrialPaymentMethod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrialPaymentMethod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TrialPaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *TrialPaymentMethod) GetExpire() time.Time {
	if o == nil || o.Expire == nil {
		var ret time.Time
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialPaymentMethod) GetExpireOk() (*time.Time, bool) {
	if o == nil || o.Expire == nil {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *TrialPaymentMethod) HasExpire() bool {
	if o != nil && o.Expire != nil {
		return true
	}

	return false
}

// SetExpire gets a reference to the given time.Time and assigns it to the Expire field.
func (o *TrialPaymentMethod) SetExpire(v time.Time) {
	o.Expire = &v
}

func (o TrialPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPaymentMethod, errPaymentMethod := json.Marshal(o.PaymentMethod)
	if errPaymentMethod != nil {
		return []byte{}, errPaymentMethod
	}
	errPaymentMethod = json.Unmarshal([]byte(serializedPaymentMethod), &toSerialize)
	if errPaymentMethod != nil {
		return []byte{}, errPaymentMethod
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Expire != nil {
		toSerialize["expire"] = o.Expire
	}
	return json.Marshal(toSerialize)
}

type NullableTrialPaymentMethod struct {
	value *TrialPaymentMethod
	isSet bool
}

func (v NullableTrialPaymentMethod) Get() *TrialPaymentMethod {
	return v.value
}

func (v *NullableTrialPaymentMethod) Set(val *TrialPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialPaymentMethod(val *TrialPaymentMethod) *NullableTrialPaymentMethod {
	return &NullableTrialPaymentMethod{value: val, isSet: true}
}

func (v NullableTrialPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


