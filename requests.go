package horizon

import (
	"fmt"
)

// Requests
type Workflow string

const (
	Enroll  Workflow = "enroll"
	Revoke  Workflow = "revoke"
	Update  Workflow = "update"
	Migrate Workflow = "migrate"
	Recover Workflow = "recover"
	Renew   Workflow = "renew"
	Import  Workflow = "import"
)

type Status string

const (
	Denied     Status = "denied"
	Pending    Status = "pending"
	Approved   Status = "approved"
	Canceled   Status = "canceled"
	Completed  Status = "completed"
	Processing Status = "processing"
)

func invalidModuleError(found, expected Module) error {
	return fmt.Errorf("invalid module (found '%s', expected '%s')", found, expected)
}

func invalidWorkflowError(found, expected Workflow) error {
	return fmt.Errorf("invalid workflow (found '%s', expected '%s')", found, expected)
}

type IndexedDNElement struct {
	Element   string `json:"element"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
	Mandatory bool   `json:"mandatory,omitempty"`
	Editable  bool   `json:"editable,omitempty"`
}

type ListSANElement struct {
	Type     string   `json:"type,omitempty"`
	Value    []string `json:"value,omitempty"`
	Editable bool     `json:"editable,omitempty"`
	Min      int      `json:"min,omitempty"`
	Max      int      `json:"max,omitempty"`
}

type ExtensionElement struct {
	Type      string `json:"type,omitempty"`
	Editable  bool   `json:"editable,omitempty"`
	Mandatory bool   `json:"mandatory,omitempty"`
	Value     string `json:"value,omitempty"`
}

func NewLabelElement(label string, value *string) *LabelElement {
	if value == nil {
		return nil
	}
	if *value == "" {
		return &LabelElement{Label: label, Value: Delete}
	}
	return &LabelElement{Label: label, Value: &String{*value}}
}

type LabelElement struct {
	Label     string  `json:"label,omitempty"`
	Mandatory bool    `json:"mandatory,omitempty"`
	Editable  bool    `json:"editable,omitempty"`
	Value     *String `json:"value"`
}

type MetadataType string

const (
	// Horizon metadata
	MetadataPkiConnector          MetadataType = "pki_connector"
	MetadataScepTransId           MetadataType = "scep_transid"
	MetadataRenewedCertificateId  MetadataType = "renewed_certificate_id"
	MetadataPreviousCertificateId MetadataType = "previous_certificate_id"
	MetadataAutomationPolicy      MetadataType = "automation_policy"

	// Third party PKI metadata
	MetadataCertEuropeId      MetadataType = "certeurope_id"
	MetadataDigiCertId        MetadataType = "digicert_id"
	MetadataDigicertOrderId   MetadataType = "digicert_order_id"
	MetadataEntrustId         MetadataType = "entrust_id"
	MetadataFCMSId            MetadataType = "fcms_id"
	MetadataGlobalSignAtlasId MetadataType = "gsatlas_id"
	MetadataGlobalSignId      MetadataType = "gs_order_id"
	MetadataMetaPKIId         MetadataType = "metapki_id"
	MetadataIDCAId            MetadataType = "eviden_idca_id"
)

func GetMetadataType(metadata string) (MetadataType, error) {
	switch MetadataType(metadata) {
	case MetadataPkiConnector:
		return MetadataPkiConnector, nil
	case MetadataScepTransId:
		return MetadataScepTransId, nil
	case MetadataRenewedCertificateId:
		return MetadataRenewedCertificateId, nil
	case MetadataPreviousCertificateId:
		return MetadataPreviousCertificateId, nil
	case MetadataAutomationPolicy:
		return MetadataAutomationPolicy, nil
	case MetadataCertEuropeId:
		return MetadataCertEuropeId, nil
	case MetadataDigiCertId:
		return MetadataDigiCertId, nil
	case MetadataDigicertOrderId:
		return MetadataDigicertOrderId, nil
	case MetadataEntrustId:
		return MetadataEntrustId, nil
	case MetadataFCMSId:
		return MetadataFCMSId, nil
	case MetadataGlobalSignAtlasId:
		return MetadataGlobalSignAtlasId, nil
	case MetadataGlobalSignId:
		return MetadataGlobalSignId, nil
	case MetadataMetaPKIId:
		return MetadataMetaPKIId, nil
	case MetadataIDCAId:
		return MetadataIDCAId, nil
	}
	return "", fmt.Errorf("invalid metadata type: %s", metadata)
}

// NewMetadataElement creates a new metadata element for Horizon:
// nil value means no change should be made
// empty value means the value was deleted
// any other value is setting the value
// note: this function does not check that metadata is a valid enum value. For this, use GetMetadataType beforehand
func NewMetadataElement(metadata string, value *string) *MetadataElement {
	if value == nil {
		return nil
	}
	if *value == "" {
		return &MetadataElement{Metadata: MetadataType(metadata), Value: Delete}
	}
	return &MetadataElement{Metadata: MetadataType(metadata), Value: &String{*value}}
}

type MetadataElement struct {
	Metadata MetadataType `json:"metadata,omitempty"`
	Editable bool         `json:"editable,omitempty"`
	Value    *String      `json:"value"`
}

func NewOwnerElement(value *string) *OwnerElement {
	if value == nil {
		return nil
	}
	if *value == "" {
		return &OwnerElement{Value: Delete}
	}
	return &OwnerElement{Value: &String{*value}}
}

type OwnerElement struct {
	Value     *String `json:"value"`
	Mandatory bool    `json:"mandatory"`
	Editable  bool    `json:"editable"`
}

func NewTeamElement(value *string) *TeamElement {
	if value == nil {
		return nil
	}
	if *value == "" {
		return &TeamElement{Value: Delete}
	}
	return &TeamElement{Value: &String{*value}}
}

type TeamElement struct {
	Value      *String  `json:"value"`
	Mandatory  bool     `json:"mandatory"`
	Authorized []string `json:"authorized,omitempty"`
	Editable   bool     `json:"editable"`
}

func NewContactEmailElement(value *string) *ContactEmailElement {
	if value == nil {
		return nil
	}
	if *value == "" {
		return &ContactEmailElement{Value: Delete}
	}
	return &ContactEmailElement{Value: &String{*value}}
}

type ContactEmailElement struct {
	Editable  bool    `json:"editable,omitempty"`
	Mandatory bool    `json:"mandatory,omitempty"`
	Value     *String `json:"value"`
}

type Secret struct {
	Value string `json:"value,omitempty"`
}

type Capabilities struct {
	Centralized              bool     `json:"centralized,omitempty"`
	Decentralized            bool     `json:"decentralized,omitempty"`
	PreferredEnrollmentMode  string   `json:"preferredEnrollmentMode,omitempty"`
	DefaultKeyType           string   `json:"defaultKeyType,omitempty"`
	AuthorizedKeyTypes       []string `json:"authorizedKeyTypes,omitempty"`
	Escrow                   bool     `json:"escrow,omitempty"`
	ShowP12PasswordOnEnroll  bool     `json:"showP12PasswordOnEnroll,omitempty"`
	ShowP12OnEnroll          bool     `json:"showP12OnEnroll,omitempty"`
	ShowP12PasswordOnRecover bool     `json:"showP12PasswordOnRecover,omitempty"`
	ShowP12OnRecover         bool     `json:"showP12OnRecover,omitempty"`
	P12PasswordMode          string   `json:"p12PasswordMode,omitempty"`
}

type Request interface {
	EnsureType() error
}

type WebRAEnrollTemplateParams struct {
	CertificatePEM string
	CertificateId  string
	Profile        string
	Csr            string
}

// Capabilities is a readonly field
type WebRAEnrollTemplate struct {
	KeyType      string               `json:"keyType,omitempty"`
	Csr          string               `json:"csr,omitempty"`
	Subject      []IndexedDNElement   `json:"subject,omitempty"`
	Sans         []ListSANElement     `json:"sans,omitempty"`
	Extensions   []ExtensionElement   `json:"extensions,omitempty"`
	Owner        *OwnerElement        `json:"owner,omitempty"`
	Team         *TeamElement         `json:"team,omitempty"`
	ContactEmail *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels       []LabelElement       `json:"labels,omitempty"`
	Metadata     []MetadataElement    `json:"metadata,omitempty"`
	Capabilities *Capabilities        `json:"capabilities,omitempty"`
}

type WebRAEnrollRequestParams struct {
	Profile  string
	Template *WebRAEnrollTemplate
	// If the request allows password set on client side, give the password here
	Password string
}

type WebRAEnrollRequest struct {
	Id                   string               `json:"_id,omitempty"`
	Workflow             Workflow             `json:"workflow"`
	Module               Module               `json:"module,omitempty"`
	Status               Status               `json:"status,omitempty"`
	Profile              string               `json:"profile,omitempty"`
	Dn                   string               `json:"dn,omitempty"`
	Requester            string               `json:"requester,omitempty"`
	Approver             string               `json:"approver,omitempty"`
	Contact              string               `json:"contact,omitempty"`
	RequesterComment     string               `json:"requesterComment,omitempty"`
	ApproverComment      string               `json:"approverComment,omitempty"`
	RegistrationDate     int64                `json:"registrationDate"`
	LastModificationDate int64                `json:"lastModificationDate"`
	ExpirationDate       int64                `json:"expirationDate"`
	RemoveAt             int64                `json:"removeAt"`
	Template             *WebRAEnrollTemplate `json:"template,omitempty"`
	CertificatePEM       string               `json:"certificatePem,omitempty"`
	CertificateId        string               `json:"certificateId,omitempty"`
	Certificate          *Certificate         `json:"certificate,omitempty"`
	Pkcs12               *Secret              `json:"pkcs12,omitempty"`
	Password             *Secret              `json:"password,omitempty"`
	Labels               []Label              `json:"labels,omitempty"`
	Metadata             []Metadata           `json:"metadata,omitempty"`
	HolderId             string               `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                  `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                  `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRAEnrollRequest) EnsureType() error {
	if r.Module != WebRA {
		return invalidModuleError(r.Module, WebRA)
	}
	if r.Workflow != Enroll {
		return invalidWorkflowError(r.Workflow, Enroll)
	}
	return nil
}

