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
	restyClient resty.Client
}

// WithRestyClient initializes the instance parameters such as its location, and authentication data.
func (c *Client) WithRestyClient(restyClient *resty.Client) {

	if restyClient == nil {
		restyClient = resty.New()
	}

	c.restyClient = *restyClient
	c.restyClient.
		SetHeader("Content-Type", "application/json").
		SetCookieJar(nil)
}

func (c *Client) WithBaseUrl(baseUrl url.URL) *Client {
	c.restyClient.SetHostURL(baseUrl.String())
	return c
}

func (c *Client) WithPasswordAuth(apiId string, apiKey string) *Client {
	c.restyClient.
		SetHeader("X-API-ID", apiId).
		SetHeader("X-API-KEY", apiKey)
	return c
}

func (c *Client) WithCertAuth(cert tls.Certificate) *Client {
	c.restyClient.SetCertificates(cert)
	return c
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
					return nil, fmt.Errorf("cannot deserialize error JSON: %s", string(body))
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

func (c *Client) Patch(path string, body []byte) (response *HorizonResponse, err error) {
	resp, err := c.newRequest().
		SetBody(body).
		Patch(path)
	if err != nil {
		return nil, err
	}

	return c.Unmarshal(resp)
}

func (c *Client) BaseUrl() url.URL {
	baseUrl, err := url.Parse(c.restyClient.HostURL)
	if err != nil {
		log.Fatal(err)
	}
	return *baseUrl
}

// SetCaBundle sets the client certificate than can be used for authentication.
func (c *Client) SetCaBundle(caBundle string) {
	c.restyClient.SetRootCertificate(caBundle)
}

// SkipTLSVerify skips the TLS verification.
// Warning: this will override any other TLS configuration, due to a limitation in the underlying library.
func (c *Client) SkipTLSVerify() {
	c.restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}

func (c *Client) SetProxy(proxyUrl url.URL) {
	c.restyClient.SetProxy(proxyUrl.String())
}

func (c *Client) newRequest() *resty.Request {
	return c.restyClient.R()
}
