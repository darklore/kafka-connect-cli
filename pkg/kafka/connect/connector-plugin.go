package connect

import (
	"context"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
)

func ListConnectorPluginOpenApi(cfg *openapi.Configuration) ([]openapi.PluginInfo, error) {
	client := openapi.NewAPIClient(cfg)
	ctx := context.Background()
	plugins, _, err := client.DefaultAPI.ListConnectorPlugins(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return plugins, nil
}
