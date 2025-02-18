/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// EventCode the model 'EventCode'
type EventCode string

// List of EventCode
const (
	EVENTCODE_ACME_ACCOUNT_KEY_CHANGE EventCode = "ACME-ACCOUNT-KEY-CHANGE"
	EVENTCODE_ACME_ACCOUNT_REGISTER EventCode = "ACME-ACCOUNT-REGISTER"
	EVENTCODE_ACME_ACCOUNT_UPDATE EventCode = "ACME-ACCOUNT-UPDATE"
	EVENTCODE_ACME_AUTHORIZATION_DEACTIVATE EventCode = "ACME-AUTHORIZATION-DEACTIVATE"
	EVENTCODE_ACME_CHALLENGE_REQUEST_VERIFY EventCode = "ACME-CHALLENGE-REQUEST-VERIFY"
	EVENTCODE_ACME_CHALLENGE_VERIFY EventCode = "ACME-CHALLENGE-VERIFY"
	EVENTCODE_ACME_ORDER_CERTIFICATE EventCode = "ACME-ORDER-CERTIFICATE"
	EVENTCODE_ACME_ORDER_FINALIZE EventCode = "ACME-ORDER-FINALIZE"
	EVENTCODE_ACME_ORDER_NEW EventCode = "ACME-ORDER-NEW"
	EVENTCODE_ACME_ORDER_UPDATE EventCode = "ACME-ORDER-UPDATE"
	EVENTCODE_ACME_REVOKE EventCode = "ACME-REVOKE"
	EVENTCODE_ACTOR EventCode = "ACTOR"
	EVENTCODE_BOOTSTRAP_ADMINISTRATOR_ACCOUNT EventCode = "BOOTSTRAP-ADMINISTRATOR-ACCOUNT"
	EVENTCODE_BOOTSTRAP_ADMINISTRATOR_PRINCIPAL EventCode = "BOOTSTRAP-ADMINISTRATOR-PRINCIPAL"
	EVENTCODE_BOOTSTRAP_GRADING_POLICY EventCode = "BOOTSTRAP-GRADING-POLICY"
	EVENTCODE_BOOTSTRAP_GRADING_RULESET EventCode = "BOOTSTRAP-GRADING-RULESET"
	EVENTCODE_BOOTSTRAP_LOCAL_IDENTITY_PROVIDER EventCode = "BOOTSTRAP-LOCAL-IDENTITY-PROVIDER"
	EVENTCODE_CA_CERT_SYNC EventCode = "CA-CERT-SYNC"
	EVENTCODE_CA_CRL_UPDATE EventCode = "CA-CRL-UPDATE"
	EVENTCODE_CONF_ADD EventCode = "CONF-ADD"
	EVENTCODE_CONF_DELETE EventCode = "CONF-DELETE"
	EVENTCODE_CONF_UPDATE EventCode = "CONF-UPDATE"
	EVENTCODE_CRMP_AUTHENTICATION EventCode = "CRMP-AUTHENTICATION"
	EVENTCODE_CRMP_BAD_REQUEST EventCode = "CRMP-BAD-REQUEST"
	EVENTCODE_CRMP_ENROLL EventCode = "CRMP-ENROLL"
	EVENTCODE_CRMP_LIST EventCode = "CRMP-LIST"
	EVENTCODE_CRMP_PROFILE_PROPERTIES EventCode = "CRMP-PROFILE-PROPERTIES"
	EVENTCODE_CRMP_RECOVER EventCode = "CRMP-RECOVER"
	EVENTCODE_CRMP_RETRIEVE EventCode = "CRMP-RETRIEVE"
	EVENTCODE_CRMP_REVOKE EventCode = "CRMP-REVOKE"
	EVENTCODE_DISCOVERY_CAMPAIGN_FLUSH EventCode = "DISCOVERY-CAMPAIGN-FLUSH"
	EVENTCODE_EST_CACERTS EventCode = "EST-CACERTS"
	EVENTCODE_EST_REVOKE_ON_RENEW EventCode = "EST-REVOKE-ON-RENEW"
	EVENTCODE_EST_SIMPLE_ENROLL EventCode = "EST-SIMPLE-ENROLL"
	EVENTCODE_EST_SIMPLE_REENROLL EventCode = "EST-SIMPLE-REENROLL"
	EVENTCODE_GRADING_END EventCode = "GRADING-END"
	EVENTCODE_GRADING_ERROR EventCode = "GRADING-ERROR"
	EVENTCODE_GRADING_START EventCode = "GRADING-START"
	EVENTCODE_LICENSE_ERROR EventCode = "LICENSE-ERROR"
	EVENTCODE_LICENSE_LIMIT_REACHED EventCode = "LICENSE-LIMIT-REACHED"
	EVENTCODE_LIFECYCLE_ENROLL EventCode = "LIFECYCLE-ENROLL"
	EVENTCODE_LIFECYCLE_ESCROW EventCode = "LIFECYCLE-ESCROW"
	EVENTCODE_LIFECYCLE_MAX_CERT_PER_HOLDER EventCode = "LIFECYCLE-MAX-CERT-PER-HOLDER"
	EVENTCODE_LIFECYCLE_MIGRATE EventCode = "LIFECYCLE-MIGRATE"
	EVENTCODE_LIFECYCLE_RECOVER EventCode = "LIFECYCLE-RECOVER"
	EVENTCODE_LIFECYCLE_RENEW EventCode = "LIFECYCLE-RENEW"
	EVENTCODE_LIFECYCLE_REVOKE EventCode = "LIFECYCLE-REVOKE"
	EVENTCODE_LIFECYCLE_UPDATE EventCode = "LIFECYCLE-UPDATE"
	EVENTCODE_PKI_CONNECTOR EventCode = "PKI-CONNECTOR"
	EVENTCODE_REQUEST_APPROVE EventCode = "REQUEST-APPROVE"
	EVENTCODE_REQUEST_CANCEL EventCode = "REQUEST-CANCEL"
	EVENTCODE_REQUEST_DENY EventCode = "REQUEST-DENY"
	EVENTCODE_REQUEST_SUBMIT EventCode = "REQUEST-SUBMIT"
	EVENTCODE_REQUEST_TEMPLATE EventCode = "REQUEST-TEMPLATE"
	EVENTCODE_SCEP_ENROLL EventCode = "SCEP-ENROLL"
	EVENTCODE_SCEP_GET_CA_CERT EventCode = "SCEP-GET-CA-CERT"
	EVENTCODE_SCEP_GET_CERT_INITIAL EventCode = "SCEP-GET-CERT-INITIAL"
	EVENTCODE_SCEP_NDES_EMULATION EventCode = "SCEP-NDES-EMULATION"
	EVENTCODE_SCEP_PKI_CLIENT EventCode = "SCEP-PKI-CLIENT"
	EVENTCODE_SCEP_PKI_OPERATION EventCode = "SCEP-PKI-OPERATION"
	EVENTCODE_SCEP_RENEW EventCode = "SCEP-RENEW"
	EVENTCODE_SCEP_REVOKE_ON_RENEW EventCode = "SCEP-REVOKE-ON-RENEW"
	EVENTCODE_SCHEDULED_TASK_COMPLETE EventCode = "SCHEDULED-TASK-COMPLETE"
	EVENTCODE_SCHEDULED_TASK_RUN EventCode = "SCHEDULED-TASK-RUN"
	EVENTCODE_SEC_AUTHENTICATION EventCode = "SEC-AUTHENTICATION"
	EVENTCODE_SEC_AUTHORIZATION_ADD EventCode = "SEC-AUTHORIZATION-ADD"
	EVENTCODE_SEC_AUTHORIZATION_DELETE EventCode = "SEC-AUTHORIZATION-DELETE"
	EVENTCODE_SEC_AUTHORIZATION_UPDATE EventCode = "SEC-AUTHORIZATION-UPDATE"
	EVENTCODE_SEC_IDENTITY_PROVIDER_ADD EventCode = "SEC-IDENTITY-PROVIDER-ADD"
	EVENTCODE_SEC_IDENTITY_PROVIDER_DELETE EventCode = "SEC-IDENTITY-PROVIDER-DELETE"
	EVENTCODE_SEC_IDENTITY_PROVIDER_UPDATE EventCode = "SEC-IDENTITY-PROVIDER-UPDATE"
	EVENTCODE_SEC_LOCAL_IDENTITY_ADD EventCode = "SEC-LOCAL-IDENTITY-ADD"
	EVENTCODE_SEC_LOCAL_IDENTITY_DELETE EventCode = "SEC-LOCAL-IDENTITY-DELETE"
	EVENTCODE_SEC_LOCAL_IDENTITY_RESET EventCode = "SEC-LOCAL-IDENTITY-RESET"
	EVENTCODE_SEC_LOCAL_IDENTITY_UPDATE EventCode = "SEC-LOCAL-IDENTITY-UPDATE"
	EVENTCODE_SEC_PASSWORD_POLICY_ADD EventCode = "SEC-PASSWORD-POLICY-ADD"
	EVENTCODE_SEC_PASSWORD_POLICY_DELETE EventCode = "SEC-PASSWORD-POLICY-DELETE"
	EVENTCODE_SEC_PASSWORD_POLICY_UPDATE EventCode = "SEC-PASSWORD-POLICY-UPDATE"
	EVENTCODE_SEC_ROLE_ADD EventCode = "SEC-ROLE-ADD"
	EVENTCODE_SEC_ROLE_DELETE EventCode = "SEC-ROLE-DELETE"
	EVENTCODE_SEC_ROLE_UPDATE EventCode = "SEC-ROLE-UPDATE"
	EVENTCODE_SEC_TEAM_ADD EventCode = "SEC-TEAM-ADD"
	EVENTCODE_SEC_TEAM_DELETE EventCode = "SEC-TEAM-DELETE"
	EVENTCODE_SEC_TEAM_SWITCH EventCode = "SEC-TEAM-SWITCH"
	EVENTCODE_SEC_TEAM_UPDATE EventCode = "SEC-TEAM-UPDATE"
	EVENTCODE_SERVICE_START EventCode = "SERVICE-START"
	EVENTCODE_SERVICE_STOP EventCode = "SERVICE-STOP"
	EVENTCODE_SUPERVISOR EventCode = "SUPERVISOR"
	EVENTCODE_SYNC_ENROLL EventCode = "SYNC-ENROLL"
	EVENTCODE_SYNC_RENEW EventCode = "SYNC-RENEW"
	EVENTCODE_SYNC_REVOKE EventCode = "SYNC-REVOKE"
	EVENTCODE_TEAM_SWITCH EventCode = "TEAM-SWITCH"
	EVENTCODE_TRIGGER_DELETE EventCode = "TRIGGER-DELETE"
	EVENTCODE_TRIGGER_EMAIL EventCode = "TRIGGER-EMAIL"
	EVENTCODE_TRIGGER_NOTIFICATION EventCode = "TRIGGER-NOTIFICATION"
	EVENTCODE_TRIGGER_PUSH EventCode = "TRIGGER-PUSH"
	EVENTCODE_TRIGGER_REMOVE EventCode = "TRIGGER-REMOVE"
	EVENTCODE_WCCE_ENROLL EventCode = "WCCE-ENROLL"
)

