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

// checks if the GitLibrary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitLibrary{}

// GitLibrary struct for GitLibrary
type GitLibrary struct {
	Library
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Path *string `json:"path,omitempty"`
	Revision *string `json:"revision,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewGitLibrary instantiates a new GitLibrary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitLibrary() *GitLibrary {
	this := GitLibrary{}
	var type_ string = "git"
	this.Type = &type_
	return &this
}

// NewGitLibraryWithDefaults instantiates a new GitLibrary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitLibraryWithDefaults() *GitLibrary {
	this := GitLibrary{}
	var type_ string = "git"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GitLibrary) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GitLibrary) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GitLibrary) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GitLibrary) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GitLibrary) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GitLibrary) SetUrl(v string) {
	o.Url = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GitLibrary) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GitLibrary) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GitLibrary) SetPath(v string) {
	o.Path = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *GitLibrary) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *GitLibrary) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *GitLibrary) SetRevision(v string) {
	o.Revision = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GitLibrary) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GitLibrary) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GitLibrary) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GitLibrary) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitLibrary) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GitLibrary) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GitLibrary) SetPassword(v string) {
	o.Password = &v
}

func (o GitLibrary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitLibrary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLibrary, errLibrary := json.Marshal(o.Library)
	if errLibrary != nil {
		return map[string]interface{}{}, errLibrary
	}
	errLibrary = json.Unmarshal([]byte(serializedLibrary), &toSerialize)
	if errLibrary != nil {
		return map[string]interface{}{}, errLibrary
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableGitLibrary struct {
	value *GitLibrary
	isSet bool
}

func (v NullableGitLibrary) Get() *GitLibrary {
	return v.value
}

func (v *NullableGitLibrary) Set(val *GitLibrary) {
	v.value = val
	v.isSet = true
}

func (v NullableGitLibrary) IsSet() bool {
	return v.isSet
}

func (v *NullableGitLibrary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitLibrary(val *GitLibrary) *NullableGitLibrary {
	return &NullableGitLibrary{value: val, isSet: true}
}

func (v NullableGitLibrary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitLibrary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


