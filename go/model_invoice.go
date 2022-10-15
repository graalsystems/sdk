/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

import (
	"encoding/json"
	"time"
)

// Invoice struct for Invoice
type Invoice struct {
	Id *string `json:"id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	StartPeriod *time.Time `json:"start_period,omitempty"`
	EndPeriod *time.Time `json:"end_period,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Amount *string `json:"amount,omitempty"`
	Metadata *map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice() *Invoice {
	this := Invoice{}
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invoice) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invoice) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invoice) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Invoice) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Invoice) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Invoice) SetCreated(v time.Time) {
	o.Created = &v
}

// GetStartPeriod returns the StartPeriod field value if set, zero value otherwise.
func (o *Invoice) GetStartPeriod() time.Time {
	if o == nil || o.StartPeriod == nil {
		var ret time.Time
		return ret
	}
	return *o.StartPeriod
}

// GetStartPeriodOk returns a tuple with the StartPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetStartPeriodOk() (*time.Time, bool) {
	if o == nil || o.StartPeriod == nil {
		return nil, false
	}
	return o.StartPeriod, true
}

// HasStartPeriod returns a boolean if a field has been set.
func (o *Invoice) HasStartPeriod() bool {
	if o != nil && o.StartPeriod != nil {
		return true
	}

	return false
}

// SetStartPeriod gets a reference to the given time.Time and assigns it to the StartPeriod field.
func (o *Invoice) SetStartPeriod(v time.Time) {
	o.StartPeriod = &v
}

// GetEndPeriod returns the EndPeriod field value if set, zero value otherwise.
func (o *Invoice) GetEndPeriod() time.Time {
	if o == nil || o.EndPeriod == nil {
		var ret time.Time
		return ret
	}
	return *o.EndPeriod
}

// GetEndPeriodOk returns a tuple with the EndPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetEndPeriodOk() (*time.Time, bool) {
	if o == nil || o.EndPeriod == nil {
		return nil, false
	}
	return o.EndPeriod, true
}

// HasEndPeriod returns a boolean if a field has been set.
func (o *Invoice) HasEndPeriod() bool {
	if o != nil && o.EndPeriod != nil {
		return true
	}

	return false
}

// SetEndPeriod gets a reference to the given time.Time and assigns it to the EndPeriod field.
func (o *Invoice) SetEndPeriod(v time.Time) {
	o.EndPeriod = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Invoice) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Invoice) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Invoice) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Invoice) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Invoice) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Invoice) SetAmount(v string) {
	o.Amount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Invoice) GetMetadata() map[string]map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetMetadataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Invoice) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *Invoice) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = &v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.StartPeriod != nil {
		toSerialize["start_period"] = o.StartPeriod
	}
	if o.EndPeriod != nil {
		toSerialize["end_period"] = o.EndPeriod
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


