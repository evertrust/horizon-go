# \Rfc5280API

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Rfc5280Detect**](Rfc5280API.md#Rfc5280Detect) | **Post** /api/v1/rfc5280/detect | Detect and parse a RFC 5280 related file
[**Rfc5280Pkcs10File**](Rfc5280API.md#Rfc5280Pkcs10File) | **Post** /api/v1/rfc5280/pkcs10 | Decode a pkcs#10 (file)
[**Rfc5280Pkcs10Pem**](Rfc5280API.md#Rfc5280Pkcs10Pem) | **Get** /api/v1/rfc5280/pkcs10/{pem} | Decode a pkcs#10 (url encoded)
[**Rfc5280Pkcs12File**](Rfc5280API.md#Rfc5280Pkcs12File) | **Post** /api/v1/rfc5280/pkcs12 | Extract the certificate and associated private key from a pkcs#12 (file)
[**Rfc5280TcFile**](Rfc5280API.md#Rfc5280TcFile) | **Post** /api/v1/rfc5280/tc | Retrieve the Trust chain from a x509 certificate (file)
[**Rfc5280TcPem**](Rfc5280API.md#Rfc5280TcPem) | **Get** /api/v1/rfc5280/tc/{pem} | Retrieve the Trust chain from a x509 certificate (url encoded)
[**Rfc5280X509File**](Rfc5280API.md#Rfc5280X509File) | **Post** /api/v1/rfc5280/x509 | Decode a x509 certificate (file)
[**Rfc5280X509Pem**](Rfc5280API.md#Rfc5280X509Pem) | **Get** /api/v1/rfc5280/x509/{pem} | Decode a x509 certificate (url encoded)



## Rfc5280Detect

> Rfc5280Detect200Response Rfc5280Detect(ctx).File(file).Execute()

Detect and parse a RFC 5280 related file



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
	file := openapiclient.grading_policy_explain_file_request_x509{ArrayOfByte: new([]byte)} // GradingPolicyExplainFileRequestX509 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280Detect(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280Detect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280Detect`: Rfc5280Detect200Response
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280Detect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280DetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [**GradingPolicyExplainFileRequestX509**](GradingPolicyExplainFileRequestX509.md) |  | 

### Return type

[**Rfc5280Detect200Response**](Rfc5280Detect200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280Pkcs10File

> CFCertificationRequestResponse Rfc5280Pkcs10File(ctx).Pkcs10(pkcs10).Execute()

Decode a pkcs#10 (file)



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
	pkcs10 := openapiclient.grading_policy_explain_file_request_x509{ArrayOfByte: new([]byte)} // GradingPolicyExplainFileRequestX509 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280Pkcs10File(context.Background()).Pkcs10(pkcs10).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280Pkcs10File``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280Pkcs10File`: CFCertificationRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280Pkcs10File`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280Pkcs10FileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkcs10** | [**GradingPolicyExplainFileRequestX509**](GradingPolicyExplainFileRequestX509.md) |  | 

### Return type

[**CFCertificationRequestResponse**](CFCertificationRequestResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280Pkcs10Pem

> CFCertificationRequestResponse Rfc5280Pkcs10Pem(ctx, pem).Execute()

Decode a pkcs#10 (url encoded)



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
	pem := "pem_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280Pkcs10Pem(context.Background(), pem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280Pkcs10Pem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280Pkcs10Pem`: CFCertificationRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280Pkcs10Pem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pem** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280Pkcs10PemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CFCertificationRequestResponse**](CFCertificationRequestResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280Pkcs12File

> Rfc5280Pkcs12ContentResponse Rfc5280Pkcs12File(ctx).Pkcs12(pkcs12).Execute()

Extract the certificate and associated private key from a pkcs#12 (file)



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
	pkcs12 := openapiclient.grading_policy_explain_file_request_x509{ArrayOfByte: new([]byte)} // GradingPolicyExplainFileRequestX509 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280Pkcs12File(context.Background()).Pkcs12(pkcs12).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280Pkcs12File``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280Pkcs12File`: Rfc5280Pkcs12ContentResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280Pkcs12File`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280Pkcs12FileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkcs12** | [**GradingPolicyExplainFileRequestX509**](GradingPolicyExplainFileRequestX509.md) |  | 

### Return type

[**Rfc5280Pkcs12ContentResponse**](Rfc5280Pkcs12ContentResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280TcFile

> []CFCertificateResponse Rfc5280TcFile(ctx).Order(order).X509(x509).Execute()

Retrieve the Trust chain from a x509 certificate (file)



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
	order := "order_example" // string |  (optional)
	x509 := openapiclient.grading_policy_explain_file_request_x509{ArrayOfByte: new([]byte)} // GradingPolicyExplainFileRequestX509 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280TcFile(context.Background()).Order(order).X509(x509).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280TcFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280TcFile`: []CFCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280TcFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280TcFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** |  | 
 **x509** | [**GradingPolicyExplainFileRequestX509**](GradingPolicyExplainFileRequestX509.md) |  | 

### Return type

[**[]CFCertificateResponse**](CFCertificateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/x-pem-file, application/x-pkcs7-certificates, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280TcPem

> []CFCertificateResponse Rfc5280TcPem(ctx, pem).Order(order).Execute()

Retrieve the Trust chain from a x509 certificate (url encoded)



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
	pem := "pem_example" // string | 
	order := "order_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280TcPem(context.Background(), pem).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280TcPem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280TcPem`: []CFCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280TcPem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pem** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280TcPemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** |  | 

### Return type

[**[]CFCertificateResponse**](CFCertificateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-pem-file, application/x-pkcs7-certificates, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280X509File

> CFCertificateResponse Rfc5280X509File(ctx).X509(x509).Execute()

Decode a x509 certificate (file)



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
	x509 := openapiclient.grading_policy_explain_file_request_x509{ArrayOfByte: new([]byte)} // GradingPolicyExplainFileRequestX509 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280X509File(context.Background()).X509(x509).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280X509File``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280X509File`: CFCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280X509File`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280X509FileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x509** | [**GradingPolicyExplainFileRequestX509**](GradingPolicyExplainFileRequestX509.md) |  | 

### Return type

[**CFCertificateResponse**](CFCertificateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/pkix-cert, application/x-pem-file, application/x-pkcs7-certificates, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rfc5280X509Pem

> CFCertificateResponse Rfc5280X509Pem(ctx, pem).Execute()

Decode a x509 certificate (url encoded)



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
	pem := "pem_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Rfc5280API.Rfc5280X509Pem(context.Background(), pem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Rfc5280API.Rfc5280X509Pem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rfc5280X509Pem`: CFCertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `Rfc5280API.Rfc5280X509Pem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pem** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRfc5280X509PemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CFCertificateResponse**](CFCertificateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth), [apiId](../README.md#apiId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pkix-cert, application/x-pem-file, application/x-pkcs7-certificates, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

