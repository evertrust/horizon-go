/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
	"gopkg.in/validator.v2"
)

// RequestTemplate200Response - struct for RequestTemplate200Response
type RequestTemplate200Response struct {
	EstEnrollRequestOnTemplateResponse    *EstEnrollRequestOnTemplateResponse
	ScepEnrollRequestOnTemplateResponse   *ScepEnrollRequestOnTemplateResponse
	WebRAEnrollRequestOnTemplateResponse  *WebRAEnrollRequestOnTemplateResponse
	WebRAImportRequestOnTemplateResponse  *WebRAImportRequestOnTemplateResponse
	WebRAMigrateRequestOnTemplateResponse *WebRAMigrateRequestOnTemplateResponse
	WebRARecoverRequestOnTemplateResponse *WebRARecoverRequestOnTemplateResponse
	WebRARenewRequestOnTemplateResponse   *WebRARenewRequestOnTemplateResponse
	WebRARevokeRequestOnTemplateResponse  *WebRARevokeRequestOnTemplateResponse
	WebRAUpdateRequestOnTemplateResponse  *WebRAUpdateRequestOnTemplateResponse
}

// EstEnrollRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns EstEnrollRequestOnTemplateResponse wrapped in RequestTemplate200Response
func EstEnrollRequestOnTemplateResponseAsRequestTemplate200Response(v *EstEnrollRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		EstEnrollRequestOnTemplateResponse: v,
	}
}

// ScepEnrollRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns ScepEnrollRequestOnTemplateResponse wrapped in RequestTemplate200Response
func ScepEnrollRequestOnTemplateResponseAsRequestTemplate200Response(v *ScepEnrollRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		ScepEnrollRequestOnTemplateResponse: v,
	}
}

// WebRAEnrollRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRAEnrollRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRAEnrollRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRAEnrollRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRAEnrollRequestOnTemplateResponse: v,
	}
}

// WebRAImportRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRAImportRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRAImportRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRAImportRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRAImportRequestOnTemplateResponse: v,
	}
}

// WebRAMigrateRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRAMigrateRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRAMigrateRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRAMigrateRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRAMigrateRequestOnTemplateResponse: v,
	}
}

// WebRARecoverRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRARecoverRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRARecoverRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRARecoverRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRARecoverRequestOnTemplateResponse: v,
	}
}

// WebRARenewRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRARenewRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRARenewRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRARenewRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRARenewRequestOnTemplateResponse: v,
	}
}

// WebRARevokeRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRARevokeRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRARevokeRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRARevokeRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRARevokeRequestOnTemplateResponse: v,
	}
}