type ScepChallengeTemplateParams struct {
	Profile string
}

type ScepChallengeTemplate struct {
	DnWhitelist  *bool                `json:"dnWhitelist,omitempty"`
	Subject      []IndexedDNElement   `json:"subject,omitempty"`
	Sans         []ListSANElement     `json:"sans,omitempty"`
	Extensions   []ExtensionElement   `json:"extensions,omitempty"`
	Owner        *OwnerElement        `json:"owner,omitempty"`
	Team         *TeamElement         `json:"team,omitempty"`
	ContactEmail *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels       []LabelElement       `json:"labels,omitempty"`
	Metadata     []MetadataElement    `json:"metadata,omitempty"`
	Capabilities *Capabilities        `json:"capabilities,omitempty"`
}

func (t *ScepChallengeTemplate) IsDnWhitelist() bool {
	return t.DnWhitelist != nil && *t.DnWhitelist
}

// Dn is mandatory if IsDnWhitelist is true for the template
type ScepChallengeRequestParams struct {
	Profile  string
	Template *ScepChallengeTemplate
	Dn       string
}

type ScepChallengeRequest struct {
	Id                   string                 `json:"_id,omitempty"`
	Workflow             Workflow               `json:"workflow"`
	Module               Module                 `json:"module,omitempty"`
	Status               Status                 `json:"status,omitempty"`
	Profile              string                 `json:"profile,omitempty"`
	Dn                   string                 `json:"dn,omitempty"`
	Requester            string                 `json:"requester,omitempty"`
	Approver             string                 `json:"approver,omitempty"`
	Contact              string                 `json:"contact,omitempty"`
	RequesterComment     string                 `json:"requesterComment,omitempty"`
	ApproverComment      string                 `json:"approverComment,omitempty"`
	RegistrationDate     int64                  `json:"registrationDate"`
	LastModificationDate int64                  `json:"lastModificationDate"`
	ExpirationDate       int64                  `json:"expirationDate"`
	RemoveAt             int64                  `json:"removeAt"`
	Template             *ScepChallengeTemplate `json:"template,omitempty"`
	CertificatePEM       string                 `json:"certificatePem,omitempty"`
	CertificateId        string                 `json:"certificateId,omitempty"`
	Certificate          *Certificate           `json:"certificate,omitempty"`
	Challenge            *Secret                `json:"password,omitempty"`
	Labels               []Label                `json:"labels,omitempty"`
	Metadata             []Metadata             `json:"metadata,omitempty"`
	HolderId             string                 `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                    `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                    `json:"profileHolderIdCount,omitempty"`
}

