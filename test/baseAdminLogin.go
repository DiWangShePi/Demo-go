package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

// 案例一：正确登录 案例二：名字错误 案例三：密码错误
func checkBaseLogin() {
	formData := url.Values{}
	formData.Set("name", "GULUGULU")
	formData.Set("password", "12345678")
	postURL := "http://localhost:9090/padmin/login"
	resp, err := sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Login 1 请求出错:", err)
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
	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Login 2 请求出错:", err)
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
	resp, err = sendPostRequest(postURL, formData)
	if err != nil {
		fmt.Println("Login 3 请求出错:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Login 3 读取响应出错:", err)
		return
	}
	fmt.Println("Login 3:", string(responseBody))
}
