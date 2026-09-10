/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
	"gopkg.in/validator.v2"
)

// ThirdPartyConnectorResponses - struct for ThirdPartyConnectorResponses
type ThirdPartyConnectorResponses struct {
	AWSConnectorResponse           *AWSConnectorResponse
	AzureKeyVaultConnectorResponse *AzureKeyVaultConnectorResponse
	F5AS3ConnectorResponse         *F5AS3ConnectorResponse
	F5ClientConnectorResponse      *F5ClientConnectorResponse
	GCMConnectorResponse           *GCMConnectorResponse
	IntuneConnectorResponse        *IntuneConnectorResponse
	IntunePKCSConnectorResponse    *IntunePKCSConnectorResponse
	JamfConnectorResponse          *JamfConnectorResponse
	LDAPConnectorResponse          *LDAPConnectorResponse
	MSADConnectorResponse          *MSADConnectorResponse
	NetscalerConnectorResponse     *NetscalerConnectorResponse
}

// AWSConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns AWSConnectorResponse wrapped in ThirdPartyConnectorResponses
func AWSConnectorResponseAsThirdPartyConnectorResponses(v *AWSConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		AWSConnectorResponse: v,
	}
}

// AzureKeyVaultConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns AzureKeyVaultConnectorResponse wrapped in ThirdPartyConnectorResponses
func AzureKeyVaultConnectorResponseAsThirdPartyConnectorResponses(v *AzureKeyVaultConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		AzureKeyVaultConnectorResponse: v,
	}
}

// F5AS3ConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns F5AS3ConnectorResponse wrapped in ThirdPartyConnectorResponses
func F5AS3ConnectorResponseAsThirdPartyConnectorResponses(v *F5AS3ConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		F5AS3ConnectorResponse: v,
	}
}

// F5ClientConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns F5ClientConnectorResponse wrapped in ThirdPartyConnectorResponses
func F5ClientConnectorResponseAsThirdPartyConnectorResponses(v *F5ClientConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		F5ClientConnectorResponse: v,
	}
}

// GCMConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns GCMConnectorResponse wrapped in ThirdPartyConnectorResponses
func GCMConnectorResponseAsThirdPartyConnectorResponses(v *GCMConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		GCMConnectorResponse: v,
	}
}

// IntuneConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns IntuneConnectorResponse wrapped in ThirdPartyConnectorResponses
func IntuneConnectorResponseAsThirdPartyConnectorResponses(v *IntuneConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		IntuneConnectorResponse: v,
	}
}

// IntunePKCSConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns IntunePKCSConnectorResponse wrapped in ThirdPartyConnectorResponses
func IntunePKCSConnectorResponseAsThirdPartyConnectorResponses(v *IntunePKCSConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		IntunePKCSConnectorResponse: v,
	}
}

// JamfConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns JamfConnectorResponse wrapped in ThirdPartyConnectorResponses
func JamfConnectorResponseAsThirdPartyConnectorResponses(v *JamfConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		JamfConnectorResponse: v,
	}
}

// LDAPConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns LDAPConnectorResponse wrapped in ThirdPartyConnectorResponses
func LDAPConnectorResponseAsThirdPartyConnectorResponses(v *LDAPConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		LDAPConnectorResponse: v,
	}
}

// MSADConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns MSADConnectorResponse wrapped in ThirdPartyConnectorResponses
func MSADConnectorResponseAsThirdPartyConnectorResponses(v *MSADConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		MSADConnectorResponse: v,
	}
}

