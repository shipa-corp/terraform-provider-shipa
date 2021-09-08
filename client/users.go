package client

import (
	"errors"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Client) GetUser(email string) (*User, error) {
	users, err := c.ListUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (c *Client) ListUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := c.get(&users, apiUsers)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) CreateUser(req *User) error {
	return c.post(req, apiUsers)
}

func (c *Client) DeleteUser(email string) error {
	params := []*QueryParam{{Key: "user", Val: email}}
	return c.deleteWithParams(params, apiUsers)
}
