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

// IppoolShadowPoolAllOf Definition of the list of properties defined in 'ippool.ShadowPool', excluding properties defined in parent classes.
type IppoolShadowPoolAllOf struct {
	IpV4Blocks *[]IppoolIpBlock  `json:"IpV4Blocks,omitempty"`
	IpV4Config *IppoolIpV4Config `json:"IpV4Config,omitempty"`
	// Number of IPv4 addresses currently in use.
	V4Assigned *int64 `json:"V4Assigned,omitempty"`
	// Number of IPv4 addresses in this pool.
	V4Size *int64 `json:"V4Size,omitempty"`
	// Number of IPv6 addresses currently in use.
	V6Assigned *int64 `json:"V6Assigned,omitempty"`
	// Number of IPv6 addresses in this pool.
	V6Size *int64                  `json:"V6Size,omitempty"`
	Pool   *IppoolPoolRelationship `json:"Pool,omitempty"`
	// An array of relationships to ippoolShadowBlock resources.
	V4BlockHeads         []IppoolShadowBlockRelationship `json:"V4BlockHeads,omitempty"`
	Vrf                  *VrfVrfRelationship             `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowPoolAllOf IppoolShadowPoolAllOf

// NewIppoolShadowPoolAllOf instantiates a new IppoolShadowPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowPoolAllOf() *IppoolShadowPoolAllOf {
	this := IppoolShadowPoolAllOf{}
	return &this
}

// NewIppoolShadowPoolAllOfWithDefaults instantiates a new IppoolShadowPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowPoolAllOfWithDefaults() *IppoolShadowPoolAllOf {
	this := IppoolShadowPoolAllOf{}
	return &this
}

// GetIpV4Blocks returns the IpV4Blocks field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetIpV4Blocks() []IppoolIpBlock {
	if o == nil || o.IpV4Blocks == nil {
		var ret []IppoolIpBlock
		return ret
	}
	return *o.IpV4Blocks
}

// GetIpV4BlocksOk returns a tuple with the IpV4Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpBlock, bool) {
	if o == nil || o.IpV4Blocks == nil {
		return nil, false
	}
	return o.IpV4Blocks, true
}

// HasIpV4Blocks returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasIpV4Blocks() bool {
	if o != nil && o.IpV4Blocks != nil {
		return true
	}

	return false
}

// SetIpV4Blocks gets a reference to the given []IppoolIpBlock and assigns it to the IpV4Blocks field.
func (o *IppoolShadowPoolAllOf) SetIpV4Blocks(v []IppoolIpBlock) {
	o.IpV4Blocks = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given IppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolShadowPoolAllOf) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config = &v
}

// GetV4Assigned returns the V4Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetV4Assigned() int64 {
	if o == nil || o.V4Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V4Assigned
}

// GetV4AssignedOk returns a tuple with the V4Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetV4AssignedOk() (*int64, bool) {
	if o == nil || o.V4Assigned == nil {
		return nil, false
	}
	return o.V4Assigned, true
}

// HasV4Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasV4Assigned() bool {
	if o != nil && o.V4Assigned != nil {
		return true
	}

	return false
}

// SetV4Assigned gets a reference to the given int64 and assigns it to the V4Assigned field.
func (o *IppoolShadowPoolAllOf) SetV4Assigned(v int64) {
	o.V4Assigned = &v
}

// GetV4Size returns the V4Size field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetV4Size() int64 {
	if o == nil || o.V4Size == nil {
		var ret int64
		return ret
	}
	return *o.V4Size
}

// GetV4SizeOk returns a tuple with the V4Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetV4SizeOk() (*int64, bool) {
	if o == nil || o.V4Size == nil {
		return nil, false
	}
	return o.V4Size, true
}

// HasV4Size returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasV4Size() bool {
	if o != nil && o.V4Size != nil {
		return true
	}

	return false
}

// SetV4Size gets a reference to the given int64 and assigns it to the V4Size field.
func (o *IppoolShadowPoolAllOf) SetV4Size(v int64) {
	o.V4Size = &v
}

// GetV6Assigned returns the V6Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetV6Assigned() int64 {
	if o == nil || o.V6Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V6Assigned
}

// GetV6AssignedOk returns a tuple with the V6Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetV6AssignedOk() (*int64, bool) {
	if o == nil || o.V6Assigned == nil {
		return nil, false
	}
	return o.V6Assigned, true
}

// HasV6Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasV6Assigned() bool {
	if o != nil && o.V6Assigned != nil {
		return true
	}

	return false
}

// SetV6Assigned gets a reference to the given int64 and assigns it to the V6Assigned field.
func (o *IppoolShadowPoolAllOf) SetV6Assigned(v int64) {
	o.V6Assigned = &v
}

// GetV6Size returns the V6Size field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetV6Size() int64 {
	if o == nil || o.V6Size == nil {
		var ret int64
		return ret
	}
	return *o.V6Size
}

// GetV6SizeOk returns a tuple with the V6Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetV6SizeOk() (*int64, bool) {
	if o == nil || o.V6Size == nil {
		return nil, false
	}
	return o.V6Size, true
}

// HasV6Size returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasV6Size() bool {
	if o != nil && o.V6Size != nil {
		return true
	}

	return false
}

// SetV6Size gets a reference to the given int64 and assigns it to the V6Size field.
func (o *IppoolShadowPoolAllOf) SetV6Size(v int64) {
	o.V6Size = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowPoolAllOf) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetV4BlockHeads returns the V4BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPoolAllOf) GetV4BlockHeads() []IppoolShadowBlockRelationship {
	if o == nil {
		var ret []IppoolShadowBlockRelationship
		return ret
	}
	return o.V4BlockHeads
}

// GetV4BlockHeadsOk returns a tuple with the V4BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPoolAllOf) GetV4BlockHeadsOk() (*[]IppoolShadowBlockRelationship, bool) {
	if o == nil || o.V4BlockHeads == nil {
		return nil, false
	}
	return &o.V4BlockHeads, true
}

// HasV4BlockHeads returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasV4BlockHeads() bool {
	if o != nil && o.V4BlockHeads != nil {
		return true
	}

	return false
}

// SetV4BlockHeads gets a reference to the given []IppoolShadowBlockRelationship and assigns it to the V4BlockHeads field.
func (o *IppoolShadowPoolAllOf) SetV4BlockHeads(v []IppoolShadowBlockRelationship) {
	o.V4BlockHeads = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolShadowPoolAllOf) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPoolAllOf) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolShadowPoolAllOf) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolShadowPoolAllOf) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolShadowPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpV4Blocks != nil {
		toSerialize["IpV4Blocks"] = o.IpV4Blocks
	}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}
	if o.V4Assigned != nil {
		toSerialize["V4Assigned"] = o.V4Assigned
	}
	if o.V4Size != nil {
		toSerialize["V4Size"] = o.V4Size
	}
	if o.V6Assigned != nil {
		toSerialize["V6Assigned"] = o.V6Assigned
	}
	if o.V6Size != nil {
		toSerialize["V6Size"] = o.V6Size
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.V4BlockHeads != nil {
		toSerialize["V4BlockHeads"] = o.V4BlockHeads
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolShadowPoolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolShadowPoolAllOf := _IppoolShadowPoolAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolShadowPoolAllOf); err == nil {
		*o = IppoolShadowPoolAllOf(varIppoolShadowPoolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IpV4Blocks")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "V4Assigned")
		delete(additionalProperties, "V4Size")
		delete(additionalProperties, "V6Assigned")
		delete(additionalProperties, "V6Size")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "V4BlockHeads")
		delete(additionalProperties, "Vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolShadowPoolAllOf struct {
	value *IppoolShadowPoolAllOf
	isSet bool
}

func (v NullableIppoolShadowPoolAllOf) Get() *IppoolShadowPoolAllOf {
	return v.value
}

func (v *NullableIppoolShadowPoolAllOf) Set(val *IppoolShadowPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowPoolAllOf(val *IppoolShadowPoolAllOf) *NullableIppoolShadowPoolAllOf {
	return &NullableIppoolShadowPoolAllOf{value: val, isSet: true}
}

func (v NullableIppoolShadowPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
