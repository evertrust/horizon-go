# \WcceAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WcceForestMappingAdd**](WcceAPI.md#WcceForestMappingAdd) | **Post** /api/v1/wcce/forests | Register a new WCCE forest mapping
[**WcceForestMappingDelete**](WcceAPI.md#WcceForestMappingDelete) | **Delete** /api/v1/wcce/forests/{name} | Delete an existing WCCE forest mapping
[**WcceForestMappingGet**](WcceAPI.md#WcceForestMappingGet) | **Get** /api/v1/wcce/forests/{name} | Retrieve an existing WCCE forest mapping
[**WcceForestMappingList**](WcceAPI.md#WcceForestMappingList) | **Get** /api/v1/wcce/forests | List the existing WCCE forest mapping(s)
[**WcceForestMappingUpdate**](WcceAPI.md#WcceForestMappingUpdate) | **Put** /api/v1/wcce/forests | Update an existing WCCE forest mapping



## WcceForestMappingAdd

> WcceForestMappingResponse WcceForestMappingAdd(ctx).WcceForestMapping(wcceForestMapping).Execute()

Register a new WCCE forest mapping



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
	wcceForestMapping := *openapiclient.NewWcceForestMapping("Forest_example", []openapiclient.WcceTemplateMapping{*openapiclient.NewWcceTemplateMapping("EnrollmentMode_example", "Profile_example", "Template_example")}) // WcceForestMapping | The WCCE forest mapping to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WcceAPI.WcceForestMappingAdd(context.Background()).WcceForestMapping(wcceForestMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WcceAPI.WcceForestMappingAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WcceForestMappingAdd`: WcceForestMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `WcceAPI.WcceForestMappingAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWcceForestMappingAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wcceForestMapping** | [**WcceForestMapping**](WcceForestMapping.md) | The WCCE forest mapping to register | 

### Return type

[**WcceForestMappingResponse**](WcceForestMappingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WcceForestMappingDelete

> WcceForestMappingDelete(ctx, name).Execute()

Delete an existing WCCE forest mapping



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WcceAPI.WcceForestMappingDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WcceAPI.WcceForestMappingDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWcceForestMappingDeleteRequest struct via the builder pattern


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


## WcceForestMappingGet

> WcceForestMappingResponse WcceForestMappingGet(ctx, name).Execute()

Retrieve an existing WCCE forest mapping



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WcceAPI.WcceForestMappingGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WcceAPI.WcceForestMappingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WcceForestMappingGet`: WcceForestMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `WcceAPI.WcceForestMappingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWcceForestMappingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WcceForestMappingResponse**](WcceForestMappingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WcceForestMappingList

> []WcceForestMappingResponse WcceForestMappingList(ctx).Execute()

List the existing WCCE forest mapping(s)



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
	resp, r, err := apiClient.WcceAPI.WcceForestMappingList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WcceAPI.WcceForestMappingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WcceForestMappingList`: []WcceForestMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `WcceAPI.WcceForestMappingList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWcceForestMappingListRequest struct via the builder pattern


### Return type

[**[]WcceForestMappingResponse**](WcceForestMappingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WcceForestMappingUpdate

> WcceForestMappingResponse WcceForestMappingUpdate(ctx).WcceForestMapping(wcceForestMapping).Execute()

Update an existing WCCE forest mapping



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
	wcceForestMapping := *openapiclient.NewWcceForestMapping("Forest_example", []openapiclient.WcceTemplateMapping{*openapiclient.NewWcceTemplateMapping("EnrollmentMode_example", "Profile_example", "Template_example")}) // WcceForestMapping | The WCCE forest mapping to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WcceAPI.WcceForestMappingUpdate(context.Background()).WcceForestMapping(wcceForestMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WcceAPI.WcceForestMappingUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WcceForestMappingUpdate`: WcceForestMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `WcceAPI.WcceForestMappingUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWcceForestMappingUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wcceForestMapping** | [**WcceForestMapping**](WcceForestMapping.md) | The WCCE forest mapping to update | 

### Return type

[**WcceForestMappingResponse**](WcceForestMappingResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

