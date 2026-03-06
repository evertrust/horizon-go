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
)

// TriggerUpdate200Response - struct for TriggerUpdate200Response
type TriggerUpdate200Response struct {
	AWSTriggerResponse           *AWSTriggerResponse
	AzureKeyVaultTriggerResponse *AzureKeyVaultTriggerResponse
	EmailNotificationResponse    *EmailNotificationResponse
	F5AS3TriggerResponse         *F5AS3TriggerResponse
	F5ClientTriggerResponse      *F5ClientTriggerResponse
	GCMTriggerResponse           *GCMTriggerResponse
	IntunePKCSTriggerResponse    *IntunePKCSTriggerResponse
	LDAPTriggerResponse          *LDAPTriggerResponse
	RESTResponse                 *RESTResponse
	WebhookNotificationResponse  *WebhookNotificationResponse
}

// AWSTriggerResponseAsTriggerUpdate200Response is a convenience function that returns AWSTriggerResponse wrapped in TriggerUpdate200Response
func AWSTriggerResponseAsTriggerUpdate200Response(v *AWSTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		AWSTriggerResponse: v,
	}
}

// AzureKeyVaultTriggerResponseAsTriggerUpdate200Response is a convenience function that returns AzureKeyVaultTriggerResponse wrapped in TriggerUpdate200Response
func AzureKeyVaultTriggerResponseAsTriggerUpdate200Response(v *AzureKeyVaultTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		AzureKeyVaultTriggerResponse: v,
	}
}

// EmailNotificationResponseAsTriggerUpdate200Response is a convenience function that returns EmailNotificationResponse wrapped in TriggerUpdate200Response
func EmailNotificationResponseAsTriggerUpdate200Response(v *EmailNotificationResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		EmailNotificationResponse: v,
	}
}

// F5AS3TriggerResponseAsTriggerUpdate200Response is a convenience function that returns F5AS3TriggerResponse wrapped in TriggerUpdate200Response
func F5AS3TriggerResponseAsTriggerUpdate200Response(v *F5AS3TriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		F5AS3TriggerResponse: v,
	}
}

// F5ClientTriggerResponseAsTriggerUpdate200Response is a convenience function that returns F5ClientTriggerResponse wrapped in TriggerUpdate200Response
func F5ClientTriggerResponseAsTriggerUpdate200Response(v *F5ClientTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		F5ClientTriggerResponse: v,
	}
}

// GCMTriggerResponseAsTriggerUpdate200Response is a convenience function that returns GCMTriggerResponse wrapped in TriggerUpdate200Response
func GCMTriggerResponseAsTriggerUpdate200Response(v *GCMTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		GCMTriggerResponse: v,
	}
}

// IntunePKCSTriggerResponseAsTriggerUpdate200Response is a convenience function that returns IntunePKCSTriggerResponse wrapped in TriggerUpdate200Response
func IntunePKCSTriggerResponseAsTriggerUpdate200Response(v *IntunePKCSTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		IntunePKCSTriggerResponse: v,
	}
}

// LDAPTriggerResponseAsTriggerUpdate200Response is a convenience function that returns LDAPTriggerResponse wrapped in TriggerUpdate200Response
func LDAPTriggerResponseAsTriggerUpdate200Response(v *LDAPTriggerResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		LDAPTriggerResponse: v,
	}
}

// RESTResponseAsTriggerUpdate200Response is a convenience function that returns RESTResponse wrapped in TriggerUpdate200Response
func RESTResponseAsTriggerUpdate200Response(v *RESTResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		RESTResponse: v,
	}
}

