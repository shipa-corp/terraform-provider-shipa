package client

import (
	"net/http"
	"net/http/httptest"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func setupServer(handlers ...clientest.Handler) (client *Client, teardown func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	c, err := NewClient(
		WithHost(server.URL),
		WithClient(server.Client()))
	if err != nil {
		panic(err)
	}

	for _, h := range handlers {
		mux.HandleFunc(h.GetEndpoint(), h.GetHandlerFunc())
	}

	return c, func() {
		server.Close()
	}
}
