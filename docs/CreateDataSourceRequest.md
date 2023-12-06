# CreateDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceConfig** | [**CreateDataSourceRequestDataSourceConfig**](CreateDataSourceRequestDataSourceConfig.md) |  | 
**DataSourceType** | **string** | A case-insensitive string value describing the type of durable storage system you wish to connect to your Ambar environment. See supported DataSourceTypes in our developer docs for more details on valid values. | 

## Methods

### NewCreateDataSourceRequest

`func NewCreateDataSourceRequest(dataSourceConfig CreateDataSourceRequestDataSourceConfig, dataSourceType string, ) *CreateDataSourceRequest`

NewCreateDataSourceRequest instantiates a new CreateDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceRequestWithDefaults

`func NewCreateDataSourceRequestWithDefaults() *CreateDataSourceRequest`

NewCreateDataSourceRequestWithDefaults instantiates a new CreateDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceConfig

`func (o *CreateDataSourceRequest) GetDataSourceConfig() CreateDataSourceRequestDataSourceConfig`

GetDataSourceConfig returns the DataSourceConfig field if non-nil, zero value otherwise.

### GetDataSourceConfigOk

`func (o *CreateDataSourceRequest) GetDataSourceConfigOk() (*CreateDataSourceRequestDataSourceConfig, bool)`

GetDataSourceConfigOk returns a tuple with the DataSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceConfig

`func (o *CreateDataSourceRequest) SetDataSourceConfig(v CreateDataSourceRequestDataSourceConfig)`

SetDataSourceConfig sets DataSourceConfig field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


