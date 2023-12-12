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

// checks if the CommentActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentActivity{}

// CommentActivity struct for CommentActivity
type CommentActivity struct {
	Comment *string `json:"comment,omitempty"`
	Id *string `json:"id,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// NewCommentActivity instantiates a new CommentActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentActivity() *CommentActivity {
	this := CommentActivity{}
	return &this
}

// NewCommentActivityWithDefaults instantiates a new CommentActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentActivityWithDefaults() *CommentActivity {
	this := CommentActivity{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CommentActivity) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentActivity) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CommentActivity) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CommentActivity) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommentActivity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentActivity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommentActivity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommentActivity) SetId(v string) {
	o.Id = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *CommentActivity) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentActivity) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *CommentActivity) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *CommentActivity) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *CommentActivity) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentActivity) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *CommentActivity) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *CommentActivity) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CommentActivity) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentActivity) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CommentActivity) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CommentActivity) SetUserId(v string) {
	o.UserId = &v
}

func (o CommentActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["last_update"] = o.LastUpdate
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableCommentActivity struct {
	value *CommentActivity
	isSet bool
}

func (v NullableCommentActivity) Get() *CommentActivity {
	return v.value
}

func (v *NullableCommentActivity) Set(val *CommentActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentActivity(val *CommentActivity) *NullableCommentActivity {
	return &NullableCommentActivity{value: val, isSet: true}
}

func (v NullableCommentActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

