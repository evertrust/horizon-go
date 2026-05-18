# \CertificateGradingPolicyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GradingPolicyExplainFile**](CertificateGradingPolicyAPI.md#GradingPolicyExplainFile) | **Post** /api/v1/certificate/grading/policies/{policy}/explain | Explain a Grading Policy
[**GradingPolicyExplainUrl**](CertificateGradingPolicyAPI.md#GradingPolicyExplainUrl) | **Get** /api/v1/certificate/grading/policies/{policy}/explain/{input} | Explain a Grading Policy
[**GradingPolicyGet**](CertificateGradingPolicyAPI.md#GradingPolicyGet) | **Get** /api/v1/certificate/grading/policies/{name} | Retrieve an existing grading policy
[**GradingPolicyList**](CertificateGradingPolicyAPI.md#GradingPolicyList) | **Get** /api/v1/certificate/grading/policies | List the existing grading policies
[**GradingPolicyRun**](CertificateGradingPolicyAPI.md#GradingPolicyRun) | **Get** /api/v1/certificate/grading/policies/{policy}/run | Run a grading policy



## GradingPolicyExplainFile

> ExplainedGradingPolicyResponse GradingPolicyExplainFile(ctx, policy).X509(x509).Execute()

Explain a Grading Policy



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
	policy := "policy_example" // string | 
	x509 := []byte("example") // []byte | The x509 certificate, PEM or DER encoded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateGradingPolicyAPI.GradingPolicyExplainFile(context.Background(), policy).X509(x509).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingPolicyAPI.GradingPolicyExplainFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingPolicyExplainFile`: ExplainedGradingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingPolicyAPI.GradingPolicyExplainFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingPolicyExplainFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **x509** | **[]byte** | The x509 certificate, PEM or DER encoded | 

### Return type

[**ExplainedGradingPolicyResponse**](ExplainedGradingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingPolicyExplainUrl

> ExplainedGradingPolicyResponse GradingPolicyExplainUrl(ctx, policy, input).Execute()

Explain a Grading Policy



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
	policy := "policy_example" // string | 
	input := "input_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateGradingPolicyAPI.GradingPolicyExplainUrl(context.Background(), policy, input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingPolicyAPI.GradingPolicyExplainUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingPolicyExplainUrl`: ExplainedGradingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingPolicyAPI.GradingPolicyExplainUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** |  | 
**input** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingPolicyExplainUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExplainedGradingPolicyResponse**](ExplainedGradingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingPolicyGet

> GradingPolicyResponse GradingPolicyGet(ctx, name).Execute()

Retrieve an existing grading policy



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
	resp, r, err := apiClient.CertificateGradingPolicyAPI.GradingPolicyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingPolicyAPI.GradingPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingPolicyGet`: GradingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingPolicyAPI.GradingPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GradingPolicyResponse**](GradingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingPolicyList

> []GradingPolicyResponse GradingPolicyList(ctx).Execute()

List the existing grading policies



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
	resp, r, err := apiClient.CertificateGradingPolicyAPI.GradingPolicyList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingPolicyAPI.GradingPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingPolicyList`: []GradingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingPolicyAPI.GradingPolicyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGradingPolicyListRequest struct via the builder pattern


### Return type

[**[]GradingPolicyResponse**](GradingPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingPolicyRun

> GradingPolicyRun(ctx, policy).Execute()

Run a grading policy



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
	policy := "policy_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateGradingPolicyAPI.GradingPolicyRun(context.Background(), policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingPolicyAPI.GradingPolicyRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingPolicyRunRequest struct via the builder pattern


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

