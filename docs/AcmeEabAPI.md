# \AcmeEabAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcmeEabAdd**](AcmeEabAPI.md#AcmeEabAdd) | **Post** /api/v1/acme/eab | Add a new External Account Binding
[**AcmeEabDelete**](AcmeEabAPI.md#AcmeEabDelete) | **Delete** /api/v1/acme/eab/{name} | Delete an External Account Binding by name
[**AcmeEabGet**](AcmeEabAPI.md#AcmeEabGet) | **Get** /api/v1/acme/eab/{name} | Get an External Account Binding by name
[**AcmeEabRenew**](AcmeEabAPI.md#AcmeEabRenew) | **Post** /api/v1/acme/eab/{name}/renew | Renew an External Account Binding
[**AcmeEabSearch**](AcmeEabAPI.md#AcmeEabSearch) | **Post** /api/v1/acme/eab/search | Search External Account Bindings
[**AcmeEabUpdate**](AcmeEabAPI.md#AcmeEabUpdate) | **Put** /api/v1/acme/eab | Update an existing External Account Binding (metadata only)
[**AcmeEabUpdateStatus**](AcmeEabAPI.md#AcmeEabUpdateStatus) | **Post** /api/v1/acme/eab/{eabName}/status | Update External Account Binding status



## AcmeEabAdd

> ExternalAccountBindingResponse AcmeEabAdd(ctx).EabAddRequest(eabAddRequest).Execute()

Add a new External Account Binding



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
	eabAddRequest := *openapiclient.NewEabAddRequest("EabPolicy_example", openapiclient.ExternalAccountBindingAlgorithm("HS256"), "Name_example") // EabAddRequest | The External Account Binding to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabAdd(context.Background()).EabAddRequest(eabAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabAdd`: ExternalAccountBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eabAddRequest** | [**EabAddRequest**](EabAddRequest.md) | The External Account Binding to add | 

### Return type

[**ExternalAccountBindingResponse**](ExternalAccountBindingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabDelete

> AcmeEabDelete(ctx, name).Execute()

Delete an External Account Binding by name



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
	name := "name_example" // string | The name of the External Account Binding

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AcmeEabAPI.AcmeEabDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the External Account Binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabDeleteRequest struct via the builder pattern


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


## AcmeEabGet

> ExternalAccountBindingResponse AcmeEabGet(ctx, name).Execute()

Get an External Account Binding by name



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
	name := "name_example" // string | The name of the External Account Binding

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabGet`: ExternalAccountBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the External Account Binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAccountBindingResponse**](ExternalAccountBindingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabRenew

> ExternalAccountBindingResponse AcmeEabRenew(ctx, name).EabRenewRequest(eabRenewRequest).Execute()

Renew an External Account Binding



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
	name := "name_example" // string | The name of the External Account Binding to renew
	eabRenewRequest := *openapiclient.NewEabRenewRequest() // EabRenewRequest | Renewal parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabRenew(context.Background(), name).EabRenewRequest(eabRenewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabRenew`: ExternalAccountBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabRenew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the External Account Binding to renew | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eabRenewRequest** | [**EabRenewRequest**](EabRenewRequest.md) | Renewal parameters | 

### Return type

[**ExternalAccountBindingResponse**](ExternalAccountBindingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabSearch

> ExternalAccountBindingSearchResults AcmeEabSearch(ctx).ExternalAccountBindingSearch(externalAccountBindingSearch).Execute()

Search External Account Bindings



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
	externalAccountBindingSearch := *openapiclient.NewExternalAccountBindingSearch() // ExternalAccountBindingSearch | Search parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabSearch(context.Background()).ExternalAccountBindingSearch(externalAccountBindingSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabSearch`: ExternalAccountBindingSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalAccountBindingSearch** | [**ExternalAccountBindingSearch**](ExternalAccountBindingSearch.md) | Search parameters | 

### Return type

[**ExternalAccountBindingSearchResults**](ExternalAccountBindingSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabUpdate

> ExternalAccountBindingResponse AcmeEabUpdate(ctx).EabUpdateRequest(eabUpdateRequest).Execute()

Update an existing External Account Binding (metadata only)



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
	eabUpdateRequest := *openapiclient.NewEabUpdateRequest("Name_example") // EabUpdateRequest | External Account Binding to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabUpdate(context.Background()).EabUpdateRequest(eabUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabUpdate`: ExternalAccountBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eabUpdateRequest** | [**EabUpdateRequest**](EabUpdateRequest.md) | External Account Binding to update | 

### Return type

[**ExternalAccountBindingResponse**](ExternalAccountBindingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabUpdateStatus

> ExternalAccountBindingResponse AcmeEabUpdateStatus(ctx, eabName).EabStatusUpdateRequest(eabStatusUpdateRequest).Execute()

Update External Account Binding status



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
	eabName := "eabName_example" // string | The name of the External Account Binding
	eabStatusUpdateRequest := *openapiclient.NewEabStatusUpdateRequest(openapiclient.ExternalAccountBindingStatus("valid")) // EabStatusUpdateRequest | Status and optional revocation date for certificates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabAPI.AcmeEabUpdateStatus(context.Background(), eabName).EabStatusUpdateRequest(eabStatusUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabAPI.AcmeEabUpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabUpdateStatus`: ExternalAccountBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabAPI.AcmeEabUpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eabName** | **string** | The name of the External Account Binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eabStatusUpdateRequest** | [**EabStatusUpdateRequest**](EabStatusUpdateRequest.md) | Status and optional revocation date for certificates | 

### Return type

[**ExternalAccountBindingResponse**](ExternalAccountBindingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

