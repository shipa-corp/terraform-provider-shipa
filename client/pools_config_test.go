package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetPoolConfig(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/pools-config/pool", PoolConfig{Name: "pool"}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetPoolConfig("pool"); err != nil {
		t.Error(err)
	}
}

func TestClient_UpdatePoolConfig(t *testing.T) {
	payload := &PoolConfig{Name: "pool"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/pools-config", payload, http.MethodPut),
	)
	defer teardown()

	if err := client.UpdatePoolConfig(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_CreatePoolConfig(t *testing.T) {
	payload := &PoolConfig{Name: "pool"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/pools-config", payload, http.MethodPost),
	)
	defer teardown()

	if err := client.CreatePoolConfig(payload); err != nil {
		t.Error(err)
	}
}
