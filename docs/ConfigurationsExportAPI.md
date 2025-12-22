# \ConfigurationsExportAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportExportItems**](ConfigurationsExportAPI.md#ExportExportItems) | **Post** /api/v1/system/configurations/export | Export configuration items
[**ExportImportItems**](ConfigurationsExportAPI.md#ExportImportItems) | **Post** /api/v1/system/configurations/import | Import configuration items
[**ExportListItems**](ConfigurationsExportAPI.md#ExportListItems) | **Get** /api/v1/system/configurations/export | List exportable configuration items



## ExportExportItems

> HorizonExport ExportExportItems(ctx).HorizonExportableItems(horizonExportableItems).Execute()

Export configuration items



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
	horizonExportableItems := *openapiclient.NewHorizonExportableItems() // HorizonExportableItems | Items to export

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsExportAPI.ExportExportItems(context.Background()).HorizonExportableItems(horizonExportableItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsExportAPI.ExportExportItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportExportItems`: HorizonExport
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsExportAPI.ExportExportItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportExportItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **horizonExportableItems** | [**HorizonExportableItems**](HorizonExportableItems.md) | Items to export | 

### Return type

[**HorizonExport**](HorizonExport.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportImportItems

> HorizonImportedItemsSummary ExportImportItems(ctx).HorizonExport(horizonExport).Execute()

Import configuration items



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
	horizonExport := *openapiclient.NewHorizonExport() // HorizonExport | Items to import

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsExportAPI.ExportImportItems(context.Background()).HorizonExport(horizonExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsExportAPI.ExportImportItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportImportItems`: HorizonImportedItemsSummary
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsExportAPI.ExportImportItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportImportItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **horizonExport** | [**HorizonExport**](HorizonExport.md) | Items to import | 

### Return type

[**HorizonImportedItemsSummary**](HorizonImportedItemsSummary.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportListItems

> HorizonExportableItems ExportListItems(ctx).Execute()

List exportable configuration items



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
	resp, r, err := apiClient.ConfigurationsExportAPI.ExportListItems(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsExportAPI.ExportListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportListItems`: HorizonExportableItems
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsExportAPI.ExportListItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportListItemsRequest struct via the builder pattern


### Return type

[**HorizonExportableItems**](HorizonExportableItems.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

