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

// PKIResponses - struct for PKIResponses
type PKIResponses struct {
	ADCSConnectorResponse           *ADCSConnectorResponse
	AWSACMPCAConnectorResponse      *AWSACMPCAConnectorResponse
	AcmeEnrollConnectorResponse     *AcmeEnrollConnectorResponse
	AcmeRevocationConnectorResponse *AcmeRevocationConnectorResponse
	CMPConnectorResponse            *CMPConnectorResponse
	CertEuropeConnectorResponse     *CertEuropeConnectorResponse
	DigiCertConnectorResponse       *DigiCertConnectorResponse
	EJBCAConnectorResponse          *EJBCAConnectorResponse
	EntrustConnectorResponse        *EntrustConnectorResponse
	EverTrustADCSConnectorResponse  *EverTrustADCSConnectorResponse
	FCMSConnectorResponse           *FCMSConnectorResponse
	GSAtlasConnectorResponse        *GSAtlasConnectorResponse
	GSMSSLConnectorResponse         *GSMSSLConnectorResponse
	IDCAConnectorResponse           *IDCAConnectorResponse
	IntegratedCAConnectorResponse   *IntegratedCAConnectorResponse
	MetaPKIConnectorResponse        *MetaPKIConnectorResponse
	NameshieldResponse              *NameshieldResponse
	NexusCMConnectorResponse        *NexusCMConnectorResponse
	OTPKIConnectorResponse          *OTPKIConnectorResponse
	SectigoCMSConnectorResponse     *SectigoCMSConnectorResponse
	StreamConnectorResponse         *StreamConnectorResponse
	SwissSignConnectorResponse      *SwissSignConnectorResponse
}

// ADCSConnectorResponseAsPKIResponses is a convenience function that returns ADCSConnectorResponse wrapped in PKIResponses
func ADCSConnectorResponseAsPKIResponses(v *ADCSConnectorResponse) PKIResponses {
	return PKIResponses{
		ADCSConnectorResponse: v,
	}
}

// AWSACMPCAConnectorResponseAsPKIResponses is a convenience function that returns AWSACMPCAConnectorResponse wrapped in PKIResponses
func AWSACMPCAConnectorResponseAsPKIResponses(v *AWSACMPCAConnectorResponse) PKIResponses {
	return PKIResponses{
		AWSACMPCAConnectorResponse: v,
	}
}

// AcmeEnrollConnectorResponseAsPKIResponses is a convenience function that returns AcmeEnrollConnectorResponse wrapped in PKIResponses
func AcmeEnrollConnectorResponseAsPKIResponses(v *AcmeEnrollConnectorResponse) PKIResponses {
	return PKIResponses{
		AcmeEnrollConnectorResponse: v,
	}
}

// AcmeRevocationConnectorResponseAsPKIResponses is a convenience function that returns AcmeRevocationConnectorResponse wrapped in PKIResponses
func AcmeRevocationConnectorResponseAsPKIResponses(v *AcmeRevocationConnectorResponse) PKIResponses {
	return PKIResponses{
		AcmeRevocationConnectorResponse: v,
	}
}

// CMPConnectorResponseAsPKIResponses is a convenience function that returns CMPConnectorResponse wrapped in PKIResponses
func CMPConnectorResponseAsPKIResponses(v *CMPConnectorResponse) PKIResponses {
	return PKIResponses{
		CMPConnectorResponse: v,
	}
}

// CertEuropeConnectorResponseAsPKIResponses is a convenience function that returns CertEuropeConnectorResponse wrapped in PKIResponses
func CertEuropeConnectorResponseAsPKIResponses(v *CertEuropeConnectorResponse) PKIResponses {
	return PKIResponses{
		CertEuropeConnectorResponse: v,
	}
}

// DigiCertConnectorResponseAsPKIResponses is a convenience function that returns DigiCertConnectorResponse wrapped in PKIResponses
func DigiCertConnectorResponseAsPKIResponses(v *DigiCertConnectorResponse) PKIResponses {
	return PKIResponses{
		DigiCertConnectorResponse: v,
	}
}

// EJBCAConnectorResponseAsPKIResponses is a convenience function that returns EJBCAConnectorResponse wrapped in PKIResponses
func EJBCAConnectorResponseAsPKIResponses(v *EJBCAConnectorResponse) PKIResponses {
	return PKIResponses{
		EJBCAConnectorResponse: v,
	}
}

