/*
NRF NFManagement Service

Testing SharedDataDocumentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_SharedDataDocumentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SharedDataDocumentAPIService DeleteSharedData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedDataId string

		httpRes, err := apiClient.SharedDataDocumentAPI.DeleteSharedData(context.Background(), sharedDataId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedDataDocumentAPIService GetSharedData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedDataId string

		resp, httpRes, err := apiClient.SharedDataDocumentAPI.GetSharedData(context.Background(), sharedDataId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedDataDocumentAPIService RegisterSharedData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedDataId string

		resp, httpRes, err := apiClient.SharedDataDocumentAPI.RegisterSharedData(context.Background(), sharedDataId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SharedDataDocumentAPIService UpdateSharedData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sharedDataId string

		resp, httpRes, err := apiClient.SharedDataDocumentAPI.UpdateSharedData(context.Background(), sharedDataId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
