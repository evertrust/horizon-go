package horizon

type EventStatus string

const (
	Success EventStatus = "success"
	Warning EventStatus = "warning"
	Failure EventStatus = "failure"
)

type DiscoveryEventCode string

const (
	DiscoveryLocalScan   DiscoveryEventCode = "LOCALSCAN"
	DiscoveryNetscan     DiscoveryEventCode = "NETSCAN"
	DiscoveryImportscan  DiscoveryEventCode = "IMPORTSCAN"
	DiscoveryLocalImport DiscoveryEventCode = "LOCALIMPORT"
	DiscoveryNetImport   DiscoveryEventCode = "NETIMPORT"
	DiscoveryEst         DiscoveryEventCode = "EST"
	DiscoveryAcme        DiscoveryEventCode = "ACME"
	DiscoveryWebRA       DiscoveryEventCode = "WEBRA"
)

// Certificate is the struct of the locally discovered certs
type DiscoveredCertificate struct {
	DiscoveryCampaign string             `json:"campaign,omitempty"`
	SessionId         string             `json:"sessionId,omitempty"`
	Certificate       string             `json:"certificate,omitempty"`
	ContactEmail      string             `json:"contactEmail,omitempty"`
	Code              DiscoveryEventCode `json:"code,omitempty"`
	DiscoveryData     *DiscoveryData     `json:"hostDiscoveryData,omitempty"`
	Metadata          []Metadata         `json:"metadata,omitempty"`
	ThirdPartyData    []ThirdPartyItem   `json:"thirdPartyData,omitempty"`
	PrivateKey        string             `json:"privateKey,omitempty"`
}

type DiscoveredCertificateParams struct {
	Certificate    string
	Code           DiscoveryEventCode
	DiscoveryData  *DiscoveryData
	Metadata       []Metadata
	PrivateKey     string
	ContactEmail   string
	ThirdPartyData []ThirdPartyItem
}

type EnforcedIdentityProviders struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type AuthorizationLevel struct {
	AccessLevel               string `json:"accessLevel,omitempty"`
	EnforcedIdentityProviders []EnforcedIdentityProviders
}

type AuthorizationLevels struct {
	Search AuthorizationLevel `json:"search,omitempty"`
	Feed   AuthorizationLevel `json:"feed,omitempty"`
}

// DiscoveryCampaign is the struct of the discovery campaign as defined in Horizon
type DiscoveryCampaign struct {
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Enabled             bool                `json:"enabled,omitempty"`
	GradingPolicies     []string            `json:"gradingPolicies,omitempty"`
	AuthorizationLevels AuthorizationLevels `json:"authorizationLevels,omitempty"`
	Id                  string              `json:"id,omitempty"`
	EventOnSuccess      bool                `json:"eventOnSuccess,omitempty"`
	EventOnFailure      bool                `json:"eventOnFailure,omitempty"`
	EventOnWarning      bool                `json:"eventOnWarning,omitempty"`
	Hosts               []string            `json:"hosts,omitempty"`
	Ports               []string            `json:"ports,omitempty"`
}

// DiscoverySession is the struct that holds info about a sessions
type DiscoverySession struct {
	Campaign       string   `json:"campaign,omitempty"`
	Id             string   `json:"id,omitempty"`
	EventOnSuccess bool     `json:"eventOnSuccess,omitempty"`
	EventOnFailure bool     `json:"eventOnFailure,omitempty"`
	EventOnWarning bool     `json:"eventOnWarning,omitempty"`
	Hosts          []string `json:"hosts,omitempty"`
	Ports          []string `json:"ports,omitempty"`
}

// DiscoveryEvent is the struct for the discovery event as defined in Horizon
type DiscoveryEvent struct {
	Id            string             `json:"_id,omitempty"`
	Code          DiscoveryEventCode `json:"code,omitempty"`
	Campaign      string             `json:"campaign,omitempty"`
	SessionId     string             `json:"sessionId,omitempty"`
	Status        EventStatus        `json:"status,omitempty"`
	ErrorCode     string             `json:"errorCode,omitempty"`
	ErrorMessage  string             `json:"errorMessage,omitempty"`
	Timestamp     int                `json:"timestamp,omitempty"`
	ClientVersion string             `json:"clientVersion,omitempty"`
	ClientIp      string             `json:"clientIp,omitempty"`
	ClientId      string             `json:"clientId,omitempty"`
	ActorId       string             `json:"actorId,omitempty"`
	CertificateId string             `json:"certificateId,omitempty"`
	Hostname      string             `json:"hostname,omitempty"`
	Ip            string             `json:"ip,omitempty"`
	Port          int                `json:"port,omitempty"`
	Source        string             `json:"source,omitempty"`
}

// DiscoveryEventParams is the struct to build a discoveryEvent
type DiscoveryEventParams struct {
	Code         DiscoveryEventCode
	Status       EventStatus
	ErrorCode    string
	ErrorMessage string
	Hostname     string
	Ip           string
	Port         int
	Source       string
}

// DiscoveryEventSearchQuery is the struct to query discovery events
type DiscoveryEventSearchQuery struct {
	Query     string       `json:"query,omitempty"`
	SortedBy  []SortFields `json:"sortedBy,omitempty"`
	PageIndex int          `json:"pageIndex,omitempty"`
	PageSize  int          `json:"pageSize,omitempty"`
	WithCount bool         `json:"withCount,omitempty"`
}
