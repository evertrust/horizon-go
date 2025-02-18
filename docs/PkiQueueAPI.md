# \PkiQueueAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PkiqueueAdd**](PkiQueueAPI.md#PkiqueueAdd) | **Post** /api/v1/pki/queues | Register a new pki queue
[**PkiqueueDelete**](PkiQueueAPI.md#PkiqueueDelete) | **Delete** /api/v1/pki/queues/{name} | Delete an existing pki queue
[**PkiqueueGet**](PkiQueueAPI.md#PkiqueueGet) | **Get** /api/v1/pki/queues/{name} | Retrieve an existing pki queue
[**PkiqueueList**](PkiQueueAPI.md#PkiqueueList) | **Get** /api/v1/pki/queues | List the existing pki queue(s)
[**PkiqueueUpdate**](PkiQueueAPI.md#PkiqueueUpdate) | **Put** /api/v1/pki/queues | Update an existing pki queue



## PkiqueueAdd

> PKIQueueResponse PkiqueueAdd(ctx).PKIQueue(pKIQueue).Execute()

Register a new pki queue



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
	pKIQueue := *openapiclient.NewPKIQueue("Name_example", false, int64(123)) // PKIQueue | The pki queue to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiQueueAPI.PkiqueueAdd(context.Background()).PKIQueue(pKIQueue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiQueueAPI.PkiqueueAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiqueueAdd`: PKIQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `PkiQueueAPI.PkiqueueAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkiqueueAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIQueue** | [**PKIQueue**](PKIQueue.md) | The pki queue to register | 

### Return type

[**PKIQueueResponse**](PKIQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiqueueDelete

> PkiqueueDelete(ctx, name).Execute()

Delete an existing pki queue



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PkiQueueAPI.PkiqueueDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiQueueAPI.PkiqueueDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPkiqueueDeleteRequest struct via the builder pattern


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


## PkiqueueGet

> PKIQueueResponse PkiqueueGet(ctx, name).Execute()

Retrieve an existing pki queue



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiQueueAPI.PkiqueueGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiQueueAPI.PkiqueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiqueueGet`: PKIQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `PkiQueueAPI.PkiqueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPkiqueueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PKIQueueResponse**](PKIQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiqueueList

> []PKIQueueResponse PkiqueueList(ctx).Execute()

List the existing pki queue(s)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiQueueAPI.PkiqueueList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiQueueAPI.PkiqueueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiqueueList`: []PKIQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `PkiQueueAPI.PkiqueueList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPkiqueueListRequest struct via the builder pattern


### Return type

[**[]PKIQueueResponse**](PKIQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiqueueUpdate

> PKIQueueResponse PkiqueueUpdate(ctx).PKIQueue(pKIQueue).Execute()

Update an existing pki queue



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
	pKIQueue := *openapiclient.NewPKIQueue("Name_example", false, int64(123)) // PKIQueue | The pki queue to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiQueueAPI.PkiqueueUpdate(context.Background()).PKIQueue(pKIQueue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiQueueAPI.PkiqueueUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiqueueUpdate`: PKIQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `PkiQueueAPI.PkiqueueUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkiqueueUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pKIQueue** | [**PKIQueue**](PKIQueue.md) | The pki queue to update | 

### Return type

[**PKIQueueResponse**](PKIQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

