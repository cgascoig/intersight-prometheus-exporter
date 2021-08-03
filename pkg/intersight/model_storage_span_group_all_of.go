/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// StorageSpanGroupAllOf Definition of the list of properties defined in 'storage.SpanGroup', excluding properties defined in parent classes.
type StorageSpanGroupAllOf struct {
	Disks                *[]StorageLocalDisk `json:"Disks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSpanGroupAllOf StorageSpanGroupAllOf

// NewStorageSpanGroupAllOf instantiates a new StorageSpanGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpanGroupAllOf() *StorageSpanGroupAllOf {
	this := StorageSpanGroupAllOf{}
	return &this
}

// NewStorageSpanGroupAllOfWithDefaults instantiates a new StorageSpanGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanGroupAllOfWithDefaults() *StorageSpanGroupAllOf {
	this := StorageSpanGroupAllOf{}
	return &this
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *StorageSpanGroupAllOf) GetDisks() []StorageLocalDisk {
	if o == nil || o.Disks == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanGroupAllOf) GetDisksOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *StorageSpanGroupAllOf) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StorageLocalDisk and assigns it to the Disks field.
func (o *StorageSpanGroupAllOf) SetDisks(v []StorageLocalDisk) {
	o.Disks = &v
}

func (o StorageSpanGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageSpanGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageSpanGroupAllOf := _StorageSpanGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageSpanGroupAllOf); err == nil {
		*o = StorageSpanGroupAllOf(varStorageSpanGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Disks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageSpanGroupAllOf struct {
	value *StorageSpanGroupAllOf
	isSet bool
}

func (v NullableStorageSpanGroupAllOf) Get() *StorageSpanGroupAllOf {
	return v.value
}

func (v *NullableStorageSpanGroupAllOf) Set(val *StorageSpanGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpanGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpanGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpanGroupAllOf(val *StorageSpanGroupAllOf) *NullableStorageSpanGroupAllOf {
	return &NullableStorageSpanGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageSpanGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpanGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
