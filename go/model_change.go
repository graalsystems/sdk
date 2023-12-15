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
)

// checks if the Change type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Change{}

// Change struct for Change
type Change struct {
	Id *string `json:"id,omitempty"`
	Author *string `json:"author,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Executed *time.Time `json:"executed,omitempty"`
	Type *string `json:"type,omitempty"`
	Order *float32 `json:"order,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// NewChange instantiates a new Change object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChange() *Change {
	this := Change{}
	return &this
}

// NewChangeWithDefaults instantiates a new Change object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWithDefaults() *Change {
	this := Change{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Change) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Change) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Change) SetId(v string) {
	o.Id = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Change) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Change) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *Change) SetAuthor(v string) {
	o.Author = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *Change) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *Change) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *Change) SetFilename(v string) {
	o.Filename = &v
}

// GetExecuted returns the Executed field value if set, zero value otherwise.
func (o *Change) GetExecuted() time.Time {
	if o == nil || IsNil(o.Executed) {
		var ret time.Time
		return ret
	}
	return *o.Executed
}

// GetExecutedOk returns a tuple with the Executed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetExecutedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Executed) {
		return nil, false
	}
	return o.Executed, true
}

// HasExecuted returns a boolean if a field has been set.
func (o *Change) HasExecuted() bool {
	if o != nil && !IsNil(o.Executed) {
		return true
	}

	return false
}

// SetExecuted gets a reference to the given time.Time and assigns it to the Executed field.
func (o *Change) SetExecuted(v time.Time) {
	o.Executed = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Change) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Change) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Change) SetType(v string) {
	o.Type = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Change) GetOrder() float32 {
	if o == nil || IsNil(o.Order) {
		var ret float32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Change) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given float32 and assigns it to the Order field.
func (o *Change) SetOrder(v float32) {
	o.Order = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Change) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Change) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Change) SetHash(v string) {
	o.Hash = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Change) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Change) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Change) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Change) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Change) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Change) SetComments(v string) {
	o.Comments = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Change) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Change) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *Change) SetTag(v string) {
	o.Tag = &v
}

func (o Change) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Change) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Executed) {
		toSerialize["executed"] = o.Executed
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableChange struct {
	value *Change
	isSet bool
}

func (v NullableChange) Get() *Change {
	return v.value
}

func (v *NullableChange) Set(val *Change) {
	v.value = val
	v.isSet = true
}

func (v NullableChange) IsSet() bool {
	return v.isSet
}

func (v *NullableChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChange(val *Change) *NullableChange {
	return &NullableChange{value: val, isSet: true}
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


