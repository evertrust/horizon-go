# \AcmeEabpolicyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcmeEabpolicyAdd**](AcmeEabpolicyAPI.md#AcmeEabpolicyAdd) | **Post** /api/v1/acme/eab-policies | Add a new External Account Binding Policy
[**AcmeEabpolicyDelete**](AcmeEabpolicyAPI.md#AcmeEabpolicyDelete) | **Delete** /api/v1/acme/eab-policies/{name} | Delete an External Account Binding Policy by name
[**AcmeEabpolicyGet**](AcmeEabpolicyAPI.md#AcmeEabpolicyGet) | **Get** /api/v1/acme/eab-policies/{name} | Get an External Account Binding Policy by name
[**AcmeEabpolicyList**](AcmeEabpolicyAPI.md#AcmeEabpolicyList) | **Post** /api/v1/acme/eab-policies/list | List all External Account Binding Policies
[**AcmeEabpolicyUpdate**](AcmeEabpolicyAPI.md#AcmeEabpolicyUpdate) | **Put** /api/v1/acme/eab-policies | Update an existing External Account Binding Policy



## AcmeEabpolicyAdd

> ExternalAccountBindingPolicyResponse AcmeEabpolicyAdd(ctx).EabPolicyAddRequest(eabPolicyAddRequest).Execute()

Add a new External Account Binding Policy



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
	eabPolicyAddRequest := *openapiclient.NewEabPolicyAddRequest("Name_example") // EabPolicyAddRequest | The External Account Binding Policy to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabpolicyAPI.AcmeEabpolicyAdd(context.Background()).EabPolicyAddRequest(eabPolicyAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabpolicyAPI.AcmeEabpolicyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabpolicyAdd`: ExternalAccountBindingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabpolicyAPI.AcmeEabpolicyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabpolicyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eabPolicyAddRequest** | [**EabPolicyAddRequest**](EabPolicyAddRequest.md) | The External Account Binding Policy to add | 

### Return type

[**ExternalAccountBindingPolicyResponse**](ExternalAccountBindingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabpolicyDelete

> AcmeEabpolicyDelete(ctx, name).Execute()

Delete an External Account Binding Policy by name



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
	name := "name_example" // string | The name of the External Account Binding Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AcmeEabpolicyAPI.AcmeEabpolicyDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabpolicyAPI.AcmeEabpolicyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the External Account Binding Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabpolicyDeleteRequest struct via the builder pattern


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


## AcmeEabpolicyGet

> ExternalAccountBindingPolicyResponse AcmeEabpolicyGet(ctx, name).Execute()

Get an External Account Binding Policy by name



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
	name := "name_example" // string | The name of the External Account Binding Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabpolicyAPI.AcmeEabpolicyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabpolicyAPI.AcmeEabpolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabpolicyGet`: ExternalAccountBindingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabpolicyAPI.AcmeEabpolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the External Account Binding Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabpolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAccountBindingPolicyResponse**](ExternalAccountBindingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabpolicyList

> []ExternalAccountBindingPolicyResponse AcmeEabpolicyList(ctx).Body(body).Execute()

List all External Account Binding Policies



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Empty request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabpolicyAPI.AcmeEabpolicyList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabpolicyAPI.AcmeEabpolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabpolicyList`: []ExternalAccountBindingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabpolicyAPI.AcmeEabpolicyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabpolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Empty request body | 

### Return type

[**[]ExternalAccountBindingPolicyResponse**](ExternalAccountBindingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeEabpolicyUpdate

> ExternalAccountBindingPolicyResponse AcmeEabpolicyUpdate(ctx).EabPolicyUpdateRequest(eabPolicyUpdateRequest).Execute()

Update an existing External Account Binding Policy



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
	eabPolicyUpdateRequest := *openapiclient.NewEabPolicyUpdateRequest("Name_example") // EabPolicyUpdateRequest | External Account Binding Policy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeEabpolicyAPI.AcmeEabpolicyUpdate(context.Background()).EabPolicyUpdateRequest(eabPolicyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeEabpolicyAPI.AcmeEabpolicyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeEabpolicyUpdate`: ExternalAccountBindingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AcmeEabpolicyAPI.AcmeEabpolicyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeEabpolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eabPolicyUpdateRequest** | [**EabPolicyUpdateRequest**](EabPolicyUpdateRequest.md) | External Account Binding Policy to update | 

### Return type

[**ExternalAccountBindingPolicyResponse**](ExternalAccountBindingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

