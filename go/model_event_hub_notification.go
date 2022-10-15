/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// EventHubNotification struct for EventHubNotification
type EventHubNotification struct {
	Notification
	Type *string `json:"type,omitempty"`
	ConnectionString *string `json:"connection_string,omitempty"`
	Data *string `json:"data,omitempty"`
}

// NewEventHubNotification instantiates a new EventHubNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHubNotification() *EventHubNotification {
	this := EventHubNotification{}
	var type_ string = "eventhub"
	this.Type = &type_
	return &this
}

// NewEventHubNotificationWithDefaults instantiates a new EventHubNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHubNotificationWithDefaults() *EventHubNotification {
	this := EventHubNotification{}
	var type_ string = "eventhub"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventHubNotification) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHubNotification) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventHubNotification) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventHubNotification) SetType(v string) {
	o.Type = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *EventHubNotification) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHubNotification) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *EventHubNotification) HasConnectionString() bool {
	if o != nil && o.ConnectionString != nil {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *EventHubNotification) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventHubNotification) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHubNotification) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventHubNotification) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *EventHubNotification) SetData(v string) {
	o.Data = &v
}

func (o EventHubNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNotification, errNotification := json.Marshal(o.Notification)
	if errNotification != nil {
		return []byte{}, errNotification
	}
	errNotification = json.Unmarshal([]byte(serializedNotification), &toSerialize)
	if errNotification != nil {
		return []byte{}, errNotification
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ConnectionString != nil {
		toSerialize["connection_string"] = o.ConnectionString
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEventHubNotification struct {
	value *EventHubNotification
	isSet bool
}

func (v NullableEventHubNotification) Get() *EventHubNotification {
	return v.value
}

func (v *NullableEventHubNotification) Set(val *EventHubNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHubNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHubNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHubNotification(val *EventHubNotification) *NullableEventHubNotification {
	return &NullableEventHubNotification{value: val, isSet: true}
}

func (v NullableEventHubNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHubNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

