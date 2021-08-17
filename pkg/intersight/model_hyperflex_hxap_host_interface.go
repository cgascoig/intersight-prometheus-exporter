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

// HyperflexHxapHostInterface A HyperFlex Application Platform compute host interface entity that can be connected by HxapHostVswitch.
type HyperflexHxapHostInterface struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                     `json:"ObjectType"`
	BondState  NullableHyperflexBondState `json:"BondState,omitempty"`
	// The UUID of the host to which this interface belongs to.
	HostName *string `json:"HostName,omitempty"`
	// The UUID of the host to which this interface belongs to.
	HostUuid *string `json:"HostUuid,omitempty"`
	// A hint of the interface type, such as \"ovs-bond\", \"device\", \"openvswitch\".
	IfType      *string  `json:"IfType,omitempty"`
	IpAddresses []string `json:"IpAddresses,omitempty"`
	// Link state information such as \"up\", \"down\". * `unknown` - The interface line is unknown. * `up` - The interface line is up. * `down` - The interface line is down. * `degraded` - For a bond/team interface, not all member interface is up.
	LinkState *string `json:"LinkState,omitempty"`
	// The MAC address of the interface.
	Mac *string `json:"Mac,omitempty"`
	// The MTU size of the interface.
	Mtu *int64 `json:"Mtu,omitempty"`
	// The name of the host to which this interface belongs to.
	Name *string `json:"Name,omitempty"`
	// A list of vlans allowed on this interface.
	Vlans                *string                            `json:"Vlans,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship  `json:"Cluster,omitempty"`
	DvUplink             *HyperflexHxapDvUplinkRelationship `json:"DvUplink,omitempty"`
	Host                 *HyperflexHxapHostRelationship     `json:"Host,omitempty"`
	Network              *HyperflexHxapNetworkRelationship  `json:"Network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapHostInterface HyperflexHxapHostInterface

// NewHyperflexHxapHostInterface instantiates a new HyperflexHxapHostInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapHostInterface(classId string, objectType string) *HyperflexHxapHostInterface {
	this := HyperflexHxapHostInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var linkState string = "unknown"
	this.LinkState = &linkState
	return &this
}

// NewHyperflexHxapHostInterfaceWithDefaults instantiates a new HyperflexHxapHostInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapHostInterfaceWithDefaults() *HyperflexHxapHostInterface {
	this := HyperflexHxapHostInterface{}
	var classId string = "hyperflex.HxapHostInterface"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapHostInterface"
	this.ObjectType = objectType
	var linkState string = "unknown"
	this.LinkState = &linkState
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapHostInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapHostInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapHostInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapHostInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBondState returns the BondState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapHostInterface) GetBondState() HyperflexBondState {
	if o == nil || o.BondState.Get() == nil {
		var ret HyperflexBondState
		return ret
	}
	return *o.BondState.Get()
}

// GetBondStateOk returns a tuple with the BondState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapHostInterface) GetBondStateOk() (*HyperflexBondState, bool) {
	if o == nil {
		return nil, false
	}
	return o.BondState.Get(), o.BondState.IsSet()
}

// HasBondState returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasBondState() bool {
	if o != nil && o.BondState.IsSet() {
		return true
	}

	return false
}

// SetBondState gets a reference to the given NullableHyperflexBondState and assigns it to the BondState field.
func (o *HyperflexHxapHostInterface) SetBondState(v HyperflexBondState) {
	o.BondState.Set(&v)
}

// SetBondStateNil sets the value for BondState to be an explicit nil
func (o *HyperflexHxapHostInterface) SetBondStateNil() {
	o.BondState.Set(nil)
}