// NetscalerConnectorResponseAsThirdPartyConnectorResponses is a convenience function that returns NetscalerConnectorResponse wrapped in ThirdPartyConnectorResponses
func NetscalerConnectorResponseAsThirdPartyConnectorResponses(v *NetscalerConnectorResponse) ThirdPartyConnectorResponses {
	return ThirdPartyConnectorResponses{
		NetscalerConnectorResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThirdPartyConnectorResponses) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.AWSConnectorResponse)
	if err == nil {
		jsonAWSConnectorResponse, _ := json.Marshal(dst.AWSConnectorResponse)
		if string(jsonAWSConnectorResponse) == "{}" { // empty struct
			dst.AWSConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.AWSConnectorResponse)
			match++
		}
	} else {
		dst.AWSConnectorResponse = nil
	}

	// try to unmarshal data into AzureKeyVaultConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.AzureKeyVaultConnectorResponse)
	if err == nil {
		jsonAzureKeyVaultConnectorResponse, _ := json.Marshal(dst.AzureKeyVaultConnectorResponse)
		if string(jsonAzureKeyVaultConnectorResponse) == "{}" { // empty struct
			dst.AzureKeyVaultConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.AzureKeyVaultConnectorResponse)
			match++
		}
	} else {
		dst.AzureKeyVaultConnectorResponse = nil
	}

	// try to unmarshal data into F5AS3ConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.F5AS3ConnectorResponse)
	if err == nil {
		jsonF5AS3ConnectorResponse, _ := json.Marshal(dst.F5AS3ConnectorResponse)
		if string(jsonF5AS3ConnectorResponse) == "{}" { // empty struct
			dst.F5AS3ConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.F5AS3ConnectorResponse)
			match++
		}
	} else {
		dst.F5AS3ConnectorResponse = nil
	}

	// try to unmarshal data into F5ClientConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.F5ClientConnectorResponse)
	if err == nil {
		jsonF5ClientConnectorResponse, _ := json.Marshal(dst.F5ClientConnectorResponse)
		if string(jsonF5ClientConnectorResponse) == "{}" { // empty struct
			dst.F5ClientConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.F5ClientConnectorResponse)
			match++
		}
	} else {
		dst.F5ClientConnectorResponse = nil
	}

	// try to unmarshal data into GCMConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.GCMConnectorResponse)
	if err == nil {
		jsonGCMConnectorResponse, _ := json.Marshal(dst.GCMConnectorResponse)
		if string(jsonGCMConnectorResponse) == "{}" { // empty struct
			dst.GCMConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.GCMConnectorResponse)
			match++
		}
	} else {
		dst.GCMConnectorResponse = nil
	}

	// try to unmarshal data into IntuneConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.IntuneConnectorResponse)
	if err == nil {
		jsonIntuneConnectorResponse, _ := json.Marshal(dst.IntuneConnectorResponse)
		if string(jsonIntuneConnectorResponse) == "{}" { // empty struct
			dst.IntuneConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.IntuneConnectorResponse)
			match++
		}
	} else {
		dst.IntuneConnectorResponse = nil
	}

	// try to unmarshal data into IntunePKCSConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.IntunePKCSConnectorResponse)
	if err == nil {
		jsonIntunePKCSConnectorResponse, _ := json.Marshal(dst.IntunePKCSConnectorResponse)
		if string(jsonIntunePKCSConnectorResponse) == "{}" { // empty struct
			dst.IntunePKCSConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.IntunePKCSConnectorResponse)
			match++
		}
	} else {
		dst.IntunePKCSConnectorResponse = nil
	}

	// try to unmarshal data into JamfConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.JamfConnectorResponse)
	if err == nil {
		jsonJamfConnectorResponse, _ := json.Marshal(dst.JamfConnectorResponse)
		if string(jsonJamfConnectorResponse) == "{}" { // empty struct
			dst.JamfConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.JamfConnectorResponse)
			match++
		}
	} else {
		dst.JamfConnectorResponse = nil
	}

	// try to unmarshal data into LDAPConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.LDAPConnectorResponse)
	if err == nil {
		jsonLDAPConnectorResponse, _ := json.Marshal(dst.LDAPConnectorResponse)
		if string(jsonLDAPConnectorResponse) == "{}" { // empty struct
			dst.LDAPConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.LDAPConnectorResponse)
			match++
		}
	} else {
		dst.LDAPConnectorResponse = nil
	}

	// try to unmarshal data into MSADConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.MSADConnectorResponse)
	if err == nil {
		jsonMSADConnectorResponse, _ := json.Marshal(dst.MSADConnectorResponse)
		if string(jsonMSADConnectorResponse) == "{}" { // empty struct
			dst.MSADConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.MSADConnectorResponse)
			match++
		}
	} else {
		dst.MSADConnectorResponse = nil
	}

	// try to unmarshal data into NetscalerConnectorResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.NetscalerConnectorResponse)
	if err == nil {
		jsonNetscalerConnectorResponse, _ := json.Marshal(dst.NetscalerConnectorResponse)
		if string(jsonNetscalerConnectorResponse) == "{}" { // empty struct
			dst.NetscalerConnectorResponse = nil
		} else {
			_ = validator.Validate(dst.NetscalerConnectorResponse)
			match++
		}
	} else {
		dst.NetscalerConnectorResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThirdPartyConnectorResponses)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThirdPartyConnectorResponses) MarshalJSON() ([]byte, error) {
	if src.AWSConnectorResponse != nil {
		return json.Marshal(&src.AWSConnectorResponse)
	}

	if src.AzureKeyVaultConnectorResponse != nil {
		return json.Marshal(&src.AzureKeyVaultConnectorResponse)
	}

	if src.F5AS3ConnectorResponse != nil {
		return json.Marshal(&src.F5AS3ConnectorResponse)
	}

	if src.F5ClientConnectorResponse != nil {
		return json.Marshal(&src.F5ClientConnectorResponse)
	}

	if src.GCMConnectorResponse != nil {
		return json.Marshal(&src.GCMConnectorResponse)
	}

	if src.IntuneConnectorResponse != nil {
		return json.Marshal(&src.IntuneConnectorResponse)
	}

	if src.IntunePKCSConnectorResponse != nil {
		return json.Marshal(&src.IntunePKCSConnectorResponse)
	}

	if src.JamfConnectorResponse != nil {
		return json.Marshal(&src.JamfConnectorResponse)
	}

	if src.LDAPConnectorResponse != nil {
		return json.Marshal(&src.LDAPConnectorResponse)
	}

	if src.MSADConnectorResponse != nil {
		return json.Marshal(&src.MSADConnectorResponse)
	}

	if src.NetscalerConnectorResponse != nil {
		return json.Marshal(&src.NetscalerConnectorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThirdPartyConnectorResponses) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSConnectorResponse != nil {
		return obj.AWSConnectorResponse
	}

	if obj.AzureKeyVaultConnectorResponse != nil {
		return obj.AzureKeyVaultConnectorResponse
	}

	if obj.F5AS3ConnectorResponse != nil {
		return obj.F5AS3ConnectorResponse
	}

	if obj.F5ClientConnectorResponse != nil {
		return obj.F5ClientConnectorResponse
	}

	if obj.GCMConnectorResponse != nil {
		return obj.GCMConnectorResponse
	}

	if obj.IntuneConnectorResponse != nil {
		return obj.IntuneConnectorResponse
	}

	if obj.IntunePKCSConnectorResponse != nil {
		return obj.IntunePKCSConnectorResponse
	}

	if obj.JamfConnectorResponse != nil {
		return obj.JamfConnectorResponse
	}

	if obj.LDAPConnectorResponse != nil {
		return obj.LDAPConnectorResponse
	}

	if obj.MSADConnectorResponse != nil {
		return obj.MSADConnectorResponse
	}

	if obj.NetscalerConnectorResponse != nil {
		return obj.NetscalerConnectorResponse
	}

	// all schemas are nil
	return nil
}

type NullableThirdPartyConnectorResponses struct {
	value *ThirdPartyConnectorResponses
	isSet bool
}

func (v NullableThirdPartyConnectorResponses) Get() *ThirdPartyConnectorResponses {
	return v.value
}

func (v *NullableThirdPartyConnectorResponses) Set(val *ThirdPartyConnectorResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyConnectorResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyConnectorResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyConnectorResponses(val *ThirdPartyConnectorResponses) *NullableThirdPartyConnectorResponses {
	return &NullableThirdPartyConnectorResponses{value: val, isSet: true}
}

func (v NullableThirdPartyConnectorResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyConnectorResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
