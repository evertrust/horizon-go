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

// RequestTemplateRequest - struct for RequestTemplateRequest
type RequestTemplateRequest struct {
	EstEnrollRequestOnTemplate    *EstEnrollRequestOnTemplate
	ScepEnrollRequestOnTemplate   *ScepEnrollRequestOnTemplate
	WebRAEnrollRequestOnTemplate  *WebRAEnrollRequestOnTemplate
	WebRAImportRequestOnTemplate  *WebRAImportRequestOnTemplate
	WebRAMigrateRequestOnTemplate *WebRAMigrateRequestOnTemplate
	WebRARecoverRequestOnTemplate *WebRARecoverRequestOnTemplate
	WebRARenewRequestOnTemplate   *WebRARenewRequestOnTemplate
	WebRARevokeRequestOnTemplate  *WebRARevokeRequestOnTemplate
	WebRAUpdateRequestOnTemplate  *WebRAUpdateRequestOnTemplate
}

// EstEnrollRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns EstEnrollRequestOnTemplate wrapped in RequestTemplateRequest
func EstEnrollRequestOnTemplateAsRequestTemplateRequest(v *EstEnrollRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		EstEnrollRequestOnTemplate: v,
	}
}

// ScepEnrollRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns ScepEnrollRequestOnTemplate wrapped in RequestTemplateRequest
func ScepEnrollRequestOnTemplateAsRequestTemplateRequest(v *ScepEnrollRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		ScepEnrollRequestOnTemplate: v,
	}
}

// WebRAEnrollRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRAEnrollRequestOnTemplate wrapped in RequestTemplateRequest
func WebRAEnrollRequestOnTemplateAsRequestTemplateRequest(v *WebRAEnrollRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRAEnrollRequestOnTemplate: v,
	}
}

// WebRAImportRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRAImportRequestOnTemplate wrapped in RequestTemplateRequest
func WebRAImportRequestOnTemplateAsRequestTemplateRequest(v *WebRAImportRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRAImportRequestOnTemplate: v,
	}
}

// WebRAMigrateRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRAMigrateRequestOnTemplate wrapped in RequestTemplateRequest
func WebRAMigrateRequestOnTemplateAsRequestTemplateRequest(v *WebRAMigrateRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRAMigrateRequestOnTemplate: v,
	}
}

// WebRARecoverRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRARecoverRequestOnTemplate wrapped in RequestTemplateRequest
func WebRARecoverRequestOnTemplateAsRequestTemplateRequest(v *WebRARecoverRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRARecoverRequestOnTemplate: v,
	}
}

// WebRARenewRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRARenewRequestOnTemplate wrapped in RequestTemplateRequest
func WebRARenewRequestOnTemplateAsRequestTemplateRequest(v *WebRARenewRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRARenewRequestOnTemplate: v,
	}
}

// WebRARevokeRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRARevokeRequestOnTemplate wrapped in RequestTemplateRequest
func WebRARevokeRequestOnTemplateAsRequestTemplateRequest(v *WebRARevokeRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRARevokeRequestOnTemplate: v,
	}
}

