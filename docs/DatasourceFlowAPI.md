# \DatasourceFlowAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatasourceFlowTemplate**](DatasourceFlowAPI.md#DatasourceFlowTemplate) | **Post** /api/v1/datasource/flows/template | Retrieve the template for a datasource flow
[**DatasourceFlowTest**](DatasourceFlowAPI.md#DatasourceFlowTest) | **Post** /api/v1/datasource/flows | Test a datasource flow



## DatasourceFlowTemplate

> []DatasourceFlowTemplate200ResponseInner DatasourceFlowTemplate(ctx).DatasourceFlow1(datasourceFlow1).Execute()

Retrieve the template for a datasource flow



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
	datasourceFlow1 := *openapiclient.NewDatasourceFlow1() // DatasourceFlow1 | Datasource flow to get template for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceFlowAPI.DatasourceFlowTemplate(context.Background()).DatasourceFlow1(datasourceFlow1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceFlowAPI.DatasourceFlowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceFlowTemplate`: []DatasourceFlowTemplate200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DatasourceFlowAPI.DatasourceFlowTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceFlowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceFlow1** | [**DatasourceFlow1**](DatasourceFlow1.md) | Datasource flow to get template for | 

### Return type

[**[]DatasourceFlowTemplate200ResponseInner**](DatasourceFlowTemplate200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceFlowTest

> []DatasourceTest200Response DatasourceFlowTest(ctx).DatasourceFlow(datasourceFlow).Execute()

Test a datasource flow



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
	datasourceFlow := *openapiclient.NewDatasourceFlow() // DatasourceFlow | Datasource flow to test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceFlowAPI.DatasourceFlowTest(context.Background()).DatasourceFlow(datasourceFlow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceFlowAPI.DatasourceFlowTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceFlowTest`: []DatasourceTest200Response
	fmt.Fprintf(os.Stdout, "Response from `DatasourceFlowAPI.DatasourceFlowTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceFlowTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceFlow** | [**DatasourceFlow**](DatasourceFlow.md) | Datasource flow to test | 

### Return type

[**[]DatasourceTest200Response**](DatasourceTest200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

