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

// HadoopOptionsAllOf struct for HadoopOptionsAllOf
type HadoopOptionsAllOf struct {
	Type *string `json:"type,omitempty"`
	ApplicationType *string `json:"application_type,omitempty"`
	ApplicationId *string `json:"application_id,omitempty"`
	JobId *string `json:"job_id,omitempty"`
	Command *string `json:"command,omitempty"`
}

// NewHadoopOptionsAllOf instantiates a new HadoopOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHadoopOptionsAllOf() *HadoopOptionsAllOf {
	this := HadoopOptionsAllOf{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// NewHadoopOptionsAllOfWithDefaults instantiates a new HadoopOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHadoopOptionsAllOfWithDefaults() *HadoopOptionsAllOf {
	this := HadoopOptionsAllOf{}
	var type_ string = "dask"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HadoopOptionsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptionsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HadoopOptionsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HadoopOptionsAllOf) SetType(v string) {
	o.Type = &v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *HadoopOptionsAllOf) GetApplicationType() string {
	if o == nil || o.ApplicationType == nil {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptionsAllOf) GetApplicationTypeOk() (*string, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *HadoopOptionsAllOf) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *HadoopOptionsAllOf) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *HadoopOptionsAllOf) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptionsAllOf) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *HadoopOptionsAllOf) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *HadoopOptionsAllOf) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *HadoopOptionsAllOf) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptionsAllOf) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *HadoopOptionsAllOf) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *HadoopOptionsAllOf) SetJobId(v string) {
	o.JobId = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *HadoopOptionsAllOf) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HadoopOptionsAllOf) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *HadoopOptionsAllOf) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *HadoopOptionsAllOf) SetCommand(v string) {
	o.Command = &v
}

func (o HadoopOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ApplicationType != nil {
		toSerialize["application_type"] = o.ApplicationType
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	return json.Marshal(toSerialize)
}

type NullableHadoopOptionsAllOf struct {
	value *HadoopOptionsAllOf
	isSet bool
}

func (v NullableHadoopOptionsAllOf) Get() *HadoopOptionsAllOf {
	return v.value
}

func (v *NullableHadoopOptionsAllOf) Set(val *HadoopOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHadoopOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHadoopOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHadoopOptionsAllOf(val *HadoopOptionsAllOf) *NullableHadoopOptionsAllOf {
	return &NullableHadoopOptionsAllOf{value: val, isSet: true}
}

func (v NullableHadoopOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHadoopOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


