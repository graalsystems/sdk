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
	"time"
)

// checks if the SearchHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHistory{}

// SearchHistory struct for SearchHistory
type SearchHistory struct {
	Id *string `json:"id,omitempty"`
	Term *string `json:"term,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

// NewSearchHistory instantiates a new SearchHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHistory() *SearchHistory {
	this := SearchHistory{}
	return &this
}

// NewSearchHistoryWithDefaults instantiates a new SearchHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHistoryWithDefaults() *SearchHistory {
	this := SearchHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchHistory) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHistory) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchHistory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchHistory) SetId(v string) {
	o.Id = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SearchHistory) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHistory) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SearchHistory) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SearchHistory) SetTerm(v string) {
	o.Term = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SearchHistory) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHistory) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SearchHistory) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SearchHistory) SetCreated(v time.Time) {
	o.Created = &v
}

func (o SearchHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableSearchHistory struct {
	value *SearchHistory
	isSet bool
}

func (v NullableSearchHistory) Get() *SearchHistory {
	return v.value
}

func (v *NullableSearchHistory) Set(val *SearchHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHistory(val *SearchHistory) *NullableSearchHistory {
	return &NullableSearchHistory{value: val, isSet: true}
}

func (v NullableSearchHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


