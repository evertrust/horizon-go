package discovery

// HrzDiscoveredCert is the struct of the locally discovered certs
type HrzDiscoveredCert struct {
	DiscoveryCampaign string                     `json:"campaign,omitempty"`
	SessionId         string                     `json:"sessionId,omitempty"`
	Certificate       string                     `json:"certificate,omitempty"`
	Code              string                     `json:"code,omitempty"`
	DiscoveryInfos    HrzDiscoveredCertsMetadata `json:"hostDiscoveryData,omitempty"`
	Metadata          []HrzCertificateMetadata   `json:"metadata,omitempty"`
	PrivateKey        string                     `json:"privateKey,omitempty"`
}

// HrzCertificateMetadata is the struct for the certificate metadata
type HrzCertificateMetadata struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// HrzDiscoveredCertsMetadata is the struct of the locally discovered certs metadata
type HrzDiscoveredCertsMetadata struct {
	CertificateLocation      []string      `json:"paths,omitempty"`
	CertificateUsageLocation []string      `json:"usages,omitempty"`
	OS                       []string      `json:"operatingSystems,omitempty"`
	Hostname                 []string      `json:"hostnames,omitempty"`
	IP                       string        `json:"ip,omitempty"`
	Source                   []string      `json:"sources,omitempty"`
	TLSPorts                 []HrzTLSPorts `json:"tlsPorts,omitempty"`
}

// HrzTLSPorts is the struct that contains the ports
type HrzTLSPorts struct {
	Port    int    `json:"port,omitempty"`
	Version string `json:"version,omitempty"`
}

// discovery campaign json:
/*
{
	"name":"mycampaign",
	"description":null,
	"enabled":true,
	"gradingPolicies":[],
	"hosts":[],
	"ports":[],
	"authorizationLevels":{
		"search":{
			"accessLevel":"authorized",
			"enforcedIdentityProviders":[{"name":"local","type":"Local"}]
		},
		"feed":{
			"accessLevel":"authorized",
			"enforcedIdentityProviders":[{"name":"local","type":"Local"}]
		}
	},
	"eventOnSuccess":false,
	"eventOnFailure":false,
	"eventOnWarning":false,
	"triggers":{}
}
*/

type DiscoveryCampaignEnforcedIdentityProviders struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type DiscoveryCampaignAuthorizationLevel struct {
	AccessLevel               string `json:"accessLevel,omitempty"`
	EnforcedIdentityProviders []DiscoveryCampaignEnforcedIdentityProviders
}

type DiscoveryCampaignAuthorizationLevels struct {
	Search DiscoveryCampaignAuthorizationLevel `json:"search,omitempty"`
	Feed   DiscoveryCampaignAuthorizationLevel `json:"feed,omitempty"`
}

// DiscoveryCampaign is the struct of the discovery campaign
type DiscoveryCampaign struct {
	Name                string                               `json:"name,omitempty"`
	Description         string                               `json:"description,omitempty"`
	Enabled             bool                                 `json:"enabled,omitempty"`
	GradingPolicies     []interface{}                        `json:"gradingPolicies,omitempty"`
	Hosts               []interface{}                        `json:"hosts,omitempty"`
	Ports               []interface{}                        `json:"ports,omitempty"`
	AuthorizationLevels DiscoveryCampaignAuthorizationLevels `json:"authorizationLevels,omitempty"`
	EventOnSuccess      bool                                 `json:"eventOnSuccess"`
	EventOnFailure      bool                                 `json:"eventOnFailure"`
	EventOnWarning      bool                                 `json:"eventOnWarning"`
	Triggers            struct{}                             `json:"triggers,omitempty"`
}

// HrzDiscoveryCampaign is the struct of the discovery campaign as defined in Horizon
type HrzDiscoveryCampaign struct {
	Campaign       string   `json:"campaign,omitempty"`
	Id             string   `json:"id,omitempty"`
	EventOnSuccess bool     `json:"eventOnSuccess,omitempty"`
	EventOnFailure bool     `json:"eventOnFailure,omitempty"`
	EventOnWarning bool     `json:"eventOnWarning,omitempty"`
	Hosts          []string `json:"hosts,omitempty"`
	Ports          []string `json:"ports,omitempty"`
}

// HrzDiscoveryEvent is the struct for the discovery event as defined in Horizon
type HrzDiscoveryEvent struct {
	Code         string `json:"code,omitempty"`
	Campaign     string `json:"campaign,omitempty"`
	SessionId    string `json:"sessionId,omitempty"`
	Status       string `json:"status,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
