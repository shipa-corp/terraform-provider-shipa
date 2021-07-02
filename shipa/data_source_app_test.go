package shipa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func TestAccApp_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShipaAppDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("shipa_app.tf", "name", "app1"),
				),
			},
		},
	})
}

func testAccShipaAppDatasourceConfig() string {
	return fmt.Sprintf(`
resource "shipa_app" "tf" {
  app {
    name = "app1"
    teamowner = shipa_team.tf.team[0].name
    framework = shipa_framework.tf.framework[0].name
    description = "test description update"
    tags = ["dev"]
  }
  depends_on = [shipa_cluster.tf]
}
`)
}

func testAccCheckAppDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(client.Client)
	for _, rs := range s.RootModule().Resources {
		app, err := c.GetApp(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("not destroyed: %v", app)
		}
	}
	return nil
}
