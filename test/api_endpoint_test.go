package test

import (
	"context"
	"testing"
	"fmt"

	endpoint "github.com/saviynt/saviynt-api-go-client/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_connections_EndpointAPIService(t *testing.T) {
	apiClient, _, _, _, err := client()
	require.Nil(t, err)

	ctx := context.Background()
	t.Run("Test Successful Endpoint Creation", func(t *testing.T) {
		createReq := endpoint.CreateEndpointRequest{
			Endpointname:   "test-endpoint-sdk5",
			Securitysystem: "TestSecuritySystem",
			DisplayName:    "test-endpoint-sdk5",
		}
	
		req := apiClient.Endpoint.CreateEndpoint(ctx).CreateEndpointRequest(createReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "0", *apiResponse.ErrorCode, "Expected errorCode 0 for success")
		assert.Equal(t, "Success", *apiResponse.Msg, "Expected success message")
	
		fmt.Printf("Create Endpoint Success Response: %+v\n", apiResponse)
	})

	t.Run("Test Creating Endpoint with Non-Existing Security System", func(t *testing.T) {
		createReq := endpoint.CreateEndpointRequest{
			Endpointname:   "test-invalid",
			Securitysystem: "non-existing-system",
			DisplayName:    "test-invalid",
		}
	
		req := apiClient.Endpoint.CreateEndpoint(ctx).CreateEndpointRequest(createReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "1", *apiResponse.ErrorCode, "Expected errorCode 1 for failure")
		assert.Equal(t, "systemname not found", *apiResponse.Msg, "Expected system not found message")
	
		fmt.Printf("Create Endpoint Non-Existing Security System Response: %+v\n", apiResponse)
	})

	t.Run("Test Creating Duplicate Endpoint", func(t *testing.T) {
		createReq := endpoint.CreateEndpointRequest{
			Endpointname:   "test-endpoint-sdk3",
			Securitysystem: "TestSecuritySystem",
			DisplayName:    "test-endpoint-sdk-duplicate3",
		}
	
		req := apiClient.Endpoint.CreateEndpoint(ctx).CreateEndpointRequest(createReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "1", *apiResponse.ErrorCode, "Expected errorCode 1 for duplicate entry")
		// assert.Contains(t, *apiResponse.Msg, "endpointname test-endpoint-sdk already exists", "Expected duplicate endpoint message")

		expectedMsg := fmt.Sprintf("endpointname %s already exists", createReq.Endpointname)
		assert.Contains(t, *apiResponse.Msg, expectedMsg, fmt.Sprintf("Expected message to contain: %s", expectedMsg))

		fmt.Printf("Create Duplicate Endpoint Response: %+v\n", apiResponse)
	})
	
	t.Run("Test Creating Endpoint Without Mandatory Parameters", func(t *testing.T) {
		createReq := endpoint.CreateEndpointRequest{
			Endpointname: "",         // Missing Endpoint Name
			Securitysystem: "TestSecuritySystem",       // Missing Security System
			DisplayName: "test-endpoint",
		}
	
		req := apiClient.Endpoint.CreateEndpoint(ctx).CreateEndpointRequest(createReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "1", *apiResponse.ErrorCode, "Expected errorCode 1 for failure")
		assert.Contains(t, *apiResponse.Msg, "endpointname not passed", "Expected missing security system error")
	
		fmt.Printf("Create Endpoint Missing Mandatory Params Response: %+v\n", apiResponse)
	})
	
	t.Run("Test Get List of Endpoints", func(t *testing.T) {
		// Create request
		getReq := endpoint.NewGetEndpointsRequest()
		apiRequest := apiClient.Endpoint.GetEndpoints(ctx).GetEndpointsRequest(*getReq)

		apiResponse, _, err := apiRequest.Execute()

		require.NoError(t, err, "Unexpected error while fetching endpoint list")
		// assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")

		assert.Equal(t, "0", *apiResponse.ErrorCode, "Expected errorCode 0 for success")
		// assert.Equal(t, "Success", apiResponse.message, "Expected message to be 'Success'")
		// assert.GreaterOrEqual(t, apiResponse.displayCount, 1, "Expected displayCount to be >= 1")
		// assert.GreaterOrEqual(t, apiResponse.totalCount, 1, "Expected totalCount to be >= 1")


		// Print for debugging
		fmt.Printf("API Response: %+v\n", apiResponse)
	})

	t.Run("Test Successful Endpoint Update", func(t *testing.T) {
		updateReq := endpoint.UpdateEndpointRequest{
			Endpointname: "test-endpoint-sdk",
		}
	
		req := apiClient.Endpoint.UpdateEndpoint(ctx).UpdateEndpointRequest(updateReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "0", *apiResponse.ErrorCode, "Expected errorCode 0 for success")
		assert.Equal(t, "Success", *apiResponse.Msg, "Expected success message")
	
		fmt.Printf("Update Success Response: %+v\n", apiResponse)
	})

	t.Run("Test Update Non-Existent Endpoint", func(t *testing.T) {
		updateReq := endpoint.UpdateEndpointRequest{
			Endpointname: "sample",
		}
	
		req := apiClient.Endpoint.UpdateEndpoint(ctx).UpdateEndpointRequest(updateReq)
		apiResponse, httpResp, err := req.Execute()
	
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResp.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "1", *apiResponse.ErrorCode, "Expected errorCode 1 for failure")
		expectedMsg := fmt.Sprintf("Endpoint %s not found", updateReq.Endpointname)
		assert.Contains(t, *apiResponse.Msg, expectedMsg, fmt.Sprintf("Expected message to contain: %s", expectedMsg))		
		// assert.Equal(t, "Endpoint sample not found", *apiResponse.Msg, "Expected endpoint not found message")
	
		fmt.Printf("Update Failure Response: %+v\n", apiResponse)
	})
}
