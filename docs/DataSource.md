# DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The UTC time at which the DataSource was created. | [optional] 
**DataSourceConfig** | Pointer to **map[string]interface{}** | The properties describing the configuration details for the given DataSourceType. | [optional] 
**DataSourceType** | Pointer to **string** | The DataSourceType describing the type of durable storage system this DataSource pulls record sequences from. | [optional] 
**ResourceId** | Pointer to **string** | The Ambar resourceId corresponding to this DataSource. | [optional] 
**State** | Pointer to **string** | The ResourceState of this DataSource. For possible values see ResourceState in our developer docs. | [optional] 

## Methods

### NewDataSource

`func NewDataSource() *DataSource`

NewDataSource instantiates a new DataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceWithDefaults

`func NewDataSourceWithDefaults() *DataSource`

NewDataSourceWithDefaults instantiates a new DataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataSource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataSource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataSource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataSource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataSourceConfig

`func (o *DataSource) GetDataSourceConfig() map[string]interface{}`

GetDataSourceConfig returns the DataSourceConfig field if non-nil, zero value otherwise.

### GetDataSourceConfigOk

`func (o *DataSource) GetDataSourceConfigOk() (*map[string]interface{}, bool)`

GetDataSourceConfigOk returns a tuple with the DataSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceConfig

`func (o *DataSource) SetDataSourceConfig(v map[string]interface{})`

SetDataSourceConfig sets DataSourceConfig field to given value.

### HasDataSourceConfig

`func (o *DataSource) HasDataSourceConfig() bool`

HasDataSourceConfig returns a boolean if a field has been set.

### GetDataSourceType

`func (o *DataSource) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *DataSource) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *DataSource) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *DataSource) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *DataSource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DataSource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DataSource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DataSource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetState

`func (o *DataSource) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataSource) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataSource) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataSource) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


