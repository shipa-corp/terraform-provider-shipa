package client

//
//func TestGetCluster(t *testing.T) {
//	c := getClient()
//
//	cluster, err := c.GetCluster("cinema-services")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//
//	printJSON(cluster)
//}
//
//func TestCreateCluster(t *testing.T) {
//	c := getClient()
//
//	req := &CreateClusterRequest{
//		Name: "ln-cl1",
//		//Provisioner: "kubernetes",
//		Endpoint: &ClusterEndpointCreate{
//			Addresses: []string{
//				"https://5e043d8f-8f9e-4dde-9cce-0a226d8b8678.us-west-1.linodelke.net:443",
//			},
//			Certificate: cert,
//			Token: token,
//
//		},
//		Resources: &ClusterResourcesCreate{
//			Frameworks: []*Framework{
//				{Name: "test-tf-14"},
//			},
//			//IngressControllers: []*IngressControllersCreate{
//			//	{
//			//		IngressIP:     "",
//			//		ServiceType:   "LoadBalancer",
//			//		Type:          "traefik",
//			//		HTTPPort:      80,
//			//		HTTPSPort:     443,
//			//		ProtectedPort: 31567,
//			//		Debug:         false,
//			//	},
//			//},
//		},
//	}
//
//	err := c.CreateCluster(req)
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//}
//
//func TestDeleteCluster(t *testing.T) {
//	c := getClient()
//
//	err := c.DeleteCluster("ln-cl1")
//	if err != nil {
//		t.Error(err.Error())
//	}
//}
