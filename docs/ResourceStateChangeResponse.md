# ResourceStateChangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewResourceStateChangeResponse

`func NewResourceStateChangeResponse(resourceId string, state string, ) *ResourceStateChangeResponse`

NewResourceStateChangeResponse instantiates a new ResourceStateChangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceStateChangeResponseWithDefaults

`func NewResourceStateChangeResponseWithDefaults() *ResourceStateChangeResponse`

NewResourceStateChangeResponseWithDefaults instantiates a new ResourceStateChangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceStateChangeResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceStateChangeResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceStateChangeResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetState

`func (o *ResourceStateChangeResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceStateChangeResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceStateChangeResponse) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


