package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
	"github.com/pkg/errors"
)

type ConnectorName = string

type ConnectorConfig = map[string]string

func CreateConnectorOpenApi(cfg *openapi.Configuration, configJson io.Reader) (*openapi.ConnectorInfo, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	var connectorReq openapi.CreateConnectorRequest
	if err := json.NewDecoder(configJson).Decode(&connectorReq); err != nil {
		return nil, errors.Wrap(err, "Failed decode json")
	}

	info, _, err := client.DefaultAPI.CreateConnector(ctx).CreateConnectorRequest(connectorReq).Execute()
	if err != nil {
		return nil, err
	}

	return info, err
}

func UpdateConnectorOpenApi(cfg *openapi.Configuration, name string, configJson io.Reader) (*openapi.ConnectorInfo, error) {
	client := openapi.NewAPIClient(cfg)

	var connectorConfig ConnectorConfig
	if err := json.NewDecoder(configJson).Decode(&connectorConfig); err != nil {
		return nil, errors.Wrap(err, "Failed decode json")
	}

	info, _, err := client.DefaultAPI.PutConnectorConfig(context.Background(), name).Body(connectorConfig).Execute()
	if err != nil {
		return nil, err
	}

	return info, err
}

func ListConnectorsOpenApi(cfg *openapi.Configuration) ([]ConnectorName, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	connectors, _, err := client.DefaultAPI.ListConnectors(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return connectors, nil
}

func GetConnectorOpenApi(cfg *openapi.Configuration, name string) (*openapi.ConnectorInfo, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	connector, _, err := client.DefaultAPI.GetConnector(ctx, name).Execute()
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func GetConnectorConfigOpenApi(cfg *openapi.Configuration, name string) (*ConnectorConfig, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	config, _, err := client.DefaultAPI.GetConnectorConfig(ctx, name).Execute()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func GetConnectorStatusOpenApi(cfg *openapi.Configuration, name string) (*openapi.ConnectorStateInfo, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	status, _, err := client.DefaultAPI.GetConnectorStatus(ctx, name).Execute()
	if err != nil {
		return nil, err
	}

	return status, nil
}

func RestartConnectorOpenApi(cfg *openapi.Configuration, name string, includeTasks, onlyFailed bool) error {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	resp, err := client.DefaultAPI.RestartConnector(ctx, name).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to do http request")
	}
	defer resp.Body.Close()

	// var r io.Reader = resp.Body
	var r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return errors.Wrap(err, "Failed to decode json")
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func PauseConnectorOpenApi(cfg *openapi.Configuration, name string) error {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	resp, err := client.DefaultAPI.PauseConnector(ctx, name).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to do http request")
	}
	defer resp.Body.Close()

	// var r io.Reader = resp.Body
	var r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return errors.Wrap(err, "Failed to decode json")
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func ResumeConnectorOpenApi(cfg *openapi.Configuration, name string) error {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	resp, err := client.DefaultAPI.ResumeConnector(ctx, name).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to do http request")
	}
	defer resp.Body.Close()

	// var r io.Reader = resp.Body
	var r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return errors.Wrap(err, "Failed to decode json")
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func DeleteConnectorOpenApi(cfg *openapi.Configuration, name string) error {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	resp, err := client.DefaultAPI.DestroyConnector(ctx, name).Execute()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// var r = resp.Body
	var r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode != http.StatusNoContent {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return errors.Wrap(err, "Failed to decode json")
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}
	return nil
}

func GetTaskConfigs(cfg *openapi.Configuration, name string) ([]openapi.TaskInfo, error) {
	client := openapi.NewAPIClient(cfg)

	tasks, _, err := client.DefaultAPI.GetTaskConfigs(context.Background(), name).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get task configs")
	}
	return tasks, nil
}

func GetTaskStatusOpenApi(cfg *openapi.Configuration, connector string, taskID int32) (*openapi.TaskState, error) {
	client := openapi.NewAPIClient(cfg)
	state, _, err := client.DefaultAPI.GetTaskStatus(context.Background(), connector, taskID).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get task status")
	}
	return state, nil
}

func RestartTaskOpenApi(cfg *openapi.Configuration, connector string, taskId int32) error {
	client := openapi.NewAPIClient(cfg)
	resp, err := client.DefaultAPI.RestartTask(context.Background(), connector, taskId).Execute()
	if err != nil {
		return errors.Wrap(err, "Failed to do http request")
	}
	defer resp.Body.Close()

	// var r io.Reader = resp.Body
	var r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return errors.Wrap(err, "Failed to decode json")
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func ListTopicsOpenApi(cfg *openapi.Configuration, name string) (*map[string]openapi.ConnectorActiveTopicsValue, error) {
	client := openapi.NewAPIClient(cfg)
	topics, _, err := client.DefaultAPI.GetConnectorActiveTopics(context.Background(), name).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to do http request")
	}
	return topics, nil
}
