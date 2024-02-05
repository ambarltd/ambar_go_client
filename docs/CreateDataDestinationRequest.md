# CreateDataDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationEndpoint** | **string** | The HTTP endpoint which Ambar will deliver messages too. The endpoint should accept a username and password credential pair for authentication. | 
**Description** | Pointer to **string** | A description for identifying this DataDestination. | [optional] 
**FilterIds** | **[]string** | The Ambar ResourceId of a the filters in your environment to use to determine what data sources to pull records from and how to filter record sequences before delivery to your endpoint. | 
**Password** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 
**Username** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 

## Methods

### NewCreateDataDestinationRequest

`func NewCreateDataDestinationRequest(destinationEndpoint string, filterIds []string, password string, username string, ) *CreateDataDestinationRequest`

NewCreateDataDestinationRequest instantiates a new CreateDataDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataDestinationRequestWithDefaults

`func NewCreateDataDestinationRequestWithDefaults() *CreateDataDestinationRequest`

NewCreateDataDestinationRequestWithDefaults instantiates a new CreateDataDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationEndpoint

`func (o *CreateDataDestinationRequest) GetDestinationEndpoint() string`

GetDestinationEndpoint returns the DestinationEndpoint field if non-nil, zero value otherwise.

### GetDestinationEndpointOk

`func (o *CreateDataDestinationRequest) GetDestinationEndpointOk() (*string, bool)`

GetDestinationEndpointOk returns a tuple with the DestinationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationEndpoint

`func (o *CreateDataDestinationRequest) SetDestinationEndpoint(v string)`

SetDestinationEndpoint sets DestinationEndpoint field to given value.


### GetDescription

`func (o *CreateDataDestinationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDataDestinationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDataDestinationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDataDestinationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterIds

`func (o *CreateDataDestinationRequest) GetFilterIds() []string`

GetFilterIds returns the FilterIds field if non-nil, zero value otherwise.

### GetFilterIdsOk

`func (o *CreateDataDestinationRequest) GetFilterIdsOk() (*[]string, bool)`

GetFilterIdsOk returns a tuple with the FilterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIds

`func (o *CreateDataDestinationRequest) SetFilterIds(v []string)`

SetFilterIds sets FilterIds field to given value.


### GetPassword

`func (o *CreateDataDestinationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDataDestinationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDataDestinationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *CreateDataDestinationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDataDestinationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDataDestinationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


