# CreateDataDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationEndpoint** | **string** | The HTTP endpoint which Ambar will deliver messages too. The endpoint should accept a username and password credential pair for authentication. | 
**FilterId** | **string** | The Ambar ResourceId of a the filter in your environment to use to determine what data sources to pull records from and how to filter record sequences before delivery to your endpoint. | 
**ProjectionName** | **string** | A name for this DataDestination projection | 
**Password** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 
**Username** | **string** | A case sensitive string for the user Ambar should use to connect to your HTTP endpoint service. | 

## Methods

### NewCreateDataDestinationRequest

`func NewCreateDataDestinationRequest(destinationEndpoint string, filterId string, projectionName string, password string, username string, ) *CreateDataDestinationRequest`

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


### GetFilterId

`func (o *CreateDataDestinationRequest) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *CreateDataDestinationRequest) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *CreateDataDestinationRequest) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.


### GetProjectionName

`func (o *CreateDataDestinationRequest) GetProjectionName() string`

GetProjectionName returns the ProjectionName field if non-nil, zero value otherwise.

### GetProjectionNameOk

`func (o *CreateDataDestinationRequest) GetProjectionNameOk() (*string, bool)`

GetProjectionNameOk returns a tuple with the ProjectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionName

`func (o *CreateDataDestinationRequest) SetProjectionName(v string)`

SetProjectionName sets ProjectionName field to given value.


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


