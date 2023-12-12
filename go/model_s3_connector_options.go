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
	"fmt"
)

// checks if the S3ConnectorOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3ConnectorOptions{}

// S3ConnectorOptions Options to connect to a S3 bucket.  Attributes ---------- bucket : str     Bucket name. access_key : str     Access key for S3. secret_key : str     Secret key for S3. session_token : str     Session token for S3. endpoint : str     URL endpoint.
type S3ConnectorOptions struct {
	// The bucket to be connected to.
	Bucket string `json:"bucket"`
	// S3 access key to access to AWS.
	AccessKey string `json:"access_key"`
	// S3 secret key to access to AWS.
	SecretKey string `json:"secret_key"`
	// S3 session token to access to AWS.
	SessionToken *string `json:"session_token,omitempty"`
	// The endpoint to be connected to.
	Endpoint *string `json:"endpoint,omitempty"`
	// The region of the endpoint.
	RegionName *string `json:"region_name,omitempty"`
}

type _S3ConnectorOptions S3ConnectorOptions

// NewS3ConnectorOptions instantiates a new S3ConnectorOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ConnectorOptions(bucket string, accessKey string, secretKey string) *S3ConnectorOptions {
	this := S3ConnectorOptions{}
	this.Bucket = bucket
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	return &this
}

// NewS3ConnectorOptionsWithDefaults instantiates a new S3ConnectorOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ConnectorOptionsWithDefaults() *S3ConnectorOptions {
	this := S3ConnectorOptions{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *S3ConnectorOptions) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *S3ConnectorOptions) SetBucket(v string) {
	o.Bucket = v
}

// GetAccessKey returns the AccessKey field value
func (o *S3ConnectorOptions) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *S3ConnectorOptions) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *S3ConnectorOptions) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *S3ConnectorOptions) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *S3ConnectorOptions) GetSessionToken() string {
	if o == nil || IsNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetSessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *S3ConnectorOptions) HasSessionToken() bool {
	if o != nil && !IsNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *S3ConnectorOptions) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *S3ConnectorOptions) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *S3ConnectorOptions) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *S3ConnectorOptions) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *S3ConnectorOptions) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ConnectorOptions) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *S3ConnectorOptions) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *S3ConnectorOptions) SetRegionName(v string) {
	o.RegionName = &v
}

func (o S3ConnectorOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3ConnectorOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	toSerialize["access_key"] = o.AccessKey
	toSerialize["secret_key"] = o.SecretKey
	if !IsNil(o.SessionToken) {
		toSerialize["session_token"] = o.SessionToken
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.RegionName) {
		toSerialize["region_name"] = o.RegionName
	}
	return toSerialize, nil
}

func (o *S3ConnectorOptions) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
		"access_key",
		"secret_key",
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

	varS3ConnectorOptions := _S3ConnectorOptions{}

	err = json.Unmarshal(bytes, &varS3ConnectorOptions)

	if err != nil {
		return err
	}

	*o = S3ConnectorOptions(varS3ConnectorOptions)

	return err
}

type NullableS3ConnectorOptions struct {
	value *S3ConnectorOptions
	isSet bool
}

func (v NullableS3ConnectorOptions) Get() *S3ConnectorOptions {
	return v.value
}

func (v *NullableS3ConnectorOptions) Set(val *S3ConnectorOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ConnectorOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ConnectorOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ConnectorOptions(val *S3ConnectorOptions) *NullableS3ConnectorOptions {
	return &NullableS3ConnectorOptions{value: val, isSet: true}
}

func (v NullableS3ConnectorOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ConnectorOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


