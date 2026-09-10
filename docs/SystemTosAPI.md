# \SystemTosAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemTosAdd**](SystemTosAPI.md#SystemTosAdd) | **Post** /api/v1/system/terms-of-services | Add a Terms of Service entry
[**SystemTosDelete**](SystemTosAPI.md#SystemTosDelete) | **Delete** /api/v1/system/terms-of-services/{name} | Delete a Terms of Service entry
[**SystemTosGet**](SystemTosAPI.md#SystemTosGet) | **Get** /api/v1/system/terms-of-services/{name} | Get a Terms of Service entry
[**SystemTosList**](SystemTosAPI.md#SystemTosList) | **Get** /api/v1/system/terms-of-services | List the existing Terms of Service entries
[**SystemTosUpdate**](SystemTosAPI.md#SystemTosUpdate) | **Put** /api/v1/system/terms-of-services | Update a Terms of Service entry



## SystemTosAdd

> TermsOfServiceResponse SystemTosAdd(ctx).TermsOfService(termsOfService).Execute()

Add a Terms of Service entry



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
	termsOfService := *openapiclient.NewTermsOfService([]openapiclient.LocalizedString{*openapiclient.NewLocalizedString("en", "Value In English")}, "default-tos") // TermsOfService | Terms of Service to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemTosAPI.SystemTosAdd(context.Background()).TermsOfService(termsOfService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemTosAPI.SystemTosAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemTosAdd`: TermsOfServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemTosAPI.SystemTosAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemTosAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **termsOfService** | [**TermsOfService**](TermsOfService.md) | Terms of Service to add | 

### Return type

[**TermsOfServiceResponse**](TermsOfServiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemTosDelete

> SystemTosDelete(ctx, name).Execute()

Delete a Terms of Service entry



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
	name := "name_example" // string | Name of the Terms of Service to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemTosAPI.SystemTosDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemTosAPI.SystemTosDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Terms of Service to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemTosDeleteRequest struct via the builder pattern


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


## SystemTosGet

> TermsOfServiceResponse SystemTosGet(ctx, name).Execute()

Get a Terms of Service entry



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
	name := "name_example" // string | Name of the Terms of Service to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemTosAPI.SystemTosGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemTosAPI.SystemTosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemTosGet`: TermsOfServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemTosAPI.SystemTosGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the Terms of Service to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemTosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TermsOfServiceResponse**](TermsOfServiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemTosList

> []TermsOfServiceResponse SystemTosList(ctx).Execute()

List the existing Terms of Service entries



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
	resp, r, err := apiClient.SystemTosAPI.SystemTosList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemTosAPI.SystemTosList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemTosList`: []TermsOfServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemTosAPI.SystemTosList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemTosListRequest struct via the builder pattern


### Return type

[**[]TermsOfServiceResponse**](TermsOfServiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemTosUpdate

> TermsOfServiceResponse SystemTosUpdate(ctx).TermsOfService(termsOfService).Execute()

Update a Terms of Service entry



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
	termsOfService := *openapiclient.NewTermsOfService([]openapiclient.LocalizedString{*openapiclient.NewLocalizedString("en", "Value In English")}, "default-tos") // TermsOfService | Terms of Service to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemTosAPI.SystemTosUpdate(context.Background()).TermsOfService(termsOfService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemTosAPI.SystemTosUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemTosUpdate`: TermsOfServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemTosAPI.SystemTosUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemTosUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **termsOfService** | [**TermsOfService**](TermsOfService.md) | Terms of Service to update | 

### Return type

[**TermsOfServiceResponse**](TermsOfServiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

