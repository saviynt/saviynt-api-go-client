# \ConnectionsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateConnection**](ConnectionsAPI.md#CreateOrUpdateConnection) | **Post** /ECM/api/v5/testConnection | Create or Update Connection
[**DeleteConnection**](ConnectionsAPI.md#DeleteConnection) | **Delete** /ECM/api/v5/deleteConnection | 
[**GetConnectionDetails**](ConnectionsAPI.md#GetConnectionDetails) | **Post** /ECM/api/v5/getConnectionDetails | Get Connection Details
[**GetConnections**](ConnectionsAPI.md#GetConnections) | **Post** /ECM/api/v5/getConnections | Get List of Connections



## CreateOrUpdateConnection

> CreateOrUpdateConnectionResponse CreateOrUpdateConnection(ctx).CreateOrUpdateConnectionRequest(createOrUpdateConnectionRequest).Execute()

Create or Update Connection



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
	createOrUpdateConnectionRequest := *openapiclient.NewCreateOrUpdateConnectionRequest("Connectiontype_example") // CreateOrUpdateConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateOrUpdateConnection(context.Background()).CreateOrUpdateConnectionRequest(createOrUpdateConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateOrUpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateConnection`: CreateOrUpdateConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateOrUpdateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateConnectionRequest** | [**CreateOrUpdateConnectionRequest**](CreateOrUpdateConnectionRequest.md) |  | 

### Return type

[**CreateOrUpdateConnectionResponse**](CreateOrUpdateConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnectionResponse DeleteConnection(ctx).DeleteConnectionRequest(deleteConnectionRequest).Execute()





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
	deleteConnectionRequest := *openapiclient.NewDeleteConnectionRequest() // DeleteConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.DeleteConnection(context.Background()).DeleteConnectionRequest(deleteConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnection`: DeleteConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.DeleteConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteConnectionRequest** | [**DeleteConnectionRequest**](DeleteConnectionRequest.md) |  | 

### Return type

[**DeleteConnectionResponse**](DeleteConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

