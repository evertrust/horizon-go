/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
	"gopkg.in/validator.v2"
)

// RequestGet200Response - struct for RequestGet200Response
type RequestGet200Response struct {
	EstEnrollRequestOnGetResponse        *EstEnrollRequestOnGetResponse
	ScepEnrollRequestOnGetResponse       *ScepEnrollRequestOnGetResponse
	WebRAEnrollRequestOnGetResponse      *WebRAEnrollRequestOnGetResponse
	WebRAImportRequestOnGetResponse      *WebRAImportRequestOnGetResponse
	WebRAMigrateRequestOnGetResponse     *WebRAMigrateRequestOnGetResponse
	WebRARecoverRequestOnApproveResponse *WebRARecoverRequestOnApproveResponse
	WebRARenewRequestOnApproveResponse   *WebRARenewRequestOnApproveResponse
	WebRARevokeRequestOnApproveResponse  *WebRARevokeRequestOnApproveResponse
	WebRAUpdateRequestOnGetResponse      *WebRAUpdateRequestOnGetResponse
}

// EstEnrollRequestOnGetResponseAsRequestGet200Response is a convenience function that returns EstEnrollRequestOnGetResponse wrapped in RequestGet200Response
func EstEnrollRequestOnGetResponseAsRequestGet200Response(v *EstEnrollRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		EstEnrollRequestOnGetResponse: v,
	}
}

// ScepEnrollRequestOnGetResponseAsRequestGet200Response is a convenience function that returns ScepEnrollRequestOnGetResponse wrapped in RequestGet200Response
func ScepEnrollRequestOnGetResponseAsRequestGet200Response(v *ScepEnrollRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		ScepEnrollRequestOnGetResponse: v,
	}
}

// WebRAEnrollRequestOnGetResponseAsRequestGet200Response is a convenience function that returns WebRAEnrollRequestOnGetResponse wrapped in RequestGet200Response
func WebRAEnrollRequestOnGetResponseAsRequestGet200Response(v *WebRAEnrollRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRAEnrollRequestOnGetResponse: v,
	}
}

// WebRAImportRequestOnGetResponseAsRequestGet200Response is a convenience function that returns WebRAImportRequestOnGetResponse wrapped in RequestGet200Response
func WebRAImportRequestOnGetResponseAsRequestGet200Response(v *WebRAImportRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRAImportRequestOnGetResponse: v,
	}
}

// WebRAMigrateRequestOnGetResponseAsRequestGet200Response is a convenience function that returns WebRAMigrateRequestOnGetResponse wrapped in RequestGet200Response
func WebRAMigrateRequestOnGetResponseAsRequestGet200Response(v *WebRAMigrateRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRAMigrateRequestOnGetResponse: v,
	}
}

// WebRARecoverRequestOnApproveResponseAsRequestGet200Response is a convenience function that returns WebRARecoverRequestOnApproveResponse wrapped in RequestGet200Response
func WebRARecoverRequestOnApproveResponseAsRequestGet200Response(v *WebRARecoverRequestOnApproveResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRARecoverRequestOnApproveResponse: v,
	}
}

// WebRARenewRequestOnApproveResponseAsRequestGet200Response is a convenience function that returns WebRARenewRequestOnApproveResponse wrapped in RequestGet200Response
func WebRARenewRequestOnApproveResponseAsRequestGet200Response(v *WebRARenewRequestOnApproveResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRARenewRequestOnApproveResponse: v,
	}
}

// WebRARevokeRequestOnApproveResponseAsRequestGet200Response is a convenience function that returns WebRARevokeRequestOnApproveResponse wrapped in RequestGet200Response
func WebRARevokeRequestOnApproveResponseAsRequestGet200Response(v *WebRARevokeRequestOnApproveResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRARevokeRequestOnApproveResponse: v,
	}
}

