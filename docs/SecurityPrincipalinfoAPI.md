# \SecurityPrincipalinfoAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScurityPrincipalInfoGet**](SecurityPrincipalinfoAPI.md#ScurityPrincipalInfoGet) | **Get** /api/v1/security/principalinfos/{identifier} | Retrieve a principal information
[**SecurityPrincipalInfoAdd**](SecurityPrincipalinfoAPI.md#SecurityPrincipalInfoAdd) | **Post** /api/v1/security/principalinfos | Create a new principal
[**SecurityPrincipalInfoDelete**](SecurityPrincipalinfoAPI.md#SecurityPrincipalInfoDelete) | **Delete** /api/v1/security/principalinfos/{identifier} | Delete a principal
[**SecurityPrincipalInfoSearch**](SecurityPrincipalinfoAPI.md#SecurityPrincipalInfoSearch) | **Post** /api/v1/security/principalinfos/search | Search for principal information
[**SecurityPrincipalInfoUpdate**](SecurityPrincipalinfoAPI.md#SecurityPrincipalInfoUpdate) | **Put** /api/v1/security/principalinfos | Update a principal&#39;s information



## ScurityPrincipalInfoGet

> PrincipalInfoResponse ScurityPrincipalInfoGet(ctx, identifier).Execute()

Retrieve a principal information



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
	identifier := "identifier_example" // string | The identifier of the principal to retrieve information of

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPrincipalinfoAPI.ScurityPrincipalInfoGet(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalinfoAPI.ScurityPrincipalInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScurityPrincipalInfoGet`: PrincipalInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPrincipalinfoAPI.ScurityPrincipalInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The identifier of the principal to retrieve information of | 

### Other Parameters

Other parameters are passed through a pointer to a apiScurityPrincipalInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalInfoResponse**](PrincipalInfoResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalInfoAdd

> PrincipalInfoResponse SecurityPrincipalInfoAdd(ctx).PrincipalInfo(principalInfo).Execute()

Create a new principal



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
	principalInfo := *openapiclient.NewPrincipalInfo(true, "administrator") // PrincipalInfo | The principal's information to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPrincipalinfoAPI.SecurityPrincipalInfoAdd(context.Background()).PrincipalInfo(principalInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalinfoAPI.SecurityPrincipalInfoAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalInfoAdd`: PrincipalInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPrincipalinfoAPI.SecurityPrincipalInfoAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalInfoAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalInfo** | [**PrincipalInfo**](PrincipalInfo.md) | The principal&#39;s information to register | 

### Return type

[**PrincipalInfoResponse**](PrincipalInfoResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalInfoDelete

> SecurityPrincipalInfoDelete(ctx, identifier).Execute()

Delete a principal



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
	identifier := "identifier_example" // string | The identifier of the principal to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityPrincipalinfoAPI.SecurityPrincipalInfoDelete(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalinfoAPI.SecurityPrincipalInfoDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | The identifier of the principal to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalInfoDeleteRequest struct via the builder pattern


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


## SecurityPrincipalInfoSearch

> PrincipalInfoSearchResultsResponse SecurityPrincipalInfoSearch(ctx).PrincipalInfoSearchQuery(principalInfoSearchQuery).Execute()

Search for principal information



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
	principalInfoSearchQuery := *openapiclient.NewPrincipalInfoSearchQuery() // PrincipalInfoSearchQuery | The principal information search request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPrincipalinfoAPI.SecurityPrincipalInfoSearch(context.Background()).PrincipalInfoSearchQuery(principalInfoSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalinfoAPI.SecurityPrincipalInfoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalInfoSearch`: PrincipalInfoSearchResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPrincipalinfoAPI.SecurityPrincipalInfoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalInfoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalInfoSearchQuery** | [**PrincipalInfoSearchQuery**](PrincipalInfoSearchQuery.md) | The principal information search request | 

### Return type

[**PrincipalInfoSearchResultsResponse**](PrincipalInfoSearchResultsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalInfoUpdate

> PrincipalInfoResponse SecurityPrincipalInfoUpdate(ctx).PrincipalInfo(principalInfo).Execute()

Update a principal's information



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
	principalInfo := *openapiclient.NewPrincipalInfo(true, "administrator") // PrincipalInfo | The principal information to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPrincipalinfoAPI.SecurityPrincipalInfoUpdate(context.Background()).PrincipalInfo(principalInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalinfoAPI.SecurityPrincipalInfoUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalInfoUpdate`: PrincipalInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPrincipalinfoAPI.SecurityPrincipalInfoUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalInfoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalInfo** | [**PrincipalInfo**](PrincipalInfo.md) | The principal information to update | 

### Return type

[**PrincipalInfoResponse**](PrincipalInfoResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

