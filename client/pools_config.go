package client

type PoolConfig struct {
	Name      string         `json:"shipaFramework" terraform:"name"`
	Resources *PoolResources `json:"resources,omitempty"`
}

type PoolResources struct {
	General *PoolGeneral `json:"general,omitempty"`
	Node    *PoolNode    `json:"shipaNode,omitempty"`
}

type PoolNode struct {
	Drivers   []string       `json:"drivers,omitempty"`
	AutoScale *PoolAutoScale `json:"autoScale,omitempty"`
}

type PoolAutoScale struct {
	MaxContainer int     `json:"maxContainer"`
	MaxMemory    int     `json:"maxMemory"`
	ScaleDown    float64 `json:"scaleDown"`
	Rebalance    bool    `json:"rebalance"`
}

type PoolGeneral struct {
	Setup           *PoolSetup           `json:"setup,omitempty"`
	Plan            *PoolPlan            `json:"plan,omitempty"`
	Security        *PoolSecurity        `json:"security,omitempty"`
	Access          *PoolServiceAccess   `json:"access,omitempty"`
	Services        *PoolServiceAccess   `json:"services,omitempty"`
	Router          string               `json:"router,omitempty"`
	Volumes         []string             `json:"volumes,omitempty"`
	AppQuota        *PoolAppQuota        `json:"appQuota,omitempty"`
	ContainerPolicy *PoolContainerPolicy `json:"containerPolicy,omitempty"`
	NetworkPolicy   *PoolNetworkPolicy   `json:"networkPolicy,omitempty"`
}

type PoolContainerPolicy struct {
	AllowedHosts []string `json:"allowedHosts,omitempty"`
}

type PoolAppQuota struct {
	Limit string `json:"limit,omitempty"`
}

type PoolServiceAccess struct {
	Append    []string `json:"append,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
}

type PoolSecurity struct {
	DisableScan        bool     `json:"disableScan"`
	ScanPlatformLayers bool     `json:"scanPlatformLayers"`
	IgnoreComponents   []string `json:"ignoreComponents,omitempty"`
	IgnoreCVES         []string `json:"ignoreCves,omitempty"`
}

type PoolPlan struct {
	Name string `json:"name,omitempty"`
}

type PoolSetup struct {
	Default             bool   `json:"default"`
	Public              bool   `json:"public"`
	Provisioner         string `json:"provisioner,omitempty"`
	KubernetesNamespace string `json:"kubernetesNamespace,omitempty"`
}

type PoolNetworkPolicy struct {
	Ingress            *NetworkPolicy `json:"ingress,omitempty"`
	Egress             *NetworkPolicy `json:"egress,omitempty"`
	DisableAppPolicies bool           `json:"disableAppPolicies"`
}

type NetworkPolicy struct {
	PolicyMode        string               `json:"policy_mode,omitempty"`
	CustomRules       []*NetworkPolicyRule `json:"custom_rules,omitempty"`
	ShipaRules        []*NetworkPolicyRule `json:"shipa_rules,omitempty"`
	ShipaRulesEnabled []string             `json:"shipa_rules_enabled,omitempty"`
}

type NetworkPolicyRule struct {
	ID          string         `json:"id,omitempty"`
	Enabled     bool           `json:"enabled"`
	Description string         `json:"description,omitempty"`
	Ports       []*NetworkPort `json:"ports,omitempty"`
	Peers       []*NetworkPeer `json:"peers,omitempty"`
	AllowedApps []string       `json:"allowed_apps,omitempty"`
}

type NetworkPort struct {
	Protocol string              `json:"protocol,omitempty"`
	Port     *NetworkPortDetails `json:"port,omitempty"`
}

type NetworkPortDetails struct {
	Type   int    `json:"type"`
	Intval int    `json:"intval"`
	Strval string `json:"strval,omitempty"`
}

type NetworkPeer struct {
	PodSelector       *NetworkPeerSelector `json:"podselector,omitempty"`
	NamespaceSelector *NetworkPeerSelector `json:"namespaceselector,omitempty"`
	IPBlock           []string             `json:"ipblock,omitempty"`
}

type NetworkPeerSelector struct {
	MatchLabels      map[string]string     `json:"matchlabels,omitempty"`
	MatchExpressions []*SelectorExpression `json:"matchexpressions,omitempty"`
}

type SelectorExpression struct {
	Key      string   `json:"key,omitempty"`
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}

func (c *Client) GetPoolConfig(name string) (*PoolConfig, error) {
	poolConfig := &PoolConfig{}
	err := c.get(poolConfig, apiPoolsConfig, name)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

func (c *Client) CreatPoolConfig(pool *PoolConfig) error {
	return c.post(pool, apiPoolsConfig)
}

func (c *Client) UpdatePoolConfig(req *PoolConfig) error {
	return c.put(req, apiPoolsConfig)
}
