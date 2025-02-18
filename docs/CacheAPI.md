# \CacheAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheCrlGet**](CacheAPI.md#CacheCrlGet) | **Get** /api/v1/caches/crls/{ca} | Retrieve the CRL cache info for a specific certificate authority
[**CacheCrlList**](CacheAPI.md#CacheCrlList) | **Get** /api/v1/caches/crls | List the CRL cache info



## CacheCrlGet

> CachedCRLInfosResponse CacheCrlGet(ctx, ca).Execute()

Retrieve the CRL cache info for a specific certificate authority



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
	ca := "ca_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CacheAPI.CacheCrlGet(context.Background(), ca).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheAPI.CacheCrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheCrlGet`: CachedCRLInfosResponse
	fmt.Fprintf(os.Stdout, "Response from `CacheAPI.CacheCrlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ca** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheCrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CachedCRLInfosResponse**](CachedCRLInfosResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheCrlList

> []CachedCRLInfosResponse CacheCrlList(ctx).Execute()

List the CRL cache info



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
	resp, r, err := apiClient.CacheAPI.CacheCrlList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CacheAPI.CacheCrlList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheCrlList`: []CachedCRLInfosResponse
	fmt.Fprintf(os.Stdout, "Response from `CacheAPI.CacheCrlList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCacheCrlListRequest struct via the builder pattern


### Return type

[**[]CachedCRLInfosResponse**](CachedCRLInfosResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

