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

// checks if the DataDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataDestination{}

// DataDestination The properties describing an Ambar Data Destination.
type DataDestination struct {
	// The UTC time at which the DataDestination was created.
	CreatedAt string `json:"createdAt"`
	// The HTTP endpoint which Ambar sends filtered record sequences to.
	DestinationEndpoint string `json:"destinationEndpoint"`
	// The port on which your endpoint service is listening for connections.
	DestinationPort string `json:"destinationPort"`
	// The description of the DataDestination given when it was created.
	Description *string `json:"description,omitempty"`
	// The Ambar resourceId for the Filters which are applied to record sequences sent to the configured DataDestination.
	FilterIds []string `json:"filterIds"`
	// The Ambar resourceId corresponding to this DataDestination.
	ResourceId string `json:"resourceId"`
	// The ResourceState of this DataDestination. For possible values see ResourceState in our developer docs.
	State string `json:"state"`
}

type _DataDestination DataDestination

// NewDataDestination instantiates a new DataDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataDestination(createdAt string, destinationEndpoint string, destinationPort string, filterIds []string, resourceId string, state string) *DataDestination {
	this := DataDestination{}
	this.CreatedAt = createdAt
	this.DestinationEndpoint = destinationEndpoint
	this.DestinationPort = destinationPort
	this.FilterIds = filterIds
	this.ResourceId = resourceId
	this.State = state
	return &this
}

// NewDataDestinationWithDefaults instantiates a new DataDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataDestinationWithDefaults() *DataDestination {
	this := DataDestination{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *DataDestination) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DataDestination) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDestinationEndpoint returns the DestinationEndpoint field value
func (o *DataDestination) GetDestinationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationEndpoint
}

// GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetDestinationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationEndpoint, true
}

// SetDestinationEndpoint sets field value
func (o *DataDestination) SetDestinationEndpoint(v string) {
	o.DestinationEndpoint = v
}

// GetDestinationPort returns the DestinationPort field value
func (o *DataDestination) GetDestinationPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetDestinationPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationPort, true
}

// SetDestinationPort sets field value
func (o *DataDestination) SetDestinationPort(v string) {
	o.DestinationPort = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DataDestination) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DataDestination) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DataDestination) SetDescription(v string) {
	o.Description = &v
}

// GetFilterIds returns the FilterIds field value
func (o *DataDestination) GetFilterIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FilterIds
}

// GetFilterIdsOk returns a tuple with the FilterIds field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetFilterIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterIds, true
}

// SetFilterIds sets field value
func (o *DataDestination) SetFilterIds(v []string) {
	o.FilterIds = v
}

// GetResourceId returns the ResourceId field value
func (o *DataDestination) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *DataDestination) SetResourceId(v string) {
	o.ResourceId = v
}

// GetState returns the State field value
func (o *DataDestination) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DataDestination) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DataDestination) SetState(v string) {
	o.State = v
}

func (o DataDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["destinationEndpoint"] = o.DestinationEndpoint
	toSerialize["destinationPort"] = o.DestinationPort
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["filterIds"] = o.FilterIds
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *DataDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"destinationEndpoint",
		"destinationPort",
		"filterIds",
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

	varDataDestination := _DataDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataDestination)

	if err != nil {
		return err
	}

	*o = DataDestination(varDataDestination)

	return err
}

type NullableDataDestination struct {
	value *DataDestination
	isSet bool
}

func (v NullableDataDestination) Get() *DataDestination {
	return v.value
}

func (v *NullableDataDestination) Set(val *DataDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDestination(val *DataDestination) *NullableDataDestination {
	return &NullableDataDestination{value: val, isSet: true}
}

func (v NullableDataDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


