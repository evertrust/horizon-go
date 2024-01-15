package discovery

import (
	"net/url"
	"os"
	"testing"

	"github.com/evertrust/horizon-go/http"
)

var client Client

var sessionId string

func init() {
	var baseClient = http.Client{}
	endpoint, _ := url.Parse(os.Getenv("ENDPOINT"))
	baseClient.
		WithBaseUrl(*endpoint).
		WithPasswordAuth(os.Getenv("APIID"), os.Getenv("APIKEY"))
	client = Client{Http: &baseClient}
}

func TestCreate(t *testing.T) {
	campaign := DiscoveryCampaign{
		Name: "testCampaign",
		AuthorizationLevels: AuthorizationLevels{
			Search: AuthorizationLevel{
				AccessLevel: "authorized",
				EnforcedIdentityProviders: []EnforcedIdentityProviders{
					{
						Name: "local",
						Type: "Local",
					},
				},
			},
			Feed: AuthorizationLevel{
				AccessLevel: "authorized",
				EnforcedIdentityProviders: []EnforcedIdentityProviders{
					{
						Name: "local",
						Type: "Local",
					},
				},
			},
		},
		EventOnSuccess: false,
		EventOnFailure: false,
		EventOnWarning: false,
		Enabled:        true,
	}
	err := client.Create(campaign)
	if err != nil {
		t.Error(err)
	}
}

func TestStart(t *testing.T) {
	campaign, err := client.Start("testCampaign")
	if err != nil {
		t.Error(err)
	}
	sessionId = campaign.Id
}

func TestFeed(t *testing.T) {
	certPem := `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`
	err := client.Feed(HrzDiscoveredCert{
		DiscoveryCampaign: "testCampaign",
		Certificate:       certPem,
		Metadata:          nil,
		DiscoveryInfos: HrzDiscoveredCertsMetadata{
			CertificateLocation:      []string{"gitlab.com"},
			CertificateUsageLocation: []string{"mysuperusage"},
			OS:                       []string{"linux"},
			Hostname:                 []string{"gitlab.com"},
			IP:                       "172.123.21.9",
			Source:                   []string{"unittesting"},
			TLSPorts: []HrzTLSPorts{
				{
					Port:    443,
					Version: "TLSv1.2",
				},
			},
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestEvent(t *testing.T) {
	err := client.Event(Event{
		Code:         "NETIMPORT",
		Campaign:     "testCampaign",
		SessionId:    sessionId,
		Status:       "failure",
		ErrorCode:    "HCL-NETIMPORT-DCC-001",
		ErrorMessage: "Could not get certificates from DigiCert CertCentral",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestStop(t *testing.T) {
	err := client.Stop("testCampaign", sessionId)
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	err := client.Delete("testCampaign")
	if err != nil {
		t.Error(err)
	}
}
