/*
Saviynt File Directory API

Testing FileDirectoryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package filedirectory

import (
	"context"
	"testing"

	openapiclient "github.com/saviynt/saviynt-api-go-client/filedirectory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_filedirectory_FileDirectoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FileDirectoryAPIService UploadNewFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FileDirectoryAPI.UploadNewFile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
