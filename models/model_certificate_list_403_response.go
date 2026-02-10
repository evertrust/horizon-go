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

// CertificateList403Response - struct for CertificateList403Response
type CertificateList403Response struct {
	AdocGet403ResponseOneOf1                         *AdocGet403ResponseOneOf1
	AdocGet403ResponseOneOf2                         *AdocGet403ResponseOneOf2
	AutomationPolicyLifecycleVerify403ResponseOneOf  *AutomationPolicyLifecycleVerify403ResponseOneOf
	AutomationPolicyLifecycleVerify403ResponseOneOf1 *AutomationPolicyLifecycleVerify403ResponseOneOf1
	AutomationPolicyLifecycleVerify403ResponseOneOf2 *AutomationPolicyLifecycleVerify403ResponseOneOf2
	AutomationPolicyLifecycleVerify403ResponseOneOf3 *AutomationPolicyLifecycleVerify403ResponseOneOf3
}

// AdocGet403ResponseOneOf1AsCertificateList403Response is a convenience function that returns AdocGet403ResponseOneOf1 wrapped in CertificateList403Response
func AdocGet403ResponseOneOf1AsCertificateList403Response(v *AdocGet403ResponseOneOf1) CertificateList403Response {
	return CertificateList403Response{
		AdocGet403ResponseOneOf1: v,
	}
}

// AdocGet403ResponseOneOf2AsCertificateList403Response is a convenience function that returns AdocGet403ResponseOneOf2 wrapped in CertificateList403Response
func AdocGet403ResponseOneOf2AsCertificateList403Response(v *AdocGet403ResponseOneOf2) CertificateList403Response {
	return CertificateList403Response{
		AdocGet403ResponseOneOf2: v,
	}
}

// AutomationPolicyLifecycleVerify403ResponseOneOfAsCertificateList403Response is a convenience function that returns AutomationPolicyLifecycleVerify403ResponseOneOf wrapped in CertificateList403Response
func AutomationPolicyLifecycleVerify403ResponseOneOfAsCertificateList403Response(v *AutomationPolicyLifecycleVerify403ResponseOneOf) CertificateList403Response {
	return CertificateList403Response{
		AutomationPolicyLifecycleVerify403ResponseOneOf: v,
	}
}

// AutomationPolicyLifecycleVerify403ResponseOneOf1AsCertificateList403Response is a convenience function that returns AutomationPolicyLifecycleVerify403ResponseOneOf1 wrapped in CertificateList403Response
func AutomationPolicyLifecycleVerify403ResponseOneOf1AsCertificateList403Response(v *AutomationPolicyLifecycleVerify403ResponseOneOf1) CertificateList403Response {
	return CertificateList403Response{
		AutomationPolicyLifecycleVerify403ResponseOneOf1: v,
	}
}

// AutomationPolicyLifecycleVerify403ResponseOneOf2AsCertificateList403Response is a convenience function that returns AutomationPolicyLifecycleVerify403ResponseOneOf2 wrapped in CertificateList403Response
func AutomationPolicyLifecycleVerify403ResponseOneOf2AsCertificateList403Response(v *AutomationPolicyLifecycleVerify403ResponseOneOf2) CertificateList403Response {
	return CertificateList403Response{
		AutomationPolicyLifecycleVerify403ResponseOneOf2: v,
	}
}

