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

// DcvProvisionerList200ResponseInner - struct for DcvProvisionerList200ResponseInner
type DcvProvisionerList200ResponseInner struct {
	AzurednsDCVProvisionerConfigResponse   *AzurednsDCVProvisionerConfigResponse
	CloudflareDCVProvisionerConfigResponse *CloudflareDCVProvisionerConfigResponse
	EfficientIpProvisionerConfigResponse   *EfficientIpProvisionerConfigResponse
	PowerDNSProvisionerConfigResponse      *PowerDNSProvisionerConfigResponse
	Route53ProvisionerConfigResponse       *Route53ProvisionerConfigResponse
}

// AzurednsDCVProvisionerConfigResponseAsDcvProvisionerList200ResponseInner is a convenience function that returns AzurednsDCVProvisionerConfigResponse wrapped in DcvProvisionerList200ResponseInner
func AzurednsDCVProvisionerConfigResponseAsDcvProvisionerList200ResponseInner(v *AzurednsDCVProvisionerConfigResponse) DcvProvisionerList200ResponseInner {
	return DcvProvisionerList200ResponseInner{
		AzurednsDCVProvisionerConfigResponse: v,
	}
}

// CloudflareDCVProvisionerConfigResponseAsDcvProvisionerList200ResponseInner is a convenience function that returns CloudflareDCVProvisionerConfigResponse wrapped in DcvProvisionerList200ResponseInner
func CloudflareDCVProvisionerConfigResponseAsDcvProvisionerList200ResponseInner(v *CloudflareDCVProvisionerConfigResponse) DcvProvisionerList200ResponseInner {
	return DcvProvisionerList200ResponseInner{
		CloudflareDCVProvisionerConfigResponse: v,
	}
}

// EfficientIpProvisionerConfigResponseAsDcvProvisionerList200ResponseInner is a convenience function that returns EfficientIpProvisionerConfigResponse wrapped in DcvProvisionerList200ResponseInner
func EfficientIpProvisionerConfigResponseAsDcvProvisionerList200ResponseInner(v *EfficientIpProvisionerConfigResponse) DcvProvisionerList200ResponseInner {
	return DcvProvisionerList200ResponseInner{
		EfficientIpProvisionerConfigResponse: v,
	}
}

// PowerDNSProvisionerConfigResponseAsDcvProvisionerList200ResponseInner is a convenience function that returns PowerDNSProvisionerConfigResponse wrapped in DcvProvisionerList200ResponseInner
func PowerDNSProvisionerConfigResponseAsDcvProvisionerList200ResponseInner(v *PowerDNSProvisionerConfigResponse) DcvProvisionerList200ResponseInner {
	return DcvProvisionerList200ResponseInner{
		PowerDNSProvisionerConfigResponse: v,
	}
}

