# \SecurityServiceAccountAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityServiceAccountsAdd**](SecurityServiceAccountAPI.md#SecurityServiceAccountsAdd) | **Post** /api/v1/security/service-accounts | Create a new service account
[**SecurityServiceAccountsDelete**](SecurityServiceAccountAPI.md#SecurityServiceAccountsDelete) | **Delete** /api/v1/security/service-accounts/{name} | Delete a service account
[**SecurityServiceAccountsGet**](SecurityServiceAccountAPI.md#SecurityServiceAccountsGet) | **Get** /api/v1/security/service-accounts/{name} | Retrieve an existing service account
[**SecurityServiceAccountsList**](SecurityServiceAccountAPI.md#SecurityServiceAccountsList) | **Get** /api/v1/security/service-accounts | List all the service accounts
[**SecurityServiceAccountsUpdate**](SecurityServiceAccountAPI.md#SecurityServiceAccountsUpdate) | **Put** /api/v1/security/service-accounts | Update an existing service account



## SecurityServiceAccountsAdd

> ServiceAccountResponse SecurityServiceAccountsAdd(ctx).ServiceAccount(serviceAccount).Execute()

Create a new service account



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
	serviceAccount := *openapiclient.NewServiceAccount("Name_example", []openapiclient.Permission{*openapiclient.NewPermission("lifecycle:*:*:enroll")}, []string{"ca-auditor"}, openapiclient.ServiceAccount_trustConfig{DynamicJWKS: openapiclient.NewDynamicJWKS()}, []string{"{{jwt.issuer}} contains "evertrust.com""}) // ServiceAccount | The service account to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityServiceAccountAPI.SecurityServiceAccountsAdd(context.Background()).ServiceAccount(serviceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityServiceAccountAPI.SecurityServiceAccountsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityServiceAccountsAdd`: ServiceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityServiceAccountAPI.SecurityServiceAccountsAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityServiceAccountsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) | The service account to register | 

### Return type

[**ServiceAccountResponse**](ServiceAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityServiceAccountsDelete

> SecurityServiceAccountsDelete(ctx, name).Execute()

Delete a service account



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
	name := "name_example" // string | The internal name of the service account to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityServiceAccountAPI.SecurityServiceAccountsDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityServiceAccountAPI.SecurityServiceAccountsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the service account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityServiceAccountsDeleteRequest struct via the builder pattern


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


## SecurityServiceAccountsGet

> ServiceAccountResponse SecurityServiceAccountsGet(ctx, name).Execute()

Retrieve an existing service account



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
	name := "name_example" // string | The internal name of the service account to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityServiceAccountAPI.SecurityServiceAccountsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityServiceAccountAPI.SecurityServiceAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityServiceAccountsGet`: ServiceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityServiceAccountAPI.SecurityServiceAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the service account to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityServiceAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountResponse**](ServiceAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityServiceAccountsList

> []ServiceAccountResponse SecurityServiceAccountsList(ctx).Execute()

List all the service accounts



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
	resp, r, err := apiClient.SecurityServiceAccountAPI.SecurityServiceAccountsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityServiceAccountAPI.SecurityServiceAccountsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityServiceAccountsList`: []ServiceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityServiceAccountAPI.SecurityServiceAccountsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityServiceAccountsListRequest struct via the builder pattern


### Return type

[**[]ServiceAccountResponse**](ServiceAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityServiceAccountsUpdate

> ServiceAccountResponse SecurityServiceAccountsUpdate(ctx).ServiceAccount(serviceAccount).Execute()

Update an existing service account



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
	serviceAccount := *openapiclient.NewServiceAccount("Name_example", []openapiclient.Permission{*openapiclient.NewPermission("lifecycle:*:*:enroll")}, []string{"ca-auditor"}, openapiclient.ServiceAccount_trustConfig{DynamicJWKS: openapiclient.NewDynamicJWKS()}, []string{"{{jwt.issuer}} contains "evertrust.com""}) // ServiceAccount | Service account to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityServiceAccountAPI.SecurityServiceAccountsUpdate(context.Background()).ServiceAccount(serviceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityServiceAccountAPI.SecurityServiceAccountsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityServiceAccountsUpdate`: ServiceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityServiceAccountAPI.SecurityServiceAccountsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityServiceAccountsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) | Service account to update | 

### Return type

[**ServiceAccountResponse**](ServiceAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

