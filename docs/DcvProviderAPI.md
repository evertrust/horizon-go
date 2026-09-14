# \DcvProviderAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DcvProviderAdd**](DcvProviderAPI.md#DcvProviderAdd) | **Post** /api/v1/dcv/providers | Register a new DCV provider configuration
[**DcvProviderDelete**](DcvProviderAPI.md#DcvProviderDelete) | **Delete** /api/v1/dcv/providers/{name} | Delete an existing DCV provider configuration
[**DcvProviderGet**](DcvProviderAPI.md#DcvProviderGet) | **Get** /api/v1/dcv/providers/{name} | Retrieve an existing DCV provider configuration
[**DcvProviderList**](DcvProviderAPI.md#DcvProviderList) | **Get** /api/v1/dcv/providers | List the existing DCV provider configurations
[**DcvProviderUpdate**](DcvProviderAPI.md#DcvProviderUpdate) | **Put** /api/v1/dcv/providers | Update an existing DCV provider configuration



## DcvProviderAdd

> DcvProviderList200ResponseInner DcvProviderAdd(ctx).DcvProviderUpdateRequest(dcvProviderUpdateRequest).Execute()

Register a new DCV provider configuration



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
	dcvProviderUpdateRequest := openapiclient.dcv_provider_update_request{DigicertDCVProviderConfig: openapiclient.NewDigicertDCVProviderConfig("Credentials_example", "https://www.digicert.com", "Name_example", "Type_example")} // DcvProviderUpdateRequest | DCV provider configuration to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvProviderAPI.DcvProviderAdd(context.Background()).DcvProviderUpdateRequest(dcvProviderUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProviderAPI.DcvProviderAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProviderAdd`: DcvProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProviderAPI.DcvProviderAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvProviderAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dcvProviderUpdateRequest** | [**DcvProviderUpdateRequest**](DcvProviderUpdateRequest.md) | DCV provider configuration to register | 

### Return type

[**DcvProviderList200ResponseInner**](DcvProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProviderDelete

> DcvProviderDelete(ctx, name).Execute()

Delete an existing DCV provider configuration



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
	r, err := apiClient.DcvProviderAPI.DcvProviderDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProviderAPI.DcvProviderDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDcvProviderDeleteRequest struct via the builder pattern


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


## DcvProviderGet

> DcvProviderList200ResponseInner DcvProviderGet(ctx, name).Execute()

Retrieve an existing DCV provider configuration



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
	resp, r, err := apiClient.DcvProviderAPI.DcvProviderGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProviderAPI.DcvProviderGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProviderGet`: DcvProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProviderAPI.DcvProviderGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvProviderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DcvProviderList200ResponseInner**](DcvProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProviderList

> []DcvProviderList200ResponseInner DcvProviderList(ctx).Execute()

List the existing DCV provider configurations



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
	resp, r, err := apiClient.DcvProviderAPI.DcvProviderList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProviderAPI.DcvProviderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProviderList`: []DcvProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProviderAPI.DcvProviderList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDcvProviderListRequest struct via the builder pattern


### Return type

[**[]DcvProviderList200ResponseInner**](DcvProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProviderUpdate

> DcvProviderList200ResponseInner DcvProviderUpdate(ctx).DcvProviderUpdateRequest(dcvProviderUpdateRequest).Execute()

Update an existing DCV provider configuration



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
	dcvProviderUpdateRequest := openapiclient.dcv_provider_update_request{DigicertDCVProviderConfig: openapiclient.NewDigicertDCVProviderConfig("Credentials_example", "https://www.digicert.com", "Name_example", "Type_example")} // DcvProviderUpdateRequest | DCV provider configuration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvProviderAPI.DcvProviderUpdate(context.Background()).DcvProviderUpdateRequest(dcvProviderUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProviderAPI.DcvProviderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProviderUpdate`: DcvProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProviderAPI.DcvProviderUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvProviderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dcvProviderUpdateRequest** | [**DcvProviderUpdateRequest**](DcvProviderUpdateRequest.md) | DCV provider configuration to update | 

### Return type

[**DcvProviderList200ResponseInner**](DcvProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

