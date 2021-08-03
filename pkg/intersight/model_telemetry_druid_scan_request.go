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

// TelemetryDruidScanRequest The Scan query returns raw Apache Druid rows in streaming mode. In addition to straightforward usage where a Scan query is issued to the Broker, the Scan query can also be issued directly to Historical processes or streaming ingestion tasks. This can be useful if you want to retrieve large amounts of data in parallel.
type TelemetryDruidScanRequest struct {
	// null
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals []string `json:"intervals"`
	// How the results are represented, list, compactedList or valueVector. Currently only list and compactedList are supported.
	ResultFormat *string               `json:"resultFormat,omitempty"`
	Filter       *TelemetryDruidFilter `json:"filter,omitempty"`
	// A String array of dimensions and metrics to scan. If left empty, all dimensions and metrics are returned.
	Columns *[]string `json:"columns,omitempty"`
	// The maximum number of rows buffered before being returned to the client.
	BatchSize *int32 `json:"batchSize,omitempty"`
	// How many rows to return. If not specified, all rows will be returned.
	Limit *int32 `json:"limit,omitempty"`
	// The ordering of returned rows based on timestamp. \"ascending\", \"descending\", and \"none\" (default) are supported. Currently, \"ascending\" and \"descending\" are only supported for queries where the __time column is included in the columns field and the requirements outlined in the time ordering section are met.
	Order *string `json:"order,omitempty"`
	// Return results consistent with the legacy \"scan-query\" contrib extension. Defaults to the value set by druid.query.scan.legacy, which in turn defaults to false.
	Legacy               *bool                       `json:"legacy,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidScanRequest TelemetryDruidScanRequest

// NewTelemetryDruidScanRequest instantiates a new TelemetryDruidScanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidScanRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string) *TelemetryDruidScanRequest {
	this := TelemetryDruidScanRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Intervals = intervals
	var resultFormat string = "list"
	this.ResultFormat = &resultFormat
	var batchSize int32 = 20480
	this.BatchSize = &batchSize
	var order string = "none"
	this.Order = &order
	var legacy bool = false
	this.Legacy = &legacy
	return &this
}

// NewTelemetryDruidScanRequestWithDefaults instantiates a new TelemetryDruidScanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidScanRequestWithDefaults() *TelemetryDruidScanRequest {
	this := TelemetryDruidScanRequest{}
	var resultFormat string = "list"
	this.ResultFormat = &resultFormat
	var batchSize int32 = 20480
	this.BatchSize = &batchSize
	var order string = "none"
	this.Order = &order
	var legacy bool = false
	this.Legacy = &legacy
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidScanRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidScanRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidScanRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidScanRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidScanRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidScanRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetResultFormat returns the ResultFormat field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetResultFormat() string {
	if o == nil || o.ResultFormat == nil {
		var ret string
		return ret
	}
	return *o.ResultFormat
}

// GetResultFormatOk returns a tuple with the ResultFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetResultFormatOk() (*string, bool) {
	if o == nil || o.ResultFormat == nil {
		return nil, false
	}
	return o.ResultFormat, true
}

// HasResultFormat returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasResultFormat() bool {
	if o != nil && o.ResultFormat != nil {
		return true
	}

	return false
}

// SetResultFormat gets a reference to the given string and assigns it to the ResultFormat field.
func (o *TelemetryDruidScanRequest) SetResultFormat(v string) {
	o.ResultFormat = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidScanRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *TelemetryDruidScanRequest) SetColumns(v []string) {
	o.Columns = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetBatchSize() int32 {
	if o == nil || o.BatchSize == nil {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetBatchSizeOk() (*int32, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasBatchSize() bool {
	if o != nil && o.BatchSize != nil {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *TelemetryDruidScanRequest) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TelemetryDruidScanRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *TelemetryDruidScanRequest) SetOrder(v string) {
	o.Order = &v
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetLegacy() bool {
	if o == nil || o.Legacy == nil {
		var ret bool
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetLegacyOk() (*bool, bool) {
	if o == nil || o.Legacy == nil {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasLegacy() bool {
	if o != nil && o.Legacy != nil {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given bool and assigns it to the Legacy field.
func (o *TelemetryDruidScanRequest) SetLegacy(v bool) {
	o.Legacy = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidScanRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidScanRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidScanRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidScanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if o.ResultFormat != nil {
		toSerialize["resultFormat"] = o.ResultFormat
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.BatchSize != nil {
		toSerialize["batchSize"] = o.BatchSize
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Legacy != nil {
		toSerialize["legacy"] = o.Legacy
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidScanRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidScanRequest := _TelemetryDruidScanRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidScanRequest); err == nil {
		*o = TelemetryDruidScanRequest(varTelemetryDruidScanRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "resultFormat")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "batchSize")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "order")
		delete(additionalProperties, "legacy")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidScanRequest struct {
	value *TelemetryDruidScanRequest
	isSet bool
}

func (v NullableTelemetryDruidScanRequest) Get() *TelemetryDruidScanRequest {
	return v.value
}

func (v *NullableTelemetryDruidScanRequest) Set(val *TelemetryDruidScanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidScanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidScanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidScanRequest(val *TelemetryDruidScanRequest) *NullableTelemetryDruidScanRequest {
	return &NullableTelemetryDruidScanRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidScanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidScanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}