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

// ContentTextParameter Concrete implementation of BaseParameter for Text content.
type ContentTextParameter struct {
	ContentBaseParameter
	// Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data.
	IsDelimiter *bool `json:"IsDelimiter,omitempty"`
	// Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match.
	IsNextCaptureOnSameLine *bool `json:"IsNextCaptureOnSameLine,omitempty"`
	// Regular expression of the line containing the data to be extracted from text content.
	RegexLine            *string `json:"RegexLine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentTextParameter ContentTextParameter

// NewContentTextParameter instantiates a new ContentTextParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentTextParameter() *ContentTextParameter {
	this := ContentTextParameter{}
	return &this
}

// NewContentTextParameterWithDefaults instantiates a new ContentTextParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTextParameterWithDefaults() *ContentTextParameter {
	this := ContentTextParameter{}
	return &this
}

// GetIsDelimiter returns the IsDelimiter field value if set, zero value otherwise.
func (o *ContentTextParameter) GetIsDelimiter() bool {
	if o == nil || o.IsDelimiter == nil {
		var ret bool
		return ret
	}
	return *o.IsDelimiter
}

// GetIsDelimiterOk returns a tuple with the IsDelimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetIsDelimiterOk() (*bool, bool) {
	if o == nil || o.IsDelimiter == nil {
		return nil, false
	}
	return o.IsDelimiter, true
}

// HasIsDelimiter returns a boolean if a field has been set.
func (o *ContentTextParameter) HasIsDelimiter() bool {
	if o != nil && o.IsDelimiter != nil {
		return true
	}

	return false
}

// SetIsDelimiter gets a reference to the given bool and assigns it to the IsDelimiter field.
func (o *ContentTextParameter) SetIsDelimiter(v bool) {
	o.IsDelimiter = &v
}

// GetIsNextCaptureOnSameLine returns the IsNextCaptureOnSameLine field value if set, zero value otherwise.
func (o *ContentTextParameter) GetIsNextCaptureOnSameLine() bool {
	if o == nil || o.IsNextCaptureOnSameLine == nil {
		var ret bool
		return ret
	}
	return *o.IsNextCaptureOnSameLine
}

// GetIsNextCaptureOnSameLineOk returns a tuple with the IsNextCaptureOnSameLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetIsNextCaptureOnSameLineOk() (*bool, bool) {
	if o == nil || o.IsNextCaptureOnSameLine == nil {
		return nil, false
	}
	return o.IsNextCaptureOnSameLine, true
}

// HasIsNextCaptureOnSameLine returns a boolean if a field has been set.
func (o *ContentTextParameter) HasIsNextCaptureOnSameLine() bool {
	if o != nil && o.IsNextCaptureOnSameLine != nil {
		return true
	}

	return false
}

// SetIsNextCaptureOnSameLine gets a reference to the given bool and assigns it to the IsNextCaptureOnSameLine field.
func (o *ContentTextParameter) SetIsNextCaptureOnSameLine(v bool) {
	o.IsNextCaptureOnSameLine = &v
}

// GetRegexLine returns the RegexLine field value if set, zero value otherwise.
func (o *ContentTextParameter) GetRegexLine() string {
	if o == nil || o.RegexLine == nil {
		var ret string
		return ret
	}
	return *o.RegexLine
}

// GetRegexLineOk returns a tuple with the RegexLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTextParameter) GetRegexLineOk() (*string, bool) {
	if o == nil || o.RegexLine == nil {
		return nil, false
	}
	return o.RegexLine, true
}

// HasRegexLine returns a boolean if a field has been set.
func (o *ContentTextParameter) HasRegexLine() bool {
	if o != nil && o.RegexLine != nil {
		return true
	}

	return false
}

// SetRegexLine gets a reference to the given string and assigns it to the RegexLine field.
func (o *ContentTextParameter) SetRegexLine(v string) {
	o.RegexLine = &v
}

func (o ContentTextParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedContentBaseParameter, errContentBaseParameter := json.Marshal(o.ContentBaseParameter)
	if errContentBaseParameter != nil {
		return []byte{}, errContentBaseParameter
	}
	errContentBaseParameter = json.Unmarshal([]byte(serializedContentBaseParameter), &toSerialize)
	if errContentBaseParameter != nil {
		return []byte{}, errContentBaseParameter
	}
	if o.IsDelimiter != nil {
		toSerialize["IsDelimiter"] = o.IsDelimiter
	}
	if o.IsNextCaptureOnSameLine != nil {
		toSerialize["IsNextCaptureOnSameLine"] = o.IsNextCaptureOnSameLine
	}
	if o.RegexLine != nil {
		toSerialize["RegexLine"] = o.RegexLine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentTextParameter) UnmarshalJSON(bytes []byte) (err error) {
	type ContentTextParameterWithoutEmbeddedStruct struct {
		// Data to be extracted from text content can be simple type or complex type or collection of simple/complex types. Complex types are group of simple or complex type. Delimiter is required to stop parsing list and complex data types. isDelimiter specifies whether given TextParameter is a delimiter or regular rule to capture the text data.
		IsDelimiter *bool `json:"IsDelimiter,omitempty"`
		// Set to true of the next value to capture resides on the same text line of current match. By default textFSM engine gets the next text line on finding the first match.
		IsNextCaptureOnSameLine *bool `json:"IsNextCaptureOnSameLine,omitempty"`
		// Regular expression of the line containing the data to be extracted from text content.
		RegexLine *string `json:"RegexLine,omitempty"`
	}

	varContentTextParameterWithoutEmbeddedStruct := ContentTextParameterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varContentTextParameterWithoutEmbeddedStruct)
	if err == nil {
		varContentTextParameter := _ContentTextParameter{}
		varContentTextParameter.IsDelimiter = varContentTextParameterWithoutEmbeddedStruct.IsDelimiter
		varContentTextParameter.IsNextCaptureOnSameLine = varContentTextParameterWithoutEmbeddedStruct.IsNextCaptureOnSameLine
		varContentTextParameter.RegexLine = varContentTextParameterWithoutEmbeddedStruct.RegexLine
		*o = ContentTextParameter(varContentTextParameter)
	} else {
		return err
	}

	varContentTextParameter := _ContentTextParameter{}

	err = json.Unmarshal(bytes, &varContentTextParameter)
	if err == nil {
		o.ContentBaseParameter = varContentTextParameter.ContentBaseParameter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IsDelimiter")
		delete(additionalProperties, "IsNextCaptureOnSameLine")
		delete(additionalProperties, "RegexLine")

		// remove fields from embedded structs
		reflectContentBaseParameter := reflect.ValueOf(o.ContentBaseParameter)
		for i := 0; i < reflectContentBaseParameter.Type().NumField(); i++ {
			t := reflectContentBaseParameter.Type().Field(i)

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

type NullableContentTextParameter struct {
	value *ContentTextParameter
	isSet bool
}

func (v NullableContentTextParameter) Get() *ContentTextParameter {
	return v.value
}

func (v *NullableContentTextParameter) Set(val *ContentTextParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContentTextParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContentTextParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentTextParameter(val *ContentTextParameter) *NullableContentTextParameter {
	return &NullableContentTextParameter{value: val, isSet: true}
}

func (v NullableContentTextParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentTextParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
