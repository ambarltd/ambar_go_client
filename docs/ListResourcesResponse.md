# ListResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | Pointer to **int32** | The next page available for listing if there is one. | [optional] 
**Resources** | Pointer to [**[]ResourceTypeDetails**](ResourceTypeDetails.md) |  | [optional] 

## Methods

### NewListResourcesResponse

`func NewListResourcesResponse() *ListResourcesResponse`

NewListResourcesResponse instantiates a new ListResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesResponseWithDefaults

`func NewListResourcesResponseWithDefaults() *ListResourcesResponse`

NewListResourcesResponseWithDefaults instantiates a new ListResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *ListResourcesResponse) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ListResourcesResponse) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ListResourcesResponse) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ListResourcesResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetResources

`func (o *ListResourcesResponse) GetResources() []ResourceTypeDetails`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListResourcesResponse) GetResourcesOk() (*[]ResourceTypeDetails, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListResourcesResponse) SetResources(v []ResourceTypeDetails)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ListResourcesResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


