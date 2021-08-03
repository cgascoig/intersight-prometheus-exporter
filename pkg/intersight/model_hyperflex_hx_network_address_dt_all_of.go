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

// HyperflexHxNetworkAddressDtAllOf Definition of the list of properties defined in 'hyperflex.HxNetworkAddressDt', excluding properties defined in parent classes.
type HyperflexHxNetworkAddressDtAllOf struct {
	Address              *string `json:"Address,omitempty"`
	Fqdn                 *string `json:"Fqdn,omitempty"`
	Ip                   *string `json:"Ip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxNetworkAddressDtAllOf HyperflexHxNetworkAddressDtAllOf

// NewHyperflexHxNetworkAddressDtAllOf instantiates a new HyperflexHxNetworkAddressDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxNetworkAddressDtAllOf() *HyperflexHxNetworkAddressDtAllOf {
	this := HyperflexHxNetworkAddressDtAllOf{}
	return &this
}

// NewHyperflexHxNetworkAddressDtAllOfWithDefaults instantiates a new HyperflexHxNetworkAddressDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxNetworkAddressDtAllOfWithDefaults() *HyperflexHxNetworkAddressDtAllOf {
	this := HyperflexHxNetworkAddressDtAllOf{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDtAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HyperflexHxNetworkAddressDtAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDtAllOf) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *HyperflexHxNetworkAddressDtAllOf) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDtAllOf) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDtAllOf) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *HyperflexHxNetworkAddressDtAllOf) SetIp(v string) {
	o.Ip = &v
}

func (o HyperflexHxNetworkAddressDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Fqdn != nil {
		toSerialize["Fqdn"] = o.Fqdn
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxNetworkAddressDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxNetworkAddressDtAllOf := _HyperflexHxNetworkAddressDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxNetworkAddressDtAllOf); err == nil {
		*o = HyperflexHxNetworkAddressDtAllOf(varHyperflexHxNetworkAddressDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Fqdn")
		delete(additionalProperties, "Ip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxNetworkAddressDtAllOf struct {
	value *HyperflexHxNetworkAddressDtAllOf
	isSet bool
}

func (v NullableHyperflexHxNetworkAddressDtAllOf) Get() *HyperflexHxNetworkAddressDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxNetworkAddressDtAllOf) Set(val *HyperflexHxNetworkAddressDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxNetworkAddressDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxNetworkAddressDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxNetworkAddressDtAllOf(val *HyperflexHxNetworkAddressDtAllOf) *NullableHyperflexHxNetworkAddressDtAllOf {
	return &NullableHyperflexHxNetworkAddressDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxNetworkAddressDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxNetworkAddressDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
