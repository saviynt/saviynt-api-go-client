# \EntitlementsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpdateEntitlement**](EntitlementsAPI.md#CreateUpdateEntitlement) | **Post** /ECM/api/v5/createUpdateEntitlement | Create/Update Entitlement
[**GetChildEntitlements**](EntitlementsAPI.md#GetChildEntitlements) | **Get** /ECM/api/v5/getChildEntitlements | Get Child Entitlements
[**GetEntitlementValuesForEndpoint**](EntitlementsAPI.md#GetEntitlementValuesForEndpoint) | **Post** /ECM/api/v5/getEntitlementValuesForEndpoint | Get Entitlement Values For Endpoint
[**GetEntitlements**](EntitlementsAPI.md#GetEntitlements) | **Post** /ECM/api/v5/getEntitlements | Get Entitlements
[**GetListOfPrivilegesForEntitlementType**](EntitlementsAPI.md#GetListOfPrivilegesForEntitlementType) | **Post** /ECM/api/v5/getListofPrivileges | Get List of Privileges for Entitlement Type
[**RemoveEntitlementFromRole**](EntitlementsAPI.md#RemoveEntitlementFromRole) | **Post** /ECM/api/v5/removeEntitlementsFromRole | Remove Entitlement From Role



## CreateUpdateEntitlement

> CreateUpdateEntitlement200Response CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(createUpdateEntitlementRequest).Execute()

Create/Update Entitlement



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
	createUpdateEntitlementRequest := *openapiclient.NewCreateUpdateEntitlementRequest("Workday", "Accounts Receivable Specialist (Unconstrained)", "Security-Groups") // CreateUpdateEntitlementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateUpdateEntitlement(context.Background()).CreateUpdateEntitlementRequest(createUpdateEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateUpdateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpdateEntitlement`: CreateUpdateEntitlement200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateUpdateEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateEntitlementRequest** | [**CreateUpdateEntitlementRequest**](CreateUpdateEntitlementRequest.md) |  | 

### Return type

[**CreateUpdateEntitlement200Response**](CreateUpdateEntitlement200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChildEntitlements

> GetChildEntitlements200Response GetChildEntitlements(ctx).Endpointkey(endpointkey).Endpointname(endpointname).Entitlementtypekey(entitlementtypekey).Entitlementtypename(entitlementtypename).Parententitlementvalue(parententitlementvalue).Parententitlementvaluekey(parententitlementvaluekey).Entquery(entquery).Childentquery(childentquery).Entitlements2query(entitlements2query).Responsefields(responsefields).Max(max).Offset(offset).Execute()

Get Child Entitlements



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
	endpointkey := "1" // string |  (optional)
	endpointname := "AWS" // string |  (optional)
	entitlementtypekey := "54" // string |  (optional)
	entitlementtypename := "AWSSecurityGroup" // string |  (optional)
	parententitlementvalue := "sg-0303987b" // string |  (optional)
	parententitlementvaluekey := "189260" // string |  (optional)
	entquery := "pev.status = 1" // string |  (optional)
	childentquery := "cev.customproperty3 = '25'" // string |  (optional)
	entitlements2query := "e2.jobId is not null" // string |  (optional)
	responsefields := []string{"Inner_example"} // []string |  (optional)
	max := int32(5) // int32 |  (optional)
	offset := int32(0) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetChildEntitlements(context.Background()).Endpointkey(endpointkey).Endpointname(endpointname).Entitlementtypekey(entitlementtypekey).Entitlementtypename(entitlementtypename).Parententitlementvalue(parententitlementvalue).Parententitlementvaluekey(parententitlementvaluekey).Entquery(entquery).Childentquery(childentquery).Entitlements2query(entitlements2query).Responsefields(responsefields).Max(max).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetChildEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChildEntitlements`: GetChildEntitlements200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetChildEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChildEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointkey** | **string** |  | 
 **endpointname** | **string** |  | 
 **entitlementtypekey** | **string** |  | 
 **entitlementtypename** | **string** |  | 
 **parententitlementvalue** | **string** |  | 
 **parententitlementvaluekey** | **string** |  | 
 **entquery** | **string** |  | 
 **childentquery** | **string** |  | 
 **entitlements2query** | **string** |  | 
 **responsefields** | **[]string** |  | 
 **max** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**GetChildEntitlements200Response**](GetChildEntitlements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementValuesForEndpoint

> GetEntitlementValuesForEndpoint200Response GetEntitlementValuesForEndpoint(ctx).GetEntitlementValuesForEndpointRequest(getEntitlementValuesForEndpointRequest).Execute()

Get Entitlement Values For Endpoint



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
	getEntitlementValuesForEndpointRequest := *openapiclient.NewGetEntitlementValuesForEndpointRequest("Workday") // GetEntitlementValuesForEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementValuesForEndpoint(context.Background()).GetEntitlementValuesForEndpointRequest(getEntitlementValuesForEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementValuesForEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementValuesForEndpoint`: GetEntitlementValuesForEndpoint200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementValuesForEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementValuesForEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEntitlementValuesForEndpointRequest** | [**GetEntitlementValuesForEndpointRequest**](GetEntitlementValuesForEndpointRequest.md) |  | 

### Return type

[**GetEntitlementValuesForEndpoint200Response**](GetEntitlementValuesForEndpoint200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlements

> GetEntitlements200Response GetEntitlements(ctx).Username(username).EntitlementType(entitlementType).Endpoint(endpoint).RequestedObject(requestedObject).Max(max).Offset(offset).EntitlementResponseFields(entitlementResponseFields).UserResponseFields(userResponseFields).Userfiltercriteria(userfiltercriteria).Accountname(accountname).Entownerwithrank(entownerwithrank).Returnentitlementmap(returnentitlementmap).Exactmatch(exactmatch).Entitlementfiltercriteria(entitlementfiltercriteria).EntQuery(entQuery).Execute()

Get Entitlements



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
	username := "jasorodriguez" // string |  (optional)
	entitlementType := "Security-Groups" // string |  (optional)
	endpoint := "AWS" // string |  (optional)
	requestedObject := "users" // string |  (optional)
	max := int32(50) // int32 |  (optional)
	offset := int32(0) // int32 |  (optional)
	entitlementResponseFields := []string{"Inner_example"} // []string |  (optional)
	userResponseFields := []string{"Inner_example"} // []string |  (optional)
	userfiltercriteria := "status = 'active'" // string |  (optional)
	accountname := "jasorodriguez" // string |  (optional)
	entownerwithrank := true // bool |  (optional)
	returnentitlementmap := true // bool |  (optional) (default to false)
	exactmatch := true // bool |  (optional) (default to true)
	entitlementfiltercriteria := "ent.description = 'Desc'" // string |  (optional)
	entQuery := "ent.entitlement_value like '%i-03d58cd60fa90b9a9%'" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlements(context.Background()).Username(username).EntitlementType(entitlementType).Endpoint(endpoint).RequestedObject(requestedObject).Max(max).Offset(offset).EntitlementResponseFields(entitlementResponseFields).UserResponseFields(userResponseFields).Userfiltercriteria(userfiltercriteria).Accountname(accountname).Entownerwithrank(entownerwithrank).Returnentitlementmap(returnentitlementmap).Exactmatch(exactmatch).Entitlementfiltercriteria(entitlementfiltercriteria).EntQuery(entQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlements`: GetEntitlements200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **entitlementType** | **string** |  | 
 **endpoint** | **string** |  | 
 **requestedObject** | **string** |  | 
 **max** | **int32** |  | 
 **offset** | **int32** |  | 
 **entitlementResponseFields** | **[]string** |  | 
 **userResponseFields** | **[]string** |  | 
 **userfiltercriteria** | **string** |  | 
 **accountname** | **string** |  | 
 **entownerwithrank** | **bool** |  | 
 **returnentitlementmap** | **bool** |  | [default to false]
 **exactmatch** | **bool** |  | [default to true]
 **entitlementfiltercriteria** | **string** |  | 
 **entQuery** | **string** |  | 

### Return type

[**GetEntitlements200Response**](GetEntitlements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfPrivilegesForEntitlementType

> GetListOfPrivilegesForEntitlementType200Response GetListOfPrivilegesForEntitlementType(ctx).Endpoint(endpoint).Entitlementtype(entitlementtype).Max(max).Offset(offset).Execute()

Get List of Privileges for Entitlement Type



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
	endpoint := "endpoint_example" // string | 
	entitlementtype := "entitlementtype_example" // string |  (optional)
	max := "max_example" // string | Maximum number of records to return. (optional)
	offset := "offset_example" // string | Pagination offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetListOfPrivilegesForEntitlementType(context.Background()).Endpoint(endpoint).Entitlementtype(entitlementtype).Max(max).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetListOfPrivilegesForEntitlementType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfPrivilegesForEntitlementType`: GetListOfPrivilegesForEntitlementType200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetListOfPrivilegesForEntitlementType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfPrivilegesForEntitlementTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **string** |  | 
 **entitlementtype** | **string** |  | 
 **max** | **string** | Maximum number of records to return. | 
 **offset** | **string** | Pagination offset. | 

### Return type

[**GetListOfPrivilegesForEntitlementType200Response**](GetListOfPrivilegesForEntitlementType200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEntitlementFromRole

> RemoveEntitlementFromRole200Response RemoveEntitlementFromRole(ctx).RemoveEntitlementFromRoleRequest(removeEntitlementFromRoleRequest).Execute()

Remove Entitlement From Role



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
	removeEntitlementFromRoleRequest := *openapiclient.NewRemoveEntitlementFromRoleRequest() // RemoveEntitlementFromRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.RemoveEntitlementFromRole(context.Background()).RemoveEntitlementFromRoleRequest(removeEntitlementFromRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.RemoveEntitlementFromRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveEntitlementFromRole`: RemoveEntitlementFromRole200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.RemoveEntitlementFromRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEntitlementFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeEntitlementFromRoleRequest** | [**RemoveEntitlementFromRoleRequest**](RemoveEntitlementFromRoleRequest.md) |  | 

### Return type

[**RemoveEntitlementFromRole200Response**](RemoveEntitlementFromRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

