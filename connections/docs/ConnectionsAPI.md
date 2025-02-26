# \ConnectionsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionDetails**](ConnectionsAPI.md#GetConnectionDetails) | **Post** /ECM/api/v5/getConnectionDetails | Get Connection Details
[**GetConnections**](ConnectionsAPI.md#GetConnections) | **Post** /ECM/api/v5/getConnections | Get List of Connections



## GetConnectionDetails

> GetConnectionDetailsResponse GetConnectionDetails(ctx).GetConnectionDetailsRequest(getConnectionDetailsRequest).Execute()

Get Connection Details



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
	getConnectionDetailsRequest := *openapiclient.NewGetConnectionDetailsRequest() // GetConnectionDetailsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetConnectionDetails(context.Background()).GetConnectionDetailsRequest(getConnectionDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnectionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionDetails`: GetConnectionDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnectionDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConnectionDetailsRequest** | [**GetConnectionDetailsRequest**](GetConnectionDetailsRequest.md) |  | 

### Return type

[**GetConnectionDetailsResponse**](GetConnectionDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> GetConnectionsResponse GetConnections(ctx).GetConnectionsRequest(getConnectionsRequest).Execute()

Get List of Connections



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
	getConnectionsRequest := *openapiclient.NewGetConnectionsRequest() // GetConnectionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetConnections(context.Background()).GetConnectionsRequest(getConnectionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnections`: GetConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConnectionsRequest** | [**GetConnectionsRequest**](GetConnectionsRequest.md) |  | 

### Return type

[**GetConnectionsResponse**](GetConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