func (r *ScepChallengeRequest) EnsureType() error {
	if r.Module != Scep {
		return invalidModuleError(r.Module, Scep)
	}
	if r.Workflow != Enroll {
		return invalidWorkflowError(r.Workflow, Enroll)
	}
	return nil
}

type EstChallengeTemplateParams struct {
	Profile string
}

type EstChallengeTemplate struct {
	DnWhitelist  *bool                `json:"dnWhitelist,omitempty"`
	Subject      []IndexedDNElement   `json:"subject,omitempty"`
	Sans         []ListSANElement     `json:"sans,omitempty"`
	Extensions   []ExtensionElement   `json:"extensions,omitempty"`
	Owner        *OwnerElement        `json:"owner,omitempty"`
	Team         *TeamElement         `json:"team,omitempty"`
	ContactEmail *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels       []LabelElement       `json:"labels,omitempty"`
	Metadata     []MetadataElement    `json:"metadata,omitempty"`
	Capabilities *Capabilities        `json:"capabilities,omitempty"`
}

func (t *EstChallengeTemplate) IsDnWhitelist() bool {
	return t.DnWhitelist != nil && *t.DnWhitelist
}

// Dn is mandatory if IsDnWhitelist is true for the template
type EstChallengeRequestParams struct {
	Profile  string
	Template *EstChallengeTemplate
	Dn       string
}

