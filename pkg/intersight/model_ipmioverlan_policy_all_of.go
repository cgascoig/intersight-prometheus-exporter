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

// IpmioverlanPolicyAllOf Definition of the list of properties defined in 'ipmioverlan.Policy', excluding properties defined in parent classes.
type IpmioverlanPolicyAllOf struct {
	// State of the IPMI Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// Indicates whether the value of the 'encryptionKey' property has been set.
	IsEncryptionKeySet *bool `json:"IsEncryptionKeySet,omitempty"`
	// The highest privilege level that can be assigned to an IPMI session on a server. * `admin` - Privilege to perform all actions available through IPMI. * `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks. * `read-only` - Privilege to view information throught IPMI but restriction on making any changes.
	Privilege    *string                               `json:"Privilege,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IpmioverlanPolicyAllOf IpmioverlanPolicyAllOf

// NewIpmioverlanPolicyAllOf instantiates a new IpmioverlanPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmioverlanPolicyAllOf() *IpmioverlanPolicyAllOf {
	this := IpmioverlanPolicyAllOf{}
	var privilege string = "admin"
	this.Privilege = &privilege
	return &this
}

// NewIpmioverlanPolicyAllOfWithDefaults instantiates a new IpmioverlanPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmioverlanPolicyAllOfWithDefaults() *IpmioverlanPolicyAllOf {
	this := IpmioverlanPolicyAllOf{}
	var privilege string = "admin"
	this.Privilege = &privilege
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IpmioverlanPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IpmioverlanPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *IpmioverlanPolicyAllOf) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyAllOf) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *IpmioverlanPolicyAllOf) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetIsEncryptionKeySet returns the IsEncryptionKeySet field value if set, zero value otherwise.
func (o *IpmioverlanPolicyAllOf) GetIsEncryptionKeySet() bool {
	if o == nil || o.IsEncryptionKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsEncryptionKeySet
}

// GetIsEncryptionKeySetOk returns a tuple with the IsEncryptionKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyAllOf) GetIsEncryptionKeySetOk() (*bool, bool) {
	if o == nil || o.IsEncryptionKeySet == nil {
		return nil, false
	}
	return o.IsEncryptionKeySet, true
}

// HasIsEncryptionKeySet returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasIsEncryptionKeySet() bool {
	if o != nil && o.IsEncryptionKeySet != nil {
		return true
	}

	return false
}

// SetIsEncryptionKeySet gets a reference to the given bool and assigns it to the IsEncryptionKeySet field.
func (o *IpmioverlanPolicyAllOf) SetIsEncryptionKeySet(v bool) {
	o.IsEncryptionKeySet = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *IpmioverlanPolicyAllOf) GetPrivilege() string {
	if o == nil || o.Privilege == nil {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyAllOf) GetPrivilegeOk() (*string, bool) {
	if o == nil || o.Privilege == nil {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasPrivilege() bool {
	if o != nil && o.Privilege != nil {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *IpmioverlanPolicyAllOf) SetPrivilege(v string) {
	o.Privilege = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IpmioverlanPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmioverlanPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IpmioverlanPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmioverlanPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmioverlanPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IpmioverlanPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *IpmioverlanPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o IpmioverlanPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.IsEncryptionKeySet != nil {
		toSerialize["IsEncryptionKeySet"] = o.IsEncryptionKeySet
	}
	if o.Privilege != nil {
		toSerialize["Privilege"] = o.Privilege
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IpmioverlanPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIpmioverlanPolicyAllOf := _IpmioverlanPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varIpmioverlanPolicyAllOf); err == nil {
		*o = IpmioverlanPolicyAllOf(varIpmioverlanPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "IsEncryptionKeySet")
		delete(additionalProperties, "Privilege")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpmioverlanPolicyAllOf struct {
	value *IpmioverlanPolicyAllOf
	isSet bool
}

func (v NullableIpmioverlanPolicyAllOf) Get() *IpmioverlanPolicyAllOf {
	return v.value
}

func (v *NullableIpmioverlanPolicyAllOf) Set(val *IpmioverlanPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmioverlanPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmioverlanPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmioverlanPolicyAllOf(val *IpmioverlanPolicyAllOf) *NullableIpmioverlanPolicyAllOf {
	return &NullableIpmioverlanPolicyAllOf{value: val, isSet: true}
}

func (v NullableIpmioverlanPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmioverlanPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
