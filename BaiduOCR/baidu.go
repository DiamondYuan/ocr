package BaiduOCR

import (
	"encoding/base64"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"bytes"
)

const (
	BAIDU_TOKEN_URL  = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s"
	BAIDU_generalUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	MAX_SIZE         = 4 * 1024 * 1024
)

type BaiduErrorMessage struct {
	Code string `json:"error"`
	Msg  string `json:"error_description"`
}
type BaiduAccessToken struct {
	BaiduErrorMessage
	Expires       int64  `json:"expires_in"` //过期时间，单位秒。一般为1个月)
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
	AccessToken   string `json:"access_token"` //要获取的Access Token；
	RefreshToken  string `json:"refresh_token"`
}
type Parameter map[string]string

func BuildParameter() Parameter {
	return make(map[string]string)
}

type BaiduSdk struct {
	AppId     string
	AppKey    string
	AppSecret string
	Token     *BaiduAccessToken
}

func (this *BaiduSdk) getAccessToken() {
	theUrl := fmt.Sprintf(BAIDU_TOKEN_URL, this.AppKey, this.AppSecret)
	token := BaiduAccessToken{}
	Post(theUrl, nil, &token)
	this.Token = &token
}

func (this *BaiduSdk) Call(theUrl string, p baiduQuery, result interface{}) {
	if this.Token == nil {
		this.getAccessToken()
	}
	if p != nil {
		parames := p.toParameter()
		parames["access_token"] = this.Token.AccessToken
		err := PostByForm(theUrl, parames, result)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func bytesTOBaiduBase64(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

type baiduQuery interface {
	toParameter() Parameter
}

func (this *BaseImageQuery) toParameter() Parameter {
	p := BuildParameter()
	p["image"] = bytesTOBaiduBase64(this.Image)
	return p
}

type RecognizeItem struct {
	Location Image_location `json:"location"`

	Text string `json:"words"`
}
type Image_location struct {
	Left   int64 `json:"left"`
	Top    int64 `json:"top"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

func (this *BaiduOcr) Init(sdk *BaiduSdk) {
	this.sdk = sdk
}

type BaseImageQuery struct {
	Image []byte
}

func (this *BaiduOcr) RecognizedImage(path string) (string, error) {
	bytes1, err := this.LoadImageToByte(path);
	if err != nil {
		return "", err
	}
	image2 := BaseImageQuery{
		Image: bytes1,
	}
	result := ImageResult{}
	this.sdk.Call(BAIDU_generalUrl, &image2, &result)
	var buf bytes.Buffer
	for _, value := range result.Words {
		buf.WriteString(value.Text)
	}
	return buf.String(), nil

}
