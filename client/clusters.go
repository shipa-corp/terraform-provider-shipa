package client

type TerraformCreateClusterRequest struct {
	Name      string                           `json:"name"`
	Endpoint  *ClusterEndpointCreate           `json:"endpoint"`
	Resources *TerraformClusterResourcesCreate `json:"resources,omitempty"`
}

type TerraformClusterResourcesCreate struct {
	Frameworks         *TerraformFramework         `json:"frameworks,omitempty"`
	IngressControllers []*IngressControllersCreate `json:"ingressControllers,omitempty" terraform:"ingress_controllers"`
}

type TerraformFramework struct {
	Name []string `json:"name,omitempty"`
}

type CreateClusterRequest struct {
	Name      string                  `json:"name"`
	Endpoint  *ClusterEndpointCreate  `json:"endpoint"`
	Resources *ClusterResourcesCreate `json:"resources,omitempty"`
}

type ClusterEndpointCreate struct {
	Addresses         []string `json:"addresses,omitempty"`
	Certificate       string   `json:"caCert,omitempty" terraform:"ca_cert"`
	ClientCertificate string   `json:"clientCert,omitempty" terraform:"client_cert"`
	ClientKey         string   `json:"clientKey,omitempty" terraform:"client_key"`
	Token             string   `json:"token,omitempty"`
	Username          string   `json:"username,omitempty"`
	Password          string   `json:"password,omitempty"`
}

type ClusterResourcesCreate struct {
	Frameworks         []*Framework                `json:"frameworks,omitempty"`
	IngressControllers []*IngressControllersCreate `json:"ingressControllers,omitempty" terraform:"ingress_controllers"`
}

type IngressControllersCreate struct {
	IngressIP     string `json:"ingressIp,omitempty" terraform:"ingress_ip"`
	ServiceType   string `json:"serviceType,omitempty" terraform:"service_type"`
	Type          string `json:"type,omitempty"`
	HTTPPort      int64  `json:"httpPort,omitempty" terraform:"http_port"`
	HTTPSPort     int64  `json:"httpsPort,omitempty" terraform:"https_port"`
	ProtectedPort int64  `json:"protectedPort,omitempty" terraform:"protected_port"`
	Debug         bool   `json:"debug"`
	AcmeEmail     string `json:"acmeEmail,omitempty" terraform:"acme_email"`
	AcmeServer    string `json:"acmeServer,omitempty" terraform:"acme_server"`
}

type Framework struct {
	Name string `json:"name,omitempty"`
}

func NewCreateClusterRequest(from *TerraformCreateClusterRequest) *CreateClusterRequest {
	out := &CreateClusterRequest{Name: from.Name}
	out.Endpoint = from.Endpoint
	if from.Resources != nil {
		out.Resources = &ClusterResourcesCreate{
			Frameworks:         convertTerraformFrameworks(from.Resources.Frameworks),
			IngressControllers: from.Resources.IngressControllers,
		}
	}

	return out
}

func convertTerraformFrameworks(input *TerraformFramework) []*Framework {
	if input == nil {
		return nil
	}

	out := make([]*Framework, 0, len(input.Name))
	for _, name := range input.Name {
		out = append(out, &Framework{Name: name})
	}
	return out
}

func NewTerraformCreateClusterRequest(from *CreateClusterRequest) *TerraformCreateClusterRequest {
	out := &TerraformCreateClusterRequest{Name: from.Name}
	out.Endpoint = from.Endpoint
	if from.Resources != nil {
		out.Resources = &TerraformClusterResourcesCreate{
			Frameworks:         convertFrameworks(from.Resources.Frameworks),
			IngressControllers: from.Resources.IngressControllers,
		}
	}

	return out
}

func convertFrameworks(input []*Framework) *TerraformFramework {
	if input == nil {
		return nil
	}

	out := &TerraformFramework{
		Name: make([]string, 0, len(input)),
	}

	for _, framework := range input {
		out.Name = append(out.Name, framework.Name)
	}
	return out
}

func (c *Client) GetCluster(name string) (*CreateClusterRequest, error) {
	cluster := &CreateClusterRequest{}
	err := c.get(&cluster, apiClusters, name)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

//func (c *Client) ListClusters() ([]*Cluster, error) {
//	// TODO: need to fix json mapping
//	clusters := make([]*Cluster, 0)
//	err := c.get(&clusters, apiClusters)
//	if err != nil {
//		return nil, err
//	}
//
//	return clusters, nil
//}

func (c *Client) CreateCluster(req *CreateClusterRequest) error {
	return c.post(req, apiClusters)
}

func (c *Client) UpdateCluster(req *CreateClusterRequest) error {
	return c.put(req, apiClusters, req.Name, "config")
}

func (c *Client) DeleteCluster(name string) error {
	return c.delete(apiClusters, name)
}
