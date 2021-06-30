package client

type NetworkPolicy struct {
	Ingress    *NetworkPolicyConfig `json:"ingress,omitempty"`
	Egress     *NetworkPolicyConfig `json:"egress,omitempty"`
	RestartApp bool                 `json:"restart_app"`
}

// CreateOrUpdateNetworkPolicy - creates or updates network policy
func (c *Client) CreateOrUpdateNetworkPolicy(app string, config *NetworkPolicy) error {
	return c.put(config, apiAppNetworkPolicy(app))
}

// DeleteNetworkPolicy - deletes network policy from the given app
func (c *Client) DeleteNetworkPolicy(app string) error {
	return c.delete(apiAppNetworkPolicy(app))
}

// GetNetworkPolicy - get current policy of an app
func (c *Client) GetNetworkPolicy(app string) (*NetworkPolicy, error) {
	config := &NetworkPolicy{}
	err := c.get(config, apiAppNetworkPolicy(app))
	if err != nil {
		return nil, err
	}

	return config, nil
}
