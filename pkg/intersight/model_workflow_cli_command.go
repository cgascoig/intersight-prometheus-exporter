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

// WorkflowCliCommand This models a single CLI command that can be executed on the end point.
type WorkflowCliCommand struct {
	WorkflowApi
	// The command to run on the device connector.
	Command *string `json:"Command,omitempty"`
	// The regex string that identifies the end of the command response.
	EndPrompt     *string                 `json:"EndPrompt,omitempty"`
	ExpectPrompts *[]WorkflowExpectPrompt `json:"ExpectPrompts,omitempty"`
	// Skips the execution status check of the terminal command. One use case for this is while exiting the terminal session from esxi host.
	SkipStatusCheck *bool `json:"SkipStatusCheck,omitempty"`
	// If this flag is set, it marks the end of the terminal session where the previous commands were executed.
	TerminalEnd *bool `json:"TerminalEnd,omitempty"`
	// If this flag is set, the execution of this command initiates a terminal session in which the subsequent CLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch.
	TerminalStart *bool `json:"TerminalStart,omitempty"`
	// The type of the command - can be interactive or non-interactive. * `NonInteractive` - The CLI command is not an interactive command. * `Interactive` - The CLI command is executed in interactive mode and the command must provide the expects andanswers.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCliCommand WorkflowCliCommand

// NewWorkflowCliCommand instantiates a new WorkflowCliCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCliCommand() *WorkflowCliCommand {
	this := WorkflowCliCommand{}
	var type_ string = "NonInteractive"
	this.Type = &type_
	return &this
}

// NewWorkflowCliCommandWithDefaults instantiates a new WorkflowCliCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCliCommandWithDefaults() *WorkflowCliCommand {
	this := WorkflowCliCommand{}
	var type_ string = "NonInteractive"
	this.Type = &type_
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *WorkflowCliCommand) SetCommand(v string) {
	o.Command = &v
}

// GetEndPrompt returns the EndPrompt field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetEndPrompt() string {
	if o == nil || o.EndPrompt == nil {
		var ret string
		return ret
	}
	return *o.EndPrompt
}

// GetEndPromptOk returns a tuple with the EndPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetEndPromptOk() (*string, bool) {
	if o == nil || o.EndPrompt == nil {
		return nil, false
	}
	return o.EndPrompt, true
}

// HasEndPrompt returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasEndPrompt() bool {
	if o != nil && o.EndPrompt != nil {
		return true
	}

	return false
}

// SetEndPrompt gets a reference to the given string and assigns it to the EndPrompt field.
func (o *WorkflowCliCommand) SetEndPrompt(v string) {
	o.EndPrompt = &v
}

// GetExpectPrompts returns the ExpectPrompts field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetExpectPrompts() []WorkflowExpectPrompt {
	if o == nil || o.ExpectPrompts == nil {
		var ret []WorkflowExpectPrompt
		return ret
	}
	return *o.ExpectPrompts
}

// GetExpectPromptsOk returns a tuple with the ExpectPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetExpectPromptsOk() (*[]WorkflowExpectPrompt, bool) {
	if o == nil || o.ExpectPrompts == nil {
		return nil, false
	}
	return o.ExpectPrompts, true
}

// HasExpectPrompts returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasExpectPrompts() bool {
	if o != nil && o.ExpectPrompts != nil {
		return true
	}

	return false
}

// SetExpectPrompts gets a reference to the given []WorkflowExpectPrompt and assigns it to the ExpectPrompts field.
func (o *WorkflowCliCommand) SetExpectPrompts(v []WorkflowExpectPrompt) {
	o.ExpectPrompts = &v
}

// GetSkipStatusCheck returns the SkipStatusCheck field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetSkipStatusCheck() bool {
	if o == nil || o.SkipStatusCheck == nil {
		var ret bool
		return ret
	}
	return *o.SkipStatusCheck
}

// GetSkipStatusCheckOk returns a tuple with the SkipStatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetSkipStatusCheckOk() (*bool, bool) {
	if o == nil || o.SkipStatusCheck == nil {
		return nil, false
	}
	return o.SkipStatusCheck, true
}

// HasSkipStatusCheck returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasSkipStatusCheck() bool {
	if o != nil && o.SkipStatusCheck != nil {
		return true
	}

	return false
}

// SetSkipStatusCheck gets a reference to the given bool and assigns it to the SkipStatusCheck field.
func (o *WorkflowCliCommand) SetSkipStatusCheck(v bool) {
	o.SkipStatusCheck = &v
}

// GetTerminalEnd returns the TerminalEnd field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetTerminalEnd() bool {
	if o == nil || o.TerminalEnd == nil {
		var ret bool
		return ret
	}
	return *o.TerminalEnd
}

// GetTerminalEndOk returns a tuple with the TerminalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetTerminalEndOk() (*bool, bool) {
	if o == nil || o.TerminalEnd == nil {
		return nil, false
	}
	return o.TerminalEnd, true
}

// HasTerminalEnd returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasTerminalEnd() bool {
	if o != nil && o.TerminalEnd != nil {
		return true
	}

	return false
}

