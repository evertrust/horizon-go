# \AcmeAccountAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcmeAccountDelete**](AcmeAccountAPI.md#AcmeAccountDelete) | **Delete** /api/v1/acme/accounts/{accountId} | Delete an ACME account by ID
[**AcmeAccountGet**](AcmeAccountAPI.md#AcmeAccountGet) | **Get** /api/v1/acme/accounts/{accountId} | Get an ACME account by ID
[**AcmeAccountSearch**](AcmeAccountAPI.md#AcmeAccountSearch) | **Post** /api/v1/acme/accounts/search | Search ACME accounts
[**AcmeAccountUpdateStatus**](AcmeAccountAPI.md#AcmeAccountUpdateStatus) | **Post** /api/v1/acme/accounts/{accountId}/status | Update ACME account status



## AcmeAccountDelete

> AcmeAccountDelete(ctx, accountId).Execute()

Delete an ACME account by ID



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
	accountId := "accountId_example" // string | The ID of the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AcmeAccountAPI.AcmeAccountDelete(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeAccountAPI.AcmeAccountDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeAccountDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeAccountGet

> Account AcmeAccountGet(ctx, accountId).Execute()

Get an ACME account by ID



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
	accountId := "accountId_example" // string | The ID of the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeAccountAPI.AcmeAccountGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeAccountAPI.AcmeAccountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeAccountGet`: Account
	fmt.Fprintf(os.Stdout, "Response from `AcmeAccountAPI.AcmeAccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeAccountSearch

> AccountSearchResults AcmeAccountSearch(ctx).AccountSearch(accountSearch).Execute()

Search ACME accounts



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
	accountSearch := *openapiclient.NewAccountSearch() // AccountSearch | Search parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeAccountAPI.AcmeAccountSearch(context.Background()).AccountSearch(accountSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeAccountAPI.AcmeAccountSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeAccountSearch`: AccountSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AcmeAccountAPI.AcmeAccountSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcmeAccountSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountSearch** | [**AccountSearch**](AccountSearch.md) | Search parameters | 

### Return type

[**AccountSearchResults**](AccountSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeAccountUpdateStatus

> Account AcmeAccountUpdateStatus(ctx, accountId).UpdateAccountRequest(updateAccountRequest).Execute()

Update ACME account status



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
	accountId := "accountId_example" // string | The ID of the account to update
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest(openapiclient.AccountStatus("valid")) // UpdateAccountRequest | Account compromise parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeAccountAPI.AcmeAccountUpdateStatus(context.Background(), accountId).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeAccountAPI.AcmeAccountUpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeAccountUpdateStatus`: Account
	fmt.Fprintf(os.Stdout, "Response from `AcmeAccountAPI.AcmeAccountUpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeAccountUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) | Account compromise parameters | 

### Return type

[**Account**](Account.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

