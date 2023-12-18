# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The UTC time at which the Filter was created. | [optional] 
**DataDestinationsUsingFilter** | Pointer to **[]string** | The list of DataDestination ResourceIds which use this Filter. | [optional] 
**DataSourcesInFilter** | Pointer to **[]string** | The list of DataSource ResourceIds which this Filter will read from. | [optional] 
**ResourceId** | Pointer to **string** | The Ambar resourceId corresponding to this Filter. | [optional] 
**Description** | Pointer to **string** | The description of the Filter given when it was created. | [optional] 
**State** | Pointer to **string** | The ResourceState of this Filter. For possible values see ResourceState in our developer docs. | [optional] 

## Methods

### NewFilter

`func NewFilter() *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Filter) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Filter) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Filter) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Filter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataDestinationsUsingFilter

`func (o *Filter) GetDataDestinationsUsingFilter() []string`

GetDataDestinationsUsingFilter returns the DataDestinationsUsingFilter field if non-nil, zero value otherwise.

### GetDataDestinationsUsingFilterOk

`func (o *Filter) GetDataDestinationsUsingFilterOk() (*[]string, bool)`

GetDataDestinationsUsingFilterOk returns a tuple with the DataDestinationsUsingFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDestinationsUsingFilter

`func (o *Filter) SetDataDestinationsUsingFilter(v []string)`

SetDataDestinationsUsingFilter sets DataDestinationsUsingFilter field to given value.

### HasDataDestinationsUsingFilter

`func (o *Filter) HasDataDestinationsUsingFilter() bool`

HasDataDestinationsUsingFilter returns a boolean if a field has been set.

### GetDataSourcesInFilter

`func (o *Filter) GetDataSourcesInFilter() []string`

GetDataSourcesInFilter returns the DataSourcesInFilter field if non-nil, zero value otherwise.

### GetDataSourcesInFilterOk

`func (o *Filter) GetDataSourcesInFilterOk() (*[]string, bool)`

GetDataSourcesInFilterOk returns a tuple with the DataSourcesInFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourcesInFilter

`func (o *Filter) SetDataSourcesInFilter(v []string)`

SetDataSourcesInFilter sets DataSourcesInFilter field to given value.

### HasDataSourcesInFilter

`func (o *Filter) HasDataSourcesInFilter() bool`

HasDataSourcesInFilter returns a boolean if a field has been set.

### GetResourceId

`func (o *Filter) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Filter) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Filter) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Filter) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetDescription

`func (o *Filter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Filter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Filter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Filter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *Filter) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Filter) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Filter) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Filter) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


