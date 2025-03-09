# \OrganizationsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationsAPI.md#CreateOrganization) | **Post** /ECM/api/v5/createOrganization | Create Organization
[**DeleteOrganization**](OrganizationsAPI.md#DeleteOrganization) | **Post** /ECM/api/v5/deleteOrganization | Delete Organization
[**GetOrganization**](OrganizationsAPI.md#GetOrganization) | **Get** /ECM/api/v5/getOrganization | Get Organization
[**GetOrganizationUserDetails**](OrganizationsAPI.md#GetOrganizationUserDetails) | **Post** /ECM/api/v5/getOrganizationUserDetails | Get User Organization
[**UpdateOrganization**](OrganizationsAPI.md#UpdateOrganization) | **Put** /ECM/api/v5/updateOrganization | Update Organization
[**UpdateOrganizationUsers**](OrganizationsAPI.md#UpdateOrganizationUsers) | **Post** /ECM/api/v5/updateOrganizationUsers | Update Organization Users



## CreateOrganization

> CreateOrganization200Response CreateOrganization(ctx).CreateOrganizationRequest(createOrganizationRequest).Execute()

Create Organization



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
	createOrganizationRequest := *openapiclient.NewCreateOrganizationRequest("TestOrg5", "awsadmin") // CreateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganization(context.Background()).CreateOrganizationRequest(createOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: CreateOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) |  | 

### Return type

[**CreateOrganization200Response**](CreateOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization200Response DeleteOrganization(ctx).DeleteOrganizationRequest(deleteOrganizationRequest).Execute()

Delete Organization



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
	deleteOrganizationRequest := *openapiclient.NewDeleteOrganizationRequest() // DeleteOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DeleteOrganization(context.Background()).DeleteOrganizationRequest(deleteOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganization`: DeleteOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOrganizationRequest** | [**DeleteOrganizationRequest**](DeleteOrganizationRequest.md) |  | 

### Return type

[**DeleteOrganization200Response**](DeleteOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> GetOrganization200Response GetOrganization(ctx).GetOrganizationRequest(getOrganizationRequest).Execute()

Get Organization



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
	getOrganizationRequest := *openapiclient.NewGetOrganizationRequest() // GetOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganization(context.Background()).GetOrganizationRequest(getOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: GetOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrganizationRequest** | [**GetOrganizationRequest**](GetOrganizationRequest.md) |  | 

### Return type

[**GetOrganization200Response**](GetOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationUserDetails

> GetOrganizationUserDetails200Response GetOrganizationUserDetails(ctx).GetOrganizationUserDetailsRequest(getOrganizationUserDetailsRequest).Execute()

Get User Organization



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
	getOrganizationUserDetailsRequest := *openapiclient.NewGetOrganizationUserDetailsRequest("smith") // GetOrganizationUserDetailsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationUserDetails(context.Background()).GetOrganizationUserDetailsRequest(getOrganizationUserDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationUserDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationUserDetails`: GetOrganizationUserDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationUserDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrganizationUserDetailsRequest** | [**GetOrganizationUserDetailsRequest**](GetOrganizationUserDetailsRequest.md) |  | 

### Return type

[**GetOrganizationUserDetails200Response**](GetOrganizationUserDetails200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> CreateOrganization200Response UpdateOrganization(ctx).UpdateOrganizationRequest(updateOrganizationRequest).Execute()

Update Organization



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
	updateOrganizationRequest := *openapiclient.NewUpdateOrganizationRequest("TestOrg5", "awsadmin") // UpdateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateOrganization(context.Background()).UpdateOrganizationRequest(updateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: CreateOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationRequest** | [**UpdateOrganizationRequest**](UpdateOrganizationRequest.md) |  | 

### Return type

[**CreateOrganization200Response**](CreateOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationUsers

> UpdateOrganizationUsers200Response UpdateOrganizationUsers(ctx).UpdateOrganizationUsersRequest(updateOrganizationUsersRequest).Execute()

Update Organization Users



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
	updateOrganizationUsersRequest := *openapiclient.NewUpdateOrganizationUsersRequest() // UpdateOrganizationUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateOrganizationUsers(context.Background()).UpdateOrganizationUsersRequest(updateOrganizationUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganizationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationUsers`: UpdateOrganizationUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganizationUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationUsersRequest** | [**UpdateOrganizationUsersRequest**](UpdateOrganizationUsersRequest.md) |  | 

### Return type

[**UpdateOrganizationUsers200Response**](UpdateOrganizationUsers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

