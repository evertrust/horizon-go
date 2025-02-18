# \DatasourceAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatasourceAdd**](DatasourceAPI.md#DatasourceAdd) | **Post** /api/v1/datasources | Register a new datasource
[**DatasourceDelete**](DatasourceAPI.md#DatasourceDelete) | **Delete** /api/v1/datasources/{name} | Delete a datasource
[**DatasourceGet**](DatasourceAPI.md#DatasourceGet) | **Get** /api/v1/datasources/{name} | Get a datasource
[**DatasourceList**](DatasourceAPI.md#DatasourceList) | **Get** /api/v1/datasources | List the existing datasource(s)
[**DatasourceTest**](DatasourceAPI.md#DatasourceTest) | **Patch** /api/v1/datasources | Test a datasource
[**DatasourceUpdate**](DatasourceAPI.md#DatasourceUpdate) | **Put** /api/v1/datasources | Update an existing datasource



## DatasourceAdd

> DatasourceList200ResponseInner DatasourceAdd(ctx).DatasourceUpdateRequest(datasourceUpdateRequest).Execute()

Register a new datasource



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
	datasourceUpdateRequest := openapiclient.datasource_update_request{DNSDatasource: openapiclient.NewDNSDatasource("Type_example", "DNS_Datasource", "Lookup_example")} // DatasourceUpdateRequest | Datasource to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceAPI.DatasourceAdd(context.Background()).DatasourceUpdateRequest(datasourceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceAdd`: DatasourceList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DatasourceAPI.DatasourceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceUpdateRequest** | [**DatasourceUpdateRequest**](DatasourceUpdateRequest.md) | Datasource to register | 

### Return type

[**DatasourceList200ResponseInner**](DatasourceList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceDelete

> DatasourceDelete(ctx, name).Execute()

Delete a datasource



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
	name := "name_example" // string | Name of the datasource to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatasourceAPI.DatasourceDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the datasource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceDeleteRequest struct via the builder pattern


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


## DatasourceGet

> DatasourceList200ResponseInner DatasourceGet(ctx, name).Execute()

Get a datasource



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
	name := "name_example" // string | Name of the datasource to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceAPI.DatasourceGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceGet`: DatasourceList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DatasourceAPI.DatasourceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the datasource to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasourceList200ResponseInner**](DatasourceList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceList

> []DatasourceList200ResponseInner DatasourceList(ctx).Execute()

List the existing datasource(s)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceAPI.DatasourceList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceList`: []DatasourceList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DatasourceAPI.DatasourceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceListRequest struct via the builder pattern


### Return type

[**[]DatasourceList200ResponseInner**](DatasourceList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceTest

> DatasourceTest200Response DatasourceTest(ctx).DataSourceTestRequest(dataSourceTestRequest).Execute()

Test a datasource



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
	dataSourceTestRequest := *openapiclient.NewDataSourceTestRequest(openapiclient.DataSourceTestRequest_ds{DNSDatasource: openapiclient.NewDNSDatasource("Type_example", "DNS_Datasource", "Lookup_example")}) // DataSourceTestRequest | Datasource to test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceAPI.DatasourceTest(context.Background()).DataSourceTestRequest(dataSourceTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceTest`: DatasourceTest200Response
	fmt.Fprintf(os.Stdout, "Response from `DatasourceAPI.DatasourceTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceTestRequest** | [**DataSourceTestRequest**](DataSourceTestRequest.md) | Datasource to test | 

### Return type

[**DatasourceTest200Response**](DatasourceTest200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceUpdate

> DatasourceList200ResponseInner DatasourceUpdate(ctx).DatasourceUpdateRequest(datasourceUpdateRequest).Execute()

Update an existing datasource



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
	datasourceUpdateRequest := openapiclient.datasource_update_request{DNSDatasource: openapiclient.NewDNSDatasource("Type_example", "DNS_Datasource", "Lookup_example")} // DatasourceUpdateRequest | Datasource to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatasourceAPI.DatasourceUpdate(context.Background()).DatasourceUpdateRequest(datasourceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatasourceAPI.DatasourceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceUpdate`: DatasourceList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DatasourceAPI.DatasourceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceUpdateRequest** | [**DatasourceUpdateRequest**](DatasourceUpdateRequest.md) | Datasource to update | 

### Return type

[**DatasourceList200ResponseInner**](DatasourceList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

