package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetPool(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/pools", []Pool{{Name: "pool"}}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetPool("pool"); err != nil {
		t.Error(err)
	}
}

func TestClient_ListPools(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/pools", []Pool{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.ListPools(); err != nil {
		t.Error(err)
	}
}

func TestClient_CreatePool(t *testing.T) {
	payload := &CreatePoolRequest{Name: "pool"}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/pools", clientest.ComparablePayload{
			Want: payload,
			Got:  payload,
		}, http.MethodPost),
	)
	defer teardown()

	if err := client.CreatePool(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_DeletePool(t *testing.T) {
	client, teardown := setupServer(
		clientest.CheckMethodHandler("/pools/pool", http.MethodDelete),
	)
	defer teardown()

	if err := client.DeletePool("pool"); err != nil {
		t.Error(err)
	}
}
