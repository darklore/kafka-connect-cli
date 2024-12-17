package connect

import "github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"

type Client struct {
	openapi *openapi.APIClient
}

func NewClient(cfg *openapi.Configuration) *Client {
	return &Client{
		openapi: openapi.NewAPIClient(cfg),
	}
}
