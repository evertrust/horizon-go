# \ReportCsvAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportCsvDelete**](ReportCsvAPI.md#ReportCsvDelete) | **Delete** /api/v1/reports/{uuid} | Delete a report CSV by its UUID
[**ReportCsvDownloadByUUID**](ReportCsvAPI.md#ReportCsvDownloadByUUID) | **Get** /reports/{uuid} | Download a CSV report by its UUID
[**ReportCsvListAllMetadata**](ReportCsvAPI.md#ReportCsvListAllMetadata) | **Get** /api/v1/reports | List all registered CSV reports
[**ReportCsvListAllMetadataByReportID**](ReportCsvAPI.md#ReportCsvListAllMetadataByReportID) | **Get** /api/v1/reports/{reportName} | List all CSV reports linked to one report



## ReportCsvDelete

> ReportCsvDelete(ctx, uuid).Execute()

Delete a report CSV by its UUID



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
	uuid := "uuid_example" // string | The UUID of the report CSV to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportCsvAPI.ReportCsvDelete(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCsvAPI.ReportCsvDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | The UUID of the report CSV to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportCsvDeleteRequest struct via the builder pattern


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


## ReportCsvDownloadByUUID

> string ReportCsvDownloadByUUID(ctx, uuid).Execute()

Download a CSV report by its UUID



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
	uuid := "uuid_example" // string | The identifier of the CSV to download (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCsvAPI.ReportCsvDownloadByUUID(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCsvAPI.ReportCsvDownloadByUUID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportCsvDownloadByUUID`: string
	fmt.Fprintf(os.Stdout, "Response from `ReportCsvAPI.ReportCsvDownloadByUUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | The identifier of the CSV to download (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportCsvDownloadByUUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportCsvListAllMetadata

> []ReportCSVMetadata ReportCsvListAllMetadata(ctx).Expired(expired).Execute()

List all registered CSV reports



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
	expired := true // bool | If set to `true` the API will return expired and available reports CSV (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCsvAPI.ReportCsvListAllMetadata(context.Background()).Expired(expired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCsvAPI.ReportCsvListAllMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportCsvListAllMetadata`: []ReportCSVMetadata
	fmt.Fprintf(os.Stdout, "Response from `ReportCsvAPI.ReportCsvListAllMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportCsvListAllMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expired** | **bool** | If set to &#x60;true&#x60; the API will return expired and available reports CSV | [default to false]

### Return type

[**[]ReportCSVMetadata**](ReportCSVMetadata.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportCsvListAllMetadataByReportID

> []ReportCSVMetadata ReportCsvListAllMetadataByReportID(ctx, reportName).Expired(expired).Execute()

List all CSV reports linked to one report



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
	reportName := "reportName_example" // string | Report identifier to list the CSV reports from
	expired := true // bool | If set to `true` the API will return expired and available reports CSV (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCsvAPI.ReportCsvListAllMetadataByReportID(context.Background(), reportName).Expired(expired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCsvAPI.ReportCsvListAllMetadataByReportID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportCsvListAllMetadataByReportID`: []ReportCSVMetadata
	fmt.Fprintf(os.Stdout, "Response from `ReportCsvAPI.ReportCsvListAllMetadataByReportID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportName** | **string** | Report identifier to list the CSV reports from | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportCsvListAllMetadataByReportIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expired** | **bool** | If set to &#x60;true&#x60; the API will return expired and available reports CSV | [default to false]

### Return type

[**[]ReportCSVMetadata**](ReportCSVMetadata.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

