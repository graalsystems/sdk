/*
Tenant API

Tenant API

API version: 0.0.1
Contact: abc@layer.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems

import (
	"encoding/json"
	"time"
)

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity Object Activity.
type Activity struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Id *string `json:"id,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	ReactionCount []Reaction `json:"reaction_count,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Activity) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Activity) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Activity) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Activity) SetId(v string) {
	o.Id = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Activity) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Activity) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *Activity) SetTime(v time.Time) {
	o.Time = &v
}

// GetReactionCount returns the ReactionCount field value if set, zero value otherwise.
func (o *Activity) GetReactionCount() []Reaction {
	if o == nil || IsNil(o.ReactionCount) {
		var ret []Reaction
		return ret
	}
	return o.ReactionCount
}

// GetReactionCountOk returns a tuple with the ReactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetReactionCountOk() ([]Reaction, bool) {
	if o == nil || IsNil(o.ReactionCount) {
		return nil, false
	}
	return o.ReactionCount, true
}

// HasReactionCount returns a boolean if a field has been set.
func (o *Activity) HasReactionCount() bool {
	if o != nil && !IsNil(o.ReactionCount) {
		return true
	}

	return false
}

// SetReactionCount gets a reference to the given []Reaction and assigns it to the ReactionCount field.
func (o *Activity) SetReactionCount(v []Reaction) {
	o.ReactionCount = v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


