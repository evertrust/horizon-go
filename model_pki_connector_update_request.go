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

// PkiConnectorUpdateRequest - struct for PkiConnectorUpdateRequest
type PkiConnectorUpdateRequest struct {
	ADCSConnector *ADCSConnector
	AWSACMPCAConnector *AWSACMPCAConnector
	AcmeRevocationConnector *AcmeRevocationConnector
	CMPConnector *CMPConnector
	CertEuropeConnector *CertEuropeConnector
	DigiCertConnector *DigiCertConnector
	EJBCAConnector *EJBCAConnector
	EntrustConnector *EntrustConnector
	EverTrustADCSConnector *EverTrustADCSConnector
	FCMSConnector *FCMSConnector
	GSAtlasConnector *GSAtlasConnector
	GSMSSLConnector *GSMSSLConnector
	IDCAConnector *IDCAConnector
	IntegratedCAConnector *IntegratedCAConnector
	MetaPKIConnector *MetaPKIConnector
	Nameshield *Nameshield
	NexusCMConnector *NexusCMConnector
	OTPKIConnector *OTPKIConnector
	SectigoCMSConnector *SectigoCMSConnector
	StreamConnector *StreamConnector
}

// ADCSConnectorAsPkiConnectorUpdateRequest is a convenience function that returns ADCSConnector wrapped in PkiConnectorUpdateRequest
func ADCSConnectorAsPkiConnectorUpdateRequest(v *ADCSConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		ADCSConnector: v,
	}
}

// AWSACMPCAConnectorAsPkiConnectorUpdateRequest is a convenience function that returns AWSACMPCAConnector wrapped in PkiConnectorUpdateRequest
func AWSACMPCAConnectorAsPkiConnectorUpdateRequest(v *AWSACMPCAConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		AWSACMPCAConnector: v,
	}
}

// AcmeRevocationConnectorAsPkiConnectorUpdateRequest is a convenience function that returns AcmeRevocationConnector wrapped in PkiConnectorUpdateRequest
func AcmeRevocationConnectorAsPkiConnectorUpdateRequest(v *AcmeRevocationConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		AcmeRevocationConnector: v,
	}
}

// CMPConnectorAsPkiConnectorUpdateRequest is a convenience function that returns CMPConnector wrapped in PkiConnectorUpdateRequest
func CMPConnectorAsPkiConnectorUpdateRequest(v *CMPConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		CMPConnector: v,
	}
}

// CertEuropeConnectorAsPkiConnectorUpdateRequest is a convenience function that returns CertEuropeConnector wrapped in PkiConnectorUpdateRequest
func CertEuropeConnectorAsPkiConnectorUpdateRequest(v *CertEuropeConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		CertEuropeConnector: v,
	}
}

// DigiCertConnectorAsPkiConnectorUpdateRequest is a convenience function that returns DigiCertConnector wrapped in PkiConnectorUpdateRequest
func DigiCertConnectorAsPkiConnectorUpdateRequest(v *DigiCertConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		DigiCertConnector: v,
	}
}

// EJBCAConnectorAsPkiConnectorUpdateRequest is a convenience function that returns EJBCAConnector wrapped in PkiConnectorUpdateRequest
func EJBCAConnectorAsPkiConnectorUpdateRequest(v *EJBCAConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		EJBCAConnector: v,
	}
}

// EntrustConnectorAsPkiConnectorUpdateRequest is a convenience function that returns EntrustConnector wrapped in PkiConnectorUpdateRequest
func EntrustConnectorAsPkiConnectorUpdateRequest(v *EntrustConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		EntrustConnector: v,
	}
}

// EverTrustADCSConnectorAsPkiConnectorUpdateRequest is a convenience function that returns EverTrustADCSConnector wrapped in PkiConnectorUpdateRequest
func EverTrustADCSConnectorAsPkiConnectorUpdateRequest(v *EverTrustADCSConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		EverTrustADCSConnector: v,
	}
}

// FCMSConnectorAsPkiConnectorUpdateRequest is a convenience function that returns FCMSConnector wrapped in PkiConnectorUpdateRequest
func FCMSConnectorAsPkiConnectorUpdateRequest(v *FCMSConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		FCMSConnector: v,
	}
}

// GSAtlasConnectorAsPkiConnectorUpdateRequest is a convenience function that returns GSAtlasConnector wrapped in PkiConnectorUpdateRequest
func GSAtlasConnectorAsPkiConnectorUpdateRequest(v *GSAtlasConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		GSAtlasConnector: v,
	}
}

