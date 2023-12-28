# CreateDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceType** | **string** | A case-insensitive string value describing the type of durable storage system you wish to connect to your Ambar environment. See supported DataSourceTypes in our developer docs for more details on valid values. | 
**Description** | Pointer to **string** | A description for identifying this DataSource. | [optional] 
**Username** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 
**Password** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 
**SerialColumn** | **string** | The name of a column which monotonically increases on database writes. | 
**PartitioningColumn** | **string** | A case sensitive string for the name of the column in your table Ambar can partition on.  Note that partition keys must be unique to a given sequence of records. | 
**DataSourceConfig** | **map[string]string** | A object containing inputs which are specific depending on the type of DataSource being created. See out developer docs for supported DataSourceTypes and corresponding configurations. | 

## Methods

### NewCreateDataSourceRequest

`func NewCreateDataSourceRequest(dataSourceType string, username string, password string, serialColumn string, partitioningColumn string, dataSourceConfig map[string]string, ) *CreateDataSourceRequest`

NewCreateDataSourceRequest instantiates a new CreateDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceRequestWithDefaults

`func NewCreateDataSourceRequestWithDefaults() *CreateDataSourceRequest`

NewCreateDataSourceRequestWithDefaults instantiates a new CreateDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceType

`func (o *CreateDataSourceRequest) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *CreateDataSourceRequest) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *CreateDataSourceRequest) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.


### GetDescription

`func (o *CreateDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDataSourceRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDataSourceRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDataSourceRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateDataSourceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDataSourceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDataSourceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSerialColumn

`func (o *CreateDataSourceRequest) GetSerialColumn() string`

GetSerialColumn returns the SerialColumn field if non-nil, zero value otherwise.

### GetSerialColumnOk

`func (o *CreateDataSourceRequest) GetSerialColumnOk() (*string, bool)`

GetSerialColumnOk returns a tuple with the SerialColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialColumn

`func (o *CreateDataSourceRequest) SetSerialColumn(v string)`

SetSerialColumn sets SerialColumn field to given value.


### GetPartitioningColumn

`func (o *CreateDataSourceRequest) GetPartitioningColumn() string`

GetPartitioningColumn returns the PartitioningColumn field if non-nil, zero value otherwise.

### GetPartitioningColumnOk

`func (o *CreateDataSourceRequest) GetPartitioningColumnOk() (*string, bool)`

GetPartitioningColumnOk returns a tuple with the PartitioningColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningColumn

`func (o *CreateDataSourceRequest) SetPartitioningColumn(v string)`

SetPartitioningColumn sets PartitioningColumn field to given value.


### GetDataSourceConfig

`func (o *CreateDataSourceRequest) GetDataSourceConfig() map[string]string`

GetDataSourceConfig returns the DataSourceConfig field if non-nil, zero value otherwise.

### GetDataSourceConfigOk

`func (o *CreateDataSourceRequest) GetDataSourceConfigOk() (*map[string]string, bool)`

GetDataSourceConfigOk returns a tuple with the DataSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceConfig

`func (o *CreateDataSourceRequest) SetDataSourceConfig(v map[string]string)`

SetDataSourceConfig sets DataSourceConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


