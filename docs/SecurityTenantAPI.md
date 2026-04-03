# \SecurityTenantAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityTenantsAdd**](SecurityTenantAPI.md#SecurityTenantsAdd) | **Post** /api/v1/security/tenants | Create a new tenant
[**SecurityTenantsDelete**](SecurityTenantAPI.md#SecurityTenantsDelete) | **Delete** /api/v1/security/tenants/{name} | Delete a tenant
[**SecurityTenantsGet**](SecurityTenantAPI.md#SecurityTenantsGet) | **Get** /api/v1/security/tenants/{name} | Retrieve tenant
[**SecurityTenantsLicense**](SecurityTenantAPI.md#SecurityTenantsLicense) | **Get** /api/v1/security/tenants/{name}/licenses | Retrieve tenant license information
[**SecurityTenantsList**](SecurityTenantAPI.md#SecurityTenantsList) | **Get** /api/v1/security/tenants | List tenants
[**SecurityTenantsReset**](SecurityTenantAPI.md#SecurityTenantsReset) | **Post** /api/v1/security/tenants/{name}/reset | Reset tenant administrator account
[**SecurityTenantsRestore**](SecurityTenantAPI.md#SecurityTenantsRestore) | **Post** /api/v1/security/tenants/{name}/restore | Restore a deleted tenant
[**SecurityTenantsUpdate**](SecurityTenantAPI.md#SecurityTenantsUpdate) | **Put** /api/v1/security/tenants | Update a tenant



## SecurityTenantsAdd

> TenantWithAccountResponse SecurityTenantsAdd(ctx).TenantCreationRequest(tenantCreationRequest).Execute()

Create a new tenant



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
	tenantCreationRequest := *openapiclient.NewTenantCreationRequest(int64(123), "Name_example") // TenantCreationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsAdd(context.Background()).TenantCreationRequest(tenantCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsAdd`: TenantWithAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantCreationRequest** | [**TenantCreationRequest**](TenantCreationRequest.md) |  | 

### Return type

[**TenantWithAccountResponse**](TenantWithAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTenantsDelete

> SecurityTenantsDelete(ctx, name).Execute()

Delete a tenant



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
	name := "tenant1" // string | Tenant name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTenantAPI.SecurityTenantsDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsDeleteRequest struct via the builder pattern


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


## SecurityTenantsGet

> TenantResponse SecurityTenantsGet(ctx, name).Execute()

Retrieve tenant



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
	name := "tenant1" // string | Tenant name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsGet`: TenantResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantResponse**](TenantResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTenantsLicense

> LicenseInfoResponse SecurityTenantsLicense(ctx, name).Execute()

Retrieve tenant license information



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
	name := "tenant1" // string | Tenant name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsLicense(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsLicense`: LicenseInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseInfoResponse**](LicenseInfoResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTenantsList

> []TenantResponse SecurityTenantsList(ctx).Execute()

List tenants



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
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsList`: []TenantResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsListRequest struct via the builder pattern


### Return type

[**[]TenantResponse**](TenantResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTenantsReset

> BootstrapAccount SecurityTenantsReset(ctx, name).Execute()

Reset tenant administrator account



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
	name := "tenant1" // string | Tenant name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsReset(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsReset`: BootstrapAccount
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BootstrapAccount**](BootstrapAccount.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTenantsRestore

> SecurityTenantsRestore(ctx, name).Execute()

Restore a deleted tenant



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
	name := "tenant1" // string | Tenant name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTenantAPI.SecurityTenantsRestore(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Tenant name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsRestoreRequest struct via the builder pattern


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


## SecurityTenantsUpdate

> TenantResponse SecurityTenantsUpdate(ctx).TenantUpdateRequest(tenantUpdateRequest).Execute()

Update a tenant



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
	tenantUpdateRequest := *openapiclient.NewTenantUpdateRequest(int64(123), "Name_example") // TenantUpdateRequest | The tenant to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTenantAPI.SecurityTenantsUpdate(context.Background()).TenantUpdateRequest(tenantUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTenantAPI.SecurityTenantsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTenantsUpdate`: TenantResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTenantAPI.SecurityTenantsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantUpdateRequest** | [**TenantUpdateRequest**](TenantUpdateRequest.md) | The tenant to update | 

### Return type

[**TenantResponse**](TenantResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

