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

// RequestApproveRequest - struct for RequestApproveRequest
type RequestApproveRequest struct {
	EstEnrollRequestOnApprove    *EstEnrollRequestOnApprove
	ScepEnrollRequestOnApprove   *ScepEnrollRequestOnApprove
	WebRAEnrollRequestOnApprove  *WebRAEnrollRequestOnApprove
	WebRAImportRequestOnApprove  *WebRAImportRequestOnApprove
	WebRAMigrateRequestOnApprove *WebRAMigrateRequestOnApprove
	WebRARecoverRequestOnApprove *WebRARecoverRequestOnApprove
	WebRARenewRequestOnApprove   *WebRARenewRequestOnApprove
	WebRARevokeRequestOnApprove  *WebRARevokeRequestOnApprove
	WebRAUpdateRequestOnApprove  *WebRAUpdateRequestOnApprove
}

// EstEnrollRequestOnApproveAsRequestApproveRequest is a convenience function that returns EstEnrollRequestOnApprove wrapped in RequestApproveRequest
func EstEnrollRequestOnApproveAsRequestApproveRequest(v *EstEnrollRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		EstEnrollRequestOnApprove: v,
	}
}

// ScepEnrollRequestOnApproveAsRequestApproveRequest is a convenience function that returns ScepEnrollRequestOnApprove wrapped in RequestApproveRequest
func ScepEnrollRequestOnApproveAsRequestApproveRequest(v *ScepEnrollRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		ScepEnrollRequestOnApprove: v,
	}
}

// WebRAEnrollRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRAEnrollRequestOnApprove wrapped in RequestApproveRequest
func WebRAEnrollRequestOnApproveAsRequestApproveRequest(v *WebRAEnrollRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRAEnrollRequestOnApprove: v,
	}
}

// WebRAImportRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRAImportRequestOnApprove wrapped in RequestApproveRequest
func WebRAImportRequestOnApproveAsRequestApproveRequest(v *WebRAImportRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRAImportRequestOnApprove: v,
	}
}

// WebRAMigrateRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRAMigrateRequestOnApprove wrapped in RequestApproveRequest
func WebRAMigrateRequestOnApproveAsRequestApproveRequest(v *WebRAMigrateRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRAMigrateRequestOnApprove: v,
	}
}

// WebRARecoverRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRARecoverRequestOnApprove wrapped in RequestApproveRequest
func WebRARecoverRequestOnApproveAsRequestApproveRequest(v *WebRARecoverRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRARecoverRequestOnApprove: v,
	}
}

// WebRARenewRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRARenewRequestOnApprove wrapped in RequestApproveRequest
func WebRARenewRequestOnApproveAsRequestApproveRequest(v *WebRARenewRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRARenewRequestOnApprove: v,
	}
}

// WebRARevokeRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRARevokeRequestOnApprove wrapped in RequestApproveRequest
func WebRARevokeRequestOnApproveAsRequestApproveRequest(v *WebRARevokeRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRARevokeRequestOnApprove: v,
	}
}

