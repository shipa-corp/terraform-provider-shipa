package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetTeam(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/teams/team", Team{Name: "team"}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetTeam("team"); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateTeam(t *testing.T) {
	payload := &Team{Name: "team"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/teams", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateTeam(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteTeam(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/teams/team", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeleteTeam("team"); err != nil {
		t.Error(err)
	}
}
