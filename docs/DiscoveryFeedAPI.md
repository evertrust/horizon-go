# \DiscoveryFeedAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryFeed**](DiscoveryFeedAPI.md#DiscoveryFeed) | **Post** /api/v1/discovery/feed | Feed a discovered certificate into a discovery campaign
[**DiscoveryFeedEventRegister**](DiscoveryFeedAPI.md#DiscoveryFeedEventRegister) | **Put** /api/v1/discovery/feed | Push a new discovery event
[**DiscoveryFeedSessionEnd**](DiscoveryFeedAPI.md#DiscoveryFeedSessionEnd) | **Delete** /api/v1/discovery/feed/{campaign}/{id} | End a discovery session
[**DiscoveryFeedSessionStart**](DiscoveryFeedAPI.md#DiscoveryFeedSessionStart) | **Get** /api/v1/discovery/feed/{name} | Create a new discovery feed session



## DiscoveryFeed

> DiscoveryFeed(ctx).DiscoveryFeed(discoveryFeed).Execute()

Feed a discovered certificate into a discovery campaign



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
	discoveryFeed := *openapiclient.NewDiscoveryFeed("Discovery-DMZ01", "-----BEGIN CERTIFICATE-----...", *openapiclient.NewHostDiscoveryData()) // DiscoveryFeed | The discovery feed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryFeedAPI.DiscoveryFeed(context.Background()).DiscoveryFeed(discoveryFeed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryFeedAPI.DiscoveryFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryFeed** | [**DiscoveryFeed**](DiscoveryFeed.md) | The discovery feed | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryFeedEventRegister

> DiscoveryFeedEventRegister(ctx).DiscoveryEvent(discoveryEvent).Execute()

Push a new discovery event



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
	discoveryEvent := *openapiclient.NewDiscoveryEvent("Discovery-DMZ01", "NETSCAN", "failure") // DiscoveryEvent | The discovery event to push

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryFeedAPI.DiscoveryFeedEventRegister(context.Background()).DiscoveryEvent(discoveryEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryFeedAPI.DiscoveryFeedEventRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryFeedEventRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryEvent** | [**DiscoveryEvent**](DiscoveryEvent.md) | The discovery event to push | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryFeedSessionEnd

> DiscoveryFeedSessionEnd(ctx, campaign, id).Execute()

End a discovery session



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
	campaign := "Discovery-DMZ01" // string | The name of the discovery campaign where the session was started
	id := "63fce2e13000003c008797c4" // string | The discovery session ID to end

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryFeedAPI.DiscoveryFeedSessionEnd(context.Background(), campaign, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryFeedAPI.DiscoveryFeedSessionEnd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaign** | **string** | The name of the discovery campaign where the session was started | 
**id** | **string** | The discovery session ID to end | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryFeedSessionEndRequest struct via the builder pattern


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


## DiscoveryFeedSessionStart

> DiscoveryFeedSessionResponse DiscoveryFeedSessionStart(ctx, name).Execute()

Create a new discovery feed session



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
	name := "name_example" // string | The name of the discovery campaign to create a new feed session for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryFeedAPI.DiscoveryFeedSessionStart(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryFeedAPI.DiscoveryFeedSessionStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryFeedSessionStart`: DiscoveryFeedSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryFeedAPI.DiscoveryFeedSessionStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the discovery campaign to create a new feed session for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryFeedSessionStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscoveryFeedSessionResponse**](DiscoveryFeedSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

