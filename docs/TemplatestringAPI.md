# \TemplatestringAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputationruleEval**](TemplatestringAPI.md#ComputationruleEval) | **Post** /api/v1/templatestring/playground | Evaluate a computation rule and its dictionary



## ComputationruleEval

> TemplateStringPlaygroundResponseResponse ComputationruleEval(ctx).TemplateStringPlaygroundRequest(templateStringPlaygroundRequest).Execute()

Evaluate a computation rule and its dictionary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go/horizon"
)

func main() {
	templateStringPlaygroundRequest := *openapiclient.NewTemplateStringPlaygroundRequest() // TemplateStringPlaygroundRequest | Playground request to evaluate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatestringAPI.ComputationruleEval(context.Background()).TemplateStringPlaygroundRequest(templateStringPlaygroundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatestringAPI.ComputationruleEval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputationruleEval`: TemplateStringPlaygroundResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `TemplatestringAPI.ComputationruleEval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputationruleEvalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateStringPlaygroundRequest** | [**TemplateStringPlaygroundRequest**](TemplateStringPlaygroundRequest.md) | Playground request to evaluate | 

### Return type

[**TemplateStringPlaygroundResponseResponse**](TemplateStringPlaygroundResponseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

