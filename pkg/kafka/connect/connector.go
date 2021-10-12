package connect

type ConnectorName string

type Connector struct {
	name ConnectorName `json:"name"`
}
