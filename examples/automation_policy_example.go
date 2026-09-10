package examples

import (
	"context"
	"fmt"
	"log"

	horizon "github.com/evertrust/horizon-go/v2"
	"github.com/evertrust/horizon-go/v2/models"
)

// ExampleCreateAutomationPolicy demonstrates how to create an automation policy.
func ExampleCreateAutomationPolicy() {
	c := horizon.NewConfiguration()
	c.Servers[0].URL = "https://horizon.example.com"
	c.SetPasswordAuth("user", "password")
	client := horizon.NewAPIClient(c)

	// Create an automation policy object
	automationPolicy := models.AutomationPolicy{
		Name:    "DummyPolicy",
		Profile: "profile",
	}

	// Request the creation of the automation policy
	creationResponse, _, err := client.AutomationPolicyAPI.AutomationPolicyAdd(context.Background()).AutomationPolicy(automationPolicy).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(creationResponse)
}
