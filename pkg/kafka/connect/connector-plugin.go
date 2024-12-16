package connect

import (
	"context"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
)

func (c *Client) ListConnectorPlugin() ([]openapi.PluginInfo, error) {
	plugins, _, err := c.openapi.DefaultAPI.ListConnectorPlugins(context.Background()).Execute()
	if err != nil {
		return nil, err
	}

	return plugins, nil
}
