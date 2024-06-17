# DataDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The UTC time at which the DataDestination was created. | 
**DestinationEndpoint** | **string** | The HTTP endpoint which Ambar sends filtered record sequences to. | 
**Description** | Pointer to **string** | The description of the DataDestination given when it was created. | [optional] 
**FilterIds** | **[]string** | The Ambar resourceId for the Filters which are applied to record sequences sent to the configured DataDestination. | 
**ResourceId** | **string** | The Ambar resourceId corresponding to this DataDestination. | 
**State** | **string** | The ResourceState of this DataDestination. For possible values see ResourceState in our developer docs. | 

## Methods

### NewDataDestination

`func NewDataDestination(createdAt string, destinationEndpoint string, filterIds []string, resourceId string, state string, ) *DataDestination`

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


### GetDescription

`func (o *DataDestination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataDestination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataDestination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataDestination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterIds

`func (o *DataDestination) GetFilterIds() []string`

GetFilterIds returns the FilterIds field if non-nil, zero value otherwise.

### GetFilterIdsOk

`func (o *DataDestination) GetFilterIdsOk() (*[]string, bool)`

GetFilterIdsOk returns a tuple with the FilterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIds

`func (o *DataDestination) SetFilterIds(v []string)`

SetFilterIds sets FilterIds field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


