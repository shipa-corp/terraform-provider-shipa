package client

import "errors"

type Pool struct {
	Name        string   `json:"name"`
	Default     bool     `json:"default"`
	Provisioner string   `json:"provisioner"`
	Public      bool     `json:"public"`
	Teams       []string `json:"teams,omitempty"`
	Allowed     *Allowed `json:"allowed,omitempty"`
}

type Allowed struct {
	Driver []string `json:"driver,omitempty"`
	Plan   []string `json:"plan,omitempty"`
	Team   []string `json:"team,omitempty"`
}

type CreatePoolRequest struct {
	Name        string `json:"name"`
	Default     bool   `json:"default"`
	Provisioner string `json:"provisioner"`
	Public      bool   `json:"public"`
	Force       bool   `json:"force"`
}

type UpdatePoolRequest struct {
	Name    string `json:"-"`
	Default bool   `json:"default"`
	Public  bool   `json:"public"`
	Force   bool   `json:"force"`
}

func (c *Client) GetPool(name string) (*Pool, error) {
	pools, err := c.ListPools()
	if err != nil {
		return nil, err
	}

	for _, pool := range pools {
		if pool.Name == name {
			return pool, nil
		}
	}

	return nil, errors.New("framework not found")
}

func (c *Client) ListPools() ([]*Pool, error) {
	pools := make([]*Pool, 0)
	err := c.get(&pools, apiPools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

func (c *Client) CreatePool(req *CreatePoolRequest) error {
	return c.post(req, apiPools)
}

func (c *Client) UpdatePool(req *UpdatePoolRequest) error {
	return c.put(req, apiPools, req.Name)
}

func (c *Client) DeletePool(name string) error {
	return c.delete(apiPools, name)
}
