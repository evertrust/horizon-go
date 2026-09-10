# \DcvProvisionerAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DcvProvisionerAdd**](DcvProvisionerAPI.md#DcvProvisionerAdd) | **Post** /api/v1/dcv/provisioners | Register a new DCV provisioner configuration
[**DcvProvisionerDelete**](DcvProvisionerAPI.md#DcvProvisionerDelete) | **Delete** /api/v1/dcv/provisioners/{name} | Delete an existing DCV provisioner configuration
[**DcvProvisionerGet**](DcvProvisionerAPI.md#DcvProvisionerGet) | **Get** /api/v1/dcv/provisioners/{name} | Retrieve an existing DCV provisioner configuration
[**DcvProvisionerList**](DcvProvisionerAPI.md#DcvProvisionerList) | **Get** /api/v1/dcv/provisioners | List the existing DCV provisioner configurations
[**DcvProvisionerUpdate**](DcvProvisionerAPI.md#DcvProvisionerUpdate) | **Put** /api/v1/dcv/provisioners | Update an existing DCV provisioner configuration



## DcvProvisionerAdd

> DcvProvisionerUpdate200Response DcvProvisionerAdd(ctx).DcvProvisionerUpdateRequest(dcvProvisionerUpdateRequest).Execute()

Register a new DCV provisioner configuration



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
	dcvProvisionerUpdateRequest := openapiclient.dcv_provisioner_update_request{AzurednsDCVProvisionerConfig: openapiclient.NewAzurednsDCVProvisionerConfig("Name_example", "ResourceGroupName_example", "SubscriptionId_example", "TenantId_example", "Ttl_example", "Type_example")} // DcvProvisionerUpdateRequest | DCV provisioner configuration to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvProvisionerAPI.DcvProvisionerAdd(context.Background()).DcvProvisionerUpdateRequest(dcvProvisionerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProvisionerAPI.DcvProvisionerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProvisionerAdd`: DcvProvisionerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DcvProvisionerAPI.DcvProvisionerAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvProvisionerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dcvProvisionerUpdateRequest** | [**DcvProvisionerUpdateRequest**](DcvProvisionerUpdateRequest.md) | DCV provisioner configuration to register | 

### Return type

[**DcvProvisionerUpdate200Response**](DcvProvisionerUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProvisionerDelete

> DcvProvisionerDelete(ctx, name).Execute()

Delete an existing DCV provisioner configuration



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
	r, err := apiClient.DcvProvisionerAPI.DcvProvisionerDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProvisionerAPI.DcvProvisionerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDcvProvisionerDeleteRequest struct via the builder pattern


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


## DcvProvisionerGet

> DcvProvisionerList200ResponseInner DcvProvisionerGet(ctx, name).Execute()

Retrieve an existing DCV provisioner configuration



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
	resp, r, err := apiClient.DcvProvisionerAPI.DcvProvisionerGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProvisionerAPI.DcvProvisionerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProvisionerGet`: DcvProvisionerList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProvisionerAPI.DcvProvisionerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvProvisionerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DcvProvisionerList200ResponseInner**](DcvProvisionerList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProvisionerList

> []DcvProvisionerList200ResponseInner DcvProvisionerList(ctx).Execute()

List the existing DCV provisioner configurations



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
	resp, r, err := apiClient.DcvProvisionerAPI.DcvProvisionerList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProvisionerAPI.DcvProvisionerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProvisionerList`: []DcvProvisionerList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DcvProvisionerAPI.DcvProvisionerList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDcvProvisionerListRequest struct via the builder pattern


### Return type

[**[]DcvProvisionerList200ResponseInner**](DcvProvisionerList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvProvisionerUpdate

> DcvProvisionerUpdate200Response DcvProvisionerUpdate(ctx).DcvProvisionerUpdateRequest(dcvProvisionerUpdateRequest).Execute()

Update an existing DCV provisioner configuration



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
	dcvProvisionerUpdateRequest := openapiclient.dcv_provisioner_update_request{AzurednsDCVProvisionerConfig: openapiclient.NewAzurednsDCVProvisionerConfig("Name_example", "ResourceGroupName_example", "SubscriptionId_example", "TenantId_example", "Ttl_example", "Type_example")} // DcvProvisionerUpdateRequest | DCV provisioner configuration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvProvisionerAPI.DcvProvisionerUpdate(context.Background()).DcvProvisionerUpdateRequest(dcvProvisionerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvProvisionerAPI.DcvProvisionerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvProvisionerUpdate`: DcvProvisionerUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `DcvProvisionerAPI.DcvProvisionerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvProvisionerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dcvProvisionerUpdateRequest** | [**DcvProvisionerUpdateRequest**](DcvProvisionerUpdateRequest.md) | DCV provisioner configuration to update | 

### Return type

[**DcvProvisionerUpdate200Response**](DcvProvisionerUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

