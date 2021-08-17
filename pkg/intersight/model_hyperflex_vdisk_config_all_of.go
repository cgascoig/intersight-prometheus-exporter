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

// HyperflexVdiskConfigAllOf Definition of the list of properties defined in 'hyperflex.VdiskConfig', excluding properties defined in parent classes.
type HyperflexVdiskConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access mode of the virtual disk. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. * `` - Unknown disk access mode.
	AccessMode *string `json:"AccessMode,omitempty"`
	// Disk capacity represented in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// File mode of the disk, example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	Mode *string `json:"Mode,omitempty"`
	// Name of the virtual disk.
	Name *string `json:"Name,omitempty"`
	// Source file path associated with virtual machine disk.
	SourceFilePath *string `json:"SourceFilePath,omitempty"`
	// Source disk name from where the clone is done.
	SourceVirtualDisk *string                     `json:"SourceVirtualDisk,omitempty"`
	Status            NullableHyperflexDiskStatus `json:"Status,omitempty"`
	// UUID of the virtual disk.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVdiskConfigAllOf HyperflexVdiskConfigAllOf

// NewHyperflexVdiskConfigAllOf instantiates a new HyperflexVdiskConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVdiskConfigAllOf(classId string, objectType string) *HyperflexVdiskConfigAllOf {
	this := HyperflexVdiskConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessMode string = "ReadWriteOnce"
	this.AccessMode = &accessMode
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewHyperflexVdiskConfigAllOfWithDefaults instantiates a new HyperflexVdiskConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVdiskConfigAllOfWithDefaults() *HyperflexVdiskConfigAllOf {
	this := HyperflexVdiskConfigAllOf{}
	var classId string = "hyperflex.VdiskConfig"
	this.ClassId = classId
	var objectType string = "hyperflex.VdiskConfig"
	this.ObjectType = objectType
	var accessMode string = "ReadWriteOnce"
	this.AccessMode = &accessMode
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVdiskConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVdiskConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVdiskConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVdiskConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *HyperflexVdiskConfigAllOf) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *HyperflexVdiskConfigAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *HyperflexVdiskConfigAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexVdiskConfigAllOf) SetName(v string) {
	o.Name = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *HyperflexVdiskConfigAllOf) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

// GetSourceVirtualDisk returns the SourceVirtualDisk field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetSourceVirtualDisk() string {
	if o == nil || o.SourceVirtualDisk == nil {
		var ret string
		return ret
	}
	return *o.SourceVirtualDisk
}

// GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetSourceVirtualDiskOk() (*string, bool) {
	if o == nil || o.SourceVirtualDisk == nil {
		return nil, false
	}
	return o.SourceVirtualDisk, true
}

// HasSourceVirtualDisk returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasSourceVirtualDisk() bool {
	if o != nil && o.SourceVirtualDisk != nil {
		return true
	}

	return false
}

// SetSourceVirtualDisk gets a reference to the given string and assigns it to the SourceVirtualDisk field.
func (o *HyperflexVdiskConfigAllOf) SetSourceVirtualDisk(v string) {
	o.SourceVirtualDisk = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVdiskConfigAllOf) GetStatus() HyperflexDiskStatus {
	if o == nil || o.Status.Get() == nil {
		var ret HyperflexDiskStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVdiskConfigAllOf) GetStatusOk() (*HyperflexDiskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableHyperflexDiskStatus and assigns it to the Status field.
func (o *HyperflexVdiskConfigAllOf) SetStatus(v HyperflexDiskStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *HyperflexVdiskConfigAllOf) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *HyperflexVdiskConfigAllOf) UnsetStatus() {
	o.Status.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexVdiskConfigAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVdiskConfigAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexVdiskConfigAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexVdiskConfigAllOf) SetUuid(v string) {
	o.Uuid = &v
}

func (o HyperflexVdiskConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessMode != nil {
		toSerialize["AccessMode"] = o.AccessMode
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}
	if o.SourceVirtualDisk != nil {
		toSerialize["SourceVirtualDisk"] = o.SourceVirtualDisk
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVdiskConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVdiskConfigAllOf := _HyperflexVdiskConfigAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVdiskConfigAllOf); err == nil {
		*o = HyperflexVdiskConfigAllOf(varHyperflexVdiskConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessMode")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SourceFilePath")
		delete(additionalProperties, "SourceVirtualDisk")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVdiskConfigAllOf struct {
	value *HyperflexVdiskConfigAllOf
	isSet bool
}

func (v NullableHyperflexVdiskConfigAllOf) Get() *HyperflexVdiskConfigAllOf {
	return v.value
}

func (v *NullableHyperflexVdiskConfigAllOf) Set(val *HyperflexVdiskConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVdiskConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVdiskConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVdiskConfigAllOf(val *HyperflexVdiskConfigAllOf) *NullableHyperflexVdiskConfigAllOf {
	return &NullableHyperflexVdiskConfigAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVdiskConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVdiskConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}