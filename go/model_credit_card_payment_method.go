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

// CreditCardPaymentMethod struct for CreditCardPaymentMethod
type CreditCardPaymentMethod struct {
	PaymentMethod
	Type *string `json:"type,omitempty"`
	Number *string `json:"number,omitempty"`
	Cvc *string `json:"cvc,omitempty"`
	ExpirationMonth *string `json:"expiration_month,omitempty"`
	ExpirationYear *string `json:"expiration_year,omitempty"`
}

// NewCreditCardPaymentMethod instantiates a new CreditCardPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardPaymentMethod() *CreditCardPaymentMethod {
	this := CreditCardPaymentMethod{}
	var type_ string = "credit_card"
	this.Type = &type_
	return &this
}

// NewCreditCardPaymentMethodWithDefaults instantiates a new CreditCardPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardPaymentMethodWithDefaults() *CreditCardPaymentMethod {
	this := CreditCardPaymentMethod{}
	var type_ string = "credit_card"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreditCardPaymentMethod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreditCardPaymentMethod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreditCardPaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CreditCardPaymentMethod) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardPaymentMethod) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CreditCardPaymentMethod) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CreditCardPaymentMethod) SetNumber(v string) {
	o.Number = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *CreditCardPaymentMethod) GetCvc() string {
	if o == nil || o.Cvc == nil {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardPaymentMethod) GetCvcOk() (*string, bool) {
	if o == nil || o.Cvc == nil {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *CreditCardPaymentMethod) HasCvc() bool {
	if o != nil && o.Cvc != nil {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *CreditCardPaymentMethod) SetCvc(v string) {
	o.Cvc = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *CreditCardPaymentMethod) GetExpirationMonth() string {
	if o == nil || o.ExpirationMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardPaymentMethod) GetExpirationMonthOk() (*string, bool) {
	if o == nil || o.ExpirationMonth == nil {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *CreditCardPaymentMethod) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth != nil {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *CreditCardPaymentMethod) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *CreditCardPaymentMethod) GetExpirationYear() string {
	if o == nil || o.ExpirationYear == nil {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCardPaymentMethod) GetExpirationYearOk() (*string, bool) {
	if o == nil || o.ExpirationYear == nil {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *CreditCardPaymentMethod) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear != nil {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *CreditCardPaymentMethod) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

func (o CreditCardPaymentMethod) MarshalJSON() ([]byte, error) {
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
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Cvc != nil {
		toSerialize["cvc"] = o.Cvc
	}
	if o.ExpirationMonth != nil {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationYear != nil {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	return json.Marshal(toSerialize)
}

type NullableCreditCardPaymentMethod struct {
	value *CreditCardPaymentMethod
	isSet bool
}

func (v NullableCreditCardPaymentMethod) Get() *CreditCardPaymentMethod {
	return v.value
}

func (v *NullableCreditCardPaymentMethod) Set(val *CreditCardPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardPaymentMethod(val *CreditCardPaymentMethod) *NullableCreditCardPaymentMethod {
	return &NullableCreditCardPaymentMethod{value: val, isSet: true}
}

func (v NullableCreditCardPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


