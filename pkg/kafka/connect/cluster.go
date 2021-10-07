package connect

type Cluster struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	ID      string `json:"kafka_cluster_id"`
}
