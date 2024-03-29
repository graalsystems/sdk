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

// checks if the WorkflowActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowActivity{}

// WorkflowActivity Object JobActivity.
type WorkflowActivity struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Id *string `json:"id,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	ReactionCount []Reaction `json:"reaction_count,omitempty"`
	Type *string `json:"type,omitempty"`
	Actor User1 `json:"actor"`
	Target Workflow1 `json:"target"`
	Verb string `json:"verb"`
}

type _WorkflowActivity WorkflowActivity

// NewWorkflowActivity instantiates a new WorkflowActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowActivity(actor User1, target Workflow1, verb string) *WorkflowActivity {
	this := WorkflowActivity{}
	var type_ string = "workflow"
	this.Type = &type_
	this.Actor = actor
	this.Target = target
	this.Verb = verb
	return &this
}

// NewWorkflowActivityWithDefaults instantiates a new WorkflowActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowActivityWithDefaults() *WorkflowActivity {
	this := WorkflowActivity{}
	var type_ string = "workflow"
	this.Type = &type_
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WorkflowActivity) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WorkflowActivity) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *WorkflowActivity) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowActivity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowActivity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowActivity) SetId(v string) {
	o.Id = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *WorkflowActivity) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *WorkflowActivity) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *WorkflowActivity) SetTime(v time.Time) {
	o.Time = &v
}

// GetReactionCount returns the ReactionCount field value if set, zero value otherwise.
func (o *WorkflowActivity) GetReactionCount() []Reaction {
	if o == nil || IsNil(o.ReactionCount) {
		var ret []Reaction
		return ret
	}
	return o.ReactionCount
}

// GetReactionCountOk returns a tuple with the ReactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetReactionCountOk() ([]Reaction, bool) {
	if o == nil || IsNil(o.ReactionCount) {
		return nil, false
	}
	return o.ReactionCount, true
}

// HasReactionCount returns a boolean if a field has been set.
func (o *WorkflowActivity) HasReactionCount() bool {
	if o != nil && !IsNil(o.ReactionCount) {
		return true
	}

	return false
}

// SetReactionCount gets a reference to the given []Reaction and assigns it to the ReactionCount field.
func (o *WorkflowActivity) SetReactionCount(v []Reaction) {
	o.ReactionCount = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowActivity) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowActivity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowActivity) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value
func (o *WorkflowActivity) GetActor() User1 {
	if o == nil {
		var ret User1
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetActorOk() (*User1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *WorkflowActivity) SetActor(v User1) {
	o.Actor = v
}

// GetTarget returns the Target field value
func (o *WorkflowActivity) GetTarget() Workflow1 {
	if o == nil {
		var ret Workflow1
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetTargetOk() (*Workflow1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *WorkflowActivity) SetTarget(v Workflow1) {
	o.Target = v
}

// GetVerb returns the Verb field value
func (o *WorkflowActivity) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *WorkflowActivity) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *WorkflowActivity) SetVerb(v string) {
	o.Verb = v
}

func (o WorkflowActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowActivity) ToMap() (map[string]interface{}, error) {
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

func (o *WorkflowActivity) UnmarshalJSON(bytes []byte) (err error) {
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

	varWorkflowActivity := _WorkflowActivity{}

	err = json.Unmarshal(bytes, &varWorkflowActivity)

	if err != nil {
		return err
	}

	*o = WorkflowActivity(varWorkflowActivity)

	return err
}

type NullableWorkflowActivity struct {
	value *WorkflowActivity
	isSet bool
}

func (v NullableWorkflowActivity) Get() *WorkflowActivity {
	return v.value
}

func (v *NullableWorkflowActivity) Set(val *WorkflowActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowActivity(val *WorkflowActivity) *NullableWorkflowActivity {
	return &NullableWorkflowActivity{value: val, isSet: true}
}

func (v NullableWorkflowActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


