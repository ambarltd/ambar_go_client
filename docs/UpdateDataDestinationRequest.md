# UpdateDataDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationEndpoint** | **string** | The HTTP endpoint which Ambar will deliver messages too. The endpoint should accept a username and password credential pair for authentication. | 
**ResourceId** | **string** | An Ambar resourceId. | 

## Methods

### NewUpdateDataDestinationRequest

`func NewUpdateDataDestinationRequest(destinationEndpoint string, resourceId string, ) *UpdateDataDestinationRequest`

NewUpdateDataDestinationRequest instantiates a new UpdateDataDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataDestinationRequestWithDefaults

`func NewUpdateDataDestinationRequestWithDefaults() *UpdateDataDestinationRequest`

NewUpdateDataDestinationRequestWithDefaults instantiates a new UpdateDataDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationEndpoint

`func (o *UpdateDataDestinationRequest) GetDestinationEndpoint() string`

GetDestinationEndpoint returns the DestinationEndpoint field if non-nil, zero value otherwise.

### GetDestinationEndpointOk

`func (o *UpdateDataDestinationRequest) GetDestinationEndpointOk() (*string, bool)`

GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationEndpoint

`func (o *UpdateDataDestinationRequest) SetDestinationEndpoint(v string)`

SetDestinationEndpoint sets DestinationEndpoint field to given value.


### GetResourceId

`func (o *UpdateDataDestinationRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateDataDestinationRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateDataDestinationRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


