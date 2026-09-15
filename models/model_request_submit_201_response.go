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
)

// RequestSubmit201Response - struct for RequestSubmit201Response
type RequestSubmit201Response struct {
	EstEnrollRequestOnSubmitResponse    *EstEnrollRequestOnSubmitResponse
	ScepEnrollRequestOnSubmitResponse   *ScepEnrollRequestOnSubmitResponse
	WebRAEnrollRequestOnSubmitResponse  *WebRAEnrollRequestOnSubmitResponse
	WebRAImportRequestOnSubmitResponse  *WebRAImportRequestOnSubmitResponse
	WebRAMigrateRequestOnSubmitResponse *WebRAMigrateRequestOnSubmitResponse
	WebRARecoverRequestOnSubmitResponse *WebRARecoverRequestOnSubmitResponse
	WebRARenewRequestOnSubmitResponse   *WebRARenewRequestOnSubmitResponse
	WebRARevokeRequestOnSubmitResponse  *WebRARevokeRequestOnSubmitResponse
	WebRAUpdateRequestOnSubmitResponse  *WebRAUpdateRequestOnSubmitResponse
}

// EstEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns EstEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func EstEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *EstEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		EstEnrollRequestOnSubmitResponse: v,
	}
}

// ScepEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns ScepEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func ScepEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *ScepEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		ScepEnrollRequestOnSubmitResponse: v,
	}
}

// WebRAEnrollRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAEnrollRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAEnrollRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAEnrollRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAEnrollRequestOnSubmitResponse: v,
	}
}

// WebRAImportRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAImportRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAImportRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAImportRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAImportRequestOnSubmitResponse: v,
	}
}

// WebRAMigrateRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAMigrateRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAMigrateRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAMigrateRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAMigrateRequestOnSubmitResponse: v,
	}
}

// WebRARecoverRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARecoverRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARecoverRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARecoverRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARecoverRequestOnSubmitResponse: v,
	}
}

// WebRARenewRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARenewRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARenewRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARenewRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARenewRequestOnSubmitResponse: v,
	}
}

// WebRARevokeRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRARevokeRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRARevokeRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRARevokeRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRARevokeRequestOnSubmitResponse: v,
	}
}

// WebRAUpdateRequestOnSubmitResponseAsRequestSubmit201Response is a convenience function that returns WebRAUpdateRequestOnSubmitResponse wrapped in RequestSubmit201Response
func WebRAUpdateRequestOnSubmitResponseAsRequestSubmit201Response(v *WebRAUpdateRequestOnSubmitResponse) RequestSubmit201Response {
	return RequestSubmit201Response{
		WebRAUpdateRequestOnSubmitResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSubmit201Response) UnmarshalJSON(data []byte) error {
	var err error
	// use the discriminator to select exactly one variant
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=est;workflow=enroll") {
		err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.EstEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as EstEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=scep;workflow=enroll") {
		err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.ScepEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as ScepEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=enroll") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=import") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAImportRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAImportRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=migrate") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAMigrateRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAMigrateRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=recover") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARecoverRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARecoverRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=renew") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARenewRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARenewRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=revoke") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARevokeRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARevokeRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "module=webra;workflow=update") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAUpdateRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAUpdateRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "EstEnrollRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.EstEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.EstEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as EstEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "ScepEnrollRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.ScepEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.ScepEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as ScepEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRAEnrollRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAEnrollRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAEnrollRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAEnrollRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRAImportRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAImportRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAImportRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAImportRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRAMigrateRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAMigrateRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAMigrateRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAMigrateRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRARecoverRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARecoverRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARecoverRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARecoverRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRARenewRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARenewRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARenewRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARenewRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRARevokeRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRARevokeRequestOnSubmitResponse)
		if err != nil {
			dst.WebRARevokeRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRARevokeRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	if utils.MatchOneOfDiscriminator(jsonDict, "workflow", "WebRAUpdateRequestOnSubmitResponse") {
		err = utils.NewStrictDecoder(data).Decode(&dst.WebRAUpdateRequestOnSubmitResponse)
		if err != nil {
			dst.WebRAUpdateRequestOnSubmitResponse = nil
			return fmt.Errorf("failed to unmarshal RequestSubmit201Response as WebRAUpdateRequestOnSubmitResponse: %s", err.Error())
		}
		return nil
	}

	return fmt.Errorf("data failed to match schemas in oneOf(RequestSubmit201Response): no variant matches discriminator 'workflow' (%v)", jsonDict["workflow"])
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSubmit201Response) MarshalJSON() ([]byte, error) {
	if src.EstEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.EstEnrollRequestOnSubmitResponse)
	}

	if src.ScepEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.ScepEnrollRequestOnSubmitResponse)
	}

	if src.WebRAEnrollRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAEnrollRequestOnSubmitResponse)
	}

	if src.WebRAImportRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAImportRequestOnSubmitResponse)
	}

	if src.WebRAMigrateRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAMigrateRequestOnSubmitResponse)
	}

	if src.WebRARecoverRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARecoverRequestOnSubmitResponse)
	}

	if src.WebRARenewRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARenewRequestOnSubmitResponse)
	}

	if src.WebRARevokeRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRARevokeRequestOnSubmitResponse)
	}

	if src.WebRAUpdateRequestOnSubmitResponse != nil {
		return json.Marshal(&src.WebRAUpdateRequestOnSubmitResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSubmit201Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EstEnrollRequestOnSubmitResponse != nil {
		return obj.EstEnrollRequestOnSubmitResponse
	}

	if obj.ScepEnrollRequestOnSubmitResponse != nil {
		return obj.ScepEnrollRequestOnSubmitResponse
	}

	if obj.WebRAEnrollRequestOnSubmitResponse != nil {
		return obj.WebRAEnrollRequestOnSubmitResponse
	}

	if obj.WebRAImportRequestOnSubmitResponse != nil {
		return obj.WebRAImportRequestOnSubmitResponse
	}

	if obj.WebRAMigrateRequestOnSubmitResponse != nil {
		return obj.WebRAMigrateRequestOnSubmitResponse
	}

	if obj.WebRARecoverRequestOnSubmitResponse != nil {
		return obj.WebRARecoverRequestOnSubmitResponse
	}

	if obj.WebRARenewRequestOnSubmitResponse != nil {
		return obj.WebRARenewRequestOnSubmitResponse
	}

	if obj.WebRARevokeRequestOnSubmitResponse != nil {
		return obj.WebRARevokeRequestOnSubmitResponse
	}

	if obj.WebRAUpdateRequestOnSubmitResponse != nil {
		return obj.WebRAUpdateRequestOnSubmitResponse
	}

	// all schemas are nil
	return nil
}

type NullableRequestSubmit201Response struct {
	value *RequestSubmit201Response
	isSet bool
}

func (v NullableRequestSubmit201Response) Get() *RequestSubmit201Response {
	return v.value
}

func (v *NullableRequestSubmit201Response) Set(val *RequestSubmit201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmit201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmit201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmit201Response(val *RequestSubmit201Response) *NullableRequestSubmit201Response {
	return &NullableRequestSubmit201Response{value: val, isSet: true}
}

func (v NullableRequestSubmit201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmit201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
