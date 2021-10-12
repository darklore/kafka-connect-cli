package connect

import (
	"encoding/json"
	"net/http"
)

type Cluster struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	ID      string `json:"kafka_cluster_id"`
}

func GetCluster(endpoint string) (*Cluster, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cluster Cluster
	if err := json.NewDecoder(resp.Body).Decode(&cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}