// EntrustConnectorResponseAsPKIResponses is a convenience function that returns EntrustConnectorResponse wrapped in PKIResponses
func EntrustConnectorResponseAsPKIResponses(v *EntrustConnectorResponse) PKIResponses {
	return PKIResponses{
		EntrustConnectorResponse: v,
	}
}

// EverTrustADCSConnectorResponseAsPKIResponses is a convenience function that returns EverTrustADCSConnectorResponse wrapped in PKIResponses
func EverTrustADCSConnectorResponseAsPKIResponses(v *EverTrustADCSConnectorResponse) PKIResponses {
	return PKIResponses{
		EverTrustADCSConnectorResponse: v,
	}
}

// FCMSConnectorResponseAsPKIResponses is a convenience function that returns FCMSConnectorResponse wrapped in PKIResponses
func FCMSConnectorResponseAsPKIResponses(v *FCMSConnectorResponse) PKIResponses {
	return PKIResponses{
		FCMSConnectorResponse: v,
	}
}

// GSAtlasConnectorResponseAsPKIResponses is a convenience function that returns GSAtlasConnectorResponse wrapped in PKIResponses
func GSAtlasConnectorResponseAsPKIResponses(v *GSAtlasConnectorResponse) PKIResponses {
	return PKIResponses{
		GSAtlasConnectorResponse: v,
	}
}

// GSMSSLConnectorResponseAsPKIResponses is a convenience function that returns GSMSSLConnectorResponse wrapped in PKIResponses
func GSMSSLConnectorResponseAsPKIResponses(v *GSMSSLConnectorResponse) PKIResponses {
	return PKIResponses{
		GSMSSLConnectorResponse: v,
	}
}

// IDCAConnectorResponseAsPKIResponses is a convenience function that returns IDCAConnectorResponse wrapped in PKIResponses
func IDCAConnectorResponseAsPKIResponses(v *IDCAConnectorResponse) PKIResponses {
	return PKIResponses{
		IDCAConnectorResponse: v,
	}
}

// IntegratedCAConnectorResponseAsPKIResponses is a convenience function that returns IntegratedCAConnectorResponse wrapped in PKIResponses
func IntegratedCAConnectorResponseAsPKIResponses(v *IntegratedCAConnectorResponse) PKIResponses {
	return PKIResponses{
		IntegratedCAConnectorResponse: v,
	}
}

// MetaPKIConnectorResponseAsPKIResponses is a convenience function that returns MetaPKIConnectorResponse wrapped in PKIResponses
func MetaPKIConnectorResponseAsPKIResponses(v *MetaPKIConnectorResponse) PKIResponses {
	return PKIResponses{
		MetaPKIConnectorResponse: v,
	}
}

// NameshieldResponseAsPKIResponses is a convenience function that returns NameshieldResponse wrapped in PKIResponses
func NameshieldResponseAsPKIResponses(v *NameshieldResponse) PKIResponses {
	return PKIResponses{
		NameshieldResponse: v,
	}
}

// NexusCMConnectorResponseAsPKIResponses is a convenience function that returns NexusCMConnectorResponse wrapped in PKIResponses
func NexusCMConnectorResponseAsPKIResponses(v *NexusCMConnectorResponse) PKIResponses {
	return PKIResponses{
		NexusCMConnectorResponse: v,
	}
}

// OTPKIConnectorResponseAsPKIResponses is a convenience function that returns OTPKIConnectorResponse wrapped in PKIResponses
func OTPKIConnectorResponseAsPKIResponses(v *OTPKIConnectorResponse) PKIResponses {
	return PKIResponses{
		OTPKIConnectorResponse: v,
	}
}

// SectigoCMSConnectorResponseAsPKIResponses is a convenience function that returns SectigoCMSConnectorResponse wrapped in PKIResponses
func SectigoCMSConnectorResponseAsPKIResponses(v *SectigoCMSConnectorResponse) PKIResponses {
	return PKIResponses{
		SectigoCMSConnectorResponse: v,
	}
}

// StreamConnectorResponseAsPKIResponses is a convenience function that returns StreamConnectorResponse wrapped in PKIResponses
func StreamConnectorResponseAsPKIResponses(v *StreamConnectorResponse) PKIResponses {
	return PKIResponses{
		StreamConnectorResponse: v,
	}
}

