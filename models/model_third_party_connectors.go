/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

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

// ThirdPartyConnectors - struct for ThirdPartyConnectors
type ThirdPartyConnectors struct {
	AWSConnector           *AWSConnector
	AzureKeyVaultConnector *AzureKeyVaultConnector
	F5AS3Connector         *F5AS3Connector
	F5ClientConnector      *F5ClientConnector
	GCMConnector           *GCMConnector
	IntuneConnector        *IntuneConnector
	IntunePKCSConnector    *IntunePKCSConnector
	JamfConnector          *JamfConnector
	LDAPConnector          *LDAPConnector
	MSADConnector          *MSADConnector
	NetscalerConnector     *NetscalerConnector
}

// AWSConnectorAsThirdPartyConnectors is a convenience function that returns AWSConnector wrapped in ThirdPartyConnectors
func AWSConnectorAsThirdPartyConnectors(v *AWSConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		AWSConnector: v,
	}
}

// AzureKeyVaultConnectorAsThirdPartyConnectors is a convenience function that returns AzureKeyVaultConnector wrapped in ThirdPartyConnectors
func AzureKeyVaultConnectorAsThirdPartyConnectors(v *AzureKeyVaultConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		AzureKeyVaultConnector: v,
	}
}

// F5AS3ConnectorAsThirdPartyConnectors is a convenience function that returns F5AS3Connector wrapped in ThirdPartyConnectors
func F5AS3ConnectorAsThirdPartyConnectors(v *F5AS3Connector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		F5AS3Connector: v,
	}
}

// F5ClientConnectorAsThirdPartyConnectors is a convenience function that returns F5ClientConnector wrapped in ThirdPartyConnectors
func F5ClientConnectorAsThirdPartyConnectors(v *F5ClientConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		F5ClientConnector: v,
	}
}

// GCMConnectorAsThirdPartyConnectors is a convenience function that returns GCMConnector wrapped in ThirdPartyConnectors
func GCMConnectorAsThirdPartyConnectors(v *GCMConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		GCMConnector: v,
	}
}

// IntuneConnectorAsThirdPartyConnectors is a convenience function that returns IntuneConnector wrapped in ThirdPartyConnectors
func IntuneConnectorAsThirdPartyConnectors(v *IntuneConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		IntuneConnector: v,
	}
}

// IntunePKCSConnectorAsThirdPartyConnectors is a convenience function that returns IntunePKCSConnector wrapped in ThirdPartyConnectors
func IntunePKCSConnectorAsThirdPartyConnectors(v *IntunePKCSConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		IntunePKCSConnector: v,
	}
}

// JamfConnectorAsThirdPartyConnectors is a convenience function that returns JamfConnector wrapped in ThirdPartyConnectors
func JamfConnectorAsThirdPartyConnectors(v *JamfConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		JamfConnector: v,
	}
}

// LDAPConnectorAsThirdPartyConnectors is a convenience function that returns LDAPConnector wrapped in ThirdPartyConnectors
func LDAPConnectorAsThirdPartyConnectors(v *LDAPConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		LDAPConnector: v,
	}
}

// MSADConnectorAsThirdPartyConnectors is a convenience function that returns MSADConnector wrapped in ThirdPartyConnectors
func MSADConnectorAsThirdPartyConnectors(v *MSADConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		MSADConnector: v,
	}
}

