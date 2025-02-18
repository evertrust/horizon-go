# \CertificateAnalyticsAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCertificateFlush**](CertificateAnalyticsAPI.md#AnalyticsCertificateFlush) | **Delete** /api/v1/analytics/certificates | Flush certificate analytics synchronization
[**AnalyticsCertificateGet**](CertificateAnalyticsAPI.md#AnalyticsCertificateGet) | **Get** /api/v1/analytics/certificates | Retrieve the certificate analytics status
[**AnalyticsCertificateUpdate**](CertificateAnalyticsAPI.md#AnalyticsCertificateUpdate) | **Patch** /api/v1/analytics/certificates | Schedule a new certificate analytics synchronization



## AnalyticsCertificateFlush

> AnalyticsCertificateFlush(ctx).Execute()

Flush certificate analytics synchronization



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
	r, err := apiClient.CertificateAnalyticsAPI.AnalyticsCertificateFlush(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAnalyticsAPI.AnalyticsCertificateFlush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCertificateFlushRequest struct via the builder pattern


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


## AnalyticsCertificateGet

> []AnalyticsStatus AnalyticsCertificateGet(ctx).Execute()

Retrieve the certificate analytics status



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
	resp, r, err := apiClient.CertificateAnalyticsAPI.AnalyticsCertificateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAnalyticsAPI.AnalyticsCertificateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCertificateGet`: []AnalyticsStatus
	fmt.Fprintf(os.Stdout, "Response from `CertificateAnalyticsAPI.AnalyticsCertificateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCertificateGetRequest struct via the builder pattern


### Return type

[**[]AnalyticsStatus**](AnalyticsStatus.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsCertificateUpdate

> AnalyticsCertificateUpdate(ctx).Execute()

Schedule a new certificate analytics synchronization



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
	r, err := apiClient.CertificateAnalyticsAPI.AnalyticsCertificateUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAnalyticsAPI.AnalyticsCertificateUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCertificateUpdateRequest struct via the builder pattern


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

