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

// checks if the ResourceTypeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceTypeDetails{}

// ResourceTypeDetails struct for ResourceTypeDetails
type ResourceTypeDetails struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Limit *ResourceTypeDetailsLimit `json:"limit,omitempty"`
	Resources []ResourceDetails `json:"resources,omitempty"`
}

// NewResourceTypeDetails instantiates a new ResourceTypeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeDetails() *ResourceTypeDetails {
	this := ResourceTypeDetails{}
	return &this
}

// NewResourceTypeDetailsWithDefaults instantiates a new ResourceTypeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeDetailsWithDefaults() *ResourceTypeDetails {
	this := ResourceTypeDetails{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceTypeDetails) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeDetails) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceTypeDetails) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourceTypeDetails) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResourceTypeDetails) GetLimit() ResourceTypeDetailsLimit {
	if o == nil || IsNil(o.Limit) {
		var ret ResourceTypeDetailsLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeDetails) GetLimitOk() (*ResourceTypeDetailsLimit, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResourceTypeDetails) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given ResourceTypeDetailsLimit and assigns it to the Limit field.
func (o *ResourceTypeDetails) SetLimit(v ResourceTypeDetailsLimit) {
	o.Limit = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceTypeDetails) GetResources() []ResourceDetails {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourceDetails
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeDetails) GetResourcesOk() ([]ResourceDetails, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceTypeDetails) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceDetails and assigns it to the Resources field.
func (o *ResourceTypeDetails) SetResources(v []ResourceDetails) {
	o.Resources = v
}

func (o ResourceTypeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceTypeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableResourceTypeDetails struct {
	value *ResourceTypeDetails
	isSet bool
}

func (v NullableResourceTypeDetails) Get() *ResourceTypeDetails {
	return v.value
}

func (v *NullableResourceTypeDetails) Set(val *ResourceTypeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeDetails(val *ResourceTypeDetails) *NullableResourceTypeDetails {
	return &NullableResourceTypeDetails{value: val, isSet: true}
}

func (v NullableResourceTypeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

