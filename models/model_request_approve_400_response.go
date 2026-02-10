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

// RequestApprove400Response - struct for RequestApprove400Response
type RequestApprove400Response struct {
	CertificateList400ResponseOneOf  *CertificateList400ResponseOneOf
	CertificateList400ResponseOneOf1 *CertificateList400ResponseOneOf1
	RequestApprove400ResponseOneOf   *RequestApprove400ResponseOneOf
	RequestApprove400ResponseOneOf1  *RequestApprove400ResponseOneOf1
	RequestApprove400ResponseOneOf10 *RequestApprove400ResponseOneOf10
	RequestApprove400ResponseOneOf11 *RequestApprove400ResponseOneOf11
	RequestApprove400ResponseOneOf12 *RequestApprove400ResponseOneOf12
	RequestApprove400ResponseOneOf13 *RequestApprove400ResponseOneOf13
	RequestApprove400ResponseOneOf14 *RequestApprove400ResponseOneOf14
	RequestApprove400ResponseOneOf15 *RequestApprove400ResponseOneOf15
	RequestApprove400ResponseOneOf16 *RequestApprove400ResponseOneOf16
	RequestApprove400ResponseOneOf17 *RequestApprove400ResponseOneOf17
	RequestApprove400ResponseOneOf18 *RequestApprove400ResponseOneOf18
	RequestApprove400ResponseOneOf19 *RequestApprove400ResponseOneOf19
	RequestApprove400ResponseOneOf2  *RequestApprove400ResponseOneOf2
	RequestApprove400ResponseOneOf20 *RequestApprove400ResponseOneOf20
	RequestApprove400ResponseOneOf21 *RequestApprove400ResponseOneOf21
	RequestApprove400ResponseOneOf22 *RequestApprove400ResponseOneOf22
	RequestApprove400ResponseOneOf23 *RequestApprove400ResponseOneOf23
	RequestApprove400ResponseOneOf24 *RequestApprove400ResponseOneOf24
	RequestApprove400ResponseOneOf25 *RequestApprove400ResponseOneOf25
	RequestApprove400ResponseOneOf26 *RequestApprove400ResponseOneOf26
	RequestApprove400ResponseOneOf27 *RequestApprove400ResponseOneOf27
	RequestApprove400ResponseOneOf28 *RequestApprove400ResponseOneOf28
	RequestApprove400ResponseOneOf29 *RequestApprove400ResponseOneOf29
	RequestApprove400ResponseOneOf3  *RequestApprove400ResponseOneOf3
	RequestApprove400ResponseOneOf30 *RequestApprove400ResponseOneOf30
	RequestApprove400ResponseOneOf31 *RequestApprove400ResponseOneOf31
	RequestApprove400ResponseOneOf32 *RequestApprove400ResponseOneOf32
	RequestApprove400ResponseOneOf33 *RequestApprove400ResponseOneOf33
	RequestApprove400ResponseOneOf34 *RequestApprove400ResponseOneOf34
	RequestApprove400ResponseOneOf4  *RequestApprove400ResponseOneOf4
	RequestApprove400ResponseOneOf5  *RequestApprove400ResponseOneOf5
	RequestApprove400ResponseOneOf6  *RequestApprove400ResponseOneOf6
	RequestApprove400ResponseOneOf7  *RequestApprove400ResponseOneOf7
	RequestApprove400ResponseOneOf8  *RequestApprove400ResponseOneOf8
	RequestApprove400ResponseOneOf9  *RequestApprove400ResponseOneOf9
	RequestCancel400ResponseOneOf    *RequestCancel400ResponseOneOf
	RequestCancel400ResponseOneOf1   *RequestCancel400ResponseOneOf1
	RequestSubmit400ResponseOneOf38  *RequestSubmit400ResponseOneOf38
	RequestSubmit400ResponseOneOf39  *RequestSubmit400ResponseOneOf39
	RequestTemplate400ResponseOneOf  *RequestTemplate400ResponseOneOf
	RequestTemplate400ResponseOneOf2 *RequestTemplate400ResponseOneOf2
}

// CertificateList400ResponseOneOfAsRequestApprove400Response is a convenience function that returns CertificateList400ResponseOneOf wrapped in RequestApprove400Response
func CertificateList400ResponseOneOfAsRequestApprove400Response(v *CertificateList400ResponseOneOf) RequestApprove400Response {
	return RequestApprove400Response{
		CertificateList400ResponseOneOf: v,
	}
}

