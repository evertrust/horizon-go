package discovery

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

// Feed feeds a certificate to the discovery campaign
func (c *Client) Feed(discoveryInfo HrzDiscoveredCert) error {
	marshalledData, err := json.Marshal(discoveryInfo)
	if err != nil {
		return err
	}
	_, err = c.Http.Post("/api/v1/discovery/feed", marshalledData)
	if err != nil {
		return err
	}
	return nil
}

// Start a discovery campaign
func (c *Client) Start(name string) (HrzDiscoveryCampaign, error) {
	res, err := c.Http.Get("/api/v1/discovery/feed/" + name)
	if err != nil {
		return HrzDiscoveryCampaign{}, err
	}
	var campaign HrzDiscoveryCampaign
	if err = res.Json().Decode(&campaign); err != nil {
		return HrzDiscoveryCampaign{}, err
	}
	return campaign, err
}

func (c *Client) Stop(name, sessionId string) (err error) {
	_, err = c.Http.Delete("/api/v1/discovery/feed/" + name + "/" + sessionId)
	return err
}

func (c *Client) Event(event HrzDiscoveryEvent) error {
	marshalledData, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = c.Http.Put("/api/v1/discovery/feed", marshalledData)
	return err
}

// Create a new discovery campaign
func (c *Client) Create(campaign DiscoveryCampaign) error {
	marshalledData, err := json.Marshal(campaign)
	if err != nil {
		return err
	}
	_, err = c.Http.Post("/api/v1/discovery/campaigns", marshalledData)
	return err
}

// Delete a discovery campaign
func (c *Client) Delete(campaignID string) error {
	_, err := c.Http.Delete("/api/v1/discovery/campaigns/" + campaignID)
	return err
}
