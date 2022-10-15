/*
 * Tenant API
 *
 * Tenant API
 *
 * API version: 0.0.1
 * Contact: abc@layer.fr
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graalsystems/sdk

import (
	"encoding/json"
	"time"
)

// BlobMetadata struct for BlobMetadata
type BlobMetadata struct {
	Key *string `json:"key,omitempty"`
	Size *string `json:"size,omitempty"`
	Filename *string `json:"filename,omitempty"`
	LastModified *time.Time `json:"last_modified,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
}

// NewBlobMetadata instantiates a new BlobMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobMetadata() *BlobMetadata {
	this := BlobMetadata{}
	return &this
}

// NewBlobMetadataWithDefaults instantiates a new BlobMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobMetadataWithDefaults() *BlobMetadata {
	this := BlobMetadata{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BlobMetadata) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobMetadata) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BlobMetadata) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BlobMetadata) SetKey(v string) {
	o.Key = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BlobMetadata) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobMetadata) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BlobMetadata) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *BlobMetadata) SetSize(v string) {
	o.Size = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *BlobMetadata) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobMetadata) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *BlobMetadata) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *BlobMetadata) SetFilename(v string) {
	o.Filename = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *BlobMetadata) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobMetadata) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *BlobMetadata) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *BlobMetadata) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *BlobMetadata) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobMetadata) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *BlobMetadata) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *BlobMetadata) SetContentType(v string) {
	o.ContentType = &v
}

func (o BlobMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.LastModified != nil {
		toSerialize["last_modified"] = o.LastModified
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableBlobMetadata struct {
	value *BlobMetadata
	isSet bool
}

func (v NullableBlobMetadata) Get() *BlobMetadata {
	return v.value
}

func (v *NullableBlobMetadata) Set(val *BlobMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobMetadata(val *BlobMetadata) *NullableBlobMetadata {
	return &NullableBlobMetadata{value: val, isSet: true}
}

func (v NullableBlobMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


