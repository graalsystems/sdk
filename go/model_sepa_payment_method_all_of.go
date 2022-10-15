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

// SepaPaymentMethodAllOf struct for SepaPaymentMethodAllOf
type SepaPaymentMethodAllOf struct {
	Type *string `json:"type,omitempty"`
	Iban *string `json:"iban,omitempty"`
}

// NewSepaPaymentMethodAllOf instantiates a new SepaPaymentMethodAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSepaPaymentMethodAllOf() *SepaPaymentMethodAllOf {
	this := SepaPaymentMethodAllOf{}
	var type_ string = "sepa"
	this.Type = &type_
	return &this
}

// NewSepaPaymentMethodAllOfWithDefaults instantiates a new SepaPaymentMethodAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSepaPaymentMethodAllOfWithDefaults() *SepaPaymentMethodAllOf {
	this := SepaPaymentMethodAllOf{}
	var type_ string = "sepa"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SepaPaymentMethodAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaPaymentMethodAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SepaPaymentMethodAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SepaPaymentMethodAllOf) SetType(v string) {
	o.Type = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *SepaPaymentMethodAllOf) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SepaPaymentMethodAllOf) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *SepaPaymentMethodAllOf) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *SepaPaymentMethodAllOf) SetIban(v string) {
	o.Iban = &v
}

func (o SepaPaymentMethodAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	return json.Marshal(toSerialize)
}

type NullableSepaPaymentMethodAllOf struct {
	value *SepaPaymentMethodAllOf
	isSet bool
}

func (v NullableSepaPaymentMethodAllOf) Get() *SepaPaymentMethodAllOf {
	return v.value
}

func (v *NullableSepaPaymentMethodAllOf) Set(val *SepaPaymentMethodAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSepaPaymentMethodAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSepaPaymentMethodAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSepaPaymentMethodAllOf(val *SepaPaymentMethodAllOf) *NullableSepaPaymentMethodAllOf {
	return &NullableSepaPaymentMethodAllOf{value: val, isSet: true}
}

func (v NullableSepaPaymentMethodAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSepaPaymentMethodAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


