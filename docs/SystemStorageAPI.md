# \SystemStorageAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemStoragesAdd**](SystemStorageAPI.md#SystemStoragesAdd) | **Post** /api/v1/system/storages | Add a storage configuration
[**SystemStoragesDelete**](SystemStorageAPI.md#SystemStoragesDelete) | **Delete** /api/v1/system/storages/{name} | Delete a storage configuration
[**SystemStoragesGet**](SystemStorageAPI.md#SystemStoragesGet) | **Get** /api/v1/system/storages/{name} | Get a storage configuration
[**SystemStoragesList**](SystemStorageAPI.md#SystemStoragesList) | **Get** /api/v1/system/storages | List the existing storage configurations
[**SystemStoragesUpdate**](SystemStorageAPI.md#SystemStoragesUpdate) | **Put** /api/v1/system/storages | Update a storage configuration



## SystemStoragesAdd

> SystemStoragesList200ResponseInner SystemStoragesAdd(ctx).SystemStoragesUpdateRequest(systemStoragesUpdateRequest).Execute()

Add a storage configuration



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
	systemStoragesUpdateRequest := openapiclient.system_storages_update_request{S3StorageBackendConfig: openapiclient.NewS3StorageBackendConfig("Bucket_example", false, "Name_example", "PartBufferSize_example", "Timeout_example", "Type_example")} // SystemStoragesUpdateRequest | Storage to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemStorageAPI.SystemStoragesAdd(context.Background()).SystemStoragesUpdateRequest(systemStoragesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemStorageAPI.SystemStoragesAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemStoragesAdd`: SystemStoragesList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemStorageAPI.SystemStoragesAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemStoragesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemStoragesUpdateRequest** | [**SystemStoragesUpdateRequest**](SystemStoragesUpdateRequest.md) | Storage to add | 

### Return type

[**SystemStoragesList200ResponseInner**](SystemStoragesList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemStoragesDelete

> SystemStoragesDelete(ctx, name).Execute()

Delete a storage configuration



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
	name := "name_example" // string | Name of the storage to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemStorageAPI.SystemStoragesDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemStorageAPI.SystemStoragesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the storage to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemStoragesDeleteRequest struct via the builder pattern


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


## SystemStoragesGet

> SystemStoragesList200ResponseInner SystemStoragesGet(ctx, name).Execute()

Get a storage configuration



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
	name := "name_example" // string | Name of the storage to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemStorageAPI.SystemStoragesGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemStorageAPI.SystemStoragesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemStoragesGet`: SystemStoragesList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemStorageAPI.SystemStoragesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the storage to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemStoragesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemStoragesList200ResponseInner**](SystemStoragesList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemStoragesList

> []SystemStoragesList200ResponseInner SystemStoragesList(ctx).Execute()

List the existing storage configurations



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
	resp, r, err := apiClient.SystemStorageAPI.SystemStoragesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemStorageAPI.SystemStoragesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemStoragesList`: []SystemStoragesList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemStorageAPI.SystemStoragesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemStoragesListRequest struct via the builder pattern


### Return type

[**[]SystemStoragesList200ResponseInner**](SystemStoragesList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemStoragesUpdate

> SystemStoragesList200ResponseInner SystemStoragesUpdate(ctx).SystemStoragesUpdateRequest(systemStoragesUpdateRequest).Execute()

Update a storage configuration



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
	systemStoragesUpdateRequest := openapiclient.system_storages_update_request{S3StorageBackendConfig: openapiclient.NewS3StorageBackendConfig("Bucket_example", false, "Name_example", "PartBufferSize_example", "Timeout_example", "Type_example")} // SystemStoragesUpdateRequest | Storage to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemStorageAPI.SystemStoragesUpdate(context.Background()).SystemStoragesUpdateRequest(systemStoragesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemStorageAPI.SystemStoragesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemStoragesUpdate`: SystemStoragesList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SystemStorageAPI.SystemStoragesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemStoragesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemStoragesUpdateRequest** | [**SystemStoragesUpdateRequest**](SystemStoragesUpdateRequest.md) | Storage to update | 

### Return type

[**SystemStoragesList200ResponseInner**](SystemStoragesList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