type EstChallengeRequest struct {
	Id                   string                `json:"_id,omitempty"`
	Workflow             Workflow              `json:"workflow"`
	Module               Module                `json:"module,omitempty"`
	Status               Status                `json:"status,omitempty"`
	Profile              string                `json:"profile,omitempty"`
	Dn                   string                `json:"dn,omitempty"`
	Requester            string                `json:"requester,omitempty"`
	Approver             string                `json:"approver,omitempty"`
	Contact              string                `json:"contact,omitempty"`
	RequesterComment     string                `json:"requesterComment,omitempty"`
	ApproverComment      string                `json:"approverComment,omitempty"`
	RegistrationDate     int64                 `json:"registrationDate"`
	LastModificationDate int64                 `json:"lastModificationDate"`
	ExpirationDate       int64                 `json:"expirationDate"`
	RemoveAt             int64                 `json:"removeAt"`
	Template             *EstChallengeTemplate `json:"template"`
	CertificatePEM       string                `json:"certificatePem,omitempty"`
	CertificateId        string                `json:"certificateId,omitempty"`
	Certificate          *Certificate          `json:"certificate,omitempty"`
	Challenge            *Secret               `json:"password,omitempty"`
	Labels               []Label               `json:"labels,omitempty"`
	Metadata             []Metadata            `json:"metadata,omitempty"`
	HolderId             string                `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                   `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                   `json:"profileHolderIdCount,omitempty"`
}

