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

// AzureLogAnalyticsNotificationAllOf struct for AzureLogAnalyticsNotificationAllOf
type AzureLogAnalyticsNotificationAllOf struct {
	Type *string `json:"type,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewAzureLogAnalyticsNotificationAllOf instantiates a new AzureLogAnalyticsNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureLogAnalyticsNotificationAllOf() *AzureLogAnalyticsNotificationAllOf {
	this := AzureLogAnalyticsNotificationAllOf{}
	var type_ string = "eventhub"
	this.Type = &type_
	return &this
}

// NewAzureLogAnalyticsNotificationAllOfWithDefaults instantiates a new AzureLogAnalyticsNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureLogAnalyticsNotificationAllOfWithDefaults() *AzureLogAnalyticsNotificationAllOf {
	this := AzureLogAnalyticsNotificationAllOf{}
	var type_ string = "eventhub"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureLogAnalyticsNotificationAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLogAnalyticsNotificationAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureLogAnalyticsNotificationAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureLogAnalyticsNotificationAllOf) SetType(v string) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AzureLogAnalyticsNotificationAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureLogAnalyticsNotificationAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AzureLogAnalyticsNotificationAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AzureLogAnalyticsNotificationAllOf) SetKey(v string) {
	o.Key = &v
}

func (o AzureLogAnalyticsNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableAzureLogAnalyticsNotificationAllOf struct {
	value *AzureLogAnalyticsNotificationAllOf
	isSet bool
}

func (v NullableAzureLogAnalyticsNotificationAllOf) Get() *AzureLogAnalyticsNotificationAllOf {
	return v.value
}

func (v *NullableAzureLogAnalyticsNotificationAllOf) Set(val *AzureLogAnalyticsNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureLogAnalyticsNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureLogAnalyticsNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureLogAnalyticsNotificationAllOf(val *AzureLogAnalyticsNotificationAllOf) *NullableAzureLogAnalyticsNotificationAllOf {
	return &NullableAzureLogAnalyticsNotificationAllOf{value: val, isSet: true}
}

func (v NullableAzureLogAnalyticsNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureLogAnalyticsNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


