# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlterConnectorOffsets**](DefaultAPI.md#AlterConnectorOffsets) | **Patch** /connectors/{connector}/offsets | Alter the offsets for the specified connector
[**CreateConnector**](DefaultAPI.md#CreateConnector) | **Post** /connectors | Create a new connector
[**DestroyConnector**](DefaultAPI.md#DestroyConnector) | **Delete** /connectors/{connector} | Delete the specified connector
[**GetConnector**](DefaultAPI.md#GetConnector) | **Get** /connectors/{connector} | Get the details for the specified connector
[**GetConnectorActiveTopics**](DefaultAPI.md#GetConnectorActiveTopics) | **Get** /connectors/{connector}/topics | Get the list of topics actively used by the specified connector
[**GetConnectorConfig**](DefaultAPI.md#GetConnectorConfig) | **Get** /connectors/{connector}/config | Get the configuration for the specified connector
[**GetConnectorConfigDef**](DefaultAPI.md#GetConnectorConfigDef) | **Get** /connector-plugins/{pluginName}/config | Get the configuration definition for the specified pluginName
[**GetConnectorStatus**](DefaultAPI.md#GetConnectorStatus) | **Get** /connectors/{connector}/status | Get the status for the specified connector
[**GetLogger**](DefaultAPI.md#GetLogger) | **Get** /admin/loggers/{logger} | Get the log level for the specified logger
[**GetOffsets**](DefaultAPI.md#GetOffsets) | **Get** /connectors/{connector}/offsets | Get the current offsets for the specified connector
[**GetTaskConfigs**](DefaultAPI.md#GetTaskConfigs) | **Get** /connectors/{connector}/tasks | List all tasks and their configurations for the specified connector
[**GetTaskStatus**](DefaultAPI.md#GetTaskStatus) | **Get** /connectors/{connector}/tasks/{task}/status | Get the state of the specified task for the specified connector
[**GetTasksConfig**](DefaultAPI.md#GetTasksConfig) | **Get** /connectors/{connector}/tasks-config | Get the configuration of all tasks for the specified connector
[**HealthCheck**](DefaultAPI.md#HealthCheck) | **Get** /health | Health check endpoint to verify worker readiness and liveness
[**ListConnectorPlugins**](DefaultAPI.md#ListConnectorPlugins) | **Get** /connector-plugins | List all connector plugins installed
[**ListConnectors**](DefaultAPI.md#ListConnectors) | **Get** /connectors | List all active connectors
[**ListLoggers**](DefaultAPI.md#ListLoggers) | **Get** /admin/loggers | List the current loggers that have their levels explicitly set and their log levels
[**PatchConnectorConfig**](DefaultAPI.md#PatchConnectorConfig) | **Patch** /connectors/{connector}/config | 
[**PauseConnector**](DefaultAPI.md#PauseConnector) | **Put** /connectors/{connector}/pause | Pause the specified connector
[**PutConnectorConfig**](DefaultAPI.md#PutConnectorConfig) | **Put** /connectors/{connector}/config | Create or reconfigure the specified connector
[**ResetConnectorActiveTopics**](DefaultAPI.md#ResetConnectorActiveTopics) | **Put** /connectors/{connector}/topics/reset | Reset the list of topics actively used by the specified connector
[**ResetConnectorOffsets**](DefaultAPI.md#ResetConnectorOffsets) | **Delete** /connectors/{connector}/offsets | Reset the offsets for the specified connector
[**RestartConnector**](DefaultAPI.md#RestartConnector) | **Post** /connectors/{connector}/restart | Restart the specified connector
[**RestartTask**](DefaultAPI.md#RestartTask) | **Post** /connectors/{connector}/tasks/{task}/restart | Restart the specified task for the specified connector
[**ResumeConnector**](DefaultAPI.md#ResumeConnector) | **Put** /connectors/{connector}/resume | Resume the specified connector
[**ServerInfo**](DefaultAPI.md#ServerInfo) | **Get** / | Get details about this Connect worker and the ID of the Kafka cluster it is connected to
[**SetLevel**](DefaultAPI.md#SetLevel) | **Put** /admin/loggers/{logger} | Set the log level for the specified logger
[**StopConnector**](DefaultAPI.md#StopConnector) | **Put** /connectors/{connector}/stop | Stop the specified connector
[**ValidateConfigs**](DefaultAPI.md#ValidateConfigs) | **Put** /connector-plugins/{pluginName}/config/validate | Validate the provided configuration against the configuration definition for the specified pluginName



## AlterConnectorOffsets

> AlterConnectorOffsets(ctx, connector).ConnectorOffsets(connectorOffsets).Execute()

Alter the offsets for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	connectorOffsets := *openapiclient.NewConnectorOffsets() // ConnectorOffsets |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AlterConnectorOffsets(context.Background(), connector).ConnectorOffsets(connectorOffsets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AlterConnectorOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlterConnectorOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorOffsets** | [**ConnectorOffsets**](ConnectorOffsets.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnector

> ConnectorInfo CreateConnector(ctx).CreateConnectorRequest(createConnectorRequest).Execute()

Create a new connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	createConnectorRequest := *openapiclient.NewCreateConnectorRequest() // CreateConnectorRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateConnector(context.Background()).CreateConnectorRequest(createConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnector`: ConnectorInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConnectorRequest** | [**CreateConnectorRequest**](CreateConnectorRequest.md) |  | 

### Return type

[**ConnectorInfo**](ConnectorInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyConnector

> DestroyConnector(ctx, connector).Execute()

Delete the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DestroyConnector(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DestroyConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnector

> ConnectorInfo GetConnector(ctx, connector).Execute()

Get the details for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConnector(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnector`: ConnectorInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorInfo**](ConnectorInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorActiveTopics

> GetConnectorActiveTopics(ctx, connector).Execute()

Get the list of topics actively used by the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetConnectorActiveTopics(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConnectorActiveTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorActiveTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorConfig

> map[string]string GetConnectorConfig(ctx, connector).Execute()

Get the configuration for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConnectorConfig(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConnectorConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorConfig`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConnectorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorConfigDef

> []ConfigKeyInfo GetConnectorConfigDef(ctx, pluginName).Execute()

Get the configuration definition for the specified pluginName

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	pluginName := "pluginName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConnectorConfigDef(context.Background(), pluginName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConnectorConfigDef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorConfigDef`: []ConfigKeyInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConnectorConfigDef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorConfigDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigKeyInfo**](ConfigKeyInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorStatus

> ConnectorStateInfo GetConnectorStatus(ctx, connector).Execute()

Get the status for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetConnectorStatus(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetConnectorStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorStatus`: ConnectorStateInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetConnectorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorStateInfo**](ConnectorStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogger

> GetLogger(ctx, logger).Execute()

Get the log level for the specified logger

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	logger := "logger_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetLogger(context.Background(), logger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLogger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logger** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffsets

> ConnectorOffsets GetOffsets(ctx, connector).Execute()

Get the current offsets for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOffsets(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOffsets`: ConnectorOffsets
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOffsets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorOffsets**](ConnectorOffsets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskConfigs

> []TaskInfo GetTaskConfigs(ctx, connector).Execute()

List all tasks and their configurations for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTaskConfigs(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTaskConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskConfigs`: []TaskInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTaskConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TaskInfo**](TaskInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStatus

> TaskState GetTaskStatus(ctx, connector, task).Execute()

Get the state of the specified task for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	task := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTaskStatus(context.Background(), connector, task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskStatus`: TaskState
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 
**task** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskState**](TaskState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksConfig

> map[string]map[string]string GetTasksConfig(ctx, connector).Execute()

Get the configuration of all tasks for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTasksConfig(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTasksConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasksConfig`: map[string]map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTasksConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]map[string]string**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheck(ctx).Execute()

Health check endpoint to verify worker readiness and liveness

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.HealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorPlugins

> []PluginInfo ListConnectorPlugins(ctx).ConnectorsOnly(connectorsOnly).Execute()

List all connector plugins installed

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connectorsOnly := true // bool | Whether to list only connectors instead of all plugins (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListConnectorPlugins(context.Background()).ConnectorsOnly(connectorsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListConnectorPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectorPlugins`: []PluginInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListConnectorPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorsOnly** | **bool** | Whether to list only connectors instead of all plugins | [default to true]

### Return type

[**[]PluginInfo**](PluginInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> []string ListConnectors(ctx).Execute()

List all active connectors

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectors`: []string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoggers

> ListLoggers(ctx).Execute()

List the current loggers that have their levels explicitly set and their log levels

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ListLoggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLoggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoggersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConnectorConfig

> PatchConnectorConfig(ctx, connector).RequestBody(requestBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PatchConnectorConfig(context.Background(), connector).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchConnectorConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchConnectorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseConnector

> PauseConnector(ctx, connector).Execute()

Pause the specified connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PauseConnector(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PauseConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorConfig

> PutConnectorConfig(ctx, connector).Body(body).Execute()

Create or reconfigure the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	body := map[string]string{ ... } // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutConnectorConfig(context.Background(), connector).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutConnectorConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetConnectorActiveTopics

> ResetConnectorActiveTopics(ctx, connector).Execute()

Reset the list of topics actively used by the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResetConnectorActiveTopics(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResetConnectorActiveTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConnectorActiveTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetConnectorOffsets

> ResetConnectorOffsets(ctx, connector).Execute()

Reset the offsets for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResetConnectorOffsets(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResetConnectorOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConnectorOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartConnector

> RestartConnector(ctx, connector).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()

Restart the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	includeTasks := true // bool | Whether to also restart tasks (optional) (default to false)
	onlyFailed := true // bool | Whether to only restart failed tasks/connectors (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RestartConnector(context.Background(), connector).IncludeTasks(includeTasks).OnlyFailed(onlyFailed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestartConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTasks** | **bool** | Whether to also restart tasks | [default to false]
 **onlyFailed** | **bool** | Whether to only restart failed tasks/connectors | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartTask

> RestartTask(ctx, connector, task).Execute()

Restart the specified task for the specified connector

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 
	task := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RestartTask(context.Background(), connector, task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestartTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 
**task** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeConnector

> ResumeConnector(ctx, connector).Execute()

Resume the specified connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResumeConnector(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResumeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServerInfo

> ServerInfo ServerInfo(ctx).Execute()

Get details about this Connect worker and the ID of the Kafka cluster it is connected to

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ServerInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ServerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServerInfo`: ServerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ServerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServerInfoRequest struct via the builder pattern


### Return type

[**ServerInfo**](ServerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLevel

> SetLevel(ctx, logger).Scope(scope).RequestBody(requestBody).Execute()

Set the log level for the specified logger

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	logger := "logger_example" // string | 
	scope := "scope_example" // string | The scope for the logging modification (single-worker, cluster-wide, etc.) (optional) (default to "worker")
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SetLevel(context.Background(), logger).Scope(scope).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SetLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logger** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | The scope for the logging modification (single-worker, cluster-wide, etc.) | [default to &quot;worker&quot;]
 **requestBody** | **map[string]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopConnector

> StopConnector(ctx, connector).Execute()

Stop the specified connector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	connector := "connector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.StopConnector(context.Background(), connector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StopConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConfigs

> ConfigInfos ValidateConfigs(ctx, pluginName).RequestBody(requestBody).Execute()

Validate the provided configuration against the configuration definition for the specified pluginName

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/darklore/kafka-connect-cli"
)

func main() {
	pluginName := "pluginName_example" // string | 
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ValidateConfigs(context.Background(), pluginName).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ValidateConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateConfigs`: ConfigInfos
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ValidateConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]string** |  | 

### Return type

[**ConfigInfos**](ConfigInfos.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

