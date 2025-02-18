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

// CertificateProfileList200ResponseInner - struct for CertificateProfileList200ResponseInner
type CertificateProfileList200ResponseInner struct {
	AcmeExternalProfileResponse *AcmeExternalProfileResponse
	AcmeProfileResponse *AcmeProfileResponse
	CrmpProfileResponse *CrmpProfileResponse
	EstProfileResponse *EstProfileResponse
	IntunePKCSProfileResponse *IntunePKCSProfileResponse
	IntuneProfileResponse *IntuneProfileResponse
	JamfProfileResponse *JamfProfileResponse
	ScepProfileResponse *ScepProfileResponse
	WcceProfileResponse *WcceProfileResponse
	WebRAProfileResponse *WebRAProfileResponse
}

// AcmeExternalProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns AcmeExternalProfileResponse wrapped in CertificateProfileList200ResponseInner
func AcmeExternalProfileResponseAsCertificateProfileList200ResponseInner(v *AcmeExternalProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		AcmeExternalProfileResponse: v,
	}
}

// AcmeProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns AcmeProfileResponse wrapped in CertificateProfileList200ResponseInner
func AcmeProfileResponseAsCertificateProfileList200ResponseInner(v *AcmeProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		AcmeProfileResponse: v,
	}
}

// CrmpProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns CrmpProfileResponse wrapped in CertificateProfileList200ResponseInner
func CrmpProfileResponseAsCertificateProfileList200ResponseInner(v *CrmpProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		CrmpProfileResponse: v,
	}
}

// EstProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns EstProfileResponse wrapped in CertificateProfileList200ResponseInner
func EstProfileResponseAsCertificateProfileList200ResponseInner(v *EstProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		EstProfileResponse: v,
	}
}

// IntunePKCSProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns IntunePKCSProfileResponse wrapped in CertificateProfileList200ResponseInner
func IntunePKCSProfileResponseAsCertificateProfileList200ResponseInner(v *IntunePKCSProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		IntunePKCSProfileResponse: v,
	}
}

// IntuneProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns IntuneProfileResponse wrapped in CertificateProfileList200ResponseInner
func IntuneProfileResponseAsCertificateProfileList200ResponseInner(v *IntuneProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		IntuneProfileResponse: v,
	}
}

// JamfProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns JamfProfileResponse wrapped in CertificateProfileList200ResponseInner
func JamfProfileResponseAsCertificateProfileList200ResponseInner(v *JamfProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		JamfProfileResponse: v,
	}
}

// ScepProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns ScepProfileResponse wrapped in CertificateProfileList200ResponseInner
func ScepProfileResponseAsCertificateProfileList200ResponseInner(v *ScepProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		ScepProfileResponse: v,
	}
}

// WcceProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns WcceProfileResponse wrapped in CertificateProfileList200ResponseInner
func WcceProfileResponseAsCertificateProfileList200ResponseInner(v *WcceProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		WcceProfileResponse: v,
	}
}

