# \DefaultAPI

All URIs are relative to *https://api.example.com/ECM/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](DefaultAPI.md#CreateEndpoint) | **Post** /createEndpoint | Create Endpoint
[**GetEndpoints**](DefaultAPI.md#GetEndpoints) | **Post** /getEndpoints | Get List of Endpoints
[**UpdateEndpoint**](DefaultAPI.md#UpdateEndpoint) | **Put** /updateEndpoint | Update Endpoint



## CreateEndpoint

> SuccessResponse CreateEndpoint(ctx).CreateEndpointRequest(createEndpointRequest).Execute()

Create Endpoint



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
	createEndpointRequest := *openapiclient.NewCreateEndpointRequest("Endpointname_example", "DisplayName_example", "Securitysystem_example") // CreateEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateEndpoint(context.Background()).CreateEndpointRequest(createEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoint`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointRequest** | [**CreateEndpointRequest**](CreateEndpointRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoints

> GetEndpointsResponse GetEndpoints(ctx).GetEndpointsRequest(getEndpointsRequest).Execute()

Get List of Endpoints



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
	getEndpointsRequest := *openapiclient.NewGetEndpointsRequest() // GetEndpointsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetEndpoints(context.Background()).GetEndpointsRequest(getEndpointsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoints`: GetEndpointsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEndpointsRequest** | [**GetEndpointsRequest**](GetEndpointsRequest.md) |  | 

### Return type

[**GetEndpointsResponse**](GetEndpointsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> SuccessResponse UpdateEndpoint(ctx).UpdateEndpointRequest(updateEndpointRequest).Execute()

Update Endpoint



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
	updateEndpointRequest := *openapiclient.NewUpdateEndpointRequest("Endpointname_example") // UpdateEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateEndpoint(context.Background()).UpdateEndpointRequest(updateEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoint`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEndpointRequest** | [**UpdateEndpointRequest**](UpdateEndpointRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

