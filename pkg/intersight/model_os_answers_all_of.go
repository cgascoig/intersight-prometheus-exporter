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

// OsAnswersAllOf Definition of the list of properties defined in 'os.Answers', excluding properties defined in parent classes.
type OsAnswersAllOf struct {
	// If the source of the answers is a static file, the content of the file is stored as value in this property. The value is mandatory only when the 'Source' property has been set to 'File'.
	AnswerFile *string `json:"AnswerFile,omitempty"`
	// Hostname to be configured for the server in the OS.
	Hostname *string `json:"Hostname,omitempty"`
	// IP configuration type. Values are Static or Dynamic configuration of IP. In case of static IP configuration, IP address, gateway and other details need to be populated. In case of dynamic the IP configuration is obtained dynamically from DHCP. * `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway. * `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.
	IpConfigType    *string            `json:"IpConfigType,omitempty"`
	IpConfiguration *OsIpConfiguration `json:"IpConfiguration,omitempty"`
	// Indicates whether the value of the 'answerFile' property has been set.
	IsAnswerFileSet *bool `json:"IsAnswerFileSet,omitempty"`
	// Enable to indicate Root Password provided is encrypted.
	IsRootPasswordCrypted *bool `json:"IsRootPasswordCrypted,omitempty"`
	// Indicates whether the value of the 'rootPassword' property has been set.
	IsRootPasswordSet *bool `json:"IsRootPasswordSet,omitempty"`
	// IP address of the name server to be configured in the OS.
	Nameserver *string `json:"Nameserver,omitempty"`
	// The product key to be used for a specific version of Windows installation.
	ProductKey *string `json:"ProductKey,omitempty"`
	// Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password. Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password. For more details on encrypting passwords, see Help Center.
	RootPassword *string `json:"RootPassword,omitempty"`
	// Answer values can be provided from three sources - Embedded in OS image, static file, or as placeholder values for an answer file template. Source of the answers is given as value, Embedded/File/Template. 'Embedded' option indicates that the answer file is embedded within the OS Image. 'File' option indicates that the answers are provided as a file. 'Template' indicates that the placeholders in the selected os.ConfigurationFile MO are replaced with values provided as os.Answers MO. * `None` - Indicates that answers is not sent and values must be populated from Install Template.   * `Embedded` - Indicates that the answer file is embedded within OS image. * `File` - Indicates that the answer file is a static content that has all thevalues populated. * `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsAnswersAllOf OsAnswersAllOf

// NewOsAnswersAllOf instantiates a new OsAnswersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsAnswersAllOf() *OsAnswersAllOf {
	this := OsAnswersAllOf{}
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var source string = "None"
	this.Source = &source
	return &this
}

// NewOsAnswersAllOfWithDefaults instantiates a new OsAnswersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsAnswersAllOfWithDefaults() *OsAnswersAllOf {
	this := OsAnswersAllOf{}
	var ipConfigType string = "static"
	this.IpConfigType = &ipConfigType
	var source string = "None"
	this.Source = &source
	return &this
}

// GetAnswerFile returns the AnswerFile field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetAnswerFile() string {
	if o == nil || o.AnswerFile == nil {
		var ret string
		return ret
	}
	return *o.AnswerFile
}

// GetAnswerFileOk returns a tuple with the AnswerFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetAnswerFileOk() (*string, bool) {
	if o == nil || o.AnswerFile == nil {
		return nil, false
	}
	return o.AnswerFile, true
}

// HasAnswerFile returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasAnswerFile() bool {
	if o != nil && o.AnswerFile != nil {
		return true
	}

	return false
}

// SetAnswerFile gets a reference to the given string and assigns it to the AnswerFile field.
func (o *OsAnswersAllOf) SetAnswerFile(v string) {
	o.AnswerFile = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *OsAnswersAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpConfigType returns the IpConfigType field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetIpConfigType() string {
	if o == nil || o.IpConfigType == nil {
		var ret string
		return ret
	}
	return *o.IpConfigType
}

// GetIpConfigTypeOk returns a tuple with the IpConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetIpConfigTypeOk() (*string, bool) {
	if o == nil || o.IpConfigType == nil {
		return nil, false
	}
	return o.IpConfigType, true
}

// HasIpConfigType returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasIpConfigType() bool {
	if o != nil && o.IpConfigType != nil {
		return true
	}

	return false
}

// SetIpConfigType gets a reference to the given string and assigns it to the IpConfigType field.
func (o *OsAnswersAllOf) SetIpConfigType(v string) {
	o.IpConfigType = &v
}

// GetIpConfiguration returns the IpConfiguration field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetIpConfiguration() OsIpConfiguration {
	if o == nil || o.IpConfiguration == nil {
		var ret OsIpConfiguration
		return ret
	}
	return *o.IpConfiguration
}

// GetIpConfigurationOk returns a tuple with the IpConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetIpConfigurationOk() (*OsIpConfiguration, bool) {
	if o == nil || o.IpConfiguration == nil {
		return nil, false
	}
	return o.IpConfiguration, true
}

// HasIpConfiguration returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasIpConfiguration() bool {
	if o != nil && o.IpConfiguration != nil {
		return true
	}

	return false
}

// SetIpConfiguration gets a reference to the given OsIpConfiguration and assigns it to the IpConfiguration field.
func (o *OsAnswersAllOf) SetIpConfiguration(v OsIpConfiguration) {
	o.IpConfiguration = &v
}

// GetIsAnswerFileSet returns the IsAnswerFileSet field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetIsAnswerFileSet() bool {
	if o == nil || o.IsAnswerFileSet == nil {
		var ret bool
		return ret
	}
	return *o.IsAnswerFileSet
}

