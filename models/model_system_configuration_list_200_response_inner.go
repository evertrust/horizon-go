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

// SystemConfigurationList200ResponseInner - struct for SystemConfigurationList200ResponseInner
type SystemConfigurationList200ResponseInner struct {
	InterfaceCustomizationConfigurationResponse *InterfaceCustomizationConfigurationResponse
	InternalMonitorConfigurationResponse        *InternalMonitorConfigurationResponse
	LicenseConfigurationResponse                *LicenseConfigurationResponse
	StorageConfigurationResponse                *StorageConfigurationResponse
}

// InterfaceCustomizationConfigurationResponseAsSystemConfigurationList200ResponseInner is a convenience function that returns InterfaceCustomizationConfigurationResponse wrapped in SystemConfigurationList200ResponseInner
func InterfaceCustomizationConfigurationResponseAsSystemConfigurationList200ResponseInner(v *InterfaceCustomizationConfigurationResponse) SystemConfigurationList200ResponseInner {
	return SystemConfigurationList200ResponseInner{
		InterfaceCustomizationConfigurationResponse: v,
	}
}

// InternalMonitorConfigurationResponseAsSystemConfigurationList200ResponseInner is a convenience function that returns InternalMonitorConfigurationResponse wrapped in SystemConfigurationList200ResponseInner
func InternalMonitorConfigurationResponseAsSystemConfigurationList200ResponseInner(v *InternalMonitorConfigurationResponse) SystemConfigurationList200ResponseInner {
	return SystemConfigurationList200ResponseInner{
		InternalMonitorConfigurationResponse: v,
	}
}

// LicenseConfigurationResponseAsSystemConfigurationList200ResponseInner is a convenience function that returns LicenseConfigurationResponse wrapped in SystemConfigurationList200ResponseInner
func LicenseConfigurationResponseAsSystemConfigurationList200ResponseInner(v *LicenseConfigurationResponse) SystemConfigurationList200ResponseInner {
	return SystemConfigurationList200ResponseInner{
		LicenseConfigurationResponse: v,
	}
}

// StorageConfigurationResponseAsSystemConfigurationList200ResponseInner is a convenience function that returns StorageConfigurationResponse wrapped in SystemConfigurationList200ResponseInner
func StorageConfigurationResponseAsSystemConfigurationList200ResponseInner(v *StorageConfigurationResponse) SystemConfigurationList200ResponseInner {
	return SystemConfigurationList200ResponseInner{
		StorageConfigurationResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SystemConfigurationList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InterfaceCustomizationConfigurationResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.InterfaceCustomizationConfigurationResponse)
	if err == nil {
		jsonInterfaceCustomizationConfigurationResponse, _ := json.Marshal(dst.InterfaceCustomizationConfigurationResponse)
		if string(jsonInterfaceCustomizationConfigurationResponse) == "{}" { // empty struct
			dst.InterfaceCustomizationConfigurationResponse = nil
		} else {
			_ = validator.Validate(dst.InterfaceCustomizationConfigurationResponse)
			match++
		}
	} else {
		dst.InterfaceCustomizationConfigurationResponse = nil
	}

	// try to unmarshal data into InternalMonitorConfigurationResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.InternalMonitorConfigurationResponse)
	if err == nil {
		jsonInternalMonitorConfigurationResponse, _ := json.Marshal(dst.InternalMonitorConfigurationResponse)
		if string(jsonInternalMonitorConfigurationResponse) == "{}" { // empty struct
			dst.InternalMonitorConfigurationResponse = nil
		} else {
			_ = validator.Validate(dst.InternalMonitorConfigurationResponse)
			match++
		}
	} else {
		dst.InternalMonitorConfigurationResponse = nil
	}

	// try to unmarshal data into LicenseConfigurationResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.LicenseConfigurationResponse)
	if err == nil {
		jsonLicenseConfigurationResponse, _ := json.Marshal(dst.LicenseConfigurationResponse)
		if string(jsonLicenseConfigurationResponse) == "{}" { // empty struct
			dst.LicenseConfigurationResponse = nil
		} else {
			_ = validator.Validate(dst.LicenseConfigurationResponse)
			match++
		}
	} else {
		dst.LicenseConfigurationResponse = nil
	}

	// try to unmarshal data into StorageConfigurationResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.StorageConfigurationResponse)
	if err == nil {
		jsonStorageConfigurationResponse, _ := json.Marshal(dst.StorageConfigurationResponse)
		if string(jsonStorageConfigurationResponse) == "{}" { // empty struct
			dst.StorageConfigurationResponse = nil
		} else {
			_ = validator.Validate(dst.StorageConfigurationResponse)
			match++
		}
	} else {
		dst.StorageConfigurationResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SystemConfigurationList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SystemConfigurationList200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.InterfaceCustomizationConfigurationResponse != nil {
		return json.Marshal(&src.InterfaceCustomizationConfigurationResponse)
	}

	if src.InternalMonitorConfigurationResponse != nil {
		return json.Marshal(&src.InternalMonitorConfigurationResponse)
	}

	if src.LicenseConfigurationResponse != nil {
		return json.Marshal(&src.LicenseConfigurationResponse)
	}

	if src.StorageConfigurationResponse != nil {
		return json.Marshal(&src.StorageConfigurationResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SystemConfigurationList200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.InterfaceCustomizationConfigurationResponse != nil {
		return obj.InterfaceCustomizationConfigurationResponse
	}

	if obj.InternalMonitorConfigurationResponse != nil {
		return obj.InternalMonitorConfigurationResponse
	}

	if obj.LicenseConfigurationResponse != nil {
		return obj.LicenseConfigurationResponse
	}

	if obj.StorageConfigurationResponse != nil {
		return obj.StorageConfigurationResponse
	}

	// all schemas are nil
	return nil
}

type NullableSystemConfigurationList200ResponseInner struct {
	value *SystemConfigurationList200ResponseInner
	isSet bool
}

func (v NullableSystemConfigurationList200ResponseInner) Get() *SystemConfigurationList200ResponseInner {
	return v.value
}

func (v *NullableSystemConfigurationList200ResponseInner) Set(val *SystemConfigurationList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigurationList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigurationList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigurationList200ResponseInner(val *SystemConfigurationList200ResponseInner) *NullableSystemConfigurationList200ResponseInner {
	return &NullableSystemConfigurationList200ResponseInner{value: val, isSet: true}
}

func (v NullableSystemConfigurationList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigurationList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
