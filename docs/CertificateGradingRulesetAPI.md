# \CertificateGradingRulesetAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GradingRulesetExplainFile**](CertificateGradingRulesetAPI.md#GradingRulesetExplainFile) | **Post** /api/v1/certificate/grading/rulesets/{ruleset}/explain | Explain a Grading Ruleset
[**GradingRulesetExplainUrl**](CertificateGradingRulesetAPI.md#GradingRulesetExplainUrl) | **Get** /api/v1/certificate/grading/rulesets/{ruleset}/explain/{input} | Explain a Grading Ruleset
[**GradingRulesetGet**](CertificateGradingRulesetAPI.md#GradingRulesetGet) | **Get** /api/v1/certificate/grading/rulesets/{name} | Retrieve an existing grading ruleset
[**GradingRulesetList**](CertificateGradingRulesetAPI.md#GradingRulesetList) | **Get** /api/v1/certificate/grading/rulesets | List the existing grading rulesets



## GradingRulesetExplainFile

> ExplainedGradingRulesetResponse GradingRulesetExplainFile(ctx, ruleset).X509(x509).Execute()

Explain a Grading Ruleset



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
	ruleset := "ruleset_example" // string | 
	x509 := os.NewFile(1234, "some_file") // []byte | The x509 certificate, PEM or DER encoded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateGradingRulesetAPI.GradingRulesetExplainFile(context.Background(), ruleset).X509(x509).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingRulesetAPI.GradingRulesetExplainFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingRulesetExplainFile`: ExplainedGradingRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingRulesetAPI.GradingRulesetExplainFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleset** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingRulesetExplainFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **x509** | **[]byte** | The x509 certificate, PEM or DER encoded | 

### Return type

[**ExplainedGradingRulesetResponse**](ExplainedGradingRulesetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingRulesetExplainUrl

> ExplainedGradingRulesetResponse GradingRulesetExplainUrl(ctx, ruleset, input).Execute()

Explain a Grading Ruleset



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
	ruleset := "ruleset_example" // string | 
	input := "input_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateGradingRulesetAPI.GradingRulesetExplainUrl(context.Background(), ruleset, input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingRulesetAPI.GradingRulesetExplainUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingRulesetExplainUrl`: ExplainedGradingRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingRulesetAPI.GradingRulesetExplainUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleset** | **string** |  | 
**input** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingRulesetExplainUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExplainedGradingRulesetResponse**](ExplainedGradingRulesetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingRulesetGet

> GradingRulesetResponse GradingRulesetGet(ctx, name).Execute()

Retrieve an existing grading ruleset



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
	resp, r, err := apiClient.CertificateGradingRulesetAPI.GradingRulesetGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingRulesetAPI.GradingRulesetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingRulesetGet`: GradingRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingRulesetAPI.GradingRulesetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGradingRulesetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GradingRulesetResponse**](GradingRulesetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GradingRulesetList

> []GradingRulesetResponse GradingRulesetList(ctx).Execute()

List the existing grading rulesets



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
	resp, r, err := apiClient.CertificateGradingRulesetAPI.GradingRulesetList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateGradingRulesetAPI.GradingRulesetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GradingRulesetList`: []GradingRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateGradingRulesetAPI.GradingRulesetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGradingRulesetListRequest struct via the builder pattern


### Return type

[**[]GradingRulesetResponse**](GradingRulesetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

