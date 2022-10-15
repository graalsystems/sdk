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
	"time"
)

// Application struct for Application
type Application struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Name *string `json:"name,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	RemoteApplicationId *string `json:"remote_application_id,omitempty"`
	RemoteApplicationSecret *string `json:"remote_application_secret,omitempty"`
	Description *string `json:"description,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Metadata *map[string]map[string]interface{} `json:"metadata,omitempty"`
	Locked *bool `json:"locked,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication() *Application {
	this := Application{}
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Application) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Application) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Application) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Application) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Application) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Application) SetUserId(v string) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Application) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Application) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Application) SetName(v string) {
	o.Name = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *Application) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *Application) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *Application) SetTenantId(v string) {
	o.TenantId = &v
}

// GetRemoteApplicationId returns the RemoteApplicationId field value if set, zero value otherwise.
func (o *Application) GetRemoteApplicationId() string {
	if o == nil || o.RemoteApplicationId == nil {
		var ret string
		return ret
	}
	return *o.RemoteApplicationId
}

// GetRemoteApplicationIdOk returns a tuple with the RemoteApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRemoteApplicationIdOk() (*string, bool) {
	if o == nil || o.RemoteApplicationId == nil {
		return nil, false
	}
	return o.RemoteApplicationId, true
}

// HasRemoteApplicationId returns a boolean if a field has been set.
func (o *Application) HasRemoteApplicationId() bool {
	if o != nil && o.RemoteApplicationId != nil {
		return true
	}

	return false
}

// SetRemoteApplicationId gets a reference to the given string and assigns it to the RemoteApplicationId field.
func (o *Application) SetRemoteApplicationId(v string) {
	o.RemoteApplicationId = &v
}

// GetRemoteApplicationSecret returns the RemoteApplicationSecret field value if set, zero value otherwise.
func (o *Application) GetRemoteApplicationSecret() string {
	if o == nil || o.RemoteApplicationSecret == nil {
		var ret string
		return ret
	}
	return *o.RemoteApplicationSecret
}

// GetRemoteApplicationSecretOk returns a tuple with the RemoteApplicationSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRemoteApplicationSecretOk() (*string, bool) {
	if o == nil || o.RemoteApplicationSecret == nil {
		return nil, false
	}
	return o.RemoteApplicationSecret, true
}

// HasRemoteApplicationSecret returns a boolean if a field has been set.
func (o *Application) HasRemoteApplicationSecret() bool {
	if o != nil && o.RemoteApplicationSecret != nil {
		return true
	}

	return false
}

// SetRemoteApplicationSecret gets a reference to the given string and assigns it to the RemoteApplicationSecret field.
func (o *Application) SetRemoteApplicationSecret(v string) {
	o.RemoteApplicationSecret = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Application) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Application) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Application) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Application) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Application) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Application) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Application) GetMetadata() map[string]map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMetadataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Application) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *Application) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *Application) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *Application) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *Application) SetLocked(v bool) {
	o.Locked = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.RemoteApplicationId != nil {
		toSerialize["remote_application_id"] = o.RemoteApplicationId
	}
	if o.RemoteApplicationSecret != nil {
		toSerialize["remote_application_secret"] = o.RemoteApplicationSecret
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

