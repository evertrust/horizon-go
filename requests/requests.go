package requests

import (
	"github.com/evertrust/horizon-go/certificateprofiles"
	"github.com/evertrust/horizon-go/certificates"
)

type RequestWorkflow string

const (
	RequestWorkflowEnroll  RequestWorkflow = "enroll"
	RequestWorkflowRevoke  RequestWorkflow = "revoke"
	RequestWorkflowUpdate  RequestWorkflow = "update"
	RequestWorkflowRecover RequestWorkflow = "recover"
)

type RequestStatus string

const (
	RequestStatusDenied    RequestStatus = "denied"
	RequestStatusPending   RequestStatus = "pending"
	RequestStatusApproved  RequestStatus = "approved"
	RequestStatusCanceled  RequestStatus = "canceled"
	RequestStatusCompleted RequestStatus = "completed"
)

type IndexedDNElement struct {
	Element string `json:"element"`
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
}

type IndexedSANElement struct {
	Element string `json:"element"`
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
}

type LabelElement struct {
	Label string `json:"label"`
	Value string `json:"value,omitempty"`
}

type CertificateOwner struct {
	Value    string `json:"value"`
	Editable bool   `json:"editable"`
}

type CertificateTeam struct {
	Value      string   `json:"value"`
	Authorized []string `json:"authorized,omitempty"`
	Editable   bool     `json:"editable"`
}

type WebRARequestTemplate struct {
	Subject      []IndexedDNElement                                 `json:"subject"`
	Sans         []IndexedSANElement                                `json:"sans"`
	Labels       []LabelElement                                     `json:"labels"`
	KeyTypes     []string                                           `json:"keyTypes"`
	Capabilities certificateprofiles.CertificateProfileCryptoPolicy `json:"capabilities"`
	Csr          string                                             `json:"csr,omitempty"`
	Owner        *CertificateOwner                                  `json:"owner,omitempty"`
	Team         *CertificateTeam                                   `json:"team,omitempty"`
}

type WebRARevokeTemplate struct {
	RevocationReason certificates.RevocationReason `json:"revocationReason,omitempty"`
}

type P12Password struct {
	HorizonKey string `json:"horizonKey"`
	Value      string `json:"value"`
	Transient  bool   `json:"transient"`
}

type HorizonRequest struct {
	Id                   string                    `json:"_id,omitempty"`
	Workflow             RequestWorkflow           `json:"workflow"`
	Module               string                    `json:"module,omitempty"`
	Status               RequestStatus             `json:"status,omitempty"`
	Profile              string                    `json:"profile"`
	Dn                   string                    `json:"dn"`
	Requester            string                    `json:"requester"`
	Approver             string                    `json:"approver"`
	Contact              string                    `json:"contact"`
	RequesterComment     string                    `json:"requesterComment"`
	ApproverComment      string                    `json:"approverComment"`
	RegistrationDate     int                       `json:"registrationDate"`
	LastModificationDate int                       `json:"lastModificationDate"`
	Template             interface{}               `json:"template"`
	CertificatePEM       string                    `json:"certificatePem,omitempty"`
	CertificateId        string                    `json:"certificateId,omitempty"`
	Certificate          *certificates.Certificate `json:"certificate,omitempty"`
	Pkcs12               struct {
		HorizonKey string `json:"horizonKey"`
		Value      string `json:"value"`
		Transient  bool   `json:"transient"`
	} `json:"pkcs12"`
	Password P12Password `json:"password"`
}
