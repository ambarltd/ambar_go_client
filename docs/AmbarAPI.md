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
[**UpdateDataDestination**](AmbarAPI.md#UpdateDataDestination) | **Put** /destination | Updates a DataDestination in your Ambar environment.
[**UpdateDataDestinationCredentials**](AmbarAPI.md#UpdateDataDestinationCredentials) | **Patch** /destination | Update the credentials associated with a DataDestination in your Ambar environment.
[**UpdateDataSource**](AmbarAPI.md#UpdateDataSource) | **Put** /source | Updates a DataSource in your Ambar environment.
[**UpdateDataSourceCredentials**](AmbarAPI.md#UpdateDataSourceCredentials) | **Patch** /source | Update the credentials associated with a DataSource in your Ambar environment.



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
	createDataDestinationRequest := *openapiclient.NewCreateDataDestinationRequest("DestinationEndpoint_example", []string{"FilterIds_example"}, "Password_example", "Username_example") // CreateDataDestinationRequest | 

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
	createDataSourceRequest := *openapiclient.NewCreateDataSourceRequest("DataSourceType_example", map[string]string{"key": "Inner_example"}) // CreateDataSourceRequest | 

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
	createFilterRequest := *openapiclient.NewCreateFilterRequest("FilterContents_example", "DataSourceId_example") // CreateFilterRequest | 

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
	deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("AMBAR-1fG45k7090") // DeleteResourceRequest | 

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
	deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("AMBAR-1fG45k7090") // DeleteResourceRequest | 

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
	deleteResourceRequest := *openapiclient.NewDeleteResourceRequest("AMBAR-1fG45k7090") // DeleteResourceRequest | 

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

> ListResourcesResponse ListResources(ctx).ListResourcesRequest(listResourcesRequest).Execute()

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
	listResourcesRequest := *openapiclient.NewListResourcesRequest() // ListResourcesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AmbarAPI.ListResources(context.Background()).ListResourcesRequest(listResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.ListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResources`: ListResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.ListResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listResourcesRequest** | [**ListResourcesRequest**](ListResourcesRequest.md) |  | 

### Return type

[**ListResourcesResponse**](ListResourcesResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataDestination

> ResourceStateChangeResponse UpdateDataDestination(ctx).UpdateDataDestinationRequest(updateDataDestinationRequest).Execute()

Updates a DataDestination in your Ambar environment.



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
	updateDataDestinationRequest := *openapiclient.NewUpdateDataDestinationRequest("DestinationEndpoint_example", "ResourceId_example") // UpdateDataDestinationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AmbarAPI.UpdateDataDestination(context.Background()).UpdateDataDestinationRequest(updateDataDestinationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.UpdateDataDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataDestination`: ResourceStateChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.UpdateDataDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataDestinationRequest** | [**UpdateDataDestinationRequest**](UpdateDataDestinationRequest.md) |  | 

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


## UpdateDataDestinationCredentials

> ResourceStateChangeResponse UpdateDataDestinationCredentials(ctx).UpdateResourceCredentialsRequest(updateResourceCredentialsRequest).Execute()

Update the credentials associated with a DataDestination in your Ambar environment.



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
	updateResourceCredentialsRequest := *openapiclient.NewUpdateResourceCredentialsRequest("ResourceId_example", "Password_example", "Username_example") // UpdateResourceCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AmbarAPI.UpdateDataDestinationCredentials(context.Background()).UpdateResourceCredentialsRequest(updateResourceCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.UpdateDataDestinationCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataDestinationCredentials`: ResourceStateChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.UpdateDataDestinationCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataDestinationCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateResourceCredentialsRequest** | [**UpdateResourceCredentialsRequest**](UpdateResourceCredentialsRequest.md) |  | 

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


## UpdateDataSource

> ResourceStateChangeResponse UpdateDataSource(ctx).UpdateDataSourceRequest(updateDataSourceRequest).Execute()

Updates a DataSource in your Ambar environment.



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
	updateDataSourceRequest := *openapiclient.NewUpdateDataSourceRequest("AMBAR-1fG45k7090") // UpdateDataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AmbarAPI.UpdateDataSource(context.Background()).UpdateDataSourceRequest(updateDataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.UpdateDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataSource`: ResourceStateChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.UpdateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataSourceRequest** | [**UpdateDataSourceRequest**](UpdateDataSourceRequest.md) |  | 

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


## UpdateDataSourceCredentials

> ResourceStateChangeResponse UpdateDataSourceCredentials(ctx).UpdateResourceCredentialsRequest(updateResourceCredentialsRequest).Execute()

Update the credentials associated with a DataSource in your Ambar environment.



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
	updateResourceCredentialsRequest := *openapiclient.NewUpdateResourceCredentialsRequest("ResourceId_example", "Password_example", "Username_example") // UpdateResourceCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AmbarAPI.UpdateDataSourceCredentials(context.Background()).UpdateResourceCredentialsRequest(updateResourceCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AmbarAPI.UpdateDataSourceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataSourceCredentials`: ResourceStateChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `AmbarAPI.UpdateDataSourceCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataSourceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateResourceCredentialsRequest** | [**UpdateResourceCredentialsRequest**](UpdateResourceCredentialsRequest.md) |  | 

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

