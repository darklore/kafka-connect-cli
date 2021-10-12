package connect

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type ConnectorPlugin struct {
	Class   string `json:"class"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

func ListConnectorPlugin(endpoint string) ([]ConnectorPlugin, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connector-plugins")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var plugins []ConnectorPlugin
	if err := json.NewDecoder(resp.Body).Decode(&plugins); err != nil {
		return nil, err
	}

	return plugins, nil
}
