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
	"reflect"
	"strings"
)

// VirtualizationVmwareFolder A folder in a VMware vCenter. Folder can be created directly under the vCenter, under a datacenter, or inside another folder.
type VirtualizationVmwareFolder struct {
	VirtualizationBaseFolder
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If a folder is internal, it will be set to true.
	Internal *bool `json:"Internal,omitempty"`
	// Inventory path to the folder. Example - /DC/myFolder.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Determines the type of folder. e.g. vCenter folder, VM and Templete Folder, StorageFolder, NetworkFolder, Host and Cluster Folder. * `Unknown` - The type of the folder is unknown. It may not represent that the folder does not exist but indicates that something might be wrong. * `VMTemplateFolder` - The folder contains VMs and VM templates. * `StorageFolder` - The folder contains storage devices. * `HostClusterFolder` - The folder contains hosts and clusters. * `NetworkFolder` - The folder contains network items. * `VcenterFolder` - The folder created under a vCenter or vCenter folder.
	TypeofFolder         *string                                     `json:"TypeofFolder,omitempty"`
	Datacenter           *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	HypervisorManager    *VirtualizationVmwareVcenterRelationship    `json:"HypervisorManager,omitempty"`
	VmwareFolder         *VirtualizationVmwareFolderRelationship     `json:"VmwareFolder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareFolder VirtualizationVmwareFolder

// NewVirtualizationVmwareFolder instantiates a new VirtualizationVmwareFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareFolder(classId string, objectType string) *VirtualizationVmwareFolder {
	this := VirtualizationVmwareFolder{}
	this.ClassId = classId
	this.ObjectType = objectType
	var typeofFolder string = "Unknown"
	this.TypeofFolder = &typeofFolder
	return &this
}

// NewVirtualizationVmwareFolderWithDefaults instantiates a new VirtualizationVmwareFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareFolderWithDefaults() *VirtualizationVmwareFolder {
	this := VirtualizationVmwareFolder{}
	var classId string = "virtualization.VmwareFolder"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareFolder"
	this.ObjectType = objectType
	var typeofFolder string = "Unknown"
	this.TypeofFolder = &typeofFolder
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareFolder) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareFolder) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareFolder) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareFolder) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *VirtualizationVmwareFolder) SetInternal(v bool) {
	o.Internal = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareFolder) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetTypeofFolder returns the TypeofFolder field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetTypeofFolder() string {
	if o == nil || o.TypeofFolder == nil {
		var ret string
		return ret
	}
	return *o.TypeofFolder
}

// GetTypeofFolderOk returns a tuple with the TypeofFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetTypeofFolderOk() (*string, bool) {
	if o == nil || o.TypeofFolder == nil {
		return nil, false
	}
	return o.TypeofFolder, true
}

// HasTypeofFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasTypeofFolder() bool {
	if o != nil && o.TypeofFolder != nil {
		return true
	}

	return false
}

// SetTypeofFolder gets a reference to the given string and assigns it to the TypeofFolder field.
func (o *VirtualizationVmwareFolder) SetTypeofFolder(v string) {
	o.TypeofFolder = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareFolder) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetHypervisorManager() VirtualizationVmwareVcenterRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given VirtualizationVmwareVcenterRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationVmwareFolder) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship) {
	o.HypervisorManager = &v
}

// GetVmwareFolder returns the VmwareFolder field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolder) GetVmwareFolder() VirtualizationVmwareFolderRelationship {
	if o == nil || o.VmwareFolder == nil {
		var ret VirtualizationVmwareFolderRelationship
		return ret
	}
	return *o.VmwareFolder
}

// GetVmwareFolderOk returns a tuple with the VmwareFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolder) GetVmwareFolderOk() (*VirtualizationVmwareFolderRelationship, bool) {
	if o == nil || o.VmwareFolder == nil {
		return nil, false
	}
	return o.VmwareFolder, true
}

// HasVmwareFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolder) HasVmwareFolder() bool {
	if o != nil && o.VmwareFolder != nil {
		return true
	}

	return false
}

// SetVmwareFolder gets a reference to the given VirtualizationVmwareFolderRelationship and assigns it to the VmwareFolder field.
func (o *VirtualizationVmwareFolder) SetVmwareFolder(v VirtualizationVmwareFolderRelationship) {
	o.VmwareFolder = &v
}

func (o VirtualizationVmwareFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseFolder, errVirtualizationBaseFolder := json.Marshal(o.VirtualizationBaseFolder)
	if errVirtualizationBaseFolder != nil {
		return []byte{}, errVirtualizationBaseFolder
	}
	errVirtualizationBaseFolder = json.Unmarshal([]byte(serializedVirtualizationBaseFolder), &toSerialize)
	if errVirtualizationBaseFolder != nil {
		return []byte{}, errVirtualizationBaseFolder
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.TypeofFolder != nil {
		toSerialize["TypeofFolder"] = o.TypeofFolder
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}
	if o.VmwareFolder != nil {
		toSerialize["VmwareFolder"] = o.VmwareFolder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareFolder) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareFolderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If a folder is internal, it will be set to true.
		Internal *bool `json:"Internal,omitempty"`
		// Inventory path to the folder. Example - /DC/myFolder.
		InventoryPath *string `json:"InventoryPath,omitempty"`
		// Determines the type of folder. e.g. vCenter folder, VM and Templete Folder, StorageFolder, NetworkFolder, Host and Cluster Folder. * `Unknown` - The type of the folder is unknown. It may not represent that the folder does not exist but indicates that something might be wrong. * `VMTemplateFolder` - The folder contains VMs and VM templates. * `StorageFolder` - The folder contains storage devices. * `HostClusterFolder` - The folder contains hosts and clusters. * `NetworkFolder` - The folder contains network items. * `VcenterFolder` - The folder created under a vCenter or vCenter folder.
		TypeofFolder      *string                                     `json:"TypeofFolder,omitempty"`
		Datacenter        *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
		HypervisorManager *VirtualizationVmwareVcenterRelationship    `json:"HypervisorManager,omitempty"`
		VmwareFolder      *VirtualizationVmwareFolderRelationship     `json:"VmwareFolder,omitempty"`
	}

	varVirtualizationVmwareFolderWithoutEmbeddedStruct := VirtualizationVmwareFolderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareFolderWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareFolder := _VirtualizationVmwareFolder{}
		varVirtualizationVmwareFolder.ClassId = varVirtualizationVmwareFolderWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareFolder.ObjectType = varVirtualizationVmwareFolderWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareFolder.Internal = varVirtualizationVmwareFolderWithoutEmbeddedStruct.Internal
		varVirtualizationVmwareFolder.InventoryPath = varVirtualizationVmwareFolderWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareFolder.TypeofFolder = varVirtualizationVmwareFolderWithoutEmbeddedStruct.TypeofFolder
		varVirtualizationVmwareFolder.Datacenter = varVirtualizationVmwareFolderWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareFolder.HypervisorManager = varVirtualizationVmwareFolderWithoutEmbeddedStruct.HypervisorManager
		varVirtualizationVmwareFolder.VmwareFolder = varVirtualizationVmwareFolderWithoutEmbeddedStruct.VmwareFolder
		*o = VirtualizationVmwareFolder(varVirtualizationVmwareFolder)
	} else {
		return err
	}

	varVirtualizationVmwareFolder := _VirtualizationVmwareFolder{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareFolder)
	if err == nil {
		o.VirtualizationBaseFolder = varVirtualizationVmwareFolder.VirtualizationBaseFolder
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "TypeofFolder")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "HypervisorManager")
		delete(additionalProperties, "VmwareFolder")

		// remove fields from embedded structs
		reflectVirtualizationBaseFolder := reflect.ValueOf(o.VirtualizationBaseFolder)
		for i := 0; i < reflectVirtualizationBaseFolder.Type().NumField(); i++ {
			t := reflectVirtualizationBaseFolder.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareFolder struct {
	value *VirtualizationVmwareFolder
	isSet bool
}

func (v NullableVirtualizationVmwareFolder) Get() *VirtualizationVmwareFolder {
	return v.value
}

func (v *NullableVirtualizationVmwareFolder) Set(val *VirtualizationVmwareFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareFolder(val *VirtualizationVmwareFolder) *NullableVirtualizationVmwareFolder {
	return &NullableVirtualizationVmwareFolder{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}