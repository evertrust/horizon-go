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

// RequestApprove500Response - struct for RequestApprove500Response
type RequestApprove500Response struct {
	RequestApprove500ResponseOneOf            *RequestApprove500ResponseOneOf
	RequestApprove500ResponseOneOf1           *RequestApprove500ResponseOneOf1
	RequestApprove500ResponseOneOf10          *RequestApprove500ResponseOneOf10
	RequestApprove500ResponseOneOf11          *RequestApprove500ResponseOneOf11
	RequestApprove500ResponseOneOf12          *RequestApprove500ResponseOneOf12
	RequestApprove500ResponseOneOf13          *RequestApprove500ResponseOneOf13
	RequestApprove500ResponseOneOf14          *RequestApprove500ResponseOneOf14
	RequestApprove500ResponseOneOf15          *RequestApprove500ResponseOneOf15
	RequestApprove500ResponseOneOf16          *RequestApprove500ResponseOneOf16
	RequestApprove500ResponseOneOf17          *RequestApprove500ResponseOneOf17
	RequestApprove500ResponseOneOf18          *RequestApprove500ResponseOneOf18
	RequestApprove500ResponseOneOf19          *RequestApprove500ResponseOneOf19
	RequestApprove500ResponseOneOf2           *RequestApprove500ResponseOneOf2
	RequestApprove500ResponseOneOf20          *RequestApprove500ResponseOneOf20
	RequestApprove500ResponseOneOf21          *RequestApprove500ResponseOneOf21
	RequestApprove500ResponseOneOf22          *RequestApprove500ResponseOneOf22
	RequestApprove500ResponseOneOf3           *RequestApprove500ResponseOneOf3
	RequestApprove500ResponseOneOf4           *RequestApprove500ResponseOneOf4
	RequestApprove500ResponseOneOf5           *RequestApprove500ResponseOneOf5
	RequestApprove500ResponseOneOf6           *RequestApprove500ResponseOneOf6
	RequestApprove500ResponseOneOf7           *RequestApprove500ResponseOneOf7
	RequestApprove500ResponseOneOf8           *RequestApprove500ResponseOneOf8
	RequestApprove500ResponseOneOf9           *RequestApprove500ResponseOneOf9
	RequestCertificateProfile500ResponseOneOf *RequestCertificateProfile500ResponseOneOf
}

// RequestApprove500ResponseOneOfAsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOfAsRequestApprove500Response(v *RequestApprove500ResponseOneOf) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf: v,
	}
}

// RequestApprove500ResponseOneOf1AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf1 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf1AsRequestApprove500Response(v *RequestApprove500ResponseOneOf1) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf1: v,
	}
}

// RequestApprove500ResponseOneOf10AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf10 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf10AsRequestApprove500Response(v *RequestApprove500ResponseOneOf10) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf10: v,
	}
}

// RequestApprove500ResponseOneOf11AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf11 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf11AsRequestApprove500Response(v *RequestApprove500ResponseOneOf11) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf11: v,
	}
}

// RequestApprove500ResponseOneOf12AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf12 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf12AsRequestApprove500Response(v *RequestApprove500ResponseOneOf12) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf12: v,
	}
}

// RequestApprove500ResponseOneOf13AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf13 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf13AsRequestApprove500Response(v *RequestApprove500ResponseOneOf13) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf13: v,
	}
}

// RequestApprove500ResponseOneOf14AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf14 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf14AsRequestApprove500Response(v *RequestApprove500ResponseOneOf14) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf14: v,
	}
}

// RequestApprove500ResponseOneOf15AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf15 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf15AsRequestApprove500Response(v *RequestApprove500ResponseOneOf15) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf15: v,
	}
}

// RequestApprove500ResponseOneOf16AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf16 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf16AsRequestApprove500Response(v *RequestApprove500ResponseOneOf16) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf16: v,
	}
}

// RequestApprove500ResponseOneOf17AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf17 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf17AsRequestApprove500Response(v *RequestApprove500ResponseOneOf17) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf17: v,
	}
}

// RequestApprove500ResponseOneOf18AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf18 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf18AsRequestApprove500Response(v *RequestApprove500ResponseOneOf18) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf18: v,
	}
}

// RequestApprove500ResponseOneOf19AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf19 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf19AsRequestApprove500Response(v *RequestApprove500ResponseOneOf19) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf19: v,
	}
}

// RequestApprove500ResponseOneOf2AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf2 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf2AsRequestApprove500Response(v *RequestApprove500ResponseOneOf2) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf2: v,
	}
}

// RequestApprove500ResponseOneOf20AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf20 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf20AsRequestApprove500Response(v *RequestApprove500ResponseOneOf20) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf20: v,
	}
}

// RequestApprove500ResponseOneOf21AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf21 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf21AsRequestApprove500Response(v *RequestApprove500ResponseOneOf21) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf21: v,
	}
}

