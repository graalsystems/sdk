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

// UserActivityAllOf struct for UserActivityAllOf
type UserActivityAllOf struct {
	Type *string `json:"type,omitempty"`
	Actor User1 `json:"actor"`
	Target OneOfUser1Project1 `json:"target"`
	Verb string `json:"verb"`
}

// NewUserActivityAllOf instantiates a new UserActivityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivityAllOf(actor User1, target OneOfUser1Project1, verb string) *UserActivityAllOf {
	this := UserActivityAllOf{}
	var type_ string = "user"
	this.Type = &type_
	this.Actor = actor
	this.Target = target
	this.Verb = verb
	return &this
}

// NewUserActivityAllOfWithDefaults instantiates a new UserActivityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityAllOfWithDefaults() *UserActivityAllOf {
	this := UserActivityAllOf{}
	var type_ string = "user"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserActivityAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserActivityAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserActivityAllOf) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value
func (o *UserActivityAllOf) GetActor() User1 {
	if o == nil {
		var ret User1
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *UserActivityAllOf) GetActorOk() (*User1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *UserActivityAllOf) SetActor(v User1) {
	o.Actor = v
}

// GetTarget returns the Target field value
// If the value is explicit nil, the zero value for OneOfUser1Project1 will be returned
func (o *UserActivityAllOf) GetTarget() OneOfUser1Project1 {
	if o == nil {
		var ret OneOfUser1Project1
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivityAllOf) GetTargetOk() (*OneOfUser1Project1, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *UserActivityAllOf) SetTarget(v OneOfUser1Project1) {
	o.Target = v
}

// GetVerb returns the Verb field value
func (o *UserActivityAllOf) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *UserActivityAllOf) GetVerbOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *UserActivityAllOf) SetVerb(v string) {
	o.Verb = v
}

func (o UserActivityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["verb"] = o.Verb
	}
	return json.Marshal(toSerialize)
}

type NullableUserActivityAllOf struct {
	value *UserActivityAllOf
	isSet bool
}

func (v NullableUserActivityAllOf) Get() *UserActivityAllOf {
	return v.value
}

func (v *NullableUserActivityAllOf) Set(val *UserActivityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivityAllOf(val *UserActivityAllOf) *NullableUserActivityAllOf {
	return &NullableUserActivityAllOf{value: val, isSet: true}
}

func (v NullableUserActivityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


