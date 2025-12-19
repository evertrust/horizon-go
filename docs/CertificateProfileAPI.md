# \CertificateProfileAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateProfileAdd**](CertificateProfileAPI.md#CertificateProfileAdd) | **Post** /api/v1/certificate/profiles | Register a new certificate profile
[**CertificateProfileDelete**](CertificateProfileAPI.md#CertificateProfileDelete) | **Delete** /api/v1/certificate/profiles/{name} | Delete a certificate profile
[**CertificateProfileGet**](CertificateProfileAPI.md#CertificateProfileGet) | **Get** /api/v1/certificate/profiles/{name} | Retrieve a specific certificate profile
[**CertificateProfileList**](CertificateProfileAPI.md#CertificateProfileList) | **Get** /api/v1/certificate/profiles | List the existing certificate profiles
[**CertificateProfileUpdate**](CertificateProfileAPI.md#CertificateProfileUpdate) | **Put** /api/v1/certificate/profiles | Update an existing certificate profile



## CertificateProfileAdd

> CertificateProfileResponses CertificateProfileAdd(ctx).CertificateProfiles(certificateProfiles).Execute()

Register a new certificate profile



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
	certificateProfiles := openapiclient.CertificateProfiles{AcmeExternalProfile: openapiclient.NewAcmeExternalProfile(*openapiclient.NewCertificateProfileAuthorizationLevels(*openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated")), []string{"AuthorizationMethods_example"}, []string{"AuthorizedCas_example"}, *openapiclient.NewManagedCertificateProfileCryptoPolicy(), false, "Module_example", "Name_example", "PkiConnector_example", *openapiclient.NewRequestsPolicy(), false, *openapiclient.NewCertificateProfileSelfPermissions())} // CertificateProfiles | Certificate profile to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfileAPI.CertificateProfileAdd(context.Background()).CertificateProfiles(certificateProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfileAPI.CertificateProfileAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfileAdd`: CertificateProfileResponses
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfileAPI.CertificateProfileAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfileAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfiles** | [**CertificateProfiles**](CertificateProfiles.md) | Certificate profile to register | 

### Return type

[**CertificateProfileResponses**](CertificateProfileResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfileDelete

> CertificateProfileDelete(ctx, name).Execute()

Delete a certificate profile



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateProfileAPI.CertificateProfileDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfileAPI.CertificateProfileDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificateProfileDeleteRequest struct via the builder pattern


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


## CertificateProfileGet

> CertificateProfileResponses CertificateProfileGet(ctx, name).Execute()

Retrieve a specific certificate profile



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfileAPI.CertificateProfileGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfileAPI.CertificateProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfileGet`: CertificateProfileResponses
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfileAPI.CertificateProfileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateProfileResponses**](CertificateProfileResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfileList

> []CertificateProfileResponses CertificateProfileList(ctx).Modules(modules).Execute()

List the existing certificate profiles



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
	modules := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfileAPI.CertificateProfileList(context.Background()).Modules(modules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfileAPI.CertificateProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfileList`: []CertificateProfileResponses
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfileAPI.CertificateProfileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modules** | **[]string** |  | 

### Return type

[**[]CertificateProfileResponses**](CertificateProfileResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateProfileUpdate

> CertificateProfileResponses CertificateProfileUpdate(ctx).CertificateProfiles(certificateProfiles).Execute()

Update an existing certificate profile



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
	certificateProfiles := openapiclient.CertificateProfiles{AcmeExternalProfile: openapiclient.NewAcmeExternalProfile(*openapiclient.NewCertificateProfileAuthorizationLevels(*openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated")), []string{"AuthorizationMethods_example"}, []string{"AuthorizedCas_example"}, *openapiclient.NewManagedCertificateProfileCryptoPolicy(), false, "Module_example", "Name_example", "PkiConnector_example", *openapiclient.NewRequestsPolicy(), false, *openapiclient.NewCertificateProfileSelfPermissions())} // CertificateProfiles | Certificate profile to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateProfileAPI.CertificateProfileUpdate(context.Background()).CertificateProfiles(certificateProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateProfileAPI.CertificateProfileUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateProfileUpdate`: CertificateProfileResponses
	fmt.Fprintf(os.Stdout, "Response from `CertificateProfileAPI.CertificateProfileUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateProfileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateProfiles** | [**CertificateProfiles**](CertificateProfiles.md) | Certificate profile to update | 

### Return type

[**CertificateProfileResponses**](CertificateProfileResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

