/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the HorizonExportableItems type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HorizonExportableItems{}

// HorizonExportableItems struct for HorizonExportableItems
type HorizonExportableItems struct {
	Automations          []HorizonExportableItem `json:"automations,omitempty"`
	Cas                  []HorizonExportableItem `json:"cas,omitempty"`
	Datasources          []HorizonExportableItem `json:"datasources,omitempty"`
	DiscoveryCampaigns   []HorizonExportableItem `json:"discoveryCampaigns,omitempty"`
	Executions           []HorizonExportableItem `json:"executions,omitempty"`
	ForestMappings       []HorizonExportableItem `json:"forestMappings,omitempty"`
	Labels               []HorizonExportableItem `json:"labels,omitempty"`
	Notifications        []HorizonExportableItem `json:"notifications,omitempty"`
	PasswordPolicies     []HorizonExportableItem `json:"passwordPolicies,omitempty"`
	PkiConnectors        []HorizonExportableItem `json:"pkiConnectors,omitempty"`
	PkiQueues            []HorizonExportableItem `json:"pkiQueues,omitempty"`
	Profiles             []HorizonExportableItem `json:"profiles,omitempty"`
	Proxies              []HorizonExportableItem `json:"proxies,omitempty"`
	Reports              []HorizonExportableItem `json:"reports,omitempty"`
	Roles                []HorizonExportableItem `json:"roles,omitempty"`
	ScimProfiles         []HorizonExportableItem `json:"scimProfiles,omitempty"`
	Storages             []HorizonExportableItem `json:"storages,omitempty"`
	Teams                []HorizonExportableItem `json:"teams,omitempty"`
	ThirdParties         []HorizonExportableItem `json:"thirdParties,omitempty"`
	Triggers             []HorizonExportableItem `json:"triggers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HorizonExportableItems HorizonExportableItems

// NewHorizonExportableItems instantiates a new HorizonExportableItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizonExportableItems() *HorizonExportableItems {
	this := HorizonExportableItems{}
	return &this
}

// NewHorizonExportableItemsWithDefaults instantiates a new HorizonExportableItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizonExportableItemsWithDefaults() *HorizonExportableItems {
	this := HorizonExportableItems{}
	return &this
}

// GetAutomations returns the Automations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetAutomations() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetAutomationsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Automations) {
		return nil, false
	}
	return o.Automations, true
}

// HasAutomations returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasAutomations() bool {
	if o != nil && !utils.IsNil(o.Automations) {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given []HorizonExportableItem and assigns it to the Automations field.
func (o *HorizonExportableItems) SetAutomations(v []HorizonExportableItem) {
	o.Automations = v
}

// GetCas returns the Cas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetCas() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetCasOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Cas) {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasCas() bool {
	if o != nil && !utils.IsNil(o.Cas) {
		return true
	}

	return false
}

// SetCas gets a reference to the given []HorizonExportableItem and assigns it to the Cas field.
func (o *HorizonExportableItems) SetCas(v []HorizonExportableItem) {
	o.Cas = v
}

// GetDatasources returns the Datasources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetDatasources() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Datasources
}

// GetDatasourcesOk returns a tuple with the Datasources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetDatasourcesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Datasources) {
		return nil, false
	}
	return o.Datasources, true
}

// HasDatasources returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasDatasources() bool {
	if o != nil && !utils.IsNil(o.Datasources) {
		return true
	}

	return false
}

// SetDatasources gets a reference to the given []HorizonExportableItem and assigns it to the Datasources field.
func (o *HorizonExportableItems) SetDatasources(v []HorizonExportableItem) {
	o.Datasources = v
}

// GetDiscoveryCampaigns returns the DiscoveryCampaigns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetDiscoveryCampaigns() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.DiscoveryCampaigns
}

// GetDiscoveryCampaignsOk returns a tuple with the DiscoveryCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetDiscoveryCampaignsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.DiscoveryCampaigns) {
		return nil, false
	}
	return o.DiscoveryCampaigns, true
}

// HasDiscoveryCampaigns returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasDiscoveryCampaigns() bool {
	if o != nil && !utils.IsNil(o.DiscoveryCampaigns) {
		return true
	}

	return false
}

// SetDiscoveryCampaigns gets a reference to the given []HorizonExportableItem and assigns it to the DiscoveryCampaigns field.
func (o *HorizonExportableItems) SetDiscoveryCampaigns(v []HorizonExportableItem) {
	o.DiscoveryCampaigns = v
}

// GetExecutions returns the Executions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetExecutions() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Executions
}

// GetExecutionsOk returns a tuple with the Executions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetExecutionsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Executions) {
		return nil, false
	}
	return o.Executions, true
}

// HasExecutions returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasExecutions() bool {
	if o != nil && !utils.IsNil(o.Executions) {
		return true
	}

	return false
}

// SetExecutions gets a reference to the given []HorizonExportableItem and assigns it to the Executions field.
func (o *HorizonExportableItems) SetExecutions(v []HorizonExportableItem) {
	o.Executions = v
}

// GetForestMappings returns the ForestMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetForestMappings() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.ForestMappings
}

// GetForestMappingsOk returns a tuple with the ForestMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetForestMappingsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.ForestMappings) {
		return nil, false
	}
	return o.ForestMappings, true
}

// HasForestMappings returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasForestMappings() bool {
	if o != nil && !utils.IsNil(o.ForestMappings) {
		return true
	}

	return false
}

// SetForestMappings gets a reference to the given []HorizonExportableItem and assigns it to the ForestMappings field.
func (o *HorizonExportableItems) SetForestMappings(v []HorizonExportableItem) {
	o.ForestMappings = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetLabels() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetLabelsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []HorizonExportableItem and assigns it to the Labels field.
func (o *HorizonExportableItems) SetLabels(v []HorizonExportableItem) {
	o.Labels = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetNotifications() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetNotificationsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasNotifications() bool {
	if o != nil && !utils.IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []HorizonExportableItem and assigns it to the Notifications field.
func (o *HorizonExportableItems) SetNotifications(v []HorizonExportableItem) {
	o.Notifications = v
}

// GetPasswordPolicies returns the PasswordPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetPasswordPolicies() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.PasswordPolicies
}

// GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetPasswordPoliciesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.PasswordPolicies) {
		return nil, false
	}
	return o.PasswordPolicies, true
}

// HasPasswordPolicies returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasPasswordPolicies() bool {
	if o != nil && !utils.IsNil(o.PasswordPolicies) {
		return true
	}

	return false
}

// SetPasswordPolicies gets a reference to the given []HorizonExportableItem and assigns it to the PasswordPolicies field.
func (o *HorizonExportableItems) SetPasswordPolicies(v []HorizonExportableItem) {
	o.PasswordPolicies = v
}

// GetPkiConnectors returns the PkiConnectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetPkiConnectors() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.PkiConnectors
}

// GetPkiConnectorsOk returns a tuple with the PkiConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetPkiConnectorsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.PkiConnectors) {
		return nil, false
	}
	return o.PkiConnectors, true
}

// HasPkiConnectors returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasPkiConnectors() bool {
	if o != nil && !utils.IsNil(o.PkiConnectors) {
		return true
	}

	return false
}

// SetPkiConnectors gets a reference to the given []HorizonExportableItem and assigns it to the PkiConnectors field.
func (o *HorizonExportableItems) SetPkiConnectors(v []HorizonExportableItem) {
	o.PkiConnectors = v
}

// GetPkiQueues returns the PkiQueues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetPkiQueues() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.PkiQueues
}

// GetPkiQueuesOk returns a tuple with the PkiQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetPkiQueuesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.PkiQueues) {
		return nil, false
	}
	return o.PkiQueues, true
}

// HasPkiQueues returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasPkiQueues() bool {
	if o != nil && !utils.IsNil(o.PkiQueues) {
		return true
	}

	return false
}

// SetPkiQueues gets a reference to the given []HorizonExportableItem and assigns it to the PkiQueues field.
func (o *HorizonExportableItems) SetPkiQueues(v []HorizonExportableItem) {
	o.PkiQueues = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetProfiles() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetProfilesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasProfiles() bool {
	if o != nil && !utils.IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []HorizonExportableItem and assigns it to the Profiles field.
func (o *HorizonExportableItems) SetProfiles(v []HorizonExportableItem) {
	o.Profiles = v
}

// GetProxies returns the Proxies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetProxies() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Proxies
}

// GetProxiesOk returns a tuple with the Proxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetProxiesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Proxies) {
		return nil, false
	}
	return o.Proxies, true
}

// HasProxies returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasProxies() bool {
	if o != nil && !utils.IsNil(o.Proxies) {
		return true
	}

	return false
}

// SetProxies gets a reference to the given []HorizonExportableItem and assigns it to the Proxies field.
func (o *HorizonExportableItems) SetProxies(v []HorizonExportableItem) {
	o.Proxies = v
}

// GetReports returns the Reports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetReports() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetReportsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasReports() bool {
	if o != nil && !utils.IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []HorizonExportableItem and assigns it to the Reports field.
func (o *HorizonExportableItems) SetReports(v []HorizonExportableItem) {
	o.Reports = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetRoles() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetRolesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []HorizonExportableItem and assigns it to the Roles field.
func (o *HorizonExportableItems) SetRoles(v []HorizonExportableItem) {
	o.Roles = v
}

// GetScimProfiles returns the ScimProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetScimProfiles() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.ScimProfiles
}

// GetScimProfilesOk returns a tuple with the ScimProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetScimProfilesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.ScimProfiles) {
		return nil, false
	}
	return o.ScimProfiles, true
}

// HasScimProfiles returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasScimProfiles() bool {
	if o != nil && !utils.IsNil(o.ScimProfiles) {
		return true
	}

	return false
}

// SetScimProfiles gets a reference to the given []HorizonExportableItem and assigns it to the ScimProfiles field.
func (o *HorizonExportableItems) SetScimProfiles(v []HorizonExportableItem) {
	o.ScimProfiles = v
}

// GetStorages returns the Storages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetStorages() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetStoragesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Storages) {
		return nil, false
	}
	return o.Storages, true
}

// HasStorages returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasStorages() bool {
	if o != nil && !utils.IsNil(o.Storages) {
		return true
	}

	return false
}

// SetStorages gets a reference to the given []HorizonExportableItem and assigns it to the Storages field.
func (o *HorizonExportableItems) SetStorages(v []HorizonExportableItem) {
	o.Storages = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetTeams() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetTeamsOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasTeams() bool {
	if o != nil && !utils.IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []HorizonExportableItem and assigns it to the Teams field.
func (o *HorizonExportableItems) SetTeams(v []HorizonExportableItem) {
	o.Teams = v
}

// GetThirdParties returns the ThirdParties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetThirdParties() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.ThirdParties
}

// GetThirdPartiesOk returns a tuple with the ThirdParties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetThirdPartiesOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.ThirdParties) {
		return nil, false
	}
	return o.ThirdParties, true
}

// HasThirdParties returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasThirdParties() bool {
	if o != nil && !utils.IsNil(o.ThirdParties) {
		return true
	}

	return false
}

// SetThirdParties gets a reference to the given []HorizonExportableItem and assigns it to the ThirdParties field.
func (o *HorizonExportableItems) SetThirdParties(v []HorizonExportableItem) {
	o.ThirdParties = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HorizonExportableItems) GetTriggers() []HorizonExportableItem {
	if o == nil {
		var ret []HorizonExportableItem
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HorizonExportableItems) GetTriggersOk() ([]HorizonExportableItem, bool) {
	if o == nil || utils.IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *HorizonExportableItems) HasTriggers() bool {
	if o != nil && !utils.IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []HorizonExportableItem and assigns it to the Triggers field.
func (o *HorizonExportableItems) SetTriggers(v []HorizonExportableItem) {
	o.Triggers = v
}

func (o HorizonExportableItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizonExportableItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Automations != nil {
		toSerialize["automations"] = o.Automations
	}
	if o.Cas != nil {
		toSerialize["cas"] = o.Cas
	}
	if o.Datasources != nil {
		toSerialize["datasources"] = o.Datasources
	}
	if o.DiscoveryCampaigns != nil {
		toSerialize["discoveryCampaigns"] = o.DiscoveryCampaigns
	}
	if o.Executions != nil {
		toSerialize["executions"] = o.Executions
	}
	if o.ForestMappings != nil {
		toSerialize["forestMappings"] = o.ForestMappings
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.PasswordPolicies != nil {
		toSerialize["passwordPolicies"] = o.PasswordPolicies
	}
	if o.PkiConnectors != nil {
		toSerialize["pkiConnectors"] = o.PkiConnectors
	}
	if o.PkiQueues != nil {
		toSerialize["pkiQueues"] = o.PkiQueues
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.Proxies != nil {
		toSerialize["proxies"] = o.Proxies
	}
	if o.Reports != nil {
		toSerialize["reports"] = o.Reports
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.ScimProfiles != nil {
		toSerialize["scimProfiles"] = o.ScimProfiles
	}
	if o.Storages != nil {
		toSerialize["storages"] = o.Storages
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.ThirdParties != nil {
		toSerialize["thirdParties"] = o.ThirdParties
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HorizonExportableItems) UnmarshalJSON(data []byte) (err error) {
	varHorizonExportableItems := _HorizonExportableItems{}

	err = json.Unmarshal(data, &varHorizonExportableItems)

	if err != nil {
		return err
	}

	*o = HorizonExportableItems(varHorizonExportableItems)

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
		delete(additionalProperties, "storages")
		delete(additionalProperties, "teams")
		delete(additionalProperties, "thirdParties")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHorizonExportableItems struct {
	value *HorizonExportableItems
	isSet bool
}

func (v NullableHorizonExportableItems) Get() *HorizonExportableItems {
	return v.value
}

func (v *NullableHorizonExportableItems) Set(val *HorizonExportableItems) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizonExportableItems) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizonExportableItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizonExportableItems(val *HorizonExportableItems) *NullableHorizonExportableItems {
	return &NullableHorizonExportableItems{value: val, isSet: true}
}

func (v NullableHorizonExportableItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizonExportableItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
