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
)

// WebHookNotificationAllOf struct for WebHookNotificationAllOf
type WebHookNotificationAllOf struct {
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Method *string `json:"method,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
}

// NewWebHookNotificationAllOf instantiates a new WebHookNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHookNotificationAllOf() *WebHookNotificationAllOf {
	this := WebHookNotificationAllOf{}
	var type_ string = "webhook"
	this.Type = &type_
	return &this
}

// NewWebHookNotificationAllOfWithDefaults instantiates a new WebHookNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookNotificationAllOfWithDefaults() *WebHookNotificationAllOf {
	this := WebHookNotificationAllOf{}
	var type_ string = "webhook"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WebHookNotificationAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotificationAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WebHookNotificationAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WebHookNotificationAllOf) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebHookNotificationAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotificationAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebHookNotificationAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebHookNotificationAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WebHookNotificationAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotificationAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WebHookNotificationAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WebHookNotificationAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *WebHookNotificationAllOf) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotificationAllOf) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WebHookNotificationAllOf) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *WebHookNotificationAllOf) SetHeaders(v map[string]string) {
	o.Headers = &v
}

func (o WebHookNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

type NullableWebHookNotificationAllOf struct {
	value *WebHookNotificationAllOf
	isSet bool
}

func (v NullableWebHookNotificationAllOf) Get() *WebHookNotificationAllOf {
	return v.value
}

func (v *NullableWebHookNotificationAllOf) Set(val *WebHookNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookNotificationAllOf(val *WebHookNotificationAllOf) *NullableWebHookNotificationAllOf {
	return &NullableWebHookNotificationAllOf{value: val, isSet: true}
}

func (v NullableWebHookNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


