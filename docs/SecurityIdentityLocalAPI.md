# \SecurityIdentityLocalAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityIdentityLocalAdd**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalAdd) | **Post** /api/v1/security/identity/locals | Create a local identity
[**SecurityIdentityLocalDelete**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalDelete) | **Delete** /api/v1/security/identity/locals/{identifier} | Delete a local identity
[**SecurityIdentityLocalGet**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalGet) | **Get** /api/v1/security/identity/locals/{identifier} | Retrieve a local identity
[**SecurityIdentityLocalList**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalList) | **Get** /api/v1/security/identity/locals | List local identities
[**SecurityIdentityLocalPasswordReset**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalPasswordReset) | **Post** /api/v1/security/identity/locals/password | Reset a password
[**SecurityIdentityLocalPasswordResetRequest**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalPasswordResetRequest) | **Get** /api/v1/security/identity/locals/password/{identifier} | Request a password reset
[**SecurityIdentityLocalPasswordSet**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalPasswordSet) | **Patch** /api/v1/security/identity/locals | Set the password for a local identity
[**SecurityIdentityLocalUpdate**](SecurityIdentityLocalAPI.md#SecurityIdentityLocalUpdate) | **Put** /api/v1/security/identity/locals | Update a local identity



## SecurityIdentityLocalAdd

> LocalIdentityResponse SecurityIdentityLocalAdd(ctx).LocalIdentityOnAdd(localIdentityOnAdd).Execute()

Create a local identity



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
	localIdentityOnAdd := *openapiclient.NewLocalIdentityOnAdd("administrator", "534169469812674870598506170552236971648310761036167896036064400452449656") // LocalIdentityOnAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalAdd(context.Background()).LocalIdentityOnAdd(localIdentityOnAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityLocalAdd`: LocalIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityLocalAPI.SecurityIdentityLocalAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localIdentityOnAdd** | [**LocalIdentityOnAdd**](LocalIdentityOnAdd.md) |  | 

### Return type

[**LocalIdentityResponse**](LocalIdentityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityLocalDelete

> SecurityIdentityLocalDelete(ctx, identifier).Execute()

Delete a local identity



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
	identifier := "administrator" // string | Local identity identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalDelete(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Local identity identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalDeleteRequest struct via the builder pattern


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


## SecurityIdentityLocalGet

> LocalIdentityResponse SecurityIdentityLocalGet(ctx, identifier).Execute()

Retrieve a local identity



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
	identifier := "administrator" // string | Local identity identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalGet(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityLocalGet`: LocalIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityLocalAPI.SecurityIdentityLocalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Local identity identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalIdentityResponse**](LocalIdentityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityLocalList

> []LocalIdentityResponse SecurityIdentityLocalList(ctx).Execute()

List local identities



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
	resp, r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityLocalList`: []LocalIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityLocalAPI.SecurityIdentityLocalList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalListRequest struct via the builder pattern


### Return type

[**[]LocalIdentityResponse**](LocalIdentityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityLocalPasswordReset

> SecurityIdentityLocalPasswordReset(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset a password



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
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("administrator", "Uuid_example", "Password_example") // ResetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordReset(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityLocalPasswordResetRequest

> SecurityIdentityLocalPasswordResetRequest(ctx, identifier).Execute()

Request a password reset



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
	identifier := "administrator" // string | Local identity identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordResetRequest(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordResetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** | Local identity identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalPasswordResetRequestRequest struct via the builder pattern


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


## SecurityIdentityLocalPasswordSet

> SecurityIdentityLocalPasswordSet(ctx).SetPasswordRequest(setPasswordRequest).Execute()

Set the password for a local identity



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
	setPasswordRequest := *openapiclient.NewSetPasswordRequest("Sup3rSecurePassw0rd") // SetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordSet(context.Background()).SetPasswordRequest(setPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalPasswordSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalPasswordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPasswordRequest** | [**SetPasswordRequest**](SetPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityIdentityLocalUpdate

> LocalIdentityResponse SecurityIdentityLocalUpdate(ctx).LocalIdentity(localIdentity).Execute()

Update a local identity



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
	localIdentity := *openapiclient.NewLocalIdentity("administrator") // LocalIdentity | Local identity to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIdentityLocalAPI.SecurityIdentityLocalUpdate(context.Background()).LocalIdentity(localIdentity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIdentityLocalAPI.SecurityIdentityLocalUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityIdentityLocalUpdate`: LocalIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityIdentityLocalAPI.SecurityIdentityLocalUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityIdentityLocalUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localIdentity** | [**LocalIdentity**](LocalIdentity.md) | Local identity to update | 

### Return type

[**LocalIdentityResponse**](LocalIdentityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

