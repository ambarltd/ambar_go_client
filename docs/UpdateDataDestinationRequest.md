# UpdateDataDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationEndpoint** | Pointer to **string** | The HTTP endpoint which Ambar will deliver messages too. The endpoint should accept a username and password credential pair for authentication. | [optional] 
**DestinationPort** | Pointer to **string** | The port on which your endpoint service is listening for connections. | [optional] 
**ResourceId** | **string** | An Ambar resourceId. | 

## Methods

### NewUpdateDataDestinationRequest

`func NewUpdateDataDestinationRequest(resourceId string, ) *UpdateDataDestinationRequest`

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

### HasDestinationEndpoint

`func (o *UpdateDataDestinationRequest) HasDestinationEndpoint() bool`

HasDestinationEndpoint returns a boolean if a field has been set.

### GetDestinationPort

`func (o *UpdateDataDestinationRequest) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *UpdateDataDestinationRequest) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *UpdateDataDestinationRequest) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *UpdateDataDestinationRequest) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

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


