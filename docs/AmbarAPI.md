# \AmbarAPI

All URIs are relative to *https://euw1.api.ambar.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataDestination**](AmbarAPI.md#CreateDataDestination) | **Post** /destination | Create a DataDestination in your Ambar environment.
[**CreateDataSource**](AmbarAPI.md#CreateDataSource) | **Post** /source | Create a DataSource in your Ambar environment.
[**CreateFilter**](AmbarAPI.md#CreateFilter) | **Post** /filter | Create a Filter in your Ambar environment.
[**DeleteDataDestination**](AmbarAPI.md#DeleteDataDestination) | **Delete** /destination | Delete a DataDestination in your Ambar environment.
[**DeleteDataSource**](AmbarAPI.md#DeleteDataSource) | **Delete** /source | Delete a DataSource in your Ambar environment.
[**DeleteFilter**](AmbarAPI.md#DeleteFilter) | **Delete** /filter | Delete a Filter in your Ambar environment.
[**DescribeDataDestination**](AmbarAPI.md#DescribeDataDestination) | **Get** /destination | Describe a DataDestination in your Ambar environment.
[**DescribeDataSource**](AmbarAPI.md#DescribeDataSource) | **Get** /source | Describe a DataSource in your Ambar environment.
[**DescribeFilter**](AmbarAPI.md#DescribeFilter) | **Get** /filter | Describes a Filter in your Ambar environment.
[**ListResources**](AmbarAPI.md#ListResources) | **Get** /resource | List the Ambar resources in your Ambar environment.



## CreateDataDestination

> ResourceStateChangeResponse CreateDataDestination(ctx).CreateDataDestinationRequest(createDataDestinationRequest).Execute()

Create a DataDestination in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createDataDestinationRequest := *openapiclient.NewCreateDataDestinationRequest("DestinationEndpoint_example", "FilterId_example", "ProjectionName_example", "Password_example", "Username_example") // CreateDataDestinationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.CreateDataDestination(context.Background()).CreateDataDestinationRequest(createDataDestinationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.CreateDataDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataDestination`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.CreateDataDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataDestinationRequest** | [**CreateDataDestinationRequest**](CreateDataDestinationRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataSource

> ResourceStateChangeResponse CreateDataSource(ctx).CreateDataSourceRequest(createDataSourceRequest).Execute()

Create a DataSource in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createDataSourceRequest := *openapiclient.NewCreateDataSourceRequest(openapiclient.CreateDataSourceRequest_dataSourceConfig{MysqlDataSource: openapiclient.NewMysqlDataSource("Username_example", "Password_example", "Hostname_example", "HostPort_example", "DatabaseName_example", "TableName_example", "GloballyUniqueColumnName_example", "IncrementingColumnName_example", "PartitioningColumnName_example", "AdditionalColumns_example", "BinLogReplicationServerId_example", "TlsTerminationOverrideHost_example")}, "DataSourceType_example") // CreateDataSourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.CreateDataSource(context.Background()).CreateDataSourceRequest(createDataSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.CreateDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataSource`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.CreateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataSourceRequest** | [**CreateDataSourceRequest**](CreateDataSourceRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFilter

> ResourceStateChangeResponse CreateFilter(ctx).CreateFilterRequest(createFilterRequest).Execute()

Create a Filter in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createFilterRequest := *openapiclient.NewCreateFilterRequest("FilterContents_example") // CreateFilterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.CreateFilter(context.Background()).CreateFilterRequest(createFilterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.CreateFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFilter`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.CreateFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFilterRequest** | [**CreateFilterRequest**](CreateFilterRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataDestination

> ResourceStateChangeResponse DeleteDataDestination(ctx).DeleteResourceRequest(deleteResourceRequest).Execute()

Delete a DataDestination in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("ResourceId_example") // DeleteResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DeleteDataDestination(context.Background()).DeleteResourceRequest(deleteResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DeleteDataDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDataDestination`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DeleteDataDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteResourceRequest** | [**DeleteResourceRequest**](DeleteResourceRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataSource

> ResourceStateChangeResponse DeleteDataSource(ctx).DeleteResourceRequest(deleteResourceRequest).Execute()

Delete a DataSource in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("ResourceId_example") // DeleteResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DeleteDataSource(context.Background()).DeleteResourceRequest(deleteResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DeleteDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDataSource`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DeleteDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteResourceRequest** | [**DeleteResourceRequest**](DeleteResourceRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFilter

> ResourceStateChangeResponse DeleteFilter(ctx).DeleteResourceRequest(deleteResourceRequest).Execute()

Delete a Filter in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("ResourceId_example") // DeleteResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DeleteFilter(context.Background()).DeleteResourceRequest(deleteResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DeleteFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFilter`: ResourceStateChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DeleteFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteResourceRequest** | [**DeleteResourceRequest**](DeleteResourceRequest.md) |  | 

### Return type

[**ResourceStateChangeResponse**](ResourceStateChangeResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDataDestination

> DataDestination DescribeDataDestination(ctx).DescribeResourceRequest(describeResourceRequest).Execute()

Describe a DataDestination in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    describeResourceRequest := *openapiclient.NewDescribeResourceRequest("ResourceId_example") // DescribeResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DescribeDataDestination(context.Background()).DescribeResourceRequest(describeResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DescribeDataDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeDataDestination`: DataDestination
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DescribeDataDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDataDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeResourceRequest** | [**DescribeResourceRequest**](DescribeResourceRequest.md) |  | 

### Return type

[**DataDestination**](DataDestination.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDataSource

> DataSource DescribeDataSource(ctx).DescribeResourceRequest(describeResourceRequest).Execute()

Describe a DataSource in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    describeResourceRequest := *openapiclient.NewDescribeResourceRequest("ResourceId_example") // DescribeResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DescribeDataSource(context.Background()).DescribeResourceRequest(describeResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DescribeDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeDataSource`: DataSource
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DescribeDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeResourceRequest** | [**DescribeResourceRequest**](DescribeResourceRequest.md) |  | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeFilter

> Filter DescribeFilter(ctx).DescribeResourceRequest(describeResourceRequest).Execute()

Describes a Filter in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    describeResourceRequest := *openapiclient.NewDescribeResourceRequest("ResourceId_example") // DescribeResourceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.DescribeFilter(context.Background()).DescribeResourceRequest(describeResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.DescribeFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeFilter`: Filter
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.DescribeFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **describeResourceRequest** | [**DescribeResourceRequest**](DescribeResourceRequest.md) |  | 

### Return type

[**Filter**](Filter.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> []ResourceTypeDetails ListResources(ctx).Execute()

List the Ambar resources in your Ambar environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AmbarAPI.ListResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.ListResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResources`: []ResourceTypeDetails
    fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.ListResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


### Return type

[**[]ResourceTypeDetails**](ResourceTypeDetails.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
