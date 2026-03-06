/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"bytes"
	"context"
	"crypto"
	"crypto/x509"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/evertrust/horizon-go/v2/models"
	"github.com/evertrust/horizon-go/v2/utils"
)

// RequestAPIService RequestAPI service
type RequestAPIService service

type RequestAPIRequestAggregateRequest struct {
	ctx                   context.Context
	ApiService            *RequestAPIService
	requestAggregateQuery *models.RequestAggregateQuery
}

// The request aggregation query
func (r RequestAPIRequestAggregateRequest) RequestAggregateQuery(requestAggregateQuery models.RequestAggregateQuery) RequestAPIRequestAggregateRequest {
	r.requestAggregateQuery = &requestAggregateQuery
	return r
}

func (r RequestAPIRequestAggregateRequest) Execute() (*models.RequestAggregateResultResponse, *http.Response, error) {
	return r.ApiService.RequestAggregateExecute(r)
}

/*
RequestAggregate Request aggregation

Send a request aggregation query and return the aggregation result

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestAggregateRequest
*/
func (a *RequestAPIService) RequestAggregate(ctx context.Context) RequestAPIRequestAggregateRequest {
	return RequestAPIRequestAggregateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestAggregateResultResponse
func (a *RequestAPIService) RequestAggregateExecute(r RequestAPIRequestAggregateRequest) (*models.RequestAggregateResultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestAggregateResultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestAggregate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestAggregateQuery == nil {
		return localVarReturnValue, nil, utils.ReportError("requestAggregateQuery is required and must be specified")
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
	localVarPostBody = r.requestAggregateQuery
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestApproveRequest struct {
	ctx                   context.Context
	ApiService            *RequestAPIService
	requestApproveRequest *models.RequestApproveRequest
}

// The request to approve
func (r RequestAPIRequestApproveRequest) RequestApproveRequest(requestApproveRequest models.RequestApproveRequest) RequestAPIRequestApproveRequest {
	r.requestApproveRequest = &requestApproveRequest
	return r
}

func (r RequestAPIRequestApproveRequest) Execute() (*models.RequestApprove200Response, *http.Response, error) {
	return r.ApiService.RequestApproveExecute(r)
}

/*
RequestApprove Approve a request

Requester that do not have the privileges to directly enroll will see their requests in the pending state after submitting them. An approver can then approve the request, which will trigger the enrollment trough the configured PKI connector.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestApproveRequest
*/
func (a *RequestAPIService) RequestApprove(ctx context.Context) RequestAPIRequestApproveRequest {
	return RequestAPIRequestApproveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestApprove200Response
func (a *RequestAPIService) RequestApproveExecute(r RequestAPIRequestApproveRequest) (*models.RequestApprove200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestApprove200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestApprove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/approve"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestApproveRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("requestApproveRequest is required and must be specified")
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
	localVarPostBody = r.requestApproveRequest
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestCancelRequest struct {
	ctx                  context.Context
	ApiService           *RequestAPIService
	requestCancelRequest *models.RequestCancelRequest
}

// The Request to cancel
func (r RequestAPIRequestCancelRequest) RequestCancelRequest(requestCancelRequest models.RequestCancelRequest) RequestAPIRequestCancelRequest {
	r.requestCancelRequest = &requestCancelRequest
	return r
}

func (r RequestAPIRequestCancelRequest) Execute() (*models.RequestApprove200Response, *http.Response, error) {
	return r.ApiService.RequestCancelExecute(r)
}

/*
RequestCancel Cancel a request

Cancel an existing request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestCancelRequest
*/
func (a *RequestAPIService) RequestCancel(ctx context.Context) RequestAPIRequestCancelRequest {
	return RequestAPIRequestCancelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestApprove200Response
func (a *RequestAPIService) RequestCancelExecute(r RequestAPIRequestCancelRequest) (*models.RequestApprove200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestApprove200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestCancelRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("requestCancelRequest is required and must be specified")
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
	localVarPostBody = r.requestCancelRequest
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestCertificateProfileRequest struct {
	ctx        context.Context
	ApiService *RequestAPIService
	module     *string
	workflow   *string
}

func (r RequestAPIRequestCertificateProfileRequest) Module(module string) RequestAPIRequestCertificateProfileRequest {
	r.module = &module
	return r
}

func (r RequestAPIRequestCertificateProfileRequest) Workflow(workflow string) RequestAPIRequestCertificateProfileRequest {
	r.workflow = &workflow
	return r
}

func (r RequestAPIRequestCertificateProfileRequest) Execute() ([]models.RequestableCertificateProfileResponse, *http.Response, error) {
	return r.ApiService.RequestCertificateProfileExecute(r)
}

/*
RequestCertificateProfile List profiles

All requests on Horizon are linked to a profile, which defines a certificate template and a PKI connector which will sign the certificate.
Before submitting a request (such as an enrollement or revocation request), you must choose the profile on which you want to perform the operation.
This endpoint lists certificate profiles on which a principal owns a given workflow capability, such as enroll or revoke.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestCertificateProfileRequest
*/
func (a *RequestAPIService) RequestCertificateProfile(ctx context.Context) RequestAPIRequestCertificateProfileRequest {
	return RequestAPIRequestCertificateProfileRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []RequestableCertificateProfileResponse
func (a *RequestAPIService) RequestCertificateProfileExecute(r RequestAPIRequestCertificateProfileRequest) ([]models.RequestableCertificateProfileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.RequestableCertificateProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestCertificateProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.module != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "module", r.module, "form", "")
	}
	if r.workflow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "workflow", r.workflow, "form", "")
	}
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestCsvRequest struct {
	ctx                context.Context
	ApiService         *RequestAPIService
	requestSearchQuery *models.RequestSearchQuery
}

// The request search query
func (r RequestAPIRequestCsvRequest) RequestSearchQuery(requestSearchQuery models.RequestSearchQuery) RequestAPIRequestCsvRequest {
	r.requestSearchQuery = &requestSearchQuery
	return r
}

func (r RequestAPIRequestCsvRequest) Execute() (*http.Response, error) {
	return r.ApiService.RequestCsvExecute(r)
}

/*
RequestCsv Export requests

Send a request search query (in HRQL format) and return the request search results in CSV format

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestCsvRequest
*/
func (a *RequestAPIService) RequestCsv(ctx context.Context) RequestAPIRequestCsvRequest {
	return RequestAPIRequestCsvRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RequestAPIService) RequestCsvExecute(r RequestAPIRequestCsvRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestCsv")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/csv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestSearchQuery == nil {
		return nil, utils.ReportError("requestSearchQuery is required and must be specified")
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
	localVarPostBody = r.requestSearchQuery
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestDenyRequest struct {
	ctx                context.Context
	ApiService         *RequestAPIService
	requestDenyRequest *models.RequestDenyRequest
}

// The request to deny
func (r RequestAPIRequestDenyRequest) RequestDenyRequest(requestDenyRequest models.RequestDenyRequest) RequestAPIRequestDenyRequest {
	r.requestDenyRequest = &requestDenyRequest
	return r
}

func (r RequestAPIRequestDenyRequest) Execute() (*models.RequestApprove200Response, *http.Response, error) {
	return r.ApiService.RequestDenyExecute(r)
}

/*
RequestDeny Deny a request

Deny an existing request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestDenyRequest
*/
func (a *RequestAPIService) RequestDeny(ctx context.Context) RequestAPIRequestDenyRequest {
	return RequestAPIRequestDenyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestApprove200Response
func (a *RequestAPIService) RequestDenyExecute(r RequestAPIRequestDenyRequest) (*models.RequestApprove200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestApprove200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestDeny")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/deny"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestDenyRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("requestDenyRequest is required and must be specified")
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
	localVarPostBody = r.requestDenyRequest
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestDictionaryRequest struct {
	ctx        context.Context
	ApiService *RequestAPIService
}

func (r RequestAPIRequestDictionaryRequest) Execute() (*models.RequestSearchDictionaryResponse, *http.Response, error) {
	return r.ApiService.RequestDictionaryExecute(r)
}

/*
RequestDictionary Retrieve the request search dictionary

Return the request search dictionary. The dictionary is computed based on the principal and includes:

  - The list of certificate profiles on which the principal is authorized to search on;

  - The list of labels the principal is authorized to search on;

  - The list of modules available on the Horizon instance;

  - The list of available teams on the Horizon instance;

  - The list of available metadata on Horizon.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return RequestAPIRequestDictionaryRequest
*/
func (a *RequestAPIService) RequestDictionary(ctx context.Context) RequestAPIRequestDictionaryRequest {
	return RequestAPIRequestDictionaryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestSearchDictionaryResponse
func (a *RequestAPIService) RequestDictionaryExecute(r RequestAPIRequestDictionaryRequest) (*models.RequestSearchDictionaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestSearchDictionaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestDictionary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/search/dictionary"

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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestGetRequest struct {
	ctx        context.Context
	ApiService *RequestAPIService
	id         string
}

func (r RequestAPIRequestGetRequest) Execute() (*models.RequestGet200Response, *http.Response, error) {
	return r.ApiService.RequestGetExecute(r)
}

/*
RequestGet Retrieve a request

Retrieve an existing request based on its id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The request ID
	@return RequestAPIRequestGetRequest
*/
func (a *RequestAPIService) RequestGet(ctx context.Context, id string) RequestAPIRequestGetRequest {
	return RequestAPIRequestGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return RequestGet200Response
func (a *RequestAPIService) RequestGetExecute(r RequestAPIRequestGetRequest) (*models.RequestGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/{id}"
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestSearchRequest struct {
	ctx                context.Context
	ApiService         *RequestAPIService
	requestSearchQuery *models.RequestSearchQuery
}

// The request search query
func (r RequestAPIRequestSearchRequest) RequestSearchQuery(requestSearchQuery models.RequestSearchQuery) RequestAPIRequestSearchRequest {
	r.requestSearchQuery = &requestSearchQuery
	return r
}

func (r RequestAPIRequestSearchRequest) Execute() (*models.RequestSearchResultsResponse, *http.Response, error) {
	return r.ApiService.RequestSearchExecute(r)
}

/*
RequestSearch Search requests

Send a request search query (in HRQL format) and return the request search results

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestSearchRequest
*/
func (a *RequestAPIService) RequestSearch(ctx context.Context) RequestAPIRequestSearchRequest {
	return RequestAPIRequestSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestSearchResultsResponse
func (a *RequestAPIService) RequestSearchExecute(r RequestAPIRequestSearchRequest) (*models.RequestSearchResultsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestSearchResultsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestSearchQuery == nil {
		return localVarReturnValue, nil, utils.ReportError("requestSearchQuery is required and must be specified")
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
	localVarPostBody = r.requestSearchQuery
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestSubmitRequest struct {
	ctx                  context.Context
	ApiService           *RequestAPIService
	requestSubmitRequest *models.RequestSubmitRequest
}

// The Request to submit
func (r RequestAPIRequestSubmitRequest) RequestSubmitRequest(requestSubmitRequest models.RequestSubmitRequest) RequestAPIRequestSubmitRequest {
	r.requestSubmitRequest = &requestSubmitRequest
	return r
}

func (r RequestAPIRequestSubmitRequest) Execute() (*models.RequestSubmit201Response, *http.Response, error) {
	return r.ApiService.RequestSubmitExecute(r)
}

/*
RequestSubmit Submit a request

Submit a new request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestSubmitRequest
*/
func (a *RequestAPIService) RequestSubmit(ctx context.Context) RequestAPIRequestSubmitRequest {
	return RequestAPIRequestSubmitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestSubmit201Response
func (a *RequestAPIService) RequestSubmitExecute(r RequestAPIRequestSubmitRequest) (*models.RequestSubmit201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestSubmit201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestSubmit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/submit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestSubmitRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("requestSubmitRequest is required and must be specified")
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
	localVarPostBody = r.requestSubmitRequest
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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

type RequestAPIRequestTemplateRequest struct {
	ctx                    context.Context
	ApiService             *RequestAPIService
	requestTemplateRequest *models.RequestTemplateRequest
}

// The request on which to return the template
func (r RequestAPIRequestTemplateRequest) RequestTemplateRequest(requestTemplateRequest models.RequestTemplateRequest) RequestAPIRequestTemplateRequest {
	r.requestTemplateRequest = &requestTemplateRequest
	return r
}

func (r RequestAPIRequestTemplateRequest) Execute() (*models.RequestTemplate200Response, *http.Response, error) {
	return r.ApiService.RequestTemplateExecute(r)
}

/*
RequestTemplate Retrieve a request template

Retrieve the template to fulfill a specific request. The template indicates the required element to include when submitting a new request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RequestAPIRequestTemplateRequest
*/
func (a *RequestAPIService) RequestTemplate(ctx context.Context) RequestAPIRequestTemplateRequest {
	return RequestAPIRequestTemplateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestTemplate200Response
func (a *RequestAPIService) RequestTemplateExecute(r RequestAPIRequestTemplateRequest) (*models.RequestTemplate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RequestTemplate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestAPIService.RequestTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/requests/template"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestTemplateRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("requestTemplateRequest is required and must be specified")
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
	localVarPostBody = r.requestTemplateRequest
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

		// JWT POP
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
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
		if jwtPopCert, ok := r.ctx.Value("jwt-pop-cert").(*x509.Certificate); ok {
			// remove the API keys from the headers to avoid account authentication to interfere with the JWT POP authentication
			delete(localVarHeaderParams, "X-API-KEY")
			delete(localVarHeaderParams, "X-API-ID")
			if jwtPopSigner, ok := r.ctx.Value("jwt-pop-signer").(crypto.Signer); ok {
				// send without the Nonce to get replay nonce
				jwt, err := utils.CreateJWT(*jwtPopCert, jwtPopSigner, "")
				if err != nil {
					return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
				}
				localVarHeaderParams["X-JWT-CERT-POP"] = jwt
				// send the request a first time but without any data to get the replay nonce
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
				if err == nil {
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
			} else {
				return localVarReturnValue, nil, &GenericOpenAPIError{error: "jwt-pop-signer is required when jwt-pop-cert is provided"}
			}
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
