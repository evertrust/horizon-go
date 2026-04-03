/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

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

// RequestSubmit201Response - struct for RequestSubmit201Response
type RequestSubmit201Response struct {
	EstEnrollRequestOnSubmitResponse    *EstEnrollRequestOnSubmitResponse
	ScepEnrollRequestOnSubmitResponse   *ScepEnrollRequestOnSubmitResponse
	WebRAEnrollRequestOnSubmitResponse  *WebRAEnrollRequestOnSubmitResponse
	WebRAImportRequestOnSubmitResponse  *WebRAImportRequestOnSubmitResponse
	WebRAMigrateRequestOnSubmitResponse *WebRAMigrateRequestOnSubmitResponse
	WebRARecoverRequestOnSubmitResponse *WebRARecoverRequestOnSubmitResponse
	WebRARenewRequestOnSubmitResponse   *WebRARenewRequestOnSubmitResponse
	WebRARevokeRequestOnSubmitResponse  *WebRARevokeRequestOnSubmitResponse
	WebRAUpdateRequestOnSubmitResponse  *WebRAUpdateRequestOnSubmitResponse
}

// EstEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns EstEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func EstEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *EstEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		EstEnrollRequestOnSubmitResponse: v,
	}
}

// ScepEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns ScepEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func ScepEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *ScepEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		ScepEnrollRequestOnSubmitResponse: v,
	}
}

// WebRAEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAEnrollRequestOnSubmitResponse: v,
	}
}

// WebRAImportRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAImportRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAImportRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAImportRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAImportRequestOnSubmitResponse: v,
	}
}

// WebRAMigrateRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAMigrateRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAMigrateRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAMigrateRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAMigrateRequestOnSubmitResponse: v,
	}
}

// WebRARecoverRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARecoverRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARecoverRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARecoverRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARecoverRequestOnSubmitResponse: v,
	}
}

// WebRARenewRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARenewRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARenewRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARenewRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARenewRequestOnSubmitResponse: v,
	}
}

// WebRARevokeRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARevokeRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARevokeRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARevokeRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARevokeRequestOnSubmitResponse: v,
	}
}

// WebRAUpdateRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAUpdateRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAUpdateRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAUpdateRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAUpdateRequestOnSubmitResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSubmit201Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnSubmitResponse)
	if err == nil {
		jsonEstEnrollRequestOnSubmitResponse, _ := json.Marshal(dst.EstEnrollRequestOnSubmitResponse)
		if string(jsonEstEnrollRequestOnSubmitResponse) == "{}" { // empty struct
			dst.EstEnrollRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.EstEnrollRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnSubmitResponse)
	if err == nil {
		jsonScepEnrollRequestOnSubmitResponse, _ := json.Marshal(dst.ScepEnrollRequestOnSubmitResponse)
		if string(jsonScepEnrollRequestOnSubmitResponse) == "{}" { // empty struct
			dst.ScepEnrollRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnSubmitResponse)
	if err == nil {
		jsonWebRAEnrollRequestOnSubmitResponse, _ := json.Marshal(dst.WebRAEnrollRequestOnSubmitResponse)
		if string(jsonWebRAEnrollRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRAImportRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnSubmitResponse)
	if err == nil {
		jsonWebRAImportRequestOnSubmitResponse, _ := json.Marshal(dst.WebRAImportRequestOnSubmitResponse)
		if string(jsonWebRAImportRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRAImportRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRAImportRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnSubmitResponse)
	if err == nil {
		jsonWebRAMigrateRequestOnSubmitResponse, _ := json.Marshal(dst.WebRAMigrateRequestOnSubmitResponse)
		if string(jsonWebRAMigrateRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnSubmitResponse)
	if err == nil {
		jsonWebRARecoverRequestOnSubmitResponse, _ := json.Marshal(dst.WebRARecoverRequestOnSubmitResponse)
		if string(jsonWebRARecoverRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRARecoverRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRARenewRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnSubmitResponse)
	if err == nil {
		jsonWebRARenewRequestOnSubmitResponse, _ := json.Marshal(dst.WebRARenewRequestOnSubmitResponse)
		if string(jsonWebRARenewRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRARenewRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRARenewRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnSubmitResponse)
	if err == nil {
		jsonWebRARevokeRequestOnSubmitResponse, _ := json.Marshal(dst.WebRARevokeRequestOnSubmitResponse)
		if string(jsonWebRARevokeRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRARevokeRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnSubmitResponse = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnSubmitResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnSubmitResponse)
	if err == nil {
		jsonWebRAUpdateRequestOnSubmitResponse, _ := json.Marshal(dst.WebRAUpdateRequestOnSubmitResponse)
		if string(jsonWebRAUpdateRequestOnSubmitResponse) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnSubmitResponse = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnSubmitResponse)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnSubmitResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestSubmit201Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSubmit201Response) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.EstEnrollRequestOnSubmitResponse)
	}

	if src.ScepEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.ScepEnrollRequestOnSubmitResponse)
	}

	if src.WebRAEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnSubmitResponse)
	}

	if src.WebRAImportRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAImportRequestOnSubmitResponse)
	}

	if src.WebRAMigrateRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnSubmitResponse)
	}

	if src.WebRARecoverRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARecoverRequestOnSubmitResponse)
	}

	if src.WebRARenewRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARenewRequestOnSubmitResponse)
	}

	if src.WebRARevokeRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARevokeRequestOnSubmitResponse)
	}

	if src.WebRAUpdateRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnSubmitResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSubmit201Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnSubmitResponse != nil {
		return obj.EstEnrollRequestOnSubmitResponse
	}

	if obj.ScepEnrollRequestOnSubmitResponse != nil {
		return obj.ScepEnrollRequestOnSubmitResponse
	}

	if obj.WebRAEnrollRequestOnSubmitResponse != nil {
		return obj.WebRAEnrollRequestOnSubmitResponse
	}

	if obj.WebRAImportRequestOnSubmitResponse != nil {
		return obj.WebRAImportRequestOnSubmitResponse
	}

	if obj.WebRAMigrateRequestOnSubmitResponse != nil {
		return obj.WebRAMigrateRequestOnSubmitResponse
	}

	if obj.WebRARecoverRequestOnSubmitResponse != nil {
		return obj.WebRARecoverRequestOnSubmitResponse
	}

	if obj.WebRARenewRequestOnSubmitResponse != nil {
		return obj.WebRARenewRequestOnSubmitResponse
	}

	if obj.WebRARevokeRequestOnSubmitResponse != nil {
		return obj.WebRARevokeRequestOnSubmitResponse
	}

	if obj.WebRAUpdateRequestOnSubmitResponse != nil {
		return obj.WebRAUpdateRequestOnSubmitResponse
	}

	// all schemas are nil
	return nil
}

type NullableRequestSubmit201Response struct {
	value *RequestSubmit201Response
	isSet bool
}

func (v NullableRequestSubmit201Response) Get() *RequestSubmit201Response {
	return v.value
}

func (v *NullableRequestSubmit201Response) Set(val *RequestSubmit201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmit201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmit201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmit201Response(val *RequestSubmit201Response) *NullableRequestSubmit201Response {
	return &NullableRequestSubmit201Response{value: val, isSet: true}
}

func (v NullableRequestSubmit201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmit201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
