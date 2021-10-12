package connect

import (
	"encoding/json"
	"net/http"
)

type Worker struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	ID      string `json:"kafka_worker_id"`
}

func GetWorker(endpoint string) (*Worker, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var worker Worker
	if err := json.NewDecoder(resp.Body).Decode(&worker); err != nil {
		return nil, err
	}

	return &worker, nil
}
