package client

import (
	"encoding/json"
	"log"
	"testing"
)


func TestCreatePermission(t *testing.T) {
	c := getClient()

	err := c.CreatePermission(&Permission{
		Role: "RoleDaniel",
		Permissions: []string{"app.read", "app.deploy"},
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetPermission(t *testing.T) {
	c := getClient()

	role, err := c.GetPermission("RoleDaniel")
	if err != nil {
		t.Error(err.Error())
	}

	data, _ := json.Marshal(role)
	log.Println("permission:", string(data))
}

func TestDeletePermission(t *testing.T) {
	c := getClient()

	err := c.DeletePermission("RoleDaniel", "app.deploy")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

