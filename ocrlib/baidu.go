package ocrlib

import (
	"github.com/DiamondYuan/ocr/utils"
	"encoding/json"
	"fmt"
	"errors"
	"math"
	"os/user"
	"io/ioutil"
)

const (
	BAIDU_TOKEN_URL  = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s"
	BAIDU_generalUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
)

func (baiduOcr *Baidu) GetResult(path string) (string, error) {
	token, err := loadAccessToken()
	if err != nil {
		return "", err
	}
	bytes, err := LoadImageToByte(path);
	if err != nil {
		return "", errors.New("读取图片失败")
	}

	p := make(map[string]string)
	p["image"] = bytes
	p["access_token"] = token

	result, err := utils.PostByForm(BAIDU_generalUrl, p)
	if err != nil {
		return "", errors.New("请求百度发生错误")
	}
	ocrResult := ImageResult{}
	json.Unmarshal([]byte(result), &ocrResult)
	finalResult := ""
	for k, v := range ocrResult.Words {
		if k != 0 {
			finalResult += " "
			pointOne := ocrResult.Words[k-1].Location
			pointTwo := v.Location
			angel := math.Atan2(float64(pointTwo.Top-pointOne.Top), float64(pointTwo.Left-pointOne.Left))
			if math.Abs(180/math.Pi*angel) > 10 {
				finalResult += "\n"
			}

		}
		finalResult += v.Text
	}
	return finalResult, nil
}

func loadAccessToken() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("Get usr error.")
	}
	baiduConfigFile, loadError := ioutil.ReadFile(usr.HomeDir + "/.ocr/baidu_config")
	if loadError != nil {
		return "", errors.New("Read config file error")
	}
	config := BaiduOcr{}
	json.Unmarshal([]byte(baiduConfigFile), &config)
	url := fmt.Sprintf(BAIDU_TOKEN_URL, config.AppKey, config.AppSecret)
	token := BaiduAccessToken{}
	result, err := utils.Get(url)
	if err != nil {
		return "", errors.New("AppKey or AppSecret error")
	}
	json.Unmarshal([]byte(result), &token)
	return token.AccessToken, nil
}
