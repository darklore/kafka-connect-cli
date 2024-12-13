package connect

import (
	"context"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
)

func GetServerInfo(cfg *openapi.Configuration) (*openapi.ServerInfo, error) {
	client := openapi.NewAPIClient(cfg)
	info, _, err := client.DefaultAPI.ServerInfo(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	return info, nil
}