// WebRAUpdateRequestOnApproveAsRequestApproveRequest is a convenience function that returns WebRAUpdateRequestOnApprove wrapped in RequestApproveRequest
func WebRAUpdateRequestOnApproveAsRequestApproveRequest(v *WebRAUpdateRequestOnApprove) RequestApproveRequest {
	return RequestApproveRequest{
		WebRAUpdateRequestOnApprove: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestApproveRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnApprove)
	if err == nil {
		jsonEstEnrollRequestOnApprove, _ := json.Marshal(dst.EstEnrollRequestOnApprove)
		if string(jsonEstEnrollRequestOnApprove) == "{}" { // empty struct
			dst.EstEnrollRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnApprove)
			match++
		}
	} else {
		dst.EstEnrollRequestOnApprove = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnApprove)
	if err == nil {
		jsonScepEnrollRequestOnApprove, _ := json.Marshal(dst.ScepEnrollRequestOnApprove)
		if string(jsonScepEnrollRequestOnApprove) == "{}" { // empty struct
			dst.ScepEnrollRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnApprove)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnApprove = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnApprove)
	if err == nil {
		jsonWebRAEnrollRequestOnApprove, _ := json.Marshal(dst.WebRAEnrollRequestOnApprove)
		if string(jsonWebRAEnrollRequestOnApprove) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnApprove)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnApprove = nil
	}

	// try to unmarshal data into WebRAImportRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnApprove)
	if err == nil {
		jsonWebRAImportRequestOnApprove, _ := json.Marshal(dst.WebRAImportRequestOnApprove)
		if string(jsonWebRAImportRequestOnApprove) == "{}" { // empty struct
			dst.WebRAImportRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnApprove)
			match++
		}
	} else {
		dst.WebRAImportRequestOnApprove = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnApprove)
	if err == nil {
		jsonWebRAMigrateRequestOnApprove, _ := json.Marshal(dst.WebRAMigrateRequestOnApprove)
		if string(jsonWebRAMigrateRequestOnApprove) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnApprove)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnApprove = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnApprove)
	if err == nil {
		jsonWebRARecoverRequestOnApprove, _ := json.Marshal(dst.WebRARecoverRequestOnApprove)
		if string(jsonWebRARecoverRequestOnApprove) == "{}" { // empty struct
			dst.WebRARecoverRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnApprove)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnApprove = nil
	}

	// try to unmarshal data into WebRARenewRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnApprove)
	if err == nil {
		jsonWebRARenewRequestOnApprove, _ := json.Marshal(dst.WebRARenewRequestOnApprove)
		if string(jsonWebRARenewRequestOnApprove) == "{}" { // empty struct
			dst.WebRARenewRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnApprove)
			match++
		}
	} else {
		dst.WebRARenewRequestOnApprove = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnApprove)
	if err == nil {
		jsonWebRARevokeRequestOnApprove, _ := json.Marshal(dst.WebRARevokeRequestOnApprove)
		if string(jsonWebRARevokeRequestOnApprove) == "{}" { // empty struct
			dst.WebRARevokeRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnApprove)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnApprove = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnApprove
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnApprove)
	if err == nil {
		jsonWebRAUpdateRequestOnApprove, _ := json.Marshal(dst.WebRAUpdateRequestOnApprove)
		if string(jsonWebRAUpdateRequestOnApprove) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnApprove = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnApprove)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnApprove = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestApproveRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestApproveRequest) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnApprove != nil {
		return json.Marshal(&src.EstEnrollRequestOnApprove)
	}

	if src.ScepEnrollRequestOnApprove != nil {
		return json.Marshal(&src.ScepEnrollRequestOnApprove)
	}

	if src.WebRAEnrollRequestOnApprove != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnApprove)
	}

	if src.WebRAImportRequestOnApprove != nil {
		return json.Marshal(&src.WebRAImportRequestOnApprove)
	}

	if src.WebRAMigrateRequestOnApprove != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnApprove)
	}

	if src.WebRARecoverRequestOnApprove != nil {
		return json.Marshal(&src.WebRARecoverRequestOnApprove)
	}

	if src.WebRARenewRequestOnApprove != nil {
		return json.Marshal(&src.WebRARenewRequestOnApprove)
	}

	if src.WebRARevokeRequestOnApprove != nil {
		return json.Marshal(&src.WebRARevokeRequestOnApprove)
	}

	if src.WebRAUpdateRequestOnApprove != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnApprove)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestApproveRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnApprove != nil {
		return obj.EstEnrollRequestOnApprove
	}

	if obj.ScepEnrollRequestOnApprove != nil {
		return obj.ScepEnrollRequestOnApprove
	}

	if obj.WebRAEnrollRequestOnApprove != nil {
		return obj.WebRAEnrollRequestOnApprove
	}

	if obj.WebRAImportRequestOnApprove != nil {
		return obj.WebRAImportRequestOnApprove
	}

	if obj.WebRAMigrateRequestOnApprove != nil {
		return obj.WebRAMigrateRequestOnApprove
	}

	if obj.WebRARecoverRequestOnApprove != nil {
		return obj.WebRARecoverRequestOnApprove
	}

	if obj.WebRARenewRequestOnApprove != nil {
		return obj.WebRARenewRequestOnApprove
	}

	if obj.WebRARevokeRequestOnApprove != nil {
		return obj.WebRARevokeRequestOnApprove
	}

	if obj.WebRAUpdateRequestOnApprove != nil {
		return obj.WebRAUpdateRequestOnApprove
	}

	// all schemas are nil
	return nil
}

type NullableRequestApproveRequest struct {
	value *RequestApproveRequest
	isSet bool
}

func (v NullableRequestApproveRequest) Get() *RequestApproveRequest {
	return v.value
}

func (v *NullableRequestApproveRequest) Set(val *RequestApproveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApproveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApproveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApproveRequest(val *RequestApproveRequest) *NullableRequestApproveRequest {
	return &NullableRequestApproveRequest{value: val, isSet: true}
}

func (v NullableRequestApproveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApproveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
