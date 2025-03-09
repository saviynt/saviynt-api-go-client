package test

import (
	"context"
	"fmt"
	"testing"

	openapi "github.com/saviynt/saviynt-api-go-client/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_AccountsAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.NoError(t, err, "Failed to initialize API client")
	if !skipTests {
		require.NotNil(t, apiClient, "apiClient should not be nil")
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts should not be nil")
	}
	ctx := context.Background()
	t.Run("Test AccountsAPIService AssignAccountToEntitlement", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")
		req := apiClient.Accounts.AssignAccountToEntitlement(ctx).Securitysystem("AD_Rashid").Endpoint("AD_Rashid").Accountname("Rose Shukla").Entitlementtype("memberOf").Entitlementvalue("CN=adtestrequest123,OU=DocTeamOU,OU=SaviyntTeams,DC=saviyntlabs,DC=org").Startdate("03-09-2024")
		resp, httpRes, err := req.Execute()

		require.NoError(t, err, "Unexpected error in AssignAccountToEntitlement")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

	t.Run("Test AccountsAPIService AssignAccountToUser", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")
		req := apiClient.Accounts.AssignAccountToUser(ctx).Securitysystem("ADSystem").Endpoint("ADEndpoint").Accountname("Rose Shukla").Username("Rose Shukla")
		resp, httpRes, err := req.Execute()
		require.NoError(t, err, "Unexpected error in AssignAccountToUser")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	t.Run("Test AccountsAPIService CreateAccount", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")

		username := "Rose Shukla"
		comment := "This is created using go api client code"
		requestor := "admin"
		createAccountsReq := openapi.CreateAccountRequest{
			Securitysystem: "ADSystem",
			Endpoint:       "ADEndpoint",
			Name:           "Rose Shukla",
			Username:       &username,
			Comments:       &comment,
			Requestor:      &requestor,
		}
		resp, httpRes, err := apiClient.Accounts.CreateAccount(ctx).CreateAccountRequest(createAccountsReq).Execute()
		require.NoError(t, err, "Unexpected error in CreateAccount")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

	t.Run("Test AccountsAPIService ExportAccount", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "aapiClient.Accounts is nil, skipping test")

		exportReq := apiClient.Accounts.ExportAccount(ctx).
			Securitysystem("ADSystem").
			Endpoint("ADEndpoint")
		resp, httpRes, err := exportReq.Execute()
		require.NoError(t, err, "Unexpected error in ExportAccount")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

	t.Run("Test AccountsAPIService GetAccounts", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")
		endpoint := "ADSI"
		getAccountsReq := openapi.GetAccountsRequest{
			Endpoint: &endpoint,
		}
		resp, httpRes, err := apiClient.Accounts.GetAccounts(ctx).
			GetAccountsRequest(getAccountsReq).
			Execute()

		require.NoError(t, err, "Unexpected error in GetAccounts")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

	t.Run("Test AccountsAPIService RemoveAccountToEntitlement", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")

		req := apiClient.Accounts.RemoveAccountToEntitlement(ctx).Securitysystem("AD_Rashid").Endpoint("AD_Rashid").Accountname("Rose Shukla").Entitlementtype("memberOf").Entitlementvalue("CN=adtestrequest123,OU=DocTeamOU,OU=SaviyntTeams,DC=saviyntlabs,DC=org")
		resp, httpRes, err := req.Execute()

		require.NoError(t, err, "Unexpected error in RemoveAccountToEntitlement")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

	t.Run("Test AccountsAPIService UpdateAccount", func(t *testing.T) {

		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		require.NotNil(t, apiClient.Accounts, "apiClient.Accounts is nil, skipping test")
		comment := "This is Updated using go api client code"
		updateAccountReq := openapi.UpdateAccountRequest{
			Securitysystem: "ADSystem",
			Endpoint:       "ADEndpoint",
			Name:           "Rose Shukla",
			Comments:       &comment,
		}
		resp, httpRes, err := apiClient.Accounts.UpdateAccount(ctx).UpdateAccountRequest(updateAccountReq).Execute()

		require.NoError(t, err, "Unexpected error in RemoveAccountToEntitlement")
		require.NotNil(t, httpRes, "httpRes should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")

	})

}
