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

// DiscoveryCampaignAPIService DiscoveryCampaignAPI service
type DiscoveryCampaignAPIService service

type DiscoveryCampaignAPIDiscoveryCampaignAddRequest struct {
	ctx               context.Context
	ApiService        *DiscoveryCampaignAPIService
	discoveryCampaign *models.DiscoveryCampaign
}

// Discovery campaign to register
func (r DiscoveryCampaignAPIDiscoveryCampaignAddRequest) DiscoveryCampaign(discoveryCampaign models.DiscoveryCampaign) DiscoveryCampaignAPIDiscoveryCampaignAddRequest {
	r.discoveryCampaign = &discoveryCampaign
	return r
}

func (r DiscoveryCampaignAPIDiscoveryCampaignAddRequest) Execute() (*models.DiscoveryCampaignResponse, *http.Response, error) {
	return r.ApiService.DiscoveryCampaignAddExecute(r)
}

/*
DiscoveryCampaignAdd Create a new discovery campaign

Create a new discovery campaign in Horizon

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryCampaignAPIDiscoveryCampaignAddRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignAdd(ctx context.Context) DiscoveryCampaignAPIDiscoveryCampaignAddRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DiscoveryCampaignResponse
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignAddExecute(r DiscoveryCampaignAPIDiscoveryCampaignAddRequest) (*models.DiscoveryCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DiscoveryCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryCampaign == nil {
		return localVarReturnValue, nil, utils.ReportError("discoveryCampaign is required and must be specified")
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
	localVarPostBody = r.discoveryCampaign
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

type DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest struct {
	ctx        context.Context
	ApiService *DiscoveryCampaignAPIService
	name       string
}

func (r DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DiscoveryCampaignDeleteExecute(r)
}

/*
DiscoveryCampaignDelete Delete a discovery campaign

Delete an existing discovery campaign based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the discovery campaign to delete
	@return DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignDelete(ctx context.Context, name string) DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignDeleteExecute(r DiscoveryCampaignAPIDiscoveryCampaignDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns/{name}"
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

type DiscoveryCampaignAPIDiscoveryCampaignFlushRequest struct {
	ctx        context.Context
	ApiService *DiscoveryCampaignAPIService
	name       string
}

func (r DiscoveryCampaignAPIDiscoveryCampaignFlushRequest) Execute() (*http.Response, error) {
	return r.ApiService.DiscoveryCampaignFlushExecute(r)
}

/*
DiscoveryCampaignFlush Flush a discovery campaign

Flush an existing discovery campaign based on its name. Flushing a discovery campaign will:

  - Remove discovery info related to the considered discovery campaign from any certificate (managed or discovered);

  - Delete any discovered certificate which is no longer associated to a discovery campaign.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @param name The name of the discovery campaign to flush
    @return DiscoveryCampaignAPIDiscoveryCampaignFlushRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignFlush(ctx context.Context, name string) DiscoveryCampaignAPIDiscoveryCampaignFlushRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignFlushRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignFlushExecute(r DiscoveryCampaignAPIDiscoveryCampaignFlushRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignFlush")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns/{name}"
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

type DiscoveryCampaignAPIDiscoveryCampaignGetRequest struct {
	ctx        context.Context
	ApiService *DiscoveryCampaignAPIService
	name       string
}

func (r DiscoveryCampaignAPIDiscoveryCampaignGetRequest) Execute() (*models.DiscoveryCampaignResponse, *http.Response, error) {
	return r.ApiService.DiscoveryCampaignGetExecute(r)
}

/*
DiscoveryCampaignGet Retrieve a discovery campaign

Retrieve a specific discovery campaign based on its name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the discovery campaign to retrieve
	@return DiscoveryCampaignAPIDiscoveryCampaignGetRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignGet(ctx context.Context, name string) DiscoveryCampaignAPIDiscoveryCampaignGetRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return DiscoveryCampaignResponse
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignGetExecute(r DiscoveryCampaignAPIDiscoveryCampaignGetRequest) (*models.DiscoveryCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DiscoveryCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns/{name}"
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

type DiscoveryCampaignAPIDiscoveryCampaignListRequest struct {
	ctx        context.Context
	ApiService *DiscoveryCampaignAPIService
}

func (r DiscoveryCampaignAPIDiscoveryCampaignListRequest) Execute() ([]models.DiscoveryCampaignResponse, *http.Response, error) {
	return r.ApiService.DiscoveryCampaignListExecute(r)
}

/*
DiscoveryCampaignList List discovery campaign(s)

List all the existing discovery campaign(s)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryCampaignAPIDiscoveryCampaignListRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignList(ctx context.Context) DiscoveryCampaignAPIDiscoveryCampaignListRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []DiscoveryCampaignResponse
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignListExecute(r DiscoveryCampaignAPIDiscoveryCampaignListRequest) ([]models.DiscoveryCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.DiscoveryCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns"

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

type DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest struct {
	ctx               context.Context
	ApiService        *DiscoveryCampaignAPIService
	discoveryCampaign *models.DiscoveryCampaign
}

// Discovery campaign to update
func (r DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest) DiscoveryCampaign(discoveryCampaign models.DiscoveryCampaign) DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest {
	r.discoveryCampaign = &discoveryCampaign
	return r
}

func (r DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest) Execute() (*models.DiscoveryCampaignResponse, *http.Response, error) {
	return r.ApiService.DiscoveryCampaignUpdateExecute(r)
}

/*
DiscoveryCampaignUpdate Update a discovery campaign

Update an existing discovery campaign's information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest
*/
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignUpdate(ctx context.Context) DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest {
	return DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DiscoveryCampaignResponse
func (a *DiscoveryCampaignAPIService) DiscoveryCampaignUpdateExecute(r DiscoveryCampaignAPIDiscoveryCampaignUpdateRequest) (*models.DiscoveryCampaignResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.DiscoveryCampaignResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveryCampaignAPIService.DiscoveryCampaignUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/discovery/campaigns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveryCampaign == nil {
		return localVarReturnValue, nil, utils.ReportError("discoveryCampaign is required and must be specified")
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
	localVarPostBody = r.discoveryCampaign
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
