# \AutomationPolicyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationPolicyAdd**](AutomationPolicyAPI.md#AutomationPolicyAdd) | **Post** /api/v1/automation/policies | Register a new automation policy
[**AutomationPolicyDelete**](AutomationPolicyAPI.md#AutomationPolicyDelete) | **Delete** /api/v1/automation/policies/{name} | Delete an existing automation policy
[**AutomationPolicyGet**](AutomationPolicyAPI.md#AutomationPolicyGet) | **Get** /api/v1/automation/policies/{name} | Retrieve an existing automation policy
[**AutomationPolicyList**](AutomationPolicyAPI.md#AutomationPolicyList) | **Get** /api/v1/automation/policies | List the existing automation policies
[**AutomationPolicyUpdate**](AutomationPolicyAPI.md#AutomationPolicyUpdate) | **Put** /api/v1/automation/policies | Update an existing automation policy



## AutomationPolicyAdd

> AutomationPolicyResponse AutomationPolicyAdd(ctx).AutomationPolicy(automationPolicy).Execute()

Register a new automation policy



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
	automationPolicy := *openapiclient.NewAutomationPolicy("Name_example", "Profile_example") // AutomationPolicy | Automation policy to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationPolicyAPI.AutomationPolicyAdd(context.Background()).AutomationPolicy(automationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationPolicyAPI.AutomationPolicyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyAdd`: AutomationPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationPolicyAPI.AutomationPolicyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationPolicy** | [**AutomationPolicy**](AutomationPolicy.md) | Automation policy to register | 

### Return type

[**AutomationPolicyResponse**](AutomationPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationPolicyDelete

> AutomationPolicyDelete(ctx, name).Execute()

Delete an existing automation policy



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
	r, err := apiClient.AutomationPolicyAPI.AutomationPolicyDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationPolicyAPI.AutomationPolicyDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutomationPolicyDeleteRequest struct via the builder pattern


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


## AutomationPolicyGet

> AutomationPolicyResponse AutomationPolicyGet(ctx, name).Execute()

Retrieve an existing automation policy



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
	resp, r, err := apiClient.AutomationPolicyAPI.AutomationPolicyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationPolicyAPI.AutomationPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyGet`: AutomationPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationPolicyAPI.AutomationPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationPolicyResponse**](AutomationPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationPolicyList

> []AutomationPolicyResponse AutomationPolicyList(ctx).Execute()

List the existing automation policies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationPolicyAPI.AutomationPolicyList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationPolicyAPI.AutomationPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyList`: []AutomationPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationPolicyAPI.AutomationPolicyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyListRequest struct via the builder pattern


### Return type

[**[]AutomationPolicyResponse**](AutomationPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomationPolicyUpdate

> AutomationPolicyResponse AutomationPolicyUpdate(ctx).AutomationPolicy(automationPolicy).Execute()

Update an existing automation policy



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
	automationPolicy := *openapiclient.NewAutomationPolicy("Name_example", "Profile_example") // AutomationPolicy | Automation policy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationPolicyAPI.AutomationPolicyUpdate(context.Background()).AutomationPolicy(automationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationPolicyAPI.AutomationPolicyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationPolicyUpdate`: AutomationPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationPolicyAPI.AutomationPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationPolicy** | [**AutomationPolicy**](AutomationPolicy.md) | Automation policy to update | 

### Return type

[**AutomationPolicyResponse**](AutomationPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