// CertificateList400ResponseOneOf1AsRequestApprove400Response is a convenience function that returns CertificateList400ResponseOneOf1 wrapped in RequestApprove400Response
func CertificateList400ResponseOneOf1AsRequestApprove400Response(v *CertificateList400ResponseOneOf1) RequestApprove400Response {
	return RequestApprove400Response{
		CertificateList400ResponseOneOf1: v,
	}
}

// RequestApprove400ResponseOneOfAsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOfAsRequestApprove400Response(v *RequestApprove400ResponseOneOf) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf: v,
	}
}

// RequestApprove400ResponseOneOf1AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf1 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf1AsRequestApprove400Response(v *RequestApprove400ResponseOneOf1) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf1: v,
	}
}

// RequestApprove400ResponseOneOf10AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf10 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf10AsRequestApprove400Response(v *RequestApprove400ResponseOneOf10) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf10: v,
	}
}

// RequestApprove400ResponseOneOf11AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf11 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf11AsRequestApprove400Response(v *RequestApprove400ResponseOneOf11) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf11: v,
	}
}

// RequestApprove400ResponseOneOf12AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf12 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf12AsRequestApprove400Response(v *RequestApprove400ResponseOneOf12) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf12: v,
	}
}

// RequestApprove400ResponseOneOf13AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf13 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf13AsRequestApprove400Response(v *RequestApprove400ResponseOneOf13) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf13: v,
	}
}

// RequestApprove400ResponseOneOf14AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf14 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf14AsRequestApprove400Response(v *RequestApprove400ResponseOneOf14) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf14: v,
	}
}

// RequestApprove400ResponseOneOf15AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf15 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf15AsRequestApprove400Response(v *RequestApprove400ResponseOneOf15) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf15: v,
	}
}

// RequestApprove400ResponseOneOf16AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf16 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf16AsRequestApprove400Response(v *RequestApprove400ResponseOneOf16) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf16: v,
	}
}

// RequestApprove400ResponseOneOf17AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf17 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf17AsRequestApprove400Response(v *RequestApprove400ResponseOneOf17) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf17: v,
	}
}

// RequestApprove400ResponseOneOf18AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf18 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf18AsRequestApprove400Response(v *RequestApprove400ResponseOneOf18) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf18: v,
	}
}

// RequestApprove400ResponseOneOf19AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf19 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf19AsRequestApprove400Response(v *RequestApprove400ResponseOneOf19) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf19: v,
	}
}

// RequestApprove400ResponseOneOf2AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf2 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf2AsRequestApprove400Response(v *RequestApprove400ResponseOneOf2) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf2: v,
	}
}

// RequestApprove400ResponseOneOf20AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf20 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf20AsRequestApprove400Response(v *RequestApprove400ResponseOneOf20) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf20: v,
	}
}

// RequestApprove400ResponseOneOf21AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf21 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf21AsRequestApprove400Response(v *RequestApprove400ResponseOneOf21) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf21: v,
	}
}

// RequestApprove400ResponseOneOf22AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf22 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf22AsRequestApprove400Response(v *RequestApprove400ResponseOneOf22) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf22: v,
	}
}

// RequestApprove400ResponseOneOf23AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf23 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf23AsRequestApprove400Response(v *RequestApprove400ResponseOneOf23) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf23: v,
	}
}

// RequestApprove400ResponseOneOf24AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf24 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf24AsRequestApprove400Response(v *RequestApprove400ResponseOneOf24) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf24: v,
	}
}

// RequestApprove400ResponseOneOf25AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf25 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf25AsRequestApprove400Response(v *RequestApprove400ResponseOneOf25) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf25: v,
	}
}

// RequestApprove400ResponseOneOf26AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf26 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf26AsRequestApprove400Response(v *RequestApprove400ResponseOneOf26) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf26: v,
	}
}

// RequestApprove400ResponseOneOf27AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf27 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf27AsRequestApprove400Response(v *RequestApprove400ResponseOneOf27) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf27: v,
	}
}

// RequestApprove400ResponseOneOf28AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf28 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf28AsRequestApprove400Response(v *RequestApprove400ResponseOneOf28) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf28: v,
	}
}

// RequestApprove400ResponseOneOf29AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf29 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf29AsRequestApprove400Response(v *RequestApprove400ResponseOneOf29) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf29: v,
	}
}

// RequestApprove400ResponseOneOf3AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf3 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf3AsRequestApprove400Response(v *RequestApprove400ResponseOneOf3) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf3: v,
	}
}

// RequestApprove400ResponseOneOf30AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf30 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf30AsRequestApprove400Response(v *RequestApprove400ResponseOneOf30) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf30: v,
	}
}

