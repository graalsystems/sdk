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
	"fmt"
)

// checks if the ProjectActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectActivity{}

// ProjectActivity Object ProjectActivity.
type ProjectActivity struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Id *string `json:"id,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	ReactionCount []Reaction `json:"reaction_count,omitempty"`
	Type *string `json:"type,omitempty"`
	Actor User1 `json:"actor"`
	Target Project1 `json:"target"`
	Verb string `json:"verb"`
}

type _ProjectActivity ProjectActivity

// NewProjectActivity instantiates a new ProjectActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectActivity(actor User1, target Project1, verb string) *ProjectActivity {
	this := ProjectActivity{}
	var type_ string = "project"
	this.Type = &type_
	this.Actor = actor
	this.Target = target
	this.Verb = verb
	return &this
}

// NewProjectActivityWithDefaults instantiates a new ProjectActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectActivityWithDefaults() *ProjectActivity {
	this := ProjectActivity{}
	var type_ string = "project"
	this.Type = &type_
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProjectActivity) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectActivity) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ProjectActivity) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectActivity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectActivity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectActivity) SetId(v string) {
	o.Id = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ProjectActivity) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ProjectActivity) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ProjectActivity) SetTime(v time.Time) {
	o.Time = &v
}

// GetReactionCount returns the ReactionCount field value if set, zero value otherwise.
func (o *ProjectActivity) GetReactionCount() []Reaction {
	if o == nil || IsNil(o.ReactionCount) {
		var ret []Reaction
		return ret
	}
	return o.ReactionCount
}

// GetReactionCountOk returns a tuple with the ReactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetReactionCountOk() ([]Reaction, bool) {
	if o == nil || IsNil(o.ReactionCount) {
		return nil, false
	}
	return o.ReactionCount, true
}

// HasReactionCount returns a boolean if a field has been set.
func (o *ProjectActivity) HasReactionCount() bool {
	if o != nil && !IsNil(o.ReactionCount) {
		return true
	}

	return false
}

// SetReactionCount gets a reference to the given []Reaction and assigns it to the ReactionCount field.
func (o *ProjectActivity) SetReactionCount(v []Reaction) {
	o.ReactionCount = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectActivity) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectActivity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectActivity) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value
func (o *ProjectActivity) GetActor() User1 {
	if o == nil {
		var ret User1
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetActorOk() (*User1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *ProjectActivity) SetActor(v User1) {
	o.Actor = v
}

// GetTarget returns the Target field value
func (o *ProjectActivity) GetTarget() Project1 {
	if o == nil {
		var ret Project1
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetTargetOk() (*Project1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ProjectActivity) SetTarget(v Project1) {
	o.Target = v
}

// GetVerb returns the Verb field value
func (o *ProjectActivity) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *ProjectActivity) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *ProjectActivity) SetVerb(v string) {
	o.Verb = v
}

func (o ProjectActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ReactionCount) {
		toSerialize["reaction_count"] = o.ReactionCount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["actor"] = o.Actor
	toSerialize["target"] = o.Target
	toSerialize["verb"] = o.Verb
	return toSerialize, nil
}

func (o *ProjectActivity) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actor",
		"target",
		"verb",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectActivity := _ProjectActivity{}

	err = json.Unmarshal(bytes, &varProjectActivity)

	if err != nil {
		return err
	}

	*o = ProjectActivity(varProjectActivity)

	return err
}

type NullableProjectActivity struct {
	value *ProjectActivity
	isSet bool
}

func (v NullableProjectActivity) Get() *ProjectActivity {
	return v.value
}

func (v *NullableProjectActivity) Set(val *ProjectActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectActivity(val *ProjectActivity) *NullableProjectActivity {
	return &NullableProjectActivity{value: val, isSet: true}
}

func (v NullableProjectActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


