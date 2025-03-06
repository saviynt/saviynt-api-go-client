package test

import (
	"context"
	"strings"
	"testing"

	"github.com/saviynt/saviynt-api-go-client/endpoints"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_endpoints_EndpointsAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()

	t.Run("Test_EndpointsAPIService_CreateEndpoint", func(t *testing.T) {
		t.Skip("skip test: unimplemented: `Test_EndpointsAPIService_CreateEndpoint`")

		resp, httpRes, err := apiClient.Endpoints.CreateEndpoint(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test_EndpointsAPIService_GetEndpoints", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		resp, httpRes, err := apiClient.Endpoints.
			GetEndpoints(ctx).
			GetEndpointsRequest(endpoints.GetEndpointsRequest{}).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", resp.ErrorCode)
	})

	t.Run("Test_EndpointsAPIService_UpdateEndpoint", func(t *testing.T) {
		t.Skip("skip test: unimplemented: `Test_EndpointsAPIService_UpdateEndpoint`")

		resp, httpRes, err := apiClient.Endpoints.UpdateEndpoint(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