// NetscalerConnectorAsThirdPartyConnectors is a convenience function that returns NetscalerConnector wrapped in ThirdPartyConnectors
func NetscalerConnectorAsThirdPartyConnectors(v *NetscalerConnector) ThirdPartyConnectors {
	return ThirdPartyConnectors{
		NetscalerConnector: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThirdPartyConnectors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.AWSConnector)
	if err == nil {
		jsonAWSConnector, _ := json.Marshal(dst.AWSConnector)
		if string(jsonAWSConnector) == "{}" { // empty struct
			dst.AWSConnector = nil
		} else {
			_ = validator.Validate(dst.AWSConnector)
			match++
		}
	} else {
		dst.AWSConnector = nil
	}

	// try to unmarshal data into AzureKeyVaultConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.AzureKeyVaultConnector)
	if err == nil {
		jsonAzureKeyVaultConnector, _ := json.Marshal(dst.AzureKeyVaultConnector)
		if string(jsonAzureKeyVaultConnector) == "{}" { // empty struct
			dst.AzureKeyVaultConnector = nil
		} else {
			_ = validator.Validate(dst.AzureKeyVaultConnector)
			match++
		}
	} else {
		dst.AzureKeyVaultConnector = nil
	}

	// try to unmarshal data into F5AS3Connector
	err = utils.NewStrictDecoder(data).Decode(&dst.F5AS3Connector)
	if err == nil {
		jsonF5AS3Connector, _ := json.Marshal(dst.F5AS3Connector)
		if string(jsonF5AS3Connector) == "{}" { // empty struct
			dst.F5AS3Connector = nil
		} else {
			_ = validator.Validate(dst.F5AS3Connector)
			match++
		}
	} else {
		dst.F5AS3Connector = nil
	}

	// try to unmarshal data into F5ClientConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.F5ClientConnector)
	if err == nil {
		jsonF5ClientConnector, _ := json.Marshal(dst.F5ClientConnector)
		if string(jsonF5ClientConnector) == "{}" { // empty struct
			dst.F5ClientConnector = nil
		} else {
			_ = validator.Validate(dst.F5ClientConnector)
			match++
		}
	} else {
		dst.F5ClientConnector = nil
	}

	// try to unmarshal data into GCMConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.GCMConnector)
	if err == nil {
		jsonGCMConnector, _ := json.Marshal(dst.GCMConnector)
		if string(jsonGCMConnector) == "{}" { // empty struct
			dst.GCMConnector = nil
		} else {
			_ = validator.Validate(dst.GCMConnector)
			match++
		}
	} else {
		dst.GCMConnector = nil
	}

	// try to unmarshal data into IntuneConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.IntuneConnector)
	if err == nil {
		jsonIntuneConnector, _ := json.Marshal(dst.IntuneConnector)
		if string(jsonIntuneConnector) == "{}" { // empty struct
			dst.IntuneConnector = nil
		} else {
			_ = validator.Validate(dst.IntuneConnector)
			match++
		}
	} else {
		dst.IntuneConnector = nil
	}

	// try to unmarshal data into IntunePKCSConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.IntunePKCSConnector)
	if err == nil {
		jsonIntunePKCSConnector, _ := json.Marshal(dst.IntunePKCSConnector)
		if string(jsonIntunePKCSConnector) == "{}" { // empty struct
			dst.IntunePKCSConnector = nil
		} else {
			_ = validator.Validate(dst.IntunePKCSConnector)
			match++
		}
	} else {
		dst.IntunePKCSConnector = nil
	}

	// try to unmarshal data into JamfConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.JamfConnector)
	if err == nil {
		jsonJamfConnector, _ := json.Marshal(dst.JamfConnector)
		if string(jsonJamfConnector) == "{}" { // empty struct
			dst.JamfConnector = nil
		} else {
			_ = validator.Validate(dst.JamfConnector)
			match++
		}
	} else {
		dst.JamfConnector = nil
	}

	// try to unmarshal data into LDAPConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.LDAPConnector)
	if err == nil {
		jsonLDAPConnector, _ := json.Marshal(dst.LDAPConnector)
		if string(jsonLDAPConnector) == "{}" { // empty struct
			dst.LDAPConnector = nil
		} else {
			_ = validator.Validate(dst.LDAPConnector)
			match++
		}
	} else {
		dst.LDAPConnector = nil
	}

	// try to unmarshal data into MSADConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.MSADConnector)
	if err == nil {
		jsonMSADConnector, _ := json.Marshal(dst.MSADConnector)
		if string(jsonMSADConnector) == "{}" { // empty struct
			dst.MSADConnector = nil
		} else {
			_ = validator.Validate(dst.MSADConnector)
			match++
		}
	} else {
		dst.MSADConnector = nil
	}

	// try to unmarshal data into NetscalerConnector
	err = utils.NewStrictDecoder(data).Decode(&dst.NetscalerConnector)
	if err == nil {
		jsonNetscalerConnector, _ := json.Marshal(dst.NetscalerConnector)
		if string(jsonNetscalerConnector) == "{}" { // empty struct
			dst.NetscalerConnector = nil
		} else {
			_ = validator.Validate(dst.NetscalerConnector)
			match++
		}
	} else {
		dst.NetscalerConnector = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThirdPartyConnectors)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThirdPartyConnectors) MarshalJSON() ([]byte, error) {
	if src.AWSConnector != nil {
		return json.Marshal(&src.AWSConnector)
	}

	if src.AzureKeyVaultConnector != nil {
		return json.Marshal(&src.AzureKeyVaultConnector)
	}

	if src.F5AS3Connector != nil {
		return json.Marshal(&src.F5AS3Connector)
	}

	if src.F5ClientConnector != nil {
		return json.Marshal(&src.F5ClientConnector)
	}

	if src.GCMConnector != nil {
		return json.Marshal(&src.GCMConnector)
	}

	if src.IntuneConnector != nil {
		return json.Marshal(&src.IntuneConnector)
	}

	if src.IntunePKCSConnector != nil {
		return json.Marshal(&src.IntunePKCSConnector)
	}

	if src.JamfConnector != nil {
		return json.Marshal(&src.JamfConnector)
	}

	if src.LDAPConnector != nil {
		return json.Marshal(&src.LDAPConnector)
	}

	if src.MSADConnector != nil {
		return json.Marshal(&src.MSADConnector)
	}

	if src.NetscalerConnector != nil {
		return json.Marshal(&src.NetscalerConnector)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThirdPartyConnectors) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSConnector != nil {
		return obj.AWSConnector
	}

	if obj.AzureKeyVaultConnector != nil {
		return obj.AzureKeyVaultConnector
	}

	if obj.F5AS3Connector != nil {
		return obj.F5AS3Connector
	}

	if obj.F5ClientConnector != nil {
		return obj.F5ClientConnector
	}

	if obj.GCMConnector != nil {
		return obj.GCMConnector
	}

	if obj.IntuneConnector != nil {
		return obj.IntuneConnector
	}

	if obj.IntunePKCSConnector != nil {
		return obj.IntunePKCSConnector
	}

	if obj.JamfConnector != nil {
		return obj.JamfConnector
	}

	if obj.LDAPConnector != nil {
		return obj.LDAPConnector
	}

	if obj.MSADConnector != nil {
		return obj.MSADConnector
	}

	if obj.NetscalerConnector != nil {
		return obj.NetscalerConnector
	}

	// all schemas are nil
	return nil
}

type NullableThirdPartyConnectors struct {
	value *ThirdPartyConnectors
	isSet bool
}

func (v NullableThirdPartyConnectors) Get() *ThirdPartyConnectors {
	return v.value
}

func (v *NullableThirdPartyConnectors) Set(val *ThirdPartyConnectors) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyConnectors) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyConnectors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyConnectors(val *ThirdPartyConnectors) *NullableThirdPartyConnectors {
	return &NullableThirdPartyConnectors{value: val, isSet: true}
}

func (v NullableThirdPartyConnectors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyConnectors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
