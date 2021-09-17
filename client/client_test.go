package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

type AuthRequest struct {
	Name      string                  `json:"password"`
	Endpoint  *ClusterEndpointCreate  `json:"endpoint"`
	Resources *ClusterResourcesCreate `json:"resources,omitempty"`
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func setupServer(handlers ...clientest.Handler) (client *Client, teardown func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	c, err := NewClient(
		WithHost(server.URL),
		WithClient(server.Client()),
		WithToken("faketoken"))
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

func TestClient_WithHostEmptyHost(t *testing.T) {
	_, err := NewClient(
		WithHost(""))
	if err == nil {
		t.Error("Passing empty host should return an error")
	}
}

func TestClient_WithTokenEmptyToken(t *testing.T) {
	_, err := NewClient(
		WithToken(""))
	if err == nil {
		t.Error("Passing empty token should return an error")
	}
}

func TestClient_WithTokenEmptyAdminEmail(t *testing.T) {
	_, err := NewClient(
		WithAdminAuth("", "abc123"))
	if err == nil {
		t.Error("Passing empty admin email should return an error")
	}
}

func TestClient_WithTokenEmptyAdminPassword(t *testing.T) {
	_, err := NewClient(
		WithAdminAuth("noone@nowhere.com", ""))
	if err == nil {
		t.Error("Passing empty admin password should return an error")
	}
}

func TestClient_WithAdminAuthAuthenticate(t *testing.T) {
	mapping, encodedData := map[string]string{"password": "abc123"}, new(bytes.Buffer)
	json.NewEncoder(encodedData).Encode(mapping)

	mockServer := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Host == "target.shipa.mock:8081" && req.URL.Path == "/users/admin@shipa.io/tokens" {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"token":"abcdef1234567890"}`)),
				Header:     make(http.Header),
			}
		}
		return nil
	})

	client, _ := NewClient(
		WithHost("https://target.shipa.mock:8081"),
		WithClient(mockServer),
		WithAdminAuth("admin@shipa.io", "abc123"))

	client.authenticate()
	if client.Token != "abcdef1234567890" {
		t.Error("Failed to get request token for admin login")
	}
}
