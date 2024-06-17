/*
Ambar OpenAPI Specification

Details about communicating with Ambar.cloud public endpoints. Supported HTTP rest endpoints and their  request and response details.

API version: 2024-06-11
Contact: contact@ambar.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ambar

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter The properties describing an Ambar Filter.
type Filter struct {
	// The UTC time at which the Filter was created.
	CreatedAt string `json:"createdAt"`
	// The list of DataDestination ResourceIds which use this Filter.
	DataDestinationsUsingFilter []string `json:"dataDestinationsUsingFilter"`
	// The DataSource ResourceId which this Filter will read from.
	DataSourceId string `json:"dataSourceId"`
	// The Ambar resourceId corresponding to this Filter.
	ResourceId string `json:"resourceId"`
	// The description of the Filter given when it was created.
	Description *string `json:"description,omitempty"`
	// The ResourceState of this Filter. For possible values see ResourceState in our developer docs.
	State string `json:"state"`
}

type _Filter Filter

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter(createdAt string, dataDestinationsUsingFilter []string, dataSourceId string, resourceId string, state string) *Filter {
	this := Filter{}
	this.CreatedAt = createdAt
	this.DataDestinationsUsingFilter = dataDestinationsUsingFilter
	this.DataSourceId = dataSourceId
	this.ResourceId = resourceId
	this.State = state
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Filter) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Filter) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Filter) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDataDestinationsUsingFilter returns the DataDestinationsUsingFilter field value
func (o *Filter) GetDataDestinationsUsingFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DataDestinationsUsingFilter
}

// GetDataDestinationsUsingFilterOk returns a tuple with the DataDestinationsUsingFilter field value
// and a boolean to check if the value has been set.
func (o *Filter) GetDataDestinationsUsingFilterOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataDestinationsUsingFilter, true
}

// SetDataDestinationsUsingFilter sets field value
func (o *Filter) SetDataDestinationsUsingFilter(v []string) {
	o.DataDestinationsUsingFilter = v
}

// GetDataSourceId returns the DataSourceId field value
func (o *Filter) GetDataSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSourceId
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value
// and a boolean to check if the value has been set.
func (o *Filter) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceId, true
}

// SetDataSourceId sets field value
func (o *Filter) SetDataSourceId(v string) {
	o.DataSourceId = v
}

// GetResourceId returns the ResourceId field value
func (o *Filter) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Filter) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Filter) SetResourceId(v string) {
	o.ResourceId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Filter) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Filter) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Filter) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value
func (o *Filter) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Filter) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Filter) SetState(v string) {
	o.State = v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["dataDestinationsUsingFilter"] = o.DataDestinationsUsingFilter
	toSerialize["dataSourceId"] = o.DataSourceId
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *Filter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"dataDestinationsUsingFilter",
		"dataSourceId",
		"resourceId",
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFilter := _Filter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilter)

	if err != nil {
		return err
	}

	*o = Filter(varFilter)

	return err
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


