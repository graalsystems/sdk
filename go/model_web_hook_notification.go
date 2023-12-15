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

// checks if the WebHookNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebHookNotification{}

// WebHookNotification struct for WebHookNotification
type WebHookNotification struct {
	Notification
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Method *string `json:"method,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
}

// NewWebHookNotification instantiates a new WebHookNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHookNotification() *WebHookNotification {
	this := WebHookNotification{}
	var type_ string = "webhook"
	this.Type = &type_
	return &this
}

// NewWebHookNotificationWithDefaults instantiates a new WebHookNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookNotificationWithDefaults() *WebHookNotification {
	this := WebHookNotification{}
	var type_ string = "webhook"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WebHookNotification) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotification) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WebHookNotification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WebHookNotification) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebHookNotification) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotification) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebHookNotification) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebHookNotification) SetUrl(v string) {
	o.Url = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WebHookNotification) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotification) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WebHookNotification) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WebHookNotification) SetMethod(v string) {
	o.Method = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *WebHookNotification) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookNotification) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WebHookNotification) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *WebHookNotification) SetHeaders(v map[string]string) {
	o.Headers = &v
}

func (o WebHookNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebHookNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotification, errNotification := json.Marshal(o.Notification)
	if errNotification != nil {
		return map[string]interface{}{}, errNotification
	}
	errNotification = json.Unmarshal([]byte(serializedNotification), &toSerialize)
	if errNotification != nil {
		return map[string]interface{}{}, errNotification
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableWebHookNotification struct {
	value *WebHookNotification
	isSet bool
}

func (v NullableWebHookNotification) Get() *WebHookNotification {
	return v.value
}

func (v *NullableWebHookNotification) Set(val *WebHookNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookNotification(val *WebHookNotification) *NullableWebHookNotification {
	return &NullableWebHookNotification{value: val, isSet: true}
}

func (v NullableWebHookNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


