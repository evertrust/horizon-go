// Package certificateprofiles provides utilities to interact with the Horizon api.certificate.profile APIs.
package certificateprofiles

type CertificateProfileCryptoPolicy struct {
	Centralized   bool `json:"centralized"`
	Decentralized bool `json:"decentralized"`
	Escrow        bool `json:"escrow"`
}

type AuthorizationLevel struct {
	AccessLevel string `json:"accessLevel"`
}

type Profile struct {
	Id                  string `json:"_id"`
	AuthorizationLevels struct {
		ApproveEnroll  AuthorizationLevel `json:"approveEnroll"`
		ApproveMigrate AuthorizationLevel `json:"approveMigrate"`
		ApproveRecover AuthorizationLevel `json:"approveRecover"`
		ApproveRevoke  AuthorizationLevel `json:"approveRevoke"`
		ApproveUpdate  AuthorizationLevel `json:"approveUpdate"`
		Enroll         AuthorizationLevel `json:"enroll"`
		Migrate        AuthorizationLevel `json:"migrate"`
		Recover        AuthorizationLevel `json:"recover"`
		RecoverApi     AuthorizationLevel `json:"recoverApi"`
		RequestEnroll  AuthorizationLevel `json:"requestEnroll"`
		RequestMigrate AuthorizationLevel `json:"requestMigrate"`
		RequestRecover AuthorizationLevel `json:"requestRecover"`
		RequestRevoke  AuthorizationLevel `json:"requestRevoke"`
		RequestUpdate  AuthorizationLevel `json:"requestUpdate"`
		Revoke         AuthorizationLevel `json:"revoke"`
		Search         AuthorizationLevel `json:"search"`
		Update         AuthorizationLevel `json:"update"`
	} `json:"authorizationLevels"`
	AuthorizationMode   string `json:"authorizationMode"`
	CA                  string `json:"ca"`
	CertificateTemplate struct {
		ContactEmailPolicy struct {
			EditableByApprover  bool `json:"editableByApprover"`
			EditableByRequester bool `json:"editableByRequester"`
			Mandatory           bool `json:"mandatory"`
		} `json:"contactEmailPolicy"`
		OwnerPolicy struct {
			EditableByApprover  bool `json:"editableByApprover"`
			EditableByRequester bool `json:"editableByRequester"`
			Mandatory           bool `json:"mandatory"`
		} `json:"ownerPolicy"`
		Sans []struct {
			EditableByApprover  bool   `json:"editableByApprover"`
			EditableByRequester bool   `json:"editableByRequester"`
			Type                string `json:"type"`
		} `json:"sans"`
		Subject []struct {
			EditableByApprover  bool   `json:"editableByApprover"`
			EditableByRequester bool   `json:"editableByRequester"`
			Type                string `json:"type"`
			Mandatory           bool   `json:"mandatory"`
		} `json:"subject"`
		TeamPolicy struct {
			EditableByApprover  bool `json:"editableByApprover"`
			EditableByRequester bool `json:"editableByRequester"`
			Mandatory           bool `json:"mandatory"`
		} `json:"teamPolicy"`
	} `json:"certificateTemplate"`
	Constraints  struct{} `json:"constraints"`
	CryptoPolicy struct {
		AuthorizedKeyTypes       []string `json:"authorizedKeyTypes"`
		Centralized              bool     `json:"centralized"`
		Decentralized            bool     `json:"decentralized"`
		DefaultKeyType           string   `json:"defaultKeyType"`
		Escrow                   bool     `json:"escrow"`
		P12PasswordMode          string   `json:"p12PasswordMode"`
		P12EncryptionType        string   `json:"p12EncryptionType"`
		ShowP12OnRecover         bool     `json:"showP12OnRecover"`
		ShowP12PasswordOnRecover bool     `json:"showP12PasswordOnRecover"`
	} `json:"cryptoPolicy"`
	CsrDataMapping       struct{}                 `json:"csrDataMapping"`
	Description          []map[string]interface{} `json:"description"`
	DisplayName          []map[string]interface{} `json:"displayName"`
	DnWhitelist          bool                     `json:"dnWhitelist"`
	Enabled              bool                     `json:"enabled"`
	EnrollAuthorizedCAs  []string                 `json:"enrollAuthorizedCas"`
	Module               string                   `json:"module"`
	Name                 string                   `json:"name"`
	PkiConnector         string                   `json:"pkiConnector"`
	RenewalAuthorizedCAs []string                 `json:"renewalAuthorizedCas"`
	RenewalPeriod        string                   `json:"renewalPeriod"`
	RequestsPolicy       struct {
		Migrate string `json:"migrate"`
		Revoke  string `json:"revoke"`
		Update  string `json:"update"`
	} `json:"requestsPolicy"`
	SelfPermissions struct {
		SelfPopRenew  bool `json:"selfPopRenew"`
		SelfPopRevoke bool `json:"selfPopRevoke"`
		SelfPopUpdate bool `json:"selfPopUpdate"`
		SelfRecover   bool `json:"selfRecover"`
		SelfRenew     bool `json:"selfRenew"`
		SelfRevoke    bool `json:"selfRevoke"`
		SelfUpdate    bool `json:"selfUpdate"`
	} `json:"selfPermissions"`
	Triggers struct{} `json:"triggers"`
}
