/*
Kafka Connect REST API

This is the documentation of the [Apache Kafka](https://kafka.apache.org) Connect REST API.

API version: 3.9.0
Contact: dev@kafka.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConnectorTaskId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorTaskId{}

// ConnectorTaskId struct for ConnectorTaskId
type ConnectorTaskId struct {
	Connector *string `json:"connector,omitempty"`
	Task *int32 `json:"task,omitempty"`
}

// NewConnectorTaskId instantiates a new ConnectorTaskId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTaskId() *ConnectorTaskId {
	this := ConnectorTaskId{}
	return &this
}

// NewConnectorTaskIdWithDefaults instantiates a new ConnectorTaskId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTaskIdWithDefaults() *ConnectorTaskId {
	this := ConnectorTaskId{}
	return &this
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *ConnectorTaskId) GetConnector() string {
	if o == nil || IsNil(o.Connector) {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTaskId) GetConnectorOk() (*string, bool) {
	if o == nil || IsNil(o.Connector) {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *ConnectorTaskId) HasConnector() bool {
	if o != nil && !IsNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *ConnectorTaskId) SetConnector(v string) {
	o.Connector = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *ConnectorTaskId) GetTask() int32 {
	if o == nil || IsNil(o.Task) {
		var ret int32
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTaskId) GetTaskOk() (*int32, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *ConnectorTaskId) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given int32 and assigns it to the Task field.
func (o *ConnectorTaskId) SetTask(v int32) {
	o.Task = &v
}

func (o ConnectorTaskId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorTaskId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	return toSerialize, nil
}

type NullableConnectorTaskId struct {
	value *ConnectorTaskId
	isSet bool
}

func (v NullableConnectorTaskId) Get() *ConnectorTaskId {
	return v.value
}

func (v *NullableConnectorTaskId) Set(val *ConnectorTaskId) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTaskId) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTaskId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTaskId(val *ConnectorTaskId) *NullableConnectorTaskId {
	return &NullableConnectorTaskId{value: val, isSet: true}
}

func (v NullableConnectorTaskId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTaskId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


