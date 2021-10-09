package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容

func Get(url string, headerMap map[string]string, queryParam map[string]string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	if len(queryParam) > 0 {
		url += "?"
		i := 0
		for key, value := range queryParam {
			if i == 0 {
				url += key + "=" + value
			} else {
				url += "&" + key + "=" + value
			}
			i++
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println(err)
	}
	for key, value := range headerMap {
		request.Header.Add(key, value)
	}
	response, _ := client.Do(request)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(response.Body)
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err == io.EOF {
			break
		} else if err != nil {
			println(err)
		}
	}
	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容

func Post(url string, data interface{}, contentType string) (string, error) {

	client := &http.Client{Timeout: 20 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		println("json marshal failed")
		return "", err
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)

	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}
