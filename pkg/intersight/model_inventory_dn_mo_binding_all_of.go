/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// InventoryDnMoBindingAllOf Definition of the list of properties defined in 'inventory.DnMoBinding', excluding properties defined in parent classes.
type InventoryDnMoBindingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Distinguished Name for this object, used to uniquely identify this object.
	Dn *string `json:"Dn,omitempty"`
	// The MO ID of the target MO for this particular Distinguished Name (dn).
	TargetMoId *string `json:"TargetMoId,omitempty"`
	// The type of the target MO for this particular Distinguished Name (dn).
	TargetMoType         *string                              `json:"TargetMoType,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryDnMoBindingAllOf InventoryDnMoBindingAllOf

// NewInventoryDnMoBindingAllOf instantiates a new InventoryDnMoBindingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDnMoBindingAllOf(classId string, objectType string) *InventoryDnMoBindingAllOf {
	this := InventoryDnMoBindingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryDnMoBindingAllOfWithDefaults instantiates a new InventoryDnMoBindingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDnMoBindingAllOfWithDefaults() *InventoryDnMoBindingAllOf {
	this := InventoryDnMoBindingAllOf{}
	var classId string = "inventory.DnMoBinding"
	this.ClassId = classId
	var objectType string = "inventory.DnMoBinding"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryDnMoBindingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryDnMoBindingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryDnMoBindingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryDnMoBindingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryDnMoBindingAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryDnMoBindingAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryDnMoBindingAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetTargetMoId returns the TargetMoId field value if set, zero value otherwise.
func (o *InventoryDnMoBindingAllOf) GetTargetMoId() string {
	if o == nil || o.TargetMoId == nil {
		var ret string
		return ret
	}
	return *o.TargetMoId
}

// GetTargetMoIdOk returns a tuple with the TargetMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetTargetMoIdOk() (*string, bool) {
	if o == nil || o.TargetMoId == nil {
		return nil, false
	}
	return o.TargetMoId, true
}

// HasTargetMoId returns a boolean if a field has been set.
func (o *InventoryDnMoBindingAllOf) HasTargetMoId() bool {
	if o != nil && o.TargetMoId != nil {
		return true
	}

	return false
}

// SetTargetMoId gets a reference to the given string and assigns it to the TargetMoId field.
func (o *InventoryDnMoBindingAllOf) SetTargetMoId(v string) {
	o.TargetMoId = &v
}

// GetTargetMoType returns the TargetMoType field value if set, zero value otherwise.
func (o *InventoryDnMoBindingAllOf) GetTargetMoType() string {
	if o == nil || o.TargetMoType == nil {
		var ret string
		return ret
	}
	return *o.TargetMoType
}

// GetTargetMoTypeOk returns a tuple with the TargetMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetTargetMoTypeOk() (*string, bool) {
	if o == nil || o.TargetMoType == nil {
		return nil, false
	}
	return o.TargetMoType, true
}

// HasTargetMoType returns a boolean if a field has been set.
func (o *InventoryDnMoBindingAllOf) HasTargetMoType() bool {
	if o != nil && o.TargetMoType != nil {
		return true
	}

	return false
}

// SetTargetMoType gets a reference to the given string and assigns it to the TargetMoType field.
func (o *InventoryDnMoBindingAllOf) SetTargetMoType(v string) {
	o.TargetMoType = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryDnMoBindingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDnMoBindingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryDnMoBindingAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryDnMoBindingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryDnMoBindingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.TargetMoId != nil {
		toSerialize["TargetMoId"] = o.TargetMoId
	}
	if o.TargetMoType != nil {
		toSerialize["TargetMoType"] = o.TargetMoType
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryDnMoBindingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryDnMoBindingAllOf := _InventoryDnMoBindingAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryDnMoBindingAllOf); err == nil {
		*o = InventoryDnMoBindingAllOf(varInventoryDnMoBindingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "TargetMoId")
		delete(additionalProperties, "TargetMoType")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryDnMoBindingAllOf struct {
	value *InventoryDnMoBindingAllOf
	isSet bool
}

func (v NullableInventoryDnMoBindingAllOf) Get() *InventoryDnMoBindingAllOf {
	return v.value
}

func (v *NullableInventoryDnMoBindingAllOf) Set(val *InventoryDnMoBindingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDnMoBindingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDnMoBindingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDnMoBindingAllOf(val *InventoryDnMoBindingAllOf) *NullableInventoryDnMoBindingAllOf {
	return &NullableInventoryDnMoBindingAllOf{value: val, isSet: true}
}

func (v NullableInventoryDnMoBindingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDnMoBindingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
