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

// S3UriAllOf struct for S3UriAllOf
type S3UriAllOf struct {
	Type *string `json:"type,omitempty"`
	AccessKey *string `json:"access_key,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewS3UriAllOf instantiates a new S3UriAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3UriAllOf() *S3UriAllOf {
	this := S3UriAllOf{}
	var type_ string = "s3"
	this.Type = &type_
	return &this
}

// NewS3UriAllOfWithDefaults instantiates a new S3UriAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3UriAllOfWithDefaults() *S3UriAllOf {
	this := S3UriAllOf{}
	var type_ string = "s3"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *S3UriAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3UriAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *S3UriAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *S3UriAllOf) SetType(v string) {
	o.Type = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *S3UriAllOf) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3UriAllOf) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *S3UriAllOf) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *S3UriAllOf) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *S3UriAllOf) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3UriAllOf) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *S3UriAllOf) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *S3UriAllOf) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *S3UriAllOf) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3UriAllOf) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *S3UriAllOf) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *S3UriAllOf) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *S3UriAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3UriAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *S3UriAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *S3UriAllOf) SetRegion(v string) {
	o.Region = &v
}

func (o S3UriAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableS3UriAllOf struct {
	value *S3UriAllOf
	isSet bool
}

func (v NullableS3UriAllOf) Get() *S3UriAllOf {
	return v.value
}

func (v *NullableS3UriAllOf) Set(val *S3UriAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableS3UriAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableS3UriAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3UriAllOf(val *S3UriAllOf) *NullableS3UriAllOf {
	return &NullableS3UriAllOf{value: val, isSet: true}
}

func (v NullableS3UriAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3UriAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


