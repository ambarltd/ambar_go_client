/*
Ambar OpenAPI Specification

Details about communicating with Ambar.cloud public endpoints. Supported HTTP rest endpoints and their  request and response details.

API version: 2025-04-02
Contact: contact@ambar.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ambar

import (
	"encoding/json"
)

// checks if the ListResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourcesRequest{}

// ListResourcesRequest struct for ListResourcesRequest
type ListResourcesRequest struct {
	// 
	ResourceType *string `json:"resourceType,omitempty"`
	// 
	Page *string `json:"page,omitempty"`
}

// NewListResourcesRequest instantiates a new ListResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourcesRequest() *ListResourcesRequest {
	this := ListResourcesRequest{}
	return &this
}

// NewListResourcesRequestWithDefaults instantiates a new ListResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourcesRequestWithDefaults() *ListResourcesRequest {
	this := ListResourcesRequest{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ListResourcesRequest) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ListResourcesRequest) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ListResourcesRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListResourcesRequest) GetPage() string {
	if o == nil || IsNil(o.Page) {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequest) GetPageOk() (*string, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListResourcesRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *ListResourcesRequest) SetPage(v string) {
	o.Page = &v
}

func (o ListResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableListResourcesRequest struct {
	value *ListResourcesRequest
	isSet bool
}

func (v NullableListResourcesRequest) Get() *ListResourcesRequest {
	return v.value
}

func (v *NullableListResourcesRequest) Set(val *ListResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourcesRequest(val *ListResourcesRequest) *NullableListResourcesRequest {
	return &NullableListResourcesRequest{value: val, isSet: true}
}

func (v NullableListResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


