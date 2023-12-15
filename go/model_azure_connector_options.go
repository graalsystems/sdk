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
	"fmt"
)

// checks if the AzureConnectorOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureConnectorOptions{}

// AzureConnectorOptions Options to connect to an Azure bucket.  Attributes ---------- storage_name : str     Name of the Azure storage. sas_token : str     SAS token to connect to Azure. container : str     Name of the container.
type AzureConnectorOptions struct {
	// Name of the Azure Blob Storage.
	StorageName string `json:"storage_name"`
	// SAS token for Azure Blob Storage connection.
	SasToken string `json:"sas_token"`
	// Name of the Azure Blob Storage container.
	Container string `json:"container"`
}

type _AzureConnectorOptions AzureConnectorOptions

// NewAzureConnectorOptions instantiates a new AzureConnectorOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureConnectorOptions(storageName string, sasToken string, container string) *AzureConnectorOptions {
	this := AzureConnectorOptions{}
	this.StorageName = storageName
	this.SasToken = sasToken
	this.Container = container
	return &this
}

// NewAzureConnectorOptionsWithDefaults instantiates a new AzureConnectorOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureConnectorOptionsWithDefaults() *AzureConnectorOptions {
	this := AzureConnectorOptions{}
	return &this
}

// GetStorageName returns the StorageName field value
func (o *AzureConnectorOptions) GetStorageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value
// and a boolean to check if the value has been set.
func (o *AzureConnectorOptions) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageName, true
}

// SetStorageName sets field value
func (o *AzureConnectorOptions) SetStorageName(v string) {
	o.StorageName = v
}

// GetSasToken returns the SasToken field value
func (o *AzureConnectorOptions) GetSasToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SasToken
}

// GetSasTokenOk returns a tuple with the SasToken field value
// and a boolean to check if the value has been set.
func (o *AzureConnectorOptions) GetSasTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SasToken, true
}

// SetSasToken sets field value
func (o *AzureConnectorOptions) SetSasToken(v string) {
	o.SasToken = v
}

// GetContainer returns the Container field value
func (o *AzureConnectorOptions) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureConnectorOptions) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureConnectorOptions) SetContainer(v string) {
	o.Container = v
}

func (o AzureConnectorOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureConnectorOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storage_name"] = o.StorageName
	toSerialize["sas_token"] = o.SasToken
	toSerialize["container"] = o.Container
	return toSerialize, nil
}

func (o *AzureConnectorOptions) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storage_name",
		"sas_token",
		"container",
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

	varAzureConnectorOptions := _AzureConnectorOptions{}

	err = json.Unmarshal(bytes, &varAzureConnectorOptions)

	if err != nil {
		return err
	}

	*o = AzureConnectorOptions(varAzureConnectorOptions)

	return err
}

type NullableAzureConnectorOptions struct {
	value *AzureConnectorOptions
	isSet bool
}

func (v NullableAzureConnectorOptions) Get() *AzureConnectorOptions {
	return v.value
}

func (v *NullableAzureConnectorOptions) Set(val *AzureConnectorOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureConnectorOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureConnectorOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureConnectorOptions(val *AzureConnectorOptions) *NullableAzureConnectorOptions {
	return &NullableAzureConnectorOptions{value: val, isSet: true}
}

func (v NullableAzureConnectorOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureConnectorOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


