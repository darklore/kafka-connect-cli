package connect

import (
	"context"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
)

func (c *Client) GetServerInfo() (*openapi.ServerInfo, error) {
	info, _, err := c.openapi.DefaultAPI.ServerInfo(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	return info, nil
}
