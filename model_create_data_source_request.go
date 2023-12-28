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

// checks if the CreateDataSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataSourceRequest{}

// CreateDataSourceRequest The request properties for creating an Ambar Data Source.
type CreateDataSourceRequest struct {
	// A case-insensitive string value describing the type of durable storage system you wish to connect to your Ambar environment. See supported DataSourceTypes in our developer docs for more details on valid values.
	DataSourceType string `json:"dataSourceType"`
	// A description for identifying this DataSource.
	Description *string `json:"description,omitempty"`
	// A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service.
	Username string `json:"username"`
	// A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service.
	Password string `json:"password"`
	// The name of a column which monotonically increases on database writes.
	SerialColumn string `json:"serialColumn"`
	// A case sensitive string for the name of the column in your table Ambar can partition on.  Note that partition keys must be unique to a given sequence of records.
	PartitioningColumn string `json:"partitioningColumn"`
	// A object containing inputs which are specific depending on the type of DataSource being created. See out developer docs for supported DataSourceTypes and corresponding configurations.
	DataSourceConfig map[string]string `json:"dataSourceConfig"`
}

// NewCreateDataSourceRequest instantiates a new CreateDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataSourceRequest(dataSourceType string, username string, password string, serialColumn string, partitioningColumn string, dataSourceConfig map[string]string) *CreateDataSourceRequest {
	this := CreateDataSourceRequest{}
	this.DataSourceType = dataSourceType
	this.Username = username
	this.Password = password
	this.SerialColumn = serialColumn
	this.PartitioningColumn = partitioningColumn
	this.DataSourceConfig = dataSourceConfig
	return &this
}

// NewCreateDataSourceRequestWithDefaults instantiates a new CreateDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataSourceRequestWithDefaults() *CreateDataSourceRequest {
	this := CreateDataSourceRequest{}
	return &this
}

// GetDataSourceType returns the DataSourceType field value
func (o *CreateDataSourceRequest) GetDataSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetDataSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceType, true
}

// SetDataSourceType sets field value
func (o *CreateDataSourceRequest) SetDataSourceType(v string) {
	o.DataSourceType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDataSourceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDataSourceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDataSourceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUsername returns the Username field value
func (o *CreateDataSourceRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateDataSourceRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CreateDataSourceRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateDataSourceRequest) SetPassword(v string) {
	o.Password = v
}

// GetSerialColumn returns the SerialColumn field value
func (o *CreateDataSourceRequest) GetSerialColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialColumn
}

// GetSerialColumnOk returns a tuple with the SerialColumn field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetSerialColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialColumn, true
}

// SetSerialColumn sets field value
func (o *CreateDataSourceRequest) SetSerialColumn(v string) {
	o.SerialColumn = v
}

// GetPartitioningColumn returns the PartitioningColumn field value
func (o *CreateDataSourceRequest) GetPartitioningColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartitioningColumn
}

// GetPartitioningColumnOk returns a tuple with the PartitioningColumn field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetPartitioningColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitioningColumn, true
}

// SetPartitioningColumn sets field value
func (o *CreateDataSourceRequest) SetPartitioningColumn(v string) {
	o.PartitioningColumn = v
}

// GetDataSourceConfig returns the DataSourceConfig field value
func (o *CreateDataSourceRequest) GetDataSourceConfig() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.DataSourceConfig
}

// GetDataSourceConfigOk returns a tuple with the DataSourceConfig field value
// and a boolean to check if the value has been set.
func (o *CreateDataSourceRequest) GetDataSourceConfigOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceConfig, true
}

// SetDataSourceConfig sets field value
func (o *CreateDataSourceRequest) SetDataSourceConfig(v map[string]string) {
	o.DataSourceConfig = v
}

func (o CreateDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataSourceType"] = o.DataSourceType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["serialColumn"] = o.SerialColumn
	toSerialize["partitioningColumn"] = o.PartitioningColumn
	toSerialize["dataSourceConfig"] = o.DataSourceConfig
	return toSerialize, nil
}

type NullableCreateDataSourceRequest struct {
	value *CreateDataSourceRequest
	isSet bool
}

func (v NullableCreateDataSourceRequest) Get() *CreateDataSourceRequest {
	return v.value
}

func (v *NullableCreateDataSourceRequest) Set(val *CreateDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataSourceRequest(val *CreateDataSourceRequest) *NullableCreateDataSourceRequest {
	return &NullableCreateDataSourceRequest{value: val, isSet: true}
}

func (v NullableCreateDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


