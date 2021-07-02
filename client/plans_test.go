package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetPlan(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/plans", []Plan{{Name: "plan"}}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetPlan("plan"); err != nil {
		t.Error(err)
	}
}

func TestClient_ListPlans(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/plans", []Plan{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.ListPlans(); err != nil {
		t.Error(err)
	}
}

func TestClient_CreatePlan(t *testing.T) {
	payload := &Plan{Name: "plan"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/plans", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreatePlan(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeletePlan(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/plans/plan", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeletePlan("plan"); err != nil {
		t.Error(err)
	}
}
