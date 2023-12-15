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

// CronScheduleAllOf struct for CronScheduleAllOf
type CronScheduleAllOf struct {
	Type *string `json:"type,omitempty"`
	CronExpression *string `json:"cron_expression,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	InfrastructureId *string `json:"infrastructure_id,omitempty"`
	DeviceId *string `json:"device_id,omitempty"`
}

// NewCronScheduleAllOf instantiates a new CronScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronScheduleAllOf() *CronScheduleAllOf {
	this := CronScheduleAllOf{}
	var type_ string = "cron"
	this.Type = &type_
	return &this
}

// NewCronScheduleAllOfWithDefaults instantiates a new CronScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronScheduleAllOfWithDefaults() *CronScheduleAllOf {
	this := CronScheduleAllOf{}
	var type_ string = "cron"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CronScheduleAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronScheduleAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CronScheduleAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CronScheduleAllOf) SetType(v string) {
	o.Type = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *CronScheduleAllOf) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronScheduleAllOf) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *CronScheduleAllOf) HasCronExpression() bool {
	if o != nil && o.CronExpression != nil {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *CronScheduleAllOf) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CronScheduleAllOf) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronScheduleAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CronScheduleAllOf) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CronScheduleAllOf) SetTimezone(v string) {
	o.Timezone = &v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *CronScheduleAllOf) GetInfrastructureId() string {
	if o == nil || o.InfrastructureId == nil {
		var ret string
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronScheduleAllOf) GetInfrastructureIdOk() (*string, bool) {
	if o == nil || o.InfrastructureId == nil {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *CronScheduleAllOf) HasInfrastructureId() bool {
	if o != nil && o.InfrastructureId != nil {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given string and assigns it to the InfrastructureId field.
func (o *CronScheduleAllOf) SetInfrastructureId(v string) {
	o.InfrastructureId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *CronScheduleAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronScheduleAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *CronScheduleAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *CronScheduleAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o CronScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CronExpression != nil {
		toSerialize["cron_expression"] = o.CronExpression
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.InfrastructureId != nil {
		toSerialize["infrastructure_id"] = o.InfrastructureId
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableCronScheduleAllOf struct {
	value *CronScheduleAllOf
	isSet bool
}

func (v NullableCronScheduleAllOf) Get() *CronScheduleAllOf {
	return v.value
}

func (v *NullableCronScheduleAllOf) Set(val *CronScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCronScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCronScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronScheduleAllOf(val *CronScheduleAllOf) *NullableCronScheduleAllOf {
	return &NullableCronScheduleAllOf{value: val, isSet: true}
}

func (v NullableCronScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


