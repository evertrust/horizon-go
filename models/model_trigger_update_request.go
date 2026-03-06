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

// TriggerUpdateRequest - struct for TriggerUpdateRequest
type TriggerUpdateRequest struct {
	AWSTrigger           *AWSTrigger
	AzureKeyVaultTrigger *AzureKeyVaultTrigger
	EmailNotification    *EmailNotification
	F5AS3Trigger         *F5AS3Trigger
	F5ClientTrigger      *F5ClientTrigger
	GCMTrigger           *GCMTrigger
	IntunePKCSTrigger    *IntunePKCSTrigger
	LDAPTrigger          *LDAPTrigger
	REST                 *REST
	WebhookNotification  *WebhookNotification
}

// AWSTriggerAsTriggerUpdateRequest is a convenience function that returns AWSTrigger wrapped in TriggerUpdateRequest
func AWSTriggerAsTriggerUpdateRequest(v *AWSTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		AWSTrigger: v,
	}
}

// AzureKeyVaultTriggerAsTriggerUpdateRequest is a convenience function that returns AzureKeyVaultTrigger wrapped in TriggerUpdateRequest
func AzureKeyVaultTriggerAsTriggerUpdateRequest(v *AzureKeyVaultTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		AzureKeyVaultTrigger: v,
	}
}

// EmailNotificationAsTriggerUpdateRequest is a convenience function that returns EmailNotification wrapped in TriggerUpdateRequest
func EmailNotificationAsTriggerUpdateRequest(v *EmailNotification) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		EmailNotification: v,
	}
}

// F5AS3TriggerAsTriggerUpdateRequest is a convenience function that returns F5AS3Trigger wrapped in TriggerUpdateRequest
func F5AS3TriggerAsTriggerUpdateRequest(v *F5AS3Trigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		F5AS3Trigger: v,
	}
}

// F5ClientTriggerAsTriggerUpdateRequest is a convenience function that returns F5ClientTrigger wrapped in TriggerUpdateRequest
func F5ClientTriggerAsTriggerUpdateRequest(v *F5ClientTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		F5ClientTrigger: v,
	}
}

// GCMTriggerAsTriggerUpdateRequest is a convenience function that returns GCMTrigger wrapped in TriggerUpdateRequest
func GCMTriggerAsTriggerUpdateRequest(v *GCMTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		GCMTrigger: v,
	}
}

// IntunePKCSTriggerAsTriggerUpdateRequest is a convenience function that returns IntunePKCSTrigger wrapped in TriggerUpdateRequest
func IntunePKCSTriggerAsTriggerUpdateRequest(v *IntunePKCSTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		IntunePKCSTrigger: v,
	}
}

// LDAPTriggerAsTriggerUpdateRequest is a convenience function that returns LDAPTrigger wrapped in TriggerUpdateRequest
func LDAPTriggerAsTriggerUpdateRequest(v *LDAPTrigger) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		LDAPTrigger: v,
	}
}

// RESTAsTriggerUpdateRequest is a convenience function that returns REST wrapped in TriggerUpdateRequest
func RESTAsTriggerUpdateRequest(v *REST) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		REST: v,
	}
}

