/*
Ambar OpenAPI Specification

Details about communicating with Ambar.cloud public endpoints. Supported HTTP rest endpoints and their  request and response details.

API version: 2023-12-01
Contact: contact@ambar.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ambar

import (
	"encoding/json"
)

// checks if the DataDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataDestination{}

// DataDestination The properties describing an Ambar Data Destination.
type DataDestination struct {
	// The UTC time at which the DataDestination was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The HTTP endpoint which Ambar sends filtered record sequences to.
	DestinationEndpoint *string `json:"destinationEndpoint,omitempty"`
	// The name given to this DataDestinations projection at creation.
	ProjectionName *string `json:"projectionName,omitempty"`
	// The Ambar resourceId for the Filter which is applied to record sequences sent to this Data Destination.
	FilterId *string `json:"filterId,omitempty"`
	// The Ambar resourceId corresponding to this DataDestination.
	ResourceId *string `json:"resourceId,omitempty"`
	// The ResourceState of this DataDestination. For possible values see ResourceState in our developer docs.
	State *string `json:"state,omitempty"`
}

// NewDataDestination instantiates a new DataDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataDestination() *DataDestination {
	this := DataDestination{}
	return &this
}

// NewDataDestinationWithDefaults instantiates a new DataDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataDestinationWithDefaults() *DataDestination {
	this := DataDestination{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataDestination) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataDestination) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DataDestination) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDestinationEndpoint returns the DestinationEndpoint field value if set, zero value otherwise.
func (o *DataDestination) GetDestinationEndpoint() string {
	if o == nil || IsNil(o.DestinationEndpoint) {
		var ret string
		return ret
	}
	return *o.DestinationEndpoint
}

// GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetDestinationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationEndpoint) {
		return nil, false
	}
	return o.DestinationEndpoint, true
}

// HasDestinationEndpoint returns a boolean if a field has been set.
func (o *DataDestination) HasDestinationEndpoint() bool {
	if o != nil && !IsNil(o.DestinationEndpoint) {
		return true
	}

	return false
}

// SetDestinationEndpoint gets a reference to the given string and assigns it to the DestinationEndpoint field.
func (o *DataDestination) SetDestinationEndpoint(v string) {
	o.DestinationEndpoint = &v
}

// GetProjectionName returns the ProjectionName field value if set, zero value otherwise.
func (o *DataDestination) GetProjectionName() string {
	if o == nil || IsNil(o.ProjectionName) {
		var ret string
		return ret
	}
	return *o.ProjectionName
}

// GetProjectionNameOk returns a tuple with the ProjectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetProjectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectionName) {
		return nil, false
	}
	return o.ProjectionName, true
}

// HasProjectionName returns a boolean if a field has been set.
func (o *DataDestination) HasProjectionName() bool {
	if o != nil && !IsNil(o.ProjectionName) {
		return true
	}

	return false
}

// SetProjectionName gets a reference to the given string and assigns it to the ProjectionName field.
func (o *DataDestination) SetProjectionName(v string) {
	o.ProjectionName = &v
}

// GetFilterId returns the FilterId field value if set, zero value otherwise.
func (o *DataDestination) GetFilterId() string {
	if o == nil || IsNil(o.FilterId) {
		var ret string
		return ret
	}
	return *o.FilterId
}

// GetFilterIdOk returns a tuple with the FilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetFilterIdOk() (*string, bool) {
	if o == nil || IsNil(o.FilterId) {
		return nil, false
	}
	return o.FilterId, true
}

// HasFilterId returns a boolean if a field has been set.
func (o *DataDestination) HasFilterId() bool {
	if o != nil && !IsNil(o.FilterId) {
		return true
	}

	return false
}

// SetFilterId gets a reference to the given string and assigns it to the FilterId field.
func (o *DataDestination) SetFilterId(v string) {
	o.FilterId = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DataDestination) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DataDestination) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DataDestination) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DataDestination) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDestination) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DataDestination) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DataDestination) SetState(v string) {
	o.State = &v
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DestinationEndpoint) {
		toSerialize["destinationEndpoint"] = o.DestinationEndpoint
	}
	if !IsNil(o.ProjectionName) {
		toSerialize["projectionName"] = o.ProjectionName
	}
	if !IsNil(o.FilterId) {
		toSerialize["filterId"] = o.FilterId
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
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

