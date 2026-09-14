# \DashboardAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityPrincipalDashboardAdd**](DashboardAPI.md#SecurityPrincipalDashboardAdd) | **Post** /api/v1/security/principals/dashboards | Add a dashboard for the authenticated principal
[**SecurityPrincipalDashboardDelete**](DashboardAPI.md#SecurityPrincipalDashboardDelete) | **Delete** /api/v1/security/principals/dashboards/{name} | Delete an existing dashboard for the authenticated principal
[**SecurityPrincipalDashboardGet**](DashboardAPI.md#SecurityPrincipalDashboardGet) | **Get** /api/v1/security/principals/dashboards/{name} | Retrieve an existing dashboard for the authenticated principal
[**SecurityPrincipalDashboardList**](DashboardAPI.md#SecurityPrincipalDashboardList) | **Get** /api/v1/security/principals/dashboards | List the authenticated principal dashboards
[**SecurityPrincipalDashboardUpdate**](DashboardAPI.md#SecurityPrincipalDashboardUpdate) | **Put** /api/v1/security/principals/dashboards | Update a dashboard for the authenticated principal



## SecurityPrincipalDashboardAdd

> DashboardResponse SecurityPrincipalDashboardAdd(ctx).Dashboard(dashboard).Execute()

Add a dashboard for the authenticated principal



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
	dashboard := *openapiclient.NewDashboard([]openapiclient.Chart{*openapiclient.NewChart([]string{"Colors_example"}, []string{"Fields_example"}, false, "Certificate status on the WebRA", openapiclient.ChartType("table"))}, "My Certificate Dashboard", "certificate") // Dashboard | Dashboard to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.SecurityPrincipalDashboardAdd(context.Background()).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.SecurityPrincipalDashboardAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalDashboardAdd`: DashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.SecurityPrincipalDashboardAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalDashboardAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | Dashboard to add | 

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalDashboardDelete

> SecurityPrincipalDashboardDelete(ctx, name).Execute()

Delete an existing dashboard for the authenticated principal



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
	name := "name_example" // string | The name of the dashboard

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardAPI.SecurityPrincipalDashboardDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.SecurityPrincipalDashboardDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalDashboardDeleteRequest struct via the builder pattern


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


## SecurityPrincipalDashboardGet

> DashboardResponse SecurityPrincipalDashboardGet(ctx, name).Execute()

Retrieve an existing dashboard for the authenticated principal



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
	name := "name_example" // string | The name of the dashboard

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.SecurityPrincipalDashboardGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.SecurityPrincipalDashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalDashboardGet`: DashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.SecurityPrincipalDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalDashboardList

> []DashboardResponse SecurityPrincipalDashboardList(ctx).Type_(type_).Execute()

List the authenticated principal dashboards



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
	type_ := "type__example" // string | The type of dashboards to filter on (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.SecurityPrincipalDashboardList(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.SecurityPrincipalDashboardList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalDashboardList`: []DashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.SecurityPrincipalDashboardList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of dashboards to filter on | 

### Return type

[**[]DashboardResponse**](DashboardResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityPrincipalDashboardUpdate

> DashboardResponse SecurityPrincipalDashboardUpdate(ctx).Dashboard(dashboard).Execute()

Update a dashboard for the authenticated principal



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
	dashboard := *openapiclient.NewDashboard([]openapiclient.Chart{*openapiclient.NewChart([]string{"Colors_example"}, []string{"Fields_example"}, false, "Certificate status on the WebRA", openapiclient.ChartType("table"))}, "My Certificate Dashboard", "certificate") // Dashboard | Dashboard to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.SecurityPrincipalDashboardUpdate(context.Background()).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.SecurityPrincipalDashboardUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityPrincipalDashboardUpdate`: DashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.SecurityPrincipalDashboardUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityPrincipalDashboardUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | Dashboard to update | 

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

