# \SecurityScimprofileAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityScimprofileAdd**](SecurityScimprofileAPI.md#SecurityScimprofileAdd) | **Post** /api/v1/security/scim/profiles | Create a Scim profile
[**SecurityScimprofileDelete**](SecurityScimprofileAPI.md#SecurityScimprofileDelete) | **Delete** /api/v1/security/scim/profiles/{name} | Delete a Scim profile
[**SecurityScimprofileGet**](SecurityScimprofileAPI.md#SecurityScimprofileGet) | **Get** /api/v1/security/scim/profiles/{name} | Retrieve a Scim profile
[**SecurityScimprofileList**](SecurityScimprofileAPI.md#SecurityScimprofileList) | **Get** /api/v1/security/scim/profiles | List Scim profiles
[**SecurityScimprofileUpdate**](SecurityScimprofileAPI.md#SecurityScimprofileUpdate) | **Put** /api/v1/security/scim/profiles | Update Scim profile



## SecurityScimprofileAdd

> ScimProfileResponse SecurityScimprofileAdd(ctx).ScimProfile(scimProfile).Execute()

Create a Scim profile



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
	scimProfile := *openapiclient.NewScimProfile("OktaScim") // ScimProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityScimprofileAPI.SecurityScimprofileAdd(context.Background()).ScimProfile(scimProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScimprofileAPI.SecurityScimprofileAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityScimprofileAdd`: ScimProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityScimprofileAPI.SecurityScimprofileAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityScimprofileAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimProfile** | [**ScimProfile**](ScimProfile.md) |  | 

### Return type

[**ScimProfileResponse**](ScimProfileResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityScimprofileDelete

> SecurityScimprofileDelete(ctx, name).Execute()

Delete a Scim profile



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
	name := "OktaScim" // string | Scim profile name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityScimprofileAPI.SecurityScimprofileDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScimprofileAPI.SecurityScimprofileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Scim profile name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityScimprofileDeleteRequest struct via the builder pattern


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


## SecurityScimprofileGet

> ScimProfileResponse SecurityScimprofileGet(ctx, name).Execute()

Retrieve a Scim profile



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
	name := "OktaScim" // string | Scim profile name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityScimprofileAPI.SecurityScimprofileGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScimprofileAPI.SecurityScimprofileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityScimprofileGet`: ScimProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityScimprofileAPI.SecurityScimprofileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Scim profile name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityScimprofileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimProfileResponse**](ScimProfileResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityScimprofileList

> []ScimProfileResponse SecurityScimprofileList(ctx).Execute()

List Scim profiles



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
	resp, r, err := apiClient.SecurityScimprofileAPI.SecurityScimprofileList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScimprofileAPI.SecurityScimprofileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityScimprofileList`: []ScimProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityScimprofileAPI.SecurityScimprofileList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityScimprofileListRequest struct via the builder pattern


### Return type

[**[]ScimProfileResponse**](ScimProfileResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityScimprofileUpdate

> ScimProfileResponse SecurityScimprofileUpdate(ctx).ScimProfile(scimProfile).Execute()

Update Scim profile



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
	scimProfile := *openapiclient.NewScimProfile("OktaScim") // ScimProfile | The Scim profile to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityScimprofileAPI.SecurityScimprofileUpdate(context.Background()).ScimProfile(scimProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityScimprofileAPI.SecurityScimprofileUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityScimprofileUpdate`: ScimProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityScimprofileAPI.SecurityScimprofileUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityScimprofileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimProfile** | [**ScimProfile**](ScimProfile.md) | The Scim profile to update | 

### Return type

[**ScimProfileResponse**](ScimProfileResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