// All allowed values of EventCode enum
var AllowedEventCodeEnumValues = []EventCode{
	"ACME-ACCOUNT-KEY-CHANGE",
	"ACME-ACCOUNT-REGISTER",
	"ACME-ACCOUNT-UPDATE",
	"ACME-AUTHORIZATION-DEACTIVATE",
	"ACME-CHALLENGE-REQUEST-VERIFY",
	"ACME-CHALLENGE-VERIFY",
	"ACME-ORDER-CERTIFICATE",
	"ACME-ORDER-FINALIZE",
	"ACME-ORDER-NEW",
	"ACME-ORDER-UPDATE",
	"ACME-REVOKE",
	"ACTOR",
	"BOOTSTRAP-ADMINISTRATOR-ACCOUNT",
	"BOOTSTRAP-ADMINISTRATOR-PRINCIPAL",
	"BOOTSTRAP-GRADING-POLICY",
	"BOOTSTRAP-GRADING-RULESET",
	"BOOTSTRAP-LOCAL-IDENTITY-PROVIDER",
	"CA-CERT-SYNC",
	"CA-CRL-UPDATE",
	"CONF-ADD",
	"CONF-DELETE",
	"CONF-UPDATE",
	"CRMP-AUTHENTICATION",
	"CRMP-BAD-REQUEST",
	"CRMP-ENROLL",
	"CRMP-LIST",
	"CRMP-PROFILE-PROPERTIES",
	"CRMP-RECOVER",
	"CRMP-RETRIEVE",
	"CRMP-REVOKE",
	"DISCOVERY-CAMPAIGN-FLUSH",
	"EST-CACERTS",
	"EST-REVOKE-ON-RENEW",
	"EST-SIMPLE-ENROLL",
	"EST-SIMPLE-REENROLL",
	"GRADING-END",
	"GRADING-ERROR",
	"GRADING-START",
	"LICENSE-ERROR",
	"LICENSE-LIMIT-REACHED",
	"LIFECYCLE-ENROLL",
	"LIFECYCLE-ESCROW",
	"LIFECYCLE-MAX-CERT-PER-HOLDER",
	"LIFECYCLE-MIGRATE",
	"LIFECYCLE-RECOVER",
	"LIFECYCLE-RENEW",
	"LIFECYCLE-REVOKE",
	"LIFECYCLE-UPDATE",
	"PKI-CONNECTOR",
	"REQUEST-APPROVE",
	"REQUEST-CANCEL",
	"REQUEST-DENY",
	"REQUEST-SUBMIT",
	"REQUEST-TEMPLATE",
	"SCEP-ENROLL",
	"SCEP-GET-CA-CERT",
	"SCEP-GET-CERT-INITIAL",
	"SCEP-NDES-EMULATION",
	"SCEP-PKI-CLIENT",
	"SCEP-PKI-OPERATION",
	"SCEP-RENEW",
	"SCEP-REVOKE-ON-RENEW",
	"SCHEDULED-TASK-COMPLETE",
	"SCHEDULED-TASK-RUN",
	"SEC-AUTHENTICATION",
	"SEC-AUTHORIZATION-ADD",
	"SEC-AUTHORIZATION-DELETE",
	"SEC-AUTHORIZATION-UPDATE",
	"SEC-IDENTITY-PROVIDER-ADD",
	"SEC-IDENTITY-PROVIDER-DELETE",
	"SEC-IDENTITY-PROVIDER-UPDATE",
	"SEC-LOCAL-IDENTITY-ADD",
	"SEC-LOCAL-IDENTITY-DELETE",
	"SEC-LOCAL-IDENTITY-RESET",
	"SEC-LOCAL-IDENTITY-UPDATE",
	"SEC-PASSWORD-POLICY-ADD",
	"SEC-PASSWORD-POLICY-DELETE",
	"SEC-PASSWORD-POLICY-UPDATE",
	"SEC-ROLE-ADD",
	"SEC-ROLE-DELETE",
	"SEC-ROLE-UPDATE",
	"SEC-TEAM-ADD",
	"SEC-TEAM-DELETE",
	"SEC-TEAM-SWITCH",
	"SEC-TEAM-UPDATE",
	"SERVICE-START",
	"SERVICE-STOP",
	"SUPERVISOR",
	"SYNC-ENROLL",
	"SYNC-RENEW",
	"SYNC-REVOKE",
	"TEAM-SWITCH",
	"TRIGGER-DELETE",
	"TRIGGER-EMAIL",
	"TRIGGER-NOTIFICATION",
	"TRIGGER-PUSH",
	"TRIGGER-REMOVE",
	"WCCE-ENROLL",
}

func (v *EventCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventCode(value)
	for _, existing := range AllowedEventCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventCode", value)
}

// NewEventCodeFromValue returns a pointer to a valid EventCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventCodeFromValue(v string) (*EventCode, error) {
	ev := EventCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventCode: valid values are %v", v, AllowedEventCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventCode) IsValid() bool {
	for _, existing := range AllowedEventCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventCode value
func (v EventCode) Ptr() *EventCode {
	return &v
}

type NullableEventCode struct {
	value *EventCode
	isSet bool
}

func (v NullableEventCode) Get() *EventCode {
	return v.value
}

func (v *NullableEventCode) Set(val *EventCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCode(val *EventCode) *NullableEventCode {
	return &NullableEventCode{value: val, isSet: true}
}

func (v NullableEventCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

