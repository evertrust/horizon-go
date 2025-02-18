# \CaAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaAdd**](CaAPI.md#CaAdd) | **Post** /api/v1/cas | Register a new certificate authority
[**CaDelete**](CaAPI.md#CaDelete) | **Delete** /api/v1/cas/{name} | Delete an existing certificate authority
[**CaGet**](CaAPI.md#CaGet) | **Get** /api/v1/cas/{name} | Retrieve an existing certificate authority
[**CaList**](CaAPI.md#CaList) | **Get** /api/v1/cas | List the existing certificate authorities
[**CaUpdate**](CaAPI.md#CaUpdate) | **Put** /api/v1/cas | Update an existing certificate authority



## CaAdd

> CertificateAuthorityResponseResponse CaAdd(ctx).CertificateAuthorityRequest(certificateAuthorityRequest).Execute()

Register a new certificate authority



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
	certificateAuthorityRequest := *openapiclient.NewCertificateAuthorityRequest("Name_example", false, false, "OutdatedRevocationStatusPolicy_example", false) // CertificateAuthorityRequest | Certificate authority to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaAPI.CaAdd(context.Background()).CertificateAuthorityRequest(certificateAuthorityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaAPI.CaAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaAdd`: CertificateAuthorityResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CaAPI.CaAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateAuthorityRequest** | [**CertificateAuthorityRequest**](CertificateAuthorityRequest.md) | Certificate authority to register | 

### Return type

[**CertificateAuthorityResponseResponse**](CertificateAuthorityResponseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaDelete

> CaDelete(ctx, name).Execute()

Delete an existing certificate authority



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CaAPI.CaDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaAPI.CaDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCaDeleteRequest struct via the builder pattern


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


## CaGet

> CertificateAuthorityResponseResponse CaGet(ctx, name).Execute()

Retrieve an existing certificate authority



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaAPI.CaGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaAPI.CaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaGet`: CertificateAuthorityResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CaAPI.CaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateAuthorityResponseResponse**](CertificateAuthorityResponseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaList

> []CertificateAuthorityResponseResponse CaList(ctx).Execute()

List the existing certificate authorities



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
	resp, r, err := apiClient.CaAPI.CaList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaAPI.CaList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaList`: []CertificateAuthorityResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CaAPI.CaList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCaListRequest struct via the builder pattern


### Return type

[**[]CertificateAuthorityResponseResponse**](CertificateAuthorityResponseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaUpdate

> CertificateAuthorityResponseResponse CaUpdate(ctx).CertificateAuthorityRequest(certificateAuthorityRequest).Execute()

Update an existing certificate authority



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
	certificateAuthorityRequest := *openapiclient.NewCertificateAuthorityRequest("Name_example", false, false, "OutdatedRevocationStatusPolicy_example", false) // CertificateAuthorityRequest | Certificate authority to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaAPI.CaUpdate(context.Background()).CertificateAuthorityRequest(certificateAuthorityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaAPI.CaUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaUpdate`: CertificateAuthorityResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CaAPI.CaUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateAuthorityRequest** | [**CertificateAuthorityRequest**](CertificateAuthorityRequest.md) | Certificate authority to update | 

### Return type

[**CertificateAuthorityResponseResponse**](CertificateAuthorityResponseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