// WebRAUpdateRequestOnTemplateResponseAsRequestTemplate200Response is a convenience function that returns WebRAUpdateRequestOnTemplateResponse wrapped in RequestTemplate200Response
func WebRAUpdateRequestOnTemplateResponseAsRequestTemplate200Response(v *WebRAUpdateRequestOnTemplateResponse) RequestTemplate200Response {
	return RequestTemplate200Response{
		WebRAUpdateRequestOnTemplateResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTemplate200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnTemplateResponse)
	if err == nil {
		jsonEstEnrollRequestOnTemplateResponse, _ := json.Marshal(dst.EstEnrollRequestOnTemplateResponse)
		if string(jsonEstEnrollRequestOnTemplateResponse) == "{}" { // empty struct
			dst.EstEnrollRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.EstEnrollRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnTemplateResponse)
	if err == nil {
		jsonScepEnrollRequestOnTemplateResponse, _ := json.Marshal(dst.ScepEnrollRequestOnTemplateResponse)
		if string(jsonScepEnrollRequestOnTemplateResponse) == "{}" { // empty struct
			dst.ScepEnrollRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnTemplateResponse)
	if err == nil {
		jsonWebRAEnrollRequestOnTemplateResponse, _ := json.Marshal(dst.WebRAEnrollRequestOnTemplateResponse)
		if string(jsonWebRAEnrollRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRAImportRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnTemplateResponse)
	if err == nil {
		jsonWebRAImportRequestOnTemplateResponse, _ := json.Marshal(dst.WebRAImportRequestOnTemplateResponse)
		if string(jsonWebRAImportRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRAImportRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRAImportRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnTemplateResponse)
	if err == nil {
		jsonWebRAMigrateRequestOnTemplateResponse, _ := json.Marshal(dst.WebRAMigrateRequestOnTemplateResponse)
		if string(jsonWebRAMigrateRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnTemplateResponse)
	if err == nil {
		jsonWebRARecoverRequestOnTemplateResponse, _ := json.Marshal(dst.WebRARecoverRequestOnTemplateResponse)
		if string(jsonWebRARecoverRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRARecoverRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRARenewRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnTemplateResponse)
	if err == nil {
		jsonWebRARenewRequestOnTemplateResponse, _ := json.Marshal(dst.WebRARenewRequestOnTemplateResponse)
		if string(jsonWebRARenewRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRARenewRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRARenewRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnTemplateResponse)
	if err == nil {
		jsonWebRARevokeRequestOnTemplateResponse, _ := json.Marshal(dst.WebRARevokeRequestOnTemplateResponse)
		if string(jsonWebRARevokeRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRARevokeRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnTemplateResponse = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnTemplateResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnTemplateResponse)
	if err == nil {
		jsonWebRAUpdateRequestOnTemplateResponse, _ := json.Marshal(dst.WebRAUpdateRequestOnTemplateResponse)
		if string(jsonWebRAUpdateRequestOnTemplateResponse) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnTemplateResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnTemplateResponse)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnTemplateResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestTemplate200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTemplate200Response) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnTemplateResponse != nil {
		return json.Marshal(&src.EstEnrollRequestOnTemplateResponse)
	}

	if src.ScepEnrollRequestOnTemplateResponse != nil {
		return json.Marshal(&src.ScepEnrollRequestOnTemplateResponse)
	}

	if src.WebRAEnrollRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnTemplateResponse)
	}

	if src.WebRAImportRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRAImportRequestOnTemplateResponse)
	}

	if src.WebRAMigrateRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnTemplateResponse)
	}

	if src.WebRARecoverRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRARecoverRequestOnTemplateResponse)
	}

	if src.WebRARenewRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRARenewRequestOnTemplateResponse)
	}

	if src.WebRARevokeRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRARevokeRequestOnTemplateResponse)
	}

	if src.WebRAUpdateRequestOnTemplateResponse != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnTemplateResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTemplate200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnTemplateResponse != nil {
		return obj.EstEnrollRequestOnTemplateResponse
	}

	if obj.ScepEnrollRequestOnTemplateResponse != nil {
		return obj.ScepEnrollRequestOnTemplateResponse
	}

	if obj.WebRAEnrollRequestOnTemplateResponse != nil {
		return obj.WebRAEnrollRequestOnTemplateResponse
	}

	if obj.WebRAImportRequestOnTemplateResponse != nil {
		return obj.WebRAImportRequestOnTemplateResponse
	}

	if obj.WebRAMigrateRequestOnTemplateResponse != nil {
		return obj.WebRAMigrateRequestOnTemplateResponse
	}

	if obj.WebRARecoverRequestOnTemplateResponse != nil {
		return obj.WebRARecoverRequestOnTemplateResponse
	}

	if obj.WebRARenewRequestOnTemplateResponse != nil {
		return obj.WebRARenewRequestOnTemplateResponse
	}

	if obj.WebRARevokeRequestOnTemplateResponse != nil {
		return obj.WebRARevokeRequestOnTemplateResponse
	}

	if obj.WebRAUpdateRequestOnTemplateResponse != nil {
		return obj.WebRAUpdateRequestOnTemplateResponse
	}

	// all schemas are nil
	return nil
}

type NullableRequestTemplate200Response struct {
	value *RequestTemplate200Response
	isSet bool
}

func (v NullableRequestTemplate200Response) Get() *RequestTemplate200Response {
	return v.value
}

func (v *NullableRequestTemplate200Response) Set(val *RequestTemplate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTemplate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTemplate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTemplate200Response(val *RequestTemplate200Response) *NullableRequestTemplate200Response {
	return &NullableRequestTemplate200Response{value: val, isSet: true}
}

func (v NullableRequestTemplate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTemplate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
