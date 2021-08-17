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
	"fmt"
)

// TelemetryDruidAggregateSearchSpec - Generic search specification descriptor
type TelemetryDruidAggregateSearchSpec struct {
	TelemetryDruidContainsSearchSpec            *TelemetryDruidContainsSearchSpec
	TelemetryDruidFragmentSearchSpec            *TelemetryDruidFragmentSearchSpec
	TelemetryDruidInsensitiveContainsSearchSpec *TelemetryDruidInsensitiveContainsSearchSpec
	TelemetryDruidRegexSearchSpec               *TelemetryDruidRegexSearchSpec
}

// TelemetryDruidContainsSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidContainsSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidContainsSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidContainsSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidContainsSearchSpec: v}
}

// TelemetryDruidFragmentSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidFragmentSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidFragmentSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidFragmentSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidFragmentSearchSpec: v}
}

// TelemetryDruidInsensitiveContainsSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidInsensitiveContainsSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidInsensitiveContainsSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidInsensitiveContainsSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidInsensitiveContainsSearchSpec: v}
}

// TelemetryDruidRegexSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidRegexSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidRegexSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidRegexSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidRegexSearchSpec: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidAggregateSearchSpec) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'contains'
	if jsonDict["type"] == "contains" {
		// try to unmarshal JSON data into TelemetryDruidContainsSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidContainsSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidContainsSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidContainsSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidContainsSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'fragment'
	if jsonDict["type"] == "fragment" {
		// try to unmarshal JSON data into TelemetryDruidFragmentSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidFragmentSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFragmentSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidFragmentSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidFragmentSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'insensitive_contains'
	if jsonDict["type"] == "insensitive_contains" {
		// try to unmarshal JSON data into TelemetryDruidInsensitiveContainsSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidInsensitiveContainsSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInsensitiveContainsSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidInsensitiveContainsSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidInsensitiveContainsSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'regex'
	if jsonDict["type"] == "regex" {
		// try to unmarshal JSON data into TelemetryDruidRegexSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidRegexSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidRegexSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidContainsSearchSpec'
	if jsonDict["type"] == "telemetry.DruidContainsSearchSpec" {
		// try to unmarshal JSON data into TelemetryDruidContainsSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidContainsSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidContainsSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidContainsSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidContainsSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidFragmentSearchSpec'
	if jsonDict["type"] == "telemetry.DruidFragmentSearchSpec" {
		// try to unmarshal JSON data into TelemetryDruidFragmentSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidFragmentSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFragmentSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidFragmentSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidFragmentSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidInsensitiveContainsSearchSpec'
	if jsonDict["type"] == "telemetry.DruidInsensitiveContainsSearchSpec" {
		// try to unmarshal JSON data into TelemetryDruidInsensitiveContainsSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidInsensitiveContainsSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInsensitiveContainsSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidInsensitiveContainsSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidInsensitiveContainsSearchSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidRegexSearchSpec'
	if jsonDict["type"] == "telemetry.DruidRegexSearchSpec" {
		// try to unmarshal JSON data into TelemetryDruidRegexSearchSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexSearchSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexSearchSpec, return on the first match
		} else {
			dst.TelemetryDruidRegexSearchSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidAggregateSearchSpec as TelemetryDruidRegexSearchSpec: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidAggregateSearchSpec) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidContainsSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidContainsSearchSpec)
	}

	if src.TelemetryDruidFragmentSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidFragmentSearchSpec)
	}

	if src.TelemetryDruidInsensitiveContainsSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidInsensitiveContainsSearchSpec)
	}

	if src.TelemetryDruidRegexSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidRegexSearchSpec)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidAggregateSearchSpec) GetActualInstance() interface{} {
	if obj.TelemetryDruidContainsSearchSpec != nil {
		return obj.TelemetryDruidContainsSearchSpec
	}

	if obj.TelemetryDruidFragmentSearchSpec != nil {
		return obj.TelemetryDruidFragmentSearchSpec
	}

	if obj.TelemetryDruidInsensitiveContainsSearchSpec != nil {
		return obj.TelemetryDruidInsensitiveContainsSearchSpec
	}

	if obj.TelemetryDruidRegexSearchSpec != nil {
		return obj.TelemetryDruidRegexSearchSpec
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidAggregateSearchSpec struct {
	value *TelemetryDruidAggregateSearchSpec
	isSet bool
}

func (v NullableTelemetryDruidAggregateSearchSpec) Get() *TelemetryDruidAggregateSearchSpec {
	return v.value
}

func (v *NullableTelemetryDruidAggregateSearchSpec) Set(val *TelemetryDruidAggregateSearchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidAggregateSearchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidAggregateSearchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidAggregateSearchSpec(val *TelemetryDruidAggregateSearchSpec) *NullableTelemetryDruidAggregateSearchSpec {
	return &NullableTelemetryDruidAggregateSearchSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidAggregateSearchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidAggregateSearchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}