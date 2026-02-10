# \DiscoveryCampaignAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryCampaignAdd**](DiscoveryCampaignAPI.md#DiscoveryCampaignAdd) | **Post** /api/v1/discovery/campaigns | Create a new discovery campaign
[**DiscoveryCampaignDelete**](DiscoveryCampaignAPI.md#DiscoveryCampaignDelete) | **Delete** /api/v1/discovery/campaigns/{name} | Delete a discovery campaign
[**DiscoveryCampaignFlush**](DiscoveryCampaignAPI.md#DiscoveryCampaignFlush) | **Patch** /api/v1/discovery/campaigns/{name} | Flush a discovery campaign
[**DiscoveryCampaignGet**](DiscoveryCampaignAPI.md#DiscoveryCampaignGet) | **Get** /api/v1/discovery/campaigns/{name} | Retrieve a discovery campaign
[**DiscoveryCampaignList**](DiscoveryCampaignAPI.md#DiscoveryCampaignList) | **Get** /api/v1/discovery/campaigns | List discovery campaign(s)
[**DiscoveryCampaignUpdate**](DiscoveryCampaignAPI.md#DiscoveryCampaignUpdate) | **Put** /api/v1/discovery/campaigns | Update a discovery campaign



## DiscoveryCampaignAdd

> DiscoveryCampaignResponse DiscoveryCampaignAdd(ctx).DiscoveryCampaign(discoveryCampaign).Execute()

Create a new discovery campaign



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
	discoveryCampaign := *openapiclient.NewDiscoveryCampaign("DiscoveryDMZ01", *openapiclient.NewDiscoveryCampaignAuthorizationLevels(*openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated")), false, false, true, true) // DiscoveryCampaign | Discovery campaign to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignAdd(context.Background()).DiscoveryCampaign(discoveryCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryCampaignAdd`: DiscoveryCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCampaignAPI.DiscoveryCampaignAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryCampaign** | [**DiscoveryCampaign**](DiscoveryCampaign.md) | Discovery campaign to register | 

### Return type

[**DiscoveryCampaignResponse**](DiscoveryCampaignResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryCampaignDelete

> DiscoveryCampaignDelete(ctx, name).Execute()

Delete a discovery campaign



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
	name := "Discovery-DMZ01" // string | The name of the discovery campaign to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the discovery campaign to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignDeleteRequest struct via the builder pattern


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


## DiscoveryCampaignFlush

> DiscoveryCampaignFlush(ctx, name).Execute()

Flush a discovery campaign



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
	name := "name_example" // string | The name of the discovery campaign to flush

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignFlush(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignFlush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the discovery campaign to flush | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignFlushRequest struct via the builder pattern


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


## DiscoveryCampaignGet

> DiscoveryCampaignResponse DiscoveryCampaignGet(ctx, name).Execute()

Retrieve a discovery campaign



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
	name := "Discovery-DMZ01" // string | The name of the discovery campaign to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryCampaignGet`: DiscoveryCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCampaignAPI.DiscoveryCampaignGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the discovery campaign to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscoveryCampaignResponse**](DiscoveryCampaignResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryCampaignList

> []DiscoveryCampaignResponse DiscoveryCampaignList(ctx).Execute()

List discovery campaign(s)



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
	resp, r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryCampaignList`: []DiscoveryCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCampaignAPI.DiscoveryCampaignList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignListRequest struct via the builder pattern


### Return type

[**[]DiscoveryCampaignResponse**](DiscoveryCampaignResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveryCampaignUpdate

> DiscoveryCampaignResponse DiscoveryCampaignUpdate(ctx).DiscoveryCampaign(discoveryCampaign).Execute()

Update a discovery campaign



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
	discoveryCampaign := *openapiclient.NewDiscoveryCampaign("DiscoveryDMZ01", *openapiclient.NewDiscoveryCampaignAuthorizationLevels(*openapiclient.NewAuthorizationLevel("authenticated"), *openapiclient.NewAuthorizationLevel("authenticated")), false, false, true, true) // DiscoveryCampaign | Discovery campaign to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryCampaignAPI.DiscoveryCampaignUpdate(context.Background()).DiscoveryCampaign(discoveryCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCampaignAPI.DiscoveryCampaignUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoveryCampaignUpdate`: DiscoveryCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCampaignAPI.DiscoveryCampaignUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveryCampaignUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveryCampaign** | [**DiscoveryCampaign**](DiscoveryCampaign.md) | Discovery campaign to update | 

### Return type

[**DiscoveryCampaignResponse**](DiscoveryCampaignResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

