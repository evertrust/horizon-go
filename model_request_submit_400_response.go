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

// RequestSubmit400Response - struct for RequestSubmit400Response
type RequestSubmit400Response struct {
	CertificateSearch400ResponseOneOf1 *CertificateSearch400ResponseOneOf1
	CertificateSearch400ResponseOneOf2 *CertificateSearch400ResponseOneOf2
	RequestSubmit400ResponseOneOf *RequestSubmit400ResponseOneOf
	RequestSubmit400ResponseOneOf1 *RequestSubmit400ResponseOneOf1
	RequestSubmit400ResponseOneOf10 *RequestSubmit400ResponseOneOf10
	RequestSubmit400ResponseOneOf11 *RequestSubmit400ResponseOneOf11
	RequestSubmit400ResponseOneOf12 *RequestSubmit400ResponseOneOf12
	RequestSubmit400ResponseOneOf13 *RequestSubmit400ResponseOneOf13
	RequestSubmit400ResponseOneOf14 *RequestSubmit400ResponseOneOf14
	RequestSubmit400ResponseOneOf15 *RequestSubmit400ResponseOneOf15
	RequestSubmit400ResponseOneOf16 *RequestSubmit400ResponseOneOf16
	RequestSubmit400ResponseOneOf17 *RequestSubmit400ResponseOneOf17
	RequestSubmit400ResponseOneOf18 *RequestSubmit400ResponseOneOf18
	RequestSubmit400ResponseOneOf19 *RequestSubmit400ResponseOneOf19
	RequestSubmit400ResponseOneOf2 *RequestSubmit400ResponseOneOf2
	RequestSubmit400ResponseOneOf20 *RequestSubmit400ResponseOneOf20
	RequestSubmit400ResponseOneOf21 *RequestSubmit400ResponseOneOf21
	RequestSubmit400ResponseOneOf22 *RequestSubmit400ResponseOneOf22
	RequestSubmit400ResponseOneOf23 *RequestSubmit400ResponseOneOf23
	RequestSubmit400ResponseOneOf24 *RequestSubmit400ResponseOneOf24
	RequestSubmit400ResponseOneOf25 *RequestSubmit400ResponseOneOf25
	RequestSubmit400ResponseOneOf26 *RequestSubmit400ResponseOneOf26
	RequestSubmit400ResponseOneOf27 *RequestSubmit400ResponseOneOf27
	RequestSubmit400ResponseOneOf28 *RequestSubmit400ResponseOneOf28
	RequestSubmit400ResponseOneOf29 *RequestSubmit400ResponseOneOf29
	RequestSubmit400ResponseOneOf3 *RequestSubmit400ResponseOneOf3
	RequestSubmit400ResponseOneOf30 *RequestSubmit400ResponseOneOf30
	RequestSubmit400ResponseOneOf31 *RequestSubmit400ResponseOneOf31
	RequestSubmit400ResponseOneOf32 *RequestSubmit400ResponseOneOf32
	RequestSubmit400ResponseOneOf33 *RequestSubmit400ResponseOneOf33
	RequestSubmit400ResponseOneOf34 *RequestSubmit400ResponseOneOf34
	RequestSubmit400ResponseOneOf35 *RequestSubmit400ResponseOneOf35
	RequestSubmit400ResponseOneOf36 *RequestSubmit400ResponseOneOf36
	RequestSubmit400ResponseOneOf37 *RequestSubmit400ResponseOneOf37
	RequestSubmit400ResponseOneOf38 *RequestSubmit400ResponseOneOf38
	RequestSubmit400ResponseOneOf39 *RequestSubmit400ResponseOneOf39
	RequestSubmit400ResponseOneOf4 *RequestSubmit400ResponseOneOf4
	RequestSubmit400ResponseOneOf40 *RequestSubmit400ResponseOneOf40
	RequestSubmit400ResponseOneOf41 *RequestSubmit400ResponseOneOf41
	RequestSubmit400ResponseOneOf5 *RequestSubmit400ResponseOneOf5
	RequestSubmit400ResponseOneOf6 *RequestSubmit400ResponseOneOf6
	RequestSubmit400ResponseOneOf7 *RequestSubmit400ResponseOneOf7
	RequestSubmit400ResponseOneOf8 *RequestSubmit400ResponseOneOf8
	RequestSubmit400ResponseOneOf9 *RequestSubmit400ResponseOneOf9
	RequestTemplate400ResponseOneOf3 *RequestTemplate400ResponseOneOf3
}

