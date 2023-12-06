# ResourceTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to [**ResourceTypeDetailsLimit**](ResourceTypeDetailsLimit.md) |  | [optional] 
**Resources** | Pointer to [**[]ResourceDetails**](ResourceDetails.md) |  | [optional] 

## Methods

### NewResourceTypeDetails

`func NewResourceTypeDetails() *ResourceTypeDetails`

NewResourceTypeDetails instantiates a new ResourceTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeDetailsWithDefaults

`func NewResourceTypeDetailsWithDefaults() *ResourceTypeDetails`

NewResourceTypeDetailsWithDefaults instantiates a new ResourceTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ResourceTypeDetails) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceTypeDetails) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceTypeDetails) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceTypeDetails) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetLimit

`func (o *ResourceTypeDetails) GetLimit() ResourceTypeDetailsLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResourceTypeDetails) GetLimitOk() (*ResourceTypeDetailsLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResourceTypeDetails) SetLimit(v ResourceTypeDetailsLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResourceTypeDetails) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetResources

`func (o *ResourceTypeDetails) GetResources() []ResourceDetails`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceTypeDetails) GetResourcesOk() (*[]ResourceDetails, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceTypeDetails) SetResources(v []ResourceDetails)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourceTypeDetails) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


