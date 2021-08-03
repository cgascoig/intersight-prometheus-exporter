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

// CapabilityIoCardDescriptorAllOf Definition of the list of properties defined in 'capability.IoCardDescriptor', excluding properties defined in parent classes.
type CapabilityIoCardDescriptorAllOf struct {
	// Number of hif ports per blade for the iocard module.
	NumHifPorts *int64 `json:"NumHifPorts,omitempty"`
	// Revision for the iocard module.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityIoCardDescriptorAllOf CapabilityIoCardDescriptorAllOf

// NewCapabilityIoCardDescriptorAllOf instantiates a new CapabilityIoCardDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityIoCardDescriptorAllOf() *CapabilityIoCardDescriptorAllOf {
	this := CapabilityIoCardDescriptorAllOf{}
	return &this
}

// NewCapabilityIoCardDescriptorAllOfWithDefaults instantiates a new CapabilityIoCardDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityIoCardDescriptorAllOfWithDefaults() *CapabilityIoCardDescriptorAllOf {
	this := CapabilityIoCardDescriptorAllOf{}
	return &this
}

// GetNumHifPorts returns the NumHifPorts field value if set, zero value otherwise.
func (o *CapabilityIoCardDescriptorAllOf) GetNumHifPorts() int64 {
	if o == nil || o.NumHifPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumHifPorts
}

// GetNumHifPortsOk returns a tuple with the NumHifPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptorAllOf) GetNumHifPortsOk() (*int64, bool) {
	if o == nil || o.NumHifPorts == nil {
		return nil, false
	}
	return o.NumHifPorts, true
}

// HasNumHifPorts returns a boolean if a field has been set.
func (o *CapabilityIoCardDescriptorAllOf) HasNumHifPorts() bool {
	if o != nil && o.NumHifPorts != nil {
		return true
	}

	return false
}

// SetNumHifPorts gets a reference to the given int64 and assigns it to the NumHifPorts field.
func (o *CapabilityIoCardDescriptorAllOf) SetNumHifPorts(v int64) {
	o.NumHifPorts = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilityIoCardDescriptorAllOf) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptorAllOf) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilityIoCardDescriptorAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilityIoCardDescriptorAllOf) SetRevision(v string) {
	o.Revision = &v
}

func (o CapabilityIoCardDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumHifPorts != nil {
		toSerialize["NumHifPorts"] = o.NumHifPorts
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityIoCardDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityIoCardDescriptorAllOf := _CapabilityIoCardDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityIoCardDescriptorAllOf); err == nil {
		*o = CapabilityIoCardDescriptorAllOf(varCapabilityIoCardDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "NumHifPorts")
		delete(additionalProperties, "Revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityIoCardDescriptorAllOf struct {
	value *CapabilityIoCardDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityIoCardDescriptorAllOf) Get() *CapabilityIoCardDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityIoCardDescriptorAllOf) Set(val *CapabilityIoCardDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityIoCardDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityIoCardDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityIoCardDescriptorAllOf(val *CapabilityIoCardDescriptorAllOf) *NullableCapabilityIoCardDescriptorAllOf {
	return &NullableCapabilityIoCardDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityIoCardDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityIoCardDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
