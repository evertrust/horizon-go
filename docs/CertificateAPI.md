# \CertificateAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateAggregate**](CertificateAPI.md#CertificateAggregate) | **Post** /api/v1/certificates/aggregate | Certificate aggregation
[**CertificateCsv**](CertificateAPI.md#CertificateCsv) | **Post** /api/v1/certificates/csv | Export certificates
[**CertificateDictionary**](CertificateAPI.md#CertificateDictionary) | **Get** /api/v1/certificates/search/dictionary | Retrieve the certificate search dictionary
[**CertificateFind**](CertificateAPI.md#CertificateFind) | **Post** /api/v1/certificates/find | Find a certificate
[**CertificateGetId**](CertificateAPI.md#CertificateGetId) | **Get** /api/v1/certificates/{id} | Retrieve a certificate
[**CertificateGetPem**](CertificateAPI.md#CertificateGetPem) | **Get** /api/v1/certificates/{pem} | Retrieve a certificate by PEM
[**CertificateList**](CertificateAPI.md#CertificateList) | **Post** /api/v1/certificates | List certificates
[**CertificateRun**](CertificateAPI.md#CertificateRun) | **Patch** /api/v1/certificates/run/{id}/{triggerName}/{event} | Run a certificate trigger
[**CertificateSearch**](CertificateAPI.md#CertificateSearch) | **Post** /api/v1/certificates/search | Search certificates



## CertificateAggregate

> CertificateAggregateResultResponse CertificateAggregate(ctx).CertificateAggregateQuery(certificateAggregateQuery).EnableAnalytics(enableAnalytics).Execute()

Certificate aggregation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	certificateAggregateQuery := *openapiclient.NewCertificateAggregateQuery() // CertificateAggregateQuery | The certificate aggregation query
	enableAnalytics := true // bool | Use the analytics database if enabled. `true` if not specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateAggregate(context.Background()).CertificateAggregateQuery(certificateAggregateQuery).EnableAnalytics(enableAnalytics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateAggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateAggregate`: CertificateAggregateResultResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateAggregate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateAggregateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateAggregateQuery** | [**CertificateAggregateQuery**](CertificateAggregateQuery.md) | The certificate aggregation query | 
 **enableAnalytics** | **bool** | Use the analytics database if enabled. &#x60;true&#x60; if not specified. | 

### Return type

[**CertificateAggregateResultResponse**](CertificateAggregateResultResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateCsv

> CertificateCsv(ctx).CertificateSearchQuery(certificateSearchQuery).EnableAnalytics(enableAnalytics).Execute()

Export certificates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	certificateSearchQuery := *openapiclient.NewCertificateSearchQuery() // CertificateSearchQuery | The certificate search query
	enableAnalytics := true // bool | Use the analytics database if enabled. `true` if not specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificateAPI.CertificateCsv(context.Background()).CertificateSearchQuery(certificateSearchQuery).EnableAnalytics(enableAnalytics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateSearchQuery** | [**CertificateSearchQuery**](CertificateSearchQuery.md) | The certificate search query | 
 **enableAnalytics** | **bool** | Use the analytics database if enabled. &#x60;true&#x60; if not specified. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateDictionary

> CertificateSearchDictionaryResponse CertificateDictionary(ctx).Execute()

Retrieve the certificate search dictionary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateDictionary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateDictionary`: CertificateSearchDictionaryResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateDictionary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateDictionaryRequest struct via the builder pattern


### Return type

[**CertificateSearchDictionaryResponse**](CertificateSearchDictionaryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateFind

> CertificateWithPermissionsResponse CertificateFind(ctx).CertificateFindRequest(certificateFindRequest).Execute()

Find a certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	certificateFindRequest := openapiclient.certificate_find_request{FindCertificateById: openapiclient.NewFindCertificateById("6448d56b310000400063f014")} // CertificateFindRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateFind(context.Background()).CertificateFindRequest(certificateFindRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateFind`: CertificateWithPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateFindRequest** | [**CertificateFindRequest**](CertificateFindRequest.md) |  | 

### Return type

[**CertificateWithPermissionsResponse**](CertificateWithPermissionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateGetId

> CertificateWithPermissionsResponse CertificateGetId(ctx, id).Execute()

Retrieve a certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	id := "644796623000003800cc6c4b" // string | The ID of the certificate to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateGetId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateGetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateGetId`: CertificateWithPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateGetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the certificate to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateGetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateWithPermissionsResponse**](CertificateWithPermissionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateGetPem

> CertificateResponse CertificateGetPem(ctx, pem).Execute()

Retrieve a certificate by PEM



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	pem := "-----BEGIN%20CERTIFICATE----- ... -----END%20CERTIFICATE-----" // string | The URL encoded PEM encoded value of the certificate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateGetPem(context.Background(), pem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateGetPem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateGetPem`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateGetPem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pem** | **string** | The URL encoded PEM encoded value of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateGetPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateList

> []CertificateWithPermissionsResponse CertificateList(ctx).RequestBody(requestBody).Execute()

List certificates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	requestBody := []string{"644796623000003800cc6c4b"} // []string | The list of certificates IDs to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateList(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateList`: []CertificateWithPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | The list of certificates IDs to fetch | 

### Return type

[**[]CertificateWithPermissionsResponse**](CertificateWithPermissionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateRun

> CertificateWithPermissionsResponse CertificateRun(ctx, id, triggerName, event).Execute()

Run a certificate trigger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	id := "644796623000003800cc6c4b" // string | The ID of the certificate
	triggerName := "TestTrigger" // string | The name of the trigger
	event := "event_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateRun(context.Background(), id, triggerName, event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateRun`: CertificateWithPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the certificate | 
**triggerName** | **string** | The name of the trigger | 
**event** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateWithPermissionsResponse**](CertificateWithPermissionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateSearch

> CertificateSearchResultsResponse CertificateSearch(ctx).CertificateSearchQuery(certificateSearchQuery).EnableAnalytics(enableAnalytics).Execute()

Search certificates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evertrust/horizon-go"
)

func main() {
	certificateSearchQuery := *openapiclient.NewCertificateSearchQuery() // CertificateSearchQuery | 
	enableAnalytics := true // bool | Use the analytics database if enabled. `true` if not specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAPI.CertificateSearch(context.Background()).CertificateSearchQuery(certificateSearchQuery).EnableAnalytics(enableAnalytics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAPI.CertificateSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateSearch`: CertificateSearchResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAPI.CertificateSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateSearchQuery** | [**CertificateSearchQuery**](CertificateSearchQuery.md) |  | 
 **enableAnalytics** | **bool** | Use the analytics database if enabled. &#x60;true&#x60; if not specified. | 

### Return type

[**CertificateSearchResultsResponse**](CertificateSearchResultsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

