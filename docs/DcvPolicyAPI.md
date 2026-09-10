# \DcvPolicyAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DcvPolicyAdd**](DcvPolicyAPI.md#DcvPolicyAdd) | **Post** /api/v1/dcv/policies | Register a new DCV policy
[**DcvPolicyDelete**](DcvPolicyAPI.md#DcvPolicyDelete) | **Delete** /api/v1/dcv/policies/{name} | Delete an existing DCV policy
[**DcvPolicyGet**](DcvPolicyAPI.md#DcvPolicyGet) | **Get** /api/v1/dcv/policies/{name} | Retrieve an existing DCV policy
[**DcvPolicyList**](DcvPolicyAPI.md#DcvPolicyList) | **Get** /api/v1/dcv/policies | List the existing DCV policies
[**DcvPolicyUpdate**](DcvPolicyAPI.md#DcvPolicyUpdate) | **Put** /api/v1/dcv/policies | Update an existing DCV policy



## DcvPolicyAdd

> DCVPolicyResponse DcvPolicyAdd(ctx).DCVPolicy(dCVPolicy).Execute()

Register a new DCV policy



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
	dCVPolicy := *openapiclient.NewDCVPolicy("ExecutionTimeout_example", "Name_example", "Provider_example", "Provisioner_example", "RetryDelay_example") // DCVPolicy | DCV policy to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvPolicyAPI.DcvPolicyAdd(context.Background()).DCVPolicy(dCVPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvPolicyAPI.DcvPolicyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvPolicyAdd`: DCVPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvPolicyAPI.DcvPolicyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvPolicyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dCVPolicy** | [**DCVPolicy**](DCVPolicy.md) | DCV policy to register | 

### Return type

[**DCVPolicyResponse**](DCVPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvPolicyDelete

> DcvPolicyDelete(ctx, name).Execute()

Delete an existing DCV policy



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
	r, err := apiClient.DcvPolicyAPI.DcvPolicyDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvPolicyAPI.DcvPolicyDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDcvPolicyDeleteRequest struct via the builder pattern


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


## DcvPolicyGet

> DCVPolicyResponse DcvPolicyGet(ctx, name).Execute()

Retrieve an existing DCV policy



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
	resp, r, err := apiClient.DcvPolicyAPI.DcvPolicyGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvPolicyAPI.DcvPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvPolicyGet`: DCVPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvPolicyAPI.DcvPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDcvPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DCVPolicyResponse**](DCVPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvPolicyList

> []DCVPolicyResponse DcvPolicyList(ctx).Execute()

List the existing DCV policies



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
	resp, r, err := apiClient.DcvPolicyAPI.DcvPolicyList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvPolicyAPI.DcvPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvPolicyList`: []DCVPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvPolicyAPI.DcvPolicyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDcvPolicyListRequest struct via the builder pattern


### Return type

[**[]DCVPolicyResponse**](DCVPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DcvPolicyUpdate

> DCVPolicyResponse DcvPolicyUpdate(ctx).DCVPolicy(dCVPolicy).Execute()

Update an existing DCV policy



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
	dCVPolicy := *openapiclient.NewDCVPolicy("ExecutionTimeout_example", "Name_example", "Provider_example", "Provisioner_example", "RetryDelay_example") // DCVPolicy | DCV policy to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DcvPolicyAPI.DcvPolicyUpdate(context.Background()).DCVPolicy(dCVPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DcvPolicyAPI.DcvPolicyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DcvPolicyUpdate`: DCVPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DcvPolicyAPI.DcvPolicyUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDcvPolicyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dCVPolicy** | [**DCVPolicy**](DCVPolicy.md) | DCV policy to update | 

### Return type

[**DCVPolicyResponse**](DCVPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

