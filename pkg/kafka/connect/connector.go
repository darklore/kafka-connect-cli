package connect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type ConnectorName = string

type Connector struct {
	Name   ConnectorName   `json:"name"`
	Config ConnectorConfig `json:"config"`
	Type   string          `json:"type"`
	Tasks  []Task          `json:"tasks"`
}

type ConnectorConfig map[string]string

type Task struct {
	Connector string `json:"connector"`
	ID        int    `json:"task"`
}

type ConnectorStatus struct {
	Name  string         `json:"name"`
	State ConnectorState `json:"connector"`
	Tasks []TaskState    `json:"tasks"`
	Type  string         `json:"type"`
}

type ConnectorState struct {
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}

type TaskState struct {
	ID       int    `json:"id"`
	State    string `json:"state"`
	WorkerID string `json:"worker_id"`
}

type TaskInfo struct {
	ID     Task       `json:"id"`
	Config TaskConfig `json:"config"`
}

type TaskConfig map[string]string

type TaskTopics map[string]Topics
type Topics struct {
	Topics []string `json:"topics"`
}

func CreateConnector(endpoint string, configJSON io.Reader) (*Connector, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors")
	resp, err := http.Post(u.String(), "application/json", configJSON)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	//r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode != http.StatusCreated {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	var connector Connector
	if err := json.NewDecoder(r).Decode(&connector); err != nil {
		return nil, err
	}

	return &connector, nil
}

func ListConnectors(endpoint string) ([]ConnectorName, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorMsg Error
		if err := json.NewDecoder(resp.Body).Decode(&errorMsg); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	var connectors []ConnectorName
	if err := json.NewDecoder(resp.Body).Decode(&connectors); err != nil {
		return nil, err
	}

	return connectors, nil
}

func GetConnector(endpoint, name string) (*Connector, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var connector Connector
	if err := json.NewDecoder(resp.Body).Decode(&connector); err != nil {
		return nil, err
	}

	return &connector, nil
}

func GetConnectorConfig(endpoint, name string) (*ConnectorConfig, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "config")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var config ConnectorConfig
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetConnectorStatus(endpoint, name string) (*ConnectorStatus, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "status")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var status ConnectorStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func RestartConnector(endpoint, name string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "connectors", name, "restart")
	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	// r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return err
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func PauseConnector(endpoint, name string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "connectors", name, "pause")
	req, err := http.NewRequest(http.MethodPut, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	//r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return err
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func ResumeConnector(endpoint, name string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "connectors", name, "resume")
	req, err := http.NewRequest(http.MethodPut, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	//r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return err
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func DeleteConnector(endpoint, name string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "connectors", name)

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var errorMsg Error
		if err := json.NewDecoder(resp.Body).Decode(&errorMsg); err != nil {
			return err
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func ListTasks(endpoint, name string) ([]TaskInfo, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "tasks")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []TaskInfo
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskStatus(endpoint, name, id string) (*TaskState, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "tasks", id, "status")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorMsg Error
		if err := json.NewDecoder(resp.Body).Decode(&errorMsg); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	var task TaskState
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func RestartTask(endpoint, name, id string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "connectors", name, "tasks", id, "restart")
	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	// r = io.TeeReader(resp.Body, os.Stderr)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorMsg Error
		if err := json.NewDecoder(r).Decode(&errorMsg); err != nil {
			return err
		}
		return fmt.Errorf("%d: %s", errorMsg.Code, errorMsg.Message)
	}

	return nil
}

func ListTopics(endpoint, name string) ([]string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "connectors", name, "topics")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var taskTopics TaskTopics
	if err := json.NewDecoder(resp.Body).Decode(&taskTopics); err != nil {
		return nil, err
	}

	return taskTopics[name].Topics, nil
}
