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
	endpoint, _ := url.Parse("https://horizon-qa.int.evertrust.fr")
	baseClient.
		SetBaseUrl(*endpoint).
		SetPasswordAuth(
			"sma",
			"test",
		).
		SkipTLSVerify()
	client = Client{http: &baseClient}
}

func TestGetParams(t *testing.T) {
	_, err := client.GetParameters(os.Getenv("AUTOMATION_POLICY"))
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

func TestGet(t *testing.T) {
	policies, err := client.Get("Azdz")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(policies)
}
