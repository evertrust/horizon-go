package automation

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

func TestGet(t *testing.T) {
	_, err := client.Get(os.Getenv("AUTOMATION_POLICY"))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCheck(t *testing.T) {
	// TODO
	t.Skip()
}

func TestList(t *testing.T) {
	policies, err := client.List()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(policies)
}

func TestNonce(t *testing.T) {
	nonce, err := client.Nonce(os.Getenv("AUTOMATION_POLICY"))
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(nonce)
}
