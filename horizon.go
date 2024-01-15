// Package horizon provides high-level methods to interact with an Horizon instance.
package horizon

import (
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

func NewHttpClient() *http.Client {
	return &http.Client{}
}

func (client *Horizon) SetDebugWriter(writer io.Writer) *Horizon {
	log.SetOutput(writer)
	mylog.LogEnabled = true
	return client
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
