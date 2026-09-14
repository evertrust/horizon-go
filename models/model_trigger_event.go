/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// TriggerEvent the model 'TriggerEvent'
type TriggerEvent string

// List of TriggerEvent
const (
	TRIGGEREVENT_ON_APPROVE_ENROLL         TriggerEvent = "on_approve_enroll"
	TRIGGEREVENT_ON_APPROVE_IMPORT         TriggerEvent = "on_approve_import"
	TRIGGEREVENT_ON_APPROVE_MIGRATE        TriggerEvent = "on_approve_migrate"
	TRIGGEREVENT_ON_APPROVE_RECOVER        TriggerEvent = "on_approve_recover"
	TRIGGEREVENT_ON_APPROVE_RENEW          TriggerEvent = "on_approve_renew"
	TRIGGEREVENT_ON_APPROVE_REVOKE         TriggerEvent = "on_approve_revoke"
	TRIGGEREVENT_ON_APPROVE_UPDATE         TriggerEvent = "on_approve_update"
	TRIGGEREVENT_ON_CANCEL_ENROLL          TriggerEvent = "on_cancel_enroll"
	TRIGGEREVENT_ON_CANCEL_IMPORT          TriggerEvent = "on_cancel_import"
	TRIGGEREVENT_ON_CANCEL_MIGRATE         TriggerEvent = "on_cancel_migrate"
	TRIGGEREVENT_ON_CANCEL_RECOVER         TriggerEvent = "on_cancel_recover"
	TRIGGEREVENT_ON_CANCEL_RENEW           TriggerEvent = "on_cancel_renew"
	TRIGGEREVENT_ON_CANCEL_REVOKE          TriggerEvent = "on_cancel_revoke"
	TRIGGEREVENT_ON_CANCEL_UPDATE          TriggerEvent = "on_cancel_update"
	TRIGGEREVENT_ON_CREDENTIALS_EXPIRATION TriggerEvent = "on_credentials_expiration"
	TRIGGEREVENT_ON_DCV_LICENSE_USAGE      TriggerEvent = "on_dcv_license_usage"
	TRIGGEREVENT_ON_DCV_POLICY_END         TriggerEvent = "on_dcv_policy_end"
	TRIGGEREVENT_ON_DCV_POLICY_START       TriggerEvent = "on_dcv_policy_start"
	TRIGGEREVENT_ON_DCV_VALIDATION_FAILURE TriggerEvent = "on_dcv_validation_failure"
	TRIGGEREVENT_ON_DCV_VALIDATION_RETRY   TriggerEvent = "on_dcv_validation_retry"
	TRIGGEREVENT_ON_DCV_VALIDATION_SUCCESS TriggerEvent = "on_dcv_validation_success"
	TRIGGEREVENT_ON_DENY_ENROLL            TriggerEvent = "on_deny_enroll"
	TRIGGEREVENT_ON_DENY_IMPORT            TriggerEvent = "on_deny_import"
	TRIGGEREVENT_ON_DENY_MIGRATE           TriggerEvent = "on_deny_migrate"
	TRIGGEREVENT_ON_DENY_RECOVER           TriggerEvent = "on_deny_recover"
	TRIGGEREVENT_ON_DENY_RENEW             TriggerEvent = "on_deny_renew"
	TRIGGEREVENT_ON_DENY_REVOKE            TriggerEvent = "on_deny_revoke"
	TRIGGEREVENT_ON_DENY_UPDATE            TriggerEvent = "on_deny_update"
	TRIGGEREVENT_ON_ENROLL                 TriggerEvent = "on_enroll"
	TRIGGEREVENT_ON_EXPIRE                 TriggerEvent = "on_expire"
	TRIGGEREVENT_ON_FAILURE_ENROLL         TriggerEvent = "on_failure_enroll"
	TRIGGEREVENT_ON_FAILURE_RENEW          TriggerEvent = "on_failure_renew"
	TRIGGEREVENT_ON_FAILURE_REVOKE         TriggerEvent = "on_failure_revoke"
	TRIGGEREVENT_ON_IMPORT                 TriggerEvent = "on_import"
	TRIGGEREVENT_ON_IN_PROGRESS_ENROLL     TriggerEvent = "on_in_progress_enroll"
	TRIGGEREVENT_ON_IN_PROGRESS_RENEW      TriggerEvent = "on_in_progress_renew"
	TRIGGEREVENT_ON_LICENSE_EXPIRATION     TriggerEvent = "on_license_expiration"
	TRIGGEREVENT_ON_LICENSE_USAGE          TriggerEvent = "on_license_usage"
	TRIGGEREVENT_ON_MIGRATE                TriggerEvent = "on_migrate"
	TRIGGEREVENT_ON_PENDING_ENROLL         TriggerEvent = "on_pending_enroll"
	TRIGGEREVENT_ON_PENDING_IMPORT         TriggerEvent = "on_pending_import"
	TRIGGEREVENT_ON_PENDING_MIGRATE        TriggerEvent = "on_pending_migrate"
	TRIGGEREVENT_ON_PENDING_RECOVER        TriggerEvent = "on_pending_recover"
	TRIGGEREVENT_ON_PENDING_RENEW          TriggerEvent = "on_pending_renew"
	TRIGGEREVENT_ON_PENDING_REVOKE         TriggerEvent = "on_pending_revoke"
	TRIGGEREVENT_ON_PENDING_UPDATE         TriggerEvent = "on_pending_update"
	TRIGGEREVENT_ON_RECOVER                TriggerEvent = "on_recover"
	TRIGGEREVENT_ON_RENEW                  TriggerEvent = "on_renew"
	TRIGGEREVENT_ON_REVOKE                 TriggerEvent = "on_revoke"
	TRIGGEREVENT_ON_SUBMIT_ENROLL          TriggerEvent = "on_submit_enroll"
	TRIGGEREVENT_ON_SUBMIT_IMPORT          TriggerEvent = "on_submit_import"
	TRIGGEREVENT_ON_SUBMIT_MIGRATE         TriggerEvent = "on_submit_migrate"
	TRIGGEREVENT_ON_SUBMIT_RECOVER         TriggerEvent = "on_submit_recover"
	TRIGGEREVENT_ON_SUBMIT_RENEW           TriggerEvent = "on_submit_renew"
	TRIGGEREVENT_ON_SUBMIT_REVOKE          TriggerEvent = "on_submit_revoke"
	TRIGGEREVENT_ON_SUBMIT_UPDATE          TriggerEvent = "on_submit_update"
	TRIGGEREVENT_ON_TEST                   TriggerEvent = "on_test"
	TRIGGEREVENT_ON_TRIGGER_ERROR          TriggerEvent = "on_trigger_error"
	TRIGGEREVENT_ON_UPDATE                 TriggerEvent = "on_update"
)

