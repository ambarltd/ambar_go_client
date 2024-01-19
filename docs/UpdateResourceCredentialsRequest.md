# UpdateResourceCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | An Ambar resourceId. | 
**Password** | **string** | The password portion of the credential to update for your resource. | 
**Username** | **string** | The username portion of the credential to update for your resource. | 

## Methods

### NewUpdateResourceCredentialsRequest

`func NewUpdateResourceCredentialsRequest(resourceId string, password string, username string, ) *UpdateResourceCredentialsRequest`

NewUpdateResourceCredentialsRequest instantiates a new UpdateResourceCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceCredentialsRequestWithDefaults

`func NewUpdateResourceCredentialsRequestWithDefaults() *UpdateResourceCredentialsRequest`

NewUpdateResourceCredentialsRequestWithDefaults instantiates a new UpdateResourceCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UpdateResourceCredentialsRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateResourceCredentialsRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateResourceCredentialsRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetPassword

`func (o *UpdateResourceCredentialsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateResourceCredentialsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateResourceCredentialsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *UpdateResourceCredentialsRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateResourceCredentialsRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateResourceCredentialsRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