// RequestApprove400ResponseOneOf31AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf31 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf31AsRequestApprove400Response(v *RequestApprove400ResponseOneOf31) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf31: v,
	}
}

// RequestApprove400ResponseOneOf32AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf32 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf32AsRequestApprove400Response(v *RequestApprove400ResponseOneOf32) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf32: v,
	}
}

// RequestApprove400ResponseOneOf33AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf33 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf33AsRequestApprove400Response(v *RequestApprove400ResponseOneOf33) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf33: v,
	}
}

// RequestApprove400ResponseOneOf34AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf34 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf34AsRequestApprove400Response(v *RequestApprove400ResponseOneOf34) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf34: v,
	}
}

// RequestApprove400ResponseOneOf4AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf4 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf4AsRequestApprove400Response(v *RequestApprove400ResponseOneOf4) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf4: v,
	}
}

// RequestApprove400ResponseOneOf5AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf5 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf5AsRequestApprove400Response(v *RequestApprove400ResponseOneOf5) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf5: v,
	}
}

// RequestApprove400ResponseOneOf6AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf6 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf6AsRequestApprove400Response(v *RequestApprove400ResponseOneOf6) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf6: v,
	}
}

// RequestApprove400ResponseOneOf7AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf7 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf7AsRequestApprove400Response(v *RequestApprove400ResponseOneOf7) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf7: v,
	}
}

// RequestApprove400ResponseOneOf8AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf8 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf8AsRequestApprove400Response(v *RequestApprove400ResponseOneOf8) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf8: v,
	}
}

// RequestApprove400ResponseOneOf9AsRequestApprove400Response is a convenience function that returns RequestApprove400ResponseOneOf9 wrapped in RequestApprove400Response
func RequestApprove400ResponseOneOf9AsRequestApprove400Response(v *RequestApprove400ResponseOneOf9) RequestApprove400Response {
	return RequestApprove400Response{
		RequestApprove400ResponseOneOf9: v,
	}
}

// RequestCancel400ResponseOneOfAsRequestApprove400Response is a convenience function that returns RequestCancel400ResponseOneOf wrapped in RequestApprove400Response
func RequestCancel400ResponseOneOfAsRequestApprove400Response(v *RequestCancel400ResponseOneOf) RequestApprove400Response {
	return RequestApprove400Response{
		RequestCancel400ResponseOneOf: v,
	}
}

// RequestCancel400ResponseOneOf1AsRequestApprove400Response is a convenience function that returns RequestCancel400ResponseOneOf1 wrapped in RequestApprove400Response
func RequestCancel400ResponseOneOf1AsRequestApprove400Response(v *RequestCancel400ResponseOneOf1) RequestApprove400Response {
	return RequestApprove400Response{
		RequestCancel400ResponseOneOf1: v,
	}
}

// RequestSubmit400ResponseOneOf38AsRequestApprove400Response is a convenience function that returns RequestSubmit400ResponseOneOf38 wrapped in RequestApprove400Response
func RequestSubmit400ResponseOneOf38AsRequestApprove400Response(v *RequestSubmit400ResponseOneOf38) RequestApprove400Response {
	return RequestApprove400Response{
		RequestSubmit400ResponseOneOf38: v,
	}
}

// RequestSubmit400ResponseOneOf39AsRequestApprove400Response is a convenience function that returns RequestSubmit400ResponseOneOf39 wrapped in RequestApprove400Response
func RequestSubmit400ResponseOneOf39AsRequestApprove400Response(v *RequestSubmit400ResponseOneOf39) RequestApprove400Response {
	return RequestApprove400Response{
		RequestSubmit400ResponseOneOf39: v,
	}
}

// RequestTemplate400ResponseOneOfAsRequestApprove400Response is a convenience function that returns RequestTemplate400ResponseOneOf wrapped in RequestApprove400Response
func RequestTemplate400ResponseOneOfAsRequestApprove400Response(v *RequestTemplate400ResponseOneOf) RequestApprove400Response {
	return RequestApprove400Response{
		RequestTemplate400ResponseOneOf: v,
	}
}

