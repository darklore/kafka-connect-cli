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

// checks if the ConfigKeyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigKeyInfo{}

// ConfigKeyInfo struct for ConfigKeyInfo
type ConfigKeyInfo struct {
	DefaultValue *string `json:"default_value,omitempty"`
	Dependents []string `json:"dependents,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Group *string `json:"group,omitempty"`
	Importance *string `json:"importance,omitempty"`
	Name *string `json:"name,omitempty"`
	Order *int32 `json:"order,omitempty"`
	OrderInGroup *int32 `json:"order_in_group,omitempty"`
	Required *bool `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
	Width *string `json:"width,omitempty"`
}

// NewConfigKeyInfo instantiates a new ConfigKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigKeyInfo() *ConfigKeyInfo {
	this := ConfigKeyInfo{}
	return &this
}

// NewConfigKeyInfoWithDefaults instantiates a new ConfigKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigKeyInfoWithDefaults() *ConfigKeyInfo {
	this := ConfigKeyInfo{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *ConfigKeyInfo) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDependents returns the Dependents field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetDependents() []string {
	if o == nil || IsNil(o.Dependents) {
		var ret []string
		return ret
	}
	return o.Dependents
}

// GetDependentsOk returns a tuple with the Dependents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetDependentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dependents) {
		return nil, false
	}
	return o.Dependents, true
}

// HasDependents returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasDependents() bool {
	if o != nil && !IsNil(o.Dependents) {
		return true
	}

	return false
}

// SetDependents gets a reference to the given []string and assigns it to the Dependents field.
func (o *ConfigKeyInfo) SetDependents(v []string) {
	o.Dependents = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConfigKeyInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *ConfigKeyInfo) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ConfigKeyInfo) SetGroup(v string) {
	o.Group = &v
}

// GetImportance returns the Importance field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetImportance() string {
	if o == nil || IsNil(o.Importance) {
		var ret string
		return ret
	}
	return *o.Importance
}

// GetImportanceOk returns a tuple with the Importance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetImportanceOk() (*string, bool) {
	if o == nil || IsNil(o.Importance) {
		return nil, false
	}
	return o.Importance, true
}

// HasImportance returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasImportance() bool {
	if o != nil && !IsNil(o.Importance) {
		return true
	}

	return false
}

// SetImportance gets a reference to the given string and assigns it to the Importance field.
func (o *ConfigKeyInfo) SetImportance(v string) {
	o.Importance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigKeyInfo) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ConfigKeyInfo) SetOrder(v int32) {
	o.Order = &v
}

// GetOrderInGroup returns the OrderInGroup field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetOrderInGroup() int32 {
	if o == nil || IsNil(o.OrderInGroup) {
		var ret int32
		return ret
	}
	return *o.OrderInGroup
}

// GetOrderInGroupOk returns a tuple with the OrderInGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetOrderInGroupOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderInGroup) {
		return nil, false
	}
	return o.OrderInGroup, true
}

// HasOrderInGroup returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasOrderInGroup() bool {
	if o != nil && !IsNil(o.OrderInGroup) {
		return true
	}

	return false
}

// SetOrderInGroup gets a reference to the given int32 and assigns it to the OrderInGroup field.
func (o *ConfigKeyInfo) SetOrderInGroup(v int32) {
	o.OrderInGroup = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ConfigKeyInfo) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConfigKeyInfo) SetType(v string) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ConfigKeyInfo) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigKeyInfo) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ConfigKeyInfo) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *ConfigKeyInfo) SetWidth(v string) {
	o.Width = &v
}

func (o ConfigKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigKeyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.Dependents) {
		toSerialize["dependents"] = o.Dependents
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Importance) {
		toSerialize["importance"] = o.Importance
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.OrderInGroup) {
		toSerialize["order_in_group"] = o.OrderInGroup
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableConfigKeyInfo struct {
	value *ConfigKeyInfo
	isSet bool
}

func (v NullableConfigKeyInfo) Get() *ConfigKeyInfo {
	return v.value
}

func (v *NullableConfigKeyInfo) Set(val *ConfigKeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigKeyInfo(val *ConfigKeyInfo) *NullableConfigKeyInfo {
	return &NullableConfigKeyInfo{value: val, isSet: true}
}

func (v NullableConfigKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

