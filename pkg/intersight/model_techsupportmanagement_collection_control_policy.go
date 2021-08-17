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

// TechsupportmanagementCollectionControlPolicy Policy to control techsupport collection for a specific account.
type TechsupportmanagementCollectionControlPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Deployment type defines whether the policy is associated with a SaaS or Appliance account. * `None` - Service deployment type None. * `SaaS` - Service deployment type SaaS. * `Appliance` - Service deployment type Appliance.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// Enable or Disable techsupport collection for a specific account. * `Enable` - Enable techsupport collection. * `Disable` - Disable techsupport collection.
	TechSupportCollection *string                 `json:"TechSupportCollection,omitempty"`
	Account               *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _TechsupportmanagementCollectionControlPolicy TechsupportmanagementCollectionControlPolicy

// NewTechsupportmanagementCollectionControlPolicy instantiates a new TechsupportmanagementCollectionControlPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementCollectionControlPolicy(classId string, objectType string) *TechsupportmanagementCollectionControlPolicy {
	this := TechsupportmanagementCollectionControlPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var deploymentType string = "None"
	this.DeploymentType = &deploymentType
	var techSupportCollection string = "Enable"
	this.TechSupportCollection = &techSupportCollection
	return &this
}

// NewTechsupportmanagementCollectionControlPolicyWithDefaults instantiates a new TechsupportmanagementCollectionControlPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementCollectionControlPolicyWithDefaults() *TechsupportmanagementCollectionControlPolicy {
	this := TechsupportmanagementCollectionControlPolicy{}
	var classId string = "techsupportmanagement.CollectionControlPolicy"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.CollectionControlPolicy"
	this.ObjectType = objectType
	var deploymentType string = "None"
	this.DeploymentType = &deploymentType
	var techSupportCollection string = "Enable"
	this.TechSupportCollection = &techSupportCollection
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementCollectionControlPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementCollectionControlPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementCollectionControlPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementCollectionControlPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicy) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicy) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicy) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *TechsupportmanagementCollectionControlPolicy) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetTechSupportCollection returns the TechSupportCollection field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicy) GetTechSupportCollection() string {
	if o == nil || o.TechSupportCollection == nil {
		var ret string
		return ret
	}
	return *o.TechSupportCollection
}

// GetTechSupportCollectionOk returns a tuple with the TechSupportCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicy) GetTechSupportCollectionOk() (*string, bool) {
	if o == nil || o.TechSupportCollection == nil {
		return nil, false
	}
	return o.TechSupportCollection, true
}

// HasTechSupportCollection returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicy) HasTechSupportCollection() bool {
	if o != nil && o.TechSupportCollection != nil {
		return true
	}

	return false
}

// SetTechSupportCollection gets a reference to the given string and assigns it to the TechSupportCollection field.
func (o *TechsupportmanagementCollectionControlPolicy) SetTechSupportCollection(v string) {
	o.TechSupportCollection = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TechsupportmanagementCollectionControlPolicy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TechsupportmanagementCollectionControlPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.TechSupportCollection != nil {
		toSerialize["TechSupportCollection"] = o.TechSupportCollection
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementCollectionControlPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type TechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Deployment type defines whether the policy is associated with a SaaS or Appliance account. * `None` - Service deployment type None. * `SaaS` - Service deployment type SaaS. * `Appliance` - Service deployment type Appliance.
		DeploymentType *string `json:"DeploymentType,omitempty"`
		// Enable or Disable techsupport collection for a specific account. * `Enable` - Enable techsupport collection. * `Disable` - Disable techsupport collection.
		TechSupportCollection *string                 `json:"TechSupportCollection,omitempty"`
		Account               *IamAccountRelationship `json:"Account,omitempty"`
	}

	varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct := TechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct)
	if err == nil {
		varTechsupportmanagementCollectionControlPolicy := _TechsupportmanagementCollectionControlPolicy{}
		varTechsupportmanagementCollectionControlPolicy.ClassId = varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct.ClassId
		varTechsupportmanagementCollectionControlPolicy.ObjectType = varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct.ObjectType
		varTechsupportmanagementCollectionControlPolicy.DeploymentType = varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct.DeploymentType
		varTechsupportmanagementCollectionControlPolicy.TechSupportCollection = varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct.TechSupportCollection
		varTechsupportmanagementCollectionControlPolicy.Account = varTechsupportmanagementCollectionControlPolicyWithoutEmbeddedStruct.Account
		*o = TechsupportmanagementCollectionControlPolicy(varTechsupportmanagementCollectionControlPolicy)
	} else {
		return err
	}

	varTechsupportmanagementCollectionControlPolicy := _TechsupportmanagementCollectionControlPolicy{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementCollectionControlPolicy)
	if err == nil {
		o.MoBaseMo = varTechsupportmanagementCollectionControlPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "TechSupportCollection")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableTechsupportmanagementCollectionControlPolicy struct {
	value *TechsupportmanagementCollectionControlPolicy
	isSet bool
}

func (v NullableTechsupportmanagementCollectionControlPolicy) Get() *TechsupportmanagementCollectionControlPolicy {
	return v.value
}

func (v *NullableTechsupportmanagementCollectionControlPolicy) Set(val *TechsupportmanagementCollectionControlPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementCollectionControlPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementCollectionControlPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementCollectionControlPolicy(val *TechsupportmanagementCollectionControlPolicy) *NullableTechsupportmanagementCollectionControlPolicy {
	return &NullableTechsupportmanagementCollectionControlPolicy{value: val, isSet: true}
}

func (v NullableTechsupportmanagementCollectionControlPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementCollectionControlPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
