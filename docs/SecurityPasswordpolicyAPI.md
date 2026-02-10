# \SecurityPasswordpolicyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PasswordPolicyAdd**](SecurityPasswordpolicyAPI.md#PasswordPolicyAdd) | **Post** /api/v1/security/passwordpolicies | Create a password policy
[**PasswordPolicyDelete**](SecurityPasswordpolicyAPI.md#PasswordPolicyDelete) | **Delete** /api/v1/security/passwordpolicies/{name} | Delete a password policy
[**PasswordPolicyGenerate**](SecurityPasswordpolicyAPI.md#PasswordPolicyGenerate) | **Get** /api/v1/security/passwordpolicies/{name}/generate | Generate a password with a password policy
[**PasswordPolicyGet**](SecurityPasswordpolicyAPI.md#PasswordPolicyGet) | **Get** /api/v1/security/passwordpolicies/{name} | Retrieve a password policy
[**PasswordPolicyList**](SecurityPasswordpolicyAPI.md#PasswordPolicyList) | **Get** /api/v1/security/passwordpolicies | List password policies
[**PasswordPolicyUpdate**](SecurityPasswordpolicyAPI.md#PasswordPolicyUpdate) | **Put** /api/v1/security/passwordpolicies | Update a password policy



## PasswordPolicyAdd

> PasswordPolicyResponse PasswordPolicyAdd(ctx).PasswordPolicy(passwordPolicy).Execute()

Create a password policy



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
	passwordPolicy := *openapiclient.NewPasswordPolicy("Horizon-Default", int64(8)) // PasswordPolicy | The password policy to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyAdd(context.Background()).PasswordPolicy(passwordPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordPolicyAdd`: PasswordPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPasswordpolicyAPI.PasswordPolicyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) | The password policy to register | 

### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordPolicyDelete

> PasswordPolicyDelete(ctx, name).Execute()

Delete a password policy



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
	name := "name_example" // string | The internal name of the password policy to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the password policy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyDeleteRequest struct via the builder pattern


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


## PasswordPolicyGenerate

> string PasswordPolicyGenerate(ctx, name).Execute()

Generate a password with a password policy



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
	name := "name_example" // string | The internal name of the password policy that the generated password must comply with

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyGenerate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordPolicyGenerate`: string
	fmt.Fprintf(os.Stdout, "Response from `SecurityPasswordpolicyAPI.PasswordPolicyGenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the password policy that the generated password must comply with | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordPolicyGet

> PasswordPolicyResponse PasswordPolicyGet(ctx, name).Execute()

Retrieve a password policy



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
	name := "name_example" // string | The internal name of the password policy to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordPolicyGet`: PasswordPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPasswordpolicyAPI.PasswordPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The internal name of the password policy to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordPolicyList

> []PasswordPolicyResponse PasswordPolicyList(ctx).Execute()

List password policies



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
	resp, r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordPolicyList`: []PasswordPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPasswordpolicyAPI.PasswordPolicyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyListRequest struct via the builder pattern


### Return type

[**[]PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordPolicyUpdate

> PasswordPolicyResponse PasswordPolicyUpdate(ctx).PasswordPolicy(passwordPolicy).Execute()

Update a password policy



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
	passwordPolicy := *openapiclient.NewPasswordPolicy("Horizon-Default", int64(8)) // PasswordPolicy | The password policy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPasswordpolicyAPI.PasswordPolicyUpdate(context.Background()).PasswordPolicy(passwordPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPasswordpolicyAPI.PasswordPolicyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordPolicyUpdate`: PasswordPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityPasswordpolicyAPI.PasswordPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) | The password policy to update | 

### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

