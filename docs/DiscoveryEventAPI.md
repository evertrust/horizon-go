# \DiscoveryEventAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryEventCsv**](DiscoveryEventAPI.md#DiscoveryEventCsv) | **Post** /api/v1/discovery/events/csv | Discovery event search
[**DiscoveryEventGet**](DiscoveryEventAPI.md#DiscoveryEventGet) | **Get** /api/v1/discovery/events/{id} | Retrieve a specific discovery event
[**DiscoveryEventSearch**](DiscoveryEventAPI.md#DiscoveryEventSearch) | **Post** /api/v1/discovery/events/search | Discovery event search



## DiscoveryEventCsv

> DiscoveryEventCsv(ctx).DiscoveryEventSearchQuery(discoveryEventSearchQuery).Execute()

Discovery event search



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
	discoveryEventSearchQuery := *openapiclient.NewDiscoveryEventSearchQuery() // DiscoveryEventSearchQuery | The discovery event search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryEventAPI.DiscoveryEventCsv(context.Background()).DiscoveryEventSearchQuery(discoveryEventSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAPI.DiscoveryEventCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryEventCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryEventSearchQuery** | [**DiscoveryEventSearchQuery**](DiscoveryEventSearchQuery.md) | The discovery event search query | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryEventGet

> DiscoveryEventResponse DiscoveryEventGet(ctx, id).Execute()

Retrieve a specific discovery event



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryEventAPI.DiscoveryEventGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAPI.DiscoveryEventGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryEventGet`: DiscoveryEventResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryEventAPI.DiscoveryEventGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryEventGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscoveryEventResponse**](DiscoveryEventResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryEventSearch

> EventSearchResultsResponse DiscoveryEventSearch(ctx).DiscoveryEventSearchQuery(discoveryEventSearchQuery).Execute()

Discovery event search



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
	discoveryEventSearchQuery := *openapiclient.NewDiscoveryEventSearchQuery() // DiscoveryEventSearchQuery | The discovery event search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryEventAPI.DiscoveryEventSearch(context.Background()).DiscoveryEventSearchQuery(discoveryEventSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEventAPI.DiscoveryEventSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryEventSearch`: EventSearchResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryEventAPI.DiscoveryEventSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryEventSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryEventSearchQuery** | [**DiscoveryEventSearchQuery**](DiscoveryEventSearchQuery.md) | The discovery event search query | 

### Return type

[**EventSearchResultsResponse**](EventSearchResultsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

