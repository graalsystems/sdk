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

// Dependency struct for Dependency
type Dependency struct {
	Id *string `json:"id,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewDependency instantiates a new Dependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependency() *Dependency {
	this := Dependency{}
	return &this
}

// NewDependencyWithDefaults instantiates a new Dependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependencyWithDefaults() *Dependency {
	this := Dependency{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dependency) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dependency) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dependency) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Dependency) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Dependency) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Dependency) SetReference(v string) {
	o.Reference = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Dependency) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Dependency) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Dependency) SetVersion(v string) {
	o.Version = &v
}

func (o Dependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDependency struct {
	value *Dependency
	isSet bool
}

func (v NullableDependency) Get() *Dependency {
	return v.value
}

func (v *NullableDependency) Set(val *Dependency) {
	v.value = val
	v.isSet = true
}

func (v NullableDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependency(val *Dependency) *NullableDependency {
	return &NullableDependency{value: val, isSet: true}
}

func (v NullableDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

