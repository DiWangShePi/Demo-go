package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

// 修改admin信息三次，两次尝试登录。前三次都应该成功 ，后两次前失败后成功
func checkAdminUpdateUser() {
	formData := url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "123456789")
	formData.Set("stedentId", "12345678")
	formData.Set("grade", "4")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")
	postURL := "http://localhost:9090/admin/updateUser"
	resp, err := sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("AdminUpdateUser 1 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("AdminUpdateUser 1 读取响应出错:", err)
		return
	}
	fmt.Println("AdminUpdateUser 1:", string(responseBody))
}
