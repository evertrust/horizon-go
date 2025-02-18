/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// SecurityIdentityProviderAdd400Response - struct for SecurityIdentityProviderAdd400Response
type SecurityIdentityProviderAdd400Response struct {
	SecurityIdentityProviderAdd400ResponseOneOf *SecurityIdentityProviderAdd400ResponseOneOf
	SecurityIdentityProviderUpdate400ResponseOneOf *SecurityIdentityProviderUpdate400ResponseOneOf
}

// SecurityIdentityProviderAdd400ResponseOneOfAsSecurityIdentityProviderAdd400Response is a convenience function that returns SecurityIdentityProviderAdd400ResponseOneOf wrapped in SecurityIdentityProviderAdd400Response
func SecurityIdentityProviderAdd400ResponseOneOfAsSecurityIdentityProviderAdd400Response(v *SecurityIdentityProviderAdd400ResponseOneOf) SecurityIdentityProviderAdd400Response {
	return SecurityIdentityProviderAdd400Response{
		SecurityIdentityProviderAdd400ResponseOneOf: v,
	}
}

// SecurityIdentityProviderUpdate400ResponseOneOfAsSecurityIdentityProviderAdd400Response is a convenience function that returns SecurityIdentityProviderUpdate400ResponseOneOf wrapped in SecurityIdentityProviderAdd400Response
func SecurityIdentityProviderUpdate400ResponseOneOfAsSecurityIdentityProviderAdd400Response(v *SecurityIdentityProviderUpdate400ResponseOneOf) SecurityIdentityProviderAdd400Response {
	return SecurityIdentityProviderAdd400Response{
		SecurityIdentityProviderUpdate400ResponseOneOf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityIdentityProviderAdd400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SecurityIdentityProviderAdd400ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.SecurityIdentityProviderAdd400ResponseOneOf)
	if err == nil {
		jsonSecurityIdentityProviderAdd400ResponseOneOf, _ := json.Marshal(dst.SecurityIdentityProviderAdd400ResponseOneOf)
		if string(jsonSecurityIdentityProviderAdd400ResponseOneOf) == "{}" { // empty struct
			dst.SecurityIdentityProviderAdd400ResponseOneOf = nil
		} else {
            _ = validator.Validate(dst.SecurityIdentityProviderAdd400ResponseOneOf)
            match++
		}
	} else {
		dst.SecurityIdentityProviderAdd400ResponseOneOf = nil
	}

	// try to unmarshal data into SecurityIdentityProviderUpdate400ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.SecurityIdentityProviderUpdate400ResponseOneOf)
	if err == nil {
		jsonSecurityIdentityProviderUpdate400ResponseOneOf, _ := json.Marshal(dst.SecurityIdentityProviderUpdate400ResponseOneOf)
		if string(jsonSecurityIdentityProviderUpdate400ResponseOneOf) == "{}" { // empty struct
			dst.SecurityIdentityProviderUpdate400ResponseOneOf = nil
		} else {
            _ = validator.Validate(dst.SecurityIdentityProviderUpdate400ResponseOneOf)
            match++
		}
	} else {
		dst.SecurityIdentityProviderUpdate400ResponseOneOf = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SecurityIdentityProviderAdd400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityIdentityProviderAdd400Response) MarshalJSON() ([]byte, error) {
	if src.SecurityIdentityProviderAdd400ResponseOneOf != nil {
		return json.Marshal(&src.SecurityIdentityProviderAdd400ResponseOneOf)
	}

	if src.SecurityIdentityProviderUpdate400ResponseOneOf != nil {
		return json.Marshal(&src.SecurityIdentityProviderUpdate400ResponseOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityIdentityProviderAdd400Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SecurityIdentityProviderAdd400ResponseOneOf != nil {
		return obj.SecurityIdentityProviderAdd400ResponseOneOf
	}

	if obj.SecurityIdentityProviderUpdate400ResponseOneOf != nil {
		return obj.SecurityIdentityProviderUpdate400ResponseOneOf
	}

	// all schemas are nil
	return nil
}

type NullableSecurityIdentityProviderAdd400Response struct {
	value *SecurityIdentityProviderAdd400Response
	isSet bool
}

func (v NullableSecurityIdentityProviderAdd400Response) Get() *SecurityIdentityProviderAdd400Response {
	return v.value
}

func (v *NullableSecurityIdentityProviderAdd400Response) Set(val *SecurityIdentityProviderAdd400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityIdentityProviderAdd400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityIdentityProviderAdd400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityIdentityProviderAdd400Response(val *SecurityIdentityProviderAdd400Response) *NullableSecurityIdentityProviderAdd400Response {
	return &NullableSecurityIdentityProviderAdd400Response{value: val, isSet: true}
}

func (v NullableSecurityIdentityProviderAdd400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityIdentityProviderAdd400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