// SetTerminalEnd gets a reference to the given bool and assigns it to the TerminalEnd field.
func (o *WorkflowCliCommand) SetTerminalEnd(v bool) {
	o.TerminalEnd = &v
}

// GetTerminalStart returns the TerminalStart field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetTerminalStart() bool {
	if o == nil || o.TerminalStart == nil {
		var ret bool
		return ret
	}
	return *o.TerminalStart
}

// GetTerminalStartOk returns a tuple with the TerminalStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetTerminalStartOk() (*bool, bool) {
	if o == nil || o.TerminalStart == nil {
		return nil, false
	}
	return o.TerminalStart, true
}

// HasTerminalStart returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasTerminalStart() bool {
	if o != nil && o.TerminalStart != nil {
		return true
	}

	return false
}

// SetTerminalStart gets a reference to the given bool and assigns it to the TerminalStart field.
func (o *WorkflowCliCommand) SetTerminalStart(v bool) {
	o.TerminalStart = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowCliCommand) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCliCommand) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowCliCommand) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowCliCommand) SetType(v string) {
	o.Type = &v
}

func (o WorkflowCliCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if o.Command != nil {
		toSerialize["Command"] = o.Command
	}
	if o.EndPrompt != nil {
		toSerialize["EndPrompt"] = o.EndPrompt
	}
	if o.ExpectPrompts != nil {
		toSerialize["ExpectPrompts"] = o.ExpectPrompts
	}
	if o.SkipStatusCheck != nil {
		toSerialize["SkipStatusCheck"] = o.SkipStatusCheck
	}
	if o.TerminalEnd != nil {
		toSerialize["TerminalEnd"] = o.TerminalEnd
	}
	if o.TerminalStart != nil {
		toSerialize["TerminalStart"] = o.TerminalStart
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCliCommand) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCliCommandWithoutEmbeddedStruct struct {
		// The command to run on the device connector.
		Command *string `json:"Command,omitempty"`
		// The regex string that identifies the end of the command response.
		EndPrompt     *string                 `json:"EndPrompt,omitempty"`
		ExpectPrompts *[]WorkflowExpectPrompt `json:"ExpectPrompts,omitempty"`
		// Skips the execution status check of the terminal command. One use case for this is while exiting the terminal session from esxi host.
		SkipStatusCheck *bool `json:"SkipStatusCheck,omitempty"`
		// If this flag is set, it marks the end of the terminal session where the previous commands were executed.
		TerminalEnd *bool `json:"TerminalEnd,omitempty"`
		// If this flag is set, the execution of this command initiates a terminal session in which the subsequent CLI commands are executed until a command with terminalEnd flag is encountered or the end of the batch.
		TerminalStart *bool `json:"TerminalStart,omitempty"`
		// The type of the command - can be interactive or non-interactive. * `NonInteractive` - The CLI command is not an interactive command. * `Interactive` - The CLI command is executed in interactive mode and the command must provide the expects andanswers.
		Type *string `json:"Type,omitempty"`
	}

	varWorkflowCliCommandWithoutEmbeddedStruct := WorkflowCliCommandWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCliCommandWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCliCommand := _WorkflowCliCommand{}
		varWorkflowCliCommand.Command = varWorkflowCliCommandWithoutEmbeddedStruct.Command
		varWorkflowCliCommand.EndPrompt = varWorkflowCliCommandWithoutEmbeddedStruct.EndPrompt
		varWorkflowCliCommand.ExpectPrompts = varWorkflowCliCommandWithoutEmbeddedStruct.ExpectPrompts
		varWorkflowCliCommand.SkipStatusCheck = varWorkflowCliCommandWithoutEmbeddedStruct.SkipStatusCheck
		varWorkflowCliCommand.TerminalEnd = varWorkflowCliCommandWithoutEmbeddedStruct.TerminalEnd
		varWorkflowCliCommand.TerminalStart = varWorkflowCliCommandWithoutEmbeddedStruct.TerminalStart
		varWorkflowCliCommand.Type = varWorkflowCliCommandWithoutEmbeddedStruct.Type
		*o = WorkflowCliCommand(varWorkflowCliCommand)
	} else {
		return err
	}

	varWorkflowCliCommand := _WorkflowCliCommand{}

	err = json.Unmarshal(bytes, &varWorkflowCliCommand)
	if err == nil {
		o.WorkflowApi = varWorkflowCliCommand.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Command")
		delete(additionalProperties, "EndPrompt")
		delete(additionalProperties, "ExpectPrompts")
		delete(additionalProperties, "SkipStatusCheck")
		delete(additionalProperties, "TerminalEnd")
		delete(additionalProperties, "TerminalStart")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowCliCommand struct {
	value *WorkflowCliCommand
	isSet bool
}

func (v NullableWorkflowCliCommand) Get() *WorkflowCliCommand {
	return v.value
}

func (v *NullableWorkflowCliCommand) Set(val *WorkflowCliCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCliCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCliCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCliCommand(val *WorkflowCliCommand) *NullableWorkflowCliCommand {
	return &NullableWorkflowCliCommand{value: val, isSet: true}
}

func (v NullableWorkflowCliCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCliCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
