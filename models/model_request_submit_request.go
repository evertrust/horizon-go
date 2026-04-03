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

// RequestSubmitRequest - struct for RequestSubmitRequest
type RequestSubmitRequest struct {
	EstEnrollRequestOnSubmit    *EstEnrollRequestOnSubmit
	ScepEnrollRequestOnSubmit   *ScepEnrollRequestOnSubmit
	WebRAEnrollRequestOnSubmit  *WebRAEnrollRequestOnSubmit
	WebRAImportRequestOnSubmit  *WebRAImportRequestOnSubmit
	WebRAMigrateRequestOnSubmit *WebRAMigrateRequestOnSubmit
	WebRARecoverRequestOnSubmit *WebRARecoverRequestOnSubmit
	WebRARenewRequestOnSubmit   *WebRARenewRequestOnSubmit
	WebRARevokeRequestOnSubmit  *WebRARevokeRequestOnSubmit
	WebRAUpdateRequestOnSubmit  *WebRAUpdateRequestOnSubmit
}

// EstEnrollRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns EstEnrollRequestOnSubmit wrapped in RequestSubmitRequest
func EstEnrollRequestOnSubmitAsRequestSubmitRequest(v *EstEnrollRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		EstEnrollRequestOnSubmit: v,
	}
}

// ScepEnrollRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns ScepEnrollRequestOnSubmit wrapped in RequestSubmitRequest
func ScepEnrollRequestOnSubmitAsRequestSubmitRequest(v *ScepEnrollRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		ScepEnrollRequestOnSubmit: v,
	}
}

// WebRAEnrollRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRAEnrollRequestOnSubmit wrapped in RequestSubmitRequest
func WebRAEnrollRequestOnSubmitAsRequestSubmitRequest(v *WebRAEnrollRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRAEnrollRequestOnSubmit: v,
	}
}

// WebRAImportRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRAImportRequestOnSubmit wrapped in RequestSubmitRequest
func WebRAImportRequestOnSubmitAsRequestSubmitRequest(v *WebRAImportRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRAImportRequestOnSubmit: v,
	}
}

// WebRAMigrateRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRAMigrateRequestOnSubmit wrapped in RequestSubmitRequest
func WebRAMigrateRequestOnSubmitAsRequestSubmitRequest(v *WebRAMigrateRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRAMigrateRequestOnSubmit: v,
	}
}

// WebRARecoverRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRARecoverRequestOnSubmit wrapped in RequestSubmitRequest
func WebRARecoverRequestOnSubmitAsRequestSubmitRequest(v *WebRARecoverRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRARecoverRequestOnSubmit: v,
	}
}

// WebRARenewRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRARenewRequestOnSubmit wrapped in RequestSubmitRequest
func WebRARenewRequestOnSubmitAsRequestSubmitRequest(v *WebRARenewRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRARenewRequestOnSubmit: v,
	}
}

// WebRARevokeRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRARevokeRequestOnSubmit wrapped in RequestSubmitRequest
func WebRARevokeRequestOnSubmitAsRequestSubmitRequest(v *WebRARevokeRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRARevokeRequestOnSubmit: v,
	}
}

