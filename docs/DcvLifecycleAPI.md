# \DcvLifecycleAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DcvLifecycleCancel**](DcvLifecycleAPI.md#DcvLifecycleCancel) | **Post** /api/v1/dcv/lifecycle/policies/{name}/cancel | Cancel an active DCV policy run
[**DcvLifecycleEventsListByPolicy**](DcvLifecycleAPI.md#DcvLifecycleEventsListByPolicy) | **Post** /api/v1/dcv/lifecycle/events/{policy} | List DCV lifecycle events for a policy
[**DcvLifecycleEventsListByPolicyAndDomain**](DcvLifecycleAPI.md#DcvLifecycleEventsListByPolicyAndDomain) | **Post** /api/v1/dcv/lifecycle/events/{policy}/{domain} | List DCV lifecycle events for a specific domain
[**DcvLifecycleGet**](DcvLifecycleAPI.md#DcvLifecycleGet) | **Get** /api/v1/dcv/lifecycle/policies/{name} | Get DCV policy status
[**DcvLifecycleList**](DcvLifecycleAPI.md#DcvLifecycleList) | **Get** /api/v1/dcv/lifecycle/policies | List DCV policies
[**DcvLifecycleRun**](DcvLifecycleAPI.md#DcvLifecycleRun) | **Post** /api/v1/dcv/lifecycle/policies/{name}/run | Trigger DCV policy run for all domains
[**DcvLifecycleTriggerDomainRun**](DcvLifecycleAPI.md#DcvLifecycleTriggerDomainRun) | **Post** /api/v1/dcv/lifecycle/policies/{name}/run/{domain} | Trigger domain validation for a specific domain



## DcvLifecycleCancel

> DcvLifecycleCancel(ctx, name).Execute()

Cancel an active DCV policy run



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
	name := "name_example" // string | Name of the DCV policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DcvLifecycleAPI.DcvLifecycleCancel(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the DCV policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleCancelRequest struct via the builder pattern


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


## DcvLifecycleEventsListByPolicy

> DCVLifecycleEventSearchResults DcvLifecycleEventsListByPolicy(ctx, policy).DCVLifecycleEventSearchQuery(dCVLifecycleEventSearchQuery).Execute()

List DCV lifecycle events for a policy



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
	policy := "policy_example" // string | Name of the DCV policy
	dCVLifecycleEventSearchQuery := *openapiclient.NewDCVLifecycleEventSearchQuery() // DCVLifecycleEventSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvLifecycleAPI.DcvLifecycleEventsListByPolicy(context.Background(), policy).DCVLifecycleEventSearchQuery(dCVLifecycleEventSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleEventsListByPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvLifecycleEventsListByPolicy`: DCVLifecycleEventSearchResults
	fmt.Fprintf(os.Stdout, "Response from `DcvLifecycleAPI.DcvLifecycleEventsListByPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Name of the DCV policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleEventsListByPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dCVLifecycleEventSearchQuery** | [**DCVLifecycleEventSearchQuery**](DCVLifecycleEventSearchQuery.md) |  | 

### Return type

[**DCVLifecycleEventSearchResults**](DCVLifecycleEventSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvLifecycleEventsListByPolicyAndDomain

> DCVLifecycleEventSearchResults DcvLifecycleEventsListByPolicyAndDomain(ctx, policy, domain).DCVLifecycleEventSearchQuery(dCVLifecycleEventSearchQuery).Execute()

List DCV lifecycle events for a specific domain



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
	policy := "policy_example" // string | Name of the DCV policy
	domain := "domain_example" // string | The domain hostname to retrieve events for
	dCVLifecycleEventSearchQuery := *openapiclient.NewDCVLifecycleEventSearchQuery() // DCVLifecycleEventSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvLifecycleAPI.DcvLifecycleEventsListByPolicyAndDomain(context.Background(), policy, domain).DCVLifecycleEventSearchQuery(dCVLifecycleEventSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleEventsListByPolicyAndDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvLifecycleEventsListByPolicyAndDomain`: DCVLifecycleEventSearchResults
	fmt.Fprintf(os.Stdout, "Response from `DcvLifecycleAPI.DcvLifecycleEventsListByPolicyAndDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string** | Name of the DCV policy | 
**domain** | **string** | The domain hostname to retrieve events for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleEventsListByPolicyAndDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dCVLifecycleEventSearchQuery** | [**DCVLifecycleEventSearchQuery**](DCVLifecycleEventSearchQuery.md) |  | 

### Return type

[**DCVLifecycleEventSearchResults**](DCVLifecycleEventSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvLifecycleGet

> DCVPolicyStatusResponse DcvLifecycleGet(ctx, name).Execute()

Get DCV policy status



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
	name := "name_example" // string | Name of the DCV policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvLifecycleAPI.DcvLifecycleGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvLifecycleGet`: DCVPolicyStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvLifecycleAPI.DcvLifecycleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the DCV policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DCVPolicyStatusResponse**](DCVPolicyStatusResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvLifecycleList

> []DCVPolicyLifecycleResponse DcvLifecycleList(ctx).Execute()

List DCV policies



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
	resp, r, err := apiClient.DcvLifecycleAPI.DcvLifecycleList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvLifecycleList`: []DCVPolicyLifecycleResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvLifecycleAPI.DcvLifecycleList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleListRequest struct via the builder pattern


### Return type

[**[]DCVPolicyLifecycleResponse**](DCVPolicyLifecycleResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvLifecycleRun

> DcvLifecycleRun(ctx, name).Execute()

Trigger DCV policy run for all domains



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
	name := "name_example" // string | Name of the DCV policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DcvLifecycleAPI.DcvLifecycleRun(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the DCV policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleRunRequest struct via the builder pattern


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


## DcvLifecycleTriggerDomainRun

> DcvLifecycleTriggerDomainRun(ctx, name, domain).Execute()

Trigger domain validation for a specific domain



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
	name := "name_example" // string | Name of the DCV policy
	domain := "domain_example" // string | The domain hostname to trigger validation for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DcvLifecycleAPI.DcvLifecycleTriggerDomainRun(context.Background(), name, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvLifecycleAPI.DcvLifecycleTriggerDomainRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the DCV policy | 
**domain** | **string** | The domain hostname to trigger validation for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvLifecycleTriggerDomainRunRequest struct via the builder pattern


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

