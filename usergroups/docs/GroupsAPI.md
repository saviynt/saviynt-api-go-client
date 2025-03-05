# \GroupsAPI

All URIs are relative to *https://dev-scrum-intgn.saviyntcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRemoveUserFromUsergroup**](GroupsAPI.md#AddRemoveUserFromUsergroup) | **Post** /addRemoveUserFromUserGroup | Add/Remove User From UserGroup
[**CreateUpdateUsergroup**](GroupsAPI.md#CreateUpdateUsergroup) | **Post** /createUpdateUserGroup | Create/Update UserGroup
[**DeleteUsergroup**](GroupsAPI.md#DeleteUsergroup) | **Post** /deleteUserGroup | Delete UserGroup
[**GetListOfUsergroups**](GroupsAPI.md#GetListOfUsergroups) | **Post** /fetchUserGroup | Get List of UserGroups



## AddRemoveUserFromUsergroup

> AddRemoveUserFromUsergroup200Response AddRemoveUserFromUsergroup(ctx).ActionType(actionType).UserGroupname(userGroupname).Username(username).Execute()

Add/Remove User From UserGroup



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
	actionType := "actionType_example" // string |  (optional)
	userGroupname := "userGroupname_example" // string |  (optional)
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddRemoveUserFromUsergroup(context.Background()).ActionType(actionType).UserGroupname(userGroupname).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddRemoveUserFromUsergroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRemoveUserFromUsergroup`: AddRemoveUserFromUsergroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddRemoveUserFromUsergroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRemoveUserFromUsergroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** |  | 
 **userGroupname** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**AddRemoveUserFromUsergroup200Response**](AddRemoveUserFromUsergroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdateUsergroup

> CreateUpdateUsergroup200Response CreateUpdateUsergroup(ctx).CreateUpdateUsergroupRequest(createUpdateUsergroupRequest).Execute()

Create/Update UserGroup



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
	createUpdateUsergroupRequest := *openapiclient.NewCreateUpdateUsergroupRequest() // CreateUpdateUsergroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateUpdateUsergroup(context.Background()).CreateUpdateUsergroupRequest(createUpdateUsergroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateUpdateUsergroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpdateUsergroup`: CreateUpdateUsergroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateUpdateUsergroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateUsergroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateUsergroupRequest** | [**CreateUpdateUsergroupRequest**](CreateUpdateUsergroupRequest.md) |  | 

### Return type

[**CreateUpdateUsergroup200Response**](CreateUpdateUsergroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsergroup

> DeleteUsergroup200Response DeleteUsergroup(ctx).DeleteUsergroupRequest(deleteUsergroupRequest).Execute()

Delete UserGroup



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
	deleteUsergroupRequest := *openapiclient.NewDeleteUsergroupRequest() // DeleteUsergroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteUsergroup(context.Background()).DeleteUsergroupRequest(deleteUsergroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteUsergroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsergroup`: DeleteUsergroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteUsergroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsergroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUsergroupRequest** | [**DeleteUsergroupRequest**](DeleteUsergroupRequest.md) |  | 

### Return type

[**DeleteUsergroup200Response**](DeleteUsergroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfUsergroups

> GetListOfUsergroups200Response GetListOfUsergroups(ctx).Systemname(systemname).Offset(offset).Connectionname(connectionname).ConnectionType(connectionType).GetListOfUsergroupsRequest(getListOfUsergroupsRequest).Execute()

Get List of UserGroups



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
	systemname := "amigopod" // string |  (optional)
	offset := "0" // string |  (optional)
	connectionname := "CN_SERP_ECC_A53" // string |  (optional)
	connectionType := "SAP" // string |  (optional)
	getListOfUsergroupsRequest := *openapiclient.NewGetListOfUsergroupsRequest() // GetListOfUsergroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetListOfUsergroups(context.Background()).Systemname(systemname).Offset(offset).Connectionname(connectionname).ConnectionType(connectionType).GetListOfUsergroupsRequest(getListOfUsergroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetListOfUsergroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfUsergroups`: GetListOfUsergroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetListOfUsergroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfUsergroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemname** | **string** |  | 
 **offset** | **string** |  | 
 **connectionname** | **string** |  | 
 **connectionType** | **string** |  | 
 **getListOfUsergroupsRequest** | [**GetListOfUsergroupsRequest**](GetListOfUsergroupsRequest.md) |  | 

### Return type

[**GetListOfUsergroups200Response**](GetListOfUsergroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

