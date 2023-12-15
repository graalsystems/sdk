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

// ProjectActivityAllOf struct for ProjectActivityAllOf
type ProjectActivityAllOf struct {
	Type *string `json:"type,omitempty"`
	Actor User1 `json:"actor"`
	Target Project1 `json:"target"`
	Verb string `json:"verb"`
}

// NewProjectActivityAllOf instantiates a new ProjectActivityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectActivityAllOf(actor User1, target Project1, verb string) *ProjectActivityAllOf {
	this := ProjectActivityAllOf{}
	var type_ string = "project"
	this.Type = &type_
	this.Actor = actor
	this.Target = target
	this.Verb = verb
	return &this
}

// NewProjectActivityAllOfWithDefaults instantiates a new ProjectActivityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectActivityAllOfWithDefaults() *ProjectActivityAllOf {
	this := ProjectActivityAllOf{}
	var type_ string = "project"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectActivityAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivityAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectActivityAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectActivityAllOf) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value
func (o *ProjectActivityAllOf) GetActor() User1 {
	if o == nil {
		var ret User1
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *ProjectActivityAllOf) GetActorOk() (*User1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *ProjectActivityAllOf) SetActor(v User1) {
	o.Actor = v
}

// GetTarget returns the Target field value
func (o *ProjectActivityAllOf) GetTarget() Project1 {
	if o == nil {
		var ret Project1
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ProjectActivityAllOf) GetTargetOk() (*Project1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ProjectActivityAllOf) SetTarget(v Project1) {
	o.Target = v
}

// GetVerb returns the Verb field value
func (o *ProjectActivityAllOf) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *ProjectActivityAllOf) GetVerbOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *ProjectActivityAllOf) SetVerb(v string) {
	o.Verb = v
}

func (o ProjectActivityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["verb"] = o.Verb
	}
	return json.Marshal(toSerialize)
}

type NullableProjectActivityAllOf struct {
	value *ProjectActivityAllOf
	isSet bool
}

func (v NullableProjectActivityAllOf) Get() *ProjectActivityAllOf {
	return v.value
}

func (v *NullableProjectActivityAllOf) Set(val *ProjectActivityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectActivityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectActivityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectActivityAllOf(val *ProjectActivityAllOf) *NullableProjectActivityAllOf {
	return &NullableProjectActivityAllOf{value: val, isSet: true}
}

func (v NullableProjectActivityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectActivityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


