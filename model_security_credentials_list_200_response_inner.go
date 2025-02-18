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

// SecurityCredentialsList200ResponseInner - struct for SecurityCredentialsList200ResponseInner
type SecurityCredentialsList200ResponseInner struct {
	CertificateCredentialsResponse *CertificateCredentialsResponse
	PasswordCredentialsResponse *PasswordCredentialsResponse
	RawCredentialsResponse *RawCredentialsResponse
}

// CertificateCredentialsResponseAsSecurityCredentialsList200ResponseInner is a convenience function that returns CertificateCredentialsResponse wrapped in SecurityCredentialsList200ResponseInner
func CertificateCredentialsResponseAsSecurityCredentialsList200ResponseInner(v *CertificateCredentialsResponse) SecurityCredentialsList200ResponseInner {
	return SecurityCredentialsList200ResponseInner{
		CertificateCredentialsResponse: v,
	}
}

// PasswordCredentialsResponseAsSecurityCredentialsList200ResponseInner is a convenience function that returns PasswordCredentialsResponse wrapped in SecurityCredentialsList200ResponseInner
func PasswordCredentialsResponseAsSecurityCredentialsList200ResponseInner(v *PasswordCredentialsResponse) SecurityCredentialsList200ResponseInner {
	return SecurityCredentialsList200ResponseInner{
		PasswordCredentialsResponse: v,
	}
}

// RawCredentialsResponseAsSecurityCredentialsList200ResponseInner is a convenience function that returns RawCredentialsResponse wrapped in SecurityCredentialsList200ResponseInner
func RawCredentialsResponseAsSecurityCredentialsList200ResponseInner(v *RawCredentialsResponse) SecurityCredentialsList200ResponseInner {
	return SecurityCredentialsList200ResponseInner{
		RawCredentialsResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityCredentialsList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CertificateCredentialsResponse
	err = newStrictDecoder(data).Decode(&dst.CertificateCredentialsResponse)
	if err == nil {
		jsonCertificateCredentialsResponse, _ := json.Marshal(dst.CertificateCredentialsResponse)
		if string(jsonCertificateCredentialsResponse) == "{}" { // empty struct
			dst.CertificateCredentialsResponse = nil
		} else {
            _ = validator.Validate(dst.CertificateCredentialsResponse)
            match++
		}
	} else {
		dst.CertificateCredentialsResponse = nil
	}

	// try to unmarshal data into PasswordCredentialsResponse
	err = newStrictDecoder(data).Decode(&dst.PasswordCredentialsResponse)
	if err == nil {
		jsonPasswordCredentialsResponse, _ := json.Marshal(dst.PasswordCredentialsResponse)
		if string(jsonPasswordCredentialsResponse) == "{}" { // empty struct
			dst.PasswordCredentialsResponse = nil
		} else {
            _ = validator.Validate(dst.PasswordCredentialsResponse)
            match++
		}
	} else {
		dst.PasswordCredentialsResponse = nil
	}

	// try to unmarshal data into RawCredentialsResponse
	err = newStrictDecoder(data).Decode(&dst.RawCredentialsResponse)
	if err == nil {
		jsonRawCredentialsResponse, _ := json.Marshal(dst.RawCredentialsResponse)
		if string(jsonRawCredentialsResponse) == "{}" { // empty struct
			dst.RawCredentialsResponse = nil
		} else {
            _ = validator.Validate(dst.RawCredentialsResponse)
            match++
		}
	} else {
		dst.RawCredentialsResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SecurityCredentialsList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityCredentialsList200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.CertificateCredentialsResponse != nil {
		return json.Marshal(&src.CertificateCredentialsResponse)
	}

	if src.PasswordCredentialsResponse != nil {
		return json.Marshal(&src.PasswordCredentialsResponse)
	}

	if src.RawCredentialsResponse != nil {
		return json.Marshal(&src.RawCredentialsResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityCredentialsList200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CertificateCredentialsResponse != nil {
		return obj.CertificateCredentialsResponse
	}

	if obj.PasswordCredentialsResponse != nil {
		return obj.PasswordCredentialsResponse
	}

	if obj.RawCredentialsResponse != nil {
		return obj.RawCredentialsResponse
	}

	// all schemas are nil
	return nil
}

type NullableSecurityCredentialsList200ResponseInner struct {
	value *SecurityCredentialsList200ResponseInner
	isSet bool
}

func (v NullableSecurityCredentialsList200ResponseInner) Get() *SecurityCredentialsList200ResponseInner {
	return v.value
}

func (v *NullableSecurityCredentialsList200ResponseInner) Set(val *SecurityCredentialsList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCredentialsList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCredentialsList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCredentialsList200ResponseInner(val *SecurityCredentialsList200ResponseInner) *NullableSecurityCredentialsList200ResponseInner {
	return &NullableSecurityCredentialsList200ResponseInner{value: val, isSet: true}
}

func (v NullableSecurityCredentialsList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCredentialsList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


