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

// BiosBootDeviceAllOf Definition of the list of properties defined in 'bios.BootDevice', excluding properties defined in parent classes.
type BiosBootDeviceAllOf struct {
	// Name of the Configured Boot Device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Type of the Configured Boot Device.
	DeviceType           *string                              `json:"DeviceType,omitempty"`
	BiosSystemBootOrder  *BiosSystemBootOrderRelationship     `json:"BiosSystemBootOrder,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosBootDeviceAllOf BiosBootDeviceAllOf

// NewBiosBootDeviceAllOf instantiates a new BiosBootDeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosBootDeviceAllOf() *BiosBootDeviceAllOf {
	this := BiosBootDeviceAllOf{}
	return &this
}

// NewBiosBootDeviceAllOfWithDefaults instantiates a new BiosBootDeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosBootDeviceAllOfWithDefaults() *BiosBootDeviceAllOf {
	this := BiosBootDeviceAllOf{}
	return &this
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BiosBootDeviceAllOf) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDeviceAllOf) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BiosBootDeviceAllOf) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BiosBootDeviceAllOf) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *BiosBootDeviceAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDeviceAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *BiosBootDeviceAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *BiosBootDeviceAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetBiosSystemBootOrder returns the BiosSystemBootOrder field value if set, zero value otherwise.
func (o *BiosBootDeviceAllOf) GetBiosSystemBootOrder() BiosSystemBootOrderRelationship {
	if o == nil || o.BiosSystemBootOrder == nil {
		var ret BiosSystemBootOrderRelationship
		return ret
	}
	return *o.BiosSystemBootOrder
}

// GetBiosSystemBootOrderOk returns a tuple with the BiosSystemBootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDeviceAllOf) GetBiosSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool) {
	if o == nil || o.BiosSystemBootOrder == nil {
		return nil, false
	}
	return o.BiosSystemBootOrder, true
}

// HasBiosSystemBootOrder returns a boolean if a field has been set.
func (o *BiosBootDeviceAllOf) HasBiosSystemBootOrder() bool {
	if o != nil && o.BiosSystemBootOrder != nil {
		return true
	}

	return false
}

// SetBiosSystemBootOrder gets a reference to the given BiosSystemBootOrderRelationship and assigns it to the BiosSystemBootOrder field.
func (o *BiosBootDeviceAllOf) SetBiosSystemBootOrder(v BiosSystemBootOrderRelationship) {
	o.BiosSystemBootOrder = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosBootDeviceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDeviceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosBootDeviceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosBootDeviceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosBootDeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.BiosSystemBootOrder != nil {
		toSerialize["BiosSystemBootOrder"] = o.BiosSystemBootOrder
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosBootDeviceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBiosBootDeviceAllOf := _BiosBootDeviceAllOf{}

	if err = json.Unmarshal(bytes, &varBiosBootDeviceAllOf); err == nil {
		*o = BiosBootDeviceAllOf(varBiosBootDeviceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "BiosSystemBootOrder")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBiosBootDeviceAllOf struct {
	value *BiosBootDeviceAllOf
	isSet bool
}

func (v NullableBiosBootDeviceAllOf) Get() *BiosBootDeviceAllOf {
	return v.value
}

func (v *NullableBiosBootDeviceAllOf) Set(val *BiosBootDeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosBootDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosBootDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosBootDeviceAllOf(val *BiosBootDeviceAllOf) *NullableBiosBootDeviceAllOf {
	return &NullableBiosBootDeviceAllOf{value: val, isSet: true}
}

func (v NullableBiosBootDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosBootDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
