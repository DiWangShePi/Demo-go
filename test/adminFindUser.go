package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

// 修改admin信息三次，两次尝试登录。前三次都应该成功 ，后两次前失败后成功
func checkAdminFindUser() {
	formData := url.Values{}
	formData.Set("stedentId", "12345678")
	postURL := "http://localhost:9090/admin/findUser"
	resp, err := sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("AdminFindUser 1 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("AdminFindUser 1 读取响应出错:", err)
		return
	}
	fmt.Println("AdminFindUser 1:", string(responseBody))
}