// RequestApprove500ResponseOneOf22AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf22 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf22AsRequestApprove500Response(v *RequestApprove500ResponseOneOf22) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf22: v,
	}
}

// RequestApprove500ResponseOneOf3AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf3 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf3AsRequestApprove500Response(v *RequestApprove500ResponseOneOf3) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf3: v,
	}
}

// RequestApprove500ResponseOneOf4AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf4 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf4AsRequestApprove500Response(v *RequestApprove500ResponseOneOf4) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf4: v,
	}
}

// RequestApprove500ResponseOneOf5AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf5 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf5AsRequestApprove500Response(v *RequestApprove500ResponseOneOf5) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf5: v,
	}
}

// RequestApprove500ResponseOneOf6AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf6 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf6AsRequestApprove500Response(v *RequestApprove500ResponseOneOf6) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf6: v,
	}
}

// RequestApprove500ResponseOneOf7AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf7 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf7AsRequestApprove500Response(v *RequestApprove500ResponseOneOf7) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf7: v,
	}
}

// RequestApprove500ResponseOneOf8AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf8 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf8AsRequestApprove500Response(v *RequestApprove500ResponseOneOf8) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf8: v,
	}
}

// RequestApprove500ResponseOneOf9AsRequestApprove500Response is a convenience function that returns RequestApprove500ResponseOneOf9 wrapped in RequestApprove500Response
func RequestApprove500ResponseOneOf9AsRequestApprove500Response(v *RequestApprove500ResponseOneOf9) RequestApprove500Response {
	return RequestApprove500Response{
		RequestApprove500ResponseOneOf9: v,
	}
}

