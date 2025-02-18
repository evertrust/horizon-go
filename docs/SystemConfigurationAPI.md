# \SystemConfigurationAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemConfigurationGet**](SystemConfigurationAPI.md#SystemConfigurationGet) | **Get** /api/v1/system/configuration/{type} | Upsert a system configuration
[**SystemConfigurationList**](SystemConfigurationAPI.md#SystemConfigurationList) | **Get** /api/v1/system/configuration | List the existing system configurations
[**SystemConfigurationUpsert**](SystemConfigurationAPI.md#SystemConfigurationUpsert) | **Put** /api/v1/system/configuration | Upsert a system configuration



## SystemConfigurationGet

> SystemConfigurationList200ResponseInner SystemConfigurationGet(ctx, type_).Execute()

Upsert a system configuration



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
	type_ := "type__example" // string | Type of the configuration entry to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigurationAPI.SystemConfigurationGet(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.SystemConfigurationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemConfigurationGet`: SystemConfigurationList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.SystemConfigurationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type of the configuration entry to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemConfigurationList200ResponseInner**](SystemConfigurationList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemConfigurationList

> []SystemConfigurationList200ResponseInner SystemConfigurationList(ctx).Execute()

List the existing system configurations



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
	resp, r, err := apiClient.SystemConfigurationAPI.SystemConfigurationList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.SystemConfigurationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemConfigurationList`: []SystemConfigurationList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.SystemConfigurationList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemConfigurationListRequest struct via the builder pattern


### Return type

[**[]SystemConfigurationList200ResponseInner**](SystemConfigurationList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemConfigurationUpsert

> SystemConfigurationList200ResponseInner SystemConfigurationUpsert(ctx).SystemConfigurationUpsertRequest(systemConfigurationUpsertRequest).Execute()

Upsert a system configuration



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
	systemConfigurationUpsertRequest := openapiclient.system_configuration_upsert_request{InterfaceCustomizationConfiguration: openapiclient.NewInterfaceCustomizationConfiguration("Type_example")} // SystemConfigurationUpsertRequest | System configuration entry to upsert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigurationAPI.SystemConfigurationUpsert(context.Background()).SystemConfigurationUpsertRequest(systemConfigurationUpsertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigurationAPI.SystemConfigurationUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemConfigurationUpsert`: SystemConfigurationList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigurationAPI.SystemConfigurationUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemConfigurationUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemConfigurationUpsertRequest** | [**SystemConfigurationUpsertRequest**](SystemConfigurationUpsertRequest.md) | System configuration entry to upsert | 

### Return type

[**SystemConfigurationList200ResponseInner**](SystemConfigurationList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

