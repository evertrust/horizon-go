/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
	"gopkg.in/validator.v2"
)

// RequestApprove200Response - struct for RequestApprove200Response
type RequestApprove200Response struct {
	EstEnrollRequestOnApproveResponse    *EstEnrollRequestOnApproveResponse
	ScepEnrollRequestOnApproveResponse   *ScepEnrollRequestOnApproveResponse
	WebRAEnrollRequestOnApproveResponse  *WebRAEnrollRequestOnApproveResponse
	WebRAImportRequestOnApproveResponse  *WebRAImportRequestOnApproveResponse
	WebRAMigrateRequestOnApproveResponse *WebRAMigrateRequestOnApproveResponse
	WebRARecoverRequestOnApproveResponse *WebRARecoverRequestOnApproveResponse
	WebRARenewRequestOnApproveResponse   *WebRARenewRequestOnApproveResponse
	WebRARevokeRequestOnApproveResponse  *WebRARevokeRequestOnApproveResponse
	WebRAUpdateRequestOnApproveResponse  *WebRAUpdateRequestOnApproveResponse
}

// EstEnrollRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns EstEnrollRequestOnApproveResponse wrapped in RequestApprove200Response
func EstEnrollRequestOnApproveResponseAsRequestApprove200Response(v *EstEnrollRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		EstEnrollRequestOnApproveResponse: v,
	}
}

// ScepEnrollRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns ScepEnrollRequestOnApproveResponse wrapped in RequestApprove200Response
func ScepEnrollRequestOnApproveResponseAsRequestApprove200Response(v *ScepEnrollRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		ScepEnrollRequestOnApproveResponse: v,
	}
}

// WebRAEnrollRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRAEnrollRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRAEnrollRequestOnApproveResponseAsRequestApprove200Response(v *WebRAEnrollRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRAEnrollRequestOnApproveResponse: v,
	}
}

// WebRAImportRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRAImportRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRAImportRequestOnApproveResponseAsRequestApprove200Response(v *WebRAImportRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRAImportRequestOnApproveResponse: v,
	}
}

// WebRAMigrateRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRAMigrateRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRAMigrateRequestOnApproveResponseAsRequestApprove200Response(v *WebRAMigrateRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRAMigrateRequestOnApproveResponse: v,
	}
}

// WebRARecoverRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRARecoverRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRARecoverRequestOnApproveResponseAsRequestApprove200Response(v *WebRARecoverRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRARecoverRequestOnApproveResponse: v,
	}
}

// WebRARenewRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRARenewRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRARenewRequestOnApproveResponseAsRequestApprove200Response(v *WebRARenewRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRARenewRequestOnApproveResponse: v,
	}
}

// WebRARevokeRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRARevokeRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRARevokeRequestOnApproveResponseAsRequestApprove200Response(v *WebRARevokeRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRARevokeRequestOnApproveResponse: v,
	}
}

// WebRAUpdateRequestOnApproveResponseAsRequestApprove200Response is a convenience function that returns WebRAUpdateRequestOnApproveResponse wrapped in RequestApprove200Response
func WebRAUpdateRequestOnApproveResponseAsRequestApprove200Response(v *WebRAUpdateRequestOnApproveResponse) RequestApprove200Response {
	return RequestApprove200Response{
		WebRAUpdateRequestOnApproveResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestApprove200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnApproveResponse)
	if err == nil {
		jsonEstEnrollRequestOnApproveResponse, _ := json.Marshal(dst.EstEnrollRequestOnApproveResponse)
		if string(jsonEstEnrollRequestOnApproveResponse) == "{}" { // empty struct
			dst.EstEnrollRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnApproveResponse)
			match++
		}
	} else {
		dst.EstEnrollRequestOnApproveResponse = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnApproveResponse)
	if err == nil {
		jsonScepEnrollRequestOnApproveResponse, _ := json.Marshal(dst.ScepEnrollRequestOnApproveResponse)
		if string(jsonScepEnrollRequestOnApproveResponse) == "{}" { // empty struct
			dst.ScepEnrollRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnApproveResponse)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnApproveResponse)
	if err == nil {
		jsonWebRAEnrollRequestOnApproveResponse, _ := json.Marshal(dst.WebRAEnrollRequestOnApproveResponse)
		if string(jsonWebRAEnrollRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRAImportRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnApproveResponse)
	if err == nil {
		jsonWebRAImportRequestOnApproveResponse, _ := json.Marshal(dst.WebRAImportRequestOnApproveResponse)
		if string(jsonWebRAImportRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRAImportRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRAImportRequestOnApproveResponse = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnApproveResponse)
	if err == nil {
		jsonWebRAMigrateRequestOnApproveResponse, _ := json.Marshal(dst.WebRAMigrateRequestOnApproveResponse)
		if string(jsonWebRAMigrateRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnApproveResponse = nil
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

	// try to unmarshal data into WebRAUpdateRequestOnApproveResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnApproveResponse)
	if err == nil {
		jsonWebRAUpdateRequestOnApproveResponse, _ := json.Marshal(dst.WebRAUpdateRequestOnApproveResponse)
		if string(jsonWebRAUpdateRequestOnApproveResponse) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnApproveResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnApproveResponse)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnApproveResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestApprove200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestApprove200Response) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnApproveResponse != nil {
		return json.Marshal(&src.EstEnrollRequestOnApproveResponse)
	}

	if src.ScepEnrollRequestOnApproveResponse != nil {
		return json.Marshal(&src.ScepEnrollRequestOnApproveResponse)
	}

	if src.WebRAEnrollRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnApproveResponse)
	}

	if src.WebRAImportRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRAImportRequestOnApproveResponse)
	}

	if src.WebRAMigrateRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnApproveResponse)
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

	if src.WebRAUpdateRequestOnApproveResponse != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnApproveResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestApprove200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnApproveResponse != nil {
		return obj.EstEnrollRequestOnApproveResponse
	}

	if obj.ScepEnrollRequestOnApproveResponse != nil {
		return obj.ScepEnrollRequestOnApproveResponse
	}

	if obj.WebRAEnrollRequestOnApproveResponse != nil {
		return obj.WebRAEnrollRequestOnApproveResponse
	}

	if obj.WebRAImportRequestOnApproveResponse != nil {
		return obj.WebRAImportRequestOnApproveResponse
	}

	if obj.WebRAMigrateRequestOnApproveResponse != nil {
		return obj.WebRAMigrateRequestOnApproveResponse
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

	if obj.WebRAUpdateRequestOnApproveResponse != nil {
		return obj.WebRAUpdateRequestOnApproveResponse
	}

	// all schemas are nil
	return nil
}

type NullableRequestApprove200Response struct {
	value *RequestApprove200Response
	isSet bool
}

func (v NullableRequestApprove200Response) Get() *RequestApprove200Response {
	return v.value
}

func (v *NullableRequestApprove200Response) Set(val *RequestApprove200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprove200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprove200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprove200Response(val *RequestApprove200Response) *NullableRequestApprove200Response {
	return &NullableRequestApprove200Response{value: val, isSet: true}
}

func (v NullableRequestApprove200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprove200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
