package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetUser(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/users", []User{{Email: "user@shipa.io"}}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetUser("user@shipa.io"); err != nil {
		t.Error(err)
	}
}

func TestClient_ListUsers(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/users", []User{{Email: "user@shipa.io"}}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.ListUsers(); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateUser(t *testing.T) {
	payload := &User{Email: "user@shipa.io"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/users", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateUser(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteUser(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/users/user", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteUser("user"); err != nil {
		t.Error(err)
	}
}