// GSMSSLConnectorAsPkiConnectorUpdateRequest is a convenience function that returns GSMSSLConnector wrapped in PkiConnectorUpdateRequest
func GSMSSLConnectorAsPkiConnectorUpdateRequest(v *GSMSSLConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		GSMSSLConnector: v,
	}
}

// IDCAConnectorAsPkiConnectorUpdateRequest is a convenience function that returns IDCAConnector wrapped in PkiConnectorUpdateRequest
func IDCAConnectorAsPkiConnectorUpdateRequest(v *IDCAConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		IDCAConnector: v,
	}
}

// IntegratedCAConnectorAsPkiConnectorUpdateRequest is a convenience function that returns IntegratedCAConnector wrapped in PkiConnectorUpdateRequest
func IntegratedCAConnectorAsPkiConnectorUpdateRequest(v *IntegratedCAConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		IntegratedCAConnector: v,
	}
}

// MetaPKIConnectorAsPkiConnectorUpdateRequest is a convenience function that returns MetaPKIConnector wrapped in PkiConnectorUpdateRequest
func MetaPKIConnectorAsPkiConnectorUpdateRequest(v *MetaPKIConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		MetaPKIConnector: v,
	}
}

// NameshieldAsPkiConnectorUpdateRequest is a convenience function that returns Nameshield wrapped in PkiConnectorUpdateRequest
func NameshieldAsPkiConnectorUpdateRequest(v *Nameshield) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		Nameshield: v,
	}
}

// NexusCMConnectorAsPkiConnectorUpdateRequest is a convenience function that returns NexusCMConnector wrapped in PkiConnectorUpdateRequest
func NexusCMConnectorAsPkiConnectorUpdateRequest(v *NexusCMConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		NexusCMConnector: v,
	}
}

// OTPKIConnectorAsPkiConnectorUpdateRequest is a convenience function that returns OTPKIConnector wrapped in PkiConnectorUpdateRequest
func OTPKIConnectorAsPkiConnectorUpdateRequest(v *OTPKIConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		OTPKIConnector: v,
	}
}

// SectigoCMSConnectorAsPkiConnectorUpdateRequest is a convenience function that returns SectigoCMSConnector wrapped in PkiConnectorUpdateRequest
func SectigoCMSConnectorAsPkiConnectorUpdateRequest(v *SectigoCMSConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		SectigoCMSConnector: v,
	}
}

