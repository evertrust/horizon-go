/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the HorizonImportedItemsSummary type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HorizonImportedItemsSummary{}

// HorizonImportedItemsSummary struct for HorizonImportedItemsSummary
type HorizonImportedItemsSummary struct {
	Automations          []string `json:"automations,omitempty"`
	Cas                  []string `json:"cas,omitempty"`
	Datasources          []string `json:"datasources,omitempty"`
	DiscoveryCampaigns   []string `json:"discoveryCampaigns,omitempty"`
	Executions           []string `json:"executions,omitempty"`
	ForestMappings       []string `json:"forestMappings,omitempty"`
	Labels               []string `json:"labels,omitempty"`
	Notifications        []string `json:"notifications,omitempty"`
	PasswordPolicies     []string `json:"passwordPolicies,omitempty"`
	PkiConnectors        []string `json:"pkiConnectors,omitempty"`
	PkiQueues            []string `json:"pkiQueues,omitempty"`
	Profiles             []string `json:"profiles,omitempty"`
	Proxies              []string `json:"proxies,omitempty"`
	Reports              []string `json:"reports,omitempty"`
	Roles                []string `json:"roles,omitempty"`
	ScimProfiles         []string `json:"scimProfiles,omitempty"`
	Teams                []string `json:"teams,omitempty"`
	ThirdParties         []string `json:"thirdParties,omitempty"`
	Triggers             []string `json:"triggers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HorizonImportedItemsSummary HorizonImportedItemsSummary

// NewHorizonImportedItemsSummary instantiates a new HorizonImportedItemsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizonImportedItemsSummary() *HorizonImportedItemsSummary {
	this := HorizonImportedItemsSummary{}
	return &this
}

// NewHorizonImportedItemsSummaryWithDefaults instantiates a new HorizonImportedItemsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizonImportedItemsSummaryWithDefaults() *HorizonImportedItemsSummary {
	this := HorizonImportedItemsSummary{}
	return &this
}

// GetAutomations returns the Automations field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetAutomations() []string {
	if o == nil || utils.IsNil(o.Automations) {
		var ret []string
		return ret
	}
	return o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetAutomationsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Automations) {
		return nil, false
	}
	return o.Automations, true
}

// HasAutomations returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasAutomations() bool {
	if o != nil && !utils.IsNil(o.Automations) {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given []string and assigns it to the Automations field.
func (o *HorizonImportedItemsSummary) SetAutomations(v []string) {
	o.Automations = v
}

// GetCas returns the Cas field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetCas() []string {
	if o == nil || utils.IsNil(o.Cas) {
		var ret []string
		return ret
	}
	return o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Cas) {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasCas() bool {
	if o != nil && !utils.IsNil(o.Cas) {
		return true
	}

	return false
}

// SetCas gets a reference to the given []string and assigns it to the Cas field.
func (o *HorizonImportedItemsSummary) SetCas(v []string) {
	o.Cas = v
}

// GetDatasources returns the Datasources field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetDatasources() []string {
	if o == nil || utils.IsNil(o.Datasources) {
		var ret []string
		return ret
	}
	return o.Datasources
}

// GetDatasourcesOk returns a tuple with the Datasources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetDatasourcesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Datasources) {
		return nil, false
	}
	return o.Datasources, true
}

// HasDatasources returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasDatasources() bool {
	if o != nil && !utils.IsNil(o.Datasources) {
		return true
	}

	return false
}

// SetDatasources gets a reference to the given []string and assigns it to the Datasources field.
func (o *HorizonImportedItemsSummary) SetDatasources(v []string) {
	o.Datasources = v
}

// GetDiscoveryCampaigns returns the DiscoveryCampaigns field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetDiscoveryCampaigns() []string {
	if o == nil || utils.IsNil(o.DiscoveryCampaigns) {
		var ret []string
		return ret
	}
	return o.DiscoveryCampaigns
}

// GetDiscoveryCampaignsOk returns a tuple with the DiscoveryCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetDiscoveryCampaignsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DiscoveryCampaigns) {
		return nil, false
	}
	return o.DiscoveryCampaigns, true
}

// HasDiscoveryCampaigns returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasDiscoveryCampaigns() bool {
	if o != nil && !utils.IsNil(o.DiscoveryCampaigns) {
		return true
	}

	return false
}

// SetDiscoveryCampaigns gets a reference to the given []string and assigns it to the DiscoveryCampaigns field.
func (o *HorizonImportedItemsSummary) SetDiscoveryCampaigns(v []string) {
	o.DiscoveryCampaigns = v
}

// GetExecutions returns the Executions field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetExecutions() []string {
	if o == nil || utils.IsNil(o.Executions) {
		var ret []string
		return ret
	}
	return o.Executions
}

// GetExecutionsOk returns a tuple with the Executions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetExecutionsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Executions) {
		return nil, false
	}
	return o.Executions, true
}

// HasExecutions returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasExecutions() bool {
	if o != nil && !utils.IsNil(o.Executions) {
		return true
	}

	return false
}

// SetExecutions gets a reference to the given []string and assigns it to the Executions field.
func (o *HorizonImportedItemsSummary) SetExecutions(v []string) {
	o.Executions = v
}

// GetForestMappings returns the ForestMappings field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetForestMappings() []string {
	if o == nil || utils.IsNil(o.ForestMappings) {
		var ret []string
		return ret
	}
	return o.ForestMappings
}

// GetForestMappingsOk returns a tuple with the ForestMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetForestMappingsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ForestMappings) {
		return nil, false
	}
	return o.ForestMappings, true
}

// HasForestMappings returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasForestMappings() bool {
	if o != nil && !utils.IsNil(o.ForestMappings) {
		return true
	}

	return false
}

// SetForestMappings gets a reference to the given []string and assigns it to the ForestMappings field.
func (o *HorizonImportedItemsSummary) SetForestMappings(v []string) {
	o.ForestMappings = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetLabels() []string {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *HorizonImportedItemsSummary) SetLabels(v []string) {
	o.Labels = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetNotifications() []string {
	if o == nil || utils.IsNil(o.Notifications) {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetNotificationsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasNotifications() bool {
	if o != nil && !utils.IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *HorizonImportedItemsSummary) SetNotifications(v []string) {
	o.Notifications = v
}

// GetPasswordPolicies returns the PasswordPolicies field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetPasswordPolicies() []string {
	if o == nil || utils.IsNil(o.PasswordPolicies) {
		var ret []string
		return ret
	}
	return o.PasswordPolicies
}

// GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetPasswordPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PasswordPolicies) {
		return nil, false
	}
	return o.PasswordPolicies, true
}

// HasPasswordPolicies returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasPasswordPolicies() bool {
	if o != nil && !utils.IsNil(o.PasswordPolicies) {
		return true
	}

	return false
}

// SetPasswordPolicies gets a reference to the given []string and assigns it to the PasswordPolicies field.
func (o *HorizonImportedItemsSummary) SetPasswordPolicies(v []string) {
	o.PasswordPolicies = v
}

// GetPkiConnectors returns the PkiConnectors field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetPkiConnectors() []string {
	if o == nil || utils.IsNil(o.PkiConnectors) {
		var ret []string
		return ret
	}
	return o.PkiConnectors
}

// GetPkiConnectorsOk returns a tuple with the PkiConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetPkiConnectorsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PkiConnectors) {
		return nil, false
	}
	return o.PkiConnectors, true
}

// HasPkiConnectors returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasPkiConnectors() bool {
	if o != nil && !utils.IsNil(o.PkiConnectors) {
		return true
	}

	return false
}

// SetPkiConnectors gets a reference to the given []string and assigns it to the PkiConnectors field.
func (o *HorizonImportedItemsSummary) SetPkiConnectors(v []string) {
	o.PkiConnectors = v
}

// GetPkiQueues returns the PkiQueues field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetPkiQueues() []string {
	if o == nil || utils.IsNil(o.PkiQueues) {
		var ret []string
		return ret
	}
	return o.PkiQueues
}

// GetPkiQueuesOk returns a tuple with the PkiQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetPkiQueuesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PkiQueues) {
		return nil, false
	}
	return o.PkiQueues, true
}

// HasPkiQueues returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasPkiQueues() bool {
	if o != nil && !utils.IsNil(o.PkiQueues) {
		return true
	}

	return false
}

// SetPkiQueues gets a reference to the given []string and assigns it to the PkiQueues field.
func (o *HorizonImportedItemsSummary) SetPkiQueues(v []string) {
	o.PkiQueues = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetProfiles() []string {
	if o == nil || utils.IsNil(o.Profiles) {
		var ret []string
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetProfilesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasProfiles() bool {
	if o != nil && !utils.IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []string and assigns it to the Profiles field.
func (o *HorizonImportedItemsSummary) SetProfiles(v []string) {
	o.Profiles = v
}

// GetProxies returns the Proxies field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetProxies() []string {
	if o == nil || utils.IsNil(o.Proxies) {
		var ret []string
		return ret
	}
	return o.Proxies
}

// GetProxiesOk returns a tuple with the Proxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetProxiesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Proxies) {
		return nil, false
	}
	return o.Proxies, true
}

// HasProxies returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasProxies() bool {
	if o != nil && !utils.IsNil(o.Proxies) {
		return true
	}

	return false
}

// SetProxies gets a reference to the given []string and assigns it to the Proxies field.
func (o *HorizonImportedItemsSummary) SetProxies(v []string) {
	o.Proxies = v
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetReports() []string {
	if o == nil || utils.IsNil(o.Reports) {
		var ret []string
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetReportsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasReports() bool {
	if o != nil && !utils.IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []string and assigns it to the Reports field.
func (o *HorizonImportedItemsSummary) SetReports(v []string) {
	o.Reports = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetRoles() []string {
	if o == nil || utils.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetRolesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *HorizonImportedItemsSummary) SetRoles(v []string) {
	o.Roles = v
}

// GetScimProfiles returns the ScimProfiles field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetScimProfiles() []string {
	if o == nil || utils.IsNil(o.ScimProfiles) {
		var ret []string
		return ret
	}
	return o.ScimProfiles
}

// GetScimProfilesOk returns a tuple with the ScimProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetScimProfilesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ScimProfiles) {
		return nil, false
	}
	return o.ScimProfiles, true
}

// HasScimProfiles returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasScimProfiles() bool {
	if o != nil && !utils.IsNil(o.ScimProfiles) {
		return true
	}

	return false
}

// SetScimProfiles gets a reference to the given []string and assigns it to the ScimProfiles field.
func (o *HorizonImportedItemsSummary) SetScimProfiles(v []string) {
	o.ScimProfiles = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetTeams() []string {
	if o == nil || utils.IsNil(o.Teams) {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetTeamsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasTeams() bool {
	if o != nil && !utils.IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *HorizonImportedItemsSummary) SetTeams(v []string) {
	o.Teams = v
}

// GetThirdParties returns the ThirdParties field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetThirdParties() []string {
	if o == nil || utils.IsNil(o.ThirdParties) {
		var ret []string
		return ret
	}
	return o.ThirdParties
}

// GetThirdPartiesOk returns a tuple with the ThirdParties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetThirdPartiesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ThirdParties) {
		return nil, false
	}
	return o.ThirdParties, true
}

// HasThirdParties returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasThirdParties() bool {
	if o != nil && !utils.IsNil(o.ThirdParties) {
		return true
	}

	return false
}

// SetThirdParties gets a reference to the given []string and assigns it to the ThirdParties field.
func (o *HorizonImportedItemsSummary) SetThirdParties(v []string) {
	o.ThirdParties = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *HorizonImportedItemsSummary) GetTriggers() []string {
	if o == nil || utils.IsNil(o.Triggers) {
		var ret []string
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HorizonImportedItemsSummary) GetTriggersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *HorizonImportedItemsSummary) HasTriggers() bool {
	if o != nil && !utils.IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []string and assigns it to the Triggers field.
func (o *HorizonImportedItemsSummary) SetTriggers(v []string) {
	o.Triggers = v
}

func (o HorizonImportedItemsSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizonImportedItemsSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Automations) {
		toSerialize["automations"] = o.Automations
	}
	if !utils.IsNil(o.Cas) {
		toSerialize["cas"] = o.Cas
	}
	if !utils.IsNil(o.Datasources) {
		toSerialize["datasources"] = o.Datasources
	}
	if !utils.IsNil(o.DiscoveryCampaigns) {
		toSerialize["discoveryCampaigns"] = o.DiscoveryCampaigns
	}
	if !utils.IsNil(o.Executions) {
		toSerialize["executions"] = o.Executions
	}
	if !utils.IsNil(o.ForestMappings) {
		toSerialize["forestMappings"] = o.ForestMappings
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !utils.IsNil(o.PasswordPolicies) {
		toSerialize["passwordPolicies"] = o.PasswordPolicies
	}
	if !utils.IsNil(o.PkiConnectors) {
		toSerialize["pkiConnectors"] = o.PkiConnectors
	}
	if !utils.IsNil(o.PkiQueues) {
		toSerialize["pkiQueues"] = o.PkiQueues
	}
	if !utils.IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !utils.IsNil(o.Proxies) {
		toSerialize["proxies"] = o.Proxies
	}
	if !utils.IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	if !utils.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !utils.IsNil(o.ScimProfiles) {
		toSerialize["scimProfiles"] = o.ScimProfiles
	}
	if !utils.IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	if !utils.IsNil(o.ThirdParties) {
		toSerialize["thirdParties"] = o.ThirdParties
	}
	if !utils.IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HorizonImportedItemsSummary) UnmarshalJSON(data []byte) (err error) {
	varHorizonImportedItemsSummary := _HorizonImportedItemsSummary{}

	err = json.Unmarshal(data, &varHorizonImportedItemsSummary)

	if err != nil {
		return err
	}

	*o = HorizonImportedItemsSummary(varHorizonImportedItemsSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "automations")
		delete(additionalProperties, "cas")
		delete(additionalProperties, "datasources")
		delete(additionalProperties, "discoveryCampaigns")
		delete(additionalProperties, "executions")
		delete(additionalProperties, "forestMappings")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "passwordPolicies")
		delete(additionalProperties, "pkiConnectors")
		delete(additionalProperties, "pkiQueues")
		delete(additionalProperties, "profiles")
		delete(additionalProperties, "proxies")
		delete(additionalProperties, "reports")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "scimProfiles")
		delete(additionalProperties, "teams")
		delete(additionalProperties, "thirdParties")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHorizonImportedItemsSummary struct {
	value *HorizonImportedItemsSummary
	isSet bool
}

func (v NullableHorizonImportedItemsSummary) Get() *HorizonImportedItemsSummary {
	return v.value
}

func (v *NullableHorizonImportedItemsSummary) Set(val *HorizonImportedItemsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizonImportedItemsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizonImportedItemsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizonImportedItemsSummary(val *HorizonImportedItemsSummary) *NullableHorizonImportedItemsSummary {
	return &NullableHorizonImportedItemsSummary{value: val, isSet: true}
}

func (v NullableHorizonImportedItemsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizonImportedItemsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
