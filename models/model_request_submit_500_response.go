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

// RequestSubmit500Response - struct for RequestSubmit500Response
type RequestSubmit500Response struct {
	RequestSubmit500ResponseOneOf   *RequestSubmit500ResponseOneOf
	RequestSubmit500ResponseOneOf1  *RequestSubmit500ResponseOneOf1
	RequestSubmit500ResponseOneOf10 *RequestSubmit500ResponseOneOf10
	RequestSubmit500ResponseOneOf11 *RequestSubmit500ResponseOneOf11
	RequestSubmit500ResponseOneOf12 *RequestSubmit500ResponseOneOf12
	RequestSubmit500ResponseOneOf13 *RequestSubmit500ResponseOneOf13
	RequestSubmit500ResponseOneOf14 *RequestSubmit500ResponseOneOf14
	RequestSubmit500ResponseOneOf15 *RequestSubmit500ResponseOneOf15
	RequestSubmit500ResponseOneOf16 *RequestSubmit500ResponseOneOf16
	RequestSubmit500ResponseOneOf17 *RequestSubmit500ResponseOneOf17
	RequestSubmit500ResponseOneOf18 *RequestSubmit500ResponseOneOf18
	RequestSubmit500ResponseOneOf19 *RequestSubmit500ResponseOneOf19
	RequestSubmit500ResponseOneOf2  *RequestSubmit500ResponseOneOf2
	RequestSubmit500ResponseOneOf20 *RequestSubmit500ResponseOneOf20
	RequestSubmit500ResponseOneOf21 *RequestSubmit500ResponseOneOf21
	RequestSubmit500ResponseOneOf3  *RequestSubmit500ResponseOneOf3
	RequestSubmit500ResponseOneOf4  *RequestSubmit500ResponseOneOf4
	RequestSubmit500ResponseOneOf5  *RequestSubmit500ResponseOneOf5
	RequestSubmit500ResponseOneOf6  *RequestSubmit500ResponseOneOf6
	RequestSubmit500ResponseOneOf7  *RequestSubmit500ResponseOneOf7
	RequestSubmit500ResponseOneOf8  *RequestSubmit500ResponseOneOf8
	RequestSubmit500ResponseOneOf9  *RequestSubmit500ResponseOneOf9
}

// RequestSubmit500ResponseOneOfAsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOfAsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf: v,
	}
}

// RequestSubmit500ResponseOneOf1AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf1 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf1AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf1) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf1: v,
	}
}

// RequestSubmit500ResponseOneOf10AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf10 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf10AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf10) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf10: v,
	}
}

// RequestSubmit500ResponseOneOf11AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf11 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf11AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf11) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf11: v,
	}
}

// RequestSubmit500ResponseOneOf12AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf12 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf12AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf12) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf12: v,
	}
}

// RequestSubmit500ResponseOneOf13AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf13 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf13AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf13) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf13: v,
	}
}

// RequestSubmit500ResponseOneOf14AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf14 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf14AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf14) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf14: v,
	}
}

// RequestSubmit500ResponseOneOf15AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf15 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf15AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf15) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf15: v,
	}
}

// RequestSubmit500ResponseOneOf16AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf16 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf16AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf16) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf16: v,
	}
}

// RequestSubmit500ResponseOneOf17AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf17 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf17AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf17) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf17: v,
	}
}

// RequestSubmit500ResponseOneOf18AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf18 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf18AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf18) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf18: v,
	}
}

// RequestSubmit500ResponseOneOf19AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf19 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf19AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf19) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf19: v,
	}
}

// RequestSubmit500ResponseOneOf2AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf2 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf2AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf2) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf2: v,
	}
}

// RequestSubmit500ResponseOneOf20AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf20 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf20AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf20) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf20: v,
	}
}

// RequestSubmit500ResponseOneOf21AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf21 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf21AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf21) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf21: v,
	}
}

// RequestSubmit500ResponseOneOf3AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf3 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf3AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf3) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf3: v,
	}
}

// RequestSubmit500ResponseOneOf4AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf4 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf4AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf4) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf4: v,
	}
}

// RequestSubmit500ResponseOneOf5AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf5 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf5AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf5) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf5: v,
	}
}

// RequestSubmit500ResponseOneOf6AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf6 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf6AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf6) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf6: v,
	}
}

// RequestSubmit500ResponseOneOf7AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf7 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf7AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf7) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf7: v,
	}
}

// RequestSubmit500ResponseOneOf8AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf8 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf8AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf8) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf8: v,
	}
}

