package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

// 修改admin信息三次，两次尝试登录。前三次都应该成功 ，后两次前失败后成功
func checkAdminUpdateAdmin() {
	formData := url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "123456789")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")
	postURL := "http://localhost:9090/admin/update"
	resp, err := sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Update 1 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login 1 读取响应出错:", err)
		return
	}
	fmt.Println("Login 1:", string(responseBody))

	formData = url.Values{}
	formData.Set("name", "GULUGUL")
	formData.Set("password", "12345678")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")
	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Update 2 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login 2 读取响应出错:", err)
		return
	}
	fmt.Println("Login 2:", string(responseBody))

	formData = url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "1234567")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")

	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Update 3 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login 3 读取响应出错:", err)
		return
	}
	fmt.Println("Login 3:", string(responseBody))

	formData = url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "12345678")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")

	postURL = "http://localhost:9090/padmin/login"
	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Login after update 1 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login after update 1 读取响应出错:", err)
		return
	}
	fmt.Println("Login after update 1:", string(responseBody))

	formData = url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "1234567")
	formData.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWxpZCI6InZhbGlkIiwidXNlcm5hbWUiOiIiLCJJZGVudGl0eSI6ImFkbWluIiwiVXNlcm5hbWUiOiJHVUxVR1VMVSIsIklEIjoiMSIsIkJ1ZmZlclRpbWUiOjM2MDAsIkNsYWltcyI6eyJpc3MiOiJibGFuayIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjk1NTM1OTY2LCJuYmYiOjE2OTU0NDk1Njd9fQ.WjA0FuV6aBR6bcZTPhhJnAhbjE3fNsrfYcEKV4EW5SI")

	postURL = "http://localhost:9090/padmin/login"
	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Login after update 2 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login after update 2 读取响应出错:", err)
		return
	}
	fmt.Println("Login after update 2:", string(responseBody))
}