func (r *EstChallengeRequest) EnsureType() error {
	if r.Module != Est {
		return invalidModuleError(r.Module, Est)
	}
	if r.Workflow != Enroll {
		return invalidWorkflowError(r.Workflow, Enroll)
	}
	return nil
}

type WebRARenewTemplateParams struct {
	CertificatePEM string
	CertificateId  string
}

// Capabilities is a readonly field
type WebRARenewTemplate struct {
	KeyType      string        `json:"keyType,omitempty"`
	Csr          string        `json:"csr,omitempty"`
	Capabilities *Capabilities `json:"capabilities,omitempty"`
}

type WebRARenewRequestParams struct {
	Template *WebRARenewTemplate
	// If the request allows password set on client side, give the password here
	Password       string
	CertToRenewId  string
	CertToRenewPem string
}

type WebRARenewRequest struct {
	Id                   string              `json:"_id,omitempty"`
	Workflow             Workflow            `json:"workflow"`
	Module               Module              `json:"module,omitempty"`
	Status               Status              `json:"status,omitempty"`
	Profile              string              `json:"profile,omitempty"`
	Dn                   string              `json:"dn,omitempty"`
	Requester            string              `json:"requester,omitempty"`
	Approver             string              `json:"approver,omitempty"`
	Contact              string              `json:"contact,omitempty"`
	RequesterComment     string              `json:"requesterComment,omitempty"`
	ApproverComment      string              `json:"approverComment,omitempty"`
	RegistrationDate     int64               `json:"registrationDate"`
	LastModificationDate int64               `json:"lastModificationDate"`
	ExpirationDate       int64               `json:"expirationDate"`
	RemoveAt             int64               `json:"removeAt"`
	Template             *WebRARenewTemplate `json:"template,omitempty"`
	CertificatePEM       string              `json:"certificatePem,omitempty"`
	CertificateId        string              `json:"certificateId,omitempty"`
	Certificate          *Certificate        `json:"certificate,omitempty"`
	Pkcs12               *Secret             `json:"pkcs12,omitempty"`
	Password             *Secret             `json:"password,omitempty"`
	Labels               []Label             `json:"labels,omitempty"`
	Metadata             []Metadata          `json:"metadata,omitempty"`
	HolderId             string              `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                 `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                 `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRARenewRequest) EnsureType() error {
	if r.Module != WebRA {
		return invalidModuleError(r.Module, WebRA)
	}
	if r.Workflow != Renew {
		return invalidWorkflowError(r.Workflow, Renew)
	}
	return nil
}

type WebRARevokeTemplate struct {
	RevocationReason RevocationReason `json:"revocationReason,omitempty"`
}

type WebRARevokeRequestParams struct {
	CertificateId    string
	CertificatePEM   string
	RevocationReason RevocationReason
}

type WebRARevokeRequest struct {
	Id                   string               `json:"_id,omitempty"`
	Workflow             Workflow             `json:"workflow"`
	Module               Module               `json:"module,omitempty"`
	Status               Status               `json:"status,omitempty"`
	Profile              string               `json:"profile,omitempty"`
	Dn                   string               `json:"dn,omitempty"`
	Requester            string               `json:"requester,omitempty"`
	Approver             string               `json:"approver,omitempty"`
	Contact              string               `json:"contact,omitempty"`
	RequesterComment     string               `json:"requesterComment,omitempty"`
	ApproverComment      string               `json:"approverComment,omitempty"`
	RegistrationDate     int64                `json:"registrationDate"`
	LastModificationDate int64                `json:"lastModificationDate"`
	ExpirationDate       int64                `json:"expirationDate"`
	RemoveAt             int64                `json:"removeAt"`
	Template             *WebRARevokeTemplate `json:"template,omitempty"`
	CertificatePEM       string               `json:"certificatePem,omitempty"`
	CertificateId        string               `json:"certificateId,omitempty"`
	Certificate          *Certificate         `json:"certificate,omitempty"`
	Password             *Secret              `json:"password,omitempty"`
	Labels               []Label              `json:"labels,omitempty"`
	Metadata             []Metadata           `json:"metadata,omitempty"`
	HolderId             string               `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                  `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                  `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRARevokeRequest) EnsureType() error {
	if r.Workflow != Revoke {
		return invalidWorkflowError(r.Workflow, Revoke)
	}
	return nil
}

