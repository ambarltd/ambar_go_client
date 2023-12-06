# DataDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The UTC time at which the DataDestination was created. | [optional] 
**DestinationEndpoint** | Pointer to **string** | The HTTP endpoint which Ambar sends filtered record sequences to. | [optional] 
**ProjectionName** | Pointer to **string** | The name given to this DataDestinations projection at creation. | [optional] 
**FilterId** | Pointer to **string** | The Ambar resourceId for the Filter which is applied to record sequences sent to this Data Destination. | [optional] 
**ResourceId** | Pointer to **string** | The Ambar resourceId corresponding to this DataDestination. | [optional] 
**State** | Pointer to **string** | The ResourceState of this DataDestination. For possible values see ResourceState in our developer docs. | [optional] 

## Methods

### NewDataDestination

`func NewDataDestination() *DataDestination`

NewDataDestination instantiates a new DataDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDestinationWithDefaults

`func NewDataDestinationWithDefaults() *DataDestination`

NewDataDestinationWithDefaults instantiates a new DataDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataDestination) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataDestination) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataDestination) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataDestination) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestinationEndpoint

`func (o *DataDestination) GetDestinationEndpoint() string`

GetDestinationEndpoint returns the DestinationEndpoint field if non-nil, zero value otherwise.

### GetDestinationEndpointOk

`func (o *DataDestination) GetDestinationEndpointOk() (*string, bool)`

GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationEndpoint

`func (o *DataDestination) SetDestinationEndpoint(v string)`

SetDestinationEndpoint sets DestinationEndpoint field to given value.

### HasDestinationEndpoint

`func (o *DataDestination) HasDestinationEndpoint() bool`

HasDestinationEndpoint returns a boolean if a field has been set.

### GetProjectionName

`func (o *DataDestination) GetProjectionName() string`

GetProjectionName returns the ProjectionName field if non-nil, zero value otherwise.

### GetProjectionNameOk

`func (o *DataDestination) GetProjectionNameOk() (*string, bool)`

GetProjectionNameOk returns a tuple with the ProjectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionName

`func (o *DataDestination) SetProjectionName(v string)`

SetProjectionName sets ProjectionName field to given value.

### HasProjectionName

`func (o *DataDestination) HasProjectionName() bool`

HasProjectionName returns a boolean if a field has been set.

### GetFilterId

`func (o *DataDestination) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *DataDestination) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *DataDestination) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *DataDestination) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetResourceId

`func (o *DataDestination) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DataDestination) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DataDestination) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DataDestination) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetState

`func (o *DataDestination) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataDestination) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataDestination) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataDestination) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


