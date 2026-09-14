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

// DashboardAPIService DashboardAPI service
type DashboardAPIService service

type DashboardAPISecurityPrincipalDashboardAddRequest struct {
	ctx        context.Context
	ApiService *DashboardAPIService
	dashboard  *models.Dashboard
}

// Dashboard to add
func (r DashboardAPISecurityPrincipalDashboardAddRequest) Dashboard(dashboard models.Dashboard) DashboardAPISecurityPrincipalDashboardAddRequest {
	r.dashboard = &dashboard
	return r
}

func (r DashboardAPISecurityPrincipalDashboardAddRequest) Execute() (*models.DashboardResponse, *http.Response, error) {
	return r.ApiService.SecurityPrincipalDashboardAddExecute(r)
}

/*
SecurityPrincipalDashboardAdd Add a dashboard for the authenticated principal

Add a dashboard for the authenticated principal

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DashboardAPISecurityPrincipalDashboardAddRequest
*/
func (a *DashboardAPIService) SecurityPrincipalDashboardAdd(ctx context.Context) DashboardAPISecurityPrincipalDashboardAddRequest {
	return DashboardAPISecurityPrincipalDashboardAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DashboardResponse
func (a *DashboardAPIService) SecurityPrincipalDashboardAddExecute(r DashboardAPISecurityPrincipalDashboardAddRequest) (*models.DashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardAPIService.SecurityPrincipalDashboardAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/principals/dashboards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dashboard == nil {
		return localVarReturnValue, nil, utils.ReportError("dashboard is required and must be specified")
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
	localVarPostBody = r.dashboard
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

type DashboardAPISecurityPrincipalDashboardDeleteRequest struct {
	ctx        context.Context
	ApiService *DashboardAPIService
	name       string
}

func (r DashboardAPISecurityPrincipalDashboardDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityPrincipalDashboardDeleteExecute(r)
}

/*
SecurityPrincipalDashboardDelete Delete an existing dashboard for the authenticated principal

Delete an existing dashboard for the authenticated principal based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the dashboard
	@return DashboardAPISecurityPrincipalDashboardDeleteRequest
*/
func (a *DashboardAPIService) SecurityPrincipalDashboardDelete(ctx context.Context, name string) DashboardAPISecurityPrincipalDashboardDeleteRequest {
	return DashboardAPISecurityPrincipalDashboardDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *DashboardAPIService) SecurityPrincipalDashboardDeleteExecute(r DashboardAPISecurityPrincipalDashboardDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardAPIService.SecurityPrincipalDashboardDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/principals/dashboards/{name}"
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

type DashboardAPISecurityPrincipalDashboardGetRequest struct {
	ctx        context.Context
	ApiService *DashboardAPIService
	name       string
}

func (r DashboardAPISecurityPrincipalDashboardGetRequest) Execute() (*models.DashboardResponse, *http.Response, error) {
	return r.ApiService.SecurityPrincipalDashboardGetExecute(r)
}

/*
SecurityPrincipalDashboardGet Retrieve an existing dashboard for the authenticated principal

Retrieve an existing dashboard for the authenticated principal based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the dashboard
	@return DashboardAPISecurityPrincipalDashboardGetRequest
*/
func (a *DashboardAPIService) SecurityPrincipalDashboardGet(ctx context.Context, name string) DashboardAPISecurityPrincipalDashboardGetRequest {
	return DashboardAPISecurityPrincipalDashboardGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return DashboardResponse
func (a *DashboardAPIService) SecurityPrincipalDashboardGetExecute(r DashboardAPISecurityPrincipalDashboardGetRequest) (*models.DashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardAPIService.SecurityPrincipalDashboardGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/principals/dashboards/{name}"
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

type DashboardAPISecurityPrincipalDashboardListRequest struct {
	ctx        context.Context
	ApiService *DashboardAPIService
	type_      *string
}

// The type of dashboards to filter on
func (r DashboardAPISecurityPrincipalDashboardListRequest) Type_(type_ string) DashboardAPISecurityPrincipalDashboardListRequest {
	r.type_ = &type_
	return r
}

func (r DashboardAPISecurityPrincipalDashboardListRequest) Execute() ([]models.DashboardResponse, *http.Response, error) {
	return r.ApiService.SecurityPrincipalDashboardListExecute(r)
}

/*
SecurityPrincipalDashboardList List the authenticated principal dashboards

List the authenticated principal dashboards, optionally filtered by dashboard type

  - certificate

  - request

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return DashboardAPISecurityPrincipalDashboardListRequest
*/
func (a *DashboardAPIService) SecurityPrincipalDashboardList(ctx context.Context) DashboardAPISecurityPrincipalDashboardListRequest {
	return DashboardAPISecurityPrincipalDashboardListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []DashboardResponse
func (a *DashboardAPIService) SecurityPrincipalDashboardListExecute(r DashboardAPISecurityPrincipalDashboardListRequest) ([]models.DashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.DashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardAPIService.SecurityPrincipalDashboardList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/principals/dashboards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

type DashboardAPISecurityPrincipalDashboardUpdateRequest struct {
	ctx        context.Context
	ApiService *DashboardAPIService
	dashboard  *models.Dashboard
}

// Dashboard to update
func (r DashboardAPISecurityPrincipalDashboardUpdateRequest) Dashboard(dashboard models.Dashboard) DashboardAPISecurityPrincipalDashboardUpdateRequest {
	r.dashboard = &dashboard
	return r
}

func (r DashboardAPISecurityPrincipalDashboardUpdateRequest) Execute() (*models.DashboardResponse, *http.Response, error) {
	return r.ApiService.SecurityPrincipalDashboardUpdateExecute(r)
}

/*
SecurityPrincipalDashboardUpdate Update a dashboard for the authenticated principal

Update a dashboard for the authenticated principal

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DashboardAPISecurityPrincipalDashboardUpdateRequest
*/
func (a *DashboardAPIService) SecurityPrincipalDashboardUpdate(ctx context.Context) DashboardAPISecurityPrincipalDashboardUpdateRequest {
	return DashboardAPISecurityPrincipalDashboardUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DashboardResponse
func (a *DashboardAPIService) SecurityPrincipalDashboardUpdateExecute(r DashboardAPISecurityPrincipalDashboardUpdateRequest) (*models.DashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardAPIService.SecurityPrincipalDashboardUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/principals/dashboards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dashboard == nil {
		return localVarReturnValue, nil, utils.ReportError("dashboard is required and must be specified")
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
	localVarPostBody = r.dashboard
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