type WebRAUpdateTemplateParams struct {
	CertificatePEM string
	CertificateId  string
}

type WebRAUpdateTemplate struct {
	Owner        *OwnerElement        `json:"owner,omitempty"`
	Team         *TeamElement         `json:"team,omitempty"`
	ContactEmail *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels       []LabelElement       `json:"labels,omitempty"`
	Metadata     []MetadataElement    `json:"metadata,omitempty"`
}

type WebRAUpdateRequestParams struct {
	CertificatePEM string
	CertificateId  string
	Template       *WebRAUpdateTemplate
}

type WebRAUpdateRequest struct {
	Id                   string               `json:"_id,omitempty"`
	Workflow             Workflow             `json:"workflow"`
	Module               Module               `json:"module,omitempty"`
	Status               Status               `json:"status,omitempty"`
	Profile              string               `json:"profile,omitempty"`
	Dn                   string               `json:"dn,omitempty"`
	Requester            string               `json:"requester,omitempty"`
	Approver             string               `json:"approver,omitempty"`
	Contact              string               `json:"contact,omitempty"`
	RequesterComment     string               `json:"requesterComment,omitempty"`
	ApproverComment      string               `json:"approverComment,omitempty"`
	RegistrationDate     int64                `json:"registrationDate"`
	LastModificationDate int64                `json:"lastModificationDate"`
	ExpirationDate       int64                `json:"expirationDate"`
	RemoveAt             int64                `json:"removeAt"`
	Template             *WebRAUpdateTemplate `json:"template,omitempty"`
	CertificatePEM       string               `json:"certificatePem,omitempty"`
	CertificateId        string               `json:"certificateId,omitempty"`
	Certificate          *Certificate         `json:"certificate,omitempty"`
	Password             *Secret              `json:"password,omitempty"`
	Labels               []Label              `json:"labels,omitempty"`
	Metadata             []Metadata           `json:"metadata,omitempty"`
	HolderId             string               `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                  `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                  `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRAUpdateRequest) EnsureType() error {
	// Do not validate module as it depends on target profile
	if r.Workflow != Update {
		return invalidWorkflowError(r.Workflow, Update)
	}
	return nil
}

type WebRAMigrateTemplateParams struct {
	CertificatePEM string
	CertificateId  string
	Profile        string
}

type WebRAMigrateTemplate struct {
	Owner        *OwnerElement        `json:"owner,omitempty"`
	Team         *TeamElement         `json:"team,omitempty"`
	ContactEmail *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels       []LabelElement       `json:"labels,omitempty"`
	Metadata     []MetadataElement    `json:"metadata,omitempty"`
}

type WebRAMigrateRequestParams struct {
	CertificatePEM string
	CertificateId  string
	Profile        string
	Template       *WebRAMigrateTemplate
}

type WebRAMigrateRequest struct {
	Id                   string                `json:"_id,omitempty"`
	Workflow             Workflow              `json:"workflow"`
	Module               Module                `json:"module,omitempty"`
	Status               Status                `json:"status,omitempty"`
	Profile              string                `json:"profile,omitempty"`
	Dn                   string                `json:"dn,omitempty"`
	Requester            string                `json:"requester,omitempty"`
	Approver             string                `json:"approver,omitempty"`
	Contact              string                `json:"contact,omitempty"`
	RequesterComment     string                `json:"requesterComment,omitempty"`
	ApproverComment      string                `json:"approverComment,omitempty"`
	RegistrationDate     int64                 `json:"registrationDate"`
	LastModificationDate int64                 `json:"lastModificationDate"`
	ExpirationDate       int64                 `json:"expirationDate"`
	RemoveAt             int64                 `json:"removeAt"`
	Template             *WebRAMigrateTemplate `json:"template,omitempty"`
	CertificatePEM       string                `json:"certificatePem,omitempty"`
	CertificateId        string                `json:"certificateId,omitempty"`
	Certificate          *Certificate          `json:"certificate,omitempty"`
	Password             *Secret               `json:"password,omitempty"`
	Labels               []Label               `json:"labels,omitempty"`
	Metadata             []Metadata            `json:"metadata,omitempty"`
	HolderId             string                `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                   `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                   `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRAMigrateRequest) EnsureType() error {
	// Do not validate module as it depends on target profile
	if r.Workflow != Migrate {
		return invalidWorkflowError(r.Workflow, Migrate)
	}
	return nil
}

