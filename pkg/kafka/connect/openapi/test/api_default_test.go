/*
Kafka Connect REST API

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService AlterConnectorOffsets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.AlterConnectorOffsets(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService CreateConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.CreateConnector(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DestroyConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.DestroyConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetConnectorActiveTopics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.GetConnectorActiveTopics(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetConnectorConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetConnectorConfig(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetConnectorConfigDef", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pluginName string

		resp, httpRes, err := apiClient.DefaultAPI.GetConnectorConfigDef(context.Background(), pluginName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetConnectorStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetConnectorStatus(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetLogger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logger string

		httpRes, err := apiClient.DefaultAPI.GetLogger(context.Background(), logger).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetOffsets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetOffsets(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetTaskConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetTaskConfigs(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetTaskStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string
		var task int32

		resp, httpRes, err := apiClient.DefaultAPI.GetTaskStatus(context.Background(), connector, task).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService GetTasksConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		resp, httpRes, err := apiClient.DefaultAPI.GetTasksConfig(context.Background(), connector).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService HealthCheck", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.HealthCheck(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListConnectorPlugins", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ListConnectorPlugins(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListConnectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.ListConnectors(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ListLoggers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DefaultAPI.ListLoggers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PatchConnectorConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.PatchConnectorConfig(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PauseConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.PauseConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PutConnectorConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.PutConnectorConfig(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ResetConnectorActiveTopics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.ResetConnectorActiveTopics(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ResetConnectorOffsets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.ResetConnectorOffsets(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RestartConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.RestartConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService RestartTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string
		var task int32

		httpRes, err := apiClient.DefaultAPI.RestartTask(context.Background(), connector, task).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ResumeConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.ResumeConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ServerInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DefaultAPI.ServerInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService SetLevel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var logger string

		httpRes, err := apiClient.DefaultAPI.SetLevel(context.Background(), logger).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService StopConnector", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connector string

		httpRes, err := apiClient.DefaultAPI.StopConnector(context.Background(), connector).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ValidateConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pluginName string

		resp, httpRes, err := apiClient.DefaultAPI.ValidateConfigs(context.Background(), pluginName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
