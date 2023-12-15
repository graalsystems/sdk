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

// checks if the Company type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Company{}

// Company struct for Company
type Company struct {
	Name *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
	City *float32 `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Vat *string `json:"vat,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Company) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Company) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Company) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Company) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Company) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Company) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Company) GetCity() float32 {
	if o == nil || IsNil(o.City) {
		var ret float32
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCityOk() (*float32, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Company) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given float32 and assigns it to the City field.
func (o *Company) SetCity(v float32) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Company) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Company) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Company) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Company) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Company) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Company) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetVat returns the Vat field value if set, zero value otherwise.
func (o *Company) GetVat() string {
	if o == nil || IsNil(o.Vat) {
		var ret string
		return ret
	}
	return *o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetVatOk() (*string, bool) {
	if o == nil || IsNil(o.Vat) {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *Company) HasVat() bool {
	if o != nil && !IsNil(o.Vat) {
		return true
	}

	return false
}

// SetVat gets a reference to the given string and assigns it to the Vat field.
func (o *Company) SetVat(v string) {
	o.Vat = &v
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Company) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Vat) {
		toSerialize["vat"] = o.Vat
	}
	return toSerialize, nil
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


