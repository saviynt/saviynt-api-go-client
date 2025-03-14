/*
Saviynt Transport API

Testing TransportAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package transport

import (
	"context"
	"testing"

	openapiclient "github.com/saviynt/saviynt-api-go-client/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_transport_TransportAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransportAPIService ExportTransportPackage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TransportAPI.ExportTransportPackage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportAPIService ImportTransportPackage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TransportAPI.ImportTransportPackage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransportAPIService TransportPackageStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TransportAPI.TransportPackageStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
