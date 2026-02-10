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

// CertificateFind400Response - struct for CertificateFind400Response
type CertificateFind400Response struct {
	CertificateSearch400ResponseOneOf  *CertificateSearch400ResponseOneOf
	CertificateSearch400ResponseOneOf1 *CertificateSearch400ResponseOneOf1
	CertificateSearch400ResponseOneOf2 *CertificateSearch400ResponseOneOf2
}

// CertificateSearch400ResponseOneOfAsCertificateFind400Response is a convenience function that returns CertificateSearch400ResponseOneOf wrapped in CertificateFind400Response
func CertificateSearch400ResponseOneOfAsCertificateFind400Response(v *CertificateSearch400ResponseOneOf) CertificateFind400Response {
	return CertificateFind400Response{
		CertificateSearch400ResponseOneOf: v,
	}
}

// CertificateSearch400ResponseOneOf1AsCertificateFind400Response is a convenience function that returns CertificateSearch400ResponseOneOf1 wrapped in CertificateFind400Response
func CertificateSearch400ResponseOneOf1AsCertificateFind400Response(v *CertificateSearch400ResponseOneOf1) CertificateFind400Response {
	return CertificateFind400Response{
		CertificateSearch400ResponseOneOf1: v,
	}
}

// CertificateSearch400ResponseOneOf2AsCertificateFind400Response is a convenience function that returns CertificateSearch400ResponseOneOf2 wrapped in CertificateFind400Response
func CertificateSearch400ResponseOneOf2AsCertificateFind400Response(v *CertificateSearch400ResponseOneOf2) CertificateFind400Response {
	return CertificateFind400Response{
		CertificateSearch400ResponseOneOf2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CertificateFind400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CertificateSearch400ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.CertificateSearch400ResponseOneOf)
	if err == nil {
		jsonCertificateSearch400ResponseOneOf, _ := json.Marshal(dst.CertificateSearch400ResponseOneOf)
		if string(jsonCertificateSearch400ResponseOneOf) == "{}" { // empty struct
			dst.CertificateSearch400ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.CertificateSearch400ResponseOneOf)
			match++
		}
	} else {
		dst.CertificateSearch400ResponseOneOf = nil
	}

	// try to unmarshal data into CertificateSearch400ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.CertificateSearch400ResponseOneOf1)
	if err == nil {
		jsonCertificateSearch400ResponseOneOf1, _ := json.Marshal(dst.CertificateSearch400ResponseOneOf1)
		if string(jsonCertificateSearch400ResponseOneOf1) == "{}" { // empty struct
			dst.CertificateSearch400ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.CertificateSearch400ResponseOneOf1)
			match++
		}
	} else {
		dst.CertificateSearch400ResponseOneOf1 = nil
	}

	// try to unmarshal data into CertificateSearch400ResponseOneOf2
	err = utils.NewStrictDecoder(data).Decode(&dst.CertificateSearch400ResponseOneOf2)
	if err == nil {
		jsonCertificateSearch400ResponseOneOf2, _ := json.Marshal(dst.CertificateSearch400ResponseOneOf2)
		if string(jsonCertificateSearch400ResponseOneOf2) == "{}" { // empty struct
			dst.CertificateSearch400ResponseOneOf2 = nil
		} else {
			_ = validator.Validate(dst.CertificateSearch400ResponseOneOf2)
			match++
		}
	} else {
		dst.CertificateSearch400ResponseOneOf2 = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CertificateFind400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CertificateFind400Response) MarshalJSON() ([]byte, error) {
	if src.CertificateSearch400ResponseOneOf != nil {
		return json.Marshal(&src.CertificateSearch400ResponseOneOf)
	}

	if src.CertificateSearch400ResponseOneOf1 != nil {
		return json.Marshal(&src.CertificateSearch400ResponseOneOf1)
	}

	if src.CertificateSearch400ResponseOneOf2 != nil {
		return json.Marshal(&src.CertificateSearch400ResponseOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CertificateFind400Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CertificateSearch400ResponseOneOf != nil {
		return obj.CertificateSearch400ResponseOneOf
	}

	if obj.CertificateSearch400ResponseOneOf1 != nil {
		return obj.CertificateSearch400ResponseOneOf1
	}

	if obj.CertificateSearch400ResponseOneOf2 != nil {
		return obj.CertificateSearch400ResponseOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableCertificateFind400Response struct {
	value *CertificateFind400Response
	isSet bool
}

func (v NullableCertificateFind400Response) Get() *CertificateFind400Response {
	return v.value
}

func (v *NullableCertificateFind400Response) Set(val *CertificateFind400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateFind400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateFind400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateFind400Response(val *CertificateFind400Response) *NullableCertificateFind400Response {
	return &NullableCertificateFind400Response{value: val, isSet: true}
}

func (v NullableCertificateFind400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateFind400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
