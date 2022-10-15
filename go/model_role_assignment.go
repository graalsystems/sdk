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

// RoleAssignment struct for RoleAssignment
type RoleAssignment struct {
	Id *string `json:"id,omitempty"`
	ResourceId *string `json:"resource_id,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	PrincipalType *string `json:"principal_type,omitempty"`
	PrincipalId *string `json:"principal_id,omitempty"`
	RoleId *string `json:"role_id,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Metadata *map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewRoleAssignment instantiates a new RoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignment() *RoleAssignment {
	this := RoleAssignment{}
	return &this
}

// NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentWithDefaults() *RoleAssignment {
	this := RoleAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignment) SetId(v string) {
	o.Id = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *RoleAssignment) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *RoleAssignment) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *RoleAssignment) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *RoleAssignment) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *RoleAssignment) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *RoleAssignment) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetPrincipalType returns the PrincipalType field value if set, zero value otherwise.
func (o *RoleAssignment) GetPrincipalType() string {
	if o == nil || o.PrincipalType == nil {
		var ret string
		return ret
	}
	return *o.PrincipalType
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetPrincipalTypeOk() (*string, bool) {
	if o == nil || o.PrincipalType == nil {
		return nil, false
	}
	return o.PrincipalType, true
}

// HasPrincipalType returns a boolean if a field has been set.
func (o *RoleAssignment) HasPrincipalType() bool {
	if o != nil && o.PrincipalType != nil {
		return true
	}

	return false
}

// SetPrincipalType gets a reference to the given string and assigns it to the PrincipalType field.
func (o *RoleAssignment) SetPrincipalType(v string) {
	o.PrincipalType = &v
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise.
func (o *RoleAssignment) GetPrincipalId() string {
	if o == nil || o.PrincipalId == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetPrincipalIdOk() (*string, bool) {
	if o == nil || o.PrincipalId == nil {
		return nil, false
	}
	return o.PrincipalId, true
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *RoleAssignment) HasPrincipalId() bool {
	if o != nil && o.PrincipalId != nil {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given string and assigns it to the PrincipalId field.
func (o *RoleAssignment) SetPrincipalId(v string) {
	o.PrincipalId = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleAssignment) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleAssignment) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleAssignment) SetRoleId(v string) {
	o.RoleId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *RoleAssignment) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *RoleAssignment) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *RoleAssignment) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RoleAssignment) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RoleAssignment) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RoleAssignment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *RoleAssignment) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *RoleAssignment) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *RoleAssignment) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RoleAssignment) GetMetadata() map[string]map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetMetadataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RoleAssignment) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *RoleAssignment) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = &v
}

func (o RoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.PrincipalType != nil {
		toSerialize["principal_type"] = o.PrincipalType
	}
	if o.PrincipalId != nil {
		toSerialize["principal_id"] = o.PrincipalId
	}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
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
	return json.Marshal(toSerialize)
}

type NullableRoleAssignment struct {
	value *RoleAssignment
	isSet bool
}

func (v NullableRoleAssignment) Get() *RoleAssignment {
	return v.value
}

func (v *NullableRoleAssignment) Set(val *RoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignment(val *RoleAssignment) *NullableRoleAssignment {
	return &NullableRoleAssignment{value: val, isSet: true}
}

func (v NullableRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


