package test

import (
	"context"
	"testing"
	"fmt"

	// saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
	usergroups "github.com/saviynt/saviynt-api-go-client/usergroups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_connections_UsergroupsAPIService(t *testing.T) {
	apiClient, _, _, _, err := client()
	require.Nil(t, err)

	ctx := context.Background()

	t.Run("Test GroupsAPIService GetListOfUsergroups", func(t *testing.T) {
		// Define parameters
		systemName := "amigopod"
		offset := "0"
		connectionName := "CN_SERP_ECC_A53"
		connectionType := "SAP"

		// Create request
		userGroupReq := apiClient.Usergroups.GetListOfUsergroups(ctx).
			Systemname(systemName).
			Offset(offset).
			Connectionname(connectionName).
			ConnectionType(connectionType).
			GetListOfUsergroupsRequest(usergroups.GetListOfUsergroupsRequest{})

		// Execute request
		response, httpResponse, err := userGroupReq.Execute()

		// Assertions
		require.NoError(t, err, "Error calling GetListOfUsergroups")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")

		// Additional JSON response validation
		assert.Equal(t, "success", *response.Msg, "Expected msg to be 'success'")
		assert.GreaterOrEqual(t, *response.Totalcount, 0, "Expected totalcount to be >= 0")

		// Validate user groups exist
		// require.NotEmpty(t, response.Usergroups, "Usergroups should not be empty")

		// Print response for debugging
		fmt.Printf("HTTP Status: %s\n", httpResponse.Status)
		// fmt.Printf("Response: %+v\n", response)
	})

	t.Run("Test Successful Group Creation", func(t *testing.T) {
		// Mandatory parameters
		usergroup := "Finance_Team_02"
		username := "admin"

		// Request payload
		createUpdateReq := usergroups.CreateUpdateUsergroupRequest{
			Usergroup: &usergroup,
			Username:  &username,
		}

		// Execute API request
		createReq := apiClient.Usergroups.CreateUpdateUsergroup(ctx).
			CreateUpdateUsergroupRequest(createUpdateReq)

		response, _, err := createReq.Execute()

		// // Assertions
		require.NoError(t, err, "Unexpected error in API call")
		// require.NotNil(t, response, "Response should not be nil")
		// assert.Equal(t, "Success", response.Msg, "Expected success message")
		assert.Equal(t, "0", *response.ErrorCode, "Expected errorCode 0 for success")

		// Print response
		fmt.Printf("Success Response: %+v\n", response)
	})

	t.Run("Test Missing Mandatory Parameters", func(t *testing.T) {
		// Request with missing mandatory fields
		createUpdateReq := usergroups.CreateUpdateUsergroupRequest{}

		// Execute API request
		createReq := apiClient.Usergroups.CreateUpdateUsergroup(ctx).
			CreateUpdateUsergroupRequest(createUpdateReq)

		response, _, err := createReq.Execute()

		// Assertions
		// require.Error(t, err, "Expected error for missing parameters")
		// assert.Nil(t, response, "Response should be nil")
		// assert.Equal(t, 400, httpResponse.StatusCode, "Expected HTTP status 400 for bad request")
		assert.Equal(t, "1", *response.ErrorCode, "Expected errorCode 1 for missing mandatory parameters")

		// Print error message
		fmt.Printf("Error Response: %v\n", err)
	})

	t.Run("Test Add User to Group Successfully", func(t *testing.T) {
		actionType := "0" // Add user
		userGroupname := "Finance_Team_02"
		username := "100450"

		request := apiClient.Usergroups.
			AddRemoveUserFromUsergroup(ctx).
			ActionType(actionType).
			UserGroupname(userGroupname).
			Username(username)

		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "saved successfully", *response.Message, "Expected success message")
		assert.Equal(t, "0", *response.ErrorCode, "Expected errorCode 0 for success")

		fmt.Printf("Success Response: %+v\n", response)
	})


	t.Run("Test Add User Already in Group", func(t *testing.T) {
		actionType := "0" // Add user
		userGroupname := "Finance_Team_02"
		username := "100450"

		request := apiClient.Usergroups.
			AddRemoveUserFromUsergroup(ctx).
			ActionType(actionType).
			UserGroupname(userGroupname).
			Username(username)

		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		require.NotNil(t, response, "Response should not be nil")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "User already Present in Group ", *response.Message, "Expected user already exists message")
		assert.Equal(t, "1", *response.ErrorCode, "Expected errorCode 1 for failure")

		fmt.Printf("User Already Present Response: %+v\n", response)
	})

	t.Run("Test Remove User from Group Successfully", func(t *testing.T) {
		actionType := "1" // Remove user
		userGroupname := "Finance_Team_02"
		username := "100450"

		request := apiClient.Usergroups.
			AddRemoveUserFromUsergroup(ctx).
			ActionType(actionType).
			UserGroupname(userGroupname).
			Username(username)

		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "delete successful", *response.Message, "Expected delete success message")
		assert.Equal(t, "0", *response.ErrorCode, "Expected errorCode 0 for success")

		fmt.Printf("Delete Success Response: %+v\n", response)
	})

	t.Run("Test Remove User Who is Not in Group", func(t *testing.T) {
		actionType := "1" // Remove user
		userGroupname := "Finance_Team_02"
		username := "0404abc"

		request := apiClient.Usergroups.
			AddRemoveUserFromUsergroup(ctx).
			ActionType(actionType).
			UserGroupname(userGroupname).
			Username(username)

		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		require.NotNil(t, response, "Response should not be nil")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "User not Present in Group", *response.Message, "Expected user not found message")
		assert.Equal(t, "1", *response.ErrorCode, "Expected errorCode 1 for failure")

		fmt.Printf("User Not in Group Response: %+v\n", response)
	})


	t.Run("Test Missing Mandatory Parameters", func(t *testing.T) {
		// Request with missing mandatory fields
		request := apiClient.Usergroups.AddRemoveUserFromUsergroup(ctx)

		response, httpResponse, err := request.Execute()

		// Assertions
		require.Error(t, err, "Expected error for missing parameters")
		assert.Nil(t, response, "Response should be nil")
		assert.Equal(t, "Username, User GroupName  is Required", *response.Message, "Missing mandatory params")
		assert.Equal(t, 400, httpResponse.StatusCode, "Expected HTTP status 400 for bad request")

		// Print error message
		fmt.Printf("Error Response: %v\n", err)
	})

	
	t.Run("Test Delete User Group Successfully", func(t *testing.T) {
		usergroup := "Finance_Team_02"
		username := "admin"

		deleteRequest := usergroups.DeleteUsergroupRequest{
			Usergroup: &usergroup,
			Username:  &username,
		}

		request := apiClient.Usergroups.DeleteUsergroup(ctx).DeleteUsergroupRequest(deleteRequest)
		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "Success", *response.Msg, "Expected success message")
		assert.Equal(t, "0", *response.ErrorCode, "Expected errorCode 0 for success")

		fmt.Printf("Delete Success Response: %+v\n", response)
	})

	t.Run("Test Delete Non-Existent User Group", func(t *testing.T) {
		usergroup := "nonExistentGroup"
		username := "admin"

		deleteRequest := usergroups.DeleteUsergroupRequest{
			Usergroup: &usergroup,
			Username:  &username,
		}

		request := apiClient.Usergroups.DeleteUsergroup(ctx).DeleteUsergroupRequest(deleteRequest)
		response, httpResponse, err := request.Execute()

		// Assertions
		require.NoError(t, err, "Unexpected error in API call")
		assert.Equal(t, 200, httpResponse.StatusCode, "Expected HTTP status 200")
		assert.Equal(t, "usergroup not found", *response.Msg, "Expected error message for non-existent group")
		assert.Equal(t, "1", *response.ErrorCode, "Expected errorCode 1 for failure")

		fmt.Printf("Delete Non-Existent Response: %+v\n", response)
	})

	t.Run("Test Missing Mandatory Parameters for Deletion", func(t *testing.T) {
		username:="admin"
		
		deleteRequest := usergroups.DeleteUsergroupRequest{
			Username: &username,
		}

		request := apiClient.Usergroups.DeleteUsergroup(ctx).DeleteUsergroupRequest(deleteRequest)
		response, _, err := request.Execute()

		if err != nil {
			t.Logf("Expected error received: %v", err)
			return
		}

		if response == nil {
			t.Fatalf("Response is nil, but no error was returned")
		}

		assert.NotNil(t, response.ErrorCode, "Expected errorCode to be present")
		assert.Equal(t, "1", *response.ErrorCode, "Expected errorCode 1 for missing param")
		assert.Equal(t, "usergroup cannot be null or blank", *response.Msg, "Expected message for missing usergroup")

		fmt.Printf("Missing Params Response: %v\n", response)
	})

}