// WebRAUpdateRequestOnTemplateAsRequestTemplateRequest is a convenience function that returns WebRAUpdateRequestOnTemplate wrapped in RequestTemplateRequest
func WebRAUpdateRequestOnTemplateAsRequestTemplateRequest(v *WebRAUpdateRequestOnTemplate) RequestTemplateRequest {
	return RequestTemplateRequest{
		WebRAUpdateRequestOnTemplate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTemplateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EstEnrollRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnTemplate)
	if err == nil {
		jsonEstEnrollRequestOnTemplate, _ := json.Marshal(dst.EstEnrollRequestOnTemplate)
		if string(jsonEstEnrollRequestOnTemplate) == "{}" { // empty struct
			dst.EstEnrollRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.EstEnrollRequestOnTemplate)
			match++
		}
	} else {
		dst.EstEnrollRequestOnTemplate = nil
	}

	// try to unmarshal data into ScepEnrollRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnTemplate)
	if err == nil {
		jsonScepEnrollRequestOnTemplate, _ := json.Marshal(dst.ScepEnrollRequestOnTemplate)
		if string(jsonScepEnrollRequestOnTemplate) == "{}" { // empty struct
			dst.ScepEnrollRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.ScepEnrollRequestOnTemplate)
			match++
		}
	} else {
		dst.ScepEnrollRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRAEnrollRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnTemplate)
	if err == nil {
		jsonWebRAEnrollRequestOnTemplate, _ := json.Marshal(dst.WebRAEnrollRequestOnTemplate)
		if string(jsonWebRAEnrollRequestOnTemplate) == "{}" { // empty struct
			dst.WebRAEnrollRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRAEnrollRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRAEnrollRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRAImportRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnTemplate)
	if err == nil {
		jsonWebRAImportRequestOnTemplate, _ := json.Marshal(dst.WebRAImportRequestOnTemplate)
		if string(jsonWebRAImportRequestOnTemplate) == "{}" { // empty struct
			dst.WebRAImportRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRAImportRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRAImportRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRAMigrateRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnTemplate)
	if err == nil {
		jsonWebRAMigrateRequestOnTemplate, _ := json.Marshal(dst.WebRAMigrateRequestOnTemplate)
		if string(jsonWebRAMigrateRequestOnTemplate) == "{}" { // empty struct
			dst.WebRAMigrateRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRAMigrateRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRAMigrateRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRARecoverRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnTemplate)
	if err == nil {
		jsonWebRARecoverRequestOnTemplate, _ := json.Marshal(dst.WebRARecoverRequestOnTemplate)
		if string(jsonWebRARecoverRequestOnTemplate) == "{}" { // empty struct
			dst.WebRARecoverRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRARecoverRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRARecoverRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRARenewRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnTemplate)
	if err == nil {
		jsonWebRARenewRequestOnTemplate, _ := json.Marshal(dst.WebRARenewRequestOnTemplate)
		if string(jsonWebRARenewRequestOnTemplate) == "{}" { // empty struct
			dst.WebRARenewRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRARenewRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRARenewRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRARevokeRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnTemplate)
	if err == nil {
		jsonWebRARevokeRequestOnTemplate, _ := json.Marshal(dst.WebRARevokeRequestOnTemplate)
		if string(jsonWebRARevokeRequestOnTemplate) == "{}" { // empty struct
			dst.WebRARevokeRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRARevokeRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRARevokeRequestOnTemplate = nil
	}

	// try to unmarshal data into WebRAUpdateRequestOnTemplate
	err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnTemplate)
	if err == nil {
		jsonWebRAUpdateRequestOnTemplate, _ := json.Marshal(dst.WebRAUpdateRequestOnTemplate)
		if string(jsonWebRAUpdateRequestOnTemplate) == "{}" { // empty struct
			dst.WebRAUpdateRequestOnTemplate = nil
		} else {
			_ = validator.Validate(dst.WebRAUpdateRequestOnTemplate)
			match++
		}
	} else {
		dst.WebRAUpdateRequestOnTemplate = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestTemplateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTemplateRequest) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnTemplate != nil {
		return json.Marshal(&src.EstEnrollRequestOnTemplate)
	}

	if src.ScepEnrollRequestOnTemplate != nil {
		return json.Marshal(&src.ScepEnrollRequestOnTemplate)
	}

	if src.WebRAEnrollRequestOnTemplate != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnTemplate)
	}

	if src.WebRAImportRequestOnTemplate != nil {
		return json.Marshal(&src.WebRAImportRequestOnTemplate)
	}

	if src.WebRAMigrateRequestOnTemplate != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnTemplate)
	}

	if src.WebRARecoverRequestOnTemplate != nil {
		return json.Marshal(&src.WebRARecoverRequestOnTemplate)
	}

	if src.WebRARenewRequestOnTemplate != nil {
		return json.Marshal(&src.WebRARenewRequestOnTemplate)
	}

	if src.WebRARevokeRequestOnTemplate != nil {
		return json.Marshal(&src.WebRARevokeRequestOnTemplate)
	}

	if src.WebRAUpdateRequestOnTemplate != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnTemplate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTemplateRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnTemplate != nil {
		return obj.EstEnrollRequestOnTemplate
	}

	if obj.ScepEnrollRequestOnTemplate != nil {
		return obj.ScepEnrollRequestOnTemplate
	}

	if obj.WebRAEnrollRequestOnTemplate != nil {
		return obj.WebRAEnrollRequestOnTemplate
	}

	if obj.WebRAImportRequestOnTemplate != nil {
		return obj.WebRAImportRequestOnTemplate
	}

	if obj.WebRAMigrateRequestOnTemplate != nil {
		return obj.WebRAMigrateRequestOnTemplate
	}

	if obj.WebRARecoverRequestOnTemplate != nil {
		return obj.WebRARecoverRequestOnTemplate
	}

	if obj.WebRARenewRequestOnTemplate != nil {
		return obj.WebRARenewRequestOnTemplate
	}

	if obj.WebRARevokeRequestOnTemplate != nil {
		return obj.WebRARevokeRequestOnTemplate
	}

	if obj.WebRAUpdateRequestOnTemplate != nil {
		return obj.WebRAUpdateRequestOnTemplate
	}

	// all schemas are nil
	return nil
}

type NullableRequestTemplateRequest struct {
	value *RequestTemplateRequest
	isSet bool
}

func (v NullableRequestTemplateRequest) Get() *RequestTemplateRequest {
	return v.value
}

func (v *NullableRequestTemplateRequest) Set(val *RequestTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTemplateRequest(val *RequestTemplateRequest) *NullableRequestTemplateRequest {
	return &NullableRequestTemplateRequest{value: val, isSet: true}
}

func (v NullableRequestTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
