package client

import "errors"

func (c *Client) GetPlan(name string) (*Plan, error) {
	plans, err := c.ListPlans()
	if err != nil {
		return nil, err
	}
	for _, plan := range plans {
		if plan.Name == name {
			return plan, nil
		}
	}

	return nil, errors.New("plan not found")
}

func (c *Client) ListPlans() ([]*Plan, error) {
	plans := make([]*Plan, 0)
	err := c.get(&plans, apiPlans)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (c *Client) CreatePlan(req *Plan) error {
	return c.post(req, apiPlans)
}

func (c *Client) DeletePlan(name string) error {
	return c.delete(apiPlans, name)
}