// Route53ProvisionerConfigResponseAsDcvProvisionerList200ResponseInner is a convenience function that returns Route53ProvisionerConfigResponse wrapped in DcvProvisionerList200ResponseInner
func Route53ProvisionerConfigResponseAsDcvProvisionerList200ResponseInner(v *Route53ProvisionerConfigResponse) DcvProvisionerList200ResponseInner {
	return DcvProvisionerList200ResponseInner{
		Route53ProvisionerConfigResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DcvProvisionerList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AzurednsDCVProvisionerConfigResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.AzurednsDCVProvisionerConfigResponse)
	if err == nil {
		jsonAzurednsDCVProvisionerConfigResponse, _ := json.Marshal(dst.AzurednsDCVProvisionerConfigResponse)
		if string(jsonAzurednsDCVProvisionerConfigResponse) == "{}" { // empty struct
			dst.AzurednsDCVProvisionerConfigResponse = nil
		} else {
			_ = validator.Validate(dst.AzurednsDCVProvisionerConfigResponse)
			match++
		}
	} else {
		dst.AzurednsDCVProvisionerConfigResponse = nil
	}

	// try to unmarshal data into CloudflareDCVProvisionerConfigResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.CloudflareDCVProvisionerConfigResponse)
	if err == nil {
		jsonCloudflareDCVProvisionerConfigResponse, _ := json.Marshal(dst.CloudflareDCVProvisionerConfigResponse)
		if string(jsonCloudflareDCVProvisionerConfigResponse) == "{}" { // empty struct
			dst.CloudflareDCVProvisionerConfigResponse = nil
		} else {
			_ = validator.Validate(dst.CloudflareDCVProvisionerConfigResponse)
			match++
		}
	} else {
		dst.CloudflareDCVProvisionerConfigResponse = nil
	}

	// try to unmarshal data into EfficientIpProvisionerConfigResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.EfficientIpProvisionerConfigResponse)
	if err == nil {
		jsonEfficientIpProvisionerConfigResponse, _ := json.Marshal(dst.EfficientIpProvisionerConfigResponse)
		if string(jsonEfficientIpProvisionerConfigResponse) == "{}" { // empty struct
			dst.EfficientIpProvisionerConfigResponse = nil
		} else {
			_ = validator.Validate(dst.EfficientIpProvisionerConfigResponse)
			match++
		}
	} else {
		dst.EfficientIpProvisionerConfigResponse = nil
	}

	// try to unmarshal data into PowerDNSProvisionerConfigResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.PowerDNSProvisionerConfigResponse)
	if err == nil {
		jsonPowerDNSProvisionerConfigResponse, _ := json.Marshal(dst.PowerDNSProvisionerConfigResponse)
		if string(jsonPowerDNSProvisionerConfigResponse) == "{}" { // empty struct
			dst.PowerDNSProvisionerConfigResponse = nil
		} else {
			_ = validator.Validate(dst.PowerDNSProvisionerConfigResponse)
			match++
		}
	} else {
		dst.PowerDNSProvisionerConfigResponse = nil
	}

	// try to unmarshal data into Route53ProvisionerConfigResponse
	err = utils.NewStrictDecoder(data).Decode(&dst.Route53ProvisionerConfigResponse)
	if err == nil {
		jsonRoute53ProvisionerConfigResponse, _ := json.Marshal(dst.Route53ProvisionerConfigResponse)
		if string(jsonRoute53ProvisionerConfigResponse) == "{}" { // empty struct
			dst.Route53ProvisionerConfigResponse = nil
		} else {
			_ = validator.Validate(dst.Route53ProvisionerConfigResponse)
			match++
		}
	} else {
		dst.Route53ProvisionerConfigResponse = nil
	}

	if match >= 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DcvProvisionerList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DcvProvisionerList200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AzurednsDCVProvisionerConfigResponse != nil {
		return json.Marshal(&src.AzurednsDCVProvisionerConfigResponse)
	}

	if src.CloudflareDCVProvisionerConfigResponse != nil {
		return json.Marshal(&src.CloudflareDCVProvisionerConfigResponse)
	}

	if src.EfficientIpProvisionerConfigResponse != nil {
		return json.Marshal(&src.EfficientIpProvisionerConfigResponse)
	}

	if src.PowerDNSProvisionerConfigResponse != nil {
		return json.Marshal(&src.PowerDNSProvisionerConfigResponse)
	}

	if src.Route53ProvisionerConfigResponse != nil {
		return json.Marshal(&src.Route53ProvisionerConfigResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DcvProvisionerList200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AzurednsDCVProvisionerConfigResponse != nil {
		return obj.AzurednsDCVProvisionerConfigResponse
	}

	if obj.CloudflareDCVProvisionerConfigResponse != nil {
		return obj.CloudflareDCVProvisionerConfigResponse
	}

	if obj.EfficientIpProvisionerConfigResponse != nil {
		return obj.EfficientIpProvisionerConfigResponse
	}

	if obj.PowerDNSProvisionerConfigResponse != nil {
		return obj.PowerDNSProvisionerConfigResponse
	}

	if obj.Route53ProvisionerConfigResponse != nil {
		return obj.Route53ProvisionerConfigResponse
	}

	// all schemas are nil
	return nil
}

type NullableDcvProvisionerList200ResponseInner struct {
	value *DcvProvisionerList200ResponseInner
	isSet bool
}

func (v NullableDcvProvisionerList200ResponseInner) Get() *DcvProvisionerList200ResponseInner {
	return v.value
}

func (v *NullableDcvProvisionerList200ResponseInner) Set(val *DcvProvisionerList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDcvProvisionerList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDcvProvisionerList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcvProvisionerList200ResponseInner(val *DcvProvisionerList200ResponseInner) *NullableDcvProvisionerList200ResponseInner {
	return &NullableDcvProvisionerList200ResponseInner{value: val, isSet: true}
}

func (v NullableDcvProvisionerList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcvProvisionerList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
