package client

type Permission struct {
	Role        string   `json:"name"`
	Permissions []string `json:"permission"`
}

type getRolePermissions struct {
	Role        string   `json:"name"`
	Context     string   `json:"context"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"scheme_names,omitempty"`
}

func (c *Client) GetPermission(role string) (*Permission, error) {
	req := &getRolePermissions{}
	err := c.get(req, apiRoles, role)
	if err != nil {
		return nil, err
	}

	return &Permission{
		Role: req.Role,
		Permissions: req.Permissions,
	}, nil
}

func (c *Client) CreatePermission(req *Permission) error {
	return c.post(req, apiRolePermissions(req.Role))
}

func (c *Client) DeletePermission(role, permission string) error {
	return c.delete(apiRolePermissions(role), permission)
}
