package license

import (
	"net/url"
	"os"
	"testing"

	"github.com/evertrust/horizon-go/http"
)

var client Client

func init() {
	var baseClient = http.Client{}
	endpoint, _ := url.Parse(os.Getenv("ENDPOINT"))
	baseClient.SetBaseUrl(*endpoint)
	baseClient.InitPasswordAuth(
		os.Getenv("APIID"),
		os.Getenv("APIKEY"),
	)
	client = Client{Http: &baseClient}
}

func TestGet(t *testing.T) {
	license, err := client.Get()
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("license: %v", license)
}