// CertificateSearch400ResponseOneOf1AsRequestSubmit400Response is a convenience function that returns CertificateSearch400ResponseOneOf1 wrapped in RequestSubmit400Response
func CertificateSearch400ResponseOneOf1AsRequestSubmit400Response(v *CertificateSearch400ResponseOneOf1) RequestSubmit400Response {
	return RequestSubmit400Response{
		CertificateSearch400ResponseOneOf1: v,
	}
}

// CertificateSearch400ResponseOneOf2AsRequestSubmit400Response is a convenience function that returns CertificateSearch400ResponseOneOf2 wrapped in RequestSubmit400Response
func CertificateSearch400ResponseOneOf2AsRequestSubmit400Response(v *CertificateSearch400ResponseOneOf2) RequestSubmit400Response {
	return RequestSubmit400Response{
		CertificateSearch400ResponseOneOf2: v,
	}
}

// RequestSubmit400ResponseOneOfAsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOfAsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf: v,
	}
}

// RequestSubmit400ResponseOneOf1AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf1 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf1AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf1) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf1: v,
	}
}

// RequestSubmit400ResponseOneOf10AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf10 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf10AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf10) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf10: v,
	}
}

// RequestSubmit400ResponseOneOf11AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf11 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf11AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf11) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf11: v,
	}
}

// RequestSubmit400ResponseOneOf12AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf12 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf12AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf12) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf12: v,
	}
}

// RequestSubmit400ResponseOneOf13AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf13 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf13AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf13) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf13: v,
	}
}

// RequestSubmit400ResponseOneOf14AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf14 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf14AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf14) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf14: v,
	}
}

// RequestSubmit400ResponseOneOf15AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf15 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf15AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf15) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf15: v,
	}
}

// RequestSubmit400ResponseOneOf16AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf16 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf16AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf16) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf16: v,
	}
}

// RequestSubmit400ResponseOneOf17AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf17 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf17AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf17) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf17: v,
	}
}

// RequestSubmit400ResponseOneOf18AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf18 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf18AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf18) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf18: v,
	}
}

// RequestSubmit400ResponseOneOf19AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf19 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf19AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf19) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf19: v,
	}
}

// RequestSubmit400ResponseOneOf2AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf2 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf2AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf2) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf2: v,
	}
}

// RequestSubmit400ResponseOneOf20AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf20 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf20AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf20) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf20: v,
	}
}

// RequestSubmit400ResponseOneOf21AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf21 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf21AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf21) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf21: v,
	}
}

// RequestSubmit400ResponseOneOf22AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf22 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf22AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf22) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf22: v,
	}
}

// RequestSubmit400ResponseOneOf23AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf23 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf23AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf23) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf23: v,
	}
}

// RequestSubmit400ResponseOneOf24AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf24 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf24AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf24) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf24: v,
	}
}

// RequestSubmit400ResponseOneOf25AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf25 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf25AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf25) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf25: v,
	}
}

// RequestSubmit400ResponseOneOf26AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf26 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf26AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf26) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf26: v,
	}
}

// RequestSubmit400ResponseOneOf27AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf27 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf27AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf27) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf27: v,
	}
}

// RequestSubmit400ResponseOneOf28AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf28 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf28AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf28) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf28: v,
	}
}

// RequestSubmit400ResponseOneOf29AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf29 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf29AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf29) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf29: v,
	}
}

// RequestSubmit400ResponseOneOf3AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf3 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf3AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf3) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf3: v,
	}
}

