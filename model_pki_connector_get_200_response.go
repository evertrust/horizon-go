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

// PkiConnectorGet200Response - struct for PkiConnectorGet200Response
type PkiConnectorGet200Response struct {
	ADCSConnectorResponse *ADCSConnectorResponse
	AWSACMPCAConnectorResponse *AWSACMPCAConnectorResponse
	CMPConnectorResponse *CMPConnectorResponse
	CertEuropeConnectorResponse *CertEuropeConnectorResponse
	DigiCertConnectorResponse *DigiCertConnectorResponse
	EJBCAConnectorResponse *EJBCAConnectorResponse
	EntrustConnectorResponse *EntrustConnectorResponse
	EverTrustADCSConnectorResponse *EverTrustADCSConnectorResponse
	FCMSConnectorResponse *FCMSConnectorResponse
	GSAtlasConnectorResponse *GSAtlasConnectorResponse
	GSMSSLConnectorResponse *GSMSSLConnectorResponse
	IntegratedCAConnectorResponse *IntegratedCAConnectorResponse
	NameshieldResponse *NameshieldResponse
	NexusCMConnectorResponse *NexusCMConnectorResponse
	OTPKIConnectorResponse *OTPKIConnectorResponse
	StreamConnectorResponse *StreamConnectorResponse
}

// ADCSConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns ADCSConnectorResponse wrapped in PkiConnectorGet200Response
func ADCSConnectorResponseAsPkiConnectorGet200Response(v *ADCSConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		ADCSConnectorResponse: v,
	}
}

// AWSACMPCAConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns AWSACMPCAConnectorResponse wrapped in PkiConnectorGet200Response
func AWSACMPCAConnectorResponseAsPkiConnectorGet200Response(v *AWSACMPCAConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		AWSACMPCAConnectorResponse: v,
	}
}

// CMPConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns CMPConnectorResponse wrapped in PkiConnectorGet200Response
func CMPConnectorResponseAsPkiConnectorGet200Response(v *CMPConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		CMPConnectorResponse: v,
	}
}

// CertEuropeConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns CertEuropeConnectorResponse wrapped in PkiConnectorGet200Response
func CertEuropeConnectorResponseAsPkiConnectorGet200Response(v *CertEuropeConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		CertEuropeConnectorResponse: v,
	}
}

// DigiCertConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns DigiCertConnectorResponse wrapped in PkiConnectorGet200Response
func DigiCertConnectorResponseAsPkiConnectorGet200Response(v *DigiCertConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		DigiCertConnectorResponse: v,
	}
}

// EJBCAConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns EJBCAConnectorResponse wrapped in PkiConnectorGet200Response
func EJBCAConnectorResponseAsPkiConnectorGet200Response(v *EJBCAConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		EJBCAConnectorResponse: v,
	}
}

// EntrustConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns EntrustConnectorResponse wrapped in PkiConnectorGet200Response
func EntrustConnectorResponseAsPkiConnectorGet200Response(v *EntrustConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		EntrustConnectorResponse: v,
	}
}

// EverTrustADCSConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns EverTrustADCSConnectorResponse wrapped in PkiConnectorGet200Response
func EverTrustADCSConnectorResponseAsPkiConnectorGet200Response(v *EverTrustADCSConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		EverTrustADCSConnectorResponse: v,
	}
}

// FCMSConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns FCMSConnectorResponse wrapped in PkiConnectorGet200Response
func FCMSConnectorResponseAsPkiConnectorGet200Response(v *FCMSConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		FCMSConnectorResponse: v,
	}
}

// GSAtlasConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns GSAtlasConnectorResponse wrapped in PkiConnectorGet200Response
func GSAtlasConnectorResponseAsPkiConnectorGet200Response(v *GSAtlasConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		GSAtlasConnectorResponse: v,
	}
}

// GSMSSLConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns GSMSSLConnectorResponse wrapped in PkiConnectorGet200Response
func GSMSSLConnectorResponseAsPkiConnectorGet200Response(v *GSMSSLConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		GSMSSLConnectorResponse: v,
	}
}

// IntegratedCAConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns IntegratedCAConnectorResponse wrapped in PkiConnectorGet200Response
func IntegratedCAConnectorResponseAsPkiConnectorGet200Response(v *IntegratedCAConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		IntegratedCAConnectorResponse: v,
	}
}

