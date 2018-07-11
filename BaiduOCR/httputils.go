package BaiduOCR

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(theUrl string, data interface{}, result interface{}) error {
	content, _ := json.Marshal(data)
	resp, err := http.Post(theUrl, "application/json;charset=utf-8", strings.NewReader(string(content)))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, result)
	return nil
}

func PostByForm(theUrl string, params map[string]string, result interface{}) error {

	var values url.Values = make(map[string][]string)
	for key, val := range params {
		values.Set(key, val)
	}
	resp, err := http.PostForm(theUrl, values)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 打印结果
	//fmt.Println(string(body))
	json.Unmarshal(body, result)
	return nil
}
