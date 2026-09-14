# \AcmeOrderAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcmeOrderGet**](AcmeOrderAPI.md#AcmeOrderGet) | **Get** /api/v1/acme/orders/{orderId} | Get an ACME order by ID
[**AcmeOrderListByAccount**](AcmeOrderAPI.md#AcmeOrderListByAccount) | **Post** /api/v1/acme/orders/account/{accountId} | List orders by account ID



## AcmeOrderGet

> Order AcmeOrderGet(ctx, orderId).Execute()

Get an ACME order by ID



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
	orderId := "orderId_example" // string | The ID of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeOrderAPI.AcmeOrderGet(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeOrderAPI.AcmeOrderGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeOrderGet`: Order
	fmt.Fprintf(os.Stdout, "Response from `AcmeOrderAPI.AcmeOrderGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeOrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Order**](Order.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcmeOrderListByAccount

> OrderSearchResults AcmeOrderListByAccount(ctx, accountId).OrderAPISearch(orderAPISearch).Execute()

List orders by account ID



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
	orderAPISearch := *openapiclient.NewOrderAPISearch() // OrderAPISearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcmeOrderAPI.AcmeOrderListByAccount(context.Background(), accountId).OrderAPISearch(orderAPISearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcmeOrderAPI.AcmeOrderListByAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcmeOrderListByAccount`: OrderSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AcmeOrderAPI.AcmeOrderListByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcmeOrderListByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderAPISearch** | [**OrderAPISearch**](OrderAPISearch.md) |  | 

### Return type

[**OrderSearchResults**](OrderSearchResults.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

