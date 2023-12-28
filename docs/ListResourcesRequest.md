# ListResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **string** |  | [optional] 

## Methods

### NewListResourcesRequest

`func NewListResourcesRequest() *ListResourcesRequest`

NewListResourcesRequest instantiates a new ListResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesRequestWithDefaults

`func NewListResourcesRequestWithDefaults() *ListResourcesRequest`

NewListResourcesRequestWithDefaults instantiates a new ListResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ListResourcesRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListResourcesRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListResourcesRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListResourcesRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetPage

`func (o *ListResourcesRequest) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListResourcesRequest) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListResourcesRequest) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListResourcesRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


