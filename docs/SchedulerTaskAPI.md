# \SchedulerTaskAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchedulerTaskAdd**](SchedulerTaskAPI.md#SchedulerTaskAdd) | **Post** /api/v1/scheduler/tasks | Register a new scheduled task
[**SchedulerTaskDelete**](SchedulerTaskAPI.md#SchedulerTaskDelete) | **Delete** /api/v1/scheduler/tasks/{id} | Delete an existing scheduled task
[**SchedulerTaskGet**](SchedulerTaskAPI.md#SchedulerTaskGet) | **Get** /api/v1/scheduler/tasks/{id} | Retrieve an existing scheduled task
[**SchedulerTaskList**](SchedulerTaskAPI.md#SchedulerTaskList) | **Get** /api/v1/scheduler/tasks | List the existing scheduled task(s)
[**SchedulerTaskRun**](SchedulerTaskAPI.md#SchedulerTaskRun) | **Get** /api/v1/scheduler/tasks/{id}/run | Run an existing scheduled task
[**SchedulerTaskUpdate**](SchedulerTaskAPI.md#SchedulerTaskUpdate) | **Put** /api/v1/scheduler/tasks | Update an existing scheduled task



## SchedulerTaskAdd

> SchedulerTaskList200ResponseInner SchedulerTaskAdd(ctx).SchedulerTaskAddRequest(schedulerTaskAddRequest).Execute()

Register a new scheduled task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	schedulerTaskAddRequest := openapiclient.scheduler_task_add_request{ReportScheduledTask: openapiclient.NewReportScheduledTask("Type_example", "Cron_example", false, "Name_example", []openapiclient.ReportRecipient{*openapiclient.NewReportRecipient("Type_example")}, "From_example", "Title_example", false, "HqlType_example")} // SchedulerTaskAddRequest | Scheduled task to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulerTaskAPI.SchedulerTaskAdd(context.Background()).SchedulerTaskAddRequest(schedulerTaskAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulerTaskAdd`: SchedulerTaskList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SchedulerTaskAPI.SchedulerTaskAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulerTaskAddRequest** | [**SchedulerTaskAddRequest**](SchedulerTaskAddRequest.md) | Scheduled task to register | 

### Return type

[**SchedulerTaskList200ResponseInner**](SchedulerTaskList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulerTaskDelete

> SchedulerTaskDelete(ctx, id).Execute()

Delete an existing scheduled task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SchedulerTaskAPI.SchedulerTaskDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulerTaskGet

> SchedulerTaskList200ResponseInner SchedulerTaskGet(ctx, id).Execute()

Retrieve an existing scheduled task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulerTaskAPI.SchedulerTaskGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulerTaskGet`: SchedulerTaskList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SchedulerTaskAPI.SchedulerTaskGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SchedulerTaskList200ResponseInner**](SchedulerTaskList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulerTaskList

> []SchedulerTaskList200ResponseInner SchedulerTaskList(ctx).ScheduledTaskType(scheduledTaskType).Execute()

List the existing scheduled task(s)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	scheduledTaskType := "scheduledTaskType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulerTaskAPI.SchedulerTaskList(context.Background()).ScheduledTaskType(scheduledTaskType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulerTaskList`: []SchedulerTaskList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SchedulerTaskAPI.SchedulerTaskList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduledTaskType** | **string** |  | 

### Return type

[**[]SchedulerTaskList200ResponseInner**](SchedulerTaskList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulerTaskRun

> SchedulerTaskRun(ctx, id).Execute()

Run an existing scheduled task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SchedulerTaskAPI.SchedulerTaskRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulerTaskUpdate

> SchedulerTaskUpdate200Response SchedulerTaskUpdate(ctx).SchedulerTaskUpdateRequest(schedulerTaskUpdateRequest).Execute()

Update an existing scheduled task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	schedulerTaskUpdateRequest := openapiclient.scheduler_task_update_request{ReportScheduledTask: openapiclient.NewReportScheduledTask("Type_example", "Cron_example", false, "Name_example", []openapiclient.ReportRecipient{*openapiclient.NewReportRecipient("Type_example")}, "From_example", "Title_example", false, "HqlType_example")} // SchedulerTaskUpdateRequest | Scheduled task to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulerTaskAPI.SchedulerTaskUpdate(context.Background()).SchedulerTaskUpdateRequest(schedulerTaskUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerTaskAPI.SchedulerTaskUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulerTaskUpdate`: SchedulerTaskUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `SchedulerTaskAPI.SchedulerTaskUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulerTaskUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulerTaskUpdateRequest** | [**SchedulerTaskUpdateRequest**](SchedulerTaskUpdateRequest.md) | Scheduled task to update | 

### Return type

[**SchedulerTaskUpdate200Response**](SchedulerTaskUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

