package test

import (
	"context"
	"strings"
	"testing"
	"fmt"

	"github.com/saviynt/saviynt-api-go-client/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EndpointsAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.NoError(t, err, "Failed to initialize API client")
	if !skipTests {
		require.NotNil(t, apiClient, "apiClient should not be nil")
	}

	ctx := context.Background()

	t.Run("Test_EndpointsAPIService_CreateEndpoint", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}

		endpointName := "hello-api-client-test"
		displayName := "hello-api-client-test"
		securitySystem := "TestSecuritySystem"

		createReq := endpoint.CreateEndpointRequest{
			Endpointname:                endpointName,
			DisplayName:                 displayName,
			Securitysystem:              securitySystem,
		}

		req := apiClient.Endpoints.CreateEndpoint(ctx).CreateEndpointRequest(createReq)

		resp, httpResp, err := req.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "0", *resp.ErrorCode, "Expected errorCode 0 for success")
		assert.Equal(t, "Success", *resp.Msg, "Expected success message")

		fmt.Printf("Create Endpoint Response: %+v\n", resp)
	})

	t.Run("Test_EndpointsAPIService_UpdateEndpoint", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}

		endpointName := "hello-api-client-test"
		displayName := "hello-api-client-test"
		securitySystem := "TestSecuritySystem"
		description := "Changed description"
		disableNewAccountReqIfAccExists := "true"
		outOfBandAction := "2"

		updateReq := endpoint.UpdateEndpointRequest{
			Endpointname:                            endpointName,
			DisplayName:                             displayName,
			Securitysystem:                          securitySystem,
			Description:                             &description,
			DisableNewAccountRequestIfAccountExists: &disableNewAccountReqIfAccExists,
			Outofbandaction:                         &outOfBandAction,
		}

		req := apiClient.Endpoints.UpdateEndpoint(ctx).UpdateEndpointRequest(updateReq)

		resp, httpResp, err := req.Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "0", *resp.ErrorCode, "Expected errorCode 0 for success")
		assert.Equal(t, "Success", *resp.Msg, "Expected success message")

		fmt.Printf("Update Endpoint Response: %+v\n", resp)
	})

	t.Run("Test_EndpointsAPIService_GetEndpoints", func(t *testing.T) {
		if skipTests {
			if skipMsg != "" {
				t.Skip(skipMsg)
			} else {
				t.Skip(MsgSkipTest)
			}
		}
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip("Skipping test based on configuration")
		}

		getReq := endpoint.GetEndpointsRequest{
			FilterCriteria: map[string]interface{}{
				"customproperty2": "aszfdgxchvjbklnm",
			},
		}

		req := apiClient.Endpoints.GetEndpoints(ctx).GetEndpointsRequest(getReq)

		resp, httpResp, err := req.Execute()

		require.Nil(t, err)
		require.NoError(t, err, "Unexpected error in GetEndpoints")
		require.NotNil(t, httpResp, "httpResp should not be nil")
		fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")

		// fmt.Printf("Get Endpoints Response: %+v\n", resp)
	})
}