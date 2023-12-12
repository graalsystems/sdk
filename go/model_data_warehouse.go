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

// checks if the DataWarehouse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataWarehouse{}

// DataWarehouse struct for DataWarehouse
type DataWarehouse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
	Options *DataWarehouseOptions `json:"options,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
	IdentityId *string `json:"identity_id,omitempty"`
	InfrastructureId *string `json:"infrastructure_id,omitempty"`
	PublicUrl *string `json:"public_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Metadata map[string]map[string]interface{} `json:"metadata,omitempty"`
}

// NewDataWarehouse instantiates a new DataWarehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataWarehouse() *DataWarehouse {
	this := DataWarehouse{}
	return &this
}

// NewDataWarehouseWithDefaults instantiates a new DataWarehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWarehouseWithDefaults() *DataWarehouse {
	this := DataWarehouse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataWarehouse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataWarehouse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataWarehouse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataWarehouse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataWarehouse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataWarehouse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DataWarehouse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DataWarehouse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DataWarehouse) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DataWarehouse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DataWarehouse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DataWarehouse) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DataWarehouse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DataWarehouse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DataWarehouse) SetVersion(v string) {
	o.Version = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DataWarehouse) GetOptions() DataWarehouseOptions {
	if o == nil || IsNil(o.Options) {
		var ret DataWarehouseOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetOptionsOk() (*DataWarehouseOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DataWarehouse) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DataWarehouseOptions and assigns it to the Options field.
func (o *DataWarehouse) SetOptions(v DataWarehouseOptions) {
	o.Options = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DataWarehouse) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DataWarehouse) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DataWarehouse) SetOwner(v string) {
	o.Owner = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DataWarehouse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DataWarehouse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DataWarehouse) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DataWarehouse) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DataWarehouse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DataWarehouse) SetUsername(v string) {
	o.Username = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *DataWarehouse) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *DataWarehouse) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *DataWarehouse) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *DataWarehouse) GetInfrastructureId() string {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret string
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetInfrastructureIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *DataWarehouse) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given string and assigns it to the InfrastructureId field.
func (o *DataWarehouse) SetInfrastructureId(v string) {
	o.InfrastructureId = &v
}

// GetPublicUrl returns the PublicUrl field value if set, zero value otherwise.
func (o *DataWarehouse) GetPublicUrl() string {
	if o == nil || IsNil(o.PublicUrl) {
		var ret string
		return ret
	}
	return *o.PublicUrl
}

// GetPublicUrlOk returns a tuple with the PublicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetPublicUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PublicUrl) {
		return nil, false
	}
	return o.PublicUrl, true
}

// HasPublicUrl returns a boolean if a field has been set.
func (o *DataWarehouse) HasPublicUrl() bool {
	if o != nil && !IsNil(o.PublicUrl) {
		return true
	}

	return false
}

// SetPublicUrl gets a reference to the given string and assigns it to the PublicUrl field.
func (o *DataWarehouse) SetPublicUrl(v string) {
	o.PublicUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataWarehouse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataWarehouse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DataWarehouse) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DataWarehouse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DataWarehouse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DataWarehouse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *DataWarehouse) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *DataWarehouse) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *DataWarehouse) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DataWarehouse) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DataWarehouse) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *DataWarehouse) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DataWarehouse) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWarehouse) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DataWarehouse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *DataWarehouse) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

func (o DataWarehouse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataWarehouse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructure_id"] = o.InfrastructureId
	}
	if !IsNil(o.PublicUrl) {
		toSerialize["public_url"] = o.PublicUrl
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
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableDataWarehouse struct {
	value *DataWarehouse
	isSet bool
}

func (v NullableDataWarehouse) Get() *DataWarehouse {
	return v.value
}

func (v *NullableDataWarehouse) Set(val *DataWarehouse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataWarehouse(val *DataWarehouse) *NullableDataWarehouse {
	return &NullableDataWarehouse{value: val, isSet: true}
}

func (v NullableDataWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


