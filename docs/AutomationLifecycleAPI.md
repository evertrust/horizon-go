# \AutomationLifecycleAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationPolicyLifecycleGet**](AutomationLifecycleAPI.md#AutomationPolicyLifecycleGet) | **Get** /api/v1/automation/lifecycle/{name} | Retrieve the enroll materials
[**AutomationPolicyLifecycleVerify**](AutomationLifecycleAPI.md#AutomationPolicyLifecycleVerify) | **Get** /api/v1/automation/lifecycle/{name}/verify | Verify the certificate against the Automation policy



## AutomationPolicyLifecycleGet

> AutomationPolicyLifecycleGet201Response AutomationPolicyLifecycleGet(ctx, name).Execute()

Retrieve the enroll materials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/v2"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationLifecycleAPI.AutomationPolicyLifecycleGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationLifecycleAPI.AutomationPolicyLifecycleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyLifecycleGet`: AutomationPolicyLifecycleGet201Response
	fmt.Fprintf(os.Stdout, "Response from `AutomationLifecycleAPI.AutomationPolicyLifecycleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyLifecycleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationPolicyLifecycleGet201Response**](AutomationPolicyLifecycleGet201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationPolicyLifecycleVerify

> AutomationReportResponse AutomationPolicyLifecycleVerify(ctx, name).Execute()

Verify the certificate against the Automation policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/v2"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationLifecycleAPI.AutomationPolicyLifecycleVerify(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationLifecycleAPI.AutomationPolicyLifecycleVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyLifecycleVerify`: AutomationReportResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationLifecycleAPI.AutomationPolicyLifecycleVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyLifecycleVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationReportResponse**](AutomationReportResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

