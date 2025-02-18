# \AutomationExecutionAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationExecutionAdd**](AutomationExecutionAPI.md#AutomationExecutionAdd) | **Post** /api/v1/automation/executions | Register a new execution policy
[**AutomationExecutionDelete**](AutomationExecutionAPI.md#AutomationExecutionDelete) | **Delete** /api/v1/automation/executions/{name} | Delete an existing execution policy
[**AutomationExecutionGet**](AutomationExecutionAPI.md#AutomationExecutionGet) | **Get** /api/v1/automation/executions/{name} | Retrieve an existing execution policy
[**AutomationExecutionList**](AutomationExecutionAPI.md#AutomationExecutionList) | **Get** /api/v1/automation/executions | List the existing execution policies
[**AutomationExecutionUpdate**](AutomationExecutionAPI.md#AutomationExecutionUpdate) | **Put** /api/v1/automation/executions | Update an existing execution policy



## AutomationExecutionAdd

> ExecutionPolicyResponse AutomationExecutionAdd(ctx).ExecutionPolicy(executionPolicy).Execute()

Register a new execution policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {
	executionPolicy := *openapiclient.NewExecutionPolicy("Name_example") // ExecutionPolicy | Execution policy to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationExecutionAPI.AutomationExecutionAdd(context.Background()).ExecutionPolicy(executionPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationExecutionAPI.AutomationExecutionAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationExecutionAdd`: ExecutionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationExecutionAPI.AutomationExecutionAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationExecutionAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionPolicy** | [**ExecutionPolicy**](ExecutionPolicy.md) | Execution policy to register | 

### Return type

[**ExecutionPolicyResponse**](ExecutionPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationExecutionDelete

> AutomationExecutionDelete(ctx, name).Execute()

Delete an existing execution policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomationExecutionAPI.AutomationExecutionDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationExecutionAPI.AutomationExecutionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationExecutionDeleteRequest struct via the builder pattern


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


## AutomationExecutionGet

> ExecutionPolicyResponse AutomationExecutionGet(ctx, name).Execute()

Retrieve an existing execution policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationExecutionAPI.AutomationExecutionGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationExecutionAPI.AutomationExecutionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationExecutionGet`: ExecutionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationExecutionAPI.AutomationExecutionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationExecutionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecutionPolicyResponse**](ExecutionPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationExecutionList

> []ExecutionPolicyResponse AutomationExecutionList(ctx).Execute()

List the existing execution policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationExecutionAPI.AutomationExecutionList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationExecutionAPI.AutomationExecutionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationExecutionList`: []ExecutionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationExecutionAPI.AutomationExecutionList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationExecutionListRequest struct via the builder pattern


### Return type

[**[]ExecutionPolicyResponse**](ExecutionPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationExecutionUpdate

> ExecutionPolicyResponse AutomationExecutionUpdate(ctx).ExecutionPolicy(executionPolicy).Execute()

Update an existing execution policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {
	executionPolicy := *openapiclient.NewExecutionPolicy("Name_example") // ExecutionPolicy | Execution policy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationExecutionAPI.AutomationExecutionUpdate(context.Background()).ExecutionPolicy(executionPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationExecutionAPI.AutomationExecutionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationExecutionUpdate`: ExecutionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationExecutionAPI.AutomationExecutionUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationExecutionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionPolicy** | [**ExecutionPolicy**](ExecutionPolicy.md) | Execution policy to update | 

### Return type

[**ExecutionPolicyResponse**](ExecutionPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

