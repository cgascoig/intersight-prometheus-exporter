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
	"reflect"
	"strings"
)

// AdapterConfigPolicy An Adapter Configuration Policy configures the Ethernet and Fibre-Channel settings for the VIC adapter.
type AdapterConfigPolicy struct {
	PolicyAbstractPolicy
	Settings     *[]AdapterAdapterConfig               `json:"Settings,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterConfigPolicy AdapterConfigPolicy

// NewAdapterConfigPolicy instantiates a new AdapterConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterConfigPolicy() *AdapterConfigPolicy {
	this := AdapterConfigPolicy{}
	return &this
}

// NewAdapterConfigPolicyWithDefaults instantiates a new AdapterConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterConfigPolicyWithDefaults() *AdapterConfigPolicy {
	this := AdapterConfigPolicy{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AdapterConfigPolicy) GetSettings() []AdapterAdapterConfig {
	if o == nil || o.Settings == nil {
		var ret []AdapterAdapterConfig
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterConfigPolicy) GetSettingsOk() (*[]AdapterAdapterConfig, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AdapterConfigPolicy) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []AdapterAdapterConfig and assigns it to the Settings field.
func (o *AdapterConfigPolicy) SetSettings(v []AdapterAdapterConfig) {
	o.Settings = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AdapterConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AdapterConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *AdapterConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterConfigPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterConfigPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *AdapterConfigPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *AdapterConfigPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o AdapterConfigPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.Settings != nil {
		toSerialize["Settings"] = o.Settings
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

func (o *AdapterConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterConfigPolicyWithoutEmbeddedStruct struct {
		Settings     *[]AdapterAdapterConfig               `json:"Settings,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varAdapterConfigPolicyWithoutEmbeddedStruct := AdapterConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varAdapterConfigPolicy := _AdapterConfigPolicy{}
		varAdapterConfigPolicy.Settings = varAdapterConfigPolicyWithoutEmbeddedStruct.Settings
		varAdapterConfigPolicy.Organization = varAdapterConfigPolicyWithoutEmbeddedStruct.Organization
		varAdapterConfigPolicy.Profiles = varAdapterConfigPolicyWithoutEmbeddedStruct.Profiles
		*o = AdapterConfigPolicy(varAdapterConfigPolicy)
	} else {
		return err
	}

	varAdapterConfigPolicy := _AdapterConfigPolicy{}

	err = json.Unmarshal(bytes, &varAdapterConfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varAdapterConfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Settings")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableAdapterConfigPolicy struct {
	value *AdapterConfigPolicy
	isSet bool
}

func (v NullableAdapterConfigPolicy) Get() *AdapterConfigPolicy {
	return v.value
}

func (v *NullableAdapterConfigPolicy) Set(val *AdapterConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterConfigPolicy(val *AdapterConfigPolicy) *NullableAdapterConfigPolicy {
	return &NullableAdapterConfigPolicy{value: val, isSet: true}
}

func (v NullableAdapterConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
