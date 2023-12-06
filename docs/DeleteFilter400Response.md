# DeleteFilter400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidParameterException** | **string** |  | 
**ResourceInvalidStateException** | **string** |  | 

## Methods

### NewDeleteFilter400Response

`func NewDeleteFilter400Response(invalidParameterException string, resourceInvalidStateException string, ) *DeleteFilter400Response`

NewDeleteFilter400Response instantiates a new DeleteFilter400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFilter400ResponseWithDefaults

`func NewDeleteFilter400ResponseWithDefaults() *DeleteFilter400Response`

NewDeleteFilter400ResponseWithDefaults instantiates a new DeleteFilter400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidParameterException

`func (o *DeleteFilter400Response) GetInvalidParameterException() string`

GetInvalidParameterException returns the InvalidParameterException field if non-nil, zero value otherwise.

### GetInvalidParameterExceptionOk

`func (o *DeleteFilter400Response) GetInvalidParameterExceptionOk() (*string, bool)`

GetInvalidParameterExceptionOk returns a tuple with the InvalidParameterException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParameterException

`func (o *DeleteFilter400Response) SetInvalidParameterException(v string)`

SetInvalidParameterException sets InvalidParameterException field to given value.


### GetResourceInvalidStateException

`func (o *DeleteFilter400Response) GetResourceInvalidStateException() string`

GetResourceInvalidStateException returns the ResourceInvalidStateException field if non-nil, zero value otherwise.

### GetResourceInvalidStateExceptionOk

`func (o *DeleteFilter400Response) GetResourceInvalidStateExceptionOk() (*string, bool)`

GetResourceInvalidStateExceptionOk returns a tuple with the ResourceInvalidStateException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInvalidStateException

`func (o *DeleteFilter400Response) SetResourceInvalidStateException(v string)`

SetResourceInvalidStateException sets ResourceInvalidStateException field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