// NameshieldResponseAsPkiConnectorGet200Response is a convenience function that returns NameshieldResponse wrapped in PkiConnectorGet200Response
func NameshieldResponseAsPkiConnectorGet200Response(v *NameshieldResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		NameshieldResponse: v,
	}
}

// NexusCMConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns NexusCMConnectorResponse wrapped in PkiConnectorGet200Response
func NexusCMConnectorResponseAsPkiConnectorGet200Response(v *NexusCMConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		NexusCMConnectorResponse: v,
	}
}

// OTPKIConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns OTPKIConnectorResponse wrapped in PkiConnectorGet200Response
func OTPKIConnectorResponseAsPkiConnectorGet200Response(v *OTPKIConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		OTPKIConnectorResponse: v,
	}
}

// StreamConnectorResponseAsPkiConnectorGet200Response is a convenience function that returns StreamConnectorResponse wrapped in PkiConnectorGet200Response
func StreamConnectorResponseAsPkiConnectorGet200Response(v *StreamConnectorResponse) PkiConnectorGet200Response {
	return PkiConnectorGet200Response{
		StreamConnectorResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PkiConnectorGet200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ADCSConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.ADCSConnectorResponse)
	if err == nil {
		jsonADCSConnectorResponse, _ := json.Marshal(dst.ADCSConnectorResponse)
		if string(jsonADCSConnectorResponse) == "{}" { // empty struct
			dst.ADCSConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.ADCSConnectorResponse)
            match++
		}
	} else {
		dst.ADCSConnectorResponse = nil
	}

	// try to unmarshal data into AWSACMPCAConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.AWSACMPCAConnectorResponse)
	if err == nil {
		jsonAWSACMPCAConnectorResponse, _ := json.Marshal(dst.AWSACMPCAConnectorResponse)
		if string(jsonAWSACMPCAConnectorResponse) == "{}" { // empty struct
			dst.AWSACMPCAConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.AWSACMPCAConnectorResponse)
            match++
		}
	} else {
		dst.AWSACMPCAConnectorResponse = nil
	}

	// try to unmarshal data into CMPConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.CMPConnectorResponse)
	if err == nil {
		jsonCMPConnectorResponse, _ := json.Marshal(dst.CMPConnectorResponse)
		if string(jsonCMPConnectorResponse) == "{}" { // empty struct
			dst.CMPConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.CMPConnectorResponse)
            match++
		}
	} else {
		dst.CMPConnectorResponse = nil
	}

	// try to unmarshal data into CertEuropeConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.CertEuropeConnectorResponse)
	if err == nil {
		jsonCertEuropeConnectorResponse, _ := json.Marshal(dst.CertEuropeConnectorResponse)
		if string(jsonCertEuropeConnectorResponse) == "{}" { // empty struct
			dst.CertEuropeConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.CertEuropeConnectorResponse)
            match++
		}
	} else {
		dst.CertEuropeConnectorResponse = nil
	}

	// try to unmarshal data into DigiCertConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.DigiCertConnectorResponse)
	if err == nil {
		jsonDigiCertConnectorResponse, _ := json.Marshal(dst.DigiCertConnectorResponse)
		if string(jsonDigiCertConnectorResponse) == "{}" { // empty struct
			dst.DigiCertConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.DigiCertConnectorResponse)
            match++
		}
	} else {
		dst.DigiCertConnectorResponse = nil
	}

	// try to unmarshal data into EJBCAConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.EJBCAConnectorResponse)
	if err == nil {
		jsonEJBCAConnectorResponse, _ := json.Marshal(dst.EJBCAConnectorResponse)
		if string(jsonEJBCAConnectorResponse) == "{}" { // empty struct
			dst.EJBCAConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.EJBCAConnectorResponse)
            match++
		}
	} else {
		dst.EJBCAConnectorResponse = nil
	}

	// try to unmarshal data into EntrustConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.EntrustConnectorResponse)
	if err == nil {
		jsonEntrustConnectorResponse, _ := json.Marshal(dst.EntrustConnectorResponse)
		if string(jsonEntrustConnectorResponse) == "{}" { // empty struct
			dst.EntrustConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.EntrustConnectorResponse)
            match++
		}
	} else {
		dst.EntrustConnectorResponse = nil
	}

	// try to unmarshal data into EverTrustADCSConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.EverTrustADCSConnectorResponse)
	if err == nil {
		jsonEverTrustADCSConnectorResponse, _ := json.Marshal(dst.EverTrustADCSConnectorResponse)
		if string(jsonEverTrustADCSConnectorResponse) == "{}" { // empty struct
			dst.EverTrustADCSConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.EverTrustADCSConnectorResponse)
            match++
		}
	} else {
		dst.EverTrustADCSConnectorResponse = nil
	}

	// try to unmarshal data into FCMSConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.FCMSConnectorResponse)
	if err == nil {
		jsonFCMSConnectorResponse, _ := json.Marshal(dst.FCMSConnectorResponse)
		if string(jsonFCMSConnectorResponse) == "{}" { // empty struct
			dst.FCMSConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.FCMSConnectorResponse)
            match++
		}
	} else {
		dst.FCMSConnectorResponse = nil
	}

	// try to unmarshal data into GSAtlasConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.GSAtlasConnectorResponse)
	if err == nil {
		jsonGSAtlasConnectorResponse, _ := json.Marshal(dst.GSAtlasConnectorResponse)
		if string(jsonGSAtlasConnectorResponse) == "{}" { // empty struct
			dst.GSAtlasConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.GSAtlasConnectorResponse)
            match++
		}
	} else {
		dst.GSAtlasConnectorResponse = nil
	}

	// try to unmarshal data into GSMSSLConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.GSMSSLConnectorResponse)
	if err == nil {
		jsonGSMSSLConnectorResponse, _ := json.Marshal(dst.GSMSSLConnectorResponse)
		if string(jsonGSMSSLConnectorResponse) == "{}" { // empty struct
			dst.GSMSSLConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.GSMSSLConnectorResponse)
            match++
		}
	} else {
		dst.GSMSSLConnectorResponse = nil
	}

	// try to unmarshal data into IntegratedCAConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.IntegratedCAConnectorResponse)
	if err == nil {
		jsonIntegratedCAConnectorResponse, _ := json.Marshal(dst.IntegratedCAConnectorResponse)
		if string(jsonIntegratedCAConnectorResponse) == "{}" { // empty struct
			dst.IntegratedCAConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.IntegratedCAConnectorResponse)
            match++
		}
	} else {
		dst.IntegratedCAConnectorResponse = nil
	}

	// try to unmarshal data into NameshieldResponse
	err = newStrictDecoder(data).Decode(&dst.NameshieldResponse)
	if err == nil {
		jsonNameshieldResponse, _ := json.Marshal(dst.NameshieldResponse)
		if string(jsonNameshieldResponse) == "{}" { // empty struct
			dst.NameshieldResponse = nil
		} else {
            _ = validator.Validate(dst.NameshieldResponse)
            match++
		}
	} else {
		dst.NameshieldResponse = nil
	}

	// try to unmarshal data into NexusCMConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.NexusCMConnectorResponse)
	if err == nil {
		jsonNexusCMConnectorResponse, _ := json.Marshal(dst.NexusCMConnectorResponse)
		if string(jsonNexusCMConnectorResponse) == "{}" { // empty struct
			dst.NexusCMConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.NexusCMConnectorResponse)
            match++
		}
	} else {
		dst.NexusCMConnectorResponse = nil
	}

	// try to unmarshal data into OTPKIConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.OTPKIConnectorResponse)
	if err == nil {
		jsonOTPKIConnectorResponse, _ := json.Marshal(dst.OTPKIConnectorResponse)
		if string(jsonOTPKIConnectorResponse) == "{}" { // empty struct
			dst.OTPKIConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.OTPKIConnectorResponse)
            match++
		}
	} else {
		dst.OTPKIConnectorResponse = nil
	}

	// try to unmarshal data into StreamConnectorResponse
	err = newStrictDecoder(data).Decode(&dst.StreamConnectorResponse)
	if err == nil {
		jsonStreamConnectorResponse, _ := json.Marshal(dst.StreamConnectorResponse)
		if string(jsonStreamConnectorResponse) == "{}" { // empty struct
			dst.StreamConnectorResponse = nil
		} else {
            _ = validator.Validate(dst.StreamConnectorResponse)
            match++
		}
	} else {
		dst.StreamConnectorResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PkiConnectorGet200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PkiConnectorGet200Response) MarshalJSON() ([]byte, error) {
	if src.ADCSConnectorResponse != nil {
		return json.Marshal(&src.ADCSConnectorResponse)
	}

	if src.AWSACMPCAConnectorResponse != nil {
		return json.Marshal(&src.AWSACMPCAConnectorResponse)
	}

	if src.CMPConnectorResponse != nil {
		return json.Marshal(&src.CMPConnectorResponse)
	}

	if src.CertEuropeConnectorResponse != nil {
		return json.Marshal(&src.CertEuropeConnectorResponse)
	}

	if src.DigiCertConnectorResponse != nil {
		return json.Marshal(&src.DigiCertConnectorResponse)
	}

	if src.EJBCAConnectorResponse != nil {
		return json.Marshal(&src.EJBCAConnectorResponse)
	}

	if src.EntrustConnectorResponse != nil {
		return json.Marshal(&src.EntrustConnectorResponse)
	}

	if src.EverTrustADCSConnectorResponse != nil {
		return json.Marshal(&src.EverTrustADCSConnectorResponse)
	}

	if src.FCMSConnectorResponse != nil {
		return json.Marshal(&src.FCMSConnectorResponse)
	}

	if src.GSAtlasConnectorResponse != nil {
		return json.Marshal(&src.GSAtlasConnectorResponse)
	}

	if src.GSMSSLConnectorResponse != nil {
		return json.Marshal(&src.GSMSSLConnectorResponse)
	}

	if src.IntegratedCAConnectorResponse != nil {
		return json.Marshal(&src.IntegratedCAConnectorResponse)
	}

	if src.NameshieldResponse != nil {
		return json.Marshal(&src.NameshieldResponse)
	}

	if src.NexusCMConnectorResponse != nil {
		return json.Marshal(&src.NexusCMConnectorResponse)
	}

	if src.OTPKIConnectorResponse != nil {
		return json.Marshal(&src.OTPKIConnectorResponse)
	}

	if src.StreamConnectorResponse != nil {
		return json.Marshal(&src.StreamConnectorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PkiConnectorGet200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ADCSConnectorResponse != nil {
		return obj.ADCSConnectorResponse
	}

	if obj.AWSACMPCAConnectorResponse != nil {
		return obj.AWSACMPCAConnectorResponse
	}

	if obj.CMPConnectorResponse != nil {
		return obj.CMPConnectorResponse
	}

	if obj.CertEuropeConnectorResponse != nil {
		return obj.CertEuropeConnectorResponse
	}

	if obj.DigiCertConnectorResponse != nil {
		return obj.DigiCertConnectorResponse
	}

	if obj.EJBCAConnectorResponse != nil {
		return obj.EJBCAConnectorResponse
	}

	if obj.EntrustConnectorResponse != nil {
		return obj.EntrustConnectorResponse
	}

	if obj.EverTrustADCSConnectorResponse != nil {
		return obj.EverTrustADCSConnectorResponse
	}

	if obj.FCMSConnectorResponse != nil {
		return obj.FCMSConnectorResponse
	}

	if obj.GSAtlasConnectorResponse != nil {
		return obj.GSAtlasConnectorResponse
	}

	if obj.GSMSSLConnectorResponse != nil {
		return obj.GSMSSLConnectorResponse
	}

	if obj.IntegratedCAConnectorResponse != nil {
		return obj.IntegratedCAConnectorResponse
	}

	if obj.NameshieldResponse != nil {
		return obj.NameshieldResponse
	}

	if obj.NexusCMConnectorResponse != nil {
		return obj.NexusCMConnectorResponse
	}

	if obj.OTPKIConnectorResponse != nil {
		return obj.OTPKIConnectorResponse
	}

	if obj.StreamConnectorResponse != nil {
		return obj.StreamConnectorResponse
	}

	// all schemas are nil
	return nil
}

type NullablePkiConnectorGet200Response struct {
	value *PkiConnectorGet200Response
	isSet bool
}

func (v NullablePkiConnectorGet200Response) Get() *PkiConnectorGet200Response {
	return v.value
}

func (v *NullablePkiConnectorGet200Response) Set(val *PkiConnectorGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiConnectorGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiConnectorGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiConnectorGet200Response(val *PkiConnectorGet200Response) *NullablePkiConnectorGet200Response {
	return &NullablePkiConnectorGet200Response{value: val, isSet: true}
}

func (v NullablePkiConnectorGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiConnectorGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


