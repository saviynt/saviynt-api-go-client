/*
Entitlements Management API

Testing Entitlements Service

This test file automatically reads credentials from an environment variable
(e.g. SAVIYNT_TEST_CREDENTIALS) so that the client can retrieve an access token.
*/

package test

import (
	"context"
	"fmt"
	"testing"

	entitlements "github.com/saviynt/saviynt-api-go-client/entitlements"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PtrString returns a pointer to a string.
func PtrString(s string) *string {
	return &s
}

func Test_EntitlementsService(t *testing.T) {
	// client() is a shared helper defined in test/client.go that returns the configured API client.
	apiClient, _, _, _, err := client()
	require.Nil(t, err)

	ctx := context.Background()

	t.Run("Test GetEntitlements without parameters", func(t *testing.T) {
		resp, httpRes, err := apiClient.Entitlements.GetEntitlements(ctx).Execute()
		require.NoError(t, err, "Unexpected error in GetEntitlements")
		fmt.Printf("HTTP Status: %d\n", httpRes.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	// ─── Test: GetChildEntitlements ──────────────────────────────────────────────
	t.Run("Test GetChildEntitlements", func(t *testing.T) {
		resp, httpRes, err := apiClient.Entitlements.GetChildEntitlements(ctx).
			Endpointname("AD_Rashid").
			Max(5).
			Offset(0).
			Execute()
		require.NoError(t, err, "Unexpected error in GetChildEntitlements")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	// ─── Test: CreateUpdateEntitlement ─────────────────────────────────────────────
	t.Run("Test CreateUpdateEntitlement", func(t *testing.T) {
		newValue := "CN=adtestrequest123,OU=DocTeamOU,OU=SaviyntTeams,DC=saviyntlabs,DC=org1"
		createReq := entitlements.CreateUpdateEntitlementRequest{
			Endpoint:            "AD_Rashid",
			Entitlementtype:     "memberOf",
			EntitlementValue:    "CN=adtestrequest123,OU=DocTeamOU,OU=SaviyntTeams,DC=saviyntlabs,DC=org",
			NewEntitlementValue: &newValue,
		}
		resp, httpRes, err := apiClient.Entitlements.CreateUpdateEntitlement(ctx).
			CreateUpdateEntitlementRequest(createReq).Execute()
		require.NoError(t, err, "Unexpected error in CreateUpdateEntitlement")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	// ─── Test: GetEntitlementValuesForEndpoint ─────────────────────────────────────
	t.Run("Test GetEntitlementValuesForEndpoint", func(t *testing.T) {
		getReq := entitlements.GetEntitlementValuesForEndpointRequest{
			Endpoint:        "AD_Rashid",
			EntitlementType: PtrString("memberOf"), // Optional field
		}
		resp, httpRes, err := apiClient.Entitlements.GetEntitlementValuesForEndpoint(ctx).
			GetEntitlementValuesForEndpointRequest(getReq).Execute()
		require.NoError(t, err, "Unexpected error in GetEntitlementValuesForEndpoint")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	// ─── Test: GetListOfPrivilegesForEntitlementType ──────────────────────────────
	t.Run("Test GetListOfPrivileges", func(t *testing.T) {
		req := apiClient.Entitlements.GetListOfPrivilegesForEntitlementType(ctx).
			Endpoint("AD_Rashid").
			Entitlementtype("memberOf")
		resp, httpRes, err := req.Execute()
		require.NoError(t, err, "Unexpected error in GetListOfPrivileges")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})

	// ─── Test: RemoveEntitlementFromRole ───────────────────────────────────────────
	t.Run("Test RemoveEntitlementFromRole", func(t *testing.T) {
		entitlementValue := "CN=adtestrequest123,OU=DocTeamOU,OU=SaviyntTeams,DC=saviyntlabs,DC=org"
		req := entitlements.RemoveEntitlementFromRoleRequest{
			Requestor: PtrString("admin"),
			Rolename:  PtrString("Shivam Verma"),
			Entitlements: []entitlements.RemoveEntitlementFromRoleRequestEntitlementsInner{
				{
					Entitlementvalue: entitlementValue,
					Entitlementtype:  "memberOf",
					Endpoint:         "AD_Rashid",
				},
			},
		}
		resp, httpRes, err := apiClient.Entitlements.RemoveEntitlementFromRole(ctx).
			RemoveEntitlementFromRoleRequest(req).Execute()
		require.NoError(t, err, "Unexpected error in RemoveEntitlementFromRole")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "Expected HTTP status 200")
	})
}