// All allowed values of TriggerEvent enum
var AllowedTriggerEventEnumValues = []TriggerEvent{
	"on_approve_enroll",
	"on_approve_import",
	"on_approve_migrate",
	"on_approve_recover",
	"on_approve_renew",
	"on_approve_revoke",
	"on_approve_update",
	"on_cancel_enroll",
	"on_cancel_import",
	"on_cancel_migrate",
	"on_cancel_recover",
	"on_cancel_renew",
	"on_cancel_revoke",
	"on_cancel_update",
	"on_credentials_expiration",
	"on_dcv_license_usage",
	"on_dcv_policy_end",
	"on_dcv_policy_start",
	"on_dcv_validation_failure",
	"on_dcv_validation_retry",
	"on_dcv_validation_success",
	"on_deny_enroll",
	"on_deny_import",
	"on_deny_migrate",
	"on_deny_recover",
	"on_deny_renew",
	"on_deny_revoke",
	"on_deny_update",
	"on_enroll",
	"on_expire",
	"on_failure_enroll",
	"on_failure_renew",
	"on_failure_revoke",
	"on_import",
	"on_in_progress_enroll",
	"on_in_progress_renew",
	"on_license_expiration",
	"on_license_usage",
	"on_migrate",
	"on_pending_enroll",
	"on_pending_import",
	"on_pending_migrate",
	"on_pending_recover",
	"on_pending_renew",
	"on_pending_revoke",
	"on_pending_update",
	"on_recover",
	"on_renew",
	"on_revoke",
	"on_submit_enroll",
	"on_submit_import",
	"on_submit_migrate",
	"on_submit_recover",
	"on_submit_renew",
	"on_submit_revoke",
	"on_submit_update",
	"on_test",
	"on_trigger_error",
	"on_update",
}

func (v *TriggerEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerEvent(value)
	for _, existing := range AllowedTriggerEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerEvent", value)
}

// NewTriggerEventFromValue returns a pointer to a valid TriggerEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerEventFromValue(v string) (*TriggerEvent, error) {
	ev := TriggerEvent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerEvent: valid values are %v", v, AllowedTriggerEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerEvent) IsValid() bool {
	for _, existing := range AllowedTriggerEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerEvent value
func (v TriggerEvent) Ptr() *TriggerEvent {
	return &v
}

type NullableTriggerEvent struct {
	value *TriggerEvent
	isSet bool
}

func (v NullableTriggerEvent) Get() *TriggerEvent {
	return v.value
}

func (v *NullableTriggerEvent) Set(val *TriggerEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerEvent(val *TriggerEvent) *NullableTriggerEvent {
	return &NullableTriggerEvent{value: val, isSet: true}
}

func (v NullableTriggerEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
