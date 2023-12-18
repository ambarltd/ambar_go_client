# Go API client for Ambar

Details about communicating with Ambar.cloud public endpoints. Supported HTTP rest endpoints and their  request and response details.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2023-12-01
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://ambar.cloud/developer/docs/help](https://ambar.cloud/developer/docs/help)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import Ambar "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), Ambar.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), Ambar.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), Ambar.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), Ambar.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://euw1.api.ambar.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AmbarAPI* | [**CreateDataDestination**](docs/AmbarAPI.md#createdatadestination) | **Post** /destination | Create a DataDestination in your Ambar environment.
*AmbarAPI* | [**CreateDataSource**](docs/AmbarAPI.md#createdatasource) | **Post** /source | Create a DataSource in your Ambar environment.
*AmbarAPI* | [**CreateFilter**](docs/AmbarAPI.md#createfilter) | **Post** /filter | Create a Filter in your Ambar environment.
*AmbarAPI* | [**DeleteDataDestination**](docs/AmbarAPI.md#deletedatadestination) | **Delete** /destination | Delete a DataDestination in your Ambar environment.
*AmbarAPI* | [**DeleteDataSource**](docs/AmbarAPI.md#deletedatasource) | **Delete** /source | Delete a DataSource in your Ambar environment.
*AmbarAPI* | [**DeleteFilter**](docs/AmbarAPI.md#deletefilter) | **Delete** /filter | Delete a Filter in your Ambar environment.
*AmbarAPI* | [**DescribeDataDestination**](docs/AmbarAPI.md#describedatadestination) | **Get** /destination | Describe a DataDestination in your Ambar environment.
*AmbarAPI* | [**DescribeDataSource**](docs/AmbarAPI.md#describedatasource) | **Get** /source | Describe a DataSource in your Ambar environment.
*AmbarAPI* | [**DescribeFilter**](docs/AmbarAPI.md#describefilter) | **Get** /filter | Describes a Filter in your Ambar environment.
*AmbarAPI* | [**ListResources**](docs/AmbarAPI.md#listresources) | **Get** /resource | List the Ambar resources in your Ambar environment.


## Documentation For Models

 - [AmbarServiceException](docs/AmbarServiceException.md)
 - [CreateDataDestinationRequest](docs/CreateDataDestinationRequest.md)
 - [CreateDataSourceRequest](docs/CreateDataSourceRequest.md)
 - [CreateDataSourceRequestDataSourceConfig](docs/CreateDataSourceRequestDataSourceConfig.md)
 - [CreateFilter400Response](docs/CreateFilter400Response.md)
 - [CreateFilterRequest](docs/CreateFilterRequest.md)
 - [DataDestination](docs/DataDestination.md)
 - [DataSource](docs/DataSource.md)
 - [DeleteFilter400Response](docs/DeleteFilter400Response.md)
 - [DeleteResourceRequest](docs/DeleteResourceRequest.md)
 - [DescribeResourceRequest](docs/DescribeResourceRequest.md)
 - [Filter](docs/Filter.md)
 - [InvalidFilterException](docs/InvalidFilterException.md)
 - [InvalidParameterException](docs/InvalidParameterException.md)
 - [ListResourcesRequest](docs/ListResourcesRequest.md)
 - [ListResourcesResponse](docs/ListResourcesResponse.md)
 - [PostgresDataSource](docs/PostgresDataSource.md)
 - [ResourceDetails](docs/ResourceDetails.md)
 - [ResourceInvalidStateException](docs/ResourceInvalidStateException.md)
 - [ResourceNotFoundException](docs/ResourceNotFoundException.md)
 - [ResourceStateChangeResponse](docs/ResourceStateChangeResponse.md)
 - [ResourceTypeDetails](docs/ResourceTypeDetails.md)
 - [ResourceTypeDetailsLimit](docs/ResourceTypeDetailsLimit.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_key

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"x-api-key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

contact@ambar.cloud

