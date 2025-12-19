# \SecurityTeamAPI

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityTeamAdd**](SecurityTeamAPI.md#SecurityTeamAdd) | **Post** /api/v1/security/teams | Create a new team
[**SecurityTeamDelete**](SecurityTeamAPI.md#SecurityTeamDelete) | **Delete** /api/v1/security/teams/{name} | Delete a team
[**SecurityTeamGet**](SecurityTeamAPI.md#SecurityTeamGet) | **Get** /api/v1/security/teams/{name} | Retrieve a team
[**SecurityTeamList**](SecurityTeamAPI.md#SecurityTeamList) | **Get** /api/v1/security/teams | List the team(s)
[**SecurityTeamMembersAdd**](SecurityTeamAPI.md#SecurityTeamMembersAdd) | **Post** /api/v1/security/teams/{name}/members | Add members to a team
[**SecurityTeamMembersGet**](SecurityTeamAPI.md#SecurityTeamMembersGet) | **Get** /api/v1/security/teams/{name}/members | Retrieve a team&#39;s members
[**SecurityTeamMembersRemove**](SecurityTeamAPI.md#SecurityTeamMembersRemove) | **Delete** /api/v1/security/teams/{name}/members | Remove members from a team
[**SecurityTeamSwitch**](SecurityTeamAPI.md#SecurityTeamSwitch) | **Patch** /api/v1/security/teams/{previousTeam}/{newTeam} | Transfer objects from a team to another
[**SecurityTeamUpdate**](SecurityTeamAPI.md#SecurityTeamUpdate) | **Put** /api/v1/security/teams | Update a team



## SecurityTeamAdd

> TeamResponse SecurityTeamAdd(ctx).Team(team).Execute()

Create a new team



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
	team := *openapiclient.NewTeam("PKIOps") // Team | Team to register

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTeamAPI.SecurityTeamAdd(context.Background()).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTeamAdd`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTeamAPI.SecurityTeamAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **team** | [**Team**](Team.md) | Team to register | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamDelete

> SecurityTeamDelete(ctx, name).Execute()

Delete a team



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
	name := "PKIOps" // string | The name of the team to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTeamAPI.SecurityTeamDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the team to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamDeleteRequest struct via the builder pattern


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


## SecurityTeamGet

> TeamResponse SecurityTeamGet(ctx, name).Execute()

Retrieve a team



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
	name := "PKIOps" // string | The name of the team to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTeamAPI.SecurityTeamGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTeamGet`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTeamAPI.SecurityTeamGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the team to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamList

> []TeamResponse SecurityTeamList(ctx).Execute()

List the team(s)



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
	resp, r, err := apiClient.SecurityTeamAPI.SecurityTeamList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTeamList`: []TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTeamAPI.SecurityTeamList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamListRequest struct via the builder pattern


### Return type

[**[]TeamResponse**](TeamResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamMembersAdd

> SecurityTeamMembersAdd(ctx, name).RequestBody(requestBody).Execute()

Add members to a team



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
	name := "SuperAdmin" // string | Team name
	requestBody := []string{"Property_example"} // []string | The members to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTeamAPI.SecurityTeamMembersAdd(context.Background(), name).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamMembersAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | The members to add | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamMembersGet

> []string SecurityTeamMembersGet(ctx, name).Execute()

Retrieve a team's members



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
	name := "SuperAdmin" // string | Team name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTeamAPI.SecurityTeamMembersGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTeamMembersGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `SecurityTeamAPI.SecurityTeamMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamMembersRemove

> SecurityTeamMembersRemove(ctx, name).RequestBody(requestBody).Execute()

Remove members from a team



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
	name := "SuperAdmin" // string | Team name
	requestBody := []string{"Property_example"} // []string | The members to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTeamAPI.SecurityTeamMembersRemove(context.Background(), name).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamMembersRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamMembersRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | The members to remove | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityTeamSwitch

> SecurityTeamSwitch(ctx, previousTeam, newTeam).Execute()

Transfer objects from a team to another



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
	previousTeam := "WinHorizon" // string | The name of the team to transfer objects from
	newTeam := "PKIOps" // string | The name of the team to transfer objects to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityTeamAPI.SecurityTeamSwitch(context.Background(), previousTeam, newTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**previousTeam** | **string** | The name of the team to transfer objects from | 
**newTeam** | **string** | The name of the team to transfer objects to | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamSwitchRequest struct via the builder pattern


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


## SecurityTeamUpdate

> TeamResponse SecurityTeamUpdate(ctx).SecurityTeamUpdateRequest(securityTeamUpdateRequest).Execute()

Update a team



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
	securityTeamUpdateRequest := *openapiclient.NewSecurityTeamUpdateRequest("PKIOps") // SecurityTeamUpdateRequest | The team to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityTeamAPI.SecurityTeamUpdate(context.Background()).SecurityTeamUpdateRequest(securityTeamUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityTeamAPI.SecurityTeamUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecurityTeamUpdate`: TeamResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityTeamAPI.SecurityTeamUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityTeamUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityTeamUpdateRequest** | [**SecurityTeamUpdateRequest**](SecurityTeamUpdateRequest.md) | The team to update | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

