package connect

type ConnectorName string

type Connector struct {
	Name   ConnectorName   `json:"name"`
	Config ConnectorConfig `json:"config"`
	Type   string          `json:"type"`
	Tasks  []Task          `json:"tasks"`
}

type ConnectorConfig struct {
	ConnectorClass string `json:"connector.class"`
}

type Task struct {
	Connector string `json:"connector"`
	Task      int    `json:"task"`
}
