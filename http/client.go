// Package http provides low-level methods to interact with the Horizon instance through its REST API.
package http

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"gopkg.in/resty.v1"
)

type Client struct {
	baseUrl     url.URL
	restyClient resty.Client
	cert        string
	key         string
	apiId       string
	apiKey      string
}

// Init initializes the instance parameters such as its location, and authentication data.
func (c *Client) Init(baseUrl url.URL, apiId string, apiKey string, cert string, key string) {
	c.baseUrl = baseUrl
	c.apiId = apiId
	c.apiKey = apiKey
	c.cert = cert
	c.key = key

	c.restyClient = *resty.New()
	// load cert and key
	clientCert, err := tls.LoadX509KeyPair(c.cert, c.key)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	c.restyClient.
		SetCertificates(clientCert).
		SetHeader("Content-Type", "application/json").
		SetHostURL(baseUrl.String()).
		SetHeader("X-API-ID", c.apiId).
		SetHeader("X-API-KEY", c.apiKey).
		SetCookieJar(nil)
}

func (c *Client) Unmarshal(r *resty.Response) (*HorizonResponse, error) {
	body := r.Body()

	if r.StatusCode() > 300 {
		if r.Header().Get("Content-Type") == "application/json" {
			// Deserialize the response to an error
			var horizonError HorizonErrorResponse
			var horizonMultiError HorizonMultipleErrorsResponse
			if err := json.Unmarshal(body, &horizonError); err != nil {
				if err := json.Unmarshal(body, &horizonMultiError); err != nil {
					log.Fatalf("(HTTP %d) error deserializing error JSON: %s", r.StatusCode(), string(body))
				}
				return &HorizonResponse{
					RestyResponse: r,
				}, &horizonMultiError
			}
			return &HorizonResponse{
				RestyResponse: r,
			}, &horizonError
		} else {
			return nil, &HorizonErrorResponse{
				Code:    "Unknown",
				Message: "Non-JSON error from Horizon",
				Detail:  string(body),
			}
		}
	}

	response := HorizonResponse{
		RestyResponse: r,
	}

	return &response, nil
}

func (c *Client) Get(path string) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().Get(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) Post(path string, body []byte) (response *HorizonResponse, err error) {
	req := c.newRequest().SetBody(body)

	resp, err := req.Post(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) PostWithJwt(path, jwt string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().SetBody(body).SetHeader("X-JWT-CERT-POP", jwt).Post(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) GetWithJwt(path, jwt string) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().SetHeader("X-JWT-CERT-POP", jwt).Get(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) Delete(path string) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().Delete(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) Put(path string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().
		SetBody(body).
		Put(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) newRequest() *resty.Request {
	return c.restyClient.R().
		SetHeader("X-API-ID", c.apiId).
		SetHeader("X-API-KEY", c.apiKey).
		SetHeader("Accept", "application/json")
}
