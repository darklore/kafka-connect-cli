package connect

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
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

func ListConnectorPlugin2(cfg *openapi.Configuration) ([]openapi.PluginInfo, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()
	req := client.DefaultAPI.ListConnectorPlugins(ctx)
	plugins, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	return plugins, nil
}
