# \SecurityPrincipalAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityPrincipalAuthenticate**](SecurityPrincipalAPI.md#SecurityPrincipalAuthenticate) | **Get** /api/v1/security/principals/authenticate | Authenticate a principal and redirect to the specified redirect URL
[**SecurityPrincipalLogout**](SecurityPrincipalAPI.md#SecurityPrincipalLogout) | **Get** /api/v1/security/principals/logout | Log out an authenticated principal and flush any cached authorization(s)
[**SecurityPrincipalSelf**](SecurityPrincipalAPI.md#SecurityPrincipalSelf) | **Get** /api/v1/security/principals/self | Return the authenticated principal



## SecurityPrincipalAuthenticate

> SecurityPrincipalAuthenticate(ctx).Redirect(redirect).Execute()

Authenticate a principal and redirect to the specified redirect URL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	redirect := "redirect_example" // string | The URL to redirect to after successful authentication. The URL must be URL encoded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityPrincipalAPI.SecurityPrincipalAuthenticate(context.Background()).Redirect(redirect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalAPI.SecurityPrincipalAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirect** | **string** | The URL to redirect to after successful authentication. The URL must be URL encoded. | 

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


## SecurityPrincipalLogout

> SecurityPrincipalLogout(ctx).Execute()

Log out an authenticated principal and flush any cached authorization(s)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityPrincipalAPI.SecurityPrincipalLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalAPI.SecurityPrincipalLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalLogoutRequest struct via the builder pattern


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


## SecurityPrincipalSelf

> PrincipalResponse SecurityPrincipalSelf(ctx).Execute()

Return the authenticated principal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPrincipalAPI.SecurityPrincipalSelf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPrincipalAPI.SecurityPrincipalSelf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalSelf`: PrincipalResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPrincipalAPI.SecurityPrincipalSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalSelfRequest struct via the builder pattern


### Return type

[**PrincipalResponse**](PrincipalResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

