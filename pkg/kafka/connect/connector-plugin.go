package connect

import (
	"context"
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
)

func (c *Client) ListConnectorPlugins() ([]openapi.PluginInfo, error) {
	plugins, _, err := c.openapi.DefaultAPI.ListConnectorPlugins(context.Background()).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list connector plugins: %w", err)
	}

	return plugins, nil
}
