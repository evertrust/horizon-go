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
	"gopkg.in/validator.v2"
)

// CacheCrlGet404Response - struct for CacheCrlGet404Response
type CacheCrlGet404Response struct {
	CacheCrlGet404ResponseOneOf *CacheCrlGet404ResponseOneOf
}

// CacheCrlGet404ResponseOneOfAsCacheCrlGet404Response is a convenience function that returns CacheCrlGet404ResponseOneOf wrapped in CacheCrlGet404Response
func CacheCrlGet404ResponseOneOfAsCacheCrlGet404Response(v *CacheCrlGet404ResponseOneOf) CacheCrlGet404Response {
	return CacheCrlGet404Response{
		CacheCrlGet404ResponseOneOf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CacheCrlGet404Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CacheCrlGet404ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.CacheCrlGet404ResponseOneOf)
	if err == nil {
		jsonCacheCrlGet404ResponseOneOf, _ := json.Marshal(dst.CacheCrlGet404ResponseOneOf)
		if string(jsonCacheCrlGet404ResponseOneOf) == "{}" { // empty struct
			dst.CacheCrlGet404ResponseOneOf = nil
		} else {
            _ = validator.Validate(dst.CacheCrlGet404ResponseOneOf)
            match++
		}
	} else {
		dst.CacheCrlGet404ResponseOneOf = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CacheCrlGet404Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CacheCrlGet404Response) MarshalJSON() ([]byte, error) {
	if src.CacheCrlGet404ResponseOneOf != nil {
		return json.Marshal(&src.CacheCrlGet404ResponseOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CacheCrlGet404Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CacheCrlGet404ResponseOneOf != nil {
		return obj.CacheCrlGet404ResponseOneOf
	}

	// all schemas are nil
	return nil
}

type NullableCacheCrlGet404Response struct {
	value *CacheCrlGet404Response
	isSet bool
}

func (v NullableCacheCrlGet404Response) Get() *CacheCrlGet404Response {
	return v.value
}

func (v *NullableCacheCrlGet404Response) Set(val *CacheCrlGet404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheCrlGet404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheCrlGet404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheCrlGet404Response(val *CacheCrlGet404Response) *NullableCacheCrlGet404Response {
	return &NullableCacheCrlGet404Response{value: val, isSet: true}
}

func (v NullableCacheCrlGet404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheCrlGet404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