// AutomationPolicyLifecycleVerify403ResponseOneOf3AsCertificateList403Response is a convenience function that returns AutomationPolicyLifecycleVerify403ResponseOneOf3 wrapped in CertificateList403Response
func AutomationPolicyLifecycleVerify403ResponseOneOf3AsCertificateList403Response(v *AutomationPolicyLifecycleVerify403ResponseOneOf3) CertificateList403Response {
	return CertificateList403Response{
		AutomationPolicyLifecycleVerify403ResponseOneOf3: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CertificateList403Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AdocGet403ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.AdocGet403ResponseOneOf1)
	if err == nil {
		jsonAdocGet403ResponseOneOf1, _ := json.Marshal(dst.AdocGet403ResponseOneOf1)
		if string(jsonAdocGet403ResponseOneOf1) == "{}" { // empty struct
			dst.AdocGet403ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.AdocGet403ResponseOneOf1)
			match++
		}
	} else {
		dst.AdocGet403ResponseOneOf1 = nil
	}

	// try to unmarshal data into AdocGet403ResponseOneOf2
	err = utils.NewStrictDecoder(data).Decode(&dst.AdocGet403ResponseOneOf2)
	if err == nil {
		jsonAdocGet403ResponseOneOf2, _ := json.Marshal(dst.AdocGet403ResponseOneOf2)
		if string(jsonAdocGet403ResponseOneOf2) == "{}" { // empty struct
			dst.AdocGet403ResponseOneOf2 = nil
		} else {
			_ = validator.Validate(dst.AdocGet403ResponseOneOf2)
			match++
		}
	} else {
		dst.AdocGet403ResponseOneOf2 = nil
	}

	// try to unmarshal data into AutomationPolicyLifecycleVerify403ResponseOneOf
	err = utils.NewStrictDecoder(data).Decode(&dst.AutomationPolicyLifecycleVerify403ResponseOneOf)
	if err == nil {
		jsonAutomationPolicyLifecycleVerify403ResponseOneOf, _ := json.Marshal(dst.AutomationPolicyLifecycleVerify403ResponseOneOf)
		if string(jsonAutomationPolicyLifecycleVerify403ResponseOneOf) == "{}" { // empty struct
			dst.AutomationPolicyLifecycleVerify403ResponseOneOf = nil
		} else {
			_ = validator.Validate(dst.AutomationPolicyLifecycleVerify403ResponseOneOf)
			match++
		}
	} else {
		dst.AutomationPolicyLifecycleVerify403ResponseOneOf = nil
	}

	// try to unmarshal data into AutomationPolicyLifecycleVerify403ResponseOneOf1
	err = utils.NewStrictDecoder(data).Decode(&dst.AutomationPolicyLifecycleVerify403ResponseOneOf1)
	if err == nil {
		jsonAutomationPolicyLifecycleVerify403ResponseOneOf1, _ := json.Marshal(dst.AutomationPolicyLifecycleVerify403ResponseOneOf1)
		if string(jsonAutomationPolicyLifecycleVerify403ResponseOneOf1) == "{}" { // empty struct
			dst.AutomationPolicyLifecycleVerify403ResponseOneOf1 = nil
		} else {
			_ = validator.Validate(dst.AutomationPolicyLifecycleVerify403ResponseOneOf1)
			match++
		}
	} else {
		dst.AutomationPolicyLifecycleVerify403ResponseOneOf1 = nil
	}

	// try to unmarshal data into AutomationPolicyLifecycleVerify403ResponseOneOf2
	err = utils.NewStrictDecoder(data).Decode(&dst.AutomationPolicyLifecycleVerify403ResponseOneOf2)
	if err == nil {
		jsonAutomationPolicyLifecycleVerify403ResponseOneOf2, _ := json.Marshal(dst.AutomationPolicyLifecycleVerify403ResponseOneOf2)
		if string(jsonAutomationPolicyLifecycleVerify403ResponseOneOf2) == "{}" { // empty struct
			dst.AutomationPolicyLifecycleVerify403ResponseOneOf2 = nil
		} else {
			_ = validator.Validate(dst.AutomationPolicyLifecycleVerify403ResponseOneOf2)
			match++
		}
	} else {
		dst.AutomationPolicyLifecycleVerify403ResponseOneOf2 = nil
	}

	// try to unmarshal data into AutomationPolicyLifecycleVerify403ResponseOneOf3
	err = utils.NewStrictDecoder(data).Decode(&dst.AutomationPolicyLifecycleVerify403ResponseOneOf3)
	if err == nil {
		jsonAutomationPolicyLifecycleVerify403ResponseOneOf3, _ := json.Marshal(dst.AutomationPolicyLifecycleVerify403ResponseOneOf3)
		if string(jsonAutomationPolicyLifecycleVerify403ResponseOneOf3) == "{}" { // empty struct
			dst.AutomationPolicyLifecycleVerify403ResponseOneOf3 = nil
		} else {
			_ = validator.Validate(dst.AutomationPolicyLifecycleVerify403ResponseOneOf3)
			match++
		}
	} else {
		dst.AutomationPolicyLifecycleVerify403ResponseOneOf3 = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CertificateList403Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CertificateList403Response) MarshalJSON() ([]byte, error) {
	if src.AdocGet403ResponseOneOf1 != nil {
		return json.Marshal(&src.AdocGet403ResponseOneOf1)
	}

	if src.AdocGet403ResponseOneOf2 != nil {
		return json.Marshal(&src.AdocGet403ResponseOneOf2)
	}

	if src.AutomationPolicyLifecycleVerify403ResponseOneOf != nil {
		return json.Marshal(&src.AutomationPolicyLifecycleVerify403ResponseOneOf)
	}

	if src.AutomationPolicyLifecycleVerify403ResponseOneOf1 != nil {
		return json.Marshal(&src.AutomationPolicyLifecycleVerify403ResponseOneOf1)
	}

	if src.AutomationPolicyLifecycleVerify403ResponseOneOf2 != nil {
		return json.Marshal(&src.AutomationPolicyLifecycleVerify403ResponseOneOf2)
	}

	if src.AutomationPolicyLifecycleVerify403ResponseOneOf3 != nil {
		return json.Marshal(&src.AutomationPolicyLifecycleVerify403ResponseOneOf3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CertificateList403Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdocGet403ResponseOneOf1 != nil {
		return obj.AdocGet403ResponseOneOf1
	}

	if obj.AdocGet403ResponseOneOf2 != nil {
		return obj.AdocGet403ResponseOneOf2
	}

	if obj.AutomationPolicyLifecycleVerify403ResponseOneOf != nil {
		return obj.AutomationPolicyLifecycleVerify403ResponseOneOf
	}

	if obj.AutomationPolicyLifecycleVerify403ResponseOneOf1 != nil {
		return obj.AutomationPolicyLifecycleVerify403ResponseOneOf1
	}

	if obj.AutomationPolicyLifecycleVerify403ResponseOneOf2 != nil {
		return obj.AutomationPolicyLifecycleVerify403ResponseOneOf2
	}

	if obj.AutomationPolicyLifecycleVerify403ResponseOneOf3 != nil {
		return obj.AutomationPolicyLifecycleVerify403ResponseOneOf3
	}

	// all schemas are nil
	return nil
}

type NullableCertificateList403Response struct {
	value *CertificateList403Response
	isSet bool
}

func (v NullableCertificateList403Response) Get() *CertificateList403Response {
	return v.value
}

func (v *NullableCertificateList403Response) Set(val *CertificateList403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateList403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateList403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateList403Response(val *CertificateList403Response) *NullableCertificateList403Response {
	return &NullableCertificateList403Response{value: val, isSet: true}
}

func (v NullableCertificateList403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateList403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
