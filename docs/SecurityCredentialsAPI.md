# \SecurityCredentialsAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCredentialsAdd**](SecurityCredentialsAPI.md#SecurityCredentialsAdd) | **Post** /api/v1/security/credentials | Create new credentials
[**SecurityCredentialsDelete**](SecurityCredentialsAPI.md#SecurityCredentialsDelete) | **Delete** /api/v1/security/credentials/{name} | Delete credentials
[**SecurityCredentialsGet**](SecurityCredentialsAPI.md#SecurityCredentialsGet) | **Get** /api/v1/security/credentials/{name} | Retrieve credentials
[**SecurityCredentialsList**](SecurityCredentialsAPI.md#SecurityCredentialsList) | **Get** /api/v1/security/credentials | List credentials
[**SecurityCredentialsUpdate**](SecurityCredentialsAPI.md#SecurityCredentialsUpdate) | **Put** /api/v1/security/credentials | Update credentials



## SecurityCredentialsAdd

> SecurityCredentialsList200ResponseInner SecurityCredentialsAdd(ctx).SecurityCredentialsUpdateRequest(securityCredentialsUpdateRequest).Execute()

Create new credentials



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
	securityCredentialsUpdateRequest := openapiclient.security_credentials_update_request{CertificateCredentials: openapiclient.NewCertificateCredentials("Type_example", *openapiclient.NewSecretStoreRequest("-----BEGIN CERTIFICATE-----...", "-----BEGIN PRIVATE KEY-----..."), "My credentials")} // SecurityCredentialsUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCredentialsAPI.SecurityCredentialsAdd(context.Background()).SecurityCredentialsUpdateRequest(securityCredentialsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCredentialsAPI.SecurityCredentialsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityCredentialsAdd`: SecurityCredentialsList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityCredentialsAPI.SecurityCredentialsAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCredentialsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityCredentialsUpdateRequest** | [**SecurityCredentialsUpdateRequest**](SecurityCredentialsUpdateRequest.md) |  | 

### Return type

[**SecurityCredentialsList200ResponseInner**](SecurityCredentialsList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCredentialsDelete

> SecurityCredentialsDelete(ctx, name).Execute()

Delete credentials



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
	name := "SuperAdmins" // string | Credentials name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityCredentialsAPI.SecurityCredentialsDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCredentialsAPI.SecurityCredentialsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Credentials name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCredentialsDeleteRequest struct via the builder pattern


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


## SecurityCredentialsGet

> SecurityCredentialsList200ResponseInner SecurityCredentialsGet(ctx, name).Execute()

Retrieve credentials



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
	name := "SuperAdmin" // string | Credentials name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCredentialsAPI.SecurityCredentialsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCredentialsAPI.SecurityCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityCredentialsGet`: SecurityCredentialsList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityCredentialsAPI.SecurityCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Credentials name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityCredentialsList200ResponseInner**](SecurityCredentialsList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCredentialsList

> []SecurityCredentialsList200ResponseInner SecurityCredentialsList(ctx).Execute()

List credentials



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
	resp, r, err := apiClient.SecurityCredentialsAPI.SecurityCredentialsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCredentialsAPI.SecurityCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityCredentialsList`: []SecurityCredentialsList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityCredentialsAPI.SecurityCredentialsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCredentialsListRequest struct via the builder pattern


### Return type

[**[]SecurityCredentialsList200ResponseInner**](SecurityCredentialsList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityCredentialsUpdate

> SecurityCredentialsList200ResponseInner SecurityCredentialsUpdate(ctx).SecurityCredentialsUpdateRequest(securityCredentialsUpdateRequest).Execute()

Update credentials



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
	securityCredentialsUpdateRequest := openapiclient.security_credentials_update_request{CertificateCredentials: openapiclient.NewCertificateCredentials("Type_example", *openapiclient.NewSecretStoreRequest("-----BEGIN CERTIFICATE-----...", "-----BEGIN PRIVATE KEY-----..."), "My credentials")} // SecurityCredentialsUpdateRequest | The credentials to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCredentialsAPI.SecurityCredentialsUpdate(context.Background()).SecurityCredentialsUpdateRequest(securityCredentialsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCredentialsAPI.SecurityCredentialsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityCredentialsUpdate`: SecurityCredentialsList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityCredentialsAPI.SecurityCredentialsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCredentialsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityCredentialsUpdateRequest** | [**SecurityCredentialsUpdateRequest**](SecurityCredentialsUpdateRequest.md) | The credentials to update | 

### Return type

[**SecurityCredentialsList200ResponseInner**](SecurityCredentialsList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

