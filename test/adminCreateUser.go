package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

// 修改admin信息三次，两次尝试登录。前三次都应该成功 ，后两次前失败后成功
func checkAdminCreateUser() {
	formData := url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "123456789")
	formData.Set("studentId", "12345678")
	formData.Set("grade", "4")
	postURL := "http://localhost:9090/admin/createUser"
	resp, err := sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("AdminCreateUser 1 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("AdminCreateUser 1 读取响应出错:", err)
		return
	}
	fmt.Println("AdminCreateUser 1:", string(responseBody))
}
