# \AccountsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAccountToEntitlement**](AccountsAPI.md#AssignAccountToEntitlement) | **Post** /ECM/api/v5/assignAccountToEntitlement | Assign Entitlement to Account
[**AssignAccountToUser**](AccountsAPI.md#AssignAccountToUser) | **Post** /ECM/api/v5/assignAccountToUser | Assign Account to User
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /ECM/api/v5/createAccount | Create Account
[**ExportAccount**](AccountsAPI.md#ExportAccount) | **Post** /ECM/api/v5/exportAccount | Export Account
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Post** /ECM/api/v5/getAccounts | Get Account Details
[**RemoveAccountToEntitlement**](AccountsAPI.md#RemoveAccountToEntitlement) | **Post** /ECM/api/v5/removeAccountToEntitlement | Remove Entitlement from Account
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Post** /ECM/api/v5/updateAccount | Update Account



## AssignAccountToEntitlement

> CreateAccount200Response AssignAccountToEntitlement(ctx).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Entitlementtype(entitlementtype).Entitlementvalue(entitlementvalue).Startdate(startdate).Execute()

Assign Entitlement to Account



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
	securitysystem := "securitysystem_example" // string | Name of the Security System (not the display name).
	endpoint := "endpoint_example" // string | Name of the Endpoint corresponding to the Security System (not the display name).
	accountname := "accountname_example" // string | Account name to which entitlements should be provisioned.
	entitlementtype := "entitlementtype_example" // string | Entitlement type (e.g., AD Groups, EBS Responsibilities, SAP Roles).
	entitlementvalue := "entitlementvalue_example" // string | Name(s) of the entitlement(s) to be provisioned.
	startdate := "startdate_example" // string | Start date for the access period in MM-DD-YYYY format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AssignAccountToEntitlement(context.Background()).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Entitlementtype(entitlementtype).Entitlementvalue(entitlementvalue).Startdate(startdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AssignAccountToEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignAccountToEntitlement`: CreateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AssignAccountToEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignAccountToEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitysystem** | **string** | Name of the Security System (not the display name). | 
 **endpoint** | **string** | Name of the Endpoint corresponding to the Security System (not the display name). | 
 **accountname** | **string** | Account name to which entitlements should be provisioned. | 
 **entitlementtype** | **string** | Entitlement type (e.g., AD Groups, EBS Responsibilities, SAP Roles). | 
 **entitlementvalue** | **string** | Name(s) of the entitlement(s) to be provisioned. | 
 **startdate** | **string** | Start date for the access period in MM-DD-YYYY format. | 

### Return type

[**CreateAccount200Response**](CreateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignAccountToUser

> CreateAccount200Response AssignAccountToUser(ctx).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Username(username).Execute()

Assign Account to User



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
	securitysystem := "securitysystem_example" // string | Name of the Security System (not the display name).
	endpoint := "endpoint_example" // string | Name of the Endpoint corresponding to the Security System (not the display name).
	accountname := "accountname_example" // string | Account name for the provisioned account.
	username := "username_example" // string | Username of the user to whom the account is to be provisioned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AssignAccountToUser(context.Background()).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AssignAccountToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignAccountToUser`: CreateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AssignAccountToUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignAccountToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitysystem** | **string** | Name of the Security System (not the display name). | 
 **endpoint** | **string** | Name of the Endpoint corresponding to the Security System (not the display name). | 
 **accountname** | **string** | Account name for the provisioned account. | 
 **username** | **string** | Username of the user to whom the account is to be provisioned. | 

### Return type

[**CreateAccount200Response**](CreateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> CreateAccount200Response CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()

Create Account



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
	createAccountRequest := *openapiclient.NewCreateAccountRequest("System1", "System1", "johnWS") // CreateAccountRequest | Request payload for creating an account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: CreateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) | Request payload for creating an account. | 

### Return type

[**CreateAccount200Response**](CreateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAccount

> ExportAccount200Response ExportAccount(ctx).Securitysystem(securitysystem).Endpoint(endpoint).Username(username).Max(max).Offset(offset).AccountQuery(accountQuery).Advsearchcriteria(advsearchcriteria).Execute()

Export Account



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
	securitysystem := "securitysystem_example" // string | The security system (not the display name).
	endpoint := "endpoint_example" // string | The endpoint (not the display name).
	username := "username_example" // string | Optional username filter. (optional)
	max := "max_example" // string | Maximum number of records to return. (optional)
	offset := "offset_example" // string | Pagination offset. (optional)
	accountQuery := "accountQuery_example" // string | A query string to filter accounts. (optional)
	advsearchcriteria := *openapiclient.NewExportAccountRequestAdvsearchcriteria() // ExportAccountRequestAdvsearchcriteria |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ExportAccount(context.Background()).Securitysystem(securitysystem).Endpoint(endpoint).Username(username).Max(max).Offset(offset).AccountQuery(accountQuery).Advsearchcriteria(advsearchcriteria).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ExportAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAccount`: ExportAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ExportAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitysystem** | **string** | The security system (not the display name). | 
 **endpoint** | **string** | The endpoint (not the display name). | 
 **username** | **string** | Optional username filter. | 
 **max** | **string** | Maximum number of records to return. | 
 **offset** | **string** | Pagination offset. | 
 **accountQuery** | **string** | A query string to filter accounts. | 
 **advsearchcriteria** | [**ExportAccountRequestAdvsearchcriteria**](ExportAccountRequestAdvsearchcriteria.md) |  | 

### Return type

[**ExportAccount200Response**](ExportAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> GetAccounts200Response GetAccounts(ctx).GetAccountsRequest(getAccountsRequest).Execute()

Get Account Details



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
	getAccountsRequest := *openapiclient.NewGetAccountsRequest() // GetAccountsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).GetAccountsRequest(getAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: GetAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAccountsRequest** | [**GetAccountsRequest**](GetAccountsRequest.md) |  | 

### Return type

[**GetAccounts200Response**](GetAccounts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAccountToEntitlement

> CreateAccount200Response RemoveAccountToEntitlement(ctx).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Entitlementtype(entitlementtype).Entitlementvalue(entitlementvalue).Execute()

Remove Entitlement from Account



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
	securitysystem := "securitysystem_example" // string | Name of the Security System (not the display name).
	endpoint := "endpoint_example" // string | Name of the Endpoint corresponding to the Security System (not the display name).
	accountname := "accountname_example" // string | Account name from which entitlements should be deprovisioned.
	entitlementtype := "entitlementtype_example" // string | Entitlement type (e.g., AD Groups, EBS Responsibilities, SAP Roles).
	entitlementvalue := "entitlementvalue_example" // string | Names of the entitlement(s) to be deprovisioned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.RemoveAccountToEntitlement(context.Background()).Securitysystem(securitysystem).Endpoint(endpoint).Accountname(accountname).Entitlementtype(entitlementtype).Entitlementvalue(entitlementvalue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.RemoveAccountToEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAccountToEntitlement`: CreateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.RemoveAccountToEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAccountToEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitysystem** | **string** | Name of the Security System (not the display name). | 
 **endpoint** | **string** | Name of the Endpoint corresponding to the Security System (not the display name). | 
 **accountname** | **string** | Account name from which entitlements should be deprovisioned. | 
 **entitlementtype** | **string** | Entitlement type (e.g., AD Groups, EBS Responsibilities, SAP Roles). | 
 **entitlementvalue** | **string** | Names of the entitlement(s) to be deprovisioned. | 

### Return type

[**CreateAccount200Response**](CreateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount200Response UpdateAccount(ctx).UpdateAccountRequest(updateAccountRequest).Execute()

Update Account



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
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest("System1", "System1", "johnWS") // UpdateAccountRequest | Request payload for updating an account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateAccount(context.Background()).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: UpdateAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) | Request payload for updating an account. | 

### Return type

[**UpdateAccount200Response**](UpdateAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

