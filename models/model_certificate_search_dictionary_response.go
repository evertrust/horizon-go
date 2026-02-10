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

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateSearchDictionaryResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateSearchDictionaryResponse{}

// CertificateSearchDictionaryResponse struct for CertificateSearchDictionaryResponse
type CertificateSearchDictionaryResponse struct {
	// The list of discovery campaign the principal is authorized to search on
	Campaigns []string `json:"campaigns,omitempty"`
	// The list of available grading policies on Horizon
	GradingPolicies []string `json:"gradingPolicies,omitempty"`
	// The list of labels the principal is authorized to search on
	Labels []CertificateLabelSearchDictionaryLocalizedEntry `json:"labels,omitempty"`
	// The list of available metadata in Horizon
	Metadata []string `json:"metadata"`
	// The list of Horizon modules available on this instance
	Modules []string `json:"modules,omitempty"`
	// The list of profiles the principal is authorized to search on
	Profiles []CertificateProfileSearchDictionaryLocalizedEntry `json:"profiles,omitempty"`
	// The list of available teams on this Horizon instance
	Teams                []TeamSearchDictionaryLocalizedEntry `json:"teams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateSearchDictionaryResponse CertificateSearchDictionaryResponse

// NewCertificateSearchDictionaryResponse instantiates a new CertificateSearchDictionaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateSearchDictionaryResponse(metadata []string) *CertificateSearchDictionaryResponse {
	this := CertificateSearchDictionaryResponse{}
	this.Metadata = metadata
	return &this
}

// NewCertificateSearchDictionaryResponseWithDefaults instantiates a new CertificateSearchDictionaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateSearchDictionaryResponseWithDefaults() *CertificateSearchDictionaryResponse {
	this := CertificateSearchDictionaryResponse{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetCampaigns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetCampaignsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasCampaigns() bool {
	if o != nil && !utils.IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []string and assigns it to the Campaigns field.
func (o *CertificateSearchDictionaryResponse) SetCampaigns(v []string) {
	o.Campaigns = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *CertificateSearchDictionaryResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetLabels() []CertificateLabelSearchDictionaryLocalizedEntry {
	if o == nil {
		var ret []CertificateLabelSearchDictionaryLocalizedEntry
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetLabelsOk() ([]CertificateLabelSearchDictionaryLocalizedEntry, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []CertificateLabelSearchDictionaryLocalizedEntry and assigns it to the Labels field.
func (o *CertificateSearchDictionaryResponse) SetLabels(v []CertificateLabelSearchDictionaryLocalizedEntry) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value
func (o *CertificateSearchDictionaryResponse) GetMetadata() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CertificateSearchDictionaryResponse) GetMetadataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *CertificateSearchDictionaryResponse) SetMetadata(v []string) {
	o.Metadata = v
}

// GetModules returns the Modules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetModules() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetModulesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasModules() bool {
	if o != nil && !utils.IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *CertificateSearchDictionaryResponse) SetModules(v []string) {
	o.Modules = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetProfiles() []CertificateProfileSearchDictionaryLocalizedEntry {
	if o == nil {
		var ret []CertificateProfileSearchDictionaryLocalizedEntry
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetProfilesOk() ([]CertificateProfileSearchDictionaryLocalizedEntry, bool) {
	if o == nil || utils.IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasProfiles() bool {
	if o != nil && !utils.IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []CertificateProfileSearchDictionaryLocalizedEntry and assigns it to the Profiles field.
func (o *CertificateSearchDictionaryResponse) SetProfiles(v []CertificateProfileSearchDictionaryLocalizedEntry) {
	o.Profiles = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchDictionaryResponse) GetTeams() []TeamSearchDictionaryLocalizedEntry {
	if o == nil {
		var ret []TeamSearchDictionaryLocalizedEntry
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchDictionaryResponse) GetTeamsOk() ([]TeamSearchDictionaryLocalizedEntry, bool) {
	if o == nil || utils.IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *CertificateSearchDictionaryResponse) HasTeams() bool {
	if o != nil && !utils.IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []TeamSearchDictionaryLocalizedEntry and assigns it to the Teams field.
func (o *CertificateSearchDictionaryResponse) SetTeams(v []TeamSearchDictionaryLocalizedEntry) {
	o.Teams = v
}

func (o CertificateSearchDictionaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateSearchDictionaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Campaigns != nil {
		toSerialize["campaigns"] = o.Campaigns
	}
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["metadata"] = o.Metadata
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateSearchDictionaryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCertificateSearchDictionaryResponse := _CertificateSearchDictionaryResponse{}

	err = json.Unmarshal(data, &varCertificateSearchDictionaryResponse)

	if err != nil {
		return err
	}

	*o = CertificateSearchDictionaryResponse(varCertificateSearchDictionaryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaigns")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "modules")
		delete(additionalProperties, "profiles")
		delete(additionalProperties, "teams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateSearchDictionaryResponse struct {
	value *CertificateSearchDictionaryResponse
	isSet bool
}

func (v NullableCertificateSearchDictionaryResponse) Get() *CertificateSearchDictionaryResponse {
	return v.value
}

func (v *NullableCertificateSearchDictionaryResponse) Set(val *CertificateSearchDictionaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateSearchDictionaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateSearchDictionaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateSearchDictionaryResponse(val *CertificateSearchDictionaryResponse) *NullableCertificateSearchDictionaryResponse {
	return &NullableCertificateSearchDictionaryResponse{value: val, isSet: true}
}

func (v NullableCertificateSearchDictionaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateSearchDictionaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
