# \ThirdpartyConnectorAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThirdpartyConnectorAdd**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorAdd) | **Post** /api/v1/thirdparty/connectors | Register a new third party connector
[**ThirdpartyConnectorDelete**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorDelete) | **Delete** /api/v1/thirdparty/connectors/{name} | Delete an existing third party connector
[**ThirdpartyConnectorGet**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorGet) | **Get** /api/v1/thirdparty/connectors/{name} | Retrieve an existing third party connector
[**ThirdpartyConnectorList**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorList) | **Get** /api/v1/thirdparty/connectors | List the existing third party connector(s)
[**ThirdpartyConnectorRetry**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorRetry) | **Patch** /api/v1/thirdparty/connectors/{name} | Retry failed triggers on a connector
[**ThirdpartyConnectorUpdate**](ThirdpartyConnectorAPI.md#ThirdpartyConnectorUpdate) | **Put** /api/v1/thirdparty/connectors | Update an existing third party connector



## ThirdpartyConnectorAdd

> ThirdPartyConnectorResponses ThirdpartyConnectorAdd(ctx).ThirdPartyConnectors(thirdPartyConnectors).Execute()

Register a new third party connector



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
	thirdPartyConnectors := openapiclient.ThirdPartyConnectors{AWSConnector: openapiclient.NewAWSConnector("Type_example", "Name_example", "5 seconds", "Region_example")} // ThirdPartyConnectors | Third party connector to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorAdd(context.Background()).ThirdPartyConnectors(thirdPartyConnectors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThirdpartyConnectorAdd`: ThirdPartyConnectorResponses
	fmt.Fprintf(os.Stdout, "Response from `ThirdpartyConnectorAPI.ThirdpartyConnectorAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyConnectors** | [**ThirdPartyConnectors**](ThirdPartyConnectors.md) | Third party connector to register | 

### Return type

[**ThirdPartyConnectorResponses**](ThirdPartyConnectorResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThirdpartyConnectorDelete

> ThirdpartyConnectorDelete(ctx, name).Execute()

Delete an existing third party connector



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorDeleteRequest struct via the builder pattern


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


## ThirdpartyConnectorGet

> ThirdpartyConnectorList200ResponseInner ThirdpartyConnectorGet(ctx, name).Execute()

Retrieve an existing third party connector



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThirdpartyConnectorGet`: ThirdpartyConnectorList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ThirdpartyConnectorAPI.ThirdpartyConnectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdpartyConnectorList200ResponseInner**](ThirdpartyConnectorList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThirdpartyConnectorList

> []ThirdpartyConnectorList200ResponseInner ThirdpartyConnectorList(ctx).Type_(type_).Module(module).Execute()

List the existing third party connector(s)



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
	type_ := "type__example" // string |  (optional)
	module := "module_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorList(context.Background()).Type_(type_).Module(module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThirdpartyConnectorList`: []ThirdpartyConnectorList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ThirdpartyConnectorAPI.ThirdpartyConnectorList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **module** | **string** |  | 

### Return type

[**[]ThirdpartyConnectorList200ResponseInner**](ThirdpartyConnectorList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThirdpartyConnectorRetry

> ThirdpartyConnectorRetry(ctx, name).Execute()

Retry failed triggers on a connector



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorRetry(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorRetryRequest struct via the builder pattern


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


## ThirdpartyConnectorUpdate

> ThirdPartyConnectorResponses ThirdpartyConnectorUpdate(ctx).ThirdPartyConnectors(thirdPartyConnectors).Execute()

Update an existing third party connector



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
	thirdPartyConnectors := openapiclient.ThirdPartyConnectors{AWSConnector: openapiclient.NewAWSConnector("Type_example", "Name_example", "5 seconds", "Region_example")} // ThirdPartyConnectors | Third party connector to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdpartyConnectorAPI.ThirdpartyConnectorUpdate(context.Background()).ThirdPartyConnectors(thirdPartyConnectors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdpartyConnectorAPI.ThirdpartyConnectorUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThirdpartyConnectorUpdate`: ThirdPartyConnectorResponses
	fmt.Fprintf(os.Stdout, "Response from `ThirdpartyConnectorAPI.ThirdpartyConnectorUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThirdpartyConnectorUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyConnectors** | [**ThirdPartyConnectors**](ThirdPartyConnectors.md) | Third party connector to update | 

### Return type

[**ThirdPartyConnectorResponses**](ThirdPartyConnectorResponses.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

