package main

func main() {
	// 发送POST请求
	//postData := []byte(`{"message": "helloWorld", "passWord": "value2", "name": "value3"}`)

	checkBaseLogin()
	checkAdminUpdateAdmin()

	//checkAdminCreateUser()
	//checkAdminUpdateUser()
	//checkAdminDeleteUser()
	//checkAdminUpdateUser()
	//checkAdminCreateUser()
	//checkAdminFindUser()
	//// 发送GET请求
	//getURL := "https://example.com/api/get-endpoint?key=value"
	//
	//resp, err = sendGETRequest(getURL)
	//if err != nil {
	//	fmt.Println("GET请求出错:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//responseBody, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("读取响应出错:", err)
	//	return
	//}
	//
	//fmt.Println("GET响应:", string(responseBody))
}
