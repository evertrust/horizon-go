/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
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

// AcmeEabAPIService AcmeEabAPI service
type AcmeEabAPIService service

type AcmeEabAPIAcmeEabAddRequest struct {
	ctx           context.Context
	ApiService    *AcmeEabAPIService
	eabAddRequest *models.EabAddRequest
}

// The External Account Binding to add
func (r AcmeEabAPIAcmeEabAddRequest) EabAddRequest(eabAddRequest models.EabAddRequest) AcmeEabAPIAcmeEabAddRequest {
	r.eabAddRequest = &eabAddRequest
	return r
}

func (r AcmeEabAPIAcmeEabAddRequest) Execute() (*models.ExternalAccountBindingResponse, *http.Response, error) {
	return r.ApiService.AcmeEabAddExecute(r)
}

/*
AcmeEabAdd Add a new External Account Binding

Create a new External Account Binding

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AcmeEabAPIAcmeEabAddRequest
*/
func (a *AcmeEabAPIService) AcmeEabAdd(ctx context.Context) AcmeEabAPIAcmeEabAddRequest {
	return AcmeEabAPIAcmeEabAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingResponse
func (a *AcmeEabAPIService) AcmeEabAddExecute(r AcmeEabAPIAcmeEabAddRequest) (*models.ExternalAccountBindingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eabAddRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("eabAddRequest is required and must be specified")
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
	localVarPostBody = r.eabAddRequest
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

type AcmeEabAPIAcmeEabDeleteRequest struct {
	ctx        context.Context
	ApiService *AcmeEabAPIService
	name       string
}

func (r AcmeEabAPIAcmeEabDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcmeEabDeleteExecute(r)
}

/*
AcmeEabDelete Delete an External Account Binding by name

Delete an External Account Binding by its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the External Account Binding
	@return AcmeEabAPIAcmeEabDeleteRequest
*/
func (a *AcmeEabAPIService) AcmeEabDelete(ctx context.Context, name string) AcmeEabAPIAcmeEabDeleteRequest {
	return AcmeEabAPIAcmeEabDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *AcmeEabAPIService) AcmeEabDeleteExecute(r AcmeEabAPIAcmeEabDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type AcmeEabAPIAcmeEabGetRequest struct {
	ctx        context.Context
	ApiService *AcmeEabAPIService
	name       string
}

func (r AcmeEabAPIAcmeEabGetRequest) Execute() (*models.ExternalAccountBindingResponse, *http.Response, error) {
	return r.ApiService.AcmeEabGetExecute(r)
}

/*
AcmeEabGet Get an External Account Binding by name

Retrieve an External Account Binding by its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the External Account Binding
	@return AcmeEabAPIAcmeEabGetRequest
*/
func (a *AcmeEabAPIService) AcmeEabGet(ctx context.Context, name string) AcmeEabAPIAcmeEabGetRequest {
	return AcmeEabAPIAcmeEabGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingResponse
func (a *AcmeEabAPIService) AcmeEabGetExecute(r AcmeEabAPIAcmeEabGetRequest) (*models.ExternalAccountBindingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type AcmeEabAPIAcmeEabRenewRequest struct {
	ctx             context.Context
	ApiService      *AcmeEabAPIService
	name            string
	eabRenewRequest *models.EabRenewRequest
}

// Renewal parameters
func (r AcmeEabAPIAcmeEabRenewRequest) EabRenewRequest(eabRenewRequest models.EabRenewRequest) AcmeEabAPIAcmeEabRenewRequest {
	r.eabRenewRequest = &eabRenewRequest
	return r
}

func (r AcmeEabAPIAcmeEabRenewRequest) Execute() (*models.ExternalAccountBindingResponse, *http.Response, error) {
	return r.ApiService.AcmeEabRenewExecute(r)
}

/*
AcmeEabRenew Renew an External Account Binding

Renew an External Account Binding by generating a new MAC key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the External Account Binding to renew
	@return AcmeEabAPIAcmeEabRenewRequest
*/
func (a *AcmeEabAPIService) AcmeEabRenew(ctx context.Context, name string) AcmeEabAPIAcmeEabRenewRequest {
	return AcmeEabAPIAcmeEabRenewRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingResponse
func (a *AcmeEabAPIService) AcmeEabRenewExecute(r AcmeEabAPIAcmeEabRenewRequest) (*models.ExternalAccountBindingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabRenew")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab/{name}/renew"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eabRenewRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("eabRenewRequest is required and must be specified")
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
	localVarPostBody = r.eabRenewRequest
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

type AcmeEabAPIAcmeEabSearchRequest struct {
	ctx                          context.Context
	ApiService                   *AcmeEabAPIService
	externalAccountBindingSearch *models.ExternalAccountBindingSearch
}

// Search parameters
func (r AcmeEabAPIAcmeEabSearchRequest) ExternalAccountBindingSearch(externalAccountBindingSearch models.ExternalAccountBindingSearch) AcmeEabAPIAcmeEabSearchRequest {
	r.externalAccountBindingSearch = &externalAccountBindingSearch
	return r
}

func (r AcmeEabAPIAcmeEabSearchRequest) Execute() (*models.ExternalAccountBindingSearchResults, *http.Response, error) {
	return r.ApiService.AcmeEabSearchExecute(r)
}

/*
AcmeEabSearch Search External Account Bindings

Search for External Account Bindings using HEABQL query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AcmeEabAPIAcmeEabSearchRequest
*/
func (a *AcmeEabAPIService) AcmeEabSearch(ctx context.Context) AcmeEabAPIAcmeEabSearchRequest {
	return AcmeEabAPIAcmeEabSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingSearchResults
func (a *AcmeEabAPIService) AcmeEabSearchExecute(r AcmeEabAPIAcmeEabSearchRequest) (*models.ExternalAccountBindingSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalAccountBindingSearch == nil {
		return localVarReturnValue, nil, utils.ReportError("externalAccountBindingSearch is required and must be specified")
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
	localVarPostBody = r.externalAccountBindingSearch
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
		if localVarHTTPResponse.StatusCode == 408 {
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

type AcmeEabAPIAcmeEabUpdateRequest struct {
	ctx              context.Context
	ApiService       *AcmeEabAPIService
	eabUpdateRequest *models.EabUpdateRequest
}

// External Account Binding to update
func (r AcmeEabAPIAcmeEabUpdateRequest) EabUpdateRequest(eabUpdateRequest models.EabUpdateRequest) AcmeEabAPIAcmeEabUpdateRequest {
	r.eabUpdateRequest = &eabUpdateRequest
	return r
}

func (r AcmeEabAPIAcmeEabUpdateRequest) Execute() (*models.ExternalAccountBindingResponse, *http.Response, error) {
	return r.ApiService.AcmeEabUpdateExecute(r)
}

/*
AcmeEabUpdate Update an existing External Account Binding (metadata only)

Update an existing External Account Binding (metadata only). To update the status, use POST /eab/{name}/status

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AcmeEabAPIAcmeEabUpdateRequest
*/
func (a *AcmeEabAPIService) AcmeEabUpdate(ctx context.Context) AcmeEabAPIAcmeEabUpdateRequest {
	return AcmeEabAPIAcmeEabUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingResponse
func (a *AcmeEabAPIService) AcmeEabUpdateExecute(r AcmeEabAPIAcmeEabUpdateRequest) (*models.ExternalAccountBindingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eabUpdateRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("eabUpdateRequest is required and must be specified")
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
	localVarPostBody = r.eabUpdateRequest
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

type AcmeEabAPIAcmeEabUpdateStatusRequest struct {
	ctx                    context.Context
	ApiService             *AcmeEabAPIService
	eabName                string
	eabStatusUpdateRequest *models.EabStatusUpdateRequest
}

// Status and optional revocation date for certificates
func (r AcmeEabAPIAcmeEabUpdateStatusRequest) EabStatusUpdateRequest(eabStatusUpdateRequest models.EabStatusUpdateRequest) AcmeEabAPIAcmeEabUpdateStatusRequest {
	r.eabStatusUpdateRequest = &eabStatusUpdateRequest
	return r
}

func (r AcmeEabAPIAcmeEabUpdateStatusRequest) Execute() (*models.ExternalAccountBindingResponse, *http.Response, error) {
	return r.ApiService.AcmeEabUpdateStatusExecute(r)
}

/*
AcmeEabUpdateStatus Update External Account Binding status

Dedicated endpoint to update the status of an External Account Binding

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eabName The name of the External Account Binding
	@return AcmeEabAPIAcmeEabUpdateStatusRequest
*/
func (a *AcmeEabAPIService) AcmeEabUpdateStatus(ctx context.Context, eabName string) AcmeEabAPIAcmeEabUpdateStatusRequest {
	return AcmeEabAPIAcmeEabUpdateStatusRequest{
		ApiService: a,
		ctx:        ctx,
		eabName:    eabName,
	}
}

// Execute executes the request
//
//	@return ExternalAccountBindingResponse
func (a *AcmeEabAPIService) AcmeEabUpdateStatusExecute(r AcmeEabAPIAcmeEabUpdateStatusRequest) (*models.ExternalAccountBindingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ExternalAccountBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AcmeEabAPIService.AcmeEabUpdateStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/acme/eab/{eabName}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"eabName"+"}", url.PathEscape(parameterValueToString(r.eabName, "eabName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eabStatusUpdateRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("eabStatusUpdateRequest is required and must be specified")
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
	localVarPostBody = r.eabStatusUpdateRequest
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
