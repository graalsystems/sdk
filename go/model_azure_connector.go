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

// checks if the AzureConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureConnector{}

// AzureConnector Connect to an Azure Blob Storage.  Attributes ---------- type : ConnectorType     Azure Blob Storage connector type. options : SQLConnectorOptions     Options to connect to an Azure Blob Storage.
type AzureConnector struct {
	// Connect to an Azure Blob Storage.
	Type *string `json:"type,omitempty"`
	Options AzureConnectorOptions `json:"options"`
}

type _AzureConnector AzureConnector

// NewAzureConnector instantiates a new AzureConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureConnector(options AzureConnectorOptions) *AzureConnector {
	this := AzureConnector{}
	var type_ string = "azure"
	this.Type = &type_
	this.Options = options
	return &this
}

// NewAzureConnectorWithDefaults instantiates a new AzureConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureConnectorWithDefaults() *AzureConnector {
	this := AzureConnector{}
	var type_ string = "azure"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureConnector) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureConnector) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureConnector) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureConnector) SetType(v string) {
	o.Type = &v
}

// GetOptions returns the Options field value
func (o *AzureConnector) GetOptions() AzureConnectorOptions {
	if o == nil {
		var ret AzureConnectorOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AzureConnector) GetOptionsOk() (*AzureConnectorOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *AzureConnector) SetOptions(v AzureConnectorOptions) {
	o.Options = v
}

func (o AzureConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

func (o *AzureConnector) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
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

	varAzureConnector := _AzureConnector{}

	err = json.Unmarshal(bytes, &varAzureConnector)

	if err != nil {
		return err
	}

	*o = AzureConnector(varAzureConnector)

	return err
}

type NullableAzureConnector struct {
	value *AzureConnector
	isSet bool
}

func (v NullableAzureConnector) Get() *AzureConnector {
	return v.value
}

func (v *NullableAzureConnector) Set(val *AzureConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureConnector(val *AzureConnector) *NullableAzureConnector {
	return &NullableAzureConnector{value: val, isSet: true}
}

func (v NullableAzureConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