// RequestCertificateProfile500ResponseOneOfAsRequestApprove500Response is a convenience function that returns RequestCertificateProfile500ResponseOneOf wrapped in RequestApprove500Response
func RequestCertificateProfile500ResponseOneOfAsRequestApprove500Response(v *RequestCertificateProfile500ResponseOneOf) RequestApprove500Response {
	return RequestApprove500Response{
		RequestCertificateProfile500ResponseOneOf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestApprove500Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RequestApprove500ResponseOneOf
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf)
	if err == nil {
		jsonRequestApprove500ResponseOneOf, _ := json.Marshal(dst.RequestApprove500ResponseOneOf)
		if string(jsonRequestApprove500ResponseOneOf) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf1
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf1)
	if err == nil {
		jsonRequestApprove500ResponseOneOf1, _ := json.Marshal(dst.RequestApprove500ResponseOneOf1)
		if string(jsonRequestApprove500ResponseOneOf1) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf10
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf10)
	if err == nil {
		jsonRequestApprove500ResponseOneOf10, _ := json.Marshal(dst.RequestApprove500ResponseOneOf10)
		if string(jsonRequestApprove500ResponseOneOf10) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf10 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf10 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf11
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf11)
	if err == nil {
		jsonRequestApprove500ResponseOneOf11, _ := json.Marshal(dst.RequestApprove500ResponseOneOf11)
		if string(jsonRequestApprove500ResponseOneOf11) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf11 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf11 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf12
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf12)
	if err == nil {
		jsonRequestApprove500ResponseOneOf12, _ := json.Marshal(dst.RequestApprove500ResponseOneOf12)
		if string(jsonRequestApprove500ResponseOneOf12) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf12 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf12 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf13
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf13)
	if err == nil {
		jsonRequestApprove500ResponseOneOf13, _ := json.Marshal(dst.RequestApprove500ResponseOneOf13)
		if string(jsonRequestApprove500ResponseOneOf13) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf13 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf13 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf14
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf14)
	if err == nil {
		jsonRequestApprove500ResponseOneOf14, _ := json.Marshal(dst.RequestApprove500ResponseOneOf14)
		if string(jsonRequestApprove500ResponseOneOf14) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf14 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf14 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf15
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf15)
	if err == nil {
		jsonRequestApprove500ResponseOneOf15, _ := json.Marshal(dst.RequestApprove500ResponseOneOf15)
		if string(jsonRequestApprove500ResponseOneOf15) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf15 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf15 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf16
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf16)
	if err == nil {
		jsonRequestApprove500ResponseOneOf16, _ := json.Marshal(dst.RequestApprove500ResponseOneOf16)
		if string(jsonRequestApprove500ResponseOneOf16) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf16 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf16 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf17
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf17)
	if err == nil {
		jsonRequestApprove500ResponseOneOf17, _ := json.Marshal(dst.RequestApprove500ResponseOneOf17)
		if string(jsonRequestApprove500ResponseOneOf17) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf17 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf17 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf18
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf18)
	if err == nil {
		jsonRequestApprove500ResponseOneOf18, _ := json.Marshal(dst.RequestApprove500ResponseOneOf18)
		if string(jsonRequestApprove500ResponseOneOf18) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf18 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf18 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf19
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf19)
	if err == nil {
		jsonRequestApprove500ResponseOneOf19, _ := json.Marshal(dst.RequestApprove500ResponseOneOf19)
		if string(jsonRequestApprove500ResponseOneOf19) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf19 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf19 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf2
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf2)
	if err == nil {
		jsonRequestApprove500ResponseOneOf2, _ := json.Marshal(dst.RequestApprove500ResponseOneOf2)
		if string(jsonRequestApprove500ResponseOneOf2) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf2 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf20
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf20)
	if err == nil {
		jsonRequestApprove500ResponseOneOf20, _ := json.Marshal(dst.RequestApprove500ResponseOneOf20)
		if string(jsonRequestApprove500ResponseOneOf20) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf20 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf20 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf21
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf21)
	if err == nil {
		jsonRequestApprove500ResponseOneOf21, _ := json.Marshal(dst.RequestApprove500ResponseOneOf21)
		if string(jsonRequestApprove500ResponseOneOf21) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf21 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf21 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf22
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf22)
	if err == nil {
		jsonRequestApprove500ResponseOneOf22, _ := json.Marshal(dst.RequestApprove500ResponseOneOf22)
		if string(jsonRequestApprove500ResponseOneOf22) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf22 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf22 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf3
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf3)
	if err == nil {
		jsonRequestApprove500ResponseOneOf3, _ := json.Marshal(dst.RequestApprove500ResponseOneOf3)
		if string(jsonRequestApprove500ResponseOneOf3) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf3 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf4
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf4)
	if err == nil {
		jsonRequestApprove500ResponseOneOf4, _ := json.Marshal(dst.RequestApprove500ResponseOneOf4)
		if string(jsonRequestApprove500ResponseOneOf4) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf4 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf5
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf5)
	if err == nil {
		jsonRequestApprove500ResponseOneOf5, _ := json.Marshal(dst.RequestApprove500ResponseOneOf5)
		if string(jsonRequestApprove500ResponseOneOf5) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf5 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf6
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf6)
	if err == nil {
		jsonRequestApprove500ResponseOneOf6, _ := json.Marshal(dst.RequestApprove500ResponseOneOf6)
		if string(jsonRequestApprove500ResponseOneOf6) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf6 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf6 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf7
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf7)
	if err == nil {
		jsonRequestApprove500ResponseOneOf7, _ := json.Marshal(dst.RequestApprove500ResponseOneOf7)
		if string(jsonRequestApprove500ResponseOneOf7) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf7 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf7 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf8
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf8)
	if err == nil {
		jsonRequestApprove500ResponseOneOf8, _ := json.Marshal(dst.RequestApprove500ResponseOneOf8)
		if string(jsonRequestApprove500ResponseOneOf8) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf8 = nil
	}

	// try to unmarshal data into RequestApprove500ResponseOneOf9
	err = json.Unmarshal(data, &dst.RequestApprove500ResponseOneOf9)
	if err == nil {
		jsonRequestApprove500ResponseOneOf9, _ := json.Marshal(dst.RequestApprove500ResponseOneOf9)
		if string(jsonRequestApprove500ResponseOneOf9) == "{}" { // empty struct
			dst.RequestApprove500ResponseOneOf9 = nil
		} else {
			match++
		}
	} else {
		dst.RequestApprove500ResponseOneOf9 = nil
	}

	// try to unmarshal data into RequestCertificateProfile500ResponseOneOf
	err = json.Unmarshal(data, &dst.RequestCertificateProfile500ResponseOneOf)
	if err == nil {
		jsonRequestCertificateProfile500ResponseOneOf, _ := json.Marshal(dst.RequestCertificateProfile500ResponseOneOf)
		if string(jsonRequestCertificateProfile500ResponseOneOf) == "{}" { // empty struct
			dst.RequestCertificateProfile500ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RequestCertificateProfile500ResponseOneOf = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestApprove500Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestApprove500Response) MarshalJSON() ([]byte, error) {
	if src.RequestApprove500ResponseOneOf != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf)
	}

	if src.RequestApprove500ResponseOneOf1 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf1)
	}

	if src.RequestApprove500ResponseOneOf10 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf10)
	}

	if src.RequestApprove500ResponseOneOf11 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf11)
	}

	if src.RequestApprove500ResponseOneOf12 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf12)
	}

	if src.RequestApprove500ResponseOneOf13 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf13)
	}

	if src.RequestApprove500ResponseOneOf14 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf14)
	}

	if src.RequestApprove500ResponseOneOf15 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf15)
	}

	if src.RequestApprove500ResponseOneOf16 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf16)
	}

	if src.RequestApprove500ResponseOneOf17 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf17)
	}

	if src.RequestApprove500ResponseOneOf18 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf18)
	}

	if src.RequestApprove500ResponseOneOf19 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf19)
	}

	if src.RequestApprove500ResponseOneOf2 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf2)
	}

	if src.RequestApprove500ResponseOneOf20 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf20)
	}

	if src.RequestApprove500ResponseOneOf21 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf21)
	}

	if src.RequestApprove500ResponseOneOf22 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf22)
	}

	if src.RequestApprove500ResponseOneOf3 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf3)
	}

	if src.RequestApprove500ResponseOneOf4 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf4)
	}

	if src.RequestApprove500ResponseOneOf5 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf5)
	}

	if src.RequestApprove500ResponseOneOf6 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf6)
	}

	if src.RequestApprove500ResponseOneOf7 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf7)
	}

	if src.RequestApprove500ResponseOneOf8 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf8)
	}

	if src.RequestApprove500ResponseOneOf9 != nil {
		return json.Marshal(&src.RequestApprove500ResponseOneOf9)
	}

	if src.RequestCertificateProfile500ResponseOneOf != nil {
		return json.Marshal(&src.RequestCertificateProfile500ResponseOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestApprove500Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestApprove500ResponseOneOf != nil {
		return obj.RequestApprove500ResponseOneOf
	}

	if obj.RequestApprove500ResponseOneOf1 != nil {
		return obj.RequestApprove500ResponseOneOf1
	}

	if obj.RequestApprove500ResponseOneOf10 != nil {
		return obj.RequestApprove500ResponseOneOf10
	}

	if obj.RequestApprove500ResponseOneOf11 != nil {
		return obj.RequestApprove500ResponseOneOf11
	}

	if obj.RequestApprove500ResponseOneOf12 != nil {
		return obj.RequestApprove500ResponseOneOf12
	}

	if obj.RequestApprove500ResponseOneOf13 != nil {
		return obj.RequestApprove500ResponseOneOf13
	}

	if obj.RequestApprove500ResponseOneOf14 != nil {
		return obj.RequestApprove500ResponseOneOf14
	}

	if obj.RequestApprove500ResponseOneOf15 != nil {
		return obj.RequestApprove500ResponseOneOf15
	}

	if obj.RequestApprove500ResponseOneOf16 != nil {
		return obj.RequestApprove500ResponseOneOf16
	}

	if obj.RequestApprove500ResponseOneOf17 != nil {
		return obj.RequestApprove500ResponseOneOf17
	}

	if obj.RequestApprove500ResponseOneOf18 != nil {
		return obj.RequestApprove500ResponseOneOf18
	}

	if obj.RequestApprove500ResponseOneOf19 != nil {
		return obj.RequestApprove500ResponseOneOf19
	}

	if obj.RequestApprove500ResponseOneOf2 != nil {
		return obj.RequestApprove500ResponseOneOf2
	}

	if obj.RequestApprove500ResponseOneOf20 != nil {
		return obj.RequestApprove500ResponseOneOf20
	}

	if obj.RequestApprove500ResponseOneOf21 != nil {
		return obj.RequestApprove500ResponseOneOf21
	}

	if obj.RequestApprove500ResponseOneOf22 != nil {
		return obj.RequestApprove500ResponseOneOf22
	}

	if obj.RequestApprove500ResponseOneOf3 != nil {
		return obj.RequestApprove500ResponseOneOf3
	}

	if obj.RequestApprove500ResponseOneOf4 != nil {
		return obj.RequestApprove500ResponseOneOf4
	}

	if obj.RequestApprove500ResponseOneOf5 != nil {
		return obj.RequestApprove500ResponseOneOf5
	}

	if obj.RequestApprove500ResponseOneOf6 != nil {
		return obj.RequestApprove500ResponseOneOf6
	}

	if obj.RequestApprove500ResponseOneOf7 != nil {
		return obj.RequestApprove500ResponseOneOf7
	}

	if obj.RequestApprove500ResponseOneOf8 != nil {
		return obj.RequestApprove500ResponseOneOf8
	}

	if obj.RequestApprove500ResponseOneOf9 != nil {
		return obj.RequestApprove500ResponseOneOf9
	}

	if obj.RequestCertificateProfile500ResponseOneOf != nil {
		return obj.RequestCertificateProfile500ResponseOneOf
	}

	// all schemas are nil
	return nil
}

type NullableRequestApprove500Response struct {
	value *RequestApprove500Response
	isSet bool
}

func (v NullableRequestApprove500Response) Get() *RequestApprove500Response {
	return v.value
}

func (v *NullableRequestApprove500Response) Set(val *RequestApprove500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprove500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprove500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprove500Response(val *RequestApprove500Response) *NullableRequestApprove500Response {
	return &NullableRequestApprove500Response{value: val, isSet: true}
}

func (v NullableRequestApprove500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprove500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
