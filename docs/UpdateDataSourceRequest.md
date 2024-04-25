# UpdateDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | An Ambar resourceId. | 
**Hostname** | Pointer to **string** | The hostname of the database which the DataSource connects to. | [optional] 
**Port** | Pointer to **string** | The port on which the database host is listening for incoming connections. | [optional] 
**TlsTerminationOverrideHost** | Pointer to **string** | The hostname of an alternative host handling TLS termination. | [optional] 

## Methods

### NewUpdateDataSourceRequest

`func NewUpdateDataSourceRequest(resourceId string, ) *UpdateDataSourceRequest`

NewUpdateDataSourceRequest instantiates a new UpdateDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataSourceRequestWithDefaults

`func NewUpdateDataSourceRequestWithDefaults() *UpdateDataSourceRequest`

NewUpdateDataSourceRequestWithDefaults instantiates a new UpdateDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UpdateDataSourceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateDataSourceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateDataSourceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetHostname

`func (o *UpdateDataSourceRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateDataSourceRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateDataSourceRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateDataSourceRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *UpdateDataSourceRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateDataSourceRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateDataSourceRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateDataSourceRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTlsTerminationOverrideHost

`func (o *UpdateDataSourceRequest) GetTlsTerminationOverrideHost() string`

GetTlsTerminationOverrideHost returns the TlsTerminationOverrideHost field if non-nil, zero value otherwise.

### GetTlsTerminationOverrideHostOk

`func (o *UpdateDataSourceRequest) GetTlsTerminationOverrideHostOk() (*string, bool)`

GetTlsTerminationOverrideHostOk returns a tuple with the TlsTerminationOverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationOverrideHost

`func (o *UpdateDataSourceRequest) SetTlsTerminationOverrideHost(v string)`

SetTlsTerminationOverrideHost sets TlsTerminationOverrideHost field to given value.

### HasTlsTerminationOverrideHost

`func (o *UpdateDataSourceRequest) HasTlsTerminationOverrideHost() bool`

HasTlsTerminationOverrideHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


