package client

//
//func TestGetPoolConfig(t *testing.T) {
//	c, teardown := setupServer(
//		clientest.PrintJsonHandler("pools", ),
//	)
//	defer teardown()
//
//
//	config, err := c.GetPoolConfig("cinema-services")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//
//	printJSON(config)
//}
//
//func TestDeletePool(t *testing.T) {
//	c, teardown := setupServer(
//		New,
//	)
//	defer teardown()
//
//	err := c.DeletePool("test-tf-14")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//}
