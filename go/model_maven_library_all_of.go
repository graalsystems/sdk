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
)

// MavenLibraryAllOf struct for MavenLibraryAllOf
type MavenLibraryAllOf struct {
	Type *string `json:"type,omitempty"`
	Repo *string `json:"repo,omitempty"`
	Dependency *string `json:"dependency,omitempty"`
}

// NewMavenLibraryAllOf instantiates a new MavenLibraryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenLibraryAllOf() *MavenLibraryAllOf {
	this := MavenLibraryAllOf{}
	var type_ string = "maven"
	this.Type = &type_
	return &this
}

// NewMavenLibraryAllOfWithDefaults instantiates a new MavenLibraryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenLibraryAllOfWithDefaults() *MavenLibraryAllOf {
	this := MavenLibraryAllOf{}
	var type_ string = "maven"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MavenLibraryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenLibraryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MavenLibraryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MavenLibraryAllOf) SetType(v string) {
	o.Type = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *MavenLibraryAllOf) GetRepo() string {
	if o == nil || o.Repo == nil {
		var ret string
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenLibraryAllOf) GetRepoOk() (*string, bool) {
	if o == nil || o.Repo == nil {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *MavenLibraryAllOf) HasRepo() bool {
	if o != nil && o.Repo != nil {
		return true
	}

	return false
}

// SetRepo gets a reference to the given string and assigns it to the Repo field.
func (o *MavenLibraryAllOf) SetRepo(v string) {
	o.Repo = &v
}

// GetDependency returns the Dependency field value if set, zero value otherwise.
func (o *MavenLibraryAllOf) GetDependency() string {
	if o == nil || o.Dependency == nil {
		var ret string
		return ret
	}
	return *o.Dependency
}

// GetDependencyOk returns a tuple with the Dependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenLibraryAllOf) GetDependencyOk() (*string, bool) {
	if o == nil || o.Dependency == nil {
		return nil, false
	}
	return o.Dependency, true
}

// HasDependency returns a boolean if a field has been set.
func (o *MavenLibraryAllOf) HasDependency() bool {
	if o != nil && o.Dependency != nil {
		return true
	}

	return false
}

// SetDependency gets a reference to the given string and assigns it to the Dependency field.
func (o *MavenLibraryAllOf) SetDependency(v string) {
	o.Dependency = &v
}

func (o MavenLibraryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Repo != nil {
		toSerialize["repo"] = o.Repo
	}
	if o.Dependency != nil {
		toSerialize["dependency"] = o.Dependency
	}
	return json.Marshal(toSerialize)
}

type NullableMavenLibraryAllOf struct {
	value *MavenLibraryAllOf
	isSet bool
}

func (v NullableMavenLibraryAllOf) Get() *MavenLibraryAllOf {
	return v.value
}

func (v *NullableMavenLibraryAllOf) Set(val *MavenLibraryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenLibraryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenLibraryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenLibraryAllOf(val *MavenLibraryAllOf) *NullableMavenLibraryAllOf {
	return &NullableMavenLibraryAllOf{value: val, isSet: true}
}

func (v NullableMavenLibraryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenLibraryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