type WebRARecoverRequestParams struct {
	CertificateId  string
	CertificatePEM string
	Password       string
	Contact        string
}

type WebRARecoverRequest struct {
	Id                   string       `json:"_id,omitempty"`
	Workflow             Workflow     `json:"workflow"`
	Module               Module       `json:"module,omitempty"`
	Status               Status       `json:"status,omitempty"`
	Profile              string       `json:"profile,omitempty"`
	Dn                   string       `json:"dn,omitempty"`
	Requester            string       `json:"requester,omitempty"`
	Approver             string       `json:"approver,omitempty"`
	Contact              string       `json:"contact,omitempty"`
	RequesterComment     string       `json:"requesterComment,omitempty"`
	ApproverComment      string       `json:"approverComment,omitempty"`
	RegistrationDate     int64        `json:"registrationDate"`
	LastModificationDate int64        `json:"lastModificationDate"`
	ExpirationDate       int64        `json:"expirationDate"`
	RemoveAt             int64        `json:"removeAt"`
	CertificatePEM       string       `json:"certificatePem,omitempty"`
	CertificateId        string       `json:"certificateId,omitempty"`
	Certificate          *Certificate `json:"certificate,omitempty"`
	Pkcs12               *Secret      `json:"pkcs12,omitempty"`
	Password             *Secret      `json:"password,omitempty"`
	Labels               []Label      `json:"labels,omitempty"`
	Metadata             []Metadata   `json:"metadata,omitempty"`
	HolderId             string       `json:"holderId,omitempty"`
	GlobalHolderIdCount  int          `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int          `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRARecoverRequest) EnsureType() error {
	if r.Module != WebRA {
		return invalidModuleError(r.Module, WebRA)
	}
	if r.Workflow != Recover {
		return invalidWorkflowError(r.Workflow, Recover)
	}
	return nil
}

type WebRAImportTemplateParams struct {
	CertificatePEM string
	CertificateId  string
	Profile        string
}

//discoveryData: Option[HostDiscoveryData] = None
//

type WebRAImportTemplate struct {
	PrivateKey     string               `json:"privateKey,omitempty"`
	Owner          *OwnerElement        `json:"owner,omitempty"`
	Team           *TeamElement         `json:"team,omitempty"`
	ContactEmail   *ContactEmailElement `json:"contactEmail,omitempty"`
	Labels         []LabelElement       `json:"labels,omitempty"`
	Metadata       []MetadataElement    `json:"metadata,omitempty"`
	ThirdPartyData []ThirdPartyItem     `json:"thirdPartyData,omitempty"`
	DiscoveryData  *DiscoveryData       `json:"discoveryData,omitempty"`
	DiscoveryInfo  *DiscoveryInfo       `json:"discoveryInfo,omitempty"`
}

