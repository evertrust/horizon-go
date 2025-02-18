# \PkiConnectorAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PkiConnectorAdd**](PkiConnectorAPI.md#PkiConnectorAdd) | **Post** /api/v1/pki/connectors | Register a new PKI connector
[**PkiConnectorDelete**](PkiConnectorAPI.md#PkiConnectorDelete) | **Delete** /api/v1/pki/connectors/{name} | Delete an existing PKI connector
[**PkiConnectorGet**](PkiConnectorAPI.md#PkiConnectorGet) | **Get** /api/v1/pki/connectors/{name} | Retrieve an existing PKI connector
[**PkiConnectorList**](PkiConnectorAPI.md#PkiConnectorList) | **Get** /api/v1/pki/connectors | List the existing PKI connector(s)
[**PkiConnectorUpdate**](PkiConnectorAPI.md#PkiConnectorUpdate) | **Put** /api/v1/pki/connectors | Update an existing PKI connector



## PkiConnectorAdd

> PkiConnectorList200ResponseInner PkiConnectorAdd(ctx).PkiConnectorUpdateRequest(pkiConnectorUpdateRequest).Execute()

Register a new PKI connector



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
	pkiConnectorUpdateRequest := openapiclient.pki_connector_update_request{ADCSConnector: openapiclient.NewADCSConnector("Name_example", "Type_example", "EndPoint_example", "Profile_example", "myPasswordCredentials", "myCertificateCredentials")} // PkiConnectorUpdateRequest | PKI connector to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiConnectorAPI.PkiConnectorAdd(context.Background()).PkiConnectorUpdateRequest(pkiConnectorUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiConnectorAPI.PkiConnectorAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiConnectorAdd`: PkiConnectorList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PkiConnectorAPI.PkiConnectorAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkiConnectorAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConnectorUpdateRequest** | [**PkiConnectorUpdateRequest**](PkiConnectorUpdateRequest.md) | PKI connector to register | 

### Return type

[**PkiConnectorList200ResponseInner**](PkiConnectorList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiConnectorDelete

> PkiConnectorDelete(ctx, name).Execute()

Delete an existing PKI connector



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PkiConnectorAPI.PkiConnectorDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiConnectorAPI.PkiConnectorDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPkiConnectorDeleteRequest struct via the builder pattern


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


## PkiConnectorGet

> PkiConnectorGet200Response PkiConnectorGet(ctx, name).Execute()

Retrieve an existing PKI connector



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiConnectorAPI.PkiConnectorGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiConnectorAPI.PkiConnectorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiConnectorGet`: PkiConnectorGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PkiConnectorAPI.PkiConnectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPkiConnectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PkiConnectorGet200Response**](PkiConnectorGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiConnectorList

> []PkiConnectorList200ResponseInner PkiConnectorList(ctx).Execute()

List the existing PKI connector(s)



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
	resp, r, err := apiClient.PkiConnectorAPI.PkiConnectorList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiConnectorAPI.PkiConnectorList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiConnectorList`: []PkiConnectorList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PkiConnectorAPI.PkiConnectorList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPkiConnectorListRequest struct via the builder pattern


### Return type

[**[]PkiConnectorList200ResponseInner**](PkiConnectorList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PkiConnectorUpdate

> PkiConnectorList200ResponseInner PkiConnectorUpdate(ctx).PkiConnectorUpdateRequest(pkiConnectorUpdateRequest).Execute()

Update an existing PKI connector



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
	pkiConnectorUpdateRequest := openapiclient.pki_connector_update_request{ADCSConnector: openapiclient.NewADCSConnector("Name_example", "Type_example", "EndPoint_example", "Profile_example", "myPasswordCredentials", "myCertificateCredentials")} // PkiConnectorUpdateRequest | PKI connector to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PkiConnectorAPI.PkiConnectorUpdate(context.Background()).PkiConnectorUpdateRequest(pkiConnectorUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PkiConnectorAPI.PkiConnectorUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PkiConnectorUpdate`: PkiConnectorList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PkiConnectorAPI.PkiConnectorUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPkiConnectorUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkiConnectorUpdateRequest** | [**PkiConnectorUpdateRequest**](PkiConnectorUpdateRequest.md) | PKI connector to update | 

### Return type

[**PkiConnectorList200ResponseInner**](PkiConnectorList200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

