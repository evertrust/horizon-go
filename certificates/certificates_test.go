package certificates

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
	baseClient.Init(
		*endpoint,
		os.Getenv("APIID"),
		os.Getenv("APIKEY"),
		"",
		"",
	)
	client = Client{Http: &baseClient}
}

func TestSearch(t *testing.T) {
	_, hasMore, count, err := client.Search("status is valid", 0, true)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("hasMore: %v", hasMore)
	t.Logf("count: %v", count)
}

func TestGet(t *testing.T) {
	certs, _, _, err := client.Search("status is valid", 0, true)
	if err != nil {
		t.Error(err.Error())
	}
	cert, err := client.Get(certs[0].Id)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("cert: %v", cert.Serial)
}

func TestUpdateMigrate(t *testing.T) {
	certs, _, _, err := client.Search("status is valid and module not in [\"discovery\"]", 0, true)
	if err != nil {
		t.Error(err.Error())
	}
	err = client.UpdateMigrate(certs[0].Id, "", "", "", "", "")
	if err != nil {
		t.Error(err.Error())
	}
}
