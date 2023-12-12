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
	"fmt"
)

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session Object that allows you to conserve an instance of a session.
type Session struct {
	TenantId string `json:"tenant_id"`
	Address *string `json:"address,omitempty"`
	WorkspaceId string `json:"workspace_id"`
	Id string `json:"id"`
	LastTimeUsed time.Time `json:"last_time_used"`
	NotebookId *string `json:"notebook_id,omitempty"`
}

type _Session Session

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(tenantId string, workspaceId string, id string, lastTimeUsed time.Time) *Session {
	this := Session{}
	this.TenantId = tenantId
	var address string = ""
	this.Address = &address
	this.WorkspaceId = workspaceId
	this.Id = id
	this.LastTimeUsed = lastTimeUsed
	var notebookId string = ""
	this.NotebookId = &notebookId
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	var address string = ""
	this.Address = &address
	var notebookId string = ""
	this.NotebookId = &notebookId
	return &this
}

// GetTenantId returns the TenantId field value
func (o *Session) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Session) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Session) SetTenantId(v string) {
	o.TenantId = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Session) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Session) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Session) SetAddress(v string) {
	o.Address = &v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Session) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Session) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Session) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetLastTimeUsed returns the LastTimeUsed field value
func (o *Session) GetLastTimeUsed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastTimeUsed
}

// GetLastTimeUsedOk returns a tuple with the LastTimeUsed field value
// and a boolean to check if the value has been set.
func (o *Session) GetLastTimeUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTimeUsed, true
}

// SetLastTimeUsed sets field value
func (o *Session) SetLastTimeUsed(v time.Time) {
	o.LastTimeUsed = v
}

// GetNotebookId returns the NotebookId field value if set, zero value otherwise.
func (o *Session) GetNotebookId() string {
	if o == nil || IsNil(o.NotebookId) {
		var ret string
		return ret
	}
	return *o.NotebookId
}

// GetNotebookIdOk returns a tuple with the NotebookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetNotebookIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotebookId) {
		return nil, false
	}
	return o.NotebookId, true
}

// HasNotebookId returns a boolean if a field has been set.
func (o *Session) HasNotebookId() bool {
	if o != nil && !IsNil(o.NotebookId) {
		return true
	}

	return false
}

// SetNotebookId gets a reference to the given string and assigns it to the NotebookId field.
func (o *Session) SetNotebookId(v string) {
	o.NotebookId = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenant_id"] = o.TenantId
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["id"] = o.Id
	toSerialize["last_time_used"] = o.LastTimeUsed
	if !IsNil(o.NotebookId) {
		toSerialize["notebook_id"] = o.NotebookId
	}
	return toSerialize, nil
}

func (o *Session) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenant_id",
		"workspace_id",
		"id",
		"last_time_used",
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

	varSession := _Session{}

	err = json.Unmarshal(bytes, &varSession)

	if err != nil {
		return err
	}

	*o = Session(varSession)

	return err
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

