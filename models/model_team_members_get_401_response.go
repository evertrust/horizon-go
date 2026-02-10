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

// TeamMembersGet401Response struct for TeamMembersGet401Response
type TeamMembersGet401Response struct {
	AdocGet401ResponseOneOf                  *AdocGet401ResponseOneOf
	AdocGet401ResponseOneOf1                 *AdocGet401ResponseOneOf1
	AdocGet401ResponseOneOf2                 *AdocGet401ResponseOneOf2
	AdocGet401ResponseOneOf3                 *AdocGet401ResponseOneOf3
	AdocGet401ResponseOneOf4                 *AdocGet401ResponseOneOf4
	AdocGet401ResponseOneOf5                 *AdocGet401ResponseOneOf5
	AdocGet401ResponseOneOf6                 *AdocGet401ResponseOneOf6
	AdocGet401ResponseOneOf7                 *AdocGet401ResponseOneOf7
	AdocGet401ResponseOneOf8                 *AdocGet401ResponseOneOf8
	ReportCsvListAllMetadata401ResponseOneOf *ReportCsvListAllMetadata401ResponseOneOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TeamMembersGet401Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AdocGet401ResponseOneOf
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf)
	if err == nil {
		jsonAdocGet401ResponseOneOf, _ := json.Marshal(dst.AdocGet401ResponseOneOf)
		if string(jsonAdocGet401ResponseOneOf) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf1
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf1)
	if err == nil {
		jsonAdocGet401ResponseOneOf1, _ := json.Marshal(dst.AdocGet401ResponseOneOf1)
		if string(jsonAdocGet401ResponseOneOf1) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf1 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf1, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf1 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf2
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf2)
	if err == nil {
		jsonAdocGet401ResponseOneOf2, _ := json.Marshal(dst.AdocGet401ResponseOneOf2)
		if string(jsonAdocGet401ResponseOneOf2) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf2 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf2, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf2 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf3
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf3)
	if err == nil {
		jsonAdocGet401ResponseOneOf3, _ := json.Marshal(dst.AdocGet401ResponseOneOf3)
		if string(jsonAdocGet401ResponseOneOf3) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf3 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf3, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf3 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf4
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf4)
	if err == nil {
		jsonAdocGet401ResponseOneOf4, _ := json.Marshal(dst.AdocGet401ResponseOneOf4)
		if string(jsonAdocGet401ResponseOneOf4) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf4 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf4, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf4 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf5
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf5)
	if err == nil {
		jsonAdocGet401ResponseOneOf5, _ := json.Marshal(dst.AdocGet401ResponseOneOf5)
		if string(jsonAdocGet401ResponseOneOf5) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf5 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf5, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf5 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf6
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf6)
	if err == nil {
		jsonAdocGet401ResponseOneOf6, _ := json.Marshal(dst.AdocGet401ResponseOneOf6)
		if string(jsonAdocGet401ResponseOneOf6) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf6 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf6, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf6 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf7
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf7)
	if err == nil {
		jsonAdocGet401ResponseOneOf7, _ := json.Marshal(dst.AdocGet401ResponseOneOf7)
		if string(jsonAdocGet401ResponseOneOf7) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf7 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf7, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf7 = nil
	}

	// try to unmarshal JSON data into AdocGet401ResponseOneOf8
	err = json.Unmarshal(data, &dst.AdocGet401ResponseOneOf8)
	if err == nil {
		jsonAdocGet401ResponseOneOf8, _ := json.Marshal(dst.AdocGet401ResponseOneOf8)
		if string(jsonAdocGet401ResponseOneOf8) == "{}" { // empty struct
			dst.AdocGet401ResponseOneOf8 = nil
		} else {
			return nil // data stored in dst.AdocGet401ResponseOneOf8, return on the first match
		}
	} else {
		dst.AdocGet401ResponseOneOf8 = nil
	}

	// try to unmarshal JSON data into ReportCsvListAllMetadata401ResponseOneOf
	err = json.Unmarshal(data, &dst.ReportCsvListAllMetadata401ResponseOneOf)
	if err == nil {
		jsonReportCsvListAllMetadata401ResponseOneOf, _ := json.Marshal(dst.ReportCsvListAllMetadata401ResponseOneOf)
		if string(jsonReportCsvListAllMetadata401ResponseOneOf) == "{}" { // empty struct
			dst.ReportCsvListAllMetadata401ResponseOneOf = nil
		} else {
			return nil // data stored in dst.ReportCsvListAllMetadata401ResponseOneOf, return on the first match
		}
	} else {
		dst.ReportCsvListAllMetadata401ResponseOneOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TeamMembersGet401Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TeamMembersGet401Response) MarshalJSON() ([]byte, error) {
	if src.AdocGet401ResponseOneOf != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf)
	}

	if src.AdocGet401ResponseOneOf1 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf1)
	}

	if src.AdocGet401ResponseOneOf2 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf2)
	}

	if src.AdocGet401ResponseOneOf3 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf3)
	}

	if src.AdocGet401ResponseOneOf4 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf4)
	}

	if src.AdocGet401ResponseOneOf5 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf5)
	}

	if src.AdocGet401ResponseOneOf6 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf6)
	}

	if src.AdocGet401ResponseOneOf7 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf7)
	}

	if src.AdocGet401ResponseOneOf8 != nil {
		return json.Marshal(&src.AdocGet401ResponseOneOf8)
	}

	if src.ReportCsvListAllMetadata401ResponseOneOf != nil {
		return json.Marshal(&src.ReportCsvListAllMetadata401ResponseOneOf)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTeamMembersGet401Response struct {
	value *TeamMembersGet401Response
	isSet bool
}

func (v NullableTeamMembersGet401Response) Get() *TeamMembersGet401Response {
	return v.value
}

func (v *NullableTeamMembersGet401Response) Set(val *TeamMembersGet401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMembersGet401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMembersGet401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMembersGet401Response(val *TeamMembersGet401Response) *NullableTeamMembersGet401Response {
	return &NullableTeamMembersGet401Response{value: val, isSet: true}
}

func (v NullableTeamMembersGet401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMembersGet401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
