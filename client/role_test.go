package client

//
//
//func TestCreateRole(t *testing.T) {
//	c := getClient()
//
//	err := c.CreateRole(&Role{
//		Name: "RoleDaniel",
//		Context: "app",
//	})
//	if err != nil {
//		t.Error(err.Error())
//	}
//}
//
//func TestGetRole(t *testing.T) {
//	c := getClient()
//
//	role, err := c.GetRole("RoleDaniel")
//	if err != nil {
//		t.Error(err.Error())
//	}
//
//	data, _ := json.Marshal(role)
//	log.Println("role:", string(data))
//}
//
//func TestDeleteRole(t *testing.T) {
//	c := getClient()
//
//	err := c.DeleteRole("RoleDaniel")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//}
//
