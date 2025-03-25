# \ConnectionsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestConnection**](ConnectionsAPI.md#TestConnection) | **Post** /ECM/api/v5/testConnection | Create a connection



## TestConnection

> TestConnection200Response TestConnection(ctx).TestConnectionRequest(testConnectionRequest).Execute()

Create a connection

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
	testConnectionRequest := openapiclient.testConnection_request{ADConnector: openapiclient.NewADConnector("PASSWORD_example", "AD")} // TestConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.TestConnection(context.Background()).TestConnectionRequest(testConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.TestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestConnection`: TestConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.TestConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testConnectionRequest** | [**TestConnectionRequest**](TestConnectionRequest.md) |  | 

### Return type

[**TestConnection200Response**](TestConnection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