type WebRAImportRequestParams struct {
	CertificatePEM string
	CertificateId  string
	Profile        string
	Template       *WebRAImportTemplate
}

type WebRAImportRequest struct {
	Id                   string               `json:"_id,omitempty"`
	Workflow             Workflow             `json:"workflow"`
	Module               Module               `json:"module,omitempty"`
	Status               Status               `json:"status,omitempty"`
	Profile              string               `json:"profile,omitempty"`
	Dn                   string               `json:"dn,omitempty"`
	Requester            string               `json:"requester,omitempty"`
	Approver             string               `json:"approver,omitempty"`
	Contact              string               `json:"contact,omitempty"`
	RequesterComment     string               `json:"requesterComment,omitempty"`
	ApproverComment      string               `json:"approverComment,omitempty"`
	RegistrationDate     int64                `json:"registrationDate"`
	LastModificationDate int64                `json:"lastModificationDate"`
	ExpirationDate       int64                `json:"expirationDate"`
	RemoveAt             int64                `json:"removeAt"`
	Template             *WebRAImportTemplate `json:"template,omitempty"`
	CertificatePEM       string               `json:"certificatePem,omitempty"`
	CertificateId        string               `json:"certificateId,omitempty"`
	Certificate          *Certificate         `json:"certificate,omitempty"`
	Password             *Secret              `json:"password,omitempty"`
	Labels               []Label              `json:"labels,omitempty"`
	Metadata             []Metadata           `json:"metadata,omitempty"`
	HolderId             string               `json:"holderId,omitempty"`
	GlobalHolderIdCount  int                  `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int                  `json:"profileHolderIdCount,omitempty"`
}

func (r *WebRAImportRequest) EnsureType() error {
	// Do not validate module as it depends on target profile
	if r.Workflow != Import {
		return invalidWorkflowError(r.Workflow, Import)
	}
	return nil
}

type SearchScope string

const (
	Search SearchScope = "search"
	Manage SearchScope = "manage"
	Self   SearchScope = "self"
)

type RequestSearchQuery struct {
	Query     string       `json:"query,omitempty"`
	Fields    []string     `json:"fields,omitempty"`
	SortedBy  []SortFields `json:"sortedBy,omitempty"`
	PageIndex int          `json:"pageIndex,omitempty"`
	PageSize  int          `json:"pageSize,omitempty"`
	WithCount bool         `json:"withCount,omitempty"`
	Scope     SearchScope  `json:"scope,omitempty"`
}

type RequestSearchResult struct {
	Id                   string       `json:"_id,omitempty"`
	Workflow             Workflow     `json:"workflow"`
	Module               Module       `json:"module,omitempty"`
	Status               Status       `json:"status,omitempty"`
	Profile              string       `json:"profile,omitempty"`
	Dn                   string       `json:"dn,omitempty"`
	Requester            string       `json:"requester,omitempty"`
	Approver             string       `json:"approver,omitempty"`
	Contact              string       `json:"contact,omitempty"`
	RequesterComment     string       `json:"requesterComment,omitempty"`
	ApproverComment      string       `json:"approverComment,omitempty"`
	RegistrationDate     int64        `json:"registrationDate"`
	LastModificationDate int64        `json:"lastModificationDate"`
	ExpirationDate       int64        `json:"expirationDate"`
	RemoveAt             int64        `json:"removeAt"`
	CertificateId        string       `json:"certificateId,omitempty"`
	Certificate          *Certificate `json:"certificate,omitempty"`
	Labels               []Label      `json:"labels,omitempty"`
	Metadata             []Metadata   `json:"metadata,omitempty"`
	HolderId             string       `json:"holderId,omitempty"`
	GlobalHolderIdCount  int          `json:"globalHolderIdCount,omitempty"`
	ProfileHolderIdCount int          `json:"profileHolderIdCount,omitempty"`
}
