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

// SubscriptionDetailsAllOf struct for SubscriptionDetailsAllOf
type SubscriptionDetailsAllOf struct {
	Type *string `json:"type,omitempty"`
	SupportPlan *string `json:"support_plan,omitempty"`
	Account *string `json:"account,omitempty"`
	Company *Company `json:"company,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

// NewSubscriptionDetailsAllOf instantiates a new SubscriptionDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDetailsAllOf() *SubscriptionDetailsAllOf {
	this := SubscriptionDetailsAllOf{}
	var type_ string = "subscription"
	this.Type = &type_
	return &this
}

// NewSubscriptionDetailsAllOfWithDefaults instantiates a new SubscriptionDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDetailsAllOfWithDefaults() *SubscriptionDetailsAllOf {
	this := SubscriptionDetailsAllOf{}
	var type_ string = "subscription"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SubscriptionDetailsAllOf) SetType(v string) {
	o.Type = &v
}

// GetSupportPlan returns the SupportPlan field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetSupportPlan() string {
	if o == nil || o.SupportPlan == nil {
		var ret string
		return ret
	}
	return *o.SupportPlan
}

// GetSupportPlanOk returns a tuple with the SupportPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetSupportPlanOk() (*string, bool) {
	if o == nil || o.SupportPlan == nil {
		return nil, false
	}
	return o.SupportPlan, true
}

// HasSupportPlan returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasSupportPlan() bool {
	if o != nil && o.SupportPlan != nil {
		return true
	}

	return false
}

// SetSupportPlan gets a reference to the given string and assigns it to the SupportPlan field.
func (o *SubscriptionDetailsAllOf) SetSupportPlan(v string) {
	o.SupportPlan = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *SubscriptionDetailsAllOf) SetAccount(v string) {
	o.Account = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetCompany() Company {
	if o == nil || o.Company == nil {
		var ret Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetCompanyOk() (*Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given Company and assigns it to the Company field.
func (o *SubscriptionDetailsAllOf) SetCompany(v Company) {
	o.Company = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetContact() Contact {
	if o == nil || o.Contact == nil {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetContactOk() (*Contact, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *SubscriptionDetailsAllOf) SetContact(v Contact) {
	o.Contact = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *SubscriptionDetailsAllOf) GetPaymentMethod() PaymentMethod {
	if o == nil || o.PaymentMethod == nil {
		var ret PaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDetailsAllOf) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *SubscriptionDetailsAllOf) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethod and assigns it to the PaymentMethod field.
func (o *SubscriptionDetailsAllOf) SetPaymentMethod(v PaymentMethod) {
	o.PaymentMethod = &v
}

func (o SubscriptionDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SupportPlan != nil {
		toSerialize["support_plan"] = o.SupportPlan
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionDetailsAllOf struct {
	value *SubscriptionDetailsAllOf
	isSet bool
}

func (v NullableSubscriptionDetailsAllOf) Get() *SubscriptionDetailsAllOf {
	return v.value
}

func (v *NullableSubscriptionDetailsAllOf) Set(val *SubscriptionDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDetailsAllOf(val *SubscriptionDetailsAllOf) *NullableSubscriptionDetailsAllOf {
	return &NullableSubscriptionDetailsAllOf{value: val, isSet: true}
}

func (v NullableSubscriptionDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


