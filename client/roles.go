package client

type Role struct {
	Name        string `json:"name"`
	Context     string `json:"context"`
	Description string `json:"description,omitempty"`
}

func (c *Client) GetRole(name string) (*Role, error) {
	role := &Role{}
	err := c.get(role, apiRoles, name)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *Client) CreateRole(req *Role) error {
	return c.post(req, apiRoles)
}

func (c *Client) DeleteRole(name string) error {
	return c.delete(apiRoles, name)
}

type Email struct {
	Email string `json:"email"`
}

func (c *Client) AssociateRoleToUser(role, email string) error {
	return c.post(&Email{email}, apiRoleUser(role))
}

func (c *Client) DisassociateRoleFromUser(role, email string) error {
	return c.deleteWithPayload(&Email{email}, nil, apiRoleUser(role), email)
}