// WebRAUpdateRequestOnSubmitAsRequestSubmitRequest is a convenience function that returns WebRAUpdateRequestOnSubmit wrapped in RequestSubmitRequest
func WebRAUpdateRequestOnSubmitAsRequestSubmitRequest(v *WebRAUpdateRequestOnSubmit) RequestSubmitRequest {
	return RequestSubmitRequest{
		WebRAUpdateRequestOnSubmit: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSubmitRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnSubmit)
	if err == nil {
		jsonEstEnrollRequestOnSubmit, _ := json.Marshal(dst.EstEnrollRequestOnSubmit)
		if string(jsonEstEnrollRequestOnSubmit) == "{}" { // empty struct
			dst.EstEnrollRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnSubmit)
			match++
		}
	} else {
		dst.EstEnrollRequestOnSubmit = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnSubmit)
	if err == nil {
		jsonScepEnrollRequestOnSubmit, _ := json.Marshal(dst.ScepEnrollRequestOnSubmit)
		if string(jsonScepEnrollRequestOnSubmit) == "{}" { // empty struct
			dst.ScepEnrollRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnSubmit)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnSubmit)
	if err == nil {
		jsonWebRAEnrollRequestOnSubmit, _ := json.Marshal(dst.WebRAEnrollRequestOnSubmit)
		if string(jsonWebRAEnrollRequestOnSubmit) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRAImportRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnSubmit)
	if err == nil {
		jsonWebRAImportRequestOnSubmit, _ := json.Marshal(dst.WebRAImportRequestOnSubmit)
		if string(jsonWebRAImportRequestOnSubmit) == "{}" { // empty struct
			dst.WebRAImportRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRAImportRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnSubmit)
	if err == nil {
		jsonWebRAMigrateRequestOnSubmit, _ := json.Marshal(dst.WebRAMigrateRequestOnSubmit)
		if string(jsonWebRAMigrateRequestOnSubmit) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnSubmit)
	if err == nil {
		jsonWebRARecoverRequestOnSubmit, _ := json.Marshal(dst.WebRARecoverRequestOnSubmit)
		if string(jsonWebRARecoverRequestOnSubmit) == "{}" { // empty struct
			dst.WebRARecoverRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRARenewRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnSubmit)
	if err == nil {
		jsonWebRARenewRequestOnSubmit, _ := json.Marshal(dst.WebRARenewRequestOnSubmit)
		if string(jsonWebRARenewRequestOnSubmit) == "{}" { // empty struct
			dst.WebRARenewRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRARenewRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnSubmit)
	if err == nil {
		jsonWebRARevokeRequestOnSubmit, _ := json.Marshal(dst.WebRARevokeRequestOnSubmit)
		if string(jsonWebRARevokeRequestOnSubmit) == "{}" { // empty struct
			dst.WebRARevokeRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnSubmit = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnSubmit
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnSubmit)
	if err == nil {
		jsonWebRAUpdateRequestOnSubmit, _ := json.Marshal(dst.WebRAUpdateRequestOnSubmit)
		if string(jsonWebRAUpdateRequestOnSubmit) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnSubmit = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnSubmit)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnSubmit = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestSubmitRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSubmitRequest) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnSubmit != nil {
		return json.Marshal(&src.EstEnrollRequestOnSubmit)
	}

	if src.ScepEnrollRequestOnSubmit != nil {
		return json.Marshal(&src.ScepEnrollRequestOnSubmit)
	}

	if src.WebRAEnrollRequestOnSubmit != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnSubmit)
	}

	if src.WebRAImportRequestOnSubmit != nil {
		return json.Marshal(&src.WebRAImportRequestOnSubmit)
	}

	if src.WebRAMigrateRequestOnSubmit != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnSubmit)
	}

	if src.WebRARecoverRequestOnSubmit != nil {
		return json.Marshal(&src.WebRARecoverRequestOnSubmit)
	}

	if src.WebRARenewRequestOnSubmit != nil {
		return json.Marshal(&src.WebRARenewRequestOnSubmit)
	}

	if src.WebRARevokeRequestOnSubmit != nil {
		return json.Marshal(&src.WebRARevokeRequestOnSubmit)
	}

	if src.WebRAUpdateRequestOnSubmit != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnSubmit)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSubmitRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnSubmit != nil {
		return obj.EstEnrollRequestOnSubmit
	}

	if obj.ScepEnrollRequestOnSubmit != nil {
		return obj.ScepEnrollRequestOnSubmit
	}

	if obj.WebRAEnrollRequestOnSubmit != nil {
		return obj.WebRAEnrollRequestOnSubmit
	}

	if obj.WebRAImportRequestOnSubmit != nil {
		return obj.WebRAImportRequestOnSubmit
	}

	if obj.WebRAMigrateRequestOnSubmit != nil {
		return obj.WebRAMigrateRequestOnSubmit
	}

	if obj.WebRARecoverRequestOnSubmit != nil {
		return obj.WebRARecoverRequestOnSubmit
	}

	if obj.WebRARenewRequestOnSubmit != nil {
		return obj.WebRARenewRequestOnSubmit
	}

	if obj.WebRARevokeRequestOnSubmit != nil {
		return obj.WebRARevokeRequestOnSubmit
	}

	if obj.WebRAUpdateRequestOnSubmit != nil {
		return obj.WebRAUpdateRequestOnSubmit
	}

	// all schemas are nil
	return nil
}

type NullableRequestSubmitRequest struct {
	value *RequestSubmitRequest
	isSet bool
}

func (v NullableRequestSubmitRequest) Get() *RequestSubmitRequest {
	return v.value
}

func (v *NullableRequestSubmitRequest) Set(val *RequestSubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmitRequest(val *RequestSubmitRequest) *NullableRequestSubmitRequest {
	return &NullableRequestSubmitRequest{value: val, isSet: true}
}

func (v NullableRequestSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
