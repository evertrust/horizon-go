package license

type ModuleLicenseInfo struct {
	Module string `json:"module"`
	Items  int    `json:"items"`
	Limit  int    `json:"limit"`
}

type LibraryInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type LicenseInfo struct {
	IsValid    bool                `json:"isValid"`
	Expiration int                 `json:"expiration"`
	Version    string              `json:"version"`
	BuildTime  int                 `json:"buildTime"`
	Modules    []ModuleLicenseInfo `json:"modules"`
	Libraries  []LibraryInfo       `json:"libraries"`
}
