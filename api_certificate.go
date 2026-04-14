/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/evertrust/horizon-go/v2/models"
	"github.com/evertrust/horizon-go/v2/utils"
)

// CertificateAPIService CertificateAPI service
type CertificateAPIService service

type CertificateAPICertificateAggregateRequest struct {
	ctx                       context.Context
	ApiService                *CertificateAPIService
	certificateAggregateQuery *models.CertificateAggregateQuery
	enableAnalytics           *bool
}

// The certificate aggregation query
func (r CertificateAPICertificateAggregateRequest) CertificateAggregateQuery(certificateAggregateQuery models.CertificateAggregateQuery) CertificateAPICertificateAggregateRequest {
	r.certificateAggregateQuery = &certificateAggregateQuery
	return r
}

// Use the analytics database if enabled. &#x60;true&#x60; if not specified.
func (r CertificateAPICertificateAggregateRequest) EnableAnalytics(enableAnalytics bool) CertificateAPICertificateAggregateRequest {
	r.enableAnalytics = &enableAnalytics
	return r
}

func (r CertificateAPICertificateAggregateRequest) Execute() (*models.CertificateAggregateResultResponse, *http.Response, error) {
	return r.ApiService.CertificateAggregateExecute(r)
}

/*
CertificateAggregate Certificate aggregation

Send a certificate aggregation query and return the aggregation result

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CertificateAPICertificateAggregateRequest
*/
func (a *CertificateAPIService) CertificateAggregate(ctx context.Context) CertificateAPICertificateAggregateRequest {
	return CertificateAPICertificateAggregateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateAggregateResultResponse
func (a *CertificateAPIService) CertificateAggregateExecute(r CertificateAPICertificateAggregateRequest) (*models.CertificateAggregateResultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateAggregateResultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateAggregate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certificateAggregateQuery == nil {
		return localVarReturnValue, nil, utils.ReportError("certificateAggregateQuery is required and must be specified")
	}

	if r.enableAnalytics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableAnalytics", r.enableAnalytics, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.certificateAggregateQuery
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateCsvRequest struct {
	ctx                    context.Context
	ApiService             *CertificateAPIService
	certificateSearchQuery *models.CertificateSearchQuery
	enableAnalytics        *bool
}

// The certificate search query
func (r CertificateAPICertificateCsvRequest) CertificateSearchQuery(certificateSearchQuery models.CertificateSearchQuery) CertificateAPICertificateCsvRequest {
	r.certificateSearchQuery = &certificateSearchQuery
	return r
}

// Use the analytics database if enabled. &#x60;true&#x60; if not specified.
func (r CertificateAPICertificateCsvRequest) EnableAnalytics(enableAnalytics bool) CertificateAPICertificateCsvRequest {
	r.enableAnalytics = &enableAnalytics
	return r
}

func (r CertificateAPICertificateCsvRequest) Execute() (*http.Response, error) {
	return r.ApiService.CertificateCsvExecute(r)
}

/*
CertificateCsv Export certificates

Send a certificate search query (in HCQL format) and return the certificate search results in CSV format

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CertificateAPICertificateCsvRequest
*/
func (a *CertificateAPIService) CertificateCsv(ctx context.Context) CertificateAPICertificateCsvRequest {
	return CertificateAPICertificateCsvRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CertificateAPIService) CertificateCsvExecute(r CertificateAPICertificateCsvRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateCsv")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/csv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certificateSearchQuery == nil {
		return nil, utils.ReportError("certificateSearchQuery is required and must be specified")
	}

	if r.enableAnalytics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableAnalytics", r.enableAnalytics, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/csv", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.certificateSearchQuery
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type CertificateAPICertificateDictionaryRequest struct {
	ctx        context.Context
	ApiService *CertificateAPIService
}

func (r CertificateAPICertificateDictionaryRequest) Execute() (*models.CertificateSearchDictionaryResponse, *http.Response, error) {
	return r.ApiService.CertificateDictionaryExecute(r)
}

/*
CertificateDictionary Retrieve the certificate search dictionary

Return the certificate search dictionary. The dictionary is computed based on the principal and includes:

  - The list of certificate profiles on which the principal is authorized to search on;

  - The list of discovery campaigns the principal is authorized to search on;

  - The list of labels the principal is authorized to search on;

  - The list of modules available on the Horizon instance;

  - The list of available teams on the Horizon instance;

  - The list of available grading policies on the Horizon instance;

  - The list of available metadata on Horizon.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return CertificateAPICertificateDictionaryRequest
*/
func (a *CertificateAPIService) CertificateDictionary(ctx context.Context) CertificateAPICertificateDictionaryRequest {
	return CertificateAPICertificateDictionaryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateSearchDictionaryResponse
func (a *CertificateAPIService) CertificateDictionaryExecute(r CertificateAPICertificateDictionaryRequest) (*models.CertificateSearchDictionaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateSearchDictionaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateDictionary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/search/dictionary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateFindRequest struct {
	ctx                    context.Context
	ApiService             *CertificateAPIService
	certificateFindRequest *models.CertificateFindRequest
}

func (r CertificateAPICertificateFindRequest) CertificateFindRequest(certificateFindRequest models.CertificateFindRequest) CertificateAPICertificateFindRequest {
	r.certificateFindRequest = &certificateFindRequest
	return r
}

func (r CertificateAPICertificateFindRequest) Execute() (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	return r.ApiService.CertificateFindExecute(r)
}

/*
CertificateFind Find a certificate

Find a certificate by its Id or PEM

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CertificateAPICertificateFindRequest
*/
func (a *CertificateAPIService) CertificateFind(ctx context.Context) CertificateAPICertificateFindRequest {
	return CertificateAPICertificateFindRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateWithPermissionsResponse
func (a *CertificateAPIService) CertificateFindExecute(r CertificateAPICertificateFindRequest) (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateWithPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateFind")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/find"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certificateFindRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("certificateFindRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.certificateFindRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateGetIdRequest struct {
	ctx        context.Context
	ApiService *CertificateAPIService
	id         string
}

func (r CertificateAPICertificateGetIdRequest) Execute() (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	return r.ApiService.CertificateGetIdExecute(r)
}

/*
CertificateGetId Retrieve a certificate

Retrieves a specific certificate based on its ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the certificate to retrieve
	@return CertificateAPICertificateGetIdRequest
*/
func (a *CertificateAPIService) CertificateGetId(ctx context.Context, id string) CertificateAPICertificateGetIdRequest {
	return CertificateAPICertificateGetIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CertificateWithPermissionsResponse
func (a *CertificateAPIService) CertificateGetIdExecute(r CertificateAPICertificateGetIdRequest) (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateWithPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateGetId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateGetPemRequest struct {
	ctx        context.Context
	ApiService *CertificateAPIService
	pem        string
}

func (r CertificateAPICertificateGetPemRequest) Execute() (*models.CertificateResponse, *http.Response, error) {
	return r.ApiService.CertificateGetPemExecute(r)
}

/*
CertificateGetPem Retrieve a certificate by PEM

Retrieve a specific certificate based on its PEM encoded value. This operation is deprecated and usage of the [find operation](#tag/certificate/operation/certificate.find) is recommended.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pem The URL encoded PEM encoded value of the certificate
	@return CertificateAPICertificateGetPemRequest

Deprecated
*/
func (a *CertificateAPIService) CertificateGetPem(ctx context.Context, pem string) CertificateAPICertificateGetPemRequest {
	return CertificateAPICertificateGetPemRequest{
		ApiService: a,
		ctx:        ctx,
		pem:        pem,
	}
}

// Execute executes the request
//
//	@return CertificateResponse
//
// Deprecated
func (a *CertificateAPIService) CertificateGetPemExecute(r CertificateAPICertificateGetPemRequest) (*models.CertificateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateGetPem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/{pem}"
	localVarPath = strings.Replace(localVarPath, "{"+"pem"+"}", url.PathEscape(parameterValueToString(r.pem, "pem")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateListRequest struct {
	ctx         context.Context
	ApiService  *CertificateAPIService
	requestBody *[]string
}

// The list of certificates IDs to fetch
func (r CertificateAPICertificateListRequest) RequestBody(requestBody []string) CertificateAPICertificateListRequest {
	r.requestBody = &requestBody
	return r
}

func (r CertificateAPICertificateListRequest) Execute() ([]models.CertificateWithPermissionsResponse, *http.Response, error) {
	return r.ApiService.CertificateListExecute(r)
}

/*
CertificateList List certificates

List certificate(s) matching any ID sent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CertificateAPICertificateListRequest
*/
func (a *CertificateAPIService) CertificateList(ctx context.Context) CertificateAPICertificateListRequest {
	return CertificateAPICertificateListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CertificateWithPermissionsResponse
func (a *CertificateAPIService) CertificateListExecute(r CertificateAPICertificateListRequest) ([]models.CertificateWithPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.CertificateWithPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, utils.ReportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateRunRequest struct {
	ctx         context.Context
	ApiService  *CertificateAPIService
	id          string
	triggerName string
	event       string
}

func (r CertificateAPICertificateRunRequest) Execute() (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	return r.ApiService.CertificateRunExecute(r)
}

/*
CertificateRun Run a certificate trigger

When a trigger fails, the user might have the ability to run the trigger manually again. This is only possible when `retryable` is set to true in the `triggerResult`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the certificate
	@param triggerName The name of the trigger
	@param event
	@return CertificateAPICertificateRunRequest
*/
func (a *CertificateAPIService) CertificateRun(ctx context.Context, id string, triggerName string, event string) CertificateAPICertificateRunRequest {
	return CertificateAPICertificateRunRequest{
		ApiService:  a,
		ctx:         ctx,
		id:          id,
		triggerName: triggerName,
		event:       event,
	}
}

// Execute executes the request
//
//	@return CertificateWithPermissionsResponse
func (a *CertificateAPIService) CertificateRunExecute(r CertificateAPICertificateRunRequest) (*models.CertificateWithPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateWithPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/run/{id}/{triggerName}/{event}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"triggerName"+"}", url.PathEscape(parameterValueToString(r.triggerName, "triggerName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event"+"}", url.PathEscape(parameterValueToString(r.event, "event")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CertificateAPICertificateSearchRequest struct {
	ctx                    context.Context
	ApiService             *CertificateAPIService
	certificateSearchQuery *models.CertificateSearchQuery
	enableAnalytics        *bool
}

func (r CertificateAPICertificateSearchRequest) CertificateSearchQuery(certificateSearchQuery models.CertificateSearchQuery) CertificateAPICertificateSearchRequest {
	r.certificateSearchQuery = &certificateSearchQuery
	return r
}

// Use the analytics database if enabled. &#x60;true&#x60; if not specified.
func (r CertificateAPICertificateSearchRequest) EnableAnalytics(enableAnalytics bool) CertificateAPICertificateSearchRequest {
	r.enableAnalytics = &enableAnalytics
	return r
}

func (r CertificateAPICertificateSearchRequest) Execute() (*models.CertificateSearchResultsResponse, *http.Response, error) {
	return r.ApiService.CertificateSearchExecute(r)
}

/*
CertificateSearch Search certificates

Send a certificate search query (in HCQL format) and return the certificate search results

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CertificateAPICertificateSearchRequest
*/
func (a *CertificateAPIService) CertificateSearch(ctx context.Context) CertificateAPICertificateSearchRequest {
	return CertificateAPICertificateSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateSearchResultsResponse
func (a *CertificateAPIService) CertificateSearchExecute(r CertificateAPICertificateSearchRequest) (*models.CertificateSearchResultsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.CertificateSearchResultsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAPIService.CertificateSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/certificates/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certificateSearchQuery == nil {
		return localVarReturnValue, nil, utils.ReportError("certificateSearchQuery is required and must be specified")
	}

	if r.enableAnalytics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableAnalytics", r.enableAnalytics, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.certificateSearchQuery
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
			}
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiId"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-ID"] = key
			}
		}
	}

	if r.ctx != nil {

		// JWT POP
		if jwtPopCert, jwtPopSigner, ok := utils.GetJWTPoP(r.ctx); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			// send without the Nonce to get replay nonce
			jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
			// send the request a first time but without any data to get the replay nonce
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
			if err != nil {
				return localVarReturnValue, nil, err
			}
			localVarHTTPResponse, err := a.client.callAPI(req)
			if err != nil || localVarHTTPResponse == nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// read the response to get the replay nonce
			localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
			localVarHTTPResponse.Body.Close()
			localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, err
			}
			// from the request read the replay nonce from the response header and resend the request
			nonce := localVarHTTPResponse.Header.Get("Replay-Nonce")
			if nonce == "" {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "no replay nonce received in response"}
			}
			jwt, err = utils.CreateJWT(*jwtPopCert, jwtPopSigner, nonce)
			if err != nil {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
			}
			localVarHeaderParams["X-JWT-CERT-POP"] = jwt
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