// WebhookNotificationResponseAsTriggerUpdate200Response is a convenience function that returns WebhookNotificationResponse wrapped in TriggerUpdate200Response
func WebhookNotificationResponseAsTriggerUpdate200Response(v *WebhookNotificationResponse) TriggerUpdate200Response {
	return TriggerUpdate200Response{
		WebhookNotificationResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TriggerUpdate200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSTriggerResponse
	err = json.Unmarshal(data, &dst.AWSTriggerResponse)
	if err == nil {
		jsonAWSTriggerResponse, _ := json.Marshal(dst.AWSTriggerResponse)
		if string(jsonAWSTriggerResponse) == "{}" { // empty struct
			dst.AWSTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.AWSTriggerResponse = nil
	}

	// try to unmarshal data into AzureKeyVaultTriggerResponse
	err = json.Unmarshal(data, &dst.AzureKeyVaultTriggerResponse)
	if err == nil {
		jsonAzureKeyVaultTriggerResponse, _ := json.Marshal(dst.AzureKeyVaultTriggerResponse)
		if string(jsonAzureKeyVaultTriggerResponse) == "{}" { // empty struct
			dst.AzureKeyVaultTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.AzureKeyVaultTriggerResponse = nil
	}

	// try to unmarshal data into EmailNotificationResponse
	err = json.Unmarshal(data, &dst.EmailNotificationResponse)
	if err == nil {
		jsonEmailNotificationResponse, _ := json.Marshal(dst.EmailNotificationResponse)
		if string(jsonEmailNotificationResponse) == "{}" { // empty struct
			dst.EmailNotificationResponse = nil
		} else {
			match++
		}
	} else {
		dst.EmailNotificationResponse = nil
	}

	// try to unmarshal data into F5AS3TriggerResponse
	err = json.Unmarshal(data, &dst.F5AS3TriggerResponse)
	if err == nil {
		jsonF5AS3TriggerResponse, _ := json.Marshal(dst.F5AS3TriggerResponse)
		if string(jsonF5AS3TriggerResponse) == "{}" { // empty struct
			dst.F5AS3TriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.F5AS3TriggerResponse = nil
	}

	// try to unmarshal data into F5ClientTriggerResponse
	err = json.Unmarshal(data, &dst.F5ClientTriggerResponse)
	if err == nil {
		jsonF5ClientTriggerResponse, _ := json.Marshal(dst.F5ClientTriggerResponse)
		if string(jsonF5ClientTriggerResponse) == "{}" { // empty struct
			dst.F5ClientTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.F5ClientTriggerResponse = nil
	}

	// try to unmarshal data into GCMTriggerResponse
	err = json.Unmarshal(data, &dst.GCMTriggerResponse)
	if err == nil {
		jsonGCMTriggerResponse, _ := json.Marshal(dst.GCMTriggerResponse)
		if string(jsonGCMTriggerResponse) == "{}" { // empty struct
			dst.GCMTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.GCMTriggerResponse = nil
	}

	// try to unmarshal data into IntunePKCSTriggerResponse
	err = json.Unmarshal(data, &dst.IntunePKCSTriggerResponse)
	if err == nil {
		jsonIntunePKCSTriggerResponse, _ := json.Marshal(dst.IntunePKCSTriggerResponse)
		if string(jsonIntunePKCSTriggerResponse) == "{}" { // empty struct
			dst.IntunePKCSTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.IntunePKCSTriggerResponse = nil
	}

	// try to unmarshal data into LDAPTriggerResponse
	err = json.Unmarshal(data, &dst.LDAPTriggerResponse)
	if err == nil {
		jsonLDAPTriggerResponse, _ := json.Marshal(dst.LDAPTriggerResponse)
		if string(jsonLDAPTriggerResponse) == "{}" { // empty struct
			dst.LDAPTriggerResponse = nil
		} else {
			match++
		}
	} else {
		dst.LDAPTriggerResponse = nil
	}

	// try to unmarshal data into RESTResponse
	err = json.Unmarshal(data, &dst.RESTResponse)
	if err == nil {
		jsonRESTResponse, _ := json.Marshal(dst.RESTResponse)
		if string(jsonRESTResponse) == "{}" { // empty struct
			dst.RESTResponse = nil
		} else {
			match++
		}
	} else {
		dst.RESTResponse = nil
	}

	// try to unmarshal data into WebhookNotificationResponse
	err = json.Unmarshal(data, &dst.WebhookNotificationResponse)
	if err == nil {
		jsonWebhookNotificationResponse, _ := json.Marshal(dst.WebhookNotificationResponse)
		if string(jsonWebhookNotificationResponse) == "{}" { // empty struct
			dst.WebhookNotificationResponse = nil
		} else {
			match++
		}
	} else {
		dst.WebhookNotificationResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TriggerUpdate200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TriggerUpdate200Response) MarshalJSON() ([]byte, error) {
	if src.AWSTriggerResponse != nil {
		return json.Marshal(&src.AWSTriggerResponse)
	}

	if src.AzureKeyVaultTriggerResponse != nil {
		return json.Marshal(&src.AzureKeyVaultTriggerResponse)
	}

	if src.EmailNotificationResponse != nil {
		return json.Marshal(&src.EmailNotificationResponse)
	}

	if src.F5AS3TriggerResponse != nil {
		return json.Marshal(&src.F5AS3TriggerResponse)
	}

	if src.F5ClientTriggerResponse != nil {
		return json.Marshal(&src.F5ClientTriggerResponse)
	}

	if src.GCMTriggerResponse != nil {
		return json.Marshal(&src.GCMTriggerResponse)
	}

	if src.IntunePKCSTriggerResponse != nil {
		return json.Marshal(&src.IntunePKCSTriggerResponse)
	}

	if src.LDAPTriggerResponse != nil {
		return json.Marshal(&src.LDAPTriggerResponse)
	}

	if src.RESTResponse != nil {
		return json.Marshal(&src.RESTResponse)
	}

	if src.WebhookNotificationResponse != nil {
		return json.Marshal(&src.WebhookNotificationResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TriggerUpdate200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSTriggerResponse != nil {
		return obj.AWSTriggerResponse
	}

	if obj.AzureKeyVaultTriggerResponse != nil {
		return obj.AzureKeyVaultTriggerResponse
	}

	if obj.EmailNotificationResponse != nil {
		return obj.EmailNotificationResponse
	}

	if obj.F5AS3TriggerResponse != nil {
		return obj.F5AS3TriggerResponse
	}

	if obj.F5ClientTriggerResponse != nil {
		return obj.F5ClientTriggerResponse
	}

	if obj.GCMTriggerResponse != nil {
		return obj.GCMTriggerResponse
	}

	if obj.IntunePKCSTriggerResponse != nil {
		return obj.IntunePKCSTriggerResponse
	}

	if obj.LDAPTriggerResponse != nil {
		return obj.LDAPTriggerResponse
	}

	if obj.RESTResponse != nil {
		return obj.RESTResponse
	}

	if obj.WebhookNotificationResponse != nil {
		return obj.WebhookNotificationResponse
	}

	// all schemas are nil
	return nil
}

type NullableTriggerUpdate200Response struct {
	value *TriggerUpdate200Response
	isSet bool
}

func (v NullableTriggerUpdate200Response) Get() *TriggerUpdate200Response {
	return v.value
}

func (v *NullableTriggerUpdate200Response) Set(val *TriggerUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerUpdate200Response(val *TriggerUpdate200Response) *NullableTriggerUpdate200Response {
	return &NullableTriggerUpdate200Response{value: val, isSet: true}
}

func (v NullableTriggerUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
