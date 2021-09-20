package client

type Team struct {
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`
}

type UpdateTeamRequest struct {
	Name string   `json:"newname,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func (c *Client) GetTeam(name string) (*Team, error) {
	team := &Team{}
	err := c.get(team, apiTeams, name)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (c *Client) CreateTeam(req *Team) error {
	return c.post(req, apiTeams)
}

func (c *Client) UpdateTeam(name string, req *UpdateTeamRequest) error {
	return c.put(req, apiTeams, name)
}

func (c *Client) DeleteTeam(name string) error {
	return c.delete(apiTeams, name)
}
