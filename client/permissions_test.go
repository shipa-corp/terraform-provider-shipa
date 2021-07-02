package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetPermission(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/roles/role", Permission{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetPermission("role"); err != nil {
		t.Error(err)
	}
}

func TestClient_CreatePermission(t *testing.T) {
	payload := &Permission{Role: "role"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/roles/role/permissions", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreatePermission(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeletePermission(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/roles/role/permissions/perm", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeletePermission("role", "perm"); err != nil {
		t.Error(err)
	}
}
