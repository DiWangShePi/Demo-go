package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func sendPostRequest(url string, data url.Values) (*http.Response, error) {
	// 构造POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendGETRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
