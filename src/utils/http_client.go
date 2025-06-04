package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Post(url string, header map[string]string, value string) (string, error) {
	httpHeader := http.Header{}
	for k, v := range header {
		httpHeader.Set(k, v)
	}
	var buff *bytes.Buffer
	if value != "" {
		buff = bytes.NewBuffer([]byte(value))
	}
	req, err := http.NewRequest(http.MethodPost, url, buff)
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return "", err
	}
	req.Header = httpHeader

	client := &http.Client{}
	resp, apiErr := client.Do(req)
	if apiErr != nil {
		fmt.Println("发送请求错误:", apiErr)
		return "", apiErr
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		var err error
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Printf("创建 gzip 读取器失败: %v\n", err)
		}
		defer reader.(*gzip.Reader).Close()
	}

	// 读取响应体
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return "", err
	}

	// 打印响应状态码和响应体
	//fmt.Println("响应状态码:", resp.StatusCode)
	//fmt.Println("响应体:", string(body))
	return string(body), err
}

func Get(url string, header map[string]string) (string, error) {
	httpHeader := http.Header{}
	for k, v := range header {
		httpHeader.Set(k, v)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return "", err
	}
	req.Header = httpHeader

	client := &http.Client{}
	resp, apiErr := client.Do(req)
	if apiErr != nil {
		fmt.Println("发送请求错误:", apiErr)
		return "", apiErr
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		var err error
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Printf("创建 gzip 读取器失败: %v\n", err)
		}
		defer reader.(*gzip.Reader).Close()
	}

	// 读取响应体
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return "", err
	}

	// 打印响应状态码和响应体
	//fmt.Println("响应状态码:", resp.StatusCode)
	//fmt.Println("响应体:", string(body))
	return string(body), err
}
