package connect

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type ConnectorName string

type Connector struct {
	Name   ConnectorName   `json:"name"`
	Config ConnectorConfig `json:"config"`
	Type   string          `json:"type"`
	Tasks  []Task          `json:"tasks"`
}

type ConnectorConfig map[string]string

type Task struct {
	Connector string `json:"connector"`
	Task      int    `json:"task"`
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
