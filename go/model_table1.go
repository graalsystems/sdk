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

// checks if the Table1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Table1{}

// Table1 struct for Table1
type Table1 struct {
	Id *string `json:"id,omitempty"`
	TechnicalName *string `json:"technical_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Uri *string `json:"uri,omitempty"`
	PublicUri *string `json:"public_uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Metadata map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewTable1 instantiates a new Table1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTable1() *Table1 {
	this := Table1{}
	return &this
}

// NewTable1WithDefaults instantiates a new Table1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTable1WithDefaults() *Table1 {
	this := Table1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Table1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Table1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Table1) SetId(v string) {
	o.Id = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *Table1) GetTechnicalName() string {
	if o == nil || IsNil(o.TechnicalName) {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetTechnicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalName) {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *Table1) HasTechnicalName() bool {
	if o != nil && !IsNil(o.TechnicalName) {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *Table1) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Table1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Table1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Table1) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Table1) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Table1) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Table1) SetUri(v string) {
	o.Uri = &v
}

// GetPublicUri returns the PublicUri field value if set, zero value otherwise.
func (o *Table1) GetPublicUri() string {
	if o == nil || IsNil(o.PublicUri) {
		var ret string
		return ret
	}
	return *o.PublicUri
}

// GetPublicUriOk returns a tuple with the PublicUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetPublicUriOk() (*string, bool) {
	if o == nil || IsNil(o.PublicUri) {
		return nil, false
	}
	return o.PublicUri, true
}

// HasPublicUri returns a boolean if a field has been set.
func (o *Table1) HasPublicUri() bool {
	if o != nil && !IsNil(o.PublicUri) {
		return true
	}

	return false
}

// SetPublicUri gets a reference to the given string and assigns it to the PublicUri field.
func (o *Table1) SetPublicUri(v string) {
	o.PublicUri = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Table1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Table1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Table1) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Table1) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Table1) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Table1) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Table1) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Table1) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Table1) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Table1) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Table1) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Table1) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *Table1) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

func (o Table1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Table1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TechnicalName) {
		toSerialize["technical_name"] = o.TechnicalName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.PublicUri) {
		toSerialize["public_uri"] = o.PublicUri
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableTable1 struct {
	value *Table1
	isSet bool
}

func (v NullableTable1) Get() *Table1 {
	return v.value
}

func (v *NullableTable1) Set(val *Table1) {
	v.value = val
	v.isSet = true
}

func (v NullableTable1) IsSet() bool {
	return v.isSet
}

func (v *NullableTable1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTable1(val *Table1) *NullableTable1 {
	return &NullableTable1{value: val, isSet: true}
}

func (v NullableTable1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTable1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

