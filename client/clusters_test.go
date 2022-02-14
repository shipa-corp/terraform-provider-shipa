package client

import (
	"net/http"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client/clientest"
)

func TestClient_GetCluster(t *testing.T) {
	client, teardown := setupServer(
		clientest.PrintJsonHandler("/provisioner/clusters/cluster", CreateClusterRequest{}, http.MethodGet),
	)
	defer teardown()

	if _, err := client.GetCluster("cluster"); err != nil {
		t.Error(err)
	}
}

func TestClient_UpdateCluster(t *testing.T) {
	payload := &CreateClusterRequest{Name: "cluster", Endpoint: &ClusterEndpointCreate{Addresses: []string{"addr"}}}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/provisioner/clusters/cluster", payload, http.MethodPut),
	)
	defer teardown()

	if err := client.UpdateCluster(payload); err != nil {
		t.Error(err)
	}
}

func TestClient_CreateCluster(t *testing.T) {
	payload := &CreateClusterRequest{
		Name: "ln-cl1",
		//Provisioner: "kubernetes",
		Endpoint: &ClusterEndpointCreate{
			Addresses: []string{
				"https://5e043d8f-8f9e-4dde-9cce-0a226d8b8678.us-west-1.linodelke.net:443",
			},
			Certificate: "",
			Token:       "",
		},
		Resources: &ClusterResourcesCreate{
			Frameworks: []*Framework{
				{Name: "test-tf-14"},
			},
			IngressController: &IngressControllerCreate{
				IngressIP:     "",
				ServiceType:   "LoadBalancer",
				Type:          "traefik",
				HTTPPort:      80,
				HTTPSPort:     443,
				ProtectedPort: 31567,
				Debug:         false,
			},
		},
	}

	client, teardown := setupServer(
		clientest.CheckPayloadHandler("/provisioner/clusters", payload, http.MethodPost),
	)
	defer teardown()

	if err := client.CreateCluster(payload); err != nil {
		t.Error(err)
	}
}
