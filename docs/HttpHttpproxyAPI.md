# \HttpHttpproxyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HttpproxyAdd**](HttpHttpproxyAPI.md#HttpproxyAdd) | **Post** /api/v1/proxy/httpproxies | Register a new HTTP proxy
[**HttpproxyDelete**](HttpHttpproxyAPI.md#HttpproxyDelete) | **Delete** /api/v1/proxy/httpproxies/{name} | Delete an existing HTTP proxy
[**HttpproxyGet**](HttpHttpproxyAPI.md#HttpproxyGet) | **Get** /api/v1/proxy/httpproxies/{name} | Retrieve an existing HTTP proxy
[**HttpproxyList**](HttpHttpproxyAPI.md#HttpproxyList) | **Get** /api/v1/proxy/httpproxies | List the existing HTTP proxy(ies)
[**HttpproxyUpdate**](HttpHttpproxyAPI.md#HttpproxyUpdate) | **Put** /api/v1/proxy/httpproxies | Update an existing HTTP proxy



## HttpproxyAdd

> HttpProxyResponse HttpproxyAdd(ctx).HttpProxy(httpProxy).Execute()

Register a new HTTP proxy



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
	httpProxy := *openapiclient.NewHttpProxy("ExternalProxy", "36.52.145.12", int64(8888)) // HttpProxy | HTTP proxy to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HttpHttpproxyAPI.HttpproxyAdd(context.Background()).HttpProxy(httpProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HttpHttpproxyAPI.HttpproxyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HttpproxyAdd`: HttpProxyResponse
	fmt.Fprintf(os.Stdout, "Response from `HttpHttpproxyAPI.HttpproxyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHttpproxyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpProxy** | [**HttpProxy**](HttpProxy.md) | HTTP proxy to register | 

### Return type

[**HttpProxyResponse**](HttpProxyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HttpproxyDelete

> HttpproxyDelete(ctx, name).Execute()

Delete an existing HTTP proxy



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
	name := "name_example" // string | Name of the HTTP Proxy to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HttpHttpproxyAPI.HttpproxyDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HttpHttpproxyAPI.HttpproxyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the HTTP Proxy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiHttpproxyDeleteRequest struct via the builder pattern


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


## HttpproxyGet

> HttpProxyResponse HttpproxyGet(ctx, name).Execute()

Retrieve an existing HTTP proxy



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
	name := "name_example" // string | Name of the HTTP Proxy to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HttpHttpproxyAPI.HttpproxyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HttpHttpproxyAPI.HttpproxyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HttpproxyGet`: HttpProxyResponse
	fmt.Fprintf(os.Stdout, "Response from `HttpHttpproxyAPI.HttpproxyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the HTTP Proxy to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiHttpproxyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpProxyResponse**](HttpProxyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HttpproxyList

> []HttpProxyResponse HttpproxyList(ctx).Execute()

List the existing HTTP proxy(ies)



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
	resp, r, err := apiClient.HttpHttpproxyAPI.HttpproxyList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HttpHttpproxyAPI.HttpproxyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HttpproxyList`: []HttpProxyResponse
	fmt.Fprintf(os.Stdout, "Response from `HttpHttpproxyAPI.HttpproxyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHttpproxyListRequest struct via the builder pattern


### Return type

[**[]HttpProxyResponse**](HttpProxyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HttpproxyUpdate

> HttpProxyResponse HttpproxyUpdate(ctx).HttpProxy(httpProxy).Execute()

Update an existing HTTP proxy



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
	httpProxy := *openapiclient.NewHttpProxy("ExternalProxy", "36.52.145.12", int64(8888)) // HttpProxy | HTTP proxy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HttpHttpproxyAPI.HttpproxyUpdate(context.Background()).HttpProxy(httpProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HttpHttpproxyAPI.HttpproxyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HttpproxyUpdate`: HttpProxyResponse
	fmt.Fprintf(os.Stdout, "Response from `HttpHttpproxyAPI.HttpproxyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHttpproxyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpProxy** | [**HttpProxy**](HttpProxy.md) | HTTP proxy to update | 

### Return type

[**HttpProxyResponse**](HttpProxyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