// WebRAProfileResponseAsCertificateProfileList200ResponseInner is a convenience function that returns WebRAProfileResponse wrapped in CertificateProfileList200ResponseInner
func WebRAProfileResponseAsCertificateProfileList200ResponseInner(v *WebRAProfileResponse) CertificateProfileList200ResponseInner {
	return CertificateProfileList200ResponseInner{
		WebRAProfileResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CertificateProfileList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AcmeExternalProfileResponse
	err = newStrictDecoder(data).Decode(&dst.AcmeExternalProfileResponse)
	if err == nil {
		jsonAcmeExternalProfileResponse, _ := json.Marshal(dst.AcmeExternalProfileResponse)
		if string(jsonAcmeExternalProfileResponse) == "{}" { // empty struct
			dst.AcmeExternalProfileResponse = nil
		} else {
            _ = validator.Validate(dst.AcmeExternalProfileResponse)
            match++
		}
	} else {
		dst.AcmeExternalProfileResponse = nil
	}

	// try to unmarshal data into AcmeProfileResponse
	err = newStrictDecoder(data).Decode(&dst.AcmeProfileResponse)
	if err == nil {
		jsonAcmeProfileResponse, _ := json.Marshal(dst.AcmeProfileResponse)
		if string(jsonAcmeProfileResponse) == "{}" { // empty struct
			dst.AcmeProfileResponse = nil
		} else {
            _ = validator.Validate(dst.AcmeProfileResponse)
            match++
		}
	} else {
		dst.AcmeProfileResponse = nil
	}

	// try to unmarshal data into CrmpProfileResponse
	err = newStrictDecoder(data).Decode(&dst.CrmpProfileResponse)
	if err == nil {
		jsonCrmpProfileResponse, _ := json.Marshal(dst.CrmpProfileResponse)
		if string(jsonCrmpProfileResponse) == "{}" { // empty struct
			dst.CrmpProfileResponse = nil
		} else {
            _ = validator.Validate(dst.CrmpProfileResponse)
            match++
		}
	} else {
		dst.CrmpProfileResponse = nil
	}

	// try to unmarshal data into EstProfileResponse
	err = newStrictDecoder(data).Decode(&dst.EstProfileResponse)
	if err == nil {
		jsonEstProfileResponse, _ := json.Marshal(dst.EstProfileResponse)
		if string(jsonEstProfileResponse) == "{}" { // empty struct
			dst.EstProfileResponse = nil
		} else {
            _ = validator.Validate(dst.EstProfileResponse)
            match++
		}
	} else {
		dst.EstProfileResponse = nil
	}

	// try to unmarshal data into IntunePKCSProfileResponse
	err = newStrictDecoder(data).Decode(&dst.IntunePKCSProfileResponse)
	if err == nil {
		jsonIntunePKCSProfileResponse, _ := json.Marshal(dst.IntunePKCSProfileResponse)
		if string(jsonIntunePKCSProfileResponse) == "{}" { // empty struct
			dst.IntunePKCSProfileResponse = nil
		} else {
            _ = validator.Validate(dst.IntunePKCSProfileResponse)
            match++
		}
	} else {
		dst.IntunePKCSProfileResponse = nil
	}

	// try to unmarshal data into IntuneProfileResponse
	err = newStrictDecoder(data).Decode(&dst.IntuneProfileResponse)
	if err == nil {
		jsonIntuneProfileResponse, _ := json.Marshal(dst.IntuneProfileResponse)
		if string(jsonIntuneProfileResponse) == "{}" { // empty struct
			dst.IntuneProfileResponse = nil
		} else {
            _ = validator.Validate(dst.IntuneProfileResponse)
            match++
		}
	} else {
		dst.IntuneProfileResponse = nil
	}

	// try to unmarshal data into JamfProfileResponse
	err = newStrictDecoder(data).Decode(&dst.JamfProfileResponse)
	if err == nil {
		jsonJamfProfileResponse, _ := json.Marshal(dst.JamfProfileResponse)
		if string(jsonJamfProfileResponse) == "{}" { // empty struct
			dst.JamfProfileResponse = nil
		} else {
            _ = validator.Validate(dst.JamfProfileResponse)
            match++
		}
	} else {
		dst.JamfProfileResponse = nil
	}

	// try to unmarshal data into ScepProfileResponse
	err = newStrictDecoder(data).Decode(&dst.ScepProfileResponse)
	if err == nil {
		jsonScepProfileResponse, _ := json.Marshal(dst.ScepProfileResponse)
		if string(jsonScepProfileResponse) == "{}" { // empty struct
			dst.ScepProfileResponse = nil
		} else {
            _ = validator.Validate(dst.ScepProfileResponse)
            match++
		}
	} else {
		dst.ScepProfileResponse = nil
	}

	// try to unmarshal data into WcceProfileResponse
	err = newStrictDecoder(data).Decode(&dst.WcceProfileResponse)
	if err == nil {
		jsonWcceProfileResponse, _ := json.Marshal(dst.WcceProfileResponse)
		if string(jsonWcceProfileResponse) == "{}" { // empty struct
			dst.WcceProfileResponse = nil
		} else {
            _ = validator.Validate(dst.WcceProfileResponse)
            match++
		}
	} else {
		dst.WcceProfileResponse = nil
	}

	// try to unmarshal data into WebRAProfileResponse
	err = newStrictDecoder(data).Decode(&dst.WebRAProfileResponse)
	if err == nil {
		jsonWebRAProfileResponse, _ := json.Marshal(dst.WebRAProfileResponse)
		if string(jsonWebRAProfileResponse) == "{}" { // empty struct
			dst.WebRAProfileResponse = nil
		} else {
            _ = validator.Validate(dst.WebRAProfileResponse)
            match++
		}
	} else {
		dst.WebRAProfileResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CertificateProfileList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CertificateProfileList200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AcmeExternalProfileResponse != nil {
		return json.Marshal(&src.AcmeExternalProfileResponse)
	}

	if src.AcmeProfileResponse != nil {
		return json.Marshal(&src.AcmeProfileResponse)
	}

	if src.CrmpProfileResponse != nil {
		return json.Marshal(&src.CrmpProfileResponse)
	}

	if src.EstProfileResponse != nil {
		return json.Marshal(&src.EstProfileResponse)
	}

	if src.IntunePKCSProfileResponse != nil {
		return json.Marshal(&src.IntunePKCSProfileResponse)
	}

	if src.IntuneProfileResponse != nil {
		return json.Marshal(&src.IntuneProfileResponse)
	}

	if src.JamfProfileResponse != nil {
		return json.Marshal(&src.JamfProfileResponse)
	}

	if src.ScepProfileResponse != nil {
		return json.Marshal(&src.ScepProfileResponse)
	}

	if src.WcceProfileResponse != nil {
		return json.Marshal(&src.WcceProfileResponse)
	}

	if src.WebRAProfileResponse != nil {
		return json.Marshal(&src.WebRAProfileResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CertificateProfileList200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AcmeExternalProfileResponse != nil {
		return obj.AcmeExternalProfileResponse
	}

	if obj.AcmeProfileResponse != nil {
		return obj.AcmeProfileResponse
	}

	if obj.CrmpProfileResponse != nil {
		return obj.CrmpProfileResponse
	}

	if obj.EstProfileResponse != nil {
		return obj.EstProfileResponse
	}

	if obj.IntunePKCSProfileResponse != nil {
		return obj.IntunePKCSProfileResponse
	}

	if obj.IntuneProfileResponse != nil {
		return obj.IntuneProfileResponse
	}

	if obj.JamfProfileResponse != nil {
		return obj.JamfProfileResponse
	}

	if obj.ScepProfileResponse != nil {
		return obj.ScepProfileResponse
	}

	if obj.WcceProfileResponse != nil {
		return obj.WcceProfileResponse
	}

	if obj.WebRAProfileResponse != nil {
		return obj.WebRAProfileResponse
	}

	// all schemas are nil
	return nil
}

type NullableCertificateProfileList200ResponseInner struct {
	value *CertificateProfileList200ResponseInner
	isSet bool
}

func (v NullableCertificateProfileList200ResponseInner) Get() *CertificateProfileList200ResponseInner {
	return v.value
}

func (v *NullableCertificateProfileList200ResponseInner) Set(val *CertificateProfileList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileList200ResponseInner(val *CertificateProfileList200ResponseInner) *NullableCertificateProfileList200ResponseInner {
	return &NullableCertificateProfileList200ResponseInner{value: val, isSet: true}
}

func (v NullableCertificateProfileList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


