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

// checks if the Notifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notifications{}

// Notifications struct for Notifications
type Notifications struct {
	OnStart   []INotification `json:"on_start,omitempty"`
	OnSuccess []INotification `json:"on_success,omitempty"`
	OnFailure []INotification `json:"on_failure,omitempty"`
}

// NewNotifications instantiates a new Notifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifications() *Notifications {
	this := Notifications{}
	return &this
}

// NewNotificationsWithDefaults instantiates a new Notifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsWithDefaults() *Notifications {
	this := Notifications{}
	return &this
}

// GetOnStart returns the OnStart field value if set, zero value otherwise.
func (o *Notifications) GetOnStart() []INotification {
	if o == nil || IsNil(o.OnStart) {
		var ret []INotification
		return ret
	}
	return o.OnStart
}

// GetOnStartOk returns a tuple with the OnStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notifications) GetOnStartOk() ([]INotification, bool) {
	if o == nil || IsNil(o.OnStart) {
		return nil, false
	}
	return o.OnStart, true
}

// HasOnStart returns a boolean if a field has been set.
func (o *Notifications) HasOnStart() bool {
	if o != nil && !IsNil(o.OnStart) {
		return true
	}

	return false
}

// SetOnStart gets a reference to the given []INotification and assigns it to the OnStart field.
func (o *Notifications) SetOnStart(v []INotification) {
	o.OnStart = v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *Notifications) GetOnSuccess() []INotification {
	if o == nil || IsNil(o.OnSuccess) {
		var ret []INotification
		return ret
	}
	return o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notifications) GetOnSuccessOk() ([]INotification, bool) {
	if o == nil || IsNil(o.OnSuccess) {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *Notifications) HasOnSuccess() bool {
	if o != nil && !IsNil(o.OnSuccess) {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given []INotification and assigns it to the OnSuccess field.
func (o *Notifications) SetOnSuccess(v []INotification) {
	o.OnSuccess = v
}

// GetOnFailure returns the OnFailure field value if set, zero value otherwise.
func (o *Notifications) GetOnFailure() []INotification {
	if o == nil || IsNil(o.OnFailure) {
		var ret []INotification
		return ret
	}
	return o.OnFailure
}

// GetOnFailureOk returns a tuple with the OnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notifications) GetOnFailureOk() ([]INotification, bool) {
	if o == nil || IsNil(o.OnFailure) {
		return nil, false
	}
	return o.OnFailure, true
}

// HasOnFailure returns a boolean if a field has been set.
func (o *Notifications) HasOnFailure() bool {
	if o != nil && !IsNil(o.OnFailure) {
		return true
	}

	return false
}

// SetOnFailure gets a reference to the given []INotification and assigns it to the OnFailure field.
func (o *Notifications) SetOnFailure(v []INotification) {
	o.OnFailure = v
}

func (o Notifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnStart) {
		toSerialize["on_start"] = o.OnStart
	}
	if !IsNil(o.OnSuccess) {
		toSerialize["on_success"] = o.OnSuccess
	}
	if !IsNil(o.OnFailure) {
		toSerialize["on_failure"] = o.OnFailure
	}
	return toSerialize, nil
}

type NullableNotifications struct {
	value *Notifications
	isSet bool
}

func (v NullableNotifications) Get() *Notifications {
	return v.value
}

func (v *NullableNotifications) Set(val *Notifications) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifications(val *Notifications) *NullableNotifications {
	return &NullableNotifications{value: val, isSet: true}
}

func (v NullableNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
