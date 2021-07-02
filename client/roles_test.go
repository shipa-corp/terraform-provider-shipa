package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetRole(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/roles/role", Role{Name: "role"}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetRole("role"); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateRole(t *testing.T) {
	payload := &Role{Name: "role"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/roles", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateRole(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteRole(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/roles/role", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteRole("role"); err != nil {
		t.Error(err)
	}
}

func TestClient_AssociateRoleToUser(t *testing.T) {
	payload := Email{Email: "shipa.io"}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiRoleUser("role"), clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.AssociateRoleToUser("role", "user@shipa.io"); err != nil {
		t.Error(err)
	}
}

func TestClient_DisassociateRoleFromUser(t *testing.T) {
	payload := Email{Email: "shipa.io"}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiRoleUser("role")+"/user@shipa.io", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodDelete),
	)
	defer teardown()

	if err := client.DisassociateRoleFromUser("role", "user@shipa.io"); err != nil {
		t.Error(err)
	}
}