// GetIsAnswerFileSetOk returns a tuple with the IsAnswerFileSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetIsAnswerFileSetOk() (*bool, bool) {
	if o == nil || o.IsAnswerFileSet == nil {
		return nil, false
	}
	return o.IsAnswerFileSet, true
}

// HasIsAnswerFileSet returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasIsAnswerFileSet() bool {
	if o != nil && o.IsAnswerFileSet != nil {
		return true
	}

	return false
}

// SetIsAnswerFileSet gets a reference to the given bool and assigns it to the IsAnswerFileSet field.
func (o *OsAnswersAllOf) SetIsAnswerFileSet(v bool) {
	o.IsAnswerFileSet = &v
}

// GetIsRootPasswordCrypted returns the IsRootPasswordCrypted field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetIsRootPasswordCrypted() bool {
	if o == nil || o.IsRootPasswordCrypted == nil {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordCrypted
}

// GetIsRootPasswordCryptedOk returns a tuple with the IsRootPasswordCrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetIsRootPasswordCryptedOk() (*bool, bool) {
	if o == nil || o.IsRootPasswordCrypted == nil {
		return nil, false
	}
	return o.IsRootPasswordCrypted, true
}

// HasIsRootPasswordCrypted returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasIsRootPasswordCrypted() bool {
	if o != nil && o.IsRootPasswordCrypted != nil {
		return true
	}

	return false
}

// SetIsRootPasswordCrypted gets a reference to the given bool and assigns it to the IsRootPasswordCrypted field.
func (o *OsAnswersAllOf) SetIsRootPasswordCrypted(v bool) {
	o.IsRootPasswordCrypted = &v
}

// GetIsRootPasswordSet returns the IsRootPasswordSet field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetIsRootPasswordSet() bool {
	if o == nil || o.IsRootPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsRootPasswordSet
}

// GetIsRootPasswordSetOk returns a tuple with the IsRootPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetIsRootPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsRootPasswordSet == nil {
		return nil, false
	}
	return o.IsRootPasswordSet, true
}

// HasIsRootPasswordSet returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasIsRootPasswordSet() bool {
	if o != nil && o.IsRootPasswordSet != nil {
		return true
	}

	return false
}

// SetIsRootPasswordSet gets a reference to the given bool and assigns it to the IsRootPasswordSet field.
func (o *OsAnswersAllOf) SetIsRootPasswordSet(v bool) {
	o.IsRootPasswordSet = &v
}

// GetNameserver returns the Nameserver field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetNameserver() string {
	if o == nil || o.Nameserver == nil {
		var ret string
		return ret
	}
	return *o.Nameserver
}

// GetNameserverOk returns a tuple with the Nameserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetNameserverOk() (*string, bool) {
	if o == nil || o.Nameserver == nil {
		return nil, false
	}
	return o.Nameserver, true
}

// HasNameserver returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasNameserver() bool {
	if o != nil && o.Nameserver != nil {
		return true
	}

	return false
}

// SetNameserver gets a reference to the given string and assigns it to the Nameserver field.
func (o *OsAnswersAllOf) SetNameserver(v string) {
	o.Nameserver = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetProductKey() string {
	if o == nil || o.ProductKey == nil {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetProductKeyOk() (*string, bool) {
	if o == nil || o.ProductKey == nil {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasProductKey() bool {
	if o != nil && o.ProductKey != nil {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *OsAnswersAllOf) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *OsAnswersAllOf) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OsAnswersAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsAnswersAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OsAnswersAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OsAnswersAllOf) SetSource(v string) {
	o.Source = &v
}

func (o OsAnswersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnswerFile != nil {
		toSerialize["AnswerFile"] = o.AnswerFile
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IpConfigType != nil {
		toSerialize["IpConfigType"] = o.IpConfigType
	}
	if o.IpConfiguration != nil {
		toSerialize["IpConfiguration"] = o.IpConfiguration
	}
	if o.IsAnswerFileSet != nil {
		toSerialize["IsAnswerFileSet"] = o.IsAnswerFileSet
	}
	if o.IsRootPasswordCrypted != nil {
		toSerialize["IsRootPasswordCrypted"] = o.IsRootPasswordCrypted
	}
	if o.IsRootPasswordSet != nil {
		toSerialize["IsRootPasswordSet"] = o.IsRootPasswordSet
	}
	if o.Nameserver != nil {
		toSerialize["Nameserver"] = o.Nameserver
	}
	if o.ProductKey != nil {
		toSerialize["ProductKey"] = o.ProductKey
	}
	if o.RootPassword != nil {
		toSerialize["RootPassword"] = o.RootPassword
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsAnswersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsAnswersAllOf := _OsAnswersAllOf{}

	if err = json.Unmarshal(bytes, &varOsAnswersAllOf); err == nil {
		*o = OsAnswersAllOf(varOsAnswersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AnswerFile")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IpConfigType")
		delete(additionalProperties, "IpConfiguration")
		delete(additionalProperties, "IsAnswerFileSet")
		delete(additionalProperties, "IsRootPasswordCrypted")
		delete(additionalProperties, "IsRootPasswordSet")
		delete(additionalProperties, "Nameserver")
		delete(additionalProperties, "ProductKey")
		delete(additionalProperties, "RootPassword")
		delete(additionalProperties, "Source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsAnswersAllOf struct {
	value *OsAnswersAllOf
	isSet bool
}

func (v NullableOsAnswersAllOf) Get() *OsAnswersAllOf {
	return v.value
}

func (v *NullableOsAnswersAllOf) Set(val *OsAnswersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsAnswersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsAnswersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsAnswersAllOf(val *OsAnswersAllOf) *NullableOsAnswersAllOf {
	return &NullableOsAnswersAllOf{value: val, isSet: true}
}

func (v NullableOsAnswersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsAnswersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