// WebRAUpdateRequestOnGetResponseAsRequestGet200Response is a convenience function that returns WebRAUpdateRequestOnGetResponse wrapped in RequestGet200Response
func WebRAUpdateRequestOnGetResponseAsRequestGet200Response(v *WebRAUpdateRequestOnGetResponse) RequestGet200Response {
	return RequestGet200Response{
		WebRAUpdateRequestOnGetResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestGet200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnGetResponse)
	if err == nil {
		jsonEstEnrollRequestOnGetResponse, _ := json.Marshal(dst.EstEnrollRequestOnGetResponse)
		if string(jsonEstEnrollRequestOnGetResponse) == "{}" { // empty struct
			dst.EstEnrollRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnGetResponse)
			match++
		}
	} else {
		dst.EstEnrollRequestOnGetResponse = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnGetResponse)
	if err == nil {
		jsonScepEnrollRequestOnGetResponse, _ := json.Marshal(dst.ScepEnrollRequestOnGetResponse)
		if string(jsonScepEnrollRequestOnGetResponse) == "{}" { // empty struct
			dst.ScepEnrollRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnGetResponse)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnGetResponse = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnGetResponse)
	if err == nil {
		jsonWebRAEnrollRequestOnGetResponse, _ := json.Marshal(dst.WebRAEnrollRequestOnGetResponse)
		if string(jsonWebRAEnrollRequestOnGetResponse) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnGetResponse)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnGetResponse = nil
	}

	// try to unmarshal data into WebRAImportRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnGetResponse)
	if err == nil {
		jsonWebRAImportRequestOnGetResponse, _ := json.Marshal(dst.WebRAImportRequestOnGetResponse)
		if string(jsonWebRAImportRequestOnGetResponse) == "{}" { // empty struct
			dst.WebRAImportRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnGetResponse)
			match++
		}
	} else {
		dst.WebRAImportRequestOnGetResponse = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnGetResponse)
	if err == nil {
		jsonWebRAMigrateRequestOnGetResponse, _ := json.Marshal(dst.WebRAMigrateRequestOnGetResponse)
		if string(jsonWebRAMigrateRequestOnGetResponse) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnGetResponse)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnGetResponse = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnApproveResponse)
	if err == nil {
		jsonWebRARecoverRequestOnApproveResponse, _ := json.Marshal(dst.WebRARecoverRequestOnApproveResponse)
		if string(jsonWebRARecoverRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRARecoverRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRARenewRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnApproveResponse)
	if err == nil {
		jsonWebRARenewRequestOnApproveResponse, _ := json.Marshal(dst.WebRARenewRequestOnApproveResponse)
		if string(jsonWebRARenewRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRARenewRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRARenewRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnApproveResponse)
	if err == nil {
		jsonWebRARevokeRequestOnApproveResponse, _ := json.Marshal(dst.WebRARevokeRequestOnApproveResponse)
		if string(jsonWebRARevokeRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRARevokeRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnGetResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnGetResponse)
	if err == nil {
		jsonWebRAUpdateRequestOnGetResponse, _ := json.Marshal(dst.WebRAUpdateRequestOnGetResponse)
		if string(jsonWebRAUpdateRequestOnGetResponse) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnGetResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnGetResponse)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnGetResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestGet200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestGet200Response) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnGetResponse != nil {
		return json.Marshal(&src.EstEnrollRequestOnGetResponse)
	}

	if src.ScepEnrollRequestOnGetResponse != nil {
		return json.Marshal(&src.ScepEnrollRequestOnGetResponse)
	}

	if src.WebRAEnrollRequestOnGetResponse != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnGetResponse)
	}

	if src.WebRAImportRequestOnGetResponse != nil {
		return json.Marshal(&src.WebRAImportRequestOnGetResponse)
	}

	if src.WebRAMigrateRequestOnGetResponse != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnGetResponse)
	}

	if src.WebRARecoverRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRARecoverRequestOnApproveResponse)
	}

	if src.WebRARenewRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRARenewRequestOnApproveResponse)
	}

	if src.WebRARevokeRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRARevokeRequestOnApproveResponse)
	}

	if src.WebRAUpdateRequestOnGetResponse != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnGetResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestGet200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnGetResponse != nil {
		return obj.EstEnrollRequestOnGetResponse
	}

	if obj.ScepEnrollRequestOnGetResponse != nil {
		return obj.ScepEnrollRequestOnGetResponse
	}

	if obj.WebRAEnrollRequestOnGetResponse != nil {
		return obj.WebRAEnrollRequestOnGetResponse
	}

	if obj.WebRAImportRequestOnGetResponse != nil {
		return obj.WebRAImportRequestOnGetResponse
	}

	if obj.WebRAMigrateRequestOnGetResponse != nil {
		return obj.WebRAMigrateRequestOnGetResponse
	}

	if obj.WebRARecoverRequestOnApproveResponse != nil {
		return obj.WebRARecoverRequestOnApproveResponse
	}

	if obj.WebRARenewRequestOnApproveResponse != nil {
		return obj.WebRARenewRequestOnApproveResponse
	}

	if obj.WebRARevokeRequestOnApproveResponse != nil {
		return obj.WebRARevokeRequestOnApproveResponse
	}

	if obj.WebRAUpdateRequestOnGetResponse != nil {
		return obj.WebRAUpdateRequestOnGetResponse
	}

	// all schemas are nil
	return nil
}

type NullableRequestGet200Response struct {
	value *RequestGet200Response
	isSet bool
}

func (v NullableRequestGet200Response) Get() *RequestGet200Response {
	return v.value
}

func (v *NullableRequestGet200Response) Set(val *RequestGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestGet200Response(val *RequestGet200Response) *NullableRequestGet200Response {
	return &NullableRequestGet200Response{value: val, isSet: true}
}

func (v NullableRequestGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
