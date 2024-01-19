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
	Module string `json:"module"`
}

func (p *InitParameter) GetModule() Module {
	return Module(p.Module)
}

type InitParameters interface {
	GetModule() Module
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

type AcmeInitParameters struct {
	Profile       string `json:"profile"`
	KeyType       string `json:"keyType"`
	TlsAlpn01Port string `json:"tlsAlpn01Port"`
	Http01Port    string `json:"http01Port"`
}

func (p *AcmeInitParameters) GetModule() Module {
	return Acme
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

type Report struct {
	IsRunnable bool `json:"runnable"`
	// If renewable, renew, else enroll
	// Is a pointer for compatibility with 2.4
	IsRenewable *bool `json:"renewable"`
}