// SwissSignConnectorResponseAsPKIResponses is a convenience function that returns SwissSignConnectorResponse wrapped in PKIResponses
func SwissSignConnectorResponseAsPKIResponses(v *SwissSignConnectorResponse) PKIResponses {
	return PKIResponses{
		SwissSignConnectorResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PKIResponses) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ADCSConnectorResponse
	err = json.Unmarshal(data, &dst.ADCSConnectorResponse)
	if err == nil {
		jsonADCSConnectorResponse, _ := json.Marshal(dst.ADCSConnectorResponse)
		if string(jsonADCSConnectorResponse) == "{}" { // empty struct
			dst.ADCSConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.ADCSConnectorResponse = nil
	}

	// try to unmarshal data into AWSACMPCAConnectorResponse
	err = json.Unmarshal(data, &dst.AWSACMPCAConnectorResponse)
	if err == nil {
		jsonAWSACMPCAConnectorResponse, _ := json.Marshal(dst.AWSACMPCAConnectorResponse)
		if string(jsonAWSACMPCAConnectorResponse) == "{}" { // empty struct
			dst.AWSACMPCAConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.AWSACMPCAConnectorResponse = nil
	}

	// try to unmarshal data into AcmeEnrollConnectorResponse
	err = json.Unmarshal(data, &dst.AcmeEnrollConnectorResponse)
	if err == nil {
		jsonAcmeEnrollConnectorResponse, _ := json.Marshal(dst.AcmeEnrollConnectorResponse)
		if string(jsonAcmeEnrollConnectorResponse) == "{}" { // empty struct
			dst.AcmeEnrollConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.AcmeEnrollConnectorResponse = nil
	}

	// try to unmarshal data into AcmeRevocationConnectorResponse
	err = json.Unmarshal(data, &dst.AcmeRevocationConnectorResponse)
	if err == nil {
		jsonAcmeRevocationConnectorResponse, _ := json.Marshal(dst.AcmeRevocationConnectorResponse)
		if string(jsonAcmeRevocationConnectorResponse) == "{}" { // empty struct
			dst.AcmeRevocationConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.AcmeRevocationConnectorResponse = nil
	}

	// try to unmarshal data into CMPConnectorResponse
	err = json.Unmarshal(data, &dst.CMPConnectorResponse)
	if err == nil {
		jsonCMPConnectorResponse, _ := json.Marshal(dst.CMPConnectorResponse)
		if string(jsonCMPConnectorResponse) == "{}" { // empty struct
			dst.CMPConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.CMPConnectorResponse = nil
	}

	// try to unmarshal data into CertEuropeConnectorResponse
	err = json.Unmarshal(data, &dst.CertEuropeConnectorResponse)
	if err == nil {
		jsonCertEuropeConnectorResponse, _ := json.Marshal(dst.CertEuropeConnectorResponse)
		if string(jsonCertEuropeConnectorResponse) == "{}" { // empty struct
			dst.CertEuropeConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.CertEuropeConnectorResponse = nil
	}

	// try to unmarshal data into DigiCertConnectorResponse
	err = json.Unmarshal(data, &dst.DigiCertConnectorResponse)
	if err == nil {
		jsonDigiCertConnectorResponse, _ := json.Marshal(dst.DigiCertConnectorResponse)
		if string(jsonDigiCertConnectorResponse) == "{}" { // empty struct
			dst.DigiCertConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.DigiCertConnectorResponse = nil
	}

	// try to unmarshal data into EJBCAConnectorResponse
	err = json.Unmarshal(data, &dst.EJBCAConnectorResponse)
	if err == nil {
		jsonEJBCAConnectorResponse, _ := json.Marshal(dst.EJBCAConnectorResponse)
		if string(jsonEJBCAConnectorResponse) == "{}" { // empty struct
			dst.EJBCAConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.EJBCAConnectorResponse = nil
	}

	// try to unmarshal data into EntrustConnectorResponse
	err = json.Unmarshal(data, &dst.EntrustConnectorResponse)
	if err == nil {
		jsonEntrustConnectorResponse, _ := json.Marshal(dst.EntrustConnectorResponse)
		if string(jsonEntrustConnectorResponse) == "{}" { // empty struct
			dst.EntrustConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.EntrustConnectorResponse = nil
	}

	// try to unmarshal data into EverTrustADCSConnectorResponse
	err = json.Unmarshal(data, &dst.EverTrustADCSConnectorResponse)
	if err == nil {
		jsonEverTrustADCSConnectorResponse, _ := json.Marshal(dst.EverTrustADCSConnectorResponse)
		if string(jsonEverTrustADCSConnectorResponse) == "{}" { // empty struct
			dst.EverTrustADCSConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.EverTrustADCSConnectorResponse = nil
	}

	// try to unmarshal data into FCMSConnectorResponse
	err = json.Unmarshal(data, &dst.FCMSConnectorResponse)
	if err == nil {
		jsonFCMSConnectorResponse, _ := json.Marshal(dst.FCMSConnectorResponse)
		if string(jsonFCMSConnectorResponse) == "{}" { // empty struct
			dst.FCMSConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.FCMSConnectorResponse = nil
	}

	// try to unmarshal data into GSAtlasConnectorResponse
	err = json.Unmarshal(data, &dst.GSAtlasConnectorResponse)
	if err == nil {
		jsonGSAtlasConnectorResponse, _ := json.Marshal(dst.GSAtlasConnectorResponse)
		if string(jsonGSAtlasConnectorResponse) == "{}" { // empty struct
			dst.GSAtlasConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.GSAtlasConnectorResponse = nil
	}

	// try to unmarshal data into GSMSSLConnectorResponse
	err = json.Unmarshal(data, &dst.GSMSSLConnectorResponse)
	if err == nil {
		jsonGSMSSLConnectorResponse, _ := json.Marshal(dst.GSMSSLConnectorResponse)
		if string(jsonGSMSSLConnectorResponse) == "{}" { // empty struct
			dst.GSMSSLConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.GSMSSLConnectorResponse = nil
	}

	// try to unmarshal data into IDCAConnectorResponse
	err = json.Unmarshal(data, &dst.IDCAConnectorResponse)
	if err == nil {
		jsonIDCAConnectorResponse, _ := json.Marshal(dst.IDCAConnectorResponse)
		if string(jsonIDCAConnectorResponse) == "{}" { // empty struct
			dst.IDCAConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.IDCAConnectorResponse = nil
	}

	// try to unmarshal data into IntegratedCAConnectorResponse
	err = json.Unmarshal(data, &dst.IntegratedCAConnectorResponse)
	if err == nil {
		jsonIntegratedCAConnectorResponse, _ := json.Marshal(dst.IntegratedCAConnectorResponse)
		if string(jsonIntegratedCAConnectorResponse) == "{}" { // empty struct
			dst.IntegratedCAConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.IntegratedCAConnectorResponse = nil
	}

	// try to unmarshal data into MetaPKIConnectorResponse
	err = json.Unmarshal(data, &dst.MetaPKIConnectorResponse)
	if err == nil {
		jsonMetaPKIConnectorResponse, _ := json.Marshal(dst.MetaPKIConnectorResponse)
		if string(jsonMetaPKIConnectorResponse) == "{}" { // empty struct
			dst.MetaPKIConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.MetaPKIConnectorResponse = nil
	}

	// try to unmarshal data into NameshieldResponse
	err = json.Unmarshal(data, &dst.NameshieldResponse)
	if err == nil {
		jsonNameshieldResponse, _ := json.Marshal(dst.NameshieldResponse)
		if string(jsonNameshieldResponse) == "{}" { // empty struct
			dst.NameshieldResponse = nil
		} else {
			match++
		}
	} else {
		dst.NameshieldResponse = nil
	}

	// try to unmarshal data into NexusCMConnectorResponse
	err = json.Unmarshal(data, &dst.NexusCMConnectorResponse)
	if err == nil {
		jsonNexusCMConnectorResponse, _ := json.Marshal(dst.NexusCMConnectorResponse)
		if string(jsonNexusCMConnectorResponse) == "{}" { // empty struct
			dst.NexusCMConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.NexusCMConnectorResponse = nil
	}

	// try to unmarshal data into OTPKIConnectorResponse
	err = json.Unmarshal(data, &dst.OTPKIConnectorResponse)
	if err == nil {
		jsonOTPKIConnectorResponse, _ := json.Marshal(dst.OTPKIConnectorResponse)
		if string(jsonOTPKIConnectorResponse) == "{}" { // empty struct
			dst.OTPKIConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.OTPKIConnectorResponse = nil
	}

	// try to unmarshal data into SectigoCMSConnectorResponse
	err = json.Unmarshal(data, &dst.SectigoCMSConnectorResponse)
	if err == nil {
		jsonSectigoCMSConnectorResponse, _ := json.Marshal(dst.SectigoCMSConnectorResponse)
		if string(jsonSectigoCMSConnectorResponse) == "{}" { // empty struct
			dst.SectigoCMSConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.SectigoCMSConnectorResponse = nil
	}

	// try to unmarshal data into StreamConnectorResponse
	err = json.Unmarshal(data, &dst.StreamConnectorResponse)
	if err == nil {
		jsonStreamConnectorResponse, _ := json.Marshal(dst.StreamConnectorResponse)
		if string(jsonStreamConnectorResponse) == "{}" { // empty struct
			dst.StreamConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.StreamConnectorResponse = nil
	}

	// try to unmarshal data into SwissSignConnectorResponse
	err = json.Unmarshal(data, &dst.SwissSignConnectorResponse)
	if err == nil {
		jsonSwissSignConnectorResponse, _ := json.Marshal(dst.SwissSignConnectorResponse)
		if string(jsonSwissSignConnectorResponse) == "{}" { // empty struct
			dst.SwissSignConnectorResponse = nil
		} else {
			match++
		}
	} else {
		dst.SwissSignConnectorResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PKIResponses)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PKIResponses) MarshalJSON() ([]byte, error) {
	if src.ADCSConnectorResponse != nil {
		return json.Marshal(&src.ADCSConnectorResponse)
	}

	if src.AWSACMPCAConnectorResponse != nil {
		return json.Marshal(&src.AWSACMPCAConnectorResponse)
	}

	if src.AcmeEnrollConnectorResponse != nil {
		return json.Marshal(&src.AcmeEnrollConnectorResponse)
	}

	if src.AcmeRevocationConnectorResponse != nil {
		return json.Marshal(&src.AcmeRevocationConnectorResponse)
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

	if src.IDCAConnectorResponse != nil {
		return json.Marshal(&src.IDCAConnectorResponse)
	}

	if src.IntegratedCAConnectorResponse != nil {
		return json.Marshal(&src.IntegratedCAConnectorResponse)
	}

	if src.MetaPKIConnectorResponse != nil {
		return json.Marshal(&src.MetaPKIConnectorResponse)
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

	if src.SectigoCMSConnectorResponse != nil {
		return json.Marshal(&src.SectigoCMSConnectorResponse)
	}

	if src.StreamConnectorResponse != nil {
		return json.Marshal(&src.StreamConnectorResponse)
	}

	if src.SwissSignConnectorResponse != nil {
		return json.Marshal(&src.SwissSignConnectorResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PKIResponses) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ADCSConnectorResponse != nil {
		return obj.ADCSConnectorResponse
	}

	if obj.AWSACMPCAConnectorResponse != nil {
		return obj.AWSACMPCAConnectorResponse
	}

	if obj.AcmeEnrollConnectorResponse != nil {
		return obj.AcmeEnrollConnectorResponse
	}

	if obj.AcmeRevocationConnectorResponse != nil {
		return obj.AcmeRevocationConnectorResponse
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

	if obj.IDCAConnectorResponse != nil {
		return obj.IDCAConnectorResponse
	}

	if obj.IntegratedCAConnectorResponse != nil {
		return obj.IntegratedCAConnectorResponse
	}

	if obj.MetaPKIConnectorResponse != nil {
		return obj.MetaPKIConnectorResponse
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

	if obj.SectigoCMSConnectorResponse != nil {
		return obj.SectigoCMSConnectorResponse
	}

	if obj.StreamConnectorResponse != nil {
		return obj.StreamConnectorResponse
	}

	if obj.SwissSignConnectorResponse != nil {
		return obj.SwissSignConnectorResponse
	}

	// all schemas are nil
	return nil
}

type NullablePKIResponses struct {
	value *PKIResponses
	isSet bool
}

func (v NullablePKIResponses) Get() *PKIResponses {
	return v.value
}

func (v *NullablePKIResponses) Set(val *PKIResponses) {
	v.value = val
	v.isSet = true
}

func (v NullablePKIResponses) IsSet() bool {
	return v.isSet
}

func (v *NullablePKIResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePKIResponses(val *PKIResponses) *NullablePKIResponses {
	return &NullablePKIResponses{value: val, isSet: true}
}

func (v NullablePKIResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePKIResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
