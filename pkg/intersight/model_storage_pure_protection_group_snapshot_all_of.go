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

// StoragePureProtectionGroupSnapshotAllOf Definition of the list of properties defined in 'storage.PureProtectionGroupSnapshot', excluding properties defined in parent classes.
type StoragePureProtectionGroupSnapshotAllOf struct {
	Array                *StoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureProtectionGroupSnapshotAllOf StoragePureProtectionGroupSnapshotAllOf

// NewStoragePureProtectionGroupSnapshotAllOf instantiates a new StoragePureProtectionGroupSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureProtectionGroupSnapshotAllOf() *StoragePureProtectionGroupSnapshotAllOf {
	this := StoragePureProtectionGroupSnapshotAllOf{}
	return &this
}

// NewStoragePureProtectionGroupSnapshotAllOfWithDefaults instantiates a new StoragePureProtectionGroupSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureProtectionGroupSnapshotAllOfWithDefaults() *StoragePureProtectionGroupSnapshotAllOf {
	this := StoragePureProtectionGroupSnapshotAllOf{}
	return &this
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureProtectionGroupSnapshotAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureProtectionGroupSnapshotAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureProtectionGroupSnapshotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureProtectionGroupSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureProtectionGroupSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureProtectionGroupSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePureProtectionGroupSnapshotAllOf := _StoragePureProtectionGroupSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePureProtectionGroupSnapshotAllOf); err == nil {
		*o = StoragePureProtectionGroupSnapshotAllOf(varStoragePureProtectionGroupSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePureProtectionGroupSnapshotAllOf struct {
	value *StoragePureProtectionGroupSnapshotAllOf
	isSet bool
}

func (v NullableStoragePureProtectionGroupSnapshotAllOf) Get() *StoragePureProtectionGroupSnapshotAllOf {
	return v.value
}

func (v *NullableStoragePureProtectionGroupSnapshotAllOf) Set(val *StoragePureProtectionGroupSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureProtectionGroupSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureProtectionGroupSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureProtectionGroupSnapshotAllOf(val *StoragePureProtectionGroupSnapshotAllOf) *NullableStoragePureProtectionGroupSnapshotAllOf {
	return &NullableStoragePureProtectionGroupSnapshotAllOf{value: val, isSet: true}
}

func (v NullableStoragePureProtectionGroupSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureProtectionGroupSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
