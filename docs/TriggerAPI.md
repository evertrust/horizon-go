# \TriggerAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TriggerAdd**](TriggerAPI.md#TriggerAdd) | **Post** /api/v1/triggers | Register a new trigger
[**TriggerDelete**](TriggerAPI.md#TriggerDelete) | **Delete** /api/v1/triggers/{name} | Delete an existing trigger
[**TriggerGet**](TriggerAPI.md#TriggerGet) | **Get** /api/v1/triggers/{name} | Retrieve an existing trigger
[**TriggerList**](TriggerAPI.md#TriggerList) | **Get** /api/v1/triggers | List the existing trigger(s)
[**TriggerTest**](TriggerAPI.md#TriggerTest) | **Patch** /api/v1/triggers | Test a trigger
[**TriggerUpdate**](TriggerAPI.md#TriggerUpdate) | **Put** /api/v1/triggers | Update an existing trigger



## TriggerAdd

> TriggerUpdate200Response TriggerAdd(ctx).TriggerUpdateRequest(triggerUpdateRequest).Execute()

Register a new trigger



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
	triggerUpdateRequest := openapiclient.trigger_update_request{AWSTrigger: openapiclient.NewAWSTrigger("Connector_example", "Name_example", "Type_example")} // TriggerUpdateRequest | The trigger to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.TriggerAdd(context.Background()).TriggerUpdateRequest(triggerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerAdd`: TriggerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerUpdateRequest** | [**TriggerUpdateRequest**](TriggerUpdateRequest.md) | The trigger to register | 

### Return type

[**TriggerUpdate200Response**](TriggerUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerDelete

> TriggerDelete(ctx, name).Execute()

Delete an existing trigger



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
	r, err := apiClient.TriggerAPI.TriggerDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTriggerDeleteRequest struct via the builder pattern


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


## TriggerGet

> TriggerUpdate200Response TriggerGet(ctx, name).Execute()

Retrieve an existing trigger



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
	resp, r, err := apiClient.TriggerAPI.TriggerGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerGet`: TriggerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TriggerUpdate200Response**](TriggerUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerList

> []TriggerList200ResponseInner TriggerList(ctx).Types(types).Module(module).Execute()

List the existing trigger(s)



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
	types := []string{"Types_example"} // []string |  (optional)
	module := "module_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.TriggerList(context.Background()).Types(types).Module(module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerList`: []TriggerList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]string** |  | 
 **module** | **string** |  | 

### Return type

[**[]TriggerList200ResponseInner**](TriggerList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerTest

> TriggerTest200Response TriggerTest(ctx).TriggerTestRequest(triggerTestRequest).Execute()

Test a trigger



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
	triggerTestRequest := *openapiclient.NewTriggerTestRequest(openapiclient.TriggerTestRequest_trigger{EmailNotification: openapiclient.NewEmailNotification(*openapiclient.NewEmailTemplate("noreply@horizon.evertrust.fr", false, "Password recovery", []openapiclient.EmailRecipient{*openapiclient.NewEmailRecipient("Type_example")}), "Type_example", []string{"Events_example"}, "NOTIFICATION_ENROLL")}) // TriggerTestRequest | Trigger to test and its dictionary

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.TriggerTest(context.Background()).TriggerTestRequest(triggerTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerTest`: TriggerTest200Response
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerTestRequest** | [**TriggerTestRequest**](TriggerTestRequest.md) | Trigger to test and its dictionary | 

### Return type

[**TriggerTest200Response**](TriggerTest200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerUpdate

> TriggerUpdate200Response TriggerUpdate(ctx).TriggerUpdateRequest(triggerUpdateRequest).Execute()

Update an existing trigger



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
	triggerUpdateRequest := openapiclient.trigger_update_request{AWSTrigger: openapiclient.NewAWSTrigger("Connector_example", "Name_example", "Type_example")} // TriggerUpdateRequest | Trigger to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.TriggerUpdate(context.Background()).TriggerUpdateRequest(triggerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerUpdate`: TriggerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerUpdateRequest** | [**TriggerUpdateRequest**](TriggerUpdateRequest.md) | Trigger to update | 

### Return type

[**TriggerUpdate200Response**](TriggerUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

