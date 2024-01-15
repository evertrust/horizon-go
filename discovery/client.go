package discovery

import (
	"encoding/json"
	"github.com/evertrust/horizon-go"
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

// Feed feeds a certificate to the discovery campaign
func (c *Client) Feed(certificate horizon.DiscoveredCertificateParams, session *horizon.DiscoverySession) error {
	cert := &horizon.DiscoveredCertificate{
		DiscoveryCampaign: session.Campaign,
		SessionId:         session.Id,
		Certificate:       certificate.Certificate,
		Code:              certificate.Code,
		DiscoveryData:     certificate.DiscoveryData,
		ContactEmail:      certificate.ContactEmail,
		ThirdPartyData:    certificate.ThirdPartyData,
		Metadata:          certificate.Metadata,
		PrivateKey:        certificate.PrivateKey,
	}
	marshalledData, err := json.Marshal(cert)
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
func (c *Client) Start(name string) (*horizon.DiscoverySession, error) {
	res, err := c.Http.Get("/api/v1/discovery/feed/" + name)
	if err != nil {
		return nil, err
	}
	var session horizon.DiscoverySession
	if err = res.Json().Decode(&session); err != nil {
		return nil, err
	}
	return &session, err
}

func (c *Client) Stop(session *horizon.DiscoverySession) (err error) {
	_, err = c.Http.Delete("/api/v1/discovery/feed/" + session.Campaign + "/" + session.Id)
	return err
}

func (c *Client) Event(event horizon.DiscoveryEventParams, session *horizon.DiscoverySession) error {
	hrzEvent := &horizon.DiscoveryEvent{
		Campaign:     session.Campaign,
		SessionId:    session.Id,
		Code:         event.Code,
		Status:       event.Status,
		ErrorCode:    event.ErrorCode,
		ErrorMessage: event.ErrorMessage,
	}
	marshalledData, err := json.Marshal(hrzEvent)
	if err != nil {
		return err
	}
	_, err = c.Http.Put("/api/v1/discovery/feed", marshalledData)
	return err
}

func (c *Client) Events(events []horizon.DiscoveryEventParams, session *horizon.DiscoverySession) error {
	var completeEvents []horizon.DiscoveryEvent
	for i := 0; i < len(events); i++ {
		hrzEvent := horizon.DiscoveryEvent{
			Campaign:     session.Campaign,
			SessionId:    session.Id,
			Code:         events[i].Code,
			Status:       events[i].Status,
			ErrorCode:    events[i].ErrorCode,
			ErrorMessage: events[i].ErrorMessage,
		}
		completeEvents = append(completeEvents, hrzEvent)
	}
	marshalledData, err := json.Marshal(completeEvents)
	if err != nil {
		return err
	}
	_, err = c.Http.Put("/api/v1/discovery/feed", marshalledData)
	return err
}

// Create a new discovery campaign
func (c *Client) Create(campaign horizon.DiscoveryCampaign) error {
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
