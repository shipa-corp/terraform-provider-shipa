package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetNetworkPolicy(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/"+apiAppNetworkPolicy("app"), NetworkPolicy{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetNetworkPolicy("app"); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateOrUpdateNetworkPolicy(t *testing.T) {
	payload := &NetworkPolicy{}
	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/"+apiAppNetworkPolicy("app"), clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPut),
	)
	defer teardown()

	if err := client.CreateOrUpdateNetworkPolicy("app", payload); err != nil {
		t.Error(err)
	}
}
