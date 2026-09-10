# \RequestAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestAggregate**](RequestAPI.md#RequestAggregate) | **Post** /api/v1/requests/aggregate | Request aggregation
[**RequestApprove**](RequestAPI.md#RequestApprove) | **Post** /api/v1/requests/approve | Approve a request
[**RequestCancel**](RequestAPI.md#RequestCancel) | **Post** /api/v1/requests/cancel | Cancel a request
[**RequestCertificateProfile**](RequestAPI.md#RequestCertificateProfile) | **Get** /api/v1/requests/profiles | List profiles
[**RequestCsv**](RequestAPI.md#RequestCsv) | **Post** /api/v1/requests/csv | Export requests
[**RequestDeny**](RequestAPI.md#RequestDeny) | **Post** /api/v1/requests/deny | Deny a request
[**RequestDictionary**](RequestAPI.md#RequestDictionary) | **Get** /api/v1/requests/search/dictionary | Retrieve the request search dictionary
[**RequestGet**](RequestAPI.md#RequestGet) | **Get** /api/v1/requests/{id} | Retrieve a request
[**RequestSearch**](RequestAPI.md#RequestSearch) | **Post** /api/v1/requests/search | Search requests
[**RequestSubmit**](RequestAPI.md#RequestSubmit) | **Post** /api/v1/requests/submit | Submit a request
[**RequestTemplate**](RequestAPI.md#RequestTemplate) | **Post** /api/v1/requests/template | Retrieve a request template



## RequestAggregate

> RequestAggregateResultResponse RequestAggregate(ctx).RequestAggregateQuery(requestAggregateQuery).Execute()

Request aggregation



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
	requestAggregateQuery := *openapiclient.NewRequestAggregateQuery() // RequestAggregateQuery | The request aggregation query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestAggregate(context.Background()).RequestAggregateQuery(requestAggregateQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestAggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestAggregate`: RequestAggregateResultResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestAggregate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAggregateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAggregateQuery** | [**RequestAggregateQuery**](RequestAggregateQuery.md) | The request aggregation query | 

### Return type

[**RequestAggregateResultResponse**](RequestAggregateResultResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestApprove

> RequestApprove200Response RequestApprove(ctx).RequestApproveRequest(requestApproveRequest).Execute()

Approve a request



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
	requestApproveRequest := openapiclient.request_approve_request{EstEnrollRequestOnApprove: openapiclient.NewEstEnrollRequestOnApprove("6448d56b310000400063f014", "Module_example", "Workflow_example")} // RequestApproveRequest | The request to approve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(requestApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestApprove`: RequestApprove200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestApproveRequest** | [**RequestApproveRequest**](RequestApproveRequest.md) | The request to approve | 

### Return type

[**RequestApprove200Response**](RequestApprove200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestCancel

> RequestApprove200Response RequestCancel(ctx).RequestCancelRequest(requestCancelRequest).Execute()

Cancel a request



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
	requestCancelRequest := *openapiclient.NewRequestCancelRequest("6448d56b310000400063f014", openapiclient.Module("webra"), openapiclient.Workflow("enroll")) // RequestCancelRequest | The Request to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(requestCancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestCancel`: RequestApprove200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCancelRequest** | [**RequestCancelRequest**](RequestCancelRequest.md) | The Request to cancel | 

### Return type

[**RequestApprove200Response**](RequestApprove200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestCertificateProfile

> []RequestableCertificateProfileResponse RequestCertificateProfile(ctx).Module(module).Workflow(workflow).Execute()

List profiles



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
	module := "module_example" // string |  (optional)
	workflow := "workflow_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestCertificateProfile(context.Background()).Module(module).Workflow(workflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestCertificateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestCertificateProfile`: []RequestableCertificateProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestCertificateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestCertificateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **module** | **string** |  | 
 **workflow** | **string** |  | 

### Return type

[**[]RequestableCertificateProfileResponse**](RequestableCertificateProfileResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestCsv

> RequestCsv(ctx).RequestSearchQuery(requestSearchQuery).Execute()

Export requests



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
	requestSearchQuery := *openapiclient.NewRequestSearchQuery() // RequestSearchQuery | The request search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestAPI.RequestCsv(context.Background()).RequestSearchQuery(requestSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSearchQuery** | [**RequestSearchQuery**](RequestSearchQuery.md) | The request search query | 

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


## RequestDeny

> RequestApprove200Response RequestDeny(ctx).RequestDenyRequest(requestDenyRequest).Execute()

Deny a request



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
	requestDenyRequest := *openapiclient.NewRequestDenyRequest("6448d56b310000400063f014", openapiclient.Module("webra"), openapiclient.Workflow("enroll")) // RequestDenyRequest | The request to deny

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(requestDenyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestDeny``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestDeny`: RequestApprove200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestDeny`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestDenyRequest** | [**RequestDenyRequest**](RequestDenyRequest.md) | The request to deny | 

### Return type

[**RequestApprove200Response**](RequestApprove200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestDictionary

> RequestSearchDictionaryResponse RequestDictionary(ctx).Execute()

Retrieve the request search dictionary



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
	resp, r, err := apiClient.RequestAPI.RequestDictionary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestDictionary`: RequestSearchDictionaryResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestDictionary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestDictionaryRequest struct via the builder pattern


### Return type

[**RequestSearchDictionaryResponse**](RequestSearchDictionaryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestGet

> RequestGet200Response RequestGet(ctx, id).Execute()

Retrieve a request



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
	id := "id_example" // string | The request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestGet`: RequestGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestGet200Response**](RequestGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSearch

> RequestSearchResultsResponse RequestSearch(ctx).RequestSearchQuery(requestSearchQuery).Execute()

Search requests



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
	requestSearchQuery := *openapiclient.NewRequestSearchQuery() // RequestSearchQuery | The request search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestSearch(context.Background()).RequestSearchQuery(requestSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestSearch`: RequestSearchResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSearchQuery** | [**RequestSearchQuery**](RequestSearchQuery.md) | The request search query | 

### Return type

[**RequestSearchResultsResponse**](RequestSearchResultsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSubmit

> RequestSubmit201Response RequestSubmit(ctx).RequestSubmitRequest(requestSubmitRequest).Execute()

Submit a request



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
	requestSubmitRequest := openapiclient.request_submit_request{EstEnrollRequestOnSubmit: openapiclient.NewEstEnrollRequestOnSubmit("Profile_example", "Module_example", "Workflow_example")} // RequestSubmitRequest | The Request to submit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(requestSubmitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestSubmit`: RequestSubmit201Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSubmitRequest** | [**RequestSubmitRequest**](RequestSubmitRequest.md) | The Request to submit | 

### Return type

[**RequestSubmit201Response**](RequestSubmit201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestTemplate

> RequestTemplate200Response RequestTemplate(ctx).RequestTemplateRequest(requestTemplateRequest).TermsOfService(termsOfService).Execute()

Retrieve a request template



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
	requestTemplateRequest := openapiclient.request_template_request{EstEnrollRequestOnTemplate: openapiclient.NewEstEnrollRequestOnTemplate("Module_example", "Workflow_example", "webra_centralized")} // RequestTemplateRequest | The request on which to return the template
	termsOfService := true // bool | If true, include the Terms of Service entry configured on the profile in the response. Only applies to WebRA, EST and SCEP enroll requests. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(requestTemplateRequest).TermsOfService(termsOfService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTemplate`: RequestTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestTemplateRequest** | [**RequestTemplateRequest**](RequestTemplateRequest.md) | The request on which to return the template | 
 **termsOfService** | **bool** | If true, include the Terms of Service entry configured on the profile in the response. Only applies to WebRA, EST and SCEP enroll requests. | [default to false]

### Return type

[**RequestTemplate200Response**](RequestTemplate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