// WebhookNotificationAsTriggerUpdateRequest is a convenience function that returns WebhookNotification wrapped in TriggerUpdateRequest
func WebhookNotificationAsTriggerUpdateRequest(v *WebhookNotification) TriggerUpdateRequest {
	return TriggerUpdateRequest{
		WebhookNotification: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TriggerUpdateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSTrigger
	err = json.Unmarshal(data, &dst.AWSTrigger)
	if err == nil {
		jsonAWSTrigger, _ := json.Marshal(dst.AWSTrigger)
		if string(jsonAWSTrigger) == "{}" { // empty struct
			dst.AWSTrigger = nil
		} else {
			match++
		}
	} else {
		dst.AWSTrigger = nil
	}

	// try to unmarshal data into AzureKeyVaultTrigger
	err = json.Unmarshal(data, &dst.AzureKeyVaultTrigger)
	if err == nil {
		jsonAzureKeyVaultTrigger, _ := json.Marshal(dst.AzureKeyVaultTrigger)
		if string(jsonAzureKeyVaultTrigger) == "{}" { // empty struct
			dst.AzureKeyVaultTrigger = nil
		} else {
			match++
		}
	} else {
		dst.AzureKeyVaultTrigger = nil
	}

	// try to unmarshal data into EmailNotification
	err = json.Unmarshal(data, &dst.EmailNotification)
	if err == nil {
		jsonEmailNotification, _ := json.Marshal(dst.EmailNotification)
		if string(jsonEmailNotification) == "{}" { // empty struct
			dst.EmailNotification = nil
		} else {
			match++
		}
	} else {
		dst.EmailNotification = nil
	}

	// try to unmarshal data into F5AS3Trigger
	err = json.Unmarshal(data, &dst.F5AS3Trigger)
	if err == nil {
		jsonF5AS3Trigger, _ := json.Marshal(dst.F5AS3Trigger)
		if string(jsonF5AS3Trigger) == "{}" { // empty struct
			dst.F5AS3Trigger = nil
		} else {
			match++
		}
	} else {
		dst.F5AS3Trigger = nil
	}

	// try to unmarshal data into F5ClientTrigger
	err = json.Unmarshal(data, &dst.F5ClientTrigger)
	if err == nil {
		jsonF5ClientTrigger, _ := json.Marshal(dst.F5ClientTrigger)
		if string(jsonF5ClientTrigger) == "{}" { // empty struct
			dst.F5ClientTrigger = nil
		} else {
			match++
		}
	} else {
		dst.F5ClientTrigger = nil
	}

	// try to unmarshal data into GCMTrigger
	err = json.Unmarshal(data, &dst.GCMTrigger)
	if err == nil {
		jsonGCMTrigger, _ := json.Marshal(dst.GCMTrigger)
		if string(jsonGCMTrigger) == "{}" { // empty struct
			dst.GCMTrigger = nil
		} else {
			match++
		}
	} else {
		dst.GCMTrigger = nil
	}

	// try to unmarshal data into IntunePKCSTrigger
	err = json.Unmarshal(data, &dst.IntunePKCSTrigger)
	if err == nil {
		jsonIntunePKCSTrigger, _ := json.Marshal(dst.IntunePKCSTrigger)
		if string(jsonIntunePKCSTrigger) == "{}" { // empty struct
			dst.IntunePKCSTrigger = nil
		} else {
			match++
		}
	} else {
		dst.IntunePKCSTrigger = nil
	}

	// try to unmarshal data into LDAPTrigger
	err = json.Unmarshal(data, &dst.LDAPTrigger)
	if err == nil {
		jsonLDAPTrigger, _ := json.Marshal(dst.LDAPTrigger)
		if string(jsonLDAPTrigger) == "{}" { // empty struct
			dst.LDAPTrigger = nil
		} else {
			match++
		}
	} else {
		dst.LDAPTrigger = nil
	}

	// try to unmarshal data into REST
	err = json.Unmarshal(data, &dst.REST)
	if err == nil {
		jsonREST, _ := json.Marshal(dst.REST)
		if string(jsonREST) == "{}" { // empty struct
			dst.REST = nil
		} else {
			match++
		}
	} else {
		dst.REST = nil
	}

	// try to unmarshal data into WebhookNotification
	err = json.Unmarshal(data, &dst.WebhookNotification)
	if err == nil {
		jsonWebhookNotification, _ := json.Marshal(dst.WebhookNotification)
		if string(jsonWebhookNotification) == "{}" { // empty struct
			dst.WebhookNotification = nil
		} else {
			match++
		}
	} else {
		dst.WebhookNotification = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TriggerUpdateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TriggerUpdateRequest) MarshalJSON() ([]byte, error) {
	if src.AWSTrigger != nil {
		return json.Marshal(&src.AWSTrigger)
	}

	if src.AzureKeyVaultTrigger != nil {
		return json.Marshal(&src.AzureKeyVaultTrigger)
	}

	if src.EmailNotification != nil {
		return json.Marshal(&src.EmailNotification)
	}

	if src.F5AS3Trigger != nil {
		return json.Marshal(&src.F5AS3Trigger)
	}

	if src.F5ClientTrigger != nil {
		return json.Marshal(&src.F5ClientTrigger)
	}

	if src.GCMTrigger != nil {
		return json.Marshal(&src.GCMTrigger)
	}

	if src.IntunePKCSTrigger != nil {
		return json.Marshal(&src.IntunePKCSTrigger)
	}

	if src.LDAPTrigger != nil {
		return json.Marshal(&src.LDAPTrigger)
	}

	if src.REST != nil {
		return json.Marshal(&src.REST)
	}

	if src.WebhookNotification != nil {
		return json.Marshal(&src.WebhookNotification)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TriggerUpdateRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSTrigger != nil {
		return obj.AWSTrigger
	}

	if obj.AzureKeyVaultTrigger != nil {
		return obj.AzureKeyVaultTrigger
	}

	if obj.EmailNotification != nil {
		return obj.EmailNotification
	}

	if obj.F5AS3Trigger != nil {
		return obj.F5AS3Trigger
	}

	if obj.F5ClientTrigger != nil {
		return obj.F5ClientTrigger
	}

	if obj.GCMTrigger != nil {
		return obj.GCMTrigger
	}

	if obj.IntunePKCSTrigger != nil {
		return obj.IntunePKCSTrigger
	}

	if obj.LDAPTrigger != nil {
		return obj.LDAPTrigger
	}

	if obj.REST != nil {
		return obj.REST
	}

	if obj.WebhookNotification != nil {
		return obj.WebhookNotification
	}

	// all schemas are nil
	return nil
}

type NullableTriggerUpdateRequest struct {
	value *TriggerUpdateRequest
	isSet bool
}

func (v NullableTriggerUpdateRequest) Get() *TriggerUpdateRequest {
	return v.value
}

func (v *NullableTriggerUpdateRequest) Set(val *TriggerUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerUpdateRequest(val *TriggerUpdateRequest) *NullableTriggerUpdateRequest {
	return &NullableTriggerUpdateRequest{value: val, isSet: true}
}

func (v NullableTriggerUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