// RequestSubmit400ResponseOneOf30AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf30 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf30AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf30) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf30: v,
	}
}

// RequestSubmit400ResponseOneOf31AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf31 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf31AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf31) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf31: v,
	}
}

// RequestSubmit400ResponseOneOf32AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf32 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf32AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf32) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf32: v,
	}
}

// RequestSubmit400ResponseOneOf33AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf33 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf33AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf33) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf33: v,
	}
}

// RequestSubmit400ResponseOneOf34AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf34 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf34AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf34) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf34: v,
	}
}

// RequestSubmit400ResponseOneOf35AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf35 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf35AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf35) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf35: v,
	}
}

// RequestSubmit400ResponseOneOf36AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf36 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf36AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf36) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf36: v,
	}
}

// RequestSubmit400ResponseOneOf37AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf37 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf37AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf37) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf37: v,
	}
}

// RequestSubmit400ResponseOneOf38AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf38 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf38AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf38) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf38: v,
	}
}

// RequestSubmit400ResponseOneOf39AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf39 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf39AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf39) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf39: v,
	}
}

// RequestSubmit400ResponseOneOf4AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf4 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf4AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf4) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf4: v,
	}
}

// RequestSubmit400ResponseOneOf40AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf40 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf40AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf40) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf40: v,
	}
}

// RequestSubmit400ResponseOneOf41AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf41 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf41AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf41) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf41: v,
	}
}

// RequestSubmit400ResponseOneOf5AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf5 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf5AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf5) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf5: v,
	}
}

// RequestSubmit400ResponseOneOf6AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf6 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf6AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf6) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf6: v,
	}
}

// RequestSubmit400ResponseOneOf7AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf7 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf7AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf7) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf7: v,
	}
}

// RequestSubmit400ResponseOneOf8AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf8 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf8AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf8) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf8: v,
	}
}

// RequestSubmit400ResponseOneOf9AsRequestSubmit400Response is a convenience function that returns RequestSubmit400ResponseOneOf9 wrapped in RequestSubmit400Response
func RequestSubmit400ResponseOneOf9AsRequestSubmit400Response(v *RequestSubmit400ResponseOneOf9) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestSubmit400ResponseOneOf9: v,
	}
}

