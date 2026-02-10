# \SecurityIdentityProviderAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityIdentityProviderAdd**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderAdd) | **Post** /api/v1/security/identity/providers | Create a new identity provider
[**SecurityIdentityProviderDelete**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderDelete) | **Delete** /api/v1/security/identity/providers/{name} | Delete an identity provider
[**SecurityIdentityProviderEnabled**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderEnabled) | **Get** /api/v1/security/identity/providers/dynamic/enabled | List the enabled identity provider(s)
[**SecurityIdentityProviderGet**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderGet) | **Get** /api/v1/security/identity/providers/{name} | Retrieve a existing identity provider
[**SecurityIdentityProviderList**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderList) | **Get** /api/v1/security/identity/providers | List all the identity provider(s)
[**SecurityIdentityProviderSearch**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderSearch) | **Post** /api/v1/security/identity/providers/search | Retrieve the provider of a principal
[**SecurityIdentityProviderUpdate**](SecurityIdentityProviderAPI.md#SecurityIdentityProviderUpdate) | **Put** /api/v1/security/identity/providers | Update an existing identity provider



## SecurityIdentityProviderAdd

> SecurityIdentityProviderAdd201Response SecurityIdentityProviderAdd(ctx).SecurityIdentityProviderAddRequest(securityIdentityProviderAddRequest).Execute()

Create a new identity provider



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
	securityIdentityProviderAddRequest := openapiclient.security_identity_provider_add_request{LocalIdentityProvider: openapiclient.NewLocalIdentityProvider("local", "Local", true, true)} // SecurityIdentityProviderAddRequest | The identity provider to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderAdd(context.Background()).SecurityIdentityProviderAddRequest(securityIdentityProviderAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderAdd`: SecurityIdentityProviderAdd201Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityIdentityProviderAddRequest** | [**SecurityIdentityProviderAddRequest**](SecurityIdentityProviderAddRequest.md) | The identity provider to register | 

### Return type

[**SecurityIdentityProviderAdd201Response**](SecurityIdentityProviderAdd201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityProviderDelete

> SecurityIdentityProviderDelete(ctx, name).Execute()

Delete an identity provider



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
	name := "name_example" // string | The internal name of the identity provider to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the identity provider to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderDeleteRequest struct via the builder pattern


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


## SecurityIdentityProviderEnabled

> []EnabledIdentityProviderResponse SecurityIdentityProviderEnabled(ctx).EnabledOnUI(enabledOnUI).Execute()

List the enabled identity provider(s)



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
	enabledOnUI := true // bool | Whether the identity provider is visible on the UI (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderEnabled(context.Background()).EnabledOnUI(enabledOnUI).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderEnabled`: []EnabledIdentityProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderEnabled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabledOnUI** | **bool** | Whether the identity provider is visible on the UI | 

### Return type

[**[]EnabledIdentityProviderResponse**](EnabledIdentityProviderResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityProviderGet

> SecurityIdentityProviderList200ResponseInner SecurityIdentityProviderGet(ctx, name).Execute()

Retrieve a existing identity provider



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
	name := "name_example" // string | The internal name of the identity provider to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderGet`: SecurityIdentityProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the identity provider to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityIdentityProviderList200ResponseInner**](SecurityIdentityProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityProviderList

> []SecurityIdentityProviderList200ResponseInner SecurityIdentityProviderList(ctx).Execute()

List all the identity provider(s)



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
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderList`: []SecurityIdentityProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderListRequest struct via the builder pattern


### Return type

[**[]SecurityIdentityProviderList200ResponseInner**](SecurityIdentityProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityProviderSearch

> []PrincipalInfoSearchResultResponse SecurityIdentityProviderSearch(ctx).PrincipalInfoSearchRequest(principalInfoSearchRequest).Execute()

Retrieve the provider of a principal



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
	principalInfoSearchRequest := *openapiclient.NewPrincipalInfoSearchRequest() // PrincipalInfoSearchRequest | The principal search request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderSearch(context.Background()).PrincipalInfoSearchRequest(principalInfoSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderSearch`: []PrincipalInfoSearchResultResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalInfoSearchRequest** | [**PrincipalInfoSearchRequest**](PrincipalInfoSearchRequest.md) | The principal search request | 

### Return type

[**[]PrincipalInfoSearchResultResponse**](PrincipalInfoSearchResultResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityProviderUpdate

> SecurityIdentityProviderList200ResponseInner SecurityIdentityProviderUpdate(ctx).SecurityIdentityProviderUpdateRequest(securityIdentityProviderUpdateRequest).Execute()

Update an existing identity provider



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
	securityIdentityProviderUpdateRequest := openapiclient.security_identity_provider_update_request{LocalIdentityProvider: openapiclient.NewLocalIdentityProvider("local", "Local", true, true)} // SecurityIdentityProviderUpdateRequest | Identity provider to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityProviderAPI.SecurityIdentityProviderUpdate(context.Background()).SecurityIdentityProviderUpdateRequest(securityIdentityProviderUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityProviderAPI.SecurityIdentityProviderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityProviderUpdate`: SecurityIdentityProviderList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityProviderAPI.SecurityIdentityProviderUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityProviderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityIdentityProviderUpdateRequest** | [**SecurityIdentityProviderUpdateRequest**](SecurityIdentityProviderUpdateRequest.md) | Identity provider to update | 

### Return type

[**SecurityIdentityProviderList200ResponseInner**](SecurityIdentityProviderList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

