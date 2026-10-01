# \ChallengeAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChallengeSubmit**](ChallengeAPI.md#ChallengeSubmit) | **Post** /api/v1/challenge/submit | Consume a challenge



## ChallengeSubmit

> WebRAChallengeSubmitResponse ChallengeSubmit(ctx).WebRAChallengeSubmitRequest(webRAChallengeSubmitRequest).Execute()

Consume a challenge



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
	webRAChallengeSubmitRequest := *openapiclient.NewWebRAChallengeSubmitRequest("8ChpvJvNSTaMxxlP", "webra_centralized", *openapiclient.NewWebRAChallengeSubmitRequestTemplate()) // WebRAChallengeSubmitRequest | The challenge to consume

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengeAPI.ChallengeSubmit(context.Background()).WebRAChallengeSubmitRequest(webRAChallengeSubmitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengeAPI.ChallengeSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeSubmit`: WebRAChallengeSubmitResponse
	fmt.Fprintf(os.Stdout, "Response from `ChallengeAPI.ChallengeSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChallengeSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webRAChallengeSubmitRequest** | [**WebRAChallengeSubmitRequest**](WebRAChallengeSubmitRequest.md) | The challenge to consume | 

### Return type

[**WebRAChallengeSubmitResponse**](WebRAChallengeSubmitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

