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

// KubernetesClusterAllOf Definition of the list of properties defined in 'kubernetes.Cluster', excluding properties defined in parent classes.
type KubernetesClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the endpoint connection of this Kubernetes cluster. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Kubeconfig for the cluster to collect inventory for.
	KubeConfig *string `json:"KubeConfig,omitempty"`
	// Name of the Kubernetes cluster.
	Name                *string                                    `json:"Name,omitempty"`
	ClusterAddonProfile *KubernetesClusterAddonProfileRelationship `json:"ClusterAddonProfile,omitempty"`
	Organization        *OrganizationOrganizationRelationship      `json:"Organization,omitempty"`
	// An array of relationships to assetDeviceRegistration resources.
	RegisteredDevices    []AssetDeviceRegistrationRelationship `json:"RegisteredDevices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterAllOf KubernetesClusterAllOf

// NewKubernetesClusterAllOf instantiates a new KubernetesClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterAllOf(classId string, objectType string) *KubernetesClusterAllOf {
	this := KubernetesClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewKubernetesClusterAllOfWithDefaults instantiates a new KubernetesClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterAllOfWithDefaults() *KubernetesClusterAllOf {
	this := KubernetesClusterAllOf{}
	var classId string = "kubernetes.Cluster"
	this.ClassId = classId
	var objectType string = "kubernetes.Cluster"
	this.ObjectType = objectType
	var connectionStatus string = ""
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *KubernetesClusterAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *KubernetesClusterAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetKubeConfig returns the KubeConfig field value if set, zero value otherwise.
func (o *KubernetesClusterAllOf) GetKubeConfig() string {
	if o == nil || o.KubeConfig == nil {
		var ret string
		return ret
	}
	return *o.KubeConfig
}

// GetKubeConfigOk returns a tuple with the KubeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetKubeConfigOk() (*string, bool) {
	if o == nil || o.KubeConfig == nil {
		return nil, false
	}
	return o.KubeConfig, true
}

// HasKubeConfig returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasKubeConfig() bool {
	if o != nil && o.KubeConfig != nil {
		return true
	}

	return false
}

// SetKubeConfig gets a reference to the given string and assigns it to the KubeConfig field.
func (o *KubernetesClusterAllOf) SetKubeConfig(v string) {
	o.KubeConfig = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesClusterAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesClusterAllOf) SetName(v string) {
	o.Name = &v
}

// GetClusterAddonProfile returns the ClusterAddonProfile field value if set, zero value otherwise.
func (o *KubernetesClusterAllOf) GetClusterAddonProfile() KubernetesClusterAddonProfileRelationship {
	if o == nil || o.ClusterAddonProfile == nil {
		var ret KubernetesClusterAddonProfileRelationship
		return ret
	}
	return *o.ClusterAddonProfile
}

// GetClusterAddonProfileOk returns a tuple with the ClusterAddonProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetClusterAddonProfileOk() (*KubernetesClusterAddonProfileRelationship, bool) {
	if o == nil || o.ClusterAddonProfile == nil {
		return nil, false
	}
	return o.ClusterAddonProfile, true
}

// HasClusterAddonProfile returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasClusterAddonProfile() bool {
	if o != nil && o.ClusterAddonProfile != nil {
		return true
	}

	return false
}

// SetClusterAddonProfile gets a reference to the given KubernetesClusterAddonProfileRelationship and assigns it to the ClusterAddonProfile field.
func (o *KubernetesClusterAllOf) SetClusterAddonProfile(v KubernetesClusterAddonProfileRelationship) {
	o.ClusterAddonProfile = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesClusterAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesClusterAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRegisteredDevices returns the RegisteredDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterAllOf) GetRegisteredDevices() []AssetDeviceRegistrationRelationship {
	if o == nil {
		var ret []AssetDeviceRegistrationRelationship
		return ret
	}
	return o.RegisteredDevices
}

// GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterAllOf) GetRegisteredDevicesOk() (*[]AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevices == nil {
		return nil, false
	}
	return &o.RegisteredDevices, true
}

// HasRegisteredDevices returns a boolean if a field has been set.
func (o *KubernetesClusterAllOf) HasRegisteredDevices() bool {
	if o != nil && o.RegisteredDevices != nil {
		return true
	}

	return false
}

// SetRegisteredDevices gets a reference to the given []AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevices field.
func (o *KubernetesClusterAllOf) SetRegisteredDevices(v []AssetDeviceRegistrationRelationship) {
	o.RegisteredDevices = v
}

func (o KubernetesClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.KubeConfig != nil {
		toSerialize["KubeConfig"] = o.KubeConfig
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ClusterAddonProfile != nil {
		toSerialize["ClusterAddonProfile"] = o.ClusterAddonProfile
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RegisteredDevices != nil {
		toSerialize["RegisteredDevices"] = o.RegisteredDevices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterAllOf := _KubernetesClusterAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterAllOf); err == nil {
		*o = KubernetesClusterAllOf(varKubernetesClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "KubeConfig")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ClusterAddonProfile")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RegisteredDevices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterAllOf struct {
	value *KubernetesClusterAllOf
	isSet bool
}

func (v NullableKubernetesClusterAllOf) Get() *KubernetesClusterAllOf {
	return v.value
}

func (v *NullableKubernetesClusterAllOf) Set(val *KubernetesClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterAllOf(val *KubernetesClusterAllOf) *NullableKubernetesClusterAllOf {
	return &NullableKubernetesClusterAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
