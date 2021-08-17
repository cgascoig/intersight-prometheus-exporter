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

// KubernetesTrustedRegistriesPolicy A policy specifying self signed docker registries and CA certificates to trust.
type KubernetesTrustedRegistriesPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string   `json:"ObjectType"`
	RootCaRegistries   []string `json:"RootCaRegistries,omitempty"`
	UnsignedRegistries []string `json:"UnsignedRegistries,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesTrustedRegistriesPolicy KubernetesTrustedRegistriesPolicy

// NewKubernetesTrustedRegistriesPolicy instantiates a new KubernetesTrustedRegistriesPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesTrustedRegistriesPolicy(classId string, objectType string) *KubernetesTrustedRegistriesPolicy {
	this := KubernetesTrustedRegistriesPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesTrustedRegistriesPolicyWithDefaults instantiates a new KubernetesTrustedRegistriesPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesTrustedRegistriesPolicyWithDefaults() *KubernetesTrustedRegistriesPolicy {
	this := KubernetesTrustedRegistriesPolicy{}
	var classId string = "kubernetes.TrustedRegistriesPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.TrustedRegistriesPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesTrustedRegistriesPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesTrustedRegistriesPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesTrustedRegistriesPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesTrustedRegistriesPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRootCaRegistries returns the RootCaRegistries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicy) GetRootCaRegistries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RootCaRegistries
}

// GetRootCaRegistriesOk returns a tuple with the RootCaRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicy) GetRootCaRegistriesOk() (*[]string, bool) {
	if o == nil || o.RootCaRegistries == nil {
		return nil, false
	}
	return &o.RootCaRegistries, true
}

// HasRootCaRegistries returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicy) HasRootCaRegistries() bool {
	if o != nil && o.RootCaRegistries != nil {
		return true
	}

	return false
}

// SetRootCaRegistries gets a reference to the given []string and assigns it to the RootCaRegistries field.
func (o *KubernetesTrustedRegistriesPolicy) SetRootCaRegistries(v []string) {
	o.RootCaRegistries = v
}

// GetUnsignedRegistries returns the UnsignedRegistries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicy) GetUnsignedRegistries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UnsignedRegistries
}

// GetUnsignedRegistriesOk returns a tuple with the UnsignedRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicy) GetUnsignedRegistriesOk() (*[]string, bool) {
	if o == nil || o.UnsignedRegistries == nil {
		return nil, false
	}
	return &o.UnsignedRegistries, true
}

// HasUnsignedRegistries returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicy) HasUnsignedRegistries() bool {
	if o != nil && o.UnsignedRegistries != nil {
		return true
	}

	return false
}

// SetUnsignedRegistries gets a reference to the given []string and assigns it to the UnsignedRegistries field.
func (o *KubernetesTrustedRegistriesPolicy) SetUnsignedRegistries(v []string) {
	o.UnsignedRegistries = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicy) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesTrustedRegistriesPolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesTrustedRegistriesPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesTrustedRegistriesPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesTrustedRegistriesPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RootCaRegistries != nil {
		toSerialize["RootCaRegistries"] = o.RootCaRegistries
	}
	if o.UnsignedRegistries != nil {
		toSerialize["UnsignedRegistries"] = o.UnsignedRegistries
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesTrustedRegistriesPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string   `json:"ObjectType"`
		RootCaRegistries   []string `json:"RootCaRegistries,omitempty"`
		UnsignedRegistries []string `json:"UnsignedRegistries,omitempty"`
		// An array of relationships to kubernetesClusterProfile resources.
		ClusterProfiles []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	}

	varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct := KubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesTrustedRegistriesPolicy := _KubernetesTrustedRegistriesPolicy{}
		varKubernetesTrustedRegistriesPolicy.ClassId = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesTrustedRegistriesPolicy.ObjectType = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesTrustedRegistriesPolicy.RootCaRegistries = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.RootCaRegistries
		varKubernetesTrustedRegistriesPolicy.UnsignedRegistries = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.UnsignedRegistries
		varKubernetesTrustedRegistriesPolicy.ClusterProfiles = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.ClusterProfiles
		varKubernetesTrustedRegistriesPolicy.Organization = varKubernetesTrustedRegistriesPolicyWithoutEmbeddedStruct.Organization
		*o = KubernetesTrustedRegistriesPolicy(varKubernetesTrustedRegistriesPolicy)
	} else {
		return err
	}

	varKubernetesTrustedRegistriesPolicy := _KubernetesTrustedRegistriesPolicy{}

	err = json.Unmarshal(bytes, &varKubernetesTrustedRegistriesPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesTrustedRegistriesPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RootCaRegistries")
		delete(additionalProperties, "UnsignedRegistries")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableKubernetesTrustedRegistriesPolicy struct {
	value *KubernetesTrustedRegistriesPolicy
	isSet bool
}

func (v NullableKubernetesTrustedRegistriesPolicy) Get() *KubernetesTrustedRegistriesPolicy {
	return v.value
}

func (v *NullableKubernetesTrustedRegistriesPolicy) Set(val *KubernetesTrustedRegistriesPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTrustedRegistriesPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTrustedRegistriesPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTrustedRegistriesPolicy(val *KubernetesTrustedRegistriesPolicy) *NullableKubernetesTrustedRegistriesPolicy {
	return &NullableKubernetesTrustedRegistriesPolicy{value: val, isSet: true}
}

func (v NullableKubernetesTrustedRegistriesPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTrustedRegistriesPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}