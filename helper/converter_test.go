package helper

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func TestNestedStruct(t *testing.T) {
	source := []interface{}{
		map[string]interface{}{
			"general": []interface{}{
				map[string]interface{}{
					"setup": []interface{}{
						map[string]interface{}{
							"public":  true,
							"default": false,
						},
					},
					"plan": []interface{}{
						map[string]interface{}{
							"name": "dev",
						},
					},
					"security": []interface{}{
						map[string]interface{}{
							"disable_scan":         true,
							"scan_platform_layers": false,
						},
					},
					"access": []interface{}{
						map[string]interface{}{
							"append": []interface{}{"dev", "test"},
						},
					},
					"router": "traefik",
					"app_quota": []interface{}{
						map[string]interface{}{
							"limit": "1",
						},
					},
				},
			},
		},
	}

	target := &client.PoolConfig{
		Name:      "test",
		Resources: &client.PoolResources{},
	}
	TerraformToStruct(source, target.Resources)
	data, err := json.Marshal(target)
	if err != nil {
		t.Error(err)
	}
	actual := string(data)
	log.Println(actual)
	expected := `{"shipaFramework":"test","resources":{"general":{"setup":{"default":false,"public":true},"plan":{"name":"dev"},"security":{"disableScan":true,"scanPlatformLayers":false},"access":{"append":["dev","test"]},"router":"traefik","appQuota":{"limit":"1"}}}}`
	if expected != actual {
		t.Error("json matching failed")
	}
}

func TestEmptyStruct(t *testing.T) {
	source := []interface{}{
		map[string]interface{}{
			"general": []interface{}{
				map[string]interface{}{
					"setup": []interface{}{},
					"plan": []interface{}{
						map[string]interface{}{
							"name": "dev",
						},
					},
				},
			},
		},
	}

	target := &client.PoolResources{}
	TerraformToStruct(source, target)
	data, err := json.Marshal(target)
	if err != nil {
		t.Error(err)
	}
	actual := string(data)
	log.Println(actual)
	expected := `{"general":{"plan":{"name":"dev"}}}`
	if expected != actual {
		t.Error("json matching failed")
	}
}

func TestConvertStructToRaw(t *testing.T) {
	source := &client.PoolConfig{
		Name: "test",
	}
	res := StructToTerraform(source)
	log.Println("res", res)
}

func TestConvertStructWithArray(t *testing.T) {
	source := &client.PoolContainerPolicy{
		AllowedHosts: []string{"one", "two"},
	}
	res := StructToTerraform(source)
	log.Println("res", res)
	log.Println("res", res[0].(map[string]interface{})["allowed_hosts"])
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertStructWithNestedStruct(t *testing.T) {
	source := &client.PoolNode{
		Drivers: []string{"test"},
		AutoScale: &client.PoolAutoScale{
			MaxContainer: 1,
		},
	}
	res := StructToTerraform(source)
	log.Println("res", res)
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertStructWithListOfNestedStructs(t *testing.T) {
	source := &client.NetworkPolicyRule{
		Ports: []*client.NetworkPort{
			{
				Protocol: "http",
			},
			{
				Protocol: "udp",
			},
		},
	}
	res := StructToTerraform(source)
	log.Println("res", res)
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertStructWithComplexStruct(t *testing.T) {
	expected := `{"shipa_pool":"test","resources":{"general":{"setup":{"default":false,"public":true},"plan":{"name":"dev"},"security":{"disable_scan":true,"scan_platform_layers":false},"access":{"append":["dev","test"]},"router":"traefik","app_quota":{"limit":"1"}}}}`
	source := &client.PoolConfig{}
	json.Unmarshal([]byte(expected), source)

	res := StructToTerraform(source)
	log.Println("res", res)
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertListOfStructs(t *testing.T) {
	source := &[]*client.NetworkPort{
		{
			Protocol: "http",
		},
		{
			Protocol: "udp",
		},
	}

	res := StructToTerraform(source)
	log.Println("res", res)
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertStructWithMap(t *testing.T) {
	source := &client.NetworkPeerSelector{
		MatchLabels: map[string]string{
			"label_1": "test",
			"label_2": "test",
		},
	}

	res := StructToTerraform(source)
	log.Println("res", res)
	data, _ := json.Marshal(res)
	log.Println("json:", string(data))
}

func TestConvertTerraformWithMap(t *testing.T) {
	expected := `{"matchLabels":{"label_1":"test","label_2":"test"}}`
	source := []interface{}{
		map[string]interface{}{
			"match_labels": map[string]interface{}{
				"label_1": "test",
				"label_2": "test",
			},
		},
	}

	target := &client.NetworkPeerSelector{}
	TerraformToStruct(source, target)
	data, err := json.Marshal(target)
	if err != nil {
		t.Error(err)
	}
	actual := string(data)
	log.Println(actual)
	if expected != actual {
		t.Error("json matching failed")
	}
}

func TestConvertMaps(t *testing.T) {
	input := `{"matchLabels":{"label_1":"test","label_2":"test"}}`

	object := &client.NetworkPeerSelector{}
	json.Unmarshal([]byte(input), object)

	rawData := StructToTerraform(object)
	log.Println("RAW data:", rawData)
	data, _ := json.Marshal(rawData)
	log.Println("JSON from RAW data:", string(data))

	object = &client.NetworkPeerSelector{}
	TerraformToStruct(rawData, object)
	data, _ = json.Marshal(object)
	actual := string(data)
	log.Println("JSON from object:", actual)
	if input != actual {
		t.Error("json matching failed")
	}
}
