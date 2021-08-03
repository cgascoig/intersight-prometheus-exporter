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

// HyperflexAutoSupportPolicy A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
type HyperflexAutoSupportPolicy struct {
	PolicyAbstractPolicy
	// Enable or disable Auto Support.
	AdminState *bool `json:"AdminState,omitempty"`
	// The recipient email address for support tickets.
	ServiceTicketReceipient *string `json:"ServiceTicketReceipient,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAutoSupportPolicy HyperflexAutoSupportPolicy

// NewHyperflexAutoSupportPolicy instantiates a new HyperflexAutoSupportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAutoSupportPolicy() *HyperflexAutoSupportPolicy {
	this := HyperflexAutoSupportPolicy{}
	return &this
}

// NewHyperflexAutoSupportPolicyWithDefaults instantiates a new HyperflexAutoSupportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAutoSupportPolicyWithDefaults() *HyperflexAutoSupportPolicy {
	this := HyperflexAutoSupportPolicy{}
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicy) GetAdminState() bool {
	if o == nil || o.AdminState == nil {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicy) GetAdminStateOk() (*bool, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicy) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexAutoSupportPolicy) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetServiceTicketReceipient returns the ServiceTicketReceipient field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicy) GetServiceTicketReceipient() string {
	if o == nil || o.ServiceTicketReceipient == nil {
		var ret string
		return ret
	}
	return *o.ServiceTicketReceipient
}

// GetServiceTicketReceipientOk returns a tuple with the ServiceTicketReceipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicy) GetServiceTicketReceipientOk() (*string, bool) {
	if o == nil || o.ServiceTicketReceipient == nil {
		return nil, false
	}
	return o.ServiceTicketReceipient, true
}

// HasServiceTicketReceipient returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicy) HasServiceTicketReceipient() bool {
	if o != nil && o.ServiceTicketReceipient != nil {
		return true
	}

	return false
}

// SetServiceTicketReceipient gets a reference to the given string and assigns it to the ServiceTicketReceipient field.
func (o *HyperflexAutoSupportPolicy) SetServiceTicketReceipient(v string) {
	o.ServiceTicketReceipient = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAutoSupportPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAutoSupportPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexAutoSupportPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexAutoSupportPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexAutoSupportPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ServiceTicketReceipient != nil {
		toSerialize["ServiceTicketReceipient"] = o.ServiceTicketReceipient
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

func (o *HyperflexAutoSupportPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexAutoSupportPolicyWithoutEmbeddedStruct struct {
		// Enable or disable Auto Support.
		AdminState *bool `json:"AdminState,omitempty"`
		// The recipient email address for support tickets.
		ServiceTicketReceipient *string `json:"ServiceTicketReceipient,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexAutoSupportPolicyWithoutEmbeddedStruct := HyperflexAutoSupportPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexAutoSupportPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAutoSupportPolicy := _HyperflexAutoSupportPolicy{}
		varHyperflexAutoSupportPolicy.AdminState = varHyperflexAutoSupportPolicyWithoutEmbeddedStruct.AdminState
		varHyperflexAutoSupportPolicy.ServiceTicketReceipient = varHyperflexAutoSupportPolicyWithoutEmbeddedStruct.ServiceTicketReceipient
		varHyperflexAutoSupportPolicy.ClusterProfiles = varHyperflexAutoSupportPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexAutoSupportPolicy.Organization = varHyperflexAutoSupportPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexAutoSupportPolicy(varHyperflexAutoSupportPolicy)
	} else {
		return err
	}

	varHyperflexAutoSupportPolicy := _HyperflexAutoSupportPolicy{}

	err = json.Unmarshal(bytes, &varHyperflexAutoSupportPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexAutoSupportPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ServiceTicketReceipient")
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

type NullableHyperflexAutoSupportPolicy struct {
	value *HyperflexAutoSupportPolicy
	isSet bool
}

func (v NullableHyperflexAutoSupportPolicy) Get() *HyperflexAutoSupportPolicy {
	return v.value
}

func (v *NullableHyperflexAutoSupportPolicy) Set(val *HyperflexAutoSupportPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAutoSupportPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAutoSupportPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAutoSupportPolicy(val *HyperflexAutoSupportPolicy) *NullableHyperflexAutoSupportPolicy {
	return &NullableHyperflexAutoSupportPolicy{value: val, isSet: true}
}

func (v NullableHyperflexAutoSupportPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAutoSupportPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
