package horizon

type CompliancePolicy struct {
	AuthorizedSigningAlgorithms []string `json:"authorizedSigningAlgorithms,omitempty"`
	AuthorizedCas               []string `json:"authorizedCas,omitempty"`
}

type Policy struct {
	Name             string            `json:"name"`
	Profile          string            `json:"profile"`
	TrustChains      []string          `json:"trustChains,omitempty"`
	ExecutionPolicy  string            `json:"executionPolicy,omitempty"`
	CompliancePolicy *CompliancePolicy `json:"compliancePolicy,omitempty"`
}

type InitParameter struct {
	Module  string `json:"module"`
	Profile string `json:"profile"`
}

func (p *InitParameter) GetModule() Module {
	return Module(p.Module)
}

func (p *InitParameter) GetProfile() string {
	return p.Profile
}

type InitParameters interface {
	GetModule() Module
	GetProfile() string
}

type EstInitParameters struct {
	Profile           string `json:"profile"`
	KeyType           string `json:"keyType"`
	AuthorizationMode string `json:"authorizationMode"`
	EnrollmentMode    string `json:"enrollmentMode"`
	CsrInfoIgnored    bool   `json:"csrInfoIgnored"`
}

func (p *EstInitParameters) GetModule() Module {
	return Est
}

func (p *EstInitParameters) GetProfile() string {
	return p.Profile
}

type AcmeExternalInitParameters struct {
	Profile                    string   `json:"profile"`
	KeyType                    string   `json:"keyType"`
	AcmeUrl                    string   `json:"acmeUrl"`
	AllowedAuthorizationMethod []string `json:"allowedAuthorizationMethod"`
	RequireEAB                 bool     `json:"requireEAB"`
}

func (p *AcmeExternalInitParameters) GetModule() Module {
	return AcmeExternal
}

func (p *AcmeExternalInitParameters) GetProfile() string {
	return p.Profile
}

type AcmeInitParameters struct {
	Profile       string `json:"profile"`
	KeyType       string `json:"keyType"`
	TlsAlpn01Port int    `json:"tlsAlpn01Port"`
	Http01Port    int    `json:"http01Port"`
}

func (p *AcmeInitParameters) GetModule() Module {
	return Acme
}

func (p *AcmeInitParameters) GetProfile() string {
	return p.Profile
}

type ScepInitParameters struct {
	Profile           string `json:"profile"`
	KeyType           string `json:"keyType"`
	CsrInfoIgnored    bool   `json:"csrInfoIgnored"`
	AuthorizationMode string `json:"authorizationMode"`
}

func (p *ScepInitParameters) GetModule() Module {
	return Scep
}

func (p *ScepInitParameters) GetProfile() string {
	return p.Profile
}

type Report struct {
	IsRunnable bool `json:"runnable"`
	// If renewable, renew, else enroll
	// Is a pointer for compatibility with 2.4
	IsRenewable *bool `json:"renewable"`
}
