package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func getClient() *Client {
	c, err := NewClient(host, authToken)
	if err != nil {
		panic(err)
	}
	return c
}

func TestGetPoolConfig(t *testing.T) {
	c := getClient()

	config, err := c.GetPoolConfig("cinema-services")
	if err != nil {
		t.Error(err.Error())
		return
	}

	printJSON(config)
}

func TestDeletePool(t *testing.T) {
	c := getClient()

	err := c.DeletePool("test-tf-14")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func printJSON(payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("ERR:", err.Error())
		return
	}
	body := &bytes.Buffer{}
	body.Write(data)

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, data, "", "  ")
	fmt.Println(string(prettyJSON.Bytes()))
}