// StreamConnectorAsPkiConnectorUpdateRequest is a convenience function that returns StreamConnector wrapped in PkiConnectorUpdateRequest
func StreamConnectorAsPkiConnectorUpdateRequest(v *StreamConnector) PkiConnectorUpdateRequest {
	return PkiConnectorUpdateRequest{
		StreamConnector: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PkiConnectorUpdateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ADCSConnector
	err = newStrictDecoder(data).Decode(&dst.ADCSConnector)
	if err == nil {
		jsonADCSConnector, _ := json.Marshal(dst.ADCSConnector)
		if string(jsonADCSConnector) == "{}" { // empty struct
			dst.ADCSConnector = nil
		} else {
            _ = validator.Validate(dst.ADCSConnector)
            match++
		}
	} else {
		dst.ADCSConnector = nil
	}

	// try to unmarshal data into AWSACMPCAConnector
	err = newStrictDecoder(data).Decode(&dst.AWSACMPCAConnector)
	if err == nil {
		jsonAWSACMPCAConnector, _ := json.Marshal(dst.AWSACMPCAConnector)
		if string(jsonAWSACMPCAConnector) == "{}" { // empty struct
			dst.AWSACMPCAConnector = nil
		} else {
            _ = validator.Validate(dst.AWSACMPCAConnector)
            match++
		}
	} else {
		dst.AWSACMPCAConnector = nil
	}

	// try to unmarshal data into AcmeRevocationConnector
	err = newStrictDecoder(data).Decode(&dst.AcmeRevocationConnector)
	if err == nil {
		jsonAcmeRevocationConnector, _ := json.Marshal(dst.AcmeRevocationConnector)
		if string(jsonAcmeRevocationConnector) == "{}" { // empty struct
			dst.AcmeRevocationConnector = nil
		} else {
            _ = validator.Validate(dst.AcmeRevocationConnector)
            match++
		}
	} else {
		dst.AcmeRevocationConnector = nil
	}

	// try to unmarshal data into CMPConnector
	err = newStrictDecoder(data).Decode(&dst.CMPConnector)
	if err == nil {
		jsonCMPConnector, _ := json.Marshal(dst.CMPConnector)
		if string(jsonCMPConnector) == "{}" { // empty struct
			dst.CMPConnector = nil
		} else {
            _ = validator.Validate(dst.CMPConnector)
            match++
		}
	} else {
		dst.CMPConnector = nil
	}

	// try to unmarshal data into CertEuropeConnector
	err = newStrictDecoder(data).Decode(&dst.CertEuropeConnector)
	if err == nil {
		jsonCertEuropeConnector, _ := json.Marshal(dst.CertEuropeConnector)
		if string(jsonCertEuropeConnector) == "{}" { // empty struct
			dst.CertEuropeConnector = nil
		} else {
            _ = validator.Validate(dst.CertEuropeConnector)
            match++
		}
	} else {
		dst.CertEuropeConnector = nil
	}

	// try to unmarshal data into DigiCertConnector
	err = newStrictDecoder(data).Decode(&dst.DigiCertConnector)
	if err == nil {
		jsonDigiCertConnector, _ := json.Marshal(dst.DigiCertConnector)
		if string(jsonDigiCertConnector) == "{}" { // empty struct
			dst.DigiCertConnector = nil
		} else {
            _ = validator.Validate(dst.DigiCertConnector)
            match++
		}
	} else {
		dst.DigiCertConnector = nil
	}

	// try to unmarshal data into EJBCAConnector
	err = newStrictDecoder(data).Decode(&dst.EJBCAConnector)
	if err == nil {
		jsonEJBCAConnector, _ := json.Marshal(dst.EJBCAConnector)
		if string(jsonEJBCAConnector) == "{}" { // empty struct
			dst.EJBCAConnector = nil
		} else {
            _ = validator.Validate(dst.EJBCAConnector)
            match++
		}
	} else {
		dst.EJBCAConnector = nil
	}

	// try to unmarshal data into EntrustConnector
	err = newStrictDecoder(data).Decode(&dst.EntrustConnector)
	if err == nil {
		jsonEntrustConnector, _ := json.Marshal(dst.EntrustConnector)
		if string(jsonEntrustConnector) == "{}" { // empty struct
			dst.EntrustConnector = nil
		} else {
            _ = validator.Validate(dst.EntrustConnector)
            match++
		}
	} else {
		dst.EntrustConnector = nil
	}

	// try to unmarshal data into EverTrustADCSConnector
	err = newStrictDecoder(data).Decode(&dst.EverTrustADCSConnector)
	if err == nil {
		jsonEverTrustADCSConnector, _ := json.Marshal(dst.EverTrustADCSConnector)
		if string(jsonEverTrustADCSConnector) == "{}" { // empty struct
			dst.EverTrustADCSConnector = nil
		} else {
            _ = validator.Validate(dst.EverTrustADCSConnector)
            match++
		}
	} else {
		dst.EverTrustADCSConnector = nil
	}

	// try to unmarshal data into FCMSConnector
	err = newStrictDecoder(data).Decode(&dst.FCMSConnector)
	if err == nil {
		jsonFCMSConnector, _ := json.Marshal(dst.FCMSConnector)
		if string(jsonFCMSConnector) == "{}" { // empty struct
			dst.FCMSConnector = nil
		} else {
            _ = validator.Validate(dst.FCMSConnector)
            match++
		}
	} else {
		dst.FCMSConnector = nil
	}

	// try to unmarshal data into GSAtlasConnector
	err = newStrictDecoder(data).Decode(&dst.GSAtlasConnector)
	if err == nil {
		jsonGSAtlasConnector, _ := json.Marshal(dst.GSAtlasConnector)
		if string(jsonGSAtlasConnector) == "{}" { // empty struct
			dst.GSAtlasConnector = nil
		} else {
            _ = validator.Validate(dst.GSAtlasConnector)
            match++
		}
	} else {
		dst.GSAtlasConnector = nil
	}

	// try to unmarshal data into GSMSSLConnector
	err = newStrictDecoder(data).Decode(&dst.GSMSSLConnector)
	if err == nil {
		jsonGSMSSLConnector, _ := json.Marshal(dst.GSMSSLConnector)
		if string(jsonGSMSSLConnector) == "{}" { // empty struct
			dst.GSMSSLConnector = nil
		} else {
            _ = validator.Validate(dst.GSMSSLConnector)
            match++
		}
	} else {
		dst.GSMSSLConnector = nil
	}

	// try to unmarshal data into IDCAConnector
	err = newStrictDecoder(data).Decode(&dst.IDCAConnector)
	if err == nil {
		jsonIDCAConnector, _ := json.Marshal(dst.IDCAConnector)
		if string(jsonIDCAConnector) == "{}" { // empty struct
			dst.IDCAConnector = nil
		} else {
            _ = validator.Validate(dst.IDCAConnector)
            match++
		}
	} else {
		dst.IDCAConnector = nil
	}

	// try to unmarshal data into IntegratedCAConnector
	err = newStrictDecoder(data).Decode(&dst.IntegratedCAConnector)
	if err == nil {
		jsonIntegratedCAConnector, _ := json.Marshal(dst.IntegratedCAConnector)
		if string(jsonIntegratedCAConnector) == "{}" { // empty struct
			dst.IntegratedCAConnector = nil
		} else {
            _ = validator.Validate(dst.IntegratedCAConnector)
            match++
		}
	} else {
		dst.IntegratedCAConnector = nil
	}

	// try to unmarshal data into MetaPKIConnector
	err = newStrictDecoder(data).Decode(&dst.MetaPKIConnector)
	if err == nil {
		jsonMetaPKIConnector, _ := json.Marshal(dst.MetaPKIConnector)
		if string(jsonMetaPKIConnector) == "{}" { // empty struct
			dst.MetaPKIConnector = nil
		} else {
            _ = validator.Validate(dst.MetaPKIConnector)
            match++
		}
	} else {
		dst.MetaPKIConnector = nil
	}

	// try to unmarshal data into Nameshield
	err = newStrictDecoder(data).Decode(&dst.Nameshield)
	if err == nil {
		jsonNameshield, _ := json.Marshal(dst.Nameshield)
		if string(jsonNameshield) == "{}" { // empty struct
			dst.Nameshield = nil
		} else {
            _ = validator.Validate(dst.Nameshield)
            match++
		}
	} else {
		dst.Nameshield = nil
	}

	// try to unmarshal data into NexusCMConnector
	err = newStrictDecoder(data).Decode(&dst.NexusCMConnector)
	if err == nil {
		jsonNexusCMConnector, _ := json.Marshal(dst.NexusCMConnector)
		if string(jsonNexusCMConnector) == "{}" { // empty struct
			dst.NexusCMConnector = nil
		} else {
            _ = validator.Validate(dst.NexusCMConnector)
            match++
		}
	} else {
		dst.NexusCMConnector = nil
	}

	// try to unmarshal data into OTPKIConnector
	err = newStrictDecoder(data).Decode(&dst.OTPKIConnector)
	if err == nil {
		jsonOTPKIConnector, _ := json.Marshal(dst.OTPKIConnector)
		if string(jsonOTPKIConnector) == "{}" { // empty struct
			dst.OTPKIConnector = nil
		} else {
            _ = validator.Validate(dst.OTPKIConnector)
            match++
		}
	} else {
		dst.OTPKIConnector = nil
	}

	// try to unmarshal data into SectigoCMSConnector
	err = newStrictDecoder(data).Decode(&dst.SectigoCMSConnector)
	if err == nil {
		jsonSectigoCMSConnector, _ := json.Marshal(dst.SectigoCMSConnector)
		if string(jsonSectigoCMSConnector) == "{}" { // empty struct
			dst.SectigoCMSConnector = nil
		} else {
            _ = validator.Validate(dst.SectigoCMSConnector)
            match++
		}
	} else {
		dst.SectigoCMSConnector = nil
	}

	// try to unmarshal data into StreamConnector
	err = newStrictDecoder(data).Decode(&dst.StreamConnector)
	if err == nil {
		jsonStreamConnector, _ := json.Marshal(dst.StreamConnector)
		if string(jsonStreamConnector) == "{}" { // empty struct
			dst.StreamConnector = nil
		} else {
            _ = validator.Validate(dst.StreamConnector)
            match++
		}
	} else {
		dst.StreamConnector = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PkiConnectorUpdateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PkiConnectorUpdateRequest) MarshalJSON() ([]byte, error) {
	if src.ADCSConnector != nil {
		return json.Marshal(&src.ADCSConnector)
	}

	if src.AWSACMPCAConnector != nil {
		return json.Marshal(&src.AWSACMPCAConnector)
	}

	if src.AcmeRevocationConnector != nil {
		return json.Marshal(&src.AcmeRevocationConnector)
	}

	if src.CMPConnector != nil {
		return json.Marshal(&src.CMPConnector)
	}

	if src.CertEuropeConnector != nil {
		return json.Marshal(&src.CertEuropeConnector)
	}

	if src.DigiCertConnector != nil {
		return json.Marshal(&src.DigiCertConnector)
	}

	if src.EJBCAConnector != nil {
		return json.Marshal(&src.EJBCAConnector)
	}

	if src.EntrustConnector != nil {
		return json.Marshal(&src.EntrustConnector)
	}

	if src.EverTrustADCSConnector != nil {
		return json.Marshal(&src.EverTrustADCSConnector)
	}

	if src.FCMSConnector != nil {
		return json.Marshal(&src.FCMSConnector)
	}

	if src.GSAtlasConnector != nil {
		return json.Marshal(&src.GSAtlasConnector)
	}

	if src.GSMSSLConnector != nil {
		return json.Marshal(&src.GSMSSLConnector)
	}

	if src.IDCAConnector != nil {
		return json.Marshal(&src.IDCAConnector)
	}

	if src.IntegratedCAConnector != nil {
		return json.Marshal(&src.IntegratedCAConnector)
	}

	if src.MetaPKIConnector != nil {
		return json.Marshal(&src.MetaPKIConnector)
	}

	if src.Nameshield != nil {
		return json.Marshal(&src.Nameshield)
	}

	if src.NexusCMConnector != nil {
		return json.Marshal(&src.NexusCMConnector)
	}

	if src.OTPKIConnector != nil {
		return json.Marshal(&src.OTPKIConnector)
	}

	if src.SectigoCMSConnector != nil {
		return json.Marshal(&src.SectigoCMSConnector)
	}

	if src.StreamConnector != nil {
		return json.Marshal(&src.StreamConnector)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PkiConnectorUpdateRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ADCSConnector != nil {
		return obj.ADCSConnector
	}

	if obj.AWSACMPCAConnector != nil {
		return obj.AWSACMPCAConnector
	}

	if obj.AcmeRevocationConnector != nil {
		return obj.AcmeRevocationConnector
	}

	if obj.CMPConnector != nil {
		return obj.CMPConnector
	}

	if obj.CertEuropeConnector != nil {
		return obj.CertEuropeConnector
	}

	if obj.DigiCertConnector != nil {
		return obj.DigiCertConnector
	}

	if obj.EJBCAConnector != nil {
		return obj.EJBCAConnector
	}

	if obj.EntrustConnector != nil {
		return obj.EntrustConnector
	}

	if obj.EverTrustADCSConnector != nil {
		return obj.EverTrustADCSConnector
	}

	if obj.FCMSConnector != nil {
		return obj.FCMSConnector
	}

	if obj.GSAtlasConnector != nil {
		return obj.GSAtlasConnector
	}

	if obj.GSMSSLConnector != nil {
		return obj.GSMSSLConnector
	}

	if obj.IDCAConnector != nil {
		return obj.IDCAConnector
	}

	if obj.IntegratedCAConnector != nil {
		return obj.IntegratedCAConnector
	}

	if obj.MetaPKIConnector != nil {
		return obj.MetaPKIConnector
	}

	if obj.Nameshield != nil {
		return obj.Nameshield
	}

	if obj.NexusCMConnector != nil {
		return obj.NexusCMConnector
	}

	if obj.OTPKIConnector != nil {
		return obj.OTPKIConnector
	}

	if obj.SectigoCMSConnector != nil {
		return obj.SectigoCMSConnector
	}

	if obj.StreamConnector != nil {
		return obj.StreamConnector
	}

	// all schemas are nil
	return nil
}

type NullablePkiConnectorUpdateRequest struct {
	value *PkiConnectorUpdateRequest
	isSet bool
}

func (v NullablePkiConnectorUpdateRequest) Get() *PkiConnectorUpdateRequest {
	return v.value
}

func (v *NullablePkiConnectorUpdateRequest) Set(val *PkiConnectorUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiConnectorUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiConnectorUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiConnectorUpdateRequest(val *PkiConnectorUpdateRequest) *NullablePkiConnectorUpdateRequest {
	return &NullablePkiConnectorUpdateRequest{value: val, isSet: true}
}

func (v NullablePkiConnectorUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiConnectorUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