// RequestSubmit500ResponseOneOf9AsRequestSubmit500Response is a convenience function that returns RequestSubmit500ResponseOneOf9 wrapped in RequestSubmit500Response
func RequestSubmit500ResponseOneOf9AsRequestSubmit500Response(v *RequestSubmit500ResponseOneOf9) RequestSubmit500Response {
	return RequestSubmit500Response{
		RequestSubmit500ResponseOneOf9: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSubmit500Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RequestSubmit500ResponseOneOf
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf)
		if string(jsonRequestSubmit500ResponseOneOf) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf1
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf1)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf1, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf1)
		if string(jsonRequestSubmit500ResponseOneOf1) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf1 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf10
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf10)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf10, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf10)
		if string(jsonRequestSubmit500ResponseOneOf10) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf10 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf10 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf11
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf11)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf11, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf11)
		if string(jsonRequestSubmit500ResponseOneOf11) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf11 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf11 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf12
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf12)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf12, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf12)
		if string(jsonRequestSubmit500ResponseOneOf12) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf12 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf12 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf13
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf13)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf13, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf13)
		if string(jsonRequestSubmit500ResponseOneOf13) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf13 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf13 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf14
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf14)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf14, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf14)
		if string(jsonRequestSubmit500ResponseOneOf14) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf14 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf14 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf15
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf15)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf15, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf15)
		if string(jsonRequestSubmit500ResponseOneOf15) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf15 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf15 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf16
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf16)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf16, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf16)
		if string(jsonRequestSubmit500ResponseOneOf16) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf16 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf16 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf17
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf17)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf17, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf17)
		if string(jsonRequestSubmit500ResponseOneOf17) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf17 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf17 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf18
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf18)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf18, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf18)
		if string(jsonRequestSubmit500ResponseOneOf18) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf18 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf18 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf19
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf19)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf19, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf19)
		if string(jsonRequestSubmit500ResponseOneOf19) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf19 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf19 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf2
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf2)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf2, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf2)
		if string(jsonRequestSubmit500ResponseOneOf2) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf2 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf20
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf20)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf20, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf20)
		if string(jsonRequestSubmit500ResponseOneOf20) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf20 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf20 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf21
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf21)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf21, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf21)
		if string(jsonRequestSubmit500ResponseOneOf21) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf21 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf21 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf3
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf3)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf3, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf3)
		if string(jsonRequestSubmit500ResponseOneOf3) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf3 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf4
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf4)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf4, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf4)
		if string(jsonRequestSubmit500ResponseOneOf4) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf4 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf5
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf5)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf5, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf5)
		if string(jsonRequestSubmit500ResponseOneOf5) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf5 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf6
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf6)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf6, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf6)
		if string(jsonRequestSubmit500ResponseOneOf6) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf6 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf6 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf7
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf7)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf7, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf7)
		if string(jsonRequestSubmit500ResponseOneOf7) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf7 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf7 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf8
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf8)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf8, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf8)
		if string(jsonRequestSubmit500ResponseOneOf8) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf8 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf8 = nil
	}

	// try to unmarshal data into RequestSubmit500ResponseOneOf9
	err = json.Unmarshal(data, &dst.RequestSubmit500ResponseOneOf9)
	if err == nil {
		jsonRequestSubmit500ResponseOneOf9, _ := json.Marshal(dst.RequestSubmit500ResponseOneOf9)
		if string(jsonRequestSubmit500ResponseOneOf9) == "{}" { // empty struct
			dst.RequestSubmit500ResponseOneOf9 = nil
		} else {
			match++
		}
	} else {
		dst.RequestSubmit500ResponseOneOf9 = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestSubmit500Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSubmit500Response) MarshalJSON() ([]byte, error) {
	if src.RequestSubmit500ResponseOneOf != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf)
	}

	if src.RequestSubmit500ResponseOneOf1 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf1)
	}

	if src.RequestSubmit500ResponseOneOf10 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf10)
	}

	if src.RequestSubmit500ResponseOneOf11 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf11)
	}

	if src.RequestSubmit500ResponseOneOf12 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf12)
	}

	if src.RequestSubmit500ResponseOneOf13 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf13)
	}

	if src.RequestSubmit500ResponseOneOf14 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf14)
	}

	if src.RequestSubmit500ResponseOneOf15 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf15)
	}

	if src.RequestSubmit500ResponseOneOf16 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf16)
	}

	if src.RequestSubmit500ResponseOneOf17 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf17)
	}

	if src.RequestSubmit500ResponseOneOf18 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf18)
	}

	if src.RequestSubmit500ResponseOneOf19 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf19)
	}

	if src.RequestSubmit500ResponseOneOf2 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf2)
	}

	if src.RequestSubmit500ResponseOneOf20 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf20)
	}

	if src.RequestSubmit500ResponseOneOf21 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf21)
	}

	if src.RequestSubmit500ResponseOneOf3 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf3)
	}

	if src.RequestSubmit500ResponseOneOf4 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf4)
	}

	if src.RequestSubmit500ResponseOneOf5 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf5)
	}

	if src.RequestSubmit500ResponseOneOf6 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf6)
	}

	if src.RequestSubmit500ResponseOneOf7 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf7)
	}

	if src.RequestSubmit500ResponseOneOf8 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf8)
	}

	if src.RequestSubmit500ResponseOneOf9 != nil {
		return json.Marshal(&src.RequestSubmit500ResponseOneOf9)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSubmit500Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestSubmit500ResponseOneOf != nil {
		return obj.RequestSubmit500ResponseOneOf
	}

	if obj.RequestSubmit500ResponseOneOf1 != nil {
		return obj.RequestSubmit500ResponseOneOf1
	}

	if obj.RequestSubmit500ResponseOneOf10 != nil {
		return obj.RequestSubmit500ResponseOneOf10
	}

	if obj.RequestSubmit500ResponseOneOf11 != nil {
		return obj.RequestSubmit500ResponseOneOf11
	}

	if obj.RequestSubmit500ResponseOneOf12 != nil {
		return obj.RequestSubmit500ResponseOneOf12
	}

	if obj.RequestSubmit500ResponseOneOf13 != nil {
		return obj.RequestSubmit500ResponseOneOf13
	}

	if obj.RequestSubmit500ResponseOneOf14 != nil {
		return obj.RequestSubmit500ResponseOneOf14
	}

	if obj.RequestSubmit500ResponseOneOf15 != nil {
		return obj.RequestSubmit500ResponseOneOf15
	}

	if obj.RequestSubmit500ResponseOneOf16 != nil {
		return obj.RequestSubmit500ResponseOneOf16
	}

	if obj.RequestSubmit500ResponseOneOf17 != nil {
		return obj.RequestSubmit500ResponseOneOf17
	}

	if obj.RequestSubmit500ResponseOneOf18 != nil {
		return obj.RequestSubmit500ResponseOneOf18
	}

	if obj.RequestSubmit500ResponseOneOf19 != nil {
		return obj.RequestSubmit500ResponseOneOf19
	}

	if obj.RequestSubmit500ResponseOneOf2 != nil {
		return obj.RequestSubmit500ResponseOneOf2
	}

	if obj.RequestSubmit500ResponseOneOf20 != nil {
		return obj.RequestSubmit500ResponseOneOf20
	}

	if obj.RequestSubmit500ResponseOneOf21 != nil {
		return obj.RequestSubmit500ResponseOneOf21
	}

	if obj.RequestSubmit500ResponseOneOf3 != nil {
		return obj.RequestSubmit500ResponseOneOf3
	}

	if obj.RequestSubmit500ResponseOneOf4 != nil {
		return obj.RequestSubmit500ResponseOneOf4
	}

	if obj.RequestSubmit500ResponseOneOf5 != nil {
		return obj.RequestSubmit500ResponseOneOf5
	}

	if obj.RequestSubmit500ResponseOneOf6 != nil {
		return obj.RequestSubmit500ResponseOneOf6
	}

	if obj.RequestSubmit500ResponseOneOf7 != nil {
		return obj.RequestSubmit500ResponseOneOf7
	}

	if obj.RequestSubmit500ResponseOneOf8 != nil {
		return obj.RequestSubmit500ResponseOneOf8
	}

	if obj.RequestSubmit500ResponseOneOf9 != nil {
		return obj.RequestSubmit500ResponseOneOf9
	}

	// all schemas are nil
	return nil
}

type NullableRequestSubmit500Response struct {
	value *RequestSubmit500Response
	isSet bool
}

func (v NullableRequestSubmit500Response) Get() *RequestSubmit500Response {
	return v.value
}

func (v *NullableRequestSubmit500Response) Set(val *RequestSubmit500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmit500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmit500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmit500Response(val *RequestSubmit500Response) *NullableRequestSubmit500Response {
	return &NullableRequestSubmit500Response{value: val, isSet: true}
}

func (v NullableRequestSubmit500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmit500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
