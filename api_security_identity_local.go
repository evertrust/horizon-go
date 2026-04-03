/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
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

// SecurityIdentityLocalAPIService SecurityIdentityLocalAPI service
type SecurityIdentityLocalAPIService service

type SecurityIdentityLocalAPISecurityIdentityLocalAddRequest struct {
	ctx                context.Context
	ApiService         *SecurityIdentityLocalAPIService
	localIdentityOnAdd *models.LocalIdentityOnAdd
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalAddRequest) LocalIdentityOnAdd(localIdentityOnAdd models.LocalIdentityOnAdd) SecurityIdentityLocalAPISecurityIdentityLocalAddRequest {
	r.localIdentityOnAdd = &localIdentityOnAdd
	return r
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalAddRequest) Execute() (*models.LocalIdentityResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityLocalAddExecute(r)
}

/*
SecurityIdentityLocalAdd Create a local identity

Create a local identity. By default, a local identity doesn't have a password and therefore cannot log in to Horizon. To set a password, call the [set password endpoint](#tag/security.identity.local/operation/security.identity.local.password.set) after creating the local identity.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityLocalAPISecurityIdentityLocalAddRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalAdd(ctx context.Context) SecurityIdentityLocalAPISecurityIdentityLocalAddRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalIdentityResponse
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalAddExecute(r SecurityIdentityLocalAPISecurityIdentityLocalAddRequest) (*models.LocalIdentityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.LocalIdentityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.localIdentityOnAdd == nil {
		return localVarReturnValue, nil, utils.ReportError("localIdentityOnAdd is required and must be specified")
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
	localVarPostBody = r.localIdentityOnAdd
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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityLocalAPIService
	identifier string
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityIdentityLocalDeleteExecute(r)
}

/*
SecurityIdentityLocalDelete Delete a local identity

Delete an existing local identity based on its identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identifier Local identity identifier.
	@return SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalDelete(ctx context.Context, identifier string) SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		identifier: identifier,
	}
}

// Execute executes the request
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalDeleteExecute(r SecurityIdentityLocalAPISecurityIdentityLocalDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterValueToString(r.identifier, "identifier")), -1)

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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalGetRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityLocalAPIService
	identifier string
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalGetRequest) Execute() (*models.LocalIdentityResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityLocalGetExecute(r)
}

/*
SecurityIdentityLocalGet Retrieve a local identity

Given an identifier, retrieve the full Local Identity object using this endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identifier Local identity identifier.
	@return SecurityIdentityLocalAPISecurityIdentityLocalGetRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalGet(ctx context.Context, identifier string) SecurityIdentityLocalAPISecurityIdentityLocalGetRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalGetRequest{
		ApiService: a,
		ctx:        ctx,
		identifier: identifier,
	}
}

// Execute executes the request
//
//	@return LocalIdentityResponse
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalGetExecute(r SecurityIdentityLocalAPISecurityIdentityLocalGetRequest) (*models.LocalIdentityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.LocalIdentityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterValueToString(r.identifier, "identifier")), -1)

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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalListRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityLocalAPIService
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalListRequest) Execute() ([]models.LocalIdentityResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityLocalListExecute(r)
}

/*
SecurityIdentityLocalList List local identities

Retrieve the list of all existing local identities.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityLocalAPISecurityIdentityLocalListRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalList(ctx context.Context) SecurityIdentityLocalAPISecurityIdentityLocalListRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []LocalIdentityResponse
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalListExecute(r SecurityIdentityLocalAPISecurityIdentityLocalListRequest) ([]models.LocalIdentityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.LocalIdentityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals"

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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest struct {
	ctx                  context.Context
	ApiService           *SecurityIdentityLocalAPIService
	resetPasswordRequest *models.ResetPasswordRequest
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest) ResetPasswordRequest(resetPasswordRequest models.ResetPasswordRequest) SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest {
	r.resetPasswordRequest = &resetPasswordRequest
	return r
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityIdentityLocalPasswordResetExecute(r)
}

/*
SecurityIdentityLocalPasswordReset Reset a password

This is the second step of the password reset flow. Following a password reset request, the user will receive a reset UUID by email. They can then send this UUID and a new password to reset their password.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordReset(ctx context.Context) SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordResetExecute(r SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalPasswordReset")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals/password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resetPasswordRequest == nil {
		return nil, utils.ReportError("resetPasswordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.resetPasswordRequest
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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityLocalAPIService
	identifier string
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityIdentityLocalPasswordResetRequestExecute(r)
}

/*
SecurityIdentityLocalPasswordResetRequest Request a password reset

This is the first step in the password reset flow. The user will receive a reset UUID by email that can be used to complete the password reset request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identifier Local identity identifier
	@return SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordResetRequest(ctx context.Context, identifier string) SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest{
		ApiService: a,
		ctx:        ctx,
		identifier: identifier,
	}
}

// Execute executes the request
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordResetRequestExecute(r SecurityIdentityLocalAPISecurityIdentityLocalPasswordResetRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalPasswordResetRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals/password/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterValueToString(r.identifier, "identifier")), -1)

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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest struct {
	ctx                context.Context
	ApiService         *SecurityIdentityLocalAPIService
	setPasswordRequest *models.SetPasswordRequest
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest) SetPasswordRequest(setPasswordRequest models.SetPasswordRequest) SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest {
	r.setPasswordRequest = &setPasswordRequest
	return r
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityIdentityLocalPasswordSetExecute(r)
}

/*
SecurityIdentityLocalPasswordSet Set the password for a local identity

You can define the password that will be used by this local identity to log in to the web UI or use APIs. You must have the right management permissions to perform this action, and the password must meet the local identity provider's password policy requirements, if any has been defined.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordSet(ctx context.Context) SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalPasswordSetExecute(r SecurityIdentityLocalAPISecurityIdentityLocalPasswordSetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalPasswordSet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setPasswordRequest == nil {
		return nil, utils.ReportError("setPasswordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.setPasswordRequest
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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest struct {
	ctx           context.Context
	ApiService    *SecurityIdentityLocalAPIService
	localIdentity *models.LocalIdentity
}

// Local identity to update
func (r SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest) LocalIdentity(localIdentity models.LocalIdentity) SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest {
	r.localIdentity = &localIdentity
	return r
}

func (r SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest) Execute() (*models.LocalIdentityResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityLocalUpdateExecute(r)
}

/*
SecurityIdentityLocalUpdate Update a local identity

Update a local identity

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest
*/
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalUpdate(ctx context.Context) SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest {
	return SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalIdentityResponse
func (a *SecurityIdentityLocalAPIService) SecurityIdentityLocalUpdateExecute(r SecurityIdentityLocalAPISecurityIdentityLocalUpdateRequest) (*models.LocalIdentityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.LocalIdentityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityLocalAPIService.SecurityIdentityLocalUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/locals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.localIdentity == nil {
		return localVarReturnValue, nil, utils.ReportError("localIdentity is required and must be specified")
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
	localVarPostBody = r.localIdentity
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
			req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "{}", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
