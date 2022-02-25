package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
	"github.com/shipa-corp/terraform-provider-shipa/helper"
)

func TestClient_GetApp(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/apps/app", App{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetApp("app"); err != nil {
		t.Error(err)
	}
}

func TestClient_ListApps(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/apps", []App{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.ListApps(); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateApp(t *testing.T) {
	payload := &App{Name: "app", Pool: "some-pool"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/apps", payload, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateApp(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_UpdateApp(t *testing.T) {
	payload := &UpdateAppRequest{Plan: "new-plan"}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/apps/app", payload, http.MethodPut),
	)
	defer teardown()

	if err := client.UpdateApp("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteApp(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/apps/app", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteApp("app"); err != nil {
		t.Error(err)
	}
}

func TestClient_DeployApp(t *testing.T) {
	payload := &AppDeploy{Image: "dockerhub-image", Port: 8080, Detach: true, Message: "test message", RegistryUser: "steve", RegistrySecret: "secret"}
	client, teardown := setupServer(
		clientest.NewHandler("/apps/app/", func(w http.ResponseWriter, request *http.Request) {
			var app AppDeployRequest
			err := json.NewDecoder(request.Body).Decode(&app)
			if err != nil {
				panic(err)
			}
			if app.Image != payload.Image {
				panic(fmt.Sprintf("got %v, expected %v", app.Image, payload.Image))
			}
			if app.Port.Number != int(payload.Port) {
				panic(fmt.Sprintf("got %v, expected %v", app.Port.Number, payload.Port))
			}
			if app.Detach != payload.Detach {
				panic(fmt.Sprintf("got %v, expected %v", app.Image, payload.Image))
			}
			if app.Message != payload.Message {
				panic(fmt.Sprintf("got %v, expected %v", app.Image, payload.Image))
			}
			if app.Registry.User != payload.RegistryUser {
				panic(fmt.Sprintf("got %v, expected %v", app.Registry.User, payload.RegistryUser))
			}
			if request.Method != http.MethodPost {
				panic(fmt.Errorf("method doesn't match, want %s, got %s", http.MethodPost, request.Method))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}),
	)
	defer teardown()

	if err := client.DeployApp("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateAppCname(t *testing.T) {
	payload := &AppCname{Cname: "shipa.io"}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppCname("app"), payload, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateAppCname("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_UpdateAppCname(t *testing.T) {
	payload := &AppCname{Cname: "shipa.io"}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppCname("app"), payload, http.MethodPut),
	)
	defer teardown()

	if err := client.UpdateAppCname("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteAppCname(t *testing.T) {
	payload := &AppCnames{Cnames: []string{"shipa.io"}}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppCname("app"), payload, http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteAppCname("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateAppEnvs(t *testing.T) {
	payload := &CreateAppEnv{NoRestart: true}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppEnvs("app"), payload, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateAppEnvs("app", payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteAppEnvs(t *testing.T) {
	payload := &DeleteAppEnv{}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppEnvs("app"), payload, http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteAppEnvs("app", payload); err != nil {
		t.Error(err)
	}
}

func TestAppToTerraformStruct(t *testing.T) {
	input := []byte(`{"name":"terraform-app-1","platform":"","teams":["dev"],"units":[],"plan":{"name":"shipa-plan","memory":0,"swap":0,"cpushare":100,"router":"traefik"},"ip":"terraform-app-1.104.200.27.23.shipa.cloud","router":"traefik","routeropts":{},"entrypoints":[],"cluster":"ln-cl2","owner":"dev@shipa.io","pool":"test-tf-142","description":"test","deploys":0,"teamowner":"dev","lock":{"Locked":false,"Reason":"","Owner":"","AcquireDate":"0001-01-01T00:00:00Z"},"tags":["dev","sandbox"],"routers":[{"name":"traefik","opts":{},"address":"terraform-app-1.104.200.27.23.shipa.cloud","type":"traefik"}],"routingsettings":null,"dependencyfilenames":["app.yaml","app.yml","Procfile","shipa-ci.yml","shipa.yaml","shipa.yml"],"status":"created","provisioner":"kubernetes","provisionerprops":{"kubernetes":{"namespace":"shipa-test-tf-142","service_account":"app-terraform-app-1","internal_dns_name":"app-terraform-app-1.shipa-test-tf-142.svc"}}}`)
	app := &App{}
	if err := json.Unmarshal(input, app); err != nil {
		t.Error(err)
	}

	rawData := helper.StructToTerraform(app)
	log.Printf("%+v", rawData)
}
