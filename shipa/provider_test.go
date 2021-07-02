package shipa

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("SHIPA_HOST"); v == "" {
		t.Fatal("SHIPA_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("SHIPA_TOKEN"); v == "" {
		t.Fatal("SHIPA_TOKEN must be set for acceptance tests")
	}
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"example": testAccProvider,
	}
}