// UnsetBondState ensures that no value is present for BondState, not even an explicit nil
func (o *HyperflexHxapHostInterface) UnsetBondState() {
	o.BondState.Unset()
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexHxapHostInterface) SetHostName(v string) {
	o.HostName = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetHostUuid() string {
	if o == nil || o.HostUuid == nil {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetHostUuidOk() (*string, bool) {
	if o == nil || o.HostUuid == nil {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasHostUuid() bool {
	if o != nil && o.HostUuid != nil {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *HyperflexHxapHostInterface) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetIfType returns the IfType field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetIfType() string {
	if o == nil || o.IfType == nil {
		var ret string
		return ret
	}
	return *o.IfType
}

// GetIfTypeOk returns a tuple with the IfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetIfTypeOk() (*string, bool) {
	if o == nil || o.IfType == nil {
		return nil, false
	}
	return o.IfType, true
}

// HasIfType returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasIfType() bool {
	if o != nil && o.IfType != nil {
		return true
	}

	return false
}

// SetIfType gets a reference to the given string and assigns it to the IfType field.
func (o *HyperflexHxapHostInterface) SetIfType(v string) {
	o.IfType = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapHostInterface) GetIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapHostInterface) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *HyperflexHxapHostInterface) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetLinkState() string {
	if o == nil || o.LinkState == nil {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetLinkStateOk() (*string, bool) {
	if o == nil || o.LinkState == nil {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasLinkState() bool {
	if o != nil && o.LinkState != nil {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *HyperflexHxapHostInterface) SetLinkState(v string) {
	o.LinkState = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *HyperflexHxapHostInterface) SetMac(v string) {
	o.Mac = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *HyperflexHxapHostInterface) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxapHostInterface) SetName(v string) {
	o.Name = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetVlans() string {
	if o == nil || o.Vlans == nil {
		var ret string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetVlansOk() (*string, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given string and assigns it to the Vlans field.
func (o *HyperflexHxapHostInterface) SetVlans(v string) {
	o.Vlans = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapHostInterface) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetDvUplink() HyperflexHxapDvUplinkRelationship {
	if o == nil || o.DvUplink == nil {
		var ret HyperflexHxapDvUplinkRelationship
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetDvUplinkOk() (*HyperflexHxapDvUplinkRelationship, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given HyperflexHxapDvUplinkRelationship and assigns it to the DvUplink field.
func (o *HyperflexHxapHostInterface) SetDvUplink(v HyperflexHxapDvUplinkRelationship) {
	o.DvUplink = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetHost() HyperflexHxapHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHxapHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetHostOk() (*HyperflexHxapHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHxapHostRelationship and assigns it to the Host field.
func (o *HyperflexHxapHostInterface) SetHost(v HyperflexHxapHostRelationship) {
	o.Host = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *HyperflexHxapHostInterface) GetNetwork() HyperflexHxapNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret HyperflexHxapNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostInterface) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *HyperflexHxapHostInterface) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given HyperflexHxapNetworkRelationship and assigns it to the Network field.
func (o *HyperflexHxapHostInterface) SetNetwork(v HyperflexHxapNetworkRelationship) {
	o.Network = &v
}

func (o HyperflexHxapHostInterface) MarshalJSON() ([]byte, error) {
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
	if o.BondState.IsSet() {
		toSerialize["BondState"] = o.BondState.Get()
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.HostUuid != nil {
		toSerialize["HostUuid"] = o.HostUuid
	}
	if o.IfType != nil {
		toSerialize["IfType"] = o.IfType
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}
	if o.LinkState != nil {
		toSerialize["LinkState"] = o.LinkState
	}
	if o.Mac != nil {
		toSerialize["Mac"] = o.Mac
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Vlans != nil {
		toSerialize["Vlans"] = o.Vlans
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapHostInterface) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapHostInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                     `json:"ObjectType"`
		BondState  NullableHyperflexBondState `json:"BondState,omitempty"`
		// The UUID of the host to which this interface belongs to.
		HostName *string `json:"HostName,omitempty"`
		// The UUID of the host to which this interface belongs to.
		HostUuid *string `json:"HostUuid,omitempty"`
		// A hint of the interface type, such as \"ovs-bond\", \"device\", \"openvswitch\".
		IfType      *string  `json:"IfType,omitempty"`
		IpAddresses []string `json:"IpAddresses,omitempty"`
		// Link state information such as \"up\", \"down\". * `unknown` - The interface line is unknown. * `up` - The interface line is up. * `down` - The interface line is down. * `degraded` - For a bond/team interface, not all member interface is up.
		LinkState *string `json:"LinkState,omitempty"`
		// The MAC address of the interface.
		Mac *string `json:"Mac,omitempty"`
		// The MTU size of the interface.
		Mtu *int64 `json:"Mtu,omitempty"`
		// The name of the host to which this interface belongs to.
		Name *string `json:"Name,omitempty"`
		// A list of vlans allowed on this interface.
		Vlans    *string                            `json:"Vlans,omitempty"`
		Cluster  *HyperflexHxapClusterRelationship  `json:"Cluster,omitempty"`
		DvUplink *HyperflexHxapDvUplinkRelationship `json:"DvUplink,omitempty"`
		Host     *HyperflexHxapHostRelationship     `json:"Host,omitempty"`
		Network  *HyperflexHxapNetworkRelationship  `json:"Network,omitempty"`
	}

	varHyperflexHxapHostInterfaceWithoutEmbeddedStruct := HyperflexHxapHostInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHostInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapHostInterface := _HyperflexHxapHostInterface{}
		varHyperflexHxapHostInterface.ClassId = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.ClassId
		varHyperflexHxapHostInterface.ObjectType = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapHostInterface.BondState = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.BondState
		varHyperflexHxapHostInterface.HostName = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.HostName
		varHyperflexHxapHostInterface.HostUuid = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.HostUuid
		varHyperflexHxapHostInterface.IfType = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.IfType
		varHyperflexHxapHostInterface.IpAddresses = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.IpAddresses
		varHyperflexHxapHostInterface.LinkState = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.LinkState
		varHyperflexHxapHostInterface.Mac = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Mac
		varHyperflexHxapHostInterface.Mtu = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Mtu
		varHyperflexHxapHostInterface.Name = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Name
		varHyperflexHxapHostInterface.Vlans = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Vlans
		varHyperflexHxapHostInterface.Cluster = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Cluster
		varHyperflexHxapHostInterface.DvUplink = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.DvUplink
		varHyperflexHxapHostInterface.Host = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Host
		varHyperflexHxapHostInterface.Network = varHyperflexHxapHostInterfaceWithoutEmbeddedStruct.Network
		*o = HyperflexHxapHostInterface(varHyperflexHxapHostInterface)
	} else {
		return err
	}

	varHyperflexHxapHostInterface := _HyperflexHxapHostInterface{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHostInterface)
	if err == nil {
		o.MoBaseMo = varHyperflexHxapHostInterface.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BondState")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "HostUuid")
		delete(additionalProperties, "IfType")
		delete(additionalProperties, "IpAddresses")
		delete(additionalProperties, "LinkState")
		delete(additionalProperties, "Mac")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Vlans")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Network")

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

type NullableHyperflexHxapHostInterface struct {
	value *HyperflexHxapHostInterface
	isSet bool
}

func (v NullableHyperflexHxapHostInterface) Get() *HyperflexHxapHostInterface {
	return v.value
}

func (v *NullableHyperflexHxapHostInterface) Set(val *HyperflexHxapHostInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapHostInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapHostInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapHostInterface(val *HyperflexHxapHostInterface) *NullableHyperflexHxapHostInterface {
	return &NullableHyperflexHxapHostInterface{value: val, isSet: true}
}

func (v NullableHyperflexHxapHostInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapHostInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
