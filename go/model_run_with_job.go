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

// RunWithJob struct for RunWithJob
type RunWithJob struct {
	Run *JobRun `json:"run,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewRunWithJob instantiates a new RunWithJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunWithJob() *RunWithJob {
	this := RunWithJob{}
	return &this
}

// NewRunWithJobWithDefaults instantiates a new RunWithJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithJobWithDefaults() *RunWithJob {
	this := RunWithJob{}
	return &this
}

// GetRun returns the Run field value if set, zero value otherwise.
func (o *RunWithJob) GetRun() JobRun {
	if o == nil || o.Run == nil {
		var ret JobRun
		return ret
	}
	return *o.Run
}

// GetRunOk returns a tuple with the Run field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWithJob) GetRunOk() (*JobRun, bool) {
	if o == nil || o.Run == nil {
		return nil, false
	}
	return o.Run, true
}

// HasRun returns a boolean if a field has been set.
func (o *RunWithJob) HasRun() bool {
	if o != nil && o.Run != nil {
		return true
	}

	return false
}

// SetRun gets a reference to the given JobRun and assigns it to the Run field.
func (o *RunWithJob) SetRun(v JobRun) {
	o.Run = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RunWithJob) GetJob() Job {
	if o == nil || o.Job == nil {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunWithJob) GetJobOk() (*Job, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RunWithJob) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RunWithJob) SetJob(v Job) {
	o.Job = &v
}

func (o RunWithJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Run != nil {
		toSerialize["run"] = o.Run
	}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableRunWithJob struct {
	value *RunWithJob
	isSet bool
}

func (v NullableRunWithJob) Get() *RunWithJob {
	return v.value
}

func (v *NullableRunWithJob) Set(val *RunWithJob) {
	v.value = val
	v.isSet = true
}

func (v NullableRunWithJob) IsSet() bool {
	return v.isSet
}

func (v *NullableRunWithJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunWithJob(val *RunWithJob) *NullableRunWithJob {
	return &NullableRunWithJob{value: val, isSet: true}
}

func (v NullableRunWithJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunWithJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


