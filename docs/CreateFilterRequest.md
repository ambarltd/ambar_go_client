# CreateFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterContents** | **string** | A base64 encoded Ambar Filter. For details about Ambar filter syntax, see our developer docs. | 
**DataSourceId** | **string** | The Resource Id of a DataSource resource, which this filter should be applied to by a DataDestination. | 
**Description** | Pointer to **string** | A description for identifying this Filter. | [optional] 

## Methods

### NewCreateFilterRequest

`func NewCreateFilterRequest(filterContents string, dataSourceId string, ) *CreateFilterRequest`

NewCreateFilterRequest instantiates a new CreateFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFilterRequestWithDefaults

`func NewCreateFilterRequestWithDefaults() *CreateFilterRequest`

NewCreateFilterRequestWithDefaults instantiates a new CreateFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterContents

`func (o *CreateFilterRequest) GetFilterContents() string`

GetFilterContents returns the FilterContents field if non-nil, zero value otherwise.

### GetFilterContentsOk

`func (o *CreateFilterRequest) GetFilterContentsOk() (*string, bool)`

GetFilterContentsOk returns a tuple with the FilterContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterContents

`func (o *CreateFilterRequest) SetFilterContents(v string)`

SetFilterContents sets FilterContents field to given value.


### GetDataSourceId

`func (o *CreateFilterRequest) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *CreateFilterRequest) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *CreateFilterRequest) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.


### GetDescription

`func (o *CreateFilterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFilterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFilterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFilterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


