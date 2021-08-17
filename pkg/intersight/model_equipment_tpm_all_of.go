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

// EquipmentTpmAllOf Definition of the list of properties defined in 'equipment.Tpm', excluding properties defined in parent classes.
type EquipmentTpmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Identifies the activation status of the TPM.
	ActivationStatus *string `json:"ActivationStatus,omitempty"`
	// Identifies the admin configured state of the TPM.
	AdminState *string `json:"AdminState,omitempty"`
	// Firmware Version of the Trusted Platform Module.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// Identifies the ownership information of the TPM.
	Ownership *string `json:"Ownership,omitempty"`
	// Enter  the ID of the trusted platform module.
	TpmId *int64 `json:"TpmId,omitempty"`
	// Identifies the version of the Trusted Platform Module.
	Version              *string                              `json:"Version,omitempty"`
	ComputeBoard         *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentTpmAllOf EquipmentTpmAllOf

// NewEquipmentTpmAllOf instantiates a new EquipmentTpmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentTpmAllOf(classId string, objectType string) *EquipmentTpmAllOf {
	this := EquipmentTpmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentTpmAllOfWithDefaults instantiates a new EquipmentTpmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentTpmAllOfWithDefaults() *EquipmentTpmAllOf {
	this := EquipmentTpmAllOf{}
	var classId string = "equipment.Tpm"
	this.ClassId = classId
	var objectType string = "equipment.Tpm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentTpmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentTpmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentTpmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentTpmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActivationStatus returns the ActivationStatus field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetActivationStatus() string {
	if o == nil || o.ActivationStatus == nil {
		var ret string
		return ret
	}
	return *o.ActivationStatus
}

// GetActivationStatusOk returns a tuple with the ActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetActivationStatusOk() (*string, bool) {
	if o == nil || o.ActivationStatus == nil {
		return nil, false
	}
	return o.ActivationStatus, true
}

// HasActivationStatus returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasActivationStatus() bool {
	if o != nil && o.ActivationStatus != nil {
		return true
	}

	return false
}

// SetActivationStatus gets a reference to the given string and assigns it to the ActivationStatus field.
func (o *EquipmentTpmAllOf) SetActivationStatus(v string) {
	o.ActivationStatus = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *EquipmentTpmAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *EquipmentTpmAllOf) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetOwnership() string {
	if o == nil || o.Ownership == nil {
		var ret string
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetOwnershipOk() (*string, bool) {
	if o == nil || o.Ownership == nil {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasOwnership() bool {
	if o != nil && o.Ownership != nil {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given string and assigns it to the Ownership field.
func (o *EquipmentTpmAllOf) SetOwnership(v string) {
	o.Ownership = &v
}

// GetTpmId returns the TpmId field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetTpmId() int64 {
	if o == nil || o.TpmId == nil {
		var ret int64
		return ret
	}
	return *o.TpmId
}

// GetTpmIdOk returns a tuple with the TpmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetTpmIdOk() (*int64, bool) {
	if o == nil || o.TpmId == nil {
		return nil, false
	}
	return o.TpmId, true
}

// HasTpmId returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasTpmId() bool {
	if o != nil && o.TpmId != nil {
		return true
	}

	return false
}

// SetTpmId gets a reference to the given int64 and assigns it to the TpmId field.
func (o *EquipmentTpmAllOf) SetTpmId(v int64) {
	o.TpmId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentTpmAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *EquipmentTpmAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentTpmAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentTpmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTpmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentTpmAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentTpmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentTpmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActivationStatus != nil {
		toSerialize["ActivationStatus"] = o.ActivationStatus
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.Ownership != nil {
		toSerialize["Ownership"] = o.Ownership
	}
	if o.TpmId != nil {
		toSerialize["TpmId"] = o.TpmId
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentTpmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentTpmAllOf := _EquipmentTpmAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentTpmAllOf); err == nil {
		*o = EquipmentTpmAllOf(varEquipmentTpmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActivationStatus")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "Ownership")
		delete(additionalProperties, "TpmId")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentTpmAllOf struct {
	value *EquipmentTpmAllOf
	isSet bool
}

func (v NullableEquipmentTpmAllOf) Get() *EquipmentTpmAllOf {
	return v.value
}

func (v *NullableEquipmentTpmAllOf) Set(val *EquipmentTpmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentTpmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentTpmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentTpmAllOf(val *EquipmentTpmAllOf) *NullableEquipmentTpmAllOf {
	return &NullableEquipmentTpmAllOf{value: val, isSet: true}
}

func (v NullableEquipmentTpmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentTpmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
