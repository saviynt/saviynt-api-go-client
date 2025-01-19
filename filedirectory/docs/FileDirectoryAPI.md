# \FileDirectoryAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadNewFile**](FileDirectoryAPI.md#UploadNewFile) | **Post** /ECM/api/v5/uploadSchemaFile | Upload File



## UploadNewFile

> UploadSchemaFileResponse UploadNewFile(ctx).UploadSchemaFileRequest(uploadSchemaFileRequest).Execute()

Upload File



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
	uploadSchemaFileRequest := *openapiclient.NewUploadSchemaFileRequest("TODO", "PathLocation_example") // UploadSchemaFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileDirectoryAPI.UploadNewFile(context.Background()).UploadSchemaFileRequest(uploadSchemaFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileDirectoryAPI.UploadNewFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadNewFile`: UploadSchemaFileResponse
	fmt.Fprintf(os.Stdout, "Response from `FileDirectoryAPI.UploadNewFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadNewFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadSchemaFileRequest** | [**UploadSchemaFileRequest**](UploadSchemaFileRequest.md) |  | 

### Return type

[**UploadSchemaFileResponse**](UploadSchemaFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

