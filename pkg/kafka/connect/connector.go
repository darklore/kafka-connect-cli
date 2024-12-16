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

type Client struct {
	openapi *openapi.APIClient
}

func NewClient(cfg *openapi.Configuration) *Client {
	return &Client{
		openapi: openapi.NewAPIClient(cfg),
	}
}

type ConnectorName = string

type ConnectorConfig = map[string]string

func (c *Client) CreateConnector(configJson io.Reader) (*openapi.ConnectorInfo, error) {
	var connectorReq openapi.CreateConnectorRequest
	if err := json.NewDecoder(configJson).Decode(&connectorReq); err != nil {
		return nil, errors.Wrap(err, "Failed decode json")
	}

	info, _, err := c.openapi.DefaultAPI.CreateConnector(context.Background()).CreateConnectorRequest(connectorReq).Execute()
	if err != nil {
		return nil, err
	}

	return info, err
}

func (c *Client) UpdateConnector(name string, configJson io.Reader) (*openapi.ConnectorInfo, error) {
	var connectorConfig ConnectorConfig
	if err := json.NewDecoder(configJson).Decode(&connectorConfig); err != nil {
		return nil, errors.Wrap(err, "Failed decode json")
	}

	info, _, err := c.openapi.DefaultAPI.PutConnectorConfig(context.Background(), name).Body(connectorConfig).Execute()
	if err != nil {
		return nil, err
	}

	return info, err
}

func (c *Client) ListConnectors() ([]ConnectorName, error) {
	ctx := context.Background()

	connectors, _, err := c.openapi.DefaultAPI.ListConnectors(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return connectors, nil
}

func (c *Client) GetConnector(name string) (*openapi.ConnectorInfo, error) {
	connector, _, err := c.openapi.DefaultAPI.GetConnector(context.Background(), name).Execute()
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func (c *Client) GetConnectorConfig(name string) (*ConnectorConfig, error) {
	config, _, err := c.openapi.DefaultAPI.GetConnectorConfig(context.Background(), name).Execute()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Client) GetConnectorStatus(name string) (*openapi.ConnectorStateInfo, error) {
	status, _, err := c.openapi.DefaultAPI.GetConnectorStatus(context.Background(), name).Execute()
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c *Client) RestartConnector(name string, includeTasks, onlyFailed bool) error {
	resp, err := c.openapi.DefaultAPI.RestartConnector(context.Background(), name).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()
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

func (c *Client) PauseConnector(name string) error {
	resp, err := c.openapi.DefaultAPI.PauseConnector(context.Background(), name).Execute()
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

func (c *Client) ResumeConnector(name string) error {
	resp, err := c.openapi.DefaultAPI.ResumeConnector(context.Background(), name).Execute()
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

func (c *Client) DeleteConnector(name string) error {
	resp, err := c.openapi.DefaultAPI.DestroyConnector(context.Background(), name).Execute()
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

func (c *Client) GetTaskConfigs(name string) ([]openapi.TaskInfo, error) {
	tasks, _, err := c.openapi.DefaultAPI.GetTaskConfigs(context.Background(), name).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get task configs")
	}
	return tasks, nil
}

func (c *Client) GetTaskStatus(connector string, taskID int32) (*openapi.TaskState, error) {
	state, _, err := c.openapi.DefaultAPI.GetTaskStatus(context.Background(), connector, taskID).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get task status")
	}
	return state, nil
}

func (c *Client) StopConnector(connector string) error {
	resp, err := c.openapi.DefaultAPI.StopConnector(context.Background(), connector).Execute()
	if err != nil {
		return err
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

func (c *Client) RestartTask(connector string, taskId int32) error {
	resp, err := c.openapi.DefaultAPI.RestartTask(context.Background(), connector, taskId).Execute()
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

func (c *Client) ListTopics(name string) (*map[string]openapi.ConnectorActiveTopicsValue, error) {
	topics, _, err := c.openapi.DefaultAPI.GetConnectorActiveTopics(context.Background(), name).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to do http request")
	}
	return topics, nil
}

func (c *Client) GetOffsets(connector string) (*openapi.ConnectorOffsets, error) {
	offsets, _, err := c.openapi.DefaultAPI.GetOffsets(context.Background(), connector).Execute()
	if err != nil {
		return nil, err
	}
	return offsets, nil
}

func (c *Client) ResetOffsets(connector string) error {
	resp, err := c.openapi.DefaultAPI.ResetConnectorOffsets(context.Background(), connector).Execute()
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
