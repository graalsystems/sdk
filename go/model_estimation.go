/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems

import (
	"encoding/json"
)

// checks if the Estimation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Estimation{}

// Estimation struct for Estimation
type Estimation struct {
	Currency *string `json:"currency,omitempty"`
	Amount *string `json:"amount,omitempty"`
	Requests []Request `json:"requests,omitempty"`
}

// NewEstimation instantiates a new Estimation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimation() *Estimation {
	this := Estimation{}
	return &this
}

// NewEstimationWithDefaults instantiates a new Estimation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimationWithDefaults() *Estimation {
	this := Estimation{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Estimation) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Estimation) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Estimation) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Estimation) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Estimation) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Estimation) SetAmount(v string) {
	o.Amount = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *Estimation) GetRequests() []Request {
	if o == nil || IsNil(o.Requests) {
		var ret []Request
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetRequestsOk() ([]Request, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *Estimation) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []Request and assigns it to the Requests field.
func (o *Estimation) SetRequests(v []Request) {
	o.Requests = v
}

func (o Estimation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Estimation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableEstimation struct {
	value *Estimation
	isSet bool
}

func (v NullableEstimation) Get() *Estimation {
	return v.value
}

func (v *NullableEstimation) Set(val *Estimation) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimation) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimation(val *Estimation) *NullableEstimation {
	return &NullableEstimation{value: val, isSet: true}
}

func (v NullableEstimation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