// RequestTemplate400ResponseOneOf2AsRequestApprove400Response is a convenience function that returns RequestTemplate400ResponseOneOf2 wrapped in RequestApprove400Response
func RequestTemplate400ResponseOneOf2AsRequestApprove400Response(v *RequestTemplate400ResponseOneOf2) RequestApprove400Response {
	return RequestApprove400Response{
		RequestTemplate400ResponseOneOf2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestApprove400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CertificateList400ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.CertificateList400ResponseOneOf)
	if err == nil {
		jsonCertificateList400ResponseOneOf, _ := json.Marshal(dst.CertificateList400ResponseOneOf)
		if string(jsonCertificateList400ResponseOneOf) == "{}" { // empty struct
			dst.CertificateList400ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.CertificateList400ResponseOneOf)
			match++
		}
	} else {
		dst.CertificateList400ResponseOneOf = nil
	}

	// try to unmarshal data into CertificateList400ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.CertificateList400ResponseOneOf1)
	if err == nil {
		jsonCertificateList400ResponseOneOf1, _ := json.Marshal(dst.CertificateList400ResponseOneOf1)
		if string(jsonCertificateList400ResponseOneOf1) == "{}" { // empty struct
			dst.CertificateList400ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.CertificateList400ResponseOneOf1)
			match++
		}
	} else {
		dst.CertificateList400ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf)
	if err == nil {
		jsonRequestApprove400ResponseOneOf, _ := json.Marshal(dst.RequestApprove400ResponseOneOf)
		if string(jsonRequestApprove400ResponseOneOf) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf1)
	if err == nil {
		jsonRequestApprove400ResponseOneOf1, _ := json.Marshal(dst.RequestApprove400ResponseOneOf1)
		if string(jsonRequestApprove400ResponseOneOf1) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf1)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf10
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf10)
	if err == nil {
		jsonRequestApprove400ResponseOneOf10, _ := json.Marshal(dst.RequestApprove400ResponseOneOf10)
		if string(jsonRequestApprove400ResponseOneOf10) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf10 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf10)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf10 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf11
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf11)
	if err == nil {
		jsonRequestApprove400ResponseOneOf11, _ := json.Marshal(dst.RequestApprove400ResponseOneOf11)
		if string(jsonRequestApprove400ResponseOneOf11) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf11 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf11)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf11 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf12
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf12)
	if err == nil {
		jsonRequestApprove400ResponseOneOf12, _ := json.Marshal(dst.RequestApprove400ResponseOneOf12)
		if string(jsonRequestApprove400ResponseOneOf12) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf12 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf12)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf12 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf13
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf13)
	if err == nil {
		jsonRequestApprove400ResponseOneOf13, _ := json.Marshal(dst.RequestApprove400ResponseOneOf13)
		if string(jsonRequestApprove400ResponseOneOf13) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf13 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf13)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf13 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf14
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf14)
	if err == nil {
		jsonRequestApprove400ResponseOneOf14, _ := json.Marshal(dst.RequestApprove400ResponseOneOf14)
		if string(jsonRequestApprove400ResponseOneOf14) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf14 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf14)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf14 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf15
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf15)
	if err == nil {
		jsonRequestApprove400ResponseOneOf15, _ := json.Marshal(dst.RequestApprove400ResponseOneOf15)
		if string(jsonRequestApprove400ResponseOneOf15) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf15 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf15)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf15 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf16
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf16)
	if err == nil {
		jsonRequestApprove400ResponseOneOf16, _ := json.Marshal(dst.RequestApprove400ResponseOneOf16)
		if string(jsonRequestApprove400ResponseOneOf16) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf16 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf16)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf16 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf17
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf17)
	if err == nil {
		jsonRequestApprove400ResponseOneOf17, _ := json.Marshal(dst.RequestApprove400ResponseOneOf17)
		if string(jsonRequestApprove400ResponseOneOf17) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf17 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf17)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf17 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf18
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf18)
	if err == nil {
		jsonRequestApprove400ResponseOneOf18, _ := json.Marshal(dst.RequestApprove400ResponseOneOf18)
		if string(jsonRequestApprove400ResponseOneOf18) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf18 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf18)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf18 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf19
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf19)
	if err == nil {
		jsonRequestApprove400ResponseOneOf19, _ := json.Marshal(dst.RequestApprove400ResponseOneOf19)
		if string(jsonRequestApprove400ResponseOneOf19) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf19 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf19)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf19 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf2
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf2)
	if err == nil {
		jsonRequestApprove400ResponseOneOf2, _ := json.Marshal(dst.RequestApprove400ResponseOneOf2)
		if string(jsonRequestApprove400ResponseOneOf2) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf2 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf2)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf2 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf20
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf20)
	if err == nil {
		jsonRequestApprove400ResponseOneOf20, _ := json.Marshal(dst.RequestApprove400ResponseOneOf20)
		if string(jsonRequestApprove400ResponseOneOf20) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf20 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf20)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf20 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf21
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf21)
	if err == nil {
		jsonRequestApprove400ResponseOneOf21, _ := json.Marshal(dst.RequestApprove400ResponseOneOf21)
		if string(jsonRequestApprove400ResponseOneOf21) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf21 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf21)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf21 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf22
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf22)
	if err == nil {
		jsonRequestApprove400ResponseOneOf22, _ := json.Marshal(dst.RequestApprove400ResponseOneOf22)
		if string(jsonRequestApprove400ResponseOneOf22) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf22 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf22)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf22 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf23
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf23)
	if err == nil {
		jsonRequestApprove400ResponseOneOf23, _ := json.Marshal(dst.RequestApprove400ResponseOneOf23)
		if string(jsonRequestApprove400ResponseOneOf23) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf23 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf23)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf23 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf24
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf24)
	if err == nil {
		jsonRequestApprove400ResponseOneOf24, _ := json.Marshal(dst.RequestApprove400ResponseOneOf24)
		if string(jsonRequestApprove400ResponseOneOf24) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf24 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf24)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf24 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf25
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf25)
	if err == nil {
		jsonRequestApprove400ResponseOneOf25, _ := json.Marshal(dst.RequestApprove400ResponseOneOf25)
		if string(jsonRequestApprove400ResponseOneOf25) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf25 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf25)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf25 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf26
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf26)
	if err == nil {
		jsonRequestApprove400ResponseOneOf26, _ := json.Marshal(dst.RequestApprove400ResponseOneOf26)
		if string(jsonRequestApprove400ResponseOneOf26) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf26 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf26)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf26 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf27
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf27)
	if err == nil {
		jsonRequestApprove400ResponseOneOf27, _ := json.Marshal(dst.RequestApprove400ResponseOneOf27)
		if string(jsonRequestApprove400ResponseOneOf27) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf27 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf27)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf27 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf28
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf28)
	if err == nil {
		jsonRequestApprove400ResponseOneOf28, _ := json.Marshal(dst.RequestApprove400ResponseOneOf28)
		if string(jsonRequestApprove400ResponseOneOf28) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf28 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf28)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf28 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf29
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf29)
	if err == nil {
		jsonRequestApprove400ResponseOneOf29, _ := json.Marshal(dst.RequestApprove400ResponseOneOf29)
		if string(jsonRequestApprove400ResponseOneOf29) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf29 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf29)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf29 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf3
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf3)
	if err == nil {
		jsonRequestApprove400ResponseOneOf3, _ := json.Marshal(dst.RequestApprove400ResponseOneOf3)
		if string(jsonRequestApprove400ResponseOneOf3) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf3 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf3)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf3 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf30
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf30)
	if err == nil {
		jsonRequestApprove400ResponseOneOf30, _ := json.Marshal(dst.RequestApprove400ResponseOneOf30)
		if string(jsonRequestApprove400ResponseOneOf30) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf30 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf30)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf30 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf31
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf31)
	if err == nil {
		jsonRequestApprove400ResponseOneOf31, _ := json.Marshal(dst.RequestApprove400ResponseOneOf31)
		if string(jsonRequestApprove400ResponseOneOf31) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf31 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf31)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf31 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf32
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf32)
	if err == nil {
		jsonRequestApprove400ResponseOneOf32, _ := json.Marshal(dst.RequestApprove400ResponseOneOf32)
		if string(jsonRequestApprove400ResponseOneOf32) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf32 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf32)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf32 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf33
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf33)
	if err == nil {
		jsonRequestApprove400ResponseOneOf33, _ := json.Marshal(dst.RequestApprove400ResponseOneOf33)
		if string(jsonRequestApprove400ResponseOneOf33) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf33 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf33)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf33 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf34
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf34)
	if err == nil {
		jsonRequestApprove400ResponseOneOf34, _ := json.Marshal(dst.RequestApprove400ResponseOneOf34)
		if string(jsonRequestApprove400ResponseOneOf34) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf34 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf34)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf34 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf4
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf4)
	if err == nil {
		jsonRequestApprove400ResponseOneOf4, _ := json.Marshal(dst.RequestApprove400ResponseOneOf4)
		if string(jsonRequestApprove400ResponseOneOf4) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf4 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf4)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf4 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf5
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf5)
	if err == nil {
		jsonRequestApprove400ResponseOneOf5, _ := json.Marshal(dst.RequestApprove400ResponseOneOf5)
		if string(jsonRequestApprove400ResponseOneOf5) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf5 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf5)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf5 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf6
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf6)
	if err == nil {
		jsonRequestApprove400ResponseOneOf6, _ := json.Marshal(dst.RequestApprove400ResponseOneOf6)
		if string(jsonRequestApprove400ResponseOneOf6) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf6 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf6)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf6 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf7
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf7)
	if err == nil {
		jsonRequestApprove400ResponseOneOf7, _ := json.Marshal(dst.RequestApprove400ResponseOneOf7)
		if string(jsonRequestApprove400ResponseOneOf7) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf7 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf7)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf7 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf8
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf8)
	if err == nil {
		jsonRequestApprove400ResponseOneOf8, _ := json.Marshal(dst.RequestApprove400ResponseOneOf8)
		if string(jsonRequestApprove400ResponseOneOf8) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf8 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf8)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf8 = nil
	}

	// try to unmarshal data into RequestApprove400ResponseOneOf9
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestApprove400ResponseOneOf9)
	if err == nil {
		jsonRequestApprove400ResponseOneOf9, _ := json.Marshal(dst.RequestApprove400ResponseOneOf9)
		if string(jsonRequestApprove400ResponseOneOf9) == "{}" { // empty struct
			dst.RequestApprove400ResponseOneOf9 = nil
		} else {
			_ = validator.Validate(dst.RequestApprove400ResponseOneOf9)
			match++
		}
	} else {
		dst.RequestApprove400ResponseOneOf9 = nil
	}

	// try to unmarshal data into RequestCancel400ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestCancel400ResponseOneOf)
	if err == nil {
		jsonRequestCancel400ResponseOneOf, _ := json.Marshal(dst.RequestCancel400ResponseOneOf)
		if string(jsonRequestCancel400ResponseOneOf) == "{}" { // empty struct
			dst.RequestCancel400ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.RequestCancel400ResponseOneOf)
			match++
		}
	} else {
		dst.RequestCancel400ResponseOneOf = nil
	}

	// try to unmarshal data into RequestCancel400ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestCancel400ResponseOneOf1)
	if err == nil {
		jsonRequestCancel400ResponseOneOf1, _ := json.Marshal(dst.RequestCancel400ResponseOneOf1)
		if string(jsonRequestCancel400ResponseOneOf1) == "{}" { // empty struct
			dst.RequestCancel400ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.RequestCancel400ResponseOneOf1)
			match++
		}
	} else {
		dst.RequestCancel400ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf38
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf38)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf38, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf38)
		if string(jsonRequestSubmit400ResponseOneOf38) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf38 = nil
		} else {
			_ = validator.Validate(dst.RequestSubmit400ResponseOneOf38)
			match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf38 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf39
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf39)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf39, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf39)
		if string(jsonRequestSubmit400ResponseOneOf39) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf39 = nil
		} else {
			_ = validator.Validate(dst.RequestSubmit400ResponseOneOf39)
			match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf39 = nil
	}

	// try to unmarshal data into RequestTemplate400ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestTemplate400ResponseOneOf)
	if err == nil {
		jsonRequestTemplate400ResponseOneOf, _ := json.Marshal(dst.RequestTemplate400ResponseOneOf)
		if string(jsonRequestTemplate400ResponseOneOf) == "{}" { // empty struct
			dst.RequestTemplate400ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.RequestTemplate400ResponseOneOf)
			match++
		}
	} else {
		dst.RequestTemplate400ResponseOneOf = nil
	}

	// try to unmarshal data into RequestTemplate400ResponseOneOf2
	err = utils.NewStrictDecoder(data).Decode(&dst.RequestTemplate400ResponseOneOf2)
	if err == nil {
		jsonRequestTemplate400ResponseOneOf2, _ := json.Marshal(dst.RequestTemplate400ResponseOneOf2)
		if string(jsonRequestTemplate400ResponseOneOf2) == "{}" { // empty struct
			dst.RequestTemplate400ResponseOneOf2 = nil
		} else {
			_ = validator.Validate(dst.RequestTemplate400ResponseOneOf2)
			match++
		}
	} else {
		dst.RequestTemplate400ResponseOneOf2 = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestApprove400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestApprove400Response) MarshalJSON() ([]byte, error) {
	if src.CertificateList400ResponseOneOf != nil {
		return json.Marshal(&src.CertificateList400ResponseOneOf)
	}

	if src.CertificateList400ResponseOneOf1 != nil {
		return json.Marshal(&src.CertificateList400ResponseOneOf1)
	}

	if src.RequestApprove400ResponseOneOf != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf)
	}

	if src.RequestApprove400ResponseOneOf1 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf1)
	}

	if src.RequestApprove400ResponseOneOf10 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf10)
	}

	if src.RequestApprove400ResponseOneOf11 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf11)
	}

	if src.RequestApprove400ResponseOneOf12 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf12)
	}

	if src.RequestApprove400ResponseOneOf13 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf13)
	}

	if src.RequestApprove400ResponseOneOf14 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf14)
	}

	if src.RequestApprove400ResponseOneOf15 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf15)
	}

	if src.RequestApprove400ResponseOneOf16 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf16)
	}

	if src.RequestApprove400ResponseOneOf17 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf17)
	}

	if src.RequestApprove400ResponseOneOf18 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf18)
	}

	if src.RequestApprove400ResponseOneOf19 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf19)
	}

	if src.RequestApprove400ResponseOneOf2 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf2)
	}

	if src.RequestApprove400ResponseOneOf20 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf20)
	}

	if src.RequestApprove400ResponseOneOf21 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf21)
	}

	if src.RequestApprove400ResponseOneOf22 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf22)
	}

	if src.RequestApprove400ResponseOneOf23 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf23)
	}

	if src.RequestApprove400ResponseOneOf24 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf24)
	}

	if src.RequestApprove400ResponseOneOf25 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf25)
	}

	if src.RequestApprove400ResponseOneOf26 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf26)
	}

	if src.RequestApprove400ResponseOneOf27 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf27)
	}

	if src.RequestApprove400ResponseOneOf28 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf28)
	}

	if src.RequestApprove400ResponseOneOf29 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf29)
	}

	if src.RequestApprove400ResponseOneOf3 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf3)
	}

	if src.RequestApprove400ResponseOneOf30 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf30)
	}

	if src.RequestApprove400ResponseOneOf31 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf31)
	}

	if src.RequestApprove400ResponseOneOf32 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf32)
	}

	if src.RequestApprove400ResponseOneOf33 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf33)
	}

	if src.RequestApprove400ResponseOneOf34 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf34)
	}

	if src.RequestApprove400ResponseOneOf4 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf4)
	}

	if src.RequestApprove400ResponseOneOf5 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf5)
	}

	if src.RequestApprove400ResponseOneOf6 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf6)
	}

	if src.RequestApprove400ResponseOneOf7 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf7)
	}

	if src.RequestApprove400ResponseOneOf8 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf8)
	}

	if src.RequestApprove400ResponseOneOf9 != nil {
		return json.Marshal(&src.RequestApprove400ResponseOneOf9)
	}

	if src.RequestCancel400ResponseOneOf != nil {
		return json.Marshal(&src.RequestCancel400ResponseOneOf)
	}

	if src.RequestCancel400ResponseOneOf1 != nil {
		return json.Marshal(&src.RequestCancel400ResponseOneOf1)
	}

	if src.RequestSubmit400ResponseOneOf38 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf38)
	}

	if src.RequestSubmit400ResponseOneOf39 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf39)
	}

	if src.RequestTemplate400ResponseOneOf != nil {
		return json.Marshal(&src.RequestTemplate400ResponseOneOf)
	}

	if src.RequestTemplate400ResponseOneOf2 != nil {
		return json.Marshal(&src.RequestTemplate400ResponseOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestApprove400Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CertificateList400ResponseOneOf != nil {
		return obj.CertificateList400ResponseOneOf
	}

	if obj.CertificateList400ResponseOneOf1 != nil {
		return obj.CertificateList400ResponseOneOf1
	}

	if obj.RequestApprove400ResponseOneOf != nil {
		return obj.RequestApprove400ResponseOneOf
	}

	if obj.RequestApprove400ResponseOneOf1 != nil {
		return obj.RequestApprove400ResponseOneOf1
	}

	if obj.RequestApprove400ResponseOneOf10 != nil {
		return obj.RequestApprove400ResponseOneOf10
	}

	if obj.RequestApprove400ResponseOneOf11 != nil {
		return obj.RequestApprove400ResponseOneOf11
	}

	if obj.RequestApprove400ResponseOneOf12 != nil {
		return obj.RequestApprove400ResponseOneOf12
	}

	if obj.RequestApprove400ResponseOneOf13 != nil {
		return obj.RequestApprove400ResponseOneOf13
	}

	if obj.RequestApprove400ResponseOneOf14 != nil {
		return obj.RequestApprove400ResponseOneOf14
	}

	if obj.RequestApprove400ResponseOneOf15 != nil {
		return obj.RequestApprove400ResponseOneOf15
	}

	if obj.RequestApprove400ResponseOneOf16 != nil {
		return obj.RequestApprove400ResponseOneOf16
	}

	if obj.RequestApprove400ResponseOneOf17 != nil {
		return obj.RequestApprove400ResponseOneOf17
	}

	if obj.RequestApprove400ResponseOneOf18 != nil {
		return obj.RequestApprove400ResponseOneOf18
	}

	if obj.RequestApprove400ResponseOneOf19 != nil {
		return obj.RequestApprove400ResponseOneOf19
	}

	if obj.RequestApprove400ResponseOneOf2 != nil {
		return obj.RequestApprove400ResponseOneOf2
	}

	if obj.RequestApprove400ResponseOneOf20 != nil {
		return obj.RequestApprove400ResponseOneOf20
	}

	if obj.RequestApprove400ResponseOneOf21 != nil {
		return obj.RequestApprove400ResponseOneOf21
	}

	if obj.RequestApprove400ResponseOneOf22 != nil {
		return obj.RequestApprove400ResponseOneOf22
	}

	if obj.RequestApprove400ResponseOneOf23 != nil {
		return obj.RequestApprove400ResponseOneOf23
	}

	if obj.RequestApprove400ResponseOneOf24 != nil {
		return obj.RequestApprove400ResponseOneOf24
	}

	if obj.RequestApprove400ResponseOneOf25 != nil {
		return obj.RequestApprove400ResponseOneOf25
	}

	if obj.RequestApprove400ResponseOneOf26 != nil {
		return obj.RequestApprove400ResponseOneOf26
	}

	if obj.RequestApprove400ResponseOneOf27 != nil {
		return obj.RequestApprove400ResponseOneOf27
	}

	if obj.RequestApprove400ResponseOneOf28 != nil {
		return obj.RequestApprove400ResponseOneOf28
	}

	if obj.RequestApprove400ResponseOneOf29 != nil {
		return obj.RequestApprove400ResponseOneOf29
	}

	if obj.RequestApprove400ResponseOneOf3 != nil {
		return obj.RequestApprove400ResponseOneOf3
	}

	if obj.RequestApprove400ResponseOneOf30 != nil {
		return obj.RequestApprove400ResponseOneOf30
	}

	if obj.RequestApprove400ResponseOneOf31 != nil {
		return obj.RequestApprove400ResponseOneOf31
	}

	if obj.RequestApprove400ResponseOneOf32 != nil {
		return obj.RequestApprove400ResponseOneOf32
	}

	if obj.RequestApprove400ResponseOneOf33 != nil {
		return obj.RequestApprove400ResponseOneOf33
	}

	if obj.RequestApprove400ResponseOneOf34 != nil {
		return obj.RequestApprove400ResponseOneOf34
	}

	if obj.RequestApprove400ResponseOneOf4 != nil {
		return obj.RequestApprove400ResponseOneOf4
	}

	if obj.RequestApprove400ResponseOneOf5 != nil {
		return obj.RequestApprove400ResponseOneOf5
	}

	if obj.RequestApprove400ResponseOneOf6 != nil {
		return obj.RequestApprove400ResponseOneOf6
	}

	if obj.RequestApprove400ResponseOneOf7 != nil {
		return obj.RequestApprove400ResponseOneOf7
	}

	if obj.RequestApprove400ResponseOneOf8 != nil {
		return obj.RequestApprove400ResponseOneOf8
	}

	if obj.RequestApprove400ResponseOneOf9 != nil {
		return obj.RequestApprove400ResponseOneOf9
	}

	if obj.RequestCancel400ResponseOneOf != nil {
		return obj.RequestCancel400ResponseOneOf
	}

	if obj.RequestCancel400ResponseOneOf1 != nil {
		return obj.RequestCancel400ResponseOneOf1
	}

	if obj.RequestSubmit400ResponseOneOf38 != nil {
		return obj.RequestSubmit400ResponseOneOf38
	}

	if obj.RequestSubmit400ResponseOneOf39 != nil {
		return obj.RequestSubmit400ResponseOneOf39
	}

	if obj.RequestTemplate400ResponseOneOf != nil {
		return obj.RequestTemplate400ResponseOneOf
	}

	if obj.RequestTemplate400ResponseOneOf2 != nil {
		return obj.RequestTemplate400ResponseOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableRequestApprove400Response struct {
	value *RequestApprove400Response
	isSet bool
}

func (v NullableRequestApprove400Response) Get() *RequestApprove400Response {
	return v.value
}

func (v *NullableRequestApprove400Response) Set(val *RequestApprove400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprove400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprove400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprove400Response(val *RequestApprove400Response) *NullableRequestApprove400Response {
	return &NullableRequestApprove400Response{value: val, isSet: true}
}

func (v NullableRequestApprove400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprove400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
