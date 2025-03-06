# \EndpointsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](EndpointsAPI.md#CreateEndpoint) | **Post** /ECM/api/v5/createEndpoint | Create Endpoint
[**GetEndpoints**](EndpointsAPI.md#GetEndpoints) | **Post** /ECM/api/v5/getEndpoints | Get List of Endpoints
[**UpdateEndpoint**](EndpointsAPI.md#UpdateEndpoint) | **Put** /ECM/api/v5/updateEndpoint | Update Endpoint



## CreateEndpoint

> CreateEndpointResponse CreateEndpoint(ctx).CreateEndpointRequest(createEndpointRequest).Execute()

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
	createEndpointRequest := *openapiclient.NewCreateEndpointRequest("DisplayName_example", "Endpointname_example", "Securitysystem_example") // CreateEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CreateEndpoint(context.Background()).CreateEndpointRequest(createEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoint`: CreateEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointRequest** | [**CreateEndpointRequest**](CreateEndpointRequest.md) |  | 

### Return type

[**CreateEndpointResponse**](CreateEndpointResponse.md)

### Authorization

No authorization required

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
	resp, r, err := apiClient.EndpointsAPI.GetEndpoints(context.Background()).GetEndpointsRequest(getEndpointsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoints`: GetEndpointsResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpoints`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> UpdateEndpointResponse UpdateEndpoint(ctx).UpdateEndpointRequest(updateEndpointRequest).Execute()

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
	resp, r, err := apiClient.EndpointsAPI.UpdateEndpoint(context.Background()).UpdateEndpointRequest(updateEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoint`: UpdateEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.UpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEndpointRequest** | [**UpdateEndpointRequest**](UpdateEndpointRequest.md) |  | 

### Return type

[**UpdateEndpointResponse**](UpdateEndpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

