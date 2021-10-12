package connect

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type ConnectorName = string

type Connector struct {
	Name   ConnectorName   `json:"name"`
	Config ConnectorConfig `json:"config"`
	Type   string          `json:"type"`
	Tasks  []Task          `json:"tasks"`
}

type ConnectorConfig map[string]string

type Task struct {
	Connector string `json:"connector"`
	ID        int    `json:"task"`
}

type ConnectorStatus struct {
	Name  string         `json:"name"`
	State ConnectorState `json:"connector"`
	Tasks []TaskState    `json:"tasks"`
	Type  string         `json:"type"`
}

type ConnectorState struct {
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}

type TaskState struct {
	ID       int    `json:"id"`
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}

func GetConnectorNames(endpoint string) ([]ConnectorName, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var connectors []ConnectorName
	if err := json.NewDecoder(resp.Body).Decode(&connectors); err != nil {
		return nil, err
	}

	return connectors, nil
}

func GetConnector(endpoint, name string) (*Connector, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var connector Connector
	if err := json.NewDecoder(resp.Body).Decode(&connector); err != nil {
		return nil, err
	}

	return &connector, nil
}

func GetConnectorConfig(endpoint, name string) (*ConnectorConfig, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "config")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var config ConnectorConfig
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetConnectorStatus(endpoint, name string) (*ConnectorStatus, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "status")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status ConnectorStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func DeleteConnector(endpoint, name string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "connectors", name)

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorMsg Error
		if err := json.NewDecoder(resp.Body).Decode(&errorMsg); err != nil {
			return err
		}
		// TODO return error
		fmt.Printf("%d: %s\n", errorMsg.Code, errorMsg.Message)
	}

	return nil
}
