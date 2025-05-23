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

// checks if the ConnectorActiveTopicsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorActiveTopicsValue{}

// ConnectorActiveTopicsValue struct for ConnectorActiveTopicsValue
type ConnectorActiveTopicsValue struct {
	Topics []string `json:"topics,omitempty"`
}

// NewConnectorActiveTopicsValue instantiates a new ConnectorActiveTopicsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorActiveTopicsValue() *ConnectorActiveTopicsValue {
	this := ConnectorActiveTopicsValue{}
	return &this
}

// NewConnectorActiveTopicsValueWithDefaults instantiates a new ConnectorActiveTopicsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorActiveTopicsValueWithDefaults() *ConnectorActiveTopicsValue {
	this := ConnectorActiveTopicsValue{}
	return &this
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *ConnectorActiveTopicsValue) GetTopics() []string {
	if o == nil || IsNil(o.Topics) {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorActiveTopicsValue) GetTopicsOk() ([]string, bool) {
	if o == nil || IsNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *ConnectorActiveTopicsValue) HasTopics() bool {
	if o != nil && !IsNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *ConnectorActiveTopicsValue) SetTopics(v []string) {
	o.Topics = v
}

func (o ConnectorActiveTopicsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorActiveTopicsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topics) {
		toSerialize["topics"] = o.Topics
	}
	return toSerialize, nil
}

type NullableConnectorActiveTopicsValue struct {
	value *ConnectorActiveTopicsValue
	isSet bool
}

func (v NullableConnectorActiveTopicsValue) Get() *ConnectorActiveTopicsValue {
	return v.value
}

func (v *NullableConnectorActiveTopicsValue) Set(val *ConnectorActiveTopicsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorActiveTopicsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorActiveTopicsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorActiveTopicsValue(val *ConnectorActiveTopicsValue) *NullableConnectorActiveTopicsValue {
	return &NullableConnectorActiveTopicsValue{value: val, isSet: true}
}

func (v NullableConnectorActiveTopicsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorActiveTopicsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


