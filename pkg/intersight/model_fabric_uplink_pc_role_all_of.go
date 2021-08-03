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

// FabricUplinkPcRoleAllOf Definition of the list of properties defined in 'fabric.UplinkPcRole', excluding properties defined in parent classes.
type FabricUplinkPcRoleAllOf struct {
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Admin configured state for UDLD for this port. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	UdldAdminState       *string `json:"UdldAdminState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricUplinkPcRoleAllOf FabricUplinkPcRoleAllOf

// NewFabricUplinkPcRoleAllOf instantiates a new FabricUplinkPcRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUplinkPcRoleAllOf() *FabricUplinkPcRoleAllOf {
	this := FabricUplinkPcRoleAllOf{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// NewFabricUplinkPcRoleAllOfWithDefaults instantiates a new FabricUplinkPcRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUplinkPcRoleAllOfWithDefaults() *FabricUplinkPcRoleAllOf {
	this := FabricUplinkPcRoleAllOf{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var udldAdminState string = "Disabled"
	this.UdldAdminState = &udldAdminState
	return &this
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricUplinkPcRoleAllOf) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRoleAllOf) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricUplinkPcRoleAllOf) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricUplinkPcRoleAllOf) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetUdldAdminState returns the UdldAdminState field value if set, zero value otherwise.
func (o *FabricUplinkPcRoleAllOf) GetUdldAdminState() string {
	if o == nil || o.UdldAdminState == nil {
		var ret string
		return ret
	}
	return *o.UdldAdminState
}

// GetUdldAdminStateOk returns a tuple with the UdldAdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRoleAllOf) GetUdldAdminStateOk() (*string, bool) {
	if o == nil || o.UdldAdminState == nil {
		return nil, false
	}
	return o.UdldAdminState, true
}

// HasUdldAdminState returns a boolean if a field has been set.
func (o *FabricUplinkPcRoleAllOf) HasUdldAdminState() bool {
	if o != nil && o.UdldAdminState != nil {
		return true
	}

	return false
}

// SetUdldAdminState gets a reference to the given string and assigns it to the UdldAdminState field.
func (o *FabricUplinkPcRoleAllOf) SetUdldAdminState(v string) {
	o.UdldAdminState = &v
}

func (o FabricUplinkPcRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.UdldAdminState != nil {
		toSerialize["UdldAdminState"] = o.UdldAdminState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricUplinkPcRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricUplinkPcRoleAllOf := _FabricUplinkPcRoleAllOf{}

	if err = json.Unmarshal(bytes, &varFabricUplinkPcRoleAllOf); err == nil {
		*o = FabricUplinkPcRoleAllOf(varFabricUplinkPcRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "UdldAdminState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricUplinkPcRoleAllOf struct {
	value *FabricUplinkPcRoleAllOf
	isSet bool
}

func (v NullableFabricUplinkPcRoleAllOf) Get() *FabricUplinkPcRoleAllOf {
	return v.value
}

func (v *NullableFabricUplinkPcRoleAllOf) Set(val *FabricUplinkPcRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUplinkPcRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUplinkPcRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUplinkPcRoleAllOf(val *FabricUplinkPcRoleAllOf) *NullableFabricUplinkPcRoleAllOf {
	return &NullableFabricUplinkPcRoleAllOf{value: val, isSet: true}
}

func (v NullableFabricUplinkPcRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUplinkPcRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
