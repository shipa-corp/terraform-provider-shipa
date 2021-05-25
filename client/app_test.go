package client

import (
	"encoding/json"
	"github.com/marcustreacy/go-terraform-provider/helper"
	"log"
	"testing"
)

func TestAppRead(t *testing.T) {
	input := []byte(`{"name":"terraform-app-1","platform":"","teams":["dev"],"units":[],"plan":{"name":"shipa-plan","memory":0,"swap":0,"cpushare":100,"router":"traefik"},"ip":"terraform-app-1.104.200.27.23.shipa.cloud","router":"traefik","routeropts":{},"entrypoints":[],"cluster":"ln-cl2","owner":"dev@shipa.io","pool":"test-tf-142","description":"test","deploys":0,"teamowner":"dev","lock":{"Locked":false,"Reason":"","Owner":"","AcquireDate":"0001-01-01T00:00:00Z"},"tags":["dev","sandbox"],"routers":[{"name":"traefik","opts":{},"address":"terraform-app-1.104.200.27.23.shipa.cloud","type":"traefik"}],"routingsettings":null,"dependencyfilenames":["app.yaml","app.yml","Procfile","shipa-ci.yml","shipa.yaml","shipa.yml"],"status":"created","provisioner":"kubernetes","provisionerprops":{"kubernetes":{"namespace":"shipa-test-tf-142","service_account":"app-terraform-app-1","internal_dns_name":"app-terraform-app-1.shipa-test-tf-142.svc"}}}`)
	app := &App{}
	json.Unmarshal(input, app)

	rawData := helper.StructToTerraform(app)
	log.Printf("%+v", rawData)
}

func TestDeployApp(t *testing.T) {
	c := getClient()

	err := c.DeployApp("terraform-app-2", &AppDeploy{
		Image: "docker.io/shipasoftware/bulletinboard:1.0",
	})
	if err != nil {
		t.Error(err.Error())
	}
}