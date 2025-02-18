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

// ThirdpartyConnectorList200ResponseInner - struct for ThirdpartyConnectorList200ResponseInner
type ThirdpartyConnectorList200ResponseInner struct {
	AWSConnectorResponse *AWSConnectorResponse
	AzureKeyVaultConnectorResponse *AzureKeyVaultConnectorResponse
	F5AS3ConnectorResponse *F5AS3ConnectorResponse
	F5ClientConnectorResponse *F5ClientConnectorResponse
	GCMConnectorResponse *GCMConnectorResponse
	IntuneConnectorResponse *IntuneConnectorResponse
	IntunePKCSConnectorResponse *IntunePKCSConnectorResponse
	JamfConnectorResponse *JamfConnectorResponse
	LDAPConnectorResponse *LDAPConnectorResponse
	MSADConnectorResponse *MSADConnectorResponse
}

// AWSConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns AWSConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func AWSConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *AWSConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		AWSConnectorResponse: v,
	}
}

// AzureKeyVaultConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns AzureKeyVaultConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func AzureKeyVaultConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *AzureKeyVaultConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		AzureKeyVaultConnectorResponse: v,
	}
}

// F5AS3ConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns F5AS3ConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func F5AS3ConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *F5AS3ConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		F5AS3ConnectorResponse: v,
	}
}

// F5ClientConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns F5ClientConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func F5ClientConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *F5ClientConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		F5ClientConnectorResponse: v,
	}
}

// GCMConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns GCMConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func GCMConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *GCMConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		GCMConnectorResponse: v,
	}
}

// IntuneConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns IntuneConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func IntuneConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *IntuneConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		IntuneConnectorResponse: v,
	}
}

// IntunePKCSConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns IntunePKCSConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func IntunePKCSConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *IntunePKCSConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		IntunePKCSConnectorResponse: v,
	}
}

// JamfConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns JamfConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func JamfConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *JamfConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		JamfConnectorResponse: v,
	}
}

// LDAPConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns LDAPConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func LDAPConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *LDAPConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		LDAPConnectorResponse: v,
	}
}

// MSADConnectorResponseAsThirdpartyConnectorList200ResponseInner is a convenience function that returns MSADConnectorResponse wrapped in ThirdpartyConnectorList200ResponseInner
func MSADConnectorResponseAsThirdpartyConnectorList200ResponseInner(v *MSADConnectorResponse) ThirdpartyConnectorList200ResponseInner {
	return ThirdpartyConnectorList200ResponseInner{
		MSADConnectorResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThirdpartyConnectorList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.AWSConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.AzureKeyVaultConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.F5AS3ConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.F5ClientConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.GCMConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.IntuneConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.IntunePKCSConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.JamfConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.LDAPConnectorResponse)
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
	err = newStrictDecoder(data).Decode(&dst.MSADConnectorResponse)
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

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ThirdpartyConnectorList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThirdpartyConnectorList200ResponseInner) MarshalJSON() ([]byte, error) {
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

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThirdpartyConnectorList200ResponseInner) GetActualInstance() (interface{}) {
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

	// all schemas are nil
	return nil
}

type NullableThirdpartyConnectorList200ResponseInner struct {
	value *ThirdpartyConnectorList200ResponseInner
	isSet bool
}

func (v NullableThirdpartyConnectorList200ResponseInner) Get() *ThirdpartyConnectorList200ResponseInner {
	return v.value
}

func (v *NullableThirdpartyConnectorList200ResponseInner) Set(val *ThirdpartyConnectorList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdpartyConnectorList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdpartyConnectorList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdpartyConnectorList200ResponseInner(val *ThirdpartyConnectorList200ResponseInner) *NullableThirdpartyConnectorList200ResponseInner {
	return &NullableThirdpartyConnectorList200ResponseInner{value: val, isSet: true}
}

func (v NullableThirdpartyConnectorList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdpartyConnectorList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


