package certificateprofiles

import (
	"encoding/json"
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
	res, err := client.Get("SSL")
	if err != nil {
		t.Error(err.Error())
	}

	jsonString, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return
	}
	t.Log(string(jsonString))
}
