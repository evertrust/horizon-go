package client

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"github.com/evertrust/horizon-go/automation"
	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/discovery"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/license"
	"github.com/evertrust/horizon-go/locals"
	mylog "github.com/evertrust/horizon-go/log"
	"github.com/evertrust/horizon-go/requests"
	"github.com/evertrust/horizon-go/rfc5280"
	"io"
	"log"
	gohttp "net/http"
	"net/url"
	"time"
)

type Horizon struct {
	Requests    *requests.Client
	License     *license.Client
	Rfc5280     *rfc5280.Client
	Certificate *certificates.Client
	Discovery   *discovery.Client
	Automation  *automation.Client
	Locals      *locals.Client
	Http        *http.Client
}

// New instantiates a new Horizon client
func New(httpClient *http.Client) *Horizon {
	var client Horizon
	if httpClient == nil {
		httpClient = &http.Client{}
		httpClient.SetHttpClient(nil)
	}
	return client.init(httpClient)
}

func (client *Horizon) SetDebugWriter(writer io.Writer) *Horizon {
	log.SetOutput(writer)
	mylog.LogEnabled = true
	return client
}

// SetHttpClient initializes the instance parameters such as its location, and authentication data.
func (c *Horizon) SetHttpClient(httpClient *gohttp.Client) *Horizon {
	c.Http.SetHttpClient(httpClient)
	return c
}

// SetBaseUrl sets the base url for the client
// This is the endpoint without any additional path
// For example: https://horizon-test.com
func (c *Horizon) SetBaseUrl(baseUrl url.URL) *Horizon {
	c.Http.SetBaseUrl(baseUrl)
	return c
}

func (c *Horizon) ClearAuth() {
	c.Http.ClearAuth()
}

func (c *Horizon) SetPasswordAuth(apiId string, apiKey string) *Horizon {
	c.Http.SetPasswordAuth(apiId, apiKey)
	return c
}

// SetCertAuth sets the client certificate than can be used for authentication.
func (c *Horizon) SetCertAuth(cert tls.Certificate) *Horizon {
	c.Http.SetCertAuth(cert)
	return c
}

func (c *Horizon) SetJwtAuth(cert x509.Certificate, key crypto.Signer) *Horizon {
	c.Http.SetJwtAuth(cert, key)
	return c
}

// SetCaBundle sets the CA bundle
func (c *Horizon) SetCaBundle(caBundle *x509.CertPool) *Horizon {
	c.Http.SetCaBundle(caBundle)
	return c
}

func (c *Horizon) GetCaBundle() *x509.CertPool {
	return c.GetTlsConfig().RootCAs
}

// SkipTLSVerify skips the TLS verification.
func (c *Horizon) SkipTLSVerify() *Horizon {
	c.Http.SkipTLSVerify()
	return c
}

// SetTimeout sets the timeout for http requests
func (c *Horizon) SetTimeout(timeout time.Duration) *Horizon {
	c.Http.SetTimeout(timeout)
	return c
}

func (c *Horizon) SetProxy(proxyUrl url.URL) *Horizon {
	c.Http.SetProxy(proxyUrl)
	return c
}

func (c *Horizon) GetTransport() *gohttp.Transport {
	return c.Http.GetTransport()
}

func (c *Horizon) GetTlsConfig() *tls.Config {
	return c.Http.GetTlsConfig()
}

func (c *Horizon) BaseUrl() (url.URL, error) {
	return c.Http.BaseUrl()
}

// init initializes the instance parameters such as its location, and authentication data.
func (client *Horizon) init(httpClient *http.Client) *Horizon {
	client.Http = httpClient
	client.Requests = requests.Init(httpClient)
	client.License = &license.Client{Http: client.Http}
	client.Rfc5280 = &rfc5280.Client{Http: client.Http}
	client.Certificate = certificates.Init(httpClient)
	client.Discovery = &discovery.Client{Http: client.Http}
	client.Automation = automation.Init(httpClient)
	client.Locals = &locals.Client{Http: client.Http}
	return client
}
