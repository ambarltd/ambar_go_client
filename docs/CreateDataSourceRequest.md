# CreateDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceType** | **string** | A case-insensitive string value describing the type of durable storage system you wish to connect to your Ambar environment. See supported DataSourceTypes in our developer docs for more details on valid values. | 
**Description** | Pointer to **string** | A description for identifying this DataSource. | [optional] 
**DataSourceConfig** | **map[string]string** | A object containing inputs which are specific depending on the type of DataSource being created. See out developer docs for supported DataSourceTypes and corresponding configurations. | 

## Methods

### NewCreateDataSourceRequest

`func NewCreateDataSourceRequest(dataSourceType string, dataSourceConfig map[string]string, ) *CreateDataSourceRequest`

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


