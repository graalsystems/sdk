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

// checks if the WorkflowRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRun{}

// WorkflowRun struct for WorkflowRun
type WorkflowRun struct {
	Id *string `json:"id,omitempty"`
	InfrastructureId *string `json:"infrastructure_id,omitempty"`
	DeviceId *string `json:"device_id,omitempty"`
	WorkflowId *string `json:"workflow_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Metadata map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewWorkflowRun instantiates a new WorkflowRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRun() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunWithDefaults() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowRun) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowRun) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowRun) SetId(v string) {
	o.Id = &v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *WorkflowRun) GetInfrastructureId() string {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret string
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetInfrastructureIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *WorkflowRun) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given string and assigns it to the InfrastructureId field.
func (o *WorkflowRun) SetInfrastructureId(v string) {
	o.InfrastructureId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *WorkflowRun) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *WorkflowRun) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *WorkflowRun) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowRun) GetWorkflowId() string {
	if o == nil || IsNil(o.WorkflowId) {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowId) {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowRun) HasWorkflowId() bool {
	if o != nil && !IsNil(o.WorkflowId) {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *WorkflowRun) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *WorkflowRun) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *WorkflowRun) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *WorkflowRun) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRun) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRun) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRun) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowRun) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowRun) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowRun) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRun) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRun) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRun) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WorkflowRun) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WorkflowRun) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WorkflowRun) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *WorkflowRun) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *WorkflowRun) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *WorkflowRun) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WorkflowRun) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WorkflowRun) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *WorkflowRun) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

func (o WorkflowRun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructure_id"] = o.InfrastructureId
	}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.WorkflowId) {
		toSerialize["workflow_id"] = o.WorkflowId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

type NullableWorkflowRun struct {
	value *WorkflowRun
	isSet bool
}

func (v NullableWorkflowRun) Get() *WorkflowRun {
	return v.value
}

func (v *NullableWorkflowRun) Set(val *WorkflowRun) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRun) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRun(val *WorkflowRun) *NullableWorkflowRun {
	return &NullableWorkflowRun{value: val, isSet: true}
}

func (v NullableWorkflowRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


