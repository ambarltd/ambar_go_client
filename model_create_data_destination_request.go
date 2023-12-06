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

// checks if the CreateDataDestinationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataDestinationRequest{}

// CreateDataDestinationRequest The request properties for creating an Ambar Data Destination.
type CreateDataDestinationRequest struct {
	// The HTTP endpoint which Ambar will deliver messages too. The endpoint should accept a username and password credential pair for authentication.
	DestinationEndpoint string `json:"destinationEndpoint"`
	// The Ambar ResourceId of a the filter in your environment to use to determine what data sources to pull records from and how to filter record sequences before delivery to your endpoint.
	FilterId string `json:"filterId"`
	// A name for this DataDestination projection
	ProjectionName string `json:"projectionName"`
	// A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service.
	Password string `json:"password"`
	// A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service.
	Username string `json:"username"`
}

// NewCreateDataDestinationRequest instantiates a new CreateDataDestinationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataDestinationRequest(destinationEndpoint string, filterId string, projectionName string, password string, username string) *CreateDataDestinationRequest {
	this := CreateDataDestinationRequest{}
	this.DestinationEndpoint = destinationEndpoint
	this.FilterId = filterId
	this.ProjectionName = projectionName
	this.Password = password
	this.Username = username
	return &this
}

// NewCreateDataDestinationRequestWithDefaults instantiates a new CreateDataDestinationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataDestinationRequestWithDefaults() *CreateDataDestinationRequest {
	this := CreateDataDestinationRequest{}
	return &this
}

// GetDestinationEndpoint returns the DestinationEndpoint field value
func (o *CreateDataDestinationRequest) GetDestinationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationEndpoint
}

// GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field value
// and a boolean to check if the value has been set.
func (o *CreateDataDestinationRequest) GetDestinationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationEndpoint, true
}

// SetDestinationEndpoint sets field value
func (o *CreateDataDestinationRequest) SetDestinationEndpoint(v string) {
	o.DestinationEndpoint = v
}

// GetFilterId returns the FilterId field value
func (o *CreateDataDestinationRequest) GetFilterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterId
}

// GetFilterIdOk returns a tuple with the FilterId field value
// and a boolean to check if the value has been set.
func (o *CreateDataDestinationRequest) GetFilterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterId, true
}

// SetFilterId sets field value
func (o *CreateDataDestinationRequest) SetFilterId(v string) {
	o.FilterId = v
}

// GetProjectionName returns the ProjectionName field value
func (o *CreateDataDestinationRequest) GetProjectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectionName
}

// GetProjectionNameOk returns a tuple with the ProjectionName field value
// and a boolean to check if the value has been set.
func (o *CreateDataDestinationRequest) GetProjectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectionName, true
}

// SetProjectionName sets field value
func (o *CreateDataDestinationRequest) SetProjectionName(v string) {
	o.ProjectionName = v
}

// GetPassword returns the Password field value
func (o *CreateDataDestinationRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateDataDestinationRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateDataDestinationRequest) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *CreateDataDestinationRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateDataDestinationRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateDataDestinationRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateDataDestinationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataDestinationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationEndpoint"] = o.DestinationEndpoint
	toSerialize["filterId"] = o.FilterId
	toSerialize["projectionName"] = o.ProjectionName
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCreateDataDestinationRequest struct {
	value *CreateDataDestinationRequest
	isSet bool
}

func (v NullableCreateDataDestinationRequest) Get() *CreateDataDestinationRequest {
	return v.value
}

func (v *NullableCreateDataDestinationRequest) Set(val *CreateDataDestinationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataDestinationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataDestinationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataDestinationRequest(val *CreateDataDestinationRequest) *NullableCreateDataDestinationRequest {
	return &NullableCreateDataDestinationRequest{value: val, isSet: true}
}

func (v NullableCreateDataDestinationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataDestinationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

