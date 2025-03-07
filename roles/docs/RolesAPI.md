# \RolesAPI

All URIs are relative to *https://dev-scrum-intgn.saviyntcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRole**](RolesAPI.md#AddRole) | **Post** /addrole | Add Role
[**GetFirefighterRoleDetails**](RolesAPI.md#GetFirefighterRoleDetails) | **Post** /getFireFighterRoles | Get FireFighter Role Details
[**GetRoleDetailsForUser**](RolesAPI.md#GetRoleDetailsForUser) | **Post** /getRoles | Get Role Details for user
[**RemoveRole**](RolesAPI.md#RemoveRole) | **Post** /removerole | Remove Role



## AddRole

> AddRole200Response AddRole(ctx).AddRoleRequest(addRoleRequest).Execute()

Add Role



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
	addRoleRequest := *openapiclient.NewAddRoleRequest() // AddRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AddRole(context.Background()).AddRoleRequest(addRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRole`: AddRole200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AddRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRoleRequest** | [**AddRoleRequest**](AddRoleRequest.md) |  | 

### Return type

[**AddRole200Response**](AddRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirefighterRoleDetails

> []GetFirefighterRoleDetails200ResponseInner GetFirefighterRoleDetails(ctx).Execute()

Get FireFighter Role Details



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
	resp, r, err := apiClient.RolesAPI.GetFirefighterRoleDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetFirefighterRoleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirefighterRoleDetails`: []GetFirefighterRoleDetails200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetFirefighterRoleDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirefighterRoleDetailsRequest struct via the builder pattern


### Return type

[**[]GetFirefighterRoleDetails200ResponseInner**](GetFirefighterRoleDetails200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleDetailsForUser

> GetRoleDetailsForUser200Response GetRoleDetailsForUser(ctx).Max(max).Offset(offset).RequestedObject(requestedObject).RoleQuery(roleQuery).Roletype(roletype).Username(username).Execute()

Get Role Details for user



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
	max := "max_example" // string |  (optional)
	offset := "offset_example" // string |  (optional)
	requestedObject := "requestedObject_example" // string |  (optional)
	roleQuery := "roleQuery_example" // string |  (optional)
	roletype := "roletype_example" // string |  (optional)
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoleDetailsForUser(context.Background()).Max(max).Offset(offset).RequestedObject(requestedObject).RoleQuery(roleQuery).Roletype(roletype).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoleDetailsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleDetailsForUser`: GetRoleDetailsForUser200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoleDetailsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleDetailsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **string** |  | 
 **offset** | **string** |  | 
 **requestedObject** | **string** |  | 
 **roleQuery** | **string** |  | 
 **roletype** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**GetRoleDetailsForUser200Response**](GetRoleDetailsForUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRole

> RemoveRole200Response RemoveRole(ctx).RemoveRoleRequest(removeRoleRequest).Execute()

Remove Role



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
	removeRoleRequest := *openapiclient.NewRemoveRoleRequest() // RemoveRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.RemoveRole(context.Background()).RemoveRoleRequest(removeRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.RemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRole`: RemoveRole200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.RemoveRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeRoleRequest** | [**RemoveRoleRequest**](RemoveRoleRequest.md) |  | 

### Return type

[**RemoveRole200Response**](RemoveRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