// RequestTemplate400ResponseOneOf3AsRequestSubmit400Response is a convenience function that returns RequestTemplate400ResponseOneOf3 wrapped in RequestSubmit400Response
func RequestTemplate400ResponseOneOf3AsRequestSubmit400Response(v *RequestTemplate400ResponseOneOf3) RequestSubmit400Response {
	return RequestSubmit400Response{
		RequestTemplate400ResponseOneOf3: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSubmit400Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CertificateSearch400ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.CertificateSearch400ResponseOneOf1)
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
	err = newStrictDecoder(data).Decode(&dst.CertificateSearch400ResponseOneOf2)
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

	// try to unmarshal data into RequestSubmit400ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf)
		if string(jsonRequestSubmit400ResponseOneOf) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf1)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf1, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf1)
		if string(jsonRequestSubmit400ResponseOneOf1) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf1 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf1)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf10
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf10)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf10, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf10)
		if string(jsonRequestSubmit400ResponseOneOf10) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf10 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf10)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf10 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf11
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf11)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf11, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf11)
		if string(jsonRequestSubmit400ResponseOneOf11) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf11 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf11)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf11 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf12
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf12)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf12, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf12)
		if string(jsonRequestSubmit400ResponseOneOf12) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf12 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf12)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf12 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf13
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf13)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf13, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf13)
		if string(jsonRequestSubmit400ResponseOneOf13) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf13 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf13)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf13 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf14
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf14)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf14, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf14)
		if string(jsonRequestSubmit400ResponseOneOf14) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf14 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf14)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf14 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf15
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf15)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf15, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf15)
		if string(jsonRequestSubmit400ResponseOneOf15) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf15 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf15)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf15 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf16
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf16)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf16, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf16)
		if string(jsonRequestSubmit400ResponseOneOf16) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf16 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf16)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf16 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf17
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf17)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf17, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf17)
		if string(jsonRequestSubmit400ResponseOneOf17) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf17 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf17)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf17 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf18
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf18)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf18, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf18)
		if string(jsonRequestSubmit400ResponseOneOf18) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf18 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf18)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf18 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf19
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf19)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf19, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf19)
		if string(jsonRequestSubmit400ResponseOneOf19) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf19 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf19)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf19 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf2
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf2)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf2, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf2)
		if string(jsonRequestSubmit400ResponseOneOf2) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf2 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf2)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf2 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf20
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf20)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf20, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf20)
		if string(jsonRequestSubmit400ResponseOneOf20) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf20 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf20)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf20 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf21
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf21)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf21, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf21)
		if string(jsonRequestSubmit400ResponseOneOf21) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf21 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf21)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf21 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf22
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf22)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf22, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf22)
		if string(jsonRequestSubmit400ResponseOneOf22) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf22 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf22)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf22 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf23
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf23)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf23, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf23)
		if string(jsonRequestSubmit400ResponseOneOf23) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf23 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf23)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf23 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf24
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf24)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf24, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf24)
		if string(jsonRequestSubmit400ResponseOneOf24) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf24 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf24)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf24 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf25
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf25)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf25, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf25)
		if string(jsonRequestSubmit400ResponseOneOf25) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf25 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf25)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf25 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf26
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf26)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf26, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf26)
		if string(jsonRequestSubmit400ResponseOneOf26) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf26 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf26)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf26 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf27
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf27)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf27, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf27)
		if string(jsonRequestSubmit400ResponseOneOf27) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf27 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf27)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf27 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf28
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf28)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf28, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf28)
		if string(jsonRequestSubmit400ResponseOneOf28) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf28 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf28)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf28 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf29
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf29)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf29, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf29)
		if string(jsonRequestSubmit400ResponseOneOf29) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf29 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf29)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf29 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf3
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf3)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf3, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf3)
		if string(jsonRequestSubmit400ResponseOneOf3) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf3 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf3)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf3 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf30
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf30)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf30, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf30)
		if string(jsonRequestSubmit400ResponseOneOf30) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf30 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf30)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf30 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf31
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf31)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf31, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf31)
		if string(jsonRequestSubmit400ResponseOneOf31) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf31 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf31)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf31 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf32
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf32)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf32, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf32)
		if string(jsonRequestSubmit400ResponseOneOf32) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf32 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf32)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf32 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf33
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf33)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf33, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf33)
		if string(jsonRequestSubmit400ResponseOneOf33) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf33 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf33)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf33 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf34
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf34)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf34, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf34)
		if string(jsonRequestSubmit400ResponseOneOf34) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf34 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf34)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf34 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf35
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf35)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf35, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf35)
		if string(jsonRequestSubmit400ResponseOneOf35) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf35 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf35)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf35 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf36
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf36)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf36, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf36)
		if string(jsonRequestSubmit400ResponseOneOf36) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf36 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf36)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf36 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf37
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf37)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf37, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf37)
		if string(jsonRequestSubmit400ResponseOneOf37) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf37 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf37)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf37 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf38
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf38)
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
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf39)
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

	// try to unmarshal data into RequestSubmit400ResponseOneOf4
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf4)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf4, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf4)
		if string(jsonRequestSubmit400ResponseOneOf4) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf4 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf4)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf4 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf40
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf40)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf40, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf40)
		if string(jsonRequestSubmit400ResponseOneOf40) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf40 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf40)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf40 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf41
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf41)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf41, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf41)
		if string(jsonRequestSubmit400ResponseOneOf41) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf41 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf41)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf41 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf5
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf5)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf5, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf5)
		if string(jsonRequestSubmit400ResponseOneOf5) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf5 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf5)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf5 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf6
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf6)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf6, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf6)
		if string(jsonRequestSubmit400ResponseOneOf6) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf6 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf6)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf6 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf7
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf7)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf7, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf7)
		if string(jsonRequestSubmit400ResponseOneOf7) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf7 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf7)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf7 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf8
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf8)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf8, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf8)
		if string(jsonRequestSubmit400ResponseOneOf8) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf8 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf8)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf8 = nil
	}

	// try to unmarshal data into RequestSubmit400ResponseOneOf9
	err = newStrictDecoder(data).Decode(&dst.RequestSubmit400ResponseOneOf9)
	if err == nil {
		jsonRequestSubmit400ResponseOneOf9, _ := json.Marshal(dst.RequestSubmit400ResponseOneOf9)
		if string(jsonRequestSubmit400ResponseOneOf9) == "{}" { // empty struct
			dst.RequestSubmit400ResponseOneOf9 = nil
		} else {
            _ = validator.Validate(dst.RequestSubmit400ResponseOneOf9)
            match++
		}
	} else {
		dst.RequestSubmit400ResponseOneOf9 = nil
	}

	// try to unmarshal data into RequestTemplate400ResponseOneOf3
	err = newStrictDecoder(data).Decode(&dst.RequestTemplate400ResponseOneOf3)
	if err == nil {
		jsonRequestTemplate400ResponseOneOf3, _ := json.Marshal(dst.RequestTemplate400ResponseOneOf3)
		if string(jsonRequestTemplate400ResponseOneOf3) == "{}" { // empty struct
			dst.RequestTemplate400ResponseOneOf3 = nil
		} else {
            _ = validator.Validate(dst.RequestTemplate400ResponseOneOf3)
            match++
		}
	} else {
		dst.RequestTemplate400ResponseOneOf3 = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestSubmit400Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSubmit400Response) MarshalJSON() ([]byte, error) {
	if src.CertificateSearch400ResponseOneOf1 != nil {
		return json.Marshal(&src.CertificateSearch400ResponseOneOf1)
	}

	if src.CertificateSearch400ResponseOneOf2 != nil {
		return json.Marshal(&src.CertificateSearch400ResponseOneOf2)
	}

	if src.RequestSubmit400ResponseOneOf != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf)
	}

	if src.RequestSubmit400ResponseOneOf1 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf1)
	}

	if src.RequestSubmit400ResponseOneOf10 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf10)
	}

	if src.RequestSubmit400ResponseOneOf11 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf11)
	}

	if src.RequestSubmit400ResponseOneOf12 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf12)
	}

	if src.RequestSubmit400ResponseOneOf13 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf13)
	}

	if src.RequestSubmit400ResponseOneOf14 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf14)
	}

	if src.RequestSubmit400ResponseOneOf15 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf15)
	}

	if src.RequestSubmit400ResponseOneOf16 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf16)
	}

	if src.RequestSubmit400ResponseOneOf17 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf17)
	}

	if src.RequestSubmit400ResponseOneOf18 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf18)
	}

	if src.RequestSubmit400ResponseOneOf19 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf19)
	}

	if src.RequestSubmit400ResponseOneOf2 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf2)
	}

	if src.RequestSubmit400ResponseOneOf20 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf20)
	}

	if src.RequestSubmit400ResponseOneOf21 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf21)
	}

	if src.RequestSubmit400ResponseOneOf22 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf22)
	}

	if src.RequestSubmit400ResponseOneOf23 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf23)
	}

	if src.RequestSubmit400ResponseOneOf24 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf24)
	}

	if src.RequestSubmit400ResponseOneOf25 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf25)
	}

	if src.RequestSubmit400ResponseOneOf26 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf26)
	}

	if src.RequestSubmit400ResponseOneOf27 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf27)
	}

	if src.RequestSubmit400ResponseOneOf28 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf28)
	}

	if src.RequestSubmit400ResponseOneOf29 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf29)
	}

	if src.RequestSubmit400ResponseOneOf3 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf3)
	}

	if src.RequestSubmit400ResponseOneOf30 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf30)
	}

	if src.RequestSubmit400ResponseOneOf31 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf31)
	}

	if src.RequestSubmit400ResponseOneOf32 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf32)
	}

	if src.RequestSubmit400ResponseOneOf33 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf33)
	}

	if src.RequestSubmit400ResponseOneOf34 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf34)
	}

	if src.RequestSubmit400ResponseOneOf35 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf35)
	}

	if src.RequestSubmit400ResponseOneOf36 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf36)
	}

	if src.RequestSubmit400ResponseOneOf37 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf37)
	}

	if src.RequestSubmit400ResponseOneOf38 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf38)
	}

	if src.RequestSubmit400ResponseOneOf39 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf39)
	}

	if src.RequestSubmit400ResponseOneOf4 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf4)
	}

	if src.RequestSubmit400ResponseOneOf40 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf40)
	}

	if src.RequestSubmit400ResponseOneOf41 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf41)
	}

	if src.RequestSubmit400ResponseOneOf5 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf5)
	}

	if src.RequestSubmit400ResponseOneOf6 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf6)
	}

	if src.RequestSubmit400ResponseOneOf7 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf7)
	}

	if src.RequestSubmit400ResponseOneOf8 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf8)
	}

	if src.RequestSubmit400ResponseOneOf9 != nil {
		return json.Marshal(&src.RequestSubmit400ResponseOneOf9)
	}

	if src.RequestTemplate400ResponseOneOf3 != nil {
		return json.Marshal(&src.RequestTemplate400ResponseOneOf3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSubmit400Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CertificateSearch400ResponseOneOf1 != nil {
		return obj.CertificateSearch400ResponseOneOf1
	}

	if obj.CertificateSearch400ResponseOneOf2 != nil {
		return obj.CertificateSearch400ResponseOneOf2
	}

	if obj.RequestSubmit400ResponseOneOf != nil {
		return obj.RequestSubmit400ResponseOneOf
	}

	if obj.RequestSubmit400ResponseOneOf1 != nil {
		return obj.RequestSubmit400ResponseOneOf1
	}

	if obj.RequestSubmit400ResponseOneOf10 != nil {
		return obj.RequestSubmit400ResponseOneOf10
	}

	if obj.RequestSubmit400ResponseOneOf11 != nil {
		return obj.RequestSubmit400ResponseOneOf11
	}

	if obj.RequestSubmit400ResponseOneOf12 != nil {
		return obj.RequestSubmit400ResponseOneOf12
	}

	if obj.RequestSubmit400ResponseOneOf13 != nil {
		return obj.RequestSubmit400ResponseOneOf13
	}

	if obj.RequestSubmit400ResponseOneOf14 != nil {
		return obj.RequestSubmit400ResponseOneOf14
	}

	if obj.RequestSubmit400ResponseOneOf15 != nil {
		return obj.RequestSubmit400ResponseOneOf15
	}

	if obj.RequestSubmit400ResponseOneOf16 != nil {
		return obj.RequestSubmit400ResponseOneOf16
	}

	if obj.RequestSubmit400ResponseOneOf17 != nil {
		return obj.RequestSubmit400ResponseOneOf17
	}

	if obj.RequestSubmit400ResponseOneOf18 != nil {
		return obj.RequestSubmit400ResponseOneOf18
	}

	if obj.RequestSubmit400ResponseOneOf19 != nil {
		return obj.RequestSubmit400ResponseOneOf19
	}

	if obj.RequestSubmit400ResponseOneOf2 != nil {
		return obj.RequestSubmit400ResponseOneOf2
	}

	if obj.RequestSubmit400ResponseOneOf20 != nil {
		return obj.RequestSubmit400ResponseOneOf20
	}

	if obj.RequestSubmit400ResponseOneOf21 != nil {
		return obj.RequestSubmit400ResponseOneOf21
	}

	if obj.RequestSubmit400ResponseOneOf22 != nil {
		return obj.RequestSubmit400ResponseOneOf22
	}

	if obj.RequestSubmit400ResponseOneOf23 != nil {
		return obj.RequestSubmit400ResponseOneOf23
	}

	if obj.RequestSubmit400ResponseOneOf24 != nil {
		return obj.RequestSubmit400ResponseOneOf24
	}

	if obj.RequestSubmit400ResponseOneOf25 != nil {
		return obj.RequestSubmit400ResponseOneOf25
	}

	if obj.RequestSubmit400ResponseOneOf26 != nil {
		return obj.RequestSubmit400ResponseOneOf26
	}

	if obj.RequestSubmit400ResponseOneOf27 != nil {
		return obj.RequestSubmit400ResponseOneOf27
	}

	if obj.RequestSubmit400ResponseOneOf28 != nil {
		return obj.RequestSubmit400ResponseOneOf28
	}

	if obj.RequestSubmit400ResponseOneOf29 != nil {
		return obj.RequestSubmit400ResponseOneOf29
	}

	if obj.RequestSubmit400ResponseOneOf3 != nil {
		return obj.RequestSubmit400ResponseOneOf3
	}

	if obj.RequestSubmit400ResponseOneOf30 != nil {
		return obj.RequestSubmit400ResponseOneOf30
	}

	if obj.RequestSubmit400ResponseOneOf31 != nil {
		return obj.RequestSubmit400ResponseOneOf31
	}

	if obj.RequestSubmit400ResponseOneOf32 != nil {
		return obj.RequestSubmit400ResponseOneOf32
	}

	if obj.RequestSubmit400ResponseOneOf33 != nil {
		return obj.RequestSubmit400ResponseOneOf33
	}

	if obj.RequestSubmit400ResponseOneOf34 != nil {
		return obj.RequestSubmit400ResponseOneOf34
	}

	if obj.RequestSubmit400ResponseOneOf35 != nil {
		return obj.RequestSubmit400ResponseOneOf35
	}

	if obj.RequestSubmit400ResponseOneOf36 != nil {
		return obj.RequestSubmit400ResponseOneOf36
	}

	if obj.RequestSubmit400ResponseOneOf37 != nil {
		return obj.RequestSubmit400ResponseOneOf37
	}

	if obj.RequestSubmit400ResponseOneOf38 != nil {
		return obj.RequestSubmit400ResponseOneOf38
	}

	if obj.RequestSubmit400ResponseOneOf39 != nil {
		return obj.RequestSubmit400ResponseOneOf39
	}

	if obj.RequestSubmit400ResponseOneOf4 != nil {
		return obj.RequestSubmit400ResponseOneOf4
	}

	if obj.RequestSubmit400ResponseOneOf40 != nil {
		return obj.RequestSubmit400ResponseOneOf40
	}

	if obj.RequestSubmit400ResponseOneOf41 != nil {
		return obj.RequestSubmit400ResponseOneOf41
	}

	if obj.RequestSubmit400ResponseOneOf5 != nil {
		return obj.RequestSubmit400ResponseOneOf5
	}

	if obj.RequestSubmit400ResponseOneOf6 != nil {
		return obj.RequestSubmit400ResponseOneOf6
	}

	if obj.RequestSubmit400ResponseOneOf7 != nil {
		return obj.RequestSubmit400ResponseOneOf7
	}

	if obj.RequestSubmit400ResponseOneOf8 != nil {
		return obj.RequestSubmit400ResponseOneOf8
	}

	if obj.RequestSubmit400ResponseOneOf9 != nil {
		return obj.RequestSubmit400ResponseOneOf9
	}

	if obj.RequestTemplate400ResponseOneOf3 != nil {
		return obj.RequestTemplate400ResponseOneOf3
	}

	// all schemas are nil
	return nil
}

type NullableRequestSubmit400Response struct {
	value *RequestSubmit400Response
	isSet bool
}

func (v NullableRequestSubmit400Response) Get() *RequestSubmit400Response {
	return v.value
}

func (v *NullableRequestSubmit400Response) Set(val *RequestSubmit400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmit400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmit400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmit400Response(val *RequestSubmit400Response) *NullableRequestSubmit400Response {
	return &NullableRequestSubmit400Response{value: val, isSet: true}
}

func (v NullableRequestSubmit400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmit400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


