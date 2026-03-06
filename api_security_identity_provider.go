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

// SecurityIdentityProviderAPIService SecurityIdentityProviderAPI service
type SecurityIdentityProviderAPIService service

type SecurityIdentityProviderAPISecurityIdentityProviderAddRequest struct {
	ctx                                context.Context
	ApiService                         *SecurityIdentityProviderAPIService
	securityIdentityProviderAddRequest *models.SecurityIdentityProviderAddRequest
}

// The identity provider to register
func (r SecurityIdentityProviderAPISecurityIdentityProviderAddRequest) SecurityIdentityProviderAddRequest(securityIdentityProviderAddRequest models.SecurityIdentityProviderAddRequest) SecurityIdentityProviderAPISecurityIdentityProviderAddRequest {
	r.securityIdentityProviderAddRequest = &securityIdentityProviderAddRequest
	return r
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderAddRequest) Execute() (*models.SecurityIdentityProviderAdd201Response, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderAddExecute(r)
}

/*
SecurityIdentityProviderAdd Create a new identity provider

Create a new identity provider. The identity provider can be either a local identity provider or an OIDC identity provider.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityProviderAPISecurityIdentityProviderAddRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderAdd(ctx context.Context) SecurityIdentityProviderAPISecurityIdentityProviderAddRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityIdentityProviderAdd201Response
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderAddExecute(r SecurityIdentityProviderAPISecurityIdentityProviderAddRequest) (*models.SecurityIdentityProviderAdd201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.SecurityIdentityProviderAdd201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityIdentityProviderAddRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("securityIdentityProviderAddRequest is required and must be specified")
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
	localVarPostBody = r.securityIdentityProviderAddRequest
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityProviderAPIService
	name       string
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SecurityIdentityProviderDeleteExecute(r)
}

/*
SecurityIdentityProviderDelete Delete an identity provider

Delete an existing identity provider based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The internal name of the identity provider to retrieve
	@return SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderDelete(ctx context.Context, name string) SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderDeleteExecute(r SecurityIdentityProviderAPISecurityIdentityProviderDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers/{name}"
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest struct {
	ctx         context.Context
	ApiService  *SecurityIdentityProviderAPIService
	enabledOnUI *bool
}

// Whether the identity provider is visible on the UI
func (r SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest) EnabledOnUI(enabledOnUI bool) SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest {
	r.enabledOnUI = &enabledOnUI
	return r
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest) Execute() ([]models.EnabledIdentityProviderResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderEnabledExecute(r)
}

/*
SecurityIdentityProviderEnabled List the enabled identity provider(s)

Return the list of enabled dynamic identity provider(s)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderEnabled(ctx context.Context) SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []EnabledIdentityProviderResponse
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderEnabledExecute(r SecurityIdentityProviderAPISecurityIdentityProviderEnabledRequest) ([]models.EnabledIdentityProviderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.EnabledIdentityProviderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderEnabled")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers/dynamic/enabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.enabledOnUI != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enabledOnUI", r.enabledOnUI, "form", "")
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderGetRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityProviderAPIService
	name       string
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderGetRequest) Execute() (*models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderGetExecute(r)
}

/*
SecurityIdentityProviderGet Retrieve a existing identity provider

Retrieve a existing identity provider based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The internal name of the identity provider to retrieve
	@return SecurityIdentityProviderAPISecurityIdentityProviderGetRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderGet(ctx context.Context, name string) SecurityIdentityProviderAPISecurityIdentityProviderGetRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return SecurityIdentityProviderList200ResponseInner
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderGetExecute(r SecurityIdentityProviderAPISecurityIdentityProviderGetRequest) (*models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.SecurityIdentityProviderList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers/{name}"
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderListRequest struct {
	ctx        context.Context
	ApiService *SecurityIdentityProviderAPIService
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderListRequest) Execute() ([]models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderListExecute(r)
}

/*
SecurityIdentityProviderList List all the identity provider(s)

List all the existing identity provider(s), regardless of whether they are enabled.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityProviderAPISecurityIdentityProviderListRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderList(ctx context.Context) SecurityIdentityProviderAPISecurityIdentityProviderListRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []SecurityIdentityProviderList200ResponseInner
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderListExecute(r SecurityIdentityProviderAPISecurityIdentityProviderListRequest) ([]models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.SecurityIdentityProviderList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers"

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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest struct {
	ctx                        context.Context
	ApiService                 *SecurityIdentityProviderAPIService
	principalInfoSearchRequest *models.PrincipalInfoSearchRequest
}

// The principal search request
func (r SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest) PrincipalInfoSearchRequest(principalInfoSearchRequest models.PrincipalInfoSearchRequest) SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest {
	r.principalInfoSearchRequest = &principalInfoSearchRequest
	return r
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest) Execute() ([]models.PrincipalInfoSearchResultResponse, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderSearchExecute(r)
}

/*
SecurityIdentityProviderSearch Retrieve the provider of a principal

Retrieve which identity provider holds the authentication information for a given identifier or a given contact e-mail.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderSearch(ctx context.Context) SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []PrincipalInfoSearchResultResponse
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderSearchExecute(r SecurityIdentityProviderAPISecurityIdentityProviderSearchRequest) ([]models.PrincipalInfoSearchResultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.PrincipalInfoSearchResultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.principalInfoSearchRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("principalInfoSearchRequest is required and must be specified")
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
	localVarPostBody = r.principalInfoSearchRequest
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest struct {
	ctx                                   context.Context
	ApiService                            *SecurityIdentityProviderAPIService
	securityIdentityProviderUpdateRequest *models.SecurityIdentityProviderUpdateRequest
}

// Identity provider to update
func (r SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest) SecurityIdentityProviderUpdateRequest(securityIdentityProviderUpdateRequest models.SecurityIdentityProviderUpdateRequest) SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest {
	r.securityIdentityProviderUpdateRequest = &securityIdentityProviderUpdateRequest
	return r
}

func (r SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest) Execute() (*models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	return r.ApiService.SecurityIdentityProviderUpdateExecute(r)
}

/*
SecurityIdentityProviderUpdate Update an existing identity provider

Update an existing identity provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest
*/
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderUpdate(ctx context.Context) SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest {
	return SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityIdentityProviderList200ResponseInner
func (a *SecurityIdentityProviderAPIService) SecurityIdentityProviderUpdateExecute(r SecurityIdentityProviderAPISecurityIdentityProviderUpdateRequest) (*models.SecurityIdentityProviderList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.SecurityIdentityProviderList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityIdentityProviderAPIService.SecurityIdentityProviderUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/security/identity/providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityIdentityProviderUpdateRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("securityIdentityProviderUpdateRequest is required and must be specified")
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
	localVarPostBody = r.securityIdentityProviderUpdateRequest
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
				req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, "", localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
