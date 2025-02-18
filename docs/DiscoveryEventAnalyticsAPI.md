# \DiscoveryEventAnalyticsAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsDiscoveryEventFlush**](DiscoveryEventAnalyticsAPI.md#AnalyticsDiscoveryEventFlush) | **Delete** /api/v1/analytics/discovery/events | Flush discovery event analytics synchronization
[**AnalyticsDiscoveryEventGet**](DiscoveryEventAnalyticsAPI.md#AnalyticsDiscoveryEventGet) | **Get** /api/v1/analytics/discovery/events | Retrieve the discovery event analytics status
[**AnalyticsDiscoveryEventUpdate**](DiscoveryEventAnalyticsAPI.md#AnalyticsDiscoveryEventUpdate) | **Patch** /api/v1/analytics/discovery/events | Schedule a new discovery event analytics synchronization



## AnalyticsDiscoveryEventFlush

> AnalyticsDiscoveryEventFlush(ctx).Execute()

Flush discovery event analytics synchronization



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
	r, err := apiClient.DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventFlush(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventFlush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDiscoveryEventFlushRequest struct via the builder pattern


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


## AnalyticsDiscoveryEventGet

> []AnalyticsStatus2 AnalyticsDiscoveryEventGet(ctx).Execute()

Retrieve the discovery event analytics status



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
	resp, r, err := apiClient.DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDiscoveryEventGet`: []AnalyticsStatus2
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDiscoveryEventGetRequest struct via the builder pattern


### Return type

[**[]AnalyticsStatus2**](AnalyticsStatus2.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsDiscoveryEventUpdate

> AnalyticsDiscoveryEventUpdate(ctx).Execute()

Schedule a new discovery event analytics synchronization



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
	r, err := apiClient.DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAnalyticsAPI.AnalyticsDiscoveryEventUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDiscoveryEventUpdateRequest struct via the builder pattern


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

